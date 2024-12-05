/*
 * Generated file models/opsmgmt/v4/reporting/reporting_model.go.
 *
 * Product version: 4.0.1
 *
 * Part of the Nutanix Cloud Management Platform APIs
 *
 * (c) 2024 Nutanix Inc.  All rights reserved
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
	import2 "github.com/nutanix/ntnx-api-golang-clients/opsmgmt-go-client/v4/models/common/v1/response"
	import1 "github.com/nutanix/ntnx-api-golang-clients/opsmgmt-go-client/v4/models/opsmgmt/v4/error"
	import3 "github.com/nutanix/ntnx-api-golang-clients/opsmgmt-go-client/v4/models/prism/v4/config"
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
	return json.Marshal(struct {
		*CustomViewProxy
		EntityType *string     `json:"entityType,omitempty"`
		Format     *ViewFormat `json:"format,omitempty"`
	}{
		CustomViewProxy: (*CustomViewProxy)(p),
		EntityType:      p.EntityType,
		Format:          p.Format,
	})
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
REST response for all response codes in API path /opsmgmt/v4.0/reporting/views/{extId} Delete operation
*/
type DeleteReportViewApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteReportViewApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDeleteReportViewApiResponse() *DeleteReportViewApiResponse {
	p := new(DeleteReportViewApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.reporting.DeleteReportViewApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteReportViewApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteReportViewApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteReportViewApiResponseData()
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
REST response for all response codes in API path /opsmgmt/v4.0/reporting/config/$actions/export Post operation
*/
type ExportReportConfigApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfExportReportConfigApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewExportReportConfigApiResponse() *ExportReportConfigApiResponse {
	p := new(ExportReportConfigApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.reporting.ExportReportConfigApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ExportReportConfigApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ExportReportConfigApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfExportReportConfigApiResponseData()
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
	return json.Marshal(struct {
		*FieldProxy
		Name *string `json:"name,omitempty"`
	}{
		FieldProxy: (*FieldProxy)(p),
		Name:       p.Name,
	})
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
	return json.Marshal(struct {
		*FileWrapperProxy
		File *string `json:"file,omitempty"`
	}{
		FileWrapperProxy: (*FileWrapperProxy)(p),
		File:             p.File,
	})
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
REST response for all response codes in API path /opsmgmt/v4.0/reporting/views/$actions/generate-report Post operation
*/
type GenerateReportFromViewApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGenerateReportFromViewApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGenerateReportFromViewApiResponse() *GenerateReportFromViewApiResponse {
	p := new(GenerateReportFromViewApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.reporting.GenerateReportFromViewApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GenerateReportFromViewApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GenerateReportFromViewApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGenerateReportFromViewApiResponseData()
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
	return json.Marshal(struct {
		*ImportExportParamsProxy
		ReportConfigIds []string `json:"reportConfigIds,omitempty"`
	}{
		ImportExportParamsProxy: (*ImportExportParamsProxy)(p),
		ReportConfigIds:         p.ReportConfigIds,
	})
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
REST response for all response codes in API path /opsmgmt/v4.0/reporting/config/$actions/import Post operation
*/
type ImportReportConfigApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfImportReportConfigApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewImportReportConfigApiResponse() *ImportReportConfigApiResponse {
	p := new(ImportReportConfigApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.reporting.ImportReportConfigApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ImportReportConfigApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ImportReportConfigApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfImportReportConfigApiResponseData()
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
	return json.Marshal(struct {
		*NotificationPolicyProxy
		Recipients []Recipient `json:"recipients,omitempty"`
	}{
		NotificationPolicyProxy: (*NotificationPolicyProxy)(p),
		Recipients:              p.Recipients,
	})
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
REST response for all response codes in API path /opsmgmt/v4.0/reporting/views/$actions/notify Post operation
*/
type NotifyReportFromViewApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfNotifyReportFromViewApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewNotifyReportFromViewApiResponse() *NotifyReportFromViewApiResponse {
	p := new(NotifyReportFromViewApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.reporting.NotifyReportFromViewApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *NotifyReportFromViewApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *NotifyReportFromViewApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfNotifyReportFromViewApiResponseData()
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
	return json.Marshal(struct {
		*RecipientProxy
		EmailAddress *string `json:"emailAddress,omitempty"`
	}{
		RecipientProxy: (*RecipientProxy)(p),
		EmailAddress:   p.EmailAddress,
	})
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
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Name of the view.
	*/
	Name *string `json:"name"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
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
	return json.Marshal(struct {
		*ViewProxy
		Name *string         `json:"name,omitempty"`
		Type *OneOfViewType_ `json:"type,omitempty"`
	}{
		ViewProxy: (*ViewProxy)(p),
		Name:      p.Name,
		Type:      p.Type,
	})
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
REST response for all response codes in API path /opsmgmt/v4.0/reporting/views Post operation
*/
type ViewApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfViewApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewViewApiResponse() *ViewApiResponse {
	p := new(ViewApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.reporting.ViewApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ViewApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ViewApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfViewApiResponseData()
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
REST response for all response codes in API path /opsmgmt/v4.0/reporting/views/$actions/create-config Post operation
*/
type ViewConfigApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfViewConfigApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewViewConfigApiResponse() *ViewConfigApiResponse {
	p := new(ViewConfigApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.reporting.ViewConfigApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ViewConfigApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ViewConfigApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfViewConfigApiResponseData()
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
	Links []import2.ApiLink `json:"links,omitempty"`
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
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  List of views used to generate a report. Number of views must be in the range of 1 to 50.
	*/
	Views []View `json:"views"`
}

func (p *ViewReport) MarshalJSON() ([]byte, error) {
	type ViewReportProxy ViewReport
	return json.Marshal(struct {
		*ViewReportProxy
		Name  *string `json:"name,omitempty"`
		Views []View  `json:"views,omitempty"`
	}{
		ViewReportProxy: (*ViewReportProxy)(p),
		Name:            p.Name,
		Views:           p.Views,
	})
}

func NewViewReport() *ViewReport {
	p := new(ViewReport)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.reporting.ViewReport"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type OneOfGenerateReportFromViewApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
}

func NewOneOfGenerateReportFromViewApiResponseData() *OneOfGenerateReportFromViewApiResponseData {
	p := new(OneOfGenerateReportFromViewApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGenerateReportFromViewApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGenerateReportFromViewApiResponseData is nil"))
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

func (p *OneOfGenerateReportFromViewApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGenerateReportFromViewApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
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
		if "opsmgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGenerateReportFromViewApiResponseData"))
}

func (p *OneOfGenerateReportFromViewApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGenerateReportFromViewApiResponseData")
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

type OneOfViewApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *View                  `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
}

func NewOneOfViewApiResponseData() *OneOfViewApiResponseData {
	p := new(OneOfViewApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfViewApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfViewApiResponseData is nil"))
	}
	switch v.(type) {
	case View:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(View)
		}
		*p.oneOfType0 = v.(View)
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

func (p *OneOfViewApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfViewApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(View)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "opsmgmt.v4.reporting.View" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(View)
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
		if "opsmgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfViewApiResponseData"))
}

func (p *OneOfViewApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfViewApiResponseData")
}

type OneOfViewConfigApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
}

func NewOneOfViewConfigApiResponseData() *OneOfViewConfigApiResponseData {
	p := new(OneOfViewConfigApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfViewConfigApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfViewConfigApiResponseData is nil"))
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

func (p *OneOfViewConfigApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfViewConfigApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
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
		if "opsmgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfViewConfigApiResponseData"))
}

func (p *OneOfViewConfigApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfViewConfigApiResponseData")
}

type OneOfExportReportConfigApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *FileWrapper           `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
}

func NewOneOfExportReportConfigApiResponseData() *OneOfExportReportConfigApiResponseData {
	p := new(OneOfExportReportConfigApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfExportReportConfigApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfExportReportConfigApiResponseData is nil"))
	}
	switch v.(type) {
	case FileWrapper:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(FileWrapper)
		}
		*p.oneOfType0 = v.(FileWrapper)
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

func (p *OneOfExportReportConfigApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfExportReportConfigApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(FileWrapper)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "opsmgmt.v4.reporting.FileWrapper" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(FileWrapper)
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
		if "opsmgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfExportReportConfigApiResponseData"))
}

func (p *OneOfExportReportConfigApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfExportReportConfigApiResponseData")
}

type OneOfNotifyReportFromViewApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
}

func NewOneOfNotifyReportFromViewApiResponseData() *OneOfNotifyReportFromViewApiResponseData {
	p := new(OneOfNotifyReportFromViewApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfNotifyReportFromViewApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfNotifyReportFromViewApiResponseData is nil"))
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

func (p *OneOfNotifyReportFromViewApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfNotifyReportFromViewApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
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
		if "opsmgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfNotifyReportFromViewApiResponseData"))
}

func (p *OneOfNotifyReportFromViewApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfNotifyReportFromViewApiResponseData")
}

type OneOfDeleteReportViewApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType1    *interface{}           `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
}

func NewOneOfDeleteReportViewApiResponseData() *OneOfDeleteReportViewApiResponseData {
	p := new(OneOfDeleteReportViewApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteReportViewApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteReportViewApiResponseData is nil"))
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

func (p *OneOfDeleteReportViewApiResponseData) GetValue() interface{} {
	if "EMPTY" == *p.Discriminator {
		return *p.oneOfType1
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDeleteReportViewApiResponseData) UnmarshalJSON(b []byte) error {
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
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "opsmgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteReportViewApiResponseData"))
}

func (p *OneOfDeleteReportViewApiResponseData) MarshalJSON() ([]byte, error) {
	if "EMPTY" == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteReportViewApiResponseData")
}

type OneOfImportReportConfigApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *ImportExportParams    `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
}

func NewOneOfImportReportConfigApiResponseData() *OneOfImportReportConfigApiResponseData {
	p := new(OneOfImportReportConfigApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfImportReportConfigApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfImportReportConfigApiResponseData is nil"))
	}
	switch v.(type) {
	case ImportExportParams:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(ImportExportParams)
		}
		*p.oneOfType0 = v.(ImportExportParams)
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

func (p *OneOfImportReportConfigApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfImportReportConfigApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(ImportExportParams)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "opsmgmt.v4.reporting.ImportExportParams" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(ImportExportParams)
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
		if "opsmgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfImportReportConfigApiResponseData"))
}

func (p *OneOfImportReportConfigApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfImportReportConfigApiResponseData")
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
