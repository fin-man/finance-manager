package handlers

import (
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

	log.Printf("redirecting %s upload requet to %s\n", proxyTarget[0], proxyBank[0])
	http.Redirect(w, r, proxyTarget[0], http.StatusTemporaryRedirect)
}
