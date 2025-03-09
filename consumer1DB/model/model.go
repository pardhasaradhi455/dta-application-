package model

import "encoding/json"

type Order struct {
	OrderId string
	OrderDate string
	OrderTotal float64
	Status Status
	Address Address
}

type Address struct {
	City, State string
	Pincode int
}

type Status struct {
	State string
	StatusDate string
}

func (o *Order) ToBytes() []byte{
	jsonData, _ := json.Marshal(o)
	return jsonData
}

func ByteToOrder(orderBytes []byte) Order{
	var order Order
	json.Unmarshal(orderBytes, &order)
	return order
}