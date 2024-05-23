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

