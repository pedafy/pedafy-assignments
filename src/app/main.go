package main

import (

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"google.golang.org/appengine"
)

func main() {
	var srv server

	if appengine.IsDevAppServer() {
		godotenv.Load("../../.env")
	}

	srv.initAPIHandler()
	srv.registerHandlers()

	appengine.Main()
}
