/*
 * Generated file models/opsmgmt/v4/reporting/reporting_model.go.
 *
 * Product version: 4.0.3
 *
 * Part of the Nutanix Cloud Management Platform APIs
 *
 * (c) 2025 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module opsmgmt.v4.reporting of Nutanix Cloud Management Platform APIs
*/
package reporting

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	import1 "github.com/nutanix/ntnx-api-golang-clients/opsmgmt-go-client/v4/models/common/v1/response"
)

/*
Custom View Parameters used to decide what data needs to be shown in the view.
*/
type CustomView struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Entity UUIDs used in the view.
	*/
	EntityIds []string `json:"entityIds,omitempty"`
	/*
	  Entity type for which data needs to be shown in the view.
	*/
	EntityType *string `json:"entityType"`
	/*
	  List of selected fields for the view.
	*/
	Fields []Field `json:"fields,omitempty"`
	/*
	  OData criteria that will be used to filter the returned data.
	*/
	FilterCriteria *string `json:"filterCriteria,omitempty"`

	Format *ViewFormat `json:"format"`
	/*
	  Column on which grouping to be used while getting data for the view.
	*/
	GroupByColumn *string `json:"groupByColumn,omitempty"`
	/*
	  Limit on the maximum number of entities to be represented in the view.
	*/
	Limit *int64 `json:"limit,omitempty"`
	/*
	  Offset to be used while getting data for the view.
	*/
	Offset *int64 `json:"offset,omitempty"`
	/*
	  Entity Property based on which the result data is to be sorted.
	*/
	SortColumn *string `json:"sortColumn,omitempty"`

	SortOrder *SortOrder `json:"sortOrder,omitempty"`
}

