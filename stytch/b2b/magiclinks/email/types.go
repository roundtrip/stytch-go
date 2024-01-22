package email

// !!!
// WARNING: This file is autogenerated
// Only modify code within MANUAL() sections
// or your changes may be overwritten later!
// !!!

import (
	"github.com/stytchauth/stytch-go/v12/stytch/b2b/organizations"
	"github.com/stytchauth/stytch-go/v12/stytch/methodoptions"
)

// InviteParams: Request type for `Email.Invite`.
type InviteParams struct {
	// OrganizationID: Globally unique UUID that identifies a specific Organization. The `organization_id` is
	// critical to perform operations on an Organization, so be sure to preserve this value.
	OrganizationID string `json:"organization_id,omitempty"`
	// EmailAddress: The email address of the Member.
	EmailAddress string `json:"email_address,omitempty"`
	// InviteRedirectURL: The URL that the Member clicks from the invite Email Magic Link. This URL should be
	// an endpoint in the backend server that verifies
	//   the request by querying Stytch's authenticate endpoint and finishes the invite flow. If this value is
	// not passed, the default `invite_redirect_url`
	//   that you set in your Dashboard is used. If you have not set a default `invite_redirect_url`, an error
	// is returned.
	InviteRedirectURL string `json:"invite_redirect_url,omitempty"`
	// InvitedByMemberID: The `member_id` of the Member who sends the invite.
	InvitedByMemberID string `json:"invited_by_member_id,omitempty"`
	// Name: The name of the Member.
	Name string `json:"name,omitempty"`
	// TrustedMetadata: An arbitrary JSON object for storing application-specific data or
	// identity-provider-specific data.
	TrustedMetadata map[string]any `json:"trusted_metadata,omitempty"`
	// UntrustedMetadata: An arbitrary JSON object of application-specific data. These fields can be edited
	// directly by the
	//   frontend SDK, and should not be used to store critical information. See the
	// [Metadata resource](https://stytch.com/docs/b2b/api/metadata)
	//   for complete field behavior details.
	UntrustedMetadata map[string]any `json:"untrusted_metadata,omitempty"`
	// InviteTemplateID: Use a custom template for invite emails. By default, it will use your default email
	// template. The template must be a template
	//   using our built-in customizations or a custom HTML email for Magic Links - Invite.
	InviteTemplateID string `json:"invite_template_id,omitempty"`
	// Locale: Used to determine which language to use when sending the user this delivery method. Parameter is
	// a [IETF BCP 47 language tag](https://www.w3.org/International/articles/language-tags/), e.g. `"en"`.
	//
	// Currently supported languages are English (`"en"`), Spanish (`"es"`), and Brazilian Portuguese
	// (`"pt-br"`); if no value is provided, the copy defaults to English.
	//
	// Request support for additional languages
	// [here](https://docs.google.com/forms/d/e/1FAIpQLScZSpAu_m2AmLXRT3F3kap-s_mcV6UTBitYn6CdyWP0-o7YjQ/viewform?usp=sf_link")!
	//
	Locale InviteRequestLocale `json:"locale,omitempty"`
	// Roles: (Coming Soon) Roles to explicitly assign to this Member. See the
	// [RBAC guide](https://stytch.com/docs/b2b/guides/rbac/role-assignment)
	//    for more information about role assignment.
	Roles []string `json:"roles,omitempty"`
}

