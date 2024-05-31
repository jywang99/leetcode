package hard

import com "jy.org/leetcode/src/common"

// 239. Sliding Window Maximum
func MaxSlidingWindow(nums []int, k int) []int {
    // monotonic decreasing queue for holding max values
    // values in decreasing order, but holding indices
    dq := com.NewDequeue(k+1)
    // max in each window (answer)
    maxs := make([]int, len(nums)-k+1)
    mi := 0

    l, r := 0, 0
    for r < len(nums) {
        // remove old left
        if !dq.IsEmpty() && l > (*dq.GetHead()).(int) {
            dq.PopHead()
        }

        // pop all that are smaller than current value, from tail
        for !dq.IsEmpty() && nums[(*dq.GetTail()).(int)] < nums[r] {
            dq.PopTail()
        }
        // append current value
        dq.AppendTail(r)

        // when an entire window is covered
        if r+1 >= k {
            // max in the window
            maxs[mi] = nums[(*dq.GetHead()).(int)]
            mi ++
            // slide window
            l ++
        }
        r ++
    }

    return maxs
}

