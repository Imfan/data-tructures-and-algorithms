package sliding_window

// 给你一个整数数组 nums，有一个大小为k的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k个数字。滑动窗口每次只向右移动一位。
//
//返回滑动窗口中的最大值
func MaxSlidingWindow(nums []int, win int) []int {
	var que, ret []int
	push := func(i int) {
		for len(que) > 0 && nums[i] >= nums[que[len(que)-1]] {
			que = que[:len(que)-1]
		}
		que = append(que, i)
	}

	for k := 0; k < win; k++ {
		push(k)
	}
	ret = append(ret, nums[que[0]])

	for k := win; k < len(nums); k++ {
		if que[0]+win == k {
			que = que[1:]
		}
		push(k)
		ret = append(ret, nums[que[0]])
	}
	return ret
}
