package main

import (
	"finance-manager/categories"
	"fmt"
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

	// fm := filemanager.FileManager{}
	// file, err := fm.OpenFile("capitalone.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)

	// if err != nil {
	// 	panic(err)
	// }

	// records := []*transactionstypes.CapitalOneTransaction{}

	// capitalOneClient := csvprocessors.NewCapitalOneClient()

	// err = capitalOneClient.Unmarshal(file, &records)

	// if err != nil {
	// 	panic(err)
	// }
	// capitalOneClient.PrintRecords(records, true)

	fmt.Println(categories.CapitalOneTransactionTypes)
}
