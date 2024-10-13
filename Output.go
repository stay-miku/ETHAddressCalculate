package main

import (
	"os"
	"sync"
)

var file *os.File
var outputMutex sync.Mutex

func initOutput() error {
	if _, err := os.Stat(Config.Output); os.IsExist(err) {
		file, err = os.OpenFile(Config.Output, os.O_APPEND|os.O_WRONLY, 0644)
		return err
	} else {
		file, err = os.OpenFile(Config.Output, os.O_CREATE|os.O_WRONLY, 0644)
		return err
	}
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
