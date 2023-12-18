// Copyright 2016 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package typeutil

import (
	"encoding/binary"
	"math/rand"
	"testing"
	"time"
)

func TestParseTimestamp(t *testing.T) {
	for i := 0; i < 3; i++ {
		t0 := time.Now().Add(time.Second * time.Duration(rand.Int31n(1000)))
		data := uint64ToBytes(uint64(t0.UnixNano()))
		nt, err := ParseTimestamp(data)
		if err != nil {
			t.Fatal(err)
		}
		if !nt.Equal(t0) {
			t.Fatal("mismatch")
		}
	}
	data := []byte("pkg")
	nt, err := ParseTimestamp(data)
	if err == nil {
		t.Fatal("should failed")
	}

	if !nt.Equal(ZeroTime) {
		t.Fatal("mismatch")
	}
}

func TestSubTimeByWallClock(t *testing.T) {
	for i := 0; i < 3; i++ {
		r := rand.Int31n(1000)
		t1 := time.Now()
		t2 := t1.Add(time.Second * time.Duration(r))
		duration := SubTimeByWallClock(t2, t1)
		if duration != time.Second*time.Duration(r) {
			t.Fatal("mismatch")
		}
	}
}

func uint64ToBytes(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, v)
	return b
}
