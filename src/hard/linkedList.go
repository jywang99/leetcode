package hard

import (
	com "jy.org/leetcode/src/common"
)

// 23. Merge k Sorted Lists
func MergeKLists(lists []*com.ListNode) *com.ListNode {
    if len(lists) == 0 {
        return nil
    }
    if len(lists) == 1 {
        return lists[0]
    }

    rem := make([]*com.ListNode, 0, len(lists)-1)
    for _, l := range lists[2:] {
        rem = append(rem, l)
    }
    nl := mergeTwoLists(lists[0], lists[1])
    rem = append(rem, nl)
    return MergeKLists(rem)
}

func mergeTwoLists(l1 *com.ListNode, l2 *com.ListNode) *com.ListNode {
    getNextN := func() *com.ListNode {
        if l1 == nil && l2 == nil {
            return nil
        }

        var n *com.ListNode
        if l1 == nil {
            n = l2
        } else if l2 == nil {
            n = l1
        } else if l2.Val < l1.Val  {
            n = l2
        } else {
            n = l1
        }

        if n == l1 {
            l1 = l1.Next
        }
        if n == l2 {
            l2 = l2.Next
        }

        return n
    }
    n := getNextN()
    head := n
    for n != nil {
        n.Next = getNextN()
        n = n.Next
    }
    return head
}

func MergeKListsFast(lists []*com.ListNode) *com.ListNode {
    n := len(lists)
    if n == 0 {
        return nil
    }
    if n == 1 {
        return lists[0]
    }

    for i:=0; i<n-1; i++ {
        lists[i+1] = mergeTwoListsFast(lists[i], lists[i+1])
    }
    return lists[n-1]
}

func mergeTwoListsFast(l1 *com.ListNode, l2 *com.ListNode) *com.ListNode {
    dummy := &com.ListNode{}

    n := dummy
    for l1 != nil && l2 != nil {
        if l1.Val < l2.Val {
            n.Next = l1
            l1 = l1.Next
        } else {
            n.Next = l2
            l2 = l2.Next
        }
        n = n.Next
    }

    if l1 != nil {
        n.Next = l1
    } else {
        n.Next = l2
    }

    return dummy.Next
}

func ReverseKGroup(head *com.ListNode, k int) *com.ListNode {
    dummy := &com.ListNode{
        Next: head,
    }
    ptail := dummy
    nhead := head

    for nhead != nil {
        // get to kth node
        kth := nhead
        for i:=0; i<k-1; i++ {
            if kth.Next == nil {
                return dummy.Next
            }
            kth = kth.Next
        }
        // update next head
        nhead = kth.Next

        // reverse range
        pn := ptail.Next
        ntail := pn // first node in range will be the new tail for the range
        nn := pn.Next
        pn.Next = nhead
        for tmp := new(com.ListNode); nn != nhead; nn = tmp {
            tmp = nn.Next
            nn.Next = pn
            pn = nn
        }
        ptail.Next = pn // link with previous segment

        // update previous tail
        ptail = ntail
    }

    return dummy.Next
}
