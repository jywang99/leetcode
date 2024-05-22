package easy_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"jy.org/leetcode/src/easy"
)

func Test(t *testing.T) {
    m := make(map[int]int)
    m[1] = 2
    m[2] = 3
    assert.Equal(t, 2, len(m))
    delete(m, 2)
    assert.Equal(t, 1, len(m))
}

func TestGroupAnagrams(t *testing.T) {
    easy.GroupAnagrams([]string{"eat","tea","tan","ate","nat","bat"})
}

func TestTopKFrequent(t *testing.T) {
    fmt.Println(easy.TopKFrequent([]int{1,1,1,2,2,3}, 2))
}

