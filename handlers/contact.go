package handlers

import (
	"contacts/models"
	"net/http"
	"sync"
)

var (
	mu       sync.Mutex
	contacts = make(map[string]models.Contact)
)

func CreateContact(w http.ResponseWriter, r *http.Request) {
}

func GetContact(w http.ResponseWriter, r *http.Request) {
}

func DeleteContact(w http.ResponseWriter, r *http.Request) {
}

func UpdateContact(w http.ResponseWriter, r *http.Request) {
}
