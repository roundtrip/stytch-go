package whatsapp

// !!!
// WARNING: This file is autogenerated
// Only modify code within MANUAL() sections
// or your changes may be overwritten later!
// !!!

import (
	"github.com/stytchauth/stytch-go/v11/stytch/consumer/attribute"
)

// LoginOrCreateParams: Request type for `Whatsapp.LoginOrCreate`.
type LoginOrCreateParams struct {
	// PhoneNumber: The phone number to use for one-time passcodes. The phone number should be in E.164 format
	// (i.e. +1XXXXXXXXXX). You may use +10000000000 to test this endpoint, see
	// [Testing](https://stytch.com/docs/home#resources_testing) for more detail.
	PhoneNumber string `json:"phone_number,omitempty"`
	// ExpirationMinutes: Set the expiration for the one-time passcode, in minutes. The minimum expiration is 1
	// minute and the maximum is 10 minutes. The default expiration is 2 minutes.
	ExpirationMinutes int32 `json:"expiration_minutes,omitempty"`
	// Attributes: Provided attributes help with fraud detection.
	Attributes *attribute.Attributes `json:"attributes,omitempty"`
	// CreateUserAsPending: Flag for whether or not to save a user as pending vs active in Stytch. Defaults to
	// false.
	//         If true, users will be saved with status pending in Stytch's backend until authenticated.
	//         If false, users will be created as active. An example usage of
	//         a true flag would be to require users to verify their phone by entering the OTP code before
	// creating
	//         an account for them.
	CreateUserAsPending bool `json:"create_user_as_pending,omitempty"`
	// Locale: Used to determine which language to use when sending the user this delivery method. Parameter is
	// a [IETF BCP 47 language tag](https://www.w3.org/International/articles/language-tags/), e.g. `"en"`.
	//
	// Currently supported languages are English (`"en"`), Spanish (`"es"`), and Brazilian Portuguese
	// (`"pt-br"`); if no value is provided, the copy defaults to English.
	//
	// Request support for additional languages
	// [here](https://docs.google.com/forms/d/e/1FAIpQLScZSpAu_m2AmLXRT3F3kap-s_mcV6UTBitYn6CdyWP0-o7YjQ/viewform?usp=sf_link")!
	//
	Locale LoginOrCreateRequestLocale `json:"locale,omitempty"`
}

// SendParams: Request type for `Whatsapp.Send`.
type SendParams struct {
	// PhoneNumber: The phone number to use for one-time passcodes. The phone number should be in E.164 format
	// (i.e. +1XXXXXXXXXX). You may use +10000000000 to test this endpoint, see
	// [Testing](https://stytch.com/docs/home#resources_testing) for more detail.
	PhoneNumber string `json:"phone_number,omitempty"`
	// ExpirationMinutes: Set the expiration for the one-time passcode, in minutes. The minimum expiration is 1
	// minute and the maximum is 10 minutes. The default expiration is 2 minutes.
	ExpirationMinutes int32 `json:"expiration_minutes,omitempty"`
	// Attributes: Provided attributes help with fraud detection.
	Attributes *attribute.Attributes `json:"attributes,omitempty"`
	// Locale: Used to determine which language to use when sending the user this delivery method. Parameter is
	// a [IETF BCP 47 language tag](https://www.w3.org/International/articles/language-tags/), e.g. `"en"`.
	//
	// Currently supported languages are English (`"en"`), Spanish (`"es"`), and Brazilian Portuguese
	// (`"pt-br"`); if no value is provided, the copy defaults to English.
	//
	// Request support for additional languages
	// [here](https://docs.google.com/forms/d/e/1FAIpQLScZSpAu_m2AmLXRT3F3kap-s_mcV6UTBitYn6CdyWP0-o7YjQ/viewform?usp=sf_link")!
	//
	Locale SendRequestLocale `json:"locale,omitempty"`
	// UserID: The unique ID of a specific User.
	UserID string `json:"user_id,omitempty"`
	// SessionToken: The `session_token` associated with a User's existing Session.
	SessionToken string `json:"session_token,omitempty"`
	// SessionJWT: The `session_jwt` associated with a User's existing Session.
	SessionJWT string `json:"session_jwt,omitempty"`
}

// LoginOrCreateResponse: Response type for `Whatsapp.LoginOrCreate`.
type LoginOrCreateResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// UserID: The unique ID of the affected User.
	UserID string `json:"user_id,omitempty"`
	// PhoneID: The unique ID for the phone number.
	PhoneID string `json:"phone_id,omitempty"`
	// UserCreated: In `login_or_create` endpoints, this field indicates whether or not a User was just created.
	UserCreated bool `json:"user_created,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
}

// SendResponse: Response type for `Whatsapp.Send`.
type SendResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// UserID: The unique ID of the affected User.
	UserID string `json:"user_id,omitempty"`
	// PhoneID: The unique ID for the phone number.
	PhoneID string `json:"phone_id,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
}

type LoginOrCreateRequestLocale string

const (
	LoginOrCreateRequestLocaleEn   LoginOrCreateRequestLocale = "en"
	LoginOrCreateRequestLocaleEs   LoginOrCreateRequestLocale = "es"
	LoginOrCreateRequestLocalePtbr LoginOrCreateRequestLocale = "pt-br"
)

type SendRequestLocale string

const (
	SendRequestLocaleEn   SendRequestLocale = "en"
	SendRequestLocaleEs   SendRequestLocale = "es"
	SendRequestLocalePtbr SendRequestLocale = "pt-br"
)
