/*
 * Generated file models/aiops/v4/config/config_model.go.
 *
 * Product version: 4.0.1
 *
 * Part of the Nutanix AIOps APIs
 *
 * (c) 2024 Nutanix Inc.  All rights reserved
 *
 */

/*
  Configure Tasks
*/
package config

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	import4 "github.com/nutanix/ntnx-api-golang-clients/aiops-go-client/v4/models/aiops/v4/common"
	import2 "github.com/nutanix/ntnx-api-golang-clients/aiops-go-client/v4/models/aiops/v4/error"
	import5 "github.com/nutanix/ntnx-api-golang-clients/aiops-go-client/v4/models/common/v1/config"
	import3 "github.com/nutanix/ntnx-api-golang-clients/aiops-go-client/v4/models/common/v1/response"
	import6 "github.com/nutanix/ntnx-api-golang-clients/aiops-go-client/v4/models/common/v1/stats"
	import1 "github.com/nutanix/ntnx-api-golang-clients/aiops-go-client/v4/models/prism/v4/config"
	"time"
)

/*
REST response for all response codes in API path /aiops/v4.0/config/ignore-windows/$actions/calculate-runway Post operation
*/
type CalculateRunwayApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCalculateRunwayApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewCalculateRunwayApiResponse() *CalculateRunwayApiResponse {
	p := new(CalculateRunwayApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.config.CalculateRunwayApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CalculateRunwayApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CalculateRunwayApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCalculateRunwayApiResponseData()
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
CapacityUpdateConfig workload.
*/
type CapacityUpdateConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	ChangeType *CapacityUpdateType `json:"changeType"`
	/*
	  Percentage change in demand in case of capacity update workload.
	*/
	PercentageChange *int `json:"percentageChange"`
}

func (p *CapacityUpdateConfig) MarshalJSON() ([]byte, error) {
	type CapacityUpdateConfigProxy CapacityUpdateConfig
	return json.Marshal(struct {
		*CapacityUpdateConfigProxy
		ChangeType       *CapacityUpdateType `json:"changeType,omitempty"`
		PercentageChange *int                `json:"percentageChange,omitempty"`
	}{
		CapacityUpdateConfigProxy: (*CapacityUpdateConfigProxy)(p),
		ChangeType:                p.ChangeType,
		PercentageChange:          p.PercentageChange,
	})
}

func NewCapacityUpdateConfig() *CapacityUpdateConfig {
	p := new(CapacityUpdateConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.config.CapacityUpdateConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Enum describing the adjusted change in workload type.
*/
type CapacityUpdateType int

const (
	CAPACITYUPDATETYPE_UNKNOWN  CapacityUpdateType = 0
	CAPACITYUPDATETYPE_REDACTED CapacityUpdateType = 1
	CAPACITYUPDATETYPE_INCREASE CapacityUpdateType = 2
	CAPACITYUPDATETYPE_DECREASE CapacityUpdateType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *CapacityUpdateType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"INCREASE",
		"DECREASE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e CapacityUpdateType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"INCREASE",
		"DECREASE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *CapacityUpdateType) index(name string) CapacityUpdateType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"INCREASE",
		"DECREASE",
	}
	for idx := range names {
		if names[idx] == name {
			return CapacityUpdateType(idx)
		}
	}
	return CAPACITYUPDATETYPE_UNKNOWN
}

func (e *CapacityUpdateType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for CapacityUpdateType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *CapacityUpdateType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e CapacityUpdateType) Ref() *CapacityUpdateType {
	return &e
}

/*
Detail of a category containing the name of the format "key/value".
*/
type CategoryDetail struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A flag that indicates whether this category is a critical out-of-box category, and might not be removed from the policy.
	*/
	IsDefault *bool `json:"isDefault,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import3.ApiLink `json:"links,omitempty"`
	/*
	  Category name of the format "key/value"
	*/
	Name *string `json:"name"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *CategoryDetail) MarshalJSON() ([]byte, error) {
	type CategoryDetailProxy CategoryDetail
	return json.Marshal(struct {
		*CategoryDetailProxy
		Name *string `json:"name,omitempty"`
	}{
		CategoryDetailProxy: (*CategoryDetailProxy)(p),
		Name:                p.Name,
	})
}

func NewCategoryDetail() *CategoryDetail {
	p := new(CategoryDetail)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.config.CategoryDetail"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	p.IsDefault = new(bool)
	*p.IsDefault = false

	return p
}

/*
Enum describing Xen provision type for Citrix.
*/
type CitrixXenProvisionType int

const (
	CITRIXXENPROVISIONTYPE_UNKNOWN  CitrixXenProvisionType = 0
	CITRIXXENPROVISIONTYPE_REDACTED CitrixXenProvisionType = 1
	CITRIXXENPROVISIONTYPE_PVS      CitrixXenProvisionType = 2
	CITRIXXENPROVISIONTYPE_MCS      CitrixXenProvisionType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *CitrixXenProvisionType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PVS",
		"MCS",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e CitrixXenProvisionType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PVS",
		"MCS",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *CitrixXenProvisionType) index(name string) CitrixXenProvisionType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PVS",
		"MCS",
	}
	for idx := range names {
		if names[idx] == name {
			return CitrixXenProvisionType(idx)
		}
	}
	return CITRIXXENPROVISIONTYPE_UNKNOWN
}

func (e *CitrixXenProvisionType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for CitrixXenProvisionType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *CitrixXenProvisionType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e CitrixXenProvisionType) Ref() *CitrixXenProvisionType {
	return &e
}

/*
Xen workload with Citrix as a vendor.
*/
type CitrixXenWorkload struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Size in GB for MCS different disk per VM.
	*/
	McsDiffSizeGb *int `json:"mcsDiffSizeGb"`

	OperatingSystem *XenOperatingSystem `json:"operatingSystem"`

	ProvisionType *CitrixXenProvisionType `json:"provisionType"`
	/*
	  Size of PVS cache write per VM in GB.
	*/
	PvsWriteCacheSizeGb *int `json:"pvsWriteCacheSizeGb"`
	/*
	  Space consumed by each Xen Server Image.
	*/
	SystemDataGb *int `json:"systemDataGb"`
	/*
	  Number of users for Xen workload.
	*/
	UserCount *int `json:"userCount"`
	/*
	  Size for the per user profile data in MB.
	*/
	UserProfileDataMb *int `json:"userProfileDataMb"`

	Vendor *CitrixXenWorkloadVendor `json:"vendor"`
}

func (p *CitrixXenWorkload) MarshalJSON() ([]byte, error) {
	type CitrixXenWorkloadProxy CitrixXenWorkload
	return json.Marshal(struct {
		*CitrixXenWorkloadProxy
		McsDiffSizeGb       *int                     `json:"mcsDiffSizeGb,omitempty"`
		OperatingSystem     *XenOperatingSystem      `json:"operatingSystem,omitempty"`
		ProvisionType       *CitrixXenProvisionType  `json:"provisionType,omitempty"`
		PvsWriteCacheSizeGb *int                     `json:"pvsWriteCacheSizeGb,omitempty"`
		SystemDataGb        *int                     `json:"systemDataGb,omitempty"`
		UserCount           *int                     `json:"userCount,omitempty"`
		UserProfileDataMb   *int                     `json:"userProfileDataMb,omitempty"`
		Vendor              *CitrixXenWorkloadVendor `json:"vendor,omitempty"`
	}{
		CitrixXenWorkloadProxy: (*CitrixXenWorkloadProxy)(p),
		McsDiffSizeGb:          p.McsDiffSizeGb,
		OperatingSystem:        p.OperatingSystem,
		ProvisionType:          p.ProvisionType,
		PvsWriteCacheSizeGb:    p.PvsWriteCacheSizeGb,
		SystemDataGb:           p.SystemDataGb,
		UserCount:              p.UserCount,
		UserProfileDataMb:      p.UserProfileDataMb,
		Vendor:                 p.Vendor,
	})
}

func NewCitrixXenWorkload() *CitrixXenWorkload {
	p := new(CitrixXenWorkload)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.config.CitrixXenWorkload"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Enum describing Xen vendor for Citrix.
*/
type CitrixXenWorkloadVendor int

const (
	CITRIXXENWORKLOADVENDOR_UNKNOWN  CitrixXenWorkloadVendor = 0
	CITRIXXENWORKLOADVENDOR_REDACTED CitrixXenWorkloadVendor = 1
	CITRIXXENWORKLOADVENDOR_XEN_APP  CitrixXenWorkloadVendor = 2
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *CitrixXenWorkloadVendor) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"XEN_APP",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e CitrixXenWorkloadVendor) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"XEN_APP",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *CitrixXenWorkloadVendor) index(name string) CitrixXenWorkloadVendor {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"XEN_APP",
	}
	for idx := range names {
		if names[idx] == name {
			return CitrixXenWorkloadVendor(idx)
		}
	}
	return CITRIXXENWORKLOADVENDOR_UNKNOWN
}

func (e *CitrixXenWorkloadVendor) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for CitrixXenWorkloadVendor:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *CitrixXenWorkloadVendor) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e CitrixXenWorkloadVendor) Ref() *CitrixXenWorkloadVendor {
	return &e
}

/*
Cluster specification.
*/
type ClusterConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	DataStoreConfig *DataStoreConfig `json:"dataStoreConfig"`
	/*
	  List containing the metadata about the nodes in a cluster. Nodes can be user added or existing.
	*/
	NodeConfigs []NodeConfig `json:"nodeConfigs,omitempty"`
}

func (p *ClusterConfig) MarshalJSON() ([]byte, error) {
	type ClusterConfigProxy ClusterConfig
	return json.Marshal(struct {
		*ClusterConfigProxy
		DataStoreConfig *DataStoreConfig `json:"dataStoreConfig,omitempty"`
	}{
		ClusterConfigProxy: (*ClusterConfigProxy)(p),
		DataStoreConfig:    p.DataStoreConfig,
	})
}

