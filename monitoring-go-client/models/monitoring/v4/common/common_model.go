/*
 * Generated file models/monitoring/v4/common/common_model.go.
 *
 * Product version: 4.0.1-beta-1
 *
 * Part of the Nutanix Monitoring Versioned APIs
 *
 * (c) 2024 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module monitoring.v4.common of Nutanix Monitoring Versioned APIs
*/
package common

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type AlertEntityReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Unique UUID of the entity.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  The name of the entity.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  The type of entity. For example, VM, node, or cluster.
	*/
	Type *string `json:"type,omitempty"`
}

func NewAlertEntityReference() *AlertEntityReference {
	p := new(AlertEntityReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.common.AlertEntityReference"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type BoolValue struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	BoolValue *bool `json:"boolValue"`
}

func (p *BoolValue) MarshalJSON() ([]byte, error) {
	type BoolValueProxy BoolValue
	return json.Marshal(struct {
		*BoolValueProxy
		BoolValue *bool `json:"boolValue,omitempty"`
	}{
		BoolValueProxy: (*BoolValueProxy)(p),
		BoolValue:      p.BoolValue,
	})
}

func NewBoolValue() *BoolValue {
	p := new(BoolValue)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.common.BoolValue"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Possible causes, resolutions and additional details to troubleshoot this alert.
*/
type CauseAndResolution struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Possible cause for the generated alert or event.
	*/
	Cause *string `json:"cause,omitempty"`
	/*
	  Possible resolution for the generated alert or event.
	*/
	Resolution *string `json:"resolution,omitempty"`
}

func NewCauseAndResolution() *CauseAndResolution {
	p := new(CauseAndResolution)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.common.CauseAndResolution"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The comparison operator used for the condition evaluation.
*/
type ComparisonOperator int

const (
	COMPARISONOPERATOR_UNKNOWN                  ComparisonOperator = 0
	COMPARISONOPERATOR_REDACTED                 ComparisonOperator = 1
	COMPARISONOPERATOR_EQUAL_TO                 ComparisonOperator = 2
	COMPARISONOPERATOR_GREATER_THAN             ComparisonOperator = 3
	COMPARISONOPERATOR_GREATER_THAN_OR_EQUAL_TO ComparisonOperator = 4
	COMPARISONOPERATOR_LESS_THAN                ComparisonOperator = 5
	COMPARISONOPERATOR_LESS_THAN_OR_EQUAL_TO    ComparisonOperator = 6
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ComparisonOperator) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"EQUAL_TO",
		"GREATER_THAN",
		"GREATER_THAN_OR_EQUAL_TO",
		"LESS_THAN",
		"LESS_THAN_OR_EQUAL_TO",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ComparisonOperator) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"EQUAL_TO",
		"GREATER_THAN",
		"GREATER_THAN_OR_EQUAL_TO",
		"LESS_THAN",
		"LESS_THAN_OR_EQUAL_TO",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ComparisonOperator) index(name string) ComparisonOperator {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"EQUAL_TO",
		"GREATER_THAN",
		"GREATER_THAN_OR_EQUAL_TO",
		"LESS_THAN",
		"LESS_THAN_OR_EQUAL_TO",
	}
	for idx := range names {
		if names[idx] == name {
			return ComparisonOperator(idx)
		}
	}
	return COMPARISONOPERATOR_UNKNOWN
}

func (e *ComparisonOperator) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ComparisonOperator:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ComparisonOperator) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ComparisonOperator) Ref() *ComparisonOperator {
	return &e
}

/*
Indicating if this symptom is caused by static threshold or anomaly (dynamic threshold) evaluation.  If an indicator is raised, there may have another indicator indicating the safeguard zone value.
*/
type ConditionType int

const (
	CONDITIONTYPE_UNKNOWN          ConditionType = 0
	CONDITIONTYPE_REDACTED         ConditionType = 1
	CONDITIONTYPE_STATIC_THRESHOLD ConditionType = 2
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ConditionType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"STATIC_THRESHOLD",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ConditionType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"STATIC_THRESHOLD",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ConditionType) index(name string) ConditionType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"STATIC_THRESHOLD",
	}
	for idx := range names {
		if names[idx] == name {
			return ConditionType(idx)
		}
	}
	return CONDITIONTYPE_UNKNOWN
}

