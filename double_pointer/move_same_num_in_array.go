package double_pointer

func moveSameNumInArray(arr []int, target int) (int, []int) {
	var slowPointer, fasterPointer int

	for fasterPointer = 0; fasterPointer < len(arr); fasterPointer++ {
		if arr[fasterPointer] != target {
			arr[slowPointer] = arr[fasterPointer]
			slowPointer++
		}
	}
	return slowPointer, arr[:slowPointer]
}
