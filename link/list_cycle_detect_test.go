package link

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestDetectCycle(t *testing.T) {
	cases := []struct {
		input []int
		want  int
	}{
		{[]int{2, 3, 4, 5, 6, 7, 8, 9}, 3},
		{[]int{2, 3, 4, 3, 6, 7, 8, 9}, 0},
		{[]int{2, 3, 4, 3, 6, 7, 8, 9}, 6},
	}

	Convey("测试 链表环的检测", t, func() {
		for _, slice := range cases {
			l := Init(0)
			// 虚拟头部
			head := l
			var tmp *List
			for k, v := range slice.input {
				l = l.AddToTail(v)
				if k == slice.want {
					tmp = l
				}
			}
			l.next = tmp
			So(DetectCycle(head.next), ShouldEqual, slice.want)
			So(HashDetectCycle(head.next), ShouldEqual, slice.want)
		}
	})
}
