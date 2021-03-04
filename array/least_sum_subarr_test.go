package array

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSubArrLenLeast(t *testing.T) {
	cases := []struct{
		input1 []int
		input2 int
		want int
	}{
		{[]int{1,2,3,4}, 1, 1},
		{[]int{1,2,3,4}, 4, 1},
		{[]int{1,2,3,4}, 5, 2},
		{[]int{1,2,3,4}, 9, 3},
	}
	Convey("测试 最小长度的数组", t, func() {
		for _, v := range cases {
			So(SubArrLenLeast(v.input1, v.input2), ShouldEqual, v.want)
		}
	})
}
