package medium

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
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
    p := nums[0]-1 // arbitrary
    rsm := make(map[string]bool)
    for i, n := range nums {
        if n == p {
            continue
        }
        p = n

        j := i+1
        k := len(nums)-1
        for j < k {
            s := nums[i]+nums[j]+nums[k]
            if s == 0 {
                rsm[fmt.Sprintf("%v,%v,%v", nums[i], nums[j], nums[k])] = true
                j++
                continue
            }
            if s > 0 {
                k --
            } else {
                j ++
            }
        }
    }
    rs := make([][]int, len(rsm))
    ri := 0
    for k := range rsm {
        r := make([]int, 3)
        sarr := strings.Split(k, ",")
        for i, s := range sarr {
            iv, _ := strconv.Atoi(s)
            r[i] = iv
        }
        rs[ri] = r
        ri ++
    }
    return rs
}