func (e *ConditionType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ConditionType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ConditionType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ConditionType) Ref() *ConditionType {
	return &e
}

/*
Data type of the metric value as stored in database.
*/
type DataType int

const (
	DATATYPE_UNKNOWN  DataType = 0
	DATATYPE_REDACTED DataType = 1
	DATATYPE_STRING   DataType = 2
	DATATYPE_BOOLEAN  DataType = 3
	DATATYPE_DOUBLE   DataType = 4
	DATATYPE_LONG     DataType = 5
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *DataType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"STRING",
		"BOOLEAN",
		"DOUBLE",
		"LONG",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e DataType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"STRING",
		"BOOLEAN",
		"DOUBLE",
		"LONG",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *DataType) index(name string) DataType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"STRING",
		"BOOLEAN",
		"DOUBLE",
		"LONG",
	}
	for idx := range names {
		if names[idx] == name {
			return DataType(idx)
		}
	}
	return DATATYPE_UNKNOWN
}

func (e *DataType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for DataType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *DataType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e DataType) Ref() *DataType {
	return &e
}

type DoubleValue struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	DoubleValue *float64 `json:"doubleValue"`
}

func (p *DoubleValue) MarshalJSON() ([]byte, error) {
	type DoubleValueProxy DoubleValue
	return json.Marshal(struct {
		*DoubleValueProxy
		DoubleValue *float64 `json:"doubleValue,omitempty"`
	}{
		DoubleValueProxy: (*DoubleValueProxy)(p),
		DoubleValue:      p.DoubleValue,
	})
}

func NewDoubleValue() *DoubleValue {
	p := new(DoubleValue)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.common.DoubleValue"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The source entity associated with the alert, event, or audit.
*/
type EntityReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Unique UUID of the entity.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  The name of the entity.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  The type of entity. For example, VM, node, or cluster.
	*/
	Type *string `json:"type,omitempty"`
}

func NewEntityReference() *EntityReference {
	p := new(EntityReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.common.EntityReference"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The impact this alert or event will have on the system. For example, availability, performance, or capacity.
*/
type ImpactType int

const (
	IMPACTTYPE_UNKNOWN          ImpactType = 0
	IMPACTTYPE_REDACTED         ImpactType = 1
	IMPACTTYPE_AVAILABILITY     ImpactType = 2
	IMPACTTYPE_CAPACITY         ImpactType = 3
	IMPACTTYPE_CONFIGURATION    ImpactType = 4
	IMPACTTYPE_PERFORMANCE      ImpactType = 5
	IMPACTTYPE_SYSTEM_INDICATOR ImpactType = 6
	IMPACTTYPE_CPU_CAPACITY     ImpactType = 7
	IMPACTTYPE_MEMORY_CAPACITY  ImpactType = 8
	IMPACTTYPE_STORAGE_CAPACITY ImpactType = 9
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ImpactType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"AVAILABILITY",
		"CAPACITY",
		"CONFIGURATION",
		"PERFORMANCE",
		"SYSTEM_INDICATOR",
		"CPU_CAPACITY",
		"MEMORY_CAPACITY",
		"STORAGE_CAPACITY",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ImpactType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"AVAILABILITY",
		"CAPACITY",
		"CONFIGURATION",
		"PERFORMANCE",
		"SYSTEM_INDICATOR",
		"CPU_CAPACITY",
		"MEMORY_CAPACITY",
		"STORAGE_CAPACITY",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ImpactType) index(name string) ImpactType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"AVAILABILITY",
		"CAPACITY",
		"CONFIGURATION",
		"PERFORMANCE",
		"SYSTEM_INDICATOR",
		"CPU_CAPACITY",
		"MEMORY_CAPACITY",
		"STORAGE_CAPACITY",
	}
	for idx := range names {
		if names[idx] == name {
			return ImpactType(idx)
		}
	}
	return IMPACTTYPE_UNKNOWN
}

func (e *ImpactType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ImpactType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ImpactType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ImpactType) Ref() *ImpactType {
	return &e
}

type IntValue struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	IntValue *int64 `json:"intValue"`
}

