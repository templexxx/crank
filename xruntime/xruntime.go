package xruntime

// ProcYield yields logic processor for a while (decided by cycles).
//
// e.g. when n = 1, it'll cost 30-40ns.
// If you want 400ns wait, n could be 10.
// (CPU-specific dependencies)
//
// It's not a good idea to use time.Sleep as ProcYield, because it may cause
// goroutine being scheduled and the cost will be unpredictable.
func ProcYield(cycles uint32) {
	pyield(cycles)
}

// TODO func ProcYieldNS(nanoseconds int64) {}
// nanoseconds -> cycles (depends on CPU family which detected in init)
