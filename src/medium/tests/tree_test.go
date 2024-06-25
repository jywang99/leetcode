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

