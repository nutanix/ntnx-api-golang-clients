/*
 * Generated file models/opsmgmt/v4/config/config_model.go.
 *
 * Product version: 4.0.1
 *
 * Part of the Nutanix Cloud Management Platform APIs
 *
 * (c) 2024 Nutanix Inc.  All rights reserved
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
	import3 "github.com/nutanix/ntnx-api-golang-clients/opsmgmt-go-client/v4/models/common/v1/response"
	import2 "github.com/nutanix/ntnx-api-golang-clients/opsmgmt-go-client/v4/models/opsmgmt/v4/error"
	import1 "github.com/nutanix/ntnx-api-golang-clients/opsmgmt-go-client/v4/models/prism/v4/config"
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
REST response for all response codes in API path /opsmgmt/v4.0/config/reports Post operation
*/
type CreateReportApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateReportApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewCreateReportApiResponse() *CreateReportApiResponse {
	p := new(CreateReportApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.config.CreateReportApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
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
REST response for all response codes in API path /opsmgmt/v4.0/config/report-configs Post operation
*/
type CreateReportConfigApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateReportConfigApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewCreateReportConfigApiResponse() *CreateReportConfigApiResponse {
	p := new(CreateReportConfigApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.config.CreateReportConfigApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
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

func NewDataCriteria() *DataCriteria {
	p := new(DataCriteria)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.config.DataCriteria"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /opsmgmt/v4.0/config/reports/{extId} Delete operation
*/
type DeleteReportApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteReportApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDeleteReportApiResponse() *DeleteReportApiResponse {
	p := new(DeleteReportApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.config.DeleteReportApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
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
REST response for all response codes in API path /opsmgmt/v4.0/config/report-configs/{extId} Delete operation
*/
type DeleteReportConfigApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteReportConfigApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDeleteReportConfigApiResponse() *DeleteReportConfigApiResponse {
	p := new(DeleteReportConfigApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.config.DeleteReportConfigApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
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
	return json.Marshal(struct {
		*EntitySelectionProxy
		EntityExtId []string    `json:"entityExtId,omitempty"`
		EntityType  *EntityType `json:"entityType,omitempty"`
	}{
		EntitySelectionProxy: (*EntitySelectionProxy)(p),
		EntityExtId:          p.EntityExtId,
		EntityType:           p.EntityType,
	})
}

func NewEntitySelection() *EntitySelection {
	p := new(EntitySelection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.config.EntitySelection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
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
REST response for all response codes in API path /opsmgmt/v4.0/config/user/{userExtId}/global-report-setting Get operation
*/
type GetGlobalReportSettingApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetGlobalReportSettingApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetGlobalReportSettingApiResponse() *GetGlobalReportSettingApiResponse {
	p := new(GetGlobalReportSettingApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.config.GetGlobalReportSettingApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
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
REST response for all response codes in API path /opsmgmt/v4.0/config/reports/{extId} Get operation
*/
type GetReportApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetReportApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetReportApiResponse() *GetReportApiResponse {
	p := new(GetReportApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.config.GetReportApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
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
REST response for all response codes in API path /opsmgmt/v4.0/config/report-configs/{extId} Get operation
*/
type GetReportConfigApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetReportConfigApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetReportConfigApiResponse() *GetReportConfigApiResponse {
	p := new(GetReportConfigApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.config.GetReportConfigApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
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
	Links []import3.ApiLink `json:"links,omitempty"`
	/*
	  Name of the global report setting.
	*/
	Name *string `json:"name"`

	NotificationPolicy *NotificationPolicy `json:"notificationPolicy,omitempty"`

	ReportCustomization *ReportCustomization `json:"reportCustomization,omitempty"`

	RetentionConfig *RetentionConfig `json:"retentionConfig,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *GlobalReportSetting) MarshalJSON() ([]byte, error) {
	type GlobalReportSettingProxy GlobalReportSetting
	return json.Marshal(struct {
		*GlobalReportSettingProxy
		Name *string `json:"name,omitempty"`
	}{
		GlobalReportSettingProxy: (*GlobalReportSettingProxy)(p),
		Name:                     p.Name,
	})
}

func NewGlobalReportSetting() *GlobalReportSetting {
	p := new(GlobalReportSetting)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.config.GlobalReportSetting"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /opsmgmt/v4.0/config/report-configs Get operation
*/
type ListReportConfigsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListReportConfigsApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListReportConfigsApiResponse() *ListReportConfigsApiResponse {
	p := new(ListReportConfigsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.config.ListReportConfigsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
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
REST response for all response codes in API path /opsmgmt/v4.0/config/reports Get operation
*/
type ListReportsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListReportsApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListReportsApiResponse() *ListReportsApiResponse {
	p := new(ListReportsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.config.ListReportsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
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
	*p.ObjectType_ = "opsmgmt.v4.config.NotificationPolicy"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /opsmgmt/v4.0/config/reports/{extId}/$actions/notify-recipients Post operation
*/
type NotifyReportApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfNotifyReportApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewNotifyReportApiResponse() *NotifyReportApiResponse {
	p := new(NotifyReportApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.config.NotifyReportApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
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
	*p.ObjectType_ = "opsmgmt.v4.config.Recipient"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
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
	return json.Marshal(struct {
		*RepeatCriteriaProxy
		EntityType *EntityType `json:"entityType,omitempty"`
	}{
		RepeatCriteriaProxy: (*RepeatCriteriaProxy)(p),
		EntityType:          p.EntityType,
	})
}

func NewRepeatCriteria() *RepeatCriteria {
	p := new(RepeatCriteria)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.config.RepeatCriteria"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
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
	Links []import3.ApiLink `json:"links,omitempty"`
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
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  Timezone in which report is to be generated. This is the list supported by pytz.all_timezones. For more info, check http://pytz.sourceforge.net
	*/
	Timezone *string `json:"timezone,omitempty"`
}

func (p *Report) MarshalJSON() ([]byte, error) {
	type ReportProxy Report
	return json.Marshal(struct {
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
	})
}

func NewReport() *Report {
	p := new(Report)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.config.Report"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
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
	Links []import3.ApiLink `json:"links,omitempty"`
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
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  The timezone in which the report will be generated. This is the list supported by pytz.all_timezones. For more info, check http://pytz.sourceforge.net
	*/
	Timezone *string `json:"timezone,omitempty"`
}

func (p *ReportConfig) MarshalJSON() ([]byte, error) {
	type ReportConfigProxy ReportConfig
	return json.Marshal(struct {
		*ReportConfigProxy
		Name     *string   `json:"name,omitempty"`
		Sections []Section `json:"sections,omitempty"`
	}{
		ReportConfigProxy: (*ReportConfigProxy)(p),
		Name:              p.Name,
		Sections:          p.Sections,
	})
}

func NewReportConfig() *ReportConfig {
	p := new(ReportConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.config.ReportConfig"
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
	*p.ObjectType_ = "opsmgmt.v4.config.ReportCustomization"
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
	Links []import3.ApiLink `json:"links,omitempty"`
	/*
	  List specifying the formats in which the report is to be sent.
	*/
	RecipientFormats []ReportFormat `json:"recipientFormats"`
	/*
	  Recipients to notify with the report.
	*/
	Recipients []Recipient `json:"recipients"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *ReportNotificationSpec) MarshalJSON() ([]byte, error) {
	type ReportNotificationSpecProxy ReportNotificationSpec
	return json.Marshal(struct {
		*ReportNotificationSpecProxy
		RecipientFormats []ReportFormat `json:"recipientFormats,omitempty"`
		Recipients       []Recipient    `json:"recipients,omitempty"`
	}{
		ReportNotificationSpecProxy: (*ReportNotificationSpecProxy)(p),
		RecipientFormats:            p.RecipientFormats,
		Recipients:                  p.Recipients,
	})
}

func NewReportNotificationSpec() *ReportNotificationSpec {
	p := new(ReportNotificationSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.config.ReportNotificationSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
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
	return json.Marshal(struct {
		*ReportScheduleProxy
		Frequency        *int              `json:"frequency,omitempty"`
		ScheduleInterval *ScheduleInterval `json:"scheduleInterval,omitempty"`
		StartTime        *time.Time        `json:"startTime,omitempty"`
	}{
		ReportScheduleProxy: (*ReportScheduleProxy)(p),
		Frequency:           p.Frequency,
		ScheduleInterval:    p.ScheduleInterval,
		StartTime:           p.StartTime,
	})
}

func NewReportSchedule() *ReportSchedule {
	p := new(ReportSchedule)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.config.ReportSchedule"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
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

func NewRetentionConfig() *RetentionConfig {
	p := new(RetentionConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.config.RetentionConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
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
	return json.Marshal(struct {
		*RowProxy
		Widgets []Widget `json:"widgets,omitempty"`
	}{
		RowProxy: (*RowProxy)(p),
		Widgets:  p.Widgets,
	})
}

func NewRow() *Row {
	p := new(Row)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.config.Row"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
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
	return json.Marshal(struct {
		*SectionProxy
		Name *string `json:"name,omitempty"`
		Rows []Row   `json:"rows,omitempty"`
	}{
		SectionProxy: (*SectionProxy)(p),
		Name:         p.Name,
		Rows:         p.Rows,
	})
}

func NewSection() *Section {
	p := new(Section)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.config.Section"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
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
REST response for all response codes in API path /opsmgmt/v4.0/config/user/{userExtId}/global-report-setting Put operation
*/
type UpdateGlobalReportSettingsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateGlobalReportSettingsApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewUpdateGlobalReportSettingsApiResponse() *UpdateGlobalReportSettingsApiResponse {
	p := new(UpdateGlobalReportSettingsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.config.UpdateGlobalReportSettingsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
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
REST response for all response codes in API path /opsmgmt/v4.0/config/report-configs/{extId} Put operation
*/
type UpdateReportConfigApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateReportConfigApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewUpdateReportConfigApiResponse() *UpdateReportConfigApiResponse {
	p := new(UpdateReportConfigApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.config.UpdateReportConfigApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
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
	return json.Marshal(struct {
		*WidgetProxy
		WidgetInfo *OneOfWidgetWidgetInfo `json:"widgetInfo,omitempty"`
	}{
		WidgetProxy: (*WidgetProxy)(p),
		WidgetInfo:  p.WidgetInfo,
	})
}

func NewWidget() *Widget {
	p := new(Widget)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.config.Widget"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
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
}

func (p *WidgetConfig) MarshalJSON() ([]byte, error) {
	type WidgetConfigProxy WidgetConfig
	return json.Marshal(struct {
		*WidgetConfigProxy
		Size *WidgetSize `json:"size,omitempty"`
		Type *WidgetType `json:"type,omitempty"`
	}{
		WidgetConfigProxy: (*WidgetConfigProxy)(p),
		Size:              p.Size,
		Type:              p.Type,
	})
}

func NewWidgetConfig() *WidgetConfig {
	p := new(WidgetConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.config.WidgetConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
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
	/*
	  Human-readable label for widget field.
	*/
	Label *string `json:"label"`
	/*
	  Entity attribute/metric to be selected for the widget.
	*/
	Name *string `json:"name"`
}

func (p *WidgetField) MarshalJSON() ([]byte, error) {
	type WidgetFieldProxy WidgetField
	return json.Marshal(struct {
		*WidgetFieldProxy
		Label *string `json:"label,omitempty"`
		Name  *string `json:"name,omitempty"`
	}{
		WidgetFieldProxy: (*WidgetFieldProxy)(p),
		Label:            p.Label,
		Name:             p.Name,
	})
}

func NewWidgetField() *WidgetField {
	p := new(WidgetField)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.config.WidgetField"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
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

func NewWidgetTemplate() *WidgetTemplate {
	p := new(WidgetTemplate)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "opsmgmt.v4.config.WidgetTemplate"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
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

type OneOfDeleteReportConfigApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType1    *interface{}           `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
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
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "opsmgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteReportConfigApiResponseData"))
}

func (p *OneOfDeleteReportConfigApiResponseData) MarshalJSON() ([]byte, error) {
	if "EMPTY" == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteReportConfigApiResponseData")
}

type OneOfListReportConfigsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
	oneOfType0    []ReportConfig         `json:"-"`
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
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
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
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<opsmgmt.v4.config.ReportConfig>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfListReportConfigsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "opsmgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
		if len(*vOneOfType0) == 0 || "opsmgmt.v4.config.ReportConfig" == *((*vOneOfType0)[0].ObjectType_) {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListReportConfigsApiResponseData"))
}

func (p *OneOfListReportConfigsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<opsmgmt.v4.config.ReportConfig>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfListReportConfigsApiResponseData")
}

type OneOfListReportsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
	oneOfType0    []Report               `json:"-"`
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
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
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
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<opsmgmt.v4.config.Report>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfListReportsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "opsmgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
		if len(*vOneOfType0) == 0 || "opsmgmt.v4.config.Report" == *((*vOneOfType0)[0].ObjectType_) {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListReportsApiResponseData"))
}

func (p *OneOfListReportsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<opsmgmt.v4.config.Report>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfListReportsApiResponseData")
}

type OneOfCreateReportConfigApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
	oneOfType0    *import1.TaskReference `json:"-"`
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
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case import1.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import1.TaskReference)
		}
		*p.oneOfType0 = v.(import1.TaskReference)
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

func (p *OneOfCreateReportConfigApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfCreateReportConfigApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "opsmgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	vOneOfType0 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import1.TaskReference)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateReportConfigApiResponseData"))
}

func (p *OneOfCreateReportConfigApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfCreateReportConfigApiResponseData")
}

type OneOfDeleteReportApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType1    *interface{}           `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
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
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "opsmgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteReportApiResponseData"))
}

func (p *OneOfDeleteReportApiResponseData) MarshalJSON() ([]byte, error) {
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
	oneOfType400  *import2.ErrorResponse `json:"-"`
	oneOfType0    []import2.AppMessage   `json:"-"`
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
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case []import2.AppMessage:
		p.oneOfType0 = v.([]import2.AppMessage)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<opsmgmt.v4.error.AppMessage>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<opsmgmt.v4.error.AppMessage>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfUpdateGlobalReportSettingsApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<opsmgmt.v4.error.AppMessage>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfUpdateGlobalReportSettingsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "opsmgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	vOneOfType0 := new([]import2.AppMessage)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "opsmgmt.v4.error.AppMessage" == *((*vOneOfType0)[0].ObjectType_) {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateGlobalReportSettingsApiResponseData"))
}

func (p *OneOfUpdateGlobalReportSettingsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<opsmgmt.v4.error.AppMessage>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateGlobalReportSettingsApiResponseData")
}

type OneOfGetGlobalReportSettingApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
	oneOfType0    *GlobalReportSetting   `json:"-"`
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
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfGetGlobalReportSettingApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfGetGlobalReportSettingApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "opsmgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	vOneOfType0 := new(GlobalReportSetting)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "opsmgmt.v4.config.GlobalReportSetting" == *vOneOfType0.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetGlobalReportSettingApiResponseData"))
}

func (p *OneOfGetGlobalReportSettingApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfGetGlobalReportSettingApiResponseData")
}

type OneOfUpdateReportConfigApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
	oneOfType0    *import1.TaskReference `json:"-"`
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
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case import1.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import1.TaskReference)
		}
		*p.oneOfType0 = v.(import1.TaskReference)
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

func (p *OneOfUpdateReportConfigApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfUpdateReportConfigApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "opsmgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	vOneOfType0 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import1.TaskReference)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateReportConfigApiResponseData"))
}

func (p *OneOfUpdateReportConfigApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateReportConfigApiResponseData")
}

type OneOfNotifyReportApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
	oneOfType0    *import1.TaskReference `json:"-"`
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
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case import1.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import1.TaskReference)
		}
		*p.oneOfType0 = v.(import1.TaskReference)
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

func (p *OneOfNotifyReportApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfNotifyReportApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "opsmgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	vOneOfType0 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import1.TaskReference)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfNotifyReportApiResponseData"))
}

func (p *OneOfNotifyReportApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfNotifyReportApiResponseData")
}

type OneOfGetReportApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
	oneOfType0    *Report                `json:"-"`
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
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfGetReportApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfGetReportApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "opsmgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	vOneOfType0 := new(Report)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "opsmgmt.v4.config.Report" == *vOneOfType0.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetReportApiResponseData"))
}

func (p *OneOfGetReportApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfGetReportApiResponseData")
}

type OneOfGetReportConfigApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
	oneOfType0    *ReportConfig          `json:"-"`
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
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfGetReportConfigApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfGetReportConfigApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "opsmgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	vOneOfType0 := new(ReportConfig)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "opsmgmt.v4.config.ReportConfig" == *vOneOfType0.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetReportConfigApiResponseData"))
}

func (p *OneOfGetReportConfigApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfGetReportConfigApiResponseData")
}

type OneOfWidgetWidgetInfo struct {
	Discriminator *string         `json:"-"`
	ObjectType_   *string         `json:"-"`
	oneOfType0    *WidgetTemplate `json:"-"`
	oneOfType1    *WidgetConfig   `json:"-"`
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
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(WidgetTemplate)
		}
		*p.oneOfType0 = v.(WidgetTemplate)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case WidgetConfig:
		if nil == p.oneOfType1 {
			p.oneOfType1 = new(WidgetConfig)
		}
		*p.oneOfType1 = v.(WidgetConfig)
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

func (p *OneOfWidgetWidgetInfo) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1
	}
	return nil
}

func (p *OneOfWidgetWidgetInfo) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(WidgetTemplate)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "opsmgmt.v4.config.WidgetTemplate" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(WidgetTemplate)
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
	vOneOfType1 := new(WidgetConfig)
	if err := json.Unmarshal(b, vOneOfType1); err == nil {
		if "opsmgmt.v4.config.WidgetConfig" == *vOneOfType1.ObjectType_ {
			if nil == p.oneOfType1 {
				p.oneOfType1 = new(WidgetConfig)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfWidgetWidgetInfo"))
}

func (p *OneOfWidgetWidgetInfo) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	return nil, errors.New("No value to marshal for OneOfWidgetWidgetInfo")
}

type OneOfCreateReportApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
	oneOfType0    *import1.TaskReference `json:"-"`
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
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case import1.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import1.TaskReference)
		}
		*p.oneOfType0 = v.(import1.TaskReference)
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

func (p *OneOfCreateReportApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfCreateReportApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "opsmgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	vOneOfType0 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import1.TaskReference)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateReportApiResponseData"))
}

func (p *OneOfCreateReportApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfCreateReportApiResponseData")
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
