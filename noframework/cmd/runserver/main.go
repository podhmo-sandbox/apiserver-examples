package main

import (
	"log"

	"net/http"

	"github.com/podhmo/apiserver-examples/noframework/app"
	"github.com/podhmo/apiserver-examples/noframework/store"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("%+v", err)
	}
}

func run() error {
	db := store.NewDB()
	server := app.NewServer(db)
	return http.ListenAndServe(":8081", server)
}
