package main

import (
	"log"
	"net/http"
	"server/configs"
	"server/router"
)

func main() {
	mux := http.NewServeMux()
	db, err := configs.Connect()
	if err != nil {
		log.Fatal(err)
	}

	mux.Handle("/", router.Handlers(db))

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())
}
