package tree

const (
	LH = 1  // 左子树高
	EH = 0  // 左右子树等高
	RH = -1 // 右子树高
)

type AvlBinaryTree struct {
	bf    int8 // 节点的平衡因子
	index int  // 整型数据
	data  interface{}
	lNode *AvlBinaryTree
	rNode *AvlBinaryTree
}

func InitAvlTree(index int, data ...interface{}) *AvlBinaryTree {
	return &AvlBinaryTree{
		bf:    0,
		index: index,
		data:  data,
		lNode: nil,
		rNode: nil,
	}
}

func (t *AvlBinaryTree) GetLNode() Tree {
	return Tree(t.lNode)
}

func (t *AvlBinaryTree) GetRNode() Tree {
	return Tree(t.rNode)
}

func (t *AvlBinaryTree) GetData() interface{} {
	return t.index
}

func (t *AvlBinaryTree) GetIndex() int {
	return t.index
}

// 右旋处理
// 对以t为根的二叉排序树做右旋处理，
func rRotate(t *AvlBinaryTree) {
	// t 为需要右旋的的二叉树
	tmp := t.lNode      // tmp指向左子树
	t.lNode = tmp.rNode // 将t的左子树改为指向 左子树的右子树
	// 必须得初始化一个地址，不然 tmp.lNode = t; *t=*tmp 会造成指针死循环
	tmp.rNode = &AvlBinaryTree{}
	*tmp.rNode = *t
	*t = *tmp // 将t的数据改为 右旋后的子树
}

// 左旋操作
// 对以t为根的二叉排序树做左旋处理,与右旋处理相反
func lRotate(t *AvlBinaryTree) {
	tmp := t.rNode
	t.rNode = tmp.lNode
	// 必须得初始化一个地址，不然 tmp.lNode = t; *t=*tmp 会造成指针死循环
	tmp.lNode = &AvlBinaryTree{}
	*tmp.lNode = *t
	*t = *tmp
}

// 对以t为根的二叉树做左平衡旋转操作
func leftBalance(t *AvlBinaryTree) {
	L := t.lNode
	switch L.bf { // 检查t的左子树的平衡因子，并做相应处理
	case LH: // L左树高了，需要单右旋操作
		rRotate(t)
		// 旋转操作后 平衡了
		t.bf = EH
		L.bf = EH
		break
	case RH: // L右树高了，要先对其左子树左旋再t右旋
		Lr := L.rNode // t的左节点的右子树节点
		// 修改t及其左子节点的平衡因子
		switch Lr.bf { // 这些状态值，在做完相应旋转操作后，都是固定的，可以画图试试
		case LH:
			t.bf = RH
			L.bf = EH
		case EH:
			t.bf = EH
			L.bf = EH
		case RH:
			t.bf = EH
			L.bf = LH
		}
		Lr.bf = EH
		lRotate(L) // 对左子树 左旋
		rRotate(t) // 对t右旋
	}
}

// 右平衡操作
func rightBalance(t *AvlBinaryTree) {
	R := t.rNode
	switch R.bf {
	case RH:
		lRotate(t) // 需要左旋转
		// 旋转操作后 平衡了
		t.bf = EH
		R.bf = EH
	case LH: //
		Rl := R.lNode
		switch Rl.bf {
		case RH:
			t.bf = LH
			R.bf = EH
		case EH:
			t.bf = EH
			R.bf = EH
		case LH:
			t.bf = EH
			R.bf = RH
		}
		Rl.bf = EH
		rRotate(R)
		lRotate(t)
	}
}

/**
ok: 插入是否成功
taller: 树是否变高
*/
func Insert(InsertNode *AvlBinaryTree, t *AvlBinaryTree) (ok, taller bool) {
	if t.GetIndex() == InsertNode.GetIndex() {
		return false, false
	}
	if InsertNode.GetIndex() < t.GetIndex() {
		if t.lNode == nil {
			t.lNode = InsertNode
			return true, true
		} else {
			ok, taller = Insert(InsertNode, t.lNode)
		}
		if taller {
			switch t.bf {
			case LH: // 原本左子树比右子树高，需要左平衡处理
				leftBalance(t)
				taller = false
			case EH: // 原本左右子树等高， 现在因为左子树增高而树增高
				t.bf = LH
				taller = true
			case RH: // 原本右子树比左子树高，现在左右子树等高
				t.bf = EH
				taller = false
			}
		}
	} else {

		if t.rNode == nil {
			t.rNode = InsertNode
			taller = true
		} else {
			ok, taller = Insert(InsertNode, t.rNode)
		}
		if taller {
			switch t.bf {
			case LH:
				// 平衡了 taller要设为false
				t.bf = EH
				taller = false
			case EH:
				t.bf = RH
				taller = true
			case RH:
				// 右平衡操作后，t的子树可能不平衡，但t肯定会平衡，所以taller设为false
				rightBalance(t)
				taller = false
			}
		}
	}
	return true, taller
}
