package queue

import (
	"daily_test/data-tructures-and-algorithms/stack"
	"strconv"
)

// 逆波兰表达式
func evalPRN(arr []string) int {
	symbol := map[string]string{
		"+": "+",
		"-": "-",
		"*": "*",
		"/": "/",
	}
	operation := func(n1, n2 int, symbol string) int {
		switch symbol {
		case "+":
			return n1 + n2
		case "-":
			return n1 - n2
		case "*":
			return n1 * n2
		case "/":
			return n1 / n2
		}
		return 0
	}
	s := stack.Init()
	for i := 0; i < len(arr); i++ {
		if v, ok := symbol[arr[i]]; ok {
			n1, _ := strconv.Atoi(s.Pop().(string))
			n2, _ := strconv.Atoi(s.Pop().(string))
			tmp := operation(n2, n1, v)
			s.Push(strconv.Itoa(tmp))
			continue
		}
		s.Push(arr[i])
	}
	r, _ := strconv.Atoi(s.Pop().(string))
	return r
}
