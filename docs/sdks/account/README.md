# Account
(*Account*)

## Overview

Account endpoints allow you to view and manage shoppers' accounts. For example,
you can add or remove addresses and payment information.


### Available Operations

* [GetDetails](#getdetails) - Retrieve account details
* [DeleteAddress](#deleteaddress) - Delete an existing address
* [Detect](#detect) - Determine the existence of a Bolt account
* [DeletePaymentMethod](#deletepaymentmethod) - Delete an existing payment method

## GetDetails

Retrieve a shopper's account details, such as addresses and payment information

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


    var xPublishableKey string = "<value>"

    ctx := context.Background()
    res, err := s.Account.GetDetails(ctx, xPublishableKey)
    if err != nil {
        log.Fatal(err)
    }
    if res.Account != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `xPublishableKey`                                                      | *string*                                                               | :heavy_check_mark:                                                     | The publicly viewable identifier used to identify a merchant division. |


### Response

**[*operations.AccountGetResponse](../../models/operations/accountgetresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 4XX                | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## DeleteAddress

Delete an existing address. Deleting an address does not invalidate transactions or
shipments that are associated with it.


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


    var id string = "D4g3h5tBuVYK9"

    var xPublishableKey string = "<value>"

    ctx := context.Background()
    res, err := s.Account.DeleteAddress(ctx, id, xPublishableKey)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            | Example                                                                |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |                                                                        |
| `id`                                                                   | *string*                                                               | :heavy_check_mark:                                                     | The ID of the address to delete                                        | D4g3h5tBuVYK9                                                          |
| `xPublishableKey`                                                      | *string*                                                               | :heavy_check_mark:                                                     | The publicly viewable identifier used to identify a merchant division. |                                                                        |


### Response

**[*operations.AccountAddressDeleteResponse](../../models/operations/accountaddressdeleteresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 4XX                | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Detect

Determine whether or not an identifier is associated with an existing Bolt account.

### Example Usage

```go
package main

import(
	boltgo "github.com/speakeasy-sdks/bolt-go"
	"github.com/speakeasy-sdks/bolt-go/models/components"
	"context"
	"log"
)

func main() {
    s := boltgo.New()


    identifier := components.Identifier{
        IdentifierType: components.IdentifierTypeEmail,
        IdentifierValue: "alice@example.com",
    }

    var xPublishableKey string = "<value>"

    ctx := context.Background()
    res, err := s.Account.Detect(ctx, identifier, xPublishableKey)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                             | Type                                                                                                  | Required                                                                                              | Description                                                                                           |
| ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                 | [context.Context](https://pkg.go.dev/context#Context)                                                 | :heavy_check_mark:                                                                                    | The context to use for the request.                                                                   |
| `identifier`                                                                                          | [components.Identifier](../../models/components/identifier.md)                                        | :heavy_check_mark:                                                                                    | A type and value combination that defines the identifier used to detect<br/>the existence of an account.<br/> |
| `xPublishableKey`                                                                                     | *string*                                                                                              | :heavy_check_mark:                                                                                    | The publicly viewable identifier used to identify a merchant division.                                |


### Response

**[*operations.AccountExistsResponse](../../models/operations/accountexistsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 4XX                | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## DeletePaymentMethod

Delete an existing payment method. Deleting a payment method does not invalidate transactions or
orders that are associated with it.


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


    var id string = "D4g3h5tBuVYK9"

    var xPublishableKey string = "<value>"

    ctx := context.Background()
    res, err := s.Account.DeletePaymentMethod(ctx, id, xPublishableKey)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            | Example                                                                |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |                                                                        |
| `id`                                                                   | *string*                                                               | :heavy_check_mark:                                                     | The ID of the payment method to delete                                 | D4g3h5tBuVYK9                                                          |
| `xPublishableKey`                                                      | *string*                                                               | :heavy_check_mark:                                                     | The publicly viewable identifier used to identify a merchant division. |                                                                        |


### Response

**[*operations.AccountPaymentMethodDeleteResponse](../../models/operations/accountpaymentmethoddeleteresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 4XX                | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
