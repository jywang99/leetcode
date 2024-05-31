package easy

// 206. Reverse Linked List
type ListNode struct {
    Val int
    Next *ListNode
}

func SliceToLinkedList(arr []int) *ListNode {
    if len(arr) == 0 {
        return nil
    }
    head := ListNode{
        Val: arr[0],
        Next: nil,
    }
    if len(arr) == 1 {
        return &head
    }
    n := &head
    for _, e := range arr[1:] {
        n.Next = &ListNode{
            Val: e,
            Next: nil,
        }
        n = n.Next
    }
    return &head
}

func ReverseList(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    p := head
    if p.Next == nil {
        return head
    }
    n := p.Next
    p.Next = nil
    var nn *ListNode
    for n != nil {
        nn = n.Next
        n.Next = p
        p = n
        n = nn
    }
    return p
}

func (head *ListNode) ToSlice() []int {
    rs := make([]int, 0)
    for n:=head; n!=nil; n=n.Next {
        rs = append(rs, n.Val)
    }
    return rs
}

// 21. Merge Two Sorted Lists
func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    n1 := list1
    n2 := list2

    getNext := func() *ListNode { // TODO handle n1 or n2 or both nil cases
        // get min node among the two linked lists
        // increment ptr for the ll selected
        var n ListNode
        if n2 == nil || n1.Val < n2.Val {
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
