package medium

// 74. Search a 2D Matrix
func SearchMatrix(matrix [][]int, target int) bool {
    sr, er := 0, len(matrix)-1
    tr := -1
    for sr <= er {
        mr := (sr + er) / 2
        // fmt.Printf("sr=%v, mr=%v, er=%v\n", sr, mr, er)
        mrow := matrix[mr]
        if mrow[0] <= target && target <= mrow[len(mrow)-1] {
            tr = mr
            break
        }
        if target < mrow[0] {
            er = mr - 1
        } else {
            sr = mr + 1
        }
    }

    if tr >= 0 && binarySearch(matrix[tr], target) >= 0 {
        return true
    }
    return false
}

func binarySearch(arr []int, t int) int {
    s, e := 0, len(arr)-1
    for s <= e {
        m := (s + e) / 2
        if arr[m] == t {
            return m
        }
        if e - s <= 0 {
            return -1
        }
        if arr[m] < t {
            s = m + 1
        } else {
            e = m - 1
        }
    }
    return -1
}

