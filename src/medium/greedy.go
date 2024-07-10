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
}

