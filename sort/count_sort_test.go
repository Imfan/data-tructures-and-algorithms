package sort

import (
	"github.com/smartystreets/goconvey/convey"
	"math/rand"
	"testing"
)

func TestCountingSort(t *testing.T) {
	for j := 0; j < 10; j++ {
		tCase := make([]int, end)
		for i := 0; i < end; i++ {
			tCase[i] = rand.Intn(randRange)
		}
		convey.Convey("测试计数排序", t, func() {
			CountingSort(tCase)
			for k := range tCase {
				if k+1 < end && tCase[k] > tCase[k+1] {
					t.Fatalf("排序失败")
				}
			}
		})
	}
}

//func TestCountingSort2(t *testing.T) {
//	for j := 0; j < 10; j++ {
//		tCase := make([]int, end)
//		for i := 0; i < end; i++ {
//			tCase[i] = rand.Intn(randRange)
//		}
//		convey.Convey("测试计数排序", t, func() {
//			tCase = CountingSort2(tCase)
//			for k := range tCase {
//				if k+1 < end && tCase[k] > tCase[k+1] {
//					t.Fatalf("排序失败")
//				}
//			}
//		})
//	}
//}

func BenchmarkFsCountingSort(b *testing.B) {
	for j := 0; j < b.N; j++ {
		tCase := make([]int, end)
		for i := 0; i < end; i++ {
			tCase[i] = rand.Intn(randRange)
		}
		CountingSort(tCase)
		//for k := range tCase {
		//	if k+1 < end && tCase[k] > tCase[k+1] {
		//		panic("排序失败")
		//	}
		//}
	}
}

//func BenchmarkFsCountingSort2(b *testing.B) {
//	for j := 0; j < b.N; j++ {
//		tCase := make([]int, end)
//		for i := 0; i < end; i++ {
//			tCase[i] = rand.Intn(randRange)
//		}
//		tCase = CountingSort2(tCase)
//		//for k := range tCase {
//		//	if k+1 < end && tCase[k] > tCase[k+1] {
//		//		panic("排序失败")
//		//	}
//		//}
//	}
//}
