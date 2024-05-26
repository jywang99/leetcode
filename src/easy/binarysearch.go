package easy

// 704. Binary Search
func Search(nums []int, target int) int {
    return searchRec(nums, 0, len(nums)-1, target)
}

func searchRec(nums []int, s, e, t int) int {
    m := (s + e) / 2
    nm := nums[m]
    if nm == t {
        return m
    }
    if e-s == 0 {
        return -1
    }
    if t < nm {
        return searchRec(nums, s, m, t)
    }
    return searchRec(nums, m+1, e, t)
}

func SearchLoop(nums []int, target int) int {
    if len(nums) == 0 {
        return -1
    }
    s, e := 0, len(nums)-1
    for s <= e {
        m := (e + s) / 2
        // fmt.Printf("s=%v, m=%v, e=%v\n", s, m, e)
        if nums[m] == target {
            return m
        }
        if nums[m] < target {
            s = m+1
        } else {
            e = m-1
        }
    }
    return -1
}

