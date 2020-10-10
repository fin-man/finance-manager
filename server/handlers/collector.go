package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fin-man/finance-manager/server/services"
)

type CollectorHandler struct {
	ProcessorService *services.ProcessorService
}

func NewCollectorHandler() *CollectorHandler {
	processorService := services.NewProcessorService()

	return &CollectorHandler{
		ProcessorService: processorService,
	}
}

func (c *CollectorHandler) RegisterNewCollector(w http.ResponseWriter, r *http.Request) {
	newCollector := make(map[string]string)

	err := json.NewDecoder(r.Body).Decode(&newCollector)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = c.ProcessorService.CreateNewCollector(newCollector["bank"], newCollector["route"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "registered new collector")
}

func (c *CollectorHandler) GetAllRegisteredCollectors(w http.ResponseWriter, r *http.Request) {

	services, err := c.CollectorService.GetAllCollectors()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "%v", services)

}

func (c *CollectorHandler) GetRegisteredCollector(w http.ResponseWriter, r *http.Request) {

	collector, ok := r.URL.Query()["collector"]

	if !ok || len(collector[0]) < 1 {
		http.Error(w, "Invalid URL param 'collector' is missing", http.StatusBadRequest)
		return
	}

	hostname, err := c.CollectorService.GetNewCollector(collector[0])

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%v", hostname)

}

func (c *CollectorHandler) Upload(w http.ResponseWriter, r *http.Request) {
	collector, ok := r.URL.Query()["collector"]

	if !ok || len(collector[0]) < 1 {
		http.Error(w, "Invalid URL param 'collector' is missing", http.StatusBadRequest)
		return
	}

	hostname, err := c.CollectorService.GetNewCollector(collector[0])

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	builtURL := fmt.Sprintf("%s/upload", hostname)

	http.Redirect(w, r, builtURL, http.StatusSeeOther)
}
