package main

import (
	"log"
	"net/http"

	common pb "githhub.com/ninoaguilar/commons/api"
)

type handler struct {
	// gateway
}

func NewHandler() *handler {
	return &handler{}
}

func (h *handler) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/v1/customers/{customer_id}/orders", h.HandleCreateOrder)
}

func (h *handler) HandleCreateOrder(w http.ResponseWriter, r *http.Request) {
	log.Println("hello world")
}