func (p *CustomView) MarshalJSON() ([]byte, error) {
	type CustomViewProxy CustomView

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*CustomViewProxy
		EntityType *string     `json:"entityType,omitempty"`
		Format     *ViewFormat `json:"format,omitempty"`
	}{
		CustomViewProxy: (*CustomViewProxy)(p),
		EntityType:      p.EntityType,
		Format:          p.Format,
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

func (p *CustomView) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias CustomView
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewCustomView()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.EntityIds != nil {
		p.EntityIds = known.EntityIds
	}
	if known.EntityType != nil {
		p.EntityType = known.EntityType
	}
	if known.Fields != nil {
		p.Fields = known.Fields
	}
	if known.FilterCriteria != nil {
		p.FilterCriteria = known.FilterCriteria
	}
	if known.Format != nil {
		p.Format = known.Format
	}
	if known.GroupByColumn != nil {
		p.GroupByColumn = known.GroupByColumn
	}
	if known.Limit != nil {
		p.Limit = known.Limit
	}
	if known.Offset != nil {
		p.Offset = known.Offset
	}
	if known.SortColumn != nil {
		p.SortColumn = known.SortColumn
	}
	if known.SortOrder != nil {
		p.SortOrder = known.SortOrder
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "entityIds")
	delete(allFields, "entityType")
	delete(allFields, "fields")
	delete(allFields, "filterCriteria")
	delete(allFields, "format")
	delete(allFields, "groupByColumn")
	delete(allFields, "limit")
	delete(allFields, "offset")
	delete(allFields, "sortColumn")
	delete(allFields, "sortOrder")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewCustomView() *CustomView {
	p := new(CustomView)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.reporting.CustomView"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Description of a field in the view. Eg: CPU Usage in a VM data table.
*/
type Field struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	AggregationOperator *Operator `json:"aggregationOperator,omitempty"`
	/*
	  Human-readable label for widget field.
	*/
	Label *string `json:"label,omitempty"`
	/*
	  Entity attribute/metric to be selected for the widget.
	*/
	Name *string `json:"name"`
}

func (p *Field) MarshalJSON() ([]byte, error) {
	type FieldProxy Field

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*FieldProxy
		Name *string `json:"name,omitempty"`
	}{
		FieldProxy: (*FieldProxy)(p),
		Name:       p.Name,
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

func (p *Field) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Field
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewField()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.AggregationOperator != nil {
		p.AggregationOperator = known.AggregationOperator
	}
	if known.Label != nil {
		p.Label = known.Label
	}
	if known.Name != nil {
		p.Name = known.Name
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "aggregationOperator")
	delete(allFields, "label")
	delete(allFields, "name")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewField() *Field {
	p := new(Field)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.reporting.Field"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The wrapper model for the binary file to be exported.
*/
type FileWrapper struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The binary of the exported report configuration(s). Default value is report-{currTime}. For eg:- report-Jan23-062758
	*/
	File *string `json:"file"`
}

func (p *FileWrapper) MarshalJSON() ([]byte, error) {
	type FileWrapperProxy FileWrapper

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*FileWrapperProxy
		File *string `json:"file,omitempty"`
	}{
		FileWrapperProxy: (*FileWrapperProxy)(p),
		File:             p.File,
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

func (p *FileWrapper) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias FileWrapper
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewFileWrapper()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.File != nil {
		p.File = known.File
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "file")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewFileWrapper() *FileWrapper {
	p := new(FileWrapper)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.reporting.FileWrapper"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Contains report configuration(s) UUIDs of those imported or to be exported.
*/
type ImportExportParams struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  UUID(s) of the report configuration to export
	*/
	ReportConfigIds []string `json:"reportConfigIds"`
}

func (p *ImportExportParams) MarshalJSON() ([]byte, error) {
	type ImportExportParamsProxy ImportExportParams

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*ImportExportParamsProxy
		ReportConfigIds []string `json:"reportConfigIds,omitempty"`
	}{
		ImportExportParamsProxy: (*ImportExportParamsProxy)(p),
		ReportConfigIds:         p.ReportConfigIds,
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

func (p *ImportExportParams) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ImportExportParams
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewImportExportParams()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ReportConfigIds != nil {
		p.ReportConfigIds = known.ReportConfigIds
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "reportConfigIds")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewImportExportParams() *ImportExportParams {
	p := new(ImportExportParams)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.reporting.ImportExportParams"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
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
	if known.Recipients != nil {
		p.Recipients = known.Recipients
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "emailBody")
	delete(allFields, "emailSubject")
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
	*p.ObjectType_ = "opsmgmt.v4.reporting.NotificationPolicy"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Used to aggregate field data from multiple values across time range.
*/
type Operator int

const (
	OPERATOR_UNKNOWN  Operator = 0
	OPERATOR_REDACTED Operator = 1
	OPERATOR_SUM      Operator = 2
	OPERATOR_MAX      Operator = 3
	OPERATOR_MIN      Operator = 4
	OPERATOR_AVG      Operator = 5
	OPERATOR_LAST     Operator = 6
	OPERATOR_COUNT    Operator = 7
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *Operator) name(index int) string {
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
func (e Operator) GetName() string {
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
func (e *Operator) index(name string) Operator {
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
			return Operator(idx)
		}
	}
	return OPERATOR_UNKNOWN
}

func (e *Operator) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for Operator:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *Operator) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e Operator) Ref() *Operator {
	return &e
}

/*
Different predefined view types supported. Allowed values are VM_EFFICIENCY, STORAGE_OVERPROVISIONING, STORAGE_SUMMARY, REBUILD_CAPACITY_RESERVATION, RECYCLE_BIN, CLUSTER_CPU_USAGE, CLUSTER_LATENCY, CLUSTER_MEMORY_USAGE, CLUSTER_QUICK_ACCESS, CLUSTER_RUNWAY, CLUSTER_STORAGE, CONTROLLER_IOPS, PERFORMANCE, ALERTS, TASKS, DISCOVERED_APPS, SECURITY, PLAYS, REPORTS, DATA_RECOVERY_STATUS.
*/
type PredefinedType int

const (
	PREDEFINEDTYPE_UNKNOWN                      PredefinedType = 0
	PREDEFINEDTYPE_REDACTED                     PredefinedType = 1
	PREDEFINEDTYPE_VM_EFFICIENCY                PredefinedType = 2
	PREDEFINEDTYPE_STORAGE_OVER_PROVISIONING    PredefinedType = 3
	PREDEFINEDTYPE_STORAGE_SUMMARY              PredefinedType = 4
	PREDEFINEDTYPE_REBUILD_CAPACITY_RESERVATION PredefinedType = 5
	PREDEFINEDTYPE_RECYCLE_BIN                  PredefinedType = 6
	PREDEFINEDTYPE_CLUSTER_CPU_USAGE            PredefinedType = 7
	PREDEFINEDTYPE_CLUSTER_LATENCY              PredefinedType = 8
	PREDEFINEDTYPE_CLUSTER_MEMORY_USAGE         PredefinedType = 9
	PREDEFINEDTYPE_CLUSTER_QUICK_ACCESS         PredefinedType = 10
	PREDEFINEDTYPE_CLUSTER_RUNWAY               PredefinedType = 11
	PREDEFINEDTYPE_CLUSTER_STORAGE              PredefinedType = 12
	PREDEFINEDTYPE_CONTROLLER_IOPS              PredefinedType = 13
	PREDEFINEDTYPE_PERFORMANCE                  PredefinedType = 14
	PREDEFINEDTYPE_ALERTS                       PredefinedType = 15
	PREDEFINEDTYPE_TASKS                        PredefinedType = 16
	PREDEFINEDTYPE_DISCOVERED_APPS              PredefinedType = 17
	PREDEFINEDTYPE_SECURITY                     PredefinedType = 18
	PREDEFINEDTYPE_PLAYS                        PredefinedType = 19
	PREDEFINEDTYPE_REPORTS                      PredefinedType = 20
	PREDEFINEDTYPE_DATA_RECOVERY_STATUS         PredefinedType = 21
	PREDEFINEDTYPE_IMPACTED_CLUSTER             PredefinedType = 22
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *PredefinedType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"VM_EFFICIENCY",
		"STORAGE_OVER-PROVISIONING",
		"STORAGE_SUMMARY",
		"REBUILD_CAPACITY_RESERVATION",
		"RECYCLE_BIN",
		"CLUSTER_CPU_USAGE",
		"CLUSTER_LATENCY",
		"CLUSTER_MEMORY_USAGE",
		"CLUSTER_QUICK_ACCESS",
		"CLUSTER_RUNWAY",
		"CLUSTER_STORAGE",
		"CONTROLLER_IOPS",
		"PERFORMANCE",
		"ALERTS",
		"TASKS",
		"DISCOVERED_APPS",
		"SECURITY",
		"PLAYS",
		"REPORTS",
		"DATA_RECOVERY_STATUS",
		"IMPACTED_CLUSTER",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e PredefinedType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"VM_EFFICIENCY",
		"STORAGE_OVER-PROVISIONING",
		"STORAGE_SUMMARY",
		"REBUILD_CAPACITY_RESERVATION",
		"RECYCLE_BIN",
		"CLUSTER_CPU_USAGE",
		"CLUSTER_LATENCY",
		"CLUSTER_MEMORY_USAGE",
		"CLUSTER_QUICK_ACCESS",
		"CLUSTER_RUNWAY",
		"CLUSTER_STORAGE",
		"CONTROLLER_IOPS",
		"PERFORMANCE",
		"ALERTS",
		"TASKS",
		"DISCOVERED_APPS",
		"SECURITY",
		"PLAYS",
		"REPORTS",
		"DATA_RECOVERY_STATUS",
		"IMPACTED_CLUSTER",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *PredefinedType) index(name string) PredefinedType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"VM_EFFICIENCY",
		"STORAGE_OVER-PROVISIONING",
		"STORAGE_SUMMARY",
		"REBUILD_CAPACITY_RESERVATION",
		"RECYCLE_BIN",
		"CLUSTER_CPU_USAGE",
		"CLUSTER_LATENCY",
		"CLUSTER_MEMORY_USAGE",
		"CLUSTER_QUICK_ACCESS",
		"CLUSTER_RUNWAY",
		"CLUSTER_STORAGE",
		"CONTROLLER_IOPS",
		"PERFORMANCE",
		"ALERTS",
		"TASKS",
		"DISCOVERED_APPS",
		"SECURITY",
		"PLAYS",
		"REPORTS",
		"DATA_RECOVERY_STATUS",
		"IMPACTED_CLUSTER",
	}
	for idx := range names {
		if names[idx] == name {
			return PredefinedType(idx)
		}
	}
	return PREDEFINEDTYPE_UNKNOWN
}

func (e *PredefinedType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for PredefinedType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *PredefinedType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e PredefinedType) Ref() *PredefinedType {
	return &e
}

type PredefinedView struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	PredefinedView *PredefinedType `json:"predefinedView,omitempty"`
}

