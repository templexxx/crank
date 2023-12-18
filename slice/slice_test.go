package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlice(t *testing.T) {
	cases := []struct {
		a      []int
		anyOf  bool
		noneOf bool
		allOf  bool
	}{
		{[]int{}, false, true, true},
		{[]int{1, 2, 3}, true, false, false},
		{[]int{1, 3}, false, true, false},
		{[]int{2, 2, 4}, true, false, true},
	}

	for _, c := range cases {
		even := func(i int) bool { return c.a[i]%2 == 0 }
		assert.Equal(t, AnyOf(c.a, even), c.anyOf)
		assert.Equal(t, AllOf(c.a, even), c.allOf)
		assert.Equal(t, NoneOf(c.a, even), c.noneOf)
	}
}
