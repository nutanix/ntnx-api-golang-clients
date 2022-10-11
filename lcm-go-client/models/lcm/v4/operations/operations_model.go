/*
 * Generated file models/lcm/v4/operations/operations_model.go.
 *
 * Product version: 4.0.1-alpha-1
 *
 * Part of the Nutanix Lcm Versioned APIs
 *
 * (c) 2022 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module to perform LCM actions 
*/
package operations
import (
  import3 "github.com/nutanix/ntnx-api-golang-clients/lcm-go-client/v4/models/common/v1/response"
  "encoding/json"
  "errors"
  "fmt"
  import2 "github.com/nutanix/ntnx-api-golang-clients/lcm-go-client/v4/models/lcm/v4/error"
  import1 "github.com/nutanix/ntnx-api-golang-clients/lcm-go-client/v4/models/prism/v4/config"
)

/**
REST response for all response codes in api path /lcm/v4.0.a1/operations/$actions/cancelupdate Post operation
*/
type CancelUpdateApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfCancelUpdateApiResponseData `json:"data,omitempty"`
  
  Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewCancelUpdateApiResponse() *CancelUpdateApiResponse {
  p := new(CancelUpdateApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.operations.CancelUpdateApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.operations.CancelUpdateApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *CancelUpdateApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *CancelUpdateApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfCancelUpdateApiResponseData()
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
REST response for all response codes in api path /lcm/v4.0.a1/operations/$actions/downloadImages Post operation
*/
type GetDownloadApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfGetDownloadApiResponseData `json:"data,omitempty"`
  
  Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewGetDownloadApiResponse() *GetDownloadApiResponse {
  p := new(GetDownloadApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.operations.GetDownloadApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.operations.GetDownloadApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *GetDownloadApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *GetDownloadApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfGetDownloadApiResponseData()
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
REST response for all response codes in api path /lcm/v4.0.a1/operations/$actions/performInventory Post operation
*/
type InventoryApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfInventoryApiResponseData `json:"data,omitempty"`
  
  Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewInventoryApiResponse() *InventoryApiResponse {
  p := new(InventoryApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.operations.InventoryApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.operations.InventoryApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *InventoryApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *InventoryApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfInventoryApiResponseData()
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
REST response for all response codes in api path /lcm/v4.0.a1/operations/$actions/performPrecheck Post operation
*/
type PrecheckApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfPrecheckApiResponseData `json:"data,omitempty"`
  
  Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewPrecheckApiResponse() *PrecheckApiResponse {
  p := new(PrecheckApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.operations.PrecheckApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.operations.PrecheckApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *PrecheckApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *PrecheckApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfPrecheckApiResponseData()
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
REST response for all response codes in api path /lcm/v4.0.a1/operations/$actions/performUpdate Post operation
*/
type UpdateApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfUpdateApiResponseData `json:"data,omitempty"`
  
  Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewUpdateApiResponse() *UpdateApiResponse {
  p := new(UpdateApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.operations.UpdateApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.operations.UpdateApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *UpdateApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *UpdateApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfUpdateApiResponseData()
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


type OneOfCancelUpdateApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *import1.TaskReference `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfCancelUpdateApiResponseData() *OneOfCancelUpdateApiResponseData {
  p := new(OneOfCancelUpdateApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfCancelUpdateApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfCancelUpdateApiResponseData is nil"))
  }
  switch v.(type) {
    case import1.TaskReference:
      if nil == p.oneOfType0 {p.oneOfType0 = new(import1.TaskReference)}
      *p.oneOfType0 = v.(import1.TaskReference)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    case import2.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = v.(import2.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfCancelUpdateApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfCancelUpdateApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new(import1.TaskReference)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(import1.TaskReference)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "lcm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCancelUpdateApiResponseData"))
}

func (p *OneOfCancelUpdateApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfCancelUpdateApiResponseData")
}

type OneOfUpdateApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *import1.TaskReference `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfUpdateApiResponseData() *OneOfUpdateApiResponseData {
  p := new(OneOfUpdateApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfUpdateApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfUpdateApiResponseData is nil"))
  }
  switch v.(type) {
    case import1.TaskReference:
      if nil == p.oneOfType0 {p.oneOfType0 = new(import1.TaskReference)}
      *p.oneOfType0 = v.(import1.TaskReference)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    case import2.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = v.(import2.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfUpdateApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfUpdateApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new(import1.TaskReference)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(import1.TaskReference)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "lcm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateApiResponseData"))
}

func (p *OneOfUpdateApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfUpdateApiResponseData")
}

type OneOfInventoryApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *import1.TaskReference `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfInventoryApiResponseData() *OneOfInventoryApiResponseData {
  p := new(OneOfInventoryApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfInventoryApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfInventoryApiResponseData is nil"))
  }
  switch v.(type) {
    case import1.TaskReference:
      if nil == p.oneOfType0 {p.oneOfType0 = new(import1.TaskReference)}
      *p.oneOfType0 = v.(import1.TaskReference)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    case import2.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = v.(import2.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfInventoryApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfInventoryApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new(import1.TaskReference)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(import1.TaskReference)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "lcm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfInventoryApiResponseData"))
}

func (p *OneOfInventoryApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfInventoryApiResponseData")
}

type OneOfPrecheckApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *import1.TaskReference `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfPrecheckApiResponseData() *OneOfPrecheckApiResponseData {
  p := new(OneOfPrecheckApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfPrecheckApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfPrecheckApiResponseData is nil"))
  }
  switch v.(type) {
    case import1.TaskReference:
      if nil == p.oneOfType0 {p.oneOfType0 = new(import1.TaskReference)}
      *p.oneOfType0 = v.(import1.TaskReference)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    case import2.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = v.(import2.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfPrecheckApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfPrecheckApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new(import1.TaskReference)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(import1.TaskReference)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "lcm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfPrecheckApiResponseData"))
}

func (p *OneOfPrecheckApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfPrecheckApiResponseData")
}

type OneOfGetDownloadApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *import1.TaskReference `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfGetDownloadApiResponseData() *OneOfGetDownloadApiResponseData {
  p := new(OneOfGetDownloadApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfGetDownloadApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfGetDownloadApiResponseData is nil"))
  }
  switch v.(type) {
    case import1.TaskReference:
      if nil == p.oneOfType0 {p.oneOfType0 = new(import1.TaskReference)}
      *p.oneOfType0 = v.(import1.TaskReference)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    case import2.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = v.(import2.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfGetDownloadApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfGetDownloadApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new(import1.TaskReference)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(import1.TaskReference)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "lcm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetDownloadApiResponseData"))
}

func (p *OneOfGetDownloadApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfGetDownloadApiResponseData")
}

