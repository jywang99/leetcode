package medium_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"jy.org/leetcode/src/medium"
)

func TestGroupAnagrams(t *testing.T) {
    rs := medium.GroupAnagrams([]string{"eat","tea","tan","ate","nat","bat"})
    assert.Contains(t, rs, []string{"eat", "tea", "ate"}, []string{"tan", "nat"}, []string{"bat"})
}

func TestTopKFrequent(t *testing.T) {
    rs := medium.TopKFrequent([]int{1,1,1,2,2,3}, 2)
    assert.Equal(t, []int{1, 2}, rs)
}

func TestProductExceptSelf(t *testing.T) {
    assert.Equal(t, []int{24,12,8,6}, medium.ProductExceptSelf([]int{1,2,3,4}))
    assert.Equal(t, []int{24,12,8,6}, medium.ProductExceptSelfSpace([]int{1,2,3,4}))
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
    assert.True(t, medium.IsValidSudoku(board))
}

func TestTwoSum(t *testing.T) {
    assert.Equal(t, []int{1,2}, medium.TwoSum([]int{2,7,11,15}, 9))
    assert.Equal(t, []int{1,3}, medium.TwoSum([]int{2,3,4}, 6))
    assert.Equal(t, []int{1,2}, medium.TwoSum([]int{-1,0}, -1))

    assert.Equal(t, []int{1,2}, medium.TwoSumTwoPtr([]int{2,7,11,15}, 9))
    assert.Equal(t, []int{1,3}, medium.TwoSumTwoPtr([]int{2,3,4}, 6))
    assert.Equal(t, []int{1,2}, medium.TwoSumTwoPtr([]int{-1,0}, -1))
}

func TestThreeSum(t *testing.T) {
    rs := medium.ThreeSum([]int{-1,0,1,2,-1,-4})
    assert.Equal(t, 2, len(rs))
    assert.Contains(t, rs, []int{-1,-1,2}, []int{-1,0,1})
    rs = medium.ThreeSum([]int{0,0,0,0})
    assert.Equal(t, 1, len(rs))
    assert.Contains(t, rs, []int{0,0,0})
}
