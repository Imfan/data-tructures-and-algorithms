package sort

import (
	"sort"
)

type tmpArr []int

func SysSort(arr tmpArr) {
	sort.Sort(arr)
}

func (arr tmpArr) Len() int {
	return len(arr)
}

func (arr tmpArr) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func (arr tmpArr) Less(i, j int) bool {
	return arr[i] < arr[j]
}
