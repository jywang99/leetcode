package easy

// 217. Contains Duplicate
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

// 242. Valid Anagram
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

// 1. Two Sum
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

