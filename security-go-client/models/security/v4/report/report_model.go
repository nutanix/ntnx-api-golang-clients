/*
 * Generated file models/security/v4/report/report_model.go.
 *
 * Product version: 4.0.1
 *
 * Part of the Nutanix Security APIs
 *
 * (c) 2025 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module security.v4.report of Nutanix Security APIs
*/
package report

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	import2 "github.com/nutanix/ntnx-api-golang-clients/security-go-client/v4/models/common/v1/response"
	import3 "github.com/nutanix/ntnx-api-golang-clients/security-go-client/v4/models/security/v4/common"
	import1 "github.com/nutanix/ntnx-api-golang-clients/security-go-client/v4/models/security/v4/error"
	"time"
)

/*
Object that has the security summary for a single cluster
*/
type IssueSummary struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The number of current issues for a type
	*/
	CurrentIssueCount *int `json:"currentIssueCount,omitempty"`
	/*
	  Historical data containing how many issues were present and when
	*/
	Trend []TrendValue `json:"trend,omitempty"`
}

func (p *IssueSummary) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias IssueSummary

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *IssueSummary) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias IssueSummary
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = IssueSummary(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "currentIssueCount")
	delete(allFields, "trend")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewIssueSummary() *IssueSummary {
	p := new(IssueSummary)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "security.v4.report.IssueSummary"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /security/v4.0/report/stig-summaries Get operation
*/
type ListStigSummariesApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListStigSummariesApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ListStigSummariesApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ListStigSummariesApiResponse

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *ListStigSummariesApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ListStigSummariesApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = ListStigSummariesApiResponse(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewListStigSummariesApiResponse() *ListStigSummariesApiResponse {
	p := new(ListStigSummariesApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "security.v4.report.ListStigSummariesApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListStigSummariesApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListStigSummariesApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListStigSummariesApiResponseData()
	}
	e := p.Data.SetValue(v)
	if nil == e {
		if nil == p.DataItemDiscriminator_ {
			p.DataItemDiscriminator_ = new(string)
		}
		*p.DataItemDiscriminator_ = *p.Data.Discriminator
	}
	return e
}

/*
REST response for all response codes in API path /security/v4.0/report/stigs Get operation
*/
type ListStigsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListStigsApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ListStigsApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ListStigsApiResponse

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *ListStigsApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ListStigsApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = ListStigsApiResponse(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewListStigsApiResponse() *ListStigsApiResponse {
	p := new(ListStigsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "security.v4.report.ListStigsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListStigsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListStigsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListStigsApiResponseData()
	}
	e := p.Data.SetValue(v)
	if nil == e {
		if nil == p.DataItemDiscriminator_ {
			p.DataItemDiscriminator_ = new(string)
		}
		*p.DataItemDiscriminator_ = *p.Data.Discriminator
	}
	return e
}

/*
Contains possible values for the severity level of a vulnerability.
*/
type Severity int

const (
	SEVERITY_UNKNOWN  Severity = 0
	SEVERITY_REDACTED Severity = 1
	SEVERITY_CRITICAL Severity = 2
	SEVERITY_HIGH     Severity = 3
	SEVERITY_MEDIUM   Severity = 4
	SEVERITY_LOW      Severity = 5
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *Severity) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CRITICAL",
		"HIGH",
		"MEDIUM",
		"LOW",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e Severity) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CRITICAL",
		"HIGH",
		"MEDIUM",
		"LOW",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *Severity) index(name string) Severity {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CRITICAL",
		"HIGH",
		"MEDIUM",
		"LOW",
	}
	for idx := range names {
		if names[idx] == name {
			return Severity(idx)
		}
	}
	return SEVERITY_UNKNOWN
}

func (e *Severity) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for Severity:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *Severity) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e Severity) Ref() *Severity {
	return &e
}

/*
Type of nutanix software, such as PC, PE, AFS, etc.
*/
type SoftwareType int

