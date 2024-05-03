package main

import (
	"contacts/router"
	"log"
	"net/http"
)

func main() {
	r := router.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", r))
}
