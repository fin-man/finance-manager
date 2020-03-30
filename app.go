package main

import (
	"finance-manager/csvprocessors"
	"finance-manager/filemanager"
	"os"
)

func main() {

	fm := filemanager.FileManager{}
	file, err := fm.OpenFile("chase.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)

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
