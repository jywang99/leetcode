package easy_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"jy.org/leetcode/src/easy"
)

func TestMap(t *testing.T) {
    m := make(map[int]int)
    m[1] = 2
    m[2] = 3
    assert.Equal(t, 2, len(m))
    delete(m, 2)
    assert.Equal(t, 1, len(m))
}

func TestValidPalindrome(t *testing.T) {
    assert.True(t, easy.IsPalindrome("A man, a plan, a canal: Panama"))
    assert.False(t, easy.IsPalindrome("0P"))
}

func TestValidParentheses(t *testing.T) {
    assert.True(t, easy.IsValid("()"))
    assert.True(t, easy.IsValid("()[]{}"))
    assert.False(t, easy.IsValid("(]"))
    assert.True(t, easy.IsValid("{[]}"))
    assert.False(t, easy.IsValid("(("))
    assert.False(t, easy.IsValid("){"))
}