func (p *IntValue) MarshalJSON() ([]byte, error) {
	type IntValueProxy IntValue
	return json.Marshal(struct {
		*IntValueProxy
		IntValue *int64 `json:"intValue,omitempty"`
	}{
		IntValueProxy: (*IntValueProxy)(p),
		IntValue:      p.IntValue,
	})
}

func NewIntValue() *IntValue {
	p := new(IntValue)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.common.IntValue"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type MetricDetail struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	ComparisonOperator *ComparisonOperator `json:"comparisonOperator,omitempty"`

	ConditionType *ConditionType `json:"conditionType,omitempty"`

	DataType *DataType `json:"dataType,omitempty"`
	/*
	  Broad category under which this metric falls. For example, Disk, CPU, or Memory.
	*/
	MetricCategory *string `json:"metricCategory,omitempty"`
	/*
	  English readable name of the metric.
	*/
	MetricDisplayName *string `json:"metricDisplayName,omitempty"`
	/*
	  The metric key. Allowed values of metrics list can be found at https://portal.nutanix.com/page/documents/details?targetId=Prism-Central-Guide-vpc_2022_9:mul-alerts-user-created-metrics-r.html
	*/
	MetricName *string `json:"metricName,omitempty"`

	MetricValueItemDiscriminator_ *string `json:"$metricValueItemDiscriminator,omitempty"`
	/*
	  Raw value of the metric when the condition threshold was exceeded.
	*/
	MetricValue *OneOfMetricDetailMetricValue `json:"metricValue,omitempty"`

	ThresholdValueItemDiscriminator_ *string `json:"$thresholdValueItemDiscriminator,omitempty"`
	/*
	  The threshold value that was used for the condition evaluation.
	*/
	ThresholdValue *OneOfMetricDetailThresholdValue `json:"thresholdValue,omitempty"`
	/*
	  The time in ISO 8601 format when the event was triggered.
	*/
	TriggerTime *time.Time `json:"triggerTime,omitempty"`
	/*
	  How long the metric breached the given condition before raising an event.
	*/
	TriggerWaitTime *int64 `json:"triggerWaitTime,omitempty"`
	/*
	  Unit of the metric. For example, percentage, ms, or usecs.
	*/
	Unit *string `json:"unit,omitempty"`
}

func NewMetricDetail() *MetricDetail {
	p := new(MetricDetail)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.common.MetricDetail"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *MetricDetail) GetMetricValue() interface{} {
	if nil == p.MetricValue {
		return nil
	}
	return p.MetricValue.GetValue()
}

func (p *MetricDetail) SetMetricValue(v interface{}) error {
	if nil == p.MetricValue {
		p.MetricValue = NewOneOfMetricDetailMetricValue()
	}
	e := p.MetricValue.SetValue(v)
	if nil == e {
		if nil == p.MetricValueItemDiscriminator_ {
			p.MetricValueItemDiscriminator_ = new(string)
		}
		*p.MetricValueItemDiscriminator_ = *p.MetricValue.Discriminator
	}
	return e
}

/*
The operation type associated with the audit. For example, CREATE, UPDATE, or DELETE.
*/
type OperationType int

