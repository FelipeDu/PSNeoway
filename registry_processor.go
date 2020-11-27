package main

import (
	"strings"
	//	"container/list"
	"log"
	"regexp"
	"strconv"
	"time"
)

type Processor interface {
	ProcessLine(id int64, line string, separator string) Registry
}

func ProcessLine(id int64, line string, separator string) Registry {

	var processedLine []string
	if separator == " " {
		processedLine = strings.Fields(line)
	} else {
		processedLine = strings.Split(line, separator)
	}

	return CreateRegistry(id, processedLine)

}

//Cria um registro com os valores da linha sendo analisada
func CreateRegistry(id int64, processedLine []string) Registry {

	var registry Registry
	registry.ID = id
	registry.ValidRegistry = true //Registro válido
	registry.ValidDocument = true
	registry.ValidFrequentStore = true
	registry.ValidLastStore = true

	var isValid bool = true
	for i, v := range processedLine {
		switch i {
		case 0: //Documento
			var processedDocument string
			if v == "null" || v == "" {
				registry.ValidRegistry = false
			} else {
				processedDocument, isValid = ProcessDocument(v)
				if !isValid {
					registry.ValidDocument = false
					registry.ValidRegistry = false
				}
				registry.PersonCompanyDocument = processedDocument
			}
		case 1: //PRIVATE
			if v == "null" || v == "" {
				registry.ValidRegistry = false
			} else if v == "0" {
				registry.Private = false
			} else {
				registry.Private = true
			}
		case 2: //INCOMPLENTO
			if v == "null" || v == "" {
				registry.ValidRegistry = false
			} else if v == "0" {
				registry.Incomplete = false
			} else {
				registry.Incomplete = true
			}
		case 3: //Data da ultima compra
			if v == "null" || v == "" {
				registry.ValidRegistry = false
			} else {
				processedDate, err := time.Parse("2006-01-02", v)
				if err != nil {
					registry.ValidRegistry = false
					//log.Printf("Invalid Date Format: %s",v)
				} else {
					registry.DateLastPurchase = processedDate
				}
			}
		case 4: //Média Ticket
			if v == "null" || v == "" {
				registry.ValidRegistry = false
			} else {
				processedValue, err := strconv.ParseFloat(strings.ReplaceAll(strings.ReplaceAll(v, ".", ""), ",", "."), 64)
				if err != nil {
					registry.ValidRegistry = false
				} else {
					registry.MedianTicket = processedValue
				}
			}
		case 5: //Ultimo Ticket
			if v == "null" || v == "" {
				registry.ValidRegistry = false
			} else {
				processedValue, err := strconv.ParseFloat(strings.ReplaceAll(strings.ReplaceAll(v, ".", ""), ",", "."), 64)
				if err != nil {
					registry.ValidRegistry = false
				} else {
					registry.LastTicket = processedValue
				}
			}
		case 6: //Loja mais frequente
			var processedDocument string
			if v == "null" || v == "" {
				registry.ValidRegistry = false
			} else {
				processedDocument, isValid = ProcessDocument(v)
				if !isValid {
					registry.ValidFrequentStore = false
					registry.ValidRegistry = false
				}
				registry.FrequentStore = processedDocument
			}
		case 7: //Ultima Loja
			var processedDocument string
			if v == "null" || v == "" {
				registry.ValidRegistry = false
			} else {
				processedDocument, isValid = ProcessDocument(v)
				if !isValid {
					registry.ValidLastStore = false
					registry.ValidRegistry = false
				}
				registry.LastStore = processedDocument
			}
		}
	}
	return registry
}

//"Limpa" e verifica o tipo de documento
func ProcessDocument(document string) (string, bool) {

	regRule, err := regexp.Compile("[^0-9]+")
	if err != nil {
		log.Print(err)
	}
	processedDocument := regRule.ReplaceAllString(document, "")
	stringSize := len(processedDocument)

	if stringSize == 11 { //Validação dos dígitos verificadores
		return processedDocument, VerifyCPF(processedDocument)
	} else if stringSize == 14 {
		return processedDocument, VerifyCNPJ(processedDocument)
	} else {
		//log.Println(stringSize, document, processedDocument)
		return processedDocument, false
	}
}

//Valida CPF, recebe a string com os dígitos "limpa"
func VerifyCPF(document string) bool {
	firstDigitsSum := 0
	secondDigitSum := 0
	for i := 0; i < 9; i++ {
		firstDigitsSum += (int(document[i]-'0') * (i + 1))
		secondDigitSum += (int(document[i]-'0') * i)
	}

	firstDigitVerifier := (firstDigitsSum % 11) % 10
	if firstDigitVerifier != int(document[9]-'0') {
		return false
	}

	secondDigitSum += firstDigitVerifier * 9
	secondDigitVerifier := (secondDigitSum % 11) % 10
	if secondDigitVerifier != int(document[10]-'0') {
		return false
	}
	return true
}

//Valida CNPJ, recebe a string com os dígitos "limpa"
func VerifyCNPJ(document string) bool {
	referenceDigitsCNPJ := [13]int{5, 6, 7, 8, 9, 2, 3, 4, 5, 6, 7, 8, 9}
	firstDigitsSum := 0
	secondDigitSum := 0
	for i := 0; i < 12; i++ {
		firstDigitsSum += (int(document[i]-'0') * referenceDigitsCNPJ[(i+1)])
		secondDigitSum += (int(document[i]-'0') * referenceDigitsCNPJ[i])
	}

	firstDigitVerifier := firstDigitsSum % 11
	if firstDigitVerifier == 10 {
		firstDigitVerifier = 0
	}
	if firstDigitVerifier != int(document[12]-'0') {
		return false
	}

	secondDigitSum += firstDigitVerifier * referenceDigitsCNPJ[12]
	secondDigitVerifier := secondDigitSum % 11
	if secondDigitVerifier == 10 {
		secondDigitVerifier = 0
	}
	if secondDigitVerifier != int(document[13]-'0') {
		return false
	}
	return true
}