func (p *PredefinedView) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias PredefinedView

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

func (p *PredefinedView) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias PredefinedView
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewPredefinedView()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.PredefinedView != nil {
		p.PredefinedView = known.PredefinedView
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "predefinedView")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewPredefinedView() *PredefinedView {
	p := new(PredefinedView)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.reporting.PredefinedView"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
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
	*p.ObjectType_ = "opsmgmt.v4.reporting.Recipient"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
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
	*p.ObjectType_ = "opsmgmt.v4.reporting.ReportCustomization"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
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
View creation/modification spec.
*/
type View struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Heading for the view, used while generating reports.
	*/
	Heading *string `json:"heading,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  Name of the view.
	*/
	Name *string `json:"name"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*

	 */
	TypeItemDiscriminator_ *string `json:"$typeItemDiscriminator,omitempty"`
	/*
	  Type for the view, which can take custom values or set of predefined views.
	*/
	Type *OneOfViewType_ `json:"type"`
}

func (p *View) MarshalJSON() ([]byte, error) {
	type ViewProxy View

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*ViewProxy
		Name *string         `json:"name,omitempty"`
		Type *OneOfViewType_ `json:"type,omitempty"`
	}{
		ViewProxy: (*ViewProxy)(p),
		Name:      p.Name,
		Type:      p.Type,
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

func (p *View) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias View
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewView()

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
	if known.Heading != nil {
		p.Heading = known.Heading
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}
	if known.TypeItemDiscriminator_ != nil {
		p.TypeItemDiscriminator_ = known.TypeItemDiscriminator_
	}
	if known.Type != nil {
		p.Type = known.Type
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "extId")
	delete(allFields, "heading")
	delete(allFields, "links")
	delete(allFields, "name")
	delete(allFields, "tenantId")
	delete(allFields, "$typeItemDiscriminator")
	delete(allFields, "type")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewView() *View {
	p := new(View)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.reporting.View"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *View) GetType() interface{} {
	if nil == p.Type {
		return nil
	}
	return p.Type.GetValue()
}

func (p *View) SetType(v interface{}) error {
	if nil == p.Type {
		p.Type = NewOneOfViewType_()
	}
	e := p.Type.SetValue(v)
	if nil == e {
		if nil == p.TypeItemDiscriminator_ {
			p.TypeItemDiscriminator_ = new(string)
		}
		*p.TypeItemDiscriminator_ = *p.Type.Discriminator
	}
	return e
}

/*
Format of the data in the view. Allowed values are LINE_CHART, DATA_TABLE, HISTOGRAM, METRIC_SUMMARY_TEXT, COUNT_SUMMARY, TEXT.
*/
type ViewFormat int

const (
	VIEWFORMAT_UNKNOWN             ViewFormat = 0
	VIEWFORMAT_REDACTED            ViewFormat = 1
	VIEWFORMAT_LINE_CHART          ViewFormat = 2
	VIEWFORMAT_DATA_TABLE          ViewFormat = 3
	VIEWFORMAT_HISTOGRAM           ViewFormat = 4
	VIEWFORMAT_METRIC_SUMMARY_TEXT ViewFormat = 5
	VIEWFORMAT_COUNT_SUMMARY       ViewFormat = 6
	VIEWFORMAT_TEXT                ViewFormat = 7
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ViewFormat) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"LINE_CHART",
		"DATA_TABLE",
		"HISTOGRAM",
		"METRIC_SUMMARY_TEXT",
		"COUNT_SUMMARY",
		"TEXT",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ViewFormat) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"LINE_CHART",
		"DATA_TABLE",
		"HISTOGRAM",
		"METRIC_SUMMARY_TEXT",
		"COUNT_SUMMARY",
		"TEXT",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ViewFormat) index(name string) ViewFormat {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"LINE_CHART",
		"DATA_TABLE",
		"HISTOGRAM",
		"METRIC_SUMMARY_TEXT",
		"COUNT_SUMMARY",
		"TEXT",
	}
	for idx := range names {
		if names[idx] == name {
			return ViewFormat(idx)
		}
	}
	return VIEWFORMAT_UNKNOWN
}

