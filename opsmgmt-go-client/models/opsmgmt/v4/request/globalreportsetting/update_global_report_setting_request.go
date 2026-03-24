package globalreportsetting

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/opsmgmt-go-client/v4/models/opsmgmt/v4/config"
)

// This file holds the request struct for the UpdateGlobalReportSetting operation.

type UpdateGlobalReportSettingRequest struct {
	// (required) User ID of the user whose global report setting needs to be fetched/updated.
	UserExtId *string

	// (required) The request body updates the report setting. Name is the only required field.
	Body *import1.GlobalReportSetting
}
