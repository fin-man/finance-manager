package main

import (
	"finance-manager/filewatcher/watcher"
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

	fw.Watch(fullPath)
}