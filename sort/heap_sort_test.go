package sort

import (
	"github.com/smartystreets/goconvey/convey"
	"math/rand"
	"testing"
)

func TestInit(t *testing.T) {
	he := &IntHeap{}
	// 初始化 堆
	Init(he)
	for i := 0; i < end; i++ {
		// 添加元素
		Push(he, rand.Intn(20000))
	}
	convey.Convey("测试小顶堆", t, func() {
		ret := make([]int, end)
		for i := 0; i < end; i++ {
			// 弹出堆顶元素
			tmp := Pop(he)
			ret[i] = tmp.(int)
		}
		// 验证排序是否成功
		for k, v := range ret {
			if k+1 < end && ret[k+1] < v {
				t.Fatalf("排序失败")
			}
		}
	})
}

func BenchmarkHeapSort(b *testing.B) {
	for j := 0; j < b.N; j++ {
		b.StopTimer()
		he := &IntHeap{}
		Init(he)
		for i := 0; i < end; i++ {
			Push(he, rand.Intn(randRange))
		}

		ret := make([]int, end)

		b.StartTimer()
		for i := 0; i < end; i++ {
			ret[i] = Pop(he).(int)
		}
		// 验证排序是否成功
		//for k, v := range ret {
		//	if k+1 < end && ret[k+1] < v {
		//		panic("排序失败")
		//	}
		//}
	}
}
