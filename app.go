package main

import (
	"finance-manager/csvprocessors"
	"finance-manager/filemanager"
	"os"
)

// type Transaction struct {
// 	Transaction string `csv:"Transaction Date"`
// 	PostDate    string `csv:"Post Date"`
// 	Description string `csv:"Description"`
// 	Category    string `csv:"Category"`
// 	Type        string `csv:"Type"`
// 	Amount      string `csv:"Amount"`
// }

func main() {

	fm := filemanager.FileManager{}
	file, err := fm.OpenFile("financials.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)

	if err != nil {
		panic(err)
	}

	records := []*csvprocessors.ChaseTransaction{}

	chaseclient := csvprocessors.NewChaseClient()

	err = chaseclient.Unmarshal(file, &records)

	if err != nil {
		panic(err)
	}
	chaseclient.PrintRecords(records)

}
