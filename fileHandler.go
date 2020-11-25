package main

import (
//	"fmt"
	"log"
	"os"
	"bufio"
//	"errors"
//	"io/ioutil"
)

type Handler interface {
	//PersistFile(fileLocation string)
}

func PersistFile(fileLocation string){

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
		//log.Fatal(err)
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

	bufferReader := bufio.NewReader(file)
	EOF := false
	lineNumber := 0
	for !EOF {
		bufferedString, err := bufferReader.ReadString('\n')
		if(err != nil){
			if(err.Error() == "EOF"){
				EOF = true
			} else {
				return err
			}
		}
		lineNumber++
		log.Printf("NUMBER: %d", lineNumber)
		log.Printf("LINE: %s", bufferedString)
	}

	return nil
}