const (
	OPERATIONTYPE_UNKNOWN                     OperationType = 0
	OPERATIONTYPE_REDACTED                    OperationType = 1
	OPERATIONTYPE_CREATE                      OperationType = 2
	OPERATIONTYPE_UPDATE                      OperationType = 3
	OPERATIONTYPE_DELETE                      OperationType = 4
	OPERATIONTYPE_POWERSTATECHANGE            OperationType = 5
	OPERATIONTYPE_HA                          OperationType = 6
	OPERATIONTYPE_RESTORE                     OperationType = 7
	OPERATIONTYPE_MIGRATE                     OperationType = 8
	OPERATIONTYPE_ENABLE                      OperationType = 9
	OPERATIONTYPE_DISABLE                     OperationType = 10
	OPERATIONTYPE_RESET                       OperationType = 11
	OPERATIONTYPE_LOGIN                       OperationType = 12
	OPERATIONTYPE_LOGOUT                      OperationType = 13
	OPERATIONTYPE_VALIDATE                    OperationType = 14
	OPERATIONTYPE_FAILOVER                    OperationType = 15
	OPERATIONTYPE_INPROGRESS                  OperationType = 16
	OPERATIONTYPE_ON                          OperationType = 17
	OPERATIONTYPE_OFF                         OperationType = 18
	OPERATIONTYPE_JOINED                      OperationType = 19
	OPERATIONTYPE_UNJOINED                    OperationType = 20
	OPERATIONTYPE_PROTECTED                   OperationType = 21
	OPERATIONTYPE_UNPROTECTED                 OperationType = 22
	OPERATIONTYPE_QUARENTINED                 OperationType = 23
	OPERATIONTYPE_UNQUARENTINED               OperationType = 24
	OPERATIONTYPE_START                       OperationType = 25
	OPERATIONTYPE_END                         OperationType = 26
	OPERATIONTYPE_TESTFAILOVER                OperationType = 27
	OPERATIONTYPE_IMPORT                      OperationType = 28
	OPERATIONTYPE_EXPORT                      OperationType = 29
	OPERATIONTYPE_IMAGEIMPORT                 OperationType = 30
	OPERATIONTYPE_IMAGETRANSFER               OperationType = 31
	OPERATIONTYPE_IMAGEREMOVE                 OperationType = 32
	OPERATIONTYPE_IMAGEUPLOAD                 OperationType = 33
	OPERATIONTYPE_TEMPLATEGUESTUPDATEINITIATE OperationType = 34
	OPERATIONTYPE_TEMPLATEGUESTUPDATECANCEL   OperationType = 35
	OPERATIONTYPE_TEMPLATEGUESTUPDATECOMPLETE OperationType = 36
	OPERATIONTYPE_TEMPLATEDEPLOY              OperationType = 37
	OPERATIONTYPE_TEMPLATEVERSIONPUBLISH      OperationType = 38
	OPERATIONTYPE_CROSSCLUSTERMIGRATE         OperationType = 39
	OPERATIONTYPE_FLOWCLUSTERMIGRATE          OperationType = 40
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *OperationType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CREATE",
		"UPDATE",
		"DELETE",
		"POWERSTATECHANGE",
		"HA",
		"RESTORE",
		"MIGRATE",
		"ENABLE",
		"DISABLE",
		"RESET",
		"LOGIN",
		"LOGOUT",
		"VALIDATE",
		"FAILOVER",
		"INPROGRESS",
		"ON",
		"OFF",
		"JOINED",
		"UNJOINED",
		"PROTECTED",
		"UNPROTECTED",
		"QUARENTINED",
		"UNQUARENTINED",
		"START",
		"END",
		"TESTFAILOVER",
		"IMPORT",
		"EXPORT",
		"IMAGEIMPORT",
		"IMAGETRANSFER",
		"IMAGEREMOVE",
		"IMAGEUPLOAD",
		"TEMPLATEGUESTUPDATEINITIATE",
		"TEMPLATEGUESTUPDATECANCEL",
		"TEMPLATEGUESTUPDATECOMPLETE",
		"TEMPLATEDEPLOY",
		"TEMPLATEVERSIONPUBLISH",
		"CROSSCLUSTERMIGRATE",
		"FLOWCLUSTERMIGRATE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e OperationType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CREATE",
		"UPDATE",
		"DELETE",
		"POWERSTATECHANGE",
		"HA",
		"RESTORE",
		"MIGRATE",
		"ENABLE",
		"DISABLE",
		"RESET",
		"LOGIN",
		"LOGOUT",
		"VALIDATE",
		"FAILOVER",
		"INPROGRESS",
		"ON",
		"OFF",
		"JOINED",
		"UNJOINED",
		"PROTECTED",
		"UNPROTECTED",
		"QUARENTINED",
		"UNQUARENTINED",
		"START",
		"END",
		"TESTFAILOVER",
		"IMPORT",
		"EXPORT",
		"IMAGEIMPORT",
		"IMAGETRANSFER",
		"IMAGEREMOVE",
		"IMAGEUPLOAD",
		"TEMPLATEGUESTUPDATEINITIATE",
		"TEMPLATEGUESTUPDATECANCEL",
		"TEMPLATEGUESTUPDATECOMPLETE",
		"TEMPLATEDEPLOY",
		"TEMPLATEVERSIONPUBLISH",
		"CROSSCLUSTERMIGRATE",
		"FLOWCLUSTERMIGRATE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *OperationType) index(name string) OperationType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CREATE",
		"UPDATE",
		"DELETE",
		"POWERSTATECHANGE",
		"HA",
		"RESTORE",
		"MIGRATE",
		"ENABLE",
		"DISABLE",
		"RESET",
		"LOGIN",
		"LOGOUT",
		"VALIDATE",
		"FAILOVER",
		"INPROGRESS",
		"ON",
		"OFF",
		"JOINED",
		"UNJOINED",
		"PROTECTED",
		"UNPROTECTED",
		"QUARENTINED",
		"UNQUARENTINED",
		"START",
		"END",
		"TESTFAILOVER",
		"IMPORT",
		"EXPORT",
		"IMAGEIMPORT",
		"IMAGETRANSFER",
		"IMAGEREMOVE",
		"IMAGEUPLOAD",
		"TEMPLATEGUESTUPDATEINITIATE",
		"TEMPLATEGUESTUPDATECANCEL",
		"TEMPLATEGUESTUPDATECOMPLETE",
		"TEMPLATEDEPLOY",
		"TEMPLATEVERSIONPUBLISH",
		"CROSSCLUSTERMIGRATE",
		"FLOWCLUSTERMIGRATE",
	}
	for idx := range names {
		if names[idx] == name {
			return OperationType(idx)
		}
	}
	return OPERATIONTYPE_UNKNOWN
}

