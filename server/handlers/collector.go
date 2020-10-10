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

// func (c *ProcessorHandler) GetRegisteredCollector(w http.ResponseWriter, r *http.Request) {

// 	collector, ok := r.URL.Query()["collector"]

// 	if !ok || len(collector[0]) < 1 {
// 		http.Error(w, "Invalid URL param 'collector' is missing", http.StatusBadRequest)
// 		return
// 	}

// 	hostname, err := c.CollectorService.GetNewCollector(collector[0])

// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	fmt.Fprintf(w, "%v", hostname)

// }

// func (c *CollectorHandler) Upload(w http.ResponseWriter, r *http.Request) {
// 	collector, ok := r.URL.Query()["collector"]

// 	if !ok || len(collector[0]) < 1 {
// 		http.Error(w, "Invalid URL param 'collector' is missing", http.StatusBadRequest)
// 		return
// 	}

// 	hostname, err := c.CollectorService.GetNewCollector(collector[0])

// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	builtURL := fmt.Sprintf("%s/upload", hostname)

// 	http.Redirect(w, r, builtURL, http.StatusSeeOther)
// }
