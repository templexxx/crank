package systimemon

import (
	"context"
	"sync/atomic"
	"testing"
	"time"

	_ "github.com/templexxx/crank/xlog/xlogtest"
	"github.com/templexxx/tsc"
)

func TestSystimeMonitor(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var jumpForward int32

	trigged := false
	go StartMonitor(ctx,
		func() int64 {
			if !trigged {
				trigged = true
				return tsc.UnixNano()
			}

			return tsc.UnixNano() - 2*int64(time.Second)
		}, func() {
			atomic.StoreInt32(&jumpForward, 1)
		})

	time.Sleep(1 * time.Second)

	if atomic.LoadInt32(&jumpForward) != 1 {
		t.Error("should detect time error")
	}
}
