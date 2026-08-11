package licenses

// This file holds the request struct for the GetSettingById operation.

type GetSettingByIdRequest struct {
	// (required) External identifier of the cluster or Prism Central whose setting is requested.
	ExtId *string
}
