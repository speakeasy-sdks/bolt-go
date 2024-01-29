// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// AddressErrorTag - The type of error returned
type AddressErrorTag string

const (
	AddressErrorTagInvalidRegion     AddressErrorTag = "invalid_region"
	AddressErrorTagInvalidPostalCode AddressErrorTag = "invalid_postal_code"
)

func (e AddressErrorTag) ToPointer() *AddressErrorTag {
	return &e
}

func (e *AddressErrorTag) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "invalid_region":
		fallthrough
	case "invalid_postal_code":
		*e = AddressErrorTag(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AddressErrorTag: %v", v)
	}
}

type AddressError struct {
	// The type of error returned
	DotTag AddressErrorTag `json:".tag"`
	// A human-readable error message, which might include information specific to
	// the request that was made.
	//
	Message string `json:"message"`
}

func (o *AddressError) GetDotTag() AddressErrorTag {
	if o == nil {
		return AddressErrorTag("")
	}
	return o.DotTag
}

func (o *AddressError) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}
