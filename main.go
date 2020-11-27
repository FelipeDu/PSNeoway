package main

import (
  "fmt"
  "log"
  "os"
  "time"
)

// const (
//   host      = "localhost"
//   port      = 5432
//   user      = "service"
//   password  = "service"
//   dbname    = "postgres"
//   tableName = "client_purchase_registry"
// )

func main() {

  dbURI := os.Getenv("DB_URI")

  if dbURI == "" {
    log.Fatal("[ERROR] Please set database connection URI env")
  }

  pathToFile := os.Args[1]

  start := time.Now()
  numArquivos := 10
  dbase = ConnectToDB(dbURI)

  for i := 1; i <= numArquivos; i++ {
    PersistFile(pathToFile, dbase)
  }

  end := time.Now()
  delta := end.Sub(start)
  CloseConnection(dbase)
  fmt.Printf("Inserido %d arquivos em %.2f\n segundos.", numArquivos, delta.Seconds())

  log.Printf("DONE\n")
}
