package common

import (
	"log"
	"net/http"
)

func HTTPError(w http.ResponseWriter, r *http.Request, error string, code int) {
	log.Printf("%s: %s", r.URL.Path, error)
	http.Error(w, error, code)
}
