package common

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func SliceToTree(src []int) *TreeNode {
    if len(src) == 0 {
        return nil
    }

    var toTree func(i int) *TreeNode
    toTree = func(i int) *TreeNode{
        if i >= len(src) {
            return nil
        }
        n := &TreeNode{
            Val: src[i],
            Left: toTree(i*2+1),
            Right: toTree(i*2+2),
        }
        return n
    }
    head := toTree(0)

    return head
}

func (t *TreeNode) ToSlice() []int {
    if t == nil {
        return []int{}
    }
    rs := make([]int, 0)
    lvl := make([]*TreeNode, 1)
    lvl[0] = t
    for len(lvl) > 0 {
        nlvl := make([]*TreeNode, 0)
        for _, n := range lvl {
            rs = append(rs, n.Val)
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

