package tree

import (
	"fmt"
	"reflect"
	"strconv"
)

type Tree interface {
	GetLNode() interface{}
	GetRNode() interface{}
	GetData() interface{}
	fmt.Stringer
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
	PreOrderTraverse(t.GetLNode().(Tree)) // 先遍历左子树
	PreOrderTraverse(t.GetRNode().(Tree)) // 遍历右子树
}

// 中序遍历
func InOrderTraverse(t Tree) {
	if reflect.ValueOf(t).IsNil() {
		//fmt.Println("中序遍历")
		return
	}
	InOrderTraverse(t.GetLNode().(Tree))
	fmt.Println(t.GetData())
	InOrderTraverse(t.GetRNode().(Tree))
}

// 后序遍历
func PostOrderTraverse(t Tree) {
	if reflect.ValueOf(t).IsNil() {
		//fmt.Println("后序遍历")
		return
	}
	PostOrderTraverse(t.GetLNode().(Tree))
	PostOrderTraverse(t.GetRNode().(Tree))
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
			d, _ := strconv.Atoi(tmp.String())
			//fmt.Println(reflect.TypeOf(tmp.GetData()))
			result[i] = append(result[i], d)
			if !reflect.ValueOf(tmp.GetLNode()).IsNil() {
				qBak = append(qBak, tmp.GetLNode().(Tree))
			}
			if !reflect.ValueOf(tmp.GetRNode()).IsNil() {
				qBak = append(qBak, tmp.GetRNode().(Tree))
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

func dfsPrint(node Tree) {
	if reflect.ValueOf(node).IsNil() {
		return
	}
	fmt.Println(node.GetData())
	dfsPrint(node.GetLNode().(Tree))
	dfsPrint(node.GetRNode().(Tree))
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
	dfsHelper(node.GetLNode().(Tree), level+1)
	dfsHelper(node.GetRNode().(Tree), level+1)
}

// 可视化 输出树结构
func Output(node Tree, prefix string, isTail bool, str *string) {
	if !reflect.ValueOf(node.GetRNode()).IsNil() {
		newPrefix := prefix
		if isTail {
			newPrefix += "│   "
		} else {
			newPrefix += "    "
		}
		Output(node.GetRNode().(Tree), newPrefix, false, str)
	}
	*str += prefix
	if isTail {
		*str += "└── "
	} else {
		*str += "┌── "
	}
	*str += node.String() + "\n"
	if !reflect.ValueOf(node.GetLNode()).IsNil() {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}
		Output(node.GetLNode().(Tree), newPrefix, true, str)
	}
}
