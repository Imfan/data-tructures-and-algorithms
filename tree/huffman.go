package tree

import "sort"

type HuffmanNodeList []*HuffmanNode

func (list HuffmanNodeList) Len() int {
	return len(list)
}

func (list HuffmanNodeList) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

func (list HuffmanNodeList) Less(i, j int) bool {
	return list[i].Weight < list[j].Weight
}

//赫夫曼节点，用于构成赫夫曼树
type HuffmanNode struct {
	Weight uint         //权重
	Data   interface{}  //数据
	Parent *HuffmanNode //父节点
	Left   *HuffmanNode //左孩子
	Right  *HuffmanNode //右孩子
}

//赫夫曼树结构，这里使用的interface作为源数据类型
type HuffmanTree struct {
	root    *HuffmanNode           //根节点
	leaf    HuffmanNodeList        //所有叶子节点（即数据对应的节点）
	src     map[interface{}]uint   //源数据，key为数据，value为权重
	codeSet map[interface{}]string //编码集，key为数据，value为通过构造赫夫曼树得到的数据的编码
}

//给定一组字符及其权重的集合，初始化出一棵赫夫曼树
func NewHuffmanTree(src map[interface{}]uint) *HuffmanTree {
	var tree = &HuffmanTree{
		src: src,
	}
	tree.init()
	tree.build()
	tree.parse()
	return tree
}

//将数据进行赫夫曼编码
func (h *HuffmanTree) Coding(target interface{}) (result string) {
	if target == nil {
		return
	}
	var s string
	switch t := target.(type) {
	case string:
		s = t
	case []byte:
		s = string(t)
	default:
		return
	}
	for _, t := range s {
		v := string(t)
		if c, ok := h.codeSet[v]; !ok {
			panic("invalid code: " + v)
		} else {
			result += c
		}
	}
	return result
}

//根据赫夫曼编码获取数据
func (h *HuffmanTree) UnCoding(target string) (result string) {
	node := h.root
	for i := 0; i < len(target); i++ {
		switch target[i] {
		case '0':
			node = node.Left
		case '1':
			node = node.Right
		}
		if node.Left == nil && node.Right == nil {
			result = result + node.Data.(string)
			node = h.root
		}
	}
	return
}

//初始化所有叶子节点
func (h *HuffmanTree) init() {
	if len(h.src) <= 1 {
		panic("invalid src length.")
	}
	// 初始化数据与编码 映射 map
	h.codeSet = make(map[interface{}]string)
	// 初始化 所有 叶节点 切片
	h.leaf = make(HuffmanNodeList, len(h.src))
	var i int
	// 将数据初始化为 赫夫曼节点
	for data, weight := range h.src {
		var node = &HuffmanNode{
			Weight: weight,
			Data:   data,
		}
		h.leaf[i] = node
		i++
	}
	//对leaf根据权值排序
	sort.Sort(h.leaf)
}

//构造赫夫曼树
//src: key为data，value为权值
func (h *HuffmanTree) build() {
	nodeList := h.leaf
	//根据huffman树的规则构造赫夫曼树
	for nodeList.Len() > 1 {
		//1. 选取权值最小的两个node构造出第一个节点
		var temp = &HuffmanNode{
			Weight: nodeList[0].Weight + nodeList[1].Weight,
			Left:   nodeList[0],
			Right:  nodeList[1],
		}
		nodeList[0].Parent = temp
		nodeList[1].Parent = temp

		//2.将生成的新节点插入节点序列中
		nodeList = regroup(nodeList[2:], temp)
	}
	h.root = nodeList[0]
}

//获取每个byte的编码，目的是为了下次需要编码的时候不用再次遍历树以获取每个byte的编码了
//在赫夫曼树中的所有节点要么没有孩子节点，要么有两个孩子节点，不存在只有一个孩子节点的节点
//此处的编码为由底至顶获取，也可以由顶至底的获取
func (h *HuffmanTree) parse() {
	if h.root == nil {
		return
	}
	var temp *HuffmanNode
	var code string
	for _, n := range h.leaf {
		temp = n
		for temp.Parent != nil {
			if temp == temp.Parent.Left {
				code = "0" + code
			} else {
				code = "1" + code
			}
			temp = temp.Parent
		}
		h.codeSet[n.Data] = code
		code = ""
	}
}

//重组，将生成的节点放入既有的list，排序后返回，权值最小的始终在最前面
func regroup(src HuffmanNodeList, temp *HuffmanNode) HuffmanNodeList {
	//将temp添加进src，然后取出weight最小的一个
	length := len(src)
	result := make(HuffmanNodeList, len(src)+1)
	if length == 0 {
		result[0] = temp
		return result
	}
	// 如果新添加的 权重大于最大的，则直接放到最后
	if src[length-1].Weight <= temp.Weight {
		copy(result, src)
		result[length] = temp
		return result
	}

	for i := range src {
		// 如果 原有的小于 新添加的则，放到result的前面
		if src[i].Weight <= temp.Weight {
			result[i] = src[i]
		} else {
			// temp该放的位置
			result[i] = temp
			// 直接复制后面的，因为是排好序的
			copy(result[i+1:], src[i:])
			break
		}
	}
	return result
}
