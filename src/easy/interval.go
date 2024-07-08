package easy

// 252. Meeting Rooms
func CanAttendMeetings(intervals [][]int) bool {
    quicksort(&intervals, func(a, b []int) int {
        return a[0] - b[0]
    })

    pend := intervals[0][1]
    for _, intv := range intervals[1:] {
        if intv[0] < pend {
            return false
        }
        pend = intv[1]
    }

    return true
}

func quicksort[T any](arr *[]T, cmpVal func(T, T) int) {
    ar := *arr
    cmp := func(i, j int) int {
        return cmpVal(ar[i], ar[j])
    }

    var sort func(int, int)
    sort = func(s, t int) {
        if t - s < 2 {
            return
        }

        l, r := s, t-1 // r = pivot
        for i:=s; i<r; i++ {
            if cmp(i, r) < 1 {
                ar[l], ar[i] = ar[i], ar[l]
                l ++
            }
        }
        ar[l], ar[r] = ar[r], ar[l]

        sort(s, l)
        sort(l+1, t)
    }

    sort(0, len(ar))
}

