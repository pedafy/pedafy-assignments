package main

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/pedafy/pedafy-assignments/src/api"
	"github.com/pedafy/pedafy-assignments/src/api/layer"
)

type server struct {
	APIHandler api.APIHandler
}

func (s *server) initAPIHandler() {
	version := os.Getenv("VERSION_API")
	if version == "" {
		version = defaultVersion
	}
	s.APIHandler = layer.NewApiManager(version)
	s.APIHandler.InitialisationDatabase(version)
}

// registerHandlers will register all the API's routes
func (s *server) registerHandlers() {
	r := mux.NewRouter()

	s.APIHandler.RegisterAPIRoutes(r)

	http.Handle("/", handlers.CombinedLoggingHandler(os.Stderr, r))
}
