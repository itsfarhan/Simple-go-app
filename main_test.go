// main_test.go
package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// TestHomeHandler tests the homeHandler function
func TestHomeHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(homeHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("homeHandler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "<title>Home</title>"
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("homeHandler returned unexpected body: got %v want to contain %v", rr.Body.String(), expected)
	}
}

// TestAboutHandler tests the aboutHandler function
func TestAboutHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/about", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(aboutHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("aboutHandler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "<title>About</title>"
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("aboutHandler returned unexpected body: got %v want to contain %v", rr.Body.String(), expected)
	}
}
