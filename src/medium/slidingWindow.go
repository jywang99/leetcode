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

    l, r, ml := 0, 1, 0
    fmap := make(map[rune]int)
    getMaxFreq := func() int {
        f := 0
        for _, v := range fmap {
            f = max(f, v)
        }
        return f
    }
    for r < n {
        c := rs[r]
        fmap[c] = fmap[c] + 1
        for (r-l+1) - getMaxFreq() > k {
            fmap[rs[l]] --
            l ++
        }
        ml = max(ml, r-l+1)
        r ++
    }

    return ml
}
