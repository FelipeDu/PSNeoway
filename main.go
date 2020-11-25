package main

import (
//	"fmt"
	"log"
	"os"
)

func main(){

	PersistFile(os.Args[1])

	log.Printf("DONE\n")
}