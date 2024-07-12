package medium

import (
	"math"

	com "jy.org/leetcode/src/common"
)

// 53. Maximum Subarray
func maxSubArray(nums []int) int {
    m, s := math.MinInt64, 0

    for _, v := range nums {
        s += v
        m = max(m, s) // upate max so far
        s = max(s, 0) // throw away negative portion
    }

    return m
}

// 55. Jump Game
func canJump(nums []int) bool {
    t := len(nums)-1
    for i:=t; i>=0; i-- {
        if i + nums[i] >= t {
            t = i
        }
    }
    return t == 0
}

// 45. Jump Game II
func jump(nums []int) int {
    d := 0
    l, r := 0, 0

    for r < len(nums)-1 {
        far := 0
        for i:=l; i<r+1; i++ {
            far = max(far, i + nums[i])
        }
        l = r + 1
        r = far
        d += 1
    }

    return d
}

// 134. Gas Station
func canCompleteCircuit(gas []int, cost []int) int {
    sum := func(a []int) int {
        s := 0
        for _, v := range a {
            s += v
        }
        return s
    }

    if sum(gas) < sum(cost) {
        return -1
    }

    total := 0
    start := 0
    for i := range gas {
        total += gas[i] - cost[i]

        // dipped below zero, start over
        if total < 0 {
            total = 0
            start = i + 1
        }
    }

    return start
}

// 846. Hand of Straights
func IsNStraightHand(hand []int, groupSize int) bool {
    if len(hand) % groupSize != 0 {
        return false
    }

    // count of each number
    cnts := make(map[int]int)
    for _, h := range hand {
        cnts[h] ++
    }

    // unique numbers
    uhand := make([]int, len(cnts))
    i := 0
    for k := range cnts {
        uhand[i] = k
        i++
    }

    // minheap
    hp := com.NewHeap(uhand, func(a, b int) int {
        return a - b
    })
    for hp.GetSize() > 0 {
        // nubmers to be processed again
        again := make([]int, 0)

        var next *int
        for i:=0; i<groupSize; i++ {
            // empty heap
            if hp.GetSize() == 0 {
                return false
            }

            // heap top is not the expected next number
            if next != nil && *next != hp.GetTop() {
                return false
            }

            // pop, update count
            n := hp.PopTop()
            cnts[n] --
            if cnts[n] > 0 {
                again = append(again, n)
            }

            // next expected number
            n ++
            next = &n
        }

        for _, v := range again {
            hp.Insert(v)
        }
    }

    return true
}

