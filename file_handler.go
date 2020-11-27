package main

import (
  "log"
  "os"
  "bufio"
  "database/sql"
  "sync"
)

var lastID int64
var containsHeader bool = true
const maxBulkSize = 10000

type Handler interface {
  PersistFile(fileLocation string, dbase *sql.DB)
}

func PersistFile(fileLocation string, dbase *sql.DB){

  lastID = GetLastID(dbase)
  var file, err = LoadFile(fileLocation)
  if err != nil {
    log.Fatal(err)
  }

  err = ParseAndInsert(file)
  if err != nil {
    log.Fatal(err)
  }

  err = CloseFile(file)
  if err != nil {
    log.Fatal(err)
  }
}

func LoadFile (fileLocation string) (*os.File, error) {
  var file, err = os.Open(fileLocation)
  if err != nil {
    return nil, err
  }
  return file, err
}

func CloseFile (file *os.File) (error) {
  err := file.Close()
  if err != nil {
    log.Fatal(err)
  }
  return err
}

func ParseAndInsert (file *os.File) (error) {

  var wg sync.WaitGroup

  bufferReader := bufio.NewReader(file)
  EOF := false
  id := lastID

  if(containsHeader){
    _, err := bufferReader.ReadString('\n')
    if(err != nil){
      if(err.Error() == "EOF"){
        EOF = true
      } else {
        return err
      }
    }
  }

  var bulkRegistry []Registry
  currentBulkSize := 0
  for !EOF {
    id++
    bufferedString, err := bufferReader.ReadString('\n')
    if(err != nil){
      if(err.Error() == "EOF"){
        EOF = true
      } else {
        return err
      }
    }
    bulkRegistry = append(bulkRegistry, ProcessLine(id,bufferedString," "))
    currentBulkSize++
    if(currentBulkSize == maxBulkSize || EOF){
      wg.Add(1)
      go BulkSendToDB(bulkRegistry,dbase,&wg)
      currentBulkSize = 0
      bulkRegistry = nil
    }
  }

  wg.Wait()

  return nil
}