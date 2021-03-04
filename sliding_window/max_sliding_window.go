package sliding_window

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
