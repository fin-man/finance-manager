package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/fin-man/finance-manager/server/services"
)

type CSVHandler struct {
	CSVService *services.CSVService
}

func NewCSVHandler() *CSVHandler {

	csvService := services.NewCSVService()

	return &CSVHandler{
		CSVService: csvService,
	}
}

func (c *CSVHandler) GetAllFields(w http.ResponseWriter, r *http.Request) {
	fields := c.CSVService.GetFields()
	enableCors(&w)
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(fields); err != nil {
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

	}
}
