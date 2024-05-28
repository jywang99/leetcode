package hard

// 84. Largest Rectangle in Histogram
func LargestRectangleArea(heights []int) int {
    type bar struct {
        idx int // start idx of the height
        height int
    }
    stk := make([]bar, 1)
    stk[0] = bar{
        idx: 0,
        height: heights[0],
    }
    marea := 0
    heights = append(heights, 0)
    for i:=1; i<len(heights); i++ {
        h := heights[i]
        th := stk[len(stk)-1].height // height on stack top

        // bar taller than stack top
        if h > th {
            // add current bar to stack
            stk = append(stk, bar{
                idx: i,
                height: h,
            })
        }

        // bar shorter than stack top
        if h < th {
            // pop until top of stack height = current height
            var pidx int
            for len(stk) > 0 && stk[len(stk)-1].height > h {
                // pop
                b := stk[len(stk)-1]
                stk = stk[:len(stk)-1]
                // calc area while popping, update marea
                a := b.height*(i-b.idx)
                marea = max(marea, a)
                // keep track of last popped bar's idx
                pidx = b.idx
            }
            // push new height to stack with starting idx = last popped
            stk = append(stk, bar{
                idx: pidx,
                height: h,
            })
        }
    }
    return marea
}

