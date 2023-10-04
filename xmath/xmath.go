// Package xmath provides useful math helper functions.
package xmath

import (
	"math"
	"math/bits"
)

// Round rounds a float64 and cuts it by n.
// n: decimal places.
// e.g.
// f = 1.001, n = 2, return 1.00
func Round(f float64, n int) float64 {
	pow10n := math.Pow10(n)
	return math.Trunc(f*pow10n+0.5) / pow10n
}

// NextPow2 returns the next power of 2 that is greater than or equal to the input number n.
// For example, if n is 1, it will return 1. If n is 2, it will return 2. If n is 3, it will return 4.
func NextPow2(n uint64) uint64 {
	if n <= 1 {
		return 1
	}

	return 1 << (64 - bits.LeadingZeros64(n-1)) // TODO may use BSR instruction.
}
