package medium_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	m "jy.org/leetcode/src/medium"
)

func TestMinStack(t *testing.T) {
    ms := m.Constructor()
    ms.Push(-2)
    ms.Push(0)
    ms.Push(-3)
    assert.Equal(t, -3, ms.GetMin())
    ms.Pop()
    assert.Equal(t, 0, ms.Top())
    assert.Equal(t, -2, ms.GetMin())
}

func TestRPN(t *testing.T) {
    assert.Equal(t, 9, m.EvalRPN([]string{"2","1","+","3","*"}))
    assert.Equal(t, 6, m.EvalRPN([]string{"4","13","5","/","+"}))
    assert.Equal(t, 22, m.EvalRPN([]string{"10","6","9","3","+","-11","*","/","*","17","+","5","+"}))
}

func TestGenerateParenthesis(t *testing.T) {
    rs := m.GenerateParenthesis(3)
    assert.Equal(t, 5, len(rs))
    assert.Contains(t, rs, "((()))","(()())","(())()","()(())","()()()")
}

func TestDailyTemperatures(t *testing.T) {
    assert.Equal(t, []int{1,1,4,2,1,1,0,0}, m.DailyTemperatures([]int{73,74,75,71,69,72,76,73}))
    assert.Equal(t, []int{1,1,1,0}, m.DailyTemperatures([]int{30,40,50,60}))
    assert.Equal(t, []int{1,1,0}, m.DailyTemperatures([]int{30,60,90}))
}
