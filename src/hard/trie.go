package hard

import (
	"errors"
)

type trie struct {
    word *string
    children map[byte]*trie
}

func NewTrie() *trie {
    return &trie{
        children: map[byte]*trie{},
    }
}

func (t *trie) insert(word string) {
    n := t
    for i:=0; i<len(word); i++ {
        b := word[i]
        nn := n.children[b]
        if nn == nil {
            nn = NewTrie()
            n.children[b] = nn
        }
        n = nn
    }
    n.word = &word
}

type tracker struct {
    m [][]bool
}

func newTracker(width, height int) *tracker {
    mtx := make([][]bool, height)
    for i := range mtx {
        mtx[i] = make([]bool, width)
    }
    return &tracker{
        m: mtx,
    }
}

func (t *tracker) mark(x, y int) error {
    if t.m[y][x] {
        return errors.New("Already visited")
    }
    t.m[y][x] = true
    return  nil
}

func (t *tracker) copy() *tracker {
    nt := newTracker(len(t.m[0]), len(t.m))
    for ri, row := range(t.m) {
        for i, v := range(row) {
            nt.m[ri][i] = v
        }
    }
    return nt
}

// 212. Word Search II
func FindWords(board [][]byte, words []string) []string {
    // trie
    root := NewTrie()
    for _, w := range words {
        root.insert(w)
    }

    width, height := len(board[0]), len(board)

    // move coordinate to a direction
    getNextCoord := func(x, y, dir int) (int, int) {
        // dir: 0=up, 1=down, 2=left, 3=right
        switch dir {
        case 0:
            y--
        case 1:
            y++
        case 2:
            x--
        case 3:
            x++
        }
        return x, y
    }

    // recursively find word from coordinate
    found := make(map[string]bool)
    var findWordFrom func(int, int, *trie, tracker)
    findWordFrom = func(x, y int, t *trie, tr tracker) {
        // out of bounds
        if x < 0 || y < 0 || x >= width || y >= height {
            return
        }

        // mark coordinate
        err := tr.mark(x, y)
        if err != nil {
            // already visited this coordinate in this recursive path
            return
        }

        // to next trie node
        t = t.children[board[y][x]]

        // deviated from trie = end
        if t == nil {
            return
        }

        // end of word = add to list
        if t.word != nil {
            found[*t.word] = true
        }

        // up, down, left, right
        for d:=0; d<=3; d++ {
            nx, ny := getNextCoord(x, y, d)
            findWordFrom(nx, ny, t, *tr.copy())
        }
    }

    // start from every coordinate
    for y := range board {
        for x := range board[y] {
            findWordFrom(x, y, root, *newTracker(width, height))
        }
    }

    rs := make([]string, 0)
    for s := range found {
        rs = append(rs, s)
    }

    return rs
}
