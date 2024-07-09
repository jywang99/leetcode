package hard

// 1851. Minimum Interval to Include Each Query
func MinInterval(intervals [][]int, queries []int) []int {
    cmpRange := func(a, b []int) int {
        sd := a[0] - b[0] // len of interval
        if sd != 0 {
            return sd
        }
        return a[1] - b[1] // end pos of interval
    }
    quicksort(&intervals, cmpRange)
    res := make([]int, len(queries))
    qmap := make(map[int]int)
    for i, q := range queries {
        qmap[q] = i
    }
    quicksort(&queries, func(a, b int) int {
        return a - b
    })

    hp := NewHeap([][]int{}, cmpRange)
    for _, q := range queries {
        // push all intervals that caontain current point into heap
        for _, intv := range intervals {
            if q < intv[0] || q > intv[1] {
                break
            }
            hp.Insert([]int{
                intv[1] - intv[0] + 1, // len of interval
                intv[1], // end pos of interval
            })
        }

        // pop all intervals that are too far left
        for hp.Len() > 0 && hp.GetTop()[1] < q {
            hp.PopTop()
        }

        // shortest remaining interval is the result for this query
        if hp.Len() > 0 {
            res[qmap[q]] = hp.GetTop()[0]
        } else {
            res[qmap[q]] = -1
        }
    }

    return res
}

func quicksort[T any](arr *[]T, cmpVal func(T, T) int) {
    ar := *arr
    cmp := func(a, b int) int {
        return cmpVal(ar[a], ar[b])
    }

    var sort func(int, int)
    sort = func(s, t int) {
        if t-s < 2 {
            return
        }
        l, r := s, t-1
        for i:=s; i<r; i++ {
            if cmp(i, r) < 0 {
                ar[l], ar[i] = ar[i], ar[l]
                l++
            }
        }
        ar[l], ar[r] = ar[r], ar[l]
        sort(s, l)
        sort(l+1, t)
    }
}

