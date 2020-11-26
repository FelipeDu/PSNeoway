package main

import (
	"time"
)

//struct com os fields que serão inseridos na tabela
type Registry struct{
	ID int64
	PersonCompanyDocument string
	Private bool
	Incomplete bool
	DateLastPurchase time.Time
	MedianTicket float64
	LastTicket float64
	FrequentStore string
	LastStore string
	IsValid bool
}