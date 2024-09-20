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

// 1899. Merge Triplets to Form Target Triplet
func MergeTriplets(triplets [][]int, target []int) bool {
    resolved := make(map[int]bool)
    for _, trip := range triplets {
        // matches are only updated to resolved if not useless
        useless := false
        match := make(map[int]bool)
        for j, v := range trip {
            // if triplet contains a column > target, skip it
            if v > target[j] {
                useless = true
                break
            }
            // don't care about already resolved column
            if resolved[j] {
                continue
            }
            if v == target[j] {
                match[j] = true
            }
        }
        // if valid triplet, update resolve status
        if !useless {
            for k, v := range match {
                resolved[k] = v
            }
            // all columns matched, done
            if len(resolved) == 3 {
                return true
            }
        }
    }
    // finished looping through all triplets but still have unresolved columns
    return false
}

// 763. Partition Labels
func PartitionLabels(s string) []int {
    sb := []byte(s)

    // get last position of all characters
    last := make(map[byte]int)
    for i, c := range sb {
        last[c] = i
    }

    // partition
    parts := make([]int, 0)
    size := 0
    stop := 0
    for i, c := range sb {
        size ++
        stop = max(stop, last[c])

        // end partition
        if i == stop {
            parts = append(parts, size)
            size = 0
        }
    }

    return parts
}
