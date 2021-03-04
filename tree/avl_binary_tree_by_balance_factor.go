package tree

const (
	LH = 1  // 左子树高
	EH = 0  // 左右子树等高
	RH = -1 // 右子树高
)

type AvlBinaryTree struct {
	bf    int8           // 节点的平衡因子
	index int            // 整型数据
	Data  string         // 字符串数据
	LNode *AvlBinaryTree // 左节点
	RNode *AvlBinaryTree // 右节点
}

// 右旋处理
// 对以t为根的二叉排序树做右旋处理，
func (t *AvlBinaryTree) rRotate() {
	// t 为需要右旋的的二叉树
	tmp := t.LNode      // tmp指向左子树
	t.LNode = tmp.RNode // 将t的左子树改为指向 左子树的右子树
	tmp.RNode = t
	*t = *tmp // 将t的数据改为 右旋后的子树
}

// 左旋操作
// 对以t为根的二叉排序树做左旋处理,与右旋处理相反
func (t *AvlBinaryTree) lRotate() {
	tmp := t.RNode
	t.RNode = tmp.LNode
	tmp.LNode = t
	*t = *tmp
}

func (t *AvlBinaryTree) leftBalance() {

}
