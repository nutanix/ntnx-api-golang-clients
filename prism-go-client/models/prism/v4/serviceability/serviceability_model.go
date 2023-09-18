/*
 * Generated file models/prism/v4/serviceability/serviceability_model.go.
 *
 * Product version: 4.0.3-alpha-2
 *
 * Part of the Nutanix Prism Versioned APIs
 *
 * (c) 2023 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module prism.v4.serviceability of Nutanix Prism Versioned APIs
*/
package serviceability

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	import4 "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/common/v1/config"
	import2 "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/common/v1/response"
	import1 "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/prism/v4/common"
	import5 "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/prism/v4/config"
	import3 "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/prism/v4/error"
	"time"
)

/*
An alert can be resolved or acknowledged.
*/
type ActionType int

const (
	ACTIONTYPE_UNKNOWN     ActionType = 0
	ACTIONTYPE_REDACTED    ActionType = 1
	ACTIONTYPE_RESOLVE     ActionType = 2
	ACTIONTYPE_ACKNOWLEDGE ActionType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ActionType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"RESOLVE",
		"ACKNOWLEDGE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ActionType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"RESOLVE",
		"ACKNOWLEDGE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ActionType) index(name string) ActionType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"RESOLVE",
		"ACKNOWLEDGE",
	}
	for idx := range names {
		if names[idx] == name {
			return ActionType(idx)
		}
	}
	return ACTIONTYPE_UNKNOWN
}

func (e *ActionType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ActionType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ActionType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ActionType) Ref() *ActionType {
	return &e
}

