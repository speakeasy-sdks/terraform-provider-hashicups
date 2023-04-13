// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"hashicups/internal/sdk/pkg/models/shared"
	"net/http"
)

type UpdateOrderRequest struct {
	// A JSON representation of the order to upsert
	Order shared.Order `request:"mediaType=application/json"`
	// The ID of the order
	OrderID int64 `pathParam:"style=simple,explode=false,name=orderID"`
}

type UpdateOrderResponse struct {
	ContentType string
	// OK
	Order       *shared.Order
	StatusCode  int
	RawResponse *http.Response
}
