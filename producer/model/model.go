package model

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