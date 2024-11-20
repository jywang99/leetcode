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

func TestMergeTriplets(t *testing.T) {
    assert.True(t, m.MergeTriplets([][]int{{2,5,3},{1,8,4},{1,7,5}}, []int{2,7,5}))
    assert.False(t, m.MergeTriplets([][]int{{3,4,5},{4,5,6}}, []int{3,2,5}))
    assert.True(t, m.MergeTriplets([][]int{{2,5,3},{2,3,4},{1,2,5},{5,2,3}}, []int{5,5,5}))
}

func TestPartitionLabels(t *testing.T) {
    assert.Equal(t, []int{9,7,8}, m.PartitionLabels("ababcbacadefegdehijhklij"))
    assert.Equal(t, []int{10}, m.PartitionLabels("eccbbbbdec"))
}

func TestCheckValidString(t *testing.T) {
    assert.True(t, m.CheckValidString("()"))
    assert.True(t, m.CheckValidString("(*)"))
    assert.True(t, m.CheckValidString("(*))"))
    assert.False(t, m.CheckValidString("(*)))"))
    assert.False(t, m.CheckValidString("())"))
    assert.False(t, m.CheckValidString(")"))
}

