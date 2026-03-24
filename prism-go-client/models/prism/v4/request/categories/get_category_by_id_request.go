package categories

// This file holds the request struct for the GetCategoryById operation.

type GetCategoryByIdRequest struct {
	// (required) A globally unique identifier of an instance that is suitable for external consumption.
	ExtId *string

	// A URL query parameter that allows clients to request related resources when a resource that satisfies a particular
	// request is retrieved. Each expanded item is evaluated relative to the entity containing the property being expanded.
	// Other query options can be applied to an expanded property by appending a semicolon-separated list of query options,
	// enclosed in parentheses, to the property name. Permissible system query options are $filter, $select and $orderby.
	Expand_ *string
}
