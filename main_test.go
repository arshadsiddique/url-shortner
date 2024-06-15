package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestShortenURL(t *testing.T) {
	urlShortener := NewURLShortener()
	r := mux.NewRouter()
	r.HandleFunc("/shorten", urlShortener.shortenURL).Methods("PUT")

	payload := map[string]string{"destination": "https://www.google.com"}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("PUT", "/shorten", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var response map[string]string
	json.NewDecoder(rr.Body).Decode(&response)
	assert.Contains(t, response, "shortened_url")
}

func TestRedirectToURL(t *testing.T) {
	urlShortener := NewURLShortener()
	r := mux.NewRouter()
	r.HandleFunc("/shorten", urlShortener.shortenURL).Methods("PUT")
	r.HandleFunc("/{shortcode}", urlShortener.redirectToURL).Methods("GET")

	// Shorten the URL
	payload := map[string]string{"destination": "https://www.google.com"}
	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("PUT", "/shorten", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	var response map[string]string
	json.NewDecoder(rr.Body).Decode(&response)
	shortenedURL := response["shortened_url"]

	// Extract shortcode from the shortened URL
	shortcode := shortenedURL[len("http://localhost:8080/"):]

	// Redirect to the original URL
	req, _ = http.NewRequest("GET", "/"+shortcode, nil)
	rr = httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusFound, rr.Code)
	assert.Equal(t, "https://www.google.com", rr.Header().Get("Location"))
}

func TestInvalidShortenURL(t *testing.T) {
	urlShortener := NewURLShortener()
	r := mux.NewRouter()
	r.HandleFunc("/shorten", urlShortener.shortenURL).Methods("PUT")

	payload := map[string]string{"destination": "invalid-url"}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("PUT", "/shorten", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}
