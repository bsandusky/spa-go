package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func checkError(err error, t *testing.T) {
	if err != nil {
		t.Errorf("An error occurred: %v", err)
	}
}

func TestServeTemplateHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	checkError(err, t)

	rr := httptest.NewRecorder()

	http.HandlerFunc(serveTemplateHandler).ServeHTTP(rr, req)

	// check status code for 200
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code not as expected. \n Expected: %d, Actual: %d", http.StatusOK, status)
	}

	// check body length
	if lenBody := len(rr.Body.String()); lenBody == 0 {
		t.Errorf("No body found. Len(rr.Body.String()) == %v", lenBody)
	}

}
