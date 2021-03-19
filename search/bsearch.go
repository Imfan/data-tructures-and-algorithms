package search

func BSearch(arr []int, target int) int {
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if arr[mid] > target {
			high = mid - 1
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

// 查找第一个等于指定值的 位置
func BSearchFirstEq(arr []int, target int) int {
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if arr[mid] > target {
			high = mid - 1
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			if (mid == 0) || (arr[mid-1] != target) {
				return mid
			} else {
				high = mid - 1
			}
		}
	}
	return -1
}

func BSearchLastEq(arr []int, target int) int {
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if target > arr[mid] {
			low = mid + 1
		} else if target < arr[mid] {
			high = mid - 1
		} else {
			if mid == len(arr)-1 || arr[mid+1] != target {
				return mid
			} else {
				low = mid + 1
			}
		}
	}
	return -1
}
