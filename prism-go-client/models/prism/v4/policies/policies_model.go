/*
 * Generated file models/prism/v4/policies/policies_model.go.
 *
 * Product version: 4.0.2-alpha-1
 *
 * Part of the Nutanix Prism Versioned APIs
 *
 * (c) 2023 Nutanix Inc.  All rights reserved
 *
 */

/*
  Create and Update Monitoring policies
*/
package policies

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	import2 "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/common/v1/response"
	import3 "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/prism/v4/common"
	import1 "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/prism/v4/error"
)

type AlertConfigExceptionGroup struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	AutoResolve *AutoResolve `json:"autoResolve,omitempty"`
	/**
	  List of cluster UUIDs on which the exceptions can be set.
	*/
	ClusterUuids []string `json:"clusterUuids,omitempty"`
	/**
	  Enable/Disable for each severity information.
	*/
	SeverityThresholdInfos []SeverityThresholdInfo `json:"severityThresholdInfos,omitempty"`
}

func NewAlertConfigExceptionGroup() *AlertConfigExceptionGroup {
	p := new(AlertConfigExceptionGroup)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.policies.AlertConfigExceptionGroup"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a1.policies.AlertConfigExceptionGroup"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
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

// returns the name of the enum given an ordinal number
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

// returns the enum type given a string value
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

/**
REST response for all response codes in api path /prism/v4.0.a1/policies/user-defined Post operation
*/
type CreateUdaPolicyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateUdaPolicyApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewCreateUdaPolicyApiResponse() *CreateUdaPolicyApiResponse {
	p := new(CreateUdaPolicyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.policies.CreateUdaPolicyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a1.policies.CreateUdaPolicyApiResponse"}
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

/**
REST response for all response codes in api path /prism/v4.0.a1/policies/user-defined/{extId} Delete operation
*/
type DeleteUdaPolicyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteUdaPolicyApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDeleteUdaPolicyApiResponse() *DeleteUdaPolicyApiResponse {
	p := new(DeleteUdaPolicyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.policies.DeleteUdaPolicyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a1.policies.DeleteUdaPolicyApiResponse"}
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

/**
REST response for all response codes in api path /prism/v4.0.a1/policies/system-defined/{entityUid} Get operation
*/
type GetSdaPolicyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetSdaPolicyApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetSdaPolicyApiResponse() *GetSdaPolicyApiResponse {
	p := new(GetSdaPolicyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.policies.GetSdaPolicyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a1.policies.GetSdaPolicyApiResponse"}
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

/**
REST response for all response codes in api path /prism/v4.0.a1/policies/user-defined/{extId} Get operation
*/
type GetUdaPolicyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetUdaPolicyApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetUdaPolicyApiResponse() *GetUdaPolicyApiResponse {
	p := new(GetUdaPolicyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.policies.GetUdaPolicyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a1.policies.GetUdaPolicyApiResponse"}
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

/**
REST response for all response codes in api path /prism/v4.0.a1/policies/$action/find-conflicts Post operation
*/
type ListConflictingUdaPoliciesApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListConflictingUdaPoliciesApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListConflictingUdaPoliciesApiResponse() *ListConflictingUdaPoliciesApiResponse {
	p := new(ListConflictingUdaPoliciesApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.policies.ListConflictingUdaPoliciesApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a1.policies.ListConflictingUdaPoliciesApiResponse"}
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

/**
REST response for all response codes in api path /prism/v4.0.a1/policies/system-defined Get operation
*/
type ListSdaPoliciesApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListSdaPoliciesApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListSdaPoliciesApiResponse() *ListSdaPoliciesApiResponse {
	p := new(ListSdaPoliciesApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.policies.ListSdaPoliciesApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a1.policies.ListSdaPoliciesApiResponse"}
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

/**
REST response for all response codes in api path /prism/v4.0.a1/policies/user-defined Get operation
*/
type ListUdaPoliciesApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListUdaPoliciesApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListUdaPoliciesApiResponse() *ListUdaPoliciesApiResponse {
	p := new(ListUdaPoliciesApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.policies.ListUdaPoliciesApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a1.policies.ListUdaPoliciesApiResponse"}
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

/**
An alert policy that is related to the entities of the current policy.
*/
type RelatedPolicy struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Uuid of the entity to which the UDA policy is associated.
	*/
	EntityUuid *string `json:"entityUuid"`
	/**
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
	*p.ObjectType_ = "prism.v4.policies.RelatedPolicy"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a1.policies.RelatedPolicy"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Enable/Disable for each severity information.
*/
type SeverityThresholdInfo struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Indicates if the system defined policy is enabled for a given severity.
	*/
	Enabled *bool `json:"enabled,omitempty"`

	Severity *import3.Severity `json:"severity,omitempty"`
}

func NewSeverityThresholdInfo() *SeverityThresholdInfo {
	p := new(SeverityThresholdInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.policies.SeverityThresholdInfo"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a1.policies.SeverityThresholdInfo"}
	p.UnknownFields_ = map[string]interface{}{}

	p.Enabled = new(bool)
	*p.Enabled = false

	return p
}

type SystemDefined struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Possible causes, resolutions and additional details to troubleshoot this alert.
	*/
	CauseAndResolutions []import3.CauseAndResolution `json:"CauseAndResolutions,omitempty"`
	/**
	  Affected entity types.
	*/
	AffectedEntityTypes []string `json:"affectedEntityTypes,omitempty"`
	/**
	  List of clusters that have their alert configurable parameters different from that of in Prism Central.
	*/
	AlertConfigExceptionGroups []AlertConfigExceptionGroup `json:"alertConfigExceptionGroups,omitempty"`
	/**
	  System defined alert policy ID.
	*/
	AlertTypeId *string `json:"alertTypeId,omitempty"`

	AutoResolve *AutoResolve `json:"autoResolve,omitempty"`
	/**
	  Various categories into which this alert type can be classified into. For example, Hardware, Storage, License and so on.
	*/
	Classifications []string `json:"classifications,omitempty"`
	/**
	  Entity UID.
	*/
	EntityUid *string `json:"entityUid,omitempty"`
	/**
	  Numbers of clusters having exceptions with respect to alert config.
	*/
	ExceptionCount *int64 `json:"exceptionCount,omitempty"`
	/**
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/**
	  Indicates whether the system defined alert policies is a default one or not.
	*/
	GlobalConfig *bool `json:"globalConfig,omitempty"`
	/**
	  Impact type on affected entities.
	*/
	ImpactTypes []import3.ImpactType `json:"impactTypes,omitempty"`
	/**
	  List of KB article numbers.
	*/
	KbArticles []string `json:"kbArticles,omitempty"`
	/**
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/**
	  Additional message associated with the event.
	*/
	Message *string `json:"message,omitempty"`
	/**
	  Name of the user who last modified the system defined alert policy.
	*/
	ModifiedByUsername *string `json:"modifiedByUsername,omitempty"`
	/**
	  Timestamp for when the system defined alert policy was last modified.
	*/
	ModifiedTimestamp *int64 `json:"modifiedTimestamp,omitempty"`
	/**
	  Enable/Disable for each severity information.
	*/
	SeverityThresholdInfos []SeverityThresholdInfo `json:"severityThresholdInfos,omitempty"`
	/**
	  Title for the alert policy which can include dynamic parameters like IP address, VM name and so on.
	*/
	SmartTitle *string `json:"smartTitle,omitempty"`
	/**
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/**
	  Indicates whether this system defined alert policy is specific to the tenant or not.
	*/
	TenantSpecific *bool `json:"tenantSpecific,omitempty"`
	/**
	  Title of the event.
	*/
	Title *string `json:"title,omitempty"`
}

func NewSystemDefined() *SystemDefined {
	p := new(SystemDefined)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.policies.SystemDefined"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a1.policies.SystemDefined"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type TriggerCondition struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Conditions to be met to trigger the alert. Conditions are expressed in FIQL.
	*/
	Condition *string `json:"condition"`

	ConditionType *import3.ConditionType `json:"conditionType"`

	SeverityLevel *import3.Severity `json:"severityLevel"`
}

func (p *TriggerCondition) MarshalJSON() ([]byte, error) {
	type TriggerConditionProxy TriggerCondition
	return json.Marshal(struct {
		*TriggerConditionProxy
		Condition     *string                `json:"condition,omitempty"`
		ConditionType *import3.ConditionType `json:"conditionType,omitempty"`
		SeverityLevel *import3.Severity      `json:"severityLevel,omitempty"`
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
	*p.ObjectType_ = "prism.v4.policies.TriggerCondition"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a1.policies.TriggerCondition"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
REST response for all response codes in api path /prism/v4.0.a1/policies/system-defined/{entityUid} Put operation
*/
type UpdateSdaPolicyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateSdaPolicyApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewUpdateSdaPolicyApiResponse() *UpdateSdaPolicyApiResponse {
	p := new(UpdateSdaPolicyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.policies.UpdateSdaPolicyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a1.policies.UpdateSdaPolicyApiResponse"}
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

/**
REST response for all response codes in api path /prism/v4.0.a1/policies/user-defined/{extId} Put operation
*/
type UpdateUdaPolicyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateUdaPolicyApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewUpdateUdaPolicyApiResponse() *UpdateUdaPolicyApiResponse {
	p := new(UpdateUdaPolicyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.policies.UpdateUdaPolicyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a1.policies.UpdateUdaPolicyApiResponse"}
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

type UserDefined struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Indicates whether the auto resolve feature is enabled for this policy.
	*/
	AutoResolve *bool `json:"autoResolve"`
	/**
	  Username of the user who created the policy.
	*/
	CreatedBy *string `json:"createdBy,omitempty"`
	/**
	  Description of the policy.
	*/
	Description *string `json:"description,omitempty"`
	/**
	  Enable/Disable flag for the policy.
	*/
	Enabled *bool `json:"enabled"`
	/**
	  Unique identifiers of the entities associated with the user defined alert policy.
	*/
	EntityIds []string `json:"entityIds,omitempty"`
	/**
	  Entity type associated with the user defined alert policy. For example, VM, Node, Cluster etc.
	*/
	EntityType *string `json:"entityType,omitempty"`
	/**
	  Error if conflicting alert policies are found.
	*/
	ErrorOnConflict *bool `json:"errorOnConflict,omitempty"`
	/**
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/**
	  Filter expression for the policy in FIQL.
	*/
	Filter *string `json:"filter,omitempty"`
	/**
	  Impact Types for the associated resulting alert.
	*/
	ImpactTypes []import3.ImpactType `json:"impactTypes,omitempty"`
	/**
	  Last updated timestamp of the policy. This value will be used as the CAS value during updates.
	*/
	LastUpdatedTimestamp *int64 `json:"lastUpdatedTimestamp,omitempty"`
	/**
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/**
	  List of IDs of the alert policies that should be overridden.
	*/
	PoliciesToOverride []string `json:"policiesToOverride,omitempty"`
	/**
	  List of alert policies that are related to the entities of the current policy.
	*/
	RelatedPolicies []RelatedPolicy `json:"relatedPolicies,omitempty"`
	/**
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/**
	  Title of the policy.
	*/
	Title *string `json:"title"`
	/**
	  Trigger conditions for the policy. If there are multiple trigger conditions, all of them will be considered during the operation.
	*/
	TriggerConditions []TriggerCondition `json:"triggerConditions"`
	/**
	  Waiting duration in seconds before triggering the alert, when the specified condition is met.
	*/
	TriggerWaitPeriod *int64 `json:"triggerWaitPeriod,omitempty"`
}

func (p *UserDefined) MarshalJSON() ([]byte, error) {
	type UserDefinedProxy UserDefined
	return json.Marshal(struct {
		*UserDefinedProxy
		AutoResolve       *bool              `json:"autoResolve,omitempty"`
		Enabled           *bool              `json:"enabled,omitempty"`
		Title             *string            `json:"title,omitempty"`
		TriggerConditions []TriggerCondition `json:"triggerConditions,omitempty"`
	}{
		UserDefinedProxy:  (*UserDefinedProxy)(p),
		AutoResolve:       p.AutoResolve,
		Enabled:           p.Enabled,
		Title:             p.Title,
		TriggerConditions: p.TriggerConditions,
	})
}

func NewUserDefined() *UserDefined {
	p := new(UserDefined)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.policies.UserDefined"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a1.policies.UserDefined"}
	p.UnknownFields_ = map[string]interface{}{}

	p.AutoResolve = new(bool)
	*p.AutoResolve = true
	p.Enabled = new(bool)
	*p.Enabled = false
	p.ErrorOnConflict = new(bool)
	*p.ErrorOnConflict = true

	return p
}

type OneOfDeleteUdaPolicyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType1    *interface{}           `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
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
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "prism.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	oneOfType400  *import1.ErrorResponse `json:"-"`
	oneOfType0    []SystemDefined        `json:"-"`
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
	case []SystemDefined:
		p.oneOfType0 = v.([]SystemDefined)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<prism.v4.policies.SystemDefined>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<prism.v4.policies.SystemDefined>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListSdaPoliciesApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<prism.v4.policies.SystemDefined>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfListSdaPoliciesApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "prism.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new([]SystemDefined)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "prism.v4.policies.SystemDefined" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<prism.v4.policies.SystemDefined>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<prism.v4.policies.SystemDefined>"
			return nil

		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListSdaPoliciesApiResponseData"))
}

func (p *OneOfListSdaPoliciesApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<prism.v4.policies.SystemDefined>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfListSdaPoliciesApiResponseData")
}

type OneOfUpdateSdaPolicyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *SystemDefined         `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
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
	case SystemDefined:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(SystemDefined)
		}
		*p.oneOfType0 = v.(SystemDefined)
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

func (p *OneOfUpdateSdaPolicyApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfUpdateSdaPolicyApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(SystemDefined)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.policies.SystemDefined" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(SystemDefined)
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
		if "prism.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateSdaPolicyApiResponseData"))
}

func (p *OneOfUpdateSdaPolicyApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateSdaPolicyApiResponseData")
}

type OneOfUpdateUdaPolicyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *UserDefined           `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
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
	case UserDefined:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(UserDefined)
		}
		*p.oneOfType0 = v.(UserDefined)
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

func (p *OneOfUpdateUdaPolicyApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfUpdateUdaPolicyApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(UserDefined)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.policies.UserDefined" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(UserDefined)
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
		if "prism.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateUdaPolicyApiResponseData"))
}

func (p *OneOfUpdateUdaPolicyApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateUdaPolicyApiResponseData")
}

type OneOfListConflictingUdaPoliciesApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	oneOfType0    []UserDefined          `json:"-"`
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
	case []UserDefined:
		p.oneOfType0 = v.([]UserDefined)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<prism.v4.policies.UserDefined>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<prism.v4.policies.UserDefined>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListConflictingUdaPoliciesApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<prism.v4.policies.UserDefined>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfListConflictingUdaPoliciesApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "prism.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new([]UserDefined)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "prism.v4.policies.UserDefined" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<prism.v4.policies.UserDefined>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<prism.v4.policies.UserDefined>"
			return nil

		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListConflictingUdaPoliciesApiResponseData"))
}

func (p *OneOfListConflictingUdaPoliciesApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<prism.v4.policies.UserDefined>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfListConflictingUdaPoliciesApiResponseData")
}

type OneOfListUdaPoliciesApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	oneOfType0    []UserDefined          `json:"-"`
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
	case []UserDefined:
		p.oneOfType0 = v.([]UserDefined)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<prism.v4.policies.UserDefined>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<prism.v4.policies.UserDefined>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListUdaPoliciesApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<prism.v4.policies.UserDefined>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfListUdaPoliciesApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "prism.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new([]UserDefined)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "prism.v4.policies.UserDefined" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<prism.v4.policies.UserDefined>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<prism.v4.policies.UserDefined>"
			return nil

		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListUdaPoliciesApiResponseData"))
}

func (p *OneOfListUdaPoliciesApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<prism.v4.policies.UserDefined>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfListUdaPoliciesApiResponseData")
}

type OneOfGetUdaPolicyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *UserDefined           `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
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
	case UserDefined:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(UserDefined)
		}
		*p.oneOfType0 = v.(UserDefined)
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

func (p *OneOfGetUdaPolicyApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetUdaPolicyApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(UserDefined)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.policies.UserDefined" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(UserDefined)
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
		if "prism.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetUdaPolicyApiResponseData"))
}

func (p *OneOfGetUdaPolicyApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetUdaPolicyApiResponseData")
}

type OneOfCreateUdaPolicyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *UserDefined           `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
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
	case UserDefined:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(UserDefined)
		}
		*p.oneOfType0 = v.(UserDefined)
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

func (p *OneOfCreateUdaPolicyApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfCreateUdaPolicyApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(UserDefined)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.policies.UserDefined" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(UserDefined)
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
		if "prism.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateUdaPolicyApiResponseData"))
}

func (p *OneOfCreateUdaPolicyApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfCreateUdaPolicyApiResponseData")
}

type OneOfGetSdaPolicyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *SystemDefined         `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
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
	case SystemDefined:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(SystemDefined)
		}
		*p.oneOfType0 = v.(SystemDefined)
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
	vOneOfType0 := new(SystemDefined)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.policies.SystemDefined" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(SystemDefined)
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
		if "prism.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
