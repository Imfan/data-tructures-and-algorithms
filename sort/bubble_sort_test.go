package sort

import (
	"github.com/smartystreets/goconvey/convey"
	"math/rand"
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	cases := []struct {
		input []int
		want  []int
	}{
		{[]int{4, 1, 3, 2}, []int{1, 2, 3, 4}},
		{[]int{1, 1, 3, 2}, []int{1, 1, 2, 3}},
	}

	convey.Convey("测试冒泡排序", t, func() {
		for _, v := range cases {
			BubbleSort(v.input)
			if !reflect.DeepEqual(v.input, v.want) {
				t.Fatalf("不相等")
			}
		}
	})

	end := 10000
	input := []int{}
	for i := 0; i < end; i++ {
		input = append(input, rand.Int())
	}

	convey.Convey("测试冒泡排序 10000个数字", t, func() {
		BubbleSort(input)
		for k, v := range input {
			if k+1 < end && input[k+1] < v {
				t.Fatalf("排序失败")
			}
		}
	})

}

func BenchmarkBubbleSort(b *testing.B) {
	end := 10000
	input := []int{}
	for i := 0; i < end; i++ {
		input = append(input, rand.Int())
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		BubbleSort(input)
	}
}
