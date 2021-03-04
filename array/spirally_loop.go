package array

// 螺旋矩阵II
//给定一个正整数 n，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的正方形矩阵。
//
//示例:
//
//输入: 3 输出: [ [ 1, 2, 3 ], [ 8, 9, 4 ], [ 7, 6, 5 ] ]
func SpirallyLoop(times int) [][]int {
	loop := int(times / 2)
	var x, y, i, startx, starty int
	offset := 1

	retArr := make([][]int, times)
	for k, _ := range retArr {
		retArr[k] = make([]int, times)
	}
	i = 1

	for ; loop > 0; loop-- {
		x, y = startx, starty
		// 向右
		for ; x < startx+times-offset; x++ {
			retArr[y][x] = i
			i++
		}

		//向下
		for ; y < starty+times-offset; y++ {
			retArr[y][x] = i
			i++
		}

		// 向左
		for ; x > startx; x-- {
			retArr[y][x] = i
			i++
		}
		//向上
		for ; y > starty; y-- {
			retArr[y][x] = i
			i++
		}
		startx++
		starty++
		offset += 2
	}
	if times%2 != 0 {
		middle := int(times / 2)
		retArr[middle][middle] = i
	}
	return retArr
}
