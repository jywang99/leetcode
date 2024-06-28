package hard

import com "jy.org/leetcode/src/common"

// 124. Binary Tree Maximum Path Sum
func MaxPathSum(root *com.TreeNode) int {
    var mv *int
    var maxRec func(n *com.TreeNode) (int)
    maxRec = func(n *com.TreeNode) (int) {
        if n == nil {
            return 0
        }

        // only this node
        if mv == nil || n.Val > *mv {
            mv = &n.Val
        }

        lns := maxRec(n.Left)
        rns := maxRec(n.Right)

        // split on this level
        sn := lns + n.Val + rns
        if sn > *mv {
            mv = &sn
        }

        // parent and left, parent and right
        lsum := n.Val + lns
        rsum := n.Val + rns
        if lsum > *mv {
            mv = &lsum
        }
        if rsum > *mv {
            mv = &rsum
        }

        // max path no split
        return max(lsum, n.Val, rsum)
    }
    maxRec(root)
    return *mv
}
