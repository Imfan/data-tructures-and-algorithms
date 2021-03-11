package sort

import (
	"github.com/smartystreets/goconvey/convey"
	"math/rand"
	"testing"
)

func TestRadixSort(t *testing.T) {
	for j := 0; j < 10; j++ {
		tCase := make([]int, end)
		for i := 0; i < end; i++ {
			tCase[i] = rand.Intn(randRange)
		}
		convey.Convey("测试插入排序", t, func() {
			RadixSort(tCase)
			for k := range tCase {
				if k+1 < end && tCase[k] > tCase[k+1] {
					t.Fatalf("排序失败")
				}
			}
		})
	}
}

func BenchmarkRadixSort(b *testing.B) {
	for j := 0; j < b.N; j++ {
		tCase := make([]int, end)
		for i := 0; i < end; i++ {
			tCase[i] = rand.Intn(randRange)
		}
		RadixSort(tCase)
	}
}
