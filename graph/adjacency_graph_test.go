package graph

import (
	"math/rand"
	"strconv"
	"testing"
)

const letter = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randStr(n int) string {
	s := make([]byte, n)
	for i := range s {
		s[i] = letter[rand.Intn(n)]
	}
	return string(s)
}
func TestGraphAdjList_AddEdge(t *testing.T) {
	g := NewGraph()
	for i := 0; i < 20; i++ {
		g.AddNode(strconv.Itoa(i), strconv.Itoa(i))

	}
	g.AddEdge("0", "10")
	g.AddEdge("0", "12")
	g.AddEdge("10", "4")
	g.AddEdge("4", "18")
	g.AddEdge("18", "9")
	//g.DFS()
	g.BFS()
}
