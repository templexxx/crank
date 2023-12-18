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
	"runtime"
)

var (
	_global *ErrorLogger
)

// InitGlobalLogger inits Global var.
// warn: It's unsafe for concurrent use.
func InitGlobalLogger(logger *ErrorLogger) {
	_global = logger
}

// Write implements io.Writer
func Write(p []byte) (n int, err error) {
	_global.Error(0, string(p))
	return len(p), nil
}

func Error(msg string) {
	_global.Error(0, msg)
}

func Info(msg string) {
	_global.Info(0, msg)
}

func Warn(msg string) {
	_global.Warn(0, msg)
}

func Debug(msg string) {
	_global.Debug(0, msg)
}

func Fatal(msg string) {
	_global.Fatal(0, msg)
}

func Panic(msg string) {
	_global.Panic(0, msg)
}

func Errorf(format string, args ...interface{}) {
	_global.Error(0, fmt.Sprintf(format, args...))
}

func Infof(format string, args ...interface{}) {
	_global.Info(0, fmt.Sprintf(format, args...))
}

func Warnf(format string, args ...interface{}) {
	_global.Warn(0, fmt.Sprintf(format, args...))
}

func Debugf(format string, args ...interface{}) {
	_global.Debugf(0, format, args...)
}

func Fatalf(format string, args ...interface{}) {
	_global.Fatal(0, fmt.Sprintf(format, args...))
}

func Panicf(format string, args ...interface{}) {
	_global.Panic(0, fmt.Sprintf(format, args...))
}

func ErrorID(reqid uint64, msg string) {
	_global.Error(reqid, msg)
}

func InfoID(reqid uint64, msg string) {
	_global.Info(reqid, msg)
}

func WarnID(reqid uint64, msg string) {
	_global.Warn(reqid, msg)
}

func DebugID(reqid uint64, msg string) {
	_global.Debug(reqid, msg)
}

func FatalID(reqid uint64, msg string) {
	_global.Fatal(reqid, msg)
}

func PanicID(reqid uint64, msg string) {
	_global.Panic(reqid, msg)
}

func ErrorIDf(reqid uint64, format string, args ...interface{}) {
	_global.Error(reqid, fmt.Sprintf(format, args...))
}

func InfoIDf(reqid uint64, format string, args ...interface{}) {
	_global.Info(reqid, fmt.Sprintf(format, args...))
}

func WarnIDf(reqid uint64, format string, args ...interface{}) {
	_global.Warn(reqid, fmt.Sprintf(format, args...))
}

func DebugIDf(reqid uint64, format string, args ...interface{}) {
	_global.Debugf(reqid, format, args)
}

func FatalIDf(reqid uint64, format string, args ...interface{}) {
	_global.Fatal(reqid, fmt.Sprintf(format, args...))
}

func PanicIDf(reqid uint64, format string, args ...interface{}) {
	_global.Panic(reqid, fmt.Sprintf(format, args...))
}

func Warningln(format string, args ...interface{}) {
	_global.Warn(0, fmt.Sprintf(format, args))
}

func Warningf(format string, args ...interface{}) {
	_global.Warn(0, fmt.Sprintf(format, args))
}

func Errorln(format string, args ...interface{}) {
	_global.Error(0, fmt.Sprintf(format, args))
}

func Fatalln(args ...interface{}) {
	_global.Fatal(0, fmt.Sprintln(args))
}

func Debugln(args ...interface{}) {
	_global.Debug(0, fmt.Sprintln(args))
}

func Tracef(format string, args ...interface{}) {
	_global.Debugf(0, format, args)
}

// Sync syncs _global.
func Sync() error {
	return _global.Sync()
}

// Close closes _global.
func Close() error {
	return _global.Close()
}

func SetLevel(level string) error {

	return _global.SetLevel(level)
}

// GetLvl returns lvl in string.
func GetLvl() string {
	return _global.GetLvl()
}

// GetLogger returns _global logger.
func GetLogger() *ErrorLogger {
	return _global
}

type GRPCLogV2 struct {
	logger *ErrorLogger
}

func (g *GRPCLogV2) Info(args ...interface{}) {

	g.logger.Info(0, fmt.Sprint(args))
}

func (g *GRPCLogV2) Infoln(args ...interface{}) {
	g.logger.Info(0, fmt.Sprintln(args))
}

func (g *GRPCLogV2) Infof(format string, args ...interface{}) {
	g.logger.Info(0, fmt.Sprintf(format, args))
}

func (g *GRPCLogV2) Warning(args ...interface{}) {
	g.logger.Warn(0, fmt.Sprint(args))
}

func (g *GRPCLogV2) Warningln(args ...interface{}) {
	g.logger.Warn(0, fmt.Sprintln(args))
}

func (g *GRPCLogV2) Warningf(format string, args ...interface{}) {
	g.logger.Warn(0, fmt.Sprintf(format, args))
}

func (g *GRPCLogV2) Warnf(format string, args ...interface{}) {
	g.logger.Warn(0, fmt.Sprintf(format, args))
}

func (g *GRPCLogV2) Error(args ...interface{}) {
	g.logger.Error(0, fmt.Sprint(args))
}

func (g *GRPCLogV2) Errorln(args ...interface{}) {
	g.logger.Error(0, fmt.Sprintln(args))
}

func (g *GRPCLogV2) Errorf(format string, args ...interface{}) {
	g.logger.Error(0, fmt.Sprintf(format, args))
}

func (g *GRPCLogV2) Fatal(args ...interface{}) {
	g.logger.Fatal(0, fmt.Sprint(args))
}

func (g *GRPCLogV2) Fatalln(args ...interface{}) {
	g.logger.Fatal(0, fmt.Sprintln(args))
}

func (g *GRPCLogV2) Fatalf(format string, args ...interface{}) {
	g.logger.Fatal(0, fmt.Sprintf(format, args))
}

func (g *GRPCLogV2) Debug(args ...interface{}) {
	g.logger.Debug(0, fmt.Sprint(args))
}

func (g *GRPCLogV2) Debugln(args ...interface{}) {
	g.logger.Debug(0, fmt.Sprintln(args))
}

func (g *GRPCLogV2) Debugf(format string, args ...interface{}) {
	g.logger.Debugf(0, format, args)
}

// no trace in xlog, using debug replaceing.

func (g *GRPCLogV2) Tracef(format string, args ...interface{}) {
	g.logger.Debugf(0, format, args)
}

func (g *GRPCLogV2) Trace(args ...interface{}) {
	g.logger.Debug(0, fmt.Sprint(args))
}

func (g *GRPCLogV2) V(l int) bool {
	return true
}

func GetGRPCLoggerV2() *GRPCLogV2 {
	return &GRPCLogV2{GetLogger()}
}

// LogPanic logs the panic reason and stack, then exit the process.
// Commonly used with a `defer`.
func LogPanic() {
	if e := recover(); e != nil {
		stackTrace := make([]byte, 1<<20)
		n := runtime.Stack(stackTrace, false)
		msg := fmt.Sprintf("panic occured: %v\nStack trace: %s", e, stackTrace[:n])
		Fatal(msg)
	}
}
