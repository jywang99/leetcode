package medium

// 57. Insert Interval
func Insert(intervals [][]int, newInterval []int) [][]int {
    rs := make([][]int, 0)

    for i, intv := range intervals {
        if newInterval[1] < intv[0] {
            rs = append(rs, newInterval)
            return append(rs, intervals[i:]...)
        }
        if intv[1] < newInterval[0] {
            rs = append(rs, intv)
        } else {
            newInterval = []int{
                min(newInterval[0], intv[0]),
                max(newInterval[1], intv[1]),
            }
        }
    }
    rs = append(rs, newInterval)

    return rs
}

// 56. Merge Intervals
func Merge(intervals [][]int) [][]int {
    rs := make([][]int, 0)

    Quicksort(&intervals, func(a, b []int) int {
        return a[0] - b[0]
    })

    current := []int{}
    for _, intv := range intervals {
        if len(current) == 0 {
            current = intv
        } else if current[1] < intv[0] {
            rs = append(rs, current)
            current = intv
        } else {
            current = []int{
                min(current[0], intv[0]),
                max(current[1], intv[1]),
            }
        }
    }
    rs = append(rs, current)

    return rs
}

func Quicksort[T any](arr *[]T, cmpVal func(T, T) int) {
    ar := *arr
    cmp := func(i, j int) int {
        return cmpVal(ar[i], ar[j])
    }

    var sort func(int, int)
    sort = func(s, t int) {
        if t - s < 2 {
            return
        }

        l := s
        r := t - 1 // pivot
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

    sort(0, len(ar))
}

