package skip_list

import (
	"math/rand"
	"time"
)

/*
	跳表特性:
	1. 底层有序链表
	2. 索引链表
	3. 节点
*/

type zLexRangeSpec struct {
	min        string //最小key
	max        string //最大key
	includeMin int    //是否包含最小
	includeMax int    //是否包含最大 默认0包含
}

type zRangeSpec struct {
	min        float64 //分数最小值
	max        float64 //分数最大值
	includeMin int     //是否包含最小值
	includeMax int     //是否包含最大值
}

//链表中的每个节点
type skipNode struct {
	eleKey   string       //节点name
	score    float64      //分数
	backward *skipNode    //后退指针
	level    []*skipLevel //索引层
}

//每一层的索引
type skipLevel struct {
	forward *skipNode //前进指针 就是本层的下一个节点的地址
	span    uint      //层跨越的节点数量
}

//有序链表本身
type skipList struct {
	header *skipNode
	tail   *skipNode
	length uint
	level  int
}

//元素节点
type ele struct {
	mem   string      //成员名
	score float64     //分值
	data  interface{} //附加数据
}

//创建一个跳表节点
func newSkipNode(level int, score float64, key string) *skipNode {
	var node = &skipNode{
		eleKey: key,
		score:  score,
		level:  make([]*skipLevel, level),
	}
	for i := 0; i < level; i++ {
		node.level[i] = &skipLevel{
			forward: nil,
			span:    0,
		}
	}
	return node
}

//创建跳表
func newSkipList() *skipList {
	var zsl = &skipList{
		level:  1,
		length: 0,
		header: newSkipNode(maxLevel, 0, ""),
	}
	// 初始化 跳表头结点的索引层
	for j := 0; j < maxLevel; j++ {
		zsl.header.level[j] = &skipLevel{
			forward: nil,
			span:    0,
		}
	}
	return zsl
}

//随机生成索引层数
func randomLevel() int {
	var level = 1
	rand.Seed(time.Now().Unix())
	for float64(rand.Int()&0xFFFF) < skipListP*0xFFFF {
		level += 1
	}
	if level > maxLevel {
		level = maxLevel
	}
	return level
}

func (e *ele) Name() string {
	return e.mem
}

func (e *ele) Score() float64 {
	return e.score
}

func (e *ele) Data() interface{} {
	return e.data
}

// 新增数据，更新跳表及其索引，并返回 新增的节点地址
func (zsl *skipList) zslInsert(score float64, mem string) *skipNode {
	var x *skipNode
	var update [maxLevel]*skipNode
	var rank [maxLevel]uint
	var i, level int

	if len(mem) == 0 {
		return nil
	}

	x = zsl.header
	// 计算每层排位，和每层需要更新前进指针的节点
	for i = zsl.level - 1; i >= 0; i-- {
		switch {
		// 如果是顶层， 则排位初始化为0，在下面会计算
		case i == zsl.level-1:
			rank[i] = 0
		default: // 非顶层的，初始化为上一层计算的结果，接着在 下面加
			rank[i] = rank[i+1]
		}
		// 找出 每层 应该插入到哪个节点之后
		for x.level[i].forward != nil && (x.level[i].forward.score < score || (x.level[i].forward.score == score && x.level[i].forward.eleKey < mem)) {
			// 计算 update[i] 的 排位， 在设置新插入节点 与 上一个节点（就是update[i]） 的跨度时会用到
			rank[i] += x.level[i].span
			// 继续看下一个节点
			x = x.level[i].forward
		}
		// 记录 在每一层需要 更新指针的节点
		update[i] = x
	}
	// 获取即将新增节点的随机层级数
	level = randomLevel()
	// 如果新增节点层级数较高，大于跳表的层级
	if level > zsl.level {
		// 记录最高层排位为0， 跳表的 头结点 在下面也需要更新
		for i = zsl.level; i < level; i++ {
			rank[i] = 0
			update[i] = zsl.header
			update[i].level[i].span = zsl.length
		}
		// 更新跳表的 最高层级
		zsl.level = level
	}
	// 初始化 新插入的节点
	x = newSkipNode(level, score, mem)
	for i = 0; i < level; i++ {
		// 在每一层 都插入 新生成的节点
		x.level[i].forward = update[i].level[i].forward
		update[i].level[i].forward = x
		// 设置 新插入节点，在每一层的跨度（或者说间隔多少个节点） 第i层的跨度-(最底层紧邻的上一个节点的排名 - 第i层上一个节点的排名)
		x.level[i].span = update[i].level[i].span - (rank[0] - rank[i])
		// 更新 每一层 上一个节点的跨度，rank[0]是 最底层 新插入节点 紧邻的上一个节点的排名，所以要加1
		update[i].level[i].span = (rank[0] - rank[i]) + 1
	}
	// 如果新插入节点的层级小于 跳表的层级
	for i = level; i < zsl.level; i++ {
		// 则  比新插入节点层级高的 更新节点 的 所有层级的跨度都应该加1
		update[i].level[i].span++
	}
	// 如果是 跳表的第一个节点
	if update[0] == zsl.header {
		// 后退指针 设为nil
		x.backward = nil
	} else {
		// 否则 设置为 最底层紧邻的上一个节点
		x.backward = update[0]
	}
	// 如果不是 最后一个节点
	if x.level[0].forward != nil {
		// 更新新插入的节点的前面节点的 后退指针为自己
		x.level[0].forward.backward = x
	} else {
		// 如果 新插入的节点是最后一个节点，则把跳表的尾部 设置为 新插入的节点
		zsl.tail = x
	}
	// 插入一个节点 长度加一
	zsl.length++
	return x
}

