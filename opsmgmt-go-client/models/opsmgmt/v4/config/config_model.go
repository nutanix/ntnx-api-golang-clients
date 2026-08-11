/*
 * Generated file models/opsmgmt/v4/config/config_model.go.
 *
 * Product version: 4.1.1-beta-1
 *
 * Part of the Nutanix Cloud Management Platform APIs
 *
 * (c) 2026 Nutanix Inc.  All rights reserved
 *
 */

/*
  Common functions for aiops, devops, secops, finops.
*/
package config

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	import4 "github.com/nutanix/ntnx-api-golang-clients/opsmgmt-go-client/v4/models/common/v1/config"
	import2 "github.com/nutanix/ntnx-api-golang-clients/opsmgmt-go-client/v4/models/common/v1/response"
	import1 "github.com/nutanix/ntnx-api-golang-clients/opsmgmt-go-client/v4/models/opsmgmt/v4/error"
	import3 "github.com/nutanix/ntnx-api-golang-clients/opsmgmt-go-client/v4/models/prism/v4/config"
	"time"
)

/*
Used to aggregate field data from multiple values across time range.
*/
type AggregateFunction int

const (
	AGGREGATEFUNCTION_UNKNOWN  AggregateFunction = 0
	AGGREGATEFUNCTION_REDACTED AggregateFunction = 1
	AGGREGATEFUNCTION_SUM      AggregateFunction = 2
	AGGREGATEFUNCTION_MAX      AggregateFunction = 3
	AGGREGATEFUNCTION_MIN      AggregateFunction = 4
	AGGREGATEFUNCTION_AVG      AggregateFunction = 5
	AGGREGATEFUNCTION_LAST     AggregateFunction = 6
	AGGREGATEFUNCTION_COUNT    AggregateFunction = 7
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *AggregateFunction) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SUM",
		"MAX",
		"MIN",
		"AVG",
		"LAST",
		"COUNT",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e AggregateFunction) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SUM",
		"MAX",
		"MIN",
		"AVG",
		"LAST",
		"COUNT",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *AggregateFunction) index(name string) AggregateFunction {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SUM",
		"MAX",
		"MIN",
		"AVG",
		"LAST",
		"COUNT",
	}
	for idx := range names {
		if names[idx] == name {
			return AggregateFunction(idx)
		}
	}
	return AGGREGATEFUNCTION_UNKNOWN
}

func (e *AggregateFunction) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for AggregateFunction:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *AggregateFunction) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e AggregateFunction) Ref() *AggregateFunction {
	return &e
}

/*
Bounds of the widget in the dashboard.
*/
type Bounds struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Height of the widget/layout in the dashboard.
	*/
	Height *int64 `json:"height"`
	/*
	  Width of the widget/layout in the dashboard.
	*/
	Width *int64 `json:"width"`
	/*
	  X position of the widget/layout in the dashboard.
	*/
	X *int64 `json:"x"`
	/*
	  Y position of the widget/layout in the dashboard.
	*/
	Y *int64 `json:"y"`
}

func (p *Bounds) MarshalJSON() ([]byte, error) {
	type BoundsProxy Bounds

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*BoundsProxy
		Height *int64 `json:"height,omitempty"`
		Width  *int64 `json:"width,omitempty"`
		X      *int64 `json:"x,omitempty"`
		Y      *int64 `json:"y,omitempty"`
	}{
		BoundsProxy: (*BoundsProxy)(p),
		Height:      p.Height,
		Width:       p.Width,
		X:           p.X,
		Y:           p.Y,
	}

	known, err := json.Marshal(baseStruct)
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

func (p *Bounds) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Bounds
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewBounds()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Height != nil {
		p.Height = known.Height
	}
	if known.Width != nil {
		p.Width = known.Width
	}
	if known.X != nil {
		p.X = known.X
	}
	if known.Y != nil {
		p.Y = known.Y
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "height")
	delete(allFields, "width")
	delete(allFields, "x")
	delete(allFields, "y")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewBounds() *Bounds {
	p := new(Bounds)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.config.Bounds"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Defines a compound metric calculated from up to two base metrics using a formula.
*/
type CompoundMetric struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Mathematical formula combining operands (A, and optionally B) using +, -, *, / operators. Numeric constants are also supported. Examples: 'A/B', 'A*100'
	*/
	Formula *string `json:"formula,omitempty"`
	/*
	  List of compound metric operands used in the compound metric formula.
	*/
	Operands []CompoundMetricOperand `json:"operands,omitempty"`
}

func (p *CompoundMetric) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias CompoundMetric

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

func (p *CompoundMetric) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias CompoundMetric
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewCompoundMetric()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Formula != nil {
		p.Formula = known.Formula
	}
	if known.Operands != nil {
		p.Operands = known.Operands
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "formula")
	delete(allFields, "operands")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewCompoundMetric() *CompoundMetric {
	p := new(CompoundMetric)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.config.CompoundMetric"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Input metric referenced by a letter (A or B) in the compound metric formula.
*/
type CompoundMetricOperand struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	AggregateFunction *AggregateFunction `json:"aggregateFunction,omitempty"`
	/*
	  Human-readable label for widget field.
	*/
	Label *string `json:"label"`
	/*
	  Entity attribute/metric to be selected for the widget.
	*/
	Name *string `json:"name"`

	Unit *Unit `json:"unit,omitempty"`
}

func (p *CompoundMetricOperand) MarshalJSON() ([]byte, error) {
	type CompoundMetricOperandProxy CompoundMetricOperand

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*CompoundMetricOperandProxy
		Label *string `json:"label,omitempty"`
		Name  *string `json:"name,omitempty"`
	}{
		CompoundMetricOperandProxy: (*CompoundMetricOperandProxy)(p),
		Label:                      p.Label,
		Name:                       p.Name,
	}

	known, err := json.Marshal(baseStruct)
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

