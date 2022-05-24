package logger

import (
	"github.com/fujiahui/talnet-challenge-payman/config"
	"go.uber.org/zap/zapcore"
	"os"
)

// How to log, by example:
// logger.Info("Importing new file, zap.String("source", filename), zap.Int("size", 1024))
// To log a stacktrace:
// logger.Error("It went wrong, zap.Stack())

// DefaultZapLogger is the default logger instance that should be used to log
// It's assigned a default value here for tests (which do not call log.Configure())
//var DefaultZapLogger = newZapLogger(false, os.Stdout, nil)
var (
	defaultZapLogger   = newZapLogger(false, os.Stdout, nil)
	defaultSugarLogger = defaultZapLogger.Sugar()
)

func init() {
	cfg := config.LogConfig{
		EncodeLogsAsJson: false,
		Directory:        "../log/",
		Filename:         "info.log",
		MaxSize:          1024,
		MaxBackups:       0,
		MaxAge:           15,
		Level:            "warn",
	}

	NewDefaultLogger(&cfg)
}

// Debug uses fmt.Sprint to construct and log a message.
func Debug(args ...interface{}) {
	defaultSugarLogger.Debug(args...)
}

// Info uses fmt.Sprint to construct and log a message.
func Info(args ...interface{}) {
	defaultSugarLogger.Info(args...)
}

// Warn uses fmt.Sprint to construct and log a message.
func Warn(args ...interface{}) {
	defaultSugarLogger.Warn(args...)
}

// Error uses fmt.Sprint to construct and log a message.
func Error(args ...interface{}) {
	defaultSugarLogger.Error(args...)
}

// DPanic uses fmt.Sprint to construct and log a message. In development, the
// logger then panics. (See DPanicLevel for details.)
func DPanic(args ...interface{}) {
	defaultSugarLogger.DPanic(args...)
}

// Panic uses fmt.Sprint to construct and log a message, then panics.
func Panic(args ...interface{}) {
	defaultSugarLogger.Panic(args...)
}

// Fatal uses fmt.Sprint to construct and log a message, then calls os.Exit.
func Fatal(args ...interface{}) {
	defaultSugarLogger.Fatal(args...)
}

// Debugf uses fmt.Sprintf to log a templated message.
func Debugf(template string, args ...interface{}) {
	defaultSugarLogger.Debugf(template, args...)
}

// Infof uses fmt.Sprintf to log a templated message.
//var Infof = defaultSugarLogger.Infof
func Infof(template string, args ...interface{}) {
	defaultSugarLogger.Infof(template, args...)
}

// Warnf uses fmt.Sprintf to log a templated message.
func Warnf(template string, args ...interface{}) {
	defaultSugarLogger.Warnf(template, args...)
}

// Errorf uses fmt.Sprintf to log a templated message.
func Errorf(template string, args ...interface{}) {
	defaultSugarLogger.Errorf(template, args...)
}

// DPanicf uses fmt.Sprintf to log a templated message. In development, the
// logger then panics. (See DPanicLevel for details.)
func DPanicf(template string, args ...interface{}) {
	defaultSugarLogger.DPanicf(template, args...)
}

// Panicf uses fmt.Sprintf to log a templated message, then panics.
func Panicf(template string, args ...interface{}) {
	defaultSugarLogger.DPanicf(template, args...)
}

// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit.
func Fatalf(template string, args ...interface{}) {
	defaultSugarLogger.DPanicf(template, args...)
}

// Debugw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
//
// When debug-level logging is disabled, this is much faster than
//  s.With(keysAndValues).Debug(msg)
func Debugw(msg string, keysAndValues ...interface{}) {
	defaultSugarLogger.Debugw(msg, keysAndValues...)
}

// Infow logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func Infow(msg string, keysAndValues ...interface{}) {
	defaultSugarLogger.Infow(msg, keysAndValues...)
}

// Warnw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func Warnw(msg string, keysAndValues ...interface{}) {
	defaultSugarLogger.Warnw(msg, keysAndValues...)
}

// Errorw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func Errorw(msg string, keysAndValues ...interface{}) {
	defaultSugarLogger.Errorw(msg, keysAndValues...)
}

// DPanicw logs a message with some additional context. In development, the
// logger then panics. (See DPanicLevel for details.) The variadic key-value
// pairs are treated as they are in With.
func DPanicw(msg string, keysAndValues ...interface{}) {
	defaultSugarLogger.DPanicw(msg, keysAndValues...)
}

// Panicw logs a message with some additional context, then panics. The
// variadic key-value pairs are treated as they are in With.
func Panicw(msg string, keysAndValues ...interface{}) {
	defaultSugarLogger.Panicw(msg, keysAndValues...)
}

// Fatalw logs a message with some additional context, then calls os.Exit. The
// variadic key-value pairs are treated as they are in With.
func Fatalw(msg string, keysAndValues ...interface{}) {
	defaultSugarLogger.Fatalw(msg, keysAndValues...)
}

// Configure sets up the logging framework
//
// In production, the container logs will be collected and file logging should be disabled. However,
// during development it's nicer to see logs as text and optionally write to a file when debugging
// problems in the containerized pipeline
//
// The output log file will be located at /var/log/auth-service/auth-service.log and
// will be rolled when it reaches 20MB with a maximum of 1 backup.

func NewDefaultLogger(cfg *config.LogConfig) {
	writers := []zapcore.WriteSyncer{
		// newRollingFile(config),
		newRotateFile(cfg),
	}

	defaultZapLogger = newZapLogger(cfg.EncodeLogsAsJson, zapcore.NewMultiWriteSyncer(writers...), cfg)
	defaultSugarLogger = defaultZapLogger.Sugar()
}
