package sort

import (
	"github.com/smartystreets/goconvey/convey"
	"math/rand"
	"testing"
)

func TestSysSort(t *testing.T) {
	for j := 0; j < 10; j++ {
		tCase := make([]int, end)
		for i := 0; i < end; i++ {
			tCase[i] = rand.Intn(randRange)
		}
		convey.Convey("测试system_sort排序", t, func() {
			SysSort(tCase)
			for k := range tCase {
				if k+1 < end && tCase[k] > tCase[k+1] {
					t.Fatalf("排序失败")
				}
			}
		})
	}
}

func BenchmarkSysSort(b *testing.B) {
	for j := 0; j < b.N; j++ {
		tCase := make([]int, end)
		for i := 0; i < end; i++ {
			tCase[i] = rand.Intn(randRange)
		}
		SysSort(tCase)
	}
}
