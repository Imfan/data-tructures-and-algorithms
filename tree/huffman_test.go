package tree

import (
	"fmt"
	"testing"
)

func TestNewHuffmanTree(t *testing.T) {
	huf := NewHuffmanTree(map[interface{}]uint{
		"a": 5,
		"b": 1,
		"c": 6,
		"e": 10,
		"d": 8,
	})

	code := huf.Coding("bcda")
	fmt.Println(code)
	str := huf.UnCoding(code)
	fmt.Println(str)

}