const (
	SOFTWARETYPE_UNKNOWN  SoftwareType = 0
	SOFTWARETYPE_REDACTED SoftwareType = 1
	SOFTWARETYPE_PC       SoftwareType = 2
	SOFTWARETYPE_AOS      SoftwareType = 3
	SOFTWARETYPE_AHV      SoftwareType = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *SoftwareType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PC",
		"AOS",
		"AHV",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e SoftwareType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PC",
		"AOS",
		"AHV",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *SoftwareType) index(name string) SoftwareType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PC",
		"AOS",
		"AHV",
	}
	for idx := range names {
		if names[idx] == name {
			return SoftwareType(idx)
		}
	}
	return SOFTWARETYPE_UNKNOWN
}

func (e *SoftwareType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for SoftwareType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *SoftwareType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e SoftwareType) Ref() *SoftwareType {
	return &e
}

/*
Contains information of a specific STIG.
*/
type Stig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of clusters that failed the STIG control.
	*/
	AffectedClusters []string `json:"affectedClusters,omitempty"`
	/*
	  Benchmark ID of the STIG rules.
	*/
	BenchmarkId *string `json:"benchmarkId,omitempty"`
	/*
	  The comments to explain why a STIG rule applies or does not apply to the cluster.
	*/
	Comments *string `json:"comments,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  The command/steps to fix the STIG rule failure.
	*/
	FixText *string `json:"fixText,omitempty"`
	/*
	  Additional identifiers used to describe this control.
	*/
	Identifiers []string `json:"identifiers,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Rule ID of the STIG control.
	*/
	RuleId *string `json:"ruleId,omitempty"`

	Severity *Severity `json:"severity,omitempty"`

	Status *StigStatus `json:"status,omitempty"`
	/*
	  STIG ID of the control.
	*/
	StigVersion *string `json:"stigVersion,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  Title of the STIG control.
	*/
	Title *string `json:"title,omitempty"`
}

func (p *Stig) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias Stig

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *Stig) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Stig
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = Stig(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "affectedClusters")
	delete(allFields, "benchmarkId")
	delete(allFields, "comments")
	delete(allFields, "extId")
	delete(allFields, "fixText")
	delete(allFields, "identifiers")
	delete(allFields, "links")
	delete(allFields, "ruleId")
	delete(allFields, "severity")
	delete(allFields, "status")
	delete(allFields, "stigVersion")
	delete(allFields, "tenantId")
	delete(allFields, "title")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewStig() *Stig {
	p := new(Stig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "security.v4.report.Stig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type StigProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of clusters that failed the STIG control.
	*/
	AffectedClusters []string `json:"affectedClusters,omitempty"`
	/*
	  Benchmark ID of the STIG rules.
	*/
	BenchmarkId *string `json:"benchmarkId,omitempty"`
	/*
	  The comments to explain why a STIG rule applies or does not apply to the cluster.
	*/
	Comments *string `json:"comments,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  The command/steps to fix the STIG rule failure.
	*/
	FixText *string `json:"fixText,omitempty"`
	/*
	  Additional identifiers used to describe this control.
	*/
	Identifiers []string `json:"identifiers,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Rule ID of the STIG control.
	*/
	RuleId *string `json:"ruleId,omitempty"`

	Severity *Severity `json:"severity,omitempty"`

	Status *StigStatus `json:"status,omitempty"`
	/*
	  STIG ID of the control.
	*/
	StigVersion *string `json:"stigVersion,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  Title of the STIG control.
	*/
	Title *string `json:"title,omitempty"`
}

