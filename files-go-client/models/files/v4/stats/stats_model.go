/*
 * Generated file models/files/v4/stats/stats_model.go.
 *
 * Product version: 4.0.1-alpha-2
 *
 * Part of the Nutanix Files Versioned APIs
 *
 * (c) 2022 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module for stats APIs for file server
*/
package stats
import (
  import2 "github.com/nutanix/ntnx-api-golang-clients/files-go-client/v4/models/common/v1/response"
  "encoding/json"
  "errors"
  import1 "github.com/nutanix/ntnx-api-golang-clients/files-go-client/v4/models/files/v4/error"
  "fmt"
)

/**
REST response for all response codes in api path /files/v4.0.a2/stats/file-server/anti-virus-servers/{antivirusServerExtId} Get operation
*/
type AntivirusServerStatsApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfAntivirusServerStatsApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewAntivirusServerStatsApiResponse() *AntivirusServerStatsApiResponse {
  p := new(AntivirusServerStatsApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.stats.AntivirusServerStatsApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.stats.AntivirusServerStatsApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *AntivirusServerStatsApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *AntivirusServerStatsApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfAntivirusServerStatsApiResponseData()
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
REST response for all response codes in api path /files/v4.0.a2/stats/file-server Get operation
*/
type FileServerStatsApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfFileServerStatsApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewFileServerStatsApiResponse() *FileServerStatsApiResponse {
  p := new(FileServerStatsApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.stats.FileServerStatsApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.stats.FileServerStatsApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *FileServerStatsApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *FileServerStatsApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfFileServerStatsApiResponseData()
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
REST response for all response codes in api path /files/v4.0.a2/stats/file-server/mount-targets/{mountTargetExtId} Get operation
*/
type MountTargetStatsApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfMountTargetStatsApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewMountTargetStatsApiResponse() *MountTargetStatsApiResponse {
  p := new(MountTargetStatsApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.stats.MountTargetStatsApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.stats.MountTargetStatsApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *MountTargetStatsApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *MountTargetStatsApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfMountTargetStatsApiResponseData()
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
Stats model
*/
type Stats struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Count of available data points. This is a read-only field.
  */
  Count *int `json:"count,omitempty"`
  /**
  Sampling interval of stats
  */
  IntervalInSecs *int `json:"intervalInSecs,omitempty"`
  /**
  Message
  */
  Message *string `json:"message,omitempty"`
  /**
  Metric to received. For example - scanned_file_count, threat_count, cleaned_file_count,quarantined_file_count, latency_ms, throughput_bytes, disconnect_count.
  */
  Metric *string `json:"metric,omitempty"`
  /**
  Start time in microseconds to retrieve all the stats generated after this timestamp. For example: 1622705280584000
  */
  StartTimeInUsecs *int64 `json:"startTimeInUsecs,omitempty"`
  /**
  Success status
  */
  Successful *bool `json:"successful,omitempty"`
  /**
  List of timestamp based data points objects.
  */
  Values []ValueMap `json:"values,omitempty"`
}


func NewStats() *Stats {
  p := new(Stats)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.stats.Stats"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.stats.Stats"}
  p.UnknownFields_ = map[string]interface{}{}

  p.Successful = new(bool)
  *p.Successful = false


  return p
}





type ValueMap struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  Time *int64 `json:"time,omitempty"`
  
  Value *int64 `json:"value,omitempty"`
}


func NewValueMap() *ValueMap {
  p := new(ValueMap)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.stats.ValueMap"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.stats.ValueMap"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}



type OneOfMountTargetStatsApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 []Stats `json:"-"`
  oneOfType400 *import1.ErrorResponse `json:"-"`
}

func NewOneOfMountTargetStatsApiResponseData() *OneOfMountTargetStatsApiResponseData {
  p := new(OneOfMountTargetStatsApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfMountTargetStatsApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfMountTargetStatsApiResponseData is nil"))
  }
  switch v.(type) {
    case []Stats:
      p.oneOfType0 = v.([]Stats)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<files.v4.stats.Stats>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<files.v4.stats.Stats>"
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

func (p *OneOfMountTargetStatsApiResponseData) GetValue() interface{} {
  if "List<files.v4.stats.Stats>" == *p.Discriminator {
    return p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfMountTargetStatsApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new([]Stats)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "files.v4.stats.Stats" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<files.v4.stats.Stats>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<files.v4.stats.Stats>"
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
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfMountTargetStatsApiResponseData"))
}

func (p *OneOfMountTargetStatsApiResponseData) MarshalJSON() ([]byte, error) {
  if "List<files.v4.stats.Stats>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfMountTargetStatsApiResponseData")
}

type OneOfAntivirusServerStatsApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 []Stats `json:"-"`
  oneOfType400 *import1.ErrorResponse `json:"-"`
}

func NewOneOfAntivirusServerStatsApiResponseData() *OneOfAntivirusServerStatsApiResponseData {
  p := new(OneOfAntivirusServerStatsApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfAntivirusServerStatsApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfAntivirusServerStatsApiResponseData is nil"))
  }
  switch v.(type) {
    case []Stats:
      p.oneOfType0 = v.([]Stats)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<files.v4.stats.Stats>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<files.v4.stats.Stats>"
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

func (p *OneOfAntivirusServerStatsApiResponseData) GetValue() interface{} {
  if "List<files.v4.stats.Stats>" == *p.Discriminator {
    return p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfAntivirusServerStatsApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new([]Stats)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "files.v4.stats.Stats" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<files.v4.stats.Stats>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<files.v4.stats.Stats>"
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
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfAntivirusServerStatsApiResponseData"))
}

func (p *OneOfAntivirusServerStatsApiResponseData) MarshalJSON() ([]byte, error) {
  if "List<files.v4.stats.Stats>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfAntivirusServerStatsApiResponseData")
}

type OneOfFileServerStatsApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 []Stats `json:"-"`
  oneOfType400 *import1.ErrorResponse `json:"-"`
}

func NewOneOfFileServerStatsApiResponseData() *OneOfFileServerStatsApiResponseData {
  p := new(OneOfFileServerStatsApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfFileServerStatsApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfFileServerStatsApiResponseData is nil"))
  }
  switch v.(type) {
    case []Stats:
      p.oneOfType0 = v.([]Stats)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<files.v4.stats.Stats>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<files.v4.stats.Stats>"
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

func (p *OneOfFileServerStatsApiResponseData) GetValue() interface{} {
  if "List<files.v4.stats.Stats>" == *p.Discriminator {
    return p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfFileServerStatsApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new([]Stats)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "files.v4.stats.Stats" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<files.v4.stats.Stats>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<files.v4.stats.Stats>"
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
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfFileServerStatsApiResponseData"))
}

func (p *OneOfFileServerStatsApiResponseData) MarshalJSON() ([]byte, error) {
  if "List<files.v4.stats.Stats>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfFileServerStatsApiResponseData")
}

