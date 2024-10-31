// Package mhlc(memory hybrid logical clock) implements hlc.HLC interface, that combines the best of
// logical clocks and physical clocks.
// It's a clock which never goes backwards in one instance during the runtime.
//
// Warn:
// After instance crash, it has chance going backwards because we have no reference time.
// But it's rare to happen because we cannot make a new device that fast(in dozens ms, which is NTP jitter).

package mhlc

import (
	"sync/atomic"
	"time"

	"github.com/templexxx/crank/xtest"
	"github.com/templexxx/crank/xtime/hlc/hlcutil"
	"github.com/templexxx/tsc"
)

type MHLC struct {
	lastTS uint64
}

// New creates an MHLC for application.
// Each instance should have one.
func New() *MHLC {

	c := &MHLC{
		lastTS: hlcutil.MakeTS(nowInMill(), 0),
	}
	return c
}

// Next returns a timestamp.
func (c *MHLC) Next() (ts uint64) {
	for {
		last, p, l, ok := c.next()
		if !ok {
			time.Sleep(200 * time.Microsecond) // Sleep for a while, wait clock move on.
			continue
		}

		ts = hlcutil.MakeTS(p, l)
		if atomic.CompareAndSwapUint64(&c.lastTS, last, ts) {
			return
		}
		xtest.DoNothing(16) // Avoiding too frequently, reducing CPU wasting.
	}
}

func (c *MHLC) next() (last, phy, logic uint64, ok bool) {
	last = atomic.LoadUint64(&c.lastTS)
	lp, ll := hlcutil.ParseTS(last)

	phy = lp

	logic = (ll + 1) & hlcutil.LogicalMask
	if logic == 0 { // Logical overflow, need new physical.
		now := nowInMill()
		if lp >= now {
			return 0, 0, 0, false // Time go backwards.
		}
		phy = now
	}
	return last, phy, logic, true
}

func nowInMill() uint64 {
	return uint64(tsc.UnixNano() / int64(time.Millisecond))
}
