// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package boltgo

type Payments struct {
	Guest    *Guest
	LoggedIn *LoggedIn

	sdkConfiguration sdkConfiguration
}

func newPayments(sdkConfig sdkConfiguration) *Payments {
	return &Payments{
		sdkConfiguration: sdkConfig,
		Guest:            newGuest(sdkConfig),
		LoggedIn:         newLoggedIn(sdkConfig),
	}
}