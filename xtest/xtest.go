// Package xtest is not a testing framework, just a collection of testing helper tools.
package xtest

import (
	"flag"
	"math/rand"

	"github.com/templexxx/tsc"
)

var _propEnabled = flag.Bool("xtest.prop", false, "enable properties testing or not")
var _randSeed = rand.New(rand.NewSource(tsc.UnixNano()))

// IsPropEnabled returns enable properties testing or not.
// Default is false.
//
// Usage:
//
// no properties testing: go test -xtest.prop=false -v or go test -v
// run properties testing: go test -xtest.prop=true -v
//
// In your testing functions: IsPropEnabled will detect passed argument.
func IsPropEnabled() bool {
	if !flag.Parsed() {
		flag.Parse()
	}

	return *_propEnabled
}

// FillRand fills p with random bytes.
func FillRand(p []byte) {
	_randSeed.Read(p)
}

// DoNothing does nothing, only for some framework testing to test pure framework cost.
// Using n to control the function total cost, actually a spin is inside this function.
//
// e.g. when n = 1, it'll cost 30-40ns.
// If you want 400ns wait, n could be 10.
//
// It's not a good idea to use time.Sleep as DoNothing, because it'll bring
// goroutine scheduler make cost unpredictable.
// TODO more testing
func DoNothing(n uint32) {
	spin(n)
}
