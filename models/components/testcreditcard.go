// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/bolt-go/internal/utils"
	"time"
)

// TestCreditCardNetwork - The credit card's network.
type TestCreditCardNetwork string

const (
	TestCreditCardNetworkVisa         TestCreditCardNetwork = "visa"
	TestCreditCardNetworkMastercard   TestCreditCardNetwork = "mastercard"
	TestCreditCardNetworkAmex         TestCreditCardNetwork = "amex"
	TestCreditCardNetworkDiscover     TestCreditCardNetwork = "discover"
	TestCreditCardNetworkJcb          TestCreditCardNetwork = "jcb"
	TestCreditCardNetworkUnionpay     TestCreditCardNetwork = "unionpay"
	TestCreditCardNetworkAlliancedata TestCreditCardNetwork = "alliancedata"
	TestCreditCardNetworkCitiplcc     TestCreditCardNetwork = "citiplcc"
)

func (e TestCreditCardNetwork) ToPointer() *TestCreditCardNetwork {
	return &e
}

func (e *TestCreditCardNetwork) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "visa":
		fallthrough
	case "mastercard":
		fallthrough
	case "amex":
		fallthrough
	case "discover":
		fallthrough
	case "jcb":
		fallthrough
	case "unionpay":
		fallthrough
	case "alliancedata":
		fallthrough
	case "citiplcc":
		*e = TestCreditCardNetwork(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TestCreditCardNetwork: %v", v)
	}
}

type TestCreditCard struct {
	// The credit card's network.
	Network TestCreditCardNetwork `json:"network"`
	// The Bank Identification Number (BIN). This is typically the first 4 to 6 digits of the account number.
	Bin string `json:"bin"`
	// The account number's last four digits.
	Last4 string `json:"last4"`
	// The token's expiration date. Tokens used past their expiration will be rejected.
	Expiration time.Time `json:"expiration"`
	// The Bolt token associated with the credit card.
	Token string `json:"token"`
}

func (t TestCreditCard) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TestCreditCard) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TestCreditCard) GetNetwork() TestCreditCardNetwork {
	if o == nil {
		return TestCreditCardNetwork("")
	}
	return o.Network
}

func (o *TestCreditCard) GetBin() string {
	if o == nil {
		return ""
	}
	return o.Bin
}

func (o *TestCreditCard) GetLast4() string {
	if o == nil {
		return ""
	}
	return o.Last4
}

func (o *TestCreditCard) GetExpiration() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.Expiration
}

func (o *TestCreditCard) GetToken() string {
	if o == nil {
		return ""
	}
	return o.Token
}
