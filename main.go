package main

import (
//	"strings"
//	"container/list"
//	"fmt"
	"log"
//	"os"
//	"database/sql"
//	"errors"
)

const (
	host     	= "localhost"
	port     	= 5432
	user     	= "service"
	password 	= "service"
	dbname   	= "dbneoway"
	tableName = "registroCompras"
)

func main(){

	PersistFile("arquivoTeste/test_file.txt")

	log.Printf("DONE\n")
}