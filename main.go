package main

import (
	"fmt"
	"log"
	"os"
	"time"
	"strconv"
)

// const (
//	 host			= "localhost"
//	 port			= 5432
//	 user			= "service"
//	 password	= "service"
//	 dbname		= "postgres"
//	 tableName = "client_purchase_registry"
// )
// URI:
// "postgres://service:service@localhost:5432/dblog?sslmode=disable"

func main() {

	fmt.Println("")
	fmt.Println("")

	dbURI := os.Getenv("DB_URI")
	pathToFile := os.Getenv("FILE")

	if dbURI == "" {
		log.Fatal("[ERROR] Please set database connection URI env")
	}
	
	start := time.Now()
	numExecutions := 1
	if(os.Getenv("NUM_EXECUTIONS") != ""){
		varExecutions, err := strconv.ParseInt(os.Getenv("NUM_EXECUTIONS"),10, 0)
		if(err == nil){
			numExecutions = int(varExecutions)
		}
	}

	dbase = ConnectToDB(dbURI)

	for i := 1; i <= numExecutions; i++ {
		PersistFile(pathToFile, dbase)
	}
	
	fmt.Println("")

	end := time.Now()
	delta := end.Sub(start)
	CloseConnection(dbase)
	fmt.Printf("[INFO] Inserido %d arquivo(s) em %.2f segundos.\n", numExecutions, delta.Seconds())

	log.Printf("DONE\n")
	
	fmt.Println("")
	fmt.Println("")
}
