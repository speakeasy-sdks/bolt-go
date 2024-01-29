# OAuth
(*OAuth*)

## Overview

Use this endpoint to retrieve an OAuth token. Use the token to allow your ecommerce server to make calls to the Account
endpoint and create a one-click checkout experience for shoppers.


<https://help.bolt.com/products/accounts/direct-api/oauth-guide/>
### Available Operations

* [GetToken](#gettoken) - Get OAuth token

## GetToken

Retrieve a new or refresh an existing OAuth token.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/bolt-go/models/components"
	boltgo "github.com/speakeasy-sdks/bolt-go"
	"context"
	"log"
)

func main() {
    s := boltgo.New(
        boltgo.WithSecurity(components.Security{
            Oauth: boltgo.String("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.OAuth.GetToken(ctx, components.GetAccessTokenRequest{
        Code: "string",
        ClientID: "string",
        ClientSecret: "string",
        GrantType: components.GrantTypeAuthorizationCode,
        Scope: []components.Scope{
            components.ScopeBoltAccountManage,
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetAccessTokenResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [components.GetAccessTokenRequest](../../models/components/getaccesstokenrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.OauthGetTokenResponse](../../models/operations/oauthgettokenresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 4XX                | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
