package unit_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"contacts/handlers"
	"contacts/models"
)

func TestCreateContact(t *testing.T) {
	contact := models.Contact{
		Name:  "John Doe",
		Email: "johndoe@example.com",
		Phone: "555-1234",
	}
	body, _ := json.Marshal(contact)
	req, _ := http.NewRequest("POST", "/contacts", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.CreateContact)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	var c models.Contact
	json.Unmarshal(rr.Body.Bytes(), &c)
	if c.Name != "John Doe" {
		t.Errorf("handler returned unexpected body: got %v name want %v name", c.Name, "John Doe")
	}
}
