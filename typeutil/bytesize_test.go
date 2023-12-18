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

	"github.com/BurntSushi/toml"
)

type byteSizeExample struct {
	Size ByteSize `json:"size" toml:"size"`
}

func TestByteSize_JSON(t *testing.T) {
	ex := &byteSizeExample{}

	text := []byte(`{"size":"1.498GiB"}`)
	err := json.Unmarshal(text, ex)
	if err != nil {
		t.Fatal(err)
	}

	if ex.Size != 1608465252 {
		t.Fatal("unmarshal mismatch")
	}

	b, err := json.Marshal(ex)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(text, b) {
		t.Fatal("marshal mismatch", string(b))
	}
}

func TestByteSize_TOML(t *testing.T) {
	ex := &byteSizeExample{}

	text := []byte(`size = "1.498GiB"`)
	err := toml.Unmarshal(text, ex)
	if err != nil {
		t.Fatal(err)
	}

	if ex.Size != 1608465252 {
		t.Fatal("unmarshal mismatch")
	}
}
