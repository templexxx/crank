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

package config

import (
	"testing"
	"time"

	"github.com/templexxx/crank/typeutil"

	"github.com/stretchr/testify/assert"
)

func TestAdjust(t *testing.T) {
	defStr := "string"
	s0 := ""
	Adjust(&s0, defStr)
	if s0 != defStr {
		t.Fatal("adjust string mismatch")
	}

	defInt := 1
	int0 := 0
	Adjust(&int0, defInt)
	if int0 != defInt {
		t.Fatal("adjust int mismatch")
	}

	var defInt64 int64 = 1
	var int640 int64
	Adjust(&int640, defInt64)
	if int640 != defInt64 {
		t.Fatal("adjust int64 mismatch")
	}

	var defUint32 uint32 = 1
	var uint320 uint32
	Adjust(&uint320, defUint32)
	if uint320 != defUint32 {
		t.Fatal("adjust uint32 mismatch")
	}

	var defDuration = time.Second
	var duration0 time.Duration
	Adjust(&duration0, defDuration)
	if duration0 != defDuration {
		t.Fatal("adjust time.Duration mismatch")
	}

	var tuduration0 typeutil.Duration
	Adjust(&tuduration0, defDuration)
	if tuduration0.Duration != defDuration {
		t.Fatal("adjust typeutil.Duration mismatch")
	}

	var defStrings = []string{"def"}
	var strings0 []string
	Adjust(&strings0, defStrings)
	assert.Equal(t, defStrings, strings0, "adjust []string mismatch")

	const defByteSize = typeutil.ByteSize(1)
	var byteSize0 typeutil.ByteSize
	Adjust(&byteSize0, defByteSize)
	assert.Equal(t, defByteSize, byteSize0, "adjust typeutil.ByteSize mismatch")
}
