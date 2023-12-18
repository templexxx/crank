// Copyright 2017 PingCAP, Inc.
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
	"bytes"
	"encoding/json"
	"testing"
)

type stringSliceExample struct {
	SS StringSlice `json:"ss" toml:"ss"`
}

func TestStringSlice_JSON(t *testing.T) {
	ex := &stringSliceExample{}

	text := []byte(`{"ss":"a,b"}`)
	err := json.Unmarshal(text, ex)
	if err != nil {
		t.Fatal(err)
	}

	if !isEqualStrings([]string{"a", "b"}, ex.SS) {
		t.Fatal("unmarshal mismatch")
	}

	b, err := json.Marshal(ex)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(text, b) {
		t.Fatal("marshal mismatch")
	}

	text = []byte(`{}`)
	ex2 := &stringSliceExample{}
	err = json.Unmarshal(text, ex2)
	if err != nil {
		t.Fatal(err)
	}

	if !isEqualStrings(nil, ex2.SS) {
		t.Fatal("unmarshal mismatch")
	}
}

func isEqualStrings(a, b []string) bool {

	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
