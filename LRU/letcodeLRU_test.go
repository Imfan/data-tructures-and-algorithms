package LRU

import (
. "github.com/smartystreets/goconvey/convey"
"strconv"
"testing"
)

func TestLRUCache(t *testing.T) {
	c := Constructor(3)
	cases := []struct{
		k, v string
	}{
		{"1", "a"},
		{"2", "b"},
		{"3", "c"},
		{"1", "aa"},
	}
	Convey("测试LRU", t, func() {
		i := 0
		for _, v := range cases {
			i ++
			Convey("Put"+strconv.Itoa(i), func() {
				c.Put(v.k, v.v)
			})
			Convey("Get("+strconv.Itoa(i)+")", func() {
				val := c.Get(v.k)
				So(val, ShouldEqual, v.v)
			})
		}
	})

	Convey("测试LRU超出时", t, func() {
		c.Get("2")
		c.Put("4", "d")

		So(c.Get("3"), ShouldEqual, "-1")
	})

}