func (p *StigProjection) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias StigProjection

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *StigProjection) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias StigProjection
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = StigProjection(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "affectedClusters")
	delete(allFields, "benchmarkId")
	delete(allFields, "comments")
	delete(allFields, "extId")
	delete(allFields, "fixText")
	delete(allFields, "identifiers")
	delete(allFields, "links")
	delete(allFields, "ruleId")
	delete(allFields, "severity")
	delete(allFields, "status")
	delete(allFields, "stigVersion")
	delete(allFields, "tenantId")
	delete(allFields, "title")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewStigProjection() *StigProjection {
	p := new(StigProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "security.v4.report.StigProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Indicates if any clusters have failed a STIG control.
*/
type StigStatus int

const (
	STIGSTATUS_UNKNOWN        StigStatus = 0
	STIGSTATUS_REDACTED       StigStatus = 1
	STIGSTATUS_APPLICABLE     StigStatus = 2
	STIGSTATUS_NOT_APPLICABLE StigStatus = 3
	STIGSTATUS_NEEDS_REVIEW   StigStatus = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *StigStatus) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"APPLICABLE",
		"NOT_APPLICABLE",
		"NEEDS_REVIEW",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e StigStatus) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"APPLICABLE",
		"NOT_APPLICABLE",
		"NEEDS_REVIEW",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *StigStatus) index(name string) StigStatus {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"APPLICABLE",
		"NOT_APPLICABLE",
		"NEEDS_REVIEW",
	}
	for idx := range names {
		if names[idx] == name {
			return StigStatus(idx)
		}
	}
	return STIGSTATUS_UNKNOWN
}

func (e *StigStatus) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for StigStatus:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *StigStatus) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e StigStatus) Ref() *StigStatus {
	return &e
}

