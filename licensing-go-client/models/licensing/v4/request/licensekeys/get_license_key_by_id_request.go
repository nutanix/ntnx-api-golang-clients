package licensekeys

// This file holds the request struct for the GetLicenseKeyById operation.

type GetLicenseKeyByIdRequest struct {
	// (required) Path parameter for the license key ID.
	ExtId *string
}
