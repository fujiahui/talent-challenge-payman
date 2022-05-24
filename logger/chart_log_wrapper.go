package logger

import (
	"github.com/fujiahui/talnet-challenge-payman/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var (
	ChartLogger = &chartStdLogger{
		chartSugaredLogger: newZapLogger(false, os.Stdout, nil).Sugar(),
	}
)

func init() {
	cfg := config.LogConfig{
		EncodeLogsAsJson: false,
		Directory:        "../log/",
		Filename:         "chart.log",
		MaxSize:          1024,
		MaxBackups:       0,
		MaxAge:           15,
		Level:            "debug",
	}

	NewChartLogger(&cfg)
}

type chartStdLogger struct {
	chartSugaredLogger *zap.SugaredLogger
}

func (ssLogger *chartStdLogger) Print(v ...interface{}) {
	// Info("chart log", zap.String("message", fmt.Sprint(v...)))
	// ChartZapLogger.Info("[chart]", zap.String("message", fmt.Sprint(v...)))
	ssLogger.chartSugaredLogger.Info(v...)
}

func (ssLogger *chartStdLogger) Printf(format string, v ...interface{}) {
	//Info("chart log", zap.String("message", fmt.Sprintf(format, v...)))
	// ChartZapLogger.Info("[chart]", zap.String("message", fmt.Sprintf(format, v...)))
	ssLogger.chartSugaredLogger.Infof(format, v...)
}

func (ssLogger *chartStdLogger) Println(v ...interface{}) {
	// Info("chart log", zap.String("message", fmt.Sprint(v...)))
	// ChartZapLogger.Info("[chart]", zap.String("message", fmt.Sprint(v...)))
	ssLogger.chartSugaredLogger.Info(v...)
}

func NewChartLogger(cfg *config.LogConfig) {
	writers := []zapcore.WriteSyncer{
		// newRollingFile(config),
		newRotateFile(cfg),
	}

	chartZapLogger := newZapLogger(cfg.EncodeLogsAsJson, zapcore.NewMultiWriteSyncer(writers...), cfg)
	ChartLogger.chartSugaredLogger = chartZapLogger.Sugar()
}
