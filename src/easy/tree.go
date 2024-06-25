package easy

import (
	"math"

	com "jy.org/leetcode/src/common"
)

// 226. Invert Binary Tree
func InvertTree(root *com.TreeNode) *com.TreeNode {
    if root == nil {
        return nil
    }
    tmp := root.Left
    root.Left = root.Right
    root.Right = tmp
    if root.Left != nil {
        InvertTree(root.Left)
    }
    if root.Right != nil {
        InvertTree(root.Right)
    }
    return root
}

// 104. Maximum Depth of Binary Tree
func MaxDepth(root *com.TreeNode) int {
    var maxd func(*com.TreeNode) int
    maxd = func(n *com.TreeNode) int {
        if n == nil {
            return 0
        }
        return max(maxd(n.Left), maxd(n.Right)) + 1
    }
    return maxd(root)
}

// 543. Diameter of Binary Tree
func DiameterOfBinaryTree(root *com.TreeNode) int {
    md := 0
    var maxd func(*com.TreeNode) int
    maxd = func(n *com.TreeNode) int {
        if n == nil {
            return 0
        }
        ld := maxd(n.Left)
        rd := maxd(n.Right)
        d := ld + rd
        if md < d {
            md = d
        }
        return max(ld, rd) + 1
    }
    maxd(root)
    return md
}

// 110. Balanced Binary Tree
func IsBalanced(root *com.TreeNode) bool {
    var balRec func(*com.TreeNode) (bool, int)
    balRec = func(n *com.TreeNode) (bool, int) {
        if n == nil {
            return true, 0
        }

        lb, lh := balRec(n.Left)
        rb, rh := balRec(n.Right)

        if !lb || !rb {
            return false, 0
        }

        diff := int(math.Abs(float64(lh) - float64(rh)))
        if diff > 1 {
            return false, 0
        }

        return true, max(lh, rh) + 1
    }
    b, _ := balRec(root)
    return b
}

// 100. Same Tree
func isSameTree(p *com.TreeNode, q *com.TreeNode) bool {
    if p == nil && q != nil || q == nil && p != nil {
        return false
    }

    if p == nil && q == nil {
        return true
    }

    return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right) && p.Val == q.Val
}

// 572. Subtree of Another Tree
func isSubtree(root *com.TreeNode, subRoot *com.TreeNode) bool {
    if subRoot == nil {
        return true
    }

    var isSameTree func (*com.TreeNode, *com.TreeNode) bool
    isSameTree = func(p, q *com.TreeNode) bool {
        if p == nil && q != nil || q == nil && p != nil {
            return false
        }

        if p == nil && q == nil {
            return true
        }

        return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right) && p.Val == q.Val

    }

    var subRec func (*com.TreeNode) bool
    subRec = func(rn *com.TreeNode) bool {
        if rn == nil {
            return false
        }
        if isSameTree(rn, subRoot) {
            return true
        }
        return subRec(rn.Left) || subRec(rn.Right)
    }

    return subRec(root)
}

