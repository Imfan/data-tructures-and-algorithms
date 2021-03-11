package sort

func MergeSort(arr []int) {
	ret := make([]int, len(arr))
	// 双闭 区间
	subMergeSort(arr, ret, 0, len(arr)-1)
}

// 递归 归并排序
func subMergeSort(arr []int, ret []int, l, r int) {
	if l >= r {
		return
	}
	c := (l + r) / 2
	// 左闭右闭
	subMergeSort(arr, ret, l, c)
	// 如果不加1会导致，死循环 左闭右闭
	subMergeSort(arr, ret, c+1, r)
	merge(arr, ret, l, c, r)
}

func IterativeMergeSort(arr []int) {
	k := 1
	length := len(arr)
	ret := make([]int, length)

	for k < length {
		subIterativeMerge(arr, ret, k, length-1)
		k = k * 2
		subIterativeMerge(arr, ret, k, length-1)
		k = k * 2
	}
}

// 各个子数组排序
func subIterativeMerge(arr, tmp []int, subLength, right int) {
	i := 1
	//for i <= right-2*subLength+1 {
	// 保证至少有两组
	for i <= right-2*subLength {
		merge(arr, tmp, i, i+subLength-1, i+2*subLength-1)
		i = i + 2*subLength
	}
	// 如果剩下两组，其中有一组是单数的
	if i < right-subLength+1 {
		// 合并剩余的两组
		merge(arr, tmp, i, i+subLength-1, right)
	} else { //只剩下一组或者一个，一定 有序的
		for j := 0; j <= right; j++ {
			arr[j] = tmp[j]
		}
	}
}

func merge(arr, ret []int, l, c, r int) {
	// 总的左右边界，双闭区间
	left, right := l, r
	// 右区间右边界
	rightRight := r
	// 右区间左边界  双闭区间
	rightLeft := c + 1
	// 左区间 右边界
	leftRight := c
	// 左区间左边界 位置
	leftLeft := l
	// 临时数组的 指针位置
	tempPointer := l
	for {
		// 左区间左边界 位置超过了 中间值 或者 右区间左边界位置超过了 整体右边界 则退出
		if rightLeft > rightRight || leftLeft > leftRight {
			break
		}
		// 依次比较。如果 左区间最小值 小于 右区间最小值，则 放到临时数组，最终是 从大到小顺序
		if smallToLargeOrder(arr[leftLeft], arr[rightLeft]) {
			ret[tempPointer] = arr[leftLeft]
			leftLeft++
		} else {
			ret[tempPointer] = arr[rightLeft]
			rightLeft++
		}
		tempPointer++
	}
	// 如果 左边还有剩余的 未放入 整合后的数组，则全部依次放入，因为 理论上 是本次 比较区间里面最大的了，直接顺序放下即可
	if leftLeft <= leftRight {
		for i := 0; i <= leftRight-leftLeft; i++ {
			ret[tempPointer+i] = arr[leftLeft+i]
		}
	}

	if rightLeft <= rightRight {
		for i := 0; i <= rightRight-rightLeft; i++ {
			ret[tempPointer+i] = arr[rightLeft+i]
		}
	}

	for ; left <= right; left++ {
		arr[left] = ret[left]
	}
	//fmt.Println(arr)

}

func smallToLargeOrder(left, right int) bool {
	return left < right && smallToLarge
}

func largeToSmallOrder(left, right int) bool {
	return left > right
}
