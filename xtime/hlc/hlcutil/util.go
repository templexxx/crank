package hlcutil

const (
	physicalShiftBits = 18
	logicalBits       = (1 << physicalShiftBits) - 1
	LogicalMask       = 0x3FFFF
)

// MakeTS makes timestamp.
func MakeTS(phy, logic uint64) uint64 {
	return phy<<physicalShiftBits | logic&LogicalMask
}

// ParseTS parses the ts to (physical,logical).
func ParseTS(ts uint64) (uint64, uint64) {
	logical := ts & logicalBits
	physical := ts >> physicalShiftBits
	// physicalTime := time.Unix(int64(physical/1000), int64(physical)%1000*time.Millisecond.Nanoseconds())
	return physical, logical
}
