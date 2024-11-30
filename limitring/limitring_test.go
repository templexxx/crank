package limitring

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

func TestLimitRing(t *testing.T) {

	r := New(0)

	for i := 0; i < MinCap; i++ {
		a := i
		err := r.Push(unsafe.Pointer(&a))
		assert.Nil(t, err)
		d, ok := r.Pop()
		assert.True(t, ok)
		assert.Equal(t, a, *(*int)(d))
	}
}
