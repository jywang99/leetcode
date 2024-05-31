package medium_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	com "jy.org/leetcode/src/common"
	m "jy.org/leetcode/src/medium"
)

func TestReorderList(t *testing.T) {
    ol := com.SliceToLinkedList([]int{1,2,3,4})
    m.ReorderList(ol)
    assert.Equal(t, []int{1,4,2,3}, ol.ToSlice())

    ol2 := com.SliceToLinkedList([]int{1,2,3,4,5})
    m.ReorderList(ol2)
    assert.Equal(t, []int{1,5,2,4,3}, ol2.ToSlice())

    ol3 := com.SliceToLinkedList([]int{1})
    m.ReorderList(ol3)
    assert.Equal(t, []int{1}, ol3.ToSlice())
}

func TestRemoveNthFromEnd(t *testing.T) {
    assert.Equal(t, []int{1,2,3,5}, m.RemoveNthFromEnd(com.SliceToLinkedList([]int{1,2,3,4,5}), 2).ToSlice())
    assert.Equal(t, []int{}, m.RemoveNthFromEnd(com.SliceToLinkedList([]int{1}), 1).ToSlice())
    assert.Equal(t, []int{1}, m.RemoveNthFromEnd(com.SliceToLinkedList([]int{1,2}), 1).ToSlice())
    assert.Equal(t, []int{2}, m.RemoveNthFromEnd(com.SliceToLinkedList([]int{1,2}), 2).ToSlice())
}
