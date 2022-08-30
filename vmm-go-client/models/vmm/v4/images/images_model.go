/*
 * Generated file models/vmm/v4/images/images_model.go.
 *
 * Product version: 4.0.1-alpha-1
 *
 * Part of the Nutanix Vmm Versioned APIs
 *
 * (c) 2022 Nutanix Inc.  All rights reserved
 *
 */

/*
  Manage Images. Configure Rate limits and Placement Policies
*/
package images
import (
  "bytes"
  import1 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/common/v1/response"
  "encoding/json"
  "errors"
  "fmt"
  import3 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/config"
  import2 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/error"
)

/**
Indicates whether the tasks running to enforce policy should be canceled. These would be checkout and uncheckout tasks to placement remove images from the cluster(s).
*/
type CancelPlacementTasks struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Boolean to indicate if tasks running to enforce image placement policy should be suspended.
  */
  CancelRunningTasks *bool `json:"cancelRunningTasks,omitempty"`
}


func NewCancelPlacementTasks() *CancelPlacementTasks {
  p := new(CancelPlacementTasks)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "vmm.v4.images.CancelPlacementTasks"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.images.CancelPlacementTasks"}
  p.UnknownFields_ = map[string]interface{}{}

  p.CancelRunningTasks = new(bool)
  *p.CancelRunningTasks = false


  return p
}




/**
Reference to a category.
*/
type CategoryReference struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  ExtId *string `json:"extId"`
}

func (p *CategoryReference) MarshalJSON() ([]byte, error) {
  type CategoryReferenceProxy CategoryReference
  return json.Marshal(struct {
    *CategoryReferenceProxy
    ExtId *string `json:"extId,omitempty"`
  }{
    CategoryReferenceProxy : (*CategoryReferenceProxy)(p),
    ExtId : p.ExtId,
  })
}

func NewCategoryReference() *CategoryReference {
  p := new(CategoryReference)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "vmm.v4.images.CategoryReference"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.images.CategoryReference"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Category based Cluster filter.
*/
type ClusterEntityFilter struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  A list of category key and list of values.
  */
  Params map[string][]string `json:"params"`
  
  Type *FilterMatchType `json:"type"`
}

func (p *ClusterEntityFilter) MarshalJSON() ([]byte, error) {
  type ClusterEntityFilterProxy ClusterEntityFilter
  return json.Marshal(struct {
    *ClusterEntityFilterProxy
    Params map[string][]string `json:"params,omitempty"`
    Type *FilterMatchType `json:"type,omitempty"`
  }{
    ClusterEntityFilterProxy : (*ClusterEntityFilterProxy)(p),
    Params : p.Params,
    Type : p.Type,
  })
}

func NewClusterEntityFilter() *ClusterEntityFilter {
  p := new(ClusterEntityFilter)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "vmm.v4.images.ClusterEntityFilter"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.images.ClusterEntityFilter"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
List of cluster extIds.
*/
type ClusterIdList struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  List of cluster extIds.
  */
  ClusterIds []string `json:"clusterIds,omitempty"`
}


func NewClusterIdList() *ClusterIdList {
  p := new(ClusterIdList)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "vmm.v4.images.ClusterIdList"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.images.ClusterIdList"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Reference to a Prism Element cluster and respective retained images.
*/
type ClusterImages struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  ClusterExtId *string `json:"clusterExtId"`
  
  ImagesExtIds []string `json:"imagesExtIds"`
}

func (p *ClusterImages) MarshalJSON() ([]byte, error) {
  type ClusterImagesProxy ClusterImages
  return json.Marshal(struct {
    *ClusterImagesProxy
    ClusterExtId *string `json:"clusterExtId,omitempty"`
    ImagesExtIds []string `json:"imagesExtIds,omitempty"`
  }{
    ClusterImagesProxy : (*ClusterImagesProxy)(p),
    ClusterExtId : p.ClusterExtId,
    ImagesExtIds : p.ImagesExtIds,
  })
}

func NewClusterImages() *ClusterImages {
  p := new(ClusterImages)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "vmm.v4.images.ClusterImages"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.images.ClusterImages"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Cluster image rate limit status.
*/
type ClusterRateLimitStatus struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Cluster extId.
  */
  ClusterId *string `json:"clusterId,omitempty"`
  /**
  Image rate limit extId currently enforced for a cluster.
  */
  EffectiveRateLimitId *string `json:"effectiveRateLimitId,omitempty"`
  /**
  The currently enforced rate limit for a cluster in KBps.
  */
  EffectiveRateLimitKbps *int `json:"effectiveRateLimitKbps,omitempty"`
  /**
  Name of the image rate limit currently enforced for a cluster.
  */
  EffectiveRateLimitName *string `json:"effectiveRateLimitName,omitempty"`
}


func NewClusterRateLimitStatus() *ClusterRateLimitStatus {
  p := new(ClusterRateLimitStatus)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "vmm.v4.images.ClusterRateLimitStatus"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.images.ClusterRateLimitStatus"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Reference to a cluster.
*/
type ClusterReference struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  ExtId *string `json:"extId"`
}

func (p *ClusterReference) MarshalJSON() ([]byte, error) {
  type ClusterReferenceProxy ClusterReference
  return json.Marshal(struct {
    *ClusterReferenceProxy
    ExtId *string `json:"extId,omitempty"`
  }{
    ClusterReferenceProxy : (*ClusterReferenceProxy)(p),
    ExtId : p.ExtId,
  })
}

