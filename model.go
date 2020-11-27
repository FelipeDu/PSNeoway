package main

import (
  "time"
)

//struct com os fields que serão inseridos na tabela
type Registry struct{
  ID int64
  PersonCompanyDocument string
  ValidDocument bool
  Private bool
  Incomplete bool
  DateLastPurchase time.Time
  MedianTicket float64
  LastTicket float64
  FrequentStore string
  ValidFrequentStore bool
  LastStore string
  ValidLastStore bool
  ValidRegistry bool
}