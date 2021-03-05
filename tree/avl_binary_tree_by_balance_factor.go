package tree

const (
	LH = 1  // 左子树高
	EH = 0  // 左右子树等高
	RH = -1 // 右子树高
)

type AvlBinaryTree struct {
	bf    int8 // 节点的平衡因子
	index int  // 整型数据
	data  int
	lNode *AvlBinaryTree
	rNode *AvlBinaryTree
}

func NewAvlTree() *AvlBinaryTree {
	return &AvlBinaryTree{}
}

func (t AvlBinaryTree) GetLNode() Tree {
	return Tree(t.lNode)
}

func (t AvlBinaryTree) GetRNode() Tree {
	return Tree(t.rNode)
}

func (t AvlBinaryTree) GetData() interface{} {
	return t.data
}

// 右旋处理
// 对以t为根的二叉排序树做右旋处理，
func (t *AvlBinaryTree) rRotate() {
	//fmt.Println("right")
	// t 为需要右旋的的二叉树
	tmp := t.lNode      // tmp指向左子树
	t.lNode = tmp.rNode // 将t的左子树改为指向 左子树的右子树
	tmp.rNode = t
	*t = *tmp // 将t的数据改为 右旋后的子树
}

// 左旋操作
// 对以t为根的二叉排序树做左旋处理,与右旋处理相反
func (t *AvlBinaryTree) lRotate() {
	//fmt.Println("left")
	tmp := t.rNode
	t.rNode = tmp.lNode
	tmp.lNode = t
	*t = *tmp
}

// 对以t为根的二叉树做左平衡旋转操作
func (t *AvlBinaryTree) leftBalance() {
	//fmt.Println("leftB")
	L := t.lNode
	switch L.bf { // 检查t的左子树的平衡因子，并做相应处理
	case LH: // 说明 新节点插入在t的左子节点的左子树上，需要单右旋操作
		t.rRotate()
		// 旋转操作后 平衡了
		t.bf = EH
		L.bf = EH
		break
	case RH: // 新加入的节点插入到了 t的左节点的右子树上，要先对其左子树左旋再t右旋
		Lr := L.rNode // t的左节点的右子树节点
		// 修改t及其左子节点的平衡因子
		switch Lr.bf {
		case LH:
			t.bf = RH //
			L.bf = EH
		case EH:
			t.bf = EH
			L.bf = EH
		case RH:
			t.bf = EH
			L.bf = LH
		}
		Lr.bf = EH
		L.lRotate() // 对左子树 左旋
		t.rRotate() // 对t右旋
	}
}

// 右平衡操作
func (t *AvlBinaryTree) rightBalance() {
	//fmt.Println("rightB")
	R := t.rNode
	switch R.bf { //
	case RH: //
		t.lRotate()
		// 旋转操作后 平衡了
		t.bf = EH
		R.bf = EH
	case LH: //
		Rl := R.lNode //
		//fmt.Println(Rl)
		//
		switch Rl.bf {
		case RH:
			t.bf = LH //
			R.bf = EH
		case EH:
			t.bf = EH
			R.bf = EH
		case LH:
			t.bf = EH
			R.bf = RH
		}
		Rl.bf = EH
		R.rRotate() //
		t.lRotate() //
	}
}

func (t *AvlBinaryTree) Insert(data int) (bool, bool) {
	//fmt.Println(data, t.data)
	var taller, ok bool
	if t.data == data {
		return false, false
	}
	if data < t.data {
		if t.lNode == nil {
			t.lNode = &AvlBinaryTree{
				data: data,
				bf:   EH,
			}
			return true, true
		}
		ok, taller = t.lNode.Insert(data)
		if !ok {
			return false, taller
		}
		//fmt.Println(taller)

		if taller {
			switch t.bf {
			case LH: // 原本左子树比右子树高，需要左平衡处理
				t.leftBalance()
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
			t.rNode = &AvlBinaryTree{
				data: data,
				bf:   EH,
			}
			return true, true
		}
		ok, taller = t.rNode.Insert(data)
		if !ok {
			return false, taller
		}
		//fmt.Println(taller)

		if taller {
			switch t.bf {
			case LH:
				t.bf = EH
				taller = false
			case EH:
				t.bf = RH
				taller = true
			case RH:
				t.rightBalance()
				taller = false
			}
		}
	}
	return true, taller
}
