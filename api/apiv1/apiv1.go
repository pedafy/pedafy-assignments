package apiv1

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pedafy/pedafy-assignments/database"
	"github.com/pedafy/pedafy-assignments/database/dblayer"
)

// APIv1 represents the first version of the API
type APIv1 struct {
	dbHandler database.DatabaseHandler
}

// InitialisationDatabase will create a new database handler
// depending on the given version
func (a *APIv1) InitialisationDatabase(versionDB string) {
	a.dbHandler = dblayer.NewDatabaseManager(versionDB)
}

// connectDatabase will connect the API's database handler with the
// server thanks to the given credentials
func (a *APIv1) connectDatabase(user, pass, url, dbname string) error {
	return a.dbHandler.ConnectionDatabase(user, pass, url, dbname)
}

// RegisterAPIRoutes will register all the API's route to
// the given router
func (a *APIv1) RegisterAPIRoutes(r *mux.Router) {

	v1 := r.PathPrefix("/v1").Subrouter()

	v1.Methods(http.MethodGet).Path("/_ah/start").HandlerFunc(a.startupHandler)
	v1.Methods(http.MethodGet).Path("/").HandlerFunc(a.homeHandler)
}
