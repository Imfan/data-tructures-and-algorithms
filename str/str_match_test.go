package str

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"math/rand"
	"testing"
)

const letter = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randStr(n int) string {
	s := make([]byte, n)
	for i := range s {
		s[i] = letter[rand.Intn(n)]
	}
	return string(s)
}
func TestBF(t *testing.T) {

	cases := []struct {
		input, input2 string
		ret1, ret2    int
	}{
		{"htewuodasfhfsda", "te", 1, 3},
		{"htewuodasfhfsda", "htew", 0, 4},
		{"htewuodasfhfsda", "fsda", 11, 15},
		{"htewuodasfhfsda", "fsd", 11, 14},
		{"htewuodasfhfsda", "fssfsfddddddddsdafssdaaada", 0, 0},
		{"htewuodasfhfsda", " dsdadaaada", 0, 0},
		{"abababccccc", " abababca", 0, 0},
		{"abababcabababcdfs", "abababcd", 7, 15},
		{"abababcabababcdfs", "aaaaa", 0, 0},
	}

	convey.Convey("测试字符串匹配", t, func() {
		for _, v := range cases {
			s, e := BF(v.input, v.input2)
			convey.So(s, convey.ShouldEqual, v.ret1)
			convey.So(e, convey.ShouldEqual, v.ret2)
			fmt.Println(v.input2, s, e)
		}
	})

	convey.Convey("测试字符串匹配", t, func() {
		for _, v := range cases {
			s, e := KMP(v.input, v.input2)
			convey.So(s, convey.ShouldEqual, v.ret1)
			convey.So(e, convey.ShouldEqual, v.ret2)
			fmt.Println(v.input2)
		}
	})

	fmt.Println(StringFind("fdsa", "fff"))
}
