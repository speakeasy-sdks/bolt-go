// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

type PaymentMethodReferenceTag string

const (
	PaymentMethodReferenceTagID PaymentMethodReferenceTag = "id"
)

func (e PaymentMethodReferenceTag) ToPointer() *PaymentMethodReferenceTag {
	return &e
}

func (e *PaymentMethodReferenceTag) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "id":
		*e = PaymentMethodReferenceTag(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PaymentMethodReferenceTag: %v", v)
	}
}

type PaymentMethodReference struct {
	DotTag PaymentMethodReferenceTag `json:".tag"`
	// Payment ID of the saved Bolt Payment method.
	ID string `json:"id"`
}

func (o *PaymentMethodReference) GetDotTag() PaymentMethodReferenceTag {
	if o == nil {
		return PaymentMethodReferenceTag("")
	}
	return o.DotTag
}

func (o *PaymentMethodReference) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}
