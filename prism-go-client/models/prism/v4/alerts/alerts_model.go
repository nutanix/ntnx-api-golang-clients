/*
 * Generated file models/prism/v4/alerts/alerts_model.go.
 *
 * Product version: 4.0.2-alpha-1
 *
 * Part of the Nutanix Prism Versioned APIs
 *
 * (c) 2023 Nutanix Inc.  All rights reserved
 *
 */

/*
  Monitor alerts
*/
package alerts

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	import2 "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/common/v1/response"
	import1 "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/prism/v4/common"
	import4 "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/prism/v4/config"
	import3 "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/prism/v4/error"
)

type ActionType int

const (
	ACTIONTYPE_UNKNOWN     ActionType = 0
	ACTIONTYPE_REDACTED    ActionType = 1
	ACTIONTYPE_RESOLVE     ActionType = 2
	ACTIONTYPE_ACKNOWLEDGE ActionType = 3
)

// returns the name of the enum given an ordinal number
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

// returns the enum type given a string value
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
	/**
	  Field specifies if the alert is acknowledged or not.
	*/
	Acknowledged *bool `json:"acknowledged,omitempty"`
	/**
	  Name of the user who acknowledged this alert.
	*/
	AcknowledgedByUsername *string `json:"acknowledgedByUsername,omitempty"`
	/**
	  The timestamp in microseconds when the alert was acknowledged.
	*/
	AcknowledgedTimestamp *int64 `json:"acknowledgedTimestamp,omitempty"`
	/**
	  List of all entities that are affected by the alert.
	*/
	AffectedEntities []import1.EntityReference `json:"affectedEntities,omitempty"`
	/**
	  A preconfigured or dynamically generated unique value for each alert type. For example, A1128  for storage pool space exceeded alerts.
	*/
	AlertType *string `json:"alertType,omitempty"`
	/**
	  Field specifies if the alert is auto resolved or not.
	*/
	AutoResolved *bool `json:"autoResolved,omitempty"`
	/**
	  Various categories into which this alert type can be classified into. For example, Hardware, Storage, License and so on.
	*/
	Classifications []string `json:"classifications,omitempty"`
	/**
	  Cluster UUID associated with the source entity of the alert.
	*/
	ClusterUUID *string `json:"clusterUUID,omitempty"`
	/**
	  Timestamp in microseconds when the alert was created.
	*/
	CreationTimestamp *int64 `json:"creationTimestamp,omitempty"`
	/**
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/**
	  The impact this alert/event will have on the system. For example, availability, performance, capacity and so on.
	*/
	ImpactTypes []import1.ImpactType `json:"impactTypes,omitempty"`
	/**
	  Timestamp in microseconds when the alert was last updated.
	*/
	LastUpdatedTimestamp *int64 `json:"lastUpdatedTimestamp,omitempty"`
	/**
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/**
	  Additional message associated with the alert.
	*/
	Message *string `json:"message,omitempty"`
	/**
	  Details of the metric for a metric-based event.
	*/
	MetricDetails []import1.MetricDetail `json:"metricDetails,omitempty"`
	/**
	  Cluster UUID associated with the cluster where the alert was first raised.
	*/
	OriginatingClusterUUID *string `json:"originatingClusterUUID,omitempty"`
	/**
	  Additional parameters associated with the alert. These parameters can be used to indicate custom key-value pairs for a given alert instance. For example, Service down in Prism Central alert can have the service name as a parameter.
	*/
	Parameters []import1.Parameter `json:"parameters,omitempty"`
	/**
	  Field specifies if the alert is resolved or not.
	*/
	Resolved *bool `json:"resolved,omitempty"`
	/**
	  Name of the user who resolved this alert.
	*/
	ResolvedByUsername *string `json:"resolvedByUsername,omitempty"`
	/**
	  The timestamp in microseconds when the alert was resolved.
	*/
	ResolvedTimestamp *int64 `json:"resolvedTimestamp,omitempty"`
	/**
	  Possible causes, resolutions and additional details to troubleshoot this alert.
	*/
	RootCauseAnalysis []RootCauseAnalysis `json:"rootCauseAnalysis,omitempty"`
	/**
	  The service that raised the alert. For internal Nutanix services, this value is set to "Nutanix".
	*/
	ServiceName *string `json:"serviceName,omitempty"`

	Severity *import1.Severity `json:"severity,omitempty"`
	/**
	  Contains information on the severity change history for alerts. If an alert was de-duplicated without change in severity, then no trail will be present.
	*/
	SeverityTrails []SeverityTrail `json:"severityTrails,omitempty"`

	SourceEntity *import1.AlertEntityReference `json:"sourceEntity,omitempty"`
	/**
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/**
	  Title of the alert.
	*/
	Title *string `json:"title,omitempty"`
	/**
	  Flag to indicate if the alert was generated from a user defined alert policy.
	*/
	UserDefined *bool `json:"userDefined,omitempty"`
}

