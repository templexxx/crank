package mhlc

import (
	"runtime"
	"testing"

	"github.com/templexxx/crank/xtime/hlc/hlcutil"

	"github.com/stretchr/testify/assert"
)

func BenchmarkLHLC_Next(b *testing.B) {
	l := New()

	for i := 0; i < b.N; i++ {
		_ = l.Next()
	}
}

func BenchmarkLHLC_NextConcurrency(b *testing.B) {

	l := New()

	b.SetParallelism(runtime.NumCPU())
	b.RunParallel(func(pb *testing.PB) {
		for i := 0; pb.Next(); i++ {
			_ = l.Next()
		}
	})
}

func TestLHLC_Next(t *testing.T) {
	l := New()

	m := make(map[uint64]bool)
	for i := 0; i < hlcutil.LogicalMask+1024; i++ {
		ts := l.Next()
		p, _ := hlcutil.ParseTS(ts)
		m[p] = true
	}
	assert.Equal(t, 2, len(m))
}
