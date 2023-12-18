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
	"errors"
	"fmt"
	"time"
)

// ZeroTime is a zero time.
var ZeroTime = time.Time{}

// ParseTimestamp returns a timestamp for a given byte slice.
func ParseTimestamp(data []byte) (time.Time, error) {
	nano, err := bytesToUint64(data)
	if err != nil {
		return ZeroTime, err
	}

	return time.Unix(0, int64(nano)), nil
}

// bytesToUint64 converts a byte slice to uint64.
func bytesToUint64(b []byte) (uint64, error) {
	if len(b) != 8 {
		return 0, errors.New(
			fmt.Sprintf("invalid data, must 8 bytes, but %d", len(b)))
	}

	return binary.BigEndian.Uint64(b), nil
}

// SubTimeByWallClock returns the duration between two different timestamps.
func SubTimeByWallClock(after time.Time, before time.Time) time.Duration {
	return time.Duration(after.UnixNano() - before.UnixNano())
}
