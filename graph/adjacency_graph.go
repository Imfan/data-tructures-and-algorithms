package graph

import "fmt"

type VertexType string // 顶点类型
type EdgeType int      // 边上的权值类型
// 边表节点
type EdgeNode struct {
	adjVex int       // 邻接点域，存储该顶点对应的下标
	weight EdgeType  // 权值
	next   *EdgeNode // 下一个邻接点
}

// 顶点表 节点
type VertexNode struct {
	key       string
	data      string    // 数据
	firstEdge *EdgeNode // 有链接的节点
}

// 邻接表 表示的图
type GraphAdjList struct {
	AdjList               []VertexNode // 表头
	index                 map[string]int
	numVertexes, numEdges int // 顶点和边的数量
}

func NewGraph() *GraphAdjList {
	return &GraphAdjList{
		AdjList:     nil,
		index:       make(map[string]int, 1),
		numVertexes: 0,
		numEdges:    0,
	}
}

func (g *GraphAdjList) Get(key string) string {
	if index, ok := g.index[key]; ok {
		return g.AdjList[index].data
	}
	return ""
}

// 添加节点
func (g *GraphAdjList) AddNode(key, data string) {
	// 在图中有没有
	srcIndex, ok := g.index[key]
	// 没有则插入图中
	if !ok {
		g.AdjList = append(g.AdjList, VertexNode{data: data, key: key})
		srcIndex = len(g.AdjList) - 1
		g.numVertexes++
		// 放入map中
		g.index[key] = srcIndex
	}
}

// 添加边,连线
func (g *GraphAdjList) AddEdge(src, dst string) error {
	var (
		srcIndex, dstIndex int
		ok                 bool
	)
	// 在图中有没有
	srcIndex, ok = g.index[src]
	if !ok {
		panic("没有节点" + src)
	}
	dstIndex, ok = g.index[dst]
	if !ok {
		panic("没有节点" + dst)
	}
	dstNode := &EdgeNode{
		adjVex: dstIndex,
		weight: 0,
		next:   nil,
	}
	node := g.AdjList[srcIndex].firstEdge
	if node == nil {
		g.AdjList[srcIndex].firstEdge = dstNode
	} else {
		for node.next != nil {
			node = node.next
		}
		node.next = dstNode
	}
	g.numEdges++
	return nil
}

var isSearch map[int]bool

// 深度遍历
func (g *GraphAdjList) DFS() {
	isSearch = make(map[int]bool, 1)
	for i := 0; i < g.numVertexes; i++ {
		isSearch[i] = false
	}
	for i := 0; i < g.numVertexes; i++ {
		if !isSearch[i] {
			g.subDFS(i)
		}
	}
}

func (g *GraphAdjList) subDFS(i int) {
	node := g.AdjList[i]
	fmt.Println(node.data)
	isSearch[i] = true
	next := node.firstEdge
	//fmt.Println(next)
	for next != nil {
		if !isSearch[next.adjVex] {
			g.subDFS(next.adjVex)
		}
		next = next.next
	}

}

// 广度优先遍历
func (g *GraphAdjList) BFS() {
	isSearch = make(map[int]bool, 1)
	for i := 0; i < g.numVertexes; i++ {
		isSearch[i] = false
	}
	q := []int{}
	for i := 0; i < g.numVertexes; i++ {
		// 如果没有 扫描过，可以
		if !isSearch[i] {
			q = append(q, i)
			for len(q) > 0 {
				qBak := []int{}
				for i := 0; i < len(q); i++ {
					// 如果扫过了，跳过
					if isSearch[q[i]] {
						continue
					}
					fmt.Println(g.AdjList[q[i]].data)
					// 记录扫过了
					isSearch[q[i]] = true
					// 添加有连接的 节点到队列里面
					p := g.AdjList[q[i]].firstEdge
					// 直到为nil
					for p != nil {
						qBak = append(qBak, p.adjVex)
						p = p.next
					}
				}
				q = qBak
			}
		}
	}
}
