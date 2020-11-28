package handlers

import (
	"fmt"
	"log"
	"net/http"
)

type ProxyHandler struct {
}

func NewProxyHandler() *ProxyHandler {
	return &ProxyHandler{}

}

func (p *ProxyHandler) UploadProxyHandler(w http.ResponseWriter, r *http.Request) {

	proxyTarget, okTarget := r.URL.Query()["target"]
	proxyBank, okBank := r.URL.Query()["bank"]

	if !okTarget || !okBank || len(proxyTarget) == 0 || len(proxyBank) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	buildTarget := fmt.Sprintf("%s/upload", proxyTarget[0])
	log.Printf("redirecting %s upload request to %s\n", buildTarget, proxyBank[0])
	http.Redirect(w, r, buildTarget, http.StatusTemporaryRedirect)
}
