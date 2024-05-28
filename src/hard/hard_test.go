package hard_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	h "jy.org/leetcode/src/hard"
)

func TestLargestRectangleArea(t *testing.T) {
    assert.Equal(t, 10, h.LargestRectangleArea([]int{2,1,5,6,2,3}))
    assert.Equal(t, 4, h.LargestRectangleArea([]int{2,4}))
}

func TestTrap(t *testing.T) {
    assert.Equal(t, 6, h.Trap([]int{0,1,0,2,1,0,1,3,2,1,2,1}))
    assert.Equal(t, 9, h.Trap([]int{4,2,0,3,2,5}))
    assert.Equal(t, 1, h.Trap([]int{5,4,1,2}))
}

