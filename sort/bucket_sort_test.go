package sort

import (
	"github.com/smartystreets/goconvey/convey"
	"math/rand"
	"testing"
)

func TestBucketSort(t *testing.T) {
	for j := 0; j < 10; j++ {
		tCase := make([]int, end)
		for i := 0; i < end; i++ {
			tCase[i] = rand.Intn(randRange)
		}
		convey.Convey("测试桶排序", t, func() {
			BucketSort(tCase, 10*20*j+10)
			for k := range tCase {
				if k+1 < end && tCase[k] > tCase[k+1] {
					t.Fatalf("排序失败")
				}
			}
		})
	}
}

func bucketSort(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		tCase := make([]int, end)
		for i := 0; i < end; i++ {
			tCase[i] = rand.Intn(randRange)
		}
		BucketSort(tCase, n)
	}
}

//func BenchmarkBucketSort10(b *testing.B) {
//	bucketSort(b, 10)
//}
//
//func BenchmarkBucketSort50(b *testing.B) {
//	bucketSort(b, 50)
//}

//func BenchmarkBucketSort100(b *testing.B) {
//	bucketSort(b, 100)
//}

func BenchmarkFsBucketSort1000(b *testing.B) {
	bucketSort(b, 1000)
}

func BenchmarkFsBucketSort5000(b *testing.B) {
	bucketSort(b, 5000)
}

func BenchmarkFsBucketSort3000(b *testing.B) {
	bucketSort(b, 3000)
}

func BenchmarkFsBucketSort8000(b *testing.B) {
	bucketSort(b, 8000)
}

func BenchmarkFsBucketSort10000(b *testing.B) {
	bucketSort(b, 10000)

}
