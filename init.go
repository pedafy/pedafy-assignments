package pedafytig

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"google.golang.org/appengine"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {

	http.HandleFunc("/", apiHomeH)
}

func initDB(user, password, dbname, connectionName string) {
	var err error

	if appengine.IsDevAppServer() {
		db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp([127.0.0.1]:3306)/%s", user, password, dbname))
	} else {
		db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@cloudsql(%s)/%s", user, password, connectionName, dbname))
	}

	if err != nil {
		log.Fatalf("Could not open db: %v", err)
	}
}

func apiHomeH(w http.ResponseWriter, r *http.Request) {

	// only '/' url
	if r.RequestURI != "/" {
		http.NotFound(w, r)
		return
	}

	// if no db, create one
	if db == nil {
		ctx := appengine.NewContext(r)
		dbInfo, _ := findDatabaseInformation(ctx)
		initDB(dbInfo.ApiUsername, dbInfo.ApiPass, "pedafy_assignments", dbInfo.InstanceName)
	}

	// set header for JSON
	w.Header().Set("Content-Type", "application/json;charset=utf8")

	// Test database stuff here
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
