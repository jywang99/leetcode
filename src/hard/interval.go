package hard

import (
	"slices"

	com "jy.org/leetcode/src/common"
)

// 1851. Minimum Interval to Include Each Query
func MinInterval(intervals [][]int, queries []int) []int {
    slices.SortFunc(intervals, func(a, b []int) int {
        return a[0] - b[0]
    })
    sq := slices.Clone(queries)
    slices.SortFunc(sq, func(a, b int) int {
        return a - b
    })

    hp := com.NewHeap([][]int{}, func(a, b []int) int {
        sd := a[0] - b[0] // len of interval, shorter = top
        if sd != 0 {
            return sd
        }
        return a[1] - b[1] // end pos of interval, earlier = top
    })
    i := 0
    resMap := make(map[int]int)
    for _, q := range sq {
        // push all intervals that start before/at q into heap
        for i < len(intervals) && intervals[i][0] <= q {
            intv := intervals[i]
            hp.Insert([]int{
                intv[1] - intv[0] + 1, // len of interval
                intv[1], // end pos of interval
            })
            i ++
        }

        // pop all intervals that are too far left
        for hp.GetSize() > 0 && hp.GetTop()[1] < q {
            hp.PopTop()
        }

        // shortest remaining interval is the result for this query
        if hp.GetSize() > 0 {
            resMap[q] = hp.GetTop()[0]
        } else {
            resMap[q] = -1
        }
    }

    res := make([]int, len(queries))
    for i, v := range queries {
        res[i] = resMap[v]
    }

    return res
}

