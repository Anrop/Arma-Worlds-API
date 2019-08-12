package server

import (
	"log"
	"net/http"
)

// Error logs error and returns 500 Internal Server Error
func (s *Server) error(w http.ResponseWriter, r *http.Request, err error) {
	log.Fatal(err)
	http.Error(w, "", http.StatusInternalServerError)
}
