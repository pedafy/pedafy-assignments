package apiv1

import "database/sql"

// APIv1 represents the first version of the API
type APIv1 struct {
	db *sql.DB
}

func (a APIv1) Test() {

}
