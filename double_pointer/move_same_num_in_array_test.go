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
		{[]int{2, 2, 5}, 2, 1},
	}

	Convey("测试moveSameNumInArray", t, func() {
		for _, v := range cases {
			So(moveSameNumInArray(v.arr, v.input), ShouldEqual, v.want)
			fmt.Println(v.arr[:v.want])
		}
	})
}
