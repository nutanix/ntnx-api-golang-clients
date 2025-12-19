/*
 * Generated file models/monitoring/v4/serviceability/serviceability_model.go.
 *
 * Product version: 4.2.1
 *
 * Part of the Nutanix Monitoring APIs
 *
 * (c) 2025 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module monitoring.v4.serviceability of Nutanix Monitoring APIs
*/
package serviceability

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	import5 "github.com/nutanix/ntnx-api-golang-clients/monitoring-go-client/v4/models/common/v1/config"
	import2 "github.com/nutanix/ntnx-api-golang-clients/monitoring-go-client/v4/models/common/v1/response"
	import1 "github.com/nutanix/ntnx-api-golang-clients/monitoring-go-client/v4/models/monitoring/v4/common"
	import4 "github.com/nutanix/ntnx-api-golang-clients/monitoring-go-client/v4/models/monitoring/v4/error"
	import3 "github.com/nutanix/ntnx-api-golang-clients/monitoring-go-client/v4/models/prism/v4/config"
	"time"
)

/*
An alert can be resolved or acknowledged using the following parameters.
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
	  List of all the entities that are affected by the alert.
	*/
	AffectedEntities []import1.EntityReference `json:"affectedEntities,omitempty"`
	/*
	  A preconfigured or dynamically generated unique value for each alert type. For example, A1128  for storage pool space exceeded alerts.
	*/
	AlertType *string `json:"alertType,omitempty"`
	/*
	  Various categories into which this alert type can be classified. For example, hardware, storage, or license.
	*/
	Classifications []string `json:"classifications,omitempty"`
	/*
	  Name of the cluster associated with the entity.
	*/
	ClusterName *string `json:"clusterName,omitempty"`
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
	  Indicates whether the alert is acknowledged or not.
	*/
	IsAcknowledged *bool `json:"isAcknowledged,omitempty"`
	/*
	  Indicates whether the alert is auto-resolved or not.
	*/
	IsAutoResolved *bool `json:"isAutoResolved,omitempty"`
	/*
	  Indicates whether the alert is resolved or not.
	*/
	IsResolved *bool `json:"isResolved,omitempty"`
	/*
	  Indicates whether the policy associated with the alert is runnable or not.
	*/
	IsRunnable *bool `json:"isRunnable,omitempty"`
	/*
	  Flag to indicate if the alert was generated from a User-Defined Alert policy.
	*/
	IsUserDefined *bool `json:"isUserDefined,omitempty"`
	/*
	  List of knowledge base article links.
	*/
	KbArticles []string `json:"kbArticles,omitempty"`
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
	  Additional parameters associated with the alert. These parameters can be used to indicate custom key-value pairs for a given alert instance. For example, a service down in Prism Central alert can have the service name as a parameter.
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
	  The service that raised the alert.
	*/
	ServiceName *string `json:"serviceName,omitempty"`

	Severity *import1.Severity `json:"severity,omitempty"`
	/*
	  Contains information on the severity change history for alerts. If an alert was de-duplicated without change in severity, then no trail will be present.
	*/
	SeverityTrails []SeverityTrail `json:"severityTrails,omitempty"`

	SourceEntity *import1.AlertEntityReference `json:"sourceEntity,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  Title of the alert.
	*/
	Title *string `json:"title,omitempty"`
}

func (p *Alert) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias Alert

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

func (p *Alert) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Alert
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewAlert()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.AcknowledgedByUsername != nil {
		p.AcknowledgedByUsername = known.AcknowledgedByUsername
	}
	if known.AcknowledgedTime != nil {
		p.AcknowledgedTime = known.AcknowledgedTime
	}
	if known.AffectedEntities != nil {
		p.AffectedEntities = known.AffectedEntities
	}
	if known.AlertType != nil {
		p.AlertType = known.AlertType
	}
	if known.Classifications != nil {
		p.Classifications = known.Classifications
	}
	if known.ClusterName != nil {
		p.ClusterName = known.ClusterName
	}
	if known.ClusterUUID != nil {
		p.ClusterUUID = known.ClusterUUID
	}
	if known.CreationTime != nil {
		p.CreationTime = known.CreationTime
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.ImpactTypes != nil {
		p.ImpactTypes = known.ImpactTypes
	}
	if known.IsAcknowledged != nil {
		p.IsAcknowledged = known.IsAcknowledged
	}
	if known.IsAutoResolved != nil {
		p.IsAutoResolved = known.IsAutoResolved
	}
	if known.IsResolved != nil {
		p.IsResolved = known.IsResolved
	}
	if known.IsRunnable != nil {
		p.IsRunnable = known.IsRunnable
	}
	if known.IsUserDefined != nil {
		p.IsUserDefined = known.IsUserDefined
	}
	if known.KbArticles != nil {
		p.KbArticles = known.KbArticles
	}
	if known.LastUpdatedTime != nil {
		p.LastUpdatedTime = known.LastUpdatedTime
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.Message != nil {
		p.Message = known.Message
	}
	if known.MetricDetails != nil {
		p.MetricDetails = known.MetricDetails
	}
	if known.OriginatingClusterUUID != nil {
		p.OriginatingClusterUUID = known.OriginatingClusterUUID
	}
	if known.Parameters != nil {
		p.Parameters = known.Parameters
	}
	if known.ResolvedByUsername != nil {
		p.ResolvedByUsername = known.ResolvedByUsername
	}
	if known.ResolvedTime != nil {
		p.ResolvedTime = known.ResolvedTime
	}
	if known.RootCauseAnalysis != nil {
		p.RootCauseAnalysis = known.RootCauseAnalysis
	}
	if known.ServiceName != nil {
		p.ServiceName = known.ServiceName
	}
	if known.Severity != nil {
		p.Severity = known.Severity
	}
	if known.SeverityTrails != nil {
		p.SeverityTrails = known.SeverityTrails
	}
	if known.SourceEntity != nil {
		p.SourceEntity = known.SourceEntity
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}
	if known.Title != nil {
		p.Title = known.Title
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "acknowledgedByUsername")
	delete(allFields, "acknowledgedTime")
	delete(allFields, "affectedEntities")
	delete(allFields, "alertType")
	delete(allFields, "classifications")
	delete(allFields, "clusterName")
	delete(allFields, "clusterUUID")
	delete(allFields, "creationTime")
	delete(allFields, "extId")
	delete(allFields, "impactTypes")
	delete(allFields, "isAcknowledged")
	delete(allFields, "isAutoResolved")
	delete(allFields, "isResolved")
	delete(allFields, "isRunnable")
	delete(allFields, "isUserDefined")
	delete(allFields, "kbArticles")
	delete(allFields, "lastUpdatedTime")
	delete(allFields, "links")
	delete(allFields, "message")
	delete(allFields, "metricDetails")
	delete(allFields, "originatingClusterUUID")
	delete(allFields, "parameters")
	delete(allFields, "resolvedByUsername")
	delete(allFields, "resolvedTime")
	delete(allFields, "rootCauseAnalysis")
	delete(allFields, "serviceName")
	delete(allFields, "severity")
	delete(allFields, "severityTrails")
	delete(allFields, "sourceEntity")
	delete(allFields, "tenantId")
	delete(allFields, "title")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewAlert() *Alert {
	p := new(Alert)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.Alert"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	p.IsAcknowledged = new(bool)
	*p.IsAcknowledged = false
	p.IsAutoResolved = new(bool)
	*p.IsAutoResolved = false
	p.IsResolved = new(bool)
	*p.IsResolved = false
	p.IsRunnable = new(bool)
	*p.IsRunnable = false
	p.IsUserDefined = new(bool)
	*p.IsUserDefined = false

	return p
}

/*
An alert can be resolved or acknowledged using the following parameters.
*/
type AlertActionSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	ActionType *ActionType `json:"actionType"`
}

func (p *AlertActionSpec) MarshalJSON() ([]byte, error) {
	type AlertActionSpecProxy AlertActionSpec

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*AlertActionSpecProxy
		ActionType *ActionType `json:"actionType,omitempty"`
	}{
		AlertActionSpecProxy: (*AlertActionSpecProxy)(p),
		ActionType:           p.ActionType,
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

func (p *AlertActionSpec) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias AlertActionSpec
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewAlertActionSpec()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ActionType != nil {
		p.ActionType = known.ActionType
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "actionType")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewAlertActionSpec() *AlertActionSpec {
	p := new(AlertActionSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.AlertActionSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Alert specific properties associated with the policy.
*/
type AlertConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	AutoResolve *AutoResolveState `json:"autoResolve,omitempty"`

	CriticalSeverity *SeverityConfig `json:"criticalSeverity,omitempty"`

	InfoSeverity *SeverityConfig `json:"infoSeverity,omitempty"`

	WarningSeverity *SeverityConfig `json:"warningSeverity,omitempty"`
}

func (p *AlertConfig) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias AlertConfig

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

func (p *AlertConfig) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias AlertConfig
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewAlertConfig()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.AutoResolve != nil {
		p.AutoResolve = known.AutoResolve
	}
	if known.CriticalSeverity != nil {
		p.CriticalSeverity = known.CriticalSeverity
	}
	if known.InfoSeverity != nil {
		p.InfoSeverity = known.InfoSeverity
	}
	if known.WarningSeverity != nil {
		p.WarningSeverity = known.WarningSeverity
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "autoResolve")
	delete(allFields, "criticalSeverity")
	delete(allFields, "infoSeverity")
	delete(allFields, "warningSeverity")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewAlertConfig() *AlertConfig {
	p := new(AlertConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.AlertConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type AlertEmailConfiguration struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Time in HH:mm format when the alert email digest is sent daily.
	*/
	AlertEmailDigestSendTime *string `json:"alertEmailDigestSendTime,omitempty"`
	/*
	  Timezone for the time at which the alert email digest is sent daily.
	*/
	AlertEmailDigestSendTimezone *string `json:"alertEmailDigestSendTimezone,omitempty"`
	/*
	  The default Nutanix email ID to which alert emails are sent.
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
	  Indicates whether alert emails are enabled or not on default Nutanix email ID.
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
	  Indicates whether alert emails are enabled or not.
	*/
	IsEnabled *bool `json:"isEnabled,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	TunnelDetails *RemoteTunnelDetails `json:"tunnelDetails,omitempty"`
}

func (p *AlertEmailConfiguration) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias AlertEmailConfiguration

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

