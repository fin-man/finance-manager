package main

import (
	"github.com/fin-man/finance-manager/server/models"
	"github.com/fin-man/finance-manager/server/services"

	"github.com/fin-man/finance-manager/server/handlers"

	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/fin-man/finance-manager/server/routers"
)

func main() {

	router := routers.NewRouter()
	_, err := models.DBInit()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Initialized DB")
	categoriesHandler := handlers.NewCategoriesHandler()
	processorHandler := handlers.NewProcessorHandler()
	processorService := services.NewProcessorService()
	processorManager := services.NewCollectorManager(processorService)
	transactionPostgresService := services.NewTransactionPostgresService()
	transactionPostgresHandler := handlers.NewTransactionPostgresHandler(transactionPostgresService)

	go processorManager.RunCollectorHealthChecks()

	// router.Router.HandleFunc("/transactions", transactionHandler.GetAllTransactions).Methods("GET")
	// router.Router.HandleFunc("/transactions/range", transactionHandler.GetTransactionsInDateRange).Methods("GET")
	router.Router.HandleFunc("/transactions", transactionPostgresHandler.CreateTransaction).Methods("POST")
	// router.Router.HandleFunc("/transactions/graph", transactionHandler.GetAllTransactionsGraph).Methods("GET")
	// router.Router.HandleFunc("/transactions/search", transactionHandler.SearchTransactions).Methods("GET")
	router.Router.HandleFunc("/processors", processorHandler.GetAllRegisteredProcessors).Methods("GET")
	router.Router.HandleFunc("/processor", processorHandler.RegisterNewProcessor).Methods("POST")
	router.Router.HandleFunc("/categories", categoriesHandler.GetAllCategories).Methods("GET")

	csvHandler := handlers.NewCSVHandler()
	router.Router.HandleFunc("/csv/fields", csvHandler.GetAllFields)
	port := os.Getenv("SERVER_PORT")

	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on port : %s\n", port)
	//     log.Fatal(http.ListenAndServe(":3000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))

	http.ListenAndServe(fmt.Sprintf(":%s", port), router.Router)
}
