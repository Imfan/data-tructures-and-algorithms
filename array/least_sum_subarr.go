package array

import "fmt"

// 长度最小的子数组
//给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的 连续 子数组，并返回其长度。如果不存在符合条件的子数组，返回 0。
const (
	MaxInt = int(^uint(0) >> 1)
	MinInt = -MaxInt - 1
)

type allCase struct {
	start, end int
}

func SubArrLenLeast(arr []int, target int) int {
	var sum int
	var all []allCase
	LeastLen := MaxInt
	slowPointer := 0
	length := len(arr) - 1
	for fastPointer := 0; fastPointer <= length; fastPointer++ {
		sum += arr[fastPointer]
		for sum >= target {
			tmpLen := fastPointer - slowPointer + 1
			if LeastLen > tmpLen && tmpLen != 0 {
				LeastLen = tmpLen
			}
			all = append(all, allCase{slowPointer, fastPointer + 1})
			sum -= arr[slowPointer]
			slowPointer++
		}
	}
	fmt.Println(all)
	return LeastLen
}
