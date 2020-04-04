package main

import (
	"finance-manager/server/handlers"
	"finance-manager/server/routers"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	router := routers.NewRouter()

	transactionHandler := handlers.NewTransactionHandler()
	router.Router.HandleFunc("/transactions", transactionHandler.GetAllTransactions).Methods("GET")
	router.Router.HandleFunc("/transactions/range", transactionHandler.GetTransactionsInDateRange).Methods("GET")
	router.Router.HandleFunc("/transactions", transactionHandler.CreateTransaction).Methods("POST")

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on port : %s\n", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), router.Router)
}
