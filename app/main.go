package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pedafy/pedafy-assignments/api/layer"

	"github.com/pedafy/pedafy-assignments/datastore"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"google.golang.org/appengine"
)

const (
	defaultVersion = layer.Version1
)

func main() {
	var srv server

	godotenv.Load("../.env")

	srv.initAPIHandler()
	srv.registerHandlers()

	appengine.Main()
}

// TODO: delete following code:

// startupH is the startup handler, Google App Engine requests
// this URL ('/_ah/start) when starting the service. It allows us to do all the
// startup, e.g: initialisation of the database.
func startupH(w http.ResponseWriter, r *http.Request) {
	if db == nil {
		ctx := appengine.NewContext(r)
		dbInfo, err := datastore.FindDatabaseInformation(ctx)
		if err != nil {
			log.Fatal(err.Error())
		}
		initDB(dbInfo.APIUsername, dbInfo.APIPass, "pedafy_assignments", dbInfo.InstanceName)
	}

	w.Write([]byte("ready"))
}

func apiHomeH(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=utf8")

	// DEBUG: Testing database connection here - Dirty code
	rows, err := db.Query("SELECT * FROM `status`")
	if err != nil {
		log.Fatal(err)
	}
	var id int
	var value string
	if rows.Next() {
		err = rows.Scan(&id, &value)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Fprintf(w, `{"id":"%d", "value":"%s"}`, id, value)
}
