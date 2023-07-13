// Package logger
// @author tabuyos
// @since 2023/7/7
// @description logger
package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"metis/config/constant"
	"metis/util/file"
	"os"
	"path"
)

var logger *zap.Logger

func init() {
	_, ok := file.IsExists(constant.LogDir)
	if !ok {
		if err := os.MkdirAll(constant.LogDir, 0766); err != nil {
			panic(err)
		}
	}
	logger = zap.New(commonCore(), zap.AddCaller())
	//logger = zap.New(teeCore(), zap.AddCaller())
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func commonCore() zapcore.Core {
	accessFile := &lumberjack.Logger{
		Filename:   path.Join(constant.LogDir, "access.log"),
		MaxSize:    20,
		MaxBackups: 0,
		MaxAge:     30,
		Compress:   false,
	}

	multiWriter := io.MultiWriter(accessFile, os.Stdout)

	return zapcore.NewCore(getEncoder(), zapcore.AddSync(multiWriter), zapcore.DebugLevel)
}

func teeCore() zapcore.Core {
	encoder := getEncoder()

	accessFile := &lumberjack.Logger{
		Filename:   path.Join(constant.LogDir, "access.log"),
		MaxSize:    10,
		MaxBackups: 0,
		MaxAge:     30,
		Compress:   false,
	}
	accessMultiWriter := io.MultiWriter(accessFile, os.Stdout)
	accessCore := zapcore.NewCore(encoder, zapcore.AddSync(accessMultiWriter), zapcore.DebugLevel)

	errorFile := &lumberjack.Logger{
		Filename:   path.Join(constant.LogDir, "access-error.log"),
		MaxSize:    10,
		MaxBackups: 0,
		MaxAge:     30,
		Compress:   false,
	}
	errorCore := zapcore.NewCore(encoder, zapcore.AddSync(errorFile), zapcore.ErrorLevel)

	return zapcore.NewTee(accessCore, errorCore)
}

func UseLogger() zap.Logger {
	return *logger
}