func (p *AlertEmailConfiguration) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias AlertEmailConfiguration
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewAlertEmailConfiguration()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.AlertEmailDigestSendTime != nil {
		p.AlertEmailDigestSendTime = known.AlertEmailDigestSendTime
	}
	if known.AlertEmailDigestSendTimezone != nil {
		p.AlertEmailDigestSendTimezone = known.AlertEmailDigestSendTimezone
	}
	if known.DefaultNutanixEmail != nil {
		p.DefaultNutanixEmail = known.DefaultNutanixEmail
	}
	if known.EmailConfigRules != nil {
		p.EmailConfigRules = known.EmailConfigRules
	}
	if known.EmailContactList != nil {
		p.EmailContactList = known.EmailContactList
	}
	if known.EmailTemplate != nil {
		p.EmailTemplate = known.EmailTemplate
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.HasDefaultNutanixEmail != nil {
		p.HasDefaultNutanixEmail = known.HasDefaultNutanixEmail
	}
	if known.IsEmailDigestEnabled != nil {
		p.IsEmailDigestEnabled = known.IsEmailDigestEnabled
	}
	if known.IsEmptyAlertEmailDigestSkipped != nil {
		p.IsEmptyAlertEmailDigestSkipped = known.IsEmptyAlertEmailDigestSkipped
	}
	if known.IsEnabled != nil {
		p.IsEnabled = known.IsEnabled
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}
	if known.TunnelDetails != nil {
		p.TunnelDetails = known.TunnelDetails
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "alertEmailDigestSendTime")
	delete(allFields, "alertEmailDigestSendTimezone")
	delete(allFields, "defaultNutanixEmail")
	delete(allFields, "emailConfigRules")
	delete(allFields, "emailContactList")
	delete(allFields, "emailTemplate")
	delete(allFields, "extId")
	delete(allFields, "hasDefaultNutanixEmail")
	delete(allFields, "isEmailDigestEnabled")
	delete(allFields, "isEmptyAlertEmailDigestSkipped")
	delete(allFields, "isEnabled")
	delete(allFields, "links")
	delete(allFields, "tenantId")
	delete(allFields, "tunnelDetails")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewAlertEmailConfiguration() *AlertEmailConfiguration {
	p := new(AlertEmailConfiguration)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.AlertEmailConfiguration"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
Parameters of the SDA that are configurable by a user.
*/
type AlertPolicyConfigurableParameter struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Equivalent name for the parameter used to display it on Prism UI.
	*/
	DisplayName *string `json:"displayName,omitempty"`
	/*
	  Unique identifier name for the parameter.
	*/
	Name *string `json:"name"`
	/*

	 */
	ParamValueItemDiscriminator_ *string `json:"$paramValueItemDiscriminator,omitempty"`
	/*
	  Captures the parameter value. Currently, it supports 4 data type, that is, string, integer, float, and boolean.
	*/
	ParamValue *OneOfAlertPolicyConfigurableParameterParamValue `json:"paramValue"`
	/*
	  Unit for the parameter. For example, sec, %, MB, GB, and so on.
	*/
	Unit *string `json:"unit,omitempty"`
}

func (p *AlertPolicyConfigurableParameter) MarshalJSON() ([]byte, error) {
	type AlertPolicyConfigurableParameterProxy AlertPolicyConfigurableParameter

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*AlertPolicyConfigurableParameterProxy
		Name       *string                                          `json:"name,omitempty"`
		ParamValue *OneOfAlertPolicyConfigurableParameterParamValue `json:"paramValue,omitempty"`
	}{
		AlertPolicyConfigurableParameterProxy: (*AlertPolicyConfigurableParameterProxy)(p),
		Name:                                  p.Name,
		ParamValue:                            p.ParamValue,
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

func (p *AlertPolicyConfigurableParameter) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias AlertPolicyConfigurableParameter
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewAlertPolicyConfigurableParameter()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DisplayName != nil {
		p.DisplayName = known.DisplayName
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.ParamValueItemDiscriminator_ != nil {
		p.ParamValueItemDiscriminator_ = known.ParamValueItemDiscriminator_
	}
	if known.ParamValue != nil {
		p.ParamValue = known.ParamValue
	}
	if known.Unit != nil {
		p.Unit = known.Unit
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "displayName")
	delete(allFields, "name")
	delete(allFields, "$paramValueItemDiscriminator")
	delete(allFields, "paramValue")
	delete(allFields, "unit")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewAlertPolicyConfigurableParameter() *AlertPolicyConfigurableParameter {
	p := new(AlertPolicyConfigurableParameter)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.AlertPolicyConfigurableParameter"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *AlertPolicyConfigurableParameter) GetParamValue() interface{} {
	if nil == p.ParamValue {
		return nil
	}
	return p.ParamValue.GetValue()
}

func (p *AlertPolicyConfigurableParameter) SetParamValue(v interface{}) error {
	if nil == p.ParamValue {
		p.ParamValue = NewOneOfAlertPolicyConfigurableParameterParamValue()
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
Arguments specifying details of the archive to be collected.
*/
type ArchiveOpts struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Custom name of the archive folder.
	*/
	ArchiveName *string `json:"archiveName,omitempty"`
	/*

	 */
	UploadParamsItemDiscriminator_ *string `json:"$uploadParamsItemDiscriminator,omitempty"`
	/*
	  This parameter specifies the target/destination location of collected archive. User must specify one of the below target/destination.
	- LocalUploadParams: Use this option to store archive on local disk.
	  Full archive path can be found in the completion_details field of the task.
	- NtnxServerUploadParams: Use this option to upload archive to Nutanix server.
	- CustomServerUploadParams: Use this option to upload archive to custom FTP/SFTP server.
	- StorageContainerUploadParams: Use this option to upload archive to a storage container.
	  Localhost storage container uploads are not supported on PC.
	*/
	UploadParams *OneOfArchiveOptsUploadParams `json:"uploadParams"`
}

func (p *ArchiveOpts) MarshalJSON() ([]byte, error) {
	type ArchiveOptsProxy ArchiveOpts

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*ArchiveOptsProxy
		UploadParams *OneOfArchiveOptsUploadParams `json:"uploadParams,omitempty"`
	}{
		ArchiveOptsProxy: (*ArchiveOptsProxy)(p),
		UploadParams:     p.UploadParams,
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

func (p *ArchiveOpts) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ArchiveOpts
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewArchiveOpts()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ArchiveName != nil {
		p.ArchiveName = known.ArchiveName
	}
	if known.UploadParamsItemDiscriminator_ != nil {
		p.UploadParamsItemDiscriminator_ = known.UploadParamsItemDiscriminator_
	}
	if known.UploadParams != nil {
		p.UploadParams = known.UploadParams
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "archiveName")
	delete(allFields, "$uploadParamsItemDiscriminator")
	delete(allFields, "uploadParams")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewArchiveOpts() *ArchiveOpts {
	p := new(ArchiveOpts)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.ArchiveOpts"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	p.ArchiveName = new(string)
	*p.ArchiveName = "NTNX-Log"

	return p
}

func (p *ArchiveOpts) GetUploadParams() interface{} {
	if nil == p.UploadParams {
		return nil
	}
	return p.UploadParams.GetValue()
}

func (p *ArchiveOpts) SetUploadParams(v interface{}) error {
	if nil == p.UploadParams {
		p.UploadParams = NewOneOfArchiveOptsUploadParams()
	}
	e := p.UploadParams.SetValue(v)
	if nil == e {
		if nil == p.UploadParamsItemDiscriminator_ {
			p.UploadParamsItemDiscriminator_ = new(string)
		}
		*p.UploadParamsItemDiscriminator_ = *p.UploadParams.Discriminator
	}
	return e
}

type Audit struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of all the entities that are affected by the event or audit.
	*/
	AffectedEntities []import1.EntityReference `json:"affectedEntities,omitempty"`
	/*
	  The unique name for a given audit type. For example, VMCloneAudit or VMDeleteAudit.
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
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	UserReference *UserReference `json:"userReference,omitempty"`
}

func (p *Audit) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias Audit

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

func (p *Audit) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Audit
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewAudit()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.AffectedEntities != nil {
		p.AffectedEntities = known.AffectedEntities
	}
	if known.AuditType != nil {
		p.AuditType = known.AuditType
	}
	if known.ClusterReference != nil {
		p.ClusterReference = known.ClusterReference
	}
	if known.CreationTime != nil {
		p.CreationTime = known.CreationTime
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.Message != nil {
		p.Message = known.Message
	}
	if known.OperationEndTime != nil {
		p.OperationEndTime = known.OperationEndTime
	}
	if known.OperationStartTime != nil {
		p.OperationStartTime = known.OperationStartTime
	}
	if known.OperationType != nil {
		p.OperationType = known.OperationType
	}
	if known.Parameters != nil {
		p.Parameters = known.Parameters
	}
	if known.ServiceName != nil {
		p.ServiceName = known.ServiceName
	}
	if known.SourceEntity != nil {
		p.SourceEntity = known.SourceEntity
	}
	if known.Status != nil {
		p.Status = known.Status
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}
	if known.UserReference != nil {
		p.UserReference = known.UserReference
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "affectedEntities")
	delete(allFields, "auditType")
	delete(allFields, "clusterReference")
	delete(allFields, "creationTime")
	delete(allFields, "extId")
	delete(allFields, "links")
	delete(allFields, "message")
	delete(allFields, "operationEndTime")
	delete(allFields, "operationStartTime")
	delete(allFields, "operationType")
	delete(allFields, "parameters")
	delete(allFields, "serviceName")
	delete(allFields, "sourceEntity")
	delete(allFields, "status")
	delete(allFields, "tenantId")
	delete(allFields, "userReference")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewAudit() *Audit {
	p := new(Audit)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.Audit"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The source entity associated with the alert, event, or audit.
*/
type AuditEntityReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  UUID of the entity.
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

func (p *AuditEntityReference) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias AuditEntityReference

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

func (p *AuditEntityReference) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias AuditEntityReference
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewAuditEntityReference()

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
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.Type != nil {
		p.Type = known.Type
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "extId")
	delete(allFields, "name")
	delete(allFields, "type")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewAuditEntityReference() *AuditEntityReference {
	p := new(AuditEntityReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.AuditEntityReference"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Indicates the auto-resolve configuration for the alert generated through the policy.
*/
type AutoResolveState int

const (
	AUTORESOLVESTATE_UNKNOWN       AutoResolveState = 0
	AUTORESOLVESTATE_REDACTED      AutoResolveState = 1
	AUTORESOLVESTATE_ENABLED       AutoResolveState = 2
	AUTORESOLVESTATE_DISABLED      AutoResolveState = 3
	AUTORESOLVESTATE_NOT_SUPPORTED AutoResolveState = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *AutoResolveState) name(index int) string {
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
func (e AutoResolveState) GetName() string {
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
func (e *AutoResolveState) index(name string) AutoResolveState {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ENABLED",
		"DISABLED",
		"NOT_SUPPORTED",
	}
	for idx := range names {
		if names[idx] == name {
			return AutoResolveState(idx)
		}
	}
	return AUTORESOLVESTATE_UNKNOWN
}

func (e *AutoResolveState) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for AutoResolveState:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *AutoResolveState) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e AutoResolveState) Ref() *AutoResolveState {
	return &e
}

/*
Captures values of the parameter when it's of type boolean.
*/
type BooleanConfigurableParamValue struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Captures the current value of the parameter.
	*/
	CurrentBoolValue *bool `json:"currentBoolValue"`
	/*
	  Captures the default value of the parameter.
	*/
	DefaultBoolValue *bool `json:"defaultBoolValue,omitempty"`
}

func (p *BooleanConfigurableParamValue) MarshalJSON() ([]byte, error) {
	type BooleanConfigurableParamValueProxy BooleanConfigurableParamValue

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*BooleanConfigurableParamValueProxy
		CurrentBoolValue *bool `json:"currentBoolValue,omitempty"`
	}{
		BooleanConfigurableParamValueProxy: (*BooleanConfigurableParamValueProxy)(p),
		CurrentBoolValue:                   p.CurrentBoolValue,
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

func (p *BooleanConfigurableParamValue) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias BooleanConfigurableParamValue
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewBooleanConfigurableParamValue()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.CurrentBoolValue != nil {
		p.CurrentBoolValue = known.CurrentBoolValue
	}
	if known.DefaultBoolValue != nil {
		p.DefaultBoolValue = known.DefaultBoolValue
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "currentBoolValue")
	delete(allFields, "defaultBoolValue")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewBooleanConfigurableParamValue() *BooleanConfigurableParamValue {
	p := new(BooleanConfigurableParamValue)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.BooleanConfigurableParamValue"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
SDA policy parameters that may differ across clusters since each cluster can run on different NCC versions. Each cluster will have an individual entry.
*/
type ClusterConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	AlertConfig *AlertConfig `json:"alertConfig,omitempty"`
	/*
	  Parameters of the SDA that are configurable by a user.
	*/
	ConfigurableParameters []AlertPolicyConfigurableParameter `json:"configurableParameters,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Indicates whether the SDA policy is enabled or not on the cluster.
	*/
	IsEnabled *bool `json:"isEnabled,omitempty"`
	/*
	  Name of the user who made the latest update to this policy. Its value will be Nutanix if the last update is due to an upgrade event.
	*/
	LastModifiedByUser *string `json:"lastModifiedByUser,omitempty"`
	/*
	  Time in ISO 8601 format when the SDA policy was last modified. It gets automatically updated by the Nutanix service from the user context during an update event.
	*/
	LastModifiedTime *time.Time `json:"lastModifiedTime,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Interval in seconds for periodically executing the SDA policy. This will not be set for policies with the type NOT_SCHEDULED & EVENT_DRIVEN.
	*/
	ScheduleIntervalSeconds *int `json:"scheduleIntervalSeconds,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *ClusterConfig) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ClusterConfig

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

func (p *ClusterConfig) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ClusterConfig
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewClusterConfig()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.AlertConfig != nil {
		p.AlertConfig = known.AlertConfig
	}
	if known.ConfigurableParameters != nil {
		p.ConfigurableParameters = known.ConfigurableParameters
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.IsEnabled != nil {
		p.IsEnabled = known.IsEnabled
	}
	if known.LastModifiedByUser != nil {
		p.LastModifiedByUser = known.LastModifiedByUser
	}
	if known.LastModifiedTime != nil {
		p.LastModifiedTime = known.LastModifiedTime
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.ScheduleIntervalSeconds != nil {
		p.ScheduleIntervalSeconds = known.ScheduleIntervalSeconds
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "alertConfig")
	delete(allFields, "configurableParameters")
	delete(allFields, "extId")
	delete(allFields, "isEnabled")
	delete(allFields, "lastModifiedByUser")
	delete(allFields, "lastModifiedTime")
	delete(allFields, "links")
	delete(allFields, "scheduleIntervalSeconds")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewClusterConfig() *ClusterConfig {
	p := new(ClusterConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.ClusterConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type ClusterConfigProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	AlertConfig *AlertConfig `json:"alertConfig,omitempty"`
	/*
	  Parameters of the SDA that are configurable by a user.
	*/
	ConfigurableParameters []AlertPolicyConfigurableParameter `json:"configurableParameters,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Indicates whether the SDA policy is enabled or not on the cluster.
	*/
	IsEnabled *bool `json:"isEnabled,omitempty"`
	/*
	  Name of the user who made the latest update to this policy. Its value will be Nutanix if the last update is due to an upgrade event.
	*/
	LastModifiedByUser *string `json:"lastModifiedByUser,omitempty"`
	/*
	  Time in ISO 8601 format when the SDA policy was last modified. It gets automatically updated by the Nutanix service from the user context during an update event.
	*/
	LastModifiedTime *time.Time `json:"lastModifiedTime,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Interval in seconds for periodically executing the SDA policy. This will not be set for policies with the type NOT_SCHEDULED & EVENT_DRIVEN.
	*/
	ScheduleIntervalSeconds *int `json:"scheduleIntervalSeconds,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *ClusterConfigProjection) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ClusterConfigProjection

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

func (p *ClusterConfigProjection) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ClusterConfigProjection
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewClusterConfigProjection()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.AlertConfig != nil {
		p.AlertConfig = known.AlertConfig
	}
	if known.ConfigurableParameters != nil {
		p.ConfigurableParameters = known.ConfigurableParameters
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.IsEnabled != nil {
		p.IsEnabled = known.IsEnabled
	}
	if known.LastModifiedByUser != nil {
		p.LastModifiedByUser = known.LastModifiedByUser
	}
	if known.LastModifiedTime != nil {
		p.LastModifiedTime = known.LastModifiedTime
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.ScheduleIntervalSeconds != nil {
		p.ScheduleIntervalSeconds = known.ScheduleIntervalSeconds
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "alertConfig")
	delete(allFields, "configurableParameters")
	delete(allFields, "extId")
	delete(allFields, "isEnabled")
	delete(allFields, "lastModifiedByUser")
	delete(allFields, "lastModifiedTime")
	delete(allFields, "links")
	delete(allFields, "scheduleIntervalSeconds")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewClusterConfigProjection() *ClusterConfigProjection {
	p := new(ClusterConfigProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.ClusterConfigProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Indicates the cluster type.
*/
type ClusterType int

const (
	CLUSTERTYPE_UNKNOWN  ClusterType = 0
	CLUSTERTYPE_REDACTED ClusterType = 1
	CLUSTERTYPE_PC       ClusterType = 2
	CLUSTERTYPE_PE       ClusterType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ClusterType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PC",
		"PE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ClusterType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PC",
		"PE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ClusterType) index(name string) ClusterType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PC",
		"PE",
	}
	for idx := range names {
		if names[idx] == name {
			return ClusterType(idx)
		}
	}
	return CLUSTERTYPE_UNKNOWN
}

func (e *ClusterType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ClusterType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ClusterType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ClusterType) Ref() *ClusterType {
	return &e
}

/*
REST response for all response codes in API path /monitoring/v4.2/serviceability/clusters/{extId}/$actions/collect-logs Post operation
*/
type CollectLogsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCollectLogsApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *CollectLogsApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias CollectLogsApiResponse

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

func (p *CollectLogsApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias CollectLogsApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewCollectLogsApiResponse()

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

func NewCollectLogsApiResponse() *CollectLogsApiResponse {
	p := new(CollectLogsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.CollectLogsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CollectLogsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CollectLogsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCollectLogsApiResponseData()
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
Status of the remote tunnel or service that is running on top of the remote tunnel. For example, pulse or alert email.
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

func (p *CommunicationStatus) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias CommunicationStatus

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

func (p *CommunicationStatus) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias CommunicationStatus
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewCommunicationStatus()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.LastChangedTime != nil {
		p.LastChangedTime = known.LastChangedTime
	}
	if known.LastCheckedTime != nil {
		p.LastCheckedTime = known.LastCheckedTime
	}
	if known.LastSuccessfulTransmissionTime != nil {
		p.LastSuccessfulTransmissionTime = known.LastSuccessfulTransmissionTime
	}
	if known.Message != nil {
		p.Message = known.Message
	}
	if known.Status != nil {
		p.Status = known.Status
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "lastChangedTime")
	delete(allFields, "lastCheckedTime")
	delete(allFields, "lastSuccessfulTransmissionTime")
	delete(allFields, "message")
	delete(allFields, "status")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewCommunicationStatus() *CommunicationStatus {
	p := new(CommunicationStatus)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.CommunicationStatus"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Conditions to be met to trigger the alert. Conditions are expressed in FIQL.
*/
type Condition struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The metric key. Allowed values of the metrics list can be found at https://portal.nutanix.com/page/documents/details?targetId=Prism-Central-Guide-vpc_2022_9:mul-alerts-user-created-metrics-r.html.
	*/
	MetricName *string `json:"metricName"`

	Operator *import1.ComparisonOperator `json:"operator"`
	/*

	 */
	ThresholdValueItemDiscriminator_ *string `json:"$thresholdValueItemDiscriminator,omitempty"`
	/*
	  The threshold value that was used for the condition evaluation.
	*/
	ThresholdValue *OneOfConditionThresholdValue `json:"thresholdValue"`
}

func (p *Condition) MarshalJSON() ([]byte, error) {
	type ConditionProxy Condition

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*ConditionProxy
		MetricName     *string                       `json:"metricName,omitempty"`
		Operator       *import1.ComparisonOperator   `json:"operator,omitempty"`
		ThresholdValue *OneOfConditionThresholdValue `json:"thresholdValue,omitempty"`
	}{
		ConditionProxy: (*ConditionProxy)(p),
		MetricName:     p.MetricName,
		Operator:       p.Operator,
		ThresholdValue: p.ThresholdValue,
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

func (p *Condition) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Condition
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewCondition()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.MetricName != nil {
		p.MetricName = known.MetricName
	}
	if known.Operator != nil {
		p.Operator = known.Operator
	}
	if known.ThresholdValueItemDiscriminator_ != nil {
		p.ThresholdValueItemDiscriminator_ = known.ThresholdValueItemDiscriminator_
	}
	if known.ThresholdValue != nil {
		p.ThresholdValue = known.ThresholdValue
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "metricName")
	delete(allFields, "operator")
	delete(allFields, "$thresholdValueItemDiscriminator")
	delete(allFields, "thresholdValue")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewCondition() *Condition {
	p := new(Condition)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.Condition"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
Valid values for the parameter.
*/
type ConfigurableParamValueRange struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	MaximumValueItemDiscriminator_ *string `json:"$maximumValueItemDiscriminator,omitempty"`
	/*
	  Captures the maximum value the parameter can have.
	*/
	MaximumValue *OneOfConfigurableParamValueRangeMaximumValue `json:"maximumValue"`
	/*

	 */
	MinimumValueItemDiscriminator_ *string `json:"$minimumValueItemDiscriminator,omitempty"`
	/*
	  Captures the minimum value the parameter can have.
	*/
	MinimumValue *OneOfConfigurableParamValueRangeMinimumValue `json:"minimumValue"`
}

func (p *ConfigurableParamValueRange) MarshalJSON() ([]byte, error) {
	type ConfigurableParamValueRangeProxy ConfigurableParamValueRange

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*ConfigurableParamValueRangeProxy
		MaximumValue *OneOfConfigurableParamValueRangeMaximumValue `json:"maximumValue,omitempty"`
		MinimumValue *OneOfConfigurableParamValueRangeMinimumValue `json:"minimumValue,omitempty"`
	}{
		ConfigurableParamValueRangeProxy: (*ConfigurableParamValueRangeProxy)(p),
		MaximumValue:                     p.MaximumValue,
		MinimumValue:                     p.MinimumValue,
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

func (p *ConfigurableParamValueRange) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ConfigurableParamValueRange
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewConfigurableParamValueRange()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.MaximumValueItemDiscriminator_ != nil {
		p.MaximumValueItemDiscriminator_ = known.MaximumValueItemDiscriminator_
	}
	if known.MaximumValue != nil {
		p.MaximumValue = known.MaximumValue
	}
	if known.MinimumValueItemDiscriminator_ != nil {
		p.MinimumValueItemDiscriminator_ = known.MinimumValueItemDiscriminator_
	}
	if known.MinimumValue != nil {
		p.MinimumValue = known.MinimumValue
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$maximumValueItemDiscriminator")
	delete(allFields, "maximumValue")
	delete(allFields, "$minimumValueItemDiscriminator")
	delete(allFields, "minimumValue")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewConfigurableParamValueRange() *ConfigurableParamValueRange {
	p := new(ConfigurableParamValueRange)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.ConfigurableParamValueRange"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ConfigurableParamValueRange) GetMaximumValue() interface{} {
	if nil == p.MaximumValue {
		return nil
	}
	return p.MaximumValue.GetValue()
}

func (p *ConfigurableParamValueRange) SetMaximumValue(v interface{}) error {
	if nil == p.MaximumValue {
		p.MaximumValue = NewOneOfConfigurableParamValueRangeMaximumValue()
	}
	e := p.MaximumValue.SetValue(v)
	if nil == e {
		if nil == p.MaximumValueItemDiscriminator_ {
			p.MaximumValueItemDiscriminator_ = new(string)
		}
		*p.MaximumValueItemDiscriminator_ = *p.MaximumValue.Discriminator
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
	  Unique UUID associated with the User-Defined Alert policy, that conflicts with the given policy.
	*/
	ExtId *string `json:"extId"`
}

func (p *ConflictingPolicy) MarshalJSON() ([]byte, error) {
	type ConflictingPolicyProxy ConflictingPolicy

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*ConflictingPolicyProxy
		ExtId *string `json:"extId,omitempty"`
	}{
		ConflictingPolicyProxy: (*ConflictingPolicyProxy)(p),
		ExtId:                  p.ExtId,
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

func (p *ConflictingPolicy) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ConflictingPolicy
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewConflictingPolicy()

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

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "extId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewConflictingPolicy() *ConflictingPolicy {
	p := new(ConflictingPolicy)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.ConflictingPolicy"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Connection status.
*/
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
REST response for all response codes in API path /monitoring/v4.2/serviceability/alerts/user-defined-policies Post operation
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

func (p *CreateUdaPolicyApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias CreateUdaPolicyApiResponse

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

func (p *CreateUdaPolicyApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias CreateUdaPolicyApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewCreateUdaPolicyApiResponse()

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

func NewCreateUdaPolicyApiResponse() *CreateUdaPolicyApiResponse {
	p := new(CreateUdaPolicyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.CreateUdaPolicyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
Credentials for the destination where the collected archive will be uploaded. Mandatory for SFTP and FTP protocol.
*/
type Credential struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Path to the key file that will be used for authentication of the archive destination.
	*/
	KeyFilePath *string `json:"keyFilePath,omitempty"`
	/*
	  Password to access the archive destination.
	*/
	Password *string `json:"password,omitempty"`
	/*
	  Username to access the archive destination. If it is not provided, the username that was used to call the API will be used.
	*/
	UserName *string `json:"userName"`
}

func (p *Credential) MarshalJSON() ([]byte, error) {
	type CredentialProxy Credential

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*CredentialProxy
		UserName *string `json:"userName,omitempty"`
	}{
		CredentialProxy: (*CredentialProxy)(p),
		UserName:        p.UserName,
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

func (p *Credential) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Credential
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewCredential()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.KeyFilePath != nil {
		p.KeyFilePath = known.KeyFilePath
	}
	if known.Password != nil {
		p.Password = known.Password
	}
	if known.UserName != nil {
		p.UserName = known.UserName
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "keyFilePath")
	delete(allFields, "password")
	delete(allFields, "userName")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewCredential() *Credential {
	p := new(Credential)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.Credential"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The parameters for uploading the archive if the destination is a custom server.
*/
type CustomServerUploadParams struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Credential *Credential `json:"credential"`
	/*
	  Path where the archive will be stored on the custom server. The path should pre-exist. If not specified, the bundle will get stored in the location as set in $HOME.
	*/
	Path *string `json:"path,omitempty"`
	/*
	  Port of the target. Mandatory for SFTP and FTP protocol.
	*/
	Port *int `json:"port,omitempty"`

	Protocol *ServerUploadProtocol `json:"protocol"`

	ServerAddress *import5.IPAddressOrFQDN `json:"serverAddress"`
}

func (p *CustomServerUploadParams) MarshalJSON() ([]byte, error) {
	type CustomServerUploadParamsProxy CustomServerUploadParams

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*CustomServerUploadParamsProxy
		Credential    *Credential              `json:"credential,omitempty"`
		Protocol      *ServerUploadProtocol    `json:"protocol,omitempty"`
		ServerAddress *import5.IPAddressOrFQDN `json:"serverAddress,omitempty"`
	}{
		CustomServerUploadParamsProxy: (*CustomServerUploadParamsProxy)(p),
		Credential:                    p.Credential,
		Protocol:                      p.Protocol,
		ServerAddress:                 p.ServerAddress,
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

func (p *CustomServerUploadParams) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias CustomServerUploadParams
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewCustomServerUploadParams()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Credential != nil {
		p.Credential = known.Credential
	}
	if known.Path != nil {
		p.Path = known.Path
	}
	if known.Port != nil {
		p.Port = known.Port
	}
	if known.Protocol != nil {
		p.Protocol = known.Protocol
	}
	if known.ServerAddress != nil {
		p.ServerAddress = known.ServerAddress
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "credential")
	delete(allFields, "path")
	delete(allFields, "port")
	delete(allFields, "protocol")
	delete(allFields, "serverAddress")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewCustomServerUploadParams() *CustomServerUploadParams {
	p := new(CustomServerUploadParams)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.CustomServerUploadParams"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /monitoring/v4.2/serviceability/alerts/user-defined-policies/{extId} Delete operation
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

func (p *DeleteUdaPolicyApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DeleteUdaPolicyApiResponse

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

func (p *DeleteUdaPolicyApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DeleteUdaPolicyApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewDeleteUdaPolicyApiResponse()

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

func NewDeleteUdaPolicyApiResponse() *DeleteUdaPolicyApiResponse {
	p := new(DeleteUdaPolicyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.DeleteUdaPolicyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
Status of the remote tunnel that is used to send alert emails.
*/
type EmailConfigurationRule struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Cluster UUIDs to which this rule applies.
	*/
	ClusterUuids []string `json:"clusterUuids,omitempty"`
	/*
	  Indicates whether to include a global email contact list.
	*/
	HasGlobalEmailContactList *bool `json:"hasGlobalEmailContactList,omitempty"`

	ImpactTypes []import1.ImpactType `json:"impactTypes,omitempty"`
	/*
	  Indicates whether the configuration rule is enabled or not.
	*/
	IsEnabled *bool `json:"isEnabled,omitempty"`
	/*
	  List of phrases to match the alert.
	*/
	MatchPhrases []string `json:"matchPhrases,omitempty"`
	/*
	  List of recipients who will receive emails.
	*/
	Recipients []string `json:"recipients,omitempty"`

	Severities []import1.Severity `json:"severities,omitempty"`
}

func (p *EmailConfigurationRule) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias EmailConfigurationRule

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

func (p *EmailConfigurationRule) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias EmailConfigurationRule
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewEmailConfigurationRule()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ClusterUuids != nil {
		p.ClusterUuids = known.ClusterUuids
	}
	if known.HasGlobalEmailContactList != nil {
		p.HasGlobalEmailContactList = known.HasGlobalEmailContactList
	}
	if known.ImpactTypes != nil {
		p.ImpactTypes = known.ImpactTypes
	}
	if known.IsEnabled != nil {
		p.IsEnabled = known.IsEnabled
	}
	if known.MatchPhrases != nil {
		p.MatchPhrases = known.MatchPhrases
	}
	if known.Recipients != nil {
		p.Recipients = known.Recipients
	}
	if known.Severities != nil {
		p.Severities = known.Severities
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "clusterUuids")
	delete(allFields, "hasGlobalEmailContactList")
	delete(allFields, "impactTypes")
	delete(allFields, "isEnabled")
	delete(allFields, "matchPhrases")
	delete(allFields, "recipients")
	delete(allFields, "severities")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewEmailConfigurationRule() *EmailConfigurationRule {
	p := new(EmailConfigurationRule)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.EmailConfigurationRule"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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

func (p *EmailTemplate) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias EmailTemplate

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

func (p *EmailTemplate) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias EmailTemplate
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewEmailTemplate()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.BodySuffix != nil {
		p.BodySuffix = known.BodySuffix
	}
	if known.SubjectPrefix != nil {
		p.SubjectPrefix = known.SubjectPrefix
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "bodySuffix")
	delete(allFields, "subjectPrefix")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewEmailTemplate() *EmailTemplate {
	p := new(EmailTemplate)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.EmailTemplate"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Entity UUID on which the User-Defined Alert policy should be set up.
*/
type EntityFilter struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Entity UUID on which the User-Defined Alert policy should be set up.
	*/
	ExtId *string `json:"extId"`
}

func (p *EntityFilter) MarshalJSON() ([]byte, error) {
	type EntityFilterProxy EntityFilter

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*EntityFilterProxy
		ExtId *string `json:"extId,omitempty"`
	}{
		EntityFilterProxy: (*EntityFilterProxy)(p),
		ExtId:             p.ExtId,
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

func (p *EntityFilter) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias EntityFilter
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewEntityFilter()

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

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "extId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewEntityFilter() *EntityFilter {
	p := new(EntityFilter)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.EntityFilter"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Entity type against which the alert is raised.
*/
type EntityType int

const (
	ENTITYTYPE_UNKNOWN                     EntityType = 0
	ENTITYTYPE_REDACTED                    EntityType = 1
	ENTITYTYPE_DISK                        EntityType = 2
	ENTITYTYPE_VIRTUAL_DISK                EntityType = 3
	ENTITYTYPE_VM                          EntityType = 4
	ENTITYTYPE_CONTAINER                   EntityType = 5
	ENTITYTYPE_HOST                        EntityType = 6
	ENTITYTYPE_CLUSTER                     EntityType = 7
	ENTITYTYPE_VOLUME_GROUP                EntityType = 8
	ENTITYTYPE_FILE_SERVER                 EntityType = 9
	ENTITYTYPE_HOST_NIC                    EntityType = 10
	ENTITYTYPE_RECOVERY_PLAN               EntityType = 11
	ENTITYTYPE_PROTECTION_DOMAIN           EntityType = 12
	ENTITYTYPE_PROTECTION_RULE             EntityType = 13
	ENTITYTYPE_ACTION_RULE_RESULT          EntityType = 14
	ENTITYTYPE_AVAILABILITY_ZONE_PHYSICAL  EntityType = 15
	ENTITYTYPE_CONSISTENCY_GROUP_CONFIG    EntityType = 16
	ENTITYTYPE_OVA                         EntityType = 17
	ENTITYTYPE_VOLUME_GROUP_CONFIG         EntityType = 18
	ENTITYTYPE_VPN_CONNECTION              EntityType = 19
	ENTITYTYPE_REMOTE_SITE                 EntityType = 20
	ENTITYTYPE_STORAGE_POOL                EntityType = 21
	ENTITYTYPE_LAYER2_STRETCH              EntityType = 22
	ENTITYTYPE_ATLAS_VIRTUAL_NETWORK       EntityType = 23
	ENTITYTYPE_IPFIX_EXPORTER              EntityType = 24
	ENTITYTYPE_ATLAS_ROUTING_POLICY        EntityType = 25
	ENTITYTYPE_BGP_SESSION                 EntityType = 26
	ENTITYTYPE_MSP                         EntityType = 27
	ENTITYTYPE_VPN_GATEWAY                 EntityType = 28
	ENTITYTYPE_ATLAS_LOAD_BALANCER_SESSION EntityType = 29
	ENTITYTYPE_ATLAS_FLOW_GATEWAY          EntityType = 30
	ENTITYTYPE_PRISM_CENTRAL               EntityType = 31
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *EntityType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"DISK",
		"VIRTUAL_DISK",
		"VM",
		"CONTAINER",
		"HOST",
		"CLUSTER",
		"VOLUME_GROUP",
		"FILE_SERVER",
		"HOST_NIC",
		"RECOVERY_PLAN",
		"PROTECTION_DOMAIN",
		"PROTECTION_RULE",
		"ACTION_RULE_RESULT",
		"AVAILABILITY_ZONE_PHYSICAL",
		"CONSISTENCY_GROUP_CONFIG",
		"OVA",
		"VOLUME_GROUP_CONFIG",
		"VPN_CONNECTION",
		"REMOTE_SITE",
		"STORAGE_POOL",
		"LAYER2_STRETCH",
		"ATLAS_VIRTUAL_NETWORK",
		"IPFIX_EXPORTER",
		"ATLAS_ROUTING_POLICY",
		"BGP_SESSION",
		"MSP",
		"VPN_GATEWAY",
		"ATLAS_LOAD_BALANCER_SESSION",
		"ATLAS_FLOW_GATEWAY",
		"PRISM_CENTRAL",
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
		"DISK",
		"VIRTUAL_DISK",
		"VM",
		"CONTAINER",
		"HOST",
		"CLUSTER",
		"VOLUME_GROUP",
		"FILE_SERVER",
		"HOST_NIC",
		"RECOVERY_PLAN",
		"PROTECTION_DOMAIN",
		"PROTECTION_RULE",
		"ACTION_RULE_RESULT",
		"AVAILABILITY_ZONE_PHYSICAL",
		"CONSISTENCY_GROUP_CONFIG",
		"OVA",
		"VOLUME_GROUP_CONFIG",
		"VPN_CONNECTION",
		"REMOTE_SITE",
		"STORAGE_POOL",
		"LAYER2_STRETCH",
		"ATLAS_VIRTUAL_NETWORK",
		"IPFIX_EXPORTER",
		"ATLAS_ROUTING_POLICY",
		"BGP_SESSION",
		"MSP",
		"VPN_GATEWAY",
		"ATLAS_LOAD_BALANCER_SESSION",
		"ATLAS_FLOW_GATEWAY",
		"PRISM_CENTRAL",
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
		"DISK",
		"VIRTUAL_DISK",
		"VM",
		"CONTAINER",
		"HOST",
		"CLUSTER",
		"VOLUME_GROUP",
		"FILE_SERVER",
		"HOST_NIC",
		"RECOVERY_PLAN",
		"PROTECTION_DOMAIN",
		"PROTECTION_RULE",
		"ACTION_RULE_RESULT",
		"AVAILABILITY_ZONE_PHYSICAL",
		"CONSISTENCY_GROUP_CONFIG",
		"OVA",
		"VOLUME_GROUP_CONFIG",
		"VPN_CONNECTION",
		"REMOTE_SITE",
		"STORAGE_POOL",
		"LAYER2_STRETCH",
		"ATLAS_VIRTUAL_NETWORK",
		"IPFIX_EXPORTER",
		"ATLAS_ROUTING_POLICY",
		"BGP_SESSION",
		"MSP",
		"VPN_GATEWAY",
		"ATLAS_LOAD_BALANCER_SESSION",
		"ATLAS_FLOW_GATEWAY",
		"PRISM_CENTRAL",
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

type Event struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of all the entities that are affected by the event or audit.
	*/
	AffectedEntities []import1.EntityReference `json:"affectedEntities,omitempty"`
	/*
	  Various categories into which this event type can be classified. For example, hardware, storage, or license.
	*/
	Classifications []string `json:"classifications,omitempty"`
	/*
	  Name of the cluster associated with the entity.
	*/
	ClusterName *string `json:"clusterName,omitempty"`
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

	OperationType *import1.OperationType `json:"operationType,omitempty"`
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
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *Event) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias Event

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

func (p *Event) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Event
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewEvent()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.AffectedEntities != nil {
		p.AffectedEntities = known.AffectedEntities
	}
	if known.Classifications != nil {
		p.Classifications = known.Classifications
	}
	if known.ClusterName != nil {
		p.ClusterName = known.ClusterName
	}
	if known.ClusterUUID != nil {
		p.ClusterUUID = known.ClusterUUID
	}
	if known.CreationTime != nil {
		p.CreationTime = known.CreationTime
	}
	if known.EventType != nil {
		p.EventType = known.EventType
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.Message != nil {
		p.Message = known.Message
	}
	if known.MetricDetails != nil {
		p.MetricDetails = known.MetricDetails
	}
	if known.OperationType != nil {
		p.OperationType = known.OperationType
	}
	if known.Parameters != nil {
		p.Parameters = known.Parameters
	}
	if known.ServiceName != nil {
		p.ServiceName = known.ServiceName
	}
	if known.SourceClusterUUID != nil {
		p.SourceClusterUUID = known.SourceClusterUUID
	}
	if known.SourceEntity != nil {
		p.SourceEntity = known.SourceEntity
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "affectedEntities")
	delete(allFields, "classifications")
	delete(allFields, "clusterName")
	delete(allFields, "clusterUUID")
	delete(allFields, "creationTime")
	delete(allFields, "eventType")
	delete(allFields, "extId")
	delete(allFields, "links")
	delete(allFields, "message")
	delete(allFields, "metricDetails")
	delete(allFields, "operationType")
	delete(allFields, "parameters")
	delete(allFields, "serviceName")
	delete(allFields, "sourceClusterUUID")
	delete(allFields, "sourceEntity")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewEvent() *Event {
	p := new(Event)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.Event"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The source entity associated with the alert, event, or audit.
*/
type EventEntityReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  UUID of the entity.
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

func (p *EventEntityReference) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias EventEntityReference

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

func (p *EventEntityReference) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias EventEntityReference
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewEventEntityReference()

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
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.Type != nil {
		p.Type = known.Type
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "extId")
	delete(allFields, "name")
	delete(allFields, "type")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewEventEntityReference() *EventEntityReference {
	p := new(EventEntityReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.EventEntityReference"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /monitoring/v4.2/serviceability/alerts/user-defined-policies/$actions/find-conflicts Post operation
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

func (p *FindConflictingUdaPoliciesApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias FindConflictingUdaPoliciesApiResponse

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

func (p *FindConflictingUdaPoliciesApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias FindConflictingUdaPoliciesApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewFindConflictingUdaPoliciesApiResponse()

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

func NewFindConflictingUdaPoliciesApiResponse() *FindConflictingUdaPoliciesApiResponse {
	p := new(FindConflictingUdaPoliciesApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.FindConflictingUdaPoliciesApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
Captures values of the parameter when it's of type float.
*/
type FloatConfigurableParamValue struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Captures the current value of the parameter.
	*/
	CurrentFloatValue *float32 `json:"currentFloatValue"`
	/*
	  Captures the default value of the parameter.
	*/
	DefaultFloatValue *float32 `json:"defaultFloatValue,omitempty"`
	/*
	  Valid values for the parameter.
	*/
	ValidValueRanges []ConfigurableParamValueRange `json:"validValueRanges,omitempty"`
}

func (p *FloatConfigurableParamValue) MarshalJSON() ([]byte, error) {
	type FloatConfigurableParamValueProxy FloatConfigurableParamValue

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*FloatConfigurableParamValueProxy
		CurrentFloatValue *float32 `json:"currentFloatValue,omitempty"`
	}{
		FloatConfigurableParamValueProxy: (*FloatConfigurableParamValueProxy)(p),
		CurrentFloatValue:                p.CurrentFloatValue,
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

func (p *FloatConfigurableParamValue) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias FloatConfigurableParamValue
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewFloatConfigurableParamValue()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.CurrentFloatValue != nil {
		p.CurrentFloatValue = known.CurrentFloatValue
	}
	if known.DefaultFloatValue != nil {
		p.DefaultFloatValue = known.DefaultFloatValue
	}
	if known.ValidValueRanges != nil {
		p.ValidValueRanges = known.ValidValueRanges
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "currentFloatValue")
	delete(allFields, "defaultFloatValue")
	delete(allFields, "validValueRanges")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewFloatConfigurableParamValue() *FloatConfigurableParamValue {
	p := new(FloatConfigurableParamValue)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.FloatConfigurableParamValue"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /monitoring/v4.2/serviceability/alerts/{extId} Get operation
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

func (p *GetAlertApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetAlertApiResponse

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

func (p *GetAlertApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetAlertApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewGetAlertApiResponse()

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

func NewGetAlertApiResponse() *GetAlertApiResponse {
	p := new(GetAlertApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.GetAlertApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
REST response for all response codes in API path /monitoring/v4.2/serviceability/alerts/email-config Get operation
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

func (p *GetAlertEmailConfigurationApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetAlertEmailConfigurationApiResponse

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

func (p *GetAlertEmailConfigurationApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetAlertEmailConfigurationApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewGetAlertEmailConfigurationApiResponse()

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

func NewGetAlertEmailConfigurationApiResponse() *GetAlertEmailConfigurationApiResponse {
	p := new(GetAlertEmailConfigurationApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.GetAlertEmailConfigurationApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
REST response for all response codes in API path /monitoring/v4.2/serviceability/audits/{extId} Get operation
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

func (p *GetAuditApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetAuditApiResponse

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

func (p *GetAuditApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetAuditApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewGetAuditApiResponse()

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

func NewGetAuditApiResponse() *GetAuditApiResponse {
	p := new(GetAuditApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.GetAuditApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
REST response for all response codes in API path /monitoring/v4.2/serviceability/alerts/system-defined-policies/{systemDefinedPolicyExtId}/cluster-configs/{extId} Get operation
*/
type GetClusterConfigApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetClusterConfigApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *GetClusterConfigApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetClusterConfigApiResponse

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

func (p *GetClusterConfigApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetClusterConfigApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewGetClusterConfigApiResponse()

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

func NewGetClusterConfigApiResponse() *GetClusterConfigApiResponse {
	p := new(GetClusterConfigApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.GetClusterConfigApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetClusterConfigApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetClusterConfigApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetClusterConfigApiResponseData()
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
REST response for all response codes in API path /monitoring/v4.2/serviceability/events/{extId} Get operation
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

func (p *GetEventApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetEventApiResponse

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

func (p *GetEventApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetEventApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewGetEventApiResponse()

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

func NewGetEventApiResponse() *GetEventApiResponse {
	p := new(GetEventApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.GetEventApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
REST response for all response codes in API path /monitoring/v4.2/serviceability/alerts/system-defined-policies/{extId} Get operation
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

func (p *GetSdaPolicyApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetSdaPolicyApiResponse

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

func (p *GetSdaPolicyApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetSdaPolicyApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewGetSdaPolicyApiResponse()

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

func NewGetSdaPolicyApiResponse() *GetSdaPolicyApiResponse {
	p := new(GetSdaPolicyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.GetSdaPolicyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
REST response for all response codes in API path /monitoring/v4.2/serviceability/alerts/user-defined-policies/{extId} Get operation
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

func (p *GetUdaPolicyApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetUdaPolicyApiResponse

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

func (p *GetUdaPolicyApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetUdaPolicyApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewGetUdaPolicyApiResponse()

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

func NewGetUdaPolicyApiResponse() *GetUdaPolicyApiResponse {
	p := new(GetUdaPolicyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.GetUdaPolicyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
The allowed entity types that can be used to filter the set of entities on which User-Defined Alert policy is set.
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

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*GroupFilterProxy
		ExtId *string          `json:"extId,omitempty"`
		Type  *GroupEntityType `json:"type,omitempty"`
	}{
		GroupFilterProxy: (*GroupFilterProxy)(p),
		ExtId:            p.ExtId,
		Type:             p.Type,
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

func (p *GroupFilter) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GroupFilter
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewGroupFilter()

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
	if known.Type != nil {
		p.Type = known.Type
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "extId")
	delete(allFields, "type")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewGroupFilter() *GroupFilter {
	p := new(GroupFilter)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.GroupFilter"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
HTTP proxy used to establish the remote tunnel.
*/
type HttpProxy struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	AddressValue *import5.IPAddressOrFQDN `json:"addressValue,omitempty"`
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
	  User name for proxy authentication.
	*/
	Username *string `json:"username,omitempty"`
}

func (p *HttpProxy) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias HttpProxy

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

func (p *HttpProxy) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias HttpProxy
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewHttpProxy()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.AddressValue != nil {
		p.AddressValue = known.AddressValue
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.Password != nil {
		p.Password = known.Password
	}
	if known.Port != nil {
		p.Port = known.Port
	}
	if known.ProxyTypes != nil {
		p.ProxyTypes = known.ProxyTypes
	}
	if known.Username != nil {
		p.Username = known.Username
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "addressValue")
	delete(allFields, "name")
	delete(allFields, "password")
	delete(allFields, "port")
	delete(allFields, "proxyTypes")
	delete(allFields, "username")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewHttpProxy() *HttpProxy {
	p := new(HttpProxy)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.HttpProxy"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Captures values of the parameter when it’s an integer type.
*/
type IntConfigurableParamValue struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Captures the current value of the parameter.
	*/
	CurrentIntValue *int64 `json:"currentIntValue"`
	/*
	  Captures the default value of the parameter.
	*/
	DefaultIntValue *int64 `json:"defaultIntValue,omitempty"`
	/*
	  Valid values for the parameter.
	*/
	ValidValueRanges []ConfigurableParamValueRange `json:"validValueRanges,omitempty"`
}

func (p *IntConfigurableParamValue) MarshalJSON() ([]byte, error) {
	type IntConfigurableParamValueProxy IntConfigurableParamValue

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*IntConfigurableParamValueProxy
		CurrentIntValue *int64 `json:"currentIntValue,omitempty"`
	}{
		IntConfigurableParamValueProxy: (*IntConfigurableParamValueProxy)(p),
		CurrentIntValue:                p.CurrentIntValue,
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

func (p *IntConfigurableParamValue) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias IntConfigurableParamValue
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewIntConfigurableParamValue()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.CurrentIntValue != nil {
		p.CurrentIntValue = known.CurrentIntValue
	}
	if known.DefaultIntValue != nil {
		p.DefaultIntValue = known.DefaultIntValue
	}
	if known.ValidValueRanges != nil {
		p.ValidValueRanges = known.ValidValueRanges
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "currentIntValue")
	delete(allFields, "defaultIntValue")
	delete(allFields, "validValueRanges")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewIntConfigurableParamValue() *IntConfigurableParamValue {
	p := new(IntConfigurableParamValue)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.IntConfigurableParamValue"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /monitoring/v4.2/serviceability/alerts Get operation
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

func (p *ListAlertsApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ListAlertsApiResponse

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

func (p *ListAlertsApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ListAlertsApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewListAlertsApiResponse()

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

func NewListAlertsApiResponse() *ListAlertsApiResponse {
	p := new(ListAlertsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.ListAlertsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
REST response for all response codes in API path /monitoring/v4.2/serviceability/audits Get operation
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

func (p *ListAuditsApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ListAuditsApiResponse

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

func (p *ListAuditsApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ListAuditsApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewListAuditsApiResponse()

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

func NewListAuditsApiResponse() *ListAuditsApiResponse {
	p := new(ListAuditsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.ListAuditsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
REST response for all response codes in API path /monitoring/v4.2/serviceability/alerts/system-defined-policies/{systemDefinedPolicyExtId}/cluster-configs Get operation
*/
type ListClusterConfigsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListClusterConfigsApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ListClusterConfigsApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ListClusterConfigsApiResponse

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

func (p *ListClusterConfigsApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ListClusterConfigsApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewListClusterConfigsApiResponse()

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

func NewListClusterConfigsApiResponse() *ListClusterConfigsApiResponse {
	p := new(ListClusterConfigsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.ListClusterConfigsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListClusterConfigsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListClusterConfigsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListClusterConfigsApiResponseData()
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
REST response for all response codes in API path /monitoring/v4.2/serviceability/events Get operation
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

func (p *ListEventsApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ListEventsApiResponse

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

func (p *ListEventsApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ListEventsApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewListEventsApiResponse()

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

func NewListEventsApiResponse() *ListEventsApiResponse {
	p := new(ListEventsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.ListEventsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
REST response for all response codes in API path /monitoring/v4.2/serviceability/alerts/system-defined-policies Get operation
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

func (p *ListSdaPoliciesApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ListSdaPoliciesApiResponse

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

func (p *ListSdaPoliciesApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ListSdaPoliciesApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewListSdaPoliciesApiResponse()

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

func NewListSdaPoliciesApiResponse() *ListSdaPoliciesApiResponse {
	p := new(ListSdaPoliciesApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.ListSdaPoliciesApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
REST response for all response codes in API path /monitoring/v4.2/serviceability/clusters/{clusterExtId}/tags Get operation
*/
type ListTagsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListTagsApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ListTagsApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ListTagsApiResponse

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

func (p *ListTagsApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ListTagsApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewListTagsApiResponse()

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

func NewListTagsApiResponse() *ListTagsApiResponse {
	p := new(ListTagsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.ListTagsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListTagsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListTagsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListTagsApiResponseData()
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
REST response for all response codes in API path /monitoring/v4.2/serviceability/alerts/user-defined-policies Get operation
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

func (p *ListUdaPoliciesApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ListUdaPoliciesApiResponse

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

func (p *ListUdaPoliciesApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ListUdaPoliciesApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewListUdaPoliciesApiResponse()

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

func NewListUdaPoliciesApiResponse() *ListUdaPoliciesApiResponse {
	p := new(ListUdaPoliciesApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.ListUdaPoliciesApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
The parameters for uploading the archive if the destination is a local server.
*/
type LocalUploadParams struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Path where the archive will be stored. The path should pre-exist. Note the Janitor service will not clean up the data from these custom paths. It only monitors the default path.
	*/
	Path *string `json:"path,omitempty"`
}

func (p *LocalUploadParams) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias LocalUploadParams

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

func (p *LocalUploadParams) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias LocalUploadParams
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewLocalUploadParams()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Path != nil {
		p.Path = known.Path
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "path")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewLocalUploadParams() *LocalUploadParams {
	p := new(LocalUploadParams)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.LocalUploadParams"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	p.Path = new(string)
	*p.Path = "/home/nutanix/data/logbay/bundles"

	return p
}

/*
Input parameters for collecting logs.
*/
type LogCollectionSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	ArchiveOpts *ArchiveOpts `json:"archiveOpts"`
	/*
	  Description to be attached with the associated log collection task.
	*/
	Description *string `json:"description,omitempty"`
	/*
	  The time marking the end of a log collection window.
	It should be in the following format:
	YYYY-MM-DDTHH:MM:SSZ or YYYY-MM-DDTHH:MM:SS+HH:MM or YYYY-MM-DDTHH:MM:SS-HH:MM.
	The end time of log collection should be within the past three months.
	For example:
	- 1985-04-12T23:20:50.52Z represents 20 minutes and 50.52 seconds after the 23rd hour of
	  April 12th, 1985 in UTC.
	- 1996-12-19T16:39:57-08:00 represents 39 minutes and 57 seconds after the 16th hour
	  of December 19th, 1996 with an offset of -08:00 from UTC (Pacific Standard Time).
	  Note that this is equivalent to 1996-12-20T00:39:57Z in UTC.
	*/
	EndTime *time.Time `json:"endTime"`
	/*
	  Exclude items associated with the tag IDs. The tag ID can be fetched from the GET tags operation.
	*/
	ExcludeTags []string `json:"excludeTags,omitempty"`
	/*
	  Collect items associated with the tag IDs. The tag ID can be fetched from the GET tags operation.
	If no tags are specified, all items will be collected by default, which will take a longer time.
	It is recommended that at least one tag is selected.
	*/
	IncludeTags []string `json:"includeTags,omitempty"`
	/*
	  List of node IP addresses from where the logs will be collected.
	*/
	NodeIpList []import5.IPv4Address `json:"nodeIpList,omitempty"`
	/*
	  Indicates whether to mask sensitive data in the collected logs or not.
	*/
	ShouldAnonymize *bool `json:"shouldAnonymize,omitempty"`
	/*
	  Indicates whether to collect logs from a node where services are down or not. This feature is not supported on PC.
	*/
	ShouldCollectFromDisabledNode *bool `json:"shouldCollectFromDisabledNode,omitempty"`
	/*
	  The time marking the start of a log collection window.
	It should be in the following format:
	YYYY-MM-DDTHH:MM:SSZ or YYYY-MM-DDTHH:MM:SS+HH:MM or YYYY-MM-DDTHH:MM:SS-HH:MM.
	The start time of log collection should be within the past three months.
	For example:
	- 1985-04-12T23:20:50.52Z represents 20 minutes and 50.52 seconds after the 23rd hour of
	  April 12th, 1985 in UTC.
	- 1996-12-19T16:39:57-08:00 represents 39 minutes and 57 seconds after the 16th hour
	  of December 19th, 1996 with an offset of -08:00 from UTC (Pacific Standard Time).
	  Note that this is equivalent to 1996-12-20T00:39:57Z in UTC.
	*/
	StartTime *time.Time `json:"startTime"`

	TagOpts *TagOpts `json:"tagOpts,omitempty"`
}

func (p *LogCollectionSpec) MarshalJSON() ([]byte, error) {
	type LogCollectionSpecProxy LogCollectionSpec

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*LogCollectionSpecProxy
		ArchiveOpts *ArchiveOpts `json:"archiveOpts,omitempty"`
		EndTime     *time.Time   `json:"endTime,omitempty"`
		StartTime   *time.Time   `json:"startTime,omitempty"`
	}{
		LogCollectionSpecProxy: (*LogCollectionSpecProxy)(p),
		ArchiveOpts:            p.ArchiveOpts,
		EndTime:                p.EndTime,
		StartTime:              p.StartTime,
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

func (p *LogCollectionSpec) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias LogCollectionSpec
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewLogCollectionSpec()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ArchiveOpts != nil {
		p.ArchiveOpts = known.ArchiveOpts
	}
	if known.Description != nil {
		p.Description = known.Description
	}
	if known.EndTime != nil {
		p.EndTime = known.EndTime
	}
	if known.ExcludeTags != nil {
		p.ExcludeTags = known.ExcludeTags
	}
	if known.IncludeTags != nil {
		p.IncludeTags = known.IncludeTags
	}
	if known.NodeIpList != nil {
		p.NodeIpList = known.NodeIpList
	}
	if known.ShouldAnonymize != nil {
		p.ShouldAnonymize = known.ShouldAnonymize
	}
	if known.ShouldCollectFromDisabledNode != nil {
		p.ShouldCollectFromDisabledNode = known.ShouldCollectFromDisabledNode
	}
	if known.StartTime != nil {
		p.StartTime = known.StartTime
	}
	if known.TagOpts != nil {
		p.TagOpts = known.TagOpts
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "archiveOpts")
	delete(allFields, "description")
	delete(allFields, "endTime")
	delete(allFields, "excludeTags")
	delete(allFields, "includeTags")
	delete(allFields, "nodeIpList")
	delete(allFields, "shouldAnonymize")
	delete(allFields, "shouldCollectFromDisabledNode")
	delete(allFields, "startTime")
	delete(allFields, "tagOpts")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewLogCollectionSpec() *LogCollectionSpec {
	p := new(LogCollectionSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.LogCollectionSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	p.Description = new(string)
	*p.Description = "Log collection via V4 API"
	p.ShouldAnonymize = new(bool)
	*p.ShouldAnonymize = false
	p.ShouldCollectFromDisabledNode = new(bool)
	*p.ShouldCollectFromDisabledNode = false

	return p
}

/*
REST response for all response codes in API path /monitoring/v4.2/serviceability/alerts/{extId}/$actions/manage-alert Post operation
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

func (p *ManageAlertApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ManageAlertApiResponse

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

func (p *ManageAlertApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ManageAlertApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewManageAlertApiResponse()

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

func NewManageAlertApiResponse() *ManageAlertApiResponse {
	p := new(ManageAlertApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.ManageAlertApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
Additional arguments that are applicable only for MSP tag.
*/
type MspOpts struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of MSP cluster UUID from where logs would be collected.
	*/
	ClusterExtIds []string `json:"clusterExtIds,omitempty"`
}

func (p *MspOpts) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias MspOpts

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

func (p *MspOpts) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias MspOpts
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewMspOpts()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ClusterExtIds != nil {
		p.ClusterExtIds = known.ClusterExtIds
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "clusterExtIds")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewMspOpts() *MspOpts {
	p := new(MspOpts)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.MspOpts"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The parameters for uploading the archive if the destination is a Nutanix server.
*/
type NtnxServerUploadParams struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Nutanix support case number under which the archive will be stored.
	*/
	CaseNumber *int64 `json:"caseNumber"`

	Protocol *ServerUploadProtocol `json:"protocol,omitempty"`
}

func (p *NtnxServerUploadParams) MarshalJSON() ([]byte, error) {
	type NtnxServerUploadParamsProxy NtnxServerUploadParams

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*NtnxServerUploadParamsProxy
		CaseNumber *int64 `json:"caseNumber,omitempty"`
	}{
		NtnxServerUploadParamsProxy: (*NtnxServerUploadParamsProxy)(p),
		CaseNumber:                  p.CaseNumber,
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

func (p *NtnxServerUploadParams) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias NtnxServerUploadParams
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewNtnxServerUploadParams()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.CaseNumber != nil {
		p.CaseNumber = known.CaseNumber
	}
	if known.Protocol != nil {
		p.Protocol = known.Protocol
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "caseNumber")
	delete(allFields, "protocol")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewNtnxServerUploadParams() *NtnxServerUploadParams {
	p := new(NtnxServerUploadParams)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.NtnxServerUploadParams"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
	/*
	  Attributes.
	*/
	Attributes []import5.KVStringPair `json:"attributes,omitempty"`
	/*
	  Message.
	*/
	Message *string `json:"message,omitempty"`
}

func (p *ParameterizedMessage) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ParameterizedMessage

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

func (p *ParameterizedMessage) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ParameterizedMessage
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewParameterizedMessage()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Attributes != nil {
		p.Attributes = known.Attributes
	}
	if known.Message != nil {
		p.Message = known.Message
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "attributes")
	delete(allFields, "message")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewParameterizedMessage() *ParameterizedMessage {
	p := new(ParameterizedMessage)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.ParameterizedMessage"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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

/*
Indicates the state of the given property.
*/
type PropertyState int

const (
	PROPERTYSTATE_UNKNOWN       PropertyState = 0
	PROPERTYSTATE_REDACTED      PropertyState = 1
	PROPERTYSTATE_ENABLED       PropertyState = 2
	PROPERTYSTATE_DISABLED      PropertyState = 3
	PROPERTYSTATE_NOT_SUPPORTED PropertyState = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *PropertyState) name(index int) string {
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
func (e PropertyState) GetName() string {
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
func (e *PropertyState) index(name string) PropertyState {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ENABLED",
		"DISABLED",
		"NOT_SUPPORTED",
	}
	for idx := range names {
		if names[idx] == name {
			return PropertyState(idx)
		}
	}
	return PROPERTYSTATE_UNKNOWN
}

func (e *PropertyState) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for PropertyState:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *PropertyState) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e PropertyState) Ref() *PropertyState {
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
	  UUID of the entity the User-Defined Alert policy is associated with.
	*/
	EntityUuid *string `json:"entityUuid"`
	/*
	  Policy IDs associated with the specified entity.
	*/
	PolicyIds []string `json:"policyIds"`
}

func (p *RelatedPolicy) MarshalJSON() ([]byte, error) {
	type RelatedPolicyProxy RelatedPolicy

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*RelatedPolicyProxy
		EntityUuid *string  `json:"entityUuid,omitempty"`
		PolicyIds  []string `json:"policyIds,omitempty"`
	}{
		RelatedPolicyProxy: (*RelatedPolicyProxy)(p),
		EntityUuid:         p.EntityUuid,
		PolicyIds:          p.PolicyIds,
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

func (p *RelatedPolicy) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias RelatedPolicy
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewRelatedPolicy()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.EntityUuid != nil {
		p.EntityUuid = known.EntityUuid
	}
	if known.PolicyIds != nil {
		p.PolicyIds = known.PolicyIds
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "entityUuid")
	delete(allFields, "policyIds")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewRelatedPolicy() *RelatedPolicy {
	p := new(RelatedPolicy)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.RelatedPolicy"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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

func (p *RemoteTunnelDetails) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias RemoteTunnelDetails

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

func (p *RemoteTunnelDetails) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias RemoteTunnelDetails
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewRemoteTunnelDetails()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ConnectionStatus != nil {
		p.ConnectionStatus = known.ConnectionStatus
	}
	if known.HttpProxy != nil {
		p.HttpProxy = known.HttpProxy
	}
	if known.ServiceCenter != nil {
		p.ServiceCenter = known.ServiceCenter
	}
	if known.TransportStatus != nil {
		p.TransportStatus = known.TransportStatus
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "connectionStatus")
	delete(allFields, "httpProxy")
	delete(allFields, "serviceCenter")
	delete(allFields, "transportStatus")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewRemoteTunnelDetails() *RemoteTunnelDetails {
	p := new(RemoteTunnelDetails)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.RemoteTunnelDetails"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
	  Possible resolutions to troubleshoot this alert.
	*/
	Resolution *string `json:"resolution,omitempty"`
}

func (p *RootCauseAnalysis) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias RootCauseAnalysis

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

func (p *RootCauseAnalysis) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias RootCauseAnalysis
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewRootCauseAnalysis()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Cause != nil {
		p.Cause = known.Cause
	}
	if known.Detail != nil {
		p.Detail = known.Detail
	}
	if known.Resolution != nil {
		p.Resolution = known.Resolution
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "cause")
	delete(allFields, "detail")
	delete(allFields, "resolution")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewRootCauseAnalysis() *RootCauseAnalysis {
	p := new(RootCauseAnalysis)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.RootCauseAnalysis"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
	  Possible resolutions to troubleshoot this alert.
	*/
	Resolution *string `json:"resolution,omitempty"`
}

func (p *RootCauseAnalysisProjection) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias RootCauseAnalysisProjection

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

func (p *RootCauseAnalysisProjection) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias RootCauseAnalysisProjection
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewRootCauseAnalysisProjection()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Cause != nil {
		p.Cause = known.Cause
	}
	if known.Detail != nil {
		p.Detail = known.Detail
	}
	if known.Resolution != nil {
		p.Resolution = known.Resolution
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "cause")
	delete(allFields, "detail")
	delete(allFields, "resolution")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewRootCauseAnalysisProjection() *RootCauseAnalysisProjection {
	p := new(RootCauseAnalysisProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.RootCauseAnalysisProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /monitoring/v4.2/serviceability/clusters/{clusterExtId}/$actions/run-system-defined-checks Post operation
*/
type RunSystemDefinedChecksApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfRunSystemDefinedChecksApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *RunSystemDefinedChecksApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias RunSystemDefinedChecksApiResponse

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

func (p *RunSystemDefinedChecksApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias RunSystemDefinedChecksApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewRunSystemDefinedChecksApiResponse()

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

func NewRunSystemDefinedChecksApiResponse() *RunSystemDefinedChecksApiResponse {
	p := new(RunSystemDefinedChecksApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.RunSystemDefinedChecksApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *RunSystemDefinedChecksApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *RunSystemDefinedChecksApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfRunSystemDefinedChecksApiResponseData()
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
Required parameters for running the System-Defined Checks.
*/
type RunSystemDefinedChecksSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A list of additional email addresses for sending the run summary. Either this should be set or shouldSendReportToConfiguredRecipients should be true. If both are set then email would be sent to all the recipients.
	*/
	AdditionalRecipients []string `json:"additionalRecipients,omitempty"`
	/*
	  List of node IP addresses where the Check will run. This field will be ignored if the check scope is a cluster.
	*/
	NodeIps []import5.IPv4Address `json:"nodeIps,omitempty"`
	/*
	  List of Check IDs to be executed. This field cannot be set simultaneously with shouldRunAllChecks; only one of them should be specified.
	*/
	SdaExtIds []string `json:"sdaExtIds,omitempty"`
	/*
	  Indicates whether to mask sensitive data in the check run summary.
	*/
	ShouldAnonymize *bool `json:"shouldAnonymize,omitempty"`
	/*
	  Indicates whether all System-Defined Checks applicable to the specified cluster should be executed. This field is mutually exclusive with the sdaExtIds parameter, meaning that only one of these should be set at a time. Please use this field with caution, as it is resource-intensive.
	*/
	ShouldRunAllChecks *bool `json:"shouldRunAllChecks,omitempty"`
	/*
	  Determines if the run summary should be sent to the configured email address associated with the cluster. Either this should be true or additionalRecipients should be provided. If both are set then email would be sent to all the recipients.
	*/
	ShouldSendReportToConfiguredRecipients *bool `json:"shouldSendReportToConfiguredRecipients,omitempty"`
}

func (p *RunSystemDefinedChecksSpec) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias RunSystemDefinedChecksSpec

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

func (p *RunSystemDefinedChecksSpec) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias RunSystemDefinedChecksSpec
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewRunSystemDefinedChecksSpec()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.AdditionalRecipients != nil {
		p.AdditionalRecipients = known.AdditionalRecipients
	}
	if known.NodeIps != nil {
		p.NodeIps = known.NodeIps
	}
	if known.SdaExtIds != nil {
		p.SdaExtIds = known.SdaExtIds
	}
	if known.ShouldAnonymize != nil {
		p.ShouldAnonymize = known.ShouldAnonymize
	}
	if known.ShouldRunAllChecks != nil {
		p.ShouldRunAllChecks = known.ShouldRunAllChecks
	}
	if known.ShouldSendReportToConfiguredRecipients != nil {
		p.ShouldSendReportToConfiguredRecipients = known.ShouldSendReportToConfiguredRecipients
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "additionalRecipients")
	delete(allFields, "nodeIps")
	delete(allFields, "sdaExtIds")
	delete(allFields, "shouldAnonymize")
	delete(allFields, "shouldRunAllChecks")
	delete(allFields, "shouldSendReportToConfiguredRecipients")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewRunSystemDefinedChecksSpec() *RunSystemDefinedChecksSpec {
	p := new(RunSystemDefinedChecksSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.RunSystemDefinedChecksSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	p.ShouldAnonymize = new(bool)
	*p.ShouldAnonymize = true
	p.ShouldRunAllChecks = new(bool)
	*p.ShouldRunAllChecks = false
	p.ShouldSendReportToConfiguredRecipients = new(bool)
	*p.ShouldSendReportToConfiguredRecipients = true

	return p
}

/*
Scope for the policy execution.
*/
type Scope int

const (
	SCOPE_UNKNOWN  Scope = 0
	SCOPE_REDACTED Scope = 1
	SCOPE_CLUSTER  Scope = 2
	SCOPE_NODE     Scope = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *Scope) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CLUSTER",
		"NODE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e Scope) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CLUSTER",
		"NODE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *Scope) index(name string) Scope {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CLUSTER",
		"NODE",
	}
	for idx := range names {
		if names[idx] == name {
			return Scope(idx)
		}
	}
	return SCOPE_UNKNOWN
}

func (e *Scope) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for Scope:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *Scope) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e Scope) Ref() *Scope {
	return &e
}

/*
Specifies the sub-type for a SDA. It is only applicable to policies with type as HEALTH_CHECK.
*/
type SdaSubType int

const (
	SDASUBTYPE_UNKNOWN       SdaSubType = 0
	SDASUBTYPE_REDACTED      SdaSubType = 1
	SDASUBTYPE_SCHEDULED     SdaSubType = 2
	SDASUBTYPE_NOT_SCHEDULED SdaSubType = 3
	SDASUBTYPE_EVENT_DRIVEN  SdaSubType = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *SdaSubType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SCHEDULED",
		"NOT_SCHEDULED",
		"EVENT_DRIVEN",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e SdaSubType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SCHEDULED",
		"NOT_SCHEDULED",
		"EVENT_DRIVEN",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *SdaSubType) index(name string) SdaSubType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SCHEDULED",
		"NOT_SCHEDULED",
		"EVENT_DRIVEN",
	}
	for idx := range names {
		if names[idx] == name {
			return SdaSubType(idx)
		}
	}
	return SDASUBTYPE_UNKNOWN
}

func (e *SdaSubType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for SdaSubType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *SdaSubType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e SdaSubType) Ref() *SdaSubType {
	return &e
}

/*
Defines the type of the System-Defined Alert Policy.
*/
type SdaType int

const (
	SDATYPE_UNKNOWN         SdaType = 0
	SDATYPE_REDACTED        SdaType = 1
	SDATYPE_HEALTH_CHECK    SdaType = 2
	SDATYPE_SERVICE_DEFINED SdaType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *SdaType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"HEALTH_CHECK",
		"SERVICE_DEFINED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e SdaType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"HEALTH_CHECK",
		"SERVICE_DEFINED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *SdaType) index(name string) SdaType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"HEALTH_CHECK",
		"SERVICE_DEFINED",
	}
	for idx := range names {
		if names[idx] == name {
			return SdaType(idx)
		}
	}
	return SDATYPE_UNKNOWN
}

func (e *SdaType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for SdaType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *SdaType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e SdaType) Ref() *SdaType {
	return &e
}

/*
Protocol used for uploading to the target destination.
*/
type ServerUploadProtocol int

const (
	SERVERUPLOADPROTOCOL_UNKNOWN  ServerUploadProtocol = 0
	SERVERUPLOADPROTOCOL_REDACTED ServerUploadProtocol = 1
	SERVERUPLOADPROTOCOL_SFTP     ServerUploadProtocol = 2
	SERVERUPLOADPROTOCOL_FTP      ServerUploadProtocol = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ServerUploadProtocol) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SFTP",
		"FTP",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ServerUploadProtocol) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SFTP",
		"FTP",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ServerUploadProtocol) index(name string) ServerUploadProtocol {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SFTP",
		"FTP",
	}
	for idx := range names {
		if names[idx] == name {
			return ServerUploadProtocol(idx)
		}
	}
	return SERVERUPLOADPROTOCOL_UNKNOWN
}

func (e *ServerUploadProtocol) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ServerUploadProtocol:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ServerUploadProtocol) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ServerUploadProtocol) Ref() *ServerUploadProtocol {
	return &e
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

func (p *ServiceCenter) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ServiceCenter

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

func (p *ServiceCenter) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ServiceCenter
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewServiceCenter()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.IpAddress != nil {
		p.IpAddress = known.IpAddress
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.Port != nil {
		p.Port = known.Port
	}
	if known.Username != nil {
		p.Username = known.Username
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "ipAddress")
	delete(allFields, "name")
	delete(allFields, "port")
	delete(allFields, "username")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewServiceCenter() *ServiceCenter {
	p := new(ServiceCenter)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.ServiceCenter"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Captures details of each of the severity types and their associated configurable parameters if applicable.
*/
type SeverityConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	State *PropertyState `json:"state,omitempty"`
	/*
	  Captures alert-related thresholds that correspond to a particular severity.
	*/
	ThresholdParameters []AlertPolicyConfigurableParameter `json:"thresholdParameters,omitempty"`
}

func (p *SeverityConfig) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias SeverityConfig

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

func (p *SeverityConfig) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias SeverityConfig
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewSeverityConfig()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.State != nil {
		p.State = known.State
	}
	if known.ThresholdParameters != nil {
		p.ThresholdParameters = known.ThresholdParameters
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "state")
	delete(allFields, "thresholdParameters")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewSeverityConfig() *SeverityConfig {
	p := new(SeverityConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.SeverityConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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

func (p *SeverityTrail) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias SeverityTrail

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

func (p *SeverityTrail) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias SeverityTrail
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewSeverityTrail()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Severity != nil {
		p.Severity = known.Severity
	}
	if known.SeverityChangeTime != nil {
		p.SeverityChangeTime = known.SeverityChangeTime
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "severity")
	delete(allFields, "severityChangeTime")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewSeverityTrail() *SeverityTrail {
	p := new(SeverityTrail)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.SeverityTrail"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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

/*
The parameters for uploading the archive if the destination is a storage container.
The local host storage container uploads are not supported on PC.
*/
type StorageContainerUploadParams struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Credential *Credential `json:"credential"`

	IpAddress *import5.IPv4Address `json:"ipAddress,omitempty"`
	/*
	  Path of storage container where the archive will be uploaded.
	*/
	Path *string `json:"path"`
	/*
	  Port of the target. Mandatory for SFTP and FTP protocol.
	*/
	Port *int `json:"port,omitempty"`
}

func (p *StorageContainerUploadParams) MarshalJSON() ([]byte, error) {
	type StorageContainerUploadParamsProxy StorageContainerUploadParams

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*StorageContainerUploadParamsProxy
		Credential *Credential `json:"credential,omitempty"`
		Path       *string     `json:"path,omitempty"`
	}{
		StorageContainerUploadParamsProxy: (*StorageContainerUploadParamsProxy)(p),
		Credential:                        p.Credential,
		Path:                              p.Path,
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

func (p *StorageContainerUploadParams) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias StorageContainerUploadParams
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewStorageContainerUploadParams()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Credential != nil {
		p.Credential = known.Credential
	}
	if known.IpAddress != nil {
		p.IpAddress = known.IpAddress
	}
	if known.Path != nil {
		p.Path = known.Path
	}
	if known.Port != nil {
		p.Port = known.Port
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "credential")
	delete(allFields, "ipAddress")
	delete(allFields, "path")
	delete(allFields, "port")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewStorageContainerUploadParams() *StorageContainerUploadParams {
	p := new(StorageContainerUploadParams)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.StorageContainerUploadParams"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	p.Port = new(int)
	*p.Port = 2222

	return p
}

/*
Captures values of the parameter when it's of type string.
*/
type StringConfigurableParamValue struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Captures the current value of the parameter.
	*/
	CurrentStrValue *string `json:"currentStrValue"`
	/*
	  Captures the default value of the parameter.
	*/
	DefaultStrValue *string `json:"defaultStrValue,omitempty"`
}

func (p *StringConfigurableParamValue) MarshalJSON() ([]byte, error) {
	type StringConfigurableParamValueProxy StringConfigurableParamValue

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*StringConfigurableParamValueProxy
		CurrentStrValue *string `json:"currentStrValue,omitempty"`
	}{
		StringConfigurableParamValueProxy: (*StringConfigurableParamValueProxy)(p),
		CurrentStrValue:                   p.CurrentStrValue,
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

func (p *StringConfigurableParamValue) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias StringConfigurableParamValue
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewStringConfigurableParamValue()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.CurrentStrValue != nil {
		p.CurrentStrValue = known.CurrentStrValue
	}
	if known.DefaultStrValue != nil {
		p.DefaultStrValue = known.DefaultStrValue
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "currentStrValue")
	delete(allFields, "defaultStrValue")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewStringConfigurableParamValue() *StringConfigurableParamValue {
	p := new(StringConfigurableParamValue)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.StringConfigurableParamValue"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Attributes of System-Defined Alert Policy across all clusters.
*/
type SystemDefinedPolicy struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Various categories into which this alert type can be classified. For example, hardware, storage, or license.
	*/
	Classifications []string `json:"classifications,omitempty"`
	/*
	  SDA policy parameters that may differ across clusters since each cluster can run on different NCC versions. Each cluster will have an individual entry.
	*/
	ClusterConfigs []ClusterConfig `json:"clusterConfigs,omitempty"`
	/*
	  System-defined alert policy description.
	*/
	Description *string `json:"description,omitempty"`

	EntityType *EntityType `json:"entityType,omitempty"`
	/*
	  Unique ID of the System-Defined Alert Policy.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Impact types to which this rule applies.
	*/
	ImpactTypes []import1.ImpactType `json:"impactTypes,omitempty"`
	/*
	  List of knowledge base article links.
	*/
	KbArticles []string `json:"kbArticles,omitempty"`
	/*
	  A HATEOAS style link for the response. Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Name of the System-Defined Alert Policy.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  Unique ID associated with the policy.
	*/
	PolicyId *string `json:"policyId,omitempty"`
	/*
	  Publisher of the policy. For example, NCC for all health check policies.
	*/
	Publisher *string `json:"publisher,omitempty"`

	Scope *Scope `json:"scope,omitempty"`

	SubType *SdaSubType `json:"subType,omitempty"`
	/*
	  Indicates the cluster type against which this policy can be executed.
	*/
	TargetClusters []ClusterType `json:"targetClusters,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  Title of a System-Defined Alert Policy.
	*/
	Title *string `json:"title,omitempty"`

	Type *SdaType `json:"type,omitempty"`
}

func (p *SystemDefinedPolicy) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias SystemDefinedPolicy

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

func (p *SystemDefinedPolicy) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias SystemDefinedPolicy
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewSystemDefinedPolicy()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Classifications != nil {
		p.Classifications = known.Classifications
	}
	if known.ClusterConfigs != nil {
		p.ClusterConfigs = known.ClusterConfigs
	}
	if known.Description != nil {
		p.Description = known.Description
	}
	if known.EntityType != nil {
		p.EntityType = known.EntityType
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.ImpactTypes != nil {
		p.ImpactTypes = known.ImpactTypes
	}
	if known.KbArticles != nil {
		p.KbArticles = known.KbArticles
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.PolicyId != nil {
		p.PolicyId = known.PolicyId
	}
	if known.Publisher != nil {
		p.Publisher = known.Publisher
	}
	if known.Scope != nil {
		p.Scope = known.Scope
	}
	if known.SubType != nil {
		p.SubType = known.SubType
	}
	if known.TargetClusters != nil {
		p.TargetClusters = known.TargetClusters
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}
	if known.Title != nil {
		p.Title = known.Title
	}
	if known.Type != nil {
		p.Type = known.Type
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "classifications")
	delete(allFields, "clusterConfigs")
	delete(allFields, "description")
	delete(allFields, "entityType")
	delete(allFields, "extId")
	delete(allFields, "impactTypes")
	delete(allFields, "kbArticles")
	delete(allFields, "links")
	delete(allFields, "name")
	delete(allFields, "policyId")
	delete(allFields, "publisher")
	delete(allFields, "scope")
	delete(allFields, "subType")
	delete(allFields, "targetClusters")
	delete(allFields, "tenantId")
	delete(allFields, "title")
	delete(allFields, "type")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewSystemDefinedPolicy() *SystemDefinedPolicy {
	p := new(SystemDefinedPolicy)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.SystemDefinedPolicy"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	p.Publisher = new(string)
	*p.Publisher = "NCC"

	return p
}

/*
Logbay tag.
*/
type Tag struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Logbay tag description.
	*/
	Description *string `json:"description,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Logbay tag name.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *Tag) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias Tag

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

func (p *Tag) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Tag
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewTag()

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
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "description")
	delete(allFields, "extId")
	delete(allFields, "links")
	delete(allFields, "name")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewTag() *Tag {
	p := new(Tag)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.Tag"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
This parameter allows the user to pass additional arguments that are specific to a particular tag.
*/
type TagOpts struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	MspOpts *MspOpts `json:"mspOpts,omitempty"`
}

func (p *TagOpts) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias TagOpts

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

func (p *TagOpts) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias TagOpts
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewTagOpts()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.MspOpts != nil {
		p.MspOpts = known.MspOpts
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "mspOpts")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewTagOpts() *TagOpts {
	p := new(TagOpts)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.TagOpts"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
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

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*TriggerConditionProxy
		Condition     *Condition             `json:"condition,omitempty"`
		ConditionType *import1.ConditionType `json:"conditionType,omitempty"`
		SeverityLevel *PolicySeverityLevel   `json:"severityLevel,omitempty"`
	}{
		TriggerConditionProxy: (*TriggerConditionProxy)(p),
		Condition:             p.Condition,
		ConditionType:         p.ConditionType,
		SeverityLevel:         p.SeverityLevel,
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

func (p *TriggerCondition) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias TriggerCondition
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewTriggerCondition()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Condition != nil {
		p.Condition = known.Condition
	}
	if known.ConditionType != nil {
		p.ConditionType = known.ConditionType
	}
	if known.SeverityLevel != nil {
		p.SeverityLevel = known.SeverityLevel
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "condition")
	delete(allFields, "conditionType")
	delete(allFields, "severityLevel")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewTriggerCondition() *TriggerCondition {
	p := new(TriggerCondition)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.TriggerCondition"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /monitoring/v4.2/serviceability/alerts/email-config Put operation
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

func (p *UpdateAlertEmailConfigurationApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias UpdateAlertEmailConfigurationApiResponse

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

func (p *UpdateAlertEmailConfigurationApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias UpdateAlertEmailConfigurationApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewUpdateAlertEmailConfigurationApiResponse()

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

func NewUpdateAlertEmailConfigurationApiResponse() *UpdateAlertEmailConfigurationApiResponse {
	p := new(UpdateAlertEmailConfigurationApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.UpdateAlertEmailConfigurationApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
REST response for all response codes in API path /monitoring/v4.2/serviceability/alerts/system-defined-policies/{systemDefinedPolicyExtId}/cluster-configs/{extId} Put operation
*/
type UpdateClusterConfigApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateClusterConfigApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *UpdateClusterConfigApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias UpdateClusterConfigApiResponse

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

func (p *UpdateClusterConfigApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias UpdateClusterConfigApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewUpdateClusterConfigApiResponse()

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

func NewUpdateClusterConfigApiResponse() *UpdateClusterConfigApiResponse {
	p := new(UpdateClusterConfigApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.UpdateClusterConfigApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdateClusterConfigApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdateClusterConfigApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdateClusterConfigApiResponseData()
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
REST response for all response codes in API path /monitoring/v4.2/serviceability/alerts/user-defined-policies/{extId} Put operation
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

func (p *UpdateUdaPolicyApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias UpdateUdaPolicyApiResponse

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

func (p *UpdateUdaPolicyApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias UpdateUdaPolicyApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewUpdateUdaPolicyApiResponse()

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

func NewUpdateUdaPolicyApiResponse() *UpdateUdaPolicyApiResponse {
	p := new(UpdateUdaPolicyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.UpdateUdaPolicyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
	  Filter criteria for narrowing down the entities on which User-Defined Alert policies can be set up.
	*/
	Filters *OneOfUserDefinedPolicyFilters `json:"filters,omitempty"`
	/*
	  Impact types for the associated resulting alert.
	*/
	ImpactTypes []import1.ImpactType `json:"impactTypes,omitempty"`
	/*
	  Indicates whether the auto-resolve feature is enabled for this policy.
	*/
	IsAutoResolved *bool `json:"isAutoResolved,omitempty"`
	/*
	  Enable/Disable flag for the policy.
	*/
	IsEnabled *bool `json:"isEnabled,omitempty"`
	/*
	  Error when conflicting alert policies are found.
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
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
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

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*UserDefinedPolicyProxy
		EntityType        *string            `json:"entityType,omitempty"`
		Title             *string            `json:"title,omitempty"`
		TriggerConditions []TriggerCondition `json:"triggerConditions,omitempty"`
	}{
		UserDefinedPolicyProxy: (*UserDefinedPolicyProxy)(p),
		EntityType:             p.EntityType,
		Title:                  p.Title,
		TriggerConditions:      p.TriggerConditions,
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

func (p *UserDefinedPolicy) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias UserDefinedPolicy
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewUserDefinedPolicy()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.CreatedBy != nil {
		p.CreatedBy = known.CreatedBy
	}
	if known.Description != nil {
		p.Description = known.Description
	}
	if known.EntityType != nil {
		p.EntityType = known.EntityType
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.FiltersItemDiscriminator_ != nil {
		p.FiltersItemDiscriminator_ = known.FiltersItemDiscriminator_
	}
	if known.Filters != nil {
		p.Filters = known.Filters
	}
	if known.ImpactTypes != nil {
		p.ImpactTypes = known.ImpactTypes
	}
	if known.IsAutoResolved != nil {
		p.IsAutoResolved = known.IsAutoResolved
	}
	if known.IsEnabled != nil {
		p.IsEnabled = known.IsEnabled
	}
	if known.IsExpectedToErrorOnConflict != nil {
		p.IsExpectedToErrorOnConflict = known.IsExpectedToErrorOnConflict
	}
	if known.LastUpdatedTime != nil {
		p.LastUpdatedTime = known.LastUpdatedTime
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.PoliciesToOverride != nil {
		p.PoliciesToOverride = known.PoliciesToOverride
	}
	if known.RelatedPolicies != nil {
		p.RelatedPolicies = known.RelatedPolicies
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}
	if known.Title != nil {
		p.Title = known.Title
	}
	if known.TriggerConditions != nil {
		p.TriggerConditions = known.TriggerConditions
	}
	if known.TriggerWaitPeriod != nil {
		p.TriggerWaitPeriod = known.TriggerWaitPeriod
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "createdBy")
	delete(allFields, "description")
	delete(allFields, "entityType")
	delete(allFields, "extId")
	delete(allFields, "$filtersItemDiscriminator")
	delete(allFields, "filters")
	delete(allFields, "impactTypes")
	delete(allFields, "isAutoResolved")
	delete(allFields, "isEnabled")
	delete(allFields, "isExpectedToErrorOnConflict")
	delete(allFields, "lastUpdatedTime")
	delete(allFields, "links")
	delete(allFields, "policiesToOverride")
	delete(allFields, "relatedPolicies")
	delete(allFields, "tenantId")
	delete(allFields, "title")
	delete(allFields, "triggerConditions")
	delete(allFields, "triggerWaitPeriod")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewUserDefinedPolicy() *UserDefinedPolicy {
	p := new(UserDefinedPolicy)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.UserDefinedPolicy"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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

func (p *UserReference) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias UserReference

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

func (p *UserReference) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias UserReference
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewUserReference()

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
	if known.IpAddress != nil {
		p.IpAddress = known.IpAddress
	}
	if known.Name != nil {
		p.Name = known.Name
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "extId")
	delete(allFields, "ipAddress")
	delete(allFields, "name")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewUserReference() *UserReference {
	p := new(UserReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "monitoring.v4.serviceability.UserReference"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type OneOfGetEventApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType0    *Event                 `json:"-"`
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
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfGetEventApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfGetEventApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "monitoring.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetEventApiResponseData"))
}

func (p *OneOfGetEventApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfGetEventApiResponseData")
}

type OneOfGetAlertEmailConfigurationApiResponseData struct {
	Discriminator *string                  `json:"-"`
	ObjectType_   *string                  `json:"-"`
	oneOfType400  *import4.ErrorResponse   `json:"-"`
	oneOfType0    *AlertEmailConfiguration `json:"-"`
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
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfGetAlertEmailConfigurationApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfGetAlertEmailConfigurationApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "monitoring.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetAlertEmailConfigurationApiResponseData"))
}

func (p *OneOfGetAlertEmailConfigurationApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfGetAlertEmailConfigurationApiResponseData")
}

type OneOfListEventsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType0    []Event                `json:"-"`
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
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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
		*p.Discriminator = "List<monitoring.v4.serviceability.Event>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<monitoring.v4.serviceability.Event>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListEventsApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<monitoring.v4.serviceability.Event>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfListEventsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "monitoring.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListEventsApiResponseData"))
}

func (p *OneOfListEventsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<monitoring.v4.serviceability.Event>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfListEventsApiResponseData")
}

type OneOfListSdaPoliciesApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
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
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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
		*p.Discriminator = "List<monitoring.v4.serviceability.SystemDefinedPolicy>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<monitoring.v4.serviceability.SystemDefinedPolicy>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListSdaPoliciesApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<monitoring.v4.serviceability.SystemDefinedPolicy>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfListSdaPoliciesApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "monitoring.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
		if len(*vOneOfType0) == 0 || "monitoring.v4.serviceability.SystemDefinedPolicy" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<monitoring.v4.serviceability.SystemDefinedPolicy>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<monitoring.v4.serviceability.SystemDefinedPolicy>"
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListSdaPoliciesApiResponseData"))
}

func (p *OneOfListSdaPoliciesApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<monitoring.v4.serviceability.SystemDefinedPolicy>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfListSdaPoliciesApiResponseData")
}

type OneOfListUdaPoliciesApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
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
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "monitoring.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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

type OneOfAlertPolicyConfigurableParameterParamValue struct {
	Discriminator *string                        `json:"-"`
	ObjectType_   *string                        `json:"-"`
	oneOfType0    *IntConfigurableParamValue     `json:"-"`
	oneOfType2    *BooleanConfigurableParamValue `json:"-"`
	oneOfType3    *StringConfigurableParamValue  `json:"-"`
	oneOfType1    *FloatConfigurableParamValue   `json:"-"`
}

func NewOneOfAlertPolicyConfigurableParameterParamValue() *OneOfAlertPolicyConfigurableParameterParamValue {
	p := new(OneOfAlertPolicyConfigurableParameterParamValue)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfAlertPolicyConfigurableParameterParamValue) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfAlertPolicyConfigurableParameterParamValue is nil"))
	}
	switch v.(type) {
	case IntConfigurableParamValue:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(IntConfigurableParamValue)
		}
		*p.oneOfType0 = v.(IntConfigurableParamValue)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case BooleanConfigurableParamValue:
		if nil == p.oneOfType2 {
			p.oneOfType2 = new(BooleanConfigurableParamValue)
		}
		*p.oneOfType2 = v.(BooleanConfigurableParamValue)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2.ObjectType_
	case StringConfigurableParamValue:
		if nil == p.oneOfType3 {
			p.oneOfType3 = new(StringConfigurableParamValue)
		}
		*p.oneOfType3 = v.(StringConfigurableParamValue)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType3.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType3.ObjectType_
	case FloatConfigurableParamValue:
		if nil == p.oneOfType1 {
			p.oneOfType1 = new(FloatConfigurableParamValue)
		}
		*p.oneOfType1 = v.(FloatConfigurableParamValue)
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

func (p *OneOfAlertPolicyConfigurableParameterParamValue) GetValue() interface{} {
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

func (p *OneOfAlertPolicyConfigurableParameterParamValue) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(IntConfigurableParamValue)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "monitoring.v4.serviceability.IntConfigurableParamValue" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(IntConfigurableParamValue)
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
	vOneOfType2 := new(BooleanConfigurableParamValue)
	if err := json.Unmarshal(b, vOneOfType2); err == nil {
		if "monitoring.v4.serviceability.BooleanConfigurableParamValue" == *vOneOfType2.ObjectType_ {
			if nil == p.oneOfType2 {
				p.oneOfType2 = new(BooleanConfigurableParamValue)
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
	vOneOfType3 := new(StringConfigurableParamValue)
	if err := json.Unmarshal(b, vOneOfType3); err == nil {
		if "monitoring.v4.serviceability.StringConfigurableParamValue" == *vOneOfType3.ObjectType_ {
			if nil == p.oneOfType3 {
				p.oneOfType3 = new(StringConfigurableParamValue)
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
	vOneOfType1 := new(FloatConfigurableParamValue)
	if err := json.Unmarshal(b, vOneOfType1); err == nil {
		if "monitoring.v4.serviceability.FloatConfigurableParamValue" == *vOneOfType1.ObjectType_ {
			if nil == p.oneOfType1 {
				p.oneOfType1 = new(FloatConfigurableParamValue)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfAlertPolicyConfigurableParameterParamValue"))
}

func (p *OneOfAlertPolicyConfigurableParameterParamValue) MarshalJSON() ([]byte, error) {
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
	return nil, errors.New("No value to marshal for OneOfAlertPolicyConfigurableParameterParamValue")
}

type OneOfRunSystemDefinedChecksApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import3.TaskReference `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfRunSystemDefinedChecksApiResponseData() *OneOfRunSystemDefinedChecksApiResponseData {
	p := new(OneOfRunSystemDefinedChecksApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfRunSystemDefinedChecksApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfRunSystemDefinedChecksApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import3.TaskReference)
		}
		*p.oneOfType2001 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfRunSystemDefinedChecksApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfRunSystemDefinedChecksApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import3.TaskReference)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "monitoring.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfRunSystemDefinedChecksApiResponseData"))
}

func (p *OneOfRunSystemDefinedChecksApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfRunSystemDefinedChecksApiResponseData")
}

type OneOfGetAlertApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *Alert                 `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
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
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "monitoring.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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

type OneOfGetSdaPolicyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType0    *SystemDefinedPolicy   `json:"-"`
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
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfGetSdaPolicyApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfGetSdaPolicyApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "monitoring.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	vOneOfType0 := new(SystemDefinedPolicy)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "monitoring.v4.serviceability.SystemDefinedPolicy" == *vOneOfType0.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetSdaPolicyApiResponseData"))
}

func (p *OneOfGetSdaPolicyApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfGetSdaPolicyApiResponseData")
}

type OneOfGetClusterConfigApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType0    *ClusterConfig         `json:"-"`
}

func NewOneOfGetClusterConfigApiResponseData() *OneOfGetClusterConfigApiResponseData {
	p := new(OneOfGetClusterConfigApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetClusterConfigApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetClusterConfigApiResponseData is nil"))
	}
	switch v.(type) {
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case ClusterConfig:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(ClusterConfig)
		}
		*p.oneOfType0 = v.(ClusterConfig)
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

func (p *OneOfGetClusterConfigApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfGetClusterConfigApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "monitoring.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	vOneOfType0 := new(ClusterConfig)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "monitoring.v4.serviceability.ClusterConfig" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(ClusterConfig)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetClusterConfigApiResponseData"))
}

func (p *OneOfGetClusterConfigApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfGetClusterConfigApiResponseData")
}

type OneOfListClusterConfigsApiResponseData struct {
	Discriminator *string                   `json:"-"`
	ObjectType_   *string                   `json:"-"`
	oneOfType400  *import4.ErrorResponse    `json:"-"`
	oneOfType0    []ClusterConfig           `json:"-"`
	oneOfType401  []ClusterConfigProjection `json:"-"`
}

func NewOneOfListClusterConfigsApiResponseData() *OneOfListClusterConfigsApiResponseData {
	p := new(OneOfListClusterConfigsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListClusterConfigsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListClusterConfigsApiResponseData is nil"))
	}
	switch v.(type) {
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case []ClusterConfig:
		p.oneOfType0 = v.([]ClusterConfig)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<monitoring.v4.serviceability.ClusterConfig>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<monitoring.v4.serviceability.ClusterConfig>"
	case []ClusterConfigProjection:
		p.oneOfType401 = v.([]ClusterConfigProjection)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<monitoring.v4.serviceability.ClusterConfigProjection>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<monitoring.v4.serviceability.ClusterConfigProjection>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListClusterConfigsApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<monitoring.v4.serviceability.ClusterConfig>" == *p.Discriminator {
		return p.oneOfType0
	}
	if "List<monitoring.v4.serviceability.ClusterConfigProjection>" == *p.Discriminator {
		return p.oneOfType401
	}
	return nil
}

func (p *OneOfListClusterConfigsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "monitoring.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	vOneOfType0 := new([]ClusterConfig)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "monitoring.v4.serviceability.ClusterConfig" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<monitoring.v4.serviceability.ClusterConfig>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<monitoring.v4.serviceability.ClusterConfig>"
			return nil
		}
	}
	vOneOfType401 := new([]ClusterConfigProjection)
	if err := json.Unmarshal(b, vOneOfType401); err == nil {
		if len(*vOneOfType401) == 0 || "monitoring.v4.serviceability.ClusterConfigProjection" == *((*vOneOfType401)[0].ObjectType_) {
			p.oneOfType401 = *vOneOfType401
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<monitoring.v4.serviceability.ClusterConfigProjection>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<monitoring.v4.serviceability.ClusterConfigProjection>"
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListClusterConfigsApiResponseData"))
}

func (p *OneOfListClusterConfigsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<monitoring.v4.serviceability.ClusterConfig>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if "List<monitoring.v4.serviceability.ClusterConfigProjection>" == *p.Discriminator {
		return json.Marshal(p.oneOfType401)
	}
	return nil, errors.New("No value to marshal for OneOfListClusterConfigsApiResponseData")
}

type OneOfConfigurableParamValueRangeMaximumValue struct {
	Discriminator *string  `json:"-"`
	ObjectType_   *string  `json:"-"`
	oneOfType0    *int64   `json:"-"`
	oneOfType1    *float32 `json:"-"`
}

func NewOneOfConfigurableParamValueRangeMaximumValue() *OneOfConfigurableParamValueRangeMaximumValue {
	p := new(OneOfConfigurableParamValueRangeMaximumValue)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfConfigurableParamValueRangeMaximumValue) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfConfigurableParamValueRangeMaximumValue is nil"))
	}
	switch v.(type) {
	case int64:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(int64)
		}
		*p.oneOfType0 = v.(int64)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "Long"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "Long"
	case float32:
		if nil == p.oneOfType1 {
			p.oneOfType1 = new(float32)
		}
		*p.oneOfType1 = v.(float32)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "Float"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "Float"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfConfigurableParamValueRangeMaximumValue) GetValue() interface{} {
	if "Long" == *p.Discriminator {
		return *p.oneOfType0
	}
	if "Float" == *p.Discriminator {
		return *p.oneOfType1
	}
	return nil
}

func (p *OneOfConfigurableParamValueRangeMaximumValue) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(int64)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(int64)
		}
		*p.oneOfType0 = *vOneOfType0
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "Long"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "Long"
		return nil
	}
	vOneOfType1 := new(float32)
	if err := json.Unmarshal(b, vOneOfType1); err == nil {
		if nil == p.oneOfType1 {
			p.oneOfType1 = new(float32)
		}
		*p.oneOfType1 = *vOneOfType1
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "Float"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "Float"
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfConfigurableParamValueRangeMaximumValue"))
}

func (p *OneOfConfigurableParamValueRangeMaximumValue) MarshalJSON() ([]byte, error) {
	if "Long" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if "Float" == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	return nil, errors.New("No value to marshal for OneOfConfigurableParamValueRangeMaximumValue")
}

type OneOfConditionThresholdValue struct {
	Discriminator *string              `json:"-"`
	ObjectType_   *string              `json:"-"`
	oneOfType0    *import1.DoubleValue `json:"-"`
	oneOfType1    *import1.IntValue    `json:"-"`
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfConditionThresholdValue) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1
	}
	return nil
}

func (p *OneOfConditionThresholdValue) UnmarshalJSON(b []byte) error {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfConditionThresholdValue"))
}

func (p *OneOfConditionThresholdValue) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	return nil, errors.New("No value to marshal for OneOfConditionThresholdValue")
}

type OneOfManageAlertApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
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
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "monitoring.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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

type OneOfArchiveOptsUploadParams struct {
	Discriminator *string                       `json:"-"`
	ObjectType_   *string                       `json:"-"`
	oneOfType2003 *CustomServerUploadParams     `json:"-"`
	oneOfType2004 *StorageContainerUploadParams `json:"-"`
	oneOfType2001 *LocalUploadParams            `json:"-"`
	oneOfType2002 *NtnxServerUploadParams       `json:"-"`
}

func NewOneOfArchiveOptsUploadParams() *OneOfArchiveOptsUploadParams {
	p := new(OneOfArchiveOptsUploadParams)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfArchiveOptsUploadParams) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfArchiveOptsUploadParams is nil"))
	}
	switch v.(type) {
	case CustomServerUploadParams:
		if nil == p.oneOfType2003 {
			p.oneOfType2003 = new(CustomServerUploadParams)
		}
		*p.oneOfType2003 = v.(CustomServerUploadParams)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2003.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2003.ObjectType_
	case StorageContainerUploadParams:
		if nil == p.oneOfType2004 {
			p.oneOfType2004 = new(StorageContainerUploadParams)
		}
		*p.oneOfType2004 = v.(StorageContainerUploadParams)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2004.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2004.ObjectType_
	case LocalUploadParams:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(LocalUploadParams)
		}
		*p.oneOfType2001 = v.(LocalUploadParams)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case NtnxServerUploadParams:
		if nil == p.oneOfType2002 {
			p.oneOfType2002 = new(NtnxServerUploadParams)
		}
		*p.oneOfType2002 = v.(NtnxServerUploadParams)
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

func (p *OneOfArchiveOptsUploadParams) GetValue() interface{} {
	if p.oneOfType2003 != nil && *p.oneOfType2003.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2003
	}
	if p.oneOfType2004 != nil && *p.oneOfType2004.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2004
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType2002 != nil && *p.oneOfType2002.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2002
	}
	return nil
}

func (p *OneOfArchiveOptsUploadParams) UnmarshalJSON(b []byte) error {
	vOneOfType2003 := new(CustomServerUploadParams)
	if err := json.Unmarshal(b, vOneOfType2003); err == nil {
		if "monitoring.v4.serviceability.CustomServerUploadParams" == *vOneOfType2003.ObjectType_ {
			if nil == p.oneOfType2003 {
				p.oneOfType2003 = new(CustomServerUploadParams)
			}
			*p.oneOfType2003 = *vOneOfType2003
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2003.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2003.ObjectType_
			return nil
		}
	}
	vOneOfType2004 := new(StorageContainerUploadParams)
	if err := json.Unmarshal(b, vOneOfType2004); err == nil {
		if "monitoring.v4.serviceability.StorageContainerUploadParams" == *vOneOfType2004.ObjectType_ {
			if nil == p.oneOfType2004 {
				p.oneOfType2004 = new(StorageContainerUploadParams)
			}
			*p.oneOfType2004 = *vOneOfType2004
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2004.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2004.ObjectType_
			return nil
		}
	}
	vOneOfType2001 := new(LocalUploadParams)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "monitoring.v4.serviceability.LocalUploadParams" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(LocalUploadParams)
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
	vOneOfType2002 := new(NtnxServerUploadParams)
	if err := json.Unmarshal(b, vOneOfType2002); err == nil {
		if "monitoring.v4.serviceability.NtnxServerUploadParams" == *vOneOfType2002.ObjectType_ {
			if nil == p.oneOfType2002 {
				p.oneOfType2002 = new(NtnxServerUploadParams)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfArchiveOptsUploadParams"))
}

func (p *OneOfArchiveOptsUploadParams) MarshalJSON() ([]byte, error) {
	if p.oneOfType2003 != nil && *p.oneOfType2003.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2003)
	}
	if p.oneOfType2004 != nil && *p.oneOfType2004.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2004)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType2002 != nil && *p.oneOfType2002.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2002)
	}
	return nil, errors.New("No value to marshal for OneOfArchiveOptsUploadParams")
}

type OneOfConfigurableParamValueRangeMinimumValue struct {
	Discriminator *string  `json:"-"`
	ObjectType_   *string  `json:"-"`
	oneOfType0    *int64   `json:"-"`
	oneOfType1    *float32 `json:"-"`
}

func NewOneOfConfigurableParamValueRangeMinimumValue() *OneOfConfigurableParamValueRangeMinimumValue {
	p := new(OneOfConfigurableParamValueRangeMinimumValue)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfConfigurableParamValueRangeMinimumValue) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfConfigurableParamValueRangeMinimumValue is nil"))
	}
	switch v.(type) {
	case int64:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(int64)
		}
		*p.oneOfType0 = v.(int64)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "Long"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "Long"
	case float32:
		if nil == p.oneOfType1 {
			p.oneOfType1 = new(float32)
		}
		*p.oneOfType1 = v.(float32)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "Float"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "Float"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfConfigurableParamValueRangeMinimumValue) GetValue() interface{} {
	if "Long" == *p.Discriminator {
		return *p.oneOfType0
	}
	if "Float" == *p.Discriminator {
		return *p.oneOfType1
	}
	return nil
}

func (p *OneOfConfigurableParamValueRangeMinimumValue) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(int64)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(int64)
		}
		*p.oneOfType0 = *vOneOfType0
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "Long"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "Long"
		return nil
	}
	vOneOfType1 := new(float32)
	if err := json.Unmarshal(b, vOneOfType1); err == nil {
		if nil == p.oneOfType1 {
			p.oneOfType1 = new(float32)
		}
		*p.oneOfType1 = *vOneOfType1
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "Float"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "Float"
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfConfigurableParamValueRangeMinimumValue"))
}

func (p *OneOfConfigurableParamValueRangeMinimumValue) MarshalJSON() ([]byte, error) {
	if "Long" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if "Float" == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	return nil, errors.New("No value to marshal for OneOfConfigurableParamValueRangeMinimumValue")
}

type OneOfGetAuditApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType0    *Audit                 `json:"-"`
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
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfGetAuditApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfGetAuditApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "monitoring.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetAuditApiResponseData"))
}

func (p *OneOfGetAuditApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfGetAuditApiResponseData")
}

type OneOfListAlertsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    []Alert                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
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
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfListAlertsApiResponseData) GetValue() interface{} {
	if "List<monitoring.v4.serviceability.Alert>" == *p.Discriminator {
		return p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfListAlertsApiResponseData) UnmarshalJSON(b []byte) error {
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "monitoring.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListAlertsApiResponseData"))
}

func (p *OneOfListAlertsApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<monitoring.v4.serviceability.Alert>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfListAlertsApiResponseData")
}

type OneOfCreateUdaPolicyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
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
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "monitoring.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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

type OneOfFindConflictingUdaPoliciesApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
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
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "monitoring.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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

type OneOfListTagsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType2001 []Tag                  `json:"-"`
}

func NewOneOfListTagsApiResponseData() *OneOfListTagsApiResponseData {
	p := new(OneOfListTagsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListTagsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListTagsApiResponseData is nil"))
	}
	switch v.(type) {
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case []Tag:
		p.oneOfType2001 = v.([]Tag)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<monitoring.v4.serviceability.Tag>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<monitoring.v4.serviceability.Tag>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListTagsApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<monitoring.v4.serviceability.Tag>" == *p.Discriminator {
		return p.oneOfType2001
	}
	return nil
}

func (p *OneOfListTagsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "monitoring.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	vOneOfType2001 := new([]Tag)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if len(*vOneOfType2001) == 0 || "monitoring.v4.serviceability.Tag" == *((*vOneOfType2001)[0].ObjectType_) {
			p.oneOfType2001 = *vOneOfType2001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<monitoring.v4.serviceability.Tag>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<monitoring.v4.serviceability.Tag>"
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListTagsApiResponseData"))
}

func (p *OneOfListTagsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<monitoring.v4.serviceability.Tag>" == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfListTagsApiResponseData")
}

type OneOfListAuditsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
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
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "monitoring.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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

type OneOfGetUdaPolicyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
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
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "monitoring.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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

type OneOfUpdateAlertEmailConfigurationApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType0    []import4.AppMessage   `json:"-"`
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
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case []import4.AppMessage:
		p.oneOfType0 = v.([]import4.AppMessage)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<monitoring.v4.error.AppMessage>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<monitoring.v4.error.AppMessage>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfUpdateAlertEmailConfigurationApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<monitoring.v4.error.AppMessage>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfUpdateAlertEmailConfigurationApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "monitoring.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	vOneOfType0 := new([]import4.AppMessage)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "monitoring.v4.error.AppMessage" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<monitoring.v4.error.AppMessage>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<monitoring.v4.error.AppMessage>"
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateAlertEmailConfigurationApiResponseData"))
}

func (p *OneOfUpdateAlertEmailConfigurationApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<monitoring.v4.error.AppMessage>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
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

type OneOfCollectLogsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import3.TaskReference `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfCollectLogsApiResponseData() *OneOfCollectLogsApiResponseData {
	p := new(OneOfCollectLogsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCollectLogsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCollectLogsApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import3.TaskReference)
		}
		*p.oneOfType2001 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfCollectLogsApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfCollectLogsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import3.TaskReference)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "monitoring.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCollectLogsApiResponseData"))
}

func (p *OneOfCollectLogsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfCollectLogsApiResponseData")
}

type OneOfDeleteUdaPolicyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType1    *interface{}           `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
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
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "monitoring.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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

type OneOfUpdateUdaPolicyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType0    []import4.AppMessage   `json:"-"`
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
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case []import4.AppMessage:
		p.oneOfType0 = v.([]import4.AppMessage)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<monitoring.v4.error.AppMessage>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<monitoring.v4.error.AppMessage>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfUpdateUdaPolicyApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<monitoring.v4.error.AppMessage>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfUpdateUdaPolicyApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "monitoring.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	vOneOfType0 := new([]import4.AppMessage)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "monitoring.v4.error.AppMessage" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<monitoring.v4.error.AppMessage>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<monitoring.v4.error.AppMessage>"
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateUdaPolicyApiResponseData"))
}

func (p *OneOfUpdateUdaPolicyApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<monitoring.v4.error.AppMessage>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateUdaPolicyApiResponseData")
}

type OneOfUpdateClusterConfigApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
}

func NewOneOfUpdateClusterConfigApiResponseData() *OneOfUpdateClusterConfigApiResponseData {
	p := new(OneOfUpdateClusterConfigApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateClusterConfigApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateClusterConfigApiResponseData is nil"))
	}
	switch v.(type) {
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfUpdateClusterConfigApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfUpdateClusterConfigApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "monitoring.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateClusterConfigApiResponseData"))
}

func (p *OneOfUpdateClusterConfigApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateClusterConfigApiResponseData")
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
