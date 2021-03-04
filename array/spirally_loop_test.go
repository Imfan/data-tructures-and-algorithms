package array

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSpirallyLoop(t *testing.T) {
	cases := []struct {
		input, x, y, want int
	}{
		{2, 0, 1, 4},
		{3, 1, 1, 9},
		{5, 2, 2, 25},
		{5, 4, 4, 9},
		{5, 0, 4, 13},
		{15, 14, 14, 29},
		{15, 7, 7, 15 * 15},
	}

	Convey("测试 螺旋矩阵", t, func() {
		for _, v := range cases {
			So(SpirallyLoop(v.input)[v.y][v.x], ShouldEqual, v.want)
		}
	})
}
