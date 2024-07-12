package medium

import (
	"testing"

	"github.com/stretchr/testify/assert"
	m "jy.org/leetcode/src/medium"
)

func TestIsNStraightHand(t *testing.T) {
    assert.Equal(t, true, m.IsNStraightHand([]int{1,2,3,6,2,3,4,7,8}, 3))
    assert.Equal(t, false, m.IsNStraightHand([]int{1,2,3,4,5}, 4))
    assert.Equal(t, false, m.IsNStraightHand([]int{1,1,2,2,3,3}, 2))
}
