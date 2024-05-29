package sdk

//go:generate go run ./dto-builder-generator/main.go

var (
	_ optionsProvider[CreateExternalOauthSecurityIntegrationOptions]               = new(CreateExternalOauthSecurityIntegrationRequest)
	_ optionsProvider[CreateOauthForPartnerApplicationsSecurityIntegrationOptions] = new(CreateOauthForPartnerApplicationsSecurityIntegrationRequest)
	_ optionsProvider[CreateOauthForCustomClientsSecurityIntegrationOptions]       = new(CreateOauthForCustomClientsSecurityIntegrationRequest)
	_ optionsProvider[CreateSaml2SecurityIntegrationOptions]                       = new(CreateSaml2SecurityIntegrationRequest)
	_ optionsProvider[CreateScimSecurityIntegrationOptions]                        = new(CreateScimSecurityIntegrationRequest)
	_ optionsProvider[AlterExternalOauthSecurityIntegrationOptions]                = new(AlterExternalOauthSecurityIntegrationRequest)
	_ optionsProvider[AlterOauthForPartnerApplicationsSecurityIntegrationOptions]  = new(AlterOauthForPartnerApplicationsSecurityIntegrationRequest)
	_ optionsProvider[AlterOauthForCustomClientsSecurityIntegrationOptions]        = new(AlterOauthForCustomClientsSecurityIntegrationRequest)
	_ optionsProvider[AlterSaml2SecurityIntegrationOptions]                        = new(AlterSaml2SecurityIntegrationRequest)
	_ optionsProvider[AlterScimSecurityIntegrationOptions]                         = new(AlterScimSecurityIntegrationRequest)
	_ optionsProvider[DropSecurityIntegrationOptions]                              = new(DropSecurityIntegrationRequest)
	_ optionsProvider[DescribeSecurityIntegrationOptions]                          = new(DescribeSecurityIntegrationRequest)
	_ optionsProvider[ShowSecurityIntegrationOptions]                              = new(ShowSecurityIntegrationRequest)
)

type CreateExternalOauthSecurityIntegrationRequest struct {
	OrReplace                                  *bool
	IfNotExists                                *bool
	name                                       AccountObjectIdentifier                                             // required
	Enabled                                    bool                                                                // required
	ExternalOauthType                          ExternalOauthSecurityIntegrationTypeOption                          // required
	ExternalOauthIssuer                        string                                                              // required
	ExternalOauthTokenUserMappingClaim         []TokenUserMappingClaim                                             // required
	ExternalOauthSnowflakeUserMappingAttribute ExternalOauthSecurityIntegrationSnowflakeUserMappingAttributeOption // required
	ExternalOauthJwsKeysUrl                    []JwsKeysUrl
	ExternalOauthBlockedRolesList              *BlockedRolesListRequest
	ExternalOauthAllowedRolesList              *AllowedRolesListRequest
	ExternalOauthRsaPublicKey                  *string
	ExternalOauthRsaPublicKey2                 *string
	ExternalOauthAudienceList                  *AudienceListRequest
	ExternalOauthAnyRoleMode                   *ExternalOauthSecurityIntegrationAnyRoleModeOption
	ExternalOauthScopeDelimiter                *string
	ExternalOauthScopeMappingAttribute         *string
	Comment                                    *string
}

type BlockedRolesListRequest struct {
	BlockedRolesList []AccountObjectIdentifier // required
}

type AllowedRolesListRequest struct {
	AllowedRolesList []AccountObjectIdentifier // required
}

type AudienceListRequest struct {
	AudienceList []AudienceListItem // required
}

type CreateOauthForPartnerApplicationsSecurityIntegrationRequest struct {
	OrReplace                 *bool
	IfNotExists               *bool
	name                      AccountObjectIdentifier              // required
	OauthClient               OauthSecurityIntegrationClientOption // required
	OauthRedirectUri          *string
	Enabled                   *bool
	OauthIssueRefreshTokens   *bool
	OauthRefreshTokenValidity *int
	OauthUseSecondaryRoles    *OauthSecurityIntegrationUseSecondaryRolesOption
	BlockedRolesList          *BlockedRolesListRequest
	Comment                   *string
}

func (r *CreateOauthForPartnerApplicationsSecurityIntegrationRequest) GetName() AccountObjectIdentifier {
	return r.name
}

