package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"
	"sync"

	"github.com/gorilla/mux"
	"github.com/rs/xid"
)

type URLShortener struct {
	mu      sync.RWMutex
	storage map[string]string
}

func NewURLShortener() *URLShortener {
	return &URLShortener{
		storage: make(map[string]string),
	}
}

func (s *URLShortener) shortenURL(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Destination string `json:"destination"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Printf("Error decoding JSON: %v\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if _, err := url.ParseRequestURI(request.Destination); err != nil {
		log.Printf("Invalid URL provided: %s\n", request.Destination)
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	shortcode := xid.New().String()
	s.mu.Lock()
	s.storage[shortcode] = request.Destination
	s.mu.Unlock()

	response := map[string]string{"shortened_url": "http://localhost:8080/" + shortcode}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	log.Printf("Shortened URL: %s -> %s\n", shortcode, request.Destination)
}

func (s *URLShortener) redirectToURL(w http.ResponseWriter, r *http.Request) {
	shortcode := mux.Vars(r)["shortcode"]

	s.mu.RLock()
	destination, ok := s.storage[shortcode]
	s.mu.RUnlock()

	if !ok {
		log.Printf("Shortcode not found: %s\n", shortcode)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "URL not found"})
		return
	}

	log.Printf("Redirecting shortcode %s to %s\n", shortcode, destination)
	http.Redirect(w, r, destination, http.StatusFound)
}

func main() {
	// Set logging to stdout and stderr
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	urlShortener := NewURLShortener()
	r := mux.NewRouter()
	r.HandleFunc("/shorten", urlShortener.shortenURL).Methods("PUT")
	r.HandleFunc("/{shortcode}", urlShortener.redirectToURL).Methods("GET")

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Server failed to start: %v\n", err)
	}
}