// 删除x节点，并更新 各层索引
func (zsl *skipList) zslDeleteNode(x *skipNode, update []*skipNode) {
	var i int
	// 更新各层缓存
	for i = 0; i < zsl.level; i++ {
		if update[i].level[i].forward == x {
			// 更新 各层 要删除节点的前一个节点的跨度和前进指针
			update[i].level[i].span += x.level[i].span - 1
			update[i].level[i].forward = x.level[i].forward
		} else {
			// 更新 没有 要删除节点的 各层的 跨度
			update[i].level[i].span -= 1
		}
	}
	// 如果不是 最后一个节点
	if x.level[0].forward != nil {
		// 更新后退指针
		x.level[0].forward.backward = x.backward
	} else {
		// 更新 尾结点
		zsl.tail = x.backward
	}
	// 如果最高层索引 为空的话，则减少层级
	for zsl.level > 1 && zsl.header.level[zsl.level-1].forward == nil {
		zsl.level--
	}
	// 总长度减1
	zsl.length--
}

// 根据 分数和key 删除 某个元素
func (zsl *skipList) zslDelete(score float64, key string) bool {
	var x *skipNode
	var update = make([]*skipNode, maxLevel)
	var i int

	x = zsl.header
	for i = zsl.level - 1; i >= 0; i-- {
		// 寻找 需要删除节点  各层的 上一个节点，需要更新
		for x.level[i].forward != nil && (x.level[i].forward.score < score ||
			(x.level[i].forward.score == score && x.level[i].forward.eleKey < key)) {
			x = x.level[i].forward
		}
		update[i] = x
	}
	// 找到需要删除的节点，就删除
	x = x.level[0].forward
	if x != nil && score == x.score && x.eleKey == key {
		// 删除需要删除的节点x，并更新 相应的节点
		zsl.zslDeleteNode(x, update)
		return true
	}
	return false
}

// 更新分数，并返回更新后的节点，只有map中存在 key 才可以更新
func (zsl *skipList) zslUpdateScore(curScore float64, key string, newScore float64) *skipNode {
	var update = make([]*skipNode, maxLevel)
	var x *skipNode
	var i int
	x = zsl.header
	for i = zsl.level - 1; i >= 0; i-- {
		// 根据当前分数和key 查找各层需要更新的节点
		for x.level[i].forward != nil && (x.level[i].forward.score < curScore || (x.level[i].forward.score == curScore && x.level[i].forward.eleKey < key)) {
			x = x.level[i].forward
		}
		update[i] = x
	}

	x = x.level[0].forward
	// 如果是首或尾节点，并且更新后 的分数 位置依旧跟之前 排位一样，则直接更新分数
	if (x.backward == nil || x.backward.score < newScore) && (x.level[0].forward == nil || x.level[0].forward.score > newScore) {
		x.score = newScore
		return x
	}
	// 先删除，再插入
	zsl.zslDeleteNode(x, update)
	newNode := zsl.zslInsert(newScore, x.eleKey)
	x = nil
	return newNode
}

