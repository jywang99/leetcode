package easy

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

