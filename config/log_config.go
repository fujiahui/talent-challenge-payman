package config

import "go.uber.org/zap/zapcore"

type LogConfig struct {
	// EncodeLogsAsJson makes the log framework log JSON
	EncodeLogsAsJson bool `yaml:"EncodeLogsAsJson"`
	// Directory to log to to when filelogging is enabled
	Directory string `yaml:"Directory"`
	// Filename is the name of the logfile which will be placed inside the directory
	Filename string `yaml:"Filename"`
	// MaxSize the max size in MB of the logfile before it's rolled
	MaxSize int `yaml:"MaxSize"`
	// MaxBackups the max number of rolled files to keep
	MaxBackups int `yaml:"MaxBackups"`
	// MaxAge the max age in days to keep a logfile
	MaxAge   int    `yaml:"MaxAge"`
	Level    string `yaml:"Level"`
	ZapLevel zapcore.Level
}
