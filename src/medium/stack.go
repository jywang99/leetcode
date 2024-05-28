package medium

import (
	"sort"
	"strconv"
)

// 155. Min Stack
/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
type MinStack struct {
    stack []int
    mStack []int
}

func Constructor() MinStack {
    return MinStack{
        stack: make([]int, 0),
        mStack: make([]int, 0),
    }
}

func (this *MinStack) Push(val int)  {
    this.stack = append(this.stack, val)
    lm := len(this.mStack)
    if lm == 0 || val <= this.mStack[lm-1] {
        this.mStack = append(this.mStack, val)
        return
    }
    this.mStack = append(this.mStack, this.mStack[lm-1])
}

func (this *MinStack) Pop()  {
    nsize := len(this.stack)-1
    this.stack = this.stack[:nsize]
    this.mStack = this.mStack[:nsize]
}

func (this *MinStack) Top() int {
    return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
    return this.mStack[len(this.mStack)-1]
}

// 150. Evaluate Reverse Polish Notation
func EvalRPN(tokens []string) int {
    ops := map[string]func(int,int)int {
        "+": func(i, j int) int { return i + j },
        "-": func(i, j int) int { return i - j },
        "*": func(i, j int) int { return i * j },
        "/": func(i, j int) int { return i / j },
    }
    vstk := make([]int, 0)
    for _, s := range tokens {
        op, e := ops[s]
        if !e {
            i, _ := strconv.Atoi(s)
            vstk = append(vstk, i)
            continue
        }
        i, j := vstk[len(vstk)-2], vstk[len(vstk)-1]
        vstk = vstk[:len(vstk)-2]
        vstk = append(vstk, op(i, j))
    }
    return vstk[0]
}

// 22. Generate Parentheses
func GenerateParenthesis(n int) []string {
    rs := make([]string, 0)
    stk := make([]rune, 0)
    var genParens func(int, int)
    genParens = func (opens, closes int) {
        if opens == closes && closes == n {
            rs = append(rs, string(stk))
        }
        if opens < n {
            stk = append(stk, '(')
            genParens(opens + 1, closes)
            stk = stk[:len(stk)-1]
        }
        if closes < opens {
            stk = append(stk, ')')
            genParens(opens, closes + 1)
            stk = stk[:len(stk)-1]
        }
    }
    genParens(0,0)
    return rs
}

// 739. Daily Temperatures
func DailyTemperatures(temperatures []int) []int {
    // unresolved entries
    // each entry a: a[0]: temperature, a[1]: index
    n := len(temperatures)
    stk := make([][]int, 0, n)
    rs := make([]int, len(temperatures), n)
    for i, t := range temperatures {
        for len(stk) > 0 && stk[len(stk)-1][0] < t {
            le := stk[len(stk)-1]
            stk = stk[:len(stk)-1]
            rs[le[1]] = i-le[1]
        }
        stk = append(stk, []int{t, i})
    }
    return rs
}

// 853. Car Fleet
func CarFleet(target int, position []int, speed []int) int {
    if len(position) == 1 {
        return 1
    }

    // position -> time it takes to get to target without joining any fleet
    tmap := make(map[int]float64)
    for i, p := range position {
        time := float64(target - p) / float64(speed[i])
        tmap[p] = time
    }
    sort.Ints(position)

    // positions
    fleets := make([]int, 1)
    // initialize with the closest car to target
    fleets[0] = position[len(position)-1]
    for i:=len(position)-2; i>=0; i-- {
        ptime := tmap[fleets[len(fleets)-1]] // time that previous fleet arrives
        if tmap[position[i]] <= ptime { // catches up, joins fleet
            continue
        }
        // , doesn't catch up, forms new fleet
        fleets = append(fleets, position[i])
    }

    return len(fleets)
}

