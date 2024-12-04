package main

import (
	"fmt"
	"net/http"

	"github.com/ninoaguilar/burgertown/common"
	pb "github.com/ninoaguilar/burgertown/common/api"
)

type handler struct {
	client pb.OrderServiceClient
}

func NewHandler(client pb.OrderServiceClient) *handler {
	return &handler{}
}

func (h *handler) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/v1/customers/{customerId}/orders", h.HandleCreateOrder)
}

func (h *handler) HandleCreateOrder(w http.ResponseWriter, r *http.Request) {
	customerId := r.PathValue("customerId")

	var items []*pb.ItemWithQuantity
	if err := common.ReadJSON(r, &items); err != nil {
		common.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	h.client.CreateOrder(r.Context(), &pb.CreateOrderRequest{
		CustomerId: customerId,
		Items:      items,
	})

	fmt.Printf("here %s --", items)
}
