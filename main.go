package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

const (
	host     	= "localhost"
	port     	= 5432
	user     	= "service"
	password 	= "service"
	dbname   	= "dbneoway"
	tableName = "client_purchase_registry"
)

func main(){

	pathToFile := os.Args[1]

	start := time.Now()
	numArquivos := 10
	dbase = ConnectToDB()

	for i := 1; i <= numArquivos; i++{
		PersistFile(pathToFile, dbase)
	}

	end := time.Now()
	delta := end.Sub(start)
	CloseConnection(dbase)
	fmt.Printf("Inserido %d arquivos em %.2f\n", numArquivos, delta.Seconds())

	log.Printf("DONE\n")
}