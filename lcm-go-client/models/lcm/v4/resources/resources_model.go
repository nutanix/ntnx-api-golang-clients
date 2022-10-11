/*
 * Generated file models/lcm/v4/resources/resources_model.go.
 *
 * Product version: 4.0.1-alpha-1
 *
 * Part of the Nutanix Lcm Versioned APIs
 *
 * (c) 2022 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module for LCM Resources
*/
package resources
import (
  "bytes"
  import2 "github.com/nutanix/ntnx-api-golang-clients/lcm-go-client/v4/models/common/v1/response"
  "encoding/json"
  "errors"
  "fmt"
  import3 "github.com/nutanix/ntnx-api-golang-clients/lcm-go-client/v4/models/lcm/v4/common"
  import1 "github.com/nutanix/ntnx-api-golang-clients/lcm-go-client/v4/models/lcm/v4/error"
  import4 "github.com/nutanix/ntnx-api-golang-clients/lcm-go-client/v4/models/prism/v4/config"
)

/**
REST response for all response codes in api path /lcm/v4.0.a1/resources/bundles/{uuid} Delete operation
*/
type DelBundleApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfDelBundleApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewDelBundleApiResponse() *DelBundleApiResponse {
  p := new(DelBundleApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.resources.DelBundleApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.resources.DelBundleApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *DelBundleApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *DelBundleApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfDelBundleApiResponseData()
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
List of deployable versions based on entity types, versions and its dependencies
*/
type DeployableVersion struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  UpdateDependencies *import3.EntityUpdateSpecs `json:"updateDependencies,omitempty"`
  
  VersionIdentifierSpec *import3.VersionIdentifierSpec `json:"versionIdentifierSpec,omitempty"`
}


func NewDeployableVersion() *DeployableVersion {
  p := new(DeployableVersion)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.resources.DeployableVersion"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.resources.DeployableVersion"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Information about an LCM entity affected by the operation.
*/
type EntityInfo struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Detailed information for the LCM entity. For example, firmware entities contain additional information about NIC and so on
  */
  ChildEntities *string `json:"childEntities,omitempty"`
  /**
  LCM entity class
  */
  EntityClass *string `json:"entityClass,omitempty"`
  /**
  LCM entity model
  */
  EntityModel *string `json:"entityModel,omitempty"`
  
  EntityType *import3.LcmEntityType `json:"entityType,omitempty"`
  
  Instances []EntityInstance `json:"instances,omitempty"`
  
  OpEnvironment *OperationEnv `json:"opEnvironment,omitempty"`
}


func NewEntityInfo() *EntityInfo {
  p := new(EntityInfo)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.resources.EntityInfo"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.resources.EntityInfo"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
LCM entity instance of an entity affected by the operation
*/
type EntityInstance struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  End time for an LCM operation
  */
  EndTime *string `json:"endTime,omitempty"`
  /**
  The location of the LCM entity, it can be at the cluster level or at the node level, in the format of "cluster:uuid" or "node:uuid"
  */
  LocationId *string `json:"locationId,omitempty"`
  /**
  The requested update version of an LCM entity.
  */
  RequestVersion *string `json:"requestVersion,omitempty"`
  /**
  Start time for an LCM operation
  */
  StartTime *string `json:"startTime,omitempty"`
  
  TaskInfo []TaskInfo `json:"taskInfo,omitempty"`
  /**
  Current version of an LCM entity
  */
  Version *string `json:"version,omitempty"`
}


func NewEntityInstance() *EntityInstance {
  p := new(EntityInstance)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.resources.EntityInstance"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.resources.EntityInstance"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST response for all response codes in api path /lcm/v4.0.a1/resources/notifications Post operation
*/
type GenUpgradeNotifApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfGenUpgradeNotifApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewGenUpgradeNotifApiResponse() *GenUpgradeNotifApiResponse {
  p := new(GenUpgradeNotifApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.resources.GenUpgradeNotifApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.resources.GenUpgradeNotifApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *GenUpgradeNotifApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *GenUpgradeNotifApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfGenUpgradeNotifApiResponseData()
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
REST response for all response codes in api path /lcm/v4.0.a1/resources/plan Post operation
*/
type GenUpgradePlanApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfGenUpgradePlanApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewGenUpgradePlanApiResponse() *GenUpgradePlanApiResponse {
  p := new(GenUpgradePlanApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.resources.GenUpgradePlanApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.resources.GenUpgradePlanApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *GenUpgradePlanApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *GenUpgradePlanApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfGenUpgradePlanApiResponseData()
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
REST response for all response codes in api path /lcm/v4.0.a1/resources/bundles/{uuid} Get operation
*/
type GetBundleInfoApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfGetBundleInfoApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewGetBundleInfoApiResponse() *GetBundleInfoApiResponse {
  p := new(GetBundleInfoApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.resources.GetBundleInfoApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.resources.GetBundleInfoApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *GetBundleInfoApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *GetBundleInfoApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfGetBundleInfoApiResponseData()
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
REST response for all response codes in api path /lcm/v4.0.a1/resources/bundles Get operation
*/
type GetBundlesApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfGetBundlesApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewGetBundlesApiResponse() *GetBundlesApiResponse {
  p := new(GetBundlesApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.resources.GetBundlesApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.resources.GetBundlesApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *GetBundlesApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *GetBundlesApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfGetBundlesApiResponseData()
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
REST response for all response codes in api path /lcm/v4.0.a1/resources/entities Get operation
*/
type GetEntitiesApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfGetEntitiesApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewGetEntitiesApiResponse() *GetEntitiesApiResponse {
  p := new(GetEntitiesApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.resources.GetEntitiesApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.resources.GetEntitiesApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *GetEntitiesApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *GetEntitiesApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfGetEntitiesApiResponseData()
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
REST response for all response codes in api path /lcm/v4.0.a1/resources/entities/{uuid} Get operation
*/
type GetEntityApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfGetEntityApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewGetEntityApiResponse() *GetEntityApiResponse {
  p := new(GetEntityApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.resources.GetEntityApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.resources.GetEntityApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *GetEntityApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *GetEntityApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfGetEntityApiResponseData()
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
REST response for all response codes in api path /lcm/v4.0.a1/resources/history Get operation
*/
type GetHistoryApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfGetHistoryApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewGetHistoryApiResponse() *GetHistoryApiResponse {
  p := new(GetHistoryApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.resources.GetHistoryApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.resources.GetHistoryApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *GetHistoryApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *GetHistoryApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfGetHistoryApiResponseData()
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
REST response for all response codes in api path /lcm/v4.0.a1/resources/history/{uuid} Get operation
*/
type GetHistoryByUuidApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfGetHistoryByUuidApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewGetHistoryByUuidApiResponse() *GetHistoryByUuidApiResponse {
  p := new(GetHistoryByUuidApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.resources.GetHistoryByUuidApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.resources.GetHistoryByUuidApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *GetHistoryByUuidApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *GetHistoryByUuidApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfGetHistoryByUuidApiResponseData()
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
REST response for all response codes in api path /lcm/v4.0.a1/resources/images Get operation
*/
type GetImagesApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfGetImagesApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewGetImagesApiResponse() *GetImagesApiResponse {
  p := new(GetImagesApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.resources.GetImagesApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.resources.GetImagesApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *GetImagesApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *GetImagesApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfGetImagesApiResponseData()
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
REST response for all response codes in api path /lcm/v4.0.a1/resources/config Get operation
*/
type GetLcmConfigApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfGetLcmConfigApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewGetLcmConfigApiResponse() *GetLcmConfigApiResponse {
  p := new(GetLcmConfigApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.resources.GetLcmConfigApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.resources.GetLcmConfigApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *GetLcmConfigApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *GetLcmConfigApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfGetLcmConfigApiResponseData()
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
REST response for all response codes in api path /lcm/v4.0.a1/resources/status Get operation
*/
type GetLcmStatusApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfGetLcmStatusApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewGetLcmStatusApiResponse() *GetLcmStatusApiResponse {
  p := new(GetLcmStatusApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.resources.GetLcmStatusApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.resources.GetLcmStatusApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *GetLcmStatusApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *GetLcmStatusApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfGetLcmStatusApiResponseData()
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
REST response for all response codes in api path /lcm/v4.0.a1/resources/logbay/items Get operation
*/
type GetLogBayItemsApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfGetLogBayItemsApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewGetLogBayItemsApiResponse() *GetLogBayItemsApiResponse {
  p := new(GetLogBayItemsApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.resources.GetLogBayItemsApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.resources.GetLogBayItemsApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *GetLogBayItemsApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *GetLogBayItemsApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfGetLogBayItemsApiResponseData()
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
REST response for all response codes in api path /lcm/v4.0.a1/resources/logbay/tags Get operation
*/
type GetLogBayTagsApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfGetLogBayTagsApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewGetLogBayTagsApiResponse() *GetLogBayTagsApiResponse {
  p := new(GetLogBayTagsApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.resources.GetLogBayTagsApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.resources.GetLogBayTagsApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *GetLogBayTagsApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *GetLogBayTagsApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfGetLogBayTagsApiResponseData()
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
REST response for all response codes in api path /lcm/v4.0.a1/resources/moduleConfig Get operation
*/
type GetModuleConfigApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfGetModuleConfigApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewGetModuleConfigApiResponse() *GetModuleConfigApiResponse {
  p := new(GetModuleConfigApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.resources.GetModuleConfigApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.resources.GetModuleConfigApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *GetModuleConfigApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *GetModuleConfigApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfGetModuleConfigApiResponseData()
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
REST response for all response codes in api path /lcm/v4.0.a1/resources/config/node-priorities Get operation
*/
type GetNodePrioritiesApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfGetNodePrioritiesApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewGetNodePrioritiesApiResponse() *GetNodePrioritiesApiResponse {
  p := new(GetNodePrioritiesApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.resources.GetNodePrioritiesApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.resources.GetNodePrioritiesApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *GetNodePrioritiesApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *GetNodePrioritiesApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfGetNodePrioritiesApiResponseData()
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
REST response for all response codes in api path /lcm/v4.0.a1/resources/config/node_priority/{nodeUuid} Get operation
*/
type GetNodePriorityNodeApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfGetNodePriorityNodeApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewGetNodePriorityNodeApiResponse() *GetNodePriorityNodeApiResponse {
  p := new(GetNodePriorityNodeApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.resources.GetNodePriorityNodeApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.resources.GetNodePriorityNodeApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *GetNodePriorityNodeApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *GetNodePriorityNodeApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfGetNodePriorityNodeApiResponseData()
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
REST response for all response codes in api path /lcm/v4.0.a1/resources/tasks/{uuid} Get operation
*/
type GetTaskByUuidApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfGetTaskByUuidApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewGetTaskByUuidApiResponse() *GetTaskByUuidApiResponse {
  p := new(GetTaskByUuidApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.resources.GetTaskByUuidApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.resources.GetTaskByUuidApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *GetTaskByUuidApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *GetTaskByUuidApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfGetTaskByUuidApiResponseData()
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
REST response for all response codes in api path /lcm/v4.0.a1/resources/recommendations Post operation
*/
type GetUpdateRecApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfGetUpdateRecApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewGetUpdateRecApiResponse() *GetUpdateRecApiResponse {
  p := new(GetUpdateRecApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.resources.GetUpdateRecApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.resources.GetUpdateRecApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *GetUpdateRecApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *GetUpdateRecApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfGetUpdateRecApiResponseData()
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
LCM upgrade plan for a node or cluster
*/
type HostUpgradePlan struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  The description of the most disruptive action that will occur on the node or cluster
  */
  Action *string `json:"action,omitempty"`
  
  HypervisorType *import3.HypervisorType `json:"hypervisorType,omitempty"`
  /**
  The location of the LCM entity, it can be at the cluster level or at the node level, in the format of "cluster:uuid" or "node:uuid"
  */
  LocationId *string `json:"locationId,omitempty"`
}


func NewHostUpgradePlan() *HostUpgradePlan {
  p := new(HostUpgradePlan)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.resources.HostUpgradePlan"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.resources.HostUpgradePlan"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Available version for an LCM entity to update
*/
type LcmAvailableVersion struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Component information for the payload based entity
  */
  ChildEntities *interface{} `json:"childEntities,omitempty"`
  /**
  List of dependencies for the available version.
  */
  Dependencies []LcmDependency `json:"dependencies,omitempty"`
  /**
  Reason for disabling the available version
  */
  DisableReason *string `json:"disableReason,omitempty"`
  /**
  Indicates if the available update is enabled
  */
  IsEnabled *bool `json:"isEnabled,omitempty"`
  /**
  Order of this available version (1 being the lowest and 6 being the highest) when multiple versions are present with different status.
  */
  Order *int64 `json:"order,omitempty"`
  /**
  Update metadata other than the release notes - like a custom message
  */
  OtherMetadata *string `json:"otherMetadata,omitempty"`
  /**
  Release date for the entities that need this information
  */
  ReleaseDate *string `json:"releaseDate,omitempty"`
  /**
  Release notes corresponding to the update
  */
  ReleaseNotes *string `json:"releaseNotes,omitempty"`
  /**
  UUID of the group that this LCM entity is part of
  */
  SingleGroupUuid *string `json:"singleGroupUuid,omitempty"`
  
  Status *import3.LcmAvailableVersionStatus `json:"status,omitempty"`
  /**
  UUID of the entity to be used
  */
  Uuid *string `json:"uuid,omitempty"`
  /**
  Available version for the LCM entity
  */
  Version *string `json:"version,omitempty"`
}


func NewLcmAvailableVersion() *LcmAvailableVersion {
  p := new(LcmAvailableVersion)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.resources.LcmAvailableVersion"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.resources.LcmAvailableVersion"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Details of the LCM bundle
*/
type LcmBundle struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Details of the LCM image
  */
  Images []LcmImageDetails `json:"images,omitempty"`
  /**
  Name of the LCM bundle
  */
  Name *string `json:"name,omitempty"`
  /**
  Size of the LCM bundle
  */
  Size *string `json:"size,omitempty"`
  /**
  LCM bundle type
  */
  Type *string `json:"type,omitempty"`
  /**
  UUID of the LCM bundle
  */
  Uuid *string `json:"uuid,omitempty"`
}


func NewLcmBundle() *LcmBundle {
  p := new(LcmBundle)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.resources.LcmBundle"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.resources.LcmBundle"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
LCM bundle creation request type
*/
type LcmBundlesSpec struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Indicates whether the LCM bundle is for a third party image or not
  */
  IsThirdParty *bool `json:"isThirdParty,omitempty"`
  
  RedistributableSpec *LcmRedistributableUploadSpec `json:"redistributableSpec,omitempty"`
  
  ThirdPartySpec *LcmThirdPartyUploadSpec `json:"thirdPartySpec,omitempty"`
}


func NewLcmBundlesSpec() *LcmBundlesSpec {
  p := new(LcmBundlesSpec)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.resources.LcmBundlesSpec"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.resources.LcmBundlesSpec"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
LCM bundle creation response
*/
type LcmBundlesStatus struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  UUID of the LCM bundle
  */
  BundleUuid *string `json:"bundleUuid,omitempty"`
  
  TaskUuid *import4.TaskReference `json:"taskUuid,omitempty"`
}


func NewLcmBundlesStatus() *LcmBundlesStatus {
  p := new(LcmBundlesStatus)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.resources.LcmBundlesStatus"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.resources.LcmBundlesStatus"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Intent to cancel an ongoing LCM update
*/
type LcmCancelIntent struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Boolean that indicates if cancel intent for LCM update is set or not
  */
  IsSet *bool `json:"isSet,omitempty"`
}


func NewLcmCancelIntent() *LcmCancelIntent {
  p := new(LcmCancelIntent)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.resources.LcmCancelIntent"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.resources.LcmCancelIntent"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
LCM configuration on the cluster
*/
type LcmConfig struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Indicates if REST APIs for LCM is enabled. The default value is True
  */
  ApiEnabled *bool `json:"apiEnabled,omitempty"`
  /**
  Indicates if the auto inventory operation is enabled. The default value is set to False
  */
  AutoInventoryEnabled *bool `json:"autoInventoryEnabled,omitempty"`
  /**
  The scheduled time in "%H:%M" 24-hour format of the next inventory execution. Used when auto_inventory_enabled is set to True. The default schedule time is 03:00(AM)
  */
  AutoInventoryScheduledTime *string `json:"autoInventoryScheduledTime,omitempty"`
  /**
  Indicates whether the LCM binaries in the URL are of type connected-site or dark-site
  */
  BuildType *string `json:"buildType,omitempty"`
  /**
  List of entities for which One-Click upgrades are not available
  */
  DeprecatedSoftwareEntities []string `json:"deprecatedSoftwareEntities,omitempty"`
  /**
  Indicates if LCM is enabled to distribute inventory across the cluster. The default value is True
  */
  DistributeInventory *bool `json:"distributeInventory,omitempty"`
  /**
  Indicates if automatic log collection on failures is enabled or not. The default value is True
  */
  EnableAutoLogCollection *bool `json:"enableAutoLogCollection,omitempty"`
  /**
  Indicates if the LCM URL has HTTPS enabled. The default value is True
  */
  EnableHttps *bool `json:"enableHttps,omitempty"`
  /**
  Number of parallel threads running inventory. The default value is set to 8
  */
  InventoryParallelLimit *int `json:"inventoryParallelLimit,omitempty"`
  /**
  Indicates if LCM is pointing to a dark site. It is set to True if the LCM URL is not pointing to the default Nutanix portal URL.
  */
  IsDarksite *bool `json:"isDarksite,omitempty"`
  /**
  LCM version installed on the cluster
  */
  LcmVersion *string `json:"lcmVersion,omitempty"`
  /**
  Indicates if metrics for LCM is enabled. The default value is True
  */
  MetricsEnabled *bool `json:"metricsEnabled,omitempty"`
  /**
  Indicates if LCM is enabled to auto-upgrade products. The default value is False.
  */
  ModuleAutoUpgradeEnabled *bool `json:"moduleAutoUpgradeEnabled,omitempty"`
  /**
  Timeout to send GET/POST requests to Prism
  */
  PrismApiTimeout *int `json:"prismApiTimeout,omitempty"`
  /**
  Indicates if Product Meta filtering is disabled. The default value is False.
  */
  ProductMetaDisabled *bool `json:"productMetaDisabled,omitempty"`
  /**
  URL for Product Meta artifact
  */
  ProductMetaUrl *string `json:"productMetaUrl,omitempty"`
  /**
  LCM can automatically recover a benign failure in the early stages of an upgrade. This value indicates if that feature is enabled. The default value is False.
  */
  RecoveryDisabled *bool `json:"recoveryDisabled,omitempty"`
  /**
  Indicates if Redfish out of band update is enabled. The default value is False.
  */
  RedfishOOBUpdateSupport *bool `json:"redfishOOBUpdateSupport,omitempty"`
  /**
  LCM semantic version installed on the cluster
  */
  SemanticVersion *string `json:"semanticVersion,omitempty"`
  /**
  List of skipped prechecks.
  */
  SkippedPrecheckNames []string `json:"skippedPrecheckNames,omitempty"`
  /**
  List of entities for which One-Click upgrades are supported
  */
  SupportedSoftwareEntities []string `json:"supportedSoftwareEntities,omitempty"`
  /**
  LCM UI version installed on the cluster
  */
  UiVersion *string `json:"uiVersion,omitempty"`
  /**
  Indicates if the bundle is uploaded or not.
  */
  UploadedBundle *bool `json:"uploadedBundle,omitempty"`
  /**
  URL of the LCM repository.
  */
  Url *string `json:"url,omitempty"`
  /**
  Indicates the user preferences.
  */
  UserPreferences []UserPreference `json:"userPreferences,omitempty"`
}


func NewLcmConfig() *LcmConfig {
  p := new(LcmConfig)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.resources.LcmConfig"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.resources.LcmConfig"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Dependency of an LCM entity available version.
*/
type LcmDependency struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  LCM entity class
  */
  EntityClass *string `json:"entityClass,omitempty"`
  /**
  LCM entity model
  */
  EntityModel *string `json:"entityModel,omitempty"`
  /**
  Information of the dependent entity versions for this available entity
  */
  EntityVersions *interface{} `json:"entityVersions,omitempty"`
  /**
  Available version for the LCM entity
  */
  Version *string `json:"version,omitempty"`
}


func NewLcmDependency() *LcmDependency {
  p := new(LcmDependency)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.resources.LcmDependency"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.resources.LcmDependency"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Details of an LCM entity
*/
type LcmEntity struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  List of available versions for an LCM entity to update
  */
  AvailableVersions []LcmAvailableVersion `json:"availableVersions,omitempty"`
  /**
  Component information for the payload based entity
  */
  ChildEntities *interface{} `json:"childEntities,omitempty"`
  /**
  LCM entity class
  */
  EntityClass *string `json:"entityClass,omitempty"`
  /**
  Description of an LCM entity
  */
  EntityDescription *string `json:"entityDescription,omitempty"`
  /**
  Detailed information for the LCM entity. For example, firmware entities contain additional information about NIC and so on
  */
  EntityDetails *interface{} `json:"entityDetails,omitempty"`
  /**
  LCM entity model
  */
  EntityModel *string `json:"entityModel,omitempty"`
  
  EntityType *import3.LcmEntityType `json:"entityType,omitempty"`
  /**
  Unique identifier of an LCM entity e.g. "HDD serial number"
  */
  Id *string `json:"id,omitempty"`
  /**
  UTC date and time in RFC-3339 format when the task was last updated.
  */
  LastUpdatedTime *string `json:"lastUpdatedTime,omitempty"`
  /**
  The location of the LCM entity, it can be at the cluster level or at the node level, in the format of "cluster:uuid" or "node:uuid"
  */
  LocationId *string `json:"locationId,omitempty"`
  /**
  The requested update version of an LCM entity.
  */
  RequestVersion *string `json:"requestVersion,omitempty"`
  /**
  UUID of the group that this LCM entity is part of
  */
  SingleGroupUuid *string `json:"singleGroupUuid,omitempty"`
  /**
  A list of sub-entities applicable to the entity.
  */
  SubEntities []SubEntity `json:"subEntities,omitempty"`
  /**
  UUID of the entity to be used
  */
  Uuid *string `json:"uuid,omitempty"`
  /**
  Current version of an LCM entity
  */
  Version *string `json:"version,omitempty"`
}


func NewLcmEntity() *LcmEntity {
  p := new(LcmEntity)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.resources.LcmEntity"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.resources.LcmEntity"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
LCM framework version information
*/
type LcmFrameworkVersionInfo struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  LCM version installed on the cluster
  */
  LocalVersion *string `json:"localVersion,omitempty"`
  /**
  LCM framework version present in the LCM URL
  */
  RemoteVersion *string `json:"remoteVersion,omitempty"`
  /**
  Boolean that indicates if LCM framework update is needed
  */
  UpdateNeeded *bool `json:"updateNeeded,omitempty"`
}


func NewLcmFrameworkVersionInfo() *LcmFrameworkVersionInfo {
  p := new(LcmFrameworkVersionInfo)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.resources.LcmFrameworkVersionInfo"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.resources.LcmFrameworkVersionInfo"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
LCM image (List of LCM image files)
*/
type LcmImage struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  LCM entity class
  */
  EntityClass *string `json:"entityClass,omitempty"`
  /**
  LCM entity model
  */
  EntityModel *string `json:"entityModel,omitempty"`
  
  EntityType *import3.LcmEntityType `json:"entityType,omitempty"`
  /**
  List of files in the image
  */
  Files []LcmImageFile `json:"files,omitempty"`
  /**
  A hardware family for a LCM entity
  */
  HardwareFamily *string `json:"hardwareFamily,omitempty"`
  /**
  Current version of an LCM entity
  */
  Version *string `json:"version,omitempty"`
}


func NewLcmImage() *LcmImage {
  p := new(LcmImage)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.resources.LcmImage"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.resources.LcmImage"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Details of the LCM image
*/
type LcmImageDetails struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  LCM entity class
  */
  EntityClass *string `json:"entityClass,omitempty"`
  /**
  LCM entity model
  */
  EntityModel *string `json:"entityModel,omitempty"`
  /**
  Release notes for the LCM image
  */
  ReleaseNotes *string `json:"releaseNotes,omitempty"`
  /**
  UUID of the entity to be used
  */
  Uuid *string `json:"uuid,omitempty"`
  /**
  LCM Image version
  */
  Version *string `json:"version,omitempty"`
}


func NewLcmImageDetails() *LcmImageDetails {
  p := new(LcmImageDetails)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.resources.LcmImageDetails"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.resources.LcmImageDetails"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Description of LCM image file
*/
type LcmImageFile struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Image file global catalog item UUID.
  */
  GlobalCatalogItemUuid *string `json:"globalCatalogItemUuid,omitempty"`
  /**
  LCM image file name
  */
  Name *string `json:"name,omitempty"`
  /**
  NFS path for the LCM image
  */
  NfsPath *string `json:"nfsPath,omitempty"`
  /**
  LCM image SHA
  */
  Shasum *string `json:"shasum,omitempty"`
  /**
  LCM image file size
  */
  SizeBytes *int64 `json:"sizeBytes,omitempty"`
}


func NewLcmImageFile() *LcmImageFile {
  p := new(LcmImageFile)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.resources.LcmImageFile"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.resources.LcmImageFile"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
List of LCM images
*/
type LcmImages struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  List of LCM images
  */
  Images []LcmImage `json:"images,omitempty"`
}


func NewLcmImages() *LcmImages {
  p := new(LcmImages)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.resources.LcmImages"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.resources.LcmImages"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Operation type and UUID of an ongoing operation in LCM
*/
type LcmInProgressOp struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Type of the operation tracked by the task.
  */
  Type *string `json:"type,omitempty"`
  /**
  Root task UUID of the operation, if it is in running state
  */
  Uuid *string `json:"uuid,omitempty"`
}


func NewLcmInProgressOp() *LcmInProgressOp {
  p := new(LcmInProgressOp)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.resources.LcmInProgressOp"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.resources.LcmInProgressOp"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Type of LCM operation that was run
*/
type LcmOpType int

const(
  LCMOPTYPE_UNKNOWN LcmOpType = 0
  LCMOPTYPE_REDACTED LcmOpType = 1
  LCMOPTYPE_KLCMPRECHECKOPERATION LcmOpType = 2
  LCMOPTYPE_KLCMDOWNLOADOPERATION LcmOpType = 3
  LCMOPTYPE_KLCMINVENTORYOPERATION LcmOpType = 4
  LCMOPTYPE_KLCMUPDATEOPERATION LcmOpType = 5
)

// returns the name of the enum given an ordinal number
func (e *LcmOpType) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "kLcmPrecheckOperation",
    "kLcmDownloadOperation",
    "kLcmInventoryOperation",
    "kLcmUpdateOperation",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *LcmOpType) index(name string) LcmOpType {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "kLcmPrecheckOperation",
    "kLcmDownloadOperation",
    "kLcmInventoryOperation",
    "kLcmUpdateOperation",
  }
  for idx := range names {
    if names[idx] == name {
      return LcmOpType(idx)
    }
  }
  return LCMOPTYPE_UNKNOWN
}

func (e *LcmOpType) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for LcmOpType:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *LcmOpType) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e LcmOpType) Ref() *LcmOpType {
  return &e
}


/**
Upload specification for a Nutanix redistributable bundle
*/
type LcmRedistributableUploadSpec struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Name of the LCM bundle
  */
  Name *string `json:"name,omitempty"`
}


func NewLcmRedistributableUploadSpec() *LcmRedistributableUploadSpec {
  p := new(LcmRedistributableUploadSpec)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.resources.LcmRedistributableUploadSpec"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.resources.LcmRedistributableUploadSpec"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Details of LCM framework status
*/
type LcmStatus struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  CancelIntent *LcmCancelIntent `json:"cancelIntent,omitempty"`
  
  FrameworkVersion *LcmFrameworkVersionInfo `json:"frameworkVersion,omitempty"`
  
  InProgressOperation *LcmInProgressOp `json:"inProgressOperation,omitempty"`
  
  UploadInfo *LcmUploadInfo `json:"uploadInfo,omitempty"`
}


func NewLcmStatus() *LcmStatus {
  p := new(LcmStatus)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.resources.LcmStatus"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.resources.LcmStatus"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Details about LCM Task
*/
type LcmTask struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Reference to the cluster the task is being executed on.
  */
  ClusterUuid *string `json:"clusterUuid,omitempty"`
  /**
  End time for an LCM operation
  */
  CompletionTime *string `json:"completionTime,omitempty"`
  /**
  UTC date and time in RFC-3339 format when the task was created.
  */
  CreationTime *string `json:"creationTime,omitempty"`
  /**
  Time in microseconds from Epoch when task was created.
  */
  CreationTimeUsecs *int64 `json:"creationTimeUsecs,omitempty"`
  /**
  List of entity identifiers being upgraded, in UUID format
  */
  EntityUuids []string `json:"entityUuids,omitempty"`
  /**
  In case of task failure this field will provide the error code.
  */
  ErrorCode *string `json:"errorCode,omitempty"`
  /**
  In case of task failure this field will provide the error description.
  */
  ErrorDetail *string `json:"errorDetail,omitempty"`
  /**
  UTC date and time in RFC-3339 format when the task was last updated.
  */
  LastUpdateTime *string `json:"lastUpdateTime,omitempty"`
  /**
  Number of times the task has been updated. The value increases sequentially with each update of the task and can be used to verify if there have been changes to the task.
  */
  LogicalTimestamp *int64 `json:"logicalTimestamp,omitempty"`
  /**
  Any custom message related to the task
  */
  Message *string `json:"message,omitempty"`
  
  OperationType *OperationType `json:"operationType,omitempty"`
  /**
  Reference to the parent task
  */
  ParentTaskUuid *string `json:"parentTaskUuid,omitempty"`
  /**
  Completion percentage of a task
  */
  PercentageComplete *int `json:"percentageComplete,omitempty"`
  /**
  Start time for an LCM operation
  */
  StartTime *string `json:"startTime,omitempty"`
  
  Status *StatusType `json:"status,omitempty"`
  /**
  List with the reference to the sub-task UUIDs
  */
  SubtaskUuids []string `json:"subtaskUuids,omitempty"`
  /**
  UUID of LCM task
  */
  Uuid *string `json:"uuid,omitempty"`
}


func NewLcmTask() *LcmTask {
  p := new(LcmTask)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.resources.LcmTask"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.resources.LcmTask"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Upload specification for a third party bundle
*/
type LcmThirdPartyUploadSpec struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  LCM entity class
  */
  EntityClass *string `json:"entityClass,omitempty"`
  /**
  LCM entity model
  */
  EntityModel *string `json:"entityModel,omitempty"`
  /**
  A hardware family for a LCM entity
  */
  HwFamilyList *string `json:"hwFamilyList,omitempty"`
  
  ImageDetails []ThirdPartyImageFileDetailsSpec `json:"imageDetails,omitempty"`
  /**
  Release notes for the LCM image
  */
  ReleaseNotes *string `json:"releaseNotes,omitempty"`
  /**
  Version specification for the third party upload metadata
  */
  SpecVersion *string `json:"specVersion,omitempty"`
  
  Status *import3.LcmAvailableVersionStatus `json:"status,omitempty"`
  /**
  LCM Image version
  */
  Version *string `json:"version,omitempty"`
}


func NewLcmThirdPartyUploadSpec() *LcmThirdPartyUploadSpec {
  p := new(LcmThirdPartyUploadSpec)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.resources.LcmThirdPartyUploadSpec"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.resources.LcmThirdPartyUploadSpec"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
LCM upload information
*/
type LcmUploadInfo struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Upload task UUID
  */
  TaskUuid *string `json:"taskUuid,omitempty"`
}


func NewLcmUploadInfo() *LcmUploadInfo {
  p := new(LcmUploadInfo)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.resources.LcmUploadInfo"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.resources.LcmUploadInfo"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
List of NCC Logbay items applicable for LCM log collection
*/
type LogbayItem struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  NCC Logbay item applicable for LCM log collection
  */
  LogbayItem *string `json:"logbayItem,omitempty"`
}


func NewLogbayItem() *LogbayItem {
  p := new(LogbayItem)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.resources.LogbayItem"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.resources.LogbayItem"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
List of NCC Logbay tags applicable for LCM log collection
*/
type LogbayTag struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  NCC Logbay tag applicable for LCM log collection
  */
  LogbayTag *string `json:"logbayTag,omitempty"`
}


func NewLogbayTag() *LogbayTag {
  p := new(LogbayTag)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.resources.LogbayTag"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.resources.LogbayTag"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
LCM module Configuration
*/
type ModuleConfig struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Indicates if redfish disabled. Default is False
  */
  RedfishDisabled *bool `json:"redfishDisabled,omitempty"`
  /**
  List of repository image module (RIM) parameters
  */
  RimParams []RimParams `json:"rimParams,omitempty"`
}


func NewModuleConfig() *ModuleConfig {
  p := new(ModuleConfig)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.resources.ModuleConfig"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.resources.ModuleConfig"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
LCM node priority config
*/
type NodePriorityConfig struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Priority node map where the key is node UUID and value is node priority
  */
  PriorityNodeMap *interface{} `json:"priorityNodeMap,omitempty"`
}


func NewNodePriorityConfig() *NodePriorityConfig {
  p := new(NodePriorityConfig)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.resources.NodePriorityConfig"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.resources.NodePriorityConfig"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Detailed LCM upgrade notification information for this entity
*/
type Notification struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  A message with notification details. The description of the most disruptive action that will occur on the node or the cluster. INFO, WARNING or NOTICE based message for an entity
  */
  Message *string `json:"message,omitempty"`
  /**
  Severity level for a notification. Permissible values are INFO, WARNING and NOTICE
  */
  SeverityLevel *string `json:"severityLevel,omitempty"`
}


func NewNotification() *Notification {
  p := new(Notification)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.resources.Notification"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.resources.Notification"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
LCM upgrade notification for a node or cluster
*/
type NotificationsSpec struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  LCM entity class
  */
  EntityClass *string `json:"entityClass,omitempty"`
  /**
  LCM entity model
  */
  EntityModel *string `json:"entityModel,omitempty"`
  /**
  Current version of an LCM entity
  */
  FromVersion *string `json:"fromVersion,omitempty"`
  
  HypervisorType *import3.HypervisorType `json:"hypervisorType,omitempty"`
  /**
  The location of the LCM entity, it can be at the cluster level or at the node level, in the format of "cluster:uuid" or "node:uuid"
  */
  LocationId *string `json:"locationId,omitempty"`
  /**
  Type of LCM upgrade notification generated. It can be any one of the values like Entity, Location, Generic or Workflow. The only types supported today are Entity and Location
  */
  NotificationType *string `json:"notificationType,omitempty"`
  /**
  List of notifications for this entity
  */
  Notifications []Notification `json:"notifications,omitempty"`
  /**
  The requested update version of an LCM entity.
  */
  ToVersion *string `json:"toVersion,omitempty"`
}


func NewNotificationsSpec() *NotificationsSpec {
  p := new(NotificationsSpec)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.resources.NotificationsSpec"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.resources.NotificationsSpec"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Environment in which the operation was run. For example, CVM, Phoenix, Host etc
*/
type OperationEnv int

const(
  OPERATIONENV_UNKNOWN OperationEnv = 0
  OPERATIONENV_REDACTED OperationEnv = 1
  OPERATIONENV_CVM OperationEnv = 2
  OPERATIONENV_HOST OperationEnv = 3
  OPERATIONENV_PHOENIX OperationEnv = 4
  OPERATIONENV_PC OperationEnv = 5
)

// returns the name of the enum given an ordinal number
func (e *OperationEnv) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "CVM",
    "Host",
    "Phoenix",
    "PC",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *OperationEnv) index(name string) OperationEnv {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "CVM",
    "Host",
    "Phoenix",
    "PC",
  }
  for idx := range names {
    if names[idx] == name {
      return OperationEnv(idx)
    }
  }
  return OPERATIONENV_UNKNOWN
}

func (e *OperationEnv) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for OperationEnv:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *OperationEnv) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e OperationEnv) Ref() *OperationEnv {
  return &e
}


/**
Details about an LCM operation that was run
*/
type OperationHistoryElement struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  End time for an LCM operation
  */
  EndTime *string `json:"endTime,omitempty"`
  
  EntityInfo []EntityInfo `json:"entityInfo,omitempty"`
  /**
  LCM version installed on the cluster
  */
  LcmVersion *string `json:"lcmVersion,omitempty"`
  /**
  Details of the operation failure
  */
  OpExceptionReport *string `json:"opExceptionReport,omitempty"`
  /**
  Boolean that indicates whether the operation was done
  */
  OperationDone *bool `json:"operationDone,omitempty"`
  /**
  UUID of an LCM operation
  */
  OperationId *string `json:"operationId,omitempty"`
  
  OperationType *LcmOpType `json:"operationType,omitempty"`
  /**
  Start time for an LCM operation
  */
  StartTime *string `json:"startTime,omitempty"`
}


func NewOperationHistoryElement() *OperationHistoryElement {
  p := new(OperationHistoryElement)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.resources.OperationHistoryElement"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.resources.OperationHistoryElement"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Details about an LCM operation that was run
*/
type OperationHistoryItem struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  End time for an LCM operation
  */
  EndTime *string `json:"endTime,omitempty"`
  
  EntityInfo []EntityInfo `json:"entityInfo,omitempty"`
  /**
  LCM version installed on the cluster
  */
  LcmVersion *string `json:"lcmVersion,omitempty"`
  /**
  Details of the operation failure
  */
  OpExceptionReport *string `json:"opExceptionReport,omitempty"`
  /**
  Boolean that indicates whether the operation was done
  */
  OperationDone *bool `json:"operationDone,omitempty"`
  /**
  UUID of an LCM operation
  */
  OperationId *string `json:"operationId,omitempty"`
  
  OperationType *LcmOpType `json:"operationType,omitempty"`
  /**
  Start time for an LCM operation
  */
  StartTime *string `json:"startTime,omitempty"`
}


func NewOperationHistoryItem() *OperationHistoryItem {
  p := new(OperationHistoryItem)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.resources.OperationHistoryItem"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.resources.OperationHistoryItem"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Type of the operation tracked by the task.
*/
type OperationType int

const(
  OPERATIONTYPE_UNKNOWN OperationType = 0
  OPERATIONTYPE_REDACTED OperationType = 1
  OPERATIONTYPE_LCMROOTTASK OperationType = 2
  OPERATIONTYPE_LCMDOWNLOADTASK OperationType = 3
  OPERATIONTYPE_LCMINVENTORYTASK OperationType = 4
  OPERATIONTYPE_LCMUPDATECLUSTERTASK OperationType = 5
  OPERATIONTYPE_LCMUPDATECLUSTERNODETASK OperationType = 6
  OPERATIONTYPE_LCMUPDATENODETASK OperationType = 7
  OPERATIONTYPE_LCMPRECHECKSTASK OperationType = 8
  OPERATIONTYPE_LCMPRODUCTMETATASK OperationType = 9
)

// returns the name of the enum given an ordinal number
func (e *OperationType) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "LcmRootTask",
    "LcmDownloadTask",
    "LcmInventoryTask",
    "LcmUpdateClusterTask",
    "LcmUpdateClusterNodeTask",
    "LcmUpdateNodeTask",
    "LcmPrechecksTask",
    "LcmProductMetaTask",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *OperationType) index(name string) OperationType {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "LcmRootTask",
    "LcmDownloadTask",
    "LcmInventoryTask",
    "LcmUpdateClusterTask",
    "LcmUpdateClusterNodeTask",
    "LcmUpdateNodeTask",
    "LcmPrechecksTask",
    "LcmProductMetaTask",
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


/**
REST response for all response codes in api path /lcm/v4.0.a1/resources/bundles Post operation
*/
type PostBundlesApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfPostBundlesApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewPostBundlesApiResponse() *PostBundlesApiResponse {
  p := new(PostBundlesApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.resources.PostBundlesApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.resources.PostBundlesApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *PostBundlesApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *PostBundlesApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfPostBundlesApiResponseData()
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
REST response for all response codes in api path /lcm/v4.0.a1/resources/config/node-priorities Post operation
*/
type PostNodePrioritiesApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfPostNodePrioritiesApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewPostNodePrioritiesApiResponse() *PostNodePrioritiesApiResponse {
  p := new(PostNodePrioritiesApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.resources.PostNodePrioritiesApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.resources.PostNodePrioritiesApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *PostNodePrioritiesApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *PostNodePrioritiesApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfPostNodePrioritiesApiResponseData()
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
REST response for all response codes in api path /lcm/v4.0.a1/resources/bundles/{uuid} Put operation
*/
type PutBundleApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfPutBundleApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewPutBundleApiResponse() *PutBundleApiResponse {
  p := new(PutBundleApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.resources.PutBundleApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.resources.PutBundleApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *PutBundleApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *PutBundleApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfPutBundleApiResponseData()
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
REST response for all response codes in api path /lcm/v4.0.a1/resources/config Put operation
*/
type PutLcmConfigApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfPutLcmConfigApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewPutLcmConfigApiResponse() *PutLcmConfigApiResponse {
  p := new(PutLcmConfigApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.resources.PutLcmConfigApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.resources.PutLcmConfigApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *PutLcmConfigApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *PutLcmConfigApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfPutLcmConfigApiResponseData()
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
REST response for all response codes in api path /lcm/v4.0.a1/resources/moduleConfig Put operation
*/
type PutModuleConfigApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfPutModuleConfigApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewPutModuleConfigApiResponse() *PutModuleConfigApiResponse {
  p := new(PutModuleConfigApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.resources.PutModuleConfigApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.resources.PutModuleConfigApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *PutModuleConfigApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *PutModuleConfigApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfPutModuleConfigApiResponseData()
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
LCM update recommendations
*/
type RecommendationResult struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  List of added LCM entities to the input recommendations specification
  */
  AddedEntities []UpdatedTargetEntityResult `json:"addedEntities,omitempty"`
  /**
  List of deployable entities and their dependencies
  */
  DeployableVersions []DeployableVersion `json:"deployableVersions,omitempty"`
  
  EntityUpdateSpecs *import3.EntityUpdateSpecs `json:"entityUpdateSpecs,omitempty"`
  /**
  List of modified LCM entities from the input recommendations specification
  */
  ModifiedVersions []UpdatedTargetEntityResult `json:"modifiedVersions,omitempty"`
  /**
  List of skipped LCM entities from the input recommendations specification
  */
  SkippedEntities []UpdatedTargetEntityResult `json:"skippedEntities,omitempty"`
}


func NewRecommendationResult() *RecommendationResult {
  p := new(RecommendationResult)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.resources.RecommendationResult"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.resources.RecommendationResult"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
LCM update recommendation specification
*/
type RecommendationSpec struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  List of deploy specifications for use in the recommendations API
  */
  EntityDeploySpecs []import3.EntityIdentifierSpec `json:"entityDeploySpecs,omitempty"`
  /**
  Type of an LCM entity
  */
  EntityTypes []import3.LcmEntityType `json:"entityTypes,omitempty"`
  /**
  List of entity update objects for getting recommendations
  */
  EntityUpdateSpecs []import3.EntityUpdateSpec `json:"entityUpdateSpecs,omitempty"`
  /**
  List of target entity objects for getting recommendations
  */
  TargetEntities []TargetEntity `json:"targetEntities,omitempty"`
}


func NewRecommendationSpec() *RecommendationSpec {
  p := new(RecommendationSpec)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.resources.RecommendationSpec"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.resources.RecommendationSpec"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
RIM parameters for one component
*/
type RimParams struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Component name of the RIM
  */
  Component *string `json:"component,omitempty"`
  /**
  Description of the RIM
  */
  Description *string `json:"description,omitempty"`
  /**
  Flag list of the RIM
  */
  Flags []string `json:"flags,omitempty"`
  /**
  Path of the RIM
  */
  Path *string `json:"path,omitempty"`
  /**
  Tag list of the RIM
  */
  Tags []string `json:"tags,omitempty"`
  /**
  URL of the RIM.
  */
  Url *string `json:"url,omitempty"`
  /**
  Indicates if the RIM URL is set by the user
  */
  UrlSetByUser *bool `json:"urlSetByUser,omitempty"`
}


func NewRimParams() *RimParams {
  p := new(RimParams)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.resources.RimParams"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.resources.RimParams"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Status of a task
*/
type Status struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Description of the current status of the operation
  */
  Description *string `json:"description,omitempty"`
  /**
  Boolean that indicates whether the operation was done
  */
  OperationDone *bool `json:"operationDone,omitempty"`
  /**
  State of the current operation
  */
  State *int `json:"state,omitempty"`
}


func NewStatus() *Status {
  p := new(Status)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.resources.Status"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.resources.Status"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Current state of the task.
*/
type StatusType int

const(
  STATUSTYPE_UNKNOWN StatusType = 0
  STATUSTYPE_REDACTED StatusType = 1
  STATUSTYPE_QUEUED StatusType = 2
  STATUSTYPE_RUNNING StatusType = 3
  STATUSTYPE_SUCCEEDED StatusType = 4
  STATUSTYPE_ABORTED StatusType = 5
  STATUSTYPE_SUSPENDED StatusType = 6
  STATUSTYPE_FAILED StatusType = 7
)

// returns the name of the enum given an ordinal number
func (e *StatusType) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "QUEUED",
    "RUNNING",
    "SUCCEEDED",
    "ABORTED",
    "SUSPENDED",
    "FAILED",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *StatusType) index(name string) StatusType {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "QUEUED",
    "RUNNING",
    "SUCCEEDED",
    "ABORTED",
    "SUSPENDED",
    "FAILED",
  }
  for idx := range names {
    if names[idx] == name {
      return StatusType(idx)
    }
  }
  return STATUSTYPE_UNKNOWN
}

func (e *StatusType) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for StatusType:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *StatusType) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e StatusType) Ref() *StatusType {
  return &e
}


/**
A partial entity that is tied to the larger entity in LCM.
*/
type SubEntity struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  LCM entity class
  */
  EntityClass *string `json:"entityClass,omitempty"`
  /**
  LCM entity model
  */
  EntityModel *string `json:"entityModel,omitempty"`
  /**
  A hardware family for a LCM entity
  */
  HardwareFamily *string `json:"hardwareFamily,omitempty"`
  /**
  The version of the sub-entity.
  */
  Version *string `json:"version,omitempty"`
}


func NewSubEntity() *SubEntity {
  p := new(SubEntity)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.resources.SubEntity"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.resources.SubEntity"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
LCM target entity for which recommendations are requested
*/
type TargetEntity struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  LCM entity class
  */
  EntityClass *string `json:"entityClass,omitempty"`
  /**
  LCM entity model
  */
  EntityModel *string `json:"entityModel,omitempty"`
  /**
  Unique identifier of an LCM entity e.g. "HDD serial number"
  */
  Id *string `json:"id,omitempty"`
  /**
  The location of the LCM entity, it can be at the cluster level or at the node level, in the format of "cluster:uuid" or "node:uuid"
  */
  LocationId *string `json:"locationId,omitempty"`
  /**
  LCM entity update version
  */
  Version *string `json:"version,omitempty"`
}


func NewTargetEntity() *TargetEntity {
  p := new(TargetEntity)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.resources.TargetEntity"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.resources.TargetEntity"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Details about LCM Task
*/
type TaskInfo struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  Status *Status `json:"status,omitempty"`
}


func NewTaskInfo() *TaskInfo {
  p := new(TaskInfo)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.resources.TaskInfo"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.resources.TaskInfo"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
File detail for a third party image
*/
type ThirdPartyImageFileDetailsSpec struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  Checksum *string `json:"checksum,omitempty"`
  /**
  Checksum type for a third party image
  */
  ChecksumType *string `json:"checksumType,omitempty"`
  /**
  File name for a third party image
  */
  Name *string `json:"name,omitempty"`
  /**
  File size for a third party image
  */
  Size *int `json:"size,omitempty"`
}


func NewThirdPartyImageFileDetailsSpec() *ThirdPartyImageFileDetailsSpec {
  p := new(ThirdPartyImageFileDetailsSpec)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.resources.ThirdPartyImageFileDetailsSpec"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.resources.ThirdPartyImageFileDetailsSpec"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Updated LCM target entity in recommendation result
*/
type UpdatedTargetEntity struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  LCM entity class
  */
  EntityClass *string `json:"entityClass,omitempty"`
  /**
  LCM entity model
  */
  EntityModel *string `json:"entityModel,omitempty"`
  /**
  UUID of the entity to be used
  */
  EntityUuid *string `json:"entityUuid,omitempty"`
  /**
  The location of the LCM entity, it can be at the cluster level or at the node level, in the format of "cluster:uuid" or "node:uuid"
  */
  LocationId *string `json:"locationId,omitempty"`
  /**
  LCM entity update version
  */
  Version *string `json:"version,omitempty"`
}


func NewUpdatedTargetEntity() *UpdatedTargetEntity {
  p := new(UpdatedTargetEntity)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.resources.UpdatedTargetEntity"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.resources.UpdatedTargetEntity"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
LCM error target entity
*/
type UpdatedTargetEntityResult struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  UUID of the entity to be used
  */
  EntityUuid *string `json:"entityUuid,omitempty"`
  /**
  Error message for the target entity that failed in the input recommendations specification
  */
  Message *string `json:"message,omitempty"`
  
  TargetEntity *UpdatedTargetEntity `json:"targetEntity,omitempty"`
}


func NewUpdatedTargetEntityResult() *UpdatedTargetEntityResult {
  p := new(UpdatedTargetEntityResult)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.resources.UpdatedTargetEntityResult"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.resources.UpdatedTargetEntityResult"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
LCM upgrade notifications
*/
type UpgradeNotifications struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  List of upgrade notifications and upgrade plan for the input entity
  */
  UpgradePlan []NotificationsSpec `json:"upgradePlan,omitempty"`
}


func NewUpgradeNotifications() *UpgradeNotifications {
  p := new(UpgradeNotifications)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.resources.UpgradeNotifications"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.resources.UpgradeNotifications"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
List of upgrade notifications and upgrade plan for the input entity
*/
type UpgradePlan struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  List of upgrade plans for the hosts
  */
  HostUpgradePlans []HostUpgradePlan `json:"hostUpgradePlans,omitempty"`
}


func NewUpgradePlan() *UpgradePlan {
  p := new(UpgradePlan)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.resources.UpgradePlan"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.resources.UpgradePlan"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Indicates if a user has changed the configuration
*/
type UserPreference struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Indicates if a user has changed the HTTP/HTTPS setting
  */
  UserChangedHttp *bool `json:"userChangedHttp,omitempty"`
}


func NewUserPreference() *UserPreference {
  p := new(UserPreference)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.resources.UserPreference"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.resources.UserPreference"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}



type OneOfGetImagesApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import1.ErrorResponse `json:"-"`
  oneOfType0 []LcmImage `json:"-"`
}

func NewOneOfGetImagesApiResponseData() *OneOfGetImagesApiResponseData {
  p := new(OneOfGetImagesApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfGetImagesApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfGetImagesApiResponseData is nil"))
  }
  switch v.(type) {
    case import1.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = v.(import1.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case []LcmImage:
      p.oneOfType0 = v.([]LcmImage)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<lcm.v4.resources.LcmImage>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<lcm.v4.resources.LcmImage>"
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfGetImagesApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if "List<lcm.v4.resources.LcmImage>" == *p.Discriminator {
    return p.oneOfType0
  }
  return nil
}

func (p *OneOfGetImagesApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import1.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "lcm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new([]LcmImage)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "lcm.v4.resources.LcmImage" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<lcm.v4.resources.LcmImage>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<lcm.v4.resources.LcmImage>"
      return nil

    }
  }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetImagesApiResponseData"))
}

func (p *OneOfGetImagesApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if "List<lcm.v4.resources.LcmImage>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfGetImagesApiResponseData")
}

type OneOfGetNodePrioritiesApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import1.ErrorResponse `json:"-"`
  oneOfType0 *NodePriorityConfig `json:"-"`
}

func NewOneOfGetNodePrioritiesApiResponseData() *OneOfGetNodePrioritiesApiResponseData {
  p := new(OneOfGetNodePrioritiesApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfGetNodePrioritiesApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfGetNodePrioritiesApiResponseData is nil"))
  }
  switch v.(type) {
    case import1.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = v.(import1.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case NodePriorityConfig:
      if nil == p.oneOfType0 {p.oneOfType0 = new(NodePriorityConfig)}
      *p.oneOfType0 = v.(NodePriorityConfig)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfGetNodePrioritiesApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  return nil
}

func (p *OneOfGetNodePrioritiesApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import1.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "lcm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new(NodePriorityConfig)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "lcm.v4.resources.NodePriorityConfig" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(NodePriorityConfig)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetNodePrioritiesApiResponseData"))
}

func (p *OneOfGetNodePrioritiesApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfGetNodePrioritiesApiResponseData")
}

type OneOfGetLcmConfigApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import1.ErrorResponse `json:"-"`
  oneOfType0 *LcmConfig `json:"-"`
}

func NewOneOfGetLcmConfigApiResponseData() *OneOfGetLcmConfigApiResponseData {
  p := new(OneOfGetLcmConfigApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfGetLcmConfigApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfGetLcmConfigApiResponseData is nil"))
  }
  switch v.(type) {
    case import1.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = v.(import1.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case LcmConfig:
      if nil == p.oneOfType0 {p.oneOfType0 = new(LcmConfig)}
      *p.oneOfType0 = v.(LcmConfig)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfGetLcmConfigApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  return nil
}

func (p *OneOfGetLcmConfigApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import1.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "lcm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new(LcmConfig)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "lcm.v4.resources.LcmConfig" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(LcmConfig)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetLcmConfigApiResponseData"))
}

func (p *OneOfGetLcmConfigApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfGetLcmConfigApiResponseData")
}

type OneOfGetModuleConfigApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import1.ErrorResponse `json:"-"`
  oneOfType0 *ModuleConfig `json:"-"`
}

func NewOneOfGetModuleConfigApiResponseData() *OneOfGetModuleConfigApiResponseData {
  p := new(OneOfGetModuleConfigApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfGetModuleConfigApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfGetModuleConfigApiResponseData is nil"))
  }
  switch v.(type) {
    case import1.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = v.(import1.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case ModuleConfig:
      if nil == p.oneOfType0 {p.oneOfType0 = new(ModuleConfig)}
      *p.oneOfType0 = v.(ModuleConfig)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfGetModuleConfigApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  return nil
}

func (p *OneOfGetModuleConfigApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import1.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "lcm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new(ModuleConfig)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "lcm.v4.resources.ModuleConfig" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(ModuleConfig)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetModuleConfigApiResponseData"))
}

func (p *OneOfGetModuleConfigApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfGetModuleConfigApiResponseData")
}

type OneOfDelBundleApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType1 *interface{} `json:"-"`
  oneOfType400 *import1.ErrorResponse `json:"-"`
}

func NewOneOfDelBundleApiResponseData() *OneOfDelBundleApiResponseData {
  p := new(OneOfDelBundleApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfDelBundleApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfDelBundleApiResponseData is nil"))
  }
  if nil == v {
    if nil == p.oneOfType1 {p.oneOfType1 = new(interface {})}
    *p.oneOfType1 = nil
    if nil == p.Discriminator {p.Discriminator = new(string)}
    *p.Discriminator = "EMPTY"
    if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
    *p.ObjectType_ = "EMPTY"
    return nil
  }
  switch v.(type) {
    case import1.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = v.(import1.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfDelBundleApiResponseData) GetValue() interface{} {
  if "EMPTY" == *p.Discriminator {
    return *p.oneOfType1
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfDelBundleApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType1 := new(interface {})
  if err := json.Unmarshal(b, vOneOfType1); err == nil {
    if nil == *vOneOfType1 {
      if nil == p.oneOfType1 {p.oneOfType1 = new(interface {})}
      *p.oneOfType1 = nil
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "EMPTY"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "EMPTY"
      return nil
    }
  }
  vOneOfType400 := new(import1.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "lcm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDelBundleApiResponseData"))
}

func (p *OneOfDelBundleApiResponseData) MarshalJSON() ([]byte, error) {
  if "EMPTY" == *p.Discriminator {
    return json.Marshal(p.oneOfType1)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfDelBundleApiResponseData")
}

type OneOfPostBundlesApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import1.ErrorResponse `json:"-"`
  oneOfType0 *LcmBundlesStatus `json:"-"`
}

func NewOneOfPostBundlesApiResponseData() *OneOfPostBundlesApiResponseData {
  p := new(OneOfPostBundlesApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfPostBundlesApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfPostBundlesApiResponseData is nil"))
  }
  switch v.(type) {
    case import1.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = v.(import1.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case LcmBundlesStatus:
      if nil == p.oneOfType0 {p.oneOfType0 = new(LcmBundlesStatus)}
      *p.oneOfType0 = v.(LcmBundlesStatus)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfPostBundlesApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  return nil
}

func (p *OneOfPostBundlesApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import1.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "lcm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new(LcmBundlesStatus)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "lcm.v4.resources.LcmBundlesStatus" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(LcmBundlesStatus)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfPostBundlesApiResponseData"))
}

func (p *OneOfPostBundlesApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfPostBundlesApiResponseData")
}

type OneOfGetLogBayTagsApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import1.ErrorResponse `json:"-"`
  oneOfType0 []LogbayTag `json:"-"`
}

func NewOneOfGetLogBayTagsApiResponseData() *OneOfGetLogBayTagsApiResponseData {
  p := new(OneOfGetLogBayTagsApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfGetLogBayTagsApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfGetLogBayTagsApiResponseData is nil"))
  }
  switch v.(type) {
    case import1.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = v.(import1.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case []LogbayTag:
      p.oneOfType0 = v.([]LogbayTag)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<lcm.v4.resources.LogbayTag>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<lcm.v4.resources.LogbayTag>"
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfGetLogBayTagsApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if "List<lcm.v4.resources.LogbayTag>" == *p.Discriminator {
    return p.oneOfType0
  }
  return nil
}

func (p *OneOfGetLogBayTagsApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import1.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "lcm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new([]LogbayTag)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "lcm.v4.resources.LogbayTag" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<lcm.v4.resources.LogbayTag>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<lcm.v4.resources.LogbayTag>"
      return nil

    }
  }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetLogBayTagsApiResponseData"))
}

func (p *OneOfGetLogBayTagsApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if "List<lcm.v4.resources.LogbayTag>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfGetLogBayTagsApiResponseData")
}

type OneOfGetUpdateRecApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import1.ErrorResponse `json:"-"`
  oneOfType0 *RecommendationResult `json:"-"`
}

func NewOneOfGetUpdateRecApiResponseData() *OneOfGetUpdateRecApiResponseData {
  p := new(OneOfGetUpdateRecApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfGetUpdateRecApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfGetUpdateRecApiResponseData is nil"))
  }
  switch v.(type) {
    case import1.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = v.(import1.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case RecommendationResult:
      if nil == p.oneOfType0 {p.oneOfType0 = new(RecommendationResult)}
      *p.oneOfType0 = v.(RecommendationResult)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfGetUpdateRecApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  return nil
}

func (p *OneOfGetUpdateRecApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import1.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "lcm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new(RecommendationResult)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "lcm.v4.resources.RecommendationResult" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(RecommendationResult)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetUpdateRecApiResponseData"))
}

func (p *OneOfGetUpdateRecApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfGetUpdateRecApiResponseData")
}

type OneOfGetEntitiesApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import1.ErrorResponse `json:"-"`
  oneOfType0 []LcmEntity `json:"-"`
}

func NewOneOfGetEntitiesApiResponseData() *OneOfGetEntitiesApiResponseData {
  p := new(OneOfGetEntitiesApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfGetEntitiesApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfGetEntitiesApiResponseData is nil"))
  }
  switch v.(type) {
    case import1.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = v.(import1.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case []LcmEntity:
      p.oneOfType0 = v.([]LcmEntity)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<lcm.v4.resources.LcmEntity>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<lcm.v4.resources.LcmEntity>"
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfGetEntitiesApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if "List<lcm.v4.resources.LcmEntity>" == *p.Discriminator {
    return p.oneOfType0
  }
  return nil
}

func (p *OneOfGetEntitiesApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import1.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "lcm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new([]LcmEntity)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "lcm.v4.resources.LcmEntity" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<lcm.v4.resources.LcmEntity>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<lcm.v4.resources.LcmEntity>"
      return nil

    }
  }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetEntitiesApiResponseData"))
}

func (p *OneOfGetEntitiesApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if "List<lcm.v4.resources.LcmEntity>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfGetEntitiesApiResponseData")
}

type OneOfPutLcmConfigApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import1.ErrorResponse `json:"-"`
  oneOfType0 *LcmConfig `json:"-"`
}

func NewOneOfPutLcmConfigApiResponseData() *OneOfPutLcmConfigApiResponseData {
  p := new(OneOfPutLcmConfigApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfPutLcmConfigApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfPutLcmConfigApiResponseData is nil"))
  }
  switch v.(type) {
    case import1.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = v.(import1.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case LcmConfig:
      if nil == p.oneOfType0 {p.oneOfType0 = new(LcmConfig)}
      *p.oneOfType0 = v.(LcmConfig)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfPutLcmConfigApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  return nil
}

func (p *OneOfPutLcmConfigApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import1.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "lcm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new(LcmConfig)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "lcm.v4.resources.LcmConfig" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(LcmConfig)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfPutLcmConfigApiResponseData"))
}

func (p *OneOfPutLcmConfigApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfPutLcmConfigApiResponseData")
}

type OneOfGetBundlesApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import1.ErrorResponse `json:"-"`
  oneOfType0 []LcmBundle `json:"-"`
}

func NewOneOfGetBundlesApiResponseData() *OneOfGetBundlesApiResponseData {
  p := new(OneOfGetBundlesApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfGetBundlesApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfGetBundlesApiResponseData is nil"))
  }
  switch v.(type) {
    case import1.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = v.(import1.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case []LcmBundle:
      p.oneOfType0 = v.([]LcmBundle)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<lcm.v4.resources.LcmBundle>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<lcm.v4.resources.LcmBundle>"
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfGetBundlesApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if "List<lcm.v4.resources.LcmBundle>" == *p.Discriminator {
    return p.oneOfType0
  }
  return nil
}

func (p *OneOfGetBundlesApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import1.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "lcm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new([]LcmBundle)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "lcm.v4.resources.LcmBundle" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<lcm.v4.resources.LcmBundle>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<lcm.v4.resources.LcmBundle>"
      return nil

    }
  }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetBundlesApiResponseData"))
}

func (p *OneOfGetBundlesApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if "List<lcm.v4.resources.LcmBundle>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfGetBundlesApiResponseData")
}

type OneOfGetLogBayItemsApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import1.ErrorResponse `json:"-"`
  oneOfType0 []LogbayItem `json:"-"`
}

func NewOneOfGetLogBayItemsApiResponseData() *OneOfGetLogBayItemsApiResponseData {
  p := new(OneOfGetLogBayItemsApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfGetLogBayItemsApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfGetLogBayItemsApiResponseData is nil"))
  }
  switch v.(type) {
    case import1.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = v.(import1.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case []LogbayItem:
      p.oneOfType0 = v.([]LogbayItem)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<lcm.v4.resources.LogbayItem>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<lcm.v4.resources.LogbayItem>"
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfGetLogBayItemsApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if "List<lcm.v4.resources.LogbayItem>" == *p.Discriminator {
    return p.oneOfType0
  }
  return nil
}

func (p *OneOfGetLogBayItemsApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import1.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "lcm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new([]LogbayItem)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "lcm.v4.resources.LogbayItem" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<lcm.v4.resources.LogbayItem>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<lcm.v4.resources.LogbayItem>"
      return nil

    }
  }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetLogBayItemsApiResponseData"))
}

func (p *OneOfGetLogBayItemsApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if "List<lcm.v4.resources.LogbayItem>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfGetLogBayItemsApiResponseData")
}

type OneOfGetEntityApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import1.ErrorResponse `json:"-"`
  oneOfType0 *LcmEntity `json:"-"`
}

func NewOneOfGetEntityApiResponseData() *OneOfGetEntityApiResponseData {
  p := new(OneOfGetEntityApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfGetEntityApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfGetEntityApiResponseData is nil"))
  }
  switch v.(type) {
    case import1.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = v.(import1.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case LcmEntity:
      if nil == p.oneOfType0 {p.oneOfType0 = new(LcmEntity)}
      *p.oneOfType0 = v.(LcmEntity)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfGetEntityApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  return nil
}

func (p *OneOfGetEntityApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import1.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "lcm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new(LcmEntity)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "lcm.v4.resources.LcmEntity" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(LcmEntity)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetEntityApiResponseData"))
}

func (p *OneOfGetEntityApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfGetEntityApiResponseData")
}

type OneOfGetNodePriorityNodeApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import1.ErrorResponse `json:"-"`
  oneOfType0 *NodePriorityConfig `json:"-"`
}

func NewOneOfGetNodePriorityNodeApiResponseData() *OneOfGetNodePriorityNodeApiResponseData {
  p := new(OneOfGetNodePriorityNodeApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfGetNodePriorityNodeApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfGetNodePriorityNodeApiResponseData is nil"))
  }
  switch v.(type) {
    case import1.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = v.(import1.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case NodePriorityConfig:
      if nil == p.oneOfType0 {p.oneOfType0 = new(NodePriorityConfig)}
      *p.oneOfType0 = v.(NodePriorityConfig)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfGetNodePriorityNodeApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  return nil
}

func (p *OneOfGetNodePriorityNodeApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import1.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "lcm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new(NodePriorityConfig)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "lcm.v4.resources.NodePriorityConfig" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(NodePriorityConfig)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetNodePriorityNodeApiResponseData"))
}

func (p *OneOfGetNodePriorityNodeApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfGetNodePriorityNodeApiResponseData")
}

type OneOfPutBundleApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import1.ErrorResponse `json:"-"`
  oneOfType0 *import4.TaskReference `json:"-"`
}

func NewOneOfPutBundleApiResponseData() *OneOfPutBundleApiResponseData {
  p := new(OneOfPutBundleApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfPutBundleApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfPutBundleApiResponseData is nil"))
  }
  switch v.(type) {
    case import1.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = v.(import1.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case import4.TaskReference:
      if nil == p.oneOfType0 {p.oneOfType0 = new(import4.TaskReference)}
      *p.oneOfType0 = v.(import4.TaskReference)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfPutBundleApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  return nil
}

func (p *OneOfPutBundleApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import1.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "lcm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new(import4.TaskReference)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(import4.TaskReference)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfPutBundleApiResponseData"))
}

func (p *OneOfPutBundleApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfPutBundleApiResponseData")
}

type OneOfGenUpgradePlanApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *UpgradePlan `json:"-"`
  oneOfType400 *import1.ErrorResponse `json:"-"`
}

func NewOneOfGenUpgradePlanApiResponseData() *OneOfGenUpgradePlanApiResponseData {
  p := new(OneOfGenUpgradePlanApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfGenUpgradePlanApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfGenUpgradePlanApiResponseData is nil"))
  }
  switch v.(type) {
    case UpgradePlan:
      if nil == p.oneOfType0 {p.oneOfType0 = new(UpgradePlan)}
      *p.oneOfType0 = v.(UpgradePlan)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    case import1.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = v.(import1.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfGenUpgradePlanApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfGenUpgradePlanApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new(UpgradePlan)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "lcm.v4.resources.UpgradePlan" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(UpgradePlan)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  vOneOfType400 := new(import1.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "lcm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGenUpgradePlanApiResponseData"))
}

func (p *OneOfGenUpgradePlanApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfGenUpgradePlanApiResponseData")
}

type OneOfGetBundleInfoApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import1.ErrorResponse `json:"-"`
  oneOfType0 *LcmBundle `json:"-"`
}

func NewOneOfGetBundleInfoApiResponseData() *OneOfGetBundleInfoApiResponseData {
  p := new(OneOfGetBundleInfoApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfGetBundleInfoApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfGetBundleInfoApiResponseData is nil"))
  }
  switch v.(type) {
    case import1.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = v.(import1.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case LcmBundle:
      if nil == p.oneOfType0 {p.oneOfType0 = new(LcmBundle)}
      *p.oneOfType0 = v.(LcmBundle)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfGetBundleInfoApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  return nil
}

func (p *OneOfGetBundleInfoApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import1.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "lcm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new(LcmBundle)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "lcm.v4.resources.LcmBundle" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(LcmBundle)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetBundleInfoApiResponseData"))
}

func (p *OneOfGetBundleInfoApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfGetBundleInfoApiResponseData")
}

type OneOfGetTaskByUuidApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import1.ErrorResponse `json:"-"`
  oneOfType0 *LcmTask `json:"-"`
}

func NewOneOfGetTaskByUuidApiResponseData() *OneOfGetTaskByUuidApiResponseData {
  p := new(OneOfGetTaskByUuidApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfGetTaskByUuidApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfGetTaskByUuidApiResponseData is nil"))
  }
  switch v.(type) {
    case import1.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = v.(import1.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case LcmTask:
      if nil == p.oneOfType0 {p.oneOfType0 = new(LcmTask)}
      *p.oneOfType0 = v.(LcmTask)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfGetTaskByUuidApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  return nil
}

func (p *OneOfGetTaskByUuidApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import1.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "lcm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new(LcmTask)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "lcm.v4.resources.LcmTask" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(LcmTask)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetTaskByUuidApiResponseData"))
}

func (p *OneOfGetTaskByUuidApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfGetTaskByUuidApiResponseData")
}

type OneOfGenUpgradeNotifApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import1.ErrorResponse `json:"-"`
  oneOfType0 *UpgradeNotifications `json:"-"`
}

func NewOneOfGenUpgradeNotifApiResponseData() *OneOfGenUpgradeNotifApiResponseData {
  p := new(OneOfGenUpgradeNotifApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfGenUpgradeNotifApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfGenUpgradeNotifApiResponseData is nil"))
  }
  switch v.(type) {
    case import1.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = v.(import1.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case UpgradeNotifications:
      if nil == p.oneOfType0 {p.oneOfType0 = new(UpgradeNotifications)}
      *p.oneOfType0 = v.(UpgradeNotifications)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfGenUpgradeNotifApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  return nil
}

func (p *OneOfGenUpgradeNotifApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import1.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "lcm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new(UpgradeNotifications)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "lcm.v4.resources.UpgradeNotifications" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(UpgradeNotifications)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGenUpgradeNotifApiResponseData"))
}

func (p *OneOfGenUpgradeNotifApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfGenUpgradeNotifApiResponseData")
}

type OneOfGetHistoryApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import1.ErrorResponse `json:"-"`
  oneOfType0 []OperationHistoryItem `json:"-"`
}

func NewOneOfGetHistoryApiResponseData() *OneOfGetHistoryApiResponseData {
  p := new(OneOfGetHistoryApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfGetHistoryApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfGetHistoryApiResponseData is nil"))
  }
  switch v.(type) {
    case import1.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = v.(import1.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case []OperationHistoryItem:
      p.oneOfType0 = v.([]OperationHistoryItem)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<lcm.v4.resources.OperationHistoryItem>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<lcm.v4.resources.OperationHistoryItem>"
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfGetHistoryApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if "List<lcm.v4.resources.OperationHistoryItem>" == *p.Discriminator {
    return p.oneOfType0
  }
  return nil
}

func (p *OneOfGetHistoryApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import1.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "lcm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new([]OperationHistoryItem)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "lcm.v4.resources.OperationHistoryItem" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<lcm.v4.resources.OperationHistoryItem>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<lcm.v4.resources.OperationHistoryItem>"
      return nil

    }
  }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetHistoryApiResponseData"))
}

func (p *OneOfGetHistoryApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if "List<lcm.v4.resources.OperationHistoryItem>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfGetHistoryApiResponseData")
}

type OneOfGetHistoryByUuidApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 []OperationHistoryElement `json:"-"`
  oneOfType400 *import1.ErrorResponse `json:"-"`
}

func NewOneOfGetHistoryByUuidApiResponseData() *OneOfGetHistoryByUuidApiResponseData {
  p := new(OneOfGetHistoryByUuidApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfGetHistoryByUuidApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfGetHistoryByUuidApiResponseData is nil"))
  }
  switch v.(type) {
    case []OperationHistoryElement:
      p.oneOfType0 = v.([]OperationHistoryElement)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<lcm.v4.resources.OperationHistoryElement>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<lcm.v4.resources.OperationHistoryElement>"
    case import1.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = v.(import1.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfGetHistoryByUuidApiResponseData) GetValue() interface{} {
  if "List<lcm.v4.resources.OperationHistoryElement>" == *p.Discriminator {
    return p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfGetHistoryByUuidApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new([]OperationHistoryElement)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "lcm.v4.resources.OperationHistoryElement" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<lcm.v4.resources.OperationHistoryElement>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<lcm.v4.resources.OperationHistoryElement>"
      return nil

    }
  }
  vOneOfType400 := new(import1.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "lcm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetHistoryByUuidApiResponseData"))
}

func (p *OneOfGetHistoryByUuidApiResponseData) MarshalJSON() ([]byte, error) {
  if "List<lcm.v4.resources.OperationHistoryElement>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfGetHistoryByUuidApiResponseData")
}

type OneOfPostNodePrioritiesApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import1.ErrorResponse `json:"-"`
  oneOfType0 *NodePriorityConfig `json:"-"`
}

func NewOneOfPostNodePrioritiesApiResponseData() *OneOfPostNodePrioritiesApiResponseData {
  p := new(OneOfPostNodePrioritiesApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfPostNodePrioritiesApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfPostNodePrioritiesApiResponseData is nil"))
  }
  switch v.(type) {
    case import1.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = v.(import1.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case NodePriorityConfig:
      if nil == p.oneOfType0 {p.oneOfType0 = new(NodePriorityConfig)}
      *p.oneOfType0 = v.(NodePriorityConfig)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfPostNodePrioritiesApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  return nil
}

func (p *OneOfPostNodePrioritiesApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import1.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "lcm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new(NodePriorityConfig)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "lcm.v4.resources.NodePriorityConfig" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(NodePriorityConfig)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfPostNodePrioritiesApiResponseData"))
}

func (p *OneOfPostNodePrioritiesApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfPostNodePrioritiesApiResponseData")
}

type OneOfPutModuleConfigApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import1.ErrorResponse `json:"-"`
  oneOfType0 *ModuleConfig `json:"-"`
}

func NewOneOfPutModuleConfigApiResponseData() *OneOfPutModuleConfigApiResponseData {
  p := new(OneOfPutModuleConfigApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfPutModuleConfigApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfPutModuleConfigApiResponseData is nil"))
  }
  switch v.(type) {
    case import1.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = v.(import1.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case ModuleConfig:
      if nil == p.oneOfType0 {p.oneOfType0 = new(ModuleConfig)}
      *p.oneOfType0 = v.(ModuleConfig)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfPutModuleConfigApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  return nil
}

func (p *OneOfPutModuleConfigApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import1.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "lcm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new(ModuleConfig)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "lcm.v4.resources.ModuleConfig" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(ModuleConfig)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfPutModuleConfigApiResponseData"))
}

func (p *OneOfPutModuleConfigApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfPutModuleConfigApiResponseData")
}

type OneOfGetLcmStatusApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import1.ErrorResponse `json:"-"`
  oneOfType0 *LcmStatus `json:"-"`
}

func NewOneOfGetLcmStatusApiResponseData() *OneOfGetLcmStatusApiResponseData {
  p := new(OneOfGetLcmStatusApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfGetLcmStatusApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfGetLcmStatusApiResponseData is nil"))
  }
  switch v.(type) {
    case import1.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = v.(import1.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case LcmStatus:
      if nil == p.oneOfType0 {p.oneOfType0 = new(LcmStatus)}
      *p.oneOfType0 = v.(LcmStatus)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfGetLcmStatusApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  return nil
}

func (p *OneOfGetLcmStatusApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import1.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "lcm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new(LcmStatus)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "lcm.v4.resources.LcmStatus" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(LcmStatus)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetLcmStatusApiResponseData"))
}

func (p *OneOfGetLcmStatusApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfGetLcmStatusApiResponseData")
}

