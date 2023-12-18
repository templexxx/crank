package xtime

import (
	"sync"
	"time"

	"github.com/templexxx/crank/xlog"
)

var nopTime = make(chan time.Time)

func init() {
	close(nopTime)
}

// GetTimerEvent gets a single time event.
// It could be used in a time event loop.
//
// It's designed for process which need accurate control of time event.
// We want the next event will come after the last event finishing in duration.
// If we use time.Ticker there, the worst case is the cost of event is high, which means
// just after the event finishing, the ticker will tick again, that's not we want.
//
// e.g.
// t := time.NewTimer(duration)
// var tChan <-chan time.Time
//
//	for {
//		var m *msg
//
//		select {
//			case m = <-msgChan:
//			case <-tChan:
//				foo()
//				tChan = nil
//				continue
//			}
//		}
//
//		if tChan == nil {
//			tChan = xtime.GetTimeEvent(t, s.FlushDelay)
//		}
//		...
func GetTimerEvent(t *time.Timer, duration time.Duration) <-chan time.Time {
	if duration <= 0 {
		return nopTime
	}

	if !t.Stop() {
		// Exhaust expired timer's chan.
		select {
		case <-t.C:
		default:
		}
	}
	t.Reset(duration)
	return t.C
}

var timerPool sync.Pool

func AcquireTimer(timeout time.Duration) *time.Timer {
	tv := timerPool.Get()
	if tv == nil {
		return time.NewTimer(timeout)
	}

	t := tv.(*time.Timer)
	if t.Reset(timeout) {
		xlog.Panic("bug: active timer trapped into AcquireTimer()")
	}
	return t
}

func ReleaseTimer(t *time.Timer) {
	if !t.Stop() {
		// Collect possibly added time from the channel
		// if timer has been stopped and nobody collected its' value.
		select {
		case <-t.C:
		default:
		}
	}

	timerPool.Put(t)
}
