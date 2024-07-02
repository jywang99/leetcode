package medium

// 208. Implement Trie (Prefix Tree)
type Trie struct {
    isWord bool
    children map[rune]*Trie
}

func TrieConstructor() Trie {
    return Trie{
        isWord: false,
        children: make(map[rune]*Trie),
    }
}

func (this *Trie) Insert(word string)  {
    t := this
    for _, c := range word {
        nt := t.children[c]
        if nt == nil {
            ct := TrieConstructor()
            nt = &ct
            t.children[c] = nt
        }
        t = nt
    }
    t.isWord = true
}

func (this *Trie) Search(word string) bool {
    t := this
    for _, c := range word {
        nt := t.children[c]
        t = nt
        if t == nil {
            return false
        }
    }
    return t.isWord
}

func (this *Trie) StartsWith(prefix string) bool {
    t := this
    for _, c := range prefix {
        nt := t.children[c]
        t = nt
        if t == nil {
            return false
        }
    }
    return true
}

// 211. Design Add and Search Words Data Structure
type WordDictionary struct {
    children map[rune]*WordDictionary
    isWord bool
}

func DictConstructor() WordDictionary {
    return WordDictionary{
        children: make(map[rune]*WordDictionary),
    }
}

func (this *WordDictionary) AddWord(word string)  {
    d := this
    for _, c := range word {
        nd := d.children[c]
        if nd == nil {
            cd := DictConstructor()
            nd = &cd
            d.children[c] = nd
        }
        d = nd
    }
    d.isWord = true
}

func (this *WordDictionary) Search(word string) bool {
    var search func(*WordDictionary, []rune) bool
    search = func(wd *WordDictionary, str []rune) bool {
        if wd == nil {
            return false
        }
        if len(str) == 0 {
            if wd.isWord {
                return true
            }
            return false
        }

        // single char
        c := str[0]
        rem := str[1:]
        if c != '.' {
            return search(wd.children[c], rem)
        }

        // .
        for _, nd := range wd.children {
            if search(nd, rem) {
                return true
            }
        }
        return false
    }
    return search(this, []rune(word))
}

