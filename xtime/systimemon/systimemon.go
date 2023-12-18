package systimemon

import (
	"context"
	"fmt"
	"time"

	"github.com/templexxx/crank/xlog"
)

// StartMonitor calls systimeErrHandler if system time jump backward.
func StartMonitor(ctx context.Context, now func() int64, systimeErrHandler func()) {
	xlog.Info("start system time monitor")
	tick := time.NewTicker(100 * time.Millisecond)
	defer tick.Stop()
	for {
		last := now()
		select {
		case <-tick.C:
			if n := now(); now() < last {
				xlog.Error(fmt.Sprintf("system time jump backward: last: %d, now: %d", last, n))
				systimeErrHandler()
			}
		case <-ctx.Done():
			return
		}
	}
}
