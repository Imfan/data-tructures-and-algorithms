package str

import "strings"

// bytes/bytes.go

func contains() {
	// 使用的是rk算法
	strings.Contains("sssssss", "s")
}

//
//// 选择非常大的一个质数16777619 作为 base
//const primeRK = 16777619
//
//// hashStr 返回子串的hash值和乘数因子
//func hashStr(sep string) (uint32, uint32) {
//	hash := uint32(0)
//	for i := 0; i < len(sep); i++ {
//		hash = hash*primeRK + uint32(sep[i])  //计算hash值
//	}
//	// 计算(最高位 + 1)位的乘数因子, 使用位移, 没有使用 i--, 可以有效减少循环次数. i >>=1 相当于遍历二进制的每一位
//	var pow, sq uint32 = 1, primeRK
//	for i := len(sep); i > 0; i >>= 1 {
//		if i&1 != 0 {
//			pow *= sq
//		}
//		sq *= sq
//	}
//	return hash, pow
//}
//
//// Index 返回sep在s里第一次匹配时的index, 无法匹配则返回-1.
//func Index(s, sep string) int {
//	n := len(sep)
//	// 先分析一些常见情况, 起到进一步加速的效果
//	switch {
//	case n == 0:
//		return 0
//	case n == 1:  //如果为一个字节,则调用IndixByte(汇编语言)
//		return IndexByte(s, sep[0])
//	case n <= shortStringLen:  //如果sep的长度小于31且大于1, 则使用汇编代码(也是一种优化).
//		return indexShortStr(s, sep)
//	case n == len(s):
//		if sep == s {
//			return 0
//		}
//		return -1
//	case n > len(s):
//		return -1
//	}
//	// 使用Rabin-Karp算法匹配
//	// 步骤1 初始计算待匹配的文本的hash值和乘数因子,
//	hashsep, pow := hashStr(sep)
//	var h uint32
//	for i := 0; i < n; i++ {
//		h = h*primeRK + uint32(s[i])  // 步骤2 计算长度跟sep一样的s子串的hash值
//	}
//	if h == hashsep && s[:n] == sep {
//		return 0
//	}
//	for i := n; i < len(s); {
//		// 利用先前的hash值, 计算新的hash值
//		h *= primeRK  // 乘以base
//		h += uint32(s[i]) // 加上下一个字符的 hash 值
//		h -= pow * uint32(s[i-n]) // 减去先前子串的第一个字符的hash值
//		i++
//		// 如果hash相等则继续使用朴素算法比较, 如果hash不一致,则直接用下一个匹配
//		if h == hashsep && s[i-n:i] == sep {
//			return i - n
//		}
//	}
//	return -1
//}
