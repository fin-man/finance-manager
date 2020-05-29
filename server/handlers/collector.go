package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fin-man/finance-manager/server/services"
)

type CollectorHandler struct {
	CollectorService *services.CollectorService
}

func NewCollectorHandler() *CollectorHandler {
	collectorService := services.NewCollectorService()

	return &CollectorHandler{
		CollectorService: collectorService,
	}
}

func (c *CollectorHandler) RegisterNewCollector(w http.ResponseWriter, r *http.Request) {
	newCollector := make(map[string]string)

	err := json.NewDecoder(r.Body).Decode(&newCollector)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = c.CollectorService.CreateNewCollector(newCollector["bank"], newCollector["route"])
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
	newCollector := make(map[string]string)

	err := json.NewDecoder(r.Body).Decode(&newCollector)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	collector, err := c.CollectorService.GetNewCollector(newCollector["bank"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%v", collector)

}
