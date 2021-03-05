package abstract

import "fmt"

type Lunch interface {
	Cook()
}

type SimpleFactory interface {
	CreateFood() Lunch
	CreateSoup() Lunch
}

// 创建
type Potato struct{}

func (p *Potato) Cook() {
	fmt.Println("cook potato")
}

type tomato struct{}

func (t *tomato) Cook() {
	fmt.Println("cook tomato")
}

type Restaurant struct{}

func (r *Restaurant) CreateFood() Lunch {
	return &tomato{}
}

func (r *Restaurant) CreateSoup() Lunch {
	return &Potato{}
}

// 抽象工厂？
func NewRestaurant() SimpleFactory {
	return &Restaurant{}
}
