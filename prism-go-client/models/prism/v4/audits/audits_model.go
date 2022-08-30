/*
 * Generated file models/prism/v4/audits/audits_model.go.
 *
 * Product version: 4.0.1-alpha-1
 *
 * Part of the Nutanix Prism Versioned APIs
 *
 * (c) 2022 Nutanix Inc.  All rights reserved
 *
 */

/*
  Audit alerts
*/
package audits
import (
  import2 "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/common/v1/response"
  "encoding/json"
  "errors"
  "fmt"
  import1 "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/prism/v4/common"
  import3 "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/prism/v4/error"
)


type Audit struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  List of all entities that are affected by the event.
  */
  AffectedEntities []import1.EntityReference `json:"affectedEntities,omitempty"`
  /**
  The unique name for a given audit type. For example, VMCloneAudit, VMDeleteAudit and so on.
  */
  AuditType *string `json:"auditType,omitempty"`
  
  ClusterReference *import1.EntityReference `json:"clusterReference,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  The IP address from where the operation was triggered.
  */
  IpAddress *string `json:"ipAddress,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import2.ApiLink `json:"links,omitempty"`
  /**
  Additional message associated with the event.
  */
  Message *string `json:"message,omitempty"`
  /**
  The audit operation end timestamp in microseconds.
  */
  OperationEndTimestamp *int64 `json:"operationEndTimestamp,omitempty"`
  /**
  The audit operation start timestamp in microseconds.
  */
  OperationStartTimestamp *int64 `json:"operationStartTimestamp,omitempty"`
  
  OperationType *import1.OperationType `json:"operationType,omitempty"`
  /**
  Additional parameters associated with the event. These parameters can be used to indicate custom key-value pairs for a given event instance. For example, Service down in Prism Central event can have the service name as a parameter.
  */
  Parameters []import1.Parameter `json:"parameters,omitempty"`
  /**
  The service which raised the event. For internal Nutanix services, this value is set to "Nutanix".
  */
  ServiceName *string `json:"serviceName,omitempty"`
  
  SourceEntity *import1.EntityReference `json:"sourceEntity,omitempty"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
  /**
  The name of the user who initiated the operation.
  */
  UserName *string `json:"userName,omitempty"`
  /**
  The UUID of the user who initiated the operation.
  */
  UserUuid *string `json:"userUuid,omitempty"`
}


func NewAudit() *Audit {
  p := new(Audit)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "prism.v4.audits.Audit"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a1.audits.Audit"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST response for all response codes in api path /prism/v4.0.a1/audits/{extId} Get operation
*/
type AuditApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfAuditApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewAuditApiResponse() *AuditApiResponse {
  p := new(AuditApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "prism.v4.audits.AuditApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a1.audits.AuditApiResponse"}
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



/**
REST response for all response codes in api path /prism/v4.0.a1/audits Get operation
*/
type AuditListApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfAuditListApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewAuditListApiResponse() *AuditListApiResponse {
  p := new(AuditListApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "prism.v4.audits.AuditListApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a1.audits.AuditListApiResponse"}
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


type OneOfAuditListApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import3.ErrorResponse `json:"-"`
  oneOfType0 []Audit `json:"-"`
}

func NewOneOfAuditListApiResponseData() *OneOfAuditListApiResponseData {
  p := new(OneOfAuditListApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfAuditListApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfAuditListApiResponseData is nil"))
  }
  switch v.(type) {
    case import3.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = v.(import3.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case []Audit:
      p.oneOfType0 = v.([]Audit)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<prism.v4.audits.Audit>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<prism.v4.audits.Audit>"
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfAuditListApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if "List<prism.v4.audits.Audit>" == *p.Discriminator {
    return p.oneOfType0
  }
  return nil
}

func (p *OneOfAuditListApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import3.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "prism.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new([]Audit)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "prism.v4.audits.Audit" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<prism.v4.audits.Audit>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<prism.v4.audits.Audit>"
      return nil

    }
  }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfAuditListApiResponseData"))
}

func (p *OneOfAuditListApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if "List<prism.v4.audits.Audit>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfAuditListApiResponseData")
}

type OneOfAuditApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *Audit `json:"-"`
  oneOfType400 *import3.ErrorResponse `json:"-"`
}

func NewOneOfAuditApiResponseData() *OneOfAuditApiResponseData {
  p := new(OneOfAuditApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfAuditApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfAuditApiResponseData is nil"))
  }
  switch v.(type) {
    case Audit:
      if nil == p.oneOfType0 {p.oneOfType0 = new(Audit)}
      *p.oneOfType0 = v.(Audit)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    case import3.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = v.(import3.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
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
    if "prism.v4.audits.Audit" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(Audit)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  vOneOfType400 := new(import3.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "prism.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
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

