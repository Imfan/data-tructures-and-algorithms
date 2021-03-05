package tree

import "fmt"

type BinaryTree struct {
	data  string
	LNode *BinaryTree
	RNode *BinaryTree
}

func (t BinaryTree) GetLNode() Tree {
	return Tree(t.LNode)
}

func (t BinaryTree) GetRNode() Tree {
	return Tree(t.RNode)
}

func (t BinaryTree) GetData() interface{} {
	return t.data
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
