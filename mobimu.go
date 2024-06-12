package main

import (
	"net/http"

	"github.com/patapancakes/mobimu/mobage"
)

func main() {
	http.HandleFunc("/{cluster}/api/{action}", mobage.HandleAPI)

	http.ListenAndServe(":80", nil)
}
