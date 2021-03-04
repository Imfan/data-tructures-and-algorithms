package tree

import "fmt"
import "strconv"

type BinaryTree struct {
	data  string
	LNode *BinaryTree
	RNode *BinaryTree
}

//二叉树前序遍历递归方式
func PreOrderTraverse(t *BinaryTree) {
	if t == nil {
		//fmt.Println("前序遍历")
		return
	}
	println(t.data)
	PreOrderTraverse(t.LNode) // 先遍历左子树
	PreOrderTraverse(t.RNode) // 遍历右子树
}

// 中序遍历
func InOrderTraverse(t *BinaryTree) {
	if t == nil {
		//fmt.Println("中序遍历")
		return
	}
	InOrderTraverse(t.LNode)
	println(t.data)
	InOrderTraverse(t.RNode)
}

// 后序遍历
func PostOrderTraverse(t *BinaryTree) {
	if t == nil {
		//fmt.Println("后序遍历")
		return
	}
	PostOrderTraverse(t.LNode)
	PostOrderTraverse(t.RNode)
	println(t.data)
}

// 层序遍历 使用迭代  广度优先搜索
func LevelTraversalByBFS(t *BinaryTree) {
	var (
		// 作为队列，广度优先次序 先进先出
		q    []*BinaryTree
		qBak []*BinaryTree
		tmp  *BinaryTree
	)
	// 到达叶子节点 直接返回
	if t == nil {
		return
	}
	// 先把根节点放进去，使其可以开始循环
	q = append(q, t)
	// 只要队列里面有元素就继续， 当到了叶子节点后，qBak就会为空
	for i := 0; len(q) > 0; i++ {
		// 每次都要初始化
		qBak = []*BinaryTree{}
		// 循环完上次 添加的所有节点，并 添加下层所有的节点到 队列中
		for j := 0; j < len(q); j++ {
			tmp = q[j]
			fmt.Println(tmp.data)
			if tmp.LNode != nil {
				qBak = append(qBak, tmp.LNode)
			}
			if tmp.RNode != nil {
				qBak = append(qBak, tmp.RNode)
			}
		}
		q = qBak
	}
}

var result [][]int

func LevelTraversalByRecursion(t *BinaryTree) [][]int {
	result = make([][]int, 0)
	if t == nil {
		return result
	}
	dfsHelper(t, 0)
	return result
}

func dfsHelper(node *BinaryTree, level int) {
	if node == nil {
		return
	}
	if len(result) < level+1 {
		result = append(result, make([]int, 0))
	}
	data, _ := strconv.Atoi(node.data)
	result[level] = append(result[level], data)
	dfsHelper(node.LNode, level+1)
	dfsHelper(node.RNode, level+1)
}

func InitBinaryTree(data string) *BinaryTree {
	return &BinaryTree{
		data: data,
	}
}

func (t *BinaryTree) Insert(data string) {
	tmp := t
	for tmp != nil {
		if tmp.data <= data {
			if tmp.RNode == nil {
				tmp.RNode = InitBinaryTree(data)
				return
			}
			tmp = tmp.RNode
		} else if tmp.data > data {
			if tmp.LNode == nil {
				tmp.LNode = InitBinaryTree(data)
				return
			}
			tmp = tmp.LNode
		}

	}
}

func (t *BinaryTree) Delete(data string) bool {

	needDel := t.Find(data)
	fmt.Println(needDel)
	if needDel != nil {
		// 右子树为空
		if needDel.RNode == nil && needDel.LNode != nil {
			*needDel = *needDel.LNode
		} else if needDel.LNode == nil && needDel.RNode != nil {
			// 左子树为空
			*needDel = *needDel.RNode
		} else {
			//  如果左右子树都不为空，修改左右子树都可以，这里选择修改左子树
			l := needDel
			r := needDel.LNode
			for r.RNode != nil {
				l = r
				r = r.RNode
			}
			// 先赋值
			needDel.data = r.data

			// 如果 要删除的节点的左子树的右子树为空，则会相等
			if l == needDel {
				// 直接把要删除的节点的左指针指向 左指针的左指针 r == needdel.LNode
				l.LNode = r.LNode
			} else {
				// 否则 将 被删除的节点的左节点，最大的一个值  放到 被删除节点的位置，
				//最大那个值的节点删了，让他的父节点指向他的左节点 l 为他的父节点
				l.RNode = r.LNode
			}

		}

	}

	return true
}

func (t *BinaryTree) Find(data string) *BinaryTree {
	tmp := t
	for tmp != nil {
		if tmp.data < data {
			tmp = tmp.RNode
		} else if tmp.data > data {
			tmp = tmp.LNode
		} else if tmp.data == data {
			return tmp
		}
	}
	return nil
}
