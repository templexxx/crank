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