func (p *CompoundMetricOperand) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias CompoundMetricOperand
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewCompoundMetricOperand()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.AggregateFunction != nil {
		p.AggregateFunction = known.AggregateFunction
	}
	if known.Label != nil {
		p.Label = known.Label
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.Unit != nil {
		p.Unit = known.Unit
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "aggregateFunction")
	delete(allFields, "label")
	delete(allFields, "name")
	delete(allFields, "unit")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewCompoundMetricOperand() *CompoundMetricOperand {
	p := new(CompoundMetricOperand)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.config.CompoundMetricOperand"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /opsmgmt/v4.1.b1/config/dashboards Post operation
*/
type CreateDashboardApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateDashboardApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *CreateDashboardApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias CreateDashboardApiResponse

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

func (p *CreateDashboardApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias CreateDashboardApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewCreateDashboardApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewCreateDashboardApiResponse() *CreateDashboardApiResponse {
	p := new(CreateDashboardApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.config.CreateDashboardApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreateDashboardApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreateDashboardApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreateDashboardApiResponseData()
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
REST response for all response codes in API path /opsmgmt/v4.1.b1/config/reports Post operation
*/
type CreateReportApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateReportApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *CreateReportApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias CreateReportApiResponse

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

func (p *CreateReportApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias CreateReportApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewCreateReportApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewCreateReportApiResponse() *CreateReportApiResponse {
	p := new(CreateReportApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.config.CreateReportApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreateReportApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreateReportApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreateReportApiResponseData()
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
REST response for all response codes in API path /opsmgmt/v4.1.b1/config/report-configs Post operation
*/
type CreateReportConfigApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateReportConfigApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *CreateReportConfigApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias CreateReportConfigApiResponse

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

func (p *CreateReportConfigApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias CreateReportConfigApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewCreateReportConfigApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewCreateReportConfigApiResponse() *CreateReportConfigApiResponse {
	p := new(CreateReportConfigApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.config.CreateReportConfigApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreateReportConfigApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreateReportConfigApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreateReportConfigApiResponseData()
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

type Dashboard struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Description of the dashboard.
	*/
	Description *string `json:"description,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  List of global filters.
	*/
	GlobalFilters []GlobalFilter `json:"globalFilters,omitempty"`
	/*
	  Indicates if the dashboard is the home dashboard.
	*/
	IsHomeDashboard *bool `json:"isHomeDashboard,omitempty"`
	/*
	  Indicates if the view is system defined or user defined.
	*/
	IsSystemDefined *bool `json:"isSystemDefined,omitempty"`
	/*
	  Timestamp when the dashboard was last updated.
	*/
	LastUpdatedTime *time.Time `json:"lastUpdatedTime,omitempty"`

	Layout *Layout `json:"layout,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Name of the dashboard.
	*/
	Name *string `json:"name"`
	/*
	  Name of the owner of the dashboard.
	*/
	OwnerName *string `json:"ownerName,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	Type *DashboardType `json:"type,omitempty"`
	/*
	  List of widget relations in the dashboard.
	*/
	WidgetRelations []WidgetRelation `json:"widgetRelations,omitempty"`
}

func (p *Dashboard) MarshalJSON() ([]byte, error) {
	type DashboardProxy Dashboard

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*DashboardProxy
		Name *string `json:"name,omitempty"`
	}{
		DashboardProxy: (*DashboardProxy)(p),
		Name:           p.Name,
	}

	known, err := json.Marshal(baseStruct)
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

func (p *Dashboard) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Dashboard
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewDashboard()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Description != nil {
		p.Description = known.Description
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.GlobalFilters != nil {
		p.GlobalFilters = known.GlobalFilters
	}
	if known.IsHomeDashboard != nil {
		p.IsHomeDashboard = known.IsHomeDashboard
	}
	if known.IsSystemDefined != nil {
		p.IsSystemDefined = known.IsSystemDefined
	}
	if known.LastUpdatedTime != nil {
		p.LastUpdatedTime = known.LastUpdatedTime
	}
	if known.Layout != nil {
		p.Layout = known.Layout
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.OwnerName != nil {
		p.OwnerName = known.OwnerName
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}
	if known.Type != nil {
		p.Type = known.Type
	}
	if known.WidgetRelations != nil {
		p.WidgetRelations = known.WidgetRelations
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "description")
	delete(allFields, "extId")
	delete(allFields, "globalFilters")
	delete(allFields, "isHomeDashboard")
	delete(allFields, "isSystemDefined")
	delete(allFields, "lastUpdatedTime")
	delete(allFields, "layout")
	delete(allFields, "links")
	delete(allFields, "name")
	delete(allFields, "ownerName")
	delete(allFields, "tenantId")
	delete(allFields, "type")
	delete(allFields, "widgetRelations")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewDashboard() *Dashboard {
	p := new(Dashboard)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.config.Dashboard"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	p.IsHomeDashboard = new(bool)
	*p.IsHomeDashboard = false

	return p
}

/*
Configuration defining geographic layers and locations for dashboard map visualizations.
*/
type DashboardGeoconfiguration struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of geographic layers in the configuration.
	*/
	Layers []Layer `json:"layers,omitempty"`
}

func (p *DashboardGeoconfiguration) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DashboardGeoconfiguration

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

func (p *DashboardGeoconfiguration) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DashboardGeoconfiguration
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewDashboardGeoconfiguration()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Layers != nil {
		p.Layers = known.Layers
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "layers")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewDashboardGeoconfiguration() *DashboardGeoconfiguration {
	p := new(DashboardGeoconfiguration)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.config.DashboardGeoconfiguration"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Global dashboard settings that apply to all users.
*/
type DashboardSettings struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Geoconfiguration *DashboardGeoconfiguration `json:"geoconfiguration,omitempty"`
}

func (p *DashboardSettings) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DashboardSettings

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

func (p *DashboardSettings) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DashboardSettings
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewDashboardSettings()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Geoconfiguration != nil {
		p.Geoconfiguration = known.Geoconfiguration
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "geoconfiguration")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewDashboardSettings() *DashboardSettings {
	p := new(DashboardSettings)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.config.DashboardSettings"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Type of the dashboard, set for system defined dashboards.
*/
type DashboardType int

const (
	DASHBOARDTYPE_UNKNOWN     DashboardType = 0
	DASHBOARDTYPE_REDACTED    DashboardType = 1
	DASHBOARDTYPE_COST        DashboardType = 2
	DASHBOARDTYPE_PERFORMANCE DashboardType = 3
	DASHBOARDTYPE_SECURITY    DashboardType = 4
	DASHBOARDTYPE_USAGE       DashboardType = 5
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *DashboardType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"COST",
		"PERFORMANCE",
		"SECURITY",
		"USAGE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e DashboardType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"COST",
		"PERFORMANCE",
		"SECURITY",
		"USAGE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *DashboardType) index(name string) DashboardType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"COST",
		"PERFORMANCE",
		"SECURITY",
		"USAGE",
	}
	for idx := range names {
		if names[idx] == name {
			return DashboardType(idx)
		}
	}
	return DASHBOARDTYPE_UNKNOWN
}

func (e *DashboardType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for DashboardType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *DashboardType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e DashboardType) Ref() *DashboardType {
	return &e
}

/*
Filtering, sorting and limit properties for the data to collected.
*/
type DataCriteria struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Generic key value pair used for custom attributes.
	*/
	CustomParameters []import4.KVPair `json:"customParameters,omitempty"`
	/*
	  OData criteria that will be used to filter the returned data.
	*/
	FilterCriteria *string `json:"filterCriteria,omitempty"`
	/*
	  Limit on the maximum number of entities to be represented in the widget.
	*/
	Limit *int `json:"limit,omitempty"`
	/*
	  Entity Property based on which the result data is to be sorted.
	*/
	SortColumn *string `json:"sortColumn,omitempty"`

	SortKey *SortKey `json:"sortKey,omitempty"`

	SortOrder *SortOrder `json:"sortOrder,omitempty"`
}

func (p *DataCriteria) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DataCriteria

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

func (p *DataCriteria) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DataCriteria
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewDataCriteria()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.CustomParameters != nil {
		p.CustomParameters = known.CustomParameters
	}
	if known.FilterCriteria != nil {
		p.FilterCriteria = known.FilterCriteria
	}
	if known.Limit != nil {
		p.Limit = known.Limit
	}
	if known.SortColumn != nil {
		p.SortColumn = known.SortColumn
	}
	if known.SortKey != nil {
		p.SortKey = known.SortKey
	}
	if known.SortOrder != nil {
		p.SortOrder = known.SortOrder
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "customParameters")
	delete(allFields, "filterCriteria")
	delete(allFields, "limit")
	delete(allFields, "sortColumn")
	delete(allFields, "sortKey")
	delete(allFields, "sortOrder")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewDataCriteria() *DataCriteria {
	p := new(DataCriteria)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.config.DataCriteria"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /opsmgmt/v4.1.b1/config/dashboards/{extId} Delete operation
*/
type DeleteDashboardApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteDashboardApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *DeleteDashboardApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DeleteDashboardApiResponse

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

func (p *DeleteDashboardApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DeleteDashboardApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewDeleteDashboardApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewDeleteDashboardApiResponse() *DeleteDashboardApiResponse {
	p := new(DeleteDashboardApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.config.DeleteDashboardApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteDashboardApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteDashboardApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteDashboardApiResponseData()
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
REST response for all response codes in API path /opsmgmt/v4.1.b1/config/reports/{extId} Delete operation
*/
type DeleteReportApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteReportApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *DeleteReportApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DeleteReportApiResponse

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

func (p *DeleteReportApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DeleteReportApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewDeleteReportApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewDeleteReportApiResponse() *DeleteReportApiResponse {
	p := new(DeleteReportApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.config.DeleteReportApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteReportApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteReportApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteReportApiResponseData()
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
REST response for all response codes in API path /opsmgmt/v4.1.b1/config/report-configs/{extId} Delete operation
*/
type DeleteReportConfigApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteReportConfigApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *DeleteReportConfigApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DeleteReportConfigApiResponse

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

func (p *DeleteReportConfigApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DeleteReportConfigApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewDeleteReportConfigApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewDeleteReportConfigApiResponse() *DeleteReportConfigApiResponse {
	p := new(DeleteReportConfigApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.config.DeleteReportConfigApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteReportConfigApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteReportConfigApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteReportConfigApiResponseData()
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
Run Time Parameters, which supply entity type and entityId for which reports need to be generated.
*/
type EntitySelection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of entity ExtId.
	*/
	EntityExtId []string `json:"entityExtId"`

	EntityType *EntityType `json:"entityType"`
}

func (p *EntitySelection) MarshalJSON() ([]byte, error) {
	type EntitySelectionProxy EntitySelection

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*EntitySelectionProxy
		EntityExtId []string    `json:"entityExtId,omitempty"`
		EntityType  *EntityType `json:"entityType,omitempty"`
	}{
		EntitySelectionProxy: (*EntitySelectionProxy)(p),
		EntityExtId:          p.EntityExtId,
		EntityType:           p.EntityType,
	}

	known, err := json.Marshal(baseStruct)
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

func (p *EntitySelection) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias EntitySelection
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewEntitySelection()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.EntityExtId != nil {
		p.EntityExtId = known.EntityExtId
	}
	if known.EntityType != nil {
		p.EntityType = known.EntityType
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "entityExtId")
	delete(allFields, "entityType")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewEntitySelection() *EntitySelection {
	p := new(EntitySelection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.config.EntitySelection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Entity type for which data needs to be shown in the widget/section.
*/
type EntityType int

const (
	ENTITYTYPE_UNKNOWN           EntityType = 0
	ENTITYTYPE_REDACTED          EntityType = 1
	ENTITYTYPE_VM                EntityType = 2
	ENTITYTYPE_CLUSTER           EntityType = 3
	ENTITYTYPE_HOST              EntityType = 4
	ENTITYTYPE_CONTAINER         EntityType = 5
	ENTITYTYPE_DISK              EntityType = 6
	ENTITYTYPE_VIRTUAL_DISK      EntityType = 7
	ENTITYTYPE_VCENTER_VM        EntityType = 8
	ENTITYTYPE_VCENTER_CLUSTER   EntityType = 9
	ENTITYTYPE_VCENTER_HOST      EntityType = 10
	ENTITYTYPE_VCENTER_DATASTORE EntityType = 11
	ENTITYTYPE_ALERT             EntityType = 12
	ENTITYTYPE_EVENT             EntityType = 13
	ENTITYTYPE_AUDIT             EntityType = 14
	ENTITYTYPE_PLAYBOOK          EntityType = 15
	ENTITYTYPE_VOLUME_GROUPS     EntityType = 16
	ENTITYTYPE_CONFIG            EntityType = 17
	ENTITYTYPE_RECOVERY_PLAN_JOB EntityType = 18
	ENTITYTYPE_CATEGORY          EntityType = 19
	ENTITYTYPE_VULNERABILITY     EntityType = 20
	ENTITYTYPE_STIG_STATS        EntityType = 21
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *EntityType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"VM",
		"CLUSTER",
		"HOST",
		"CONTAINER",
		"DISK",
		"VIRTUAL_DISK",
		"VCENTER_VM",
		"VCENTER_CLUSTER",
		"VCENTER_HOST",
		"VCENTER_DATASTORE",
		"ALERT",
		"EVENT",
		"AUDIT",
		"PLAYBOOK",
		"VOLUME_GROUPS",
		"CONFIG",
		"RECOVERY_PLAN_JOB",
		"CATEGORY",
		"VULNERABILITY",
		"STIG_STATS",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e EntityType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"VM",
		"CLUSTER",
		"HOST",
		"CONTAINER",
		"DISK",
		"VIRTUAL_DISK",
		"VCENTER_VM",
		"VCENTER_CLUSTER",
		"VCENTER_HOST",
		"VCENTER_DATASTORE",
		"ALERT",
		"EVENT",
		"AUDIT",
		"PLAYBOOK",
		"VOLUME_GROUPS",
		"CONFIG",
		"RECOVERY_PLAN_JOB",
		"CATEGORY",
		"VULNERABILITY",
		"STIG_STATS",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *EntityType) index(name string) EntityType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"VM",
		"CLUSTER",
		"HOST",
		"CONTAINER",
		"DISK",
		"VIRTUAL_DISK",
		"VCENTER_VM",
		"VCENTER_CLUSTER",
		"VCENTER_HOST",
		"VCENTER_DATASTORE",
		"ALERT",
		"EVENT",
		"AUDIT",
		"PLAYBOOK",
		"VOLUME_GROUPS",
		"CONFIG",
		"RECOVERY_PLAN_JOB",
		"CATEGORY",
		"VULNERABILITY",
		"STIG_STATS",
	}
	for idx := range names {
		if names[idx] == name {
			return EntityType(idx)
		}
	}
	return ENTITYTYPE_UNKNOWN
}

func (e *EntityType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for EntityType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *EntityType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e EntityType) Ref() *EntityType {
	return &e
}

/*
Geometric data defining location coordinates.
*/
type Geometry struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Array of coordinates [longitude, latitude].
	*/
	Coordinates []float64 `json:"coordinates"`

	Type *GeometryType `json:"type"`
}

func (p *Geometry) MarshalJSON() ([]byte, error) {
	type GeometryProxy Geometry

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*GeometryProxy
		Coordinates []float64     `json:"coordinates,omitempty"`
		Type        *GeometryType `json:"type,omitempty"`
	}{
		GeometryProxy: (*GeometryProxy)(p),
		Coordinates:   p.Coordinates,
		Type:          p.Type,
	}

	known, err := json.Marshal(baseStruct)
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

func (p *Geometry) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Geometry
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewGeometry()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Coordinates != nil {
		p.Coordinates = known.Coordinates
	}
	if known.Type != nil {
		p.Type = known.Type
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "coordinates")
	delete(allFields, "type")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewGeometry() *Geometry {
	p := new(Geometry)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.config.Geometry"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Enumeration of supported geometry types.
*/
type GeometryType int

const (
	GEOMETRYTYPE_UNKNOWN  GeometryType = 0
	GEOMETRYTYPE_REDACTED GeometryType = 1
	GEOMETRYTYPE_POINT    GeometryType = 2
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *GeometryType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"POINT",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e GeometryType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"POINT",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *GeometryType) index(name string) GeometryType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"POINT",
	}
	for idx := range names {
		if names[idx] == name {
			return GeometryType(idx)
		}
	}
	return GEOMETRYTYPE_UNKNOWN
}

func (e *GeometryType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for GeometryType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *GeometryType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e GeometryType) Ref() *GeometryType {
	return &e
}

/*
REST response for all response codes in API path /opsmgmt/v4.1.b1/config/dashboards/{extId} Get operation
*/
type GetDashboardApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetDashboardApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *GetDashboardApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetDashboardApiResponse

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

func (p *GetDashboardApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetDashboardApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewGetDashboardApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewGetDashboardApiResponse() *GetDashboardApiResponse {
	p := new(GetDashboardApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.config.GetDashboardApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetDashboardApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetDashboardApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetDashboardApiResponseData()
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
REST response for all response codes in API path /opsmgmt/v4.1.b1/config/dashboard-settings Get operation
*/
type GetDashboardSettingsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetDashboardSettingsApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *GetDashboardSettingsApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetDashboardSettingsApiResponse

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

func (p *GetDashboardSettingsApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetDashboardSettingsApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewGetDashboardSettingsApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewGetDashboardSettingsApiResponse() *GetDashboardSettingsApiResponse {
	p := new(GetDashboardSettingsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.config.GetDashboardSettingsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetDashboardSettingsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetDashboardSettingsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetDashboardSettingsApiResponseData()
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
REST response for all response codes in API path /opsmgmt/v4.1.b1/config/user/{userExtId}/global-report-setting Get operation
*/
type GetGlobalReportSettingApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetGlobalReportSettingApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *GetGlobalReportSettingApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetGlobalReportSettingApiResponse

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

func (p *GetGlobalReportSettingApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetGlobalReportSettingApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewGetGlobalReportSettingApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewGetGlobalReportSettingApiResponse() *GetGlobalReportSettingApiResponse {
	p := new(GetGlobalReportSettingApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.config.GetGlobalReportSettingApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetGlobalReportSettingApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetGlobalReportSettingApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetGlobalReportSettingApiResponseData()
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
REST response for all response codes in API path /opsmgmt/v4.1.b1/config/reports/{extId} Get operation
*/
type GetReportApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetReportApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *GetReportApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetReportApiResponse

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

func (p *GetReportApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetReportApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewGetReportApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewGetReportApiResponse() *GetReportApiResponse {
	p := new(GetReportApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.config.GetReportApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetReportApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetReportApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetReportApiResponseData()
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
REST response for all response codes in API path /opsmgmt/v4.1.b1/config/report-configs/{extId} Get operation
*/
type GetReportConfigApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetReportConfigApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *GetReportConfigApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetReportConfigApiResponse

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

func (p *GetReportConfigApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetReportConfigApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewGetReportConfigApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewGetReportConfigApiResponse() *GetReportConfigApiResponse {
	p := new(GetReportConfigApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.config.GetReportConfigApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetReportConfigApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetReportConfigApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetReportConfigApiResponseData()
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
Global filters applicable to the entire dashboard
*/
type GlobalFilter int

const (
	GLOBALFILTER_UNKNOWN    GlobalFilter = 0
	GLOBALFILTER_REDACTED   GlobalFilter = 1
	GLOBALFILTER_PROVIDER   GlobalFilter = 2
	GLOBALFILTER_ACCOUNT    GlobalFilter = 3
	GLOBALFILTER_CATEGORIES GlobalFilter = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *GlobalFilter) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PROVIDER",
		"ACCOUNT",
		"CATEGORIES",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e GlobalFilter) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PROVIDER",
		"ACCOUNT",
		"CATEGORIES",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *GlobalFilter) index(name string) GlobalFilter {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PROVIDER",
		"ACCOUNT",
		"CATEGORIES",
	}
	for idx := range names {
		if names[idx] == name {
			return GlobalFilter(idx)
		}
	}
	return GLOBALFILTER_UNKNOWN
}

func (e *GlobalFilter) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for GlobalFilter:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *GlobalFilter) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e GlobalFilter) Ref() *GlobalFilter {
	return &e
}

type GlobalReportSetting struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Name of the global report setting.
	*/
	Name *string `json:"name"`

	NotificationPolicy *NotificationPolicy `json:"notificationPolicy,omitempty"`

	ReportCustomization *ReportCustomization `json:"reportCustomization,omitempty"`

	RetentionConfig *RetentionConfig `json:"retentionConfig,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *GlobalReportSetting) MarshalJSON() ([]byte, error) {
	type GlobalReportSettingProxy GlobalReportSetting

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*GlobalReportSettingProxy
		Name *string `json:"name,omitempty"`
	}{
		GlobalReportSettingProxy: (*GlobalReportSettingProxy)(p),
		Name:                     p.Name,
	}

	known, err := json.Marshal(baseStruct)
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

func (p *GlobalReportSetting) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GlobalReportSetting
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewGlobalReportSetting()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.NotificationPolicy != nil {
		p.NotificationPolicy = known.NotificationPolicy
	}
	if known.ReportCustomization != nil {
		p.ReportCustomization = known.ReportCustomization
	}
	if known.RetentionConfig != nil {
		p.RetentionConfig = known.RetentionConfig
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "extId")
	delete(allFields, "links")
	delete(allFields, "name")
	delete(allFields, "notificationPolicy")
	delete(allFields, "reportCustomization")
	delete(allFields, "retentionConfig")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewGlobalReportSetting() *GlobalReportSetting {
	p := new(GlobalReportSetting)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.config.GlobalReportSetting"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
A geographic layer containing labeled locations.
*/
type Layer struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Label identifying the layer type (e.g., Region, Country, City).
	*/
	Label *string `json:"label"`
	/*
	  List of geographic locations in this layer.
	*/
	Locations []Location `json:"locations"`
}

func (p *Layer) MarshalJSON() ([]byte, error) {
	type LayerProxy Layer

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*LayerProxy
		Label     *string    `json:"label,omitempty"`
		Locations []Location `json:"locations,omitempty"`
	}{
		LayerProxy: (*LayerProxy)(p),
		Label:      p.Label,
		Locations:  p.Locations,
	}

	known, err := json.Marshal(baseStruct)
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

func (p *Layer) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Layer
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewLayer()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Label != nil {
		p.Label = known.Label
	}
	if known.Locations != nil {
		p.Locations = known.Locations
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "label")
	delete(allFields, "locations")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewLayer() *Layer {
	p := new(Layer)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.config.Layer"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Layout of the widgets in the dashboard. Maximum allowed nesting level is 2
*/
type Layout struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Bounds of the widget in the dashboard.
	*/
	Bounds []Bounds `json:"bounds,omitempty"`
	/*
	  Name of the layout in the dashboard.
	*/
	Heading *string `json:"heading,omitempty"`
	/*
	  List of nodes in the dashboard.
	*/
	Nodes []Node `json:"nodes,omitempty"`

	Type *LayoutType `json:"type"`
}

func (p *Layout) MarshalJSON() ([]byte, error) {
	type LayoutProxy Layout

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*LayoutProxy
		Type *LayoutType `json:"type,omitempty"`
	}{
		LayoutProxy: (*LayoutProxy)(p),
		Type:        p.Type,
	}

	known, err := json.Marshal(baseStruct)
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

func (p *Layout) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Layout
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewLayout()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Bounds != nil {
		p.Bounds = known.Bounds
	}
	if known.Heading != nil {
		p.Heading = known.Heading
	}
	if known.Nodes != nil {
		p.Nodes = known.Nodes
	}
	if known.Type != nil {
		p.Type = known.Type
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "bounds")
	delete(allFields, "heading")
	delete(allFields, "nodes")
	delete(allFields, "type")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewLayout() *Layout {
	p := new(Layout)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.config.Layout"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Type of the layout in the dashboard.
*/
type LayoutType int

const (
	LAYOUTTYPE_UNKNOWN     LayoutType = 0
	LAYOUTTYPE_REDACTED    LayoutType = 1
	LAYOUTTYPE_GRID_LAYOUT LayoutType = 2
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *LayoutType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"GRID_LAYOUT",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e LayoutType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"GRID_LAYOUT",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *LayoutType) index(name string) LayoutType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"GRID_LAYOUT",
	}
	for idx := range names {
		if names[idx] == name {
			return LayoutType(idx)
		}
	}
	return LAYOUTTYPE_UNKNOWN
}

func (e *LayoutType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for LayoutType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *LayoutType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e LayoutType) Ref() *LayoutType {
	return &e
}

/*
REST response for all response codes in API path /opsmgmt/v4.1.b1/config/dashboards Get operation
*/
type ListDashboardApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListDashboardApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ListDashboardApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ListDashboardApiResponse

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

func (p *ListDashboardApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ListDashboardApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewListDashboardApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewListDashboardApiResponse() *ListDashboardApiResponse {
	p := new(ListDashboardApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.config.ListDashboardApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListDashboardApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListDashboardApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListDashboardApiResponseData()
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
REST response for all response codes in API path /opsmgmt/v4.1.b1/config/report-configs Get operation
*/
type ListReportConfigsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListReportConfigsApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ListReportConfigsApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ListReportConfigsApiResponse

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

func (p *ListReportConfigsApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ListReportConfigsApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewListReportConfigsApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewListReportConfigsApiResponse() *ListReportConfigsApiResponse {
	p := new(ListReportConfigsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.config.ListReportConfigsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListReportConfigsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListReportConfigsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListReportConfigsApiResponseData()
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
REST response for all response codes in API path /opsmgmt/v4.1.b1/config/reports Get operation
*/
type ListReportsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListReportsApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ListReportsApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ListReportsApiResponse

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

func (p *ListReportsApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ListReportsApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewListReportsApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewListReportsApiResponse() *ListReportsApiResponse {
	p := new(ListReportsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.config.ListReportsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListReportsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListReportsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListReportsApiResponseData()
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
A geographic location with coordinates and metadata.
*/
type Location struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	ChildrenItemDiscriminator_ *string `json:"$childrenItemDiscriminator,omitempty"`
	/*
	  Child layer containing nested locations.
	*/
	Children *OneOfLocationChildren `json:"children,omitempty"`

	Geometry *Geometry `json:"geometry"`
	/*
	  Unique identifier for the location.
	*/
	Id *string `json:"id"`
	/*
	  Display label for the location.
	*/
	Label *string `json:"label"`
	/*
	  Key-value properties associated with the location.
	*/
	Properties []import4.KVPair `json:"properties,omitempty"`
}

func (p *Location) MarshalJSON() ([]byte, error) {
	type LocationProxy Location

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*LocationProxy
		Geometry *Geometry `json:"geometry,omitempty"`
		Id       *string   `json:"id,omitempty"`
		Label    *string   `json:"label,omitempty"`
	}{
		LocationProxy: (*LocationProxy)(p),
		Geometry:      p.Geometry,
		Id:            p.Id,
		Label:         p.Label,
	}

	known, err := json.Marshal(baseStruct)
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

func (p *Location) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Location
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewLocation()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ChildrenItemDiscriminator_ != nil {
		p.ChildrenItemDiscriminator_ = known.ChildrenItemDiscriminator_
	}
	if known.Children != nil {
		p.Children = known.Children
	}
	if known.Geometry != nil {
		p.Geometry = known.Geometry
	}
	if known.Id != nil {
		p.Id = known.Id
	}
	if known.Label != nil {
		p.Label = known.Label
	}
	if known.Properties != nil {
		p.Properties = known.Properties
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$childrenItemDiscriminator")
	delete(allFields, "children")
	delete(allFields, "geometry")
	delete(allFields, "id")
	delete(allFields, "label")
	delete(allFields, "properties")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewLocation() *Location {
	p := new(Location)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.config.Location"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *Location) GetChildren() interface{} {
	if nil == p.Children {
		return nil
	}
	return p.Children.GetValue()
}

func (p *Location) SetChildren(v interface{}) error {
	if nil == p.Children {
		p.Children = NewOneOfLocationChildren()
	}
	e := p.Children.SetValue(v)
	if nil == e {
		if nil == p.ChildrenItemDiscriminator_ {
			p.ChildrenItemDiscriminator_ = new(string)
		}
		*p.ChildrenItemDiscriminator_ = *p.Children.Discriminator
	}
	return e
}

/*
Node of the dashboard.
*/
type Node struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	NodeInfoItemDiscriminator_ *string `json:"$nodeInfoItemDiscriminator,omitempty"`
	/*
	  Information of the node in the dashboard.
	*/
	NodeInfo *OneOfNodeNodeInfo `json:"nodeInfo,omitempty"`
}

func (p *Node) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias Node

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

func (p *Node) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Node
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewNode()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.NodeInfoItemDiscriminator_ != nil {
		p.NodeInfoItemDiscriminator_ = known.NodeInfoItemDiscriminator_
	}
	if known.NodeInfo != nil {
		p.NodeInfo = known.NodeInfo
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$nodeInfoItemDiscriminator")
	delete(allFields, "nodeInfo")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewNode() *Node {
	p := new(Node)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.config.Node"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *Node) GetNodeInfo() interface{} {
	if nil == p.NodeInfo {
		return nil
	}
	return p.NodeInfo.GetValue()
}

func (p *Node) SetNodeInfo(v interface{}) error {
	if nil == p.NodeInfo {
		p.NodeInfo = NewOneOfNodeNodeInfo()
	}
	e := p.NodeInfo.SetValue(v)
	if nil == e {
		if nil == p.NodeInfoItemDiscriminator_ {
			p.NodeInfoItemDiscriminator_ = new(string)
		}
		*p.NodeInfoItemDiscriminator_ = *p.NodeInfo.Discriminator
	}
	return e
}

/*
Notification policy for sending the email of the generated report.
*/
type NotificationPolicy struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Content of the email body.
	*/
	EmailBody *string `json:"emailBody,omitempty"`
	/*
	  Subject of the email to be sent for the report.
	*/
	EmailSubject *string `json:"emailSubject,omitempty"`
	/*
	  List specifying the formats in which the report is to be sent.
	*/
	RecipientFormats []ReportFormat `json:"recipientFormats,omitempty"`
	/*
	  Email recipients list.
	*/
	Recipients []Recipient `json:"recipients"`
}

func (p *NotificationPolicy) MarshalJSON() ([]byte, error) {
	type NotificationPolicyProxy NotificationPolicy

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*NotificationPolicyProxy
		Recipients []Recipient `json:"recipients,omitempty"`
	}{
		NotificationPolicyProxy: (*NotificationPolicyProxy)(p),
		Recipients:              p.Recipients,
	}

	known, err := json.Marshal(baseStruct)
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

func (p *NotificationPolicy) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias NotificationPolicy
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewNotificationPolicy()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.EmailBody != nil {
		p.EmailBody = known.EmailBody
	}
	if known.EmailSubject != nil {
		p.EmailSubject = known.EmailSubject
	}
	if known.RecipientFormats != nil {
		p.RecipientFormats = known.RecipientFormats
	}
	if known.Recipients != nil {
		p.Recipients = known.Recipients
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "emailBody")
	delete(allFields, "emailSubject")
	delete(allFields, "recipientFormats")
	delete(allFields, "recipients")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewNotificationPolicy() *NotificationPolicy {
	p := new(NotificationPolicy)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.config.NotificationPolicy"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /opsmgmt/v4.1.b1/config/reports/{extId}/$actions/notify-recipients Post operation
*/
type NotifyReportApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfNotifyReportApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *NotifyReportApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias NotifyReportApiResponse

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

func (p *NotifyReportApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias NotifyReportApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewNotifyReportApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewNotifyReportApiResponse() *NotifyReportApiResponse {
	p := new(NotifyReportApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.config.NotifyReportApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *NotifyReportApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *NotifyReportApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfNotifyReportApiResponseData()
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
Recipient of the report email.
*/
type Recipient struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Email address of the recipient.
	*/
	EmailAddress *string `json:"emailAddress"`
	/*
	  Name of the recipient.
	*/
	RecipientName *string `json:"recipientName,omitempty"`
}

func (p *Recipient) MarshalJSON() ([]byte, error) {
	type RecipientProxy Recipient

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*RecipientProxy
		EmailAddress *string `json:"emailAddress,omitempty"`
	}{
		RecipientProxy: (*RecipientProxy)(p),
		EmailAddress:   p.EmailAddress,
	}

	known, err := json.Marshal(baseStruct)
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

func (p *Recipient) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Recipient
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewRecipient()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.EmailAddress != nil {
		p.EmailAddress = known.EmailAddress
	}
	if known.RecipientName != nil {
		p.RecipientName = known.RecipientName
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "emailAddress")
	delete(allFields, "recipientName")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewRecipient() *Recipient {
	p := new(Recipient)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.config.Recipient"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Criteria for repeating a widget/section.
*/
type RepeatCriteria struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	EntityType *EntityType `json:"entityType"`
	/*
	  Rule based on which the widget/section will be repeated.
	*/
	RepetitionRule *string `json:"repetitionRule,omitempty"`
}

func (p *RepeatCriteria) MarshalJSON() ([]byte, error) {
	type RepeatCriteriaProxy RepeatCriteria

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*RepeatCriteriaProxy
		EntityType *EntityType `json:"entityType,omitempty"`
	}{
		RepeatCriteriaProxy: (*RepeatCriteriaProxy)(p),
		EntityType:          p.EntityType,
	}

	known, err := json.Marshal(baseStruct)
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

func (p *RepeatCriteria) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias RepeatCriteria
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewRepeatCriteria()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.EntityType != nil {
		p.EntityType = known.EntityType
	}
	if known.RepetitionRule != nil {
		p.RepetitionRule = known.RepetitionRule
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "entityType")
	delete(allFields, "repetitionRule")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewRepeatCriteria() *RepeatCriteria {
	p := new(RepeatCriteria)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.config.RepeatCriteria"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type Report struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of formats in which report generation was successful.
	*/
	AvailableFormats []ReportFormat `json:"availableFormats,omitempty"`
	/*
	  UUID for the report configuration for which report needs to be generated.
	*/
	ConfigExtId *string `json:"configExtId"`
	/*
	  Time in ISO 8601 format when the report instance was created.
	*/
	CreationTime *time.Time `json:"creationTime,omitempty"`
	/*
	  Description of the report. This will be part of generated report.
	*/
	Description *string `json:"description,omitempty"`
	/*
	  UTC date and time in "%Y-%m-%d %H:%M:%S" format for data collection end point. Eg:- 2023-10-23 11:34:45
	*/
	EndTime *time.Time `json:"endTime"`

	EntitySelection *EntitySelection `json:"entitySelection,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Generated report saved or not
	*/
	IsPersistent *bool `json:"isPersistent,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Report instance name
	*/
	Name *string `json:"name"`
	/*
	  List specifying the formats in which report is to be created. This overrides the supportedFormats defined in the report configuration.
	*/
	OverrideSupportedFormats []ReportFormat `json:"overrideSupportedFormats,omitempty"`
	/*
	  Owner UUID for the report instance."
	*/
	OwnerExtId *string `json:"ownerExtId,omitempty"`
	/*
	  List specifying the formats in which the report is to be sent.
	*/
	RecipientFormats []ReportFormat `json:"recipientFormats,omitempty"`
	/*
	  Recipients in addition to the ones specified on the report configuration.
	*/
	Recipients []Recipient `json:"recipients,omitempty"`
	/*
	  UTC date and time in "%Y-%m-%d %H:%M:%S" format for data collection start point. Eg:- 2023-10-23 11:34:45
	*/
	StartTime *time.Time `json:"startTime"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  Timezone in which report is to be generated. This is the list supported by pytz.all_timezones. For more info, check http://pytz.sourceforge.net
	*/
	Timezone *string `json:"timezone,omitempty"`
}

func (p *Report) MarshalJSON() ([]byte, error) {
	type ReportProxy Report

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*ReportProxy
		ConfigExtId *string    `json:"configExtId,omitempty"`
		EndTime     *time.Time `json:"endTime,omitempty"`
		Name        *string    `json:"name,omitempty"`
		StartTime   *time.Time `json:"startTime,omitempty"`
	}{
		ReportProxy: (*ReportProxy)(p),
		ConfigExtId: p.ConfigExtId,
		EndTime:     p.EndTime,
		Name:        p.Name,
		StartTime:   p.StartTime,
	}

	known, err := json.Marshal(baseStruct)
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

func (p *Report) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Report
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewReport()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.AvailableFormats != nil {
		p.AvailableFormats = known.AvailableFormats
	}
	if known.ConfigExtId != nil {
		p.ConfigExtId = known.ConfigExtId
	}
	if known.CreationTime != nil {
		p.CreationTime = known.CreationTime
	}
	if known.Description != nil {
		p.Description = known.Description
	}
	if known.EndTime != nil {
		p.EndTime = known.EndTime
	}
	if known.EntitySelection != nil {
		p.EntitySelection = known.EntitySelection
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.IsPersistent != nil {
		p.IsPersistent = known.IsPersistent
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.OverrideSupportedFormats != nil {
		p.OverrideSupportedFormats = known.OverrideSupportedFormats
	}
	if known.OwnerExtId != nil {
		p.OwnerExtId = known.OwnerExtId
	}
	if known.RecipientFormats != nil {
		p.RecipientFormats = known.RecipientFormats
	}
	if known.Recipients != nil {
		p.Recipients = known.Recipients
	}
	if known.StartTime != nil {
		p.StartTime = known.StartTime
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}
	if known.Timezone != nil {
		p.Timezone = known.Timezone
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "availableFormats")
	delete(allFields, "configExtId")
	delete(allFields, "creationTime")
	delete(allFields, "description")
	delete(allFields, "endTime")
	delete(allFields, "entitySelection")
	delete(allFields, "extId")
	delete(allFields, "isPersistent")
	delete(allFields, "links")
	delete(allFields, "name")
	delete(allFields, "overrideSupportedFormats")
	delete(allFields, "ownerExtId")
	delete(allFields, "recipientFormats")
	delete(allFields, "recipients")
	delete(allFields, "startTime")
	delete(allFields, "tenantId")
	delete(allFields, "timezone")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewReport() *Report {
	p := new(Report)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.config.Report"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type ReportConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Time in ISO 8601 format when the report configuration was created.
	*/
	CreationTime *time.Time `json:"creationTime,omitempty"`

	DefaultSectionEntityType *EntityType `json:"defaultSectionEntityType,omitempty"`
	/*
	  Description of the report configuration.
	*/
	Description *string `json:"description,omitempty"`
	/*
	  Offset for end time for data collection during report generation.
	*/
	EndTimeOffsetSecs *int64 `json:"endTimeOffsetSecs,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Flag specifying if the report configuration is imported.
	*/
	IsImported *bool `json:"isImported,omitempty"`
	/*
	  Flag specifying if the report configuration is private.
	*/
	IsPrivate *bool `json:"isPrivate,omitempty"`
	/*
	  Flag specifying if the report configuration is pre-defined.
	*/
	IsSystemDefined *bool `json:"isSystemDefined,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Name of the report configuration.
	*/
	Name *string `json:"name"`

	NotificationPolicy *NotificationPolicy `json:"notificationPolicy,omitempty"`

	ReportCustomization *ReportCustomization `json:"reportCustomization,omitempty"`

	RetentionConfig *RetentionConfig `json:"retentionConfig,omitempty"`

	Schedule *ReportSchedule `json:"schedule,omitempty"`
	/*
	  List of sections in the report.
	*/
	Sections []Section `json:"sections"`
	/*
	  Offset for start time for data collection during report generation.
	*/
	StartTimeOffsetSecs *int64 `json:"startTimeOffsetSecs,omitempty"`
	/*
	  List specifying the formats in which the report will be created.
	*/
	SupportedFormats []ReportFormat `json:"supportedFormats,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  The timezone in which the report will be generated. This is the list supported by pytz.all_timezones. For more info, check http://pytz.sourceforge.net
	*/
	Timezone *string `json:"timezone,omitempty"`
}

func (p *ReportConfig) MarshalJSON() ([]byte, error) {
	type ReportConfigProxy ReportConfig

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*ReportConfigProxy
		Name     *string   `json:"name,omitempty"`
		Sections []Section `json:"sections,omitempty"`
	}{
		ReportConfigProxy: (*ReportConfigProxy)(p),
		Name:              p.Name,
		Sections:          p.Sections,
	}

	known, err := json.Marshal(baseStruct)
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

func (p *ReportConfig) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ReportConfig
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewReportConfig()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.CreationTime != nil {
		p.CreationTime = known.CreationTime
	}
	if known.DefaultSectionEntityType != nil {
		p.DefaultSectionEntityType = known.DefaultSectionEntityType
	}
	if known.Description != nil {
		p.Description = known.Description
	}
	if known.EndTimeOffsetSecs != nil {
		p.EndTimeOffsetSecs = known.EndTimeOffsetSecs
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.IsImported != nil {
		p.IsImported = known.IsImported
	}
	if known.IsPrivate != nil {
		p.IsPrivate = known.IsPrivate
	}
	if known.IsSystemDefined != nil {
		p.IsSystemDefined = known.IsSystemDefined
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.NotificationPolicy != nil {
		p.NotificationPolicy = known.NotificationPolicy
	}
	if known.ReportCustomization != nil {
		p.ReportCustomization = known.ReportCustomization
	}
	if known.RetentionConfig != nil {
		p.RetentionConfig = known.RetentionConfig
	}
	if known.Schedule != nil {
		p.Schedule = known.Schedule
	}
	if known.Sections != nil {
		p.Sections = known.Sections
	}
	if known.StartTimeOffsetSecs != nil {
		p.StartTimeOffsetSecs = known.StartTimeOffsetSecs
	}
	if known.SupportedFormats != nil {
		p.SupportedFormats = known.SupportedFormats
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}
	if known.Timezone != nil {
		p.Timezone = known.Timezone
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "creationTime")
	delete(allFields, "defaultSectionEntityType")
	delete(allFields, "description")
	delete(allFields, "endTimeOffsetSecs")
	delete(allFields, "extId")
	delete(allFields, "isImported")
	delete(allFields, "isPrivate")
	delete(allFields, "isSystemDefined")
	delete(allFields, "links")
	delete(allFields, "name")
	delete(allFields, "notificationPolicy")
	delete(allFields, "reportCustomization")
	delete(allFields, "retentionConfig")
	delete(allFields, "schedule")
	delete(allFields, "sections")
	delete(allFields, "startTimeOffsetSecs")
	delete(allFields, "supportedFormats")
	delete(allFields, "tenantId")
	delete(allFields, "timezone")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewReportConfig() *ReportConfig {
	p := new(ReportConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.config.ReportConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Report-level customizations.
*/
type ReportCustomization struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Global cascadable style for the report.
	*/
	CssStyleSheet *string `json:"cssStyleSheet,omitempty"`
	/*
	  Custom footer HTML for the report.
	*/
	FooterHtml *string `json:"footerHtml,omitempty"`
	/*
	  Custom header HTML for the report.
	*/
	HeaderHtml *string `json:"headerHtml,omitempty"`
	/*
	  Custom logo for the report as selected by the user. The logo extID can be fetched using report artifact list API.
	*/
	LogoImageExtId *string `json:"logoImageExtId,omitempty"`
}

func (p *ReportCustomization) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ReportCustomization

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

func (p *ReportCustomization) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ReportCustomization
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewReportCustomization()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.CssStyleSheet != nil {
		p.CssStyleSheet = known.CssStyleSheet
	}
	if known.FooterHtml != nil {
		p.FooterHtml = known.FooterHtml
	}
	if known.HeaderHtml != nil {
		p.HeaderHtml = known.HeaderHtml
	}
	if known.LogoImageExtId != nil {
		p.LogoImageExtId = known.LogoImageExtId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "cssStyleSheet")
	delete(allFields, "footerHtml")
	delete(allFields, "headerHtml")
	delete(allFields, "logoImageExtId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewReportCustomization() *ReportCustomization {
	p := new(ReportCustomization)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.config.ReportCustomization"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
List specifying the formats in which the report will be created.
*/
type ReportFormat int

const (
	REPORTFORMAT_UNKNOWN  ReportFormat = 0
	REPORTFORMAT_REDACTED ReportFormat = 1
	REPORTFORMAT_PDF      ReportFormat = 2
	REPORTFORMAT_CSV      ReportFormat = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ReportFormat) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PDF",
		"CSV",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ReportFormat) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PDF",
		"CSV",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ReportFormat) index(name string) ReportFormat {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PDF",
		"CSV",
	}
	for idx := range names {
		if names[idx] == name {
			return ReportFormat(idx)
		}
	}
	return REPORTFORMAT_UNKNOWN
}

func (e *ReportFormat) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ReportFormat:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ReportFormat) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ReportFormat) Ref() *ReportFormat {
	return &e
}

type ReportNotificationSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  List specifying the formats in which the report is to be sent.
	*/
	RecipientFormats []ReportFormat `json:"recipientFormats"`
	/*
	  Recipients to notify with the report.
	*/
	Recipients []Recipient `json:"recipients"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *ReportNotificationSpec) MarshalJSON() ([]byte, error) {
	type ReportNotificationSpecProxy ReportNotificationSpec

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*ReportNotificationSpecProxy
		RecipientFormats []ReportFormat `json:"recipientFormats,omitempty"`
		Recipients       []Recipient    `json:"recipients,omitempty"`
	}{
		ReportNotificationSpecProxy: (*ReportNotificationSpecProxy)(p),
		RecipientFormats:            p.RecipientFormats,
		Recipients:                  p.Recipients,
	}

	known, err := json.Marshal(baseStruct)
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

func (p *ReportNotificationSpec) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ReportNotificationSpec
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewReportNotificationSpec()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.RecipientFormats != nil {
		p.RecipientFormats = known.RecipientFormats
	}
	if known.Recipients != nil {
		p.Recipients = known.Recipients
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "extId")
	delete(allFields, "links")
	delete(allFields, "recipientFormats")
	delete(allFields, "recipients")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewReportNotificationSpec() *ReportNotificationSpec {
	p := new(ReportNotificationSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.config.ReportNotificationSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Defines the parameters for schdeuling report creation from the report configuration.
*/
type ReportSchedule struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  End time of the schedule. The value should be in extended ISO-8601 format. For example, 2022-04-23T01:23:45.678+09:00 represents 1:23:45.678 on the 23rd of April 2022. Details around ISO-8601 format can be found at https://www.iso.org/standard/70907.html
	*/
	EndTime *time.Time `json:"endTime,omitempty"`
	/*
	  Multiple of scheduleInterval. For example, if the scheduleInterval is set to daily and frequency is set to 2, the schedule will run every 2 days.
	*/
	Frequency *int `json:"frequency"`

	ScheduleInterval *ScheduleInterval `json:"scheduleInterval"`
	/*
	  Start time of the schedule. The value should be in extended ISO-8601 format. For example, 2022-04-23T01:23:45.678+09:00 represents 1:23:45.678 on the 23rd of April 2022. Details around ISO-8601 format can be found at https://www.iso.org/standard/70907.html
	*/
	StartTime *time.Time `json:"startTime"`
}

func (p *ReportSchedule) MarshalJSON() ([]byte, error) {
	type ReportScheduleProxy ReportSchedule

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*ReportScheduleProxy
		Frequency        *int              `json:"frequency,omitempty"`
		ScheduleInterval *ScheduleInterval `json:"scheduleInterval,omitempty"`
		StartTime        *time.Time        `json:"startTime,omitempty"`
	}{
		ReportScheduleProxy: (*ReportScheduleProxy)(p),
		Frequency:           p.Frequency,
		ScheduleInterval:    p.ScheduleInterval,
		StartTime:           p.StartTime,
	}

	known, err := json.Marshal(baseStruct)
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

func (p *ReportSchedule) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ReportSchedule
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewReportSchedule()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.EndTime != nil {
		p.EndTime = known.EndTime
	}
	if known.Frequency != nil {
		p.Frequency = known.Frequency
	}
	if known.ScheduleInterval != nil {
		p.ScheduleInterval = known.ScheduleInterval
	}
	if known.StartTime != nil {
		p.StartTime = known.StartTime
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "endTime")
	delete(allFields, "frequency")
	delete(allFields, "scheduleInterval")
	delete(allFields, "startTime")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewReportSchedule() *ReportSchedule {
	p := new(ReportSchedule)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.config.ReportSchedule"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Defines how long to retain a report generated from the report configuration. Only one of the retentionPeriodSeconds and retentionCount should be specified.
*/
type RetentionConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Number of reports to be retained.
	*/
	RetentionCount *int `json:"retentionCount,omitempty"`
	/*
	  Retention period in seconds for the generated reports.
	*/
	RetentionPeriodSeconds *int64 `json:"retentionPeriodSeconds,omitempty"`
}

func (p *RetentionConfig) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias RetentionConfig

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

func (p *RetentionConfig) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias RetentionConfig
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewRetentionConfig()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.RetentionCount != nil {
		p.RetentionCount = known.RetentionCount
	}
	if known.RetentionPeriodSeconds != nil {
		p.RetentionPeriodSeconds = known.RetentionPeriodSeconds
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "retentionCount")
	delete(allFields, "retentionPeriodSeconds")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewRetentionConfig() *RetentionConfig {
	p := new(RetentionConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.config.RetentionConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
A row is a list of maximum 3 widgets.
*/
type Row struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of widgets in a row.
	*/
	Widgets []Widget `json:"widgets"`
}

func (p *Row) MarshalJSON() ([]byte, error) {
	type RowProxy Row

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*RowProxy
		Widgets []Widget `json:"widgets,omitempty"`
	}{
		RowProxy: (*RowProxy)(p),
		Widgets:  p.Widgets,
	}

	known, err := json.Marshal(baseStruct)
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

func (p *Row) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Row
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewRow()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Widgets != nil {
		p.Widgets = known.Widgets
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "widgets")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewRow() *Row {
	p := new(Row)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.config.Row"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Interval at which report generation should be scheduled.
*/
type ScheduleInterval int

const (
	SCHEDULEINTERVAL_UNKNOWN  ScheduleInterval = 0
	SCHEDULEINTERVAL_REDACTED ScheduleInterval = 1
	SCHEDULEINTERVAL_NONE     ScheduleInterval = 2
	SCHEDULEINTERVAL_DAILY    ScheduleInterval = 3
	SCHEDULEINTERVAL_WEEKLY   ScheduleInterval = 4
	SCHEDULEINTERVAL_MONTHLY  ScheduleInterval = 5
	SCHEDULEINTERVAL_YEARLY   ScheduleInterval = 6
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ScheduleInterval) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NONE",
		"DAILY",
		"WEEKLY",
		"MONTHLY",
		"YEARLY",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ScheduleInterval) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NONE",
		"DAILY",
		"WEEKLY",
		"MONTHLY",
		"YEARLY",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ScheduleInterval) index(name string) ScheduleInterval {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NONE",
		"DAILY",
		"WEEKLY",
		"MONTHLY",
		"YEARLY",
	}
	for idx := range names {
		if names[idx] == name {
			return ScheduleInterval(idx)
		}
	}
	return SCHEDULEINTERVAL_UNKNOWN
}

func (e *ScheduleInterval) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ScheduleInterval:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ScheduleInterval) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ScheduleInterval) Ref() *ScheduleInterval {
	return &e
}

/*
A section is a group of rows consisting of widgets.
*/
type Section struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Description of the section.
	*/
	Description *string `json:"description,omitempty"`
	/*
	  Name of the section.
	*/
	Name *string `json:"name"`

	RepeatCriteria *RepeatCriteria `json:"repeatCriteria,omitempty"`
	/*
	  List of row in the section.
	*/
	Rows []Row `json:"rows"`

	TimeFilter *TimeFilter `json:"timeFilter,omitempty"`
}

func (p *Section) MarshalJSON() ([]byte, error) {
	type SectionProxy Section

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*SectionProxy
		Name *string `json:"name,omitempty"`
		Rows []Row   `json:"rows,omitempty"`
	}{
		SectionProxy: (*SectionProxy)(p),
		Name:         p.Name,
		Rows:         p.Rows,
	}

	known, err := json.Marshal(baseStruct)
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

func (p *Section) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Section
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewSection()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Description != nil {
		p.Description = known.Description
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.RepeatCriteria != nil {
		p.RepeatCriteria = known.RepeatCriteria
	}
	if known.Rows != nil {
		p.Rows = known.Rows
	}
	if known.TimeFilter != nil {
		p.TimeFilter = known.TimeFilter
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "description")
	delete(allFields, "name")
	delete(allFields, "repeatCriteria")
	delete(allFields, "rows")
	delete(allFields, "timeFilter")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewSection() *Section {
	p := new(Section)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.config.Section"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Aggregation to be used while sorting time-series data.
*/
type SortKey int

const (
	SORTKEY_UNKNOWN  SortKey = 0
	SORTKEY_REDACTED SortKey = 1
	SORTKEY_MAX      SortKey = 2
	SORTKEY_MIN      SortKey = 3
	SORTKEY_FIRST    SortKey = 4
	SORTKEY_LAST     SortKey = 5
	SORTKEY_LATEST   SortKey = 6
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *SortKey) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"MAX",
		"MIN",
		"FIRST",
		"LAST",
		"LATEST",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e SortKey) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"MAX",
		"MIN",
		"FIRST",
		"LAST",
		"LATEST",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *SortKey) index(name string) SortKey {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"MAX",
		"MIN",
		"FIRST",
		"LAST",
		"LATEST",
	}
	for idx := range names {
		if names[idx] == name {
			return SortKey(idx)
		}
	}
	return SORTKEY_UNKNOWN
}

func (e *SortKey) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for SortKey:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *SortKey) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e SortKey) Ref() *SortKey {
	return &e
}

/*
Order of sorting.
*/
type SortOrder int

const (
	SORTORDER_UNKNOWN    SortOrder = 0
	SORTORDER_REDACTED   SortOrder = 1
	SORTORDER_ASCENDING  SortOrder = 2
	SORTORDER_DESCENDING SortOrder = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *SortOrder) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ASCENDING",
		"DESCENDING",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e SortOrder) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ASCENDING",
		"DESCENDING",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *SortOrder) index(name string) SortOrder {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ASCENDING",
		"DESCENDING",
	}
	for idx := range names {
		if names[idx] == name {
			return SortOrder(idx)
		}
	}
	return SORTORDER_UNKNOWN
}

func (e *SortOrder) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for SortOrder:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *SortOrder) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e SortOrder) Ref() *SortOrder {
	return &e
}

/*
Threshold configuration.
*/
type Threshold struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Color of the threshold.
	*/
	Color *string `json:"color"`
	/*
	  End range of the threshold.
	*/
	EndRange *float64 `json:"endRange"`
	/*
	  Start range of the threshold.
	*/
	StartRange *float64 `json:"startRange"`
}

func (p *Threshold) MarshalJSON() ([]byte, error) {
	type ThresholdProxy Threshold

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*ThresholdProxy
		Color      *string  `json:"color,omitempty"`
		EndRange   *float64 `json:"endRange,omitempty"`
		StartRange *float64 `json:"startRange,omitempty"`
	}{
		ThresholdProxy: (*ThresholdProxy)(p),
		Color:          p.Color,
		EndRange:       p.EndRange,
		StartRange:     p.StartRange,
	}

	known, err := json.Marshal(baseStruct)
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

func (p *Threshold) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Threshold
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewThreshold()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Color != nil {
		p.Color = known.Color
	}
	if known.EndRange != nil {
		p.EndRange = known.EndRange
	}
	if known.StartRange != nil {
		p.StartRange = known.StartRange
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "color")
	delete(allFields, "endRange")
	delete(allFields, "startRange")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewThreshold() *Threshold {
	p := new(Threshold)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.config.Threshold"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Time based filtering, can be done on created and last occured timestamp.
*/
type TimeFilter int

const (
	TIMEFILTER_UNKNOWN            TimeFilter = 0
	TIMEFILTER_REDACTED           TimeFilter = 1
	TIMEFILTER_CREATED_TIME       TimeFilter = 2
	TIMEFILTER_LAST_OCCURRED_TIME TimeFilter = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *TimeFilter) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CREATED_TIME",
		"LAST_OCCURRED_TIME",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e TimeFilter) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CREATED_TIME",
		"LAST_OCCURRED_TIME",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *TimeFilter) index(name string) TimeFilter {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CREATED_TIME",
		"LAST_OCCURRED_TIME",
	}
	for idx := range names {
		if names[idx] == name {
			return TimeFilter(idx)
		}
	}
	return TIMEFILTER_UNKNOWN
}

func (e *TimeFilter) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for TimeFilter:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *TimeFilter) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e TimeFilter) Ref() *TimeFilter {
	return &e
}

/*
Unit of the metric data in the report.
*/
type Unit int

const (
	UNIT_UNKNOWN              Unit = 0
	UNIT_REDACTED             Unit = 1
	UNIT_BYTES                Unit = 2
	UNIT_KIBIBYTES            Unit = 3
	UNIT_MEBIBYTES            Unit = 4
	UNIT_GIBIBYTES            Unit = 5
	UNIT_TEBIBYTES            Unit = 6
	UNIT_PEBIBYTES            Unit = 7
	UNIT_PERCENT              Unit = 8
	UNIT_HERTZ                Unit = 9
	UNIT_KILOHERTZ            Unit = 10
	UNIT_MEGAHERTZ            Unit = 11
	UNIT_GIGAHERTZ            Unit = 12
	UNIT_MICROSECONDS         Unit = 13
	UNIT_MILLISECONDS         Unit = 14
	UNIT_SECONDS              Unit = 15
	UNIT_KILOBYTES_PER_SECOND Unit = 16
	UNIT_MEGABYTES_PER_SECOND Unit = 17
	UNIT_GIGABYTES_PER_SECOND Unit = 18
	UNIT_ABSOLUTE             Unit = 19
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *Unit) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"BYTES",
		"KIBIBYTES",
		"MEBIBYTES",
		"GIBIBYTES",
		"TEBIBYTES",
		"PEBIBYTES",
		"PERCENT",
		"HERTZ",
		"KILOHERTZ",
		"MEGAHERTZ",
		"GIGAHERTZ",
		"MICROSECONDS",
		"MILLISECONDS",
		"SECONDS",
		"KILOBYTES_PER_SECOND",
		"MEGABYTES_PER_SECOND",
		"GIGABYTES_PER_SECOND",
		"ABSOLUTE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e Unit) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"BYTES",
		"KIBIBYTES",
		"MEBIBYTES",
		"GIBIBYTES",
		"TEBIBYTES",
		"PEBIBYTES",
		"PERCENT",
		"HERTZ",
		"KILOHERTZ",
		"MEGAHERTZ",
		"GIGAHERTZ",
		"MICROSECONDS",
		"MILLISECONDS",
		"SECONDS",
		"KILOBYTES_PER_SECOND",
		"MEGABYTES_PER_SECOND",
		"GIGABYTES_PER_SECOND",
		"ABSOLUTE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *Unit) index(name string) Unit {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"BYTES",
		"KIBIBYTES",
		"MEBIBYTES",
		"GIBIBYTES",
		"TEBIBYTES",
		"PEBIBYTES",
		"PERCENT",
		"HERTZ",
		"KILOHERTZ",
		"MEGAHERTZ",
		"GIGAHERTZ",
		"MICROSECONDS",
		"MILLISECONDS",
		"SECONDS",
		"KILOBYTES_PER_SECOND",
		"MEGABYTES_PER_SECOND",
		"GIGABYTES_PER_SECOND",
		"ABSOLUTE",
	}
	for idx := range names {
		if names[idx] == name {
			return Unit(idx)
		}
	}
	return UNIT_UNKNOWN
}

func (e *Unit) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for Unit:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *Unit) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e Unit) Ref() *Unit {
	return &e
}

/*
REST response for all response codes in API path /opsmgmt/v4.1.b1/config/dashboards/{extId} Put operation
*/
type UpdateDashboardApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateDashboardApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *UpdateDashboardApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias UpdateDashboardApiResponse

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

func (p *UpdateDashboardApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias UpdateDashboardApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewUpdateDashboardApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewUpdateDashboardApiResponse() *UpdateDashboardApiResponse {
	p := new(UpdateDashboardApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.config.UpdateDashboardApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdateDashboardApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdateDashboardApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdateDashboardApiResponseData()
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
REST response for all response codes in API path /opsmgmt/v4.1.b1/config/user/{userExtId}/global-report-setting Put operation
*/
type UpdateGlobalReportSettingsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateGlobalReportSettingsApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *UpdateGlobalReportSettingsApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias UpdateGlobalReportSettingsApiResponse

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

func (p *UpdateGlobalReportSettingsApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias UpdateGlobalReportSettingsApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewUpdateGlobalReportSettingsApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewUpdateGlobalReportSettingsApiResponse() *UpdateGlobalReportSettingsApiResponse {
	p := new(UpdateGlobalReportSettingsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.config.UpdateGlobalReportSettingsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdateGlobalReportSettingsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdateGlobalReportSettingsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdateGlobalReportSettingsApiResponseData()
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
REST response for all response codes in API path /opsmgmt/v4.1.b1/config/report-configs/{extId} Put operation
*/
type UpdateReportConfigApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateReportConfigApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *UpdateReportConfigApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias UpdateReportConfigApiResponse

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

func (p *UpdateReportConfigApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias UpdateReportConfigApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewUpdateReportConfigApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewUpdateReportConfigApiResponse() *UpdateReportConfigApiResponse {
	p := new(UpdateReportConfigApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.config.UpdateReportConfigApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdateReportConfigApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdateReportConfigApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdateReportConfigApiResponseData()
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
REST response for all response codes in API path /opsmgmt/v4.1.b1/config/dashboard-settings/$actions/upload-geoconfigurations Post operation
*/
type UploadDashboardGeoconfigurationApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUploadDashboardGeoconfigurationApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *UploadDashboardGeoconfigurationApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias UploadDashboardGeoconfigurationApiResponse

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

func (p *UploadDashboardGeoconfigurationApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias UploadDashboardGeoconfigurationApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewUploadDashboardGeoconfigurationApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewUploadDashboardGeoconfigurationApiResponse() *UploadDashboardGeoconfigurationApiResponse {
	p := new(UploadDashboardGeoconfigurationApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.config.UploadDashboardGeoconfigurationApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UploadDashboardGeoconfigurationApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UploadDashboardGeoconfigurationApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUploadDashboardGeoconfigurationApiResponseData()
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
Element that displays information in the report.
*/
type Widget struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	WidgetInfoItemDiscriminator_ *string `json:"$widgetInfoItemDiscriminator,omitempty"`
	/*
	  This describes the details of the widget.
	*/
	WidgetInfo *OneOfWidgetWidgetInfo `json:"widgetInfo"`
}

func (p *Widget) MarshalJSON() ([]byte, error) {
	type WidgetProxy Widget

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*WidgetProxy
		WidgetInfo *OneOfWidgetWidgetInfo `json:"widgetInfo,omitempty"`
	}{
		WidgetProxy: (*WidgetProxy)(p),
		WidgetInfo:  p.WidgetInfo,
	}

	known, err := json.Marshal(baseStruct)
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

func (p *Widget) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Widget
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewWidget()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.WidgetInfoItemDiscriminator_ != nil {
		p.WidgetInfoItemDiscriminator_ = known.WidgetInfoItemDiscriminator_
	}
	if known.WidgetInfo != nil {
		p.WidgetInfo = known.WidgetInfo
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$widgetInfoItemDiscriminator")
	delete(allFields, "widgetInfo")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewWidget() *Widget {
	p := new(Widget)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.config.Widget"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *Widget) GetWidgetInfo() interface{} {
	if nil == p.WidgetInfo {
		return nil
	}
	return p.WidgetInfo.GetValue()
}

func (p *Widget) SetWidgetInfo(v interface{}) error {
	if nil == p.WidgetInfo {
		p.WidgetInfo = NewOneOfWidgetWidgetInfo()
	}
	e := p.WidgetInfo.SetValue(v)
	if nil == e {
		if nil == p.WidgetInfoItemDiscriminator_ {
			p.WidgetInfoItemDiscriminator_ = new(string)
		}
		*p.WidgetInfoItemDiscriminator_ = *p.WidgetInfo.Discriminator
	}
	return e
}

/*
Configuration of the widget.
*/
type WidgetConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	DataCriteria *DataCriteria `json:"dataCriteria,omitempty"`
	/*
	  Description of the widget.
	*/
	Description *string `json:"description,omitempty"`

	EntityType *EntityType `json:"entityType,omitempty"`
	/*
	  List of selected entity attributes/metrics for the widget.
	*/
	Fields []WidgetField `json:"fields,omitempty"`
	/*
	  Heading for a widget.
	*/
	Heading *string `json:"heading,omitempty"`

	RepeatCriteria *RepeatCriteria `json:"repeatCriteria,omitempty"`

	Size *WidgetSize `json:"size"`

	TimeFilter *TimeFilter `json:"timeFilter,omitempty"`

	Type *WidgetType `json:"type"`
	/*
	  Widget type specific configurations.
	*/
	TypeSpecificConfigs []import4.KVPair `json:"typeSpecificConfigs,omitempty"`
	/*
	  Identifier for a predefined widget.
	*/
	WidgetId *string `json:"widgetId,omitempty"`
}

func (p *WidgetConfig) MarshalJSON() ([]byte, error) {
	type WidgetConfigProxy WidgetConfig

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*WidgetConfigProxy
		Size *WidgetSize `json:"size,omitempty"`
		Type *WidgetType `json:"type,omitempty"`
	}{
		WidgetConfigProxy: (*WidgetConfigProxy)(p),
		Size:              p.Size,
		Type:              p.Type,
	}

	known, err := json.Marshal(baseStruct)
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

func (p *WidgetConfig) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias WidgetConfig
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewWidgetConfig()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataCriteria != nil {
		p.DataCriteria = known.DataCriteria
	}
	if known.Description != nil {
		p.Description = known.Description
	}
	if known.EntityType != nil {
		p.EntityType = known.EntityType
	}
	if known.Fields != nil {
		p.Fields = known.Fields
	}
	if known.Heading != nil {
		p.Heading = known.Heading
	}
	if known.RepeatCriteria != nil {
		p.RepeatCriteria = known.RepeatCriteria
	}
	if known.Size != nil {
		p.Size = known.Size
	}
	if known.TimeFilter != nil {
		p.TimeFilter = known.TimeFilter
	}
	if known.Type != nil {
		p.Type = known.Type
	}
	if known.TypeSpecificConfigs != nil {
		p.TypeSpecificConfigs = known.TypeSpecificConfigs
	}
	if known.WidgetId != nil {
		p.WidgetId = known.WidgetId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "dataCriteria")
	delete(allFields, "description")
	delete(allFields, "entityType")
	delete(allFields, "fields")
	delete(allFields, "heading")
	delete(allFields, "repeatCriteria")
	delete(allFields, "size")
	delete(allFields, "timeFilter")
	delete(allFields, "type")
	delete(allFields, "typeSpecificConfigs")
	delete(allFields, "widgetId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewWidgetConfig() *WidgetConfig {
	p := new(WidgetConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.config.WidgetConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Selected entity attribute/metric for the widget.
*/
type WidgetField struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	AggregateFunction *AggregateFunction `json:"aggregateFunction,omitempty"`

	CompoundMetric *CompoundMetric `json:"compoundMetric,omitempty"`
	/*
	  Human-readable label for widget field.
	*/
	Label *string `json:"label"`
	/*
	  Entity attribute/metric to be selected for the widget.
	*/
	Name *string `json:"name"`
	/*
	  List of thresholds.
	*/
	Thresholds []Threshold `json:"thresholds,omitempty"`

	Unit *Unit `json:"unit,omitempty"`
}

func (p *WidgetField) MarshalJSON() ([]byte, error) {
	type WidgetFieldProxy WidgetField

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*WidgetFieldProxy
		Label *string `json:"label,omitempty"`
		Name  *string `json:"name,omitempty"`
	}{
		WidgetFieldProxy: (*WidgetFieldProxy)(p),
		Label:            p.Label,
		Name:             p.Name,
	}

	known, err := json.Marshal(baseStruct)
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

func (p *WidgetField) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias WidgetField
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewWidgetField()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.AggregateFunction != nil {
		p.AggregateFunction = known.AggregateFunction
	}
	if known.CompoundMetric != nil {
		p.CompoundMetric = known.CompoundMetric
	}
	if known.Label != nil {
		p.Label = known.Label
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.Thresholds != nil {
		p.Thresholds = known.Thresholds
	}
	if known.Unit != nil {
		p.Unit = known.Unit
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "aggregateFunction")
	delete(allFields, "compoundMetric")
	delete(allFields, "label")
	delete(allFields, "name")
	delete(allFields, "thresholds")
	delete(allFields, "unit")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewWidgetField() *WidgetField {
	p := new(WidgetField)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.config.WidgetField"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Relation of the widget in the dashboard.
*/
type WidgetRelation struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  ID of the child widget in the dashboard.
	*/
	ChildWidgetId *string `json:"childWidgetId,omitempty"`
	/*
	  ID of the parent widget in the dashboard.
	*/
	ParentWidgetId *string `json:"parentWidgetId,omitempty"`
	/*
	  Attribute of the relation b/w two widgets in the dashboard.
	*/
	RelationMappings []import4.KVPair `json:"relationMappings,omitempty"`
}

func (p *WidgetRelation) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias WidgetRelation

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

func (p *WidgetRelation) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias WidgetRelation
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewWidgetRelation()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ChildWidgetId != nil {
		p.ChildWidgetId = known.ChildWidgetId
	}
	if known.ParentWidgetId != nil {
		p.ParentWidgetId = known.ParentWidgetId
	}
	if known.RelationMappings != nil {
		p.RelationMappings = known.RelationMappings
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "childWidgetId")
	delete(allFields, "parentWidgetId")
	delete(allFields, "relationMappings")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewWidgetRelation() *WidgetRelation {
	p := new(WidgetRelation)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.config.WidgetRelation"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Size of the widget.
*/
type WidgetSize int

const (
	WIDGETSIZE_UNKNOWN  WidgetSize = 0
	WIDGETSIZE_REDACTED WidgetSize = 1
	WIDGETSIZE_SMALL    WidgetSize = 2
	WIDGETSIZE_LARGE    WidgetSize = 3
	WIDGETSIZE_FULLSPAN WidgetSize = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *WidgetSize) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SMALL",
		"LARGE",
		"FULLSPAN",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e WidgetSize) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SMALL",
		"LARGE",
		"FULLSPAN",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *WidgetSize) index(name string) WidgetSize {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SMALL",
		"LARGE",
		"FULLSPAN",
	}
	for idx := range names {
		if names[idx] == name {
			return WidgetSize(idx)
		}
	}
	return WIDGETSIZE_UNKNOWN
}

func (e *WidgetSize) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for WidgetSize:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *WidgetSize) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e WidgetSize) Ref() *WidgetSize {
	return &e
}

/*
Predefined widget.
*/
type WidgetTemplate struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	WidgetTemplate *WidgetTemplateType `json:"widgetTemplate,omitempty"`
}

func (p *WidgetTemplate) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias WidgetTemplate

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

func (p *WidgetTemplate) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias WidgetTemplate
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewWidgetTemplate()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.WidgetTemplate != nil {
		p.WidgetTemplate = known.WidgetTemplate
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "widgetTemplate")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewWidgetTemplate() *WidgetTemplate {
	p := new(WidgetTemplate)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.config.WidgetTemplate"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
List of predefined widgets supported.
*/
type WidgetTemplateType int

const (
	WIDGETTEMPLATETYPE_UNKNOWN                           WidgetTemplateType = 0
	WIDGETTEMPLATETYPE_REDACTED                          WidgetTemplateType = 1
	WIDGETTEMPLATETYPE_CLUSTER_LICENSE_TABLE             WidgetTemplateType = 2
	WIDGETTEMPLATETYPE_MULTICLUSTER_LICENSE_SUMMARY      WidgetTemplateType = 3
	WIDGETTEMPLATETYPE_ALERTS_HISTOGRAM                  WidgetTemplateType = 4
	WIDGETTEMPLATETYPE_BLOCKS_SUMMARY                    WidgetTemplateType = 5
	WIDGETTEMPLATETYPE_IGNORE_TIME_WINDOW_TABLE          WidgetTemplateType = 6
	WIDGETTEMPLATETYPE_CLUSTER_CPU_RUNWAY_CHART          WidgetTemplateType = 7
	WIDGETTEMPLATETYPE_CLUSTER_MEMORY_RUNWAY_CHART       WidgetTemplateType = 8
	WIDGETTEMPLATETYPE_CLUSTER_STORAGE_RUNWAY_CHART      WidgetTemplateType = 9
	WIDGETTEMPLATETYPE_CLUSTER_INACTIVE_VM_TABLE         WidgetTemplateType = 10
	WIDGETTEMPLATETYPE_CLUSTER_CONSTRAINED_VM_TABLE      WidgetTemplateType = 11
	WIDGETTEMPLATETYPE_CLUSTER_OVERPROVISIONED_VM_TABLE  WidgetTemplateType = 12
	WIDGETTEMPLATETYPE_CLUSTER_BULLY_VM_TABLE            WidgetTemplateType = 13
	WIDGETTEMPLATETYPE_CLUSTER_POTENTIAL_CPU_RECLAIM     WidgetTemplateType = 14
	WIDGETTEMPLATETYPE_CLUSTER_POTENTIAL_MEMORY_RECLAIM  WidgetTemplateType = 15
	WIDGETTEMPLATETYPE_CLUSTER_POTENTIAL_STORAGE_RECLAIM WidgetTemplateType = 16
	WIDGETTEMPLATETYPE_VCENTER_BLOCKS_SUMMARY            WidgetTemplateType = 17
	WIDGETTEMPLATETYPE_VCENTER_CPU_RUNWAY_CHART          WidgetTemplateType = 18
	WIDGETTEMPLATETYPE_VCENTER_MEMORY_RUNWAY_CHART       WidgetTemplateType = 19
	WIDGETTEMPLATETYPE_VCENTER_INACTIVE_VM_TABLE         WidgetTemplateType = 20
	WIDGETTEMPLATETYPE_VCENTER_CONSTRAINED_VM_TABLE      WidgetTemplateType = 21
	WIDGETTEMPLATETYPE_VCENTER_OVERPROVISIONED_VM_TABLE  WidgetTemplateType = 22
	WIDGETTEMPLATETYPE_VCENTER_BULLY_VM_TABLE            WidgetTemplateType = 23
	WIDGETTEMPLATETYPE_VCENTER_POTENTIAL_CPU_RECLAIM     WidgetTemplateType = 24
	WIDGETTEMPLATETYPE_VCENTER_POTENTIAL_MEMORY_RECLAIM  WidgetTemplateType = 25
	WIDGETTEMPLATETYPE_VCENTER_POTENTIAL_STORAGE_RECLAIM WidgetTemplateType = 26
	WIDGETTEMPLATETYPE_ACCOUNT_OVERVIEW                  WidgetTemplateType = 27
	WIDGETTEMPLATETYPE_ALERTS_TIMELINE                   WidgetTemplateType = 28
	WIDGETTEMPLATETYPE_ALERTS_COUNT                      WidgetTemplateType = 29
	WIDGETTEMPLATETYPE_ANOMALIES_COUNT                   WidgetTemplateType = 30
	WIDGETTEMPLATETYPE_IMPACTED_CLUSTER                  WidgetTemplateType = 31
	WIDGETTEMPLATETYPE_VM_EFFICIENCY                     WidgetTemplateType = 32
	WIDGETTEMPLATETYPE_CLUSTER_RUNWAY                    WidgetTemplateType = 33
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *WidgetTemplateType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CLUSTER_LICENSE_TABLE",
		"MULTICLUSTER_LICENSE_SUMMARY",
		"ALERTS_HISTOGRAM",
		"BLOCKS_SUMMARY",
		"IGNORE_TIME_WINDOW_TABLE",
		"CLUSTER_CPU_RUNWAY_CHART",
		"CLUSTER_MEMORY_RUNWAY_CHART",
		"CLUSTER_STORAGE_RUNWAY_CHART",
		"CLUSTER_INACTIVE_VM_TABLE",
		"CLUSTER_CONSTRAINED_VM_TABLE",
		"CLUSTER_OVERPROVISIONED_VM_TABLE",
		"CLUSTER_BULLY_VM_TABLE",
		"CLUSTER_POTENTIAL_CPU_RECLAIM",
		"CLUSTER_POTENTIAL_MEMORY_RECLAIM",
		"CLUSTER_POTENTIAL_STORAGE_RECLAIM",
		"VCENTER_BLOCKS_SUMMARY",
		"VCENTER_CPU_RUNWAY_CHART",
		"VCENTER_MEMORY_RUNWAY_CHART",
		"VCENTER_INACTIVE_VM_TABLE",
		"VCENTER_CONSTRAINED_VM_TABLE",
		"VCENTER_OVERPROVISIONED_VM_TABLE",
		"VCENTER_BULLY_VM_TABLE",
		"VCENTER_POTENTIAL_CPU_RECLAIM",
		"VCENTER_POTENTIAL_MEMORY_RECLAIM",
		"VCENTER_POTENTIAL_STORAGE_RECLAIM",
		"ACCOUNT_OVERVIEW",
		"ALERTS_TIMELINE",
		"ALERTS_COUNT",
		"ANOMALIES_COUNT",
		"IMPACTED_CLUSTER",
		"VM_EFFICIENCY",
		"CLUSTER_RUNWAY",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e WidgetTemplateType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CLUSTER_LICENSE_TABLE",
		"MULTICLUSTER_LICENSE_SUMMARY",
		"ALERTS_HISTOGRAM",
		"BLOCKS_SUMMARY",
		"IGNORE_TIME_WINDOW_TABLE",
		"CLUSTER_CPU_RUNWAY_CHART",
		"CLUSTER_MEMORY_RUNWAY_CHART",
		"CLUSTER_STORAGE_RUNWAY_CHART",
		"CLUSTER_INACTIVE_VM_TABLE",
		"CLUSTER_CONSTRAINED_VM_TABLE",
		"CLUSTER_OVERPROVISIONED_VM_TABLE",
		"CLUSTER_BULLY_VM_TABLE",
		"CLUSTER_POTENTIAL_CPU_RECLAIM",
		"CLUSTER_POTENTIAL_MEMORY_RECLAIM",
		"CLUSTER_POTENTIAL_STORAGE_RECLAIM",
		"VCENTER_BLOCKS_SUMMARY",
		"VCENTER_CPU_RUNWAY_CHART",
		"VCENTER_MEMORY_RUNWAY_CHART",
		"VCENTER_INACTIVE_VM_TABLE",
		"VCENTER_CONSTRAINED_VM_TABLE",
		"VCENTER_OVERPROVISIONED_VM_TABLE",
		"VCENTER_BULLY_VM_TABLE",
		"VCENTER_POTENTIAL_CPU_RECLAIM",
		"VCENTER_POTENTIAL_MEMORY_RECLAIM",
		"VCENTER_POTENTIAL_STORAGE_RECLAIM",
		"ACCOUNT_OVERVIEW",
		"ALERTS_TIMELINE",
		"ALERTS_COUNT",
		"ANOMALIES_COUNT",
		"IMPACTED_CLUSTER",
		"VM_EFFICIENCY",
		"CLUSTER_RUNWAY",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *WidgetTemplateType) index(name string) WidgetTemplateType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CLUSTER_LICENSE_TABLE",
		"MULTICLUSTER_LICENSE_SUMMARY",
		"ALERTS_HISTOGRAM",
		"BLOCKS_SUMMARY",
		"IGNORE_TIME_WINDOW_TABLE",
		"CLUSTER_CPU_RUNWAY_CHART",
		"CLUSTER_MEMORY_RUNWAY_CHART",
		"CLUSTER_STORAGE_RUNWAY_CHART",
		"CLUSTER_INACTIVE_VM_TABLE",
		"CLUSTER_CONSTRAINED_VM_TABLE",
		"CLUSTER_OVERPROVISIONED_VM_TABLE",
		"CLUSTER_BULLY_VM_TABLE",
		"CLUSTER_POTENTIAL_CPU_RECLAIM",
		"CLUSTER_POTENTIAL_MEMORY_RECLAIM",
		"CLUSTER_POTENTIAL_STORAGE_RECLAIM",
		"VCENTER_BLOCKS_SUMMARY",
		"VCENTER_CPU_RUNWAY_CHART",
		"VCENTER_MEMORY_RUNWAY_CHART",
		"VCENTER_INACTIVE_VM_TABLE",
		"VCENTER_CONSTRAINED_VM_TABLE",
		"VCENTER_OVERPROVISIONED_VM_TABLE",
		"VCENTER_BULLY_VM_TABLE",
		"VCENTER_POTENTIAL_CPU_RECLAIM",
		"VCENTER_POTENTIAL_MEMORY_RECLAIM",
		"VCENTER_POTENTIAL_STORAGE_RECLAIM",
		"ACCOUNT_OVERVIEW",
		"ALERTS_TIMELINE",
		"ALERTS_COUNT",
		"ANOMALIES_COUNT",
		"IMPACTED_CLUSTER",
		"VM_EFFICIENCY",
		"CLUSTER_RUNWAY",
	}
	for idx := range names {
		if names[idx] == name {
			return WidgetTemplateType(idx)
		}
	}
	return WIDGETTEMPLATETYPE_UNKNOWN
}

func (e *WidgetTemplateType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for WidgetTemplateType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *WidgetTemplateType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e WidgetTemplateType) Ref() *WidgetTemplateType {
	return &e
}

/*
Type of widget configuration.
*/
type WidgetType int

const (
	WIDGETTYPE_UNKNOWN              WidgetType = 0
	WIDGETTYPE_REDACTED             WidgetType = 1
	WIDGETTYPE_BAR_CHART            WidgetType = 2
	WIDGETTYPE_LINE_CHART           WidgetType = 3
	WIDGETTYPE_HISTOGRAM            WidgetType = 4
	WIDGETTYPE_DATA_TABLE           WidgetType = 5
	WIDGETTYPE_CONFIG_SUMMARY       WidgetType = 6
	WIDGETTYPE_METRIC_SUMMARY_TEXT  WidgetType = 7
	WIDGETTYPE_METRIC_SUMMARY_CHART WidgetType = 8
	WIDGETTYPE_COUNT_SUMMARY        WidgetType = 9
	WIDGETTYPE_TEXT                 WidgetType = 10
	WIDGETTYPE_STATS_SUMMARY        WidgetType = 11
	WIDGETTYPE_PIE_CHART            WidgetType = 12
	WIDGETTYPE_GEO_MAP              WidgetType = 13
	WIDGETTYPE_ALERT_LIST           WidgetType = 14
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *WidgetType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"BAR_CHART",
		"LINE_CHART",
		"HISTOGRAM",
		"DATA_TABLE",
		"CONFIG_SUMMARY",
		"METRIC_SUMMARY_TEXT",
		"METRIC_SUMMARY_CHART",
		"COUNT_SUMMARY",
		"TEXT",
		"STATS_SUMMARY",
		"PIE_CHART",
		"GEO_MAP",
		"ALERT_LIST",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e WidgetType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"BAR_CHART",
		"LINE_CHART",
		"HISTOGRAM",
		"DATA_TABLE",
		"CONFIG_SUMMARY",
		"METRIC_SUMMARY_TEXT",
		"METRIC_SUMMARY_CHART",
		"COUNT_SUMMARY",
		"TEXT",
		"STATS_SUMMARY",
		"PIE_CHART",
		"GEO_MAP",
		"ALERT_LIST",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *WidgetType) index(name string) WidgetType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"BAR_CHART",
		"LINE_CHART",
		"HISTOGRAM",
		"DATA_TABLE",
		"CONFIG_SUMMARY",
		"METRIC_SUMMARY_TEXT",
		"METRIC_SUMMARY_CHART",
		"COUNT_SUMMARY",
		"TEXT",
		"STATS_SUMMARY",
		"PIE_CHART",
		"GEO_MAP",
		"ALERT_LIST",
	}
	for idx := range names {
		if names[idx] == name {
			return WidgetType(idx)
		}
	}
	return WIDGETTYPE_UNKNOWN
}

func (e *WidgetType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for WidgetType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *WidgetType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e WidgetType) Ref() *WidgetType {
	return &e
}

type OneOfNodeNodeInfo struct {
	Discriminator *string `json:"-"`
	ObjectType_   *string `json:"-"`
	oneOfType1001 *Layout `json:"-"`
	oneOfType1002 *Widget `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfNodeNodeInfo() *OneOfNodeNodeInfo {
	p := new(OneOfNodeNodeInfo)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfNodeNodeInfo) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfNodeNodeInfo is nil"))
	}
	switch v.(type) {
	case Layout:
		if nil == p.oneOfType1001 {
			p.oneOfType1001 = new(Layout)
		}
		*p.oneOfType1001 = v.(Layout)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType1001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType1001.ObjectType_
	case Widget:
		if nil == p.oneOfType1002 {
			p.oneOfType1002 = new(Widget)
		}
		*p.oneOfType1002 = v.(Widget)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType1002.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType1002.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfNodeNodeInfo) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType1001 != nil && *p.oneOfType1001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1001
	}
	if p.oneOfType1002 != nil && *p.oneOfType1002.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1002
	}
	return nil
}

func (p *OneOfNodeNodeInfo) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType1001 := new(Layout)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType1001)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType1001.ObjectType_ != nil && "opsmgmt.v4.config.Layout" == *vOneOfType1001.ObjectType_ {
							if nil == p.oneOfType1001 {
								p.oneOfType1001 = new(Layout)
							}
							*p.oneOfType1001 = *vOneOfType1001
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType1001.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType1001.ObjectType_
							return nil
						}
					}
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType1002 := new(Widget)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType1002)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType1002.ObjectType_ != nil && "opsmgmt.v4.config.Widget" == *vOneOfType1002.ObjectType_ {
							if nil == p.oneOfType1002 {
								p.oneOfType1002 = new(Widget)
							}
							*p.oneOfType1002 = *vOneOfType1002
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType1002.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType1002.ObjectType_
							return nil
						}
					}
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType1001 := new(Layout)
	if err := json.Unmarshal(b, vOneOfType1001); err == nil {
		if vOneOfType1001.ObjectType_ != nil && "opsmgmt.v4.config.Layout" == *vOneOfType1001.ObjectType_ {
			if nil == p.oneOfType1001 {
				p.oneOfType1001 = new(Layout)
			}
			*p.oneOfType1001 = *vOneOfType1001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType1001.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType1001.ObjectType_
			return nil
		}
	}
	vOneOfType1002 := new(Widget)
	if err := json.Unmarshal(b, vOneOfType1002); err == nil {
		if vOneOfType1002.ObjectType_ != nil && "opsmgmt.v4.config.Widget" == *vOneOfType1002.ObjectType_ {
			if nil == p.oneOfType1002 {
				p.oneOfType1002 = new(Widget)
			}
			*p.oneOfType1002 = *vOneOfType1002
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType1002.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType1002.ObjectType_
			return nil
		}
	}
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfNodeNodeInfo"))
}

func (p *OneOfNodeNodeInfo) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType1001 != nil && *p.oneOfType1001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1001)
	}
	if p.oneOfType1002 != nil && *p.oneOfType1002.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1002)
	}
	return nil, errors.New("No value to marshal for OneOfNodeNodeInfo")
}

type OneOfUpdateDashboardApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 []import1.AppMessage   `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfUpdateDashboardApiResponseData() *OneOfUpdateDashboardApiResponseData {
	p := new(OneOfUpdateDashboardApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateDashboardApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateDashboardApiResponseData is nil"))
	}
	switch v.(type) {
	case []import1.AppMessage:
		p.oneOfType2001 = v.([]import1.AppMessage)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<opsmgmt.v4.error.AppMessage>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<opsmgmt.v4.error.AppMessage>"
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfUpdateDashboardApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if "List<opsmgmt.v4.error.AppMessage>" == *p.Discriminator {
		return p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfUpdateDashboardApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["List<opsmgmt.v4.error.AppMessage>"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType2001 := new([]import1.AppMessage)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType2001)
					if unmarshalErr == nil {
						// For arrays, verify the array item ObjectType matches
						if vOneOfType2001 == nil || len(*vOneOfType2001) == 0 || ((*vOneOfType2001)[0].ObjectType_ != nil && "opsmgmt.v4.error.AppMessage" == *((*vOneOfType2001)[0].ObjectType_)) {
							p.oneOfType2001 = *vOneOfType2001
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = "List<opsmgmt.v4.error.AppMessage>"
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = "List<opsmgmt.v4.error.AppMessage>"
							return nil
						}
					}
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import1.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "opsmgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType2001 := new([]import1.AppMessage)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if len(*vOneOfType2001) == 0 || (vOneOfType2001 != nil && (*vOneOfType2001)[0].ObjectType_ != nil && "opsmgmt.v4.error.AppMessage" == *((*vOneOfType2001)[0].ObjectType_)) {
			p.oneOfType2001 = *vOneOfType2001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<opsmgmt.v4.error.AppMessage>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<opsmgmt.v4.error.AppMessage>"
			return nil
		}
	}
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "opsmgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateDashboardApiResponseData"))
}

func (p *OneOfUpdateDashboardApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if "List<opsmgmt.v4.error.AppMessage>" == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateDashboardApiResponseData")
}

type OneOfWidgetWidgetInfo struct {
	Discriminator *string         `json:"-"`
	ObjectType_   *string         `json:"-"`
	oneOfType2001 *WidgetTemplate `json:"-"`
	oneOfType2002 *WidgetConfig   `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfWidgetWidgetInfo() *OneOfWidgetWidgetInfo {
	p := new(OneOfWidgetWidgetInfo)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfWidgetWidgetInfo) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfWidgetWidgetInfo is nil"))
	}
	switch v.(type) {
	case WidgetTemplate:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(WidgetTemplate)
		}
		*p.oneOfType2001 = v.(WidgetTemplate)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case WidgetConfig:
		if nil == p.oneOfType2002 {
			p.oneOfType2002 = new(WidgetConfig)
		}
		*p.oneOfType2002 = v.(WidgetConfig)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2002.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2002.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfWidgetWidgetInfo) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType2002 != nil && *p.oneOfType2002.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2002
	}
	return nil
}

func (p *OneOfWidgetWidgetInfo) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType2001 := new(WidgetTemplate)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType2001)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType2001.ObjectType_ != nil && "opsmgmt.v4.config.WidgetTemplate" == *vOneOfType2001.ObjectType_ {
							if nil == p.oneOfType2001 {
								p.oneOfType2001 = new(WidgetTemplate)
							}
							*p.oneOfType2001 = *vOneOfType2001
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType2001.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType2001.ObjectType_
							return nil
						}
					}
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType2002 := new(WidgetConfig)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType2002)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType2002.ObjectType_ != nil && "opsmgmt.v4.config.WidgetConfig" == *vOneOfType2002.ObjectType_ {
							if nil == p.oneOfType2002 {
								p.oneOfType2002 = new(WidgetConfig)
							}
							*p.oneOfType2002 = *vOneOfType2002
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType2002.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType2002.ObjectType_
							return nil
						}
					}
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType2001 := new(WidgetTemplate)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if vOneOfType2001.ObjectType_ != nil && "opsmgmt.v4.config.WidgetTemplate" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(WidgetTemplate)
			}
			*p.oneOfType2001 = *vOneOfType2001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2001.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2001.ObjectType_
			return nil
		}
	}
	vOneOfType2002 := new(WidgetConfig)
	if err := json.Unmarshal(b, vOneOfType2002); err == nil {
		if vOneOfType2002.ObjectType_ != nil && "opsmgmt.v4.config.WidgetConfig" == *vOneOfType2002.ObjectType_ {
			if nil == p.oneOfType2002 {
				p.oneOfType2002 = new(WidgetConfig)
			}
			*p.oneOfType2002 = *vOneOfType2002
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2002.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2002.ObjectType_
			return nil
		}
	}
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfWidgetWidgetInfo"))
}

