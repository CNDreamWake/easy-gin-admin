package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

var logger *zap.Logger

func GetLog() *zap.Logger {
	return logger
}

// Initialization 初始化日志组件。
// conf: 日志配置信息，包括日志级别和输出路径等。
// 返回值: 返回一个配置好的zap.Logger实例，用于后续的日志记录。
func Initialization(conf *Log) {
	// 解析日志级别，若失败则默认使用InfoLevel
	level, err := zapcore.ParseLevel(conf.Level)
	if err != nil {
		level = zap.InfoLevel
	}

	// 创建日志输出的写入同步器，若失败则默认输出到Stdout
	var ws zapcore.WriteSyncer
	ws, err = createWriter(conf, "2006-01-02-15-04")
	if err != nil {
		ws = zapcore.AddSync(os.Stdout)
	}

	// 根据配置创建并返回一个日志器
	logger = zap.New(zapcore.NewCore(getDefaultEnc(), ws, level), zap.WithCaller(true))
}

// getDefaultEnc 创建并返回一个默认的日志编码器。
func getDefaultEnc() zapcore.Encoder {
	return zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     customTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	})
}

// customTimeEncoder 格式化日志输出时间
func customTimeEncoder(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
	encoder.AppendString(t.Format("2006/01/02 - 15:04:05.000"))
}
