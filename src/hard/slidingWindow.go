package hard

import "errors"

type Dequeue struct {
    arr []int
    head int
    tail int
    size int
}

func NewDequeue(size int) *Dequeue {
    d := &Dequeue{
        arr: make([]int, size),
        head: 0,
        tail: 0,
        size: 0,
    }
    return d
}

func (dq *Dequeue) IsEmpty() bool {
    return dq.size == 0
}

func (dq *Dequeue) IsFull() bool {
    return dq.size == len(dq.arr)
}

func (dq *Dequeue) GetSize() int {
    return dq.size
}

func (dq *Dequeue) InsertHead(v int) error {
    if dq.IsFull() {
        return errors.New("Dequeue is full!")
    }
    if dq.IsEmpty() {
        dq.size ++
        dq.arr[dq.head] = v
        return nil
    }
    if dq.head == 0 {
        dq.head = len(dq.arr)-1
    } else {
        dq.head --
    }
    dq.size ++
    dq.arr[dq.head] = v
    return nil
}

func (dq *Dequeue) AppendTail(v int) error {
    if dq.IsFull() {
        return errors.New("Dequeue is full!")
    }
    if dq.IsEmpty() {
        dq.size ++
        dq.arr[dq.tail] = v
        return nil
    }
    if dq.tail == len(dq.arr)-1 {
        dq.tail = 0
    } else {
        dq.tail ++
    }
    dq.size ++
    dq.arr[dq.tail] = v
    return nil
}

func (dq *Dequeue) PopHead() *int {
    if dq.IsEmpty() {
        return nil
    }
    v := dq.arr[dq.head]
    dq.size --
    if dq.size == 0 {
        dq.head = 0
        dq.tail = 0
        return &v
    }
    if dq.head < len(dq.arr)-1 {
        dq.head ++
        return &v
    }
    dq.head = 0
    return &v
}

func (dq *Dequeue) PopTail() *int {
    if dq.IsEmpty() {
        return nil
    }
    v := dq.arr[dq.tail]
    dq.size --
    if dq.size == 0 {
        dq.head = 0
        dq.tail = 0
        return &v
    }
    if dq.tail > 0 {
        dq.tail --
        return &v
    }
    dq.tail = len(dq.arr)-1
    return &v
}

func (dq *Dequeue) GetHead() *int {
    if dq.IsEmpty() {
        return nil
    }
    v := dq.arr[dq.head]
    return &v
}

func (dq *Dequeue) GetTail() *int {
    if dq.IsEmpty() {
        return nil
    }
    v := dq.arr[dq.tail]
    return &v
}

// 239. Sliding Window Maximum
func MaxSlidingWindow(nums []int, k int) []int {
    // monotonic decreasing queue for holding max values
    // values in decreasing order, but holding indices
    dq := NewDequeue(k+1)
    // max in each window (answer)
    maxs := make([]int, len(nums)-k+1)
    mi := 0

    l, r := 0, 0
    for r < len(nums) {
        // remove old left
        if !dq.IsEmpty() && l > *dq.GetHead() {
            dq.PopHead()
        }

        // pop all that are smaller than current value, from tail
        for !dq.IsEmpty() && nums[*dq.GetTail()] < nums[r] {
            dq.PopTail()
        }
        // append current value
        dq.AppendTail(r)

        // when an entire window is covered
        if r+1 >= k {
            // max in the window
            maxs[mi] = nums[*dq.GetHead()]
            mi ++
            // slide window
            l ++
        }
        r ++
    }

    return maxs
}

