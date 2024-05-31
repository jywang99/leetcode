package medium

import (
	com "jy.org/leetcode/src/common"
)

// 143. Reorder List
func ReorderList(head *com.ListNode) {
    if head == nil || head.Next == nil {
        return
    }

    // find middle node
    s, ps := head, head
    f := head
    for f != nil {
        ps = s
        s = s.Next
        f = f.Next
        if f != nil {
            f = f.Next
        }
    }
    ps.Next = nil

    // reverse links for the nodes to be wrapped around
    p := s
    n := p.Next
    p.Next = nil
    var nn *com.ListNode
    for n != nil {
        nn = n.Next
        n.Next = p
        p = n
        n = nn
    }

    // merge two lists
    fn := head // start from head
    var nfn *com.ListNode // next forward node
    rn := p
    var nrn *com.ListNode // next backward node
    // fmt.Println(fn.ToSlice())
    // fmt.Println(rn.ToSlice())
    for rn != nil {
        // link: fn -> rn -> nfn
        nfn = fn.Next
        nrn = rn.Next
        fn.Next = rn
        rn.Next = nfn
        // update next nodes
        fn = nfn
        rn = nrn
    }
}

// 19. Remove Nth Node From End of List
func RemoveNthFromEnd(head *com.ListNode, n int) *com.ListNode {
    if head.Next == nil {
        return nil
    }

    // find last n+1 nodes
    dq := com.NewDequeue(n+1)
    for n:=head; n!=nil; n=n.Next {
        if dq.IsFull() {
            dq.PopHead()
        }
        dq.AppendTail(n)
    }

    // remove head
    if dq.Size() == n {
        return head.Next
    }

    prev := (*dq.PopHead()).(*com.ListNode)
    dq.PopHead() // node to discard
    np := dq.PopHead() // next node, can be nil
    var next *com.ListNode
    if np != nil {
        next = (*np).(*com.ListNode)
    }
    prev.Next = next

    return head
}

