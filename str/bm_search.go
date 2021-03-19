package str

import (
	"strings"
)

type stringFinder struct {
	pattern        string
	badCharSkip    [256]int
	goodSuffixSkip []int
}

func BM(pattern, text string) int {
	return makeStringFinder(pattern).next(text)
}

// 创建 好字符 规则和 坏字符 规则,比较时是 从后向前 比较
func makeStringFinder(pattern string) *stringFinder {
	f := &stringFinder{
		pattern:        pattern,
		goodSuffixSkip: make([]int, len(pattern)),
	}
	// last 是pattern最后一个字符的索引
	last := len(pattern) - 1

	// 创建坏字符表，记录不匹配时T的i指针移动步数
	// 第一阶段，初始化256个字符全部移动 len(pattern) 步。以ASCII码作为键，比较时直接传主串不相等的字符ASCII码即可
	for i := range f.badCharSkip {
		f.badCharSkip[i] = len(pattern)
	}

	// 第二阶段：从左到右遍历pattern，更新其索引与P末尾的距离，结果就是该字符到末尾的最小距离
	// 没有计算last byte的距离, 因为移动至少要一步。 没有0步。
	for i := 0; i < last; i++ {
		// 后面相同的ASCII码会覆盖前面的，以防止 移动距离过大，错过
		f.badCharSkip[pattern[i]] = last - i
	}

	// 创建 好后缀表
	// 第一阶段: 此时pattern[i+1:]都是已经匹配的，且 好后缀只出现了一次
	// 计算T中的指针要移动的步数
	lastPrefix := last
	for i := last; i >= 0; i-- {
		// 如果有相同的前缀和后缀
		if strings.HasPrefix(pattern, pattern[i+1:]) {
			// 记录 每次 相同后缀的起点
			lastPrefix = i + 1
		}
		// 好后缀时T的指针移动分两步，首先移动到与 pattern的末尾对齐，即 last - i
		// lastPrefix 用来记录 pattern[i+1:]中所有后缀与同等长度的前缀相等时的最大索引
		// 然后移动 lastPrefix步
		// 先对齐
		tm := last - i
		// 然后移动 lastPrefix次， 就将相同的字符对齐了。相同后缀的起点到子串的头部的距离 跟  相同前缀的终点到子串的尾部距离 是一样的
		f.goodSuffixSkip[i] = lastPrefix + tm
	}
	// 第二阶段: 除去前缀，好后缀 在pattern前面还出现过, 如下计算相应的移动步数
	// 举例： "mississi" 中好后缀是issi, 在pattern[1]处出现过，移动步数为 last-i  +  lenSuffix
	for i := 0; i < last; i++ {
		// 前缀已经算过了，这里计算抛去第一个字符，部分子串与后缀相同的长度
		lenSuffix := longestCommonSuffix(pattern, pattern[1:i+1])
		// 略过 某个字符一直重复的  子串,还有尾部跟头部相同时，略过
		if pattern[i-lenSuffix] != pattern[last-lenSuffix] {
			// 这里的是从前往后数的，所以 这里的 lenSuffix是对齐 主串和pattern子串 尾部， last-i 是重复子串尾部到 pattern子串的尾部的距离
			f.goodSuffixSkip[last-lenSuffix] = lenSuffix + last - i
		}
	}
	return f
}

// 返回两个字符串的共同后缀的长度, 没有则为0
func longestCommonSuffix(a, b string) (i int) {
	for ; i < len(a) && i < len(b); i++ {
		if a[len(a)-1-i] != b[len(b)-1-i] {
			break
		}
	}
	return
}

// next 主要返回p在text里第一次匹配时的索引, 不匹配则返回-1
func (f *stringFinder) next(text string) int {
	// i 是T(即变量text)中要检查的字符索引, j为P中要检查的字符索引

	// 因从后向前比较, 所以i初始化为P的最后一位索引
	i := len(f.pattern) - 1
	for i < len(text) {
		// 每次比较时都从p的最后一位开始比较
		j := len(f.pattern) - 1
		for j >= 0 && i >= 0 && text[i] == f.pattern[j] {
			i--
			j--
		}
		// j为负数,说明匹配成功, 则直接返回 i+ 1
		if j < 0 {
			return i + 1
		}
		// j为非负, 表明text[i] != f.pattern[j], 则从坏字符表和好后缀表中获取分别获取i需要移动的步数, 取最大值并使移动到新位置
		i += max(f.badCharSkip[text[i]], f.goodSuffixSkip[j])
	}
	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