func NewClusterReference() *ClusterReference {
  p := new(ClusterReference)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "vmm.v4.images.ClusterReference"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.images.ClusterReference"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Compliance status for a placement policy.
*/
type ComplianceStatus int

const(
  COMPLIANCESTATUS_UNKNOWN ComplianceStatus = 0
  COMPLIANCESTATUS_REDACTED ComplianceStatus = 1
  COMPLIANCESTATUS_COMPLIANT ComplianceStatus = 2
  COMPLIANCESTATUS_NON_COMPLIANT ComplianceStatus = 3
)

// returns the name of the enum given an ordinal number
func (e *ComplianceStatus) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "COMPLIANT",
    "NON_COMPLIANT",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *ComplianceStatus) index(name string) ComplianceStatus {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "COMPLIANT",
    "NON_COMPLIANT",
  }
  for idx := range names {
    if names[idx] == name {
      return ComplianceStatus(idx)
    }
  }
  return COMPLIANCESTATUS_UNKNOWN
}

func (e *ComplianceStatus) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for ComplianceStatus:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *ComplianceStatus) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e ComplianceStatus) Ref() *ComplianceStatus {
  return &e
}


/**
The enforced rate limit on cluster when multiple image rate limit policies apply.
*/
type EffectiveRateLimit struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  ClusterId *string `json:"clusterId,omitempty"`
  
  RateLimit *RateLimit `json:"rateLimit,omitempty"`
}


func NewEffectiveRateLimit() *EffectiveRateLimit {
  p := new(EffectiveRateLimit)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "vmm.v4.images.EffectiveRateLimit"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.images.EffectiveRateLimit"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Indicates whether placement policy enforcement is ongoing or has failed.
*/
type EnforcementMode int

const(
  ENFORCEMENTMODE_UNKNOWN EnforcementMode = 0
  ENFORCEMENTMODE_REDACTED EnforcementMode = 1
  ENFORCEMENTMODE_ENFORCING EnforcementMode = 2
  ENFORCEMENTMODE_ENFORCEMENT_FAILED EnforcementMode = 3
)

// returns the name of the enum given an ordinal number
func (e *EnforcementMode) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "ENFORCING",
    "ENFORCEMENT_FAILED",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *EnforcementMode) index(name string) EnforcementMode {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "ENFORCING",
    "ENFORCEMENT_FAILED",
  }
  for idx := range names {
    if names[idx] == name {
      return EnforcementMode(idx)
    }
  }
  return ENFORCEMENTMODE_UNKNOWN
}

func (e *EnforcementMode) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for EnforcementMode:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *EnforcementMode) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e EnforcementMode) Ref() *EnforcementMode {
  return &e
}


/**
Enforcement status of the image placement policy.
*/
type EnforcementState int

const(
  ENFORCEMENTSTATE_UNKNOWN EnforcementState = 0
  ENFORCEMENTSTATE_REDACTED EnforcementState = 1
  ENFORCEMENTSTATE_ACTIVE EnforcementState = 2
  ENFORCEMENTSTATE_SUSPENDED EnforcementState = 3
)

// returns the name of the enum given an ordinal number
func (e *EnforcementState) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "ACTIVE",
    "SUSPENDED",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *EnforcementState) index(name string) EnforcementState {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "ACTIVE",
    "SUSPENDED",
  }
  for idx := range names {
    if names[idx] == name {
      return EnforcementState(idx)
    }
  }
  return ENFORCEMENTSTATE_UNKNOWN
}

func (e *EnforcementState) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for EnforcementState:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *EnforcementState) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e EnforcementState) Ref() *EnforcementState {
  return &e
}


/**
Category-based entity filter.
*/
type EntityFilter struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  A list of category key and list of values.
  */
  Params map[string][]string `json:"params"`
  
  Type *FilterMatchType `json:"type"`
}

func (p *EntityFilter) MarshalJSON() ([]byte, error) {
  type EntityFilterProxy EntityFilter
  return json.Marshal(struct {
    *EntityFilterProxy
    Params map[string][]string `json:"params,omitempty"`
    Type *FilterMatchType `json:"type,omitempty"`
  }{
    EntityFilterProxy : (*EntityFilterProxy)(p),
    Params : p.Params,
    Type : p.Type,
  })
}

func NewEntityFilter() *EntityFilter {
  p := new(EntityFilter)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "vmm.v4.images.EntityFilter"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.images.EntityFilter"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
The matching done by the filter.
*/
type FilterMatchType int

const(
  FILTERMATCHTYPE_UNKNOWN FilterMatchType = 0
  FILTERMATCHTYPE_REDACTED FilterMatchType = 1
  FILTERMATCHTYPE_CATEGORIES_MATCH_ANY FilterMatchType = 2
  FILTERMATCHTYPE_CATEGORIES_MATCH_ALL FilterMatchType = 3
)

// returns the name of the enum given an ordinal number
func (e *FilterMatchType) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "CATEGORIES_MATCH_ANY",
    "CATEGORIES_MATCH_ALL",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *FilterMatchType) index(name string) FilterMatchType {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "CATEGORIES_MATCH_ANY",
    "CATEGORIES_MATCH_ALL",
  }
  for idx := range names {
    if names[idx] == name {
      return FilterMatchType(idx)
    }
  }
  return FILTERMATCHTYPE_UNKNOWN
}

func (e *FilterMatchType) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for FilterMatchType:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *FilterMatchType) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e FilterMatchType) Ref() *FilterMatchType {
  return &e
}



