// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/bolt-go/models/components"
	"net/http"
)

type OauthGetTokenResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Access token is successfully fetched
	GetAccessTokenResponse *components.GetAccessTokenResponse
}

func (o *OauthGetTokenResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *OauthGetTokenResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *OauthGetTokenResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *OauthGetTokenResponse) GetGetAccessTokenResponse() *components.GetAccessTokenResponse {
	if o == nil {
		return nil
	}
	return o.GetAccessTokenResponse
}