type CreateOauthForCustomClientsSecurityIntegrationRequest struct {
	OrReplace                   *bool
	IfNotExists                 *bool
	name                        AccountObjectIdentifier                  // required
	OauthClientType             OauthSecurityIntegrationClientTypeOption // required
	OauthRedirectUri            string                                   // required
	Enabled                     *bool
	OauthAllowNonTlsRedirectUri *bool
	OauthEnforcePkce            *bool
	OauthUseSecondaryRoles      *OauthSecurityIntegrationUseSecondaryRolesOption
	PreAuthorizedRolesList      *PreAuthorizedRolesListRequest
	BlockedRolesList            *BlockedRolesListRequest
	OauthIssueRefreshTokens     *bool
	OauthRefreshTokenValidity   *int
	NetworkPolicy               *AccountObjectIdentifier
	OauthClientRsaPublicKey     *string
	OauthClientRsaPublicKey2    *string
	Comment                     *string
}

func (r *CreateOauthForCustomClientsSecurityIntegrationRequest) GetName() AccountObjectIdentifier {
	return r.name
}

type PreAuthorizedRolesListRequest struct {
	PreAuthorizedRolesList []AccountObjectIdentifier
}

type CreateSaml2SecurityIntegrationRequest struct {
	OrReplace                      *bool
	IfNotExists                    *bool
	name                           AccountObjectIdentifier // required
	Enabled                        bool                    // required
	Saml2Issuer                    string                  // required
	Saml2SsoUrl                    string                  // required
	Saml2Provider                  string                  // required
	Saml2X509Cert                  string                  // required
	AllowedUserDomains             []UserDomain
	AllowedEmailPatterns           []EmailPattern
	Saml2SpInitiatedLoginPageLabel *string
	Saml2EnableSpInitiated         *bool
	Saml2SnowflakeX509Cert         *string
	Saml2SignRequest               *bool
	Saml2RequestedNameidFormat     *string
	Saml2PostLogoutRedirectUrl     *string
	Saml2ForceAuthn                *bool
	Saml2SnowflakeIssuerUrl        *string
	Saml2SnowflakeAcsUrl           *string
	Comment                        *string
}

func (r *CreateSaml2SecurityIntegrationRequest) GetName() AccountObjectIdentifier {
	return r.name
}

type CreateScimSecurityIntegrationRequest struct {
	OrReplace     *bool
	IfNotExists   *bool
	name          AccountObjectIdentifier                 // required
	Enabled       bool                                    // required
	ScimClient    ScimSecurityIntegrationScimClientOption // required
	RunAsRole     ScimSecurityIntegrationRunAsRoleOption  // required
	NetworkPolicy *AccountObjectIdentifier
	SyncPassword  *bool
	Comment       *string
}

func (r *CreateScimSecurityIntegrationRequest) GetName() AccountObjectIdentifier {
	return r.name
}

type AlterExternalOauthSecurityIntegrationRequest struct {
	IfExists  *bool
	name      AccountObjectIdentifier // required
	SetTags   []TagAssociation
	UnsetTags []ObjectIdentifier
	Set       *ExternalOauthIntegrationSetRequest
	Unset     *ExternalOauthIntegrationUnsetRequest
}

type ExternalOauthIntegrationSetRequest struct {
	Enabled                                    *bool
	ExternalOauthType                          *ExternalOauthSecurityIntegrationTypeOption
	ExternalOauthIssuer                        *string
	ExternalOauthTokenUserMappingClaim         []TokenUserMappingClaim
	ExternalOauthSnowflakeUserMappingAttribute *ExternalOauthSecurityIntegrationSnowflakeUserMappingAttributeOption
	ExternalOauthJwsKeysUrl                    []JwsKeysUrl
	ExternalOauthBlockedRolesList              *BlockedRolesListRequest
	ExternalOauthAllowedRolesList              *AllowedRolesListRequest
	ExternalOauthRsaPublicKey                  *string
	ExternalOauthRsaPublicKey2                 *string
	ExternalOauthAudienceList                  *AudienceListRequest
	ExternalOauthAnyRoleMode                   *ExternalOauthSecurityIntegrationAnyRoleModeOption
	ExternalOauthScopeDelimiter                *string
	Comment                                    *string
}

type ExternalOauthIntegrationUnsetRequest struct {
	Enabled                   *bool
	ExternalOauthAudienceList *bool
}

type AlterOauthForPartnerApplicationsSecurityIntegrationRequest struct {
	IfExists  *bool
	name      AccountObjectIdentifier // required
	SetTags   []TagAssociation
	UnsetTags []ObjectIdentifier
	Set       *OauthForPartnerApplicationsIntegrationSetRequest
	Unset     *OauthForPartnerApplicationsIntegrationUnsetRequest
}

