package server

import (
	"github.com/Anrop/Arma-Worlds-API/config"
	"github.com/Anrop/Arma-Worlds-API/database"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

// Server object wrapping required components
type Server struct {
	config   *config.Config
	database *database.Database
	router   *mux.Router
}

// New creates a new server instance
func New(config *config.Config, database *database.Database) (*Server, error) {
	router := mux.NewRouter()

	server := &Server{
		config:   config,
		database: database,
		router:   router,
	}

	server.setupRoutes()

	return server, nil
}

// Serve binds the server to port
func (s *Server) Serve(port string) error {
	var handler http.Handler
	handler = handlers.CORS()(s.router)
	handler = handlers.CompressHandler(handler)

	return http.ListenAndServe(":"+port, handler)
}
