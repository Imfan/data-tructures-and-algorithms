package sort

import (
	"github.com/smartystreets/goconvey/convey"
	"math/rand"
	"testing"
)

func TestMergeSort(t *testing.T) {
	for j := 0; j < 10; j++ {
		tCase := make([]int, end)
		for i := 0; i < end; i++ {
			tCase[i] = rand.Intn(randRange)
		}
		convey.Convey("测试插入排序", t, func() {
			MergeSort(tCase)
			for k := range tCase {
				if k+1 < end && tCase[k] > tCase[k+1] {
					t.Fatalf("排序失败")
				}
			}
		})
	}
}

func TestIterativeMergeSort(t *testing.T) {
	for j := 0; j < 10; j++ {
		tCase := make([]int, end)
		for i := 0; i < end; i++ {
			tCase[i] = rand.Intn(randRange)
		}
		convey.Convey("测试插入排序", t, func() {
			IterativeMergeSort(tCase)
			for k := range tCase {
				if k+1 < end && tCase[k] > tCase[k+1] {
					t.Fatalf("排序失败")
				}
			}
		})
	}
}

func BenchmarkMergeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tCase := make([]int, end)
		for i := 0; i < end; i++ {
			tCase[i] = rand.Intn(randRange)
		}
		MergeSort(tCase)
		//for k := range tCase {
		//	if k+1 < end && tCase[k] > tCase[k+1] {
		//		panic("排序失败")
		//	}
		//}
	}
}

//func BenchmarkIterativeMergeSort(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		tCase := make([]int, end)
//		for i := 0; i < end; i++ {
//			tCase[i] = rand.Intn(randRange)
//		}
//		IterativeMergeSort(tCase)
//		//for k := range tCase {
//		//	if k+1 < end && tCase[k] > tCase[k+1] {
//		//		panic("排序失败")
//		//	}
//		//}
//	}
//}
