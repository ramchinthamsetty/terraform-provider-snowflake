---
page_title: "snowflake_api_authentication_integration_with_jwt_bearer Resource - terraform-provider-snowflake"
subcategory: ""
description: |-
  Resource used to manage api authentication security integration objects with jwt bearer. For more information, check security integrations documentation https://docs.snowflake.com/en/sql-reference/sql/create-security-integration-api-auth.
---

!> **V1 release candidate** This resource was reworked and is a release candidate for the V1. We do not expect significant changes in it before the V1. We will welcome any feedback and adjust the resource if needed. Any errors reported will be resolved with a higher priority. We encourage checking this resource out before the V1 release. Please follow the [migration guide](https://github.com/Snowflake-Labs/terraform-provider-snowflake/blob/main/MIGRATION_GUIDE.md#v0920--v0930) to use it.

# snowflake_api_authentication_integration_with_jwt_bearer (Resource)

Resource used to manage api authentication security integration objects with jwt bearer. For more information, check [security integrations documentation](https://docs.snowflake.com/en/sql-reference/sql/create-security-integration-api-auth).

## Example Usage

```terraform
# basic resource
resource "snowflake_api_authentication_integration_with_jwt_bearer" "test" {
  enabled                = true
  name                   = "test"
  oauth_client_id        = "sn-oauth-134o9erqfedlc"
  oauth_client_secret    = "eb9vaXsrcEvrFdfcvCaoijhilj4fc"
  oauth_assertion_issuer = "issuer"
}
# resource with all fields set
resource "snowflake_api_authentication_integration_with_jwt_bearer" "test" {
  comment                      = "comment"
  enabled                      = true
  name                         = "test"
  oauth_access_token_validity  = 42
  oauth_authorization_endpoint = "https://example.com"
  oauth_client_auth_method     = "CLIENT_SECRET_POST"
  oauth_client_id              = "sn-oauth-134o9erqfedlc"
  oauth_client_secret          = "eb9vaXsrcEvrFdfcvCaoijhilj4fc"
  oauth_refresh_token_validity = 42
  oauth_token_endpoint         = "https://example.com"
  oauth_assertion_issuer       = "issuer"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `enabled` (Boolean) Specifies whether this security integration is enabled or disabled.
- `name` (String) Specifies the identifier (i.e. name) for the integration. This value must be unique in your account.
- `oauth_assertion_issuer` (String)
- `oauth_client_id` (String) Specifies the client ID for the OAuth application in the external service.
- `oauth_client_secret` (String) Specifies the client secret for the OAuth application in the ServiceNow instance from the previous step. The connector uses this to request an access token from the ServiceNow instance.

### Optional

- `comment` (String) Specifies a comment for the integration.
- `oauth_access_token_validity` (Number) Specifies the default lifetime of the OAuth access token (in seconds) issued by an OAuth server.
- `oauth_authorization_endpoint` (String) Specifies the URL for authenticating to the external service.
- `oauth_client_auth_method` (String) Specifies that POST is used as the authentication method to the external service. If removed from the config, the resource is recreated. Valid values are (case-insensitive): `CLIENT_SECRET_POST`.
- `oauth_refresh_token_validity` (Number) Specifies the value to determine the validity of the refresh token obtained from the OAuth server.
- `oauth_token_endpoint` (String) Specifies the token endpoint used by the client to obtain an access token by presenting its authorization grant or refresh token. The token endpoint is used with every authorization grant except for the implicit grant type (since an access token is issued directly). If removed from the config, the resource is recreated.

### Read-Only

- `describe_output` (List of Object) Outputs the result of `DESCRIBE SECURITY INTEGRATIONS` for the given security integration. (see [below for nested schema](#nestedatt--describe_output))
- `id` (String) The ID of this resource.
- `show_output` (List of Object) Outputs the result of `SHOW SECURITY INTEGRATIONS` for the given security integration. (see [below for nested schema](#nestedatt--show_output))

<a id="nestedatt--describe_output"></a>
### Nested Schema for `describe_output`

Read-Only:

- `auth_type` (List of Object) (see [below for nested schema](#nestedobjatt--describe_output--auth_type))
- `comment` (List of Object) (see [below for nested schema](#nestedobjatt--describe_output--comment))
- `enabled` (List of Object) (see [below for nested schema](#nestedobjatt--describe_output--enabled))
- `oauth_access_token_validity` (List of Object) (see [below for nested schema](#nestedobjatt--describe_output--oauth_access_token_validity))
- `oauth_allowed_scopes` (List of Object) (see [below for nested schema](#nestedobjatt--describe_output--oauth_allowed_scopes))
- `oauth_authorization_endpoint` (List of Object) (see [below for nested schema](#nestedobjatt--describe_output--oauth_authorization_endpoint))
- `oauth_client_auth_method` (List of Object) (see [below for nested schema](#nestedobjatt--describe_output--oauth_client_auth_method))
- `oauth_client_id` (List of Object) (see [below for nested schema](#nestedobjatt--describe_output--oauth_client_id))
- `oauth_grant` (List of Object) (see [below for nested schema](#nestedobjatt--describe_output--oauth_grant))
- `oauth_refresh_token_validity` (List of Object) (see [below for nested schema](#nestedobjatt--describe_output--oauth_refresh_token_validity))
- `oauth_token_endpoint` (List of Object) (see [below for nested schema](#nestedobjatt--describe_output--oauth_token_endpoint))
- `parent_integration` (List of Object) (see [below for nested schema](#nestedobjatt--describe_output--parent_integration))

<a id="nestedobjatt--describe_output--auth_type"></a>
### Nested Schema for `describe_output.auth_type`

Read-Only:

- `default` (String)
- `name` (String)
- `type` (String)
- `value` (String)


<a id="nestedobjatt--describe_output--comment"></a>
### Nested Schema for `describe_output.comment`

Read-Only:

- `default` (String)
- `name` (String)
- `type` (String)
- `value` (String)


<a id="nestedobjatt--describe_output--enabled"></a>
### Nested Schema for `describe_output.enabled`

Read-Only:

- `default` (String)
- `name` (String)
- `type` (String)
- `value` (String)


<a id="nestedobjatt--describe_output--oauth_access_token_validity"></a>
### Nested Schema for `describe_output.oauth_access_token_validity`

Read-Only:

- `default` (String)
- `name` (String)
- `type` (String)
- `value` (String)


<a id="nestedobjatt--describe_output--oauth_allowed_scopes"></a>
### Nested Schema for `describe_output.oauth_allowed_scopes`

Read-Only:

- `default` (String)
- `name` (String)
- `type` (String)
- `value` (String)


<a id="nestedobjatt--describe_output--oauth_authorization_endpoint"></a>
### Nested Schema for `describe_output.oauth_authorization_endpoint`

Read-Only:

- `default` (String)
- `name` (String)
- `type` (String)
- `value` (String)


<a id="nestedobjatt--describe_output--oauth_client_auth_method"></a>
### Nested Schema for `describe_output.oauth_client_auth_method`

Read-Only:

- `default` (String)
- `name` (String)
- `type` (String)
- `value` (String)


<a id="nestedobjatt--describe_output--oauth_client_id"></a>
### Nested Schema for `describe_output.oauth_client_id`

Read-Only:

- `default` (String)
- `name` (String)
- `type` (String)
- `value` (String)


<a id="nestedobjatt--describe_output--oauth_grant"></a>
### Nested Schema for `describe_output.oauth_grant`

Read-Only:

- `default` (String)
- `name` (String)
- `type` (String)
- `value` (String)


<a id="nestedobjatt--describe_output--oauth_refresh_token_validity"></a>
### Nested Schema for `describe_output.oauth_refresh_token_validity`

Read-Only:

- `default` (String)
- `name` (String)
- `type` (String)
- `value` (String)


<a id="nestedobjatt--describe_output--oauth_token_endpoint"></a>
### Nested Schema for `describe_output.oauth_token_endpoint`

Read-Only:

- `default` (String)
- `name` (String)
- `type` (String)
- `value` (String)


<a id="nestedobjatt--describe_output--parent_integration"></a>
### Nested Schema for `describe_output.parent_integration`

Read-Only:

- `default` (String)
- `name` (String)
- `type` (String)
- `value` (String)



<a id="nestedatt--show_output"></a>
### Nested Schema for `show_output`

Read-Only:

- `category` (String)
- `comment` (String)
- `created_on` (String)
- `enabled` (Boolean)
- `integration_type` (String)
- `name` (String)

## Import

Import is supported using the following syntax:

```shell
terraform import snowflake_api_authentication_integration_with_jwt_bearer.example "name"
```