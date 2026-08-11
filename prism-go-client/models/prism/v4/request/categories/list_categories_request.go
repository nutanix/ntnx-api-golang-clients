package categories

// This file holds the request struct for the ListCategories operation.

type ListCategoriesRequest struct {
	// A URL query parameter that specifies the page number of the result set. It must be a positive integer between 0 and the
	// maximum number of pages that are available for that resource. Any number out of this range might lead to no results. If
	// both $page and $apply query parameters are present, $page will be applied on entities within the groups.
	Page_ *int

	// A URL query parameter that specifies the total number of records returned in the result set.  Must be a positive integer
	// between 1 and 100. Any number out of this range will lead to a validation error. If the limit is not provided, a default
	// value of 50 records will be returned in the result set. If both $limit and $apply query parameters are present, $limit
	// will be applied on entities within the groups. Default value of limit with $apply will be 5.
	Limit_ *int

	// A URL query parameter that allows clients to filter a collection of resources. The expression specified with $filter is
	// evaluated for each resource in the collection, and only items where the expression evaluates to true are included in the
	// response. Expression specified with the $filter must conform to the [OData
	// V4.01](https://docs.oasis-open.org/odata/odata/v4.01/odata-v4.01-part1-protocol.html) URL conventions.
	// For example, filter **$filter=name eq 'karbon-ntnx-1.0'** would filter the result on cluster name 'karbon-ntnx1.0',
	// filter **$filter=startswith(name, 'C')** would filter on cluster name starting with 'C'. If both $filter and $apply
	// query parameters are present, $filter will be applied on entities within the groups.
	Filter_ *string

	// A URL query parameter that allows clients to specify the sort criteria for the returned list of objects. Resources can
	// be sorted in ascending order using asc or descending order using desc. If asc or desc are not specified, the resources
	// will be sorted in ascending order by default. For example, '$orderby=templateName desc' would get all templates sorted
	// by templateName in descending order. If both $orderby and $apply query parameters are present, $orderby will be applied
	// on entities within the groups.
	Orderby_ *string

	// A URL query parameter that allows clients to specify a sequence of transformations to the entity set, such as groupby,
	// orderby, aggregate etc. For example '$apply=groupby((property))' will get the entity set grouped by "property". The
	// sequence must start with groupby, which may include multiple nested aggregates. It may optionally be followed by orderby
	// and/or pagination parameters. For example '$apply=groupby((property),aggregate(property with count as
	// totalCount))/orderby(property)/page(0)/limit(2)' will group the entity set by "property", aggregate the number of
	// entities in each group, sort the groups by "property" and return the first two groups. The term "property" is a
	// placeholder in these examples and should be substituted with the actual property name of the resource.
	Apply_ *string

	// A URL query parameter that allows clients to request related resources when a resource that satisfies a particular
	// request is retrieved. Each expanded item is evaluated relative to the entity containing the property being expanded.
	// Other query options can be applied to an expanded property by appending a semicolon-separated list of query options,
	// enclosed in parentheses, to the property name. Permissible system query options are $filter, $select and $orderby. If
	// both $expand and $apply query parameters are present, $expand will be applied on entities within the groups.
	Expand_ *string

	// A URL query parameter that allows clients to request a specific set of properties for each entity or complex type.
	// Expression specified with the $select must conform to the [OData
	// V4.01](https://docs.oasis-open.org/odata/odata/v4.01/odata-v4.01-part1-protocol.html) URL conventions. If a $select
	// expression consists of a single select item that is an asterisk (i.e., *), then all properties on the matching resource
	// will be returned. If both $select and $apply query parameters are present, $select will be applied on entities within
	// the groups.
	Select_ *string
}
