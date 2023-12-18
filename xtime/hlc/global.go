package hlc

import "github.com/templexxx/crank/xtime/hlc/mhlc"

var (
	_global HLC = mhlc.New()
)

// InitGlobalHLC Inits Global var.
// warn: It's unsafe for concurrent use.
func InitGlobalHLC(h HLC) {
	_global = h
}

// Next returns timestamp by global HLC.
func Next() uint64 {
	return _global.Next()
}
