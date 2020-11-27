package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"
	"github.com/lib/pq"
)

type DBConnector interface {
	ConnectToDB()
	GetLastID(*sql.DB) int64
	CloseConnection(*sql.DB)
	BulkSendToDB([]Registry, *sql.DB, *sync.WaitGroup) error
}

var dbase *sql.DB

func ConnectToDB(uri string) *sql.DB {

	// psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	var err error

	dbase, err = sql.Open("postgres", uri)
	if err != nil {
		log.Fatal(err)
	}

	dbase.SetMaxIdleConns(10)
	dbase.SetMaxOpenConns(10)
	dbase.SetConnMaxLifetime(0)

	return dbase
}

func GetLastID(dbase *sql.DB) int64 {
	tableName := "client_purchase_registry"
	prepQuery := fmt.Sprintf("select max(id) from %s", tableName)
	line, err := dbase.Query(prepQuery)
	if err != nil {
		errorString := fmt.Sprintf("relation \"%s\" does not exist", tableName)
		if strings.Contains(err.Error(), errorString) {
			log.Printf("[INFO] Table \"%s\" does not exist. Creating Table", tableName)
			err = CreateTable(dbase)
			if err != nil {
				log.Fatal(err)
			}
			return 0
		}
		log.Fatal(err)
	}

	var lastID int64 = 0
	for line.Next() {
		var maxID sql.NullString
		err = line.Scan(&maxID)
		if err != nil {
			log.Print(err)
		}
		if maxID.Valid {
			lastID, err = strconv.ParseInt(maxID.String, 10, 64)
			if err != nil {
				log.Print(err)
			}
			//log.Printf("[INFO] LastId Obtained Successfully: %d", lastID)
		}
	}

	return lastID
}

func CreateTable(dbase *sql.DB) error {
	tableName := "client_purchase_registry"
	prepTableCreation := fmt.Sprintf("create table %s (id integer primary key, document varchar(14) not null, valid_document boolean, private boolean, incomplete boolean, date_of_last_purchase date, median_ticket numeric, last_ticket numeric, frequent_store char(14), valid_frequent_store boolean, last_store char(14), valid_last_store boolean, valid_registry boolean)", tableName)
	_, err := dbase.Query(prepTableCreation)
	return err
}

func CloseConnection(dbase *sql.DB) {
	dbase.Close()
}

func BulkSendToDB(bulkRegistry []Registry, dbase *sql.DB, wg *sync.WaitGroup) error {
	tableName := "client_purchase_registry"

	trsc, err := dbase.Begin()
	if err != nil {
		wg.Done()
		return err
	}

	stmt, err := trsc.Prepare(pq.CopyIn(tableName, "id", "document", "valid_document", "private", "incomplete", "date_of_last_purchase", "median_ticket", "last_ticket", "frequent_store", "valid_frequent_store", "last_store", "valid_last_store", "valid_registry"))
	if err != nil {
		wg.Done()
		return err
	}

	for i := range bulkRegistry {
		fields := bulkRegistry[i]
		_, err := stmt.Exec(fields.ID,
			fields.PersonCompanyDocument,
			fields.ValidDocument,
			fields.Private,
			fields.Incomplete,
			fields.DateLastPurchase,
			fields.MedianTicket,
			fields.LastTicket,
			fields.FrequentStore,
			fields.ValidFrequentStore,
			fields.LastStore,
			fields.ValidLastStore,
			fields.ValidRegistry)
		if err != nil {
			wg.Done()
			return err
		}
	}

	_, err = stmt.Exec()
	if err != nil {
		wg.Done()
		return err
	}

	err = stmt.Close()
	if err != nil {
		wg.Done()
		return err
	}

	err = trsc.Commit()
	if err != nil {
		wg.Done()
		return err
	}

	wg.Done()

	return nil
}