// LoginOrSignupParams: Request type for `Email.LoginOrSignup`.
type LoginOrSignupParams struct {
	// OrganizationID: Globally unique UUID that identifies a specific Organization. The `organization_id` is
	// critical to perform operations on an Organization, so be sure to preserve this value.
	OrganizationID string `json:"organization_id,omitempty"`
	// EmailAddress: The email address of the Member.
	EmailAddress string `json:"email_address,omitempty"`
	// LoginRedirectURL: The URL that the Member clicks from the login Email Magic Link. This URL should be an
	// endpoint in the backend server that
	//   verifies the request by querying Stytch's authenticate endpoint and finishes the login. If this value
	// is not passed, the default login
	//   redirect URL that you set in your Dashboard is used. If you have not set a default login redirect URL,
	// an error is returned.
	LoginRedirectURL string `json:"login_redirect_url,omitempty"`
	// SignupRedirectURL: The URL the Member clicks from the signup Email Magic Link. This URL should be an
	// endpoint in the backend server that verifies
	//   the request by querying Stytch's authenticate endpoint and finishes the login. If this value is not
	// passed, the default sign-up redirect URL
	//   that you set in your Dashboard is used. If you have not set a default sign-up redirect URL, an error
	// is returned.
	SignupRedirectURL string `json:"signup_redirect_url,omitempty"`
	// PkceCodeChallenge: A base64url encoded SHA256 hash of a one time secret used to validate that the
	// request starts and ends on the same device.
	PkceCodeChallenge string `json:"pkce_code_challenge,omitempty"`
	// LoginTemplateID: Use a custom template for login emails. By default, it will use your default email
	// template. The template must be from Stytch's
	// built-in customizations or a custom HTML email for Magic Links - Login.
	LoginTemplateID string `json:"login_template_id,omitempty"`
	// SignupTemplateID: Use a custom template for signup emails. By default, it will use your default email
	// template. The template must be from Stytch's
	// built-in customizations or a custom HTML email for Magic Links - Signup.
	SignupTemplateID string `json:"signup_template_id,omitempty"`
	// Locale: Used to determine which language to use when sending the user this delivery method. Parameter is
	// a [IETF BCP 47 language tag](https://www.w3.org/International/articles/language-tags/), e.g. `"en"`.
	//
	// Currently supported languages are English (`"en"`), Spanish (`"es"`), and Brazilian Portuguese
	// (`"pt-br"`); if no value is provided, the copy defaults to English.
	//
	// Request support for additional languages
	// [here](https://docs.google.com/forms/d/e/1FAIpQLScZSpAu_m2AmLXRT3F3kap-s_mcV6UTBitYn6CdyWP0-o7YjQ/viewform?usp=sf_link")!
	//
	Locale LoginOrSignupRequestLocale `json:"locale,omitempty"`
}

// InviteRequestOptions:
type InviteRequestOptions struct {
	// Authorization: Optional authorization object.
	// Pass in an active Stytch Member session token or session JWT and the request
	// will be run using that member's permissions.
	Authorization methodoptions.Authorization `json:"authorization,omitempty"`
}

func (o *InviteRequestOptions) AddHeaders(headers map[string][]string) map[string][]string {
	headers = o.Authorization.AddHeaders(headers)
	return headers
}

// InviteResponse: Response type for `Email.Invite`.
type InviteResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// MemberID: Globally unique UUID that identifies a specific Member.
	MemberID string `json:"member_id,omitempty"`
	// Member: The [Member object](https://stytch.com/docs/b2b/api/member-object)
	Member organizations.Member `json:"member,omitempty"`
	// Organization: The [Organization object](https://stytch.com/docs/b2b/api/organization-object).
	Organization organizations.Organization `json:"organization,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
}

// LoginOrSignupResponse: Response type for `Email.LoginOrSignup`.
type LoginOrSignupResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// MemberID: Globally unique UUID that identifies a specific Member.
	MemberID string `json:"member_id,omitempty"`
	// MemberCreated: A flag indicating `true` if a new Member object was created and `false` if the Member
	// object already existed.
	MemberCreated bool `json:"member_created,omitempty"`
	// Member: The [Member object](https://stytch.com/docs/b2b/api/member-object)
	Member organizations.Member `json:"member,omitempty"`
	// Organization: The [Organization object](https://stytch.com/docs/b2b/api/organization-object).
	Organization organizations.Organization `json:"organization,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
}

type InviteRequestLocale string

const (
	InviteRequestLocaleEn   InviteRequestLocale = "en"
	InviteRequestLocaleEs   InviteRequestLocale = "es"
	InviteRequestLocalePtbr InviteRequestLocale = "pt-br"
)

type LoginOrSignupRequestLocale string

const (
	LoginOrSignupRequestLocaleEn   LoginOrSignupRequestLocale = "en"
	LoginOrSignupRequestLocaleEs   LoginOrSignupRequestLocale = "es"
	LoginOrSignupRequestLocalePtbr LoginOrSignupRequestLocale = "pt-br"
)
