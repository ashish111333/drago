package api

import (
	"fmt"
	"net/http"
)

func JobApiServer(port int) *http.ServeMux {
	s := http.ServeMux{}
	s.Handle("GET /", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("drago server running on port : %d", port)))

	}))
	// get job by ID
	s.Handle("GET /job/id", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	}))
	// get jobs by status
	s.Handle("GET /job/status/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	}))
	s.Handle("POST /job", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	}))
	s.Handle("DELETE /job/status", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	s.Handle("DELETE /job/id", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	return &s
}
