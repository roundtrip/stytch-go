package discovery

// !!!
// WARNING: This file is autogenerated
// Only modify code within MANUAL() sections
// or your changes may be overwritten later!
// !!!

import (
	"github.com/stytchauth/stytch-go/v8/stytch/b2b/organizations"
)

// DiscoveredOrganization:
type DiscoveredOrganization struct {
	// MemberAuthenticated: Indicates whether or not the discovery magic link initiated session is valid for
	// the organization's allowed auth method settings.
	//   If not, the member needs to perform additional authentication before logging in - such as password or
	// SSO auth.
	MemberAuthenticated bool `json:"member_authenticated,omitempty"`
	// Organization: The [Organization object](https://stytch.com/docs/b2b/api/organization-object).
	Organization organizations.Organization `json:"organization,omitempty"`
	// Membership: Information about the membership.
	Membership Membership `json:"membership,omitempty"`
}

// Membership:
type Membership struct {
	// Type: Either `active_member`, `pending_member`, `invited_member`, or `eligible_to_join_by_email_domain`
	Type string `json:"type,omitempty"`
	// Details: An object containing additional metadata about the membership, if available.
	Details map[string]any `json:"details,omitempty"`
	// Member: The [Member object](https://stytch.com/docs/b2b/api/member-object) if one already exists, or
	// null if one does not.
	Member organizations.Member `json:"member,omitempty"`
}