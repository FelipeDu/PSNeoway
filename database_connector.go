package main

import (
	"database/sql"
	"github.com/lib/pq"
	"log"
	"fmt"
	"strconv"
	"strings"
)
/*
const (
	host     = "localhost"
	port     = 5432
	user     = "admin"
	password = "admin"
	dbname   = "dbneoway"
)
*/

type DBConnector interface {
	ConnectToDB()
	GetLastID()(int64)
	CloseConnection()
}

var dbase *sql.DB

func ConnectToDB(){

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	var err error

	dbase, err = sql.Open("postgres", psqlconn)
	if(err != nil){
		log.Fatal(err)
	}
	
	dbase.SetMaxIdleConns(10)
	dbase.SetMaxOpenConns(10)
	dbase.SetConnMaxLifetime(0)

}

func GetLastID()(int64){
	prepQuery := fmt.Sprintf("select max(id) from %s",tableName)
	line, err := dbase.Query(prepQuery)
	if(err != nil){
		//errorString := fmt.Sprintf("pq: relation \"%s\" does not exist",tableName)
		errorString := "does not exist"
		if(strings.Contains(err.Error(), errorString)){
			log.Printf("Table \"%s\" does not exist. Creating Table",tableName)
			err = CreateTable()
			if(err != nil){
				log.Fatal(err)
			}
			return 0
		}
		log.Fatal(err)
	}

	var lastID int64
	for line.Next(){
		var maxID string
		err = line.Scan(&maxID)
		if(err != nil){
			log.Fatal(err)
		}
		lastID,_ = strconv.ParseInt(maxID,10,64)
	}

	return lastID
}

func CreateTable()(error){
	prepTableCreation := fmt.Sprintf("create table %s (id integer primary key, documento varchar(14), private boolean, incomplete boolean, dateOfLastPurchase date, medianTicket numeric, lastTicket numeric, frequentStore char(14), lastStore char(14), isValid boolean)",tableName)
	_, err := dbase.Query(prepTableCreation)
	return err
}

func CloseConnection(){
	dbase.Close()
}

func BulkSendToDB(bulkRegistry []Registry) (error){

	trsc, err := dbase.Begin()
	if(err != nil){
		log.Fatal(err)
	}

	stmt, err := trsc.Prepare(pq.CopyIn(tableName, "id", "documento", "private", "incomplete", "dateOfLastPurchase", "medianTicket", "lastTicket", "frequentStore", "lastStore", "isValid"))
	if(err != nil){
		log.Fatal(err)
	}

	for i := range bulkRegistry{
		fields := bulkRegistry[i]
		_, err := stmt.Exec(fields.ID,
												fields.PersonCompanyDocument,
												fields.Private,
												fields.Incomplete,
												fields.DateLastPurchase,
												fields.MedianTicket,
												fields.LastTicket,
												fields.FrequentStore,
												fields.LastStore,
												fields.IsValid)
		if err != nil {
			log.Fatal(err)
		}
	}

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

	return nil
}