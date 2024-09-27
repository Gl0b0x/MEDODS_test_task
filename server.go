package server

import (
	"net/http"
	"time"
)

// Server представляет сервер HTTP.
type Server struct {
	httpServer *http.Server
}

// Run запускает сервер на указанном порту.
func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:         ":" + port,
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	return s.httpServer.ListenAndServe()
}
