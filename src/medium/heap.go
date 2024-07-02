package medium

import (
	"math"
)

type Heap[T any] struct {
    arr *[]T
    cmp func(int, int) int // return positive = first element should be prioritized (popped first), negative = otherwise
}

func NewHeap[T any](arr []T, cmp func(T, T) int) *Heap[T] {
    hp := &Heap[T]{
        arr: &arr,
        cmp: func(i, j int) int {
            return cmp(arr[i], arr[j])
        },
    }

    for i:=len(arr)-1; i>=0; i-- {
        hp.heapify(i)
    }

    return hp
}

func (hp *Heap[T]) swap(i, j int) {
    (*hp.arr)[j], (*hp.arr)[i] = (*hp.arr)[i], (*hp.arr)[j]
}

func (hp *Heap[T]) heapify(i int) {
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

    if mi != i {
        hp.swap(mi, i)
        hp.heapify(mi)
    }
}

func (hp *Heap[T]) Insert(val T) {
    *hp.arr = append(*hp.arr, val)

    i := len(*hp.arr) - 1
    for i > 0 {
        pi := (i - 1) / 2
        if pi < 0 || hp.cmp(pi, i) > 0 {
            break
        }
        hp.swap(pi, i)
        i = pi
    }
}

func (hp *Heap[T]) PopTop() T {
    l := hp.GetSize()

    // no element in heap
    if l == 0 {
        return *new(T)
    }

    // return value = heap top
    res := (*hp.arr)[0]

    // last element
    if l == 1 {
        *hp.arr = []T{}
        return res
    }

    // put last element at top, and heapify
    (*hp.arr)[0] = (*hp.arr)[l-1]
    *hp.arr = (*hp.arr)[:l-1]
    hp.heapify(0)

    return res
}

func (hp *Heap[T]) GetTop() T {
    return (*hp.arr)[0]
}

func (hp *Heap[T]) GetSize() int {
    return len(*hp.arr)
}

// 973. K Closest Points to Origin
func KClosest(points [][]int, k int) [][]int {
    getDst := func(p []int) float64 {
        fx := float64(p[0])
        fy := float64(p[1])
        return float64(math.Sqrt(math.Pow(fx, 2) + math.Pow(fy, 2)))
    }

    hp := NewHeap(points, func(pa, pb []int) int {
        ad, bd := getDst(pa), getDst(pb)
        if ad > bd {
            return -1
        }
        return 1
    })

    res := make([][]int, 0)
    for i:=0; i<k; i++ {
        res = append(res, hp.PopTop())
    }

    return res
}

// 215. Kth Largest Element in an Array
func FindKthLargest(nums []int, k int) int {
    hp := NewHeap(nums, func(a, b int) int {
        if a == b {
            return 0
        }
        if a < b { // smaller = prioritized
            return 1
        }
        return -1
    })
    for hp.GetSize() > k {
        hp.PopTop()
    }
    return hp.GetTop()
}

// 621. Task Scheduler
func LeastInterval(tasks []byte, n int) int {
    type tnode struct {
        char byte
        freq int
        next int
    }

    // frequency map
    tmap := make(map[byte]int)
    for _, t := range tasks {
        v, e := tmap[t]
        if !e {
            tmap[t] = 1
            continue
        }
        tmap[t] = v + 1
    }

    // heap, frequent nodes prioritized
    thp := NewHeap([]*tnode{}, func(a, b *tnode) int {
        if a.freq == b.freq {
            return 0
        }
        if a.freq > b.freq {
            return 1
        }
        return -1
    })
    for t, f := range tmap {
        thp.Insert(&tnode{
            char: t,
            freq: f,
        })
    }
    tq := []*tnode{}

    // start clock
    time := 0
    for thp.GetSize() > 0 || len(tq) > 0 {
        // check if any task in queue is ready again
        if len(tq) > 0 {
            qhead := tq[0]
            if qhead.next == time {
                tq = tq[1:]
                thp.Insert(qhead)
            }
        }

        // only if there's task to process at this time
        if thp.GetSize() > 0 {
            // process task
            task := thp.PopTop()
            task.freq --

            // if not the last occurrence, push it to queue
            task.next = time + n + 1
            if task.freq > 0 {
                tq = append(tq, task)
            }
        }

        time ++
    }

    return time
}

// 355. Design Twitter
type Twitter struct {
    
}


func TwitterConstructor() Twitter {
    
}


func (this *Twitter) PostTweet(userId int, tweetId int)  {
    
}


func (this *Twitter) GetNewsFeed(userId int) []int {
    
}


func (this *Twitter) Follow(followerId int, followeeId int)  {
    
}


func (this *Twitter) Unfollow(followerId int, followeeId int)  {
    
}

