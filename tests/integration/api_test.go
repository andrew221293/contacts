package integration_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"contacts/models"
	"contacts/router"

	"github.com/gorilla/mux"
)

func setupRouter() *mux.Router {
	return router.NewRouter()
}

func TestAPI(t *testing.T) {
	r := setupRouter()

	contact := models.Contact{
		Name:  "Jane Doe",
		Email: "janedoe@example.com",
		Phone: "555-4321",
	}
	body, _ := json.Marshal(contact)
	req, _ := http.NewRequest("POST", "/contacts", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("POST /contacts returned wrong status code: got %v want %v", rr.Code, http.StatusCreated)
	}
}