type Alert struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Name of the user who acknowledged this alert.
	*/
	AcknowledgedByUsername *string `json:"acknowledgedByUsername,omitempty"`
	/*
	  The time in ISO 8601 format when the alert was acknowledged.
	*/
	AcknowledgedTime *time.Time `json:"acknowledgedTime,omitempty"`
	/*
	  List of all entities that are affected by the alert.
	*/
	AffectedEntities []import1.EntityReference `json:"affectedEntities,omitempty"`
	/*
	  A preconfigured or dynamically generated unique value for each alert type. For example, A1128  for storage pool space exceeded alerts.
	*/
	AlertType *string `json:"alertType,omitempty"`
	/*
	  Various categories into which this alert type can be classified into. For example, Hardware, Storage, License and so on.
	*/
	Classifications []string `json:"classifications,omitempty"`
	/*
	  Cluster UUID associated with the source entity of the alert.
	*/
	ClusterUUID *string `json:"clusterUUID,omitempty"`
	/*
	  Time in ISO 8601 format when the alert was created.
	*/
	CreationTime *time.Time `json:"creationTime,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  The impact this alert/event will have on the system. For example, availability, performance, capacity and so on.
	*/
	ImpactTypes []import1.ImpactType `json:"impactTypes,omitempty"`
	/*
	  Field specifies if the alert is acknowledged or not.
	*/
	IsAcknowledged *bool `json:"isAcknowledged,omitempty"`
	/*
	  Field specifies if the alert is auto resolved or not.
	*/
	IsAutoResolved *bool `json:"isAutoResolved,omitempty"`
	/*
	  Field specifies if the alert is resolved or not.
	*/
	IsResolved *bool `json:"isResolved,omitempty"`
	/*
	  Flag to indicate if the alert was generated from a user defined alert policy.
	*/
	IsUserDefined *bool `json:"isUserDefined,omitempty"`
	/*
	  Time in ISO 8601 format when the alert was last updated.
	*/
	LastUpdatedTime *time.Time `json:"lastUpdatedTime,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Additional message associated with the alert.
	*/
	Message *string `json:"message,omitempty"`
	/*
	  Details of the metric for a metric-based event.
	*/
	MetricDetails []import1.MetricDetail `json:"metricDetails,omitempty"`
	/*
	  Cluster UUID associated with the cluster where the alert was first raised.
	*/
	OriginatingClusterUUID *string `json:"originatingClusterUUID,omitempty"`
	/*
	  Additional parameters associated with the alert. These parameters can be used to indicate custom key-value pairs for a given alert instance. For example, Service down in Prism Central alert can have the service name as a parameter.
	*/
	Parameters []import1.Parameter `json:"parameters,omitempty"`
	/*
	  Name of the user who resolved this alert.
	*/
	ResolvedByUsername *string `json:"resolvedByUsername,omitempty"`
	/*
	  The time in ISO 8601 format when the alert was resolved.
	*/
	ResolvedTime *time.Time `json:"resolvedTime,omitempty"`
	/*
	  Possible causes, resolutions and additional details to troubleshoot this alert.
	*/
	RootCauseAnalysis []RootCauseAnalysis `json:"rootCauseAnalysis,omitempty"`
	/*
	  The service that raised the alert. For internal Nutanix services, this value is set to "Nutanix".
	*/
	ServiceName *string `json:"serviceName,omitempty"`

	Severity *import1.Severity `json:"severity,omitempty"`
	/*
	  Contains information on the severity change history for alerts. If an alert was de-duplicated without change in severity, then no trail will be present.
	*/
	SeverityTrails []SeverityTrail `json:"severityTrails,omitempty"`

	SourceEntity *import1.AlertEntityReference `json:"sourceEntity,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  Title of the alert.
	*/
	Title *string `json:"title,omitempty"`
}

func NewAlert() *Alert {
	p := new(Alert)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.serviceability.Alert"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a2.serviceability.Alert"}
	p.UnknownFields_ = map[string]interface{}{}

	p.IsAcknowledged = new(bool)
	*p.IsAcknowledged = false
	p.IsAutoResolved = new(bool)
	*p.IsAutoResolved = false
	p.IsResolved = new(bool)
	*p.IsResolved = false
	p.IsUserDefined = new(bool)
	*p.IsUserDefined = false

	return p
}

/*
An alert can be resolved or acknowledged.
*/
type AlertActionSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	ActionType *ActionType `json:"actionType"`
}

func (p *AlertActionSpec) MarshalJSON() ([]byte, error) {
	type AlertActionSpecProxy AlertActionSpec
	return json.Marshal(struct {
		*AlertActionSpecProxy
		ActionType *ActionType `json:"actionType,omitempty"`
	}{
		AlertActionSpecProxy: (*AlertActionSpecProxy)(p),
		ActionType:           p.ActionType,
	})
}

func NewAlertActionSpec() *AlertActionSpec {
	p := new(AlertActionSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.serviceability.AlertActionSpec"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a2.serviceability.AlertActionSpec"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /prism/v4.0.a2/serviceability/alerts/{extId} Get operation
*/
type AlertApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfAlertApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewAlertApiResponse() *AlertApiResponse {
	p := new(AlertApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.serviceability.AlertApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a2.serviceability.AlertApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *AlertApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *AlertApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfAlertApiResponseData()
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

type AlertConfigExceptionGroup struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	AutoResolve *AutoResolve `json:"autoResolve,omitempty"`
	/*
	  List of cluster UUIDs on which the exceptions can be set.
	*/
	ClusterUuids []string `json:"clusterUuids,omitempty"`
	/*
	  Enable/Disable for each severity information.
	*/
	SeverityThresholdInfos []SeverityThresholdInfo `json:"severityThresholdInfos,omitempty"`
}

func NewAlertConfigExceptionGroup() *AlertConfigExceptionGroup {
	p := new(AlertConfigExceptionGroup)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.serviceability.AlertConfigExceptionGroup"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a2.serviceability.AlertConfigExceptionGroup"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type AlertDb struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	AcknowledgedByUsername *string `json:"acknowledgedByUsername,omitempty"`

	AcknowledgedTime *time.Time `json:"acknowledgedTime,omitempty"`

	AlertType *string `json:"alertType,omitempty"`

	Classifications []string `json:"classifications,omitempty"`

	Cluster []string `json:"cluster,omitempty"`

	CreationTime *time.Time `json:"creationTime,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`

	ImpactTypes []string `json:"impactTypes,omitempty"`

	IsAcknowledged *bool `json:"isAcknowledged,omitempty"`

	IsAutoResolved *bool `json:"isAutoResolved,omitempty"`

	IsResolved *bool `json:"isResolved,omitempty"`

	IsUserDefined *bool `json:"isUserDefined,omitempty"`

	LastUpdatedTime *time.Time `json:"lastUpdatedTime,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`

	OriginatingClusterUUID *string `json:"originatingClusterUUID,omitempty"`

	Protobuf *string `json:"protobuf,omitempty"`

	ResolvedByUsername *string `json:"resolvedByUsername,omitempty"`

	ResolvedTime *time.Time `json:"resolvedTime,omitempty"`

	ServiceName *string `json:"serviceName,omitempty"`

	Severity *import1.Severity `json:"severity,omitempty"`

	SourceEntity *import1.AlertEntityReference `json:"sourceEntity,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewAlertDb() *AlertDb {
	p := new(AlertDb)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.serviceability.AlertDb"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a2.serviceability.AlertDb"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type AlertDbProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	AcknowledgedByUsername *string `json:"acknowledgedByUsername,omitempty"`

	AcknowledgedTime *time.Time `json:"acknowledgedTime,omitempty"`

	AlertType *string `json:"alertType,omitempty"`

	Classifications []string `json:"classifications,omitempty"`

	Cluster []string `json:"cluster,omitempty"`

	CreationTime *time.Time `json:"creationTime,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`

	ImpactTypes []string `json:"impactTypes,omitempty"`

	IsAcknowledged *bool `json:"isAcknowledged,omitempty"`

	IsAutoResolved *bool `json:"isAutoResolved,omitempty"`

	IsResolved *bool `json:"isResolved,omitempty"`

	IsUserDefined *bool `json:"isUserDefined,omitempty"`

	LastUpdatedTime *time.Time `json:"lastUpdatedTime,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`

	OriginatingClusterUUID *string `json:"originatingClusterUUID,omitempty"`

	Protobuf *string `json:"protobuf,omitempty"`

	ResolvedByUsername *string `json:"resolvedByUsername,omitempty"`

	ResolvedTime *time.Time `json:"resolvedTime,omitempty"`

	RootCauseAnalysisProjection *RootCauseAnalysisProjection `json:"rootCauseAnalysisProjection,omitempty"`

	ServiceName *string `json:"serviceName,omitempty"`

	Severity *import1.Severity `json:"severity,omitempty"`

	SourceEntity *import1.AlertEntityReference `json:"sourceEntity,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewAlertDbProjection() *AlertDbProjection {
	p := new(AlertDbProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.serviceability.AlertDbProjection"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a2.serviceability.AlertDbProjection"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type AlertEmailConfiguration struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Time in HH:mm format when the alert email digest is to be sent daily.
	*/
	AlertEmailDigestSendTime *string `json:"alertEmailDigestSendTime,omitempty"`
	/*
	  Timezone for the time at which the alert email digest is to be sent daily.
	*/
	AlertEmailDigestSendTimezone *string `json:"alertEmailDigestSendTimezone,omitempty"`
	/*
	  The default Nutanix email ID to which alert emails would be sent.
	*/
	DefaultNutanixEmail *string `json:"defaultNutanixEmail,omitempty"`
	/*
	  Rules for email configuration.
	*/
	EmailConfigRules []EmailConfigurationRule `json:"emailConfigRules,omitempty"`
	/*
	  List of email contacts.
	*/
	EmailContactList []string `json:"emailContactList,omitempty"`

	EmailTemplate *EmailTemplate `json:"emailTemplate,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Indicates whether alert emails should be enabled or not on default Nutanix email ID.
	*/
	HasDefaultNutanixEmail *bool `json:"hasDefaultNutanixEmail,omitempty"`
	/*
	  Indicates whether alert email digest should be enabled or not.
	*/
	IsEmailDigestEnabled *bool `json:"isEmailDigestEnabled,omitempty"`
	/*
	  Send alert email digest only if there are one or more alerts.
	*/
	IsEmptyAlertEmailDigestSkipped *bool `json:"isEmptyAlertEmailDigestSkipped,omitempty"`
	/*
	  Indicates whether alert emails should be enabled or not.
	*/
	IsEnabled *bool `json:"isEnabled,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	TunnelDetails *RemoteTunnelDetails `json:"tunnelDetails,omitempty"`
}

func NewAlertEmailConfiguration() *AlertEmailConfiguration {
	p := new(AlertEmailConfiguration)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.serviceability.AlertEmailConfiguration"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a2.serviceability.AlertEmailConfiguration"}
	p.UnknownFields_ = map[string]interface{}{}

	p.HasDefaultNutanixEmail = new(bool)
	*p.HasDefaultNutanixEmail = false
	p.IsEmailDigestEnabled = new(bool)
	*p.IsEmailDigestEnabled = false
	p.IsEmptyAlertEmailDigestSkipped = new(bool)
	*p.IsEmptyAlertEmailDigestSkipped = false
	p.IsEnabled = new(bool)
	*p.IsEnabled = false

	return p
}

/*
REST response for all response codes in API path /prism/v4.0.a2/serviceability/alerts/email-config Put operation
*/
type AlertEmailConfigurationApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfAlertEmailConfigurationApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewAlertEmailConfigurationApiResponse() *AlertEmailConfigurationApiResponse {
	p := new(AlertEmailConfigurationApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.serviceability.AlertEmailConfigurationApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a2.serviceability.AlertEmailConfigurationApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *AlertEmailConfigurationApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *AlertEmailConfigurationApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfAlertEmailConfigurationApiResponseData()
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
REST response for all response codes in API path /prism/v4.0.a2/serviceability/alerts Get operation
*/
type AlertListApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfAlertListApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewAlertListApiResponse() *AlertListApiResponse {
	p := new(AlertListApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.serviceability.AlertListApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a2.serviceability.AlertListApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *AlertListApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *AlertListApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfAlertListApiResponseData()
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

type Audit struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of all entities that are affected by the event/audit.
	*/
	AffectedEntities []import1.EntityReference `json:"affectedEntities,omitempty"`
	/*
	  The unique name for a given audit type. For example, VMCloneAudit, VMDeleteAudit and so on.
	*/
	AuditType *string `json:"auditType,omitempty"`

	ClusterReference *import1.EntityReference `json:"clusterReference,omitempty"`
	/*
	  The time in ISO 8601 format when the audit was created.
	*/
	CreationTime *time.Time `json:"creationTime,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Additional message associated with the audit.
	*/
	Message *string `json:"message,omitempty"`
	/*
	  The audit operation end time in ISO 8601 format.
	*/
	OperationEndTime *time.Time `json:"operationEndTime,omitempty"`
	/*
	  The audit operation start time in ISO 8601 format.
	*/
	OperationStartTime *time.Time `json:"operationStartTime,omitempty"`

	OperationType *import1.OperationType `json:"operationType,omitempty"`
	/*
	  Additional parameters associated with the audit. These parameters can be used to indicate custom key-value pairs for a given audit instance. For example, a service down audit in Prism Central can have the service name as a parameter.
	*/
	Parameters []import1.Parameter `json:"parameters,omitempty"`
	/*
	  The service which raised the event/audit. For internal Nutanix services, this value is set to "Nutanix".
	*/
	ServiceName *string `json:"serviceName,omitempty"`

	SourceEntity *import1.EntityReference `json:"sourceEntity,omitempty"`

	Status *Status `json:"status,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	UserReference *UserReference `json:"userReference,omitempty"`
}

func NewAudit() *Audit {
	p := new(Audit)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.serviceability.Audit"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a2.serviceability.Audit"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /prism/v4.0.a2/serviceability/audits/{extId} Get operation
*/
type AuditApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfAuditApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewAuditApiResponse() *AuditApiResponse {
	p := new(AuditApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.serviceability.AuditApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a2.serviceability.AuditApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *AuditApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *AuditApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfAuditApiResponseData()
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

type AuditDb struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	AuditType *string `json:"auditType,omitempty"`

	CreationTime *time.Time `json:"creationTime,omitempty"`

	ExtId *string `json:"extId,omitempty"`

	Message *string `json:"message,omitempty"`

	OperationEndTime *time.Time `json:"operationEndTime,omitempty"`

	OperationStartTime *time.Time `json:"operationStartTime,omitempty"`

	OperationType *import1.OperationType `json:"operationType,omitempty"`

	Protobuf *string `json:"protobuf,omitempty"`

	SourceEntity *AuditEntityReference `json:"sourceEntity,omitempty"`

	Status *Status `json:"status,omitempty"`

	UserReference *UserReference `json:"userReference,omitempty"`
}

func NewAuditDb() *AuditDb {
	p := new(AuditDb)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.serviceability.AuditDb"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a2.serviceability.AuditDb"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type AuditDbProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	AuditType *string `json:"auditType,omitempty"`

	CreationTime *time.Time `json:"creationTime,omitempty"`

	ExtId *string `json:"extId,omitempty"`

	Message *string `json:"message,omitempty"`

	OperationEndTime *time.Time `json:"operationEndTime,omitempty"`

	OperationStartTime *time.Time `json:"operationStartTime,omitempty"`

	OperationType *import1.OperationType `json:"operationType,omitempty"`

	Protobuf *string `json:"protobuf,omitempty"`

	SourceEntity *AuditEntityReference `json:"sourceEntity,omitempty"`

	Status *Status `json:"status,omitempty"`

	UserReference *UserReference `json:"userReference,omitempty"`
}

func NewAuditDbProjection() *AuditDbProjection {
	p := new(AuditDbProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.serviceability.AuditDbProjection"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a2.serviceability.AuditDbProjection"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type AuditEntityReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	ExtId *string `json:"extId,omitempty"`

	Name *string `json:"name,omitempty"`

	Type *string `json:"type,omitempty"`
}

func NewAuditEntityReference() *AuditEntityReference {
	p := new(AuditEntityReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.serviceability.AuditEntityReference"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a2.serviceability.AuditEntityReference"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /prism/v4.0.a2/serviceability/audits Get operation
*/
type AuditListApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfAuditListApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewAuditListApiResponse() *AuditListApiResponse {
	p := new(AuditListApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.serviceability.AuditListApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a2.serviceability.AuditListApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *AuditListApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *AuditListApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfAuditListApiResponseData()
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
Auto resolve status for this system defined alert policy on specified cluster.
*/
type AutoResolve int

const (
	AUTORESOLVE_UNKNOWN       AutoResolve = 0
	AUTORESOLVE_REDACTED      AutoResolve = 1
	AUTORESOLVE_ENABLED       AutoResolve = 2
	AUTORESOLVE_DISABLED      AutoResolve = 3
	AUTORESOLVE_NOT_SUPPORTED AutoResolve = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *AutoResolve) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ENABLED",
		"DISABLED",
		"NOT_SUPPORTED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e AutoResolve) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ENABLED",
		"DISABLED",
		"NOT_SUPPORTED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *AutoResolve) index(name string) AutoResolve {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ENABLED",
		"DISABLED",
		"NOT_SUPPORTED",
	}
	for idx := range names {
		if names[idx] == name {
			return AutoResolve(idx)
		}
	}
	return AUTORESOLVE_UNKNOWN
}

func (e *AutoResolve) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for AutoResolve:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *AutoResolve) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e AutoResolve) Ref() *AutoResolve {
	return &e
}

/*
Status of the remote tunnel or service that is running on top of the remote tunnel. For example, pulse, alert email and so on.
*/
type CommunicationStatus struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Last changed time.
	*/
	LastChangedTime *time.Time `json:"lastChangedTime,omitempty"`
	/*
	  Last checked time.
	*/
	LastCheckedTime *time.Time `json:"lastCheckedTime,omitempty"`
	/*
	  Last successful transmission time.
	*/
	LastSuccessfulTransmissionTime *time.Time `json:"lastSuccessfulTransmissionTime,omitempty"`

	Message *ParameterizedMessage `json:"message,omitempty"`

	Status *ConnectionStatus `json:"status,omitempty"`
}

func NewCommunicationStatus() *CommunicationStatus {
	p := new(CommunicationStatus)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.serviceability.CommunicationStatus"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a2.serviceability.CommunicationStatus"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type Condition struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The metric key. Allowed values of metrics list can be found at https://portal.nutanix.com/page/documents/details?targetId=Prism-Central-Guide-vpc_2022_9:mul-alerts-user-created-metrics-r.html
	*/
	MetricName *string `json:"metricName"`

	Operator *import1.ComparisonOperator `json:"operator"`

	ThresholdValueItemDiscriminator_ *string `json:"$thresholdValueItemDiscriminator,omitempty"`
	/*
	  The threshold value that was used for the condition evaluation.
	*/
	ThresholdValue *OneOfConditionThresholdValue `json:"thresholdValue"`
}

func (p *Condition) MarshalJSON() ([]byte, error) {
	type ConditionProxy Condition
	return json.Marshal(struct {
		*ConditionProxy
		MetricName     *string                       `json:"metricName,omitempty"`
		Operator       *import1.ComparisonOperator   `json:"operator,omitempty"`
		ThresholdValue *OneOfConditionThresholdValue `json:"thresholdValue,omitempty"`
	}{
		ConditionProxy: (*ConditionProxy)(p),
		MetricName:     p.MetricName,
		Operator:       p.Operator,
		ThresholdValue: p.ThresholdValue,
	})
}

func NewCondition() *Condition {
	p := new(Condition)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.serviceability.Condition"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a2.serviceability.Condition"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *Condition) GetThresholdValue() interface{} {
	if nil == p.ThresholdValue {
		return nil
	}
	return p.ThresholdValue.GetValue()
}

func (p *Condition) SetThresholdValue(v interface{}) error {
	if nil == p.ThresholdValue {
		p.ThresholdValue = NewOneOfConditionThresholdValue()
	}
	e := p.ThresholdValue.SetValue(v)
	if nil == e {
		if nil == p.ThresholdValueItemDiscriminator_ {
			p.ThresholdValueItemDiscriminator_ = new(string)
		}
		*p.ThresholdValueItemDiscriminator_ = *p.ThresholdValue.Discriminator
	}
	return e
}

/*
Details of the policy conflicting with a given user-defined alert policy.
*/
type ConflictingPolicy struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Unique UUID associated with the user-defined alert policy which conflicts with the given policy.
	*/
	ExtId *string `json:"extId"`
}

func (p *ConflictingPolicy) MarshalJSON() ([]byte, error) {
	type ConflictingPolicyProxy ConflictingPolicy
	return json.Marshal(struct {
		*ConflictingPolicyProxy
		ExtId *string `json:"extId,omitempty"`
	}{
		ConflictingPolicyProxy: (*ConflictingPolicyProxy)(p),
		ExtId:                  p.ExtId,
	})
}

func NewConflictingPolicy() *ConflictingPolicy {
	p := new(ConflictingPolicy)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.serviceability.ConflictingPolicy"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a2.serviceability.ConflictingPolicy"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type ConnectionStatus int

const (
	CONNECTIONSTATUS_UNKNOWN  ConnectionStatus = 0
	CONNECTIONSTATUS_REDACTED ConnectionStatus = 1
	CONNECTIONSTATUS_SUCCESS  ConnectionStatus = 2
	CONNECTIONSTATUS_FAILED   ConnectionStatus = 3
	CONNECTIONSTATUS_RETRYING ConnectionStatus = 4
	CONNECTIONSTATUS_DISABLED ConnectionStatus = 5
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ConnectionStatus) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SUCCESS",
		"FAILED",
		"RETRYING",
		"DISABLED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ConnectionStatus) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SUCCESS",
		"FAILED",
		"RETRYING",
		"DISABLED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ConnectionStatus) index(name string) ConnectionStatus {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SUCCESS",
		"FAILED",
		"RETRYING",
		"DISABLED",
	}
	for idx := range names {
		if names[idx] == name {
			return ConnectionStatus(idx)
		}
	}
	return CONNECTIONSTATUS_UNKNOWN
}

func (e *ConnectionStatus) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ConnectionStatus:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ConnectionStatus) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ConnectionStatus) Ref() *ConnectionStatus {
	return &e
}

/*
REST response for all response codes in API path /prism/v4.0.a2/serviceability/alerts/user-defined-policies Post operation
*/
type CreateUdaPolicyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateUdaPolicyApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewCreateUdaPolicyApiResponse() *CreateUdaPolicyApiResponse {
	p := new(CreateUdaPolicyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.serviceability.CreateUdaPolicyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a2.serviceability.CreateUdaPolicyApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreateUdaPolicyApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreateUdaPolicyApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreateUdaPolicyApiResponseData()
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
REST response for all response codes in API path /prism/v4.0.a2/serviceability/alerts/user-defined-policies/{extId} Delete operation
*/
type DeleteUdaPolicyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteUdaPolicyApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDeleteUdaPolicyApiResponse() *DeleteUdaPolicyApiResponse {
	p := new(DeleteUdaPolicyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.serviceability.DeleteUdaPolicyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a2.serviceability.DeleteUdaPolicyApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteUdaPolicyApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteUdaPolicyApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteUdaPolicyApiResponseData()
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
Status of remote tunnel that is used to send alert emails.
*/
type EmailConfigurationRule struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Cluster UUIDs to which this rule applies to.
	*/
	ClusterUuids []string `json:"clusterUuids,omitempty"`

	HasGlobalEmailContactList *bool `json:"hasGlobalEmailContactList,omitempty"`

	ImpactTypes []import1.ImpactType `json:"impactTypes,omitempty"`
	/*
	  Indicates if the configuration rule is enabled or not.
	*/
	IsEnabled *bool `json:"isEnabled,omitempty"`
	/*
	  List of phrases to match the alert with.
	*/
	MatchPhrases []string `json:"matchPhrases,omitempty"`
	/*
	  List of recipients who will receive emails.
	*/
	Recipients []string `json:"recipients,omitempty"`

	Severities []import1.Severity `json:"severities,omitempty"`
}

func NewEmailConfigurationRule() *EmailConfigurationRule {
	p := new(EmailConfigurationRule)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.serviceability.EmailConfigurationRule"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a2.serviceability.EmailConfigurationRule"}
	p.UnknownFields_ = map[string]interface{}{}

	p.HasGlobalEmailContactList = new(bool)
	*p.HasGlobalEmailContactList = false
	p.IsEnabled = new(bool)
	*p.IsEnabled = true

	return p
}

/*
Email template.
*/
type EmailTemplate struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Suffix for the email body.
	*/
	BodySuffix *string `json:"bodySuffix,omitempty"`
	/*
	  Prefix for the email subject.
	*/
	SubjectPrefix *string `json:"subjectPrefix,omitempty"`
}

func NewEmailTemplate() *EmailTemplate {
	p := new(EmailTemplate)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.serviceability.EmailTemplate"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a2.serviceability.EmailTemplate"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Entity UUID on which the user-defined alert policy should be set up. This should be UUID related to the entity type.
*/
type EntityFilter struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Entity UUID on which the user-defined alert policy should be set up. This should be UUID related to the entity type.
	*/
	ExtId *string `json:"extId"`
}

func (p *EntityFilter) MarshalJSON() ([]byte, error) {
	type EntityFilterProxy EntityFilter
	return json.Marshal(struct {
		*EntityFilterProxy
		ExtId *string `json:"extId,omitempty"`
	}{
		EntityFilterProxy: (*EntityFilterProxy)(p),
		ExtId:             p.ExtId,
	})
}

func NewEntityFilter() *EntityFilter {
	p := new(EntityFilter)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.serviceability.EntityFilter"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a2.serviceability.EntityFilter"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type Event struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of all entities that are affected by the event/audit.
	*/
	AffectedEntities []import1.EntityReference `json:"affectedEntities,omitempty"`
	/*
	  Various categories into which this event type can be classified. For example, Hardware, Storage, License and so on.
	*/
	Classifications []string `json:"classifications,omitempty"`
	/*
	  Cluster UUID associated with the cluster where the event was first raised.
	*/
	ClusterUUID *string `json:"clusterUUID,omitempty"`
	/*
	  The time in ISO 8601 format when the event was created.
	*/
	CreationTime *time.Time `json:"creationTime,omitempty"`
	/*
	  A preconfigured or dynamically generated unique value for each event type.
	*/
	EventType *string `json:"eventType,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Additional message associated with the event.
	*/
	Message *string `json:"message,omitempty"`
	/*
	  Details of the metric for a metric-based event.
	*/
	MetricDetails []import1.MetricDetail `json:"metricDetails,omitempty"`
	/*
	  Additional parameters associated with the event. These parameters can be used to indicate custom key-value pairs for a given event instance. For example, a service down event in Prism Central can have the service name as a parameter.
	*/
	Parameters []import1.Parameter `json:"parameters,omitempty"`
	/*
	  The service which raised the event/audit. For internal Nutanix services, this value is set to "Nutanix".
	*/
	ServiceName *string `json:"serviceName,omitempty"`
	/*
	  Cluster UUID associated with the source entity of the event.
	*/
	SourceClusterUUID *string `json:"sourceClusterUUID,omitempty"`

	SourceEntity *import1.EntityReference `json:"sourceEntity,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewEvent() *Event {
	p := new(Event)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.serviceability.Event"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a2.serviceability.Event"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /prism/v4.0.a2/serviceability/events/{extId} Get operation
*/
type EventApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfEventApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewEventApiResponse() *EventApiResponse {
	p := new(EventApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.serviceability.EventApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a2.serviceability.EventApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *EventApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *EventApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfEventApiResponseData()
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

type EventDb struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Classifications []string `json:"classifications,omitempty"`

	CreationTime *time.Time `json:"creationTime,omitempty"`

	EventType *string `json:"eventType,omitempty"`

	ExtId *string `json:"extId,omitempty"`

	Message *string `json:"message,omitempty"`

	Protobuf *string `json:"protobuf,omitempty"`

	SourceEntity *import1.AlertEntityReference `json:"sourceEntity,omitempty"`
}

func NewEventDb() *EventDb {
	p := new(EventDb)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.serviceability.EventDb"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a2.serviceability.EventDb"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type EventDbProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Classifications []string `json:"classifications,omitempty"`

	CreationTime *time.Time `json:"creationTime,omitempty"`

	EventType *string `json:"eventType,omitempty"`

	ExtId *string `json:"extId,omitempty"`

	Message *string `json:"message,omitempty"`

	Protobuf *string `json:"protobuf,omitempty"`

	SourceEntity *import1.AlertEntityReference `json:"sourceEntity,omitempty"`
}

func NewEventDbProjection() *EventDbProjection {
	p := new(EventDbProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.serviceability.EventDbProjection"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a2.serviceability.EventDbProjection"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /prism/v4.0.a2/serviceability/events Get operation
*/
type EventListApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfEventListApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewEventListApiResponse() *EventListApiResponse {
	p := new(EventListApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.serviceability.EventListApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a2.serviceability.EventListApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *EventListApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *EventListApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfEventListApiResponseData()
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
REST response for all response codes in API path /prism/v4.0.a2/serviceability/alerts/system-defined-policies/{extId} Get operation
*/
type GetSdaPolicyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetSdaPolicyApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetSdaPolicyApiResponse() *GetSdaPolicyApiResponse {
	p := new(GetSdaPolicyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.serviceability.GetSdaPolicyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a2.serviceability.GetSdaPolicyApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetSdaPolicyApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetSdaPolicyApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetSdaPolicyApiResponseData()
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
REST response for all response codes in API path /prism/v4.0.a2/serviceability/alerts/user-defined-policies/{extId} Get operation
*/
type GetUdaPolicyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetUdaPolicyApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetUdaPolicyApiResponse() *GetUdaPolicyApiResponse {
	p := new(GetUdaPolicyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.serviceability.GetUdaPolicyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a2.serviceability.GetUdaPolicyApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetUdaPolicyApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetUdaPolicyApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetUdaPolicyApiResponseData()
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
The allowed entity types which can be used to filter the set of entities on which user-defined alert policy is setup.
*/
type GroupEntityType int

const (
	GROUPENTITYTYPE_UNKNOWN  GroupEntityType = 0
	GROUPENTITYTYPE_REDACTED GroupEntityType = 1
	GROUPENTITYTYPE_CATEGORY GroupEntityType = 2
	GROUPENTITYTYPE_CLUSTER  GroupEntityType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *GroupEntityType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CATEGORY",
		"CLUSTER",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e GroupEntityType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CATEGORY",
		"CLUSTER",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *GroupEntityType) index(name string) GroupEntityType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CATEGORY",
		"CLUSTER",
	}
	for idx := range names {
		if names[idx] == name {
			return GroupEntityType(idx)
		}
	}
	return GROUPENTITYTYPE_UNKNOWN
}

func (e *GroupEntityType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for GroupEntityType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *GroupEntityType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e GroupEntityType) Ref() *GroupEntityType {
	return &e
}

/*
Group of entities on which the user-defined alert policy should be set up.
*/
type GroupFilter struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Entity UUID of the group entity type on which the user-defined alert policy should be set up.
	*/
	ExtId *string `json:"extId"`

	Type *GroupEntityType `json:"type"`
}

func (p *GroupFilter) MarshalJSON() ([]byte, error) {
	type GroupFilterProxy GroupFilter
	return json.Marshal(struct {
		*GroupFilterProxy
		ExtId *string          `json:"extId,omitempty"`
		Type  *GroupEntityType `json:"type,omitempty"`
	}{
		GroupFilterProxy: (*GroupFilterProxy)(p),
		ExtId:            p.ExtId,
		Type:             p.Type,
	})
}

func NewGroupFilter() *GroupFilter {
	p := new(GroupFilter)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.serviceability.GroupFilter"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a2.serviceability.GroupFilter"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Http proxy used to establish remote tunnel.
*/
type HttpProxy struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	AddressValue *import4.IPAddressOrFQDN `json:"addressValue,omitempty"`
	/*
	  Proxy name.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  Password for proxy authentication.
	*/
	Password *string `json:"password,omitempty"`
	/*
	  Port on which proxy is binding.
	*/
	Port *int `json:"port,omitempty"`
	/*
	  Proxy types to send applicable traffic.
	*/
	ProxyTypes []ProxyType `json:"proxyTypes,omitempty"`
	/*
	  Username for proxy authentication.
	*/
	Username *string `json:"username,omitempty"`
}

func NewHttpProxy() *HttpProxy {
	p := new(HttpProxy)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.serviceability.HttpProxy"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a2.serviceability.HttpProxy"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /prism/v4.0.a2/serviceability/alerts/user-defined-policies/$actions/find-conflicts Post operation
*/
type ListConflictingUdaPoliciesApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListConflictingUdaPoliciesApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListConflictingUdaPoliciesApiResponse() *ListConflictingUdaPoliciesApiResponse {
	p := new(ListConflictingUdaPoliciesApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.serviceability.ListConflictingUdaPoliciesApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a2.serviceability.ListConflictingUdaPoliciesApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListConflictingUdaPoliciesApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListConflictingUdaPoliciesApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListConflictingUdaPoliciesApiResponseData()
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
REST response for all response codes in API path /prism/v4.0.a2/serviceability/alerts/system-defined-policies Get operation
*/
type ListSdaPoliciesApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListSdaPoliciesApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListSdaPoliciesApiResponse() *ListSdaPoliciesApiResponse {
	p := new(ListSdaPoliciesApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.serviceability.ListSdaPoliciesApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a2.serviceability.ListSdaPoliciesApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListSdaPoliciesApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListSdaPoliciesApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListSdaPoliciesApiResponseData()
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
REST response for all response codes in API path /prism/v4.0.a2/serviceability/alerts/user-defined-policies Get operation
*/
type ListUdaPoliciesApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListUdaPoliciesApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListUdaPoliciesApiResponse() *ListUdaPoliciesApiResponse {
	p := new(ListUdaPoliciesApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.serviceability.ListUdaPoliciesApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a2.serviceability.ListUdaPoliciesApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListUdaPoliciesApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListUdaPoliciesApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListUdaPoliciesApiResponseData()
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
REST response for all response codes in API path /prism/v4.0.a2/serviceability/alerts/$actions/manage-alerts-bulk Post operation
*/
type ManageAlertTaskApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfManageAlertTaskApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewManageAlertTaskApiResponse() *ManageAlertTaskApiResponse {
	p := new(ManageAlertTaskApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.serviceability.ManageAlertTaskApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a2.serviceability.ManageAlertTaskApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ManageAlertTaskApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ManageAlertTaskApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfManageAlertTaskApiResponseData()
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

type ManageAlertsBulk struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	ActionType *ActionType `json:"actionType"`
	/*
	  Unique identifiers for alerts that can be resolved or acknowledged.
	*/
	Uuids []string `json:"uuids"`
}

func (p *ManageAlertsBulk) MarshalJSON() ([]byte, error) {
	type ManageAlertsBulkProxy ManageAlertsBulk
	return json.Marshal(struct {
		*ManageAlertsBulkProxy
		ActionType *ActionType `json:"actionType,omitempty"`
		Uuids      []string    `json:"uuids,omitempty"`
	}{
		ManageAlertsBulkProxy: (*ManageAlertsBulkProxy)(p),
		ActionType:            p.ActionType,
		Uuids:                 p.Uuids,
	})
}

func NewManageAlertsBulk() *ManageAlertsBulk {
	p := new(ManageAlertsBulk)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.serviceability.ManageAlertsBulk"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a2.serviceability.ManageAlertsBulk"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
To store any custom message and attributes.
*/
type ParameterizedMessage struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Attributes []import4.KVStringPair `json:"attributes,omitempty"`

	Message *string `json:"message,omitempty"`
}

func NewParameterizedMessage() *ParameterizedMessage {
	p := new(ParameterizedMessage)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.serviceability.ParameterizedMessage"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a2.serviceability.ParameterizedMessage"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type ProxyType int

const (
	PROXYTYPE_UNKNOWN  ProxyType = 0
	PROXYTYPE_REDACTED ProxyType = 1
	PROXYTYPE_HTTP     ProxyType = 2
	PROXYTYPE_HTTPS    ProxyType = 3
	PROXYTYPE_SOCKS    ProxyType = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ProxyType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"HTTP",
		"HTTPS",
		"SOCKS",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ProxyType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"HTTP",
		"HTTPS",
		"SOCKS",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ProxyType) index(name string) ProxyType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"HTTP",
		"HTTPS",
		"SOCKS",
	}
	for idx := range names {
		if names[idx] == name {
			return ProxyType(idx)
		}
	}
	return PROXYTYPE_UNKNOWN
}

func (e *ProxyType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ProxyType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ProxyType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ProxyType) Ref() *ProxyType {
	return &e
}

/*
An alert policy that is related to the entities of the current policy.
*/
type RelatedPolicy struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Uuid of the entity to which the UDA policy is associated.
	*/
	EntityUuid *string `json:"entityUuid"`
	/*
	  Policy IDs that are related to the specified entity.
	*/
	PolicyIds []string `json:"policyIds"`
}

func (p *RelatedPolicy) MarshalJSON() ([]byte, error) {
	type RelatedPolicyProxy RelatedPolicy
	return json.Marshal(struct {
		*RelatedPolicyProxy
		EntityUuid *string  `json:"entityUuid,omitempty"`
		PolicyIds  []string `json:"policyIds,omitempty"`
	}{
		RelatedPolicyProxy: (*RelatedPolicyProxy)(p),
		EntityUuid:         p.EntityUuid,
		PolicyIds:          p.PolicyIds,
	})
}

func NewRelatedPolicy() *RelatedPolicy {
	p := new(RelatedPolicy)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.serviceability.RelatedPolicy"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a2.serviceability.RelatedPolicy"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Remote tunnel details associated with the email configuration.
*/
type RemoteTunnelDetails struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	ConnectionStatus *CommunicationStatus `json:"connectionStatus,omitempty"`

	HttpProxy *HttpProxy `json:"httpProxy,omitempty"`

	ServiceCenter *ServiceCenter `json:"serviceCenter,omitempty"`

	TransportStatus *CommunicationStatus `json:"transportStatus,omitempty"`
}

func NewRemoteTunnelDetails() *RemoteTunnelDetails {
	p := new(RemoteTunnelDetails)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.serviceability.RemoteTunnelDetails"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a2.serviceability.RemoteTunnelDetails"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type RootCauseAnalysis struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Possible causes to this alert.
	*/
	Cause *string `json:"cause,omitempty"`
	/*
	  Additional details to troubleshoot this alert.
	*/
	Detail *string `json:"detail,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Possible resolutions to troubleshoot this alert.
	*/
	Resolution *string `json:"resolution,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewRootCauseAnalysis() *RootCauseAnalysis {
	p := new(RootCauseAnalysis)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.serviceability.RootCauseAnalysis"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a2.serviceability.RootCauseAnalysis"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type RootCauseAnalysisProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Possible causes to this alert.
	*/
	Cause *string `json:"cause,omitempty"`
	/*
	  Additional details to troubleshoot this alert.
	*/
	Detail *string `json:"detail,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Possible resolutions to troubleshoot this alert.
	*/
	Resolution *string `json:"resolution,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewRootCauseAnalysisProjection() *RootCauseAnalysisProjection {
	p := new(RootCauseAnalysisProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.serviceability.RootCauseAnalysisProjection"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a2.serviceability.RootCauseAnalysisProjection"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Service center to which remote tunnel is connected at Nutanix's end.
*/
type ServiceCenter struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  IP address of the service center.
	*/
	IpAddress *string `json:"ipAddress,omitempty"`
	/*
	  Name of service center.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  Port number.
	*/
	Port *int `json:"port,omitempty"`
	/*
	  Username.
	*/
	Username *string `json:"username,omitempty"`
}

func NewServiceCenter() *ServiceCenter {
	p := new(ServiceCenter)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.serviceability.ServiceCenter"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a2.serviceability.ServiceCenter"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Enable/Disable for each severity information.
*/
type SeverityThresholdInfo struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Indicates if the system defined policy is enabled for a given severity.
	*/
	IsEnabled *bool `json:"isEnabled,omitempty"`

	Severity *import1.Severity `json:"severity,omitempty"`
}

func NewSeverityThresholdInfo() *SeverityThresholdInfo {
	p := new(SeverityThresholdInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.serviceability.SeverityThresholdInfo"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a2.serviceability.SeverityThresholdInfo"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Contains information on the severity change history for alerts. If an alert was de-duplicated without change in severity, then no trail will be present.
*/
type SeverityTrail struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Severity *import1.Severity `json:"severity,omitempty"`
	/*
	  The time in ISO 8601 format when the severity of the alert was changed.
	*/
	SeverityChangeTime *time.Time `json:"severityChangeTime,omitempty"`
}

func NewSeverityTrail() *SeverityTrail {
	p := new(SeverityTrail)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.serviceability.SeverityTrail"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a2.serviceability.SeverityTrail"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The status of the operation captured by the audit.
*/
type Status int

const (
	STATUS_UNKNOWN   Status = 0
	STATUS_REDACTED  Status = 1
	STATUS_SUCCEEDED Status = 2
	STATUS_FAILED    Status = 3
	STATUS_ABORTED   Status = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *Status) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SUCCEEDED",
		"FAILED",
		"ABORTED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e Status) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SUCCEEDED",
		"FAILED",
		"ABORTED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *Status) index(name string) Status {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SUCCEEDED",
		"FAILED",
		"ABORTED",
	}
	for idx := range names {
		if names[idx] == name {
			return Status(idx)
		}
	}
	return STATUS_UNKNOWN
}

func (e *Status) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for Status:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *Status) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e Status) Ref() *Status {
	return &e
}

type SystemDefinedPolicy struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Affected entity types.
	*/
	AffectedEntityTypes []string `json:"affectedEntityTypes,omitempty"`
	/*
	  List of clusters that have their alert configurable parameters different from that of in Prism Central.
	*/
	AlertConfigExceptionGroups []AlertConfigExceptionGroup `json:"alertConfigExceptionGroups,omitempty"`
	/*
	  System defined alert policy ID.
	*/
	AlertTypeId *string `json:"alertTypeId,omitempty"`

	AutoResolve *AutoResolve `json:"autoResolve,omitempty"`
	/*
	  Possible causes, resolutions and additional details to troubleshoot this alert.
	*/
	CauseAndResolutions []import1.CauseAndResolution `json:"causeAndResolutions,omitempty"`
	/*
	  Various categories into which this alert type can be classified into. For example, Hardware, Storage, License and so on.
	*/
	Classifications []string `json:"classifications,omitempty"`
	/*
	  Numbers of clusters having exceptions with respect to alert config.
	*/
	ExceptionCount *int64 `json:"exceptionCount,omitempty"`
	/*
	  UID of the system defined alert policy.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Impact types to which this rule applies to.
	*/
	ImpactTypes []import1.ImpactType `json:"impactTypes,omitempty"`
	/*
	  Indicates whether this system defined alert policy is specific to the tenant or not.
	*/
	IsTenantSpecific *bool `json:"isTenantSpecific,omitempty"`
	/*
	  List of KB article numbers.
	*/
	KbArticles []string `json:"kbArticles,omitempty"`
	/*
	  Message for a system defined alert policy.
	*/
	Message *string `json:"message,omitempty"`
	/*
	  Name of the user who last modified the system defined alert policy.
	*/
	ModifiedByUsername *string `json:"modifiedByUsername,omitempty"`
	/*
	  Time in ISO 8601 format for when the system defined alert policy was last modified.
	*/
	ModifiedTime *time.Time `json:"modifiedTime,omitempty"`
	/*
	  Enable/Disable for each severity information.
	*/
	SeverityThresholdInfos []SeverityThresholdInfo `json:"severityThresholdInfos,omitempty"`
	/*
	  Title for the alert policy which can include dynamic parameters like IP address, VM name and so on.
	*/
	SmartTitle *string `json:"smartTitle,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  Title of a system defined alert policy.
	*/
	Title *string `json:"title,omitempty"`
}

