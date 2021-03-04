package queue

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestEvalPRN(t *testing.T) {
	cases := []struct {
		input []string
		want  int
	}{
		{[]string{"2", "1", "+", "3", "*"}, 9},
		{[]string{"4", "13", "5", "/", "+"}, 6},
	}

	convey.Convey("测试 逆波兰表达式", t, func() {
		for _, v := range cases {
			convey.So(evalPRN(v.input), convey.ShouldEqual, v.want)
		}
	})

}
