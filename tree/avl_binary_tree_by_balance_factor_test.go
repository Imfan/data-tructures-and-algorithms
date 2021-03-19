package tree

import (
	"fmt"
	"testing"
)

func TestAvlBinaryTree_Insert(t *testing.T) {
	avl := InitAvlTree(1)
	for i := 8; i >= 1; i-- {
		Insert(InitAvlTree(i, []int{i}), avl)
	}
	//Insert(1, avl)
	//Insert(2, avl)
	//Insert(3, avl)
	//LevelTraversalByBFS(avl)
	str := ""
	Output(avl, "", true, &str)
	fmt.Println(str)

	//PreOrderTraverse(avl)
	//InOrderTraverse(avl)
	PostOrderTraverse(avl)

}
