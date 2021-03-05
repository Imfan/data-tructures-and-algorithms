package factory

import "fmt"

type Restaurant interface {
	GetFood()
}

type A struct {
}

func (a *A) GetFood() {
	fmt.Println("A饭店的食物")
}

type B struct {
}

func (b *B) GetFood() {
	fmt.Println("B饭店的食物")
}

func NewRestaurant(name string) Restaurant {
	switch name {
	case "a":
		return &A{}
	case "b":
		return &B{}
	}
	return nil
}
