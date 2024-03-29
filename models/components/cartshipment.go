// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type CartShipment struct {
	Address *AddressReference `json:"address,omitempty"`
	// A monetary amount, i.e. a base unit amount and a supported currency.
	Cost *Amount `json:"cost,omitempty"`
	// The name of the carrier selected.
	Carrier *string `json:"carrier,omitempty"`
}

func (o *CartShipment) GetAddress() *AddressReference {
	if o == nil {
		return nil
	}
	return o.Address
}

func (o *CartShipment) GetAddressID() *Schemas {
	if v := o.GetAddress(); v != nil {
		return v.Schemas
	}
	return nil
}

func (o *CartShipment) GetAddressExplicit() *SchemasInput {
	if v := o.GetAddress(); v != nil {
		return v.SchemasInput
	}
	return nil
}

func (o *CartShipment) GetCost() *Amount {
	if o == nil {
		return nil
	}
	return o.Cost
}

func (o *CartShipment) GetCarrier() *string {
	if o == nil {
		return nil
	}
	return o.Carrier
}
