package sort

func QuickSort(arr []int) {
	// 双闭区间，排序 0-(len-1)的元素
	qSort(arr, 0, len(arr)-1)
}

func qSort(arr []int, low, high int) {
	if low < high {
		// 枢轴(shuzhou) 选取的关键字（作为分隔元素的）位置，在它左边的比它小（或大），右边的比它大（或小）
		pivot := partition(arr, low, high)
		// 继续分别对 左区间和右区间 排序
		qSort(arr, low, pivot-1)
		qSort(arr, pivot+1, high)
	}
}

// 整理数据并返回 分割点、枢轴
func partition(arr []int, low, high int) int {
	// 选取第一个元素为 分割点
	pivotKey := arr[low]
	// low大于等于high，否则会陷入死循环
	for low < high {
		for low < high && arr[high] >= pivotKey {
			// 如果右边的元素大于 分割点元素，则 -- 比较下一个，直到遇到 小于的，则执行下一步，交换到左边
			high--
		}
		// 将小于 分割点元素的 放到左边
		arr[low], arr[high] = arr[high], arr[low]
		// 如果左边的元素 小于等于 分割点，则比较下一个，否则 交换到右边
		for low < high && arr[low] <= pivotKey {
			low++
		}
		arr[low], arr[high] = arr[high], arr[low]
	}
	// 返回的位置是：在它左边的都小于 pivotKey，右边都大于pivotKey
	return low
}
