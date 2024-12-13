package randring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestU64Ring_TryPopOnlyOne(t *testing.T) {
	r := New(4)
	v, ok := r.TryPop()
	assert.False(t, ok)
	assert.Equal(t, uint64(0), v)
	r.Push(1)
	v, ok = r.TryPop()
	assert.True(t, ok)
	assert.Equal(t, uint64(1), v)
	r.Push(2)
	v, ok = r.TryPop()
	assert.True(t, ok)
	assert.Equal(t, uint64(2), v)
}

func TestU64Ring_TryPop(t *testing.T) {

	r := New(4)
	for i := 0; i < 1<<4; i++ {
		r.Push(uint64(i))
	}

	assert.Equal(t, uint64(16-1), r.writeIndex)

	cnt := 0
	for {
		_, ok := r.TryPop()
		if !ok {
			break
		}
		cnt++
	}
	assert.Equal(t, 1, cnt)
}

func TestU64Ring_TryPopSlowPop(t *testing.T) {

	r := New(4)

	for i := 0; i < 1<<4; i++ {
		r.Push(uint64(i))
	}
	r.TryPop()
	r.TryPop()
	for i := 16; i < 1<<5; i++ {
		r.Push(uint64(i))
	}

	assert.Equal(t, uint64(31), r.writeIndex)

	cnt := 0
	for {
		_, ok := r.TryPop()
		if !ok {
			break
		}
		cnt++
	}
	assert.Equal(t, 1, cnt)
}
