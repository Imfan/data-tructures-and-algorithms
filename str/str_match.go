package str

// 朴素算法
// 返回第一个匹配的字符串 起始位置，左闭右开
func BF(str, matchStr string) (int, int) {
	for i := 0; i+len(matchStr) <= len(str); i++ {
		tmpS := str[i:]
		j := 0
		jj := 0
		//fmt.Println(jj, string(matchStr[0]), matchStr)
		//fmt.Println(matchStr[jj])
		for jj < len(matchStr) && tmpS[j] == matchStr[jj] {
			j++
			jj++
		}
		//fmt.Println(jj)
		// 如果全匹配 直接返回
		if jj == len(matchStr) {
			return i, jj + i
		}
	}
	return 0, 0
}

//KMP算法
func KMP(str, matchStr string) (int, int) {
	// 匹配表
	next := getNext(matchStr)
	j := 0
	// 主串索引
	i := 0
	// 如果剩余未比较的， 超过了 总串就没比较的必要了
	for i+len(matchStr)-j <= len(str) && j < len(matchStr) {
		// 如果相等比较下一个字符串
		if j == -1 || str[i] == matchStr[j] {
			// 主串的索引 一直是增加的
			i++
			j++
		} else {
			// 只有不相等的时候 j才会回溯
			j = next[j]
		}

		if j == len(matchStr) {
			// 返回起始位置，左闭右开
			return i - j, i
		}
	}
	return 0, 0
}

// 求next数组的过程完全可以看成字符串匹配的过程，即以模式字符串为主字符串，以模式字符串的前缀为目标字符串，一旦字符串匹配成功，那么当前的next值就是匹配成功的字符串的长度。
// 获取字符串匹配表
func getNext(strM string) []int {
	match := make([]int, len(strM))
	// 是为了判断 j又回到起点了
	match[0] = -1
	// 被匹配的字符串的 指针
	i := 0
	// 匹配字符串的指针
	j := -1
	for i < len(strM) {
		if j == -1 || strM[i] == strM[j] {
			// 这个i++不能放到 if条件前面，那样判断逻辑就不对了
			// 是为了将 子串的 PMT值 放到 后面一位，方便 总串查找的时候调用
			i++
			j++ // PMT值，又多了一个相等的，所以+1
			// 这里如果不判断会 越界
			if i < len(match) {
				match[i] = j
			}
		} else { // 若不相等，则j将回溯
			// 回溯到 第一次出现某个前缀的位置，或者 头部
			j = match[j]
		}
	}
	//fmt.Println(match)
	return match
}

func BM(strM string) (int, int) {

	return 0, 0
}
