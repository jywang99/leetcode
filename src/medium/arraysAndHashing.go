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

// space complexity of O(n) except for the returned slice
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

func IsValidSudoku(board [][]byte) bool {
    rows := make([]map[byte]bool, 9)
    cols := make([]map[byte]bool, 9)
    squares := make([][]map[byte]bool, 3, 3)

    for ri, row := range board {
        for ci, c := range row {
            if string(c) == "." {
                continue
            }

            if rows[ri] == nil {
                rows[ri] = make(map[byte]bool)
            }
            if rows[ri][c] {
                return false
            } else {
                rows[ri][c] = true
            }
            if cols[ci] == nil {
                cols[ci] = make(map[byte]bool)
            }
            if cols[ci][c] {
                return false
            } else {
                cols[ci][c] = true
            }

            sri := ri / 3
            sci := ci / 3
            if squares[sri] == nil {
                squares[sri] = make([]map[byte]bool, 3)
            }
            sqmap := squares[sri][sci]
            if sqmap == nil {
                squares[sri][sci] = make(map[byte]bool)
            }
            if sqmap[c] {
                return false
            } else {
                squares[sri][sci][c] = true
            }
        }
    }

    return true
}

func LongestConsecutive(nums []int) int {
    nmap := make(map[int]bool)
    for _, n := range nums {
        nmap[n] = true
    }
    ll := 0
    for _, n := range nums {
        if !nmap[n-1] {
            l := 1
            for i:=n+1; nmap[i]; i++ {
                l++
            }
            ll = max(ll, l)
        }
    }
    return ll
}

