package medium

import (
	"sort"
)

func GroupAnagrams(strs []string) [][]string {
    m := make(map[string][]int)

    for i, s := range strs {
        sr := []rune(s)
        sort.Slice(sr, func(i, j int) bool {
            return sr[i] < sr[j]
        })

        ss := string(sr)
        _, e := m[ss]
        if e {
            m[ss] = append(m[ss], i)
            continue
        }
        m[ss] = []int{i}
    }

    rs := make([][]string, len(m))
    i := 0
    for _, anas := range m {
        es := make([]string, len(anas))
        for i, ai := range(anas) {
            es[i] = strs[ai]
        }
        rs[i] = es
        i++
    }

    return rs
}

func TopKFrequent(nums []int, k int) []int {
    // vals -> freq
    valfreq := make(map[int]int)
    for _, n := range nums {
        f, e := valfreq[n]
        nf := 1
        if e {
            nf += f
        }
        valfreq[n] = nf
    }

    // freq -> vals
    freqvals := make(map[int][]int)
    maxFreq := 0
    for v, f := range valfreq {
        vals, e := freqvals[f]
        if !e {
            vals = make([]int, 0)
        }
        freqvals[f] = append(vals, v)
        if f > maxFreq {
            maxFreq = f
        }
    }

    cnt := 0
    rs := make([]int, k)
    for f:=maxFreq; f>0; f-- {
        vals, e := freqvals[f]
        if e {
            for _, v := range vals {
                rs[cnt] = v
                cnt ++
                if cnt == k {
                    return rs
                }
            }
        }
    }

    return rs
}

func ProductExceptSelf(nums []int) []int {
    pre := make([]int, len(nums))
    p := 1
    for i, n := range nums {
        p =  p * n
        pre[i] = p
    }

    post := make([]int, len(nums))
    p = 1
    for i:=len(nums)-1; i>=0; i-- {
        n := nums[i]
        p =  p * n
        post[i] = p
    }

    rs := make([]int, len(nums))
    for i := range rs {
        r := 1
        if i > 0 {
            r *= pre[i-1]
        }
        if i < len(nums)-1 {
            r *= post[i+1]
        }
        rs[i] = r
    }

    return rs
}

func ProductExceptSelfSpace(nums []int) []int {
    rs := make([]int, len(nums))
    rs[0] = 1
    for i:=0; i<len(nums)-1; i++ {
        rs[i+1] = rs[i] * nums[i]
    }

    post := make([]int, len(nums))
    post[len(nums)-1] = 1
    for i:=len(nums)-1; i>0; i-- {
        post[i-1] = post[i] * nums[i]
        rs[i-1] = rs[i-1] * post[i-1]
    }

    return rs
}
