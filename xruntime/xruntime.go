package xruntime

// ProcYield yields logic processor for a while (decided by cycles).
// Providing stable and low-energy intervals for spin scenarios (e.g. spin lock).
//
// The cost is CPU-specific dependencies.
//
// It's not a good idea to use time.Sleep as ProcYield, because it may cause
// goroutine being scheduled and the cost will be unpredictable.
func ProcYield(cycles uint32) {
	pyield(cycles)
}

// TODO func ProcYieldNS(nanoseconds int64) {} will it work?
// nanoseconds -> cycles (depends on CPU family which detected in init)

// TODO DoSpins (with default cycles for ProcYield)
