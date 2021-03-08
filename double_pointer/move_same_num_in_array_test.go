package double_pointer

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestMoveSameNumInArray(t *testing.T) {
	cases := []struct {
		arr         []int
		input, want int
	}{
		{[]int{2, 3, 4, 5}, 2, 3},
		{[]int{2, 2, 4, 5}, 2, 2},
		{[]int{2, 3, 3, 5}, 3, 2},
		{[]int{2, 3, 3, 5}, 5, 3},
		{[]int{10, 7, 5, 12, 30, 3, 5}, 5, 5},
		{[]int{2, 2, 5}, 2, 1},
	}

	Convey("测试moveSameNumInArray", t, func() {
		for _, v := range cases {
			t, w := moveSameNumInArray(v.arr, v.input)
			So(t, ShouldEqual, v.want)
			fmt.Println(w)
		}
	})
}
