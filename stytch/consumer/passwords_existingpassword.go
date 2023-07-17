package consumer

// !!!
// WARNING: This file is autogenerated
// Only modify code within MANUAL() sections
// or your changes may be overwritten later!
// !!!

import (
	"context"
	"encoding/json"

	"github.com/stytchauth/stytch-go/v9/stytch"
	"github.com/stytchauth/stytch-go/v9/stytch/consumer/passwords/existingpassword"
	"github.com/stytchauth/stytch-go/v9/stytch/stytcherror"
)

type PasswordsExistingPasswordClient struct {
	C stytch.Client
}

func NewPasswordsExistingPasswordClient(c stytch.Client) *PasswordsExistingPasswordClient {
	return &PasswordsExistingPasswordClient{
		C: c,
	}
}

// Reset the User’s password using their existing password.
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
		"/v1/passwords/existing_password/reset",
		nil,
		jsonBody,
		&retVal,
	)
	return &retVal, err
}
