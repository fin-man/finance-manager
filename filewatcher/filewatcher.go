package main

import (
	"finance-manager/filewatcher/watcher"
	"fmt"
	"log"
	"os"
)

func main() {
	log.Println("Starting a new filewatcher ")
	fw := watcher.NewFileWatcher()

	pwd, err := os.Getwd()

	if err != nil {
		panic(err)
	}
	fullPath := pwd + "/data/transactions"

	fw.Watch(fullPath, ProcessFile)
}

func ProcessFile(data ...interface{}) error {
	fmt.Printf("FilePath : %s \n", data[0])
	fmt.Printf("FileName : %s \n", data[1])

	return nil
}
