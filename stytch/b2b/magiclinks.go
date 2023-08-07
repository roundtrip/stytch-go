package b2b

// !!!
// WARNING: This file is autogenerated
// Only modify code within MANUAL() sections
// or your changes may be overwritten later!
// !!!

import (
	"context"
	"encoding/json"

	"github.com/stytchauth/stytch-go/v10/stytch"
	"github.com/stytchauth/stytch-go/v10/stytch/b2b/magiclinks"
	"github.com/stytchauth/stytch-go/v10/stytch/stytcherror"
)

type MagicLinksClient struct {
	C         stytch.Client
	Email     *MagicLinksEmailClient
	Discovery *MagicLinksDiscoveryClient
}

func NewMagicLinksClient(c stytch.Client) *MagicLinksClient {
	return &MagicLinksClient{
		C:         c,
		Email:     NewMagicLinksEmailClient(c),
		Discovery: NewMagicLinksDiscoveryClient(c),
	}
}

// Authenticate a Member with a Magic Link. This endpoint requires a Magic Link token that is not expired
// or previously used. If the Member’s status is `pending` or `invited`, they will be updated to `active`.
// Provide the `session_duration_minutes` parameter to set the lifetime of the session. If the
// `session_duration_minutes` parameter is not specified, a Stytch session will be created with a 60 minute
// duration.
//
// (Coming Soon) If the Member is required to complete MFA to log in to the Organization, the returned
// value of `member_authenticated` will be `false`, and an `intermediate_session_token` will be returned.
// The `intermediate_session_token` can be passed into the
// [OTP SMS Authenticate endpoint](https://stytch.com/docs/b2b/api/authenticate-otp-sms) to complete the
// MFA step and acquire a full member session.
// The `intermediate_session_token` can also be used with the
// [Exchange Intermediate Session endpoint](https://stytch.com/docs/b2b/api/exchange-intermediate-session)
// or the
// [Create Organization via Discovery endpoint](https://stytch.com/docs/b2b/api/create-organization-via-discovery) to join a different Organization or create a new one.
// The `session_duration_minutes` and `session_custom_claims` parameters will be ignored.
//
// If a valid `session_token` or `session_jwt` is passed in, the Member will not be required to complete an
// MFA step.
func (c *MagicLinksClient) Authenticate(
	ctx context.Context,
	body *magiclinks.AuthenticateParams,
) (*magiclinks.AuthenticateResponse, error) {
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("error marshaling request body")
		}
	}

	var retVal magiclinks.AuthenticateResponse
	err = c.C.NewRequest(
		ctx,
		"POST",
		"/v1/b2b/magic_links/authenticate",
		nil,
		jsonBody,
		&retVal,
	)
	return &retVal, err
}
