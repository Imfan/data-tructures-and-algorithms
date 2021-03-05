package factory

import "testing"

func TestNewRestaurant(t *testing.T) {
	NewRestaurant("a").GetFood()
	NewRestaurant("b").GetFood()
}
