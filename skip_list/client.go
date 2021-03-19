package skip_list

import (
	"fmt"
	"sync"
)

type SkipList interface {
	Len() int                                                                                                                 //ZCARD: 获取有序集合的成员数
	InsertByEle(member string, score float64, addData interface{}) Node                                                       //ZADD: 向有序集合添加一个成员，或者更新已存在成员的分数
	InsertByEleArray(scoreMemAddValues ...interface{}) ([]Node, error)                                                        //ZADD: 向有序集合添加多个成员，或者更新已存在成员的分数
	IncrBy(member string, increment float64) (bool, error)                                                                    //ZINCRBY: 有序集合中对指定成员的分数加上增量 increment
	DeleteByMem(member ...string) (ok bool, err error)                                                                        //ZREM: 移除有序集合中的一个或多个成员
	DeleteByScore(startScore, endScore float64) (removed uint)                                                                //ZREMRANGEBYSCORE: 移除有序集合中分数区间[startScore, endScore]的所有成员
	DeleteByRank(startRank, endRank uint) (removed uint)                                                                      //ZREMRANGEBYRANK: 移除有序集合中给定的排名区间的所有成员
	UpdateScore(member string, score float64) (ok bool, err error)                                                            //更新指定成员的分数
	UpdateAddData(member string, data interface{}) (ok bool, err error)                                                       //更新指定成员的附加数据
	GetRank(members ...string) (rankInfo map[string]uint, err error)                                                          //ZRANK: 返回有序集合中指定成员的索引
	GetElementByMem(members ...string) (nodes []Node, err error)                                                              //根据成员名获取成员信息
	GetElementByRank(ranks ...uint) (nodes []Node, err error)                                                                 //根据成员排名获取成员信息
	SafeRange(startRank, endRank uint, isReverse bool, do func(mem string, score float64, addData interface{}) error) error   //safe range
	UnSafeRange(startRank, endRank uint, isReverse bool, do func(mem string, score float64, addData interface{}) error) error //unsafe range
}

type Node interface {
	Name() string
	Score() float64
	Data() interface{}
}

type zSet struct {
	lock     sync.Mutex
	dict     map[string]*ele
	skipList *skipList
}

// 实例化
func NewSkipList() SkipList {
	return &zSet{
		lock:     sync.Mutex{},
		dict:     make(map[string]*ele),
		skipList: newSkipList(),
	}
}

func (z *zSet) Len() int {
	return int(z.skipList.length)
}

// 增
// 插入对象，如果分数相同则 直接返回 Node，分数不同则更新 对应member的分数
func (z *zSet) InsertByEle(member string, score float64, addData interface{}) Node {
	z.lock.Lock()
	defer z.lock.Unlock()

	data, ok := z.dict[member]
	switch ok {
	case true:
		//已经存在，分数有变更时才更新
		if data.score != score {
			_ = z.skipList.zslUpdateScore(data.score, member, score)
			data.score = score
		}

	default:
		// 不存在则插入
		_ = z.skipList.zslInsert(score, member)
		data = &ele{
			mem:   member,
			score: score,
			data:  addData,
		}
		z.dict[member] = data
	}

	return data
}

// 参数必须是3个3个，第一个是 key，第二个是分数，第三个是额外数据
func (z *zSet) InsertByEleArray(scoreMemAddValues ...interface{}) ([]Node, error) {
	valueCount := len(scoreMemAddValues)
	// 一定要是 3的倍数，
	if valueCount%3 != 0 {
		return nil, errInvalidScoreMemArray
	}

	z.lock.Lock()
	defer z.lock.Unlock()

	var result []Node = make([]Node, valueCount/3)
	var mem string
	var score float64
	var addData interface{}
	for i := 0; i < valueCount; i += 3 {
		// 初始化 数据
		mem = scoreMemAddValues[i].(string)
		addData = scoreMemAddValues[i+2]
		switch t := scoreMemAddValues[i+1].(type) {
		case int:
			score = float64(t)
		case float64:
			score = t
		default:
			return nil, errInvalidScoreType
		}
		//判断是否已经添加了，如果是，则直接更新
		data, ok := z.dict[mem]
		if ok {
			if score != data.score { //如果分数不同，需要更新分数
				z.skipList.zslUpdateScore(data.score, mem, score)
				data.score = score
			}
			//更新addData
			data.data = addData
		} else { //如果不存在，则直接insert
			node := z.skipList.zslInsert(score, mem)
			if node == nil {
				continue
			}
			data = &ele{
				mem:   mem,
				score: score,
				data:  addData,
			}
			z.dict[mem] = data
		}
		result[i/3] = data
	}
	return result, nil
}

// 删
// 根据 key 删除多个 元素
func (z *zSet) DeleteByMem(members ...string) (bool, error) {
	if len(members) == 0 {
		return false, nil
	}

	z.lock.Lock()
	defer z.lock.Unlock()

	for _, member := range members {
		data, ok := z.dict[member]
		if !ok {
			return false, fmt.Errorf("no mem: %s", member)
		}
		ok = z.skipList.zslDelete(data.score, data.mem)
		if !ok {
			return false, fmt.Errorf("delete %s fail", member)
		}
		delete(z.dict, member)
	}
	return true, nil
}

