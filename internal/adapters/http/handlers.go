package http

import (
	"net/http"
)

func addHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/-/ready", healthCheckHandler)
}
