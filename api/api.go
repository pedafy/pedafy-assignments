package api

import (
	"github.com/gorilla/mux"
)

// APIHandler interface every version of the API
type APIHandler interface {
	RegisterAPIRoutes(r *mux.Router)
	InitialisationDatabase(versionDB string)
}
