package medium

import (
	"sort"
)

// 167. Two Sum II - Input Array Is Sorted
func TwoSum(numbers []int, target int) []int {
    cs := make(map[int]int)
    for i, n := range numbers {
        oi, e := cs[n]
        if e {
            return []int{oi+1, i+1}
        }
        cs[target - n] = i
    }
    return []int{}
}

func TwoSumTwoPtr(numbers []int, target int) []int {
    i := 0
    j := len(numbers)-1
    s := 0

    for i < j {
        s = numbers[i] + numbers[j]
        if s == target {
            return []int{i+1, j+1}
        }
        if s < target {
            i ++
        }
        if s > target {
            j --
        }
    }
    return []int{}
}

func ThreeSum(nums []int) [][]int {
    sort.Ints(nums)
    rs := make([][]int, 0)
    ln := len(nums)
    for i, n := range nums[:ln-2] {
        // skip all duplicate for i
        if i > 0 && n == nums[i-1] {
            continue
        }
        // all possible j, k values are too big
        if n + nums[i+1] + nums[i+2] > 0 {
            break
        }
        // all possible j, k values are too small
        if n + nums[ln-2] + nums[ln-1] < 0 {
            continue
        }

        // double pointer two sum
        j := i+1
        k := ln-1
        for j < k {
            s := n + nums[j] + nums[k]
            if s < 0 {
                j ++
            } 
            if s > 0 {
                k --
            }
            if s == 0 {
                rs = append(rs, []int{n, nums[j], nums[k]})
                j, k = j+1, k-1
                // skip all duplicates for j
                for j<k && nums[j] == nums[j-1] {
                    j++
                }
                // skip all duplicates for k
                for k>j && nums[k] == nums[k+1] {
                    k--
                }
            }
        }
    }
    return rs
}

// 11. Container With Most Water
func MaxArea(height []int) int {
    marea := 0
    i := 0
    j := len(height)-1
    for i < j {
        a := (j-i) * min(height[i], height[j])
        marea = max(marea, a)
        if height[i] < height[j] {
            i++
        } else {
            j--
        }
    }
    return marea
}

