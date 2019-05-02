package pedafytig

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"google.golang.org/appengine"

	"github.com/joho/godotenv"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	// Load the environment
	godotenv.Load(".env")

	var (
		connectionName = os.Getenv("INSTANCE")
		user           = os.Getenv("DBUSER")
		password       = os.Getenv("DBPASS")
		dbname         = os.Getenv("DBNAME")
	)

	var err error

	if appengine.IsDevAppServer() {
		db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp([127.0.0.1]:3306)/%s", user, password, dbname))
	} else {
		db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@cloudsql(%s)/%s", user, password, connectionName, dbname))
	}

	if err != nil {
		log.Fatalf("Could not open db: %v", err)
	}

	http.HandleFunc("/", apiHomeH)
	http.HandleFunc("/datastore", datastoreH)
}

func apiHomeH(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=utf8")

	rows, err := db.Query("SELECT * FROM `test`")
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
