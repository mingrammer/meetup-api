package web

import (
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
)

func HealthCheck(w rest.ResponseWriter, r *rest.Request) {
	w.WriteHeader(http.StatusOK)
}
