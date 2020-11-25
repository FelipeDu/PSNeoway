package main

import (
	"strings"
	"container/list"
	"time"
	"log"
)

type Processor interface {
	ProcessLine(line string)(*list.List)
}

func ProcessLine(id int64, line string)(*list.List){
	processedLine := strings.Fields(line)
	var lineFields = list.New()
	lineFields.PushBack(id)

	var isValid bool = true
	for i, v := range processedLine{
		switch {
			case i == 0:
				var processedDocument string
				processedDocument, isValid = ProcessDocument(v)
				lineFields.PushBack(processedDocument)
			case (i == 1 || i == 2):
				if(v == "null" || v == ""){
					isValid = false
					lineFields.PushBack(nil)
				} else if(v == "0"){
					lineFields.PushBack(false)
				} else {
					lineFields.PushBack(true)
				}
			case i == 3:
				if(v == "null" || v == ""){
					lineFields.PushBack(nil)
					isValid = false
				} else {
					processedDate, err := time.Parse("1800-01-01",v)
					if(err != nil){
						isValid = false
						lineFields.PushBack(nil)
						log.Printf("Invalid Date Format")
					} else {
						lineFields.PushBack(processedDate)
					}
				}
			case (i == 4 || i == 5):
		}
	}
	lineFields.PushBack(isValid)
	return lineFields
}

func ProcessDocument(document string)(string,bool){
	processedDocument := "TESTE"
	return processedDocument, false
}
