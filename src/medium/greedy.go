package medium

import "math"

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

