package sort

import (
	"math/rand"
	"time"
)

var (
	end       = 1 + randInt()
	randRange = end * 10
	input     = []int{}
	// 从小到大顺序排序
	smallToLarge = true
)

func randInt() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(9)
}
