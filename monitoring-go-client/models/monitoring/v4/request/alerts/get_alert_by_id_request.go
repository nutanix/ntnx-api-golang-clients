package alerts

// This file holds the request struct for the GetAlertById operation.

type GetAlertByIdRequest struct {
	// (required) UUID of the generated alert.
	ExtId *string

	// A URL query parameter that allows clients to request related resources when a resource that satisfies a particular
	// request is retrieved. Each expanded item is evaluated relative to the entity containing the property being expanded.
	// Other query options can be applied to an expanded property by appending a semicolon-separated list of query options,
	// enclosed in parentheses, to the property name. Permissible system query options are $filter, $select and $orderby. RBAC
	// access is required on the underlying entity that an expand value resolves to; if the caller lacks permission, the
	// corresponding fields will be omitted from the response. The supported expand values are listed below: - `project`,
	// `shared_with_projects`
	Expand_ *string
}
