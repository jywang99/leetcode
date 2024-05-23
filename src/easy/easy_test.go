package easy_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
    m := make(map[int]int)
    m[1] = 2
    m[2] = 3
    assert.Equal(t, 2, len(m))
    delete(m, 2)
    assert.Equal(t, 1, len(m))
}

