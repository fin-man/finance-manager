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
	router.Router.HandleFunc("/transactions/graph", transactionHandler.GetAllTransactionsGraph).Methods("GET")

	categoriesHandler := handlers.NewCategoriesHandler()
	router.Router.HandleFunc("/categories", categoriesHandler.GetAllCategories).Methods("GET")
	port := os.Getenv("SERVER_PORT")

	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on port : %s\n", port)
	//     log.Fatal(http.ListenAndServe(":3000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))

	http.ListenAndServe(fmt.Sprintf(":%s", port), router.Router)
}
