package database

// DatabaseHandler interfaces all the version of the database code
type DatabaseHandler interface {
	ConnectionDatabase(user, pass, dbname, url string) error
	IsNewDB() bool
	IsOldDB() bool
}