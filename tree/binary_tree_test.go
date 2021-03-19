package tree

import (
	"fmt"
	"strconv"
	"testing"
)

func TestTree(t *testing.T) {
	t1 := InitBinaryTree("1")
	for i := 2; i < 10; i++ {
		t1.Insert(strconv.Itoa(i))
	}

	str := "RedBlackTree\n"
	//if !t1.Empty() {
	Output(t1, "", true, &str)
	fmt.Println(str)
	//}
	//PreOrderTraverse(t1)
	//PostOrderTraverse(t1)
	//InOrderTraverse(t1)
	LevelTraversalByBFS(t1)
	//dfsPrint(t1)
}
