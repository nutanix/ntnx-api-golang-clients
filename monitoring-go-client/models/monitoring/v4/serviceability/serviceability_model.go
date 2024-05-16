/*
 * Generated file models/monitoring/v4/serviceability/serviceability_model.go.
 *
 * Product version: 4.0.1-beta-1
 *
 * Part of the Nutanix Monitoring Versioned APIs
 *
 * (c) 2024 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module monitoring.v4.serviceability of Nutanix Monitoring Versioned APIs
*/
package serviceability

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	import4 "github.com/nutanix/ntnx-api-golang-clients/monitoring-go-client/v4/models/common/v1/config"
	import2 "github.com/nutanix/ntnx-api-golang-clients/monitoring-go-client/v4/models/common/v1/response"
	import1 "github.com/nutanix/ntnx-api-golang-clients/monitoring-go-client/v4/models/monitoring/v4/common"
	import3 "github.com/nutanix/ntnx-api-golang-clients/monitoring-go-client/v4/models/monitoring/v4/error"
	import5 "github.com/nutanix/ntnx-api-golang-clients/monitoring-go-client/v4/models/prism/v4/config"
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
	  Various categories into which this alert type can be classified. For example, Hardware, Storage, or License.
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
	  The impact this alert or event will have on the system. For example, availability, performance, or capacity.
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
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
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
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
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
	*p.ObjectType_ = "monitoring.v4.serviceability.Alert"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
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
	*p.ObjectType_ = "monitoring.v4.serviceability.AlertActionSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
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
	*p.ObjectType_ = "monitoring.v4.serviceability.AlertConfigExceptionGroup"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
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

	ImpactTypes []import1.ImpactType `json:"impactTypes,omitempty"`

	IsAcknowledged *bool `json:"isAcknowledged,omitempty"`

	IsAutoResolved *bool `json:"isAutoResolved,omitempty"`

	IsResolved *bool `json:"isResolved,omitempty"`

	IsUserDefined *bool `json:"isUserDefined,omitempty"`

	LastUpdatedTime *time.Time `json:"lastUpdatedTime,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
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
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewAlertDb() *AlertDb {
	p := new(AlertDb)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.AlertDb"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
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

	ImpactTypes []import1.ImpactType `json:"impactTypes,omitempty"`

	IsAcknowledged *bool `json:"isAcknowledged,omitempty"`

	IsAutoResolved *bool `json:"isAutoResolved,omitempty"`

	IsResolved *bool `json:"isResolved,omitempty"`

	IsUserDefined *bool `json:"isUserDefined,omitempty"`

	LastUpdatedTime *time.Time `json:"lastUpdatedTime,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
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
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewAlertDbProjection() *AlertDbProjection {
	p := new(AlertDbProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.AlertDbProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
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
	  Indicates whether alert emails is enabled or not on default Nutanix email ID.
	*/
	HasDefaultNutanixEmail *bool `json:"hasDefaultNutanixEmail,omitempty"`
	/*
	  Indicates whether alert email digest is enabled or not.
	*/
	IsEmailDigestEnabled *bool `json:"isEmailDigestEnabled,omitempty"`
	/*
	  Send alert email digest only if there are one or more alerts.
	*/
	IsEmptyAlertEmailDigestSkipped *bool `json:"isEmptyAlertEmailDigestSkipped,omitempty"`
	/*
	  Indicates whether alert emails is enabled or not.
	*/
	IsEnabled *bool `json:"isEnabled,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	TunnelDetails *RemoteTunnelDetails `json:"tunnelDetails,omitempty"`
}

func NewAlertEmailConfiguration() *AlertEmailConfiguration {
	p := new(AlertEmailConfiguration)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.AlertEmailConfiguration"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
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

type Audit struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of all entities that are affected by the event or audit.
	*/
	AffectedEntities []import1.EntityReference `json:"affectedEntities,omitempty"`
	/*
	  The unique name for a given audit type. For example, VMCloneAudit, or VMDeleteAudit.
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
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
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
	  The service which raised the event or audit. For internal Nutanix services, this value is set to "Nutanix".
	*/
	ServiceName *string `json:"serviceName,omitempty"`

	SourceEntity *AuditEntityReference `json:"sourceEntity,omitempty"`

	Status *Status `json:"status,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	UserReference *UserReference `json:"userReference,omitempty"`
}

func NewAudit() *Audit {
	p := new(Audit)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.Audit"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
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
	*p.ObjectType_ = "monitoring.v4.serviceability.AuditDb"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
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
	*p.ObjectType_ = "monitoring.v4.serviceability.AuditDbProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type AuditEntityReference struct {
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

func NewAuditEntityReference() *AuditEntityReference {
	p := new(AuditEntityReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.AuditEntityReference"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Auto resolve status for this System-Defined Alert policy on specified cluster.
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
Status of the remote tunnel or service that is running on top of the remote tunnel. For example, pulse, or alert email.
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
	*p.ObjectType_ = "monitoring.v4.serviceability.CommunicationStatus"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
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
	*p.ObjectType_ = "monitoring.v4.serviceability.Condition"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
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
Details of the policy conflicting with a given User-Defined Alert policy.
*/
type ConflictingPolicy struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Unique UUID associated with the User-Defined Alert policy which conflicts with the given policy.
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
	*p.ObjectType_ = "monitoring.v4.serviceability.ConflictingPolicy"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
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
REST response for all response codes in API path /monitoring/v4.0.b1/serviceability/alerts/user-defined-policies Post operation
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
	*p.ObjectType_ = "monitoring.v4.serviceability.CreateUdaPolicyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
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
REST response for all response codes in API path /monitoring/v4.0.b1/serviceability/alerts/user-defined-policies/{extId} Delete operation
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
	*p.ObjectType_ = "monitoring.v4.serviceability.DeleteUdaPolicyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
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
	  Cluster UUIDs to which this rule applies.
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
	*p.ObjectType_ = "monitoring.v4.serviceability.EmailConfigurationRule"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
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
	*p.ObjectType_ = "monitoring.v4.serviceability.EmailTemplate"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Entity UUID on which the User-Defined Alert policy should be set up. This should be UUID related to the entity type.
*/
type EntityFilter struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Entity UUID on which the User-Defined Alert policy should be set up. This should be UUID related to the entity type.
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
	*p.ObjectType_ = "monitoring.v4.serviceability.EntityFilter"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type Event struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of all entities that are affected by the event or audit.
	*/
	AffectedEntities []import1.EntityReference `json:"affectedEntities,omitempty"`
	/*
	  Various categories into which this event type can be classified. For example, Hardware, Storage, or License.
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
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
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
	  The service which raised the event or audit. For internal Nutanix services, this value is set to "Nutanix".
	*/
	ServiceName *string `json:"serviceName,omitempty"`
	/*
	  Cluster UUID associated with the source entity of the event.
	*/
	SourceClusterUUID *string `json:"sourceClusterUUID,omitempty"`

	SourceEntity *EventEntityReference `json:"sourceEntity,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewEvent() *Event {
	p := new(Event)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.Event"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
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

	SourceEntity *EventEntityReference `json:"sourceEntity,omitempty"`
}

func NewEventDb() *EventDb {
	p := new(EventDb)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.EventDb"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
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

	SourceEntity *EventEntityReference `json:"sourceEntity,omitempty"`
}

func NewEventDbProjection() *EventDbProjection {
	p := new(EventDbProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.EventDbProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type EventEntityReference struct {
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

func NewEventEntityReference() *EventEntityReference {
	p := new(EventEntityReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.EventEntityReference"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /monitoring/v4.0.b1/serviceability/alerts/user-defined-policies/$actions/find-conflicts Post operation
*/
type FindConflictingUdaPoliciesApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfFindConflictingUdaPoliciesApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewFindConflictingUdaPoliciesApiResponse() *FindConflictingUdaPoliciesApiResponse {
	p := new(FindConflictingUdaPoliciesApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.FindConflictingUdaPoliciesApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *FindConflictingUdaPoliciesApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *FindConflictingUdaPoliciesApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfFindConflictingUdaPoliciesApiResponseData()
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
REST response for all response codes in API path /monitoring/v4.0.b1/serviceability/alerts/{extId} Get operation
*/
type GetAlertApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetAlertApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetAlertApiResponse() *GetAlertApiResponse {
	p := new(GetAlertApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.GetAlertApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetAlertApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetAlertApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetAlertApiResponseData()
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
REST response for all response codes in API path /monitoring/v4.0.b1/serviceability/alerts/email-config Get operation
*/
type GetAlertEmailConfigurationApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetAlertEmailConfigurationApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetAlertEmailConfigurationApiResponse() *GetAlertEmailConfigurationApiResponse {
	p := new(GetAlertEmailConfigurationApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.GetAlertEmailConfigurationApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetAlertEmailConfigurationApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetAlertEmailConfigurationApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetAlertEmailConfigurationApiResponseData()
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
REST response for all response codes in API path /monitoring/v4.0.b1/serviceability/audits/{extId} Get operation
*/
type GetAuditApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetAuditApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetAuditApiResponse() *GetAuditApiResponse {
	p := new(GetAuditApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.GetAuditApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetAuditApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetAuditApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetAuditApiResponseData()
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
REST response for all response codes in API path /monitoring/v4.0.b1/serviceability/events/{extId} Get operation
*/
type GetEventApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetEventApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetEventApiResponse() *GetEventApiResponse {
	p := new(GetEventApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.GetEventApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetEventApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetEventApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetEventApiResponseData()
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
REST response for all response codes in API path /monitoring/v4.0.b1/serviceability/alerts/user-defined-policies/{extId} Get operation
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
	*p.ObjectType_ = "monitoring.v4.serviceability.GetUdaPolicyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
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
The allowed entity types which can be used to filter the set of entities on which User-Defined Alert policy is setup.
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
Group of entities on which the User-Defined Alert policy should be set up.
*/
type GroupFilter struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Entity UUID of the group entity type on which the User-Defined Alert policy should be set up.
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
	*p.ObjectType_ = "monitoring.v4.serviceability.GroupFilter"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
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
	*p.ObjectType_ = "monitoring.v4.serviceability.HttpProxy"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /monitoring/v4.0.b1/serviceability/alerts Get operation
*/
type ListAlertsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListAlertsApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListAlertsApiResponse() *ListAlertsApiResponse {
	p := new(ListAlertsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.ListAlertsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListAlertsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListAlertsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListAlertsApiResponseData()
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
REST response for all response codes in API path /monitoring/v4.0.b1/serviceability/audits Get operation
*/
type ListAuditsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListAuditsApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListAuditsApiResponse() *ListAuditsApiResponse {
	p := new(ListAuditsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.ListAuditsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListAuditsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListAuditsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListAuditsApiResponseData()
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
REST response for all response codes in API path /monitoring/v4.0.b1/serviceability/events Get operation
*/
type ListEventsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListEventsApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListEventsApiResponse() *ListEventsApiResponse {
	p := new(ListEventsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.ListEventsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListEventsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListEventsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListEventsApiResponseData()
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
REST response for all response codes in API path /monitoring/v4.0.b1/serviceability/alerts/user-defined-policies Get operation
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
	*p.ObjectType_ = "monitoring.v4.serviceability.ListUdaPoliciesApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
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
REST response for all response codes in API path /monitoring/v4.0.b1/serviceability/alerts/{extId}/$actions/manage-alert Post operation
*/
type ManageAlertApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfManageAlertApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewManageAlertApiResponse() *ManageAlertApiResponse {
	p := new(ManageAlertApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.ManageAlertApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ManageAlertApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ManageAlertApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfManageAlertApiResponseData()
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
REST response for all response codes in API path /monitoring/v4.0.b1/serviceability/alerts/$actions/manage-alerts-bulk Post operation
*/
type ManageAlertsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfManageAlertsApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewManageAlertsApiResponse() *ManageAlertsApiResponse {
	p := new(ManageAlertsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.ManageAlertsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ManageAlertsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ManageAlertsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfManageAlertsApiResponseData()
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
	*p.ObjectType_ = "monitoring.v4.serviceability.ParameterizedMessage"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Severity of an alert.
*/
type PolicySeverityLevel int

const (
	POLICYSEVERITYLEVEL_UNKNOWN  PolicySeverityLevel = 0
	POLICYSEVERITYLEVEL_REDACTED PolicySeverityLevel = 1
	POLICYSEVERITYLEVEL_WARNING  PolicySeverityLevel = 2
	POLICYSEVERITYLEVEL_CRITICAL PolicySeverityLevel = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *PolicySeverityLevel) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"WARNING",
		"CRITICAL",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e PolicySeverityLevel) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"WARNING",
		"CRITICAL",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *PolicySeverityLevel) index(name string) PolicySeverityLevel {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"WARNING",
		"CRITICAL",
	}
	for idx := range names {
		if names[idx] == name {
			return PolicySeverityLevel(idx)
		}
	}
	return POLICYSEVERITYLEVEL_UNKNOWN
}

func (e *PolicySeverityLevel) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for PolicySeverityLevel:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *PolicySeverityLevel) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e PolicySeverityLevel) Ref() *PolicySeverityLevel {
	return &e
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
	  UUID of the entity to which the UDA policy is associated.
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
	*p.ObjectType_ = "monitoring.v4.serviceability.RelatedPolicy"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
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
	*p.ObjectType_ = "monitoring.v4.serviceability.RemoteTunnelDetails"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type RootCauseAnalysis struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Possible causes of this alert.
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
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Possible resolutions to troubleshoot this alert.
	*/
	Resolution *string `json:"resolution,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewRootCauseAnalysis() *RootCauseAnalysis {
	p := new(RootCauseAnalysis)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.RootCauseAnalysis"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type RootCauseAnalysisProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Possible causes of this alert.
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
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Possible resolutions to troubleshoot this alert.
	*/
	Resolution *string `json:"resolution,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewRootCauseAnalysisProjection() *RootCauseAnalysisProjection {
	p := new(RootCauseAnalysisProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.RootCauseAnalysisProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
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
	*p.ObjectType_ = "monitoring.v4.serviceability.ServiceCenter"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
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
	*p.ObjectType_ = "monitoring.v4.serviceability.SeverityThresholdInfo"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
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
	*p.ObjectType_ = "monitoring.v4.serviceability.SeverityTrail"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
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

type TriggerCondition struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Condition *Condition `json:"condition"`

	ConditionType *import1.ConditionType `json:"conditionType"`

	SeverityLevel *PolicySeverityLevel `json:"severityLevel"`
}

func (p *TriggerCondition) MarshalJSON() ([]byte, error) {
	type TriggerConditionProxy TriggerCondition
	return json.Marshal(struct {
		*TriggerConditionProxy
		Condition     *Condition             `json:"condition,omitempty"`
		ConditionType *import1.ConditionType `json:"conditionType,omitempty"`
		SeverityLevel *PolicySeverityLevel   `json:"severityLevel,omitempty"`
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
	*p.ObjectType_ = "monitoring.v4.serviceability.TriggerCondition"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /monitoring/v4.0.b1/serviceability/alerts/email-config Put operation
*/
type UpdateAlertEmailConfigurationApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateAlertEmailConfigurationApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewUpdateAlertEmailConfigurationApiResponse() *UpdateAlertEmailConfigurationApiResponse {
	p := new(UpdateAlertEmailConfigurationApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.UpdateAlertEmailConfigurationApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdateAlertEmailConfigurationApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdateAlertEmailConfigurationApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdateAlertEmailConfigurationApiResponseData()
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
REST response for all response codes in API path /monitoring/v4.0.b1/serviceability/alerts/system-defined-policies/{extId} Put operation
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
	*p.ObjectType_ = "monitoring.v4.serviceability.UpdateSdaPolicyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
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
REST response for all response codes in API path /monitoring/v4.0.b1/serviceability/alerts/user-defined-policies/{extId} Put operation
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
	*p.ObjectType_ = "monitoring.v4.serviceability.UpdateUdaPolicyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
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
	  Entity type associated with the User-Defined Alert policy. Allowed values are VM, node and cluster.
	*/
	EntityType *string `json:"entityType"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`

	FiltersItemDiscriminator_ *string `json:"$filtersItemDiscriminator,omitempty"`
	/*
	  Filter criteria for narrowing down the entities on which User-Defined Alert policies can be setup.
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
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
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
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
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
	  Waiting duration in seconds before triggering the alert, when the specified condition is met. It is set to 600s by default.
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
	*p.ObjectType_ = "monitoring.v4.serviceability.UserDefinedPolicy"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	p.IsAutoResolved = new(bool)
	*p.IsAutoResolved = true
	p.IsEnabled = new(bool)
	*p.IsEnabled = false
	p.IsExpectedToErrorOnConflict = new(bool)
	*p.IsExpectedToErrorOnConflict = true

	return p
}

func (p *UserDefinedPolicy) GetFilters() interface{} {
	if nil == p.Filters {
		return nil
	}
	return p.Filters.GetValue()
}

func (p *UserDefinedPolicy) SetFilters(v interface{}) error {
	if nil == p.Filters {
		p.Filters = NewOneOfUserDefinedPolicyFilters()
	}
	e := p.Filters.SetValue(v)
	if nil == e {
		if nil == p.FiltersItemDiscriminator_ {
			p.FiltersItemDiscriminator_ = new(string)
		}
		*p.FiltersItemDiscriminator_ = *p.Filters.Discriminator
	}
	return e
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
	*p.ObjectType_ = "monitoring.v4.serviceability.UserReference"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type OneOfGetEventApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *Event                 `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfGetEventApiResponseData() *OneOfGetEventApiResponseData {
	p := new(OneOfGetEventApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetEventApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetEventApiResponseData is nil"))
	}
	switch v.(type) {
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

func (p *OneOfGetEventApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetEventApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(Event)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "monitoring.v4.serviceability.Event" == *vOneOfType0.ObjectType_ {
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "monitoring.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetEventApiResponseData"))
}

func (p *OneOfGetEventApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetEventApiResponseData")
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
		if "monitoring.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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

type OneOfUpdateAlertEmailConfigurationApiResponseData struct {
	Discriminator *string                  `json:"-"`
	ObjectType_   *string                  `json:"-"`
	oneOfType0    *AlertEmailConfiguration `json:"-"`
	oneOfType400  *import3.ErrorResponse   `json:"-"`
}

func NewOneOfUpdateAlertEmailConfigurationApiResponseData() *OneOfUpdateAlertEmailConfigurationApiResponseData {
	p := new(OneOfUpdateAlertEmailConfigurationApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateAlertEmailConfigurationApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateAlertEmailConfigurationApiResponseData is nil"))
	}
	switch v.(type) {
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

func (p *OneOfUpdateAlertEmailConfigurationApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfUpdateAlertEmailConfigurationApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(AlertEmailConfiguration)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "monitoring.v4.serviceability.AlertEmailConfiguration" == *vOneOfType0.ObjectType_ {
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "monitoring.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateAlertEmailConfigurationApiResponseData"))
}

func (p *OneOfUpdateAlertEmailConfigurationApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateAlertEmailConfigurationApiResponseData")
}

type OneOfUserDefinedPolicyFilters struct {
	Discriminator *string        `json:"-"`
	ObjectType_   *string        `json:"-"`
	oneOfType0    []EntityFilter `json:"-"`
	oneOfType1    []GroupFilter  `json:"-"`
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
	case []EntityFilter:
		p.oneOfType0 = v.([]EntityFilter)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<monitoring.v4.serviceability.EntityFilter>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<monitoring.v4.serviceability.EntityFilter>"
	case []GroupFilter:
		p.oneOfType1 = v.([]GroupFilter)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<monitoring.v4.serviceability.GroupFilter>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<monitoring.v4.serviceability.GroupFilter>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfUserDefinedPolicyFilters) GetValue() interface{} {
	if "List<monitoring.v4.serviceability.EntityFilter>" == *p.Discriminator {
		return p.oneOfType0
	}
	if "List<monitoring.v4.serviceability.GroupFilter>" == *p.Discriminator {
		return p.oneOfType1
	}
	return nil
}

func (p *OneOfUserDefinedPolicyFilters) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new([]EntityFilter)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {

		if len(*vOneOfType0) == 0 || "monitoring.v4.serviceability.EntityFilter" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<monitoring.v4.serviceability.EntityFilter>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<monitoring.v4.serviceability.EntityFilter>"
			return nil

		}
	}
	vOneOfType1 := new([]GroupFilter)
	if err := json.Unmarshal(b, vOneOfType1); err == nil {

		if len(*vOneOfType1) == 0 || "monitoring.v4.serviceability.GroupFilter" == *((*vOneOfType1)[0].ObjectType_) {
			p.oneOfType1 = *vOneOfType1
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<monitoring.v4.serviceability.GroupFilter>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<monitoring.v4.serviceability.GroupFilter>"
			return nil

		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUserDefinedPolicyFilters"))
}

func (p *OneOfUserDefinedPolicyFilters) MarshalJSON() ([]byte, error) {
	if "List<monitoring.v4.serviceability.EntityFilter>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if "List<monitoring.v4.serviceability.GroupFilter>" == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	return nil, errors.New("No value to marshal for OneOfUserDefinedPolicyFilters")
}

type OneOfListEventsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    []Event                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfListEventsApiResponseData() *OneOfListEventsApiResponseData {
	p := new(OneOfListEventsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListEventsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListEventsApiResponseData is nil"))
	}
	switch v.(type) {
	case []Event:
		p.oneOfType0 = v.([]Event)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<monitoring.v4.serviceability.Event>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<monitoring.v4.serviceability.Event>"
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

func (p *OneOfListEventsApiResponseData) GetValue() interface{} {
	if "List<monitoring.v4.serviceability.Event>" == *p.Discriminator {
		return p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfListEventsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new([]Event)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {

		if len(*vOneOfType0) == 0 || "monitoring.v4.serviceability.Event" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<monitoring.v4.serviceability.Event>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<monitoring.v4.serviceability.Event>"
			return nil

		}
	}
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "monitoring.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListEventsApiResponseData"))
}

func (p *OneOfListEventsApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<monitoring.v4.serviceability.Event>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfListEventsApiResponseData")
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
		if "monitoring.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
		if "monitoring.v4.serviceability.UserDefinedPolicy" == *vOneOfType0.ObjectType_ {
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
		if "monitoring.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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

type OneOfGetAlertApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *Alert                 `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfGetAlertApiResponseData() *OneOfGetAlertApiResponseData {
	p := new(OneOfGetAlertApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetAlertApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetAlertApiResponseData is nil"))
	}
	switch v.(type) {
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

func (p *OneOfGetAlertApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetAlertApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(Alert)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "monitoring.v4.serviceability.Alert" == *vOneOfType0.ObjectType_ {
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "monitoring.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetAlertApiResponseData"))
}

func (p *OneOfGetAlertApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetAlertApiResponseData")
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
		if "monitoring.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
		if "monitoring.v4.serviceability.UserDefinedPolicy" == *vOneOfType0.ObjectType_ {
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

type OneOfListAlertsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType0    []Alert                `json:"-"`
}

func NewOneOfListAlertsApiResponseData() *OneOfListAlertsApiResponseData {
	p := new(OneOfListAlertsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListAlertsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListAlertsApiResponseData is nil"))
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
	case []Alert:
		p.oneOfType0 = v.([]Alert)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<monitoring.v4.serviceability.Alert>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<monitoring.v4.serviceability.Alert>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListAlertsApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<monitoring.v4.serviceability.Alert>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfListAlertsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "monitoring.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new([]Alert)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {

		if len(*vOneOfType0) == 0 || "monitoring.v4.serviceability.Alert" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<monitoring.v4.serviceability.Alert>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<monitoring.v4.serviceability.Alert>"
			return nil

		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListAlertsApiResponseData"))
}

func (p *OneOfListAlertsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<monitoring.v4.serviceability.Alert>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfListAlertsApiResponseData")
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
		if "monitoring.v4.common.IntValue" == *vOneOfType1.ObjectType_ {
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
		if "monitoring.v4.common.DoubleValue" == *vOneOfType0.ObjectType_ {
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

type OneOfGetAuditApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *Audit                 `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfGetAuditApiResponseData() *OneOfGetAuditApiResponseData {
	p := new(OneOfGetAuditApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetAuditApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetAuditApiResponseData is nil"))
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

func (p *OneOfGetAuditApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetAuditApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(Audit)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "monitoring.v4.serviceability.Audit" == *vOneOfType0.ObjectType_ {
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
		if "monitoring.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetAuditApiResponseData"))
}

func (p *OneOfGetAuditApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetAuditApiResponseData")
}

type OneOfGetAlertEmailConfigurationApiResponseData struct {
	Discriminator *string                  `json:"-"`
	ObjectType_   *string                  `json:"-"`
	oneOfType0    *AlertEmailConfiguration `json:"-"`
	oneOfType400  *import3.ErrorResponse   `json:"-"`
}

func NewOneOfGetAlertEmailConfigurationApiResponseData() *OneOfGetAlertEmailConfigurationApiResponseData {
	p := new(OneOfGetAlertEmailConfigurationApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetAlertEmailConfigurationApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetAlertEmailConfigurationApiResponseData is nil"))
	}
	switch v.(type) {
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

func (p *OneOfGetAlertEmailConfigurationApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetAlertEmailConfigurationApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(AlertEmailConfiguration)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "monitoring.v4.serviceability.AlertEmailConfiguration" == *vOneOfType0.ObjectType_ {
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "monitoring.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetAlertEmailConfigurationApiResponseData"))
}

func (p *OneOfGetAlertEmailConfigurationApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetAlertEmailConfigurationApiResponseData")
}

type OneOfFindConflictingUdaPoliciesApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType0    []ConflictingPolicy    `json:"-"`
}

func NewOneOfFindConflictingUdaPoliciesApiResponseData() *OneOfFindConflictingUdaPoliciesApiResponseData {
	p := new(OneOfFindConflictingUdaPoliciesApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfFindConflictingUdaPoliciesApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfFindConflictingUdaPoliciesApiResponseData is nil"))
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
	case []ConflictingPolicy:
		p.oneOfType0 = v.([]ConflictingPolicy)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<monitoring.v4.serviceability.ConflictingPolicy>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<monitoring.v4.serviceability.ConflictingPolicy>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfFindConflictingUdaPoliciesApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<monitoring.v4.serviceability.ConflictingPolicy>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfFindConflictingUdaPoliciesApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "monitoring.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new([]ConflictingPolicy)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {

		if len(*vOneOfType0) == 0 || "monitoring.v4.serviceability.ConflictingPolicy" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<monitoring.v4.serviceability.ConflictingPolicy>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<monitoring.v4.serviceability.ConflictingPolicy>"
			return nil

		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfFindConflictingUdaPoliciesApiResponseData"))
}

func (p *OneOfFindConflictingUdaPoliciesApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<monitoring.v4.serviceability.ConflictingPolicy>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfFindConflictingUdaPoliciesApiResponseData")
}

type OneOfListUdaPoliciesApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType0    []UserDefinedPolicy    `json:"-"`
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
	case []UserDefinedPolicy:
		p.oneOfType0 = v.([]UserDefinedPolicy)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<monitoring.v4.serviceability.UserDefinedPolicy>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<monitoring.v4.serviceability.UserDefinedPolicy>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListUdaPoliciesApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<monitoring.v4.serviceability.UserDefinedPolicy>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfListUdaPoliciesApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "monitoring.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new([]UserDefinedPolicy)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {

		if len(*vOneOfType0) == 0 || "monitoring.v4.serviceability.UserDefinedPolicy" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<monitoring.v4.serviceability.UserDefinedPolicy>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<monitoring.v4.serviceability.UserDefinedPolicy>"
			return nil

		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListUdaPoliciesApiResponseData"))
}

func (p *OneOfListUdaPoliciesApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<monitoring.v4.serviceability.UserDefinedPolicy>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfListUdaPoliciesApiResponseData")
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
		if "monitoring.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
		if "monitoring.v4.serviceability.UserDefinedPolicy" == *vOneOfType0.ObjectType_ {
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

type OneOfManageAlertsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType0    *import5.TaskReference `json:"-"`
}

func NewOneOfManageAlertsApiResponseData() *OneOfManageAlertsApiResponseData {
	p := new(OneOfManageAlertsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfManageAlertsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfManageAlertsApiResponseData is nil"))
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

func (p *OneOfManageAlertsApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfManageAlertsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "monitoring.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfManageAlertsApiResponseData"))
}

func (p *OneOfManageAlertsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfManageAlertsApiResponseData")
}

type OneOfManageAlertApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType0    *import5.TaskReference `json:"-"`
}

func NewOneOfManageAlertApiResponseData() *OneOfManageAlertApiResponseData {
	p := new(OneOfManageAlertApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfManageAlertApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfManageAlertApiResponseData is nil"))
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

func (p *OneOfManageAlertApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfManageAlertApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "monitoring.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfManageAlertApiResponseData"))
}

func (p *OneOfManageAlertApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfManageAlertApiResponseData")
}

type OneOfListAuditsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType0    []Audit                `json:"-"`
}

func NewOneOfListAuditsApiResponseData() *OneOfListAuditsApiResponseData {
	p := new(OneOfListAuditsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListAuditsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListAuditsApiResponseData is nil"))
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
		*p.Discriminator = "List<monitoring.v4.serviceability.Audit>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<monitoring.v4.serviceability.Audit>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListAuditsApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<monitoring.v4.serviceability.Audit>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfListAuditsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "monitoring.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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

		if len(*vOneOfType0) == 0 || "monitoring.v4.serviceability.Audit" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<monitoring.v4.serviceability.Audit>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<monitoring.v4.serviceability.Audit>"
			return nil

		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListAuditsApiResponseData"))
}

func (p *OneOfListAuditsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<monitoring.v4.serviceability.Audit>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfListAuditsApiResponseData")
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
