package hard

import com "jy.org/leetcode/src/common"

// 295. Find Median from Data Stream
type MedianFinder struct {
    lheap *com.Heap[int]
    gheap *com.Heap[int]
}

func MfConstructor() MedianFinder {
    lh := com.NewHeap([]int{}, func(a, b int) int {
        return b - a
    })
    gh := com.NewHeap([]int{}, func(a, b int) int {
        return a - b
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
    ll := this.lheap.GetSize()
    gl := this.gheap.GetSize()
    if ll - gl > 1 {
        this.gheap.Insert(this.lheap.PopTop())
    }
    if gl - ll > 1 {
        this.lheap.Insert(this.gheap.PopTop())
    }
}

func (this *MedianFinder) FindMedian() float64 {
    ll := this.lheap.GetSize()
    gl := this.gheap.GetSize()
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

