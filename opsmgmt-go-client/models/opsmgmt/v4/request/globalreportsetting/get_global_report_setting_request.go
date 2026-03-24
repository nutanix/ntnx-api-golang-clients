package globalreportsetting

// This file holds the request struct for the GetGlobalReportSetting operation.

type GetGlobalReportSettingRequest struct {
	// (required) User ID of the user whose global report setting needs to be fetched/updated.
	UserExtId *string
}
