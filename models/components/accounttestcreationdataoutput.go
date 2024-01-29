// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

type EmailState string

const (
	EmailStateMissing    EmailState = "missing"
	EmailStateUnverified EmailState = "unverified"
	EmailStateVerified   EmailState = "verified"
)

func (e EmailState) ToPointer() *EmailState {
	return &e
}

func (e *EmailState) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "missing":
		fallthrough
	case "unverified":
		fallthrough
	case "verified":
		*e = EmailState(v)
		return nil
	default:
		return fmt.Errorf("invalid value for EmailState: %v", v)
	}
}

type PhoneState string

const (
	PhoneStateMissing    PhoneState = "missing"
	PhoneStateUnverified PhoneState = "unverified"
	PhoneStateVerified   PhoneState = "verified"
)

func (e PhoneState) ToPointer() *PhoneState {
	return &e
}

func (e *PhoneState) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "missing":
		fallthrough
	case "unverified":
		fallthrough
	case "verified":
		*e = PhoneState(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PhoneState: %v", v)
	}
}

type AccountTestCreationDataOutput struct {
	Email      string     `json:"email"`
	EmailState EmailState `json:"email_state"`
	Phone      string     `json:"phone"`
	PhoneState PhoneState `json:"phone_state"`
	OtpCode    string     `json:"otp_code"`
	OauthCode  string     `json:"oauth_code"`
}

func (o *AccountTestCreationDataOutput) GetEmail() string {
	if o == nil {
		return ""
	}
	return o.Email
}

func (o *AccountTestCreationDataOutput) GetEmailState() EmailState {
	if o == nil {
		return EmailState("")
	}
	return o.EmailState
}

func (o *AccountTestCreationDataOutput) GetPhone() string {
	if o == nil {
		return ""
	}
	return o.Phone
}

func (o *AccountTestCreationDataOutput) GetPhoneState() PhoneState {
	if o == nil {
		return PhoneState("")
	}
	return o.PhoneState
}

func (o *AccountTestCreationDataOutput) GetOtpCode() string {
	if o == nil {
		return ""
	}
	return o.OtpCode
}

func (o *AccountTestCreationDataOutput) GetOauthCode() string {
	if o == nil {
		return ""
	}
	return o.OauthCode
}

type AccountTestCreationData struct {
	EmailState EmailState `json:"email_state"`
	PhoneState PhoneState `json:"phone_state"`
	IsMigrated *bool      `json:"is_migrated,omitempty"`
	HasAddress *bool      `json:"has_address,omitempty"`
}

func (o *AccountTestCreationData) GetEmailState() EmailState {
	if o == nil {
		return EmailState("")
	}
	return o.EmailState
}

func (o *AccountTestCreationData) GetPhoneState() PhoneState {
	if o == nil {
		return PhoneState("")
	}
	return o.PhoneState
}

func (o *AccountTestCreationData) GetIsMigrated() *bool {
	if o == nil {
		return nil
	}
	return o.IsMigrated
}

func (o *AccountTestCreationData) GetHasAddress() *bool {
	if o == nil {
		return nil
	}
	return o.HasAddress
}
