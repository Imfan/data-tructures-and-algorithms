package stack

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestIsValid(t *testing.T) {
	cases := []struct {
		input string
		want  bool
	}{
		{"(中文){}{{}}", true},
		{"({}{{}})", true},
		{"[({}{{}})]", true},
	}

	convey.Convey("测试 符号是否对称", t, func() {
		for _, v := range cases {
			convey.So(IsValid(v.input), convey.ShouldEqual, v.want)
		}
	})
}