//删除某个分数区间的（闭区间 包含两端）的节点，返回删除了几个节点
func (z *zSet) DeleteByScore(startScore, endScore float64) (removed uint) {
	spec := &zRangeSpec{
		min:        startScore,
		max:        endScore,
		includeMin: 0,
		includeMax: 0,
	}
	z.lock.Lock()
	defer z.lock.Unlock()
	removed = z.skipList.zslDeleteRangeByScore(spec, z.dict)
	return removed
}

//删除某个排位区间（闭区间，包含两端）的节点，并返回移除了多少个
func (z *zSet) DeleteByRank(startRank, endRank uint) (removed uint) {
	if startRank == 0 {
		startRank = 1
	}
	if endRank > z.skipList.length {
		endRank = z.skipList.length
	}
	// 分数越高，排名越高， 要 从高 往低计算排名
	startRank = z.skipList.length - startRank + 1
	endRank = z.skipList.length - endRank + 1
	if startRank > endRank {
		startRank, endRank = endRank, startRank
	}
	z.lock.Lock()
	defer z.lock.Unlock()
	removed = z.skipList.zslDeleteRangeByRank(startRank, endRank, z.dict)
	return
}

// 改
//对指定成员的分数加上增量increment
func (z *zSet) IncrBy(member string, increment float64) (bool, error) {
	z.lock.Lock()
	defer z.lock.Unlock()

	data, ok := z.dict[member]
	if !ok {
		return false, errNoKey
	}
	z.skipList.zslUpdateScore(data.score, member, data.score+increment)
	data.score = data.score + increment
	return true, nil
}

// 更新指定节点的分数
func (z *zSet) UpdateScore(member string, score float64) (ok bool, err error) {
	z.lock.Lock()
	defer z.lock.Unlock()

	data, ok := z.dict[member]
	if !ok {
		return false, errNoKey
	}
	if data.score == score {
		return false, nil
	}
	_ = z.skipList.zslUpdateScore(data.score, member, score)
	// map中也维护相应的分数，用于快速判断是否相等
	data.score = score
	return true, nil
}

// 更新额外数据
func (z *zSet) UpdateAddData(member string, data interface{}) (ok bool, err error) {
	z.lock.Lock()
	defer z.lock.Unlock()

	ele, ok := z.dict[member]
	if !ok {
		return false, errNoKey
	}
	ele.data = data
	return true, nil
}

// 查
// 获取指定 key的排名
func (z *zSet) GetRank(members ...string) (rankInfo map[string]uint, err error) {
	z.lock.Lock()
	defer z.lock.Unlock()

	rankInfo = make(map[string]uint)
	var rank uint
	for _, m := range members {
		ele, ok := z.dict[m]
		if !ok {
			err = errNoKey
			return
		}
		rank = z.skipList.zslGetRank(ele.score, ele.mem)
		rank = z.skipList.length - rank + 1
		rankInfo[m] = rank
	}
	return
}

func (z *zSet) GetElementByMem(members ...string) (nodes []Node, err error) {
	z.lock.Lock()
	defer z.lock.Unlock()

	for _, m := range members {
		ele, ok := z.dict[m]
		if !ok {
			err = fmt.Errorf("no mem: %s", m)
			nodes = nil
			return
		}
		nodes = append(nodes, ele)
	}
	return
}

func (z *zSet) GetElementByRank(ranks ...uint) (nodes []Node, err error) {
	z.lock.Lock()
	defer z.lock.Unlock()
	for _, rank := range ranks {
		r := z.skipList.length - rank + 1
		sNode := z.skipList.zslGetElementByRank(r)
		if sNode == nil {
			nodes = nil
			err = fmt.Errorf("invalid rank: %d", rank)
			return
		}
		data, ok := z.dict[sNode.eleKey]
		if !ok {
			nodes = nil
			z.skipList.zslDelete(sNode.score, sNode.eleKey)
			err = fmt.Errorf("invalid rank: %d", rank)
			return
		}
		nodes = append(nodes, data)
	}
	return
}

// 返回 指定排名区间的 元素
func (z *zSet) SafeRange(startRank, endRank uint, isReverse bool, do func(mem string, score float64, addData interface{}) error) error {
	z.lock.Lock()
	defer z.lock.Unlock()

	err := z.UnSafeRange(startRank, endRank, isReverse, do)

	return err
}

// isReverse : false 是否倒叙输出
// startRank: 默认是从第一名开始，不是第0名
func (z *zSet) UnSafeRange(startRank, endRank uint, isReverse bool, do func(mem string, score float64, addData interface{}) error) error {
	if z.skipList.length == 0 {
		return nil
	}
	if startRank == 0 {
		startRank = 1
	}
	if endRank > z.skipList.length {
		endRank = z.skipList.length
	}
	if startRank > endRank {
		startRank, endRank = endRank, startRank
	}
	loopCnt := endRank - startRank + 1
	var node *skipNode
	switch isReverse {
	case true:
		node = z.skipList.zslGetElementByRank(startRank)
	default:
		node = z.skipList.zslGetElementByRank(endRank)
	}
	if node == nil {
		return errInvalidRank
	}
	for loopCnt > 0 {
		ele := z.dict[node.eleKey]
		if err := do(node.eleKey, node.score, ele.data); err != nil {
			return err
		}
		switch isReverse {
		case true:
			node = node.level[0].forward
		default:
			node = node.backward
		}
		loopCnt--
	}
	return nil
}
