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

// Package config implements ability of loading app's config.
package config

import (
	"flag"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

var path *string

const cFlag = "c"

const defaultConfPath = "/usr/local/zai"

// Init app's config path.
func Init(appName string) {
	path = flag.String(cFlag,
		filepath.Join(defaultConfPath, appName+".toml"),
		"the config file")
}

// Load app's config returns metadata for parsing/adjusting outside,
// process will exit if any error happens.
func Load(cfg interface{}) *toml.MetaData {
	if !flag.Parsed() {
		flag.Parse()
	}

	data, err := ioutil.ReadFile(*path)
	if err != nil {
		log.Fatalln(err)
	}

	meta, err := toml.Decode(string(data), cfg)
	if err != nil {
		log.Fatalln(err)
	}
	return &meta
}
