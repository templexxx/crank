package xbytes

import "unsafe"

// AlignSize calculates size after align.
// The return will be the multiple of align.
func AlignSize(n int64, align int64) int64 {
	if n <= 0 {
		return 0
	}
	return (n + align - 1) &^ (align - 1)
}

// Alignment returns Alignment of the block in memory
// with reference to alignSize.
//
// Can't check Alignment of a zero sized block as &block[0] is invalid.
func Alignment(block []byte, alignSize int) int {
	return int(uintptr(unsafe.Pointer(&block[0])) & uintptr(alignSize-1))
}

// MakeAlignedBlock returns []byte of size BlockSize aligned to a multiple
// of alignSize in memory (must be power of two).
func MakeAlignedBlock(blockSize, alignSize int) []byte {
	block := make([]byte, blockSize+alignSize)
	if alignSize == 0 {
		return block
	}
	a := Alignment(block, alignSize)
	offset := 0
	if a != 0 {
		offset = alignSize - a
	}
	block = block[offset : offset+blockSize]
	// Can't check Alignment of a zero sized block.
	if blockSize != 0 {
		a = Alignment(block, alignSize)
		if a != 0 {
			panic("failed to align block")
		}
	}
	return block
}
