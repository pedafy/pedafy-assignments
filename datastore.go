package pedafytig

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

type databaseInformation struct {
	ApiUsername  string `datastore:"API_USER_NAME"`
	ApiPass      string `datastore:"API_USER_PASS"`
	InstanceName string `datastore:"INSTANCE_NAME"`
}

// Create a new database information key
func createDatabaseInformation(ctx context.Context, username, pass, instance string) error {
	dbInfo := databaseInformation{username, pass, instance}

	key := datastore.NewIncompleteKey(ctx, "DATABASE_INFORMATION", nil)

	key, err := datastore.Put(ctx, key, &dbInfo)
	if err != nil {
		return err
	}
	return nil
}

// Find on database information
func findDatabaseInformation(ctx context.Context) (databaseInformation, error) {
	q := datastore.NewQuery("DATABASE_INFORMATION").Limit(1)
	iterator := q.Run(ctx)

	var entity databaseInformation
	_, err := iterator.Next(&entity)
	log.Println(entity)
	if err != nil {
		return databaseInformation{}, err
	}
	return entity, nil
}

// DEBUG: trying things here
func datastoreH(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	log.Println(appengine.ModuleName(ctx))

	err := createDatabaseInformation(ctx, os.Getenv("DBUSER"), os.Getenv("DBPASS"), os.Getenv("INSTANCE"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dbInfo2, err := findDatabaseInformation(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "New database information: %q", dbInfo2.ApiUsername)
}