type Image struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Categories for an image.
  */
  Categories []CategoryReference `json:"categories,omitempty"`
  
  Checksum *ImageChecksum `json:"checksum,omitempty"`
  
  Description *string `json:"description,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  Cluster details for the image location.
  */
  InitialClusterLocations []ClusterReference `json:"initialClusterLocations,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import1.ApiLink `json:"links,omitempty"`
  
  Name *string `json:"name"`
  /**
  The size in bytes of an image file.
  */
  SizeBytes *int64 `json:"sizeBytes,omitempty"`
  
  SourceItemDiscriminator_ *string `json:"$sourceItemDiscriminator,omitempty"`
  /**
  The source of an image. It can be a VM disk or a URL. If not mentioned, the image will be inactive until a file is uploaded for it.
  */
  Source *OneOfImageSource `json:"source,omitempty"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
  
  Type *ImageType `json:"type"`
}

func (p *Image) MarshalJSON() ([]byte, error) {
  type ImageProxy Image
  return json.Marshal(struct {
    *ImageProxy
    Name *string `json:"name,omitempty"`
    Type *ImageType `json:"type,omitempty"`
  }{
    ImageProxy : (*ImageProxy)(p),
    Name : p.Name,
    Type : p.Type,
  })
}

func NewImage() *Image {
  p := new(Image)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "vmm.v4.images.Image"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.images.Image"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST response for all response codes in api path /vmm/v4.0.a1/images/{extId} Get operation
*/
type ImageApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfImageApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewImageApiResponse() *ImageApiResponse {
  p := new(ImageApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "vmm.v4.images.ImageApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.images.ImageApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *ImageApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *ImageApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfImageApiResponseData()
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
REST response for all response codes in api path /vmm/v4.0.a1/images/{extId}/categories Get operation
*/
type ImageCategoriesApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfImageCategoriesApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewImageCategoriesApiResponse() *ImageCategoriesApiResponse {
  p := new(ImageCategoriesApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "vmm.v4.images.ImageCategoriesApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.images.ImageCategoriesApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *ImageCategoriesApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *ImageCategoriesApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfImageCategoriesApiResponseData()
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
The checksum of an image.
*/
type ImageChecksum struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  HexDigest *string `json:"hexDigest"`
  
  Type *ImageChecksumType `json:"type"`
}

func (p *ImageChecksum) MarshalJSON() ([]byte, error) {
  type ImageChecksumProxy ImageChecksum
  return json.Marshal(struct {
    *ImageChecksumProxy
    HexDigest *string `json:"hexDigest,omitempty"`
    Type *ImageChecksumType `json:"type,omitempty"`
  }{
    ImageChecksumProxy : (*ImageChecksumProxy)(p),
    HexDigest : p.HexDigest,
    Type : p.Type,
  })
}

func NewImageChecksum() *ImageChecksum {
  p := new(ImageChecksum)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "vmm.v4.images.ImageChecksum"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.images.ImageChecksum"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
The checksum algorithm of an image.
*/
type ImageChecksumType int

const(
  IMAGECHECKSUMTYPE_UNKNOWN ImageChecksumType = 0
  IMAGECHECKSUMTYPE_REDACTED ImageChecksumType = 1
  IMAGECHECKSUMTYPE_SHA_256 ImageChecksumType = 2
  IMAGECHECKSUMTYPE_SHA_1 ImageChecksumType = 3
)

// returns the name of the enum given an ordinal number
func (e *ImageChecksumType) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "SHA_256",
    "SHA_1",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *ImageChecksumType) index(name string) ImageChecksumType {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "SHA_256",
    "SHA_1",
  }
  for idx := range names {
    if names[idx] == name {
      return ImageChecksumType(idx)
    }
  }
  return IMAGECHECKSUMTYPE_UNKNOWN
}

func (e *ImageChecksumType) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for ImageChecksumType:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *ImageChecksumType) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e ImageChecksumType) Ref() *ImageChecksumType {
  return &e
}


/**
REST response for all response codes in api path /vmm/v4.0.a1/images/{extId}/cluster-locations Get operation
*/
type ImageClusterLocationListApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfImageClusterLocationListApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewImageClusterLocationListApiResponse() *ImageClusterLocationListApiResponse {
  p := new(ImageClusterLocationListApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "vmm.v4.images.ImageClusterLocationListApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.images.ImageClusterLocationListApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *ImageClusterLocationListApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *ImageClusterLocationListApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfImageClusterLocationListApiResponseData()
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
REST response for all response codes in api path /vmm/v4.0.a1/images Get operation
*/
type ImageListApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfImageListApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewImageListApiResponse() *ImageListApiResponse {
  p := new(ImageListApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "vmm.v4.images.ImageListApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.images.ImageListApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *ImageListApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *ImageListApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfImageListApiResponseData()
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
REST response for all response codes in api path /vmm/v4.0.a1/images/{extId}/placement-policies Get operation
*/
type ImagePlacementPolicyListApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfImagePlacementPolicyListApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewImagePlacementPolicyListApiResponse() *ImagePlacementPolicyListApiResponse {
  p := new(ImagePlacementPolicyListApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "vmm.v4.images.ImagePlacementPolicyListApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.images.ImagePlacementPolicyListApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *ImagePlacementPolicyListApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *ImagePlacementPolicyListApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfImagePlacementPolicyListApiResponseData()
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
Reference to an image.
*/
type ImageReference struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  ExtId *string `json:"extId"`
}

func (p *ImageReference) MarshalJSON() ([]byte, error) {
  type ImageReferenceProxy ImageReference
  return json.Marshal(struct {
    *ImageReferenceProxy
    ExtId *string `json:"extId,omitempty"`
  }{
    ImageReferenceProxy : (*ImageReferenceProxy)(p),
    ExtId : p.ExtId,
  })
}

func NewImageReference() *ImageReference {
  p := new(ImageReference)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "vmm.v4.images.ImageReference"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.images.ImageReference"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
The type of an image.
*/
type ImageType int

const(
  IMAGETYPE_UNKNOWN ImageType = 0
  IMAGETYPE_REDACTED ImageType = 1
  IMAGETYPE_DISK_IMAGE ImageType = 2
  IMAGETYPE_ISO_IMAGE ImageType = 3
)

// returns the name of the enum given an ordinal number
func (e *ImageType) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "DISK_IMAGE",
    "ISO_IMAGE",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *ImageType) index(name string) ImageType {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "DISK_IMAGE",
    "ISO_IMAGE",
  }
  for idx := range names {
    if names[idx] == name {
      return ImageType(idx)
    }
  }
  return IMAGETYPE_UNKNOWN
}

func (e *ImageType) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for ImageType:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *ImageType) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e ImageType) Ref() *ImageType {
  return &e
}


/**
REST response for all response codes in api path /vmm/v4.0.a1/images/{extId}/categories Put operation
*/
type ImagesTaskApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfImagesTaskApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewImagesTaskApiResponse() *ImagesTaskApiResponse {
  p := new(ImagesTaskApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "vmm.v4.images.ImagesTaskApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.images.ImagesTaskApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *ImagesTaskApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *ImagesTaskApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfImagesTaskApiResponseData()
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




type PlacementPolicy struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  ClusterEntityFilter *EntityFilter `json:"clusterEntityFilter"`
  /**
  Description of the image placement policy.
  */
  Description *string `json:"description,omitempty"`
  
  EnforcementState *EnforcementState `json:"enforcementState,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  
  ImageEntityFilter *EntityFilter `json:"imageEntityFilter"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import1.ApiLink `json:"links,omitempty"`
  /**
  Name of the image placement policy.
  */
  Name *string `json:"name"`
  
  PlacementType *PlacementType `json:"placementType"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
}

