package main

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/pedafy/pedafy-assignments/api"
	"github.com/pedafy/pedafy-assignments/api/layer"
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
}

func (s *server) registerHandlers() {
	r := mux.NewRouter()
	r.HandleFunc("/", apiHomeH)
	r.HandleFunc("/_ah/start", startupH)

	http.Handle("/", handlers.CombinedLoggingHandler(os.Stderr, r))
}
