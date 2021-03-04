package dichotomy

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test2(t *testing.T)  {
	fmt.Println("I'm Test2")
}

func TestSearchOrInsertPosition(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Error(err)
		}
	}()
	data := []int{1,2,3,5,7,9}
	cases := []struct {
		Name string
		input, want int
	}{
		{"超过左边", -1,  0},
		{Name: "和某个相等", input: 1, want: 0},
		{Name: "插入其中某个", input: 4, want: 3},
		{Name: "超过右边", input: 40, want: 6},
	}

	for  _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			got := SearchOrInsertPosition(data, c.input)
			if got != c.want {
				t.Errorf("excepted %v, got %v", c.want, got)
			}
		})

	}
}

func TestSearchOrInsertPosition2(t *testing.T) {
	data := []int{1,2,3,5,7,9}
	cases := []struct {
		Name string
		input, want int
	}{
		{"超过左边", -1,  0},
		{Name: "和某个相等", input: 1, want: 0},
		{Name: "插入其中某个", input: 4, want: 3},
		{Name: "超过右边", input: 40, want: 6},
	}
	Convey("测试search_insert", t, func() {
		for _, c := range cases {
			Convey(c.Name, func() {
				So(SearchOrInsertPosition(data, c.input), ShouldEqual, c.want)
			})
		}
	})
}
