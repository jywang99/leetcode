package hard

// 42. Trapping Rain Water
func Trap(height []int) int {
    total := 0
    i := 0
    j := len(height)-1
    maxih := height[i] // max height on left side
    maxjh := height[j] // max height on right side
    for i < j {
        var h int
        var a int
        if height[i] < height[j] {
            i++
            // fmt.Println("i++")
            h = height[i]
            a = maxih - h
            maxih = max(maxih, h)
        } else {
            j--
            // fmt.Println("j--")
            h = height[j]
            a = maxjh - h
            maxjh = max(maxjh, h)
        }

        if a > 0 {
            // fmt.Printf("i=%v, j=%v, maxih=%v, maxjh=%v, a=%v\n", i, j, maxih, maxjh, a)
            total += a
        }
    }
    return total
}

