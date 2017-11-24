package contains

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {

	strings := []string{"Hey", "there"}
	assert.True(t, StringSlice(strings).Contains("Hey"))
	assert.True(t, StringSlice(strings).Contains("there"))
	assert.False(t, StringSlice(strings).Contains(""))
	assert.False(t, StringSlice(strings).Contains("foo"))

	ints := []int{1, 2, 3, 4, 5}
	for i := 1; i <= 5; i++ {
		assert.True(t, IntSlice(ints).Contains(i))
	}
	assert.False(t, IntSlice(ints).Contains(6))

	int32s := []int32{1, 2, 3, 4, 5}
	for i := int32(1); i <= 5; i++ {
		assert.True(t, Int32Slice(int32s).Contains(i))
	}
	assert.False(t, Int32Slice(int32s).Contains(6))

	int64s := []int64{1, 2, 3, 4, 5}
	for i := int64(1); i <= 5; i++ {
		assert.True(t, Int64Slice(int64s).Contains(i))
	}
	assert.False(t, Int64Slice(int64s).Contains(6))
}
