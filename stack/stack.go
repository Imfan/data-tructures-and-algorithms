package stack

type StackByArray struct {
	stack     []interface{}
	base, top int
}

func Init() *StackByArray {
	return &StackByArray{
		stack: []interface{}{},
	}
}

func (s *StackByArray) Push(val interface{}) {
	s.stack = append(s.stack, val)
	s.top++
}

func (s *StackByArray) Pop() interface{} {
	if s.top-1 < 0 {
		return false
	}
	s.top--
	ret := s.stack[s.top]
	//fmt.Println(ret)
	s.stack = s.stack[:s.top]
	return ret
}
