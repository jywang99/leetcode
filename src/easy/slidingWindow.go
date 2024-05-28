package easy

// 121. Best Time to Buy and Sell Stock
func MaxProfit(prices []int) int {
    l, r, mp := 0, 1, 0
    for r < len(prices) {
        if prices[l] < prices[r] {
            mp = max(mp, prices[r] - prices[l])
        } else {
            l = r
        }
        r ++
    }
    return mp
}

func MaxProfitSliding(prices []int) int {
    n := len(prices)
    mins, maxs := make([]int, n), make([]int, n)

    // minimum values up until i
    mins[0] = prices[0]
    for i:=1; i<n; i++ {
        mins[i] = min(mins[i-1], prices[i])
    }

    // maximum values up until i, from end
    maxs[n-1] = prices[n-1]
    for i:=n-2; i>=0; i-- {
        maxs[i] = max(maxs[i+1], prices[i])
    }

    // compare max price obtained at each index
    mp := 0
    for i:=0; i<n; i++ {
        mp = max(mp, maxs[i] - mins[i])
    }

    return mp
}

