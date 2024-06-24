package easy_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	com "jy.org/leetcode/src/common"
	e "jy.org/leetcode/src/easy"
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
    assert.True(t, e.IsPalindrome("A man, a plan, a canal: Panama"))
    assert.False(t, e.IsPalindrome("0P"))
}

func TestValidParentheses(t *testing.T) {
    assert.True(t, e.IsValid("()"))
    assert.True(t, e.IsValid("()[]{}"))
    assert.False(t, e.IsValid("(]"))
    assert.True(t, e.IsValid("{[]}"))
    assert.False(t, e.IsValid("(("))
    assert.False(t, e.IsValid("){"))
}

func TestBSearch(t *testing.T) {
    arr := []int{-1,0,3,5,9,12}
    assert.Equal(t, 4, e.Search(arr, 9))
    assert.Equal(t, -1, e.Search(arr, 2))
    assert.Equal(t, -1, e.Search(arr, 13))
    assert.Equal(t, 1, e.Search([]int{2,5}, 5))

    assert.Equal(t, 4, e.SearchLoop(arr, 9))
    assert.Equal(t, -1, e.SearchLoop(arr, 2))
    assert.Equal(t, -1, e.SearchLoop(arr, 13))
    assert.Equal(t, 1, e.SearchLoop([]int{2,5}, 5))
}

func TestMaxProfit(t *testing.T) {
    assert.Equal(t, 5, e.MaxProfit([]int{7,1,5,3,6,4}))
    assert.Equal(t, 0, e.MaxProfit([]int{7,6,4,3,1}))

    assert.Equal(t, 5, e.MaxProfitSliding([]int{7,1,5,3,6,4}))
    assert.Equal(t, 0, e.MaxProfitSliding([]int{7,6,4,3,1}))
}

func TestReverseList(t *testing.T) {
    a := []int{1,2,3,4,5}
    ll := com.SliceToLinkedList(a)
    assert.Equal(t, a, ll.ToSlice())

    llr := e.ReverseList(ll)
    assert.Equal(t, []int{5,4,3,2,1}, llr.ToSlice())
    assert.Equal(t, []int{2,1}, e.ReverseList(com.SliceToLinkedList([]int{1,2})).ToSlice())
    assert.Equal(t, []int{}, e.ReverseList(com.SliceToLinkedList([]int{})).ToSlice())
}

func TestMergeTwoLists(t *testing.T) {
    assert.Equal(t, []int{1,1,2,3,4,4}, e.MergeTwoLists(com.SliceToLinkedList([]int{1,2,4}), com.SliceToLinkedList([]int{1,3,4})).ToSlice())
    assert.Equal(t, []int{}, e.MergeTwoLists(com.SliceToLinkedList([]int{}), com.SliceToLinkedList([]int{})).ToSlice())
    assert.Equal(t, []int{0}, e.MergeTwoLists(com.SliceToLinkedList([]int{}), com.SliceToLinkedList([]int{0})).ToSlice())
}

func TestHasCycle(t *testing.T) {
    assert.True(t, e.HasCycle(SliceToLinkedListCycle([]int{3,2,0,-4}, 1)))
    assert.True(t, e.HasCycle(SliceToLinkedListCycle([]int{1,2}, 0)))
    assert.False(t, e.HasCycle(com.SliceToLinkedList([]int{1})))

    assert.True(t, e.HasCycleCompact(SliceToLinkedListCycle([]int{3,2,0,-4}, 1)))
    assert.True(t, e.HasCycleCompact(SliceToLinkedListCycle([]int{1,2}, 0)))
    assert.False(t, e.HasCycleCompact(com.SliceToLinkedList([]int{1,2,3,4})))
    assert.False(t, e.HasCycleCompact(com.SliceToLinkedList([]int{1})))
}

func SliceToLinkedListCycle(arr []int, pos int) *com.ListNode {
    n := com.SliceToLinkedList(arr)
    h := n
    i := 0
    var ln *com.ListNode
    for n.Next != nil { // find loop node
        if i == pos {
            ln = n
        }
        n = n.Next
        i++
    }
    n.Next = ln // create loop from tail to loop node
    return h
}

func TestInvertTree(t *testing.T) {
    assert.Equal(t, []int{4,7,2,9,6,3,1}, e.InvertTree(com.SliceToTree([]int{4,2,7,1,3,6,9})).ToSlice())
    assert.Equal(t, []int{2,1,3}, e.InvertTree(com.SliceToTree([]int{2,3,1})).ToSlice())
    assert.Equal(t, []int{}, e.InvertTree(com.SliceToTree([]int{})).ToSlice())
}

func TestDiameterOfBinaryTree(t *testing.T) {
    n := &com.TreeNode{
        Val: 1,
        Right: &com.TreeNode{
            Val: 3,
        },
        Left: &com.TreeNode{
            Val: 2,
            Left: &com.TreeNode{
                Val: 4,
            },
            Right: &com.TreeNode{
                Val: 5,
            },
        },
    }
    assert.Equal(t, 3, e.DiameterOfBinaryTree(n))

    n = &com.TreeNode{
        Val: 1,
        Left: &com.TreeNode{
            Val: 2,
        },
    }
    assert.Equal(t, 1, e.DiameterOfBinaryTree(n))
}

func TestIsBalanced(t *testing.T) {
    n := &com.TreeNode{
        Val: 3,
        Right: &com.TreeNode{
            Val: 20,
            Left: &com.TreeNode{
                Val: 15,
            },
            Right: &com.TreeNode{
                Val: 7,
            },
        },
        Left: &com.TreeNode{
            Val: 9,
        },
    }
    assert.True(t, e.IsBalanced(n))

    n = &com.TreeNode{
        Val: 1,
        Left: &com.TreeNode{
            Val: 2,
            Left: &com.TreeNode{
                Val: 3,
                Left: &com.TreeNode{
                    Val: 4,
                },
                Right: &com.TreeNode{
                    Val: 4,
                },
            },
            Right: &com.TreeNode{
                Val: 3,
            },
        },
        Right: &com.TreeNode{
            Val: 2,
        },
    }
    assert.False(t, e.IsBalanced(n))

    n = nil
    assert.True(t, e.IsBalanced(n))

    n = &com.TreeNode{
        Val: 1,
        Right: &com.TreeNode{
            Val: 2,
            Right: &com.TreeNode{
                Val: 3,
            },
        },
    }
    assert.False(t, e.IsBalanced(n))
}
