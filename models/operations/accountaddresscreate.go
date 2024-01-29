// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/bolt-go/models/components"
	"net/http"
)

type AccountAddressCreateRequest struct {
	// The publicly viewable identifier used to identify a merchant division.
	XPublishableKey string                         `header:"style=simple,explode=false,name=X-Publishable-Key"`
	AddressListing  components.AddressListingInput `request:"mediaType=application/json"`
}

func (o *AccountAddressCreateRequest) GetXPublishableKey() string {
	if o == nil {
		return ""
	}
	return o.XPublishableKey
}

func (o *AccountAddressCreateRequest) GetAddressListing() components.AddressListingInput {
	if o == nil {
		return components.AddressListingInput{}
	}
	return o.AddressListing
}

type AccountAddressCreateResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The address was successfully added
	AddressListing *components.AddressListing
}

func (o *AccountAddressCreateResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *AccountAddressCreateResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *AccountAddressCreateResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *AccountAddressCreateResponse) GetAddressListing() *components.AddressListing {
	if o == nil {
		return nil
	}
	return o.AddressListing
}
