package easy

type Heap[T any] struct {
    arr *[]T
    cmp func(int, int) int
}

func NewHeap[T any](arr []T, cmp func(T, T) int) *Heap[T] {
    hp := &Heap[T]{
        arr: &arr,
        cmp: func(a, b int) int {
            return cmp(arr[a], arr[b])
        },
    }
    for i:=len(arr)-1; i>=0; i-- {
        hp.heapify(i)
    }
    return hp
}

func (h *Heap[T]) heapify(i int) {
    // indices
    mi := i // index that should be placed at the top
    li := i * 2 + 1
    ri := i * 2 + 2
    l := len(*h.arr)

    // check left child
    if li < l && h.cmp(mi, li) > 0 {
        mi = li
    }

    // check right child
    if ri < l && h.cmp(mi, ri) > 0 {
        mi = ri
    }

    // recursively heapify the affected subtree
    if mi != i {
        (*h.arr)[i], (*h.arr)[mi] = (*h.arr)[mi], (*h.arr)[i]
        h.heapify(mi)
    }
}

func (h *Heap[T]) GetSize() int {
    return len(*h.arr)
}

func (h *Heap[T]) GetTop() T {
    return (*h.arr)[0]
}

func (h *Heap[T]) PopTop() T {
    if len((*h.arr)) == 0 {
        return *new(T)
    }

    res := (*h.arr)[0]
    l := len(*h.arr)

    if l == 1 {
        (*h.arr) = (*h.arr)[1:]
        return res
    }

    (*h.arr)[0] = (*h.arr)[l-1]
    (*h.arr) = (*h.arr)[:l-1]
    h.heapify(0)

    return res
}

func (h *Heap[T]) Insert(val T) {
    // append to end
    (*h.arr) = append((*h.arr), val)

    // swim it
    i := len(*h.arr)-1
    for i > 0 {
        pi := (i - 1) / 2
		if h.cmp(i, pi) >= 0 {
			break
		}
        (*h.arr)[pi], (*h.arr)[i] = (*h.arr)[i], (*h.arr)[pi] 
        i = pi
    }
}

// 703. Kth Largest Element in a Stream
type KthLargest struct {
    k int
    hp *Heap[int]
}

func KthConstructor(k int, nums []int) KthLargest {
    hp := NewHeap(nums, func(a, b int) int {
        if a == b {
            return 0
        }
        if a > b {
            return 1
        }
        return -1
    })

    for hp.GetSize() > k {
        hp.PopTop()
    }

    return KthLargest{
        k: k,
        hp: hp,
    }
}

func (this *KthLargest) Add(val int) int {
	if this.hp.GetSize() < this.k {
		this.hp.Insert(val)
	} else if this.hp.GetTop() < val {
		this.hp.PopTop()
		this.hp.Insert(val)
	}
	return this.hp.GetTop()
}

// 1046. Last Stone Weight
func lastStoneWeight(stones []int) int {
    hp := NewHeap(stones, func(a, b int) int {
        if a == b {
            return 0
        }
        if a > b {
            return -1
        }
        return 1
    })

    var r1, r2 int
    for hp.GetSize() > 1 {
        r1, r2 = hp.PopTop(), hp.PopTop()
        if r1 == r2 {
            continue
        }
        hp.Insert(abs(r1 - r2))
    }

    if hp.GetSize() == 0 {
        return 0
    }
    return hp.GetTop()
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
