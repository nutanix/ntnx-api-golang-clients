package loadbalancersessionstats

import (
	import12 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/common/v1/stats"
	"time"
)

// This file holds the request struct for the GetLoadBalancerSessionStats operation.

type GetLoadBalancerSessionStatsRequest struct {
	// (required) The UUID of the load balancer subnet.
	ExtId *string

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
	StatType_ *import12.DownSamplingOperator

	// A URL query parameter that allows clients to request a specific set of properties for each entity or complex type.
	// Expression specified with the $select must conform to the [OData
	// V4.01](https://docs.oasis-open.org/odata/odata/v4.01/odata-v4.01-part1-protocol.html) URL conventions. If a $select
	// expression consists of a single select item that is an asterisk (i.e., *), then all properties on the matching resource
	// will be returned.
	Select_ *string
}
