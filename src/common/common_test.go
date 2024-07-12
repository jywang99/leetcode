package common_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	com "jy.org/leetcode/src/common"
)

func getSample() (*com.TreeNode, []int) {
    return &com.TreeNode{
        Val: 1,
        Left: &com.TreeNode{
            Val: 2,
            Left: &com.TreeNode{
                Val: 4,
            },
            Right: &com.TreeNode{
                Val: 5,
            },
        },
        Right: &com.TreeNode{
            Val: 3,
            Left: &com.TreeNode{
                Val: 6,
            },
            Right: &com.TreeNode{
                Val: 7,
            },
        },
    }, []int{1,2,3,4,5,6,7}
}

func TestTreeSlice(t *testing.T) {
    head, arr := getSample()
    assert.Equal(t, arr, head.ToSlice())
}

func TestSliceToTree(t *testing.T) {
    head, arr := getSample()
    assert.Equal(t, head, com.SliceToTree(arr))
}

func TestHeap(t *testing.T) {
    // min heap
    h := com.NewHeap([]int{5,6,1,3,2}, func(a, b int) int {
        return a - b
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
    h1 := com.NewHeap([]int{5,6,1,3,2}, func(a, b int) int {
        return b - a
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

