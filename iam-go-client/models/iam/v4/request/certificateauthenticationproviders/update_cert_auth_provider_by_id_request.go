package certificateauthenticationproviders

import (
	import3 "github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/models/common/v1/response"
	import4 "github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/models/iam/v4/authn"
	"time"
)

// This file holds the request struct for the UpdateCertAuthProviderById operation.

type UpdateCertAuthProviderByIdRequest struct {
	// (required) UUID V5 created for the certificate-based authentication provider.
	ExtId *string

	// (required)
	ClientCaChain *string

	// (required)
	CaCertFileName *string

	// (required)
	IsCertAuthEnabled *bool

	// (required)
	Name *string

	// (required)
	IsCacEnabled *bool

	//
	DirSvcExtID *string

	//
	CertRevocationInfo *import4.CertRevocationInfo

	//
	CreatedBy *string

	//
	TenantId *string

	//
	CreatedTime *time.Time

	//
	Links *[]import3.ApiLink

	//
	LastUpdatedTime *time.Time

	// UUID V5 created for the certificate-based authentication provider.
	ExtId2 *string
}
