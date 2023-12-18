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

// Package xlogtest implements functions to generate temporary loggers.
package xlogtest

import (
	"io/ioutil"
	"os"

	"github.com/templexxx/crank/xlog"
	"github.com/zaibyte/nanozap"
	"github.com/zaibyte/nanozap/zapcore"
)

func init() {
	New(true)
}

// New init global logger with stdout.
func New(discard bool) {

	var syncer zapcore.WriteSyncer
	if !discard {
		syncer = &Stdouter{}
	} else {
		syncer = &Discarder{}
	}

	lvl := nanozap.NewAtomicLevel()
	lvl.SetLevel(zapcore.InfoLevel)
	core := zapcore.NewCore(zapcore.NewJSONEncoder(xlog.DefaultEncoderConfig()), syncer, lvl)

	el := &xlog.ErrorLogger{
		Logger:   nanozap.New(core),
		Lvl:      lvl,
		Rotation: nil,
	}
	xlog.InitGlobalLogger(el)
	_ = xlog.SetLevel("debug")
}

// A Syncer is a spy for the Sync portion of zapcore.WriteSyncer.
type Syncer struct {
	err    error
	called bool
}

// SetError sets the error that the Sync method will return.
func (s *Syncer) SetError(err error) {
	s.err = err
}

// Sync records that it was called, then returns the user-supplied error (if
// any).
func (s *Syncer) Sync() error {
	s.called = true
	return s.err
}

// Called reports whether the Sync method was called.
func (s *Syncer) Called() bool {
	return s.called
}

// A Discarder sends all writes to ioutil.Discard.
type Discarder struct{ Syncer }

// Write implements io.Writer.
func (d *Discarder) Write(b []byte) (int, error) {
	return ioutil.Discard.Write(b)
}

type Stdouter struct {
	Syncer
}

func (s *Stdouter) Write(b []byte) (int, error) {
	return os.Stdout.Write(b)
}

func SetLevel(lvl string) {
	_ = xlog.SetLevel(lvl)
}

// Sync syncs logger buffer.
func Sync() {
	_ = xlog.Sync()
}

// Close closes logger.
func Close() {
	_ = xlog.Close()
}
