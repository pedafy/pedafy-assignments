package dblayer

import (
	"github.com/pedafy/pedafy-assignments/database"
	"github.com/pedafy/pedafy-assignments/database/dbv1"
	"github.com/pedafy/pedafy-assignments/version"
)

// NewDatabaseManager returns the database version corresponding to the
// given version
func NewDatabaseManager(dbVersion string) database.DatabaseHandler {
	switch dbVersion {
	case version.Version1:
		return &dbv1.DBV1{}
	}
	return nil
}
