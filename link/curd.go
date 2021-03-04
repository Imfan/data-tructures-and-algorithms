package link

type List struct {
	next  *List
	value int
}

func Init(val int) *List {
	return &List{
		value: val,
	}
}

func (this *List) AddToHead(val int) *List {
	return &List{
		next:  this,
		value: val,
	}
}

func (this *List) AddToTail(val int) *List {
	l := &List{
		value: val,
	}
	this.next = l
	return l
}

func (this *List) UpdateNextPointer(l *List) {
	this.next = l
}
