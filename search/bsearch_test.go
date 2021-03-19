package search

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestBSearch(t *testing.T) {
	cases := []struct {
		input  []int
		target int
		want   int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7}, 4, 3},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 5, 4},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 7, 6},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 1, 0},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 11, -1},
	}

	convey.Convey("测试二分查找", t, func() {
		for _, v := range cases {
			convey.So(BSearch(v.input, v.target), convey.ShouldEqual, v.want)
		}
	})
}

func TestBSearchFirstEq(t *testing.T) {
	cases := []struct {
		input  []int
		target int
		want   int
	}{
		{[]int{1, 2, 3, 3, 5, 6, 7}, 3, 2},
		{[]int{1, 2, 3, 4, 5, 7, 7}, 7, 5},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 7, 6},
		{[]int{1, 1, 3, 4, 5, 6, 7}, 1, 0},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 11, -1},
	}

	convey.Convey("测试二分查找", t, func() {
		for _, v := range cases {
			convey.So(BSearchFirstEq(v.input, v.target), convey.ShouldEqual, v.want)
		}
	})
}

func TestBSearchLastEq(t *testing.T) {
	cases := []struct {
		input  []int
		target int
		want   int
	}{
		{[]int{1, 2, 3, 3, 5, 6, 7}, 3, 3},
		{[]int{1, 2, 3, 4, 5, 7, 7}, 7, 6},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 7, 6},
		{[]int{1, 1, 3, 4, 5, 6, 7}, 1, 1},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 11, -1},
	}

	convey.Convey("测试二分查找", t, func() {
		for _, v := range cases {
			convey.So(BSearchLastEq(v.input, v.target), convey.ShouldEqual, v.want)
		}
	})
}
