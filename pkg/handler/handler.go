package handler

import (
	"MEDODS/pkg/service"
	"log"
	"net/http"
)

type Handler struct {
	services *service.Service
	logger   *log.Logger
}

func NewHandler(service *service.Service, logger *log.Logger) *Handler {
	return &Handler{services: service, logger: logger}
}

func (h *Handler) InitRoutes() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("GET /refresh", h.refresh)
	router.HandleFunc("GET /login", h.login)
	return router
}