func (p *OneOfWidgetWidgetInfo) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType2002 != nil && *p.oneOfType2002.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2002)
	}
	return nil, errors.New("No value to marshal for OneOfWidgetWidgetInfo")
}

type OneOfCreateDashboardApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType3001 *Dashboard             `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfCreateDashboardApiResponseData() *OneOfCreateDashboardApiResponseData {
	p := new(OneOfCreateDashboardApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateDashboardApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateDashboardApiResponseData is nil"))
	}
	switch v.(type) {
	case Dashboard:
		if nil == p.oneOfType3001 {
			p.oneOfType3001 = new(Dashboard)
		}
		*p.oneOfType3001 = v.(Dashboard)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType3001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType3001.ObjectType_
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfCreateDashboardApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType3001 != nil && *p.oneOfType3001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType3001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfCreateDashboardApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType3001 := new(Dashboard)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType3001)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType3001.ObjectType_ != nil && "opsmgmt.v4.config.Dashboard" == *vOneOfType3001.ObjectType_ {
							if nil == p.oneOfType3001 {
								p.oneOfType3001 = new(Dashboard)
							}
							*p.oneOfType3001 = *vOneOfType3001
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType3001.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType3001.ObjectType_
							return nil
						}
					}
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import1.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "opsmgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType3001 := new(Dashboard)
	if err := json.Unmarshal(b, vOneOfType3001); err == nil {
		if vOneOfType3001.ObjectType_ != nil && "opsmgmt.v4.config.Dashboard" == *vOneOfType3001.ObjectType_ {
			if nil == p.oneOfType3001 {
				p.oneOfType3001 = new(Dashboard)
			}
			*p.oneOfType3001 = *vOneOfType3001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType3001.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType3001.ObjectType_
			return nil
		}
	}
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "opsmgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateDashboardApiResponseData"))
}

func (p *OneOfCreateDashboardApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType3001 != nil && *p.oneOfType3001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType3001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfCreateDashboardApiResponseData")
}

type OneOfCreateReportApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfCreateReportApiResponseData() *OneOfCreateReportApiResponseData {
	p := new(OneOfCreateReportApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateReportApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateReportApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.TaskReference)
		}
		*p.oneOfType0 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfCreateReportApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfCreateReportApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType0 := new(import3.TaskReference)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType0)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType0.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
							if nil == p.oneOfType0 {
								p.oneOfType0 = new(import3.TaskReference)
							}
							*p.oneOfType0 = *vOneOfType0
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType0.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType0.ObjectType_
							return nil
						}
					}
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import1.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "opsmgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if vOneOfType0.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.TaskReference)
			}
			*p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType0.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType0.ObjectType_
			return nil
		}
	}
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "opsmgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateReportApiResponseData"))
}

func (p *OneOfCreateReportApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfCreateReportApiResponseData")
}

type OneOfUpdateReportConfigApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfUpdateReportConfigApiResponseData() *OneOfUpdateReportConfigApiResponseData {
	p := new(OneOfUpdateReportConfigApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateReportConfigApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateReportConfigApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.TaskReference)
		}
		*p.oneOfType0 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfUpdateReportConfigApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfUpdateReportConfigApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType0 := new(import3.TaskReference)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType0)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType0.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
							if nil == p.oneOfType0 {
								p.oneOfType0 = new(import3.TaskReference)
							}
							*p.oneOfType0 = *vOneOfType0
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType0.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType0.ObjectType_
							return nil
						}
					}
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import1.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "opsmgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if vOneOfType0.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.TaskReference)
			}
			*p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType0.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType0.ObjectType_
			return nil
		}
	}
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "opsmgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateReportConfigApiResponseData"))
}

func (p *OneOfUpdateReportConfigApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateReportConfigApiResponseData")
}

type OneOfGetReportApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *Report                `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfGetReportApiResponseData() *OneOfGetReportApiResponseData {
	p := new(OneOfGetReportApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetReportApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetReportApiResponseData is nil"))
	}
	switch v.(type) {
	case Report:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(Report)
		}
		*p.oneOfType0 = v.(Report)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfGetReportApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetReportApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType0 := new(Report)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType0)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType0.ObjectType_ != nil && "opsmgmt.v4.config.Report" == *vOneOfType0.ObjectType_ {
							if nil == p.oneOfType0 {
								p.oneOfType0 = new(Report)
							}
							*p.oneOfType0 = *vOneOfType0
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType0.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType0.ObjectType_
							return nil
						}
					}
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import1.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "opsmgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType0 := new(Report)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if vOneOfType0.ObjectType_ != nil && "opsmgmt.v4.config.Report" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(Report)
			}
			*p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType0.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType0.ObjectType_
			return nil
		}
	}
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "opsmgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetReportApiResponseData"))
}

