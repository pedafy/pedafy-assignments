package apiv1

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pedafy/pedafy-assignments/src/database"
	"github.com/pedafy/pedafy-assignments/src/database/dblayer"
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

// RegisterAPIRoutes will register all the API's route to
// the given router
func (a *APIv1) RegisterAPIRoutes(r *mux.Router) {

	r.Methods(http.MethodGet).Path("/_ah/start").HandlerFunc(a.startupHandler)
	r.Methods(http.MethodGet).Path("/").HandlerFunc(a.homeHandler)

	r.Use(a.setJSONResponse)

	// Status
	r.Methods(http.MethodGet).Path("/status").HandlerFunc(a.statusGetAllHandler)
	r.Methods(http.MethodGet).Path("/status/{id:[0-9]+}").HandlerFunc(a.statusGetByIDHandler)
	r.Methods(http.MethodGet).Path("/status/{name}").HandlerFunc(a.statusGetByNameHandler)
}

// connectDatabase will connect the API's database handler with the
// server thanks to the given credentials
func (a *APIv1) connectDatabase(user, pass, url, dbname string) error {
	return a.dbHandler.ConnectionDatabase(user, pass, url, dbname)
}

func (a *APIv1) setJSONResponse(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json;charset=utf8")
		next.ServeHTTP(w, r)
	})
}
