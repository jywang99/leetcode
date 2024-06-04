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
