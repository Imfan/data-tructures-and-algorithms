package sort

func BubbleSort(arr []int) {
	for k, _ := range arr {
		k++
		isSwap := false
		length := len(arr)
		//fmt.Println(arr)
		for i := 0; i < length-k; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				isSwap = true
			}
		}
		if !isSwap {
			break
		}
	}
}
