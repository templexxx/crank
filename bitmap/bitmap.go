// The MIT License (MIT)
//
// Copyright (c) 2015 Bol Christophe

package bitmap

var (
	tA = [8]byte{1, 2, 4, 8, 16, 32, 64, 128}
	tB = [8]byte{254, 253, 251, 247, 239, 223, 191, 127}
)

func dataOrCopy(d []byte, c bool) []byte {
	if !c {
		return d
	}
	ndata := make([]byte, len(d))
	copy(ndata, d)
	return ndata
}

// NewSlice creates a new byteslice with length l (in bits).
// The actual size in bits might be up to 7 bits larger because
// they are stored in a byteslice.
func NewSlice(l int) []byte {
	remainder := l % 8
	if remainder != 0 {
		remainder = 1
	}
	return make([]byte, l/8+remainder)
}

// Get returns the value of bit i from map m.
// It doesn't check the bounds of the slice.
func Get(m []byte, i int) bool {
	return m[i/8]&tA[i%8] != 0
}

// Set sets bit i of map m to value v.
// It doesn't check the bounds of the slice.
func Set(m []byte, i int, v bool) {
	index := i / 8
	bit := i % 8
	if v {
		m[index] = m[index] | tA[bit]
	} else {
		m[index] = m[index] & tB[bit]
	}
}

// GetBit returns the value of bit i of byte b.
// The bit index must be between 0 and 7.
func GetBit(b byte, i int) bool {
	return b&tA[i] != 0
}

// SetBit sets bit i of byte b to value v.
// The bit index must be between 0 and 7.
func SetBit(b byte, i int, v bool) byte {
	if v {
		return b | tA[i]
	}
	return b & tB[i]
}

// SetBitRef sets bit i of byte *b to value v.
func SetBitRef(b *byte, i int, v bool) {
	if v {
		*b = *b | tA[i]
	} else {
		*b = *b & tB[i]
	}
}

// Len returns the length (in bits) of the provided byteslice.
// It will always be a multipile of 8 bits.
func Len(m []byte) int {
	return len(m) * 8
}

// Bitmap is a byteslice with bitmap functions.
// Creating one form existing data is as simple as bitmap := Bitmap(data).
type Bitmap []byte

// New creates a new Bitmap instance with length l (in bits).
func New(l int) Bitmap {
	return NewSlice(l)
}

// Len wraps around the Len function.
func (b Bitmap) Len() int {
	return Len(b)
}

// Get wraps around the Get function.
func (b Bitmap) Get(i int) bool {
	return Get(b, i)
}

// Set wraps around the Set function.
func (b Bitmap) Set(i int, v bool) {
	Set(b, i, v)
}

// Data returns the data of the bitmap.
// If copy is false the actual underlying slice will be returned.
func (b Bitmap) Data(copy bool) []byte {
	return dataOrCopy(b, copy)
}
