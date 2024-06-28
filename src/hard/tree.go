package hard

import (
	"fmt"
	"strconv"
	"strings"

	com "jy.org/leetcode/src/common"
)

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

// 297. Serialize and Deserialize Binary Tree
type Codec struct {
    Sep byte
    Nil byte
}

func CodecConstructor() Codec {
    return Codec{Sep: ',', Nil: '#'}
}

// Serializes a tree to a single string.
func (this *Codec) Serialize(root *com.TreeNode) string {
    rs := ""
    if root == nil {
        return rs
    }

    lvl := make([]*com.TreeNode, 1)
    lvl[0] = root

    for len(lvl) > 0 {
        nlvl := make([]*com.TreeNode, 0)
        for _, n := range lvl {
            var s string
            if n != nil {
                s = fmt.Sprintf("%d", n.Val)
                nlvl = append(nlvl, n.Left)
                nlvl = append(nlvl, n.Right)
            } else {
                s = fmt.Sprintf("%c", this.Nil)
            }
            rs += fmt.Sprintf("%s%c", s, this.Sep)
        }
        lvl = nlvl
    }

    var lni int
    for i:=len(rs)-1; i>=0; i-- {
        b := rs[i]
        if b >= '0' && b <= '9' {
            lni = i
            break
        }
    }

    return rs[:lni+1]
}

// Deserializes your encoded data to tree.
func (this *Codec) Deserialize(data string) *com.TreeNode {
    if len(data) == 0 {
        return nil
    }

    sarr := strings.Split(data, fmt.Sprintf("%c", this.Sep))
    varr := make([]*int, len(sarr))
    for i, v := range sarr {
        iv, err := strconv.Atoi(v)
        if err == nil {
            varr[i] = &iv
        } else {
            varr[i] = nil
        }
    }

    root := &com.TreeNode{
        Val: *varr[0],
    }
    queue := []*com.TreeNode{root}
    i := 1
    for i < len(varr) {
        n := queue[0]
        queue = queue[1:]

        if i < len(varr) && varr[i] != nil {
            n.Left = &com.TreeNode{
                Val: *varr[i],
            }
            queue = append(queue, n.Left)
        }
        i++
        if i < len(varr) && varr[i] != nil {
            n.Right = &com.TreeNode{
                Val: *varr[i],
            }
            queue = append(queue, n.Right)
        }
        i++
    }

    return root
}