func (p *OneOfGetReportApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetReportApiResponseData")
}

type OneOfCreateReportConfigApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfCreateReportConfigApiResponseData() *OneOfCreateReportConfigApiResponseData {
	p := new(OneOfCreateReportConfigApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateReportConfigApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateReportConfigApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.TaskReference)
		}
		*p.oneOfType0 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfCreateReportConfigApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfCreateReportConfigApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType0 := new(import3.TaskReference)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType0)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType0.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
							if nil == p.oneOfType0 {
								p.oneOfType0 = new(import3.TaskReference)
							}
							*p.oneOfType0 = *vOneOfType0
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType0.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType0.ObjectType_
							return nil
						}
					}
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import1.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "opsmgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if vOneOfType0.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.TaskReference)
			}
			*p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType0.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType0.ObjectType_
			return nil
		}
	}
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "opsmgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateReportConfigApiResponseData"))
}

func (p *OneOfCreateReportConfigApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfCreateReportConfigApiResponseData")
}

type OneOfNotifyReportApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfNotifyReportApiResponseData() *OneOfNotifyReportApiResponseData {
	p := new(OneOfNotifyReportApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfNotifyReportApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfNotifyReportApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.TaskReference)
		}
		*p.oneOfType0 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfNotifyReportApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfNotifyReportApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType0 := new(import3.TaskReference)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType0)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType0.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
							if nil == p.oneOfType0 {
								p.oneOfType0 = new(import3.TaskReference)
							}
							*p.oneOfType0 = *vOneOfType0
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType0.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType0.ObjectType_
							return nil
						}
					}
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import1.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "opsmgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if vOneOfType0.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.TaskReference)
			}
			*p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType0.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType0.ObjectType_
			return nil
		}
	}
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "opsmgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfNotifyReportApiResponseData"))
}

