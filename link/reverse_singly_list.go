package link

// 反转一个单链表。
//
//示例: 输入: 1->2->3->4->5->NULL
//输出: 5->4->3->2->1->NULL
type SinglyList struct {
	next  *SinglyList
	value int
}

func ReverseSinglyList(list *SinglyList) *SinglyList {
	var slowPointer, fastPointer *SinglyList
	fastPointer = list
	for fastPointer != nil {
		// 存储fastPointer下一个位置，之后会改变next指针
		tmp := fastPointer.next
		// 改变next指针，反转
		fastPointer.next = slowPointer
		// slowPointer 向前一步
		slowPointer = fastPointer
		// fasterPinter 向前一步
		fastPointer = tmp
	}
	return slowPointer
}
