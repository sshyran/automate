# Authentication & Authorization

Our work on Authentication and Authorization is based on industry best-practices.

- [OpenID Connect](http://openid.net/connect/)
- [IAM Policies](http://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html)

We evaluated a set of tools that would meet our requirements in [aaa-spike](https://github.com/chef/aaa-spike) and decided for:

- [Dex](https://github.com/coreos/dex) as OpenID solution

We added our own [authn-service](https://github.com/chef/authn-service) as glue between those services to add support for non-human clients. If you're interessted, please read the [design documentation](https://github.com/chef/authn-service/blob/master/DESIGN.md).

![authn diagram](https://github.com/chef/authn-service/raw/master/docs/images/authn_service_situation.png)

For further questions, the team can be reached at [#auth-team](https://chefio.slack.com/messages/C61F9HHKK/)

## Supported Identity Management Systems

OKTA with SAML and AD with LDAP are the most common configurations among our users.
We have also had success with the following configurations:

### SAML

- OKTA
- Ping
- OneLogin
- Office365
- Tivoli Federated Identity Manager

### LDAP

- Microsoft Active Directory (MSAD)

### Configurations with Known Issues

- Azure AD with SAML: we cannot support SAML with Azure AD because Automate cannot consume SAML IdP metadata, which Azure relies on to rotate keys.
- Appleconnect with SAML: invalid signature error.
- In general, dex does not support idP initiated SSO. There is a [PR](https://github.com/dexidp/dex/pull/1514) in progress to add this functionality, last updated Sep 24, 2019.

## authn-service

This is the service used to authenticate requests coming through automate-gateway.

It tries to authenticate each request using the configured authenticators.
The service currently defaults to the following two:

- `oidc`: this expects an `Authorization: bearer JWT` header, where JWT is a Dex-issued JWT ID token
- `header-token`: this expects a `api-token: TOKEN` header, where TOKEN is a token set up
  with the corresponding ("clients") token service. This can be done via automate-ui, or the API.


## session-service

This services provides backs the session management in Automate 2.

It provides a set of simple HTTP endpoints, for consumption by browsers (automate-ui).

For getting a session, it will initiate the _OIDC authorization code flow_, and, when successful, redirect the browser to `/signin#id_token=XYZ&state=ROUTE`.
Also, it stores the `refresh_token` gathered during that login process, and allows for getting a new session using this `refresh_token`, shortcutting the interactive login process.

It acts as an OIDC client to Dex:

- `client_id`: automate-session
- `client_secret`: secretChangeMe
- `redirect_url`: `/session/callback`

For further details, please see [`components/session-service/`](../components/session-service).


## How do our inspec tests work?

When executing the [`a2-api-smoke`](../inspec/a2-api-smoke/) or
[`a2-api-integration`](../inspec/a2-api-integration/), the `automate_api_request` resources provided
by [`a2-resource-pack`](../inspec/a2-resource-pack/) takes care of signing you in.
This happens _for every request_, and does an "actual login", with going to dex, logging in using
the (hardcoded) local user credentials `admin:chefautomate`, and reading the `id_token` from the
final redirect.

The OIDC client configuration used here is

- `client_id`: automate-api
- `client_secret`: *none* (public client)
- `redirect_url`: `http://localhost[:X][/path]` or `urn:ietf:wg:oauth:2.0:oob`

It is using the _OIDC authorization code flow_ as a public client, which ends in an exchange of the `id_token` for the `code`.
The resource plucks the `id_token` out of that response.

* Why doesn't the `automate_api_request` resource use session-service?

  Because it's a few redirects more, and, annoyingly, depends on cookies to work correctly.
  (This doesn't mean that we don't pick that route in the future.)

* There's an example Go method in [`docs/examples/api_request.go`](examples/api_request.go).