func NewSystemDefinedPolicy() *SystemDefinedPolicy {
	p := new(SystemDefinedPolicy)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.serviceability.SystemDefinedPolicy"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a2.serviceability.SystemDefinedPolicy"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type TriggerCondition struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Condition *Condition `json:"condition"`

	ConditionType *import1.ConditionType `json:"conditionType"`

	SeverityLevel *import1.Severity `json:"severityLevel"`
}

func (p *TriggerCondition) MarshalJSON() ([]byte, error) {
	type TriggerConditionProxy TriggerCondition
	return json.Marshal(struct {
		*TriggerConditionProxy
		Condition     *Condition             `json:"condition,omitempty"`
		ConditionType *import1.ConditionType `json:"conditionType,omitempty"`
		SeverityLevel *import1.Severity      `json:"severityLevel,omitempty"`
	}{
		TriggerConditionProxy: (*TriggerConditionProxy)(p),
		Condition:             p.Condition,
		ConditionType:         p.ConditionType,
		SeverityLevel:         p.SeverityLevel,
	})
}

func NewTriggerCondition() *TriggerCondition {
	p := new(TriggerCondition)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.serviceability.TriggerCondition"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a2.serviceability.TriggerCondition"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /prism/v4.0.a2/serviceability/alerts/system-defined-policies/{extId} Put operation
*/
type UpdateSdaPolicyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateSdaPolicyApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewUpdateSdaPolicyApiResponse() *UpdateSdaPolicyApiResponse {
	p := new(UpdateSdaPolicyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.serviceability.UpdateSdaPolicyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a2.serviceability.UpdateSdaPolicyApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdateSdaPolicyApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdateSdaPolicyApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdateSdaPolicyApiResponseData()
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
REST response for all response codes in API path /prism/v4.0.a2/serviceability/alerts/user-defined-policies/{extId} Put operation
*/
type UpdateUdaPolicyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateUdaPolicyApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewUpdateUdaPolicyApiResponse() *UpdateUdaPolicyApiResponse {
	p := new(UpdateUdaPolicyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.serviceability.UpdateUdaPolicyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a2.serviceability.UpdateUdaPolicyApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdateUdaPolicyApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdateUdaPolicyApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdateUdaPolicyApiResponseData()
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

type UserDefinedPolicy struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Username of the user who created the policy.
	*/
	CreatedBy *string `json:"createdBy,omitempty"`
	/*
	  Description of the policy.
	*/
	Description *string `json:"description,omitempty"`
	/*
	  Entity type associated with the user-defined alert policy. Allowed values are vm, node and cluster.
	*/
	EntityType *string `json:"entityType"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`

	FiltersItemDiscriminator_ *string `json:"$filtersItemDiscriminator,omitempty"`
	/*
	  Filter criteria for narrowing down the entities on which user-defined alert policies can be setup.
	*/
	Filters *OneOfUserDefinedPolicyFilters `json:"filters,omitempty"`
	/*
	  Impact Types for the associated resulting alert.
	*/
	ImpactTypes []import1.ImpactType `json:"impactTypes,omitempty"`
	/*
	  Indicates whether the auto resolve feature is enabled for this policy.
	*/
	IsAutoResolved *bool `json:"isAutoResolved,omitempty"`
	/*
	  Enable/Disable flag for the policy.
	*/
	IsEnabled *bool `json:"isEnabled,omitempty"`
	/*
	  Error if conflicting alert policies are found.
	*/
	IsExpectedToErrorOnConflict *bool `json:"isExpectedToErrorOnConflict,omitempty"`
	/*
	  Last updated time of the policy in ISO 8601 format. This value will be used as the CAS value during updates.
	*/
	LastUpdatedTime *time.Time `json:"lastUpdatedTime,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  List of IDs of the alert policies that should be overridden.
	*/
	PoliciesToOverride []string `json:"policiesToOverride,omitempty"`
	/*
	  List of alert policies that are related to the entities of the current policy.
	*/
	RelatedPolicies []RelatedPolicy `json:"relatedPolicies,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  Title of the policy.
	*/
	Title *string `json:"title"`
	/*
	  Trigger conditions for the policy. If there are multiple trigger conditions, all of them will be considered during the operation.
	*/
	TriggerConditions []TriggerCondition `json:"triggerConditions"`
	/*
	  Waiting duration in seconds before triggering the alert, when the specified condition is met.
	*/
	TriggerWaitPeriod *int64 `json:"triggerWaitPeriod,omitempty"`
}

func (p *UserDefinedPolicy) MarshalJSON() ([]byte, error) {
	type UserDefinedPolicyProxy UserDefinedPolicy
	return json.Marshal(struct {
		*UserDefinedPolicyProxy
		EntityType        *string            `json:"entityType,omitempty"`
		Title             *string            `json:"title,omitempty"`
		TriggerConditions []TriggerCondition `json:"triggerConditions,omitempty"`
	}{
		UserDefinedPolicyProxy: (*UserDefinedPolicyProxy)(p),
		EntityType:             p.EntityType,
		Title:                  p.Title,
		TriggerConditions:      p.TriggerConditions,
	})
}

func NewUserDefinedPolicy() *UserDefinedPolicy {
	p := new(UserDefinedPolicy)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.serviceability.UserDefinedPolicy"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a2.serviceability.UserDefinedPolicy"}
	p.UnknownFields_ = map[string]interface{}{}

	p.IsAutoResolved = new(bool)
	*p.IsAutoResolved = true
	p.IsEnabled = new(bool)
	*p.IsEnabled = false
	p.IsExpectedToErrorOnConflict = new(bool)
	*p.IsExpectedToErrorOnConflict = true

	return p
}

/*
Reference to the user who initiated the operation being audited.
*/
type UserReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Unique UUID of the user who initiated the operation.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  The IP address from where the operation was triggered.
	*/
	IpAddress *string `json:"ipAddress,omitempty"`
	/*
	  The name of the user who initiated the operation.
	*/
	Name *string `json:"name,omitempty"`
}

func NewUserReference() *UserReference {
	p := new(UserReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.serviceability.UserReference"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a2.serviceability.UserReference"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type OneOfListUdaPoliciesApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    []UserDefinedPolicy    `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfListUdaPoliciesApiResponseData() *OneOfListUdaPoliciesApiResponseData {
	p := new(OneOfListUdaPoliciesApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListUdaPoliciesApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListUdaPoliciesApiResponseData is nil"))
	}
	switch v.(type) {
	case []UserDefinedPolicy:
		p.oneOfType0 = v.([]UserDefinedPolicy)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<prism.v4.serviceability.UserDefinedPolicy>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<prism.v4.serviceability.UserDefinedPolicy>"
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
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

func (p *OneOfListUdaPoliciesApiResponseData) GetValue() interface{} {
	if "List<prism.v4.serviceability.UserDefinedPolicy>" == *p.Discriminator {
		return p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfListUdaPoliciesApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new([]UserDefinedPolicy)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "prism.v4.serviceability.UserDefinedPolicy" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<prism.v4.serviceability.UserDefinedPolicy>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<prism.v4.serviceability.UserDefinedPolicy>"
			return nil

		}
	}
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "prism.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListUdaPoliciesApiResponseData"))
}

func (p *OneOfListUdaPoliciesApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<prism.v4.serviceability.UserDefinedPolicy>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfListUdaPoliciesApiResponseData")
}

type OneOfEventApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType0    *Event                 `json:"-"`
}

func NewOneOfEventApiResponseData() *OneOfEventApiResponseData {
	p := new(OneOfEventApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfEventApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfEventApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case Event:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(Event)
		}
		*p.oneOfType0 = v.(Event)
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

func (p *OneOfEventApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfEventApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "prism.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	vOneOfType0 := new(Event)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.serviceability.Event" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(Event)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfEventApiResponseData"))
}

func (p *OneOfEventApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfEventApiResponseData")
}

type OneOfAuditListApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType0    []Audit                `json:"-"`
}

func NewOneOfAuditListApiResponseData() *OneOfAuditListApiResponseData {
	p := new(OneOfAuditListApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfAuditListApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfAuditListApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case []Audit:
		p.oneOfType0 = v.([]Audit)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<prism.v4.serviceability.Audit>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<prism.v4.serviceability.Audit>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfAuditListApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<prism.v4.serviceability.Audit>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfAuditListApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "prism.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	vOneOfType0 := new([]Audit)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "prism.v4.serviceability.Audit" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<prism.v4.serviceability.Audit>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<prism.v4.serviceability.Audit>"
			return nil

		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfAuditListApiResponseData"))
}

func (p *OneOfAuditListApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<prism.v4.serviceability.Audit>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfAuditListApiResponseData")
}

type OneOfCreateUdaPolicyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType0    *UserDefinedPolicy     `json:"-"`
}

func NewOneOfCreateUdaPolicyApiResponseData() *OneOfCreateUdaPolicyApiResponseData {
	p := new(OneOfCreateUdaPolicyApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateUdaPolicyApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateUdaPolicyApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case UserDefinedPolicy:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(UserDefinedPolicy)
		}
		*p.oneOfType0 = v.(UserDefinedPolicy)
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

func (p *OneOfCreateUdaPolicyApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfCreateUdaPolicyApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "prism.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	vOneOfType0 := new(UserDefinedPolicy)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.serviceability.UserDefinedPolicy" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(UserDefinedPolicy)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateUdaPolicyApiResponseData"))
}

func (p *OneOfCreateUdaPolicyApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfCreateUdaPolicyApiResponseData")
}

type OneOfConditionThresholdValue struct {
	Discriminator *string              `json:"-"`
	ObjectType_   *string              `json:"-"`
	oneOfType1    *import1.IntValue    `json:"-"`
	oneOfType0    *import1.DoubleValue `json:"-"`
}

func NewOneOfConditionThresholdValue() *OneOfConditionThresholdValue {
	p := new(OneOfConditionThresholdValue)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfConditionThresholdValue) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfConditionThresholdValue is nil"))
	}
	switch v.(type) {
	case import1.IntValue:
		if nil == p.oneOfType1 {
			p.oneOfType1 = new(import1.IntValue)
		}
		*p.oneOfType1 = v.(import1.IntValue)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType1.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType1.ObjectType_
	case import1.DoubleValue:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import1.DoubleValue)
		}
		*p.oneOfType0 = v.(import1.DoubleValue)
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

func (p *OneOfConditionThresholdValue) GetValue() interface{} {
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfConditionThresholdValue) UnmarshalJSON(b []byte) error {
	vOneOfType1 := new(import1.IntValue)
	if err := json.Unmarshal(b, vOneOfType1); err == nil {
		if "prism.v4.common.IntValue" == *vOneOfType1.ObjectType_ {
			if nil == p.oneOfType1 {
				p.oneOfType1 = new(import1.IntValue)
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
	vOneOfType0 := new(import1.DoubleValue)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.common.DoubleValue" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import1.DoubleValue)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfConditionThresholdValue"))
}

func (p *OneOfConditionThresholdValue) MarshalJSON() ([]byte, error) {
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfConditionThresholdValue")
}

type OneOfAlertListApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    []Alert                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfAlertListApiResponseData() *OneOfAlertListApiResponseData {
	p := new(OneOfAlertListApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfAlertListApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfAlertListApiResponseData is nil"))
	}
	switch v.(type) {
	case []Alert:
		p.oneOfType0 = v.([]Alert)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<prism.v4.serviceability.Alert>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<prism.v4.serviceability.Alert>"
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
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

func (p *OneOfAlertListApiResponseData) GetValue() interface{} {
	if "List<prism.v4.serviceability.Alert>" == *p.Discriminator {
		return p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfAlertListApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new([]Alert)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "prism.v4.serviceability.Alert" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<prism.v4.serviceability.Alert>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<prism.v4.serviceability.Alert>"
			return nil

		}
	}
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "prism.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfAlertListApiResponseData"))
}

func (p *OneOfAlertListApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<prism.v4.serviceability.Alert>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfAlertListApiResponseData")
}

type OneOfAuditApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *Audit                 `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfAuditApiResponseData() *OneOfAuditApiResponseData {
	p := new(OneOfAuditApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfAuditApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfAuditApiResponseData is nil"))
	}
	switch v.(type) {
	case Audit:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(Audit)
		}
		*p.oneOfType0 = v.(Audit)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
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

func (p *OneOfAuditApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfAuditApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(Audit)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.serviceability.Audit" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(Audit)
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "prism.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfAuditApiResponseData"))
}

func (p *OneOfAuditApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfAuditApiResponseData")
}

type OneOfDeleteUdaPolicyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType1    *interface{}           `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfDeleteUdaPolicyApiResponseData() *OneOfDeleteUdaPolicyApiResponseData {
	p := new(OneOfDeleteUdaPolicyApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteUdaPolicyApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteUdaPolicyApiResponseData is nil"))
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
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
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

func (p *OneOfDeleteUdaPolicyApiResponseData) GetValue() interface{} {
	if "EMPTY" == *p.Discriminator {
		return *p.oneOfType1
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDeleteUdaPolicyApiResponseData) UnmarshalJSON(b []byte) error {
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "prism.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteUdaPolicyApiResponseData"))
}

func (p *OneOfDeleteUdaPolicyApiResponseData) MarshalJSON() ([]byte, error) {
	if "EMPTY" == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteUdaPolicyApiResponseData")
}

type OneOfListSdaPoliciesApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType0    []SystemDefinedPolicy  `json:"-"`
}

func NewOneOfListSdaPoliciesApiResponseData() *OneOfListSdaPoliciesApiResponseData {
	p := new(OneOfListSdaPoliciesApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListSdaPoliciesApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListSdaPoliciesApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case []SystemDefinedPolicy:
		p.oneOfType0 = v.([]SystemDefinedPolicy)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<prism.v4.serviceability.SystemDefinedPolicy>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<prism.v4.serviceability.SystemDefinedPolicy>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListSdaPoliciesApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<prism.v4.serviceability.SystemDefinedPolicy>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfListSdaPoliciesApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "prism.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	vOneOfType0 := new([]SystemDefinedPolicy)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "prism.v4.serviceability.SystemDefinedPolicy" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<prism.v4.serviceability.SystemDefinedPolicy>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<prism.v4.serviceability.SystemDefinedPolicy>"
			return nil

		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListSdaPoliciesApiResponseData"))
}

func (p *OneOfListSdaPoliciesApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<prism.v4.serviceability.SystemDefinedPolicy>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfListSdaPoliciesApiResponseData")
}

type OneOfUpdateSdaPolicyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType0    *import5.TaskReference `json:"-"`
}

func NewOneOfUpdateSdaPolicyApiResponseData() *OneOfUpdateSdaPolicyApiResponseData {
	p := new(OneOfUpdateSdaPolicyApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateSdaPolicyApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateSdaPolicyApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case import5.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import5.TaskReference)
		}
		*p.oneOfType0 = v.(import5.TaskReference)
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

func (p *OneOfUpdateSdaPolicyApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfUpdateSdaPolicyApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "prism.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	vOneOfType0 := new(import5.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import5.TaskReference)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateSdaPolicyApiResponseData"))
}

func (p *OneOfUpdateSdaPolicyApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateSdaPolicyApiResponseData")
}

type OneOfGetSdaPolicyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *SystemDefinedPolicy   `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfGetSdaPolicyApiResponseData() *OneOfGetSdaPolicyApiResponseData {
	p := new(OneOfGetSdaPolicyApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetSdaPolicyApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetSdaPolicyApiResponseData is nil"))
	}
	switch v.(type) {
	case SystemDefinedPolicy:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(SystemDefinedPolicy)
		}
		*p.oneOfType0 = v.(SystemDefinedPolicy)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
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

func (p *OneOfGetSdaPolicyApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetSdaPolicyApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(SystemDefinedPolicy)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.serviceability.SystemDefinedPolicy" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(SystemDefinedPolicy)
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "prism.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetSdaPolicyApiResponseData"))
}

func (p *OneOfGetSdaPolicyApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetSdaPolicyApiResponseData")
}

type OneOfAlertApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType0    *Alert                 `json:"-"`
}

func NewOneOfAlertApiResponseData() *OneOfAlertApiResponseData {
	p := new(OneOfAlertApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfAlertApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfAlertApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case Alert:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(Alert)
		}
		*p.oneOfType0 = v.(Alert)
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

func (p *OneOfAlertApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfAlertApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "prism.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	vOneOfType0 := new(Alert)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.serviceability.Alert" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(Alert)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfAlertApiResponseData"))
}

func (p *OneOfAlertApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfAlertApiResponseData")
}

type OneOfUserDefinedPolicyFilters struct {
	Discriminator *string        `json:"-"`
	ObjectType_   *string        `json:"-"`
	oneOfType1    []GroupFilter  `json:"-"`
	oneOfType0    []EntityFilter `json:"-"`
}

func NewOneOfUserDefinedPolicyFilters() *OneOfUserDefinedPolicyFilters {
	p := new(OneOfUserDefinedPolicyFilters)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUserDefinedPolicyFilters) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUserDefinedPolicyFilters is nil"))
	}
	switch v.(type) {
	case []GroupFilter:
		p.oneOfType1 = v.([]GroupFilter)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<prism.v4.serviceability.GroupFilter>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<prism.v4.serviceability.GroupFilter>"
	case []EntityFilter:
		p.oneOfType0 = v.([]EntityFilter)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<prism.v4.serviceability.EntityFilter>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<prism.v4.serviceability.EntityFilter>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfUserDefinedPolicyFilters) GetValue() interface{} {
	if "List<prism.v4.serviceability.GroupFilter>" == *p.Discriminator {
		return p.oneOfType1
	}
	if "List<prism.v4.serviceability.EntityFilter>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfUserDefinedPolicyFilters) UnmarshalJSON(b []byte) error {
	vOneOfType1 := new([]GroupFilter)
	if err := json.Unmarshal(b, vOneOfType1); err == nil {
		if len(*vOneOfType1) == 0 || "prism.v4.serviceability.GroupFilter" == *((*vOneOfType1)[0].ObjectType_) {
			p.oneOfType1 = *vOneOfType1
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<prism.v4.serviceability.GroupFilter>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<prism.v4.serviceability.GroupFilter>"
			return nil

		}
	}
	vOneOfType0 := new([]EntityFilter)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "prism.v4.serviceability.EntityFilter" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<prism.v4.serviceability.EntityFilter>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<prism.v4.serviceability.EntityFilter>"
			return nil

		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUserDefinedPolicyFilters"))
}

func (p *OneOfUserDefinedPolicyFilters) MarshalJSON() ([]byte, error) {
	if "List<prism.v4.serviceability.GroupFilter>" == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	if "List<prism.v4.serviceability.EntityFilter>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfUserDefinedPolicyFilters")
}

type OneOfUpdateUdaPolicyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType0    *UserDefinedPolicy     `json:"-"`
}

func NewOneOfUpdateUdaPolicyApiResponseData() *OneOfUpdateUdaPolicyApiResponseData {
	p := new(OneOfUpdateUdaPolicyApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateUdaPolicyApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateUdaPolicyApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case UserDefinedPolicy:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(UserDefinedPolicy)
		}
		*p.oneOfType0 = v.(UserDefinedPolicy)
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

func (p *OneOfUpdateUdaPolicyApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfUpdateUdaPolicyApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "prism.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	vOneOfType0 := new(UserDefinedPolicy)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.serviceability.UserDefinedPolicy" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(UserDefinedPolicy)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateUdaPolicyApiResponseData"))
}

func (p *OneOfUpdateUdaPolicyApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateUdaPolicyApiResponseData")
}

type OneOfListConflictingUdaPoliciesApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    []ConflictingPolicy    `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfListConflictingUdaPoliciesApiResponseData() *OneOfListConflictingUdaPoliciesApiResponseData {
	p := new(OneOfListConflictingUdaPoliciesApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListConflictingUdaPoliciesApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListConflictingUdaPoliciesApiResponseData is nil"))
	}
	switch v.(type) {
	case []ConflictingPolicy:
		p.oneOfType0 = v.([]ConflictingPolicy)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<prism.v4.serviceability.ConflictingPolicy>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<prism.v4.serviceability.ConflictingPolicy>"
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
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

func (p *OneOfListConflictingUdaPoliciesApiResponseData) GetValue() interface{} {
	if "List<prism.v4.serviceability.ConflictingPolicy>" == *p.Discriminator {
		return p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfListConflictingUdaPoliciesApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new([]ConflictingPolicy)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "prism.v4.serviceability.ConflictingPolicy" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<prism.v4.serviceability.ConflictingPolicy>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<prism.v4.serviceability.ConflictingPolicy>"
			return nil

		}
	}
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "prism.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListConflictingUdaPoliciesApiResponseData"))
}

func (p *OneOfListConflictingUdaPoliciesApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<prism.v4.serviceability.ConflictingPolicy>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfListConflictingUdaPoliciesApiResponseData")
}

type OneOfGetUdaPolicyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType0    *UserDefinedPolicy     `json:"-"`
}

func NewOneOfGetUdaPolicyApiResponseData() *OneOfGetUdaPolicyApiResponseData {
	p := new(OneOfGetUdaPolicyApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetUdaPolicyApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetUdaPolicyApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case UserDefinedPolicy:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(UserDefinedPolicy)
		}
		*p.oneOfType0 = v.(UserDefinedPolicy)
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

func (p *OneOfGetUdaPolicyApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfGetUdaPolicyApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "prism.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	vOneOfType0 := new(UserDefinedPolicy)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.serviceability.UserDefinedPolicy" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(UserDefinedPolicy)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetUdaPolicyApiResponseData"))
}

func (p *OneOfGetUdaPolicyApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfGetUdaPolicyApiResponseData")
}

type OneOfAlertEmailConfigurationApiResponseData struct {
	Discriminator *string                  `json:"-"`
	ObjectType_   *string                  `json:"-"`
	oneOfType400  *import3.ErrorResponse   `json:"-"`
	oneOfType0    *AlertEmailConfiguration `json:"-"`
}

func NewOneOfAlertEmailConfigurationApiResponseData() *OneOfAlertEmailConfigurationApiResponseData {
	p := new(OneOfAlertEmailConfigurationApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfAlertEmailConfigurationApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfAlertEmailConfigurationApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case AlertEmailConfiguration:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(AlertEmailConfiguration)
		}
		*p.oneOfType0 = v.(AlertEmailConfiguration)
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

func (p *OneOfAlertEmailConfigurationApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfAlertEmailConfigurationApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "prism.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	vOneOfType0 := new(AlertEmailConfiguration)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.serviceability.AlertEmailConfiguration" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(AlertEmailConfiguration)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfAlertEmailConfigurationApiResponseData"))
}

func (p *OneOfAlertEmailConfigurationApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfAlertEmailConfigurationApiResponseData")
}

type OneOfEventListApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType0    []Event                `json:"-"`
}

func NewOneOfEventListApiResponseData() *OneOfEventListApiResponseData {
	p := new(OneOfEventListApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfEventListApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfEventListApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case []Event:
		p.oneOfType0 = v.([]Event)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<prism.v4.serviceability.Event>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<prism.v4.serviceability.Event>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfEventListApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<prism.v4.serviceability.Event>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfEventListApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "prism.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	vOneOfType0 := new([]Event)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "prism.v4.serviceability.Event" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<prism.v4.serviceability.Event>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<prism.v4.serviceability.Event>"
			return nil

		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfEventListApiResponseData"))
}

func (p *OneOfEventListApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<prism.v4.serviceability.Event>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfEventListApiResponseData")
}

type OneOfManageAlertTaskApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType0    *import5.TaskReference `json:"-"`
}

func NewOneOfManageAlertTaskApiResponseData() *OneOfManageAlertTaskApiResponseData {
	p := new(OneOfManageAlertTaskApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfManageAlertTaskApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfManageAlertTaskApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case import5.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import5.TaskReference)
		}
		*p.oneOfType0 = v.(import5.TaskReference)
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

func (p *OneOfManageAlertTaskApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfManageAlertTaskApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "prism.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	vOneOfType0 := new(import5.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import5.TaskReference)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfManageAlertTaskApiResponseData"))
}

func (p *OneOfManageAlertTaskApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfManageAlertTaskApiResponseData")
}
