package router

import (
	"contacts/handlers"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/contacts", handlers.CreateContact).Methods("POST")
	router.HandleFunc("/contacts/{id}", handlers.GetContact).Methods("GET")
	router.HandleFunc("/contacts/{id}", handlers.DeleteContact).Methods("DELETE")
	router.HandleFunc("/contacts/{id}", handlers.UpdateContact).Methods("PUT")

	return router
}
