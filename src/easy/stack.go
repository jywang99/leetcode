package easy

// 20. Valid Parentheses
func IsValid(s string) bool {
    ls := len(s)
    if ls > 0 && ls % 2 != 0 {
        return false
    }

    cps := map[rune]rune {
        '(': ')',
        '[': ']',
        '{': '}',
    }
    stk := make([]rune, 0)
    rs := []rune(s)
    for _, r := range rs {
        c, e := cps[r]
        if e {
            // new pair started
            stk = append(stk, c)
            continue
        }
        // invalid closing
        lst := len(stk)
        if lst == 0 || stk[lst-1] != r {
            return false
        }
        // valid closing
        stk = stk[:len(stk)-1]
    }
    return len(stk) == 0
}

