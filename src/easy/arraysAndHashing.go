package easy

import (
	"sort"
)

func containsDuplicate(nums []int) bool {
    m := make(map[int]bool)
    for _, n := range nums {
        if m[n] {
            return true
        }
        m[n] = true
    }
    return false
}

func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    rs := []rune(s)
    rt := []rune(t)

    cmap := make(map[rune]int)
    for _, c := range rs {
        cmap[c] ++
    }

    for _, c := range rt {
        cmap[c] --
        if cmap[c] == 0 {
            delete(cmap, c)
        }
    }

    return len(cmap) == 0
}


func twoSum(nums []int, target int) []int {
    m := make(map[int]int)

    for i, n := range nums {
        ci, e := m[n]
        if e {
            return []int{ci, i}
        }
        m[target-n] = i
    }

    return []int{}
}

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

func inList(nums []int, n int) bool {
    for _, m := range nums {
        if n == m {
            return true
        }
    }
    return false
}
