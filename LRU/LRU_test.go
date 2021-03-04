package LRU

import (
	. "github.com/smartystreets/goconvey/convey"
	"strconv"
	"testing"
)

func TestCache(t *testing.T) {
	c := NewCache(3)
	cases := []struct{
		k, v string
		st bool
	}{
		{"1", "a", true},
		{"2", "b", true},
		{"3", "c", true},
		{"1", "aa", true},
	}
	Convey("测试LRU", t, func() {
		i := 0
		for _, v := range cases {
			i ++
			Convey("Set"+strconv.Itoa(i), func() {
				r, b := c.Set(v.k, v.v)
				So(r, ShouldEqual, v.v)
				So(b, ShouldEqual, v.st)
			})
			Convey("Get("+strconv.Itoa(i)+")", func() {
				val, b := c.Get(v.k)
				So(val, ShouldEqual, v.v)
				So(b, ShouldEqual, v.st)
			})
		}
	})

	Convey("测试LRU超出时", t, func() {
		c.Get("2")
		r, st := c.Set("4", "d")
		So(r, ShouldEqual, "d")
		So(st, ShouldEqual, true)
		r, st = c.Get("3")
		So(r, ShouldEqual, "")
		So(st, ShouldEqual, false)
	})

	Convey("测试LRU.Del",t, func() {
		for _, v := range cases {
			So(c.Delete(v.k), ShouldEqual, v.st)
			_, st := c.Get(v.k)
			So(st, ShouldEqual, false)
		}
	})
}
