// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

type PaymentActionRequestTag string

const (
	PaymentActionRequestTagFinalize PaymentActionRequestTag = "finalize"
)

func (e PaymentActionRequestTag) ToPointer() *PaymentActionRequestTag {
	return &e
}

func (e *PaymentActionRequestTag) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "finalize":
		*e = PaymentActionRequestTag(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PaymentActionRequestTag: %v", v)
	}
}

type PaymentActionRequest struct {
	DotTag PaymentActionRequestTag `json:".tag"`
}

func (o *PaymentActionRequest) GetDotTag() PaymentActionRequestTag {
	if o == nil {
		return PaymentActionRequestTag("")
	}
	return o.DotTag
}