func (p *OneOfNotifyReportApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfNotifyReportApiResponseData")
}

type OneOfGetGlobalReportSettingApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *GlobalReportSetting   `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfGetGlobalReportSettingApiResponseData() *OneOfGetGlobalReportSettingApiResponseData {
	p := new(OneOfGetGlobalReportSettingApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetGlobalReportSettingApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetGlobalReportSettingApiResponseData is nil"))
	}
	switch v.(type) {
	case GlobalReportSetting:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(GlobalReportSetting)
		}
		*p.oneOfType0 = v.(GlobalReportSetting)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfGetGlobalReportSettingApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetGlobalReportSettingApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType0 := new(GlobalReportSetting)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType0)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType0.ObjectType_ != nil && "opsmgmt.v4.config.GlobalReportSetting" == *vOneOfType0.ObjectType_ {
							if nil == p.oneOfType0 {
								p.oneOfType0 = new(GlobalReportSetting)
							}
							*p.oneOfType0 = *vOneOfType0
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType0.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType0.ObjectType_
							return nil
						}
					}
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import1.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "opsmgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType0 := new(GlobalReportSetting)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if vOneOfType0.ObjectType_ != nil && "opsmgmt.v4.config.GlobalReportSetting" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(GlobalReportSetting)
			}
			*p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType0.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType0.ObjectType_
			return nil
		}
	}
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "opsmgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetGlobalReportSettingApiResponseData"))
}

func (p *OneOfGetGlobalReportSettingApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetGlobalReportSettingApiResponseData")
}

type OneOfLocationChildren struct {
	Discriminator *string `json:"-"`
	ObjectType_   *string `json:"-"`
	oneOfType1001 *Layer  `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfLocationChildren() *OneOfLocationChildren {
	p := new(OneOfLocationChildren)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfLocationChildren) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfLocationChildren is nil"))
	}
	switch v.(type) {
	case Layer:
		if nil == p.oneOfType1001 {
			p.oneOfType1001 = new(Layer)
		}
		*p.oneOfType1001 = v.(Layer)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType1001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType1001.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfLocationChildren) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType1001 != nil && *p.oneOfType1001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1001
	}
	return nil
}

