package medium_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	m "jy.org/leetcode/src/medium"
)

func TestSearchMatrix(t *testing.T) {
    mtx := [][]int{
        {1,3,5,7},
        {10,11,16,20},
        {23,30,34,60},
    }
    assert.Equal(t, true, m.SearchMatrix(mtx, 3))
    assert.Equal(t, false, m.SearchMatrix(mtx, 13))
}

