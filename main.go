package main

import (
	"strings"
//	"container/list"
	"fmt"
	"log"
//	"os"
//	"database/sql"
	_ "github.com/lib/pq"
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


	stringTest := "1     ad  451,45      0"
	campos := strings.Fields(stringTest)

	for i, v := range campos{
		fmt.Println(i, v)
	}

	fmt.Println(campos[0])
	fmt.Println(campos[1])
	fmt.Println(campos[2])
	fmt.Println(campos[3])

//	PersistFile(os.Args[1])
/*
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	dbase, err := sql.Open("postgres", psqlconn)
	if(err != nil){
		log.Fatal(err)
	}

	dbase.SetMaxIdleConns(10)
	dbase.SetMaxOpenConns(10)
	dbase.SetConnMaxLifetime(0)
	err = dbase.Ping()
	if(err != nil){
		log.Fatal(err)
	}

	prepQuery := fmt.Sprintf("select max(id) from %s",tableName)
	lines, err := dbase.Query(prepQuery)
	if(err != nil){
		log.Printf(err.Error())
		if(err.Error() == "pq: relation \"registrocompras\" does not exist"){
			log.Printf("YEP")
			log.Fatal(err.Error())
		}else{
			log.Fatal(err.Error())
		}
	}
	for lines.Next(){
		var linha string
		err = lines.Scan(&linha)
		if(err != nil){
			log.Fatal(err)
		}
		log.Printf(linha)
	}

	trsc, err := dbase.Begin()
	if(err != nil){
		log.Fatal(err)
	}

	////
	
	stmt, err := trsc.Prepare(pq.CopyIn("testcase", "id", "name", "whatever")) // MessageDetailRecord is the table name
	if(err != nil){
		log.Fatal(err)
	}

	for id := 1; id <= 1000; id++{
		_, err := stmt.Exec(id, "FELIPE", "WHATEVER")
		if err != nil {
			log.Fatal(err)
		}
	}

	////

	_, err = stmt.Exec()
	if(err != nil){
		log.Fatal(err)
	}

	err = stmt.Close()
	if(err != nil){
		log.Fatal(err)
	}

	err = trsc.Commit()
	if(err != nil){
		log.Fatal(err)
	}

	err = dbase.Close()
	if(err != nil){
		log.Fatal(err)
	}
*/
	log.Printf("DONE\n")
}