func (p *OneOfLocationChildren) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType1001 := new(Layer)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType1001)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType1001.ObjectType_ != nil && "opsmgmt.v4.config.Layer" == *vOneOfType1001.ObjectType_ {
							if nil == p.oneOfType1001 {
								p.oneOfType1001 = new(Layer)
							}
							*p.oneOfType1001 = *vOneOfType1001
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType1001.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType1001.ObjectType_
							return nil
						}
					}
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType1001 := new(Layer)
	if err := json.Unmarshal(b, vOneOfType1001); err == nil {
		if vOneOfType1001.ObjectType_ != nil && "opsmgmt.v4.config.Layer" == *vOneOfType1001.ObjectType_ {
			if nil == p.oneOfType1001 {
				p.oneOfType1001 = new(Layer)
			}
			*p.oneOfType1001 = *vOneOfType1001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType1001.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType1001.ObjectType_
			return nil
		}
	}
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfLocationChildren"))
}

func (p *OneOfLocationChildren) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType1001 != nil && *p.oneOfType1001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1001)
	}
	return nil, errors.New("No value to marshal for OneOfLocationChildren")
}

type OneOfGetDashboardApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	oneOfType2001 *Dashboard             `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfGetDashboardApiResponseData() *OneOfGetDashboardApiResponseData {
	p := new(OneOfGetDashboardApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetDashboardApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetDashboardApiResponseData is nil"))
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
	case Dashboard:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(Dashboard)
		}
		*p.oneOfType2001 = v.(Dashboard)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfGetDashboardApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	return nil
}

func (p *OneOfGetDashboardApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import1.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "opsmgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType2001 := new(Dashboard)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType2001)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType2001.ObjectType_ != nil && "opsmgmt.v4.config.Dashboard" == *vOneOfType2001.ObjectType_ {
							if nil == p.oneOfType2001 {
								p.oneOfType2001 = new(Dashboard)
							}
							*p.oneOfType2001 = *vOneOfType2001
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType2001.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType2001.ObjectType_
							return nil
						}
					}
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "opsmgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType2001 := new(Dashboard)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if vOneOfType2001.ObjectType_ != nil && "opsmgmt.v4.config.Dashboard" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(Dashboard)
			}
			*p.oneOfType2001 = *vOneOfType2001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2001.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2001.ObjectType_
			return nil
		}
	}
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetDashboardApiResponseData"))
}

func (p *OneOfGetDashboardApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfGetDashboardApiResponseData")
}

type OneOfGetDashboardSettingsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	oneOfType2001 *DashboardSettings     `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfGetDashboardSettingsApiResponseData() *OneOfGetDashboardSettingsApiResponseData {
	p := new(OneOfGetDashboardSettingsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetDashboardSettingsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetDashboardSettingsApiResponseData is nil"))
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
	case DashboardSettings:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(DashboardSettings)
		}
		*p.oneOfType2001 = v.(DashboardSettings)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfGetDashboardSettingsApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	return nil
}

func (p *OneOfGetDashboardSettingsApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import1.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "opsmgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType2001 := new(DashboardSettings)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType2001)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType2001.ObjectType_ != nil && "opsmgmt.v4.config.DashboardSettings" == *vOneOfType2001.ObjectType_ {
							if nil == p.oneOfType2001 {
								p.oneOfType2001 = new(DashboardSettings)
							}
							*p.oneOfType2001 = *vOneOfType2001
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType2001.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType2001.ObjectType_
							return nil
						}
					}
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "opsmgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType2001 := new(DashboardSettings)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if vOneOfType2001.ObjectType_ != nil && "opsmgmt.v4.config.DashboardSettings" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(DashboardSettings)
			}
			*p.oneOfType2001 = *vOneOfType2001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2001.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2001.ObjectType_
			return nil
		}
	}
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetDashboardSettingsApiResponseData"))
}

func (p *OneOfGetDashboardSettingsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfGetDashboardSettingsApiResponseData")
}

type OneOfListReportConfigsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	oneOfType0    []ReportConfig         `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfListReportConfigsApiResponseData() *OneOfListReportConfigsApiResponseData {
	p := new(OneOfListReportConfigsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListReportConfigsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListReportConfigsApiResponseData is nil"))
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
	case []ReportConfig:
		p.oneOfType0 = v.([]ReportConfig)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<opsmgmt.v4.config.ReportConfig>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<opsmgmt.v4.config.ReportConfig>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListReportConfigsApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<opsmgmt.v4.config.ReportConfig>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfListReportConfigsApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import1.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "opsmgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["List<opsmgmt.v4.config.ReportConfig>"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType0 := new([]ReportConfig)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType0)
					if unmarshalErr == nil {
						// For arrays, verify the array item ObjectType matches
						if vOneOfType0 == nil || len(*vOneOfType0) == 0 || ((*vOneOfType0)[0].ObjectType_ != nil && "opsmgmt.v4.config.ReportConfig" == *((*vOneOfType0)[0].ObjectType_)) {
							p.oneOfType0 = *vOneOfType0
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = "List<opsmgmt.v4.config.ReportConfig>"
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = "List<opsmgmt.v4.config.ReportConfig>"
							return nil
						}
					}
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "opsmgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new([]ReportConfig)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || (vOneOfType0 != nil && (*vOneOfType0)[0].ObjectType_ != nil && "opsmgmt.v4.config.ReportConfig" == *((*vOneOfType0)[0].ObjectType_)) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<opsmgmt.v4.config.ReportConfig>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<opsmgmt.v4.config.ReportConfig>"
			return nil
		}
	}
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListReportConfigsApiResponseData"))
}

func (p *OneOfListReportConfigsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<opsmgmt.v4.config.ReportConfig>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfListReportConfigsApiResponseData")
}

type OneOfDeleteDashboardApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	oneOfType1008 *interface{}           `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfDeleteDashboardApiResponseData() *OneOfDeleteDashboardApiResponseData {
	p := new(OneOfDeleteDashboardApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteDashboardApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteDashboardApiResponseData is nil"))
	}
	if nil == v {
		if nil == p.oneOfType1008 {
			p.oneOfType1008 = new(interface{})
		}
		*p.oneOfType1008 = nil
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "EMPTY"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "EMPTY"
		return nil
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfDeleteDashboardApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "EMPTY" == *p.Discriminator {
		return *p.oneOfType1008
	}
	return nil
}

func (p *OneOfDeleteDashboardApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType1008 := new(interface{})
	if err := json.Unmarshal(b, vOneOfType1008); err == nil {
		if nil == *vOneOfType1008 {
			if nil == p.oneOfType1008 {
				p.oneOfType1008 = new(interface{})
			}
			*p.oneOfType1008 = nil
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "EMPTY"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "EMPTY"
			return nil
		}
	}
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import1.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "opsmgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "opsmgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteDashboardApiResponseData"))
}

func (p *OneOfDeleteDashboardApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "EMPTY" == *p.Discriminator {
		return json.Marshal(p.oneOfType1008)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteDashboardApiResponseData")
}

type OneOfDeleteReportConfigApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType1    *interface{}           `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfDeleteReportConfigApiResponseData() *OneOfDeleteReportConfigApiResponseData {
	p := new(OneOfDeleteReportConfigApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteReportConfigApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteReportConfigApiResponseData is nil"))
	}
	if nil == v {
		if nil == p.oneOfType1 {
			p.oneOfType1 = new(interface{})
		}
		*p.oneOfType1 = nil
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "EMPTY"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "EMPTY"
		return nil
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfDeleteReportConfigApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if "EMPTY" == *p.Discriminator {
		return *p.oneOfType1
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDeleteReportConfigApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType1 := new(interface{})
	if err := json.Unmarshal(b, vOneOfType1); err == nil {
		if nil == *vOneOfType1 {
			if nil == p.oneOfType1 {
				p.oneOfType1 = new(interface{})
			}
			*p.oneOfType1 = nil
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "EMPTY"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "EMPTY"
			return nil
		}
	}
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import1.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "opsmgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "opsmgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteReportConfigApiResponseData"))
}

func (p *OneOfDeleteReportConfigApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if "EMPTY" == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteReportConfigApiResponseData")
}

type OneOfListReportsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	oneOfType0    []Report               `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfListReportsApiResponseData() *OneOfListReportsApiResponseData {
	p := new(OneOfListReportsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListReportsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListReportsApiResponseData is nil"))
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
	case []Report:
		p.oneOfType0 = v.([]Report)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<opsmgmt.v4.config.Report>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<opsmgmt.v4.config.Report>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListReportsApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<opsmgmt.v4.config.Report>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfListReportsApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import1.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "opsmgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["List<opsmgmt.v4.config.Report>"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType0 := new([]Report)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType0)
					if unmarshalErr == nil {
						// For arrays, verify the array item ObjectType matches
						if vOneOfType0 == nil || len(*vOneOfType0) == 0 || ((*vOneOfType0)[0].ObjectType_ != nil && "opsmgmt.v4.config.Report" == *((*vOneOfType0)[0].ObjectType_)) {
							p.oneOfType0 = *vOneOfType0
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = "List<opsmgmt.v4.config.Report>"
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = "List<opsmgmt.v4.config.Report>"
							return nil
						}
					}
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "opsmgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new([]Report)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || (vOneOfType0 != nil && (*vOneOfType0)[0].ObjectType_ != nil && "opsmgmt.v4.config.Report" == *((*vOneOfType0)[0].ObjectType_)) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<opsmgmt.v4.config.Report>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<opsmgmt.v4.config.Report>"
			return nil
		}
	}
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListReportsApiResponseData"))
}

func (p *OneOfListReportsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<opsmgmt.v4.config.Report>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfListReportsApiResponseData")
}

type OneOfListDashboardApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	oneOfType2001 []Dashboard            `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfListDashboardApiResponseData() *OneOfListDashboardApiResponseData {
	p := new(OneOfListDashboardApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListDashboardApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListDashboardApiResponseData is nil"))
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
	case []Dashboard:
		p.oneOfType2001 = v.([]Dashboard)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<opsmgmt.v4.config.Dashboard>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<opsmgmt.v4.config.Dashboard>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListDashboardApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<opsmgmt.v4.config.Dashboard>" == *p.Discriminator {
		return p.oneOfType2001
	}
	return nil
}

func (p *OneOfListDashboardApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import1.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "opsmgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["List<opsmgmt.v4.config.Dashboard>"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType2001 := new([]Dashboard)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType2001)
					if unmarshalErr == nil {
						// For arrays, verify the array item ObjectType matches
						if vOneOfType2001 == nil || len(*vOneOfType2001) == 0 || ((*vOneOfType2001)[0].ObjectType_ != nil && "opsmgmt.v4.config.Dashboard" == *((*vOneOfType2001)[0].ObjectType_)) {
							p.oneOfType2001 = *vOneOfType2001
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = "List<opsmgmt.v4.config.Dashboard>"
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = "List<opsmgmt.v4.config.Dashboard>"
							return nil
						}
					}
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "opsmgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType2001 := new([]Dashboard)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if len(*vOneOfType2001) == 0 || (vOneOfType2001 != nil && (*vOneOfType2001)[0].ObjectType_ != nil && "opsmgmt.v4.config.Dashboard" == *((*vOneOfType2001)[0].ObjectType_)) {
			p.oneOfType2001 = *vOneOfType2001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<opsmgmt.v4.config.Dashboard>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<opsmgmt.v4.config.Dashboard>"
			return nil
		}
	}
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListDashboardApiResponseData"))
}

func (p *OneOfListDashboardApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<opsmgmt.v4.config.Dashboard>" == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfListDashboardApiResponseData")
}

type OneOfUploadDashboardGeoconfigurationApiResponseData struct {
	Discriminator *string                    `json:"-"`
	ObjectType_   *string                    `json:"-"`
	oneOfType400  *import1.ErrorResponse     `json:"-"`
	oneOfType2080 *DashboardGeoconfiguration `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfUploadDashboardGeoconfigurationApiResponseData() *OneOfUploadDashboardGeoconfigurationApiResponseData {
	p := new(OneOfUploadDashboardGeoconfigurationApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUploadDashboardGeoconfigurationApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUploadDashboardGeoconfigurationApiResponseData is nil"))
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
	case DashboardGeoconfiguration:
		if nil == p.oneOfType2080 {
			p.oneOfType2080 = new(DashboardGeoconfiguration)
		}
		*p.oneOfType2080 = v.(DashboardGeoconfiguration)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2080.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2080.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfUploadDashboardGeoconfigurationApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType2080 != nil && *p.oneOfType2080.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2080
	}
	return nil
}

func (p *OneOfUploadDashboardGeoconfigurationApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import1.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "opsmgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType2080 := new(DashboardGeoconfiguration)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType2080)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType2080.ObjectType_ != nil && "opsmgmt.v4.config.DashboardGeoconfiguration" == *vOneOfType2080.ObjectType_ {
							if nil == p.oneOfType2080 {
								p.oneOfType2080 = new(DashboardGeoconfiguration)
							}
							*p.oneOfType2080 = *vOneOfType2080
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType2080.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType2080.ObjectType_
							return nil
						}
					}
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "opsmgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType2080 := new(DashboardGeoconfiguration)
	if err := json.Unmarshal(b, vOneOfType2080); err == nil {
		if vOneOfType2080.ObjectType_ != nil && "opsmgmt.v4.config.DashboardGeoconfiguration" == *vOneOfType2080.ObjectType_ {
			if nil == p.oneOfType2080 {
				p.oneOfType2080 = new(DashboardGeoconfiguration)
			}
			*p.oneOfType2080 = *vOneOfType2080
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2080.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2080.ObjectType_
			return nil
		}
	}
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUploadDashboardGeoconfigurationApiResponseData"))
}

