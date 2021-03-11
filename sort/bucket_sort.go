package sort

// 桶排序
func BucketSort(numArr []int, bucketSize int) {
	if len(numArr) <= 100 {
		InsertSort(numArr)
		return
	}
	minValue := numArr[0]
	maxValue := numArr[0]
	// 确定最大值最小值
	for i := 0; i < len(numArr); i++ {
		if minValue > numArr[i] {
			minValue = numArr[i]
		}
		if maxValue < numArr[i] {
			maxValue = numArr[i]
		}
	}
	// 桶切片初始化
	buckets := make([][]int, (maxValue-minValue)/bucketSize+1)
	// 数据入桶
	for i := 0; i < len(numArr); i++ {
		// 必须用 /，如果用取余 不能保证顺序性
		buckets[(numArr[i]-minValue)/bucketSize] = append(buckets[(numArr[i]-minValue)/bucketSize], numArr[i])
	}
	//start := 0
	key := 0
	// 对每个桶进行排序
	for _, bucket := range buckets {
		if len(bucket) <= 0 {
			continue
		}
		QuickSort(bucket)
		//copy(numArr[start:], bucket)
		//start += len(bucket)
		for _, value := range bucket {
			numArr[key] = value
			key = key + 1
		}
	}

}
