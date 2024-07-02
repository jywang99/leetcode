package medium_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	m "jy.org/leetcode/src/medium"
)

func TestTrie(t *testing.T) {
    trie := m.TrieConstructor()
    trie.Insert("apple")
    assert.True(t, trie.Search("apple"))
    assert.False(t, trie.Search("app"))
    assert.True(t, trie.StartsWith("app"))
    trie.Insert("app")
    assert.True(t, trie.Search("app"))
}

func TestWordDictionary(t *testing.T) {
    wd := m.DictConstructor()
    wd.AddWord("bad")
    wd.AddWord("dad")
    wd.AddWord("mad")
    assert.False(t, wd.Search("pad"))
    assert.True(t, wd.Search("bad"))
    assert.True(t, wd.Search(".ad"))
    assert.True(t, wd.Search("b.."))
}