func zslValueGteMin(value float64, spec *zRangeSpec) bool {
	if spec.includeMin != 0 {
		return value > spec.min
	}
	return value >= spec.min
}

func zslValueLteMax(value float64, spec *zRangeSpec) bool {
	if spec.includeMax != 0 {
		return value < spec.max
	}
	return value <= spec.max
}

func (zsl *skipList) zslIsInRange(spec *zRangeSpec) bool {
	var x *skipNode

	if spec.min > spec.max || (spec.min == spec.max && (spec.includeMin != 0 || spec.includeMax != 0)) {
		return false
	}
	x = zsl.tail
	if x == nil || !zslValueGteMin(x.score, spec) {
		return false
	}
	x = zsl.header.level[0].forward
	if x == nil || !zslValueLteMax(x.score, spec) {
		return false
	}
	return true
}

func (zsl *skipList) zslFirstInRange(spec *zRangeSpec) *skipNode {
	var x *skipNode
	var i int

	if !zsl.zslIsInRange(spec) {
		return nil
	}

	x = zsl.header
	for i = zsl.level - 1; i >= 0; i-- {
		for x.level[i].forward != nil && !zslValueGteMin(x.level[i].forward.score, spec) {
			x = x.level[i].forward
		}
	}

	x = x.level[0].forward
	if !zslValueLteMax(x.score, spec) {
		return nil
	}
	return x
}

func (zsl *skipList) zslLastInRange(spec *zRangeSpec) *skipNode {
	var x *skipNode
	var i int

	if !zsl.zslIsInRange(spec) {
		return nil
	}

	x = zsl.header
	for i = zsl.level - 1; i >= 0; i-- {
		for x.level[i].forward != nil && zslValueLteMax(x.level[i].forward.score, spec) {
			x = x.level[i].forward
		}
	}
	if !zslValueGteMin(x.score, spec) {
		return nil
	}
	return x
}

// 根据分数范围删除
func (zsl *skipList) zslDeleteRangeByScore(spec *zRangeSpec, dict map[string]*ele) uint {
	var x *skipNode
	var update = make([]*skipNode, maxLevel)
	var i int
	var removed uint = 0

	x = zsl.header
	for i = zsl.level - 1; i >= 0; i-- {
		// 寻找 各层 要删除的最小节点的 前一个节点，它们需要更新前进指针、跨度
		for x.level[i].forward != nil && ((spec.includeMin != 0 && x.level[i].forward.score <= spec.min) || (spec.includeMin == 0 && x.level[i].forward.score < spec.min)) {
			x = x.level[i].forward
		}
		update[i] = x
	}

	x = x.level[0].forward
	// 删除范围内的节点
	for x != nil && ((spec.includeMax != 0 && x.score < spec.max) || (spec.includeMax == 0 && x.score <= spec.max)) {
		var next = x.level[0].forward
		zsl.zslDeleteNode(x, update)
		delete(dict, x.eleKey)
		removed++
		x = next
	}
	return removed
}

func zslLexValueGteMin(key string, spec *zLexRangeSpec) bool {
	if spec.includeMin != 0 {
		return key > spec.min
	}
	return key >= spec.min
}

func zslLexValueLteMax(value string, spec *zLexRangeSpec) bool {
	if spec.includeMax != 0 {
		return value < spec.max
	}
	return value <= spec.max
}

