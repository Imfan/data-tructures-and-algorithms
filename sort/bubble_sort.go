package sort

// 冒泡排序
func BubbleSort(arr []int) {
	for k, _ := range arr {
		isSwap := false
		length := len(arr)
		//fmt.Println(arr)
		for i := 1; i < length-k; i++ {
			if arr[i-1] > arr[i] {
				arr[i-1], arr[i] = arr[i], arr[i-1]
				isSwap = true
			}
		}
		if !isSwap {
			break
		}
	}
}

func NoSkipBubbleSort(arr []int) {
	for k, _ := range arr {
		length := len(arr)
		//fmt.Println(arr)
		for i := 1; i < length-k; i++ {
			if arr[i-1] > arr[i] {
				arr[i-1], arr[i] = arr[i], arr[i-1]
			}
		}

	}
}
