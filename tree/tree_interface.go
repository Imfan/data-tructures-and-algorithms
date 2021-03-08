package tree

import (
	"fmt"
	"reflect"
)

type Tree interface {
	GetLNode() Tree
	GetRNode() Tree
	GetData() interface{}
}

type AVLInsert interface {
	GetIndex() int
	rRotate()
	lRotate()
	leftBalance()
	rightBalance()
	Tree
}

//二叉树前序遍历递归方式
func PreOrderTraverse(t Tree) {
	if reflect.ValueOf(t).IsNil() {
		//fmt.Println("前序遍历")
		return
	}
	fmt.Println(t.GetData())
	PreOrderTraverse(t.GetLNode()) // 先遍历左子树
	PreOrderTraverse(t.GetRNode()) // 遍历右子树
}

// 中序遍历
func InOrderTraverse(t Tree) {
	if reflect.ValueOf(t).IsNil() {
		//fmt.Println("中序遍历")
		return
	}
	InOrderTraverse(t.GetLNode())
	fmt.Println(t.GetData())
	InOrderTraverse(t.GetRNode())
}

// 后序遍历
func PostOrderTraverse(t Tree) {
	if reflect.ValueOf(t).IsNil() {
		//fmt.Println("后序遍历")
		return
	}
	PostOrderTraverse(t.GetLNode())
	PostOrderTraverse(t.GetRNode())
	fmt.Println(t.GetData())
}

// 层序遍历 使用迭代  广度优先搜索
func LevelTraversalByBFS(t Tree) {
	var (
		// 作为队列，广度优先次序 先进先出
		q    []Tree
		qBak []Tree
		tmp  Tree
	)
	// 到达叶子节点 直接返回
	if reflect.ValueOf(t).IsNil() {
		return
	}
	// 先把根节点放进去，使其可以开始循环
	q = append(q, t)
	// 只要队列里面有元素就继续， 当到了叶子节点后，qBak就会为空
	for i := 0; len(q) > 0; i++ {
		result = append(result, []int{})
		// 每次都要初始化
		qBak = []Tree{}
		// 循环完上次 添加的所有节点，并 添加下层所有的节点到 队列中
		for j := 0; j < len(q); j++ {
			tmp = q[j]
			d, _ := tmp.GetData().(int)
			//fmt.Println(d)
			result[i] = append(result[i], d)
			if !reflect.ValueOf(tmp.GetLNode()).IsNil() {
				qBak = append(qBak, tmp.GetLNode())
			}
			if !reflect.ValueOf(tmp.GetRNode()).IsNil() {
				qBak = append(qBak, tmp.GetRNode())
			}
		}
		q = qBak
	}
	fmt.Println(result)
}

var result [][]int

// 层序遍历
func LevelTraversalByRecursion(t Tree) [][]int {
	if reflect.ValueOf(t).IsNil() {
		return [][]int{}
	}
	dfsHelper(t, 0)
	return [][]int{}
}

func dfsHelper(node Tree, level int) {
	if reflect.ValueOf(node).IsNil() {
		return
	}
	if len(result) < level+1 {
		result = append(result, make([]int, 0))
	}
	data, _ := node.GetData().(int)
	result[level] = append(result[level], data)
	dfsHelper(node.GetLNode(), level+1)
	dfsHelper(node.GetRNode(), level+1)
}
