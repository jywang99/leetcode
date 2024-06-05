package easy

import com "jy.org/leetcode/src/common"

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
    return 0
}
