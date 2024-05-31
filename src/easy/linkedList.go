package easy

import com "jy.org/leetcode/src/common"

// 206. Reverse Linked List
func ReverseList(head *com.ListNode) *com.ListNode {
    if head == nil {
        return nil
    }
    p := head
    if p.Next == nil {
        return head
    }
    n := p.Next
    p.Next = nil
    var nn *com.ListNode
    for n != nil {
        nn = n.Next
        n.Next = p
        p = n
        n = nn
    }
    return p
}

// 21. Merge Two Sorted Lists
func MergeTwoLists(list1 *com.ListNode, list2 *com.ListNode) *com.ListNode {
    n1 := list1
    n2 := list2

    if n1 == nil {
        return n2
    }
    if n2 == nil {
        return n1
    }

    getNext := func() *com.ListNode {
        // get min node among the two linked lists
        // increment ptr for the ll selected
        var n com.ListNode
        if n1.Val < n2.Val {
            n = *n1
            n1 = n1.Next
        } else {
            n = *n2
            n2 = n2.Next
        }
        return &n
    }

    // first node
    n := getNext()
    head := n
    for n1 != nil && n2 != nil {
        n.Next = getNext()
        n = n.Next
    }

    // remaining nodes
    if n1 != nil {
        n.Next = n1
    } 
    if n2 != nil {
        n.Next = n2
    }

    return head
}
