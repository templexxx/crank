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

// Package xlog provides logger features.
//
// All log entries are encoded in JSON,
package xlog

import (
	"path/filepath"

	"github.com/templexxx/crank/config"
)

const (
	// DefaultLogRoot is the default log files path root.
	// e.g.:
	// <DefaultLogRoot>/<appName>/access.log
	// & <DefaultLogRoot>/<appName>/error.log
	DefaultLogRoot = "/var/log/xlog"
)

// Config is the log configs of a zai application.
type Config struct {
	Output string       `toml:"output"`
	Level  string       `toml:"level"`
	Rotate RotateConfig `toml:"rotate"`
}

// MakeLogger init global error logger and returns logger for application.
func (c *Config) MakeLogger(appName string) (el *ErrorLogger, err error) {

	config.Adjust(&c.Output, filepath.Join(DefaultLogRoot, appName, "error.log"))
	config.Adjust(&c.Level, "info")

	el, err = NewErrorLogger(c.Output, c.Level, &c.Rotate)
	if err != nil {
		return
	}

	InitGlobalLogger(el)

	return
}
