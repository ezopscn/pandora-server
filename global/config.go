package global

import "go.uber.org/zap/zapcore"

// 配置引用
var Config Configuration

// 配置结构体
type Configuration struct {
	System SystemConfiguration `mapstructure:"system" json:"system"`
	Log    LogConfiguration    `mapstructure:"log" json:"log"`
}

// 系统配置
type SystemConfiguration struct {
	Listen string `mapstructure:"listen" json:"listen"`
	Port   string `mapstructure:"port" json:"port"`
}

// 日志配置
type LogConfiguration struct {
	System LoggerConfiguration `mapstructure:"system" json:"system"`
	Access LoggerConfiguration `mapstructure:"access" json:"access"`
}

// 日志类型配置
type LoggerConfiguration struct {
	Enabled    bool          `mapstructure:"enabled" json:"enabled"`
	Level      zapcore.Level `mapstructure:"level" json:"level"`
	Path       string        `mapstructure:"path" json:"path"`
	MaxSize    int           `mapstructure:"max-size" json:"MaxSize"`
	MaxAge     int           `mapstructure:"max-age" json:"maxAge"`
	MaxBackups int           `mapstructure:"max-backups" json:"maxBackups"`
	Compress   bool          `mapstructure:"compress" json:"compress"`
}
