package api

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type createOrderRequest struct {
	ProductName    string `json:"product_name"`
	Quantity int `json:"quantity"`
}

type createOrderResponse struct {
	OrderId int `json:"order_id,omitempty"`
	Email string `json:"email,omitempty"`
	TotalPrice int `json:"total_price,omitempty"` // quantity * price
	Err   string `json:"err,omitempty"` // errors don't JSON-marshal, so we use a string
}

func makeCreateOrderEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(createOrderRequest)
		order_id, email, total_price, err := svc.CreateOrder(ctx, req.ProductName, req.Quantity)
		if err != nil {
			return createOrderResponse{-1, "", -1, err.Error()}, err
		}
		return createOrderResponse{order_id, email, total_price, ""}, err
	}
}
