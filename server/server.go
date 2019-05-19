package server

import (
	"github.com/Anrop/Arma-Worlds-API/database"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

type Server struct {
	database *database.Database
	router   *mux.Router
}

func New(database *database.Database) (*Server, error) {
	router := mux.NewRouter()

	server := &Server{
		database: database,
		router:   router,
	}

	server.SetupRoutes()

	return server, nil
}

func (s *Server) Serve(port string) error {
	var handler http.Handler
	handler = handlers.CORS()(s.router)
	handler = handlers.CompressHandler(handler)

	return http.ListenAndServe(":"+port, handler)
}
