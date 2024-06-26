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

// 199. Binary Tree Right Side View
func RightSideView(root *com.TreeNode) []int {
    rs := make([]int, 0)
    if root == nil {
        return rs
    }

    lvl := make([]*com.TreeNode, 1)
    lvl[0] = root

    for len(lvl) > 0 {
        rs = append(rs, lvl[len(lvl)-1].Val)
        nlvl := make([]*com.TreeNode, 0)
        for _, n := range lvl {
            if n.Left != nil {
                nlvl = append(nlvl, n.Left)
            }
            if n.Right != nil {
                nlvl = append(nlvl, n.Right)
            }
        }
        lvl = nlvl
    }

    return rs
}

// 1448. Count Good Nodes in Binary Tree
func goodNodes(root *com.TreeNode) int {
    cnt := 0
    var goodRec func(*com.TreeNode, int)
    goodRec = func(n *com.TreeNode, m int) {
        if n == nil {
            return
        }
        if n.Val >= m {
            cnt ++
            m = n.Val
        }
        goodRec(n.Left, m)
        goodRec(n.Right, m)
    }
    goodRec(root, root.Val)
    return cnt
}

// 98. Validate Binary Search Tree
func IsValidBST(root *com.TreeNode) bool {
    var validRec func(n *com.TreeNode) (bool, *int, *int)
    validRec = func(n *com.TreeNode) (bool, *int, *int) {
        if n == nil {
            return true, nil, nil
        }

        lv, lmax, lmin := validRec(n.Left)
        if !lv || lmax != nil && *lmax >= n.Val {
            return false, nil, nil
        }
        rv, rmax, rmin := validRec(n.Right)
        if !rv ||  rmin != nil &&*rmin <= n.Val {
            return false, nil, nil
        }

        if lmin == nil {
            lmin = &n.Val
        }
        if rmax == nil {
            rmax = &n.Val
        }

        return true, rmax, lmin
    }
    rs, _, _ := validRec(root)
    return rs
}