func (p *PlacementPolicy) MarshalJSON() ([]byte, error) {
  type PlacementPolicyProxy PlacementPolicy
  return json.Marshal(struct {
    *PlacementPolicyProxy
    ClusterEntityFilter *EntityFilter `json:"clusterEntityFilter,omitempty"`
    ImageEntityFilter *EntityFilter `json:"imageEntityFilter,omitempty"`
    Name *string `json:"name,omitempty"`
    PlacementType *PlacementType `json:"placementType,omitempty"`
  }{
    PlacementPolicyProxy : (*PlacementPolicyProxy)(p),
    ClusterEntityFilter : p.ClusterEntityFilter,
    ImageEntityFilter : p.ImageEntityFilter,
    Name : p.Name,
    PlacementType : p.PlacementType,
  })
}

func NewPlacementPolicy() *PlacementPolicy {
  p := new(PlacementPolicy)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "vmm.v4.images.PlacementPolicy"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.images.PlacementPolicy"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST response for all response codes in api path /vmm/v4.0.a1/images/placement-policies/{extId} Get operation
*/
type PlacementPolicyApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfPlacementPolicyApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewPlacementPolicyApiResponse() *PlacementPolicyApiResponse {
  p := new(PlacementPolicyApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "vmm.v4.images.PlacementPolicyApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.images.PlacementPolicyApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *PlacementPolicyApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *PlacementPolicyApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfPlacementPolicyApiResponseData()
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
REST response for all response codes in api path /vmm/v4.0.a1/images/placement-policies Get operation
*/
type PlacementPolicyListApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfPlacementPolicyListApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewPlacementPolicyListApiResponse() *PlacementPolicyListApiResponse {
  p := new(PlacementPolicyListApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "vmm.v4.images.PlacementPolicyListApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.images.PlacementPolicyListApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *PlacementPolicyListApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *PlacementPolicyListApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfPlacementPolicyListApiResponseData()
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
Reference to a placement policy.
*/
type PlacementPolicyReference struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  ExtId *string `json:"extId"`
}

func (p *PlacementPolicyReference) MarshalJSON() ([]byte, error) {
  type PlacementPolicyReferenceProxy PlacementPolicyReference
  return json.Marshal(struct {
    *PlacementPolicyReferenceProxy
    ExtId *string `json:"extId,omitempty"`
  }{
    PlacementPolicyReferenceProxy : (*PlacementPolicyReferenceProxy)(p),
    ExtId : p.ExtId,
  })
}

func NewPlacementPolicyReference() *PlacementPolicyReference {
  p := new(PlacementPolicyReference)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "vmm.v4.images.PlacementPolicyReference"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.images.PlacementPolicyReference"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Status of an image placement policy.
*/
type PlacementPolicyStatus struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  ComplianceStatus *ComplianceStatus `json:"complianceStatus"`
  /**
  Placement policies that conflict with the current one.
  */
  ConflictingPolicies []PlacementPolicyReference `json:"conflictingPolicies,omitempty"`
  /**
  Clusters for the enforced placement policy.
  */
  EnforcedClusters []ClusterReference `json:"enforcedClusters,omitempty"`
  
  EnforcementMode *EnforcementMode `json:"enforcementMode"`
  
  ExtId *string `json:"extId"`
  /**
  Clusters details of the image location for the enforced placement policy.
  */
  PolicyClusters []ClusterReference `json:"policyClusters,omitempty"`
}

func (p *PlacementPolicyStatus) MarshalJSON() ([]byte, error) {
  type PlacementPolicyStatusProxy PlacementPolicyStatus
  return json.Marshal(struct {
    *PlacementPolicyStatusProxy
    ComplianceStatus *ComplianceStatus `json:"complianceStatus,omitempty"`
    EnforcementMode *EnforcementMode `json:"enforcementMode,omitempty"`
    ExtId *string `json:"extId,omitempty"`
  }{
    PlacementPolicyStatusProxy : (*PlacementPolicyStatusProxy)(p),
    ComplianceStatus : p.ComplianceStatus,
    EnforcementMode : p.EnforcementMode,
    ExtId : p.ExtId,
  })
}

func NewPlacementPolicyStatus() *PlacementPolicyStatus {
  p := new(PlacementPolicyStatus)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "vmm.v4.images.PlacementPolicyStatus"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.images.PlacementPolicyStatus"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Type of the image placement policy.
*/
type PlacementType int

const(
  PLACEMENTTYPE_UNKNOWN PlacementType = 0
  PLACEMENTTYPE_REDACTED PlacementType = 1
  PLACEMENTTYPE_AT_LEAST PlacementType = 2
  PLACEMENTTYPE_EXACTLY PlacementType = 3
)

// returns the name of the enum given an ordinal number
func (e *PlacementType) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "AT_LEAST",
    "EXACTLY",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *PlacementType) index(name string) PlacementType {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "AT_LEAST",
    "EXACTLY",
  }
  for idx := range names {
    if names[idx] == name {
      return PlacementType(idx)
    }
  }
  return PLACEMENTTYPE_UNKNOWN
}

func (e *PlacementType) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for PlacementType:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *PlacementType) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e PlacementType) Ref() *PlacementType {
  return &e
}



type RateLimit struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  ClusterEntityFilter *ClusterEntityFilter `json:"clusterEntityFilter"`
  /**
  Image rate limit status of each affected cluster.
  */
  ClusterStatusList []ClusterRateLimitStatus `json:"clusterStatusList,omitempty"`
  /**
  A description for the current image rate limit.
  */
  Description *string `json:"description,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import1.ApiLink `json:"links,omitempty"`
  /**
  Name of the image rate limit policy.
  */
  Name *string `json:"name"`
  /**
  Network bandwidth in KBps that the rate limited image operation can utilize.
  */
  RateLimitKbps *int `json:"rateLimitKbps"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
  /**
  List of cluster extIds matched by the policy which do not support image rate limits.
  */
  UnsupportedClustersIdList []string `json:"unsupportedClustersIdList,omitempty"`
}