func (zsl *skipList) zslIsInLexRange(spec *zLexRangeSpec) bool {
	var x *skipNode

	if spec.min > spec.max || (spec.min == spec.max && (spec.includeMin != 0 || spec.includeMax != 0)) {
		return false
	}
	x = zsl.tail
	if x == nil || !zslLexValueGteMin(x.eleKey, spec) {
		return false
	}
	x = zsl.header.level[0].forward
	if x == nil || !zslLexValueLteMax(x.eleKey, spec) {
		return false
	}
	return true
}

func (zsl *skipList) zslFirstInLexRange(spec *zLexRangeSpec) *skipNode {
	var x *skipNode
	var i int

	if !zsl.zslIsInLexRange(spec) {
		return nil
	}
	x = zsl.header
	for i = zsl.level - 1; i >= 0; i-- {
		for x.level[i].forward != nil && !zslLexValueGteMin(x.level[i].forward.eleKey, spec) {
			x = x.level[i].forward
		}
	}

	x = x.level[0].forward
	if !zslLexValueLteMax(x.eleKey, spec) {
		return nil
	}
	return x
}

func (zsl *skipList) zslLastInLexRange(spec *zLexRangeSpec) *skipNode {
	var x *skipNode
	var i int

	if !zsl.zslIsInLexRange(spec) {
		return nil
	}

	x = zsl.header
	for i = zsl.level - 1; i >= 0; i-- {
		for x.level[i].forward != nil && zslLexValueLteMax(x.level[i].forward.eleKey, spec) {
			x = x.level[i].forward
		}
	}

	if !zslLexValueGteMin(x.eleKey, spec) {
		return nil
	}
	return x
}

//所有节点分数相同的情况下，以字典序删除区间范围内的节点
func (zsl *skipList) zslDeleteRangeByLex(spec *zLexRangeSpec, dict map[string]*ele) uint {
	var x *skipNode
	var update = make([]*skipNode, maxLevel)
	var removed uint
	var i int

	x = zsl.header
	for i = zsl.level - 1; i >= 0; i-- {
		for x.level[i].forward != nil && !zslLexValueGteMin(x.level[i].forward.eleKey, spec) {
			x = x.level[i].forward
		}
		update[i] = x
	}

	x = x.level[0].forward

	for x != nil && zslLexValueLteMax(x.eleKey, spec) {
		next := x.level[0].forward
		zsl.zslDeleteNode(x, update)
		delete(dict, x.eleKey)
		removed++
		x = next
	}
	return removed
}

func (zsl *skipList) zslDeleteRangeByRank(start, end uint, dict map[string]*ele) uint {
	var x *skipNode
	var update = make([]*skipNode, maxLevel)
	var i int
	var traversed, removed uint

	x = zsl.header
	// 找各层需要更新的节点
	for i = zsl.level - 1; i >= 0; i-- {
		for x.level[i].forward != nil && (traversed+x.level[i].span) < start {
			traversed += x.level[i].span
			x = x.level[i].forward
		}
		update[i] = x
	}
	// 排名，要从 第一开始
	traversed++
	// 开始删除，直到最后一个
	x = x.level[0].forward
	for x != nil && traversed <= end {
		next := x.level[0].forward
		zsl.zslDeleteNode(x, update)
		delete(dict, x.eleKey)
		removed++
		traversed++
		x = next
	}
	return removed
}

func (zsl *skipList) zslGetRank(score float64, ele string) uint {
	var x *skipNode
	var i int
	var rank uint

	x = zsl.header
	for i = zsl.level - 1; i >= 0; i-- {
		for x.level[i].forward != nil && (x.level[i].forward.score < score ||
			(x.level[i].forward.score == score && x.level[i].forward.eleKey <= ele)) {
			rank += x.level[i].span
			x = x.level[i].forward
		}
		if len(x.eleKey) > 0 && x.eleKey == ele {
			return rank
		}
	}
	return 0
}

func (zsl *skipList) zslGetElementByRank(rank uint) *skipNode {
	var x *skipNode
	var traversed uint
	var i int

	x = zsl.header
	for i = zsl.level - 1; i >= 0; i-- {
		for x.level[i].forward != nil && ((traversed + x.level[i].span) <= rank) {
			traversed += x.level[i].span
			x = x.level[i].forward
		}
		if traversed == rank {
			return x
		}
	}
	return nil
}
