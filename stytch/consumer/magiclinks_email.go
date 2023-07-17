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
	"github.com/stytchauth/stytch-go/v9/stytch/consumer/magiclinks/email"
	"github.com/stytchauth/stytch-go/v9/stytch/stytcherror"
)

type MagicLinksEmailClient struct {
	C stytch.Client
}

func NewMagicLinksEmailClient(c stytch.Client) *MagicLinksEmailClient {
	return &MagicLinksEmailClient{
		C: c,
	}
}

// Send a magic link to an existing Stytch user using their email address. If you'd like to create a user
// and send them a magic link by email with one request, use our
// [log in or create endpoint](https://stytch.com/docs/api/log-in-or-create-user-by-email).
//
// ### Add an email to an existing user
// This endpoint also allows you to add a new email to an existing Stytch User. Including a `user_id`,
// `session_token`, or `session_jwt` in the request will add the email to the pre-existing Stytch User upon
// successful authentication.
//
// Adding a new email to an existing Stytch User requires the user to be present and validate the email via
// magic link. This requirement is in place to prevent account takeover attacks.
//
// ### Next steps
// The user is emailed a magic link which redirects them to the provided
// [redirect URL](https://stytch.com/docs/magic-links#email-magic-links_redirect-routing). Collect the
// `token` from the URL query parameters, and call
// [Authenticate magic link](https://stytch.com/docs/api/authenticate-magic-link) to complete
// authentication.
func (c *MagicLinksEmailClient) Send(
	ctx context.Context,
	body *email.SendParams,
) (*email.SendResponse, error) {
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("error marshaling request body")
		}
	}

	var retVal email.SendResponse
	err = c.C.NewRequest(
		ctx,
		"POST",
		"/v1/magic_links/email/send",
		nil,
		jsonBody,
		&retVal,
	)
	return &retVal, err
}

// LoginOrCreate: Send either a login or signup Magic Link to the User based on if the email is associated
// with a User already. A new or pending User will receive a signup Magic Link. An active User will receive
// a login Magic Link. For more information on how to control the status your Users are created in see the
// `create_user_as_pending` flag.
//
// ### Next steps
// The User is emailed a Magic Link which redirects them to the provided
// [redirect URL](https://stytch.com/docs/magic-links#email-magic-links_redirect-routing). Collect the
// `token` from the URL query parameters and call
// [Authenticate Magic Link](https://stytch.com/docs/api/authenticate-magic-link) to complete
// authentication.
func (c *MagicLinksEmailClient) LoginOrCreate(
	ctx context.Context,
	body *email.LoginOrCreateParams,
) (*email.LoginOrCreateResponse, error) {
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("error marshaling request body")
		}
	}

	var retVal email.LoginOrCreateResponse
	err = c.C.NewRequest(
		ctx,
		"POST",
		"/v1/magic_links/email/login_or_create",
		nil,
		jsonBody,
		&retVal,
	)
	return &retVal, err
}

// Invite: Create a User and send an invite Magic Link to the provided `email`. The User will be created
// with a `pending` status until they click the Magic Link in the invite email.
//
// ### Next steps
// The User is emailed a Magic Link which redirects them to the provided
// [redirect URL](https://stytch.com/docs/magic-links#email-magic-links_redirect-routing). Collect the
// `token` from the URL query parameters and call
// [Authenticate Magic Link](https://stytch.com/docs/api/authenticate-magic-link) to complete
// authentication.
func (c *MagicLinksEmailClient) Invite(
	ctx context.Context,
	body *email.InviteParams,
) (*email.InviteResponse, error) {
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("error marshaling request body")
		}
	}

	var retVal email.InviteResponse
	err = c.C.NewRequest(
		ctx,
		"POST",
		"/v1/magic_links/email/invite",
		nil,
		jsonBody,
		&retVal,
	)
	return &retVal, err
}

// RevokeInvite: Revoke a pending invite based on the `email` provided.
func (c *MagicLinksEmailClient) RevokeInvite(
	ctx context.Context,
	body *email.RevokeInviteParams,
) (*email.RevokeInviteResponse, error) {
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("error marshaling request body")
		}
	}

	var retVal email.RevokeInviteResponse
	err = c.C.NewRequest(
		ctx,
		"POST",
		"/v1/magic_links/email/revoke_invite",
		nil,
		jsonBody,
		&retVal,
	)
	return &retVal, err
}
