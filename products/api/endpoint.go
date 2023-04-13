package api

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type decreaseProductQuantityRequest struct {
	ProductName    string `json:"product_name"`
	Quantity int `json:"quantity"`
}

type decreaseProductQuantityResponse struct {
	Quantity int `json:"quantity,omitempty"`
	Price int `json:"price,omitempty"`
	Err   string `json:"err,omitempty"` // errors don't JSON-marshal, so we use a string
}

func makeDecreaseProductQuantityEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(decreaseProductQuantityRequest)
		quantity, price, err := svc.DecreaseProductQuantity(ctx, req.ProductName, req.Quantity)
		if err != nil {
			return decreaseProductQuantityResponse{-1, -1, err.Error()}, err
		}
		return decreaseProductQuantityResponse{quantity, price, ""}, err
	}
}
