package main

import (
	"encoding/json"
	"finance-manager/categories"
	"finance-manager/csvprocessors"
	"finance-manager/filemanager"
	"finance-manager/transactionstypes"
	"fmt"
	"log"
	"os"
)

func main() {

	// fm := filemanager.FileManager{}
	// file, err := fm.OpenFile("chase.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)

	// if err != nil {
	// 	panic(err)
	// }

	// records := []*transactionstypes.ChaseTransaction{}

	// chaseclient := csvprocessors.NewChaseClient()

	// err = chaseclient.Unmarshal(file, &records)

	// if err != nil {
	// 	panic(err)
	// }
	// chaseclient.PrintRecords(records, true)

	fm := filemanager.FileManager{}
	file, err := fm.OpenFile("capitalone.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)

	if err != nil {
		panic(err)
	}

	records := []*transactionstypes.CapitalOneTransaction{}

	capitalOneClient := csvprocessors.NewCapitalOneClient()

	err = capitalOneClient.Unmarshal(file, &records)

	if err != nil {
		panic(err)
	}

	datas := capitalOneClient.ProcessCSV(records)

	bytes, err := json.Marshal(&datas)

	if err != nil {
		log.Fatal(err)
	}

	fm.SaveFile("data.json", "", bytes)

	fmt.Println(categories.CapitalOneTransactionTypes)
}
