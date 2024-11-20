package medium

// 78. Subsets
func Subsets(nums []int) [][]int {
    sets := make([][]int, 0)

    var dfs func(int, []int)
    dfs = func(i int, ss []int) {
        if i >= len(nums) {
            tmp := make([]int, len(ss))
            copy(tmp, ss)
            sets = append(sets, tmp)
            return
        }

        dfs(i+1, ss)
        dfs(i+1, append(ss, nums[i]))
    }
    dfs(0, []int{})

    return sets
}

// 39. Combination Sum
func sumArray(arr []int) int {
    s := 0
    for _, val := range(arr) {
        s += val
    }
    return s
}

func CombinationSum(candidates []int, target int) [][]int {
    res := make([][]int, 0)

    var dfs func(int, []int)
    dfs = func(i int, cur []int) {
        sum := sumArray(cur)
        if sum == target {
            tmp := make([]int, len(cur))
            copy(tmp, cur)
            res = append(res, tmp)
            return
        }
        if sum > target || i >= len(candidates) {
            return
        }

        dfs(i, append(cur, candidates[i]))
        dfs(i+1, cur)
    }
    dfs(0, make([]int, 0))

    return res
}

