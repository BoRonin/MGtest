package main

import (
	"log"
	"mgtest/internal/app"
)

func main() {
	app := app.New()
	app.Log.Debug("starting server...")
	defer func() {
		app.S.PS.Close()
	}()
	if err := app.Server.ListenAndServe(); err != nil {
		log.Fatal("server couldn't run")
	}
	app.Log.Debug("serving...")
}
