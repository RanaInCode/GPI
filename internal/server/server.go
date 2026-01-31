package server

import (
	"net/http"
)

type Server struct {
	router http.Handler
}

func New () *Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK Anuj"))
	})
	return &Server{router: mux}
}

func (s *Server) Start(addr string ) error {
	return http.ListenAndServe(addr, s.router)
}