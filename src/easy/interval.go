package easy

import com "jy.org/leetcode/src/common"

// 252. Meeting Rooms
func CanAttendMeetings(intervals [][]int) bool {
    intervals = com.Quicksort(intervals, func(a, b []int) int {
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