func (e *ViewFormat) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ViewFormat:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ViewFormat) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ViewFormat) Ref() *ViewFormat {
	return &e
}

/*
Generated report instance from given list of views.
*/
type ViewReport struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The end time of the period for which stats should be reported. The value should be in extended ISO-8601 format. For example, end time of 2022-04-23T013:23:45.678+09:00 would consider all stats till 13:23:45 .678 on the 23rd of April 2022. Details around ISO-8601 format can be found at https://www.iso.org/standard/70907.html
	*/
	EndTime *string `json:"endTime,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  Name of the report generated.
	*/
	Name *string `json:"name"`

	NotificationPolicy *NotificationPolicy `json:"notificationPolicy,omitempty"`
	/*
	  List specifying the formats in which the report will be created.
	*/
	ReportFormat []ReportFormat `json:"reportFormat,omitempty"`
	/*
	  The start time of the period for which stats should be reported. The value should be in extended ISO-8601 format. For example, start time of 2022-04-23T01:23:45.678+09:00 would consider all stats starting at 1:23:45.678 on the 23rd of April 2022. Details around ISO-8601 format can be found at https://www.iso.org/standard/70907.html
	*/
	StartTime *string `json:"startTime,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  List of views used to generate a report. Number of views must be in the range of 1 to 50.
	*/
	Views []View `json:"views"`
}

