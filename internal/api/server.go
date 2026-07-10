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
	s.Handle("GET /job/id")
	s.Handle("GET /job/status/")
	s.Handle("POST /job")

	return &s
}
