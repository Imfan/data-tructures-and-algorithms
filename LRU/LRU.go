package LRU

// 最近最少使用 淘汰法

type Cache struct {
	hash map[string]*link
	head, tail *link
	len, cap int
}

type link struct {
	prev, next *link
	key, val string
}


func (c *Cache) Set(k, v string) (string,bool) {
	return c.addToLinkHead(k, v)
}

// 增加节点到 链表头部
func (c *Cache) addToLinkHead(k, v string) (string, bool) {
	l, ok := c.hash[k]
	if  ok {
		// 如果存在，则更新，并放到头部
		c.updateNode(l, v)
		c.moveToLinkHead(l.key)
	}else {
		// 如果长度大于等于容量，等于容量说明多了一个，移除尾部
		if c.len >= c.cap {
			c.delLinkTail()
		}
		// 新增节点
		l = c.addNode(k,v, c.head)
		c.hash[k] = l
	}
	return l.val, true
}

// 将原有绑定解除，绑定到头部
func (c *Cache) moveToLinkHead(k string)  {
	node := c.hash[k]
	c.unBound(node)
	c.boundNode(node, c.head)
}

func (c *Cache) Get(k string) (string, bool)  {
	if link, ok := c.hash[k]; ok {
		c.moveToLinkHead(k)
		return link.val, true
	}
	return "", false
}

func (c *Cache) Delete(k string) bool  {
	c.removeNode(k)
	return  c.delHash(k)
}

func NewCache(cap int) *Cache  {
	if cap <= 0 {
		cap = 1
	}
	h := make(map[string]*link, cap)
	head := initNode("", "")
	tail := initNode("", "")
	head.next = tail
	tail.prev = head
	return &Cache{
		hash: h,
		head: head,
		tail: tail,
		cap: cap,
	}
}

// 初始化一个节点
func initNode(k, v string) *link  {
	return &link{
		key: k,
		val: v,
	}
}
// 新增节点
func (c *Cache) addNode(k, v string, node *link) *link  {
	l := initNode(k,v)
	// 绑定
	c.boundNode(l, node)
	c.len++
	return l
}
// 移除节点
func (c *Cache) removeNode(k string) {
	if l, ok := c.hash[k]; ok {
		c.unBound(l)
		c.len--
	}
}
// 更新节点
func (c *Cache) updateNode(l *link, v string) {
	l.val = v
}

// 解除绑定
func (c *Cache) unBound(l *link)  {
	l.next.prev = l.prev
	l.prev.next = l.next
	l.next = nil
	l.prev = nil
}
// 将node 绑定到 boundTo节点后面
func (c *Cache) boundNode(node, boundTo *link)  {
	boundTo.next.prev = node
	node.next = boundTo.next
	boundTo.next = node
	node.prev = boundTo
}

// 删除尾部节点
func (c *Cache) delLinkTail() bool {
	k := c.tail.prev.key
	c.removeNode(k)
	return c.delHash(k)
}
// 删除hash
func (c *Cache) delHash(k string) bool {
	delete(c.hash, k)
	_, ok := c.hash[k]
	return !ok
}






