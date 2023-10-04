package xmath

import (
	"testing"
)

func TestRound(t *testing.T) {
	f := 1.1
	var i float64
	for i = 0; i < 0.05; i += 0.01 {
		if Round(f+i, 1) != 1.1 {
			testRound(t, f+i, 1.1, Round(f+i, 1), 1)
		}
	}
	for i = 0.05; i < 0.1; i += 0.01 {
		if Round(f+i, 1) != 1.2 {
			testRound(t, f+i, 1.2, Round(f+i, 1), 1)
		}
	}
}

func testRound(t *testing.T, input, exp, got float64, decimal int) {
	if exp != got {
		t.Fatalf("mismatch: input=%f, exp=%f, got=%f, decimal=%d", input, exp, got, decimal)
	}
}

func TestNextPow2(t *testing.T) {

	testNextPow2(t, 0, 1, NextPow2(0))
	testNextPow2(t, 1, 1, NextPow2(1))
	testNextPow2(t, 2, 2, NextPow2(2))
	testNextPow2(t, 3, 4, NextPow2(3))
	testNextPow2(t, 4, 4, NextPow2(4))

	for i := 5; i <= 1025; i++ {
		testNextPow2(t, uint64(i), slowNextPow2(uint64(i)), NextPow2(uint64(i)))
	}
}

func slowNextPow2(n uint64) uint64 {
	var p uint64 = 1
	for {
		if p < n {
			p *= 2
		} else {
			break
		}
	}
	return p
}

func testNextPow2(t *testing.T, n, exp, got uint64) {
	if exp != got {
		t.Fatalf("mismatch: n=%d, exp=%d, got=%d", n, exp, got)
	}
}
