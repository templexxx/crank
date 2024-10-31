package uid

import (
	"math/rand"
	"time"

	"github.com/templexxx/tsc"
)

const (
	// epoch is an Unix time.
	// 2020-06-03T08:39:34.000+0800.
	epoch     int64 = 1591144774
	epochNano       = epoch * int64(time.Second)

	randIDBits    = 2
	timestampBits = 62 // more than 100 years after epoch
)

// reqid struct:
// +-----------+---------------+
// | randID(2) | timestamp(62) |
// +-----------+---------------+
//
// Total length: 8B (After hex encoding, it's 16B).
//
// randID: 2bit
// timestamp: 62bit
//
// Because timestamp's precision is nanosecond (details see tsc.UnixNano()),
// and getting timestamp has cost too,
// so it's almost impossible to find two same reqid with 2bit randID.

var _randID = uint64(rand.New(rand.NewSource(tsc.UnixNano())).Int63n(1 << randIDBits))

// MakeReqID makes a request ID.
// Request ID is encoded in 64bit unsigned integer.
//
// Warn:
// Maybe not unique but it's acceptable.
func MakeReqID() uint64 {

	return _randID<<timestampBits | (uint64(tsc.UnixNano()) - uint64(epochNano))
}

const reqTSMask = (1 << timestampBits) - 1

// ParseRequestTimestamp gets unix timestamp from reqID.
func ParseRequestTimestamp(reqID uint64) int64 {
	reqTS := reqID & reqTSMask
	return toUnixTS(reqTS)
}

// toUnixTS converts ReqID timestamp to unix timestamp.
func toUnixTS(ts uint64) int64 {
	return int64(ts) + epochNano
}
