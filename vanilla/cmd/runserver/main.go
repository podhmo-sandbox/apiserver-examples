package main

import (
	"log"

	"net/http"

	"github.com/podhmo/apiserver-examples/vanilla/app"
	"github.com/podhmo/apiserver-examples/vanilla/middleware"
	"github.com/podhmo/apiserver-examples/vanilla/store"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("%+v", err)
	}
}

func run() error {
	var (
		db     = store.NewDB()
		server http.Handler
	)
	server = app.NewServer(db)
	server = middleware.Recover(server)
	return http.ListenAndServe(":8081", server)
}
