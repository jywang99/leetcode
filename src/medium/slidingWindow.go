package medium

// 3. Longest Substring Without Repeating Characters
func LengthOfLongestSubstring(s string) int {
    rs := []rune(s)
    n := len(rs)
    if n <= 1 {
        return n
    }

    // left, right ptr, local max len, global max len
    l, r, ll, ml := 0, 1, 1, 0
    // local char -> last seen index
    lset := make(map[rune]int)
    lset[rs[l]] = 0
    for r < n {
        i, e := lset[rs[r]]
        if e {
            // seen this rune before = end of substring
            l, r = i+1, i+1
            // re-init lset, retaining last char pos
            lset = make(map[rune]int)
            lset[rs[l]] = l
            // update global max len
            ml = max(ml, ll)
            ll = 1
        } else {
            // never seen this rune before = mark and continue
            lset[rs[r]] = r
            ll ++
        }
        r ++
    }
    // last substring
    if len(lset) > 0 {
        ml = max(ml, ll)
    }

    return ml
}

func LengthOfLongestSubstringFast(s string) int {
    rs := []rune(s)
    n := len(rs)
    if n <= 1 {
        return n
    }

    // left, right ptr, global max len
    l, r, ml := 0, 1, 0
    // local char -> last seen index
    pmap := make(map[rune]int)
    pmap[rs[l]] = 0
    for r < n {
        c := rs[r]
        pi, e := pmap[c]

        if !e || pi == -1 {
            // never seen this rune before = mark and continue
            pmap[c] = r
            r ++
            continue
        }

        // seen this rune before = end of substring
        // update max len
        ml = max(ml, r-l)

        // re-init map
        // forget about all occurrences before the dup
        for l <= pi {
            pmap[rs[l]] = -1
            l++
        }
        // l = pi + 1

        // right pointer can continue
        pmap[c] = r
        r ++
    }

    // last substring
    if r-l > 0 {
        ml = max(ml, r-l)
    }

    return ml
}

// 424. Longest Repeating Character Replacement
func CharacterReplacement(s string, k int) int {
    rs := []rune(s)
    n := len(rs)
    if n <= 1 {
        return n
    }

    fmap := make(map[rune]int)
    getMaxFreq := func() int {
        f := 0
        for _, v := range fmap {
            f = max(f, v)
        }
        return f
    }
    l, ml := 0, 0
    for r := range rs {
        fmap[rs[r]] ++
        for (r-l+1) - getMaxFreq() > k {
            fmap[rs[l]] --
            l ++
        }
        ml = max(ml, r-l+1)
    }

    return ml
}

// 567. Permutation in String
func CheckInclusion(s1 string, s2 string) bool {
    n1, n2 := len(s1), len(s2)
    if n1 > n2 {
        return false
    }

    // maps to track character occurrences
    pf := make([]int, 26)
    cf := make([]int, 26)
    isMatch := func() bool {
        for i, f := range pf {
            if f != cf[i] {
                return false
            }
        }
        return true
    }

    // init freq map for s1
    for i := range s1 {
        pf[s1[i] - 'a'] ++
    }

    // init freq map for first window
    for i:=0; i<n1; i++ {
        cf[s2[i] - 'a'] ++
    }
    // check first window
    if isMatch() {
        return true
    }

    l, r := 0, n1-1
    for r < n2-1 {
        // update current char map
        cf[s2[l] - 'a'] --
        l, r = l+1, r+1
        cf[s2[r] - 'a'] ++

        if isMatch() {
            return true
        }
    }

    return false
}

// 76. Minimum Window Substring
func MinWindow(s string, t string) string {
    rs, rt := []rune(s), []rune(t)
    ns := len(rs)
    if len(t) > ns {
        return ""
    }

    // char occurrences
    tmap := make(map[rune]int)
    for _, c := range rt {
        tmap[c] ++
    }
    need := len(tmap) // # of conditions to satisfy

    // current window info
    cmap := make(map[rune]int)
    have := 0 // # of conditions met
    addAndCheck := func(c rune) bool {
        if tmap[c] == 0 { // no occurrence in pattern = ignore
            return have == need
        }
        cmap[c] ++
        if cmap[c] == tmap[c] { // have same # only after adding
            have ++
        }
        return have == need
    }
    removeAndCheck := func(c rune) bool {
        if tmap[c] == 0 { // no occurrence in pattern = ignore
            return have == need
        }
        if cmap[c] == tmap[c] { // have same # before removing
            have --
        }
        if cmap[c] > 0 {
            cmap[c] --
        }
        return have == need
    }

    // result
    ml, ms, mt := ns+1, -1, -1
    updateResult := func(l, r int) {
        nl := r-l+1
        if ml > nl {
            ml, ms, mt = nl, l, r
        }
    }
    
    l, r := 0, 0
    for r < ns {
        if addAndCheck(rs[r]) {
            updateResult(l, r)
            for ok := true; ok; {
                l ++
                ok = removeAndCheck(rs[l-1])
                if ok {
                    updateResult(l, r)
                }
            }
        }
        r ++
    }

    if ms == -1 {
        return ""
    }
    return string(rs[ms:mt+1])
}

