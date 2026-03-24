package esxistats

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/common/v1/stats"
	"time"
)

// This file holds the request struct for the ListVmStats operation.

type ListVmStatsRequest struct {
	// (required) The start time of the period for which stats should be reported. The value should be in extended ISO-8601 format. For
	// example, start time of 2022-04-23T01:23:45.678+09:00 would consider all stats starting at 1:23:45.678 on the 23rd of
	// April 2022. Details around ISO-8601 format can be found at https://www.iso.org/standard/70907.html
	StartTime_ *time.Time

	// (required) The end time of the period for which stats should be reported. The value should be in extended ISO-8601 format. For
	// example, end time of 2022-04-23T013:23:45.678+09:00 would consider all stats till 13:23:45 .678 on the 23rd of April
	// 2022. Details around ISO-8601 format can be found at https://www.iso.org/standard/70907.html
	EndTime_ *time.Time

	// The sampling interval in seconds at which statistical data should be collected. For example, if you want performance
	// statistics every 30 seconds, then provide the value as 30.
	SamplingInterval_ *int

	// The operator to use while performing down-sampling on stats data. Allowed values are SUM, MIN, MAX, AVG, COUNT and LAST.
	StatType_ *import1.DownSamplingOperator

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

	// A URL query parameter that allows clients to request a specific set of properties for each entity or complex type.
	// Expression specified with the $select must conform to the [OData
	// V4.01](https://docs.oasis-open.org/odata/odata/v4.01/odata-v4.01-part1-protocol.html) URL conventions. If a $select
	// expression consists of a single select item that is an asterisk (i.e., *), then all properties on the matching resource
	// will be returned.
	Select_ *string
}
