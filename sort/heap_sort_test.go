package sort

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestInit(t *testing.T) {
	h := &IntHeap{2, 1, 5, 11}
	Init(h)
	Push(h, 3)
	Push(h, 3)
	Push(h, 3)
	Push(h, 3)

	want := []int{1, 2, 3, 3, 3, 3, 5, 11}

	fmt.Printf("minimum: %d\n", (*h)[0])
	i := 0
	convey.Convey("测试小顶堆", t, func() {
		for h.Len() > 0 {
			convey.So(Pop(h), convey.ShouldEqual, want[i])
			i++
		}
	})

}
