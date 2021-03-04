package stack

func IsValid(str string) bool {
	leftSymbol := map[rune]rune{'(': ')', '{': '}', '<': '>', '[': ']'}
	rightSymbol := map[rune]rune{')': '(', '}': '{', '>': '<', ']': '['}
	stack := Init()
	for _, v := range str {
		if _, ok := leftSymbol[v]; ok {
			stack.Push(v)
		}
		if k, ok := rightSymbol[v]; ok {
			if stack.Pop() != k {
				return false
			}
		}
	}
	if stack.Pop() == false {
		return true
	}
	return false
}
