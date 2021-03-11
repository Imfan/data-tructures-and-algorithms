package sort

// 计数排序
func CountingSort(arr []int) {
	maxVal := 0
	// 寻找最大值和最小值
	for _, v := range arr {
		if v < 0 {
			panic("不能小于0")
		}
		if maxVal < v {
			maxVal = v
		}
	}
	bucketLen := maxVal + 1
	bucket := make([]int, bucketLen) // 初始化桶的数量
	length := len(arr)
	// 把数据放进桶内
	for i := 0; i < length; i++ {
		// 以值为 bucket的键，每多一个重复的就加1
		bucket[arr[i]] += 1
	}

	sortedIndex := 0
	// 目测复杂度为O(n)
	// 从 0桶开始遍历，
	for j := 0; j < bucketLen; j++ {
		// 如果有数据就放到 数组中，直到这个桶没数据了
		for bucket[j] > 0 {
			// j就是原始arr里面存的值
			arr[sortedIndex] = j
			// 把数据依次往后 一个一个放下
			sortedIndex++
			bucket[j]--
		}
	}
}

//func CountingSort2(arr []int) []int {
//	maxVal := 0
//	// 寻找最大值
//	for _, v := range arr {
//		if v < 0 {
//			panic("不能小于0")
//		}
//		if maxVal < v {
//			maxVal = v
//		}
//	}
//	bucketLen := maxVal + 1
//	// 桶的长度
//	bucket := make([]int, bucketLen) // 初始为0的数组
//	// 数组长度
//	length := len(arr)
//	for i := 0; i < length; i++ {
//		bucket[arr[i]] += 1
//	}
//	for i := 1; i < bucketLen; i++ {
//		bucket[i] = bucket[i-1] + bucket[i]
//	}
//	ret := make([]int, length) // 初始 跟 原数组长度相同的数组
//
//	for j := length - 1; j >= 0; j-- {
//		bucketIndex := arr[j]
//		ret[bucket[bucketIndex]-1] = bucketIndex
//		bucket[bucketIndex] -= 1
//	}
//	return ret
//}
