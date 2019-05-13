package apiv1

import (
	"log"
	"net/http"
	"time"

	"github.com/pedafy/pedafy-assignments/src/datastore"
	"google.golang.org/appengine"
)

// startupHandler is the startup handler, Google App Engine requests
// this URL ('/_ah/start) when starting the service. It allows us to do all the
// startup, e.g: initialisation of the database.
func (a *APIv1) startupHandler(w http.ResponseWriter, r *http.Request) {
	if a.dbHandler.IsNewDB() {
		ctx := appengine.NewContext(r)
		dbInfo, err := datastore.FindDatabaseInformation(ctx)
		if err != nil {
			log.Fatal(err.Error())
		}
		err = a.connectDatabase(dbInfo.APIUsername, dbInfo.APIPass, "pedafy_assignments", dbInfo.InstanceName)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
	if a.authToken == "" {
		ctx := appengine.NewContext(r)
		var err error
		a.authToken, err = datastore.FindAPITokenInformation(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}
	w.Write([]byte("ready"))
}

func (a *APIv1) homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(time.Now().Format(time.RFC850)))
}
