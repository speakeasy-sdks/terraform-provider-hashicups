// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// OrderItem - An OrderItem is an item/quantity in an order
type OrderItem struct {
	Coffee   *Coffee `json:"coffee,omitempty"`
	CoffeeID *int64  `json:"coffee_id,omitempty"`
	ID       *int64  `json:"id,omitempty"`
	OrderID  *int64  `json:"order_id,omitempty"`
	Quantity *int64  `json:"quantity,omitempty"`
}
