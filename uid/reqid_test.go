package uid

import (
	"github.com/templexxx/tsc"
	"testing"
	"time"
)

func TestMakeParseReqID(t *testing.T) {

	// Because it's fast, second ts may not change.
	expTime := time.Unix(0, tsc.UnixNano())
	reqID := MakeReqID()
	ts := ParseRequestTimestamp(reqID)
	if expTime.Unix() != ts/int64(time.Second) &&
		expTime.Unix()+1 != ts/int64(time.Second) { // May meet critical point.
		t.Fatal("mismatch", expTime.Unix(), ts/int64(time.Second))
	}
}

func BenchmarkMakeReqID(b *testing.B) {

	for i := 0; i < b.N; i++ {
		_ = MakeReqID()
	}
}
