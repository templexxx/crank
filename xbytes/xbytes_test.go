package xbytes

import (
	"testing"
	"unsafe"
)

func TestAlignSize(t *testing.T) {
	var align int64 = 1 << 12
	var i int64
	for i = 1; i <= align; i++ {
		n := AlignSize(i, align)
		if n != align {
			t.Fatal("align mismatch", n, i)
		}
	}
	for i = align + 1; i <= align*2; i++ {
		n := AlignSize(i, align)
		if n != align*2 {
			t.Fatal("align mismatch")
		}
	}
}

func TestAlignedBlock(t *testing.T) {
	for i := 1; i < 33; i++ {
		b := MakeAlignedBlock(i, 16)
		if uintptr(unsafe.Pointer(&b[0]))&15 != 0 {
			t.Fatal("aligned mismatch")
		}
	}
}
