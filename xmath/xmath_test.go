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
