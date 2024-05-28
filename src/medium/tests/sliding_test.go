package medium_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	m "jy.org/leetcode/src/medium"
)

func TestLengthOfLongestSubstring(t *testing.T) {
    assert.Equal(t, 3, m.LengthOfLongestSubstring("abcabcbb"))
    assert.Equal(t, 1, m.LengthOfLongestSubstring("bbbbb"))
    assert.Equal(t, 3, m.LengthOfLongestSubstring("pwwkew"))
    assert.Equal(t, 2, m.LengthOfLongestSubstring("au"))
    assert.Equal(t, 3, m.LengthOfLongestSubstring("dvdf"))

    assert.Equal(t, 3, m.LengthOfLongestSubstringFast("abcabcbb"))
    assert.Equal(t, 1, m.LengthOfLongestSubstringFast("bbbbb"))
    assert.Equal(t, 3, m.LengthOfLongestSubstringFast("pwwkew"))
    assert.Equal(t, 2, m.LengthOfLongestSubstringFast("au"))
    assert.Equal(t, 3, m.LengthOfLongestSubstringFast("dvdf"))
}

func TestCharacterReplacement(t *testing.T) {
    assert.Equal(t, 4, m.CharacterReplacement("ABAB", 2))
    assert.Equal(t, 4, m.CharacterReplacement("AABABBA", 1))
}
