package str

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestBF(t *testing.T) {

	testCases := []struct {
		pat, text string
		index     int
	}{
		{"", "", 0},
		{"", "abc", 0},
		{"abc", "", -1},
		{"abc", "abc", 0},
		{"d", "abcdefg", 3},
		{"nan", "banana", 2},
		{"na", "banana", 2},
		{"pan", "anpanman", 2},
		{"nnaaman", "anpanmanam", -1},
		{"abcd", "abc", -1},
		{"abcd", "bcd", -1},
		{"bcd", "abcd", 1},
		{"abc", "acca", -1},
		{"aa", "aaa", 0},
		{"aaaaaa", "aaaaaaaaaa", 0},
		{"a", "aaaab", 0},
		{"baa", "aaaaa", -1},
		{"at that", "which finally halts.  at that point", 22},
	}

	//convey.Convey("测试字符串匹配", t, func() {
	//	for _, v := range testCases {
	//		s := BF(v.pat, v.text)
	//		convey.So(s, convey.ShouldEqual, v.index)
	//		fmt.Println(v.text, s)
	//	}
	//})

	//convey.Convey("测试KMP字符串匹配", t, func() {
	//	for _, v := range testCases {
	//		s := KMP(v.pat, v.text)
	//		convey.So(s, convey.ShouldEqual, v.index)
	//		fmt.Println(v.text)
	//	}
	//})

	convey.Convey("测试BM字符串匹配", t, func() {
		for _, v := range testCases {
			s := BM(v.pat, v.text)
			convey.So(s, convey.ShouldEqual, v.index)
			//convey.So(e, convey.ShouldEqual, v.ret2)
			fmt.Println(v.pat)
		}
	})

	//fmt.Println(StringFind("fdsa", "fff"))
}