func (e *OperationType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for OperationType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *OperationType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e OperationType) Ref() *OperationType {
	return &e
}

type Parameter struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Name or key of additional parameter for an instance.
	*/
	ParamName *string `json:"paramName,omitempty"`

	ParamValueItemDiscriminator_ *string `json:"$paramValueItemDiscriminator,omitempty"`
	/*
	  Value of additional parameter for an instance.
	*/
	ParamValue *OneOfParameterParamValue `json:"paramValue,omitempty"`
}

func NewParameter() *Parameter {
	p := new(Parameter)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.common.Parameter"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *Parameter) GetParamValue() interface{} {
	if nil == p.ParamValue {
		return nil
	}
	return p.ParamValue.GetValue()
}

func (p *Parameter) SetParamValue(v interface{}) error {
	if nil == p.ParamValue {
		p.ParamValue = NewOneOfParameterParamValue()
	}
	e := p.ParamValue.SetValue(v)
	if nil == e {
		if nil == p.ParamValueItemDiscriminator_ {
			p.ParamValueItemDiscriminator_ = new(string)
		}
		*p.ParamValueItemDiscriminator_ = *p.ParamValue.Discriminator
	}
	return e
}

/*
Severity of an alert.
*/
type Severity int

const (
	SEVERITY_UNKNOWN  Severity = 0
	SEVERITY_REDACTED Severity = 1
	SEVERITY_INFO     Severity = 2
	SEVERITY_WARNING  Severity = 3
	SEVERITY_CRITICAL Severity = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *Severity) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"INFO",
		"WARNING",
		"CRITICAL",
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
		"INFO",
		"WARNING",
		"CRITICAL",
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
		"INFO",
		"WARNING",
		"CRITICAL",
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

type StringValue struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	StringValue *string `json:"stringValue"`
}

func (p *StringValue) MarshalJSON() ([]byte, error) {
	type StringValueProxy StringValue
	return json.Marshal(struct {
		*StringValueProxy
		StringValue *string `json:"stringValue,omitempty"`
	}{
		StringValueProxy: (*StringValueProxy)(p),
		StringValue:      p.StringValue,
	})
}

func NewStringValue() *StringValue {
	p := new(StringValue)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.common.StringValue"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type OneOfMetricDetailMetricValue struct {
	Discriminator *string      `json:"-"`
	ObjectType_   *string      `json:"-"`
	oneOfType0    *StringValue `json:"-"`
	oneOfType2    *DoubleValue `json:"-"`
	oneOfType3    *IntValue    `json:"-"`
	oneOfType1    *BoolValue   `json:"-"`
}

