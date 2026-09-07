package main

import (
	"AuthInGo/app"
	config "AuthInGo/config/env"
	dbConfig "AuthInGo/config/db"
)

func main() {

	config.Load()

	cfg:=app.NewConfig()
	app:=app.NewApplication(cfg)
	dbConfig.SetupDB()

	app.Run()
}