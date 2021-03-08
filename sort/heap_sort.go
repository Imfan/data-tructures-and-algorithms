package sort

import (
	"sort"
)

type Interface interface {
	sort.Interface
	Push(x interface{}) // add x as element Len()
	Pop() interface{}   // remove and return element Len() - 1.
}

func Init(h Interface) {
	// 堆化
	n := h.Len()
	// 从非叶子节点开始，因为非叶子结点不需要堆化，n/2-1 表示是从0开始的
	for i := n/2 - 1; i >= 0; i-- {
		down(h, i, n)
	}
}

// Push pushes the element x onto the
// The complexity is O(log n) where n = h.Len().
func Push(h Interface, x interface{}) {
	// 追加到最后一个节点，并使其 上浮
	h.Push(x)
	up(h, h.Len()-1)
}

// Pop removes and returns the minimum element (according to Less) from the
// The complexity is O(log n) where n = h.Len().
// Pop is equivalent to Remove(h, 0).
func Pop(h Interface) interface{} {
	n := h.Len() - 1
	h.Swap(0, n)
	// 把 换过来的最后一个元素，放到合适的位置，因为已经堆化了，所以只需要 下沉一次
	down(h, 0, n)
	return h.Pop()
}

// 使元素上浮
func up(h Interface, j int) {
	for {
		i := (j - 1) / 2 // parent
		// 当 j 和它的父节点 不满足条件或 i==j的时候，就退出循环
		if i == j || !h.Less(j, i) {
			break
		}
		// 如果 h.Less(j,i)返回true 就交换，往上走
		h.Swap(i, j)
		j = i
	}
}

// h 为需要堆化是切片，i0,要堆化的元素，n下标最大值
// 元素 i ，将其与它的子节点 2i+1 和 2i+2比较，如果元素 i 比它的子节点小，则将元素 i 与两个子节点中较小的节点交换（j），从而保证满足最小树的要求
func down(h Interface, i0, n int) bool {
	i := i0
	// 从下往上堆化的，在到 上面的时候，需要继续 将值向下比较，放到合适的位置
	// 如果子节点 j 可能也有它的子节点，继续比较、交换，直到数组末尾，或者元素 i 比它的两个子节点都小，跳出循环（）。
	// 如果没有子节点，则越界判断会跳出循环
	for {
		// 左子节点，因为是从0开始的所以 +1是左节点，书上大部分为了讲解方便，都是从1开始标号的
		j1 := 2*i + 1
		// j2为右节点，如果  右子树比左子树小，则将 j 等于 右子树的索引，否则还是左子树的索引
		j2 := j1 + 1
		// 越界判断
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		// 赋值，下面与i比较要用
		j := j1 // left child

		// j1和j2谁符合条件，j就赋值为谁的值
		// j2 < n：表示 如果有右节点则和右节点比较，没有就在下面只和左节点比较
		if j2 < n && h.Less(j2, j1) {
			j = j2 // = 2*i + 2  // right child
		}
		// 将i节点与 左右子节点中最小（或最大）的值比较，如果 索引j 的值 大于（小于） 索引 i 的值，则说明 索引 i的值是最小（最大）的，直接返回
		// 为什么元素 i 比它的两个子节点都小，就可以跳出循环，不再继续下去呢？
		// 这是由于，在Init函数中，第一个开始down（下沉）的元素是第 n/2 - 1 个，可以保证总是从最后一棵子树开始down
		// 如果元素 i 比它的两个子节点都小，那么该元素对应的子树，就是最小堆。
		if !h.Less(j, i) {
			break
		}
		// 上面判断为真，则交换i和j的值
		h.Swap(i, j)
		// 虽说 到这 i 和它的左右子节点比较符合条件了，但是不能确定，它跟它子树里面的所有节点比是不是都符合条件
		// 所以 让 i=j,得到 i交换后对应的位置，继续 down，比较并下沉，知道 它不用交换值和它的左右节点比较也符合条件
		i = j
	}
	// 返回有没有交换值，如果交换过了，则返回true
	return i > i0
}

// 先将要删除的节点 i 与末尾节点 n 交换，然后将交换后新的节点 i 下沉或上浮到合适的位置
func Remove(h Interface, i int) interface{} {
	n := h.Len() - 1
	if n != i {
		h.Swap(i, n)
		if !down(h, i, n) {
			up(h, i)
		}
	}
	return h.Pop()
}

// Fix re-establishes the heap ordering after the element at index i has changed its value.
// Changing the value of the element at index i and then calling Fix is equivalent to,
// but less expensive than, calling Remove(h, i) followed by a Push of the new value.
// The complexity is O(log n) where n = h.Len().
// 更新某个值之后，使其 下沉或上浮到何时的的位置，修正 它的位置
func Fix(h Interface, i int) {
	if !down(h, i, h.Len()) {
		up(h, i)
	}
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	// Push和Pop使用指针接收者
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
