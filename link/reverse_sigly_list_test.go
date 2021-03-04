package link

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestReverseSinglyList(t *testing.T) {
	cases := []struct {
		input, want []int
	}{
		{[]int{1, 2, 3, 4}, []int{4, 3, 2, 1}},
		{[]int{7, 3, 0, 1, 5, 7}, []int{7, 5, 1, 0, 3, 7}},
	}

	Convey("测试单链表反转", t, func() {

		for _, slice := range cases {
			lists := &SinglyList{}
			tail := lists
			head := lists
			for _, v := range slice.input {
				lists = tail
				lists.next = &SinglyList{
					value: v,
					next:  nil,
				}
				tail = lists.next
			}
			i := 0
			head = head.next
			head = ReverseSinglyList(head)
			for head.next != nil {
				So(head.value, ShouldEqual, slice.want[i])
				i++
				head = head.next
			}

		}
	})

}
