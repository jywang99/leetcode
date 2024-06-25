package medium

import com "jy.org/leetcode/src/common"

// 235. Lowest Common Ancestor of a Binary Search Tree
func lowestCommonAncestor(root, p, q *com.TreeNode) *com.TreeNode {
    var getPath func(*com.TreeNode, *com.TreeNode, *[]*com.TreeNode)
    getPath = func(n , t *com.TreeNode, path *[]*com.TreeNode) {
        *path = append(*path, n)
        if n == t {
            return
        }

        var nn *com.TreeNode
        if n.Val > t.Val {
            nn = n.Left
        } else {
            nn = n.Right
        }

        getPath(nn, t, path)
    }

    ppath := make([]*com.TreeNode, 0)
    getPath(root, p, &ppath)
    qpath := make([]*com.TreeNode, 0)
    getPath(root, q, &qpath)

    cp := ppath[0]
    for i:=1; i<min(len(ppath), len(qpath)); i++ {
        if ppath[i] != qpath[i] {
            break
        }
        cp = ppath[i]
    }

    return cp
}

// 102. Binary Tree Level Order Traversal
func LevelOrder(root *com.TreeNode) [][]int {
    rs := make([][]int, 0)
    if root == nil {
        return rs
    }

    lvl := make([]*com.TreeNode, 1)
    lvl[0] = root

    for len(lvl) > 0 {
        ilvl := make([]int, 0, len(lvl))
        nlvl := make([]*com.TreeNode, 0)
        for _, n := range lvl {
            ilvl = append(ilvl, n.Val)
            if n.Left != nil {
                nlvl = append(nlvl, n.Left)
            }
            if n.Right != nil {
                nlvl = append(nlvl, n.Right)
            }
        }

        rs = append(rs, ilvl)
        lvl = nlvl
    }

    return rs
}

