package main

import (
	"finance-manager/categories"
	"finance-manager/clients/recordcreator"
	"finance-manager/csvprocessors"
	"finance-manager/filemanager"
	"finance-manager/filewatcher/watcher"
	"finance-manager/transactionstypes"
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
	// fileName := data[1].(string)
	filePath := data[0].(string)

	recordCreator := recordcreator.NewRecordCreator()

	err := HandleOverall(filePath, recordCreator)
	if err != nil {
		return err
	}

	log.Printf("Unknown File Found ..")

	return nil
}

func HandleOverall(filePath string, recordCreator *recordcreator.RecordCreator) error {
	fm := filemanager.FileManager{}
	file, err := fm.OpenFile(filePath, os.O_RDWR|os.O_CREATE, os.ModePerm)

	if err != nil {
		return err
	}

	defer file.Close()

	records := []*categories.NormalizedTransaction{}

	csvProcessor := csvprocessors.NewCSVprocessor()

	err = csvProcessor.Unmarshal(file, &records)

	fmt.Println(records)
	if err != nil {

		//file is prolly dont match the format
		return err
	}

	for _, v := range records {

		//	_, ok := categories.OverallTransactionTypes[string(v.Category)]
		err = recordCreator.CreateNewRecord(v)
		if err != nil {
			log.Println(err)
		}
	}

	return nil

}

func HandleChase(filePath string, recordCreator *recordcreator.RecordCreator) error {

	fm := filemanager.FileManager{}

	file, err := fm.OpenFile(filePath, os.O_RDWR|os.O_CREATE, os.ModePerm)

	if err != nil {
		return err
	}

	defer file.Close()

	records := []*transactionstypes.ChaseTransaction{}
	chaseClient := csvprocessors.NewChaseClient()

	err = chaseClient.Unmarshal(file, &records)

	if err != nil {
		return err
	}

	normalizedRecords := chaseClient.ProcessCSV(records)
	for _, v := range normalizedRecords {
		err = recordCreator.CreateNewRecord(v)
		if err != nil {
			log.Println(err)
		}
	}

	return nil
}

func HandleCapitalOne(filePath string, recordCreator *recordcreator.RecordCreator) error {
	fm := filemanager.FileManager{}

	file, err := fm.OpenFile(filePath, os.O_RDWR|os.O_CREATE, os.ModePerm)

	if err != nil {
		return err
	}

	defer file.Close()

	records := []*transactionstypes.CapitalOneTransaction{}
	capitalOneClient := csvprocessors.NewCapitalOneClient()

	err = capitalOneClient.Unmarshal(file, &records)

	if err != nil {
		return err
	}

	normalizedRecords := capitalOneClient.ProcessCSV(records)
	for _, v := range normalizedRecords {
		err = recordCreator.CreateNewRecord(v)
		if err != nil {
			log.Println(err)
		}
	}

	return nil
}
