package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/fin-man/finance-manager/server/services"
)

type CategoriesHandler struct {
	CategoriesService *services.CategoriesService
}

func NewCategoriesHandler() *CategoriesHandler {
	return &CategoriesHandler{}
}
func (c *CategoriesHandler) GetAllCategories(w http.ResponseWriter, r *http.Request) {
	categories := c.CategoriesService.GetCategories()

	enableCors(&w)

	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(categories); err != nil {
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

	}
}
