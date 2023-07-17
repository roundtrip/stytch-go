package magiclinks

// !!!
// WARNING: This file is autogenerated
// Only modify code within MANUAL() sections
// or your changes may be overwritten later!
// !!!

import (
	"github.com/stytchauth/stytch-go/v9/stytch/b2b/organizations"
	"github.com/stytchauth/stytch-go/v9/stytch/b2b/sessions"
)

// AuthenticateParams: Request type for `MagicLinks.Authenticate`.
type AuthenticateParams struct {
	// MagicLinksToken: The Email Magic Link token to authenticate.
	MagicLinksToken string `json:"magic_links_token,omitempty"`
	// PkceCodeVerifier: A base64url encoded one time secret used to validate that the request starts and ends
	// on the same device.
	PkceCodeVerifier string `json:"pkce_code_verifier,omitempty"`
	// SessionToken: Reuse an existing session instead of creating a new one. If you provide a `session_token`,
	// Stytch will update the session.
	//       If the `session_token` and `magic_links_token` belong to different Members, the `session_token`
	// will be ignored. This endpoint will error if
	//       both `session_token` and `session_jwt` are provided.
	SessionToken string `json:"session_token,omitempty"`
	// SessionJWT: Reuse an existing session instead of creating a new one. If you provide a `session_jwt`,
	// Stytch will update the session. If the `session_jwt`
	//       and `magic_links_token` belong to different Members, the `session_jwt` will be ignored. This
	// endpoint will error if both `session_token` and `session_jwt`
	//       are provided.
	SessionJWT string `json:"session_jwt,omitempty"`
	// SessionDurationMinutes: Set the session lifetime to be this many minutes from now. This will start a new
	// session if one doesn't already exist,
	//   returning both an opaque `session_token` and `session_jwt` for this session. Remember that the
	// `session_jwt` will have a fixed lifetime of
	//   five minutes regardless of the underlying session duration, and will need to be refreshed over time.
	//
	//   This value must be a minimum of 5 and a maximum of 527040 minutes (366 days).
	//
	//   If a `session_token` or `session_jwt` is provided then a successful authentication will continue to
	// extend the session this many minutes.
	//
	//   If the `session_duration_minutes` parameter is not specified, a Stytch session will be created with a
	// 60 minute duration. If you don't want
	//   to use the Stytch session product, you can ignore the session fields in the response.
	SessionDurationMinutes int32 `json:"session_duration_minutes,omitempty"`
	// SessionCustomClaims: Add a custom claims map to the Session being authenticated. Claims are only created
	// if a Session is initialized by providing a value in
	//   `session_duration_minutes`. Claims will be included on the Session object and in the JWT. To update a
	// key in an existing Session, supply a new value. To
	//   delete a key, supply a null value. Custom claims made with reserved claims (`iss`, `sub`, `aud`,
	// `exp`, `nbf`, `iat`, `jti`) will be ignored.
	//   Total custom claims size cannot exceed four kilobytes.
	SessionCustomClaims map[string]any `json:"session_custom_claims,omitempty"`
}

// AuthenticateResponse: Response type for `MagicLinks.Authenticate`.
type AuthenticateResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// MemberID: Globally unique UUID that identifies a specific Member.
	MemberID string `json:"member_id,omitempty"`
	// MethodID: The email or device involved in the authentication.
	MethodID string `json:"method_id,omitempty"`
	// ResetSessions: Indicates if all Sessions linked to the Member need to be reset. You should check this
	// field if you aren't using
	//     Stytch's Session product. If you are using Stytch's Session product, we revoke the Member’s other
	// Sessions for you.
	ResetSessions bool `json:"reset_sessions,omitempty"`
	// OrganizationID: Globally unique UUID that identifies a specific Organization. The `organization_id` is
	// critical to perform operations on an Organization, so be sure to preserve this value.
	OrganizationID string `json:"organization_id,omitempty"`
	// Member: The [Member object](https://stytch.com/docs/b2b/api/member-object).
	Member organizations.Member `json:"member,omitempty"`
	// SessionToken: A secret token for a given Stytch Session.
	SessionToken string `json:"session_token,omitempty"`
	// SessionJWT: The JSON Web Token (JWT) for a given Stytch Session.
	SessionJWT string `json:"session_jwt,omitempty"`
	// MemberSession: The [Session object](https://stytch.com/docs/b2b/api/session-object).
	MemberSession sessions.MemberSession `json:"member_session,omitempty"`
	// Organization: The [Organization object](https://stytch.com/docs/b2b/api/organization-object).
	Organization organizations.Organization `json:"organization,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
}
