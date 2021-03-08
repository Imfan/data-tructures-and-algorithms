package queue

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestQueueByArr_Pop(t *testing.T) {
	cases := []struct {
		input []int
		cap   int
	}{
		{[]int{1, 2, 3, 4, 5, 6}, 2},
	}

	convey.Convey("测试队列ByArr", t, func() {
		for _, v := range cases {
			que := Init(v.cap)
			for v1 := range v.input {
				que.Push(v1)
			}
			i := 0
			for {
				v2, ok := que.Pop()
				if !ok {
					break
				}
				convey.So(v2, convey.ShouldEqual, v.input[i])
			}

		}
	})
}
