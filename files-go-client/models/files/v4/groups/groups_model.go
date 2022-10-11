/*
 * Generated file models/files/v4/groups/groups_model.go.
 *
 * Product version: 4.0.1-alpha-2
 *
 * Part of the Nutanix Files Versioned APIs
 *
 * (c) 2022 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module for proxy APIs for alerts and events
*/
package groups
import (
  import2 "github.com/nutanix/ntnx-api-golang-clients/files-go-client/v4/models/common/v1/response"
  "encoding/json"
  "errors"
  import1 "github.com/nutanix/ntnx-api-golang-clients/files-go-client/v4/models/files/v4/error"
  "fmt"
)

/**
Group query attribute
*/
type Attribute struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  Attribute *string `json:"attribute,omitempty"`
}


func NewAttribute() *Attribute {
  p := new(Attribute)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.groups.Attribute"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.groups.Attribute"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST response for all response codes in api path /files/v4.0.a2/groups/file-server Post operation
*/
type GroupsQueryApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfGroupsQueryApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewGroupsQueryApiResponse() *GroupsQueryApiResponse {
  p := new(GroupsQueryApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.groups.GroupsQueryApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.groups.GroupsQueryApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *GroupsQueryApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *GroupsQueryApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfGroupsQueryApiResponseData()
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
Alerts group request
*/
type GroupsQueryRequest struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Entity type to get the alerts.
  */
  EntityType *string `json:"entityType,omitempty"`
  /**
  Filter criteria for entries in group query response.
  */
  FilterCriteria *string `json:"filterCriteria,omitempty"`
  /**
  Attributes of entries returned in group query response.
  */
  GroupMemberAttributes []Attribute `json:"groupMemberAttributes,omitempty"`
  /**
  Count of the entries requested in response.
  */
  GroupMemberCount *int64 `json:"groupMemberCount,omitempty"`
  /**
  Offset of the entries returned in group query response.
  */
  GroupMemberOffset *int64 `json:"groupMemberOffset,omitempty"`
  /**
  Attribute name for sorting the results in group query response.
  */
  GroupMemberSortAttribute *string `json:"groupMemberSortAttribute,omitempty"`
  /**
  Attribute name for ordering the results in group query response.
  */
  GroupMemberSortOrder *string `json:"groupMemberSortOrder,omitempty"`
}


func NewGroupsQueryRequest() *GroupsQueryRequest {
  p := new(GroupsQueryRequest)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.groups.GroupsQueryRequest"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.groups.GroupsQueryRequest"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Group query response
*/
type GroupsQueryResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Entity type to get the alerts.
  */
  EntityType *string `json:"entityType,omitempty"`
  /**
  Count of the filtered entities returned in group query response.
  */
  FilteredEntityCount *int64 `json:"filteredEntityCount,omitempty"`
  /**
  Count of the filtered groups returned in group query response.
  */
  FilteredGroupCount *int64 `json:"filteredGroupCount,omitempty"`
  /**
  List of objects returned in group query response.
  */
  GroupResults []interface{} `json:"groupResults,omitempty"`
  /**
  Total entity count returned in group query response.
  */
  TotalEntityCount *int64 `json:"totalEntityCount,omitempty"`
  /**
  Total group count returned in group query response.
  */
  TotalGroupCount *int64 `json:"totalGroupCount,omitempty"`
}


func NewGroupsQueryResponse() *GroupsQueryResponse {
  p := new(GroupsQueryResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.groups.GroupsQueryResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.groups.GroupsQueryResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}



type OneOfGroupsQueryApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *GroupsQueryResponse `json:"-"`
  oneOfType400 *import1.ErrorResponse `json:"-"`
}

func NewOneOfGroupsQueryApiResponseData() *OneOfGroupsQueryApiResponseData {
  p := new(OneOfGroupsQueryApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfGroupsQueryApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfGroupsQueryApiResponseData is nil"))
  }
  switch v.(type) {
    case GroupsQueryResponse:
      if nil == p.oneOfType0 {p.oneOfType0 = new(GroupsQueryResponse)}
      *p.oneOfType0 = v.(GroupsQueryResponse)
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

func (p *OneOfGroupsQueryApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfGroupsQueryApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new(GroupsQueryResponse)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "files.v4.groups.GroupsQueryResponse" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(GroupsQueryResponse)}
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
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGroupsQueryApiResponseData"))
}

func (p *OneOfGroupsQueryApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfGroupsQueryApiResponseData")
}

