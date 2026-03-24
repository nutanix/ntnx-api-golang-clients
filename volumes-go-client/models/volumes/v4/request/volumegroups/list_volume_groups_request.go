package volumegroups

// This file holds the request struct for the ListVolumeGroups operation.

type ListVolumeGroupsRequest struct {
	// A URL query parameter that specifies the page number of the result set. It must be a positive integer between 0 and the
	// maximum number of pages that are available for that resource. Any number out of this range might lead to no results.
	Page_ *int

	// A URL query parameter that specifies the total number of records returned in the result set.  Must be a positive integer
	// between 1 and 100. Any number out of this range will lead to a validation error. If the limit is not provided, a default
	// value of 50 records will be returned in the result set.
	Limit_ *int

	// A URL query parameter that allows clients to filter a collection of resources. The expression specified with $filter is
	// evaluated for each resource in the collection, and only items where the expression evaluates to true are included in the
	// response. Expression specified with the $filter must conform to the [OData
	// V4.01](https://docs.oasis-open.org/odata/odata/v4.01/odata-v4.01-part1-protocol.html) URL conventions.
	// For example, filter **$filter=name eq 'karbon-ntnx-1.0'** would filter the result on cluster name 'karbon-ntnx1.0',
	// filter **$filter=startswith(name, 'C')** would filter on cluster name starting with 'C'.
	Filter_ *string

	// A URL query parameter that allows clients to specify the sort criteria for the returned list of objects. Resources can
	// be sorted in ascending order using asc or descending order using desc. If asc or desc are not specified, the resources
	// will be sorted in ascending order by default. For example, '$orderby=templateName desc' would get all templates sorted
	// by templateName in descending order.
	Orderby_ *string

	// A URL query parameter that allows clients to request related resources when a resource that satisfies a particular
	// request is retrieved. Each expanded item is evaluated relative to the entity containing the property being expanded.
	// Other query options can be applied to an expanded property by appending a semicolon-separated list of query options,
	// enclosed in parentheses, to the property name. Permissible system query options are $filter, $select and $orderby.
	Expand_ *string

	// A URL query parameter that allows clients to request a specific set of properties for each entity or complex type.
	// Expression specified with the $select must conform to the [OData
	// V4.01](https://docs.oasis-open.org/odata/odata/v4.01/odata-v4.01-part1-protocol.html) URL conventions. If a $select
	// expression consists of a single select item that is an asterisk (i.e., *), then all properties on the matching resource
	// will be returned.
	Select_ *string
}