func NewOneOfMetricDetailMetricValue() *OneOfMetricDetailMetricValue {
	p := new(OneOfMetricDetailMetricValue)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfMetricDetailMetricValue) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfMetricDetailMetricValue is nil"))
	}
	switch v.(type) {
	case StringValue:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(StringValue)
		}
		*p.oneOfType0 = v.(StringValue)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case DoubleValue:
		if nil == p.oneOfType2 {
			p.oneOfType2 = new(DoubleValue)
		}
		*p.oneOfType2 = v.(DoubleValue)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2.ObjectType_
	case IntValue:
		if nil == p.oneOfType3 {
			p.oneOfType3 = new(IntValue)
		}
		*p.oneOfType3 = v.(IntValue)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType3.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType3.ObjectType_
	case BoolValue:
		if nil == p.oneOfType1 {
			p.oneOfType1 = new(BoolValue)
		}
		*p.oneOfType1 = v.(BoolValue)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType1.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType1.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfMetricDetailMetricValue) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType2 != nil && *p.oneOfType2.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2
	}
	if p.oneOfType3 != nil && *p.oneOfType3.ObjectType_ == *p.Discriminator {
		return *p.oneOfType3
	}
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1
	}
	return nil
}

func (p *OneOfMetricDetailMetricValue) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(StringValue)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "monitoring.v4.common.StringValue" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(StringValue)
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
	vOneOfType2 := new(DoubleValue)
	if err := json.Unmarshal(b, vOneOfType2); err == nil {
		if "monitoring.v4.common.DoubleValue" == *vOneOfType2.ObjectType_ {
			if nil == p.oneOfType2 {
				p.oneOfType2 = new(DoubleValue)
			}
			*p.oneOfType2 = *vOneOfType2
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2.ObjectType_
			return nil
		}
	}
	vOneOfType3 := new(IntValue)
	if err := json.Unmarshal(b, vOneOfType3); err == nil {
		if "monitoring.v4.common.IntValue" == *vOneOfType3.ObjectType_ {
			if nil == p.oneOfType3 {
				p.oneOfType3 = new(IntValue)
			}
			*p.oneOfType3 = *vOneOfType3
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType3.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType3.ObjectType_
			return nil
		}
	}
	vOneOfType1 := new(BoolValue)
	if err := json.Unmarshal(b, vOneOfType1); err == nil {
		if "monitoring.v4.common.BoolValue" == *vOneOfType1.ObjectType_ {
			if nil == p.oneOfType1 {
				p.oneOfType1 = new(BoolValue)
			}
			*p.oneOfType1 = *vOneOfType1
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType1.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType1.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfMetricDetailMetricValue"))
}

func (p *OneOfMetricDetailMetricValue) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType2 != nil && *p.oneOfType2.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2)
	}
	if p.oneOfType3 != nil && *p.oneOfType3.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType3)
	}
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	return nil, errors.New("No value to marshal for OneOfMetricDetailMetricValue")
}

type OneOfParameterParamValue struct {
	Discriminator *string      `json:"-"`
	ObjectType_   *string      `json:"-"`
	oneOfType2    *IntValue    `json:"-"`
	oneOfType0    *StringValue `json:"-"`
	oneOfType1    *BoolValue   `json:"-"`
}

func NewOneOfParameterParamValue() *OneOfParameterParamValue {
	p := new(OneOfParameterParamValue)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfParameterParamValue) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfParameterParamValue is nil"))
	}
	switch v.(type) {
	case IntValue:
		if nil == p.oneOfType2 {
			p.oneOfType2 = new(IntValue)
		}
		*p.oneOfType2 = v.(IntValue)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2.ObjectType_
	case StringValue:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(StringValue)
		}
		*p.oneOfType0 = v.(StringValue)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case BoolValue:
		if nil == p.oneOfType1 {
			p.oneOfType1 = new(BoolValue)
		}
		*p.oneOfType1 = v.(BoolValue)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType1.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType1.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfParameterParamValue) GetValue() interface{} {
	if p.oneOfType2 != nil && *p.oneOfType2.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1
	}
	return nil
}

