package main

import (
	"os"
	"sync"
)

var file *os.File
var outputMutex sync.Mutex

func initOutput() error {
	var err error
	file, err = os.OpenFile("result.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	return err
}

func writeOutput(data string) error {
	outputMutex.Lock()
	_, err := file.WriteString(data)
	outputMutex.Unlock()
	return err
}

func writeResult(privateKey, address string) error {
	return writeOutput(privateKey + "," + address + "\n")
}
