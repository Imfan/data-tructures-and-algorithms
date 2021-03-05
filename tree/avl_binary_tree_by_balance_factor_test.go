package tree

import (
	"testing"
)

func TestAvlBinaryTree_Insert(t *testing.T) {
	avl := NewAvlTree()
	for i := 8; i > 1; i-- {
		avl.Insert(i)
	}
	LevelTraversalByBFS(avl)
}