func (p *OneOfParameterParamValue) UnmarshalJSON(b []byte) error {
	vOneOfType2 := new(IntValue)
	if err := json.Unmarshal(b, vOneOfType2); err == nil {
		if "monitoring.v4.common.IntValue" == *vOneOfType2.ObjectType_ {
			if nil == p.oneOfType2 {
				p.oneOfType2 = new(IntValue)
			}
			*p.oneOfType2 = *vOneOfType2
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2.ObjectType_
			return nil
		}
	}
	vOneOfType0 := new(StringValue)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "monitoring.v4.common.StringValue" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(StringValue)
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
	vOneOfType1 := new(BoolValue)
	if err := json.Unmarshal(b, vOneOfType1); err == nil {
		if "monitoring.v4.common.BoolValue" == *vOneOfType1.ObjectType_ {
			if nil == p.oneOfType1 {
				p.oneOfType1 = new(BoolValue)
			}
			*p.oneOfType1 = *vOneOfType1
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType1.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType1.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfParameterParamValue"))
}

func (p *OneOfParameterParamValue) MarshalJSON() ([]byte, error) {
	if p.oneOfType2 != nil && *p.oneOfType2.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	return nil, errors.New("No value to marshal for OneOfParameterParamValue")
}

type OneOfMetricDetailThresholdValue struct {
	Discriminator *string      `json:"-"`
	ObjectType_   *string      `json:"-"`
	oneOfType0    *StringValue `json:"-"`
	oneOfType2    *DoubleValue `json:"-"`
	oneOfType3    *IntValue    `json:"-"`
	oneOfType1    *BoolValue   `json:"-"`
}

func NewOneOfMetricDetailThresholdValue() *OneOfMetricDetailThresholdValue {
	p := new(OneOfMetricDetailThresholdValue)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfMetricDetailThresholdValue) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfMetricDetailThresholdValue is nil"))
	}
	switch v.(type) {
	case StringValue:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(StringValue)
		}
		*p.oneOfType0 = v.(StringValue)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case DoubleValue:
		if nil == p.oneOfType2 {
			p.oneOfType2 = new(DoubleValue)
		}
		*p.oneOfType2 = v.(DoubleValue)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2.ObjectType_
	case IntValue:
		if nil == p.oneOfType3 {
			p.oneOfType3 = new(IntValue)
		}
		*p.oneOfType3 = v.(IntValue)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType3.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType3.ObjectType_
	case BoolValue:
		if nil == p.oneOfType1 {
			p.oneOfType1 = new(BoolValue)
		}
		*p.oneOfType1 = v.(BoolValue)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType1.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType1.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfMetricDetailThresholdValue) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType2 != nil && *p.oneOfType2.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2
	}
	if p.oneOfType3 != nil && *p.oneOfType3.ObjectType_ == *p.Discriminator {
		return *p.oneOfType3
	}
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1
	}
	return nil
}

func (p *OneOfMetricDetailThresholdValue) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(StringValue)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "monitoring.v4.common.StringValue" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(StringValue)
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
	vOneOfType2 := new(DoubleValue)
	if err := json.Unmarshal(b, vOneOfType2); err == nil {
		if "monitoring.v4.common.DoubleValue" == *vOneOfType2.ObjectType_ {
			if nil == p.oneOfType2 {
				p.oneOfType2 = new(DoubleValue)
			}
			*p.oneOfType2 = *vOneOfType2
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2.ObjectType_
			return nil
		}
	}
	vOneOfType3 := new(IntValue)
	if err := json.Unmarshal(b, vOneOfType3); err == nil {
		if "monitoring.v4.common.IntValue" == *vOneOfType3.ObjectType_ {
			if nil == p.oneOfType3 {
				p.oneOfType3 = new(IntValue)
			}
			*p.oneOfType3 = *vOneOfType3
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType3.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType3.ObjectType_
			return nil
		}
	}
	vOneOfType1 := new(BoolValue)
	if err := json.Unmarshal(b, vOneOfType1); err == nil {
		if "monitoring.v4.common.BoolValue" == *vOneOfType1.ObjectType_ {
			if nil == p.oneOfType1 {
				p.oneOfType1 = new(BoolValue)
			}
			*p.oneOfType1 = *vOneOfType1
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType1.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType1.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfMetricDetailThresholdValue"))
}

func (p *OneOfMetricDetailThresholdValue) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType2 != nil && *p.oneOfType2.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2)
	}
	if p.oneOfType3 != nil && *p.oneOfType3.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType3)
	}
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	return nil, errors.New("No value to marshal for OneOfMetricDetailThresholdValue")
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
