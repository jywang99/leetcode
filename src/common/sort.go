package common

func Quicksort[T any](arr []T, cmpVal func(T, T) int) []T {
    cmp := func(a, b int) int {
        return cmpVal(arr[a], arr[b])
    }

    var sort func(int, int)
    sort = func(s, t int) {
        if t-s < 2 {
            return
        }
        l, r := s, t-1
        for i:=s; i<r; i++ {
            if cmp(i, r) < 0 {
                arr[l], arr[i] = arr[i], arr[l]
                l++
            }
        }
        arr[l], arr[r] = arr[r], arr[l]
        sort(s, l)
        sort(l+1, t)
    }
    sort(0, len(arr))

    return arr
}

