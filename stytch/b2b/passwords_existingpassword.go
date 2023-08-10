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
	"github.com/stytchauth/stytch-go/v10/stytch/b2b/passwords/existingpassword"
	"github.com/stytchauth/stytch-go/v10/stytch/stytcherror"
)

type PasswordsExistingPasswordClient struct {
	C stytch.Client
}

func NewPasswordsExistingPasswordClient(c stytch.Client) *PasswordsExistingPasswordClient {
	return &PasswordsExistingPasswordClient{
		C: c,
	}
}

// Reset the member’s password using their existing password.
//
// This endpoint adapts to your Project's password strength configuration.
// If you're using [zxcvbn](https://stytch.com/docs/guides/passwords/strength-policy), the default, your
// passwords are considered valid
// if the strength score is >= 3. If you're using
// [LUDS](https://stytch.com/docs/guides/passwords/strength-policy), your passwords are
// considered valid if they meet the requirements that you've set with Stytch.
// You may update your password strength configuration in the
// [stytch dashboard](https://stytch.com/dashboard/password-strength-config).
//
// If the Member is required to complete MFA to log in to the Organization, the returned value of
// `member_authenticated` will be `false`, and an `intermediate_session_token` will be returned.
// The `intermediate_session_token` can be passed into the
// [OTP SMS Authenticate endpoint](https://stytch.com/docs/b2b/api/authenticate-otp-sms) to complete the
// MFA step and acquire a full member session.
// The `session_duration_minutes` and `session_custom_claims` parameters will be ignored.
//
// If a valid `session_token` or `session_jwt` is passed in, the Member will not be required to complete an
// MFA step.
func (c *PasswordsExistingPasswordClient) Reset(
	ctx context.Context,
	body *existingpassword.ResetParams,
) (*existingpassword.ResetResponse, error) {
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("error marshaling request body")
		}
	}

	var retVal existingpassword.ResetResponse
	err = c.C.NewRequest(
		ctx,
		"POST",
		"/v1/b2b/passwords/existing_password/reset",
		nil,
		jsonBody,
		&retVal,
	)
	return &retVal, err
}
