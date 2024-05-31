package common


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

func (head *ListNode) ToSlice() []int {
    rs := make([]int, 0)
    for n:=head; n!=nil; n=n.Next {
        rs = append(rs, n.Val)
    }
    return rs
}

