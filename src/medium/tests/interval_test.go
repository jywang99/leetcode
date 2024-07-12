package medium_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	com "jy.org/leetcode/src/common"
	m "jy.org/leetcode/src/medium"
)

func TestInsert(t *testing.T) {
    assert.Equal(t, [][]int{{1,5}, {6,9}}, m.Insert([][]int{{1,3}, {6,9}}, []int{2,5}))
    assert.Equal(t, [][]int{{5,7}}, m.Insert([][]int{}, []int{5,7}))
}

func TestMerge(t *testing.T) {
    assert.Equal(t, [][]int{{1,6},{8,10},{15,18}}, m.Merge([][]int{{1,3},{2,6},{8,10},{15,18}}))
    assert.Equal(t, [][]int{{1,5}}, m.Merge([][]int{{1,4},{4,5}}))
}

func TestQuicksort(t *testing.T) {
    arr := []int{4,5,2,1,3,6}
    arr = com.Quicksort(arr, func(i1, i2 int) int {
        return i1 - i2
    })
    assert.Equal(t, []int{1,2,3,4,5,6}, arr)
}

func TestMinMeetingRooms(t *testing.T) {
    assert.Equal(t, 2, m.MinMeetingRooms([][]int{{0,40},{5,10},{15,20}}))
    assert.Equal(t, 1, m.MinMeetingRooms([][]int{{4,9}}))
}