func NewClusterConfig() *ClusterConfig {
	p := new(ClusterConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.config.ClusterConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /aiops/v4.0/config/ignore-windows Post operation
*/
type CreateIgnoreWindowApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateIgnoreWindowApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewCreateIgnoreWindowApiResponse() *CreateIgnoreWindowApiResponse {
	p := new(CreateIgnoreWindowApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.config.CreateIgnoreWindowApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreateIgnoreWindowApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreateIgnoreWindowApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreateIgnoreWindowApiResponseData()
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
REST response for all response codes in API path /aiops/v4.0/config/scenarios Post operation
*/
type CreateScenarioApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateScenarioApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewCreateScenarioApiResponse() *CreateScenarioApiResponse {
	p := new(CreateScenarioApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.config.CreateScenarioApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreateScenarioApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreateScenarioApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreateScenarioApiResponseData()
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
REST response for all response codes in API path /aiops/v4.0/config/simulations Post operation
*/
type CreateSimulationApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateSimulationApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewCreateSimulationApiResponse() *CreateSimulationApiResponse {
	p := new(CreateSimulationApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.config.CreateSimulationApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreateSimulationApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreateSimulationApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreateSimulationApiResponseData()
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
The data store configuration. For example, compression saving percentage, CPU over-commit ratio and so on.
*/
type DataStoreConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Compression saving percentage of the cluster resources.
	*/
	CompressionSavingPercent *float64 `json:"compressionSavingPercent"`
	/*
	  CPU overcommit ratio.
	*/
	CpuOverCommitRatio *float64 `json:"cpuOverCommitRatio"`
	/*
	  CPU reservation percentage.
	*/
	CpuReservationPercentage *float64 `json:"cpuReservationPercentage"`
	/*
	  De-dupe saving percentage of the cluster resources.
	*/
	DedupSavingPercent *float64 `json:"dedupSavingPercent"`
	/*
	  Erasure coding saving percentage of the cluster resources.
	*/
	ErasureCodingSavingPercent *float64 `json:"erasureCodingSavingPercent"`
	/*
	  Overall saving percentage of the cluster resources.
	*/
	OverallSavingPercent *float64 `json:"overallSavingPercent"`
	/*
	  RAM overcommit ratio.
	*/
	RamOverCommitRatio *float64 `json:"ramOverCommitRatio"`
	/*
	  RAM reservation percentage.
	*/
	RamReservationPercentage *float64 `json:"ramReservationPercentage"`

	ReplicationFactor *ReplicationFactor `json:"replicationFactor"`
	/*
	  Storage reservation percentage.
	*/
	StorageReservationPercentage *float64 `json:"storageReservationPercentage"`
}

func (p *DataStoreConfig) MarshalJSON() ([]byte, error) {
	type DataStoreConfigProxy DataStoreConfig
	return json.Marshal(struct {
		*DataStoreConfigProxy
		CompressionSavingPercent     *float64           `json:"compressionSavingPercent,omitempty"`
		CpuOverCommitRatio           *float64           `json:"cpuOverCommitRatio,omitempty"`
		CpuReservationPercentage     *float64           `json:"cpuReservationPercentage,omitempty"`
		DedupSavingPercent           *float64           `json:"dedupSavingPercent,omitempty"`
		ErasureCodingSavingPercent   *float64           `json:"erasureCodingSavingPercent,omitempty"`
		OverallSavingPercent         *float64           `json:"overallSavingPercent,omitempty"`
		RamOverCommitRatio           *float64           `json:"ramOverCommitRatio,omitempty"`
		RamReservationPercentage     *float64           `json:"ramReservationPercentage,omitempty"`
		ReplicationFactor            *ReplicationFactor `json:"replicationFactor,omitempty"`
		StorageReservationPercentage *float64           `json:"storageReservationPercentage,omitempty"`
	}{
		DataStoreConfigProxy:         (*DataStoreConfigProxy)(p),
		CompressionSavingPercent:     p.CompressionSavingPercent,
		CpuOverCommitRatio:           p.CpuOverCommitRatio,
		CpuReservationPercentage:     p.CpuReservationPercentage,
		DedupSavingPercent:           p.DedupSavingPercent,
		ErasureCodingSavingPercent:   p.ErasureCodingSavingPercent,
		OverallSavingPercent:         p.OverallSavingPercent,
		RamOverCommitRatio:           p.RamOverCommitRatio,
		RamReservationPercentage:     p.RamReservationPercentage,
		ReplicationFactor:            p.ReplicationFactor,
		StorageReservationPercentage: p.StorageReservationPercentage,
	})
}

func NewDataStoreConfig() *DataStoreConfig {
	p := new(DataStoreConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.config.DataStoreConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	p.CompressionSavingPercent = new(float64)
	*p.CompressionSavingPercent = 25.81037
	p.CpuOverCommitRatio = new(float64)
	*p.CpuOverCommitRatio = 1.0
	p.CpuReservationPercentage = new(float64)
	*p.CpuReservationPercentage = 0.0
	p.DedupSavingPercent = new(float64)
	*p.DedupSavingPercent = 35.864
	p.ErasureCodingSavingPercent = new(float64)
	*p.ErasureCodingSavingPercent = 15.16939
	p.OverallSavingPercent = new(float64)
	*p.OverallSavingPercent = 59.64
	p.RamOverCommitRatio = new(float64)
	*p.RamOverCommitRatio = 1.0
	p.RamReservationPercentage = new(float64)
	*p.RamReservationPercentage = 0.0
	p.StorageReservationPercentage = new(float64)
	*p.StorageReservationPercentage = 0.0

	return p
}

/*
REST response for all response codes in API path /aiops/v4.0/config/ignore-windows/{extId} Delete operation
*/
type DeleteIgnoreWindowApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteIgnoreWindowApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDeleteIgnoreWindowApiResponse() *DeleteIgnoreWindowApiResponse {
	p := new(DeleteIgnoreWindowApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.config.DeleteIgnoreWindowApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteIgnoreWindowApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteIgnoreWindowApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteIgnoreWindowApiResponseData()
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
REST response for all response codes in API path /aiops/v4.0/config/scenarios/{extId} Delete operation
*/
type DeleteScenarioApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteScenarioApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDeleteScenarioApiResponse() *DeleteScenarioApiResponse {
	p := new(DeleteScenarioApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.config.DeleteScenarioApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteScenarioApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteScenarioApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteScenarioApiResponseData()
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
REST response for all response codes in API path /aiops/v4.0/config/simulations/{extId} Delete operation
*/
type DeleteSimulationApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteSimulationApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDeleteSimulationApiResponse() *DeleteSimulationApiResponse {
	p := new(DeleteSimulationApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.config.DeleteSimulationApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteSimulationApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteSimulationApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteSimulationApiResponseData()
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

type EntityDescriptor struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Readable name for the entity.
	*/
	DisplayName *string `json:"displayName,omitempty"`
	/*
	  Entity type of the data supported for a given source. For example VM, cluster etc.
	*/
	EntityType *string `json:"entityType,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import3.ApiLink `json:"links,omitempty"`
	/*
	  Metrics data for the entity.
	*/
	Metrics []MetricDescriptor `json:"metrics,omitempty"`
	/*
	  Parent entity types for the given entity.
	*/
	Parents []EntityDescriptor `json:"parents,omitempty"`
	/*
	  Source name for the vendors. For example 'Nutanix'.
	*/
	Source *string `json:"source,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewEntityDescriptor() *EntityDescriptor {
	p := new(EntityDescriptor)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.config.EntityDescriptor"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /aiops/v4.0/config/sources/{sourceExtId}/entity-descriptors Get operation
*/
type EntityDescriptorListApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfEntityDescriptorListApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewEntityDescriptorListApiResponse() *EntityDescriptorListApiResponse {
	p := new(EntityDescriptorListApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.config.EntityDescriptorListApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *EntityDescriptorListApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *EntityDescriptorListApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfEntityDescriptorListApiResponseData()
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

type EntityType struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Entity type of the data supported for a given source. For example VM, cluster etc.
	*/
	EntityTypeName *string `json:"entityTypeName,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import3.ApiLink `json:"links,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewEntityType() *EntityType {
	p := new(EntityType)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.config.EntityType"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /aiops/v4.0/config/sources/{sourceExtId}/entity-types Get operation
*/
type EntityTypeListApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfEntityTypeListApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewEntityTypeListApiResponse() *EntityTypeListApiResponse {
	p := new(EntityTypeListApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.config.EntityTypeListApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *EntityTypeListApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *EntityTypeListApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfEntityTypeListApiResponseData()
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
REST response for all response codes in API path /aiops/v4.0/config/scenarios/{extId}/$actions/generate-recommendation Post operation
*/
type GenerateRecommendationApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGenerateRecommendationApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGenerateRecommendationApiResponse() *GenerateRecommendationApiResponse {
	p := new(GenerateRecommendationApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.config.GenerateRecommendationApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GenerateRecommendationApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GenerateRecommendationApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGenerateRecommendationApiResponseData()
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
REST response for all response codes in API path /aiops/v4.0/config/scenarios/{extId}/$actions/generate-report Post operation
*/
type GenerateReportApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGenerateReportApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGenerateReportApiResponse() *GenerateReportApiResponse {
	p := new(GenerateReportApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.config.GenerateReportApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GenerateReportApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GenerateReportApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGenerateReportApiResponseData()
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
REST response for all response codes in API path /aiops/v4.0/config/scenarios/{extId}/$actions/generate-runway Post operation
*/
type GenerateRunwayApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGenerateRunwayApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGenerateRunwayApiResponse() *GenerateRunwayApiResponse {
	p := new(GenerateRunwayApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.config.GenerateRunwayApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GenerateRunwayApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GenerateRunwayApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGenerateRunwayApiResponseData()
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
REST response for all response codes in API path /aiops/v4.0/config/scenarios/{extId} Get operation
*/
type GetScenarioApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetScenarioApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetScenarioApiResponse() *GetScenarioApiResponse {
	p := new(GetScenarioApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.config.GetScenarioApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetScenarioApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetScenarioApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetScenarioApiResponseData()
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
REST response for all response codes in API path /aiops/v4.0/config/scenarios/{scenarioExtId}/reports Get operation
*/
type GetScenarioReportApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetScenarioReportApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetScenarioReportApiResponse() *GetScenarioReportApiResponse {
	p := new(GetScenarioReportApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.config.GetScenarioReportApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetScenarioReportApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetScenarioReportApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetScenarioReportApiResponseData()
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
REST response for all response codes in API path /aiops/v4.0/config/simulations/{extId} Get operation
*/
type GetSimulationApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetSimulationApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetSimulationApiResponse() *GetSimulationApiResponse {
	p := new(GetSimulationApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.config.GetSimulationApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetSimulationApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetSimulationApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetSimulationApiResponseData()
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
Definition of an ignore window.
*/
type IgnoreWindow struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	ClusterType *import4.ClusterType `json:"clusterType"`
	/*
	  UUID of the cluster to which the ignore window is to be applied.
	*/
	ClusterUuid *string `json:"clusterUuid"`
	/*
	  End timestamp of the ignore window. It can be a maximum of one year in the past from the current date but should be at least one day after the startDateTime.
	*/
	EndDateTime *time.Time `json:"endDateTime"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import3.ApiLink `json:"links,omitempty"`
	/*
	  Name of an ignore window period.
	*/
	Name *string `json:"name"`
	/*
	  Start timestamp of the ignore window. It can be maximum of one year in the past from the current date.
	*/
	StartDateTime *time.Time `json:"startDateTime"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *IgnoreWindow) MarshalJSON() ([]byte, error) {
	type IgnoreWindowProxy IgnoreWindow
	return json.Marshal(struct {
		*IgnoreWindowProxy
		ClusterType   *import4.ClusterType `json:"clusterType,omitempty"`
		ClusterUuid   *string              `json:"clusterUuid,omitempty"`
		EndDateTime   *time.Time           `json:"endDateTime,omitempty"`
		Name          *string              `json:"name,omitempty"`
		StartDateTime *time.Time           `json:"startDateTime,omitempty"`
	}{
		IgnoreWindowProxy: (*IgnoreWindowProxy)(p),
		ClusterType:       p.ClusterType,
		ClusterUuid:       p.ClusterUuid,
		EndDateTime:       p.EndDateTime,
		Name:              p.Name,
		StartDateTime:     p.StartDateTime,
	})
}

func NewIgnoreWindow() *IgnoreWindow {
	p := new(IgnoreWindow)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.config.IgnoreWindow"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type IgnoreWindowProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	ClusterType *import4.ClusterType `json:"clusterType"`
	/*
	  UUID of the cluster to which the ignore window is to be applied.
	*/
	ClusterUuid *string `json:"clusterUuid"`
	/*
	  End timestamp of the ignore window. It can be a maximum of one year in the past from the current date but should be at least one day after the startDateTime.
	*/
	EndDateTime *time.Time `json:"endDateTime"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import3.ApiLink `json:"links,omitempty"`
	/*
	  Name of an ignore window period.
	*/
	Name *string `json:"name"`
	/*
	  Start timestamp of the ignore window. It can be maximum of one year in the past from the current date.
	*/
	StartDateTime *time.Time `json:"startDateTime"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *IgnoreWindowProjection) MarshalJSON() ([]byte, error) {
	type IgnoreWindowProjectionProxy IgnoreWindowProjection
	return json.Marshal(struct {
		*IgnoreWindowProjectionProxy
		ClusterType   *import4.ClusterType `json:"clusterType,omitempty"`
		ClusterUuid   *string              `json:"clusterUuid,omitempty"`
		EndDateTime   *time.Time           `json:"endDateTime,omitempty"`
		Name          *string              `json:"name,omitempty"`
		StartDateTime *time.Time           `json:"startDateTime,omitempty"`
	}{
		IgnoreWindowProjectionProxy: (*IgnoreWindowProjectionProxy)(p),
		ClusterType:                 p.ClusterType,
		ClusterUuid:                 p.ClusterUuid,
		EndDateTime:                 p.EndDateTime,
		Name:                        p.Name,
		StartDateTime:               p.StartDateTime,
	})
}

func NewIgnoreWindowProjection() *IgnoreWindowProjection {
	p := new(IgnoreWindowProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.config.IgnoreWindowProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /aiops/v4.0/config/scenarios Get operation
*/
type ListScenariosApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListScenariosApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListScenariosApiResponse() *ListScenariosApiResponse {
	p := new(ListScenariosApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.config.ListScenariosApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListScenariosApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListScenariosApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListScenariosApiResponseData()
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
REST response for all response codes in API path /aiops/v4.0/config/simulations Get operation
*/
type ListSimulationsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListSimulationsApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListSimulationsApiResponse() *ListSimulationsApiResponse {
	p := new(ListSimulationsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.config.ListSimulationsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListSimulationsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListSimulationsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListSimulationsApiResponseData()
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

type MetricDescriptor struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Additional properties for metric descriptor definition. Like anomaly metric, alert metric and search UI metric.
	*/
	AdditionalProperties []import5.KVPair `json:"additionalProperties,omitempty"`
	/*
	  Default value of the metric.
	*/
	DefaultValue *string `json:"defaultValue,omitempty"`
	/*
	  Readable name for the entity.
	*/
	DisplayName *string `json:"displayName,omitempty"`

	DownsamplingOperator *import6.DownSamplingOperator `json:"downsamplingOperator,omitempty"`
	/*
	  Indicates whether it is an attribute or not.
	*/
	IsAttribute *bool `json:"isAttribute,omitempty"`
	/*
	  Indicator to specify whether the attribute should be persisted as time series data or not.
	*/
	IsAttributePersistedAsTimeSeries *bool `json:"isAttributePersistedAsTimeSeries,omitempty"`
	/*
	  Name of the metric.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  The interval value is used to resample the queried time series data by using down_sampling_operator operator. The default is 86400 seconds.
	*/
	SamplingIntervalSecs *int `json:"samplingIntervalSecs,omitempty"`
	/*
	  Unit for the metric.
	*/
	Unit *string `json:"unit,omitempty"`

	ValueRange *ValueRange `json:"valueRange,omitempty"`

	ValueType *ValueType `json:"valueType,omitempty"`
}

func NewMetricDescriptor() *MetricDescriptor {
	p := new(MetricDescriptor)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.config.MetricDescriptor"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Enum describing Xen provision type for Microsoft.
*/
type MicrosoftXenProvisionType int

const (
	MICROSOFTXENPROVISIONTYPE_UNKNOWN  MicrosoftXenProvisionType = 0
	MICROSOFTXENPROVISIONTYPE_REDACTED MicrosoftXenProvisionType = 1
	MICROSOFTXENPROVISIONTYPE_VM_CLONE MicrosoftXenProvisionType = 2
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *MicrosoftXenProvisionType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"VM_CLONE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e MicrosoftXenProvisionType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"VM_CLONE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *MicrosoftXenProvisionType) index(name string) MicrosoftXenProvisionType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"VM_CLONE",
	}
	for idx := range names {
		if names[idx] == name {
			return MicrosoftXenProvisionType(idx)
		}
	}
	return MICROSOFTXENPROVISIONTYPE_UNKNOWN
}

func (e *MicrosoftXenProvisionType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for MicrosoftXenProvisionType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *MicrosoftXenProvisionType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e MicrosoftXenProvisionType) Ref() *MicrosoftXenProvisionType {
	return &e
}

/*
Xen workload with Microsoft as a vendor.
*/
type MicrosoftXenWorkload struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Size in GB for MCS different disk per VM.
	*/
	McsDiffSizeGb *int `json:"mcsDiffSizeGb"`

	OperatingSystem *XenOperatingSystem `json:"operatingSystem"`

	ProvisionType *MicrosoftXenProvisionType `json:"provisionType"`
	/*
	  Size of PVS cache write per VM in GB.
	*/
	PvsWriteCacheSizeGb *int `json:"pvsWriteCacheSizeGb"`
	/*
	  Space consumed by each Xen Server Image.
	*/
	SystemDataGb *int `json:"systemDataGb"`
	/*
	  Number of users for Xen workload.
	*/
	UserCount *int `json:"userCount"`
	/*
	  Size for the per user profile data in MB.
	*/
	UserProfileDataMb *int `json:"userProfileDataMb"`

	Vendor *MicrosoftXenWorkloadVendor `json:"vendor"`
}

func (p *MicrosoftXenWorkload) MarshalJSON() ([]byte, error) {
	type MicrosoftXenWorkloadProxy MicrosoftXenWorkload
	return json.Marshal(struct {
		*MicrosoftXenWorkloadProxy
		McsDiffSizeGb       *int                        `json:"mcsDiffSizeGb,omitempty"`
		OperatingSystem     *XenOperatingSystem         `json:"operatingSystem,omitempty"`
		ProvisionType       *MicrosoftXenProvisionType  `json:"provisionType,omitempty"`
		PvsWriteCacheSizeGb *int                        `json:"pvsWriteCacheSizeGb,omitempty"`
		SystemDataGb        *int                        `json:"systemDataGb,omitempty"`
		UserCount           *int                        `json:"userCount,omitempty"`
		UserProfileDataMb   *int                        `json:"userProfileDataMb,omitempty"`
		Vendor              *MicrosoftXenWorkloadVendor `json:"vendor,omitempty"`
	}{
		MicrosoftXenWorkloadProxy: (*MicrosoftXenWorkloadProxy)(p),
		McsDiffSizeGb:             p.McsDiffSizeGb,
		OperatingSystem:           p.OperatingSystem,
		ProvisionType:             p.ProvisionType,
		PvsWriteCacheSizeGb:       p.PvsWriteCacheSizeGb,
		SystemDataGb:              p.SystemDataGb,
		UserCount:                 p.UserCount,
		UserProfileDataMb:         p.UserProfileDataMb,
		Vendor:                    p.Vendor,
	})
}

func NewMicrosoftXenWorkload() *MicrosoftXenWorkload {
	p := new(MicrosoftXenWorkload)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.config.MicrosoftXenWorkload"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Enum describing Xen vendor for Microsoft.
*/
type MicrosoftXenWorkloadVendor int

const (
	MICROSOFTXENWORKLOADVENDOR_UNKNOWN        MicrosoftXenWorkloadVendor = 0
	MICROSOFTXENWORKLOADVENDOR_REDACTED       MicrosoftXenWorkloadVendor = 1
	MICROSOFTXENWORKLOADVENDOR_MICROSOFT_RDSH MicrosoftXenWorkloadVendor = 2
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *MicrosoftXenWorkloadVendor) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"MICROSOFT_RDSH",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e MicrosoftXenWorkloadVendor) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"MICROSOFT_RDSH",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *MicrosoftXenWorkloadVendor) index(name string) MicrosoftXenWorkloadVendor {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"MICROSOFT_RDSH",
	}
	for idx := range names {
		if names[idx] == name {
			return MicrosoftXenWorkloadVendor(idx)
		}
	}
	return MICROSOFTXENWORKLOADVENDOR_UNKNOWN
}

func (e *MicrosoftXenWorkloadVendor) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for MicrosoftXenWorkloadVendor:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *MicrosoftXenWorkloadVendor) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e MicrosoftXenWorkloadVendor) Ref() *MicrosoftXenWorkloadVendor {
	return &e
}

/*
Node specification.
*/
type NodeConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Flag to indicate if node is to be taken into account while performing capacity scenario analysis.
	*/
	IsEnabled *bool `json:"isEnabled"`
	/*
	  Model name of a node.
	*/
	Model *string `json:"model"`
	/*
	  Number of nodes in a model.
	*/
	NodeCount *int `json:"nodeCount"`

	NodeResourceCapacity *ResourceCapacity `json:"nodeResourceCapacity,omitempty"`

	NodeSource *NodeSource `json:"nodeSource"`
	/*
	  Recommended time when a node should be live in a cluster.
	*/
	TargetOnlineTime *time.Time `json:"targetOnlineTime,omitempty"`
}

func (p *NodeConfig) MarshalJSON() ([]byte, error) {
	type NodeConfigProxy NodeConfig
	return json.Marshal(struct {
		*NodeConfigProxy
		IsEnabled  *bool       `json:"isEnabled,omitempty"`
		Model      *string     `json:"model,omitempty"`
		NodeCount  *int        `json:"nodeCount,omitempty"`
		NodeSource *NodeSource `json:"nodeSource,omitempty"`
	}{
		NodeConfigProxy: (*NodeConfigProxy)(p),
		IsEnabled:       p.IsEnabled,
		Model:           p.Model,
		NodeCount:       p.NodeCount,
		NodeSource:      p.NodeSource,
	})
}

func NewNodeConfig() *NodeConfig {
	p := new(NodeConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.config.NodeConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	p.IsEnabled = new(bool)
	*p.IsEnabled = false

	return p
}

/*
Source of the node added.
*/
type NodeSource int

const (
	NODESOURCE_UNKNOWN     NodeSource = 0
	NODESOURCE_REDACTED    NodeSource = 1
	NODESOURCE_EXISTING    NodeSource = 2
	NODESOURCE_USER_ADDED  NodeSource = 3
	NODESOURCE_RECOMMENDED NodeSource = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *NodeSource) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"EXISTING",
		"USER_ADDED",
		"RECOMMENDED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e NodeSource) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"EXISTING",
		"USER_ADDED",
		"RECOMMENDED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *NodeSource) index(name string) NodeSource {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"EXISTING",
		"USER_ADDED",
		"RECOMMENDED",
	}
	for idx := range names {
		if names[idx] == name {
			return NodeSource(idx)
		}
	}
	return NODESOURCE_UNKNOWN
}

func (e *NodeSource) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for NodeSource:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *NodeSource) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e NodeSource) Ref() *NodeSource {
	return &e
}

/*
Policy type Enum values. The permissible values are inefficiency_exclusion or anomaly_exclusion.
*/
type PolicyType int

const (
	POLICYTYPE_UNKNOWN                PolicyType = 0
	POLICYTYPE_REDACTED               PolicyType = 1
	POLICYTYPE_INEFFICIENCY_EXCLUSION PolicyType = 2
	POLICYTYPE_ANOMALY_EXCLUSION      PolicyType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *PolicyType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"INEFFICIENCY_EXCLUSION",
		"ANOMALY_EXCLUSION",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e PolicyType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"INEFFICIENCY_EXCLUSION",
		"ANOMALY_EXCLUSION",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *PolicyType) index(name string) PolicyType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"INEFFICIENCY_EXCLUSION",
		"ANOMALY_EXCLUSION",
	}
	for idx := range names {
		if names[idx] == name {
			return PolicyType(idx)
		}
	}
	return POLICYTYPE_UNKNOWN
}

func (e *PolicyType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for PolicyType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *PolicyType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e PolicyType) Ref() *PolicyType {
	return &e
}

/*
Replication factor of the cluster resources.
*/
type ReplicationFactor int

const (
	REPLICATIONFACTOR_UNKNOWN  ReplicationFactor = 0
	REPLICATIONFACTOR_REDACTED ReplicationFactor = 1
	REPLICATIONFACTOR_RF_2     ReplicationFactor = 2
	REPLICATIONFACTOR_RF_3     ReplicationFactor = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ReplicationFactor) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"RF_2",
		"RF_3",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ReplicationFactor) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"RF_2",
		"RF_3",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ReplicationFactor) index(name string) ReplicationFactor {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"RF_2",
		"RF_3",
	}
	for idx := range names {
		if names[idx] == name {
			return ReplicationFactor(idx)
		}
	}
	return REPLICATIONFACTOR_UNKNOWN
}

func (e *ReplicationFactor) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ReplicationFactor:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ReplicationFactor) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ReplicationFactor) Ref() *ReplicationFactor {
	return &e
}

/*
Resource specification.
*/
type ResourceCapacity struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  CPU capacity in GHz.
	*/
	CpuGhz *float64 `json:"cpuGhz"`
	/*
	  HDD capacity in GB.
	*/
	HddGb *float64 `json:"hddGb"`
	/*
	  NVMe capacity in GB.
	*/
	NvmeGb *float64 `json:"nvmeGb"`
	/*
	  RAM capacity in GB.
	*/
	RamGb *float64 `json:"ramGb"`
	/*
	  SSD capacity in GB.
	*/
	SsdGb *float64 `json:"ssdGb"`
}

func (p *ResourceCapacity) MarshalJSON() ([]byte, error) {
	type ResourceCapacityProxy ResourceCapacity
	return json.Marshal(struct {
		*ResourceCapacityProxy
		CpuGhz *float64 `json:"cpuGhz,omitempty"`
		HddGb  *float64 `json:"hddGb,omitempty"`
		NvmeGb *float64 `json:"nvmeGb,omitempty"`
		RamGb  *float64 `json:"ramGb,omitempty"`
		SsdGb  *float64 `json:"ssdGb,omitempty"`
	}{
		ResourceCapacityProxy: (*ResourceCapacityProxy)(p),
		CpuGhz:                p.CpuGhz,
		HddGb:                 p.HddGb,
		NvmeGb:                p.NvmeGb,
		RamGb:                 p.RamGb,
		SsdGb:                 p.SsdGb,
	})
}

func NewResourceCapacity() *ResourceCapacity {
	p := new(ResourceCapacity)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.config.ResourceCapacity"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Runway in a scenario.
*/
type Runway struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Number of days the CPU can sustain the added and existing workload.
	*/
	CpuRunwayDays *int `json:"cpuRunwayDays"`
	/*
	  Number of days the memory can sustain the added and existing workload.
	*/
	MemoryRunwayDays *int `json:"memoryRunwayDays"`
	/*
	  Number of days the cluster will be able to sustain the existing and added workloads.
	*/
	MinimumRunwayDays *int `json:"minimumRunwayDays"`
	/*
	  The starting date of the usage, capacity, and effective capacity forecasts. It is used as the start-time in stats API to query scenario statistics.
	*/
	RunwayStartDate *time.Time `json:"runwayStartDate,omitempty"`
	/*
	  Number of days the storage can sustain the added and existing workload.
	*/
	StorageRunwayDays *int `json:"storageRunwayDays"`
}

func (p *Runway) MarshalJSON() ([]byte, error) {
	type RunwayProxy Runway
	return json.Marshal(struct {
		*RunwayProxy
		CpuRunwayDays     *int   `json:"cpuRunwayDays,omitempty"`
		MemoryRunwayDays  *int   `json:"memoryRunwayDays,omitempty"`
		MinimumRunwayDays *int   `json:"minimumRunwayDays,omitempty"`
		StorageRunwayDays *int   `json:"storageRunwayDays,omitempty"`
		RunwayStartDate   string `json:"runwayStartDate,omitempty"`
	}{
		RunwayProxy:       (*RunwayProxy)(p),
		CpuRunwayDays:     p.CpuRunwayDays,
		MemoryRunwayDays:  p.MemoryRunwayDays,
		MinimumRunwayDays: p.MinimumRunwayDays,
		StorageRunwayDays: p.StorageRunwayDays,
		RunwayStartDate: func() string {
			if p.RunwayStartDate != nil {
				return p.RunwayStartDate.Format("2006-01-02")
			} else {
				return ""
			}
		}(),
	})
}

func (p *Runway) UnmarshalJSON(b []byte) error {
	type CustomRunway struct {
		ObjectType_       *string                `json:"$objectType,omitempty"`
		Reserved_         map[string]interface{} `json:"$reserved,omitempty"`
		UnknownFields_    map[string]interface{} `json:"$unknownFields,omitempty"`
		CpuRunwayDays     *int                   `json:"cpuRunwayDays"`
		MemoryRunwayDays  *int                   `json:"memoryRunwayDays"`
		MinimumRunwayDays *int                   `json:"minimumRunwayDays"`
		RunwayStartDate   string                 `json:"runwayStartDate,omitempty"`
		StorageRunwayDays *int                   `json:"storageRunwayDays"`
	}

	var customRunway CustomRunway
	err := json.Unmarshal(b, &customRunway)
	if err != nil {
		return err
	}

	p.ObjectType_ = customRunway.ObjectType_
	p.Reserved_ = customRunway.Reserved_
	p.UnknownFields_ = customRunway.UnknownFields_
	p.CpuRunwayDays = customRunway.CpuRunwayDays
	p.MemoryRunwayDays = customRunway.MemoryRunwayDays
	p.MinimumRunwayDays = customRunway.MinimumRunwayDays
	// Custom date parsing logic for Date field
	if customRunway.RunwayStartDate != "" {
		parsedRunwayStartDate, err := time.Parse("2006-01-02", customRunway.RunwayStartDate)
		if err != nil {
			return errors.New(fmt.Sprintf("Unable to unmarshal field RunwayStartDate in struct Runway: %s", err))
		}
		p.RunwayStartDate = &parsedRunwayStartDate
	}
	p.StorageRunwayDays = customRunway.StorageRunwayDays

	return nil
}

func NewRunway() *Runway {
	p := new(Runway)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.config.Runway"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Required details to trigger a calculate runway task.
*/
type RunwayParams struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	ClusterType *import4.ClusterType `json:"clusterType"`
	/*
	  UUID of the cluster to which the ignore window is to be applied.
	*/
	ClusterUuid *string `json:"clusterUuid"`
}

func (p *RunwayParams) MarshalJSON() ([]byte, error) {
	type RunwayParamsProxy RunwayParams
	return json.Marshal(struct {
		*RunwayParamsProxy
		ClusterType *import4.ClusterType `json:"clusterType,omitempty"`
		ClusterUuid *string              `json:"clusterUuid,omitempty"`
	}{
		RunwayParamsProxy: (*RunwayParamsProxy)(p),
		ClusterType:       p.ClusterType,
		ClusterUuid:       p.ClusterUuid,
	})
}

func NewRunwayParams() *RunwayParams {
	p := new(RunwayParams)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.config.RunwayParams"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type RunwayParamsProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	ClusterType *import4.ClusterType `json:"clusterType"`
	/*
	  UUID of the cluster to which the ignore window is to be applied.
	*/
	ClusterUuid *string `json:"clusterUuid"`
}

func (p *RunwayParamsProjection) MarshalJSON() ([]byte, error) {
	type RunwayParamsProjectionProxy RunwayParamsProjection
	return json.Marshal(struct {
		*RunwayParamsProjectionProxy
		ClusterType *import4.ClusterType `json:"clusterType,omitempty"`
		ClusterUuid *string              `json:"clusterUuid,omitempty"`
	}{
		RunwayParamsProjectionProxy: (*RunwayParamsProjectionProxy)(p),
		ClusterType:                 p.ClusterType,
		ClusterUuid:                 p.ClusterUuid,
	})
}

func NewRunwayParamsProjection() *RunwayParamsProjection {
	p := new(RunwayParamsProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.config.RunwayParamsProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
WhatIf Scenario.
*/
type Scenario struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	ClusterConfig *ClusterConfig `json:"clusterConfig,omitempty"`
	/*
	  UUID of the cluster for which Whatif analysis is being performed.
	*/
	ClusterExtId *string `json:"clusterExtId,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import3.ApiLink `json:"links,omitempty"`
	/*
	  Name of capacity planning scenario.
	*/
	Name *string `json:"name,omitempty"`

	Runway *Runway `json:"runway,omitempty"`
	/*
	  Number of days a cluster is expected to sustain the workload in a capacity planning scenario.
	*/
	TargetRunwayDays *int `json:"targetRunwayDays,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  Timestamp when the capacity planning scenario is updated. The value should be in ISO-8601 format. For example, 2022-02-20T00:00:00.458Z.
	*/
	UpdatedTime *time.Time `json:"updatedTime,omitempty"`
	/*
	  A list of allowed vendors whose model can be requested to sustain the workload in a capacity planning scenario.
	*/
	Vendors []Vendor `json:"vendors,omitempty"`
	/*
	  List of workloads for which runway analysis is being done. It can be considered an additional resource requirement required to run a specific use case. For example, a SQL server.
	*/
	Workloads []Workload `json:"workloads,omitempty"`
}

func NewScenario() *Scenario {
	p := new(Scenario)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.config.Scenario"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type ScenarioProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	ClusterConfig *ClusterConfig `json:"clusterConfig,omitempty"`
	/*
	  UUID of the cluster for which Whatif analysis is being performed.
	*/
	ClusterExtId *string `json:"clusterExtId,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import3.ApiLink `json:"links,omitempty"`
	/*
	  Name of capacity planning scenario.
	*/
	Name *string `json:"name,omitempty"`

	Runway *Runway `json:"runway,omitempty"`
	/*
	  Number of days a cluster is expected to sustain the workload in a capacity planning scenario.
	*/
	TargetRunwayDays *int `json:"targetRunwayDays,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  Timestamp when the capacity planning scenario is updated. The value should be in ISO-8601 format. For example, 2022-02-20T00:00:00.458Z.
	*/
	UpdatedTime *time.Time `json:"updatedTime,omitempty"`
	/*
	  A list of allowed vendors whose model can be requested to sustain the workload in a capacity planning scenario.
	*/
	Vendors []Vendor `json:"vendors,omitempty"`
	/*
	  List of workloads for which runway analysis is being done. It can be considered an additional resource requirement required to run a specific use case. For example, a SQL server.
	*/
	Workloads []Workload `json:"workloads,omitempty"`
}

func NewScenarioProjection() *ScenarioProjection {
	p := new(ScenarioProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.config.ScenarioProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Simulated resource specification for a VM workload.
*/
type SimulatedVmResourceSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  HDD capacity in GB.
	*/
	HddGb *float64 `json:"hddGb"`
	/*
	  RAM capacity in GB.
	*/
	RamGb *float64 `json:"ramGb"`
	/*
	  SSD capacity in GB.
	*/
	SsdGb *float64 `json:"ssdGb,omitempty"`
	/*
	  Number of vCPUs.
	*/
	VcpuCount *int `json:"vcpuCount"`
}

func (p *SimulatedVmResourceSpec) MarshalJSON() ([]byte, error) {
	type SimulatedVmResourceSpecProxy SimulatedVmResourceSpec
	return json.Marshal(struct {
		*SimulatedVmResourceSpecProxy
		HddGb     *float64 `json:"hddGb,omitempty"`
		RamGb     *float64 `json:"ramGb,omitempty"`
		VcpuCount *int     `json:"vcpuCount,omitempty"`
	}{
		SimulatedVmResourceSpecProxy: (*SimulatedVmResourceSpecProxy)(p),
		HddGb:                        p.HddGb,
		RamGb:                        p.RamGb,
		VcpuCount:                    p.VcpuCount,
	})
}

func NewSimulatedVmResourceSpec() *SimulatedVmResourceSpec {
	p := new(SimulatedVmResourceSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.config.SimulatedVmResourceSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type SimulatedVmResourceSpecProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  HDD capacity in GB.
	*/
	HddGb *float64 `json:"hddGb"`
	/*
	  RAM capacity in GB.
	*/
	RamGb *float64 `json:"ramGb"`
	/*
	  SSD capacity in GB.
	*/
	SsdGb *float64 `json:"ssdGb,omitempty"`
	/*
	  Number of vCPUs.
	*/
	VcpuCount *int `json:"vcpuCount"`
}

func (p *SimulatedVmResourceSpecProjection) MarshalJSON() ([]byte, error) {
	type SimulatedVmResourceSpecProjectionProxy SimulatedVmResourceSpecProjection
	return json.Marshal(struct {
		*SimulatedVmResourceSpecProjectionProxy
		HddGb     *float64 `json:"hddGb,omitempty"`
		RamGb     *float64 `json:"ramGb,omitempty"`
		VcpuCount *int     `json:"vcpuCount,omitempty"`
	}{
		SimulatedVmResourceSpecProjectionProxy: (*SimulatedVmResourceSpecProjectionProxy)(p),
		HddGb:                                  p.HddGb,
		RamGb:                                  p.RamGb,
		VcpuCount:                              p.VcpuCount,
	})
}

func NewSimulatedVmResourceSpecProjection() *SimulatedVmResourceSpecProjection {
	p := new(SimulatedVmResourceSpecProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.config.SimulatedVmResourceSpecProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Simulation of a workload and its resource requirements.
*/
type Simulation struct {
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
	  Name of the resource used in a scenario.
	*/
	Name *string `json:"name,omitempty"`

	SimulationSpec *SimulatedVmResourceSpec `json:"simulationSpec,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewSimulation() *Simulation {
	p := new(Simulation)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.config.Simulation"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type SimulationProjection struct {
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
	  Name of the resource used in a scenario.
	*/
	Name *string `json:"name,omitempty"`

	SimulationSpec *SimulatedVmResourceSpec `json:"simulationSpec,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewSimulationProjection() *SimulationProjection {
	p := new(SimulationProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.config.SimulationProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type Source struct {
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
	  Source name for the vendors. For example 'Nutanix'.
	*/
	SourceName *string `json:"sourceName,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewSource() *Source {
	p := new(Source)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.config.Source"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /aiops/v4.0/config/sources Get operation
*/
type SourceListApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfSourceListApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewSourceListApiResponse() *SourceListApiResponse {
	p := new(SourceListApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.config.SourceListApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *SourceListApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *SourceListApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfSourceListApiResponseData()
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
Splunk workload description.
*/
type SplunkWorkload struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Number of days to retain cold data.
	*/
	ColdRetentionDays *int `json:"coldRetentionDays"`
	/*
	  Daily average index size in Splunk workload.
	*/
	DailyAverageIndexingRateGb *int `json:"dailyAverageIndexingRateGb"`
	/*
	  Number of days to retain hot data.
	*/
	HotRetentionDays *int `json:"hotRetentionDays"`
	/*
	  Number of users of Splunk workload.
	*/
	UserCount *int `json:"userCount"`
}

func (p *SplunkWorkload) MarshalJSON() ([]byte, error) {
	type SplunkWorkloadProxy SplunkWorkload
	return json.Marshal(struct {
		*SplunkWorkloadProxy
		ColdRetentionDays          *int `json:"coldRetentionDays,omitempty"`
		DailyAverageIndexingRateGb *int `json:"dailyAverageIndexingRateGb,omitempty"`
		HotRetentionDays           *int `json:"hotRetentionDays,omitempty"`
		UserCount                  *int `json:"userCount,omitempty"`
	}{
		SplunkWorkloadProxy:        (*SplunkWorkloadProxy)(p),
		ColdRetentionDays:          p.ColdRetentionDays,
		DailyAverageIndexingRateGb: p.DailyAverageIndexingRateGb,
		HotRetentionDays:           p.HotRetentionDays,
		UserCount:                  p.UserCount,
	})
}

func NewSplunkWorkload() *SplunkWorkload {
	p := new(SplunkWorkload)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.config.SplunkWorkload"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	p.ColdRetentionDays = new(int)
	*p.ColdRetentionDays = 60
	p.DailyAverageIndexingRateGb = new(int)
	*p.DailyAverageIndexingRateGb = 500
	p.HotRetentionDays = new(int)
	*p.HotRetentionDays = 7
	p.UserCount = new(int)
	*p.UserCount = 5

	return p
}

/*
Enum for SQL profile type.
*/
type SqlProfileType int

const (
	SQLPROFILETYPE_UNKNOWN  SqlProfileType = 0
	SQLPROFILETYPE_REDACTED SqlProfileType = 1
	SQLPROFILETYPE_SMALL    SqlProfileType = 2
	SQLPROFILETYPE_MEDIUM   SqlProfileType = 3
	SQLPROFILETYPE_LARGE    SqlProfileType = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *SqlProfileType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SMALL",
		"MEDIUM",
		"LARGE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e SqlProfileType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SMALL",
		"MEDIUM",
		"LARGE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *SqlProfileType) index(name string) SqlProfileType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SMALL",
		"MEDIUM",
		"LARGE",
	}
	for idx := range names {
		if names[idx] == name {
			return SqlProfileType(idx)
		}
	}
	return SQLPROFILETYPE_UNKNOWN
}

func (e *SqlProfileType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for SqlProfileType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *SqlProfileType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e SqlProfileType) Ref() *SqlProfileType {
	return &e
}

/*
Enum for SQL transaction type.
*/
type SqlTransactionType int

const (
	SQLTRANSACTIONTYPE_UNKNOWN  SqlTransactionType = 0
	SQLTRANSACTIONTYPE_REDACTED SqlTransactionType = 1
	SQLTRANSACTIONTYPE_OLAP     SqlTransactionType = 2
	SQLTRANSACTIONTYPE_OLTP     SqlTransactionType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *SqlTransactionType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"OLAP",
		"OLTP",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e SqlTransactionType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"OLAP",
		"OLTP",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *SqlTransactionType) index(name string) SqlTransactionType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"OLAP",
		"OLTP",
	}
	for idx := range names {
		if names[idx] == name {
			return SqlTransactionType(idx)
		}
	}
	return SQLTRANSACTIONTYPE_UNKNOWN
}

func (e *SqlTransactionType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for SqlTransactionType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *SqlTransactionType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e SqlTransactionType) Ref() *SqlTransactionType {
	return &e
}

/*
SQL workload description.
*/
type SqlWorkload struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Number of SQL instances.
	*/
	DbCount *int `json:"dbCount"`
	/*
	  Whether the SQL instance is critical to business or not.
	*/
	IsBusinessCritical *bool `json:"isBusinessCritical"`

	ProfileType *SqlProfileType `json:"profileType"`

	TransactionType *SqlTransactionType `json:"transactionType"`
}

func (p *SqlWorkload) MarshalJSON() ([]byte, error) {
	type SqlWorkloadProxy SqlWorkload
	return json.Marshal(struct {
		*SqlWorkloadProxy
		DbCount            *int                `json:"dbCount,omitempty"`
		IsBusinessCritical *bool               `json:"isBusinessCritical,omitempty"`
		ProfileType        *SqlProfileType     `json:"profileType,omitempty"`
		TransactionType    *SqlTransactionType `json:"transactionType,omitempty"`
	}{
		SqlWorkloadProxy:   (*SqlWorkloadProxy)(p),
		DbCount:            p.DbCount,
		IsBusinessCritical: p.IsBusinessCritical,
		ProfileType:        p.ProfileType,
		TransactionType:    p.TransactionType,
	})
}

func NewSqlWorkload() *SqlWorkload {
	p := new(SqlWorkload)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.config.SqlWorkload"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	p.IsBusinessCritical = new(bool)
	*p.IsBusinessCritical = false

	return p
}

/*
REST response for all response codes in API path /aiops/v4.0/config/ignore-windows/{extId} Put operation
*/
type UpdateIgnoreWindowApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateIgnoreWindowApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewUpdateIgnoreWindowApiResponse() *UpdateIgnoreWindowApiResponse {
	p := new(UpdateIgnoreWindowApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.config.UpdateIgnoreWindowApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdateIgnoreWindowApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdateIgnoreWindowApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdateIgnoreWindowApiResponseData()
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
REST response for all response codes in API path /aiops/v4.0/config/xfit-policies/{extId} Put operation
*/
type UpdatePolicyTaskApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdatePolicyTaskApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewUpdatePolicyTaskApiResponse() *UpdatePolicyTaskApiResponse {
	p := new(UpdatePolicyTaskApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.config.UpdatePolicyTaskApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdatePolicyTaskApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdatePolicyTaskApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdatePolicyTaskApiResponseData()
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
REST response for all response codes in API path /aiops/v4.0/config/scenarios/{extId} Put operation
*/
type UpdateScenarioApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateScenarioApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewUpdateScenarioApiResponse() *UpdateScenarioApiResponse {
	p := new(UpdateScenarioApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.config.UpdateScenarioApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdateScenarioApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdateScenarioApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdateScenarioApiResponseData()
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
REST response for all response codes in API path /aiops/v4.0/config/simulations/{extId} Put operation
*/
type UpdateSimulationApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateSimulationApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewUpdateSimulationApiResponse() *UpdateSimulationApiResponse {
	p := new(UpdateSimulationApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.config.UpdateSimulationApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdateSimulationApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdateSimulationApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdateSimulationApiResponseData()
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
Range of the values for the metric.
*/
type ValueRange struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Maximum value in the value range.
	*/
	Max *float64 `json:"max,omitempty"`
	/*
	  Minimum value in the value range.
	*/
	Min *float64 `json:"min,omitempty"`
}

func NewValueRange() *ValueRange {
	p := new(ValueRange)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.config.ValueRange"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Data type of the values for the metric.
*/
type ValueType int

const (
	VALUETYPE_UNKNOWN     ValueType = 0
	VALUETYPE_REDACTED    ValueType = 1
	VALUETYPE_BOOL        ValueType = 2
	VALUETYPE_INT         ValueType = 3
	VALUETYPE_DOUBLE      ValueType = 4
	VALUETYPE_STRING      ValueType = 5
	VALUETYPE_BOOL_LIST   ValueType = 6
	VALUETYPE_INT_LIST    ValueType = 7
	VALUETYPE_DOUBLE_LIST ValueType = 8
	VALUETYPE_STRING_LIST ValueType = 9
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ValueType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"BOOL",
		"INT",
		"DOUBLE",
		"STRING",
		"BOOL_LIST",
		"INT_LIST",
		"DOUBLE_LIST",
		"STRING_LIST",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ValueType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"BOOL",
		"INT",
		"DOUBLE",
		"STRING",
		"BOOL_LIST",
		"INT_LIST",
		"DOUBLE_LIST",
		"STRING_LIST",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ValueType) index(name string) ValueType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"BOOL",
		"INT",
		"DOUBLE",
		"STRING",
		"BOOL_LIST",
		"INT_LIST",
		"DOUBLE_LIST",
		"STRING_LIST",
	}
	for idx := range names {
		if names[idx] == name {
			return ValueType(idx)
		}
	}
	return VALUETYPE_UNKNOWN
}

func (e *ValueType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ValueType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ValueType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ValueType) Ref() *ValueType {
	return &e
}

/*
Enum describing VDI provision types.
*/
type VdiProvisionType int

const (
	VDIPROVISIONTYPE_UNKNOWN                   VdiProvisionType = 0
	VDIPROVISIONTYPE_REDACTED                  VdiProvisionType = 1
	VDIPROVISIONTYPE_FULL_CLONES               VdiProvisionType = 2
	VDIPROVISIONTYPE_V2V_P2V                   VdiProvisionType = 3
	VDIPROVISIONTYPE_PROVISIONING_SERVICES     VdiProvisionType = 4
	VDIPROVISIONTYPE_MACHINE_CREATION_SERVICES VdiProvisionType = 5
	VDIPROVISIONTYPE_LINKED_CLONES             VdiProvisionType = 6
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *VdiProvisionType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"FULL_CLONES",
		"V2V_P2V",
		"PROVISIONING_SERVICES",
		"MACHINE_CREATION_SERVICES",
		"LINKED_CLONES",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e VdiProvisionType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"FULL_CLONES",
		"V2V_P2V",
		"PROVISIONING_SERVICES",
		"MACHINE_CREATION_SERVICES",
		"LINKED_CLONES",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *VdiProvisionType) index(name string) VdiProvisionType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"FULL_CLONES",
		"V2V_P2V",
		"PROVISIONING_SERVICES",
		"MACHINE_CREATION_SERVICES",
		"LINKED_CLONES",
	}
	for idx := range names {
		if names[idx] == name {
			return VdiProvisionType(idx)
		}
	}
	return VDIPROVISIONTYPE_UNKNOWN
}

func (e *VdiProvisionType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for VdiProvisionType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *VdiProvisionType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e VdiProvisionType) Ref() *VdiProvisionType {
	return &e
}

/*
Enum describing VDI user types.
*/
type VdiUserType int

const (
	VDIUSERTYPE_UNKNOWN          VdiUserType = 0
	VDIUSERTYPE_REDACTED         VdiUserType = 1
	VDIUSERTYPE_TASK_WORKER      VdiUserType = 2
	VDIUSERTYPE_KNOWLEDGE_WORKER VdiUserType = 3
	VDIUSERTYPE_POWER_USER       VdiUserType = 4
	VDIUSERTYPE_DEVELOPER        VdiUserType = 5
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *VdiUserType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"TASK_WORKER",
		"KNOWLEDGE_WORKER",
		"POWER_USER",
		"DEVELOPER",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e VdiUserType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"TASK_WORKER",
		"KNOWLEDGE_WORKER",
		"POWER_USER",
		"DEVELOPER",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *VdiUserType) index(name string) VdiUserType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"TASK_WORKER",
		"KNOWLEDGE_WORKER",
		"POWER_USER",
		"DEVELOPER",
	}
	for idx := range names {
		if names[idx] == name {
			return VdiUserType(idx)
		}
	}
	return VDIUSERTYPE_UNKNOWN
}

func (e *VdiUserType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for VdiUserType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *VdiUserType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e VdiUserType) Ref() *VdiUserType {
	return &e
}

/*
Enum describing VDI vendors.
*/
type VdiVendor int

const (
	VDIVENDOR_UNKNOWN     VdiVendor = 0
	VDIVENDOR_REDACTED    VdiVendor = 1
	VDIVENDOR_VIEW        VdiVendor = 2
	VDIVENDOR_XEN_DESKTOP VdiVendor = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *VdiVendor) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"VIEW",
		"XEN_DESKTOP",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e VdiVendor) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"VIEW",
		"XEN_DESKTOP",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *VdiVendor) index(name string) VdiVendor {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"VIEW",
		"XEN_DESKTOP",
	}
	for idx := range names {
		if names[idx] == name {
			return VdiVendor(idx)
		}
	}
	return VDIVENDOR_UNKNOWN
}

func (e *VdiVendor) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for VdiVendor:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *VdiVendor) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e VdiVendor) Ref() *VdiVendor {
	return &e
}

/*
VDI workload description.
*/
type VdiWorkload struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	ProvisionType *VdiProvisionType `json:"provisionType"`
	/*
	  Number of VDI workload users.
	*/
	UserCount *int `json:"userCount"`

	UserType *VdiUserType `json:"userType"`

	Vendor *VdiVendor `json:"vendor"`
}

func (p *VdiWorkload) MarshalJSON() ([]byte, error) {
	type VdiWorkloadProxy VdiWorkload
	return json.Marshal(struct {
		*VdiWorkloadProxy
		ProvisionType *VdiProvisionType `json:"provisionType,omitempty"`
		UserCount     *int              `json:"userCount,omitempty"`
		UserType      *VdiUserType      `json:"userType,omitempty"`
		Vendor        *VdiVendor        `json:"vendor,omitempty"`
	}{
		VdiWorkloadProxy: (*VdiWorkloadProxy)(p),
		ProvisionType:    p.ProvisionType,
		UserCount:        p.UserCount,
		UserType:         p.UserType,
		Vendor:           p.Vendor,
	})
}

func NewVdiWorkload() *VdiWorkload {
	p := new(VdiWorkload)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.config.VdiWorkload"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Allowed vendors in a WhatIf Scenario.
*/
type Vendor int

const (
	VENDOR_UNKNOWN  Vendor = 0
	VENDOR_REDACTED Vendor = 1
	VENDOR_NUTANIX  Vendor = 2
	VENDOR_DELL     Vendor = 3
	VENDOR_LENOVO   Vendor = 4
	VENDOR_CISCO    Vendor = 5
	VENDOR_IBM      Vendor = 6
	VENDOR_HPE_DX   Vendor = 7
	VENDOR_AWS      Vendor = 8
	VENDOR_FUJITSU  Vendor = 9
	VENDOR_AZURE    Vendor = 10
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *Vendor) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NUTANIX",
		"DELL",
		"LENOVO",
		"CISCO",
		"IBM",
		"HPE_DX",
		"AWS",
		"FUJITSU",
		"AZURE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e Vendor) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NUTANIX",
		"DELL",
		"LENOVO",
		"CISCO",
		"IBM",
		"HPE_DX",
		"AWS",
		"FUJITSU",
		"AZURE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *Vendor) index(name string) Vendor {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NUTANIX",
		"DELL",
		"LENOVO",
		"CISCO",
		"IBM",
		"HPE_DX",
		"AWS",
		"FUJITSU",
		"AZURE",
	}
	for idx := range names {
		if names[idx] == name {
			return Vendor(idx)
		}
	}
	return VENDOR_UNKNOWN
}

func (e *Vendor) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for Vendor:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *Vendor) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e Vendor) Ref() *Vendor {
	return &e
}

/*
VM category workload.
*/
type VmCategoryWorkload struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  External identifier of the VM category.
	*/
	CategoryExtId *string `json:"categoryExtId"`
	/*
	  Number of VMs in category.
	*/
	CurrentVmCount *int `json:"currentVmCount"`
	/*
	  Number of VMs considered in category.
	*/
	TargetVmCount *int `json:"targetVmCount"`
}

func (p *VmCategoryWorkload) MarshalJSON() ([]byte, error) {
	type VmCategoryWorkloadProxy VmCategoryWorkload
	return json.Marshal(struct {
		*VmCategoryWorkloadProxy
		CategoryExtId  *string `json:"categoryExtId,omitempty"`
		CurrentVmCount *int    `json:"currentVmCount,omitempty"`
		TargetVmCount  *int    `json:"targetVmCount,omitempty"`
	}{
		VmCategoryWorkloadProxy: (*VmCategoryWorkloadProxy)(p),
		CategoryExtId:           p.CategoryExtId,
		CurrentVmCount:          p.CurrentVmCount,
		TargetVmCount:           p.TargetVmCount,
	})
}

func NewVmCategoryWorkload() *VmCategoryWorkload {
	p := new(VmCategoryWorkload)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.config.VmCategoryWorkload"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
VM workload description.
*/
type VmWorkload struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The UUID of the simulation that is created manually.
	*/
	SimulationExtId *string `json:"simulationExtId"`
	/*
	  Number of VMs in a VM workload.
	*/
	VmCount *int `json:"vmCount"`
}

func (p *VmWorkload) MarshalJSON() ([]byte, error) {
	type VmWorkloadProxy VmWorkload
	return json.Marshal(struct {
		*VmWorkloadProxy
		SimulationExtId *string `json:"simulationExtId,omitempty"`
		VmCount         *int    `json:"vmCount,omitempty"`
	}{
		VmWorkloadProxy: (*VmWorkloadProxy)(p),
		SimulationExtId: p.SimulationExtId,
		VmCount:         p.VmCount,
	})
}

func NewVmWorkload() *VmWorkload {
	p := new(VmWorkload)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.config.VmWorkload"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Workload in a scenario.
*/
type Workload struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Flag to indicate whether the added workload in the planned capacity scenario is to be included in whatif analysis or to be ignored.
	*/
	IsEnabled *bool `json:"isEnabled"`

	ProjectedResourceRequirement *ResourceCapacity `json:"projectedResourceRequirement,omitempty"`
	/*
	  Time since the workload is planned to run on a cluster.
	*/
	ScheduleDate *time.Time `json:"scheduleDate"`
	/*

	 */
	WorkloadPropertiesItemDiscriminator_ *string `json:"$workloadPropertiesItemDiscriminator,omitempty"`
	/*
	  Metadata about the workload in a capacity planning scenario. For example, if SQL is added as a workload, then the number of users and database transaction type will be metadata for it.
	*/
	WorkloadProperties *OneOfWorkloadWorkloadProperties `json:"workloadProperties"`
}

func (p *Workload) MarshalJSON() ([]byte, error) {
	type WorkloadProxy Workload
	return json.Marshal(struct {
		*WorkloadProxy
		IsEnabled          *bool                            `json:"isEnabled,omitempty"`
		ScheduleDate       string                           `json:"scheduleDate,omitempty"`
		WorkloadProperties *OneOfWorkloadWorkloadProperties `json:"workloadProperties,omitempty"`
	}{
		WorkloadProxy: (*WorkloadProxy)(p),
		IsEnabled:     p.IsEnabled,
		ScheduleDate: func() string {
			if p.ScheduleDate != nil {
				return p.ScheduleDate.Format("2006-01-02")
			} else {
				return ""
			}
		}(),
		WorkloadProperties: p.WorkloadProperties,
	})
}

func (p *Workload) UnmarshalJSON(b []byte) error {
	type CustomWorkload struct {
		ObjectType_                          *string                          `json:"$objectType,omitempty"`
		Reserved_                            map[string]interface{}           `json:"$reserved,omitempty"`
		UnknownFields_                       map[string]interface{}           `json:"$unknownFields,omitempty"`
		IsEnabled                            *bool                            `json:"isEnabled"`
		ProjectedResourceRequirement         *ResourceCapacity                `json:"projectedResourceRequirement,omitempty"`
		ScheduleDate                         string                           `json:"scheduleDate"`
		WorkloadPropertiesItemDiscriminator_ *string                          `json:"$workloadPropertiesItemDiscriminator,omitempty"`
		WorkloadProperties                   *OneOfWorkloadWorkloadProperties `json:"workloadProperties"`
	}

	var customWorkload CustomWorkload
	err := json.Unmarshal(b, &customWorkload)
	if err != nil {
		return err
	}

	p.ObjectType_ = customWorkload.ObjectType_
	p.Reserved_ = customWorkload.Reserved_
	p.UnknownFields_ = customWorkload.UnknownFields_
	p.IsEnabled = customWorkload.IsEnabled
	p.ProjectedResourceRequirement = customWorkload.ProjectedResourceRequirement
	// Custom date parsing logic for Date field
	if customWorkload.ScheduleDate != "" {
		parsedScheduleDate, err := time.Parse("2006-01-02", customWorkload.ScheduleDate)
		if err != nil {
			return errors.New(fmt.Sprintf("Unable to unmarshal field ScheduleDate in struct Workload: %s", err))
		}
		p.ScheduleDate = &parsedScheduleDate
	}
	p.WorkloadPropertiesItemDiscriminator_ = customWorkload.WorkloadPropertiesItemDiscriminator_
	p.WorkloadProperties = customWorkload.WorkloadProperties

	return nil
}

func NewWorkload() *Workload {
	p := new(Workload)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.config.Workload"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	p.IsEnabled = new(bool)
	*p.IsEnabled = false

	return p
}

func (p *Workload) GetWorkloadProperties() interface{} {
	if nil == p.WorkloadProperties {
		return nil
	}
	return p.WorkloadProperties.GetValue()
}

func (p *Workload) SetWorkloadProperties(v interface{}) error {
	if nil == p.WorkloadProperties {
		p.WorkloadProperties = NewOneOfWorkloadWorkloadProperties()
	}
	e := p.WorkloadProperties.SetValue(v)
	if nil == e {
		if nil == p.WorkloadPropertiesItemDiscriminator_ {
			p.WorkloadPropertiesItemDiscriminator_ = new(string)
		}
		*p.WorkloadPropertiesItemDiscriminator_ = *p.WorkloadProperties.Discriminator
	}
	return e
}

/*
Enum describing Xen operating system types.
*/
type XenOperatingSystem int

const (
	XENOPERATINGSYSTEM_UNKNOWN        XenOperatingSystem = 0
	XENOPERATINGSYSTEM_REDACTED       XenOperatingSystem = 1
	XENOPERATINGSYSTEM_WINDOWS_2012R2 XenOperatingSystem = 2
	XENOPERATINGSYSTEM_WINDOWS_2008R2 XenOperatingSystem = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *XenOperatingSystem) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"WINDOWS_2012R2",
		"WINDOWS_2008R2",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e XenOperatingSystem) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"WINDOWS_2012R2",
		"WINDOWS_2008R2",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *XenOperatingSystem) index(name string) XenOperatingSystem {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"WINDOWS_2012R2",
		"WINDOWS_2008R2",
	}
	for idx := range names {
		if names[idx] == name {
			return XenOperatingSystem(idx)
		}
	}
	return XENOPERATINGSYSTEM_UNKNOWN
}

func (e *XenOperatingSystem) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for XenOperatingSystem:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *XenOperatingSystem) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e XenOperatingSystem) Ref() *XenOperatingSystem {
	return &e
}

/*
Xen workload description.
*/
type XenWorkload struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Size in GB for MCS different disk per VM.
	*/
	McsDiffSizeGb *int `json:"mcsDiffSizeGb"`

	OperatingSystem *XenOperatingSystem `json:"operatingSystem"`
	/*
	  Size of PVS cache write per VM in GB.
	*/
	PvsWriteCacheSizeGb *int `json:"pvsWriteCacheSizeGb"`
	/*
	  Space consumed by each Xen Server Image.
	*/
	SystemDataGb *int `json:"systemDataGb"`
	/*
	  Number of users for Xen workload.
	*/
	UserCount *int `json:"userCount"`
	/*
	  Size for the per user profile data in MB.
	*/
	UserProfileDataMb *int `json:"userProfileDataMb"`
}

func (p *XenWorkload) MarshalJSON() ([]byte, error) {
	type XenWorkloadProxy XenWorkload
	return json.Marshal(struct {
		*XenWorkloadProxy
		McsDiffSizeGb       *int                `json:"mcsDiffSizeGb,omitempty"`
		OperatingSystem     *XenOperatingSystem `json:"operatingSystem,omitempty"`
		PvsWriteCacheSizeGb *int                `json:"pvsWriteCacheSizeGb,omitempty"`
		SystemDataGb        *int                `json:"systemDataGb,omitempty"`
		UserCount           *int                `json:"userCount,omitempty"`
		UserProfileDataMb   *int                `json:"userProfileDataMb,omitempty"`
	}{
		XenWorkloadProxy:    (*XenWorkloadProxy)(p),
		McsDiffSizeGb:       p.McsDiffSizeGb,
		OperatingSystem:     p.OperatingSystem,
		PvsWriteCacheSizeGb: p.PvsWriteCacheSizeGb,
		SystemDataGb:        p.SystemDataGb,
		UserCount:           p.UserCount,
		UserProfileDataMb:   p.UserProfileDataMb,
	})
}

func NewXenWorkload() *XenWorkload {
	p := new(XenWorkload)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.config.XenWorkload"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type XfitPolicy struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of associated category details for the policy.
	*/
	Categories []CategoryDetail `json:"categories"`
	/*
	  The description of the policy.
	*/
	Description *string `json:"description,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Indicates whether the policy is applied by default.
	*/
	IsDefault *bool `json:"isDefault,omitempty"`
	/*
	  Indicates if the policy is enabled or disabled.
	*/
	IsEnabled *bool `json:"isEnabled"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import3.ApiLink `json:"links,omitempty"`
	/*
	  The name of the policy.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  Key value pairs for the relevant properties of a policy. For example, it can hold the lookback period for an Inefficiency exclusion policy.
	*/
	Parameters []import5.KVStringPair `json:"parameters"`

	PolicyType *PolicyType `json:"policyType,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  Last updated time of the policy in ISO-8601 format.
	*/
	UpdatedTime *time.Time `json:"updatedTime,omitempty"`
}

func (p *XfitPolicy) MarshalJSON() ([]byte, error) {
	type XfitPolicyProxy XfitPolicy
	return json.Marshal(struct {
		*XfitPolicyProxy
		Categories []CategoryDetail       `json:"categories,omitempty"`
		IsEnabled  *bool                  `json:"isEnabled,omitempty"`
		Parameters []import5.KVStringPair `json:"parameters,omitempty"`
	}{
		XfitPolicyProxy: (*XfitPolicyProxy)(p),
		Categories:      p.Categories,
		IsEnabled:       p.IsEnabled,
		Parameters:      p.Parameters,
	})
}

func NewXfitPolicy() *XfitPolicy {
	p := new(XfitPolicy)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.config.XfitPolicy"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	p.IsDefault = new(bool)
	*p.IsDefault = false
	p.IsEnabled = new(bool)
	*p.IsEnabled = true

	return p
}

type XfitPolicyProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of associated category details for the policy.
	*/
	Categories []CategoryDetail `json:"categories"`
	/*
	  The description of the policy.
	*/
	Description *string `json:"description,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Indicates whether the policy is applied by default.
	*/
	IsDefault *bool `json:"isDefault,omitempty"`
	/*
	  Indicates if the policy is enabled or disabled.
	*/
	IsEnabled *bool `json:"isEnabled"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import3.ApiLink `json:"links,omitempty"`
	/*
	  The name of the policy.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  Key value pairs for the relevant properties of a policy. For example, it can hold the lookback period for an Inefficiency exclusion policy.
	*/
	Parameters []import5.KVStringPair `json:"parameters"`

	PolicyType *PolicyType `json:"policyType,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  Last updated time of the policy in ISO-8601 format.
	*/
	UpdatedTime *time.Time `json:"updatedTime,omitempty"`
}

func (p *XfitPolicyProjection) MarshalJSON() ([]byte, error) {
	type XfitPolicyProjectionProxy XfitPolicyProjection
	return json.Marshal(struct {
		*XfitPolicyProjectionProxy
		Categories []CategoryDetail       `json:"categories,omitempty"`
		IsEnabled  *bool                  `json:"isEnabled,omitempty"`
		Parameters []import5.KVStringPair `json:"parameters,omitempty"`
	}{
		XfitPolicyProjectionProxy: (*XfitPolicyProjectionProxy)(p),
		Categories:                p.Categories,
		IsEnabled:                 p.IsEnabled,
		Parameters:                p.Parameters,
	})
}

func NewXfitPolicyProjection() *XfitPolicyProjection {
	p := new(XfitPolicyProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.config.XfitPolicyProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	p.IsDefault = new(bool)
	*p.IsDefault = false
	p.IsEnabled = new(bool)
	*p.IsEnabled = true

	return p
}

type OneOfUpdateScenarioApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfUpdateScenarioApiResponseData() *OneOfUpdateScenarioApiResponseData {
	p := new(OneOfUpdateScenarioApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateScenarioApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateScenarioApiResponseData is nil"))
	}
	switch v.(type) {
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

func (p *OneOfUpdateScenarioApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfUpdateScenarioApiResponseData) UnmarshalJSON(b []byte) error {
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "aiops.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateScenarioApiResponseData"))
}

func (p *OneOfUpdateScenarioApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateScenarioApiResponseData")
}

type OneOfEntityTypeListApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    []EntityType           `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfEntityTypeListApiResponseData() *OneOfEntityTypeListApiResponseData {
	p := new(OneOfEntityTypeListApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfEntityTypeListApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfEntityTypeListApiResponseData is nil"))
	}
	switch v.(type) {
	case []EntityType:
		p.oneOfType0 = v.([]EntityType)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<aiops.v4.config.EntityType>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<aiops.v4.config.EntityType>"
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

func (p *OneOfEntityTypeListApiResponseData) GetValue() interface{} {
	if "List<aiops.v4.config.EntityType>" == *p.Discriminator {
		return p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfEntityTypeListApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new([]EntityType)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "aiops.v4.config.EntityType" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<aiops.v4.config.EntityType>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<aiops.v4.config.EntityType>"
			return nil
		}
	}
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "aiops.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfEntityTypeListApiResponseData"))
}

func (p *OneOfEntityTypeListApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<aiops.v4.config.EntityType>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfEntityTypeListApiResponseData")
}

type OneOfGetSimulationApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *Simulation            `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfGetSimulationApiResponseData() *OneOfGetSimulationApiResponseData {
	p := new(OneOfGetSimulationApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetSimulationApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetSimulationApiResponseData is nil"))
	}
	switch v.(type) {
	case Simulation:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(Simulation)
		}
		*p.oneOfType0 = v.(Simulation)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
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

func (p *OneOfGetSimulationApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetSimulationApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(Simulation)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "aiops.v4.config.Simulation" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(Simulation)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "aiops.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetSimulationApiResponseData"))
}

func (p *OneOfGetSimulationApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetSimulationApiResponseData")
}

type OneOfListScenariosApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType401  []ScenarioProjection   `json:"-"`
	oneOfType0    []Scenario             `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfListScenariosApiResponseData() *OneOfListScenariosApiResponseData {
	p := new(OneOfListScenariosApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListScenariosApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListScenariosApiResponseData is nil"))
	}
	switch v.(type) {
	case []ScenarioProjection:
		p.oneOfType401 = v.([]ScenarioProjection)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<aiops.v4.config.ScenarioProjection>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<aiops.v4.config.ScenarioProjection>"
	case []Scenario:
		p.oneOfType0 = v.([]Scenario)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<aiops.v4.config.Scenario>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<aiops.v4.config.Scenario>"
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

func (p *OneOfListScenariosApiResponseData) GetValue() interface{} {
	if "List<aiops.v4.config.ScenarioProjection>" == *p.Discriminator {
		return p.oneOfType401
	}
	if "List<aiops.v4.config.Scenario>" == *p.Discriminator {
		return p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfListScenariosApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType401 := new([]ScenarioProjection)
	if err := json.Unmarshal(b, vOneOfType401); err == nil {
		if len(*vOneOfType401) == 0 || "aiops.v4.config.ScenarioProjection" == *((*vOneOfType401)[0].ObjectType_) {
			p.oneOfType401 = *vOneOfType401
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<aiops.v4.config.ScenarioProjection>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<aiops.v4.config.ScenarioProjection>"
			return nil
		}
	}
	vOneOfType0 := new([]Scenario)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "aiops.v4.config.Scenario" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<aiops.v4.config.Scenario>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<aiops.v4.config.Scenario>"
			return nil
		}
	}
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "aiops.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListScenariosApiResponseData"))
}

func (p *OneOfListScenariosApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<aiops.v4.config.ScenarioProjection>" == *p.Discriminator {
		return json.Marshal(p.oneOfType401)
	}
	if "List<aiops.v4.config.Scenario>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfListScenariosApiResponseData")
}

type OneOfDeleteSimulationApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType1    *interface{}           `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfDeleteSimulationApiResponseData() *OneOfDeleteSimulationApiResponseData {
	p := new(OneOfDeleteSimulationApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteSimulationApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteSimulationApiResponseData is nil"))
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

func (p *OneOfDeleteSimulationApiResponseData) GetValue() interface{} {
	if "EMPTY" == *p.Discriminator {
		return *p.oneOfType1
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDeleteSimulationApiResponseData) UnmarshalJSON(b []byte) error {
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
		if "aiops.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteSimulationApiResponseData"))
}

func (p *OneOfDeleteSimulationApiResponseData) MarshalJSON() ([]byte, error) {
	if "EMPTY" == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteSimulationApiResponseData")
}

type OneOfDeleteIgnoreWindowApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType1    *interface{}           `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfDeleteIgnoreWindowApiResponseData() *OneOfDeleteIgnoreWindowApiResponseData {
	p := new(OneOfDeleteIgnoreWindowApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteIgnoreWindowApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteIgnoreWindowApiResponseData is nil"))
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

func (p *OneOfDeleteIgnoreWindowApiResponseData) GetValue() interface{} {
	if "EMPTY" == *p.Discriminator {
		return *p.oneOfType1
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDeleteIgnoreWindowApiResponseData) UnmarshalJSON(b []byte) error {
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
		if "aiops.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteIgnoreWindowApiResponseData"))
}

func (p *OneOfDeleteIgnoreWindowApiResponseData) MarshalJSON() ([]byte, error) {
	if "EMPTY" == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteIgnoreWindowApiResponseData")
}

type OneOfCreateScenarioApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfCreateScenarioApiResponseData() *OneOfCreateScenarioApiResponseData {
	p := new(OneOfCreateScenarioApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateScenarioApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateScenarioApiResponseData is nil"))
	}
	switch v.(type) {
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

func (p *OneOfCreateScenarioApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfCreateScenarioApiResponseData) UnmarshalJSON(b []byte) error {
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "aiops.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateScenarioApiResponseData"))
}

func (p *OneOfCreateScenarioApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfCreateScenarioApiResponseData")
}

type OneOfSourceListApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    []Source               `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfSourceListApiResponseData() *OneOfSourceListApiResponseData {
	p := new(OneOfSourceListApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfSourceListApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfSourceListApiResponseData is nil"))
	}
	switch v.(type) {
	case []Source:
		p.oneOfType0 = v.([]Source)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<aiops.v4.config.Source>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<aiops.v4.config.Source>"
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

func (p *OneOfSourceListApiResponseData) GetValue() interface{} {
	if "List<aiops.v4.config.Source>" == *p.Discriminator {
		return p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfSourceListApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new([]Source)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "aiops.v4.config.Source" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<aiops.v4.config.Source>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<aiops.v4.config.Source>"
			return nil
		}
	}
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "aiops.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfSourceListApiResponseData"))
}

func (p *OneOfSourceListApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<aiops.v4.config.Source>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfSourceListApiResponseData")
}

type OneOfUpdateIgnoreWindowApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *IgnoreWindow          `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfUpdateIgnoreWindowApiResponseData() *OneOfUpdateIgnoreWindowApiResponseData {
	p := new(OneOfUpdateIgnoreWindowApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateIgnoreWindowApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateIgnoreWindowApiResponseData is nil"))
	}
	switch v.(type) {
	case IgnoreWindow:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(IgnoreWindow)
		}
		*p.oneOfType0 = v.(IgnoreWindow)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
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

func (p *OneOfUpdateIgnoreWindowApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfUpdateIgnoreWindowApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(IgnoreWindow)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "aiops.v4.config.IgnoreWindow" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(IgnoreWindow)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "aiops.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateIgnoreWindowApiResponseData"))
}

func (p *OneOfUpdateIgnoreWindowApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateIgnoreWindowApiResponseData")
}

type OneOfCreateIgnoreWindowApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *IgnoreWindow          `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfCreateIgnoreWindowApiResponseData() *OneOfCreateIgnoreWindowApiResponseData {
	p := new(OneOfCreateIgnoreWindowApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateIgnoreWindowApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateIgnoreWindowApiResponseData is nil"))
	}
	switch v.(type) {
	case IgnoreWindow:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(IgnoreWindow)
		}
		*p.oneOfType0 = v.(IgnoreWindow)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
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

func (p *OneOfCreateIgnoreWindowApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfCreateIgnoreWindowApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(IgnoreWindow)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "aiops.v4.config.IgnoreWindow" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(IgnoreWindow)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "aiops.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateIgnoreWindowApiResponseData"))
}

func (p *OneOfCreateIgnoreWindowApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfCreateIgnoreWindowApiResponseData")
}

type OneOfWorkloadWorkloadProperties struct {
	Discriminator *string               `json:"-"`
	ObjectType_   *string               `json:"-"`
	oneOfType1    *VmWorkload           `json:"-"`
	oneOfType3    *SplunkWorkload       `json:"-"`
	oneOfType4    *CitrixXenWorkload    `json:"-"`
	oneOfType5    *MicrosoftXenWorkload `json:"-"`
	oneOfType0    *SqlWorkload          `json:"-"`
	oneOfType6    *CapacityUpdateConfig `json:"-"`
	oneOfType2    *VdiWorkload          `json:"-"`
	oneOfType7    *VmCategoryWorkload   `json:"-"`
}

func NewOneOfWorkloadWorkloadProperties() *OneOfWorkloadWorkloadProperties {
	p := new(OneOfWorkloadWorkloadProperties)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfWorkloadWorkloadProperties) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfWorkloadWorkloadProperties is nil"))
	}
	switch v.(type) {
	case VmWorkload:
		if nil == p.oneOfType1 {
			p.oneOfType1 = new(VmWorkload)
		}
		*p.oneOfType1 = v.(VmWorkload)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType1.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType1.ObjectType_
	case SplunkWorkload:
		if nil == p.oneOfType3 {
			p.oneOfType3 = new(SplunkWorkload)
		}
		*p.oneOfType3 = v.(SplunkWorkload)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType3.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType3.ObjectType_
	case CitrixXenWorkload:
		if nil == p.oneOfType4 {
			p.oneOfType4 = new(CitrixXenWorkload)
		}
		*p.oneOfType4 = v.(CitrixXenWorkload)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType4.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType4.ObjectType_
	case MicrosoftXenWorkload:
		if nil == p.oneOfType5 {
			p.oneOfType5 = new(MicrosoftXenWorkload)
		}
		*p.oneOfType5 = v.(MicrosoftXenWorkload)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType5.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType5.ObjectType_
	case SqlWorkload:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(SqlWorkload)
		}
		*p.oneOfType0 = v.(SqlWorkload)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case CapacityUpdateConfig:
		if nil == p.oneOfType6 {
			p.oneOfType6 = new(CapacityUpdateConfig)
		}
		*p.oneOfType6 = v.(CapacityUpdateConfig)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType6.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType6.ObjectType_
	case VdiWorkload:
		if nil == p.oneOfType2 {
			p.oneOfType2 = new(VdiWorkload)
		}
		*p.oneOfType2 = v.(VdiWorkload)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2.ObjectType_
	case VmCategoryWorkload:
		if nil == p.oneOfType7 {
			p.oneOfType7 = new(VmCategoryWorkload)
		}
		*p.oneOfType7 = v.(VmCategoryWorkload)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType7.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType7.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfWorkloadWorkloadProperties) GetValue() interface{} {
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1
	}
	if p.oneOfType3 != nil && *p.oneOfType3.ObjectType_ == *p.Discriminator {
		return *p.oneOfType3
	}
	if p.oneOfType4 != nil && *p.oneOfType4.ObjectType_ == *p.Discriminator {
		return *p.oneOfType4
	}
	if p.oneOfType5 != nil && *p.oneOfType5.ObjectType_ == *p.Discriminator {
		return *p.oneOfType5
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType6 != nil && *p.oneOfType6.ObjectType_ == *p.Discriminator {
		return *p.oneOfType6
	}
	if p.oneOfType2 != nil && *p.oneOfType2.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2
	}
	if p.oneOfType7 != nil && *p.oneOfType7.ObjectType_ == *p.Discriminator {
		return *p.oneOfType7
	}
	return nil
}

func (p *OneOfWorkloadWorkloadProperties) UnmarshalJSON(b []byte) error {
	vOneOfType1 := new(VmWorkload)
	if err := json.Unmarshal(b, vOneOfType1); err == nil {
		if "aiops.v4.config.VmWorkload" == *vOneOfType1.ObjectType_ {
			if nil == p.oneOfType1 {
				p.oneOfType1 = new(VmWorkload)
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
	vOneOfType3 := new(SplunkWorkload)
	if err := json.Unmarshal(b, vOneOfType3); err == nil {
		if "aiops.v4.config.SplunkWorkload" == *vOneOfType3.ObjectType_ {
			if nil == p.oneOfType3 {
				p.oneOfType3 = new(SplunkWorkload)
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
	vOneOfType4 := new(CitrixXenWorkload)
	if err := json.Unmarshal(b, vOneOfType4); err == nil {
		if "aiops.v4.config.CitrixXenWorkload" == *vOneOfType4.ObjectType_ {
			if nil == p.oneOfType4 {
				p.oneOfType4 = new(CitrixXenWorkload)
			}
			*p.oneOfType4 = *vOneOfType4
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType4.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType4.ObjectType_
			return nil
		}
	}
	vOneOfType5 := new(MicrosoftXenWorkload)
	if err := json.Unmarshal(b, vOneOfType5); err == nil {
		if "aiops.v4.config.MicrosoftXenWorkload" == *vOneOfType5.ObjectType_ {
			if nil == p.oneOfType5 {
				p.oneOfType5 = new(MicrosoftXenWorkload)
			}
			*p.oneOfType5 = *vOneOfType5
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType5.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType5.ObjectType_
			return nil
		}
	}
	vOneOfType0 := new(SqlWorkload)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "aiops.v4.config.SqlWorkload" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(SqlWorkload)
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
	vOneOfType6 := new(CapacityUpdateConfig)
	if err := json.Unmarshal(b, vOneOfType6); err == nil {
		if "aiops.v4.config.CapacityUpdateConfig" == *vOneOfType6.ObjectType_ {
			if nil == p.oneOfType6 {
				p.oneOfType6 = new(CapacityUpdateConfig)
			}
			*p.oneOfType6 = *vOneOfType6
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType6.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType6.ObjectType_
			return nil
		}
	}
	vOneOfType2 := new(VdiWorkload)
	if err := json.Unmarshal(b, vOneOfType2); err == nil {
		if "aiops.v4.config.VdiWorkload" == *vOneOfType2.ObjectType_ {
			if nil == p.oneOfType2 {
				p.oneOfType2 = new(VdiWorkload)
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
	vOneOfType7 := new(VmCategoryWorkload)
	if err := json.Unmarshal(b, vOneOfType7); err == nil {
		if "aiops.v4.config.VmCategoryWorkload" == *vOneOfType7.ObjectType_ {
			if nil == p.oneOfType7 {
				p.oneOfType7 = new(VmCategoryWorkload)
			}
			*p.oneOfType7 = *vOneOfType7
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType7.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType7.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfWorkloadWorkloadProperties"))
}

func (p *OneOfWorkloadWorkloadProperties) MarshalJSON() ([]byte, error) {
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	if p.oneOfType3 != nil && *p.oneOfType3.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType3)
	}
	if p.oneOfType4 != nil && *p.oneOfType4.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType4)
	}
	if p.oneOfType5 != nil && *p.oneOfType5.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType5)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType6 != nil && *p.oneOfType6.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType6)
	}
	if p.oneOfType2 != nil && *p.oneOfType2.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2)
	}
	if p.oneOfType7 != nil && *p.oneOfType7.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType7)
	}
	return nil, errors.New("No value to marshal for OneOfWorkloadWorkloadProperties")
}

type OneOfUpdatePolicyTaskApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfUpdatePolicyTaskApiResponseData() *OneOfUpdatePolicyTaskApiResponseData {
	p := new(OneOfUpdatePolicyTaskApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdatePolicyTaskApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdatePolicyTaskApiResponseData is nil"))
	}
	switch v.(type) {
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

func (p *OneOfUpdatePolicyTaskApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfUpdatePolicyTaskApiResponseData) UnmarshalJSON(b []byte) error {
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "aiops.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdatePolicyTaskApiResponseData"))
}

func (p *OneOfUpdatePolicyTaskApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfUpdatePolicyTaskApiResponseData")
}

type OneOfGenerateRunwayApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfGenerateRunwayApiResponseData() *OneOfGenerateRunwayApiResponseData {
	p := new(OneOfGenerateRunwayApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGenerateRunwayApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGenerateRunwayApiResponseData is nil"))
	}
	switch v.(type) {
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

func (p *OneOfGenerateRunwayApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGenerateRunwayApiResponseData) UnmarshalJSON(b []byte) error {
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "aiops.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGenerateRunwayApiResponseData"))
}

func (p *OneOfGenerateRunwayApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGenerateRunwayApiResponseData")
}

type OneOfUpdateSimulationApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    []import2.AppMessage   `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfUpdateSimulationApiResponseData() *OneOfUpdateSimulationApiResponseData {
	p := new(OneOfUpdateSimulationApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateSimulationApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateSimulationApiResponseData is nil"))
	}
	switch v.(type) {
	case []import2.AppMessage:
		p.oneOfType0 = v.([]import2.AppMessage)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<aiops.v4.error.AppMessage>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<aiops.v4.error.AppMessage>"
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

func (p *OneOfUpdateSimulationApiResponseData) GetValue() interface{} {
	if "List<aiops.v4.error.AppMessage>" == *p.Discriminator {
		return p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfUpdateSimulationApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new([]import2.AppMessage)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "aiops.v4.error.AppMessage" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<aiops.v4.error.AppMessage>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<aiops.v4.error.AppMessage>"
			return nil
		}
	}
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "aiops.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateSimulationApiResponseData"))
}

func (p *OneOfUpdateSimulationApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<aiops.v4.error.AppMessage>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateSimulationApiResponseData")
}

type OneOfGenerateRecommendationApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfGenerateRecommendationApiResponseData() *OneOfGenerateRecommendationApiResponseData {
	p := new(OneOfGenerateRecommendationApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGenerateRecommendationApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGenerateRecommendationApiResponseData is nil"))
	}
	switch v.(type) {
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

func (p *OneOfGenerateRecommendationApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGenerateRecommendationApiResponseData) UnmarshalJSON(b []byte) error {
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "aiops.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGenerateRecommendationApiResponseData"))
}

func (p *OneOfGenerateRecommendationApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGenerateRecommendationApiResponseData")
}

type OneOfCreateSimulationApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *Simulation            `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfCreateSimulationApiResponseData() *OneOfCreateSimulationApiResponseData {
	p := new(OneOfCreateSimulationApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateSimulationApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateSimulationApiResponseData is nil"))
	}
	switch v.(type) {
	case Simulation:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(Simulation)
		}
		*p.oneOfType0 = v.(Simulation)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
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

func (p *OneOfCreateSimulationApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfCreateSimulationApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(Simulation)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "aiops.v4.config.Simulation" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(Simulation)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "aiops.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateSimulationApiResponseData"))
}

func (p *OneOfCreateSimulationApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfCreateSimulationApiResponseData")
}

type OneOfGetScenarioApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *Scenario              `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfGetScenarioApiResponseData() *OneOfGetScenarioApiResponseData {
	p := new(OneOfGetScenarioApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetScenarioApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetScenarioApiResponseData is nil"))
	}
	switch v.(type) {
	case Scenario:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(Scenario)
		}
		*p.oneOfType0 = v.(Scenario)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
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

func (p *OneOfGetScenarioApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetScenarioApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(Scenario)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "aiops.v4.config.Scenario" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(Scenario)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "aiops.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetScenarioApiResponseData"))
}

func (p *OneOfGetScenarioApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetScenarioApiResponseData")
}

type OneOfListSimulationsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    []Simulation           `json:"-"`
	oneOfType401  []SimulationProjection `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfListSimulationsApiResponseData() *OneOfListSimulationsApiResponseData {
	p := new(OneOfListSimulationsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListSimulationsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListSimulationsApiResponseData is nil"))
	}
	switch v.(type) {
	case []Simulation:
		p.oneOfType0 = v.([]Simulation)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<aiops.v4.config.Simulation>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<aiops.v4.config.Simulation>"
	case []SimulationProjection:
		p.oneOfType401 = v.([]SimulationProjection)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<aiops.v4.config.SimulationProjection>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<aiops.v4.config.SimulationProjection>"
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

func (p *OneOfListSimulationsApiResponseData) GetValue() interface{} {
	if "List<aiops.v4.config.Simulation>" == *p.Discriminator {
		return p.oneOfType0
	}
	if "List<aiops.v4.config.SimulationProjection>" == *p.Discriminator {
		return p.oneOfType401
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfListSimulationsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new([]Simulation)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "aiops.v4.config.Simulation" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<aiops.v4.config.Simulation>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<aiops.v4.config.Simulation>"
			return nil
		}
	}
	vOneOfType401 := new([]SimulationProjection)
	if err := json.Unmarshal(b, vOneOfType401); err == nil {
		if len(*vOneOfType401) == 0 || "aiops.v4.config.SimulationProjection" == *((*vOneOfType401)[0].ObjectType_) {
			p.oneOfType401 = *vOneOfType401
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<aiops.v4.config.SimulationProjection>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<aiops.v4.config.SimulationProjection>"
			return nil
		}
	}
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "aiops.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListSimulationsApiResponseData"))
}

func (p *OneOfListSimulationsApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<aiops.v4.config.Simulation>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if "List<aiops.v4.config.SimulationProjection>" == *p.Discriminator {
		return json.Marshal(p.oneOfType401)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfListSimulationsApiResponseData")
}

type OneOfDeleteScenarioApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType1    *interface{}           `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfDeleteScenarioApiResponseData() *OneOfDeleteScenarioApiResponseData {
	p := new(OneOfDeleteScenarioApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteScenarioApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteScenarioApiResponseData is nil"))
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

func (p *OneOfDeleteScenarioApiResponseData) GetValue() interface{} {
	if "EMPTY" == *p.Discriminator {
		return *p.oneOfType1
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDeleteScenarioApiResponseData) UnmarshalJSON(b []byte) error {
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
		if "aiops.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteScenarioApiResponseData"))
}

func (p *OneOfDeleteScenarioApiResponseData) MarshalJSON() ([]byte, error) {
	if "EMPTY" == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteScenarioApiResponseData")
}

type OneOfGenerateReportApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfGenerateReportApiResponseData() *OneOfGenerateReportApiResponseData {
	p := new(OneOfGenerateReportApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGenerateReportApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGenerateReportApiResponseData is nil"))
	}
	switch v.(type) {
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

func (p *OneOfGenerateReportApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGenerateReportApiResponseData) UnmarshalJSON(b []byte) error {
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "aiops.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGenerateReportApiResponseData"))
}

func (p *OneOfGenerateReportApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGenerateReportApiResponseData")
}

type OneOfGetScenarioReportApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *FileDetail            `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfGetScenarioReportApiResponseData() *OneOfGetScenarioReportApiResponseData {
	p := new(OneOfGetScenarioReportApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetScenarioReportApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetScenarioReportApiResponseData is nil"))
	}
	switch v.(type) {
	case FileDetail:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(FileDetail)
		}
		*p.oneOfType0 = v.(FileDetail)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "FileDetail"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "FileDetail"
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

func (p *OneOfGetScenarioReportApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && "FileDetail" == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetScenarioReportApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(FileDetail)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(FileDetail)
		}
		*p.oneOfType0 = *vOneOfType0
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "FileDetail"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "FileDetail"
		return nil
	}
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "aiops.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetScenarioReportApiResponseData"))
}

func (p *OneOfGetScenarioReportApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && "FileDetail" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetScenarioReportApiResponseData")
}

type OneOfCalculateRunwayApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfCalculateRunwayApiResponseData() *OneOfCalculateRunwayApiResponseData {
	p := new(OneOfCalculateRunwayApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCalculateRunwayApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCalculateRunwayApiResponseData is nil"))
	}
	switch v.(type) {
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

func (p *OneOfCalculateRunwayApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfCalculateRunwayApiResponseData) UnmarshalJSON(b []byte) error {
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "aiops.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCalculateRunwayApiResponseData"))
}

func (p *OneOfCalculateRunwayApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfCalculateRunwayApiResponseData")
}

type OneOfEntityDescriptorListApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    []EntityDescriptor     `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfEntityDescriptorListApiResponseData() *OneOfEntityDescriptorListApiResponseData {
	p := new(OneOfEntityDescriptorListApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfEntityDescriptorListApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfEntityDescriptorListApiResponseData is nil"))
	}
	switch v.(type) {
	case []EntityDescriptor:
		p.oneOfType0 = v.([]EntityDescriptor)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<aiops.v4.config.EntityDescriptor>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<aiops.v4.config.EntityDescriptor>"
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

func (p *OneOfEntityDescriptorListApiResponseData) GetValue() interface{} {
	if "List<aiops.v4.config.EntityDescriptor>" == *p.Discriminator {
		return p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfEntityDescriptorListApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new([]EntityDescriptor)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "aiops.v4.config.EntityDescriptor" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<aiops.v4.config.EntityDescriptor>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<aiops.v4.config.EntityDescriptor>"
			return nil
		}
	}
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "aiops.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfEntityDescriptorListApiResponseData"))
}

func (p *OneOfEntityDescriptorListApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<aiops.v4.config.EntityDescriptor>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfEntityDescriptorListApiResponseData")
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
