package dichotomy

// 题目：编号35：搜索插入位置
//给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置

/**
这道题目，要在数组中插入目标值，无非是这四种情况。
目标值在数组所有元素之前
目标值等于数组中某一个元素
目标值插入数组中的位置
目标值在数组所有元素之后
 */
func SearchOrInsertPosition(data []int, target int) int {
	length := len(data)
	left := 0
	right := length - 1 // 假设target在左闭右闭的区间里，[left, right]
	for left <= right { // 当left==right，区间[left, right]依然有效
		middle := (int)((left + right) / 2)
		if target > data[middle] {
			left = middle + 1 // target 在右区间，所以[middle + 1, right]
		}else if target < data[middle] {
			right = middle - 1 // target 在左区间，所以[left, middle - 1]
		}else {
			return middle // 目标值等于数组中某一个元素
		}
	}

	// 分别处理如下四种情况
	// 目标值在数组所有元素之前  return right + 1
	// 目标值等于数组中某一个元素  return middle;
	// 目标值插入数组中的位置 [left, right]，return  right + 1
	// 目标值在数组所有元素之后的情况 [left, right]， return right + 1
	return right + 1
}
