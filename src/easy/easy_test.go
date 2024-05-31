package easy_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"jy.org/leetcode/src/easy"
)

func TestMap(t *testing.T) {
    m := make(map[int]int)
    m[1] = 2
    m[2] = 3
    assert.Equal(t, 2, len(m))
    delete(m, 2)
    assert.Equal(t, 1, len(m))
}

func TestValidPalindrome(t *testing.T) {
    assert.True(t, easy.IsPalindrome("A man, a plan, a canal: Panama"))
    assert.False(t, easy.IsPalindrome("0P"))
}

func TestValidParentheses(t *testing.T) {
    assert.True(t, easy.IsValid("()"))
    assert.True(t, easy.IsValid("()[]{}"))
    assert.False(t, easy.IsValid("(]"))
    assert.True(t, easy.IsValid("{[]}"))
    assert.False(t, easy.IsValid("(("))
    assert.False(t, easy.IsValid("){"))
}

func TestBSearch(t *testing.T) {
    arr := []int{-1,0,3,5,9,12}
    assert.Equal(t, 4, easy.Search(arr, 9))
    assert.Equal(t, -1, easy.Search(arr, 2))
    assert.Equal(t, -1, easy.Search(arr, 13))
    assert.Equal(t, 1, easy.Search([]int{2,5}, 5))

    assert.Equal(t, 4, easy.SearchLoop(arr, 9))
    assert.Equal(t, -1, easy.SearchLoop(arr, 2))
    assert.Equal(t, -1, easy.SearchLoop(arr, 13))
    assert.Equal(t, 1, easy.SearchLoop([]int{2,5}, 5))
}

func TestMaxProfit(t *testing.T) {
    assert.Equal(t, 5, easy.MaxProfit([]int{7,1,5,3,6,4}))
    assert.Equal(t, 0, easy.MaxProfit([]int{7,6,4,3,1}))

    assert.Equal(t, 5, easy.MaxProfitSliding([]int{7,1,5,3,6,4}))
    assert.Equal(t, 0, easy.MaxProfitSliding([]int{7,6,4,3,1}))
}

func TestReverseList(t *testing.T) {
    a := []int{1,2,3,4,5}
    ll := easy.SliceToLinkedList(a)
    assert.Equal(t, a, ll.ToSlice())

    llr := easy.ReverseList(ll)
    assert.Equal(t, []int{5,4,3,2,1}, llr.ToSlice())
    assert.Equal(t, []int{2,1}, easy.ReverseList(easy.SliceToLinkedList([]int{1,2})).ToSlice())
    assert.Equal(t, []int{}, easy.ReverseList(easy.SliceToLinkedList([]int{})).ToSlice())
}

func TestMergeTwoLists(t *testing.T) {
    // assert.Equal(t, []int{1,1,2,3,4,4}, easy.MergeTwoLists(easy.SliceToLinkedList([]int{1,2,4}), easy.SliceToLinkedList([]int{1,3,4})).ToSlice())
    // assert.Equal(t, []int{}, easy.MergeTwoLists(easy.SliceToLinkedList([]int{}), easy.SliceToLinkedList([]int{})).ToSlice())
    assert.Equal(t, []int{0}, easy.MergeTwoLists(easy.SliceToLinkedList([]int{}), easy.SliceToLinkedList([]int{0})).ToSlice())
}
