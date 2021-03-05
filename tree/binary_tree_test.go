package tree

import (
	"strconv"
	"testing"
)

func TestTree(t *testing.T) {
	t1 := InitBinaryTree("1")
	for i := 2; i < 10; i++ {
		t1.Insert(strconv.Itoa(i))
	}
	//PreOrderTraverse(t1)
	//PostOrderTraverse(t1)
	//InOrderTraverse(t1)
	LevelTraversalByBFS(t1)
}
