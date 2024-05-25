package easy

import "unicode"

func IsPalindrome(s string) bool {
    r := []rune(s)
    i := 0
    j := len(s)-1
    for ; ; {
        // skip non-alphabet chars
        for ; i<j && !isRuneAN(r[i]); {
            i++
        }
        for ; i<j && !isRuneAN(r[j]); {
            j--
        }
        if j <= i {
            return true
        }
        if unicode.ToLower(r[i]) != unicode.ToLower(r[j]) {
            return false
        }
        i++
        j--
    }
}

func isRuneAN(r rune) bool {
    return unicode.IsLetter(r) || unicode.IsNumber(r)
}

