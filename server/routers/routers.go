package routers

import (
	"github.com/gorilla/mux"
)

type Routers struct {
	Router *mux.Router
}

func NewRouter() *Routers {
	return &Routers{
		Router: mux.NewRouter(),
	}
}
