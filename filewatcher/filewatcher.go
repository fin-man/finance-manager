package main

import (
	"finance-manager/csvprocessors"
	"finance-manager/filemanager"
	"finance-manager/filewatcher/watcher"
	"finance-manager/transactionstypes"
	"fmt"
	"log"
	"os"
	"strings"
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
	fileName := data[1].(string)
	filePath := data[0].(string)

	if strings.Contains(fileName, "chase") {
		log.Println("Detected a new Chase file")

		err := HandleChase(filePath)

		if err != nil {
			log.Println(err)
			return err
		}

		return nil

	} else if strings.Contains(fileName, "capital_one") {
		log.Println("Detected a new Chase file")
		err := HandleCapitalOne(filePath)

		if err != nil {
			log.Println(err)
			return err
		}

		return nil
	}

	log.Printf("Unknown File Found ..")
	//TO DO alert

	return nil
}

func HandleChase(filePath string) error {

	fm := filemanager.FileManager{}

	file, err := fm.OpenFile(filePath, os.O_RDWR|os.O_CREATE, os.ModePerm)

	if err != nil {
		return err
	}

	records := []*transactionstypes.ChaseTransaction{}
	chaseClient := csvprocessors.NewChaseClient()

	err = chaseClient.Unmarshal(file, &records)

	if err != nil {
		return err
	}

	normalizedRecords := chaseClient.ProcessCSV(records)
	for _, v := range normalizedRecords {
		fmt.Printf("ChaseRecord : %v \n", v)
	}

	return nil
}

func HandleCapitalOne(filePath string) error {
	fm := filemanager.FileManager{}

	file, err := fm.OpenFile(filePath, os.O_RDWR|os.O_CREATE, os.ModePerm)

	if err != nil {
		return err
	}

	records := []*transactionstypes.CapitalOneTransaction{}
	capitalOneClient := csvprocessors.NewCapitalOneClient()

	err = capitalOneClient.Unmarshal(file, &records)

	if err != nil {
		return err
	}

	normalizedRecords := capitalOneClient.ProcessCSV(records)
	for _, v := range normalizedRecords {
		fmt.Printf("CapitalOneRecord : %v \n", v)
	}

	return nil
}