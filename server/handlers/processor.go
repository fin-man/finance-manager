package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fin-man/finance-manager/server/models"

	"github.com/fin-man/finance-manager/server/services"
)

type ProcessorHandler struct {
	ProcessorService *services.ProcessorService
}

func NewProcessorHandler() *ProcessorHandler {
	processorService := services.NewProcessorService()

	return &ProcessorHandler{
		ProcessorService: processorService,
	}
}

func (c *ProcessorHandler) RegisterNewProcessor(w http.ResponseWriter, r *http.Request) {
	var processor models.ProcessorModel

	err := json.NewDecoder(r.Body).Decode(&processor)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = c.ProcessorService.CreateProcessor(&processor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "registered new collector")
}

func (c *ProcessorHandler) GetAllRegisteredProcessors(w http.ResponseWriter, r *http.Request) {

	processors, err := c.ProcessorService.GetAllProcessors()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%v", processors)

}
