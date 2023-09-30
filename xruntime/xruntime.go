// Package xruntime provides low-level runtime operations.
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
