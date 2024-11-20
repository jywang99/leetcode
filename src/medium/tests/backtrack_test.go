package medium

import (
	"testing"

	"github.com/stretchr/testify/assert"
	m "jy.org/leetcode/src/medium"
)

func TestSubsets(t *testing.T) {
    assert.ElementsMatch(t, [][]int{{},{1},{2},{1,2},{3},{1,3},{2,3},{1,2,3}}, m.Subsets([]int{1,2,3}))
    assert.ElementsMatch(t, [][]int{{},{0}}, m.Subsets([]int{0}))
}

