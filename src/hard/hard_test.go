package hard_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	h "jy.org/leetcode/src/hard"
)

func TestLargestRectangleArea(t *testing.T) {
    assert.Equal(t, 10, h.LargestRectangleArea([]int{2,1,5,6,2,3}))
    assert.Equal(t, 4, h.LargestRectangleArea([]int{2,4}))
}

func TestTrap(t *testing.T) {
    assert.Equal(t, 6, h.Trap([]int{0,1,0,2,1,0,1,3,2,1,2,1}))
    assert.Equal(t, 9, h.Trap([]int{4,2,0,3,2,5}))
    assert.Equal(t, 1, h.Trap([]int{5,4,1,2}))
}

func TestDequeue(t *testing.T) {
    // construct
    dq := h.NewDequeue(5)
    dq.AppendTail(3)
    dq.InsertHead(2)
    dq.AppendTail(4)
    dq.InsertHead(1)
    dq.AppendTail(5)

    // get
    assert.Equal(t, 5, dq.GetSize())
    assert.Equal(t, 1, *dq.GetHead())
    assert.Equal(t, 5, *dq.GetTail())

    // add beyond cap
    e := dq.AppendTail(5)
    assert.NotNil(t, e)
    e = dq.InsertHead(5)
    assert.NotNil(t, e)
    assert.Equal(t, 5, dq.GetSize())

    // pop head
    assert.Equal(t, 1, *dq.PopHead())
    assert.Equal(t, 4, dq.GetSize())
    assert.Equal(t, 2, *dq.GetHead())

    // insert head
    dq.InsertHead(6)
    assert.Equal(t, 6, *dq.GetHead())
    assert.Equal(t, 5, dq.GetSize())

    // pop tail
    assert.Equal(t, 5, *dq.PopTail())
    assert.Equal(t, 4, dq.GetSize())
    assert.Equal(t, 4, *dq.GetTail())

    // append tail
    dq.AppendTail(7)
    assert.Equal(t, 7, *dq.GetTail())
    assert.Equal(t, 5, dq.GetSize())

    // pop until size=0
    dq.PopTail()
    dq.PopTail()
    assert.Equal(t, 3, dq.GetSize())
    dq.PopHead()
    dq.PopHead()
    assert.Equal(t, 1, dq.GetSize())
    dq.PopTail()
    assert.Equal(t, 0, dq.GetSize())

    // add to empty queue
    dq.AppendTail(3)
    dq.InsertHead(2)
    assert.Equal(t, 2, dq.GetSize())
    assert.Equal(t, 2, *dq.GetHead())
    assert.Equal(t, 3, *dq.GetTail())
}

func TestMaxSlidingWindow(t *testing.T) {
    assert.Equal(t, []int{3,3,5,5,6,7}, h.MaxSlidingWindow([]int{1,3,-1,-3,5,3,6,7}, 3))
    assert.Equal(t, []int{1}, h.MaxSlidingWindow([]int{1}, 1))
    assert.Equal(t, []int{1,-1}, h.MaxSlidingWindow([]int{1,-1}, 1))
    assert.Equal(t, []int{7,4}, h.MaxSlidingWindow([]int{7,2,4}, 2))
    assert.Equal(t, []int{3,3,2,5}, h.MaxSlidingWindow([]int{1,3,1,2,0,5}, 3))
}

