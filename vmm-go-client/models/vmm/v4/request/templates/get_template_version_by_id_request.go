package templates

// This file holds the request struct for the GetTemplateVersionById operation.

type GetTemplateVersionByIdRequest struct {
	// (required) The identifier of a template.
	TemplateExtId *string

	// (required) The identifier of a version.
	ExtId *string
}
