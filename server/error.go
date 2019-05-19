package server

import (
	"log"
	"net/http"
)

func ServerError(w http.ResponseWriter, r *http.Request, err error) {
	log.Fatal(err)
	http.Error(w, "", http.StatusInternalServerError)
}