type OauthForPartnerApplicationsIntegrationSetRequest struct {
	Enabled                   *bool
	OauthIssueRefreshTokens   *bool
	OauthRedirectUri          *string
	OauthRefreshTokenValidity *int
	OauthUseSecondaryRoles    *OauthSecurityIntegrationUseSecondaryRolesOption
	BlockedRolesList          *BlockedRolesListRequest
	Comment                   *string
}

type OauthForPartnerApplicationsIntegrationUnsetRequest struct {
	Enabled                *bool
	OauthUseSecondaryRoles *bool
}

type AlterOauthForCustomClientsSecurityIntegrationRequest struct {
	IfExists  *bool
	name      AccountObjectIdentifier // required
	SetTags   []TagAssociation
	UnsetTags []ObjectIdentifier
	Set       *OauthForCustomClientsIntegrationSetRequest
	Unset     *OauthForCustomClientsIntegrationUnsetRequest
}

type OauthForCustomClientsIntegrationSetRequest struct {
	Enabled                     *bool
	OauthRedirectUri            *string
	OauthAllowNonTlsRedirectUri *bool
	OauthEnforcePkce            *bool
	PreAuthorizedRolesList      *PreAuthorizedRolesListRequest
	BlockedRolesList            *BlockedRolesListRequest
	OauthIssueRefreshTokens     *bool
	OauthRefreshTokenValidity   *int
	OauthUseSecondaryRoles      *OauthSecurityIntegrationUseSecondaryRolesOption
	NetworkPolicy               *AccountObjectIdentifier
	OauthClientRsaPublicKey     *string
	OauthClientRsaPublicKey2    *string
	Comment                     *string
}

type OauthForCustomClientsIntegrationUnsetRequest struct {
	Enabled                  *bool
	NetworkPolicy            *bool
	OauthClientRsaPublicKey  *bool
	OauthClientRsaPublicKey2 *bool
	OauthUseSecondaryRoles   *bool
}

type AlterSaml2SecurityIntegrationRequest struct {
	IfExists                        *bool
	name                            AccountObjectIdentifier // required
	SetTags                         []TagAssociation
	UnsetTags                       []ObjectIdentifier
	Set                             *Saml2IntegrationSetRequest
	Unset                           *Saml2IntegrationUnsetRequest
	RefreshSaml2SnowflakePrivateKey *bool
}

type Saml2IntegrationSetRequest struct {
	Enabled                        *bool
	Saml2Issuer                    *string
	Saml2SsoUrl                    *string
	Saml2Provider                  *string
	Saml2X509Cert                  *string
	AllowedUserDomains             []UserDomain
	AllowedEmailPatterns           []EmailPattern
	Saml2SpInitiatedLoginPageLabel *string
	Saml2EnableSpInitiated         *bool
	Saml2SnowflakeX509Cert         *string
	Saml2SignRequest               *bool
	Saml2RequestedNameidFormat     *string
	Saml2PostLogoutRedirectUrl     *string
	Saml2ForceAuthn                *bool
	Saml2SnowflakeIssuerUrl        *string
	Saml2SnowflakeAcsUrl           *string
	Comment                        *string
}

type Saml2IntegrationUnsetRequest struct {
	Saml2ForceAuthn            *bool
	Saml2RequestedNameidFormat *bool
	Saml2PostLogoutRedirectUrl *bool
	Comment                    *bool
}

type AlterScimSecurityIntegrationRequest struct {
	IfExists  *bool
	name      AccountObjectIdentifier // required
	SetTags   []TagAssociation
	UnsetTags []ObjectIdentifier
	Set       *ScimIntegrationSetRequest
	Unset     *ScimIntegrationUnsetRequest
}

type ScimIntegrationSetRequest struct {
	Enabled       *bool
	NetworkPolicy *AccountObjectIdentifier
	SyncPassword  *bool
	Comment       *string
}

type ScimIntegrationUnsetRequest struct {
	Enabled       *bool
	NetworkPolicy *bool
	SyncPassword  *bool
	Comment       *bool
}

type DropSecurityIntegrationRequest struct {
	IfExists *bool
	name     AccountObjectIdentifier // required
}

type DescribeSecurityIntegrationRequest struct {
	name AccountObjectIdentifier // required
}

type ShowSecurityIntegrationRequest struct {
	Like *Like
}