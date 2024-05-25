package medium_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	m "jy.org/leetcode/src/medium"
)

func TestGroupAnagrams(t *testing.T) {
    rs := m.GroupAnagrams([]string{"eat","tea","tan","ate","nat","bat"})
    assert.Contains(t, rs, []string{"eat", "tea", "ate"}, []string{"tan", "nat"}, []string{"bat"})
}

func TestTopKFrequent(t *testing.T) {
    rs := m.TopKFrequent([]int{1,1,1,2,2,3}, 2)
    assert.Equal(t, []int{1, 2}, rs)
}

func TestProductExceptSelf(t *testing.T) {
    assert.Equal(t, []int{24,12,8,6}, m.ProductExceptSelf([]int{1,2,3,4}))
    assert.Equal(t, []int{24,12,8,6}, m.ProductExceptSelfSpace([]int{1,2,3,4}))
}

func TestIsValidSudoku(t *testing.T) {
    board := [][]byte{
        []byte("53..7...."),
        []byte("6..195..."),
        []byte(".98....6."),
        []byte("8...6...3"),
        []byte("4..8.3..1"),
        []byte("7...2...6"),
        []byte(".6....28."),
        []byte("...419..5"),
        []byte("....8..79"),
    }
    assert.True(t, m.IsValidSudoku(board))
}

func TestTwoSum(t *testing.T) {
    assert.Equal(t, []int{1,2}, m.TwoSum([]int{2,7,11,15}, 9))
    assert.Equal(t, []int{1,3}, m.TwoSum([]int{2,3,4}, 6))
    assert.Equal(t, []int{1,2}, m.TwoSum([]int{-1,0}, -1))

    assert.Equal(t, []int{1,2}, m.TwoSumTwoPtr([]int{2,7,11,15}, 9))
    assert.Equal(t, []int{1,3}, m.TwoSumTwoPtr([]int{2,3,4}, 6))
    assert.Equal(t, []int{1,2}, m.TwoSumTwoPtr([]int{-1,0}, -1))
}

func TestThreeSum(t *testing.T) {
    rs := m.ThreeSum([]int{-1,0,1,2,-1,-4})
    assert.Equal(t, 2, len(rs))
    assert.Contains(t, rs, []int{-1,-1,2}, []int{-1,0,1})
    rs = m.ThreeSum([]int{0,0,0,0})
    assert.Equal(t, 1, len(rs))
    assert.Contains(t, rs, []int{0,0,0})
}

func TestMaxArea(t *testing.T) {
    assert.Equal(t, 49, m.MaxArea([]int{1,8,6,2,5,4,8,3,7}))
    assert.Equal(t, 1, m.MaxArea([]int{1,1}))
}

func TestTrap(t *testing.T) {
    assert.Equal(t, 6, m.Trap([]int{0,1,0,2,1,0,1,3,2,1,2,1}))
    assert.Equal(t, 9, m.Trap([]int{4,2,0,3,2,5}))
    assert.Equal(t, 1, m.Trap([]int{5,4,1,2}))
}