func NewAlert() *Alert {
	p := new(Alert)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.alerts.Alert"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a1.alerts.Alert"}
	p.UnknownFields_ = map[string]interface{}{}

	p.Acknowledged = new(bool)
	*p.Acknowledged = false
	p.AutoResolved = new(bool)
	*p.AutoResolved = false
	p.Resolved = new(bool)
	*p.Resolved = false
	p.UserDefined = new(bool)
	*p.UserDefined = false

	return p
}

type AlertAction struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	ActionType *ActionType `json:"actionType"`
	/**
	  Unique identifiers for alerts that can be resolved or acknowledged.
	*/
	Uuids []string `json:"uuids"`
}

func (p *AlertAction) MarshalJSON() ([]byte, error) {
	type AlertActionProxy AlertAction
	return json.Marshal(struct {
		*AlertActionProxy
		ActionType *ActionType `json:"actionType,omitempty"`
		Uuids      []string    `json:"uuids,omitempty"`
	}{
		AlertActionProxy: (*AlertActionProxy)(p),
		ActionType:       p.ActionType,
		Uuids:            p.Uuids,
	})
}

func NewAlertAction() *AlertAction {
	p := new(AlertAction)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.alerts.AlertAction"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a1.alerts.AlertAction"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
REST response for all response codes in api path /prism/v4.0.a1/alerts/{extId} Get operation
*/
type AlertApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfAlertApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewAlertApiResponse() *AlertApiResponse {
	p := new(AlertApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.alerts.AlertApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a1.alerts.AlertApiResponse"}
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

type AlertDb struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Acknowledged *bool `json:"acknowledged,omitempty"`

	AlertType *string `json:"alertType,omitempty"`

	AutoResolved *bool `json:"autoResolved,omitempty"`

	Classifications []string `json:"classifications,omitempty"`

	Cluster []string `json:"cluster,omitempty"`

	ClusterUUID *string `json:"clusterUUID,omitempty"`

	CreationTimestamp *int64 `json:"creationTimestamp,omitempty"`
	/**
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`

	ImpactTypes []string `json:"impactTypes,omitempty"`

	LastUpdatedTimestamp *int64 `json:"lastUpdatedTimestamp,omitempty"`
	/**
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`

	Protobuf *string `json:"protobuf,omitempty"`

	Resolved *bool `json:"resolved,omitempty"`

	ServiceName *string `json:"serviceName,omitempty"`

	Severity *import1.Severity `json:"severity,omitempty"`

	SourceEntity *import1.AlertEntityReference `json:"sourceEntity,omitempty"`
	/**
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewAlertDb() *AlertDb {
	p := new(AlertDb)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.alerts.AlertDb"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a1.alerts.AlertDb"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type AlertDbProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Acknowledged *bool `json:"acknowledged,omitempty"`

	AlertType *string `json:"alertType,omitempty"`

	AutoResolved *bool `json:"autoResolved,omitempty"`

	Classifications []string `json:"classifications,omitempty"`

	Cluster []string `json:"cluster,omitempty"`

	ClusterUUID *string `json:"clusterUUID,omitempty"`

	CreationTimestamp *int64 `json:"creationTimestamp,omitempty"`
	/**
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`

	ImpactTypes []string `json:"impactTypes,omitempty"`

	LastUpdatedTimestamp *int64 `json:"lastUpdatedTimestamp,omitempty"`
	/**
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`

	Protobuf *string `json:"protobuf,omitempty"`

	Resolved *bool `json:"resolved,omitempty"`

	RootCauseAnalysisProjection *RootCauseAnalysisProjection `json:"rootCauseAnalysisProjection,omitempty"`

	ServiceName *string `json:"serviceName,omitempty"`

	Severity *import1.Severity `json:"severity,omitempty"`

	SourceEntity *import1.AlertEntityReference `json:"sourceEntity,omitempty"`
	/**
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewAlertDbProjection() *AlertDbProjection {
	p := new(AlertDbProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.alerts.AlertDbProjection"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a1.alerts.AlertDbProjection"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
REST response for all response codes in api path /prism/v4.0.a1/alerts Get operation
*/
type AlertListApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfAlertListApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewAlertListApiResponse() *AlertListApiResponse {
	p := new(AlertListApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.alerts.AlertListApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a1.alerts.AlertListApiResponse"}
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

/**
REST response for all response codes in api path /prism/v4.0.a1/alerts/$action/performAction Post operation
*/
type AlertsActionTaskApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfAlertsActionTaskApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewAlertsActionTaskApiResponse() *AlertsActionTaskApiResponse {
	p := new(AlertsActionTaskApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.alerts.AlertsActionTaskApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a1.alerts.AlertsActionTaskApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *AlertsActionTaskApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *AlertsActionTaskApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfAlertsActionTaskApiResponseData()
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

type RootCauseAnalysis struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Cause *string `json:"cause,omitempty"`

	Detail *string `json:"detail,omitempty"`
	/**
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/**
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`

	Resolution *string `json:"resolution,omitempty"`
	/**
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewRootCauseAnalysis() *RootCauseAnalysis {
	p := new(RootCauseAnalysis)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.alerts.RootCauseAnalysis"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a1.alerts.RootCauseAnalysis"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type RootCauseAnalysisProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Cause *string `json:"cause,omitempty"`

	Detail *string `json:"detail,omitempty"`
	/**
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/**
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`

	Resolution *string `json:"resolution,omitempty"`
	/**
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewRootCauseAnalysisProjection() *RootCauseAnalysisProjection {
	p := new(RootCauseAnalysisProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.alerts.RootCauseAnalysisProjection"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a1.alerts.RootCauseAnalysisProjection"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Contains information on the severity change history for alerts. If an alert was de-duplicated without change in severity, then no trail will be present.
*/
type SeverityTrail struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Severity *import1.Severity `json:"severity,omitempty"`
	/**
	  The timestamp in microseconds when the severity of the alert was changed.
	*/
	SeverityChangeTimestamp *int64 `json:"severityChangeTimestamp,omitempty"`
}

func NewSeverityTrail() *SeverityTrail {
	p := new(SeverityTrail)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.alerts.SeverityTrail"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a1.alerts.SeverityTrail"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type OneOfAlertsActionTaskApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType0    *import4.TaskReference `json:"-"`
}

func NewOneOfAlertsActionTaskApiResponseData() *OneOfAlertsActionTaskApiResponseData {
	p := new(OneOfAlertsActionTaskApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfAlertsActionTaskApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfAlertsActionTaskApiResponseData is nil"))
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
	case import4.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import4.TaskReference)
		}
		*p.oneOfType0 = v.(import4.TaskReference)
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

func (p *OneOfAlertsActionTaskApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfAlertsActionTaskApiResponseData) UnmarshalJSON(b []byte) error {
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
	vOneOfType0 := new(import4.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import4.TaskReference)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfAlertsActionTaskApiResponseData"))
}

func (p *OneOfAlertsActionTaskApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfAlertsActionTaskApiResponseData")
}

type OneOfAlertApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *Alert                 `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
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

func (p *OneOfAlertApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfAlertApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(Alert)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.alerts.Alert" == *vOneOfType0.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfAlertApiResponseData"))
}

func (p *OneOfAlertApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfAlertApiResponseData")
}

type OneOfAlertListApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType0    []Alert                `json:"-"`
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
		*p.Discriminator = "List<prism.v4.alerts.Alert>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<prism.v4.alerts.Alert>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfAlertListApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<prism.v4.alerts.Alert>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfAlertListApiResponseData) UnmarshalJSON(b []byte) error {
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
	vOneOfType0 := new([]Alert)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "prism.v4.alerts.Alert" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<prism.v4.alerts.Alert>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<prism.v4.alerts.Alert>"
			return nil

		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfAlertListApiResponseData"))
}

func (p *OneOfAlertListApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<prism.v4.alerts.Alert>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfAlertListApiResponseData")
}
