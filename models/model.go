package models

import (
	"github.com/google/uuid"
)

type Customer struct {
	CustomerId    uuid.UUID `json:"customeruuid" db:"customeruuid"`
	FirstName     string    `json:"firstname" db:"firstname"`
	MiddleInitial string    `json:"middleinitial" db:"middleinitial"`
	LastName      string    `json:"lastname" db:"lastname"`
}

type Employee struct {
	employeeId    uuid.UUID
	firstName     string
	middleInitial string
	lastName      string
}

type Product struct {
	productId uuid.UUID
	name      string
	price     float64
}

type Sale struct {
	saleId       uuid.UUID
	salePersonId int64
	customerId   int64
	productId    int64
	quantity     int64
}
