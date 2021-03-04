package link

func HashDetectCycle(head *List) int {
	seen := map[*List]int{}
	i := 0
	for head != nil {
		if _, ok := seen[head]; ok {
			return seen[head]
		}
		seen[head] = i
		head = head.next
		i++
	}
	return -1
}

func DetectCycle(l *List) int {
	var slowPointer, fasterPointer *List
	slowPointer, fasterPointer = l, l
	for fasterPointer != nil {
		slowPointer = slowPointer.next
		fasterPointer = fasterPointer.next.next
		if slowPointer == fasterPointer {
			index1 := l
			index2 := slowPointer
			i := 0
			for index1 != index2 { // 这里都是 next
				index1 = index1.next
				index2 = index2.next
				i++
			}
			// 返回第几个节点是 环入口
			return i
		}
	}
	return -1
}
