package b2b

// !!!
// WARNING: This file is autogenerated
// Only modify code within MANUAL() sections
// or your changes may be overwritten later!
// !!!

import (
	"context"
	"fmt"

	"github.com/stytchauth/stytch-go/v13/stytch"
	"github.com/stytchauth/stytch-go/v13/stytch/b2b/organizations/members/oauthproviders"
)

type OrganizationsMembersOAuthProvidersClient struct {
	C stytch.Client
}

func NewOrganizationsMembersOAuthProvidersClient(c stytch.Client) *OrganizationsMembersOAuthProvidersClient {
	return &OrganizationsMembersOAuthProvidersClient{
		C: c,
	}
}

// Google: Retrieve the saved Google access token and ID token for a member. After a successful OAuth
// login, Stytch will save the
// issued access token and ID token from the identity provider. If a refresh token has been issued, Stytch
// will refresh the
// access token automatically.
//
// __Note:__ Google does not issue a refresh token on every login, and refresh tokens may expire if unused.
// To force a refresh token to be issued, pass the `?provider_prompt=consent` query param into the
// [Start Google OAuth flow](https://stytch.com/docs/b2b/api/oauth-google-start) endpoint.
func (c *OrganizationsMembersOAuthProvidersClient) Google(
	ctx context.Context,
	body *oauthproviders.ProviderInformationParams,
) (*oauthproviders.GoogleResponse, error) {
	queryParams := make(map[string]string)
	if body != nil {
		if body.IncludeRefreshToken {
			queryParams["include_refresh_token"] = "true"
		} else {
			queryParams["include_refresh_token"] = "false"
		}
	}

	headers := make(map[string][]string)

	var retVal oauthproviders.GoogleResponse
	err := c.C.NewRequest(
		ctx,
		"GET",
		fmt.Sprintf("/v1/b2b/organizations/%s/members/%s/oauth_providers/google", body.OrganizationID, body.MemberID),
		queryParams,
		nil,
		&retVal,
		headers,
	)
	return &retVal, err
}

// Microsoft: Retrieve the saved Microsoft access token and ID token for a member. After a successful OAuth
// login, Stytch will save the
// issued access token and ID token from the identity provider. If a refresh token has been issued, Stytch
// will refresh the
// access token automatically.
func (c *OrganizationsMembersOAuthProvidersClient) Microsoft(
	ctx context.Context,
	body *oauthproviders.ProviderInformationParams,
) (*oauthproviders.MicrosoftResponse, error) {
	queryParams := make(map[string]string)
	if body != nil {
		if body.IncludeRefreshToken {
			queryParams["include_refresh_token"] = "true"
		} else {
			queryParams["include_refresh_token"] = "false"
		}
	}

	headers := make(map[string][]string)

	var retVal oauthproviders.MicrosoftResponse
	err := c.C.NewRequest(
		ctx,
		"GET",
		fmt.Sprintf("/v1/b2b/organizations/%s/members/%s/oauth_providers/microsoft", body.OrganizationID, body.MemberID),
		queryParams,
		nil,
		&retVal,
		headers,
	)
	return &retVal, err
}
