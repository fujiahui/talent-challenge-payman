package logger

import (
	"github.com/fujiahui/talnet-challenge-payman/config"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path"
	"time"
)

func newRotateFile(config *config.LogConfig) zapcore.WriteSyncer {
	filename := path.Join(config.Directory, config.Filename)
	rotateLogger, err := rotatelogs.New(
		filename+".%Y%m%d",
		rotatelogs.WithLinkName(filename),
		rotatelogs.WithMaxAge(time.Hour*24*14),
		rotatelogs.WithRotationSize(int64(config.MaxSize*1024*1024)),
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	if err != nil {
		return nil
	}

	return zapcore.AddSync(rotateLogger)
}

func newRollingFile(cfg *config.LogConfig) zapcore.WriteSyncer {
	if err := os.MkdirAll(cfg.Directory, os.ModePerm); err != nil {
		Error("failed create log directory", zap.Error(err), zap.String("path", cfg.Directory))
		return nil
	}

	jackLogger := &lumberjack.Logger{
		Filename:   path.Join(cfg.Directory, cfg.Filename),
		MaxSize:    cfg.MaxSize,    //megabytes
		MaxAge:     cfg.MaxAge,     //days
		MaxBackups: cfg.MaxBackups, //files
		LocalTime:  true,
	}

	jackLogger.Rotate()
	return zapcore.AddSync(jackLogger)
}

func newZapLogger(encodeAsJSON bool, output zapcore.WriteSyncer, cfg *config.LogConfig) *zap.Logger {
	encCfg := zapcore.EncoderConfig{
		TimeKey:       "@timestamp",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "caller",
		FunctionKey:   zapcore.OmitKey,
		MessageKey:    "msg",
		StacktraceKey: "stacktrace",
		//LineEnding: 	zapcore.DefaultLineEnding,
		LineEnding:       "",
		EncodeLevel:      zapcore.CapitalLevelEncoder,
		EncodeTime:       zapcore.ISO8601TimeEncoder,
		EncodeDuration:   zapcore.NanosDurationEncoder,
		EncodeCaller:     zapcore.ShortCallerEncoder,
		ConsoleSeparator: " | ",
	}

	encoder := zapcore.NewConsoleEncoder(encCfg)
	if encodeAsJSON {
		encoder = zapcore.NewJSONEncoder(encCfg)
	}

	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	if cfg != nil {
		atomicLevel.SetLevel(cfg.ZapLevel)
	} else {
		atomicLevel.SetLevel(zapcore.DebugLevel)
	}

	return zap.New(zapcore.NewCore(encoder, output, atomicLevel), zap.AddCallerSkip(1), zap.AddCaller())
}
