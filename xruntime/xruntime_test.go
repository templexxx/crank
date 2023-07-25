package xruntime

import "testing"

func BenchmarkDoNothing(b *testing.B) {

	for i := 0; i < b.N; i++ {
		ProcYield(1000)
	}
}

// TODO using TSC register to report cost (as property testing)