/*
Cluster STIG summary.
*/
type StigSummary struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  UUID of the cluster.
	*/
	ClusterExtId *string `json:"clusterExtId,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Count of resources with failed STIG tests.
	*/
	FailedCount *int `json:"failedCount,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Count of resources with STIG test not applicable.
	*/
	NotApplicableCount *int `json:"notApplicableCount,omitempty"`
	/*
	  Count of total resources in the cluster.
	*/
	PassedCount *int `json:"passedCount,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *StigSummary) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias StigSummary

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *StigSummary) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias StigSummary
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = StigSummary(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "clusterExtId")
	delete(allFields, "extId")
	delete(allFields, "failedCount")
	delete(allFields, "links")
	delete(allFields, "notApplicableCount")
	delete(allFields, "passedCount")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewStigSummary() *StigSummary {
	p := new(StigSummary)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "security.v4.report.StigSummary"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type StigSummaryProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  UUID of the cluster.
	*/
	ClusterExtId *string `json:"clusterExtId,omitempty"`

	ClusterProjection *import3.ClusterProjection `json:"clusterProjection,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Count of resources with failed STIG tests.
	*/
	FailedCount *int `json:"failedCount,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Count of resources with STIG test not applicable.
	*/
	NotApplicableCount *int `json:"notApplicableCount,omitempty"`
	/*
	  Count of total resources in the cluster.
	*/
	PassedCount *int `json:"passedCount,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *StigSummaryProjection) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias StigSummaryProjection

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *StigSummaryProjection) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias StigSummaryProjection
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = StigSummaryProjection(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "clusterExtId")
	delete(allFields, "clusterProjection")
	delete(allFields, "extId")
	delete(allFields, "failedCount")
	delete(allFields, "links")
	delete(allFields, "notApplicableCount")
	delete(allFields, "passedCount")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewStigSummaryProjection() *StigSummaryProjection {
	p := new(StigSummaryProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "security.v4.report.StigSummaryProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Number of security related issues for a single cluster
*/
type Summary struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  UUID of the cluster.
	*/
	ClusterExtId *string `json:"clusterExtId,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  The date the dashboard information was last refreshed
	*/
	LastRefreshTime *time.Time `json:"lastRefreshTime,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`

	PasswordSummary *IssueSummary `json:"passwordSummary,omitempty"`

	SecurityConfigSummary *IssueSummary `json:"securityConfigSummary,omitempty"`

	StigSummary *IssueSummary `json:"stigSummary,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	TrendType *import3.Timescale `json:"trendType,omitempty"`

	VulnerabilitiesSummary *IssueSummary `json:"vulnerabilitiesSummary,omitempty"`
}

func (p *Summary) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias Summary

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *Summary) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Summary
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = Summary(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "clusterExtId")
	delete(allFields, "extId")
	delete(allFields, "lastRefreshTime")
	delete(allFields, "links")
	delete(allFields, "passwordSummary")
	delete(allFields, "securityConfigSummary")
	delete(allFields, "stigSummary")
	delete(allFields, "tenantId")
	delete(allFields, "trendType")
	delete(allFields, "vulnerabilitiesSummary")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewSummary() *Summary {
	p := new(Summary)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "security.v4.report.Summary"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type SummaryProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  UUID of the cluster.
	*/
	ClusterExtId *string `json:"clusterExtId,omitempty"`

	ClusterProjection *import3.ClusterProjection `json:"clusterProjection,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  The date the dashboard information was last refreshed
	*/
	LastRefreshTime *time.Time `json:"lastRefreshTime,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`

	PasswordSummary *IssueSummary `json:"passwordSummary,omitempty"`

	SecurityConfigSummary *IssueSummary `json:"securityConfigSummary,omitempty"`

	StigSummary *IssueSummary `json:"stigSummary,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	TrendType *import3.Timescale `json:"trendType,omitempty"`

	VulnerabilitiesSummary *IssueSummary `json:"vulnerabilitiesSummary,omitempty"`
}

func (p *SummaryProjection) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias SummaryProjection

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *SummaryProjection) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias SummaryProjection
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = SummaryProjection(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "clusterExtId")
	delete(allFields, "clusterProjection")
	delete(allFields, "extId")
	delete(allFields, "lastRefreshTime")
	delete(allFields, "links")
	delete(allFields, "passwordSummary")
	delete(allFields, "securityConfigSummary")
	delete(allFields, "stigSummary")
	delete(allFields, "tenantId")
	delete(allFields, "trendType")
	delete(allFields, "vulnerabilitiesSummary")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewSummaryProjection() *SummaryProjection {
	p := new(SummaryProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "security.v4.report.SummaryProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Historical data containing how many issues were present and when
*/
type TrendValue struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The number of previous issues at that date
	*/
	Trend *int `json:"trend,omitempty"`
	/*
	  Date value for trends
	*/
	TrendDate *time.Time `json:"trendDate,omitempty"`
}

func (p *TrendValue) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias TrendValue

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *TrendValue) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias TrendValue
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = TrendValue(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "trend")
	delete(allFields, "trendDate")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewTrendValue() *TrendValue {
	p := new(TrendValue)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "security.v4.report.TrendValue"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Contains all relevant information for a vulnerability
*/
type Vulnerability struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The CESA id
	*/
	CesaId *string `json:"cesaId,omitempty"`
	/*
	  The CVE ids associated with this CESA
	*/
	CveIds []string `json:"cveIds,omitempty"`
	/*
	  The short description of the vulnerability
	*/
	Description *string `json:"description,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  The AOS versions where the vulnerability was fixed
	*/
	FixVersions []string `json:"fixVersions,omitempty"`
	/*
	  If the vulnerability is critical or not
	*/
	IsCritical *bool `json:"isCritical,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Date of creation of NXVD data used to calculate vulnerabilities
	*/
	NxvdCreatedTime *time.Time `json:"nxvdCreatedTime,omitempty"`
	/*
	  Version of NXVD data used to calculate vulnerabilities
	*/
	NxvdVersion *string `json:"nxvdVersion,omitempty"`

	Severity *Severity `json:"severity,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *Vulnerability) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias Vulnerability

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *Vulnerability) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Vulnerability
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = Vulnerability(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "cesaId")
	delete(allFields, "cveIds")
	delete(allFields, "description")
	delete(allFields, "extId")
	delete(allFields, "fixVersions")
	delete(allFields, "isCritical")
	delete(allFields, "links")
	delete(allFields, "nxvdCreatedTime")
	delete(allFields, "nxvdVersion")
	delete(allFields, "severity")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewVulnerability() *Vulnerability {
	p := new(Vulnerability)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "security.v4.report.Vulnerability"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type VulnerabilityProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The CESA id
	*/
	CesaId *string `json:"cesaId,omitempty"`
	/*
	  The CVE ids associated with this CESA
	*/
	CveIds []string `json:"cveIds,omitempty"`
	/*
	  The short description of the vulnerability
	*/
	Description *string `json:"description,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  The AOS versions where the vulnerability was fixed
	*/
	FixVersions []string `json:"fixVersions,omitempty"`
	/*
	  If the vulnerability is critical or not
	*/
	IsCritical *bool `json:"isCritical,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Date of creation of NXVD data used to calculate vulnerabilities
	*/
	NxvdCreatedTime *time.Time `json:"nxvdCreatedTime,omitempty"`
	/*
	  Version of NXVD data used to calculate vulnerabilities
	*/
	NxvdVersion *string `json:"nxvdVersion,omitempty"`

	Severity *Severity `json:"severity,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *VulnerabilityProjection) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias VulnerabilityProjection

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *VulnerabilityProjection) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias VulnerabilityProjection
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = VulnerabilityProjection(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "cesaId")
	delete(allFields, "cveIds")
	delete(allFields, "description")
	delete(allFields, "extId")
	delete(allFields, "fixVersions")
	delete(allFields, "isCritical")
	delete(allFields, "links")
	delete(allFields, "nxvdCreatedTime")
	delete(allFields, "nxvdVersion")
	delete(allFields, "severity")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewVulnerabilityProjection() *VulnerabilityProjection {
	p := new(VulnerabilityProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "security.v4.report.VulnerabilityProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Cluster vulnerability report
*/
type VulnerabilitySummary struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  UUID of the cluster.
	*/
	ClusterExtId *string `json:"clusterExtId,omitempty"`
	/*
	  The CVE ids associated with this CESA
	*/
	CveIds []string `json:"cveIds,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  AOS version on that cluster
	*/
	InstalledVersion *string `json:"installedVersion,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`

	SoftwareType *SoftwareType `json:"softwareType,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *VulnerabilitySummary) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias VulnerabilitySummary

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *VulnerabilitySummary) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias VulnerabilitySummary
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = VulnerabilitySummary(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "clusterExtId")
	delete(allFields, "cveIds")
	delete(allFields, "extId")
	delete(allFields, "installedVersion")
	delete(allFields, "links")
	delete(allFields, "softwareType")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewVulnerabilitySummary() *VulnerabilitySummary {
	p := new(VulnerabilitySummary)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "security.v4.report.VulnerabilitySummary"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type VulnerabilitySummaryProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  UUID of the cluster.
	*/
	ClusterExtId *string `json:"clusterExtId,omitempty"`
	/*
	  The CVE ids associated with this CESA
	*/
	CveIds []string `json:"cveIds,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  AOS version on that cluster
	*/
	InstalledVersion *string `json:"installedVersion,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`

	SoftwareType *SoftwareType `json:"softwareType,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *VulnerabilitySummaryProjection) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias VulnerabilitySummaryProjection

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *VulnerabilitySummaryProjection) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias VulnerabilitySummaryProjection
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = VulnerabilitySummaryProjection(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "clusterExtId")
	delete(allFields, "cveIds")
	delete(allFields, "extId")
	delete(allFields, "installedVersion")
	delete(allFields, "links")
	delete(allFields, "softwareType")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewVulnerabilitySummaryProjection() *VulnerabilitySummaryProjection {
	p := new(VulnerabilitySummaryProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "security.v4.report.VulnerabilitySummaryProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type OneOfListStigSummariesApiResponseData struct {
	Discriminator *string                 `json:"-"`
	ObjectType_   *string                 `json:"-"`
	oneOfType400  *import1.ErrorResponse  `json:"-"`
	oneOfType2006 []StigSummary           `json:"-"`
	oneOfType401  []StigSummaryProjection `json:"-"`
}

func NewOneOfListStigSummariesApiResponseData() *OneOfListStigSummariesApiResponseData {
	p := new(OneOfListStigSummariesApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListStigSummariesApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListStigSummariesApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import1.ErrorResponse)
		}
		*p.oneOfType400 = v.(import1.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case []StigSummary:
		p.oneOfType2006 = v.([]StigSummary)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<security.v4.report.StigSummary>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<security.v4.report.StigSummary>"
	case []StigSummaryProjection:
		p.oneOfType401 = v.([]StigSummaryProjection)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<security.v4.report.StigSummaryProjection>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<security.v4.report.StigSummaryProjection>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListStigSummariesApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<security.v4.report.StigSummary>" == *p.Discriminator {
		return p.oneOfType2006
	}
	if "List<security.v4.report.StigSummaryProjection>" == *p.Discriminator {
		return p.oneOfType401
	}
	return nil
}

func (p *OneOfListStigSummariesApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "security.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import1.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType400.ObjectType_
			return nil
		}
	}
	vOneOfType2006 := new([]StigSummary)
	if err := json.Unmarshal(b, vOneOfType2006); err == nil {
		if len(*vOneOfType2006) == 0 || "security.v4.report.StigSummary" == *((*vOneOfType2006)[0].ObjectType_) {
			p.oneOfType2006 = *vOneOfType2006
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<security.v4.report.StigSummary>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<security.v4.report.StigSummary>"
			return nil
		}
	}
	vOneOfType401 := new([]StigSummaryProjection)
	if err := json.Unmarshal(b, vOneOfType401); err == nil {
		if len(*vOneOfType401) == 0 || "security.v4.report.StigSummaryProjection" == *((*vOneOfType401)[0].ObjectType_) {
			p.oneOfType401 = *vOneOfType401
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<security.v4.report.StigSummaryProjection>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<security.v4.report.StigSummaryProjection>"
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListStigSummariesApiResponseData"))
}

func (p *OneOfListStigSummariesApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<security.v4.report.StigSummary>" == *p.Discriminator {
		return json.Marshal(p.oneOfType2006)
	}
	if "List<security.v4.report.StigSummaryProjection>" == *p.Discriminator {
		return json.Marshal(p.oneOfType401)
	}
	return nil, errors.New("No value to marshal for OneOfListStigSummariesApiResponseData")
}

type OneOfListStigsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	oneOfType2006 []Stig                 `json:"-"`
	oneOfType401  []StigProjection       `json:"-"`
}

func NewOneOfListStigsApiResponseData() *OneOfListStigsApiResponseData {
	p := new(OneOfListStigsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListStigsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListStigsApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import1.ErrorResponse)
		}
		*p.oneOfType400 = v.(import1.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case []Stig:
		p.oneOfType2006 = v.([]Stig)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<security.v4.report.Stig>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<security.v4.report.Stig>"
	case []StigProjection:
		p.oneOfType401 = v.([]StigProjection)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<security.v4.report.StigProjection>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<security.v4.report.StigProjection>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListStigsApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<security.v4.report.Stig>" == *p.Discriminator {
		return p.oneOfType2006
	}
	if "List<security.v4.report.StigProjection>" == *p.Discriminator {
		return p.oneOfType401
	}
	return nil
}

func (p *OneOfListStigsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "security.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import1.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType400.ObjectType_
			return nil
		}
	}
	vOneOfType2006 := new([]Stig)
	if err := json.Unmarshal(b, vOneOfType2006); err == nil {
		if len(*vOneOfType2006) == 0 || "security.v4.report.Stig" == *((*vOneOfType2006)[0].ObjectType_) {
			p.oneOfType2006 = *vOneOfType2006
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<security.v4.report.Stig>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<security.v4.report.Stig>"
			return nil
		}
	}
	vOneOfType401 := new([]StigProjection)
	if err := json.Unmarshal(b, vOneOfType401); err == nil {
		if len(*vOneOfType401) == 0 || "security.v4.report.StigProjection" == *((*vOneOfType401)[0].ObjectType_) {
			p.oneOfType401 = *vOneOfType401
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<security.v4.report.StigProjection>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<security.v4.report.StigProjection>"
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListStigsApiResponseData"))
}

func (p *OneOfListStigsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<security.v4.report.Stig>" == *p.Discriminator {
		return json.Marshal(p.oneOfType2006)
	}
	if "List<security.v4.report.StigProjection>" == *p.Discriminator {
		return json.Marshal(p.oneOfType401)
	}
	return nil, errors.New("No value to marshal for OneOfListStigsApiResponseData")
}

type FileDetail struct {
	Path        *string `json:"-"`
	ObjectType_ *string `json:"-"`
}

func NewFileDetail() *FileDetail {
	p := new(FileDetail)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "FileDetail"

	return p
}
