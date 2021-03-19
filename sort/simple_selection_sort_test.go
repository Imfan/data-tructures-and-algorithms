package sort

import (
	"github.com/smartystreets/goconvey/convey"
	"math/rand"
	"testing"
)

func TestSelectSort(t *testing.T) {
	end := 10000
	input := []int{}
	for i := 0; i < end; i++ {
		input = append(input, rand.Int())
	}

	convey.Convey("测试选择排序 10000个数字", t, func() {
		SelectSort(input)
		for k, v := range input {
			if k+1 < end && input[k+1] < v {
				t.Fatalf("排序失败")
			}
		}
	})

}

func BenchmarkSLSelectSort(b *testing.B) {
	for j := 0; j < b.N; j++ {
		tCase := make([]int, end)
		for i := 0; i < end; i++ {
			tCase[i] = rand.Intn(randRange)
		}
		SelectSort(tCase)
	}
}
