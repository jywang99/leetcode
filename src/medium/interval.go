package medium

import com "jy.org/leetcode/src/common"

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

    intervals = com.Quicksort(intervals, func(a, b []int) int {
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

// 435. Non-overlapping Intervals
func eraseOverlapIntervals(intervals [][]int) int {
    intervals = com.Quicksort(intervals, func(a, b []int) int {
        return a[0] - b[0]
    })

    cnt := 0
    lend := intervals[0][1]
    for _, intv := range intervals[1:] {
        // no overlap
        if intv[0] >= lend {
            lend = intv[1]
            continue
        }
        // overlap
        cnt ++
        // remove one, keep the one that ends first
        if intv[1] < lend {
            lend = intv[1]
        }
    }

    return cnt
}

// 253. Meeting Rooms II
func MinMeetingRooms(intervals [][]int) int {
    times := make([]int, 0, 2*len(intervals))
    start := make(map[int]bool)
    for _, intv := range intervals {
        s, e := intv[0], intv[1]
        times = append(times, s)
        times = append(times, e)
        start[s] = true
        start[e] = false
    }
    times = com.Quicksort(times, func(a, b int) int {
        return a - b
    })

    mo := 1
    co := 0
    for _, t := range times {
        // end of meeting
        if !start[t] {
            co --
            continue
        }
        // start of meeting
        co ++
        mo = max(mo, co)
    }

    return mo
}

