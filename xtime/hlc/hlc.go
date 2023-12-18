package hlc

import (
	"github.com/templexxx/tsc"
)

// HLC (hybrid logical clock), that combines the best of logical clocks and physical clocks.
// It's a clock which never goes backwards in one instance.
type HLC interface {
	// Next gets a unique timestamp.
	Next() uint64
}

// WallClock uses wall clock as HLC source,
// we could use it in testing or development env.
type WallClock struct{}

// NewWallClock creates a WallClock.
func NewWallClock() *WallClock {
	return new(WallClock)
}

func (c *WallClock) Next() uint64 {
	return uint64(tsc.UnixNano())
}
