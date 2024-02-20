// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package boltgo

import (
	"bytes"
	"context"
	"fmt"
	"github.com/speakeasy-sdks/bolt-go/internal/hooks"
	"github.com/speakeasy-sdks/bolt-go/internal/utils"
	"github.com/speakeasy-sdks/bolt-go/models/components"
	"github.com/speakeasy-sdks/bolt-go/models/operations"
	"github.com/speakeasy-sdks/bolt-go/models/sdkerrors"
	"io"
	"net/http"
	"net/url"
)

// OAuth - Use this endpoint to retrieve an OAuth token. Use the token to allow your ecommerce server to make calls to the Account
// endpoint and create a one-click checkout experience for shoppers.
//
// https://help.bolt.com/products/accounts/direct-api/oauth-guide/
type OAuth struct {
	sdkConfiguration sdkConfiguration
}

func newOAuth(sdkConfig sdkConfiguration) *OAuth {
	return &OAuth{
		sdkConfiguration: sdkConfig,
	}
}

// GetToken - Get OAuth token
// Retrieve a new or refresh an existing OAuth token.
func (s *OAuth) GetToken(ctx context.Context, request components.GetAccessTokenRequest) (*operations.OauthGetTokenResponse, error) {
	hookCtx := hooks.HookContext{OperationID: "oauthGetToken"}

	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	opURL, err := url.JoinPath(baseURL, "/oauth/token")
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, false, false, "Request", "form", `request:"mediaType=application/x-www-form-urlencoded"`)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", opURL, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", s.sdkConfiguration.UserAgent)
	req.Header.Set("Content-Type", reqContentType)

	req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{hookCtx}, req)
	if err != nil {
		return nil, err
	}

	client := s.sdkConfiguration.SecurityClient

	httpRes, err := client.Do(req)
	if err != nil || httpRes == nil {
		if err != nil {
			err = fmt.Errorf("error sending request: %w", err)
		} else {
			err = fmt.Errorf("error sending request: no response")
		}

		_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{hookCtx}, nil, err)
		return nil, err
	} else if utils.MatchStatusCodes([]string{"4XX", "5XX"}, httpRes.StatusCode) {
		httpRes, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{hookCtx}, httpRes, nil)
		if err != nil {
			return nil, err
		}
	} else {
		httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{hookCtx}, httpRes)
		if err != nil {
			return nil, err
		}
	}
	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.OauthGetTokenResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out components.GetAccessTokenResponse
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.GetAccessTokenResponse = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out sdkerrors.Error
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}
			return nil, &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		return nil, sdkerrors.NewSDKError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	default:
	}

	return res, nil
}
