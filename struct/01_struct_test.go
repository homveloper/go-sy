package struct_test

import (
	"fmt"
	"testing"
)

type HouseType uint8

const (
	Apartment HouseType = iota
)

type House struct {
	address   string
	size      int
	price     float64
	houseType HouseType
}

func TestStruct(t *testing.T) {
	var house House
	house.address = "서울 강남구 테헤란로 419"
	house.size = 28
	house.price = 9.8
	house.houseType = Apartment

	fmt.Println(house)

	var house1 House = House{
		address:   "서울 강남구 테헤란로 419",
		size:      28,
		price:     9.8,
		houseType: Apartment,
	}

	fmt.Println(house1)

}
