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

// 138. Copy List with Random Pointer
type Node struct {
    Val int
    Next *Node
    Random *Node
}

func CopyRandomList(head *Node) *Node {
    pmap := make(map[*Node]*Node)
    for n:=head; n!=nil; n=n.Next {
        pmap[n] = &Node{
            Val: n.Val,
        }
    }

    getCloneN := func(on *Node) *Node {
        if on == nil {
            return nil
        }
        return pmap[on]
    }

    var ncur *Node
    var nnext *Node
    var nrand *Node
    for n:=head; n!=nil; n=n.Next {
        ncur = getCloneN(n)
        nnext = getCloneN(n.Next)
        nrand = getCloneN(n.Random)
        ncur.Next = nnext
        ncur.Random = nrand
    }

    return pmap[head]
}

// 2. Add Two Numbers
func AddTwoNumbers(l1 *com.ListNode, l2 *com.ListNode) *com.ListNode {
    carry := false
    getNextN := func() *com.ListNode {
        var n *com.ListNode
        ncarry := false // carry for next digit
        if l1 == nil && l2 == nil {
            if carry {
                n = &com.ListNode{
                    Val: 0,
                }
            }
            // else nil
        } else if l1 == nil {
            n = l2
        } else if l2 == nil {
            n = l1
        } else {
            sum := l1.Val + l2.Val
            n = &com.ListNode{
                Val: sum % 10,
            }
            ncarry = sum > 9
        }

        // apply carry from pervious digit
        if carry {
            if n.Val < 9 {
                n.Val ++
            } else {
                n.Val = 0
                ncarry = true
            }
        }
        carry = ncarry

        // move pointers to next digit
        if l1 != nil {
            l1 = l1.Next
        }
        if l2 != nil {
            l2 = l2.Next
        }

        return n
    }

    dummy, n := l1, l1
    for n != nil {
        n.Next = getNextN()
        n = n.Next
    }

    return dummy.Next
}

// 287. Find the Duplicate Number
func FindDuplicate(nums []int) int {
    m := make(map[int]bool)
    for _, n := range nums {
        if m[n] {
            return n
        }
        m[n] = true
    }
    return -1
}

// constant space
func FindDuplicateConstant(nums []int) int {
    if len(nums) == 2 {
        return nums[0]
    }
    s, f := 0, 0
    for {
        s = nums[s]
        f = nums[f]
        f = nums[f]
        if s == f {
            break
        }
    }

    s2 := 0
    for {
        s = nums[s]
        s2 = nums[s2]
        if s == s2 {
            return s
        }
    }
}

type DlNode struct {
    Key int
    Data any
    Prev *DlNode
    Next *DlNode
}

type LRUCache struct {
    Data map[int]*DlNode
    Cap int
    Lru *DlNode
    Mru *DlNode
}

func LruConstructor(capacity int) LRUCache {
    lfu := &DlNode{}
    mfu := &DlNode{
        Prev: lfu,
    }
    lfu.Next = mfu

    return LRUCache{
        Data: make(map[int]*DlNode),
        Cap: capacity,
        Lru: lfu,
        Mru: mfu,
    }
}

func (this *LRUCache) Get(key int) int {
    n := this.Data[key]
    if n == nil {
        return -1
    }
    this.moveToMru(n)
    return n.Data.(int)
}

func (this *LRUCache) moveToMru(n *DlNode)  {
    // take out n
    op := n.Prev
    on := n.Next
    op.Next = on
    on.Prev = op
    // put it to MRU
    op = this.Mru.Prev
    op.Next = n
    n.Prev = op
    this.Mru.Prev = n
    n.Next = this.Mru
}

func (this *LRUCache) Put(key int, value int)  {
    n := this.Data[key]
    if n != nil {
        // key exists, update
        n.Data = value
        this.moveToMru(n)
        return
    }

    // key doesn't exist, new node
    n = &DlNode{
        Key: key,
        Data: value,
    }

    // if full, remove LRU
    if len(this.Data) >= this.Cap {
        ol := this.Lru.Next
        this.Lru.Next = ol.Next
        ol.Next.Prev = this.Lru
        delete(this.Data, ol.Key)
    }

    // original MRU
    om := this.Mru.Prev
    // n <-> MRU
    this.Mru.Prev = n
    n.Next = this.Mru
    // om <-> n
    om.Next = n
    n.Prev = om

    this.Data[key] = n
}

