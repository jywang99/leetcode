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

func TestMinEatingSpeed(t *testing.T) {
    assert.Equal(t, 4, m.MinEatingSpeed([]int{3,6,7,11}, 8))
    assert.Equal(t, 30, m.MinEatingSpeed([]int{30,11,23,4,20}, 5))
    assert.Equal(t, 23, m.MinEatingSpeed([]int{30,11,23,4,20}, 6))
}

func TestFindMin(t *testing.T) {
    assert.Equal(t, 1, m.FindMin([]int{3,4,5,1,2}))
    assert.Equal(t, 0, m.FindMin([]int{4,5,6,7,0,1,2}))
    assert.Equal(t, 11, m.FindMin([]int{11,13,15,17}))
    assert.Equal(t, 1, m.FindMin([]int{2,1}))

    assert.Equal(t, 1, m.FindMinFast([]int{3,4,5,1,2}))
    assert.Equal(t, 0, m.FindMinFast([]int{4,5,6,7,0,1,2}))
    assert.Equal(t, 11, m.FindMinFast([]int{11,13,15,17}))
    assert.Equal(t, 1, m.FindMinFast([]int{2,1}))
}

func TestSearchRotated(t *testing.T) {
    assert.Equal(t, 4, m.Search([]int{4,5,6,7,0,1,2}, 0))
    assert.Equal(t, -1, m.Search([]int{4,5,6,7,0,1,2}, 3))
    assert.Equal(t, -1, m.Search([]int{1}, 0))
    assert.Equal(t, 4, m.Search([]int{4,5,6,7,8,1,2,3}, 8))
    assert.Equal(t, 1, m.Search([]int{3,1}, 1))
}

func TestTimeMap(t *testing.T) {
    timeMap := m.TMapConstructor()
    timeMap.Set("foo", "bar", 1)
    assert.Equal(t, "bar", timeMap.Get("foo", 1))
    assert.Equal(t, "bar", timeMap.Get("foo", 3))
    timeMap.Set("foo", "bar2", 4)
    assert.Equal(t, "bar2", timeMap.Get("foo", 4))
    assert.Equal(t, "bar2", timeMap.Get("foo", 5))

    timeMap2 := m.TMapConstructor()
    timeMap2.Set("love", "high", 10)
    timeMap2.Set("love", "low", 20)
    assert.Equal(t, "", timeMap2.Get("love", 5))
    assert.Equal(t, "high", timeMap2.Get("love", 10))
    assert.Equal(t, "high", timeMap2.Get("love", 15))
    assert.Equal(t, "low", timeMap2.Get("love", 20))
    assert.Equal(t, "low", timeMap2.Get("love", 25))
}
