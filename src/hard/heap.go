package hard

type heap[T any] struct {
    arr *[]T
    cmp func(int, int) int // >0 = first should be on top
}

func NewHeap[T any](arr []T, cmp func(T, T) int) *heap[T] {
    hp := &heap[T]{
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

func (hp *heap[T]) heapify(i int) {
    mi := i
    li := i*2 + 1
    ri := i*2 + 2
    l := len(*hp.arr)

    if li < l && hp.cmp(mi, li) < 0 {
        mi = li
    }

    if ri < l && hp.cmp(mi, ri) < 0 {
        mi = ri
    }

    if i != mi {
        hp.swap(i, mi)
        hp.heapify(mi)
    }
}

func (hp *heap[T]) Insert(val T) {
    *hp.arr = append(*hp.arr, val)
    i := len(*hp.arr) - 1
    for i > 0 {
        pi := (i-1) / 2
        if hp.cmp(pi, i) > 0 {
            break
        }
        hp.swap(pi, i)
        i = pi
    }
}

func (hp *heap[T]) swap(i, j int) {
    (*hp.arr)[i], (*hp.arr)[j] = (*hp.arr)[j], (*hp.arr)[i]
}

func (hp *heap[T]) GetTop() T {
    return (*hp.arr)[0]
}

func (hp *heap[T]) PopTop() T {
    res := (*hp.arr)[0]

    l := len(*hp.arr)
    (*hp.arr)[0] = (*hp.arr)[l-1]
    (*hp.arr) = (*hp.arr)[:l-1]
    hp.heapify(0)

    return res
}

func (hp *heap[T]) Len() int {
    return len(*hp.arr)
}

// 295. Find Median from Data Stream
type MedianFinder struct {
    lheap *heap[int]
    gheap *heap[int]
}

func MfConstructor() MedianFinder {
    lh := NewHeap([]int{}, func(a, b int) int {
        return a - b
    })
    gh := NewHeap([]int{}, func(a, b int) int {
        return b - a
    })
    return MedianFinder{
        lheap: lh,
        gheap: gh,
    }
}

func (this *MedianFinder) AddNum(num int) {
    m := this.FindMedian()
    if float64(num) > m {
        this.gheap.Insert(num)
    } else {
        this.lheap.Insert(num)
    }
    ll := this.lheap.Len()
    gl := this.gheap.Len()
    if ll - gl > 1 {
        this.gheap.Insert(this.lheap.PopTop())
    }
    if gl - ll > 1 {
        this.lheap.Insert(this.gheap.PopTop())
    }
}

func (this *MedianFinder) FindMedian() float64 {
    ll := this.lheap.Len()
    gl := this.gheap.Len()
    if ll == 0 && gl == 0 {
        return 0
    }
    if ll > gl {
        return float64(this.lheap.GetTop())
    }
    if gl > ll {
        return float64(this.gheap.GetTop())
    }
    return (float64(this.lheap.GetTop()) + float64(this.gheap.GetTop())) / 2
}

