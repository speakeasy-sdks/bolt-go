// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type PaymentUpdateRequest struct {
	Cart *Cart `json:"cart,omitempty"`
}

func (o *PaymentUpdateRequest) GetCart() *Cart {
	if o == nil {
		return nil
	}
	return o.Cart
}
