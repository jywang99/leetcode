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

func TestAddTwoNumbers(t *testing.T) {
    assert.Equal(t, []int{7,0,8}, m.AddTwoNumbers(com.SliceToLinkedList([]int{2,4,3}), com.SliceToLinkedList([]int{5,6,4})).ToSlice())
    assert.Equal(t, []int{0}, m.AddTwoNumbers(com.SliceToLinkedList([]int{0}), com.SliceToLinkedList([]int{0})).ToSlice())
    assert.Equal(t, []int{8,9,9,9,0,0,0,1}, m.AddTwoNumbers(com.SliceToLinkedList([]int{9,9,9,9,9,9,9}), com.SliceToLinkedList([]int{9,9,9,9})).ToSlice())
}

func TestFindDuplicate(t *testing.T) {
    assert.Equal(t, 2, m.FindDuplicate([]int{1,3,4,2,2}))
    assert.Equal(t, 3, m.FindDuplicate([]int{3,1,3,4,2}))
    assert.Equal(t, 3, m.FindDuplicate([]int{3,3,3,3,3}))

    assert.Equal(t, 2, m.FindDuplicateConstant([]int{1,3,4,2,2}))
    assert.Equal(t, 3, m.FindDuplicateConstant([]int{3,1,3,4,2}))
    assert.Equal(t, 3, m.FindDuplicateConstant([]int{3,3,3,3,3}))
}

func TestLruCache(t *testing.T) {
    ca := m.LruConstructor(2);
    ca.Put(1, 1) // cache is {1=1}
    ca.Put(2, 2) // cache is {1=1, 2=2}
    assert.Equal(t, 1, ca.Get(1))
    ca.Put(3, 3) // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
    assert.Equal(t, -1, ca.Get(2))
    ca.Put(4, 4) // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
    assert.Equal(t, -1, ca.Get(1))
    assert.Equal(t, 3, ca.Get(3))
    assert.Equal(t, 4, ca.Get(4))
}

