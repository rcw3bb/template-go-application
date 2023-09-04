// Package logger provides the logging capability of the application.
// Author: Ron Webb
package logger

import (
	"go-app-template/config/appinfo"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"runtime"
)

type logger struct {
	logLogger *zap.Logger
	logFile   *os.File
	logLevel  zapcore.Level
}

const (
	logDir      string = "./logs"
	logFilename string = appinfo.Application + ".log"
)

var log = &logger{
	logLevel: zap.InfoLevel,
}

var IsDebugEnabled bool

func createLogDir() {
	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		panic(err)
	}
}

func InitLogger() {

	if log.logFile != nil {
		return
	}

	if IsDebugEnabled {
		log.logLevel = zap.DebugLevel
	}

	createLogDir()

	// Create write syncer that directs logger entries to both the console and a logger file.
	// Console output.
	consoleDebugging := zapcore.Lock(os.Stdout)

	logFile := logDir + "/" + logFilename

	fileOutput, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic("Failed to open logger file: " + err.Error())
	}

	// File output.
	log.logFile = fileOutput

	fileDebugging := zapcore.AddSync(log.logFile)

	// Create a zap logger configuration.
	encoderCfg := zapcore.EncoderConfig{
		TimeKey:    "timestamp",
		EncodeTime: zapcore.ISO8601TimeEncoder,
		LevelKey:   "level",
		NameKey:    "loggerName",
		MessageKey: "message",
		// Customize the output format.
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
	}

	// Customize the encoder to include the goroutine ID.
	encoderCfg.EncodeCaller = func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(caller.TrimmedPath())
		enc.AppendString(" (goroutine " + log.getGoroutineID() + ")")
	}

	// Configure console output core.
	consoleCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderCfg), // Use customized console encoder.
		consoleDebugging,
		zap.NewAtomicLevelAt(log.logLevel), // Set logging level for console.
	)

	// Configure file output core.
	fileCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg), // Use JSON encoder for file output.
		fileDebugging,
		zap.NewAtomicLevelAt(log.logLevel), // Set logging level for file.
	)

	// Combine the two cores to output logs to both console and file.
	combinedCore := zapcore.NewTee(consoleCore, fileCore)

	// Create the global logger instance.
	log.logLogger = zap.New(combinedCore)
}

// Close must be called to close logger properly.
func Close() {
	log.logFile.Close()
}

// Function to get the current goroutine ID (thread name).
func (logger *logger) getGoroutineID() string {
	buf := make([]byte, 64)
	buf = buf[:runtime.Stack(buf, false)]
	return string(buf)
}

// Warn use to log warning.
func Warn(msg string, fields ...zapcore.Field) {
	log.logLogger.Warn(msg, fields...)
}

// Debug use to log debug information.
func Debug(msg string, fields ...zapcore.Field) {
	log.logLogger.Debug(msg, fields...)
}

// Info use to log information.
func Info(msg string, fields ...zapcore.Field) {
	log.logLogger.Info(msg, fields...)
}

// Error use to log error.
func Error(msg string, fields ...zapcore.Field) {
	log.logLogger.Error(msg, fields...)
}

// Panic use to log panic.
func Panic(msg string, fields ...zapcore.Field) {
	log.logLogger.Panic(msg, fields...)
}

// Fatal use to log fatal.
func Fatal(msg string, fields ...zapcore.Field) {
	log.logLogger.Fatal(msg, fields...)
}
