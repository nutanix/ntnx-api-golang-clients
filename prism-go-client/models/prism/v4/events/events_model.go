/*
 * Generated file models/prism/v4/events/events_model.go.
 *
 * Product version: 4.0.1-alpha-1
 *
 * Part of the Nutanix Prism Versioned APIs
 *
 * (c) 2022 Nutanix Inc.  All rights reserved
 *
 */

/*
  Monitor events
*/
package events
import (
  import2 "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/common/v1/response"
  "encoding/json"
  "errors"
  "fmt"
  import1 "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/prism/v4/common"
  import3 "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/prism/v4/error"
)


type Event struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  List of all entities that are affected by the event.
  */
  AffectedEntities []import1.EntityReference `json:"affectedEntities,omitempty"`
  /**
  Classification of the system defined alert policy.
  */
  Classifications []string `json:"classifications,omitempty"`
  /**
  Cluster UUID associated with the source entity of the event.
  */
  ClusterUUID *string `json:"clusterUUID,omitempty"`
  /**
  Timestamp in microseconds when the event was created.
  */
  CreationTimestampInUsecs *string `json:"creationTimestampInUsecs,omitempty"`
  /**
  A preconfigured or dynamically generated unique value for each event type.
  */
  EventType *string `json:"eventType,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import2.ApiLink `json:"links,omitempty"`
  /**
  Additional message associated with the event.
  */
  Message *string `json:"message,omitempty"`
  /**
  Details of the metric for a metric-based event.
  */
  MetricDetails []import1.MetricDetail `json:"metricDetails,omitempty"`
  /**
  Cluster UUID associated with the cluster where the event was first raised.
  */
  OriginatingClusterUUID *string `json:"originatingClusterUUID,omitempty"`
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
  Title of the event.
  */
  Title *string `json:"title,omitempty"`
}


func NewEvent() *Event {
  p := new(Event)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "prism.v4.events.Event"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a1.events.Event"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST response for all response codes in api path /prism/v4.0.a1/events/{extId} Get operation
*/
type EventApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfEventApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewEventApiResponse() *EventApiResponse {
  p := new(EventApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "prism.v4.events.EventApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a1.events.EventApiResponse"}
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



/**
REST response for all response codes in api path /prism/v4.0.a1/events Get operation
*/
type EventListApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfEventListApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewEventListApiResponse() *EventListApiResponse {
  p := new(EventListApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "prism.v4.events.EventListApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a1.events.EventListApiResponse"}
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


type OneOfEventApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import3.ErrorResponse `json:"-"`
  oneOfType0 *Event `json:"-"`
}

func NewOneOfEventApiResponseData() *OneOfEventApiResponseData {
  p := new(OneOfEventApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfEventApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfEventApiResponseData is nil"))
  }
  switch v.(type) {
    case import3.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = v.(import3.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case Event:
      if nil == p.oneOfType0 {p.oneOfType0 = new(Event)}
      *p.oneOfType0 = v.(Event)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
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
          if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new(Event)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "prism.v4.events.Event" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(Event)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
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

type OneOfEventListApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import3.ErrorResponse `json:"-"`
  oneOfType0 []Event `json:"-"`
}

func NewOneOfEventListApiResponseData() *OneOfEventListApiResponseData {
  p := new(OneOfEventListApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfEventListApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfEventListApiResponseData is nil"))
  }
  switch v.(type) {
    case import3.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = v.(import3.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case []Event:
      p.oneOfType0 = v.([]Event)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<prism.v4.events.Event>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<prism.v4.events.Event>"
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfEventListApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if "List<prism.v4.events.Event>" == *p.Discriminator {
    return p.oneOfType0
  }
  return nil
}

func (p *OneOfEventListApiResponseData) UnmarshalJSON(b []byte) error {
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
  vOneOfType0 := new([]Event)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "prism.v4.events.Event" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<prism.v4.events.Event>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<prism.v4.events.Event>"
      return nil

    }
  }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfEventListApiResponseData"))
}

func (p *OneOfEventListApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if "List<prism.v4.events.Event>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfEventListApiResponseData")
}