func (p *RateLimit) MarshalJSON() ([]byte, error) {
  type RateLimitProxy RateLimit
  return json.Marshal(struct {
    *RateLimitProxy
    ClusterEntityFilter *ClusterEntityFilter `json:"clusterEntityFilter,omitempty"`
    Name *string `json:"name,omitempty"`
    RateLimitKbps *int `json:"rateLimitKbps,omitempty"`
  }{
    RateLimitProxy : (*RateLimitProxy)(p),
    ClusterEntityFilter : p.ClusterEntityFilter,
    Name : p.Name,
    RateLimitKbps : p.RateLimitKbps,
  })
}

func NewRateLimit() *RateLimit {
  p := new(RateLimit)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "vmm.v4.images.RateLimit"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.images.RateLimit"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST response for all response codes in api path /vmm/v4.0.a1/images/rate-limits/{extId} Get operation
*/
type RateLimitApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfRateLimitApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewRateLimitApiResponse() *RateLimitApiResponse {
  p := new(RateLimitApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "vmm.v4.images.RateLimitApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.images.RateLimitApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *RateLimitApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *RateLimitApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfRateLimitApiResponseData()
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
List of image rate limits extIds.
*/
type RateLimitIdList struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  List of image rate limits extIds.
  */
  RateLimitIds []string `json:"rateLimitIds,omitempty"`
}


func NewRateLimitIdList() *RateLimitIdList {
  p := new(RateLimitIdList)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "vmm.v4.images.RateLimitIdList"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.images.RateLimitIdList"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST response for all response codes in api path /vmm/v4.0.a1/images/rate-limits Get operation
*/
type RateLimitListApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfRateLimitListApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewRateLimitListApiResponse() *RateLimitListApiResponse {
  p := new(RateLimitListApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "vmm.v4.images.RateLimitListApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.images.RateLimitListApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *RateLimitListApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *RateLimitListApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfRateLimitListApiResponseData()
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
Object with a rate limit and task extIds.
*/
type RateLimitTask struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Image rate limit extId.
  */
  RateLimitId *string `json:"rateLimitId,omitempty"`
  /**
  Task extId for image rate limits operations.
  */
  TaskId *string `json:"taskId,omitempty"`
}


func NewRateLimitTask() *RateLimitTask {
  p := new(RateLimitTask)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "vmm.v4.images.RateLimitTask"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.images.RateLimitTask"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST response for all response codes in api path /vmm/v4.0.a1/images/rate-limits/$actions/delete Post operation
*/
type RateLimitTaskListApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfRateLimitTaskListApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewRateLimitTaskListApiResponse() *RateLimitTaskListApiResponse {
  p := new(RateLimitTaskListApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "vmm.v4.images.RateLimitTaskListApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.images.RateLimitTaskListApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *RateLimitTaskListApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *RateLimitTaskListApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfRateLimitTaskListApiResponseData()
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
REST response for all response codes in api path /vmm/v4.0.a1/images/rate-limits/$actions/resolve Post operation
*/
type ResolveRateLimitsApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfResolveRateLimitsApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewResolveRateLimitsApiResponse() *ResolveRateLimitsApiResponse {
  p := new(ResolveRateLimitsApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "vmm.v4.images.ResolveRateLimitsApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.images.ResolveRateLimitsApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *ResolveRateLimitsApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *ResolveRateLimitsApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfResolveRateLimitsApiResponseData()
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
URL for creating an image.
*/
type UrlSource struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Ignore the certificate errors, if the value is true. Default is false.
  */
  AllowInsecure *bool `json:"allowInsecure,omitempty"`
  
  Url *string `json:"url"`
}

func (p *UrlSource) MarshalJSON() ([]byte, error) {
  type UrlSourceProxy UrlSource
  return json.Marshal(struct {
    *UrlSourceProxy
    Url *string `json:"url,omitempty"`
  }{
    UrlSourceProxy : (*UrlSourceProxy)(p),
    Url : p.Url,
  })
}

func NewUrlSource() *UrlSource {
  p := new(UrlSource)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "vmm.v4.images.UrlSource"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.images.UrlSource"}
  p.UnknownFields_ = map[string]interface{}{}

  p.AllowInsecure = new(bool)
  *p.AllowInsecure = false


  return p
}




/**
VM disk for creating an image. VM disk extId is available from the `disks` list in VMs APIs.
*/
type VmDiskSource struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  ExtId *string `json:"extId"`
}

func (p *VmDiskSource) MarshalJSON() ([]byte, error) {
  type VmDiskSourceProxy VmDiskSource
  return json.Marshal(struct {
    *VmDiskSourceProxy
    ExtId *string `json:"extId,omitempty"`
  }{
    VmDiskSourceProxy : (*VmDiskSourceProxy)(p),
    ExtId : p.ExtId,
  })
}

func NewVmDiskSource() *VmDiskSource {
  p := new(VmDiskSource)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "vmm.v4.images.VmDiskSource"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.images.VmDiskSource"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}



type OneOfRateLimitApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
  oneOfType0 *RateLimit `json:"-"`
}

func NewOneOfRateLimitApiResponseData() *OneOfRateLimitApiResponseData {
  p := new(OneOfRateLimitApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfRateLimitApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfRateLimitApiResponseData is nil"))
  }
  switch v.(type) {
    case import2.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = v.(import2.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case RateLimit:
      if nil == p.oneOfType0 {p.oneOfType0 = new(RateLimit)}
      *p.oneOfType0 = v.(RateLimit)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfRateLimitApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  return nil
}

func (p *OneOfRateLimitApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new(RateLimit)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "vmm.v4.images.RateLimit" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(RateLimit)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfRateLimitApiResponseData"))
}

func (p *OneOfRateLimitApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfRateLimitApiResponseData")
}

type OneOfImageApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
  oneOfType0 *Image `json:"-"`
}

func NewOneOfImageApiResponseData() *OneOfImageApiResponseData {
  p := new(OneOfImageApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfImageApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfImageApiResponseData is nil"))
  }
  switch v.(type) {
    case import2.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = v.(import2.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case Image:
      if nil == p.oneOfType0 {p.oneOfType0 = new(Image)}
      *p.oneOfType0 = v.(Image)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfImageApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  return nil
}

func (p *OneOfImageApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new(Image)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "vmm.v4.images.Image" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(Image)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfImageApiResponseData"))
}

func (p *OneOfImageApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfImageApiResponseData")
}

type OneOfResolveRateLimitsApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
  oneOfType0 []EffectiveRateLimit `json:"-"`
}

func NewOneOfResolveRateLimitsApiResponseData() *OneOfResolveRateLimitsApiResponseData {
  p := new(OneOfResolveRateLimitsApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfResolveRateLimitsApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfResolveRateLimitsApiResponseData is nil"))
  }
  switch v.(type) {
    case import2.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = v.(import2.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case []EffectiveRateLimit:
      p.oneOfType0 = v.([]EffectiveRateLimit)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<vmm.v4.images.EffectiveRateLimit>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<vmm.v4.images.EffectiveRateLimit>"
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfResolveRateLimitsApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if "List<vmm.v4.images.EffectiveRateLimit>" == *p.Discriminator {
    return p.oneOfType0
  }
  return nil
}

func (p *OneOfResolveRateLimitsApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new([]EffectiveRateLimit)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "vmm.v4.images.EffectiveRateLimit" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<vmm.v4.images.EffectiveRateLimit>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<vmm.v4.images.EffectiveRateLimit>"
      return nil

    }
  }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfResolveRateLimitsApiResponseData"))
}

func (p *OneOfResolveRateLimitsApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if "List<vmm.v4.images.EffectiveRateLimit>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfResolveRateLimitsApiResponseData")
}

type OneOfPlacementPolicyApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
  oneOfType0 *PlacementPolicy `json:"-"`
}

func NewOneOfPlacementPolicyApiResponseData() *OneOfPlacementPolicyApiResponseData {
  p := new(OneOfPlacementPolicyApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfPlacementPolicyApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfPlacementPolicyApiResponseData is nil"))
  }
  switch v.(type) {
    case import2.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = v.(import2.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case PlacementPolicy:
      if nil == p.oneOfType0 {p.oneOfType0 = new(PlacementPolicy)}
      *p.oneOfType0 = v.(PlacementPolicy)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfPlacementPolicyApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  return nil
}

func (p *OneOfPlacementPolicyApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new(PlacementPolicy)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "vmm.v4.images.PlacementPolicy" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(PlacementPolicy)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfPlacementPolicyApiResponseData"))
}

func (p *OneOfPlacementPolicyApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfPlacementPolicyApiResponseData")
}

type OneOfPlacementPolicyListApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
  oneOfType0 []PlacementPolicy `json:"-"`
}

func NewOneOfPlacementPolicyListApiResponseData() *OneOfPlacementPolicyListApiResponseData {
  p := new(OneOfPlacementPolicyListApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfPlacementPolicyListApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfPlacementPolicyListApiResponseData is nil"))
  }
  switch v.(type) {
    case import2.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = v.(import2.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case []PlacementPolicy:
      p.oneOfType0 = v.([]PlacementPolicy)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<vmm.v4.images.PlacementPolicy>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<vmm.v4.images.PlacementPolicy>"
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfPlacementPolicyListApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if "List<vmm.v4.images.PlacementPolicy>" == *p.Discriminator {
    return p.oneOfType0
  }
  return nil
}

func (p *OneOfPlacementPolicyListApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new([]PlacementPolicy)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "vmm.v4.images.PlacementPolicy" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<vmm.v4.images.PlacementPolicy>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<vmm.v4.images.PlacementPolicy>"
      return nil

    }
  }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfPlacementPolicyListApiResponseData"))
}

func (p *OneOfPlacementPolicyListApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if "List<vmm.v4.images.PlacementPolicy>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfPlacementPolicyListApiResponseData")
}

type OneOfImageListApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
  oneOfType0 []Image `json:"-"`
}

func NewOneOfImageListApiResponseData() *OneOfImageListApiResponseData {
  p := new(OneOfImageListApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfImageListApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfImageListApiResponseData is nil"))
  }
  switch v.(type) {
    case import2.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = v.(import2.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case []Image:
      p.oneOfType0 = v.([]Image)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<vmm.v4.images.Image>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<vmm.v4.images.Image>"
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfImageListApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if "List<vmm.v4.images.Image>" == *p.Discriminator {
    return p.oneOfType0
  }
  return nil
}

func (p *OneOfImageListApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new([]Image)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "vmm.v4.images.Image" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<vmm.v4.images.Image>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<vmm.v4.images.Image>"
      return nil

    }
  }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfImageListApiResponseData"))
}

func (p *OneOfImageListApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if "List<vmm.v4.images.Image>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfImageListApiResponseData")
}

type OneOfImagePlacementPolicyListApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 []PlacementPolicyStatus `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfImagePlacementPolicyListApiResponseData() *OneOfImagePlacementPolicyListApiResponseData {
  p := new(OneOfImagePlacementPolicyListApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfImagePlacementPolicyListApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfImagePlacementPolicyListApiResponseData is nil"))
  }
  switch v.(type) {
    case []PlacementPolicyStatus:
      p.oneOfType0 = v.([]PlacementPolicyStatus)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<vmm.v4.images.PlacementPolicyStatus>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<vmm.v4.images.PlacementPolicyStatus>"
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

func (p *OneOfImagePlacementPolicyListApiResponseData) GetValue() interface{} {
  if "List<vmm.v4.images.PlacementPolicyStatus>" == *p.Discriminator {
    return p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfImagePlacementPolicyListApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new([]PlacementPolicyStatus)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "vmm.v4.images.PlacementPolicyStatus" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<vmm.v4.images.PlacementPolicyStatus>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<vmm.v4.images.PlacementPolicyStatus>"
      return nil

    }
  }
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfImagePlacementPolicyListApiResponseData"))
}

func (p *OneOfImagePlacementPolicyListApiResponseData) MarshalJSON() ([]byte, error) {
  if "List<vmm.v4.images.PlacementPolicyStatus>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfImagePlacementPolicyListApiResponseData")
}

type OneOfImageClusterLocationListApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
  oneOfType0 []ClusterReference `json:"-"`
}

func NewOneOfImageClusterLocationListApiResponseData() *OneOfImageClusterLocationListApiResponseData {
  p := new(OneOfImageClusterLocationListApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfImageClusterLocationListApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfImageClusterLocationListApiResponseData is nil"))
  }
  switch v.(type) {
    case import2.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = v.(import2.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case []ClusterReference:
      p.oneOfType0 = v.([]ClusterReference)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<vmm.v4.images.ClusterReference>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<vmm.v4.images.ClusterReference>"
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfImageClusterLocationListApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if "List<vmm.v4.images.ClusterReference>" == *p.Discriminator {
    return p.oneOfType0
  }
  return nil
}

func (p *OneOfImageClusterLocationListApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new([]ClusterReference)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "vmm.v4.images.ClusterReference" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<vmm.v4.images.ClusterReference>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<vmm.v4.images.ClusterReference>"
      return nil

    }
  }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfImageClusterLocationListApiResponseData"))
}

func (p *OneOfImageClusterLocationListApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if "List<vmm.v4.images.ClusterReference>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfImageClusterLocationListApiResponseData")
}

type OneOfImagesTaskApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
  oneOfType0 *import3.Task `json:"-"`
}

func NewOneOfImagesTaskApiResponseData() *OneOfImagesTaskApiResponseData {
  p := new(OneOfImagesTaskApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfImagesTaskApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfImagesTaskApiResponseData is nil"))
  }
  switch v.(type) {
    case import2.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = v.(import2.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case import3.Task:
      if nil == p.oneOfType0 {p.oneOfType0 = new(import3.Task)}
      *p.oneOfType0 = v.(import3.Task)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfImagesTaskApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  return nil
}

func (p *OneOfImagesTaskApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new(import3.Task)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "vmm.v4.config.Task" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(import3.Task)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfImagesTaskApiResponseData"))
}

func (p *OneOfImagesTaskApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfImagesTaskApiResponseData")
}

type OneOfRateLimitTaskListApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 []RateLimitTask `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfRateLimitTaskListApiResponseData() *OneOfRateLimitTaskListApiResponseData {
  p := new(OneOfRateLimitTaskListApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfRateLimitTaskListApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfRateLimitTaskListApiResponseData is nil"))
  }
  switch v.(type) {
    case []RateLimitTask:
      p.oneOfType0 = v.([]RateLimitTask)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<vmm.v4.images.RateLimitTask>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<vmm.v4.images.RateLimitTask>"
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

func (p *OneOfRateLimitTaskListApiResponseData) GetValue() interface{} {
  if "List<vmm.v4.images.RateLimitTask>" == *p.Discriminator {
    return p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfRateLimitTaskListApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new([]RateLimitTask)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "vmm.v4.images.RateLimitTask" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<vmm.v4.images.RateLimitTask>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<vmm.v4.images.RateLimitTask>"
      return nil

    }
  }
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfRateLimitTaskListApiResponseData"))
}

func (p *OneOfRateLimitTaskListApiResponseData) MarshalJSON() ([]byte, error) {
  if "List<vmm.v4.images.RateLimitTask>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfRateLimitTaskListApiResponseData")
}

type OneOfRateLimitListApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
  oneOfType0 []RateLimit `json:"-"`
}

func NewOneOfRateLimitListApiResponseData() *OneOfRateLimitListApiResponseData {
  p := new(OneOfRateLimitListApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfRateLimitListApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfRateLimitListApiResponseData is nil"))
  }
  switch v.(type) {
    case import2.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = v.(import2.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case []RateLimit:
      p.oneOfType0 = v.([]RateLimit)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<vmm.v4.images.RateLimit>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<vmm.v4.images.RateLimit>"
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfRateLimitListApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if "List<vmm.v4.images.RateLimit>" == *p.Discriminator {
    return p.oneOfType0
  }
  return nil
}

func (p *OneOfRateLimitListApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new([]RateLimit)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "vmm.v4.images.RateLimit" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<vmm.v4.images.RateLimit>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<vmm.v4.images.RateLimit>"
      return nil

    }
  }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfRateLimitListApiResponseData"))
}

func (p *OneOfRateLimitListApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if "List<vmm.v4.images.RateLimit>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfRateLimitListApiResponseData")
}

type OneOfImageSource struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType1 *VmDiskSource `json:"-"`
  oneOfType0 *UrlSource `json:"-"`
}

func NewOneOfImageSource() *OneOfImageSource {
  p := new(OneOfImageSource)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfImageSource) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfImageSource is nil"))
  }
  switch v.(type) {
    case VmDiskSource:
      if nil == p.oneOfType1 {p.oneOfType1 = new(VmDiskSource)}
      *p.oneOfType1 = v.(VmDiskSource)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType1.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType1.ObjectType_
    case UrlSource:
      if nil == p.oneOfType0 {p.oneOfType0 = new(UrlSource)}
      *p.oneOfType0 = v.(UrlSource)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfImageSource) GetValue() interface{} {
  if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
    return *p.oneOfType1
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  return nil
}

func (p *OneOfImageSource) UnmarshalJSON(b []byte) error {
  vOneOfType1 := new(VmDiskSource)
  if err := json.Unmarshal(b, vOneOfType1); err == nil {
    if "vmm.v4.images.VmDiskSource" == *vOneOfType1.ObjectType_ {
          if nil == p.oneOfType1 {p.oneOfType1 = new(VmDiskSource)}
      *p.oneOfType1 = *vOneOfType1
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType1.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType1.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new(UrlSource)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "vmm.v4.images.UrlSource" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(UrlSource)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfImageSource"))
}

func (p *OneOfImageSource) MarshalJSON() ([]byte, error) {
  if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType1)
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfImageSource")
}

type OneOfImageCategoriesApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 []CategoryReference `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfImageCategoriesApiResponseData() *OneOfImageCategoriesApiResponseData {
  p := new(OneOfImageCategoriesApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfImageCategoriesApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfImageCategoriesApiResponseData is nil"))
  }
  switch v.(type) {
    case []CategoryReference:
      p.oneOfType0 = v.([]CategoryReference)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<vmm.v4.images.CategoryReference>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<vmm.v4.images.CategoryReference>"
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

func (p *OneOfImageCategoriesApiResponseData) GetValue() interface{} {
  if "List<vmm.v4.images.CategoryReference>" == *p.Discriminator {
    return p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfImageCategoriesApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new([]CategoryReference)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "vmm.v4.images.CategoryReference" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<vmm.v4.images.CategoryReference>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<vmm.v4.images.CategoryReference>"
      return nil

    }
  }
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfImageCategoriesApiResponseData"))
}

func (p *OneOfImageCategoriesApiResponseData) MarshalJSON() ([]byte, error) {
  if "List<vmm.v4.images.CategoryReference>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfImageCategoriesApiResponseData")
}