func (p *ViewReport) MarshalJSON() ([]byte, error) {
	type ViewReportProxy ViewReport

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*ViewReportProxy
		Name  *string `json:"name,omitempty"`
		Views []View  `json:"views,omitempty"`
	}{
		ViewReportProxy: (*ViewReportProxy)(p),
		Name:            p.Name,
		Views:           p.Views,
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

func (p *ViewReport) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ViewReport
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewViewReport()

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
	if known.ReportFormat != nil {
		p.ReportFormat = known.ReportFormat
	}
	if known.StartTime != nil {
		p.StartTime = known.StartTime
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}
	if known.Views != nil {
		p.Views = known.Views
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "endTime")
	delete(allFields, "extId")
	delete(allFields, "links")
	delete(allFields, "name")
	delete(allFields, "notificationPolicy")
	delete(allFields, "reportFormat")
	delete(allFields, "startTime")
	delete(allFields, "tenantId")
	delete(allFields, "views")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewViewReport() *ViewReport {
	p := new(ViewReport)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.reporting.ViewReport"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type OneOfViewType_ struct {
	Discriminator *string         `json:"-"`
	ObjectType_   *string         `json:"-"`
	oneOfType1    *PredefinedView `json:"-"`
	oneOfType0    *CustomView     `json:"-"`
}

func NewOneOfViewType_() *OneOfViewType_ {
	p := new(OneOfViewType_)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfViewType_) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfViewType_ is nil"))
	}
	switch v.(type) {
	case PredefinedView:
		if nil == p.oneOfType1 {
			p.oneOfType1 = new(PredefinedView)
		}
		*p.oneOfType1 = v.(PredefinedView)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType1.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType1.ObjectType_
	case CustomView:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(CustomView)
		}
		*p.oneOfType0 = v.(CustomView)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfViewType_) GetValue() interface{} {
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfViewType_) UnmarshalJSON(b []byte) error {
	vOneOfType1 := new(PredefinedView)
	if err := json.Unmarshal(b, vOneOfType1); err == nil {
		if "opsmgmt.v4.reporting.PredefinedView" == *vOneOfType1.ObjectType_ {
			if nil == p.oneOfType1 {
				p.oneOfType1 = new(PredefinedView)
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
	vOneOfType0 := new(CustomView)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "opsmgmt.v4.reporting.CustomView" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(CustomView)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfViewType_"))
}

func (p *OneOfViewType_) MarshalJSON() ([]byte, error) {
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfViewType_")
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