func (p *OneOfUploadDashboardGeoconfigurationApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType2080 != nil && *p.oneOfType2080.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2080)
	}
	return nil, errors.New("No value to marshal for OneOfUploadDashboardGeoconfigurationApiResponseData")
}

type OneOfGetReportConfigApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *ReportConfig          `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfGetReportConfigApiResponseData() *OneOfGetReportConfigApiResponseData {
	p := new(OneOfGetReportConfigApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetReportConfigApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetReportConfigApiResponseData is nil"))
	}
	switch v.(type) {
	case ReportConfig:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(ReportConfig)
		}
		*p.oneOfType0 = v.(ReportConfig)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfGetReportConfigApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetReportConfigApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType0 := new(ReportConfig)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType0)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType0.ObjectType_ != nil && "opsmgmt.v4.config.ReportConfig" == *vOneOfType0.ObjectType_ {
							if nil == p.oneOfType0 {
								p.oneOfType0 = new(ReportConfig)
							}
							*p.oneOfType0 = *vOneOfType0
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType0.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType0.ObjectType_
							return nil
						}
					}
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import1.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "opsmgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType0 := new(ReportConfig)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if vOneOfType0.ObjectType_ != nil && "opsmgmt.v4.config.ReportConfig" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(ReportConfig)
			}
			*p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType0.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType0.ObjectType_
			return nil
		}
	}
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "opsmgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetReportConfigApiResponseData"))
}

func (p *OneOfGetReportConfigApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetReportConfigApiResponseData")
}

type OneOfDeleteReportApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType1    *interface{}           `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfDeleteReportApiResponseData() *OneOfDeleteReportApiResponseData {
	p := new(OneOfDeleteReportApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteReportApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteReportApiResponseData is nil"))
	}
	if nil == v {
		if nil == p.oneOfType1 {
			p.oneOfType1 = new(interface{})
		}
		*p.oneOfType1 = nil
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "EMPTY"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "EMPTY"
		return nil
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfDeleteReportApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if "EMPTY" == *p.Discriminator {
		return *p.oneOfType1
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDeleteReportApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType1 := new(interface{})
	if err := json.Unmarshal(b, vOneOfType1); err == nil {
		if nil == *vOneOfType1 {
			if nil == p.oneOfType1 {
				p.oneOfType1 = new(interface{})
			}
			*p.oneOfType1 = nil
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "EMPTY"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "EMPTY"
			return nil
		}
	}
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import1.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "opsmgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "opsmgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteReportApiResponseData"))
}

func (p *OneOfDeleteReportApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if "EMPTY" == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteReportApiResponseData")
}

type OneOfUpdateGlobalReportSettingsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    []import1.AppMessage   `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfUpdateGlobalReportSettingsApiResponseData() *OneOfUpdateGlobalReportSettingsApiResponseData {
	p := new(OneOfUpdateGlobalReportSettingsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateGlobalReportSettingsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateGlobalReportSettingsApiResponseData is nil"))
	}
	switch v.(type) {
	case []import1.AppMessage:
		p.oneOfType0 = v.([]import1.AppMessage)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<opsmgmt.v4.error.AppMessage>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<opsmgmt.v4.error.AppMessage>"
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfUpdateGlobalReportSettingsApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if "List<opsmgmt.v4.error.AppMessage>" == *p.Discriminator {
		return p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfUpdateGlobalReportSettingsApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["List<opsmgmt.v4.error.AppMessage>"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType0 := new([]import1.AppMessage)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType0)
					if unmarshalErr == nil {
						// For arrays, verify the array item ObjectType matches
						if vOneOfType0 == nil || len(*vOneOfType0) == 0 || ((*vOneOfType0)[0].ObjectType_ != nil && "opsmgmt.v4.error.AppMessage" == *((*vOneOfType0)[0].ObjectType_)) {
							p.oneOfType0 = *vOneOfType0
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = "List<opsmgmt.v4.error.AppMessage>"
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = "List<opsmgmt.v4.error.AppMessage>"
							return nil
						}
					}
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import1.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "opsmgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType0 := new([]import1.AppMessage)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || (vOneOfType0 != nil && (*vOneOfType0)[0].ObjectType_ != nil && "opsmgmt.v4.error.AppMessage" == *((*vOneOfType0)[0].ObjectType_)) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<opsmgmt.v4.error.AppMessage>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<opsmgmt.v4.error.AppMessage>"
			return nil
		}
	}
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "opsmgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateGlobalReportSettingsApiResponseData"))
}

func (p *OneOfUpdateGlobalReportSettingsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if "List<opsmgmt.v4.error.AppMessage>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateGlobalReportSettingsApiResponseData")
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
