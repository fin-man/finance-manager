package main

import (
	"finance-manager/elasticsearch"
	"fmt"
)

func main() {

	esClient := elasticsearch.NewElasticSearchClient()

	fmt.Println(esClient)

	// watcher := filewatcher.NewFileWatcher()

	// pwd, err := os.Getwd()

	// fmt.Println(pwd)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// watcher.Watch(pwd + "/data/transactions")

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

	// datas := chaseclient.ProcessCSV(records)

	// bytes, err := json.Marshal(&datas)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fm.SaveFile("dataChase.json", "", bytes)

	////---------

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

	// datas := capitalOneClient.ProcessCSV(records)

	// bytes, err := json.Marshal(&datas)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fm.SaveFile("data.json", "", bytes)

	// fmt.Println(categories.CapitalOneTransactionTypes)
}
