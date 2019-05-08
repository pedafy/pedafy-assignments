package main

import (

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/pedafy/pedafy-assignments/version"
	"google.golang.org/appengine"
)

const (
	// TODO: remove this and use version.Default() instead
	defaultVersion = version.Version1
)

func main() {
	var srv server

	// TODO: Only in development environments
	godotenv.Load("../.env")

	srv.initAPIHandler()
	srv.registerHandlers()

	appengine.Main()
}
