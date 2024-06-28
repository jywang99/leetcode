package medium_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	com "jy.org/leetcode/src/common"
	m "jy.org/leetcode/src/medium"
)

func TestLevelOrder(t *testing.T) {
    n := &com.TreeNode{
        Val: 3,
        Left: &com.TreeNode{
            Val: 9,
        },
        Right: &com.TreeNode{
            Val: 20,
            Left: &com.TreeNode{
                Val: 15,
            },
            Right: &com.TreeNode{
                Val: 7,
            },
        },
    }
    exp := [][]int{
        {3},
        {9,20},
        {15,7},
    }
    assert.Equal(t, exp, m.LevelOrder(n))

    n2 := &com.TreeNode{
        Val: 1,
    }
    exp2 := [][]int{
        {1},
    }
    assert.Equal(t, exp2, m.LevelOrder(n2))

    var n3 *com.TreeNode
    exp3 := [][]int{}
    assert.Equal(t, exp3, m.LevelOrder(n3))
}

func TestRightSideView(t *testing.T) {
    n := &com.TreeNode{
        Val: 1,
        Left: &com.TreeNode{
            Val: 2,
            Right: &com.TreeNode{
                Val: 5,
            },
        },
        Right: &com.TreeNode{
            Val: 3,
            Right: &com.TreeNode{
                Val: 4,
            },
        },
    }
    assert.Equal(t, []int{1,3,4}, m.RightSideView(n))

    n1 := &com.TreeNode{
        Val: 1,
        Right: &com.TreeNode{
            Val: 3,
        },
    }
    assert.Equal(t, []int{1,3}, m.RightSideView(n1))

    var n2 *com.TreeNode
    assert.Equal(t, []int{}, m.RightSideView(n2))
}

func TestIsValidBST(t *testing.T) {
    n := &com.TreeNode{
        Val: 2,
        Left: &com.TreeNode{
            Val: 2,
        },
        Right: &com.TreeNode{
            Val: 2,
        },
    }
    assert.False(t, m.IsValidBST(n))

    n1 := &com.TreeNode{
        Val: 5,
        Left: &com.TreeNode{
            Val: 4,
        },
        Right: &com.TreeNode{
            Val: 6,
            Left: &com.TreeNode{
                Val: 3,
            },
            Right: &com.TreeNode{
                Val: 7,
            },
        },
    }
    assert.False(t, m.IsValidBST(n1))

    n2 := &com.TreeNode{
        Val: 1,
        Left: &com.TreeNode{
            Val: 1,
        },
    }
    assert.False(t, m.IsValidBST(n2))
}

func TestBuildTree(t *testing.T) {
    r := &com.TreeNode{
        Val: 3,
        Left: &com.TreeNode{
            Val: 9,
        },
        Right: &com.TreeNode{
            Val: 20,
            Left: &com.TreeNode{
                Val: 15,
            },
            Right: &com.TreeNode{
                Val: 7,
            },
        },
    }
    assert.Equal(t, r, m.BuildTree([]int{3,9,20,15,7}, []int{9,3,15,20,7}))
}
