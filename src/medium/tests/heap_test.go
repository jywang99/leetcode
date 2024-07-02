package medium_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	m "jy.org/leetcode/src/medium"
)

func TestHeap(t *testing.T) {
    // min heap
    h := m.NewHeap([]int{5,6,1,3,2}, func(a, b int) int {
        if a == b {
            return 0
        }
        if a > b {
            return -1
        }
        return 1
    })
    assert.Equal(t, 1, h.PopTop())
    assert.Equal(t, 2, h.GetTop())
    assert.Equal(t, 2, h.PopTop())
    assert.Equal(t, 3, h.PopTop())
    h.Insert(0)
    assert.Equal(t, 0, h.PopTop())
    h.Insert(2)
    assert.Equal(t, 2, h.PopTop())
    assert.Equal(t, 5, h.PopTop())

    // max heap
    h1 := m.NewHeap([]int{5,6,1,3,2}, func(a, b int) int {
        if a == b {
            return 0
        }
        if a > b {
            return 1
        }
        return -1
    })
    assert.Equal(t, 6, h1.PopTop())
    assert.Equal(t, 5, h1.GetTop())
    assert.Equal(t, 5, h1.PopTop())
    assert.Equal(t, 3, h1.PopTop())
    h1.Insert(7)
    assert.Equal(t, 7, h1.PopTop())
    h1.Insert(4)
    assert.Equal(t, 4, h1.PopTop())
    assert.Equal(t, 2, h1.PopTop())
}

func TestKClosest(t *testing.T) {
    assert.ElementsMatch(t, [][]int{{3,3},{-2,4}}, m.KClosest([][]int{{3,3}, {5,-1}, {-2,4}}, 2))
    assert.ElementsMatch(t, [][]int{{-5,4},{4,6}}, m.KClosest([][]int{{-5,4}, {-6,-5}, {4,6}}, 2))
    assert.ElementsMatch(t, 
        [][]int{{0,1},{-2,0},{-3,2},{0,-5},{-5,4},{-4,-6}}, 
        m.KClosest([][]int{{-5,4},{-3,2},{0,1},{-3,7},{-2,0},{-4,-6},{0,-5}}, 6),
    )
}

