/*
 * Copyright (c) 2020. Temple3x (temple3x@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package xlog

import (
	"fmt"

	"github.com/zaibyte/nanozap"
	"github.com/zaibyte/nanozap/zapcore"
	"github.com/zaibyte/nanozap/zaproll"
)

// RotateConfig is partly copy from zaproll's Config,
// hiding details in zaproll.
type RotateConfig struct {
	// Maximum size of a log file before it gets rotated.
	// Unit is MB.
	MaxSize int64 `toml:"max_size"`
	// Maximum number of backup log files to retain.
	MaxBackups int
	// Timestamp in backup log file. Default(false) is UTC time.
	LocalTime bool `toml:"local_time"`
}

// ErrorLogger is used for recording the common application log,
// xlog also provides global logger for more convenient.
// In practice, ErrorLogger is just a global logger's container,
// it won't be used directly.
type ErrorLogger struct {
	Logger   *nanozap.Logger
	Lvl      nanozap.AtomicLevel
	Rotation *zaproll.Rotation
}

// ErrLogFields shows error logger output fields.
//
// Warn:
// Sometimes, there is no "x-zai-request-id"" or "x-zai-box-id".
// (It's not from any request)
type ErrLogFields struct {
	Level string `json:"level"`
	Time  string `json:"time"`
	Msg   string `json:"msg"`
	ReqID string `json:"x-zai-request-id"`
	BoxID int64  `json:"x-zai-box-id"`
}

// NewErrorLogger returns a logger with its properties.
//
// Legal Levels:
// info: "info", "INFO", ""
// debug: "debug", "DEBUG"
// warn: "warn", "WARN"
// error: "error", "ERROR"
// panic: "panic", "PANIC"
// fatal: "fatal", "FATAL"
func NewErrorLogger(outputPath, level string, rCfg *RotateConfig) (logger *ErrorLogger, err error) {

	r, err := zaproll.New(&zaproll.Config{
		OutputPath: outputPath,
		MaxSize:    rCfg.MaxSize,
		MaxBackups: rCfg.MaxBackups,
		LocalTime:  rCfg.LocalTime,
	})
	if err != nil {
		return
	}

	lvl := nanozap.NewAtomicLevel()
	err = lvl.UnmarshalText([]byte(level))
	if err != nil {
		return
	}

	core := zapcore.NewCore(zapcore.NewJSONEncoder(DefaultEncoderConfig()), r, lvl)

	return &ErrorLogger{
		Logger:   nanozap.New(core),
		Lvl:      lvl,
		Rotation: r,
	}, nil
}

// DefaultEncoderConfig is the default logger config.
func DefaultEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		TimeKey:        "time",
		ReqIDKey:       "reqid",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
	}
}

// Write implements io.Writer
func (l *ErrorLogger) Write(p []byte) (n int, err error) {
	l.Error(0, string(p))
	return len(p), nil
}

func (l *ErrorLogger) Error(reqid uint64, msg string) {
	l.Logger.Error(reqid, msg)
}

func (l *ErrorLogger) Info(reqid uint64, msg string) {
	l.Logger.Info(reqid, msg)
}

func (l *ErrorLogger) Warn(reqid uint64, msg string) {
	l.Logger.Warn(reqid, msg)
}

func (l *ErrorLogger) Debug(reqid uint64, msg string) {
	l.Logger.Debug(reqid, msg)
}

func (l *ErrorLogger) Fatal(reqid uint64, msg string) {
	l.Logger.Fatal(reqid, msg)
}

func (l *ErrorLogger) Panic(reqid uint64, msg string) {
	l.Logger.Panic(reqid, msg)
}

func (l *ErrorLogger) Errorf(reqid uint64, format string, args ...interface{}) {
	l.Logger.Error(reqid, fmt.Sprintf(format, args...))
}

func (l *ErrorLogger) Infof(reqid uint64, format string, args ...interface{}) {
	l.Logger.Info(reqid, fmt.Sprintf(format, args...))
}

func (l *ErrorLogger) Warnf(reqid uint64, format string, args ...interface{}) {
	l.Logger.Warn(reqid, fmt.Sprintf(format, args...))
}

func (l *ErrorLogger) Debugf(reqid uint64, format string, args ...interface{}) {
	l.Logger.Debugf(reqid, format, args...)
}

func (l *ErrorLogger) Fatalf(reqid uint64, format string, args ...interface{}) {
	l.Logger.Fatal(reqid, fmt.Sprintf(format, args...))
}

func (l *ErrorLogger) Panicf(reqid uint64, format string, args ...interface{}) {
	l.Logger.Panic(reqid, fmt.Sprintf(format, args...))
}

func (l *ErrorLogger) Printf(format string, args ...interface{}) {
	l.Logger.Errorf(0, fmt.Sprintf(format, args...))
}

// Sync syncs ErrorLogger.
func (l *ErrorLogger) Sync() error {
	return l.Logger.Sync()
}

func (l *ErrorLogger) Close() error {
	l.Logger.Close()
	if l.Rotation != nil {
		return l.Rotation.Close()
	}
	return nil
}

func (l *ErrorLogger) SetLevel(level string) error {
	lvl := nanozap.NewAtomicLevel()
	err := lvl.UnmarshalText([]byte(level))
	if err != nil {
		return err
	}

	l.Lvl.SetLevel(lvl.Level())
	return nil
}

// GetLvl return lvl in string.
func (l *ErrorLogger) GetLvl() string {
	return l.Lvl.String()
}
