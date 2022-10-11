/*
 * Generated file models/files/v4/config/config_model.go.
 *
 * Product version: 4.0.1-alpha-2
 *
 * Part of the Nutanix Files Versioned APIs
 *
 * (c) 2022 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module for config APIs for file server
*/
package config
import (
  "bytes"
  import3 "github.com/nutanix/ntnx-api-golang-clients/files-go-client/v4/models/common/v1/config"
  import1 "github.com/nutanix/ntnx-api-golang-clients/files-go-client/v4/models/common/v1/response"
  "encoding/json"
  "errors"
  import2 "github.com/nutanix/ntnx-api-golang-clients/files-go-client/v4/models/files/v4/error"
  "fmt"
  import4 "github.com/nutanix/ntnx-api-golang-clients/files-go-client/v4/models/prism/v4/config"
)

/**
Active directory DNS failover
*/
type ADDnsFailover struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import1.ApiLink `json:"links,omitempty"`
  
  PrimaryAdCredential *Credential `json:"primaryAdCredential,omitempty"`
  /**
  Use a specific domain controller for the join-domain operation in a multi DC active directory setup. By default, AFS discovers a site local domain controller for join-domain operation. The preferred domain controller cannot be an IP address. It has to be FQDN (example: dc_name.dom.companyname.com)
  */
  PrimaryAdPreferredDomainController *string `json:"primaryAdPreferredDomainController,omitempty"`
  
  PrimaryDns *Dns `json:"primaryDns,omitempty"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
}


func NewADDnsFailover() *ADDnsFailover {
  p := new(ADDnsFailover)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.ADDnsFailover"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.ADDnsFailover"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Access type of the mount target. The acceptable values are "READ_WRITE", "READ_ONLY", "NO_ACCESS".
*/
type AccessType int

const(
  ACCESSTYPE_UNKNOWN AccessType = 0
  ACCESSTYPE_REDACTED AccessType = 1
  ACCESSTYPE_READ_WRITE AccessType = 2
  ACCESSTYPE_READ_ONLY AccessType = 3
  ACCESSTYPE_NO_ACCESS AccessType = 4
)

// returns the name of the enum given an ordinal number
func (e *AccessType) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "READ_WRITE",
    "READ_ONLY",
    "NO_ACCESS",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *AccessType) index(name string) AccessType {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "READ_WRITE",
    "READ_ONLY",
    "NO_ACCESS",
  }
  for idx := range names {
    if names[idx] == name {
      return AccessType(idx)
    }
  }
  return ACCESSTYPE_UNKNOWN
}

func (e *AccessType) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for AccessType:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *AccessType) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e AccessType) Ref() *AccessType {
  return &e
}


/**
REST response for all response codes in api path /files/v4.0.a2/config/file-server/alerts/$actions/acknowledge Post operation
*/
type AcknowledgeAlertsApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfAcknowledgeAlertsApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewAcknowledgeAlertsApiResponse() *AcknowledgeAlertsApiResponse {
  p := new(AcknowledgeAlertsApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.AcknowledgeAlertsApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.AcknowledgeAlertsApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *AcknowledgeAlertsApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *AcknowledgeAlertsApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfAcknowledgeAlertsApiResponseData()
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
Response for the acknowledge alert request.
*/
type AcknowledgeAlertsResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Alert UUID
  */
  Id *string `json:"id,omitempty"`
  /**
  Message corresponding to alert acknowledge request.
  */
  Message *string `json:"message,omitempty"`
  /**
  This field indicates if the alert is acknowledged.
  */
  Successful *bool `json:"successful,omitempty"`
}


func NewAcknowledgeAlertsResponse() *AcknowledgeAlertsResponse {
  p := new(AcknowledgeAlertsResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.AcknowledgeAlertsResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.AcknowledgeAlertsResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Failover status and action details.
*/
type ActionStatus struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  Action *FailOverAction `json:"action,omitempty"`
  /**
  A code that uniquely identifies a message.
  */
  Code *string `json:"code,omitempty"`
  /**
  The locale for the message description.
  */
  Locale *string `json:"locale,omitempty"`
  
  Message *string `json:"message,omitempty"`
  
  Severity *import3.MessageSeverity `json:"severity,omitempty"`
  
  Status *FailOverStatus `json:"status,omitempty"`
}


func NewActionStatus() *ActionStatus {
  p := new(ActionStatus)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.ActionStatus"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.ActionStatus"}
  p.UnknownFields_ = map[string]interface{}{}

  p.Locale = new(string)
  *p.Locale = "en_US"


  return p
}




/**
REST response for all response codes in api path /files/v4.0.a2/config/file-server/data-protection/$actions/activate Post operation
*/
type ActivateFileServerApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfActivateFileServerApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewActivateFileServerApiResponse() *ActivateFileServerApiResponse {
  p := new(ActivateFileServerApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.ActivateFileServerApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.ActivateFileServerApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *ActivateFileServerApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *ActivateFileServerApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfActivateFileServerApiResponseData()
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
Ad domain associated with the file server.
*/
type Ad struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Indicates whether the Ad credentials will be validated or skipped. This field will be deferred.
  */
  ValidateAdCredential *bool `json:"ValidateAdCredential,omitempty"`
  /**
  Ad user or group name as 'name' or 'NETBIOS\name' format.
  */
  AddUserAsFsAdmin *bool `json:"addUserAsFsAdmin,omitempty"`
  /**
  Ad domain name associated with the file server.
  */
  DomainName *string `json:"domainName,omitempty"`
  /**
  If the joined active directory is down or not reachable, then use this option to leave the domain by bypassing all Ad interactions. Please ensure to remove the file server machine account manually afterwards.
  */
  ForceLeaveDomain *bool `json:"forceLeaveDomain,omitempty"`
  /**
  Indicates whether the Ad credentials will be validated or skipped.
  */
  IsValidateAdCredential *bool `json:"isValidateAdCredential,omitempty"`
  /**
  Ad domain organizational unit associated with the file server.
  */
  OrganizationUnit *string `json:"organizationUnit,omitempty"`
  /**
  If a machine account with the same name as file server name is present on the Ad, then overwrite it during join-domain operation.
  */
  OverwriteUserAccount *bool `json:"overwriteUserAccount,omitempty"`
  /**
  Ad domain password associated with the file server.
  */
  Password *string `json:"password,omitempty"`
  /**
  Use a specific domain controller for the join-domain operation in a multi DC active directory setup. By default, AFS discovers a site local domain controller for join-domain operation. The preferred domain controller cannot be an IP address. It has to be FQDN (example: dc_name.dom.companyname.com)
  */
  PreferredDomainController *string `json:"preferredDomainController,omitempty"`
  
  ProtocolType *ProtocolType `json:"protocolType,omitempty"`
  /**
  RFC 2307 ENABLED (true, false).
  */
  Rfc2307Enabled *bool `json:"rfc2307Enabled,omitempty"`
  /**
  Use the same Ad credential for DNS.
  */
  UseSameCredentialsForDns *bool `json:"useSameCredentialsForDns,omitempty"`
  /**
  Ad domain username associated with the file server.
  */
  Username *string `json:"username,omitempty"`
}


func NewAd() *Ad {
  p := new(Ad)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.Ad"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.Ad"}
  p.UnknownFields_ = map[string]interface{}{}

  p.ValidateAdCredential = new(bool)
  *p.ValidateAdCredential = false
  p.AddUserAsFsAdmin = new(bool)
  *p.AddUserAsFsAdmin = false
  p.ForceLeaveDomain = new(bool)
  *p.ForceLeaveDomain = false
  p.IsValidateAdCredential = new(bool)
  *p.IsValidateAdCredential = false
  p.OverwriteUserAccount = new(bool)
  *p.OverwriteUserAccount = false
  p.Rfc2307Enabled = new(bool)
  *p.Rfc2307Enabled = false
  p.UseSameCredentialsForDns = new(bool)
  *p.UseSameCredentialsForDns = false


  return p
}




/**
REST response for all response codes in api path /files/v4.0.a2/config/file-server/data-protection/$actions/ad-dns-failover Post operation
*/
type AdDnsFailoverApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfAdDnsFailoverApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewAdDnsFailoverApiResponse() *AdDnsFailoverApiResponse {
  p := new(AdDnsFailoverApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.AdDnsFailoverApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.AdDnsFailoverApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *AdDnsFailoverApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *AdDnsFailoverApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfAdDnsFailoverApiResponseData()
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
REST response for all response codes in api path /files/v4.0.a2/config/file-server/$actions/add-dns-entries Post operation
*/
type AddDnsEntriesApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfAddDnsEntriesApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewAddDnsEntriesApiResponse() *AddDnsEntriesApiResponse {
  p := new(AddDnsEntriesApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.AddDnsEntriesApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.AddDnsEntriesApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *AddDnsEntriesApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *AddDnsEntriesApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfAddDnsEntriesApiResponseData()
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
Admin role. The acceptable values are "ADMIN" or "BACKUP_OPERATOR".
*/
type AdminRole int

const(
  ADMINROLE_UNKNOWN AdminRole = 0
  ADMINROLE_REDACTED AdminRole = 1
  ADMINROLE_ADMIN AdminRole = 2
  ADMINROLE_BACKUP_OPERATOR AdminRole = 3
)

// returns the name of the enum given an ordinal number
func (e *AdminRole) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "ADMIN",
    "BACKUP_OPERATOR",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *AdminRole) index(name string) AdminRole {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "ADMIN",
    "BACKUP_OPERATOR",
  }
  for idx := range names {
    if names[idx] == name {
      return AdminRole(idx)
    }
  }
  return ADMINROLE_UNKNOWN
}

func (e *AdminRole) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for AdminRole:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *AdminRole) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e AdminRole) Ref() *AdminRole {
  return &e
}


/**
Admin user model
*/
type AdminUser struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import1.ApiLink `json:"links,omitempty"`
  /**
  Admin user name
  */
  Name *string `json:"name"`
  
  Role *AdminRole `json:"role"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
}

func (p *AdminUser) MarshalJSON() ([]byte, error) {
  type AdminUserProxy AdminUser
  return json.Marshal(struct {
    *AdminUserProxy
    Name *string `json:"name,omitempty"`
    Role *AdminRole `json:"role,omitempty"`
  }{
    AdminUserProxy : (*AdminUserProxy)(p),
    Name : p.Name,
    Role : p.Role,
  })
}

func NewAdminUser() *AdminUser {
  p := new(AdminUser)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.AdminUser"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.AdminUser"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST response for all response codes in api path /files/v4.0.a2/config/file-server/admin-users/{adminUserExtId} Get operation
*/
type AdminUserApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfAdminUserApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewAdminUserApiResponse() *AdminUserApiResponse {
  p := new(AdminUserApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.AdminUserApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.AdminUserApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *AdminUserApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *AdminUserApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfAdminUserApiResponseData()
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
REST response for all response codes in api path /files/v4.0.a2/config/file-server/admin-users Get operation
*/
type AdminUserListApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfAdminUserListApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewAdminUserListApiResponse() *AdminUserListApiResponse {
  p := new(AdminUserListApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.AdminUserListApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.AdminUserListApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *AdminUserListApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *AdminUserListApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfAdminUserListApiResponseData()
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
Alert object details
*/
type Alert struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Alerts which have acknowledged status.
  */
  Acknowledged *bool `json:"acknowledged,omitempty"`
  /**
  Username who acknowledged the alerts.
  */
  AcknowledgedByUsername *string `json:"acknowledgedByUsername,omitempty"`
  /**
  Acknowledged timestamp in microseconds.
  */
  AcknowledgedTimeStampInUsecs *int64 `json:"acknowledgedTimeStampInUsecs,omitempty"`
  /**
  List of affected entities.
  */
  AffectedEntities []interface{} `json:"affectedEntities,omitempty"`
  /**
  Alert description
  */
  AlertDetails *interface{} `json:"alertDetails,omitempty"`
  /**
  Title of the alert.
  */
  AlertTitle *string `json:"alertTitle,omitempty"`
  /**
  Alert type UUID
  */
  AlertTypeUuid *string `json:"alertTypeUuid,omitempty"`
  /**
  Alerts which have auto-resolve status.
  */
  AutoResolved *bool `json:"autoResolved,omitempty"`
  /**
  Alerts associated with the categories. Categories can be SystemIndicator, Cluster.
  */
  Categories []string `json:"categories,omitempty"`
  /**
  Alert check UUID
  */
  CheckId *string `json:"checkId,omitempty"`
  /**
  Alerts associated with the listed classification.
  */
  Classifications []string `json:"classifications,omitempty"`
  /**
  The cluster associated with the file server. This contains AOS cluster UUID. This is a read-only field.
  */
  ClusterUuid *string `json:"clusterUuid,omitempty"`
  /**
  List of alert context types.
  */
  ContextTypes []string `json:"contextTypes,omitempty"`
  /**
  List of alert context values.
  */
  ContextValues []string `json:"contextValues,omitempty"`
  /**
  Created timestamp in microseconds.
  */
  CreatedTimeStampInUsecs *int64 `json:"createdTimeStampInUsecs,omitempty"`
  /**
  Detailed alerts message.
  */
  DetailedMessage *string `json:"detailedMessage,omitempty"`
  /**
  List of entity extIds.
  */
  EntityIds []string `json:"entityIds,omitempty"`
  /**
  Entity type to get the alerts.
  */
  EntityTypes []string `json:"entityTypes,omitempty"`
  /**
  List of UUID for entities.
  */
  EntityUuids []string `json:"entityUuids,omitempty"`
  /**
  Alert UUID
  */
  Id *string `json:"id,omitempty"`
  /**
  Alert impact
  */
  Impact *string `json:"impact,omitempty"`
  /**
  Alerts associated with the impact types. Impact types can be Availability, Capacity, Configuration, Performance, System Indicator, Controller VM.
  */
  ImpactTypes []string `json:"impactTypes,omitempty"`
  /**
  Last occurred timestamp in microseconds.
  */
  LastOccurrenceTimeStampInUsecs *int64 `json:"lastOccurrenceTimeStampInUsecs,omitempty"`
  /**
  Alerts message
  */
  Message *string `json:"message,omitempty"`
  /**
  Node UUID
  */
  NodeUuid *string `json:"nodeUuid,omitempty"`
  /**
  Operation type of alerts.
  */
  OperationType *string `json:"operationType,omitempty"`
  /**
  Originating cluster UUID
  */
  OriginatingClusterUuid *string `json:"originatingClusterUuid,omitempty"`
  /**
  The possible causes for this alert.
  */
  PossibleCauses []interface{} `json:"possibleCauses,omitempty"`
  /**
  Used as a filter to fetch only the resolved alerts.
  */
  Resolved *bool `json:"resolved,omitempty"`
  /**
  Username who resolved the alerts.
  */
  ResolvedByUsername *string `json:"resolvedByUsername,omitempty"`
  /**
  Resolved timestamp in microseconds.
  */
  ResolvedTimeStampInUsecs *int64 `json:"resolvedTimeStampInUsecs,omitempty"`
  /**
  Service VM Id
  */
  ServiceVMId *string `json:"serviceVMId,omitempty"`
  /**
  Severity level of the alert. The values can be INFO, WARNING or CRITICAL.
  */
  Severity *string `json:"severity,omitempty"`
  /**
  Indicates whether the alert is user-defined.
  */
  UserDefined *bool `json:"userDefined,omitempty"`
}


func NewAlert() *Alert {
  p := new(Alert)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.Alert"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.Alert"}
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




/**
REST response for all response codes in api path /files/v4.0.a2/config/file-server/alerts Get operation
*/
type AlertApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfAlertApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewAlertApiResponse() *AlertApiResponse {
  p := new(AlertApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.AlertApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.AlertApiResponse"}
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



/**
Anonymous identifier of the mount target.
*/
type AnonymousIdentifier struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Group identifier. Default value of the group identifier is -2.
  */
  Gid *int64 `json:"gid,omitempty"`
  /**
  User identifier. Default value of the user identifier is -2.
  */
  Uid *int64 `json:"uid,omitempty"`
}


func NewAnonymousIdentifier() *AnonymousIdentifier {
  p := new(AnonymousIdentifier)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.AnonymousIdentifier"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.AnonymousIdentifier"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Antivirus server model
*/
type AntivirusServer struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  Address *import3.IPAddressOrFQDN `json:"address"`
  
  ConnectionStatus *ConnectionStatus `json:"connectionStatus,omitempty"`
  /**
  Antivirus server description
  */
  Description *string `json:"description,omitempty"`
  /**
  This field indicates if the antivirus server is enabled. This is a read-only field. This field will be deprecated.
  */
  Enable *bool `json:"enable,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  ICAP service name
  */
  IcapServiceName *string `json:"icapServiceName,omitempty"`
  /**
  This field indicates if the antivirus server is enabled. This is a read-only field.
  */
  IsEnable *bool `json:"isEnable,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import1.ApiLink `json:"links,omitempty"`
  /**
  Partner details of the antivirus server. This is a read-only field.
  */
  Partner *string `json:"partner,omitempty"`
  /**
  Antivirus server port
  */
  Port *int `json:"port"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
}

func (p *AntivirusServer) MarshalJSON() ([]byte, error) {
  type AntivirusServerProxy AntivirusServer
  return json.Marshal(struct {
    *AntivirusServerProxy
    Address *import3.IPAddressOrFQDN `json:"address,omitempty"`
    Port *int `json:"port,omitempty"`
  }{
    AntivirusServerProxy : (*AntivirusServerProxy)(p),
    Address : p.Address,
    Port : p.Port,
  })
}

func NewAntivirusServer() *AntivirusServer {
  p := new(AntivirusServer)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.AntivirusServer"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.AntivirusServer"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST response for all response codes in api path /files/v4.0.a2/config/file-server/anti-virus-servers/{antivirusServerExtId} Get operation
*/
type AntivirusServerApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfAntivirusServerApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewAntivirusServerApiResponse() *AntivirusServerApiResponse {
  p := new(AntivirusServerApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.AntivirusServerApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.AntivirusServerApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *AntivirusServerApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *AntivirusServerApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfAntivirusServerApiResponseData()
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
REST response for all response codes in api path /files/v4.0.a2/config/file-server/anti-virus-servers Get operation
*/
type AntivirusServerListApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfAntivirusServerListApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewAntivirusServerListApiResponse() *AntivirusServerListApiResponse {
  p := new(AntivirusServerListApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.AntivirusServerListApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.AntivirusServerListApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *AntivirusServerListApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *AntivirusServerListApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfAntivirusServerListApiResponseData()
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
Auth Type Details.
*/
type AuthType int

const(
  AUTHTYPE_UNKNOWN AuthType = 0
  AUTHTYPE_REDACTED AuthType = 1
  AUTHTYPE_KRB5 AuthType = 2
  AUTHTYPE_NONE AuthType = 3
  AUTHTYPE_SYS AuthType = 4
  AUTHTYPE_KRB5P AuthType = 5
  AUTHTYPE_KRB5I AuthType = 6
)

// returns the name of the enum given an ordinal number
func (e *AuthType) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "KRB5",
    "NONE",
    "SYS",
    "KRB5P",
    "KRB5I",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *AuthType) index(name string) AuthType {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "KRB5",
    "NONE",
    "SYS",
    "KRB5P",
    "KRB5I",
  }
  for idx := range names {
    if names[idx] == name {
      return AuthType(idx)
    }
  }
  return AUTHTYPE_UNKNOWN
}

func (e *AuthType) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for AuthType:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *AuthType) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e AuthType) Ref() *AuthType {
  return &e
}


/**
Protected file server pair model.
*/
type BaseProtectedFileServerPair struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import1.ApiLink `json:"links,omitempty"`
  
  PrimaryAdCredential *Credential `json:"primaryAdCredential,omitempty"`
  /**
  Use a specific domain controller for the join-domain operation in a multi DC active directory setup. By default, AFS discovers a site local domain controller for join-domain operation. The preferred domain controller cannot be an IP address. It has to be FQDN (example: dc_name.dom.companyname.com)
  */
  PrimaryAdPreferredDomainController *string `json:"primaryAdPreferredDomainController,omitempty"`
  
  PrimaryDns *Dns `json:"primaryDns,omitempty"`
  /**
  The {extId} of the primary file server. This is a read-only field.
  */
  PrimaryFileServerExtId *string `json:"primaryFileServerExtId,omitempty"`
  /**
  The {extId} of the target file server. This is a read-only field.
  */
  TargetFileServerExtId *string `json:"targetFileServerExtId,omitempty"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
}


func NewBaseProtectedFileServerPair() *BaseProtectedFileServerPair {
  p := new(BaseProtectedFileServerPair)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.BaseProtectedFileServerPair"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.BaseProtectedFileServerPair"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Blocked clients due to read-only or no access.
*/
type BlockedClients struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  No access clients
  */
  NoAccessFilters []ClientBlockingFilter `json:"noAccessFilters,omitempty"`
  /**
  Read only access clients
  */
  RoAccessFilters []ClientBlockingFilter `json:"roAccessFilters,omitempty"`
  /**
  Read write access clients
  */
  RwAccessFilters []ClientBlockingFilter `json:"rwAccessFilters,omitempty"`
}


func NewBlockedClients() *BlockedClients {
  p := new(BlockedClients)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.BlockedClients"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.BlockedClients"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST response for all response codes in api path /files/v4.0.a2/config/file-server/$actions/clear-recommendation Post operation
*/
type ClearLoadBalanceFileServerApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfClearLoadBalanceFileServerApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewClearLoadBalanceFileServerApiResponse() *ClearLoadBalanceFileServerApiResponse {
  p := new(ClearLoadBalanceFileServerApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.ClearLoadBalanceFileServerApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.ClearLoadBalanceFileServerApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *ClearLoadBalanceFileServerApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *ClearLoadBalanceFileServerApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfClearLoadBalanceFileServerApiResponseData()
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
Access type of the mount target.
*/
type ClientAccessType struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  ClientExceptions []ClientException `json:"clientExceptions,omitempty"`
  
  DefaultAccessType *AccessType `json:"defaultAccessType,omitempty"`
}


func NewClientAccessType() *ClientAccessType {
  p := new(ClientAccessType)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.ClientAccessType"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.ClientAccessType"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Blocked clients filter
*/
type ClientBlockingFilter struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Block all IP addresses. This field will be deprecated.
  */
  BlockAllIps *bool `json:"blockAllIps,omitempty"`
  /**
  GID list
  */
  GidList []int64 `json:"gidList,omitempty"`
  /**
  IP address list
  */
  IpList []import3.IPAddress `json:"ipList,omitempty"`
  /**
  Block all IP addresses
  */
  IsBlockAllIps *bool `json:"isBlockAllIps,omitempty"`
  /**
  SID list
  */
  SidList []string `json:"sidList,omitempty"`
  /**
  UID list
  */
  UidList []int64 `json:"uidList,omitempty"`
  /**
  Vendor name
  */
  VendorName *string `json:"vendorName,omitempty"`
}


func NewClientBlockingFilter() *ClientBlockingFilter {
  p := new(ClientBlockingFilter)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.ClientBlockingFilter"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.ClientBlockingFilter"}
  p.UnknownFields_ = map[string]interface{}{}

  p.BlockAllIps = new(bool)
  *p.BlockAllIps = false
  p.IsBlockAllIps = new(bool)
  *p.IsBlockAllIps = false


  return p
}




/**
Access exception client list of the mount target.
*/
type ClientException struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  No-access client list.
  */
  NoAccesses *string `json:"noAccesses,omitempty"`
  /**
  Read-only client access list.
  */
  ReadOnlyAccesses *string `json:"readOnlyAccesses,omitempty"`
  /**
  Read-write client access list.
  */
  ReadWriteAccesses *string `json:"readWriteAccesses,omitempty"`
}


func NewClientException() *ClientException {
  p := new(ClientException)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.ClientException"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.ClientException"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST response for all response codes in api path /files/v4.0.a2/config/file-server/object-store/profiles Post operation
*/
type ConfigObjectStoreApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfConfigObjectStoreApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewConfigObjectStoreApiResponse() *ConfigObjectStoreApiResponse {
  p := new(ConfigObjectStoreApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.ConfigObjectStoreApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.ConfigObjectStoreApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *ConfigObjectStoreApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *ConfigObjectStoreApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfConfigObjectStoreApiResponseData()
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
REST response for all response codes in api path /files/v4.0.a2/config/file-server/$actions/configure-name-services Post operation
*/
type ConfigureNameServiceApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfConfigureNameServiceApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewConfigureNameServiceApiResponse() *ConfigureNameServiceApiResponse {
  p := new(ConfigureNameServiceApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.ConfigureNameServiceApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.ConfigureNameServiceApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *ConfigureNameServiceApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *ConfigureNameServiceApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfConfigureNameServiceApiResponseData()
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
Connection status values.The acceptable values are "NOT_TESTED", "NOT_REACHABLE", "OK".
*/
type ConnectionStatus int

const(
  CONNECTIONSTATUS_UNKNOWN ConnectionStatus = 0
  CONNECTIONSTATUS_REDACTED ConnectionStatus = 1
  CONNECTIONSTATUS_NOT_TESTED ConnectionStatus = 2
  CONNECTIONSTATUS_NOT_REACHABLE ConnectionStatus = 3
  CONNECTIONSTATUS_OK ConnectionStatus = 4
)

// returns the name of the enum given an ordinal number
func (e *ConnectionStatus) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "NOT_TESTED",
    "NOT_REACHABLE",
    "OK",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *ConnectionStatus) index(name string) ConnectionStatus {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "NOT_TESTED",
    "NOT_REACHABLE",
    "OK",
  }
  for idx := range names {
    if names[idx] == name {
      return ConnectionStatus(idx)
    }
  }
  return CONNECTIONSTATUS_UNKNOWN
}

func (e *ConnectionStatus) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for ConnectionStatus:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *ConnectionStatus) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e ConnectionStatus) Ref() *ConnectionStatus {
  return &e
}


/**
This contains the source paths configuration belong to the source file server.
*/
type ConsolidationEntity struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Absolute source path.

For the source paths ending with '/', the contents will be copied directly into the target folder.\n
For the paths, ending with name, a new folder with the same name will be created on the
target server and then the data is synced."
  */
  SourcePath *string `json:"sourcePath"`
}

func (p *ConsolidationEntity) MarshalJSON() ([]byte, error) {
  type ConsolidationEntityProxy ConsolidationEntity
  return json.Marshal(struct {
    *ConsolidationEntityProxy
    SourcePath *string `json:"sourcePath,omitempty"`
  }{
    ConsolidationEntityProxy : (*ConsolidationEntityProxy)(p),
    SourcePath : p.SourcePath,
  })
}

func NewConsolidationEntity() *ConsolidationEntity {
  p := new(ConsolidationEntity)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.ConsolidationEntity"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.ConsolidationEntity"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Consolidation sub policy represents source, source paths and schedules.
*/
type ConsolidationSubPolicy struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  ExtId *string `json:"extId,omitempty"`
  
  PolicyStatus *PolicyStatus `json:"policyStatus,omitempty"`
  /**
  This contains details of snapshot schedules. It contains the start time of the replication and frequency of the replication.
  */
  Schedules []Schedule `json:"schedules"`
  /**
  The {extId} of the file server. This is a read-only field.
  */
  SourceFileServerExtId *string `json:"sourceFileServerExtId"`
  /**
  This contains list of source paths which belong to the source file server.
  */
  SyncEntities []ConsolidationEntity `json:"syncEntities,omitempty"`
}

func (p *ConsolidationSubPolicy) MarshalJSON() ([]byte, error) {
  type ConsolidationSubPolicyProxy ConsolidationSubPolicy
  return json.Marshal(struct {
    *ConsolidationSubPolicyProxy
    Schedules []Schedule `json:"schedules,omitempty"`
    SourceFileServerExtId *string `json:"sourceFileServerExtId,omitempty"`
  }{
    ConsolidationSubPolicyProxy : (*ConsolidationSubPolicyProxy)(p),
    Schedules : p.Schedules,
    SourceFileServerExtId : p.SourceFileServerExtId,
  })
}

func NewConsolidationSubPolicy() *ConsolidationSubPolicy {
  p := new(ConsolidationSubPolicy)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.ConsolidationSubPolicy"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.ConsolidationSubPolicy"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST response for all response codes in api path /files/v4.0.a2/config/file-server/admin-users Post operation
*/
type CreateAdminUserApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfCreateAdminUserApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewCreateAdminUserApiResponse() *CreateAdminUserApiResponse {
  p := new(CreateAdminUserApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.CreateAdminUserApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.CreateAdminUserApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *CreateAdminUserApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *CreateAdminUserApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfCreateAdminUserApiResponseData()
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
REST response for all response codes in api path /files/v4.0.a2/config/file-server/anti-virus-servers Post operation
*/
type CreateAntivirusServerApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfCreateAntivirusServerApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewCreateAntivirusServerApiResponse() *CreateAntivirusServerApiResponse {
  p := new(CreateAntivirusServerApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.CreateAntivirusServerApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.CreateAntivirusServerApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *CreateAntivirusServerApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *CreateAntivirusServerApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfCreateAntivirusServerApiResponseData()
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
REST response for all response codes in api path /files/v4.0.a2/config/file-server/data-consolidation-policies Post operation
*/
type CreateDataConsolidationPolicyApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfCreateDataConsolidationPolicyApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewCreateDataConsolidationPolicyApiResponse() *CreateDataConsolidationPolicyApiResponse {
  p := new(CreateDataConsolidationPolicyApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.CreateDataConsolidationPolicyApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.CreateDataConsolidationPolicyApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *CreateDataConsolidationPolicyApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *CreateDataConsolidationPolicyApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfCreateDataConsolidationPolicyApiResponseData()
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
REST response for all response codes in api path /files/v4.0.a2/config/file-server/data-protection-policies Post operation
*/
type CreateDataProtectionPolicyApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfCreateDataProtectionPolicyApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewCreateDataProtectionPolicyApiResponse() *CreateDataProtectionPolicyApiResponse {
  p := new(CreateDataProtectionPolicyApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.CreateDataProtectionPolicyApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.CreateDataProtectionPolicyApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *CreateDataProtectionPolicyApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *CreateDataProtectionPolicyApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfCreateDataProtectionPolicyApiResponseData()
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
REST response for all response codes in api path /files/v4.0.a2/config/file-server/users Post operation
*/
type CreateFileServerUserApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfCreateFileServerUserApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewCreateFileServerUserApiResponse() *CreateFileServerUserApiResponse {
  p := new(CreateFileServerUserApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.CreateFileServerUserApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.CreateFileServerUserApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *CreateFileServerUserApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *CreateFileServerUserApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfCreateFileServerUserApiResponseData()
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
REST response for all response codes in api path /files/v4.0.a2/config/file-server/mount-targets Post operation
*/
type CreateMountTargetApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfCreateMountTargetApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewCreateMountTargetApiResponse() *CreateMountTargetApiResponse {
  p := new(CreateMountTargetApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.CreateMountTargetApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.CreateMountTargetApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *CreateMountTargetApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *CreateMountTargetApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfCreateMountTargetApiResponseData()
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
REST response for all response codes in api path /files/v4.0.a2/config/file-server/mount-targets/{mountTargetExtId}/virus-scan-policy Post operation
*/
type CreateMountTargetVirusScanPolicyApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfCreateMountTargetVirusScanPolicyApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewCreateMountTargetVirusScanPolicyApiResponse() *CreateMountTargetVirusScanPolicyApiResponse {
  p := new(CreateMountTargetVirusScanPolicyApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.CreateMountTargetVirusScanPolicyApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.CreateMountTargetVirusScanPolicyApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *CreateMountTargetVirusScanPolicyApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *CreateMountTargetVirusScanPolicyApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfCreateMountTargetVirusScanPolicyApiResponseData()
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
REST response for all response codes in api path /files/v4.0.a2/config/file-server/networks Post operation
*/
type CreateNetworkApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfCreateNetworkApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewCreateNetworkApiResponse() *CreateNetworkApiResponse {
  p := new(CreateNetworkApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.CreateNetworkApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.CreateNetworkApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *CreateNetworkApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *CreateNetworkApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfCreateNetworkApiResponseData()
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
REST response for all response codes in api path /files/v4.0.a2/config/file-server/mount-targets/{mountTargetExtId}/quota-policies Post operation
*/
type CreateQuotaPolicyApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfCreateQuotaPolicyApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewCreateQuotaPolicyApiResponse() *CreateQuotaPolicyApiResponse {
  p := new(CreateQuotaPolicyApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.CreateQuotaPolicyApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.CreateQuotaPolicyApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *CreateQuotaPolicyApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *CreateQuotaPolicyApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfCreateQuotaPolicyApiResponseData()
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
REST response for all response codes in api path /files/v4.0.a2/config/file-server/ssr-snapshot-schedules Post operation
*/
type CreateSsrSnapshotScheduleApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfCreateSsrSnapshotScheduleApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewCreateSsrSnapshotScheduleApiResponse() *CreateSsrSnapshotScheduleApiResponse {
  p := new(CreateSsrSnapshotScheduleApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.CreateSsrSnapshotScheduleApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.CreateSsrSnapshotScheduleApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *CreateSsrSnapshotScheduleApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *CreateSsrSnapshotScheduleApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfCreateSsrSnapshotScheduleApiResponseData()
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
REST response for all response codes in api path /files/v4.0.a2/config/file-server/tier-config Post operation
*/
type CreateTieringConfigApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfCreateTieringConfigApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewCreateTieringConfigApiResponse() *CreateTieringConfigApiResponse {
  p := new(CreateTieringConfigApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.CreateTieringConfigApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.CreateTieringConfigApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *CreateTieringConfigApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *CreateTieringConfigApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfCreateTieringConfigApiResponseData()
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
REST response for all response codes in api path /files/v4.0.a2/config/file-server/tier-policies Post operation
*/
type CreateTieringPolicyApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfCreateTieringPolicyApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewCreateTieringPolicyApiResponse() *CreateTieringPolicyApiResponse {
  p := new(CreateTieringPolicyApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.CreateTieringPolicyApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.CreateTieringPolicyApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *CreateTieringPolicyApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *CreateTieringPolicyApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfCreateTieringPolicyApiResponseData()
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
REST response for all response codes in api path /files/v4.0.a2/config/file-server/user-mapping Post operation
*/
type CreateUserMappingApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfCreateUserMappingApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewCreateUserMappingApiResponse() *CreateUserMappingApiResponse {
  p := new(CreateUserMappingApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.CreateUserMappingApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.CreateUserMappingApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *CreateUserMappingApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *CreateUserMappingApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfCreateUserMappingApiResponseData()
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
REST response for all response codes in api path /files/v4.0.a2/config/file-server/virus-scan-policy Post operation
*/
type CreateVirusScanPolicyApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfCreateVirusScanPolicyApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewCreateVirusScanPolicyApiResponse() *CreateVirusScanPolicyApiResponse {
  p := new(CreateVirusScanPolicyApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.CreateVirusScanPolicyApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.CreateVirusScanPolicyApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *CreateVirusScanPolicyApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *CreateVirusScanPolicyApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfCreateVirusScanPolicyApiResponseData()
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
Details of the user actions to be performed.
*/
type Credential struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Ad domain password associated with the file server.
  */
  Password *string `json:"password"`
  /**
  Ad domain username associated with the file server.
  */
  Username *string `json:"username"`
}

func (p *Credential) MarshalJSON() ([]byte, error) {
  type CredentialProxy Credential
  return json.Marshal(struct {
    *CredentialProxy
    Password *string `json:"password,omitempty"`
    Username *string `json:"username,omitempty"`
  }{
    CredentialProxy : (*CredentialProxy)(p),
    Password : p.Password,
    Username : p.Username,
  })
}

func NewCredential() *Credential {
  p := new(Credential)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.Credential"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.Credential"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Data consolidation policy
*/
type DataConsolidationPolicy struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Policy description
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
  Consolidation policy name
  */
  Name *string `json:"name"`
  
  PolicyStatus *PolicyStatus `json:"policyStatus,omitempty"`
  /**
  Consolidation sub policy represents source, source paths and schedules.
  */
  SubPolicies []ConsolidationSubPolicy `json:"subPolicies"`
  /**
  The {extId} of the target file server. This is a read-only field.
  */
  TargetFileServerExtId *string `json:"targetFileServerExtId"`
  /**
  Absolute path for the target. For example: '/share1/dir1'. Path should always start with a root folder.
  */
  TargetPath *string `json:"targetPath"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
}

func (p *DataConsolidationPolicy) MarshalJSON() ([]byte, error) {
  type DataConsolidationPolicyProxy DataConsolidationPolicy
  return json.Marshal(struct {
    *DataConsolidationPolicyProxy
    Name *string `json:"name,omitempty"`
    SubPolicies []ConsolidationSubPolicy `json:"subPolicies,omitempty"`
    TargetFileServerExtId *string `json:"targetFileServerExtId,omitempty"`
    TargetPath *string `json:"targetPath,omitempty"`
  }{
    DataConsolidationPolicyProxy : (*DataConsolidationPolicyProxy)(p),
    Name : p.Name,
    SubPolicies : p.SubPolicies,
    TargetFileServerExtId : p.TargetFileServerExtId,
    TargetPath : p.TargetPath,
  })
}

func NewDataConsolidationPolicy() *DataConsolidationPolicy {
  p := new(DataConsolidationPolicy)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.DataConsolidationPolicy"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.DataConsolidationPolicy"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST response for all response codes in api path /files/v4.0.a2/config/file-server/data-consolidation-policies/{dataConsolidationPolicyExtId} Get operation
*/
type DataConsolidationPolicyApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfDataConsolidationPolicyApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewDataConsolidationPolicyApiResponse() *DataConsolidationPolicyApiResponse {
  p := new(DataConsolidationPolicyApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.DataConsolidationPolicyApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.DataConsolidationPolicyApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *DataConsolidationPolicyApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *DataConsolidationPolicyApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfDataConsolidationPolicyApiResponseData()
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
REST response for all response codes in api path /files/v4.0.a2/config/file-server/data-consolidation-policies Get operation
*/
type DataConsolidationPolicyListApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfDataConsolidationPolicyListApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewDataConsolidationPolicyListApiResponse() *DataConsolidationPolicyListApiResponse {
  p := new(DataConsolidationPolicyListApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.DataConsolidationPolicyListApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.DataConsolidationPolicyListApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *DataConsolidationPolicyListApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *DataConsolidationPolicyListApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfDataConsolidationPolicyListApiResponseData()
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
Data protection policy
*/
type DataProtectionPolicy struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Policy description
  */
  Description *string `json:"description,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  This flag indicates if this policy has been used for reverse replication.
  */
  IsReversePolicy *bool `json:"isReversePolicy,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import1.ApiLink `json:"links,omitempty"`
  /**
  Name of the policy
  */
  Name *string `json:"name"`
  
  PolicyStatus *PolicyStatus `json:"policyStatus,omitempty"`
  
  ReplicationSummary *ReplicationSummary `json:"replicationSummary,omitempty"`
  /**
  Data location represents combination of source, target, source target map, schedules, policy status and replication summary status.
  */
  SubPolicies []SubPolicy `json:"subPolicies,omitempty"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
  
  Type *DataProtectionPolicyType `json:"type"`
}

func (p *DataProtectionPolicy) MarshalJSON() ([]byte, error) {
  type DataProtectionPolicyProxy DataProtectionPolicy
  return json.Marshal(struct {
    *DataProtectionPolicyProxy
    Name *string `json:"name,omitempty"`
    Type *DataProtectionPolicyType `json:"type,omitempty"`
  }{
    DataProtectionPolicyProxy : (*DataProtectionPolicyProxy)(p),
    Name : p.Name,
    Type : p.Type,
  })
}

func NewDataProtectionPolicy() *DataProtectionPolicy {
  p := new(DataProtectionPolicy)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.DataProtectionPolicy"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.DataProtectionPolicy"}
  p.UnknownFields_ = map[string]interface{}{}

  p.IsReversePolicy = new(bool)
  *p.IsReversePolicy = false


  return p
}




/**
REST response for all response codes in api path /files/v4.0.a2/config/file-server/data-protection-policies/{dataProtectionPolicyExtId} Get operation
*/
type DataProtectionPolicyApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfDataProtectionPolicyApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewDataProtectionPolicyApiResponse() *DataProtectionPolicyApiResponse {
  p := new(DataProtectionPolicyApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.DataProtectionPolicyApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.DataProtectionPolicyApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *DataProtectionPolicyApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *DataProtectionPolicyApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfDataProtectionPolicyApiResponseData()
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
REST response for all response codes in api path /files/v4.0.a2/config/file-server/data-protection-policies Get operation
*/
type DataProtectionPolicyListApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfDataProtectionPolicyListApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewDataProtectionPolicyListApiResponse() *DataProtectionPolicyListApiResponse {
  p := new(DataProtectionPolicyListApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.DataProtectionPolicyListApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.DataProtectionPolicyListApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *DataProtectionPolicyListApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *DataProtectionPolicyListApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfDataProtectionPolicyListApiResponseData()
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
Policy type value. The acceptable values are "DATAPROTECTIONPOLICYTYPE_DR", "DATAPROTECTIONPOLICYTYPE_SYNC".
*/
type DataProtectionPolicyType int

const(
  DATAPROTECTIONPOLICYTYPE_UNKNOWN DataProtectionPolicyType = 0
  DATAPROTECTIONPOLICYTYPE_REDACTED DataProtectionPolicyType = 1
  DATAPROTECTIONPOLICYTYPE_DATAPROTECTIONPOLICYTYPE_DR DataProtectionPolicyType = 2
  DATAPROTECTIONPOLICYTYPE_DATAPROTECTIONPOLICYTYPE_SYNC DataProtectionPolicyType = 3
)

// returns the name of the enum given an ordinal number
func (e *DataProtectionPolicyType) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "DATAPROTECTIONPOLICYTYPE_DR",
    "DATAPROTECTIONPOLICYTYPE_SYNC",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *DataProtectionPolicyType) index(name string) DataProtectionPolicyType {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "DATAPROTECTIONPOLICYTYPE_DR",
    "DATAPROTECTIONPOLICYTYPE_SYNC",
  }
  for idx := range names {
    if names[idx] == name {
      return DataProtectionPolicyType(idx)
    }
  }
  return DATAPROTECTIONPOLICYTYPE_UNKNOWN
}

func (e *DataProtectionPolicyType) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for DataProtectionPolicyType:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *DataProtectionPolicyType) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e DataProtectionPolicyType) Ref() *DataProtectionPolicyType {
  return &e
}


/**
REST response for all response codes in api path /files/v4.0.a2/config/file-server/data-protection/$actions/deactivate Post operation
*/
type DeActivateFileServerApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfDeActivateFileServerApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewDeActivateFileServerApiResponse() *DeActivateFileServerApiResponse {
  p := new(DeActivateFileServerApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.DeActivateFileServerApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.DeActivateFileServerApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *DeActivateFileServerApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *DeActivateFileServerApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfDeActivateFileServerApiResponseData()
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
DeActivation involves, making the mount targets for the provided file server pair read-only and replicating the final delta, if needed.
*/
type DeactivateSpec struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import1.ApiLink `json:"links,omitempty"`
  /**
  The {extId} of the primary file server. This is a read-only field.
  */
  PrimaryFileServerExtId *string `json:"primaryFileServerExtId,omitempty"`
  /**
  Flag to indicate whether the final delta needs to be replicated or not.
  */
  ReplicateFinalDelta *bool `json:"replicateFinalDelta,omitempty"`
  /**
  The {extId} of the target file server. This is a read-only field.
  */
  TargetFileServerExtId *string `json:"targetFileServerExtId,omitempty"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
}


func NewDeactivateSpec() *DeactivateSpec {
  p := new(DeactivateSpec)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.DeactivateSpec"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.DeactivateSpec"}
  p.UnknownFields_ = map[string]interface{}{}

  p.ReplicateFinalDelta = new(bool)
  *p.ReplicateFinalDelta = false


  return p
}




/**
REST response for all response codes in api path /files/v4.0.a2/config/file-server/admin-users/{adminUserExtId} Delete operation
*/
type DeleteAdminUserApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfDeleteAdminUserApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewDeleteAdminUserApiResponse() *DeleteAdminUserApiResponse {
  p := new(DeleteAdminUserApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.DeleteAdminUserApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.DeleteAdminUserApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *DeleteAdminUserApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *DeleteAdminUserApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfDeleteAdminUserApiResponseData()
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
REST response for all response codes in api path /files/v4.0.a2/config/file-server/anti-virus-servers/{antivirusServerExtId} Delete operation
*/
type DeleteAntivirusServerApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfDeleteAntivirusServerApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewDeleteAntivirusServerApiResponse() *DeleteAntivirusServerApiResponse {
  p := new(DeleteAntivirusServerApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.DeleteAntivirusServerApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.DeleteAntivirusServerApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *DeleteAntivirusServerApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *DeleteAntivirusServerApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfDeleteAntivirusServerApiResponseData()
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
REST response for all response codes in api path /files/v4.0.a2/config/file-server/data-consolidation-policies/{dataConsolidationPolicyExtId} Delete operation
*/
type DeleteDataConsolidationPolicyApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfDeleteDataConsolidationPolicyApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewDeleteDataConsolidationPolicyApiResponse() *DeleteDataConsolidationPolicyApiResponse {
  p := new(DeleteDataConsolidationPolicyApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.DeleteDataConsolidationPolicyApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.DeleteDataConsolidationPolicyApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *DeleteDataConsolidationPolicyApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *DeleteDataConsolidationPolicyApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfDeleteDataConsolidationPolicyApiResponseData()
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
REST response for all response codes in api path /files/v4.0.a2/config/file-server/data-protection-policies/{dataProtectionPolicyExtId} Delete operation
*/
type DeleteDataProtectionPolicyApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfDeleteDataProtectionPolicyApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewDeleteDataProtectionPolicyApiResponse() *DeleteDataProtectionPolicyApiResponse {
  p := new(DeleteDataProtectionPolicyApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.DeleteDataProtectionPolicyApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.DeleteDataProtectionPolicyApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *DeleteDataProtectionPolicyApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *DeleteDataProtectionPolicyApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfDeleteDataProtectionPolicyApiResponseData()
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
REST response for all response codes in api path /files/v4.0.a2/config/file-server/users/{fileServerUserExtId} Delete operation
*/
type DeleteFileServerUserApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfDeleteFileServerUserApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewDeleteFileServerUserApiResponse() *DeleteFileServerUserApiResponse {
  p := new(DeleteFileServerUserApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.DeleteFileServerUserApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.DeleteFileServerUserApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *DeleteFileServerUserApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *DeleteFileServerUserApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfDeleteFileServerUserApiResponseData()
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
REST response for all response codes in api path /files/v4.0.a2/config/file-server/mount-targets/{mountTargetExtId} Delete operation
*/
type DeleteMountTargetApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfDeleteMountTargetApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewDeleteMountTargetApiResponse() *DeleteMountTargetApiResponse {
  p := new(DeleteMountTargetApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.DeleteMountTargetApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.DeleteMountTargetApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *DeleteMountTargetApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *DeleteMountTargetApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfDeleteMountTargetApiResponseData()
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
REST response for all response codes in api path /files/v4.0.a2/config/file-server/mount-targets/{mountTargetExtId}/snapshots/{snapshotExtId} Delete operation
*/
type DeleteMountTargetSnapshotApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfDeleteMountTargetSnapshotApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewDeleteMountTargetSnapshotApiResponse() *DeleteMountTargetSnapshotApiResponse {
  p := new(DeleteMountTargetSnapshotApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.DeleteMountTargetSnapshotApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.DeleteMountTargetSnapshotApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *DeleteMountTargetSnapshotApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *DeleteMountTargetSnapshotApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfDeleteMountTargetSnapshotApiResponseData()
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
REST response for all response codes in api path /files/v4.0.a2/config/file-server/mount-targets/{mountTargetExtId}/virus-scan-policy/{virusScanPolicyExtId} Delete operation
*/
type DeleteMountTargetVirusScanPolicyApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfDeleteMountTargetVirusScanPolicyApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewDeleteMountTargetVirusScanPolicyApiResponse() *DeleteMountTargetVirusScanPolicyApiResponse {
  p := new(DeleteMountTargetVirusScanPolicyApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.DeleteMountTargetVirusScanPolicyApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.DeleteMountTargetVirusScanPolicyApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *DeleteMountTargetVirusScanPolicyApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *DeleteMountTargetVirusScanPolicyApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfDeleteMountTargetVirusScanPolicyApiResponseData()
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
REST response for all response codes in api path /files/v4.0.a2/config/file-server/networks/{extId} Delete operation
*/
type DeleteNetworkApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfDeleteNetworkApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewDeleteNetworkApiResponse() *DeleteNetworkApiResponse {
  p := new(DeleteNetworkApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.DeleteNetworkApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.DeleteNetworkApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *DeleteNetworkApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *DeleteNetworkApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfDeleteNetworkApiResponseData()
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
REST response for all response codes in api path /files/v4.0.a2/config/file-server/mount-targets/{mountTargetExtId}/quota-policies/{quotaPolicyExtId} Delete operation
*/
type DeleteQuotaPolicyApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfDeleteQuotaPolicyApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewDeleteQuotaPolicyApiResponse() *DeleteQuotaPolicyApiResponse {
  p := new(DeleteQuotaPolicyApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.DeleteQuotaPolicyApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.DeleteQuotaPolicyApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *DeleteQuotaPolicyApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *DeleteQuotaPolicyApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfDeleteQuotaPolicyApiResponseData()
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
REST response for all response codes in api path /files/v4.0.a2/config/file-server/ssr-snapshot-schedules/{ssrSnapshotScheduleExtId} Delete operation
*/
type DeleteSsrSnapshotScheduleApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfDeleteSsrSnapshotScheduleApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewDeleteSsrSnapshotScheduleApiResponse() *DeleteSsrSnapshotScheduleApiResponse {
  p := new(DeleteSsrSnapshotScheduleApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.DeleteSsrSnapshotScheduleApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.DeleteSsrSnapshotScheduleApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *DeleteSsrSnapshotScheduleApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *DeleteSsrSnapshotScheduleApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfDeleteSsrSnapshotScheduleApiResponseData()
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
REST response for all response codes in api path /files/v4.0.a2/config/file-server/tier-config/{extId} Delete operation
*/
type DeleteTieringConfigApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfDeleteTieringConfigApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewDeleteTieringConfigApiResponse() *DeleteTieringConfigApiResponse {
  p := new(DeleteTieringConfigApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.DeleteTieringConfigApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.DeleteTieringConfigApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *DeleteTieringConfigApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *DeleteTieringConfigApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfDeleteTieringConfigApiResponseData()
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
REST response for all response codes in api path /files/v4.0.a2/config/file-server/tier-policies/{extId} Delete operation
*/
type DeleteTieringPolicyApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfDeleteTieringPolicyApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewDeleteTieringPolicyApiResponse() *DeleteTieringPolicyApiResponse {
  p := new(DeleteTieringPolicyApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.DeleteTieringPolicyApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.DeleteTieringPolicyApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *DeleteTieringPolicyApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *DeleteTieringPolicyApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfDeleteTieringPolicyApiResponseData()
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
REST response for all response codes in api path /files/v4.0.a2/config/file-server/data-consolidation-policies/{dataConsolidationPolicyExtId}/$actions/disable Post operation
*/
type DisableDataConsolidationPolicyApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfDisableDataConsolidationPolicyApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewDisableDataConsolidationPolicyApiResponse() *DisableDataConsolidationPolicyApiResponse {
  p := new(DisableDataConsolidationPolicyApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.DisableDataConsolidationPolicyApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.DisableDataConsolidationPolicyApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *DisableDataConsolidationPolicyApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *DisableDataConsolidationPolicyApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfDisableDataConsolidationPolicyApiResponseData()
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
REST response for all response codes in api path /files/v4.0.a2/config/file-server/data-protection-policies/{dataProtectionPolicyExtId}/$actions/disable Post operation
*/
type DisableDataProtectionPolicyApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfDisableDataProtectionPolicyApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewDisableDataProtectionPolicyApiResponse() *DisableDataProtectionPolicyApiResponse {
  p := new(DisableDataProtectionPolicyApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.DisableDataProtectionPolicyApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.DisableDataProtectionPolicyApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *DisableDataProtectionPolicyApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *DisableDataProtectionPolicyApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfDisableDataProtectionPolicyApiResponseData()
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
Disables the policy with cancelOnGoingJobs as user given input, if the input is not provided then the default value is set to false.
*/
type DisablePolicySpec struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  This parameter indicates whether all the ongoing jobs needed to be
cancelled during the disable policy operation. It is false by default.
  */
  CancelOngoingJobs *bool `json:"cancelOngoingJobs,omitempty"`
  /**
  A code that uniquely identifies a message.
  */
  Code *string `json:"code,omitempty"`
  /**
  The locale for the message description.
  */
  Locale *string `json:"locale,omitempty"`
  
  Message *string `json:"message,omitempty"`
  
  Severity *import3.MessageSeverity `json:"severity,omitempty"`
}


func NewDisablePolicySpec() *DisablePolicySpec {
  p := new(DisablePolicySpec)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.DisablePolicySpec"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.DisablePolicySpec"}
  p.UnknownFields_ = map[string]interface{}{}

  p.CancelOngoingJobs = new(bool)
  *p.CancelOngoingJobs = false
  p.Locale = new(string)
  *p.Locale = "en_US"


  return p
}




/**
DNS detail for the file server.
*/
type Dns struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  DNS server UUID. This is a read-only field. Example:28f78959-14a6-4c47-b5db-920460c4b668.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  Host name for the DNS server. This is a read-only field.
  */
  HostName *string `json:"hostName,omitempty"`
  /**
  IPV4 address of the DNS server. This is a read-only field.
  */
  IpAddress *string `json:"ipAddress,omitempty"`
  /**
  This field indicates if the DNS server is verified. This is a read-only field.
  */
  IsVerified *bool `json:"isVerified,omitempty"`
  
  OperationType *DnsOperationType `json:"operationType,omitempty"`
  /**
  Ad domain password associated with the file server.
  */
  Password *string `json:"password"`
  /**
  Preferred name server for a file server.
  */
  PreferredNameServer *string `json:"preferredNameServer,omitempty"`
  /**
  Ad domain username associated with the file server.
  */
  Username *string `json:"username"`
}

func (p *Dns) MarshalJSON() ([]byte, error) {
  type DnsProxy Dns
  return json.Marshal(struct {
    *DnsProxy
    Password *string `json:"password,omitempty"`
    Username *string `json:"username,omitempty"`
  }{
    DnsProxy : (*DnsProxy)(p),
    Password : p.Password,
    Username : p.Username,
  })
}

func NewDns() *Dns {
  p := new(Dns)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.Dns"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.Dns"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
DNS detail for the file server.
*/
type DnsDetails struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  Host name for the DNS server. This is a read-only field.
  */
  HostName *string `json:"hostName,omitempty"`
  /**
  IPV4 address of the DNS server. This is a read-only field.
  */
  IpAddress *string `json:"ipAddress,omitempty"`
  /**
  This field indicates if the DNS server is verified. This is a read-only field.
  */
  IsVerified *bool `json:"isVerified,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import1.ApiLink `json:"links,omitempty"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
}


func NewDnsDetails() *DnsDetails {
  p := new(DnsDetails)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.DnsDetails"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.DnsDetails"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST response for all response codes in api path /files/v4.0.a2/config/file-server/dns-entries Get operation
*/
type DnsEntriesListApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfDnsEntriesListApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewDnsEntriesListApiResponse() *DnsEntriesListApiResponse {
  p := new(DnsEntriesListApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.DnsEntriesListApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.DnsEntriesListApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *DnsEntriesListApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *DnsEntriesListApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfDnsEntriesListApiResponseData()
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
DNS operation type. The acceptable values are "MANUAL", "MS_DNS", "DNSOPERATIONTYPE_MANUAL", "DNSOPERATIONTYPE_MS_DNS".
*/
type DnsOperationType int

const(
  DNSOPERATIONTYPE_UNKNOWN DnsOperationType = 0
  DNSOPERATIONTYPE_REDACTED DnsOperationType = 1
  DNSOPERATIONTYPE_MANUAL DnsOperationType = 2
  DNSOPERATIONTYPE_MS_DNS DnsOperationType = 3
  DNSOPERATIONTYPE_DNSOPERATIONTYPE_MANUAL DnsOperationType = 4
  DNSOPERATIONTYPE_DNSOPERATIONTYPE_MS_DNS DnsOperationType = 5
)

// returns the name of the enum given an ordinal number
func (e *DnsOperationType) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "MANUAL",
    "MS_DNS",
    "DNSOPERATIONTYPE_MANUAL",
    "DNSOPERATIONTYPE_MS_DNS",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *DnsOperationType) index(name string) DnsOperationType {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "MANUAL",
    "MS_DNS",
    "DNSOPERATIONTYPE_MANUAL",
    "DNSOPERATIONTYPE_MS_DNS",
  }
  for idx := range names {
    if names[idx] == name {
      return DnsOperationType(idx)
    }
  }
  return DNSOPERATIONTYPE_UNKNOWN
}

func (e *DnsOperationType) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for DnsOperationType:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *DnsOperationType) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e DnsOperationType) Ref() *DnsOperationType {
  return &e
}


/**
REST response for all response codes in api path /files/v4.0.a2/config/file-server/data-consolidation-policies/{dataConsolidationPolicyExtId}/$actions/enable Post operation
*/
type EnableDataConsolidationPolicyApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfEnableDataConsolidationPolicyApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewEnableDataConsolidationPolicyApiResponse() *EnableDataConsolidationPolicyApiResponse {
  p := new(EnableDataConsolidationPolicyApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.EnableDataConsolidationPolicyApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.EnableDataConsolidationPolicyApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *EnableDataConsolidationPolicyApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *EnableDataConsolidationPolicyApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfEnableDataConsolidationPolicyApiResponseData()
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
REST response for all response codes in api path /files/v4.0.a2/config/file-server/data-protection-policies/{dataProtectionPolicyExtId}/$actions/enable Post operation
*/
type EnableDataProtectionPolicyApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfEnableDataProtectionPolicyApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewEnableDataProtectionPolicyApiResponse() *EnableDataProtectionPolicyApiResponse {
  p := new(EnableDataProtectionPolicyApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.EnableDataProtectionPolicyApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.EnableDataProtectionPolicyApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *EnableDataProtectionPolicyApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *EnableDataProtectionPolicyApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfEnableDataProtectionPolicyApiResponseData()
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
Configuration to enable current and future mount targets.The acceptable values are "ALL_CURRENT_MOUNT_TARGETS", "ALL_CURRENT_FUTURE_MOUNT_TARGETS",  "ALL_FUTURE_MOUNT_TARGETS", "NONE".
*/
type EnableMountTargets int

const(
  ENABLEMOUNTTARGETS_UNKNOWN EnableMountTargets = 0
  ENABLEMOUNTTARGETS_REDACTED EnableMountTargets = 1
  ENABLEMOUNTTARGETS_ALL_CURRENT_MOUNT_TARGETS EnableMountTargets = 2
  ENABLEMOUNTTARGETS_ALL_CURRENT_FUTURE_MOUNT_TARGETS EnableMountTargets = 3
  ENABLEMOUNTTARGETS_ALL_FUTURE_MOUNT_TARGETS EnableMountTargets = 4
  ENABLEMOUNTTARGETS_NONE EnableMountTargets = 5
)

// returns the name of the enum given an ordinal number
func (e *EnableMountTargets) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "ALL_CURRENT_MOUNT_TARGETS",
    "ALL_CURRENT_FUTURE_MOUNT_TARGETS",
    "ALL_FUTURE_MOUNT_TARGETS",
    "NONE",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *EnableMountTargets) index(name string) EnableMountTargets {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "ALL_CURRENT_MOUNT_TARGETS",
    "ALL_CURRENT_FUTURE_MOUNT_TARGETS",
    "ALL_FUTURE_MOUNT_TARGETS",
    "NONE",
  }
  for idx := range names {
    if names[idx] == name {
      return EnableMountTargets(idx)
    }
  }
  return ENABLEMOUNTTARGETS_UNKNOWN
}

func (e *EnableMountTargets) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for EnableMountTargets:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *EnableMountTargets) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e EnableMountTargets) Ref() *EnableMountTargets {
  return &e
}


/**
Enforcement types for the quota policy. The acceptable values are "soft" or "hard".
*/
type EnforcementType int

const(
  ENFORCEMENTTYPE_UNKNOWN EnforcementType = 0
  ENFORCEMENTTYPE_REDACTED EnforcementType = 1
  ENFORCEMENTTYPE_SOFT EnforcementType = 2
  ENFORCEMENTTYPE_HARD EnforcementType = 3
)

// returns the name of the enum given an ordinal number
func (e *EnforcementType) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "SOFT",
    "HARD",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *EnforcementType) index(name string) EnforcementType {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "SOFT",
    "HARD",
  }
  for idx := range names {
    if names[idx] == name {
      return EnforcementType(idx)
    }
  }
  return ENFORCEMENTTYPE_UNKNOWN
}

func (e *EnforcementType) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for EnforcementType:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *EnforcementType) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e EnforcementType) Ref() *EnforcementType {
  return &e
}


/**
Explicit mapping model.
*/
type ExplicitIdentityMapping struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  DeniedAccessNfsGroups []NfsIdentitiesMapping `json:"deniedAccessNfsGroups,omitempty"`
  
  DeniedAccessNfsUsers []NfsIdentitiesMapping `json:"deniedAccessNfsUsers,omitempty"`
  
  DeniedAccessSmbGroups []SmbIdentitiesMapping `json:"deniedAccessSmbGroups,omitempty"`
  
  DeniedAccessSmbUsers []SmbIdentitiesMapping `json:"deniedAccessSmbUsers,omitempty"`
  
  OneToOneMappings []OneToOneMapping `json:"oneToOneMappings,omitempty"`
  
  WildcardMappings []WildCardMapping `json:"wildcardMappings,omitempty"`
}


func NewExplicitIdentityMapping() *ExplicitIdentityMapping {
  p := new(ExplicitIdentityMapping)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.ExplicitIdentityMapping"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.ExplicitIdentityMapping"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Details of the user actions to be performed. The acceptable values are "FAILOVERACTION_PLANNED_FAILOVER", "FAILOVERACTION_UNPLANNED_FAILOVER", "FAILOVERACTION_FAILBACK", "FAILOVERACTION_RESUME_REPLICATION", "FAILOVERACTION_RESUME_REPLICATION_FROM_PRIMARY", "FAILOVERACTION_RESUME_REPLICATION_FROM_STANDBY", "FAILOVERACTION_CONFIGURE_AD_DNS"
*/
type FailOverAction int

const(
  FAILOVERACTION_UNKNOWN FailOverAction = 0
  FAILOVERACTION_REDACTED FailOverAction = 1
  FAILOVERACTION_FAILOVERACTION_PLANNED_FAILOVER FailOverAction = 2
  FAILOVERACTION_FAILOVERACTION_UNPLANNED_FAILOVER FailOverAction = 3
  FAILOVERACTION_FAILOVERACTION_FAILBACK FailOverAction = 4
  FAILOVERACTION_FAILOVERACTION_RESUME_REPLICATION FailOverAction = 5
  FAILOVERACTION_FAILOVERACTION_RESUME_REPLICATION_FROM_PRIMARY FailOverAction = 6
  FAILOVERACTION_FAILOVERACTION_RESUME_REPLICATION_FROM_STANDBY FailOverAction = 7
  FAILOVERACTION_FAILOVERACTION_CONFIGURE_AD_DNS FailOverAction = 8
)

// returns the name of the enum given an ordinal number
func (e *FailOverAction) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "FAILOVERACTION_PLANNED_FAILOVER",
    "FAILOVERACTION_UNPLANNED_FAILOVER",
    "FAILOVERACTION_FAILBACK",
    "FAILOVERACTION_RESUME_REPLICATION",
    "FAILOVERACTION_RESUME_REPLICATION_FROM_PRIMARY",
    "FAILOVERACTION_RESUME_REPLICATION_FROM_STANDBY",
    "FAILOVERACTION_CONFIGURE_AD_DNS",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *FailOverAction) index(name string) FailOverAction {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "FAILOVERACTION_PLANNED_FAILOVER",
    "FAILOVERACTION_UNPLANNED_FAILOVER",
    "FAILOVERACTION_FAILBACK",
    "FAILOVERACTION_RESUME_REPLICATION",
    "FAILOVERACTION_RESUME_REPLICATION_FROM_PRIMARY",
    "FAILOVERACTION_RESUME_REPLICATION_FROM_STANDBY",
    "FAILOVERACTION_CONFIGURE_AD_DNS",
  }
  for idx := range names {
    if names[idx] == name {
      return FailOverAction(idx)
    }
  }
  return FAILOVERACTION_UNKNOWN
}

func (e *FailOverAction) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for FailOverAction:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *FailOverAction) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e FailOverAction) Ref() *FailOverAction {
  return &e
}


/**
Failover status. The acceptable values are "FAILOVERSTATUS_ERROR", "FAILOVERSTATUS_FAILOVER_IN_PROGRESS", "FAILOVERSTATUS_FAILBACK_IN_PROGRESS", "FAILOVERSTATUS_RESUME_REPLICATION_IN_PROGRESS"
*/
type FailOverStatus int

const(
  FAILOVERSTATUS_UNKNOWN FailOverStatus = 0
  FAILOVERSTATUS_REDACTED FailOverStatus = 1
  FAILOVERSTATUS_FAILOVERSTATUS_ERROR FailOverStatus = 2
  FAILOVERSTATUS_FAILOVERSTATUS_FAILOVER_IN_PROGRESS FailOverStatus = 3
  FAILOVERSTATUS_FAILOVERSTATUS_FAILBACK_IN_PROGRESS FailOverStatus = 4
  FAILOVERSTATUS_FAILOVERSTATUS_RESUME_REPLICATION_IN_PROGRESS FailOverStatus = 5
)

// returns the name of the enum given an ordinal number
func (e *FailOverStatus) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "FAILOVERSTATUS_ERROR",
    "FAILOVERSTATUS_FAILOVER_IN_PROGRESS",
    "FAILOVERSTATUS_FAILBACK_IN_PROGRESS",
    "FAILOVERSTATUS_RESUME_REPLICATION_IN_PROGRESS",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *FailOverStatus) index(name string) FailOverStatus {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "FAILOVERSTATUS_ERROR",
    "FAILOVERSTATUS_FAILOVER_IN_PROGRESS",
    "FAILOVERSTATUS_FAILBACK_IN_PROGRESS",
    "FAILOVERSTATUS_RESUME_REPLICATION_IN_PROGRESS",
  }
  for idx := range names {
    if names[idx] == name {
      return FailOverStatus(idx)
    }
  }
  return FAILOVERSTATUS_UNKNOWN
}

func (e *FailOverStatus) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for FailOverStatus:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *FailOverStatus) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e FailOverStatus) Ref() *FailOverStatus {
  return &e
}


/**
Fail over type .The acceptable values are "FAILOVERTYPE_PLANNED_FAILOVER", "FAILOVERTYPE_UNPLANNED_FAILOVER", "FAILOVERTYPE_FAILBACK", "FAILOVERTYPE_RESUME_REPLICATION", "FAILOVERTYPE_RESUME_REPLICATION_FROM_PRIMARY", "FAILOVERTYPE_RESUME_REPLICATION_FROM_STANDBY".
*/
type FailOverType int

const(
  FAILOVERTYPE_UNKNOWN FailOverType = 0
  FAILOVERTYPE_REDACTED FailOverType = 1
  FAILOVERTYPE_FAILOVERTYPE_PLANNED_FAILOVER FailOverType = 2
  FAILOVERTYPE_FAILOVERTYPE_UNPLANNED_FAILOVER FailOverType = 3
  FAILOVERTYPE_FAILOVERTYPE_FAILBACK FailOverType = 4
  FAILOVERTYPE_FAILOVERTYPE_RESUME_REPLICATION FailOverType = 5
  FAILOVERTYPE_FAILOVERTYPE_RESUME_REPLICATION_FROM_PRIMARY FailOverType = 6
  FAILOVERTYPE_FAILOVERTYPE_RESUME_REPLICATION_FROM_STANDBY FailOverType = 7
)

// returns the name of the enum given an ordinal number
func (e *FailOverType) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "FAILOVERTYPE_PLANNED_FAILOVER",
    "FAILOVERTYPE_UNPLANNED_FAILOVER",
    "FAILOVERTYPE_FAILBACK",
    "FAILOVERTYPE_RESUME_REPLICATION",
    "FAILOVERTYPE_RESUME_REPLICATION_FROM_PRIMARY",
    "FAILOVERTYPE_RESUME_REPLICATION_FROM_STANDBY",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *FailOverType) index(name string) FailOverType {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "FAILOVERTYPE_PLANNED_FAILOVER",
    "FAILOVERTYPE_UNPLANNED_FAILOVER",
    "FAILOVERTYPE_FAILBACK",
    "FAILOVERTYPE_RESUME_REPLICATION",
    "FAILOVERTYPE_RESUME_REPLICATION_FROM_PRIMARY",
    "FAILOVERTYPE_RESUME_REPLICATION_FROM_STANDBY",
  }
  for idx := range names {
    if names[idx] == name {
      return FailOverType(idx)
    }
  }
  return FAILOVERTYPE_UNKNOWN
}

func (e *FailOverType) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for FailOverType:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *FailOverType) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e FailOverType) Ref() *FailOverType {
  return &e
}


/**
Failed file details of migration plan
*/
type FailedFile struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Migration of file failed with error code.
  */
  ErrorCode *string `json:"errorCode,omitempty"`
  /**
  Migration of file failed with an error message.
  */
  ErrorMsg *string `json:"errorMsg,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  
  FilesList []string `json:"filesList,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import1.ApiLink `json:"links,omitempty"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
}


func NewFailedFile() *FailedFile {
  p := new(FailedFile)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.FailedFile"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.FailedFile"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Failed files list of migration plan
*/
type FailedFilesList struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  
  FailedFiles []FailedFile `json:"failedFiles"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import1.ApiLink `json:"links,omitempty"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
}

func (p *FailedFilesList) MarshalJSON() ([]byte, error) {
  type FailedFilesListProxy FailedFilesList
  return json.Marshal(struct {
    *FailedFilesListProxy
    FailedFiles []FailedFile `json:"failedFiles,omitempty"`
  }{
    FailedFilesListProxy : (*FailedFilesListProxy)(p),
    FailedFiles : p.FailedFiles,
  })
}

func NewFailedFilesList() *FailedFilesList {
  p := new(FailedFilesList)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.FailedFilesList"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.FailedFilesList"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST response for all response codes in api path /files/v4.0.a2/config/file-server/data-protection/$actions/failover Post operation
*/
type FailoverApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfFailoverApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewFailoverApiResponse() *FailoverApiResponse {
  p := new(FailoverApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.FailoverApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.FailoverApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *FailoverApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *FailoverApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfFailoverApiResponseData()
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
Failover model.
*/
type FailoverSpec struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  
  FileServerState *FileServerState `json:"fileServerState,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import1.ApiLink `json:"links,omitempty"`
  /**
  The {extId} of the owning file server. This is a read-only field
  */
  OwningFileServerExtId *string `json:"owningFileServerExtId,omitempty"`
  /**
  The {extId} of the data protection policy.
  */
  PolicyExtIds []string `json:"policyExtIds,omitempty"`
  
  PrimaryAdCredential *Credential `json:"primaryAdCredential,omitempty"`
  /**
  Use a specific domain controller for the join-domain operation in a multi DC active directory setup. By default, AFS discovers a site local domain controller for join-domain operation. The preferred domain controller cannot be an IP address. It has to be FQDN (example: dc_name.dom.companyname.com)
  */
  PrimaryAdPreferredDomainController *string `json:"primaryAdPreferredDomainController,omitempty"`
  
  PrimaryDns *Dns `json:"primaryDns,omitempty"`
  /**
  The {extId} of the primary file server. This is a read-only field.
  */
  PrimaryFileServerExtId *string `json:"primaryFileServerExtId,omitempty"`
  
  ReverseReplicationPolicy *DataProtectionPolicy `json:"reverseReplicationPolicy,omitempty"`
  
  Status *ActionStatus `json:"status,omitempty"`
  /**
  The {extId} of the target file server. This is a read-only field.
  */
  TargetFileServerExtId *string `json:"targetFileServerExtId,omitempty"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
  
  Type *FailOverType `json:"type,omitempty"`
}


func NewFailoverSpec() *FailoverSpec {
  p := new(FailoverSpec)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.FailoverSpec"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.FailoverSpec"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Information related to a single file sent and received back.
*/
type FileInfoSpec struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Error while retrieving the mimetype of a file.
  */
  ErrorMsg *string `json:"errorMsg,omitempty"`
  /**
  genId of a file.
  */
  GenId *int64 `json:"genId,omitempty"`
  /**
  Inode number of a file.
  */
  Inode *int64 `json:"inode,omitempty"`
  /**
  MimeType of a file.
  */
  MimeType *string `json:"mimeType,omitempty"`
}


func NewFileInfoSpec() *FileInfoSpec {
  p := new(FileInfoSpec)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.FileInfoSpec"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.FileInfoSpec"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
File server object
*/
type FileServer struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  The cluster associated with the file server. This contains AOS cluster UUID. This is a read-only field.
  */
  ClusterExtId *string `json:"clusterExtId,omitempty"`
  /**
  Filesystem compression value. For file servers greater than 3.6 version, the filesystem compression is enabled by default. This is a read-only field.
  */
  CompressionEnabled *bool `json:"compressionEnabled,omitempty"`
  /**
  The container associated with the file server. This is a read-only field.
  */
  ContainerExtId *string `json:"containerExtId,omitempty"`
  /**
  List of CVM IP addresses associated with the file server. This is a read-only field.
  */
  CvmIpAddresses []import3.IPv4Address `json:"cvmIpAddresses,omitempty"`
  /**
  Fully qualified domain name (file server namespace). This, along with the file server name, constitutes the namespace of the file server. Example: fileserver_name.corp.companyname.com. This is also used to create file server DNS entries on the nameservers so that clients can access the file server using its name.
  */
  DnsDomainName *string `json:"dnsDomainName,omitempty"`
  /**
  List of DNS servers associated with the file server.
  */
  DnsServers []import3.IPv4Address `json:"dnsServers,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  List of external networks associated with the file server.
  */
  ExternalNetworks []Network `json:"externalNetworks,omitempty"`
  /**
  Network Id of external primary network.
  */
  ExternalPrimaryNetworkExtId *string `json:"externalPrimaryNetworkExtId,omitempty"`
  /**
  List of file blocking extensions or patterns in a comma separated list of strings. For Ex: [".db",".txt",".mp3"].
  */
  FileBlockingExtensions []string `json:"fileBlockingExtensions,omitempty"`
  /**
  List of internal networks associated with the file server.
  */
  InternalNetworks []Network `json:"internalNetworks,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import1.ApiLink `json:"links,omitempty"`
  /**
  Memory associated with the file server VM.
  */
  MemoryGib *int64 `json:"memoryGib,omitempty"`
  /**
  File server name.
  */
  Name *string `json:"name"`
  /**
  The list of IP or Expand FQDN of the NTP servers associated with the file server.
  */
  NtpServers []import3.IPAddressOrFQDN `json:"ntpServers,omitempty"`
  /**
  Total number of file server VMs
  */
  NvmsCount *int `json:"nvmsCount,omitempty"`
  /**
  This flag indicates if file server recommendations are available and the user has to perform actions. This is a read-only field.
  */
  RebalanceNeeded *bool `json:"rebalanceNeeded,omitempty"`
  /**
  File server size.
  */
  SizeInGib *float64 `json:"sizeInGib,omitempty"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
  /**
  CPU associated with the file server VM.
  */
  Vcpus *int64 `json:"vcpus,omitempty"`
  /**
  File server version. This is a read-only field.
  */
  Version *string `json:"version,omitempty"`
  /**
  List of file server VMs associated with the file server. This is a read-only field.
  */
  Vms []VM `json:"vms,omitempty"`
}

func (p *FileServer) MarshalJSON() ([]byte, error) {
  type FileServerProxy FileServer
  return json.Marshal(struct {
    *FileServerProxy
    Name *string `json:"name,omitempty"`
  }{
    FileServerProxy : (*FileServerProxy)(p),
    Name : p.Name,
  })
}

func NewFileServer() *FileServer {
  p := new(FileServer)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.FileServer"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.FileServer"}
  p.UnknownFields_ = map[string]interface{}{}

  p.MemoryGib = new(int64)
  *p.MemoryGib = 12
  p.Vcpus = new(int64)
  *p.Vcpus = 4


  return p
}




/**
REST response for all response codes in api path /files/v4.0.a2/config/file-server Get operation
*/
type FileServerGetApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfFileServerGetApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewFileServerGetApiResponse() *FileServerGetApiResponse {
  p := new(FileServerGetApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.FileServerGetApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.FileServerGetApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *FileServerGetApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *FileServerGetApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfFileServerGetApiResponseData()
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
Defines the state of the file server. The acceptable values are "FILESERVERSTATE_ACTIVE" or "FILESERVERSTATE_STANDBY".
*/
type FileServerState int

const(
  FILESERVERSTATE_UNKNOWN FileServerState = 0
  FILESERVERSTATE_REDACTED FileServerState = 1
  FILESERVERSTATE_FILESERVERSTATE_ACTIVE FileServerState = 2
  FILESERVERSTATE_FILESERVERSTATE_STANDBY FileServerState = 3
)

// returns the name of the enum given an ordinal number
func (e *FileServerState) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "FILESERVERSTATE_ACTIVE",
    "FILESERVERSTATE_STANDBY",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *FileServerState) index(name string) FileServerState {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "FILESERVERSTATE_ACTIVE",
    "FILESERVERSTATE_STANDBY",
  }
  for idx := range names {
    if names[idx] == name {
      return FileServerState(idx)
    }
  }
  return FILESERVERSTATE_UNKNOWN
}

func (e *FileServerState) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for FileServerState:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *FileServerState) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e FileServerState) Ref() *FileServerState {
  return &e
}


/**
File server user model.
*/
type FileServerUser struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import1.ApiLink `json:"links,omitempty"`
  /**
  File server user name.
  */
  Name *string `json:"name"`
  /**
  Password of the file server user.
  */
  Password *string `json:"password"`
  /**
  List of file server user roles.
  */
  Roles []FileServerUserRole `json:"roles,omitempty"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
}

func (p *FileServerUser) MarshalJSON() ([]byte, error) {
  type FileServerUserProxy FileServerUser
  return json.Marshal(struct {
    *FileServerUserProxy
    Name *string `json:"name,omitempty"`
    Password *string `json:"password,omitempty"`
  }{
    FileServerUserProxy : (*FileServerUserProxy)(p),
    Name : p.Name,
    Password : p.Password,
  })
}

func NewFileServerUser() *FileServerUser {
  p := new(FileServerUser)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.FileServerUser"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.FileServerUser"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST response for all response codes in api path /files/v4.0.a2/config/file-server/users/{fileServerUserExtId} Get operation
*/
type FileServerUserApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfFileServerUserApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewFileServerUserApiResponse() *FileServerUserApiResponse {
  p := new(FileServerUserApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.FileServerUserApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.FileServerUserApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *FileServerUserApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *FileServerUserApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfFileServerUserApiResponseData()
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
REST response for all response codes in api path /files/v4.0.a2/config/file-server/users Get operation
*/
type FileServerUserListApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfFileServerUserListApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewFileServerUserListApiResponse() *FileServerUserListApiResponse {
  p := new(FileServerUserListApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.FileServerUserListApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.FileServerUserListApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *FileServerUserListApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *FileServerUserListApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfFileServerUserListApiResponseData()
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
File server user role. The acceptable values are "ROLE_CLUSTER_ADMIN", "ROLE_CLUSTER_VIEWER" or "ROLE_USER_ADMIN".
*/
type FileServerUserRole int

const(
  FILESERVERUSERROLE_UNKNOWN FileServerUserRole = 0
  FILESERVERUSERROLE_REDACTED FileServerUserRole = 1
  FILESERVERUSERROLE_ROLE_CLUSTER_ADMIN FileServerUserRole = 2
  FILESERVERUSERROLE_ROLE_CLUSTER_VIEWER FileServerUserRole = 3
  FILESERVERUSERROLE_ROLE_USER_ADMIN FileServerUserRole = 4
  FILESERVERUSERROLE_ROLE_MULTICLUSTER_ADMIN FileServerUserRole = 5
)

// returns the name of the enum given an ordinal number
func (e *FileServerUserRole) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "ROLE_CLUSTER_ADMIN",
    "ROLE_CLUSTER_VIEWER",
    "ROLE_USER_ADMIN",
    "ROLE_MULTICLUSTER_ADMIN",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *FileServerUserRole) index(name string) FileServerUserRole {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "ROLE_CLUSTER_ADMIN",
    "ROLE_CLUSTER_VIEWER",
    "ROLE_USER_ADMIN",
    "ROLE_MULTICLUSTER_ADMIN",
  }
  for idx := range names {
    if names[idx] == name {
      return FileServerUserRole(idx)
    }
  }
  return FILESERVERUSERROLE_UNKNOWN
}

func (e *FileServerUserRole) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for FileServerUserRole:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *FileServerUserRole) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e FileServerUserRole) Ref() *FileServerUserRole {
  return &e
}


/**
Authentication type of the mount target. The acceptable values are "KERBEROS5", "NONE", "SYSTEM", "KERBEROS5P", "KERBEROS5I".
*/
type FilesAuthenticationType int

const(
  FILESAUTHENTICATIONTYPE_UNKNOWN FilesAuthenticationType = 0
  FILESAUTHENTICATIONTYPE_REDACTED FilesAuthenticationType = 1
  FILESAUTHENTICATIONTYPE_KERBEROS5 FilesAuthenticationType = 2
  FILESAUTHENTICATIONTYPE_NONE FilesAuthenticationType = 3
  FILESAUTHENTICATIONTYPE_SYSTEM FilesAuthenticationType = 4
  FILESAUTHENTICATIONTYPE_KERBEROS5P FilesAuthenticationType = 5
  FILESAUTHENTICATIONTYPE_KERBEROS5I FilesAuthenticationType = 6
)

// returns the name of the enum given an ordinal number
func (e *FilesAuthenticationType) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "KERBEROS5",
    "NONE",
    "SYSTEM",
    "KERBEROS5P",
    "KERBEROS5I",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *FilesAuthenticationType) index(name string) FilesAuthenticationType {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "KERBEROS5",
    "NONE",
    "SYSTEM",
    "KERBEROS5P",
    "KERBEROS5I",
  }
  for idx := range names {
    if names[idx] == name {
      return FilesAuthenticationType(idx)
    }
  }
  return FILESAUTHENTICATIONTYPE_UNKNOWN
}

func (e *FilesAuthenticationType) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for FilesAuthenticationType:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *FilesAuthenticationType) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e FilesAuthenticationType) Ref() *FilesAuthenticationType {
  return &e
}


/**
REST response for all response codes in api path /files/v4.0.a2/config/file-server/object-store/profiles/{profileExtId} Get operation
*/
type GetObjectStoreApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfGetObjectStoreApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewGetObjectStoreApiResponse() *GetObjectStoreApiResponse {
  p := new(GetObjectStoreApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.GetObjectStoreApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.GetObjectStoreApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *GetObjectStoreApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *GetObjectStoreApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfGetObjectStoreApiResponseData()
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
REST response for all response codes in api path /files/v4.0.a2/config/file-server/quota-email-config Get operation
*/
type GetQuotaEmailConfigApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfGetQuotaEmailConfigApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewGetQuotaEmailConfigApiResponse() *GetQuotaEmailConfigApiResponse {
  p := new(GetQuotaEmailConfigApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.GetQuotaEmailConfigApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.GetQuotaEmailConfigApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *GetQuotaEmailConfigApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *GetQuotaEmailConfigApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfGetQuotaEmailConfigApiResponseData()
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
REST response for all response codes in api path /files/v4.0.a2/config/file-server/tier-config/{extId} Get operation
*/
type GetTieringConfigApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfGetTieringConfigApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewGetTieringConfigApiResponse() *GetTieringConfigApiResponse {
  p := new(GetTieringConfigApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.GetTieringConfigApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.GetTieringConfigApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *GetTieringConfigApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *GetTieringConfigApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfGetTieringConfigApiResponseData()
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
REST response for all response codes in api path /files/v4.0.a2/config/file-server/tier-config Get operation
*/
type GetTieringConfigListApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfGetTieringConfigListApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewGetTieringConfigListApiResponse() *GetTieringConfigListApiResponse {
  p := new(GetTieringConfigListApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.GetTieringConfigListApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.GetTieringConfigListApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *GetTieringConfigListApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *GetTieringConfigListApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfGetTieringConfigListApiResponseData()
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
REST response for all response codes in api path /files/v4.0.a2/config/file-server/tier-policies/{extId} Get operation
*/
type GetTieringPolicyApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfGetTieringPolicyApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewGetTieringPolicyApiResponse() *GetTieringPolicyApiResponse {
  p := new(GetTieringPolicyApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.GetTieringPolicyApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.GetTieringPolicyApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *GetTieringPolicyApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *GetTieringPolicyApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfGetTieringPolicyApiResponseData()
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
REST response for all response codes in api path /files/v4.0.a2/config/file-server/tier-policies Get operation
*/
type GetTieringPolicyListApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfGetTieringPolicyListApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewGetTieringPolicyListApiResponse() *GetTieringPolicyListApiResponse {
  p := new(GetTieringPolicyListApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.GetTieringPolicyListApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.GetTieringPolicyListApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *GetTieringPolicyListApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *GetTieringPolicyListApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfGetTieringPolicyListApiResponseData()
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
Blocked clients due to read-only or no access.
*/
type GlobalBlockedClient struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Flag to enable or disable global blocked clients.
  */
  BlockingEnabled *bool `json:"blockingEnabled,omitempty"`
  /**
  No access clients
  */
  NoAccessFilters []ClientBlockingFilter `json:"noAccessFilters,omitempty"`
  /**
  Read only access clients
  */
  RoAccessFilters []ClientBlockingFilter `json:"roAccessFilters,omitempty"`
}


func NewGlobalBlockedClient() *GlobalBlockedClient {
  p := new(GlobalBlockedClient)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.GlobalBlockedClient"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.GlobalBlockedClient"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST response for all response codes in api path /files/v4.0.a2/config/file-server/blocked-client Get operation
*/
type GlobalBlockedClientApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfGlobalBlockedClientApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewGlobalBlockedClientApiResponse() *GlobalBlockedClientApiResponse {
  p := new(GlobalBlockedClientApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.GlobalBlockedClientApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.GlobalBlockedClientApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *GlobalBlockedClientApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *GlobalBlockedClientApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfGlobalBlockedClientApiResponseData()
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
Identity mapping action type. The acceptable values are "DENY_ACCESS" or "MAP_IDENTITIES".
*/
type IdentityActionType int

const(
  IDENTITYACTIONTYPE_UNKNOWN IdentityActionType = 0
  IDENTITYACTIONTYPE_REDACTED IdentityActionType = 1
  IDENTITYACTIONTYPE_DENY_ACCESS IdentityActionType = 2
  IDENTITYACTIONTYPE_MAP_IDENTITIES IdentityActionType = 3
)

// returns the name of the enum given an ordinal number
func (e *IdentityActionType) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "DENY_ACCESS",
    "MAP_IDENTITIES",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *IdentityActionType) index(name string) IdentityActionType {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "DENY_ACCESS",
    "MAP_IDENTITIES",
  }
  for idx := range names {
    if names[idx] == name {
      return IdentityActionType(idx)
    }
  }
  return IDENTITYACTIONTYPE_UNKNOWN
}

func (e *IdentityActionType) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for IdentityActionType:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *IdentityActionType) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e IdentityActionType) Ref() *IdentityActionType {
  return &e
}


/**
Identity mapping option type. The acceptable values are "DENY", "ONE_TO_ONE", "WILD_CARD", "DEFAULT", "TEMPLATE".
*/
type IdentityMappingOptionType int

const(
  IDENTITYMAPPINGOPTIONTYPE_UNKNOWN IdentityMappingOptionType = 0
  IDENTITYMAPPINGOPTIONTYPE_REDACTED IdentityMappingOptionType = 1
  IDENTITYMAPPINGOPTIONTYPE_DENY IdentityMappingOptionType = 2
  IDENTITYMAPPINGOPTIONTYPE_ONE_TO_ONE IdentityMappingOptionType = 3
  IDENTITYMAPPINGOPTIONTYPE_WILD_CARD IdentityMappingOptionType = 4
  IDENTITYMAPPINGOPTIONTYPE_DEFAULT IdentityMappingOptionType = 5
  IDENTITYMAPPINGOPTIONTYPE_TEMPLATE IdentityMappingOptionType = 6
)

// returns the name of the enum given an ordinal number
func (e *IdentityMappingOptionType) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "DENY",
    "ONE_TO_ONE",
    "WILD_CARD",
    "DEFAULT",
    "TEMPLATE",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *IdentityMappingOptionType) index(name string) IdentityMappingOptionType {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "DENY",
    "ONE_TO_ONE",
    "WILD_CARD",
    "DEFAULT",
    "TEMPLATE",
  }
  for idx := range names {
    if names[idx] == name {
      return IdentityMappingOptionType(idx)
    }
  }
  return IDENTITYMAPPINGOPTIONTYPE_UNKNOWN
}

func (e *IdentityMappingOptionType) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for IdentityMappingOptionType:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *IdentityMappingOptionType) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e IdentityMappingOptionType) Ref() *IdentityMappingOptionType {
  return &e
}


/**
Identity mapping model. The acceptable values are "USER" or "GROUP".
*/
type IdentityMappingType int

const(
  IDENTITYMAPPINGTYPE_UNKNOWN IdentityMappingType = 0
  IDENTITYMAPPINGTYPE_REDACTED IdentityMappingType = 1
  IDENTITYMAPPINGTYPE_USER IdentityMappingType = 2
  IDENTITYMAPPINGTYPE_GROUP IdentityMappingType = 3
)

// returns the name of the enum given an ordinal number
func (e *IdentityMappingType) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "USER",
    "GROUP",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *IdentityMappingType) index(name string) IdentityMappingType {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "USER",
    "GROUP",
  }
  for idx := range names {
    if names[idx] == name {
      return IdentityMappingType(idx)
    }
  }
  return IDENTITYMAPPINGTYPE_UNKNOWN
}

func (e *IdentityMappingType) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for IdentityMappingType:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *IdentityMappingType) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e IdentityMappingType) Ref() *IdentityMappingType {
  return &e
}


/**
Infected files model
*/
type InfectedFile struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  File path
  */
  FilePath *string `json:"filePath,omitempty"`
  /**
  The {extId} of the file server. This is a read-only field.
  */
  FileServerExtId *string `json:"fileServerExtId,omitempty"`
  /**
  ICAP server details
  */
  IcapServer *string `json:"icapServer,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import1.ApiLink `json:"links,omitempty"`
  /**
  The {extId} of the mount target. Example:9c1e537d-6777-4c22-5d41-ddd0c3337aa9.
  */
  MountTargetExtId *string `json:"mountTargetExtId,omitempty"`
  /**
  Quarantine infected files.
  */
  Quarantined *bool `json:"quarantined,omitempty"`
  /**
  Scan timestamp in sec.s
  */
  ScanTimestampUsecs *string `json:"scanTimestampUsecs,omitempty"`
  /**
  Share name
  */
  ShareName *string `json:"shareName,omitempty"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
  /**
  Threat description
  */
  ThreatDescription *string `json:"threatDescription,omitempty"`
}


func NewInfectedFile() *InfectedFile {
  p := new(InfectedFile)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.InfectedFile"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.InfectedFile"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Comma-separated list of infected file UUID.
*/
type InfectedFileSpec struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Comma-separated list of infected file UUID.
  */
  InfectedFileUuids []string `json:"infectedFileUuids"`
}

func (p *InfectedFileSpec) MarshalJSON() ([]byte, error) {
  type InfectedFileSpecProxy InfectedFileSpec
  return json.Marshal(struct {
    *InfectedFileSpecProxy
    InfectedFileUuids []string `json:"infectedFileUuids,omitempty"`
  }{
    InfectedFileSpecProxy : (*InfectedFileSpecProxy)(p),
    InfectedFileUuids : p.InfectedFileUuids,
  })
}

func NewInfectedFileSpec() *InfectedFileSpec {
  p := new(InfectedFileSpec)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.InfectedFileSpec"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.InfectedFileSpec"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST response for all response codes in api path /files/v4.0.a2/config/file-server/infected-files/{infectedFileExtId} Get operation
*/
type InfectedFilesApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfInfectedFilesApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewInfectedFilesApiResponse() *InfectedFilesApiResponse {
  p := new(InfectedFilesApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.InfectedFilesApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.InfectedFilesApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *InfectedFilesApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *InfectedFilesApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfInfectedFilesApiResponseData()
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
REST response for all response codes in api path /files/v4.0.a2/config/file-server/infected-files Get operation
*/
type InfectedFilesListApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfInfectedFilesListApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewInfectedFilesListApiResponse() *InfectedFilesListApiResponse {
  p := new(InfectedFilesListApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.InfectedFilesListApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.InfectedFilesListApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *InfectedFilesListApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *InfectedFilesListApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfInfectedFilesListApiResponseData()
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
Status of each job. The acceptable values are "JOBSTATUS_RUNNING", "JOBSTATUS_QUEUED", "JOBSTATUS_SUCCEEDED", "JOBSTATUS_CANCELLED", "JOBSTATUS_JOB_PAUSED", "JOBSTATUS_JOB_FAILED".
*/
type JobStatus int

const(
  JOBSTATUS_UNKNOWN JobStatus = 0
  JOBSTATUS_REDACTED JobStatus = 1
  JOBSTATUS_JOBSTATUS_RUNNING JobStatus = 2
  JOBSTATUS_JOBSTATUS_QUEUED JobStatus = 3
  JOBSTATUS_JOBSTATUS_SUCCEEDED JobStatus = 4
  JOBSTATUS_JOBSTATUS_CANCELLED JobStatus = 5
  JOBSTATUS_JOBSTATUS_JOB_PAUSED JobStatus = 6
  JOBSTATUS_JOBSTATUS_JOB_FAILED JobStatus = 7
)

// returns the name of the enum given an ordinal number
func (e *JobStatus) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "JOBSTATUS_RUNNING",
    "JOBSTATUS_QUEUED",
    "JOBSTATUS_SUCCEEDED",
    "JOBSTATUS_CANCELLED",
    "JOBSTATUS_JOB_PAUSED",
    "JOBSTATUS_JOB_FAILED",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *JobStatus) index(name string) JobStatus {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "JOBSTATUS_RUNNING",
    "JOBSTATUS_QUEUED",
    "JOBSTATUS_SUCCEEDED",
    "JOBSTATUS_CANCELLED",
    "JOBSTATUS_JOB_PAUSED",
    "JOBSTATUS_JOB_FAILED",
  }
  for idx := range names {
    if names[idx] == name {
      return JobStatus(idx)
    }
  }
  return JOBSTATUS_UNKNOWN
}

func (e *JobStatus) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for JobStatus:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *JobStatus) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e JobStatus) Ref() *JobStatus {
  return &e
}


/**
LDAP domain associated with the file server.
*/
type Ldap struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Base DN
  */
  Base *string `json:"base,omitempty"`
  /**
  Bind DN
  */
  Binddn *string `json:"binddn,omitempty"`
  /**
  LDAP password
  */
  Bindpw *string `json:"bindpw,omitempty"`
  
  ProtocolType *ProtocolType `json:"protocolType"`
  /**
  LDAP SSH certificate key
  */
  TlsCacertContent *string `json:"tlsCacertContent,omitempty"`
  
  TlsReqCert *TLSReqCertStatus `json:"tlsReqCert,omitempty"`
  /**
  URI of the ldap domain.
  */
  Uri *string `json:"uri,omitempty"`
}

func (p *Ldap) MarshalJSON() ([]byte, error) {
  type LdapProxy Ldap
  return json.Marshal(struct {
    *LdapProxy
    ProtocolType *ProtocolType `json:"protocolType,omitempty"`
  }{
    LdapProxy : (*LdapProxy)(p),
    ProtocolType : p.ProtocolType,
  })
}

func NewLdap() *Ldap {
  p := new(Ldap)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.Ldap"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.Ldap"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST response for all response codes in api path /files/v4.0.a2/config/file-server/mount-targets/{mountTargetExtId}/virus-scan-policy Get operation
*/
type ListMountTargetVirusScanPolicyApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfListMountTargetVirusScanPolicyApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewListMountTargetVirusScanPolicyApiResponse() *ListMountTargetVirusScanPolicyApiResponse {
  p := new(ListMountTargetVirusScanPolicyApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.ListMountTargetVirusScanPolicyApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.ListMountTargetVirusScanPolicyApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *ListMountTargetVirusScanPolicyApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *ListMountTargetVirusScanPolicyApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfListMountTargetVirusScanPolicyApiResponseData()
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
REST response for all response codes in api path /files/v4.0.a2/config/file-server/object-store/profiles Get operation
*/
type ListObjectStoreApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfListObjectStoreApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewListObjectStoreApiResponse() *ListObjectStoreApiResponse {
  p := new(ListObjectStoreApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.ListObjectStoreApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.ListObjectStoreApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *ListObjectStoreApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *ListObjectStoreApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfListObjectStoreApiResponseData()
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
REST response for all response codes in api path /files/v4.0.a2/config/file-server/virus-scan-policy Get operation
*/
type ListVirusScanPolicyApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfListVirusScanPolicyApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewListVirusScanPolicyApiResponse() *ListVirusScanPolicyApiResponse {
  p := new(ListVirusScanPolicyApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.ListVirusScanPolicyApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.ListVirusScanPolicyApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *ListVirusScanPolicyApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *ListVirusScanPolicyApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfListVirusScanPolicyApiResponseData()
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
REST response for all response codes in api path /files/v4.0.a2/config/file-server/recommendations/{recommendationExtId}/$actions/perform-recommendation Post operation
*/
type LoadBalanceFileServerApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfLoadBalanceFileServerApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewLoadBalanceFileServerApiResponse() *LoadBalanceFileServerApiResponse {
  p := new(LoadBalanceFileServerApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.LoadBalanceFileServerApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.LoadBalanceFileServerApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *LoadBalanceFileServerApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *LoadBalanceFileServerApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfLoadBalanceFileServerApiResponseData()
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
Model associated with load balance file server.
*/
type LoadBalanceSpec struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  ScaleOut *ScaleOutSpec `json:"scaleOut,omitempty"`
  
  ScaleUp *ScaleUpSpec `json:"scaleUp,omitempty"`
}


func NewLoadBalanceSpec() *LoadBalanceSpec {
  p := new(LoadBalanceSpec)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.LoadBalanceSpec"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.LoadBalanceSpec"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Local domain object.
*/
type LocalDomain struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  ProtocolType *ProtocolType `json:"protocolType"`
}

func (p *LocalDomain) MarshalJSON() ([]byte, error) {
  type LocalDomainProxy LocalDomain
  return json.Marshal(struct {
    *LocalDomainProxy
    ProtocolType *ProtocolType `json:"protocolType,omitempty"`
  }{
    LocalDomainProxy : (*LocalDomainProxy)(p),
    ProtocolType : p.ProtocolType,
  })
}

func NewLocalDomain() *LocalDomain {
  p := new(LocalDomain)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.LocalDomain"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.LocalDomain"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Migration pair details
*/
type MigrationPair struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import1.ApiLink `json:"links,omitempty"`
  /**
  Absolute source path.

For the source paths ending with '/', the contents will be copied directly into the target folder.\n
For the paths, ending with name, a new folder with the same name will be created on the
target server and then the data is synced."
  */
  SourcePath *string `json:"sourcePath"`
  /**
  Absolute path for the target. For example: '/share1/dir1'. Path should always start with a root folder.
  */
  TargetPath *string `json:"targetPath,omitempty"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
}

func (p *MigrationPair) MarshalJSON() ([]byte, error) {
  type MigrationPairProxy MigrationPair
  return json.Marshal(struct {
    *MigrationPairProxy
    SourcePath *string `json:"sourcePath,omitempty"`
  }{
    MigrationPairProxy : (*MigrationPairProxy)(p),
    SourcePath : p.SourcePath,
  })
}

func NewMigrationPair() *MigrationPair {
  p := new(MigrationPair)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.MigrationPair"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.MigrationPair"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Migration plan details
*/
type MigrationPlan struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Description of migration plan
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
  Migration plan name
  */
  Name *string `json:"name"`
  /**
  Source ExtId for the migration plan
  */
  SourceExtId *string `json:"sourceExtId"`
  
  SubPlans []MigrationSubPlan `json:"subPlans"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
}

func (p *MigrationPlan) MarshalJSON() ([]byte, error) {
  type MigrationPlanProxy MigrationPlan
  return json.Marshal(struct {
    *MigrationPlanProxy
    Name *string `json:"name,omitempty"`
    SourceExtId *string `json:"sourceExtId,omitempty"`
    SubPlans []MigrationSubPlan `json:"subPlans,omitempty"`
  }{
    MigrationPlanProxy : (*MigrationPlanProxy)(p),
    Name : p.Name,
    SourceExtId : p.SourceExtId,
    SubPlans : p.SubPlans,
  })
}

func NewMigrationPlan() *MigrationPlan {
  p := new(MigrationPlan)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.MigrationPlan"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.MigrationPlan"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
An enum that represents the plan status
*/
type MigrationPlanStatus int

const(
  MIGRATIONPLANSTATUS_UNKNOWN MigrationPlanStatus = 0
  MIGRATIONPLANSTATUS_REDACTED MigrationPlanStatus = 1
  MIGRATIONPLANSTATUS_DEFAULT MigrationPlanStatus = 2
  MIGRATIONPLANSTATUS_QUEUED MigrationPlanStatus = 3
  MIGRATIONPLANSTATUS_RUNNING MigrationPlanStatus = 4
  MIGRATIONPLANSTATUS_SUCCEEDED MigrationPlanStatus = 5
  MIGRATIONPLANSTATUS_CANCELLED MigrationPlanStatus = 6
  MIGRATIONPLANSTATUS_PAUSED MigrationPlanStatus = 7
  MIGRATIONPLANSTATUS_FAILED MigrationPlanStatus = 8
  MIGRATIONPLANSTATUS_INPROGRESS MigrationPlanStatus = 9
)

// returns the name of the enum given an ordinal number
func (e *MigrationPlanStatus) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "DEFAULT",
    "QUEUED",
    "RUNNING",
    "SUCCEEDED",
    "CANCELLED",
    "PAUSED",
    "FAILED",
    "INPROGRESS",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *MigrationPlanStatus) index(name string) MigrationPlanStatus {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "DEFAULT",
    "QUEUED",
    "RUNNING",
    "SUCCEEDED",
    "CANCELLED",
    "PAUSED",
    "FAILED",
    "INPROGRESS",
  }
  for idx := range names {
    if names[idx] == name {
      return MigrationPlanStatus(idx)
    }
  }
  return MIGRATIONPLANSTATUS_UNKNOWN
}

func (e *MigrationPlanStatus) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for MigrationPlanStatus:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *MigrationPlanStatus) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e MigrationPlanStatus) Ref() *MigrationPlanStatus {
  return &e
}


/**
Migration source details
*/
type MigrationSource struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Address object.
  */
  Address *string `json:"address"`
  /**
  Migration source alias
  */
  Alias *string `json:"alias"`
  /**
  Description for migration source
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
  DNS password.
  */
  Password *string `json:"password,omitempty"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
  /**
  User for the migration source
  */
  User *string `json:"user,omitempty"`
}

func (p *MigrationSource) MarshalJSON() ([]byte, error) {
  type MigrationSourceProxy MigrationSource
  return json.Marshal(struct {
    *MigrationSourceProxy
    Address *string `json:"address,omitempty"`
    Alias *string `json:"alias,omitempty"`
  }{
    MigrationSourceProxy : (*MigrationSourceProxy)(p),
    Address : p.Address,
    Alias : p.Alias,
  })
}

func NewMigrationSource() *MigrationSource {
  p := new(MigrationSource)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.MigrationSource"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.MigrationSource"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Migration plan stats details
*/
type MigrationStats struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  The average data throughput for migration sub-plan in bytes/sec.
  */
  AverageDataThroughput *float32 `json:"averageDataThroughput,omitempty"`
  /**
  The average entities throughput for migration sub-plan in entities/sec.
  */
  AverageEntitiesThroughput *int64 `json:"averageEntitiesThroughput,omitempty"`
  /**
  The average average file scan rate for migration sub-plan in files/sec.
  */
  AverageFilesScannedRate *int64 `json:"averageFilesScannedRate,omitempty"`
  /**
  The number of migration sub-plan bytes metadata transferred.
  */
  BytesMetadataTransferred *int64 `json:"bytesMetadataTransferred,omitempty"`
  /**
  The transferred bytes for the migration sub-plan.
  */
  BytesTransferred *int64 `json:"bytesTransferred,omitempty"`
  /**
  The number of the transferred directories for the migration sub-plan.
  */
  DirectoriesTransferred *int64 `json:"directoriesTransferred,omitempty"`
  /**
  End time for the migration sub-plan.
  */
  EndTime *string `json:"endTime,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  The number of the transferred failed directories for the migration sub-plan.
  */
  FailedDirectoriesCount *int64 `json:"failedDirectoriesCount,omitempty"`
  /**
  The number of the transferred failed files for the migration sub-plan.
  */
  FailedFilesCount *int64 `json:"failedFilesCount,omitempty"`
  /**
  The number of migration sub-plan failed streams.
  */
  FailedStreamsCount *int64 `json:"failedStreamsCount,omitempty"`
  /**
  The number of migration sub-plan files scanned.
  */
  FilesScanned *int64 `json:"filesScanned,omitempty"`
  /**
  The number of migration sub-plan files skipped.
  */
  FilesSkipped *int64 `json:"filesSkipped,omitempty"`
  /**
  The number of the transferred files for the migration sub-plan.
  */
  FilesTransferred *int64 `json:"filesTransferred,omitempty"`
  /**
  The iteration number for Migration sub-plan.
  */
  IterationNumber *int64 `json:"iterationNumber,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import1.ApiLink `json:"links,omitempty"`
  /**
  Start time for the migration sub-plan.
  */
  StartTime *string `json:"startTime,omitempty"`
  
  StatusInfo *StatusInfo `json:"statusInfo,omitempty"`
  /**
  The number of migration sub-plan streams transferred.
  */
  StreamsTransferred *int64 `json:"streamsTransferred,omitempty"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
}


func NewMigrationStats() *MigrationStats {
  p := new(MigrationStats)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.MigrationStats"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.MigrationStats"}
  p.UnknownFields_ = map[string]interface{}{}

  p.AverageDataThroughput = new(float32)
  *p.AverageDataThroughput = 0
  p.AverageFilesScannedRate = new(int64)
  *p.AverageFilesScannedRate = 0
  p.BytesMetadataTransferred = new(int64)
  *p.BytesMetadataTransferred = 0
  p.BytesTransferred = new(int64)
  *p.BytesTransferred = 0
  p.DirectoriesTransferred = new(int64)
  *p.DirectoriesTransferred = 0
  p.FailedDirectoriesCount = new(int64)
  *p.FailedDirectoriesCount = 0
  p.FailedFilesCount = new(int64)
  *p.FailedFilesCount = 0
  p.FailedStreamsCount = new(int64)
  *p.FailedStreamsCount = 0
  p.FilesScanned = new(int64)
  *p.FilesScanned = 0
  p.FilesSkipped = new(int64)
  *p.FilesSkipped = 0
  p.FilesTransferred = new(int64)
  *p.FilesTransferred = 0
  p.StreamsTransferred = new(int64)
  *p.StreamsTransferred = 0


  return p
}




/**
Migration sub plan details
*/
type MigrationSubPlan struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  Flag for sync type. If it is false, then source will be only synced. If it is true, then source will be mirrored on target.
  */
  IsSyncTypeMirror *bool `json:"isSyncTypeMirror,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import1.ApiLink `json:"links,omitempty"`
  
  MigrationPair *MigrationPair `json:"migrationPair"`
  /**
  Migration sub plan name
  */
  Name *string `json:"name,omitempty"`
  
  NfsOptions *NfsOptions `json:"nfsOptions,omitempty"`
  
  Protocol *SourceShareProtocolType `json:"protocol"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
}

func (p *MigrationSubPlan) MarshalJSON() ([]byte, error) {
  type MigrationSubPlanProxy MigrationSubPlan
  return json.Marshal(struct {
    *MigrationSubPlanProxy
    MigrationPair *MigrationPair `json:"migrationPair,omitempty"`
    Protocol *SourceShareProtocolType `json:"protocol,omitempty"`
  }{
    MigrationSubPlanProxy : (*MigrationSubPlanProxy)(p),
    MigrationPair : p.MigrationPair,
    Protocol : p.Protocol,
  })
}

func NewMigrationSubPlan() *MigrationSubPlan {
  p := new(MigrationSubPlan)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.MigrationSubPlan"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.MigrationSubPlan"}
  p.UnknownFields_ = map[string]interface{}{}

  p.IsSyncTypeMirror = new(bool)
  *p.IsSyncTypeMirror = false


  return p
}




/**
Response for the get mimetype of files.
*/
type MimeTypeResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  FileSpecs []FileInfoSpec `json:"fileSpecs,omitempty"`
}


func NewMimeTypeResponse() *MimeTypeResponse {
  p := new(MimeTypeResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.MimeTypeResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.MimeTypeResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST response for all response codes in api path /files/v4.0.a2/config/file-server/mount-targets/{mountTargetExtId}/files Get operation
*/
type MimeTypeResponseApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfMimeTypeResponseApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewMimeTypeResponseApiResponse() *MimeTypeResponseApiResponse {
  p := new(MimeTypeResponseApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.MimeTypeResponseApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.MimeTypeResponseApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *MimeTypeResponseApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *MimeTypeResponseApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfMimeTypeResponseApiResponseData()
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
Mount target object.
*/
type MountTarget struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  BlockedClients *BlockedClients `json:"blockedClients,omitempty"`
  /**
  Mount target description.
  */
  Description *string `json:"description,omitempty"`
  /**
  Flag to enable compression. This field will be deprecated.
  */
  EnableCompression *bool `json:"enableCompression,omitempty"`
  /**
  Flag to enable windows previous version. This field will be deprecated.
  */
  EnablePreviousVersion *bool `json:"enablePreviousVersion,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  Comma-separated list of file blocking extensions.
  */
  FileBlockingExtensions []string `json:"fileBlockingExtensions,omitempty"`
  /**
  The {extId} of the file server. This is a read-only field.
  */
  FileServerExtId *string `json:"fileServerExtId,omitempty"`
  /**
  Flag to enable compression.
  */
  IsEnableCompression *bool `json:"isEnableCompression,omitempty"`
  /**
  Flag to enable windows previous version.
  */
  IsEnablePreviousVersion *bool `json:"isEnablePreviousVersion,omitempty"`
  /**
  Enable long name support.
  */
  IsLongnameEnabled *bool `json:"isLongnameEnabled,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import1.ApiLink `json:"links,omitempty"`
  /**
  Enable long name support. This field will be deprecated.
  */
  LongnameEnabled *bool `json:"longnameEnabled,omitempty"`
  /**
  Maximum size of mount target in GiB.
  */
  MaxSizeGib *int64 `json:"maxSizeGib,omitempty"`
  
  MultiProtocolProperties *MultiProtocol `json:"multiProtocolProperties,omitempty"`
  /**
  Mount target name.
  */
  Name *string `json:"name"`
  
  NfsProperties *NfsProtocol `json:"nfsProperties,omitempty"`
  /**
  {extId} of the parent mount target. This is a read-only field.
  */
  ParentMountTargetExtId *string `json:"parentMountTargetExtId,omitempty"`
  /**
  Path of the nested mount target.
  */
  Path *string `json:"path,omitempty"`
  
  Protocol *ProtocolTypeMountTarget `json:"protocol"`
  /**
  List of secondary protocol type for the mount target. The acceptable values are "SMB", "NFS", "NONE", "INCOMPATIBLE".
  */
  SecondaryProtocol []ProtocolTypeMountTarget `json:"secondaryProtocol,omitempty"`
  
  SmbProperties *SmbProtocol `json:"smbProperties,omitempty"`
  
  State *MountTargetState `json:"state,omitempty"`
  
  StatusType *StatusType `json:"statusType,omitempty"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
  
  Type *MountTargetType `json:"type"`
  
  WorkloadType *MountTargetWorkloadType `json:"workloadType,omitempty"`
  
  WormSpec *WormSpec `json:"wormSpec,omitempty"`
}

func (p *MountTarget) MarshalJSON() ([]byte, error) {
  type MountTargetProxy MountTarget
  return json.Marshal(struct {
    *MountTargetProxy
    Name *string `json:"name,omitempty"`
    Protocol *ProtocolTypeMountTarget `json:"protocol,omitempty"`
    Type *MountTargetType `json:"type,omitempty"`
  }{
    MountTargetProxy : (*MountTargetProxy)(p),
    Name : p.Name,
    Protocol : p.Protocol,
    Type : p.Type,
  })
}

func NewMountTarget() *MountTarget {
  p := new(MountTarget)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.MountTarget"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.MountTarget"}
  p.UnknownFields_ = map[string]interface{}{}

  p.EnablePreviousVersion = new(bool)
  *p.EnablePreviousVersion = false
  p.IsEnablePreviousVersion = new(bool)
  *p.IsEnablePreviousVersion = false
  p.IsLongnameEnabled = new(bool)
  *p.IsLongnameEnabled = false
  p.LongnameEnabled = new(bool)
  *p.LongnameEnabled = false


  return p
}




/**
REST response for all response codes in api path /files/v4.0.a2/config/file-server/mount-targets/{mountTargetExtId} Get operation
*/
type MountTargetApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfMountTargetApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewMountTargetApiResponse() *MountTargetApiResponse {
  p := new(MountTargetApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.MountTargetApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.MountTargetApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *MountTargetApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *MountTargetApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfMountTargetApiResponseData()
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
REST response for all response codes in api path /files/v4.0.a2/config/file-server/mount-targets Get operation
*/
type MountTargetListApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfMountTargetListApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewMountTargetListApiResponse() *MountTargetListApiResponse {
  p := new(MountTargetListApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.MountTargetListApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.MountTargetListApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *MountTargetListApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *MountTargetListApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfMountTargetListApiResponseData()
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
REST response for all response codes in api path /files/v4.0.a2/config/file-server/mount-targets/{mountTargetExtId}/snapshots/{snapshotExtId} Get operation
*/
type MountTargetSnapshotApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfMountTargetSnapshotApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewMountTargetSnapshotApiResponse() *MountTargetSnapshotApiResponse {
  p := new(MountTargetSnapshotApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.MountTargetSnapshotApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.MountTargetSnapshotApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *MountTargetSnapshotApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *MountTargetSnapshotApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfMountTargetSnapshotApiResponseData()
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
REST response for all response codes in api path /files/v4.0.a2/config/file-server/mount-targets/{mountTargetExtId}/snapshots Get operation
*/
type MountTargetSnapshotListApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfMountTargetSnapshotListApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewMountTargetSnapshotListApiResponse() *MountTargetSnapshotListApiResponse {
  p := new(MountTargetSnapshotListApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.MountTargetSnapshotListApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.MountTargetSnapshotListApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *MountTargetSnapshotListApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *MountTargetSnapshotListApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfMountTargetSnapshotListApiResponseData()
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
State of the mount target. The acceptable values are "ONLINE" or "OFFLINE". This is a read-only field.
*/
type MountTargetState int

const(
  MOUNTTARGETSTATE_UNKNOWN MountTargetState = 0
  MOUNTTARGETSTATE_REDACTED MountTargetState = 1
  MOUNTTARGETSTATE_ONLINE MountTargetState = 2
  MOUNTTARGETSTATE_OFFLINE MountTargetState = 3
)

// returns the name of the enum given an ordinal number
func (e *MountTargetState) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "ONLINE",
    "OFFLINE",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *MountTargetState) index(name string) MountTargetState {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "ONLINE",
    "OFFLINE",
  }
  for idx := range names {
    if names[idx] == name {
      return MountTargetState(idx)
    }
  }
  return MOUNTTARGETSTATE_UNKNOWN
}

func (e *MountTargetState) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for MountTargetState:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *MountTargetState) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e MountTargetState) Ref() *MountTargetState {
  return &e
}


/**
Supported mount target types. The acceptable values are "HOMES", "GENERAL", "DISTRIBUTED", "STANDARD".
*/
type MountTargetType int

const(
  MOUNTTARGETTYPE_UNKNOWN MountTargetType = 0
  MOUNTTARGETTYPE_REDACTED MountTargetType = 1
  MOUNTTARGETTYPE_HOMES MountTargetType = 2
  MOUNTTARGETTYPE_GENERAL MountTargetType = 3
  MOUNTTARGETTYPE_DISTRIBUTED MountTargetType = 4
  MOUNTTARGETTYPE_STANDARD MountTargetType = 5
)

// returns the name of the enum given an ordinal number
func (e *MountTargetType) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "HOMES",
    "GENERAL",
    "DISTRIBUTED",
    "STANDARD",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *MountTargetType) index(name string) MountTargetType {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "HOMES",
    "GENERAL",
    "DISTRIBUTED",
    "STANDARD",
  }
  for idx := range names {
    if names[idx] == name {
      return MountTargetType(idx)
    }
  }
  return MOUNTTARGETTYPE_UNKNOWN
}

func (e *MountTargetType) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for MountTargetType:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *MountTargetType) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e MountTargetType) Ref() *MountTargetType {
  return &e
}


/**
REST response for all response codes in api path /files/v4.0.a2/config/file-server/mount-targets/{mountTargetExtId}/virus-scan-policy/{virusScanPolicyExtId} Get operation
*/
type MountTargetVirusScanPolicyApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfMountTargetVirusScanPolicyApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewMountTargetVirusScanPolicyApiResponse() *MountTargetVirusScanPolicyApiResponse {
  p := new(MountTargetVirusScanPolicyApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.MountTargetVirusScanPolicyApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.MountTargetVirusScanPolicyApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *MountTargetVirusScanPolicyApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *MountTargetVirusScanPolicyApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfMountTargetVirusScanPolicyApiResponseData()
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
Supported mount target workload types. The acceptable values are "RANDOM", "DEFAULT", "SEQUENTIAL".
*/
type MountTargetWorkloadType int

const(
  MOUNTTARGETWORKLOADTYPE_UNKNOWN MountTargetWorkloadType = 0
  MOUNTTARGETWORKLOADTYPE_REDACTED MountTargetWorkloadType = 1
  MOUNTTARGETWORKLOADTYPE_UNDEFINED MountTargetWorkloadType = 2
  MOUNTTARGETWORKLOADTYPE_RANDOM MountTargetWorkloadType = 3
  MOUNTTARGETWORKLOADTYPE_DEFAULT MountTargetWorkloadType = 4
  MOUNTTARGETWORKLOADTYPE_SEQUENTIAL MountTargetWorkloadType = 5
)

// returns the name of the enum given an ordinal number
func (e *MountTargetWorkloadType) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "UNDEFINED",
    "RANDOM",
    "DEFAULT",
    "SEQUENTIAL",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *MountTargetWorkloadType) index(name string) MountTargetWorkloadType {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "UNDEFINED",
    "RANDOM",
    "DEFAULT",
    "SEQUENTIAL",
  }
  for idx := range names {
    if names[idx] == name {
      return MountTargetWorkloadType(idx)
    }
  }
  return MOUNTTARGETWORKLOADTYPE_UNKNOWN
}

func (e *MountTargetWorkloadType) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for MountTargetWorkloadType:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *MountTargetWorkloadType) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e MountTargetWorkloadType) Ref() *MountTargetWorkloadType {
  return &e
}


/**
Multi protocol object.
*/
type MultiProtocol struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Flag to enable case sensitive name space. This field will be deprecated.
  */
  EnableCasesensitiveNamespace *bool `json:"enableCasesensitiveNamespace,omitempty"`
  /**
  Flag to enable concurrent reads. This field will be deprecated.
  */
  EnableConcurrentReads *bool `json:"enableConcurrentReads,omitempty"`
  /**
  Flag to enable simultaneous access. This field will be deprecated.
  */
  EnableSimulteneousAccess *bool `json:"enableSimulteneousAccess,omitempty"`
  /**
  Flag to enable symlink creation. This field will be deprecated.
  */
  EnableSymlinkCreation *bool `json:"enableSymlinkCreation,omitempty"`
  /**
  Flag to enable case sensitive name space.
  */
  IsEnableCasesensitiveNamespace *bool `json:"isEnableCasesensitiveNamespace,omitempty"`
  /**
  Flag to enable concurrent reads.
  */
  IsEnableConcurrentReads *bool `json:"isEnableConcurrentReads,omitempty"`
  /**
  Flag to enable simultaneous access.
  */
  IsEnableSimulteneousAccess *bool `json:"isEnableSimulteneousAccess,omitempty"`
  /**
  Flag to enable symlink creation.
  */
  IsEnableSymlinkCreation *bool `json:"isEnableSymlinkCreation,omitempty"`
}


func NewMultiProtocol() *MultiProtocol {
  p := new(MultiProtocol)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.MultiProtocol"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.MultiProtocol"}
  p.UnknownFields_ = map[string]interface{}{}

  p.EnableCasesensitiveNamespace = new(bool)
  *p.EnableCasesensitiveNamespace = false
  p.EnableSimulteneousAccess = new(bool)
  *p.EnableSimulteneousAccess = false
  p.EnableSymlinkCreation = new(bool)
  *p.EnableSymlinkCreation = false
  p.IsEnableCasesensitiveNamespace = new(bool)
  *p.IsEnableCasesensitiveNamespace = false
  p.IsEnableSimulteneousAccess = new(bool)
  *p.IsEnableSimulteneousAccess = false
  p.IsEnableSymlinkCreation = new(bool)
  *p.IsEnableSymlinkCreation = false


  return p
}




/**
Configure the name services model.
*/
type NameServiceSpec struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  AdDomain *Ad `json:"adDomain,omitempty"`
  
  LdapDomain *Ldap `json:"ldapDomain,omitempty"`
  
  LocalDomain *LocalDomain `json:"localDomain,omitempty"`
  /**
  NFS V4 domain associated with the file server.
  */
  NfsV4Domain *string `json:"nfsV4Domain,omitempty"`
  
  NfsVersion *NfsVersion `json:"nfsVersion,omitempty"`
}


func NewNameServiceSpec() *NameServiceSpec {
  p := new(NameServiceSpec)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.NameServiceSpec"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.NameServiceSpec"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST response for all response codes in api path /files/v4.0.a2/config/file-server/name-services Get operation
*/
type NameServicesApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfNameServicesApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewNameServicesApiResponse() *NameServicesApiResponse {
  p := new(NameServicesApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.NameServicesApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.NameServicesApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *NameServicesApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *NameServicesApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfNameServicesApiResponseData()
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
REST response for all response codes in api path /files/v4.0.a2/config/file-server/mount-targets/{mountTargetExtId}/nested-mount-targets Get operation
*/
type NestedMountTargetListApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfNestedMountTargetListApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewNestedMountTargetListApiResponse() *NestedMountTargetListApiResponse {
  p := new(NestedMountTargetListApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.NestedMountTargetListApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.NestedMountTargetListApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *NestedMountTargetListApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *NestedMountTargetListApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfNestedMountTargetListApiResponseData()
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
File server network object.
*/
type Network struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  DefaultGateway *import3.IPAddress `json:"defaultGateway,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  IPV4 address range to be used in this network.
  */
  IpAddresses []import3.IPAddress `json:"ipAddresses,omitempty"`
  /**
  Prefix length for IPV6 address
  */
  Ipv6PrefixLength *int `json:"ipv6PrefixLength,omitempty"`
  /**
  Indicates whether the current file server network is managed or unmanaged. This is a read-only field.
  */
  IsManaged *bool `json:"isManaged,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import1.ApiLink `json:"links,omitempty"`
  /**
  UUID of the the file server network.
  */
  NetworkExtId *string `json:"networkExtId,omitempty"`
  
  SubnetMask *import3.IPAddress `json:"subnetMask,omitempty"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
  
  VirtualIpAddress *import3.IPAddress `json:"virtualIpAddress,omitempty"`
  /**
  File server network name. This is a read-only field.
  */
  VirtualNetworkName *string `json:"virtualNetworkName,omitempty"`
  /**
  File server network VLAN Id.
  */
  VlanId *int `json:"vlanId,omitempty"`
}


func NewNetwork() *Network {
  p := new(Network)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.Network"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.Network"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST response for all response codes in api path /files/v4.0.a2/config/file-server/networks/{extId} Get operation
*/
type NetworkGetByExtIdApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfNetworkGetByExtIdApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewNetworkGetByExtIdApiResponse() *NetworkGetByExtIdApiResponse {
  p := new(NetworkGetByExtIdApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.NetworkGetByExtIdApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.NetworkGetByExtIdApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *NetworkGetByExtIdApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *NetworkGetByExtIdApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfNetworkGetByExtIdApiResponseData()
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
REST response for all response codes in api path /files/v4.0.a2/config/file-server/networks Get operation
*/
type NetworkListApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfNetworkListApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewNetworkListApiResponse() *NetworkListApiResponse {
  p := new(NetworkListApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.NetworkListApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.NetworkListApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *NetworkListApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *NetworkListApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfNetworkListApiResponseData()
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
NFS mapping model.
*/
type NfsIdentitiesMapping struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  ActionType *IdentityActionType `json:"actionType,omitempty"`
  
  MappingType *IdentityMappingType `json:"mappingType,omitempty"`
  /**
  NFS user name.
  */
  Name *string `json:"name,omitempty"`
}


func NewNfsIdentitiesMapping() *NfsIdentitiesMapping {
  p := new(NfsIdentitiesMapping)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.NfsIdentitiesMapping"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.NfsIdentitiesMapping"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
NFS options details.
*/
type NfsOptions struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  AuthType *AuthType `json:"authType,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  Hard link optimization for NFS.
  */
  HardlinkOptimization *bool `json:"hardlinkOptimization,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import1.ApiLink `json:"links,omitempty"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
}


func NewNfsOptions() *NfsOptions {
  p := new(NfsOptions)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.NfsOptions"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.NfsOptions"}
  p.UnknownFields_ = map[string]interface{}{}

  p.HardlinkOptimization = new(bool)
  *p.HardlinkOptimization = true


  return p
}




/**
NFS protocol object.
*/
type NfsProtocol struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  AnonymousIdentifier *AnonymousIdentifier `json:"anonymousIdentifier,omitempty"`
  
  AuthenticationType *FilesAuthenticationType `json:"authenticationType,omitempty"`
  
  ClientAccessType *ClientAccessType `json:"clientAccessType,omitempty"`
  
  SquashType *SquashType `json:"squashType,omitempty"`
}


func NewNfsProtocol() *NfsProtocol {
  p := new(NfsProtocol)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.NfsProtocol"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.NfsProtocol"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
NFS version associated with the file server. The acceptable values are "NFSV4", "NFSV3", "NFSV3V4".
*/
type NfsVersion int

const(
  NFSVERSION_UNKNOWN NfsVersion = 0
  NFSVERSION_REDACTED NfsVersion = 1
  NFSVERSION_NFSV4 NfsVersion = 2
  NFSVERSION_NFSV3 NfsVersion = 3
  NFSVERSION_NFSV3V4 NfsVersion = 4
)

// returns the name of the enum given an ordinal number
func (e *NfsVersion) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "NFSV4",
    "NFSV3",
    "NFSV3V4",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *NfsVersion) index(name string) NfsVersion {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "NFSV4",
    "NFSV3",
    "NFSV3V4",
  }
  for idx := range names {
    if names[idx] == name {
      return NfsVersion(idx)
    }
  }
  return NFSVERSION_UNKNOWN
}

func (e *NfsVersion) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for NfsVersion:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *NfsVersion) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e NfsVersion) Ref() *NfsVersion {
  return &e
}


/**
Schedule for each day of week.
*/
type OneScheduleSpec struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  The time difference between the start and the end time of schedule in minutes
  */
  DurationMinutes *int64 `json:"durationMinutes"`
  /**
  Start time hour of the schedule
  */
  StartTimeHours *int64 `json:"startTimeHours"`
  /**
  Start time minute of the schedule
  */
  StartTimeMinutes *int64 `json:"startTimeMinutes"`
}

func (p *OneScheduleSpec) MarshalJSON() ([]byte, error) {
  type OneScheduleSpecProxy OneScheduleSpec
  return json.Marshal(struct {
    *OneScheduleSpecProxy
    DurationMinutes *int64 `json:"durationMinutes,omitempty"`
    StartTimeHours *int64 `json:"startTimeHours,omitempty"`
    StartTimeMinutes *int64 `json:"startTimeMinutes,omitempty"`
  }{
    OneScheduleSpecProxy : (*OneScheduleSpecProxy)(p),
    DurationMinutes : p.DurationMinutes,
    StartTimeHours : p.StartTimeHours,
    StartTimeMinutes : p.StartTimeMinutes,
  })
}

func NewOneScheduleSpec() *OneScheduleSpec {
  p := new(OneScheduleSpec)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.OneScheduleSpec"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.OneScheduleSpec"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
One to one mapping model.
*/
type OneToOneMapping struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  IdentityMappingType *IdentityMappingType `json:"identityMappingType,omitempty"`
  
  NfsMapping *NfsIdentitiesMapping `json:"nfsMapping,omitempty"`
  
  SmbMapping *SmbIdentitiesMapping `json:"smbMapping,omitempty"`
}


func NewOneToOneMapping() *OneToOneMapping {
  p := new(OneToOneMapping)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.OneToOneMapping"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.OneToOneMapping"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Per day tiering schedule model
*/
type PerDayScheduleSpec struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Day of the week for tiering schedule. 1 starting with sunday.
  */
  DayOfWeek *int64 `json:"dayOfWeek"`
  
  EndTime *ScheduleTime `json:"endTime,omitempty"`
  
  Schedules []OneScheduleSpec `json:"schedules,omitempty"`
  
  StartTime *ScheduleTime `json:"startTime,omitempty"`
}

func (p *PerDayScheduleSpec) MarshalJSON() ([]byte, error) {
  type PerDayScheduleSpecProxy PerDayScheduleSpec
  return json.Marshal(struct {
    *PerDayScheduleSpecProxy
    DayOfWeek *int64 `json:"dayOfWeek,omitempty"`
  }{
    PerDayScheduleSpecProxy : (*PerDayScheduleSpecProxy)(p),
    DayOfWeek : p.DayOfWeek,
  })
}

func NewPerDayScheduleSpec() *PerDayScheduleSpec {
  p := new(PerDayScheduleSpec)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.PerDayScheduleSpec"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.PerDayScheduleSpec"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Depicts the overall state of the policy. It is a readonly field.

The following policy status is supported
- STATUS_ENABLED: All the subpolicies in a policy are in an enabled state.
- STATUS_DISABLED: All the subpolicies in a policy are in a disabled state.
- STATUS_PARTIALLY_DISABLED: Some of the subpolicies in a policy are in a disabled state
  and other are in an enabled state.
*/
type PolicyStatus struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  A code that uniquely identifies a message.
  */
  Code *string `json:"code,omitempty"`
  /**
  The locale for the message description.
  */
  Locale *string `json:"locale,omitempty"`
  
  Message *string `json:"message,omitempty"`
  
  Severity *import3.MessageSeverity `json:"severity,omitempty"`
  
  Status *Status `json:"status,omitempty"`
}


func NewPolicyStatus() *PolicyStatus {
  p := new(PolicyStatus)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.PolicyStatus"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.PolicyStatus"}
  p.UnknownFields_ = map[string]interface{}{}

  p.Locale = new(string)
  *p.Locale = "en_US"


  return p
}




/**
Principle type of quota policy. The acceptable values are "USER" or "GROUP".
*/
type PrincipalType int

const(
  PRINCIPALTYPE_UNKNOWN PrincipalType = 0
  PRINCIPALTYPE_REDACTED PrincipalType = 1
  PRINCIPALTYPE_USER PrincipalType = 2
  PRINCIPALTYPE_GROUP PrincipalType = 3
)

// returns the name of the enum given an ordinal number
func (e *PrincipalType) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "USER",
    "GROUP",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *PrincipalType) index(name string) PrincipalType {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "USER",
    "GROUP",
  }
  for idx := range names {
    if names[idx] == name {
      return PrincipalType(idx)
    }
  }
  return PRINCIPALTYPE_UNKNOWN
}

func (e *PrincipalType) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for PrincipalType:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *PrincipalType) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e PrincipalType) Ref() *PrincipalType {
  return &e
}


/**
Validate principal type and value for the provided mount target.
*/
type PrincipalTypeByPrincipalValue struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  PrincipalType *PrincipalType `json:"principalType"`
  /**
  Principal value is the name of the user or group, assigned to the principal type.
  */
  PrincipalValue *string `json:"principalValue"`
  
  ProtocolType *ProtocolTypeMountTarget `json:"protocolType"`
}

func (p *PrincipalTypeByPrincipalValue) MarshalJSON() ([]byte, error) {
  type PrincipalTypeByPrincipalValueProxy PrincipalTypeByPrincipalValue
  return json.Marshal(struct {
    *PrincipalTypeByPrincipalValueProxy
    PrincipalType *PrincipalType `json:"principalType,omitempty"`
    PrincipalValue *string `json:"principalValue,omitempty"`
    ProtocolType *ProtocolTypeMountTarget `json:"protocolType,omitempty"`
  }{
    PrincipalTypeByPrincipalValueProxy : (*PrincipalTypeByPrincipalValueProxy)(p),
    PrincipalType : p.PrincipalType,
    PrincipalValue : p.PrincipalValue,
    ProtocolType : p.ProtocolType,
  })
}

func NewPrincipalTypeByPrincipalValue() *PrincipalTypeByPrincipalValue {
  p := new(PrincipalTypeByPrincipalValue)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.PrincipalTypeByPrincipalValue"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.PrincipalTypeByPrincipalValue"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST response for all response codes in api path /files/v4.0.a2/config/file-server/mount-targets/{mountTargetExtId}/principal-type Post operation
*/
type PrincipalTypeByPrincipalValueApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfPrincipalTypeByPrincipalValueApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewPrincipalTypeByPrincipalValueApiResponse() *PrincipalTypeByPrincipalValueApiResponse {
  p := new(PrincipalTypeByPrincipalValueApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.PrincipalTypeByPrincipalValueApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.PrincipalTypeByPrincipalValueApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *PrincipalTypeByPrincipalValueApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *PrincipalTypeByPrincipalValueApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfPrincipalTypeByPrincipalValueApiResponseData()
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
This contains source and target configuration belonging to a sub policy.
*/
type ProtectedEntity struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  The {extId} of the mount target. Example:9c1e537d-6777-4c22-5d41-ddd0c3337aa9.
  */
  SourceMountTargetExtId *string `json:"sourceMountTargetExtId,omitempty"`
  /**
  Path of the nested mount target.
  */
  SourceMountTargetPath *string `json:"sourceMountTargetPath,omitempty"`
  /**
  The {extId} of the mount target. Example:9c1e537d-6777-4c22-5d41-ddd0c3337aa9.
  */
  TargetMountTargetExtId *string `json:"targetMountTargetExtId,omitempty"`
}


func NewProtectedEntity() *ProtectedEntity {
  p := new(ProtectedEntity)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.ProtectedEntity"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.ProtectedEntity"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Protected file server pair model.
*/
type ProtectedFileServerPair struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  
  FileServerState *FileServerState `json:"fileServerState,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import1.ApiLink `json:"links,omitempty"`
  /**
  The {extId} of the owning file server. This is a read-only field
  */
  OwningFileServerExtId *string `json:"owningFileServerExtId,omitempty"`
  /**
  The {extId} of the data protection policy.
  */
  PolicyExtIds []string `json:"policyExtIds,omitempty"`
  
  PrimaryAdCredential *Credential `json:"primaryAdCredential,omitempty"`
  /**
  Use a specific domain controller for the join-domain operation in a multi DC active directory setup. By default, AFS discovers a site local domain controller for join-domain operation. The preferred domain controller cannot be an IP address. It has to be FQDN (example: dc_name.dom.companyname.com)
  */
  PrimaryAdPreferredDomainController *string `json:"primaryAdPreferredDomainController,omitempty"`
  
  PrimaryDns *Dns `json:"primaryDns,omitempty"`
  /**
  The {extId} of the primary file server. This is a read-only field.
  */
  PrimaryFileServerExtId *string `json:"primaryFileServerExtId,omitempty"`
  
  Status *ActionStatus `json:"status,omitempty"`
  /**
  The {extId} of the target file server. This is a read-only field.
  */
  TargetFileServerExtId *string `json:"targetFileServerExtId,omitempty"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
}


func NewProtectedFileServerPair() *ProtectedFileServerPair {
  p := new(ProtectedFileServerPair)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.ProtectedFileServerPair"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.ProtectedFileServerPair"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST response for all response codes in api path /files/v4.0.a2/config/file-server/data-protection/protected-pairs/{protectedPairExtId} Get operation
*/
type ProtectedPairsApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfProtectedPairsApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewProtectedPairsApiResponse() *ProtectedPairsApiResponse {
  p := new(ProtectedPairsApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.ProtectedPairsApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.ProtectedPairsApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *ProtectedPairsApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *ProtectedPairsApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfProtectedPairsApiResponseData()
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
REST response for all response codes in api path /files/v4.0.a2/config/file-server/data-protection/protected-pairs Get operation
*/
type ProtectedPairsListApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfProtectedPairsListApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewProtectedPairsListApiResponse() *ProtectedPairsListApiResponse {
  p := new(ProtectedPairsListApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.ProtectedPairsListApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.ProtectedPairsListApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *ProtectedPairsListApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *ProtectedPairsListApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfProtectedPairsListApiResponseData()
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
Protocol type for file server. The acceptable values are "SMB", "NFS", "SMB_NFS", "NONE".
*/
type ProtocolType int

const(
  PROTOCOLTYPE_UNKNOWN ProtocolType = 0
  PROTOCOLTYPE_REDACTED ProtocolType = 1
  PROTOCOLTYPE_SMB ProtocolType = 2
  PROTOCOLTYPE_NFS ProtocolType = 3
  PROTOCOLTYPE_SMB_NFS ProtocolType = 4
  PROTOCOLTYPE_NONE ProtocolType = 5
)

// returns the name of the enum given an ordinal number
func (e *ProtocolType) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "SMB",
    "NFS",
    "SMB_NFS",
    "NONE",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *ProtocolType) index(name string) ProtocolType {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "SMB",
    "NFS",
    "SMB_NFS",
    "NONE",
  }
  for idx := range names {
    if names[idx] == name {
      return ProtocolType(idx)
    }
  }
  return PROTOCOLTYPE_UNKNOWN
}

func (e *ProtocolType) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for ProtocolType:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *ProtocolType) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e ProtocolType) Ref() *ProtocolType {
  return &e
}


/**
Protocol type of the mount target. The acceptable values are "SMB", "NFS", "NONE", "INCOMPATIBLE".
*/
type ProtocolTypeMountTarget int

const(
  PROTOCOLTYPEMOUNTTARGET_UNKNOWN ProtocolTypeMountTarget = 0
  PROTOCOLTYPEMOUNTTARGET_REDACTED ProtocolTypeMountTarget = 1
  PROTOCOLTYPEMOUNTTARGET_SMB ProtocolTypeMountTarget = 2
  PROTOCOLTYPEMOUNTTARGET_NFS ProtocolTypeMountTarget = 3
  PROTOCOLTYPEMOUNTTARGET_NONE ProtocolTypeMountTarget = 4
  PROTOCOLTYPEMOUNTTARGET_INCOMPATIBLE ProtocolTypeMountTarget = 5
)

// returns the name of the enum given an ordinal number
func (e *ProtocolTypeMountTarget) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "SMB",
    "NFS",
    "NONE",
    "INCOMPATIBLE",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *ProtocolTypeMountTarget) index(name string) ProtocolTypeMountTarget {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "SMB",
    "NFS",
    "NONE",
    "INCOMPATIBLE",
  }
  for idx := range names {
    if names[idx] == name {
      return ProtocolTypeMountTarget(idx)
    }
  }
  return PROTOCOLTYPEMOUNTTARGET_UNKNOWN
}

func (e *ProtocolTypeMountTarget) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for ProtocolTypeMountTarget:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *ProtocolTypeMountTarget) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e ProtocolTypeMountTarget) Ref() *ProtocolTypeMountTarget {
  return &e
}


/**
REST response for all response codes in api path /files/v4.0.a2/config/file-server/$actions/quarantine-infected-files Post operation
*/
type QuarantineInfectedFilesApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfQuarantineInfectedFilesApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewQuarantineInfectedFilesApiResponse() *QuarantineInfectedFilesApiResponse {
  p := new(QuarantineInfectedFilesApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.QuarantineInfectedFilesApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.QuarantineInfectedFilesApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *QuarantineInfectedFilesApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *QuarantineInfectedFilesApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfQuarantineInfectedFilesApiResponseData()
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
Quota policy email configuration
*/
type QuotaEmailConfig struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Email content
  */
  Content *string `json:"content,omitempty"`
  /**
  Email subject
  */
  Subject *string `json:"subject,omitempty"`
}


func NewQuotaEmailConfig() *QuotaEmailConfig {
  p := new(QuotaEmailConfig)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.QuotaEmailConfig"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.QuotaEmailConfig"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Quota policy model.
*/
type QuotaPolicy struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Enable email notifications for the user or group defined in the principal type. Notification will be sent, only if the user or groups are close to the quota provided. This field will be deprecated.
  */
  EnableNotifications *bool `json:"enableNotifications,omitempty"`
  
  EnforcementType *EnforcementType `json:"enforcementType"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  Enable email notifications for the user or group defined in the principal type. Notification will be sent, only if the user or groups are close to the quota provided.
  */
  IsEnableNotifications *bool `json:"isEnableNotifications,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import1.ApiLink `json:"links,omitempty"`
  /**
  The {extId} of the mount target. Example:9c1e537d-6777-4c22-5d41-ddd0c3337aa9.
  */
  MountTargetExtId *string `json:"mountTargetExtId,omitempty"`
  /**
  List of recipients emails.
  */
  NotificationRecipients []string `json:"notificationRecipients,omitempty"`
  
  PrincipalType *PrincipalType `json:"principalType"`
  /**
  Principal value is the name of the user or group, assigned to the principal type.
  */
  PrincipalValue *string `json:"principalValue"`
  /**
  Quota size in bytes.
  */
  SizeInBytes *int64 `json:"sizeInBytes"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
}

func (p *QuotaPolicy) MarshalJSON() ([]byte, error) {
  type QuotaPolicyProxy QuotaPolicy
  return json.Marshal(struct {
    *QuotaPolicyProxy
    EnforcementType *EnforcementType `json:"enforcementType,omitempty"`
    PrincipalType *PrincipalType `json:"principalType,omitempty"`
    PrincipalValue *string `json:"principalValue,omitempty"`
    SizeInBytes *int64 `json:"sizeInBytes,omitempty"`
  }{
    QuotaPolicyProxy : (*QuotaPolicyProxy)(p),
    EnforcementType : p.EnforcementType,
    PrincipalType : p.PrincipalType,
    PrincipalValue : p.PrincipalValue,
    SizeInBytes : p.SizeInBytes,
  })
}

func NewQuotaPolicy() *QuotaPolicy {
  p := new(QuotaPolicy)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.QuotaPolicy"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.QuotaPolicy"}
  p.UnknownFields_ = map[string]interface{}{}

  p.EnableNotifications = new(bool)
  *p.EnableNotifications = true
  p.IsEnableNotifications = new(bool)
  *p.IsEnableNotifications = true
  p.SizeInBytes = new(int64)
  *p.SizeInBytes = 1


  return p
}




/**
REST response for all response codes in api path /files/v4.0.a2/config/file-server/mount-targets/{mountTargetExtId}/quota-policies/{quotaPolicyExtId} Get operation
*/
type QuotaPolicyApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfQuotaPolicyApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewQuotaPolicyApiResponse() *QuotaPolicyApiResponse {
  p := new(QuotaPolicyApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.QuotaPolicyApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.QuotaPolicyApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *QuotaPolicyApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *QuotaPolicyApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfQuotaPolicyApiResponseData()
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
REST response for all response codes in api path /files/v4.0.a2/config/file-server/mount-targets/{mountTargetExtId}/quota-policies Get operation
*/
type QuotaPolicyListApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfQuotaPolicyListApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewQuotaPolicyListApiResponse() *QuotaPolicyListApiResponse {
  p := new(QuotaPolicyListApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.QuotaPolicyListApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.QuotaPolicyListApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *QuotaPolicyListApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *QuotaPolicyListApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfQuotaPolicyListApiResponseData()
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
Ransomware configuration information.
*/
type RansomwareConfig struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Enable / disable ransomware protection.
  */
  RansomwareEnabled *bool `json:"ransomwareEnabled,omitempty"`
  /**
  Ransomware file patterns. Ex - *.txt, ?.db
  */
  RansomwarePatterns []string `json:"ransomwarePatterns,omitempty"`
}


func NewRansomwareConfig() *RansomwareConfig {
  p := new(RansomwareConfig)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.RansomwareConfig"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.RansomwareConfig"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST response for all response codes in api path /files/v4.0.a2/config/file-server/ransomware-config Get operation
*/
type RansomwareConfigApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfRansomwareConfigApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewRansomwareConfigApiResponse() *RansomwareConfigApiResponse {
  p := new(RansomwareConfigApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.RansomwareConfigApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.RansomwareConfigApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *RansomwareConfigApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *RansomwareConfigApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfRansomwareConfigApiResponseData()
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
File server recommendation description.
*/
type Recommendation struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  The {extId} of the file server recommendation. This is a read-only field.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  The recommended number of file server VMs.
  */
  NumberOfExpectedVms *int `json:"numberOfExpectedVms,omitempty"`
  
  RecommendationType *RecommendationType `json:"recommendationType,omitempty"`
}


func NewRecommendation() *Recommendation {
  p := new(Recommendation)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.Recommendation"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.Recommendation"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST response for all response codes in api path /files/v4.0.a2/config/file-server/recommendations Get operation
*/
type RecommendationApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfRecommendationApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewRecommendationApiResponse() *RecommendationApiResponse {
  p := new(RecommendationApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.RecommendationApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.RecommendationApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *RecommendationApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *RecommendationApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfRecommendationApiResponseData()
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
File server recommendation type details. The acceptable values are "SCALEUP", "SCALEOUT", "REBALANCE".
*/
type RecommendationType int

const(
  RECOMMENDATIONTYPE_UNKNOWN RecommendationType = 0
  RECOMMENDATIONTYPE_REDACTED RecommendationType = 1
  RECOMMENDATIONTYPE_SCALEUP RecommendationType = 2
  RECOMMENDATIONTYPE_SCALEOUT RecommendationType = 3
  RECOMMENDATIONTYPE_REBALANCE RecommendationType = 4
)

// returns the name of the enum given an ordinal number
func (e *RecommendationType) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "SCALEUP",
    "SCALEOUT",
    "REBALANCE",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *RecommendationType) index(name string) RecommendationType {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "SCALEUP",
    "SCALEOUT",
    "REBALANCE",
  }
  for idx := range names {
    if names[idx] == name {
      return RecommendationType(idx)
    }
  }
  return RECOMMENDATIONTYPE_UNKNOWN
}

func (e *RecommendationType) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for RecommendationType:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *RecommendationType) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e RecommendationType) Ref() *RecommendationType {
  return &e
}


/**
REST response for all response codes in api path /files/v4.0.a2/config/file-server/$actions/remove-dns-entries Post operation
*/
type RemoveDnsEntriesApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfRemoveDnsEntriesApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewRemoveDnsEntriesApiResponse() *RemoveDnsEntriesApiResponse {
  p := new(RemoveDnsEntriesApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.RemoveDnsEntriesApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.RemoveDnsEntriesApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *RemoveDnsEntriesApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *RemoveDnsEntriesApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfRemoveDnsEntriesApiResponseData()
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
REST response for all response codes in api path /files/v4.0.a2/config/file-server/$actions/remove-infected-files Post operation
*/
type RemoveInfectedFilesApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfRemoveInfectedFilesApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewRemoveInfectedFilesApiResponse() *RemoveInfectedFilesApiResponse {
  p := new(RemoveInfectedFilesApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.RemoveInfectedFilesApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.RemoveInfectedFilesApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *RemoveInfectedFilesApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *RemoveInfectedFilesApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfRemoveInfectedFilesApiResponseData()
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
Contains details of replicated job.
*/
type ReplicationJob struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  The time difference between the start and the end time of the job.
  */
  Duration *string `json:"duration,omitempty"`
  /**
  The time difference between the start and the end time of the job in ISO format.
  */
  DurationISO8601 *string `json:"durationISO8601,omitempty"`
  /**
  End time for the replication job.
  */
  EndTime *string `json:"endTime,omitempty"`
  /**
  End time for the replication job in ISO format.
  */
  EndTimeISO8601 *string `json:"endTimeISO8601,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import1.ApiLink `json:"links,omitempty"`
  /**
  The {extId} of the mount target. Example:9c1e537d-6777-4c22-5d41-ddd0c3337aa9.
  */
  MountTargetExtId *string `json:"mountTargetExtId,omitempty"`
  /**
  ExtId of the policy.
  */
  PolicyExtId *string `json:"policyExtId,omitempty"`
  /**
  Start time for the replication job.
  */
  StartTime *string `json:"startTime,omitempty"`
  /**
  Start time for the replication job in ISO format.
  */
  StartTimeISO8601 *string `json:"startTimeISO8601,omitempty"`
  
  Status *JobStatus `json:"status,omitempty"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
}


func NewReplicationJob() *ReplicationJob {
  p := new(ReplicationJob)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.ReplicationJob"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.ReplicationJob"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST response for all response codes in api path /files/v4.0.a2/config/file-server/data-replication-jobs/{replicationJobExtId} Get operation
*/
type ReplicationJobApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfReplicationJobApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewReplicationJobApiResponse() *ReplicationJobApiResponse {
  p := new(ReplicationJobApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.ReplicationJobApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.ReplicationJobApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *ReplicationJobApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *ReplicationJobApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfReplicationJobApiResponseData()
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
REST response for all response codes in api path /files/v4.0.a2/config/file-server/data-replication-jobs Get operation
*/
type ReplicationJobListApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfReplicationJobListApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewReplicationJobListApiResponse() *ReplicationJobListApiResponse {
  p := new(ReplicationJobListApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.ReplicationJobListApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.ReplicationJobListApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *ReplicationJobListApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *ReplicationJobListApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfReplicationJobListApiResponseData()
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
Replication summary as per the last running jobs of the targeted mount targets. It is a readonly field.
*/
type ReplicationSummary struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  A code that uniquely identifies a message.
  */
  Code *string `json:"code,omitempty"`
  /**
  The locale for the message description.
  */
  Locale *string `json:"locale,omitempty"`
  
  Message *string `json:"message,omitempty"`
  
  Severity *import3.MessageSeverity `json:"severity,omitempty"`
  
  Status *Summary `json:"status,omitempty"`
}


func NewReplicationSummary() *ReplicationSummary {
  p := new(ReplicationSummary)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.ReplicationSummary"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.ReplicationSummary"}
  p.UnknownFields_ = map[string]interface{}{}

  p.Locale = new(string)
  *p.Locale = "en_US"


  return p
}




/**
REST response for all response codes in api path /files/v4.0.a2/config/file-server/$actions/rescan-infected-files Post operation
*/
type RescanInfectedFilesApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfRescanInfectedFilesApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewRescanInfectedFilesApiResponse() *RescanInfectedFilesApiResponse {
  p := new(RescanInfectedFilesApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.RescanInfectedFilesApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.RescanInfectedFilesApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *RescanInfectedFilesApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *RescanInfectedFilesApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfRescanInfectedFilesApiResponseData()
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
REST response for all response codes in api path /files/v4.0.a2/config/file-server/$actions/reset-infected-files Post operation
*/
type ResetInfectedFilesApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfResetInfectedFilesApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewResetInfectedFilesApiResponse() *ResetInfectedFilesApiResponse {
  p := new(ResetInfectedFilesApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.ResetInfectedFilesApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.ResetInfectedFilesApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *ResetInfectedFilesApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *ResetInfectedFilesApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfResetInfectedFilesApiResponseData()
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
REST response for all response codes in api path /files/v4.0.a2/config/file-server/alerts/$actions/resolve Post operation
*/
type ResolveAlertsApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfResolveAlertsApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewResolveAlertsApiResponse() *ResolveAlertsApiResponse {
  p := new(ResolveAlertsApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.ResolveAlertsApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.ResolveAlertsApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *ResolveAlertsApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *ResolveAlertsApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfResolveAlertsApiResponseData()
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
Response for resolve alert
*/
type ResolveAlertsResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Alert UUID
  */
  Id *string `json:"id,omitempty"`
  /**
  Message corresponding to alert resolve request.
  */
  Message *string `json:"message,omitempty"`
  /**
  This field indicates if the alert is resolved.
  */
  Successful *bool `json:"successful,omitempty"`
}


func NewResolveAlertsResponse() *ResolveAlertsResponse {
  p := new(ResolveAlertsResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.ResolveAlertsResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.ResolveAlertsResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Generic CRUD response details
*/
type Response struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  Generic CRUD response details
  */
  IsStatus *bool `json:"isStatus,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import1.ApiLink `json:"links,omitempty"`
  /**
  Generic response message
  */
  Message *string `json:"message,omitempty"`
  /**
  Generic CRUD response details. This field will be deprecated.
  */
  Status *bool `json:"status,omitempty"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
}


func NewResponse() *Response {
  p := new(Response)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.Response"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.Response"}
  p.UnknownFields_ = map[string]interface{}{}

  p.IsStatus = new(bool)
  *p.IsStatus = false
  p.Status = new(bool)
  *p.Status = false


  return p
}




/**
Rule based mapping model.
*/
type RuleBasedMapping struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  NfsWithNoMappings []NfsIdentitiesMapping `json:"nfsWithNoMappings,omitempty"`
  
  SmbWithNoMappings []SmbIdentitiesMapping `json:"smbWithNoMappings,omitempty"`
  
  TemplateMapping *TemplateMappingType `json:"templateMapping,omitempty"`
}


func NewRuleBasedMapping() *RuleBasedMapping {
  p := new(RuleBasedMapping)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.RuleBasedMapping"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.RuleBasedMapping"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Scale out of a file server.
*/
type ScaleOutSpec struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  List of external network IP addresses.
  */
  ExternalIpAddresses []import3.IPAddress `json:"externalIpAddresses,omitempty"`
  /**
  List of internal network IP addresses.
  */
  InternalIpAdresses []import3.IPAddress `json:"internalIpAdresses,omitempty"`
  /**
  The recommended number of file server VMs.
  */
  NumberOfExpectedVms *int `json:"numberOfExpectedVms,omitempty"`
}


func NewScaleOutSpec() *ScaleOutSpec {
  p := new(ScaleOutSpec)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.ScaleOutSpec"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.ScaleOutSpec"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Scale up of a file server.
*/
type ScaleUpSpec struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Memory associated with the file server VM.
  */
  MemoryGib *int64 `json:"memoryGib,omitempty"`
  /**
  CPU associated with the file server VM.
  */
  Vcpus *int64 `json:"vcpus,omitempty"`
}


func NewScaleUpSpec() *ScaleUpSpec {
  p := new(ScaleUpSpec)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.ScaleUpSpec"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.ScaleUpSpec"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
This contains the schedule for all the policy. This has the start time and frequency information of the policy.
*/
type Schedule struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  RPO (Integer values are multiplied by schedule type to get the next scheduled repetition time).
  */
  Multiple *int `json:"multiple"`
  
  ScheduleType *ScheduleType `json:"scheduleType"`
  /**
  Time of the day, the policy will be started. If the value not set, the policy will be applicable immediately.
  */
  StartTime *string `json:"startTime,omitempty"`
  /**
  Time of the day, the policy will be started. If the value not set, the policy will be applicable immediately.
  */
  StartTimeISO8601 *string `json:"startTimeISO8601,omitempty"`
}

func (p *Schedule) MarshalJSON() ([]byte, error) {
  type ScheduleProxy Schedule
  return json.Marshal(struct {
    *ScheduleProxy
    Multiple *int `json:"multiple,omitempty"`
    ScheduleType *ScheduleType `json:"scheduleType,omitempty"`
  }{
    ScheduleProxy : (*ScheduleProxy)(p),
    Multiple : p.Multiple,
    ScheduleType : p.ScheduleType,
  })
}

func NewSchedule() *Schedule {
  p := new(Schedule)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.Schedule"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.Schedule"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Time of the day for starting or ending the tier schedule.
*/
type ScheduleTime struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Hour of the day
  */
  Hour *int64 `json:"hour"`
  /**
  Minute value
  */
  Minute *int64 `json:"minute"`
}

func (p *ScheduleTime) MarshalJSON() ([]byte, error) {
  type ScheduleTimeProxy ScheduleTime
  return json.Marshal(struct {
    *ScheduleTimeProxy
    Hour *int64 `json:"hour,omitempty"`
    Minute *int64 `json:"minute,omitempty"`
  }{
    ScheduleTimeProxy : (*ScheduleTimeProxy)(p),
    Hour : p.Hour,
    Minute : p.Minute,
  })
}

func NewScheduleTime() *ScheduleTime {
  p := new(ScheduleTime)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.ScheduleTime"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.ScheduleTime"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Schedule type defines the following values

- SCHEDULETYPE_YEARLY: To repeat the schedule on an `Yearly` basis.
- SCHEDULETYPE_MONTHLY: To repeat the schedule on- a `Monthly` basis.
- SCHEDULETYPE_WEEKLY: To repeat the schedule on a `Weekly` basis.
- SCHEDULETYPE_DAILY: To repeat the schedule on a `Daily` basis.
- SCHEDULETYPE_HOURLY: To repeat the schedule on an `Hourly` basis.
- SCHEDULETYPE_MINUTELY: To repeat the schedule on a `Minute` basis.
- SCHEDULETYPE_SECONDLY: To repeat the schedule on a `Second` basis.
*/
type ScheduleType int

const(
  SCHEDULETYPE_UNKNOWN ScheduleType = 0
  SCHEDULETYPE_REDACTED ScheduleType = 1
  SCHEDULETYPE_SCHEDULETYPE_YEARLY ScheduleType = 2
  SCHEDULETYPE_SCHEDULETYPE_MONTHLY ScheduleType = 3
  SCHEDULETYPE_SCHEDULETYPE_WEEKLY ScheduleType = 4
  SCHEDULETYPE_SCHEDULETYPE_DAILY ScheduleType = 5
  SCHEDULETYPE_SCHEDULETYPE_HOURLY ScheduleType = 6
  SCHEDULETYPE_SCHEDULETYPE_MINUTELY ScheduleType = 7
  SCHEDULETYPE_SCHEDULETYPE_SECONDLY ScheduleType = 8
)

// returns the name of the enum given an ordinal number
func (e *ScheduleType) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "SCHEDULETYPE_YEARLY",
    "SCHEDULETYPE_MONTHLY",
    "SCHEDULETYPE_WEEKLY",
    "SCHEDULETYPE_DAILY",
    "SCHEDULETYPE_HOURLY",
    "SCHEDULETYPE_MINUTELY",
    "SCHEDULETYPE_SECONDLY",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *ScheduleType) index(name string) ScheduleType {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "SCHEDULETYPE_YEARLY",
    "SCHEDULETYPE_MONTHLY",
    "SCHEDULETYPE_WEEKLY",
    "SCHEDULETYPE_DAILY",
    "SCHEDULETYPE_HOURLY",
    "SCHEDULETYPE_MINUTELY",
    "SCHEDULETYPE_SECONDLY",
  }
  for idx := range names {
    if names[idx] == name {
      return ScheduleType(idx)
    }
  }
  return SCHEDULETYPE_UNKNOWN
}

func (e *ScheduleType) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for ScheduleType:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *ScheduleType) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e ScheduleType) Ref() *ScheduleType {
  return &e
}


/**
Search mapping type. The acceptable values are "SMB_TO_NFS" or "NFS_TO_SMB".
*/
type SearchMappingType int

const(
  SEARCHMAPPINGTYPE_UNKNOWN SearchMappingType = 0
  SEARCHMAPPINGTYPE_REDACTED SearchMappingType = 1
  SEARCHMAPPINGTYPE_SMB_TO_NFS SearchMappingType = 2
  SEARCHMAPPINGTYPE_NFS_TO_SMB SearchMappingType = 3
)

// returns the name of the enum given an ordinal number
func (e *SearchMappingType) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "SMB_TO_NFS",
    "NFS_TO_SMB",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *SearchMappingType) index(name string) SearchMappingType {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "SMB_TO_NFS",
    "NFS_TO_SMB",
  }
  for idx := range names {
    if names[idx] == name {
      return SearchMappingType(idx)
    }
  }
  return SEARCHMAPPINGTYPE_UNKNOWN
}

func (e *SearchMappingType) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for SearchMappingType:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *SearchMappingType) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e SearchMappingType) Ref() *SearchMappingType {
  return &e
}


/**
REST response for all response codes in api path /files/v4.0.a2/config/file-server/$actions/search-user-mapping Post operation
*/
type SearchUserMappingPOSTResponses struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfSearchUserMappingPOSTResponsesData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewSearchUserMappingPOSTResponses() *SearchUserMappingPOSTResponses {
  p := new(SearchUserMappingPOSTResponses)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.SearchUserMappingPOSTResponses"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.SearchUserMappingPOSTResponses"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *SearchUserMappingPOSTResponses) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *SearchUserMappingPOSTResponses) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfSearchUserMappingPOSTResponsesData()
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
Search user mapping model
*/
type SearchUserMappingRequestSpec struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  ActionType *IdentityActionType `json:"actionType,omitempty"`
  
  MappingType *IdentityMappingType `json:"mappingType,omitempty"`
  /**
  Name or Id for searching the user mapping.
  */
  NameIds []string `json:"nameIds"`
  
  SearchMappingType *SearchMappingType `json:"searchMappingType,omitempty"`
}

func (p *SearchUserMappingRequestSpec) MarshalJSON() ([]byte, error) {
  type SearchUserMappingRequestSpecProxy SearchUserMappingRequestSpec
  return json.Marshal(struct {
    *SearchUserMappingRequestSpecProxy
    NameIds []string `json:"nameIds,omitempty"`
  }{
    SearchUserMappingRequestSpecProxy : (*SearchUserMappingRequestSpecProxy)(p),
    NameIds : p.NameIds,
  })
}

func NewSearchUserMappingRequestSpec() *SearchUserMappingRequestSpec {
  p := new(SearchUserMappingRequestSpec)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.SearchUserMappingRequestSpec"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.SearchUserMappingRequestSpec"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Details of the queried user mapping.
*/
type SearchUserMappingResponseSpec struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  MappedEntry []SearchUserMappingRequestSpec `json:"mappedEntry,omitempty"`
  
  OptionType *IdentityMappingOptionType `json:"optionType,omitempty"`
  
  SearchedEntry *SearchUserMappingRequestSpec `json:"searchedEntry,omitempty"`
}


func NewSearchUserMappingResponseSpec() *SearchUserMappingResponseSpec {
  p := new(SearchUserMappingResponseSpec)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.SearchUserMappingResponseSpec"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.SearchUserMappingResponseSpec"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
SMB mapping details
*/
type SmbIdentitiesMapping struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  ActionType *IdentityActionType `json:"actionType,omitempty"`
  /**
  SMB group Id
  */
  GroupId *int `json:"groupId,omitempty"`
  
  MappingType *IdentityMappingType `json:"mappingType,omitempty"`
  /**
  SMB user name
  */
  Name *string `json:"name,omitempty"`
  /**
  SMB user Id
  */
  UserId *int `json:"userId,omitempty"`
}


func NewSmbIdentitiesMapping() *SmbIdentitiesMapping {
  p := new(SmbIdentitiesMapping)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.SmbIdentitiesMapping"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.SmbIdentitiesMapping"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
SMB protocol object.
*/
type SmbProtocol struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Flag to enable access based enumeration. This field will be deprecated.
  */
  EnableAccessBasedEnumeration *bool `json:"enableAccessBasedEnumeration,omitempty"`
  /**
  Flag to enable continuous availability feature for SMB mount targets. This field will be deprecated.
  */
  EnableCa *bool `json:"enableCa,omitempty"`
  /**
  Flag to enable SMB3 encryption. This field will be deprecated.
  */
  EnableSmb3Encryption *bool `json:"enableSmb3Encryption,omitempty"`
  /**
  Flag to enable access based enumeration.
  */
  IsEnableAccessBasedEnumeration *bool `json:"isEnableAccessBasedEnumeration,omitempty"`
  /**
  Flag to enable continuous availability feature for SMB mount targets.
  */
  IsEnableCa *bool `json:"isEnableCa,omitempty"`
  /**
  Flag to enable SMB3 encryption.
  */
  IsEnableSmb3Encryption *bool `json:"isEnableSmb3Encryption,omitempty"`
}


func NewSmbProtocol() *SmbProtocol {
  p := new(SmbProtocol)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.SmbProtocol"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.SmbProtocol"}
  p.UnknownFields_ = map[string]interface{}{}

  p.EnableAccessBasedEnumeration = new(bool)
  *p.EnableAccessBasedEnumeration = false
  p.EnableCa = new(bool)
  *p.EnableCa = false
  p.EnableSmb3Encryption = new(bool)
  *p.EnableSmb3Encryption = false
  p.IsEnableAccessBasedEnumeration = new(bool)
  *p.IsEnableAccessBasedEnumeration = false
  p.IsEnableCa = new(bool)
  *p.IsEnableCa = false
  p.IsEnableSmb3Encryption = new(bool)
  *p.IsEnableSmb3Encryption = false


  return p
}




/**
Mount target snapshot model.
*/
type Snapshot struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Mount target snapshot created time.
  */
  CreateTime *int64 `json:"createTime,omitempty"`
  
  Creator *SnapshotCreator `json:"creator,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  The {extId} of the file server. This is a read-only field.
  */
  FileServerExtId *string `json:"fileServerExtId,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import1.ApiLink `json:"links,omitempty"`
  /**
  The {extId} of the mount target. Example:9c1e537d-6777-4c22-5d41-ddd0c3337aa9.
  */
  MountTargetExtId *string `json:"mountTargetExtId,omitempty"`
  /**
  Mount target snapshot name.
  */
  Name *string `json:"name,omitempty"`
  /**
  Mount target snapshot reclaimable space in bytes.
  */
  ReclaimableSpaceBytes *int64 `json:"reclaimableSpaceBytes,omitempty"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
  /**
  Mount target snapshot total space in bytes.
  */
  TotalSpaceBytes *int64 `json:"totalSpaceBytes,omitempty"`
}


func NewSnapshot() *Snapshot {
  p := new(Snapshot)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.Snapshot"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.Snapshot"}
  p.UnknownFields_ = map[string]interface{}{}

  p.ReclaimableSpaceBytes = new(int64)
  *p.ReclaimableSpaceBytes = 0
  p.TotalSpaceBytes = new(int64)
  *p.TotalSpaceBytes = 0


  return p
}




/**
Mount target snapshot creator type. The acceptable values are "SSR_SNAPSHOT", "BACKUP_SNAPSHOT", "REPLICATION_SNAPSHOT".
*/
type SnapshotCreator int

const(
  SNAPSHOTCREATOR_UNKNOWN SnapshotCreator = 0
  SNAPSHOTCREATOR_REDACTED SnapshotCreator = 1
  SNAPSHOTCREATOR_SSR_SNAPSHOT SnapshotCreator = 2
  SNAPSHOTCREATOR_BACKUP_SNAPSHOT SnapshotCreator = 3
  SNAPSHOTCREATOR_REPLICATION_SNAPSHOT SnapshotCreator = 4
)

// returns the name of the enum given an ordinal number
func (e *SnapshotCreator) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "SSR_SNAPSHOT",
    "BACKUP_SNAPSHOT",
    "REPLICATION_SNAPSHOT",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *SnapshotCreator) index(name string) SnapshotCreator {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "SSR_SNAPSHOT",
    "BACKUP_SNAPSHOT",
    "REPLICATION_SNAPSHOT",
  }
  for idx := range names {
    if names[idx] == name {
      return SnapshotCreator(idx)
    }
  }
  return SNAPSHOTCREATOR_UNKNOWN
}

func (e *SnapshotCreator) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for SnapshotCreator:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *SnapshotCreator) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e SnapshotCreator) Ref() *SnapshotCreator {
  return &e
}


/**
SSR snapshot schedule schedule type. The acceptable values are "HOURLY", "DAILY", "WEEKLY", "MONTHLY". If user passes this value as WEEKLY, then daysOfWeek becomes a required field. If user passes this value as MONTHLY, then daysOfMonth becomes a required field.
*/
type SnapshotScheduleType int

const(
  SNAPSHOTSCHEDULETYPE_UNKNOWN SnapshotScheduleType = 0
  SNAPSHOTSCHEDULETYPE_REDACTED SnapshotScheduleType = 1
  SNAPSHOTSCHEDULETYPE_HOURLY SnapshotScheduleType = 2
  SNAPSHOTSCHEDULETYPE_DAILY SnapshotScheduleType = 3
  SNAPSHOTSCHEDULETYPE_WEEKLY SnapshotScheduleType = 4
  SNAPSHOTSCHEDULETYPE_MONTHLY SnapshotScheduleType = 5
)

// returns the name of the enum given an ordinal number
func (e *SnapshotScheduleType) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "HOURLY",
    "DAILY",
    "WEEKLY",
    "MONTHLY",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *SnapshotScheduleType) index(name string) SnapshotScheduleType {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "HOURLY",
    "DAILY",
    "WEEKLY",
    "MONTHLY",
  }
  for idx := range names {
    if names[idx] == name {
      return SnapshotScheduleType(idx)
    }
  }
  return SNAPSHOTSCHEDULETYPE_UNKNOWN
}

func (e *SnapshotScheduleType) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for SnapshotScheduleType:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *SnapshotScheduleType) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e SnapshotScheduleType) Ref() *SnapshotScheduleType {
  return &e
}


/**
An enum that represents the source share protocol type
*/
type SourceShareProtocolType int

const(
  SOURCESHAREPROTOCOLTYPE_UNKNOWN SourceShareProtocolType = 0
  SOURCESHAREPROTOCOLTYPE_REDACTED SourceShareProtocolType = 1
  SOURCESHAREPROTOCOLTYPE_SMB SourceShareProtocolType = 2
  SOURCESHAREPROTOCOLTYPE_NFS SourceShareProtocolType = 3
)

// returns the name of the enum given an ordinal number
func (e *SourceShareProtocolType) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "SMB",
    "NFS",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *SourceShareProtocolType) index(name string) SourceShareProtocolType {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "SMB",
    "NFS",
  }
  for idx := range names {
    if names[idx] == name {
      return SourceShareProtocolType(idx)
    }
  }
  return SOURCESHAREPROTOCOLTYPE_UNKNOWN
}

func (e *SourceShareProtocolType) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for SourceShareProtocolType:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *SourceShareProtocolType) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e SourceShareProtocolType) Ref() *SourceShareProtocolType {
  return &e
}


/**
Squash type of the mount target. The acceptable values are "ROOT_SQUASH", "ALL_SQUASH", "NONE".
*/
type SquashType int

const(
  SQUASHTYPE_UNKNOWN SquashType = 0
  SQUASHTYPE_REDACTED SquashType = 1
  SQUASHTYPE_ROOT_SQUASH SquashType = 2
  SQUASHTYPE_ALL_SQUASH SquashType = 3
  SQUASHTYPE_NONE SquashType = 4
)

// returns the name of the enum given an ordinal number
func (e *SquashType) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "ROOT_SQUASH",
    "ALL_SQUASH",
    "NONE",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *SquashType) index(name string) SquashType {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "ROOT_SQUASH",
    "ALL_SQUASH",
    "NONE",
  }
  for idx := range names {
    if names[idx] == name {
      return SquashType(idx)
    }
  }
  return SQUASHTYPE_UNKNOWN
}

func (e *SquashType) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for SquashType:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *SquashType) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e SquashType) Ref() *SquashType {
  return &e
}


/**
SSR snapshot policy model
*/
type SsrSnapshotPolicy struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  SSR snapshot policy description
  */
  Description *string `json:"description,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  Indicates if this is a default SSR snapshot policy
  */
  IsDefault *bool `json:"isDefault,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import1.ApiLink `json:"links,omitempty"`
  /**
  SSR snapshot policy name
  */
  Name *string `json:"name"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
}

func (p *SsrSnapshotPolicy) MarshalJSON() ([]byte, error) {
  type SsrSnapshotPolicyProxy SsrSnapshotPolicy
  return json.Marshal(struct {
    *SsrSnapshotPolicyProxy
    Name *string `json:"name,omitempty"`
  }{
    SsrSnapshotPolicyProxy : (*SsrSnapshotPolicyProxy)(p),
    Name : p.Name,
  })
}

func NewSsrSnapshotPolicy() *SsrSnapshotPolicy {
  p := new(SsrSnapshotPolicy)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.SsrSnapshotPolicy"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.SsrSnapshotPolicy"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
SSR snapshot schedule model
*/
type SsrSnapshotSchedule struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Number of days of month for snapshot to be taken.
  */
  DaysOfMonth []int `json:"daysOfMonth,omitempty"`
  /**
  Number of days of week for snapshot to be taken.
  */
  DaysOfWeek []int `json:"daysOfWeek,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import1.ApiLink `json:"links,omitempty"`
  /**
  The maximum number of snapshots to be kept locally.
  */
  LocalMaxSnapshots *int `json:"localMaxSnapshots"`
  /**
  Frequency of snapshots.
  */
  SnapshotScheduleFrequency *int `json:"snapshotScheduleFrequency"`
  
  SnapshotScheduleType *SnapshotScheduleType `json:"snapshotScheduleType"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
}

func (p *SsrSnapshotSchedule) MarshalJSON() ([]byte, error) {
  type SsrSnapshotScheduleProxy SsrSnapshotSchedule
  return json.Marshal(struct {
    *SsrSnapshotScheduleProxy
    LocalMaxSnapshots *int `json:"localMaxSnapshots,omitempty"`
    SnapshotScheduleFrequency *int `json:"snapshotScheduleFrequency,omitempty"`
    SnapshotScheduleType *SnapshotScheduleType `json:"snapshotScheduleType,omitempty"`
  }{
    SsrSnapshotScheduleProxy : (*SsrSnapshotScheduleProxy)(p),
    LocalMaxSnapshots : p.LocalMaxSnapshots,
    SnapshotScheduleFrequency : p.SnapshotScheduleFrequency,
    SnapshotScheduleType : p.SnapshotScheduleType,
  })
}

func NewSsrSnapshotSchedule() *SsrSnapshotSchedule {
  p := new(SsrSnapshotSchedule)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.SsrSnapshotSchedule"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.SsrSnapshotSchedule"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST response for all response codes in api path /files/v4.0.a2/config/file-server/ssr-snapshot-schedules/{ssrSnapshotScheduleExtId} Get operation
*/
type SsrSnapshotScheduleApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfSsrSnapshotScheduleApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewSsrSnapshotScheduleApiResponse() *SsrSnapshotScheduleApiResponse {
  p := new(SsrSnapshotScheduleApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.SsrSnapshotScheduleApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.SsrSnapshotScheduleApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *SsrSnapshotScheduleApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *SsrSnapshotScheduleApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfSsrSnapshotScheduleApiResponseData()
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
REST response for all response codes in api path /files/v4.0.a2/config/file-server/ssr-snapshot-schedules Get operation
*/
type SsrSnapshotScheduleListApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfSsrSnapshotScheduleListApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewSsrSnapshotScheduleListApiResponse() *SsrSnapshotScheduleListApiResponse {
  p := new(SsrSnapshotScheduleListApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.SsrSnapshotScheduleListApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.SsrSnapshotScheduleListApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *SsrSnapshotScheduleListApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *SsrSnapshotScheduleListApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfSsrSnapshotScheduleListApiResponseData()
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
Status for policy. The acceptable values are  `PROGRESS`, `PASSED`,`FAILED`, `PENDING`, `SUCCEEDED`, `ENABLED`, `DISABLED`, `PARTIALLY_DISABLED`.
*/
type Status int

const(
  STATUS_UNKNOWN Status = 0
  STATUS_REDACTED Status = 1
  STATUS_STATUS_PROGRESS Status = 2
  STATUS_STATUS_PASSED Status = 3
  STATUS_STATUS_FAILED Status = 4
  STATUS_STATUS_PENDING Status = 5
  STATUS_STATUS_ENABLED Status = 6
  STATUS_STATUS_DISABLED Status = 7
  STATUS_STATUS_PARTIALLY_DISABLED Status = 8
)

// returns the name of the enum given an ordinal number
func (e *Status) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "STATUS_PROGRESS",
    "STATUS_PASSED",
    "STATUS_FAILED",
    "STATUS_PENDING",
    "STATUS_ENABLED",
    "STATUS_DISABLED",
    "STATUS_PARTIALLY_DISABLED",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *Status) index(name string) Status {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "STATUS_PROGRESS",
    "STATUS_PASSED",
    "STATUS_FAILED",
    "STATUS_PENDING",
    "STATUS_ENABLED",
    "STATUS_DISABLED",
    "STATUS_PARTIALLY_DISABLED",
  }
  for idx := range names {
    if names[idx] == name {
      return Status(idx)
    }
  }
  return STATUS_UNKNOWN
}

func (e *Status) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for Status:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *Status) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e Status) Ref() *Status {
  return &e
}


/**
Status info details
*/
type StatusInfo struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import1.ApiLink `json:"links,omitempty"`
  
  Status *MigrationPlanStatus `json:"status"`
  /**
  Status message
  */
  StatusMsg *string `json:"statusMsg"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
}

func (p *StatusInfo) MarshalJSON() ([]byte, error) {
  type StatusInfoProxy StatusInfo
  return json.Marshal(struct {
    *StatusInfoProxy
    Status *MigrationPlanStatus `json:"status,omitempty"`
    StatusMsg *string `json:"statusMsg,omitempty"`
  }{
    StatusInfoProxy : (*StatusInfoProxy)(p),
    Status : p.Status,
    StatusMsg : p.StatusMsg,
  })
}

func NewStatusInfo() *StatusInfo {
  p := new(StatusInfo)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.StatusInfo"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.StatusInfo"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Status type. The acceptable values are "AVAILABLE" or "UNAVAILABLE".
*/
type StatusType int

const(
  STATUSTYPE_UNKNOWN StatusType = 0
  STATUSTYPE_REDACTED StatusType = 1
  STATUSTYPE_AVAILABLE StatusType = 2
  STATUSTYPE_UNAVAILABLE StatusType = 3
)

// returns the name of the enum given an ordinal number
func (e *StatusType) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "AVAILABLE",
    "UNAVAILABLE",
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
    "AVAILABLE",
    "UNAVAILABLE",
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
Subplan list details
*/
type SubPlanList struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import1.ApiLink `json:"links,omitempty"`
  
  SubPlanExtIds []string `json:"subPlanExtIds"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
}

func (p *SubPlanList) MarshalJSON() ([]byte, error) {
  type SubPlanListProxy SubPlanList
  return json.Marshal(struct {
    *SubPlanListProxy
    SubPlanExtIds []string `json:"subPlanExtIds,omitempty"`
  }{
    SubPlanListProxy : (*SubPlanListProxy)(p),
    SubPlanExtIds : p.SubPlanExtIds,
  })
}

func NewSubPlanList() *SubPlanList {
  p := new(SubPlanList)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.SubPlanList"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.SubPlanList"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
This contains list of sub policies. Sub policy is combination of sourceEntity, targetEntity and sourceTargetMap.
*/
type SubPolicy struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  ExtId *string `json:"extId,omitempty"`
  /**
  This field talks about if all the mount targets in file server are included in the policy. By default it is always true.If specific mount targets are selected then we should mark this field as false.
  */
  IncludeNewMountTargets *bool `json:"includeNewMountTargets,omitempty"`
  /**
  This flag indicates if this policy has been used for reverse replication.
  */
  IsReversePolicy *bool `json:"isReversePolicy,omitempty"`
  
  PolicyStatus *PolicyStatus `json:"policyStatus,omitempty"`
  /**
  This contains list of source target entity belongs to a sub policy.
  */
  ProtectedEntities []ProtectedEntity `json:"protectedEntities,omitempty"`
  
  ReplicationSummary *ReplicationSummary `json:"replicationSummary,omitempty"`
  /**
  This contains details of snapshot schedules. It contains the start time of the replication and frequency of the replication.
  */
  Schedules []Schedule `json:"schedules,omitempty"`
  /**
  The {extId} of the file server. This is a read-only field.
  */
  SourceFileServerExtId *string `json:"sourceFileServerExtId"`
  /**
  The {extId} of the target file server. This is a read-only field.
  */
  TargetFileServerExtId *string `json:"targetFileServerExtId"`
}

func (p *SubPolicy) MarshalJSON() ([]byte, error) {
  type SubPolicyProxy SubPolicy
  return json.Marshal(struct {
    *SubPolicyProxy
    SourceFileServerExtId *string `json:"sourceFileServerExtId,omitempty"`
    TargetFileServerExtId *string `json:"targetFileServerExtId,omitempty"`
  }{
    SubPolicyProxy : (*SubPolicyProxy)(p),
    SourceFileServerExtId : p.SourceFileServerExtId,
    TargetFileServerExtId : p.TargetFileServerExtId,
  })
}

func NewSubPolicy() *SubPolicy {
  p := new(SubPolicy)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.SubPolicy"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.SubPolicy"}
  p.UnknownFields_ = map[string]interface{}{}

  p.IncludeNewMountTargets = new(bool)
  *p.IncludeNewMountTargets = true
  p.IsReversePolicy = new(bool)
  *p.IsReversePolicy = false


  return p
}




/**
This contains summary for replication and contains values. The acceptable values are "SUMMARY_MET", "SUMMARY_MISSED", "SUMMARY_PAUSED", "SUMMARY_REPLICATION_ERROR".
*/
type Summary int

const(
  SUMMARY_UNKNOWN Summary = 0
  SUMMARY_REDACTED Summary = 1
  SUMMARY_SUMMARY_MET Summary = 2
  SUMMARY_SUMMARY_MISSED Summary = 3
  SUMMARY_SUMMARY_PAUSED Summary = 4
  SUMMARY_SUMMARY_REPLICATION_ERROR Summary = 5
)

// returns the name of the enum given an ordinal number
func (e *Summary) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "SUMMARY_MET",
    "SUMMARY_MISSED",
    "SUMMARY_PAUSED",
    "SUMMARY_REPLICATION_ERROR",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *Summary) index(name string) Summary {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "SUMMARY_MET",
    "SUMMARY_MISSED",
    "SUMMARY_PAUSED",
    "SUMMARY_REPLICATION_ERROR",
  }
  for idx := range names {
    if names[idx] == name {
      return Summary(idx)
    }
  }
  return SUMMARY_UNKNOWN
}

func (e *Summary) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for Summary:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *Summary) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e Summary) Ref() *Summary {
  return &e
}


/**
Data synchronization path entity.
*/
type SynchronizationPathEntity struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Data synchronization path
  */
  SyncPath *string `json:"syncPath"`
}

func (p *SynchronizationPathEntity) MarshalJSON() ([]byte, error) {
  type SynchronizationPathEntityProxy SynchronizationPathEntity
  return json.Marshal(struct {
    *SynchronizationPathEntityProxy
    SyncPath *string `json:"syncPath,omitempty"`
  }{
    SynchronizationPathEntityProxy : (*SynchronizationPathEntityProxy)(p),
    SyncPath : p.SyncPath,
  })
}

func NewSynchronizationPathEntity() *SynchronizationPathEntity {
  p := new(SynchronizationPathEntity)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.SynchronizationPathEntity"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.SynchronizationPathEntity"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Data synchronization path entity validation response.
*/
type SynchronizationPathEntityResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Data synchronization path error message.
  */
  ErrorMessage *string `json:"errorMessage,omitempty"`
  /**
  Data synchronization path
  */
  SyncPath *string `json:"syncPath,omitempty"`
}


func NewSynchronizationPathEntityResponse() *SynchronizationPathEntityResponse {
  p := new(SynchronizationPathEntityResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.SynchronizationPathEntityResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.SynchronizationPathEntityResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Data synchronization paths.
*/
type SynchronizationPaths struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Data synchronization path entities.
  */
  SyncPathEntities []SynchronizationPathEntity `json:"syncPathEntities"`
}

func (p *SynchronizationPaths) MarshalJSON() ([]byte, error) {
  type SynchronizationPathsProxy SynchronizationPaths
  return json.Marshal(struct {
    *SynchronizationPathsProxy
    SyncPathEntities []SynchronizationPathEntity `json:"syncPathEntities,omitempty"`
  }{
    SynchronizationPathsProxy : (*SynchronizationPathsProxy)(p),
    SyncPathEntities : p.SyncPathEntities,
  })
}

func NewSynchronizationPaths() *SynchronizationPaths {
  p := new(SynchronizationPaths)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.SynchronizationPaths"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.SynchronizationPaths"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST response for all response codes in api path /files/v4.0.a2/config/file-server/data-synchronization/$actions/validate Post operation
*/
type SynchronizationPathsValidationApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfSynchronizationPathsValidationApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewSynchronizationPathsValidationApiResponse() *SynchronizationPathsValidationApiResponse {
  p := new(SynchronizationPathsValidationApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.SynchronizationPathsValidationApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.SynchronizationPathsValidationApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *SynchronizationPathsValidationApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *SynchronizationPathsValidationApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfSynchronizationPathsValidationApiResponseData()
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
Data synchronization path validation response.
*/
type SynchronizationPathsValidationResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Data synchronization path entities validation response.
  */
  SyncPathEntitiesResponse []SynchronizationPathEntityResponse `json:"syncPathEntitiesResponse,omitempty"`
}


func NewSynchronizationPathsValidationResponse() *SynchronizationPathsValidationResponse {
  p := new(SynchronizationPathsValidationResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.SynchronizationPathsValidationResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.SynchronizationPathsValidationResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
TLSReqCert Type corresponding to TLSReqCertStatus. The acceptable values are "ALLOW", "DEMAND", "NEVER".
*/
type TLSReqCertStatus int

const(
  TLSREQCERTSTATUS_UNKNOWN TLSReqCertStatus = 0
  TLSREQCERTSTATUS_REDACTED TLSReqCertStatus = 1
  TLSREQCERTSTATUS_ALLOW TLSReqCertStatus = 2
  TLSREQCERTSTATUS_DEMAND TLSReqCertStatus = 3
  TLSREQCERTSTATUS_NEVER TLSReqCertStatus = 4
)

// returns the name of the enum given an ordinal number
func (e *TLSReqCertStatus) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "ALLOW",
    "DEMAND",
    "NEVER",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *TLSReqCertStatus) index(name string) TLSReqCertStatus {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "ALLOW",
    "DEMAND",
    "NEVER",
  }
  for idx := range names {
    if names[idx] == name {
      return TLSReqCertStatus(idx)
    }
  }
  return TLSREQCERTSTATUS_UNKNOWN
}

func (e *TLSReqCertStatus) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for TLSReqCertStatus:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *TLSReqCertStatus) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e TLSReqCertStatus) Ref() *TLSReqCertStatus {
  return &e
}


/**
Template mapping type. The acceptable values are "NO_TEMPLATE_MAPPING", "SMB_NAME_NFS_NAME", "NO_MAPPING".
*/
type TemplateMappingType int

const(
  TEMPLATEMAPPINGTYPE_UNKNOWN TemplateMappingType = 0
  TEMPLATEMAPPINGTYPE_REDACTED TemplateMappingType = 1
  TEMPLATEMAPPINGTYPE_NO_TEMPLATE_MAPPING TemplateMappingType = 2
  TEMPLATEMAPPINGTYPE_SMB_NAME_NFS_NAME TemplateMappingType = 3
  TEMPLATEMAPPINGTYPE_NO_MAPPING TemplateMappingType = 4
)

// returns the name of the enum given an ordinal number
func (e *TemplateMappingType) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "NO_TEMPLATE_MAPPING",
    "SMB_NAME_NFS_NAME",
    "NO_MAPPING",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *TemplateMappingType) index(name string) TemplateMappingType {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "NO_TEMPLATE_MAPPING",
    "SMB_NAME_NFS_NAME",
    "NO_MAPPING",
  }
  for idx := range names {
    if names[idx] == name {
      return TemplateMappingType(idx)
    }
  }
  return TEMPLATEMAPPINGTYPE_UNKNOWN
}

func (e *TemplateMappingType) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for TemplateMappingType:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *TemplateMappingType) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e TemplateMappingType) Ref() *TemplateMappingType {
  return &e
}


/**
REST response for all response codes in api path /files/v4.0.a2/config/file-server/anti-virus-servers/{antivirusServerExtId}/$actions/test-connection Post operation
*/
type TestConnectionAntivirusServerApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfTestConnectionAntivirusServerApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewTestConnectionAntivirusServerApiResponse() *TestConnectionAntivirusServerApiResponse {
  p := new(TestConnectionAntivirusServerApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.TestConnectionAntivirusServerApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.TestConnectionAntivirusServerApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *TestConnectionAntivirusServerApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *TestConnectionAntivirusServerApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfTestConnectionAntivirusServerApiResponseData()
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
Model for tier data.
*/
type TierData struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Duration (in seconds) for which the tiering action should run. If not provided tiering will run until the memory has reached the threshold capacity value.
  */
  DurationInSeconds *int64 `json:"durationInSeconds,omitempty"`
}


func NewTierData() *TierData {
  p := new(TierData)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.TierData"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.TierData"}
  p.UnknownFields_ = map[string]interface{}{}

  p.DurationInSeconds = new(int64)
  *p.DurationInSeconds = 0


  return p
}




/**
REST response for all response codes in api path /files/v4.0.a2/config/file-server/$actions/tier-data Post operation
*/
type TierDataApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfTierDataApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewTierDataApiResponse() *TierDataApiResponse {
  p := new(TierDataApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.TierDataApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.TierDataApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *TierDataApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *TierDataApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfTierDataApiResponseData()
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
Request body specification for tier, recall and tier status APIs.
*/
type TierFileInfo struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Path of the file to be tiered relative to root Mount target.
  */
  FilePath []string `json:"filePath,omitempty"`
  /**
  List of inode numbers strings. Inode number should be in the format Inode/Inode:GenId.
  */
  INodeNumber []string `json:"iNodeNumber,omitempty"`
  /**
  Boolean to indicate if all files belong to the same VG.
  */
  IsSingleDirTree *bool `json:"isSingleDirTree,omitempty"`
  /**
  The {extId} of the mount target. Example:9c1e537d-6777-4c22-5d41-ddd0c3337aa9.
  */
  MountTargetExtId *string `json:"mountTargetExtId"`
  /**
  Boolean to indicate if all files belong to the same VG. This field will be deferred in future.
  */
  SingleDirTree *bool `json:"singleDirTree,omitempty"`
}

func (p *TierFileInfo) MarshalJSON() ([]byte, error) {
  type TierFileInfoProxy TierFileInfo
  return json.Marshal(struct {
    *TierFileInfoProxy
    MountTargetExtId *string `json:"mountTargetExtId,omitempty"`
  }{
    TierFileInfoProxy : (*TierFileInfoProxy)(p),
    MountTargetExtId : p.MountTargetExtId,
  })
}

func NewTierFileInfo() *TierFileInfo {
  p := new(TierFileInfo)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.TierFileInfo"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.TierFileInfo"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST response for all response codes in api path /files/v4.0.a2/config/file-server/$actions/recall-files Post operation
*/
type TierFilesApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfTierFilesApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewTierFilesApiResponse() *TierFilesApiResponse {
  p := new(TierFilesApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.TierFilesApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.TierFilesApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *TierFilesApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *TierFilesApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfTierFilesApiResponseData()
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
REST response for all response codes in api path /files/v4.0.a2/config/file-server/$actions/refresh-cold-data Post operation
*/
type TierRefreshColdDataApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfTierRefreshColdDataApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewTierRefreshColdDataApiResponse() *TierRefreshColdDataApiResponse {
  p := new(TierRefreshColdDataApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.TierRefreshColdDataApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.TierRefreshColdDataApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *TierRefreshColdDataApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *TierRefreshColdDataApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfTierRefreshColdDataApiResponseData()
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
Request body specification for tier, recall and tier status APIs.
*/
type TierRequestSpec struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  The {extId} of the file server. This is a read-only field.
  */
  FileServerExtId *string `json:"fileServerExtId,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import1.ApiLink `json:"links,omitempty"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
  /**
  Contains file paths and Mount target information for tier request.
  */
  TierFileInfo []TierFileInfo `json:"tierFileInfo"`
}

func (p *TierRequestSpec) MarshalJSON() ([]byte, error) {
  type TierRequestSpecProxy TierRequestSpec
  return json.Marshal(struct {
    *TierRequestSpecProxy
    TierFileInfo []TierFileInfo `json:"tierFileInfo,omitempty"`
  }{
    TierRequestSpecProxy : (*TierRequestSpecProxy)(p),
    TierFileInfo : p.TierFileInfo,
  })
}

func NewTierRequestSpec() *TierRequestSpec {
  p := new(TierRequestSpec)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.TierRequestSpec"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.TierRequestSpec"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Tier request status. The acceptable values are "ACCEPTED" or "NOT_ACCEPTED".
*/
type TierRequestStatus int

const(
  TIERREQUESTSTATUS_UNKNOWN TierRequestStatus = 0
  TIERREQUESTSTATUS_REDACTED TierRequestStatus = 1
  TIERREQUESTSTATUS_ACCEPTED TierRequestStatus = 2
  TIERREQUESTSTATUS_NOT_ACCEPTED TierRequestStatus = 3
)

// returns the name of the enum given an ordinal number
func (e *TierRequestStatus) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "ACCEPTED",
    "NOT_ACCEPTED",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *TierRequestStatus) index(name string) TierRequestStatus {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "ACCEPTED",
    "NOT_ACCEPTED",
  }
  for idx := range names {
    if names[idx] == name {
      return TierRequestStatus(idx)
    }
  }
  return TIERREQUESTSTATUS_UNKNOWN
}

func (e *TierRequestStatus) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for TierRequestStatus:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *TierRequestStatus) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e TierRequestStatus) Ref() *TierRequestStatus {
  return &e
}


/**
Response body for tier and recall APIs.
*/
type TierResponseSpec struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Error code for tiering.
  */
  ErrorCode *int `json:"errorCode,omitempty"`
  /**
  Error message for tiering.
  */
  ErrorMessage *string `json:"errorMessage,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import1.ApiLink `json:"links,omitempty"`
  
  Status *TierRequestStatus `json:"status,omitempty"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
}


func NewTierResponseSpec() *TierResponseSpec {
  p := new(TierResponseSpec)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.TierResponseSpec"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.TierResponseSpec"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Request body specification for tiering configuration APIs.
*/
type TieringConfig struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Capacity threshold for tiering.
  */
  CapacityThreshold *int64 `json:"capacityThreshold,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import1.ApiLink `json:"links,omitempty"`
  /**
  Tiering schedule for on-prem tiering.
  */
  Schedule []PerDayScheduleSpec `json:"schedule,omitempty"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
}


func NewTieringConfig() *TieringConfig {
  p := new(TieringConfig)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.TieringConfig"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.TieringConfig"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Object store profile specification.
*/
type TieringObjectStoreProfile struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Access key
  */
  AccessKey *string `json:"accessKey,omitempty"`
  /**
  Base URL
  */
  BaseUrl *string `json:"baseUrl,omitempty"`
  /**
  Bucket name
  */
  BucketName *string `json:"bucketName,omitempty"`
  /**
  CA certificate
  */
  CaCertContent *string `json:"caCertContent,omitempty"`
  
  EnableMountTargets *EnableMountTargets `json:"enableMountTargets,omitempty"`
  /**
  Enable SSL verify peer certificate
  */
  EnableSslVerifyPeer *bool `json:"enableSslVerifyPeer,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  The {extId} of the file server. This is a read-only field.
  */
  FileServerExtId *string `json:"fileServerExtId,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import1.ApiLink `json:"links,omitempty"`
  /**
  The {extId} of the mount target. Example:9c1e537d-6777-4c22-5d41-ddd0c3337aa9.
  */
  MountTargetExtIds []string `json:"mountTargetExtIds,omitempty"`
  /**
  Object store cloud profile name
  */
  ProfileName *string `json:"profileName,omitempty"`
  /**
  HTTP proxy server address
  */
  ProxyServer *string `json:"proxyServer,omitempty"`
  /**
  Time in days for which the data will be maintained in the cloud after it is deleted from local storage on the file server.
  */
  RetentionPeriod *int `json:"retentionPeriod,omitempty"`
  /**
  Secret key
  */
  SecretKey *string `json:"secretKey,omitempty"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
  
  TieringObjectStoreType *TieringObjectStoreType `json:"tieringObjectStoreType,omitempty"`
}


func NewTieringObjectStoreProfile() *TieringObjectStoreProfile {
  p := new(TieringObjectStoreProfile)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.TieringObjectStoreProfile"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.TieringObjectStoreProfile"}
  p.UnknownFields_ = map[string]interface{}{}

  p.RetentionPeriod = new(int)
  *p.RetentionPeriod = 1825


  return p
}




/**
Object store. The acceptable values are "NUTANIX", "AWS", "GCP", "AZURE", "AWS_STANDARD_IA", "OTHER".
*/
type TieringObjectStoreType int

const(
  TIERINGOBJECTSTORETYPE_UNKNOWN TieringObjectStoreType = 0
  TIERINGOBJECTSTORETYPE_REDACTED TieringObjectStoreType = 1
  TIERINGOBJECTSTORETYPE_NUTANIX TieringObjectStoreType = 2
  TIERINGOBJECTSTORETYPE_AWS TieringObjectStoreType = 3
  TIERINGOBJECTSTORETYPE_GCP TieringObjectStoreType = 4
  TIERINGOBJECTSTORETYPE_AZURE TieringObjectStoreType = 5
  TIERINGOBJECTSTORETYPE_AWS_STANDARD_IA TieringObjectStoreType = 6
  TIERINGOBJECTSTORETYPE_OTHER TieringObjectStoreType = 7
  TIERINGOBJECTSTORETYPE_AWS_GLACIER_IR TieringObjectStoreType = 8
)

// returns the name of the enum given an ordinal number
func (e *TieringObjectStoreType) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "NUTANIX",
    "AWS",
    "GCP",
    "AZURE",
    "AWS_STANDARD_IA",
    "OTHER",
    "AWS_GLACIER_IR",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *TieringObjectStoreType) index(name string) TieringObjectStoreType {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "NUTANIX",
    "AWS",
    "GCP",
    "AZURE",
    "AWS_STANDARD_IA",
    "OTHER",
    "AWS_GLACIER_IR",
  }
  for idx := range names {
    if names[idx] == name {
      return TieringObjectStoreType(idx)
    }
  }
  return TIERINGOBJECTSTORETYPE_UNKNOWN
}

func (e *TieringObjectStoreType) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for TieringObjectStoreType:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *TieringObjectStoreType) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e TieringObjectStoreType) Ref() *TieringObjectStoreType {
  return &e
}


/**
Tiering policy model
*/
type TieringPolicy struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Cool off period in seconds for tiering. Files older than a this time will be considered for tiering.
  */
  CooloffPeriodInSeconds *int64 `json:"cooloffPeriodInSeconds,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  Indicates whether to include future shares to the tiering policy.
  */
  IsIncludeFutureShares *bool `json:"isIncludeFutureShares,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import1.ApiLink `json:"links,omitempty"`
  /**
  Minimum file size for tiering. Files size greater than this will be considered for tiering.
  */
  MinimumFileSizeInBytes *int64 `json:"minimumFileSizeInBytes,omitempty"`
  /**
  The {extId} of the mount target. Example:9c1e537d-6777-4c22-5d41-ddd0c3337aa9.
  */
  MountTargetExtIds []string `json:"mountTargetExtIds,omitempty"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
}


func NewTieringPolicy() *TieringPolicy {
  p := new(TieringPolicy)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.TieringPolicy"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.TieringPolicy"}
  p.UnknownFields_ = map[string]interface{}{}

  p.MinimumFileSizeInBytes = new(int64)
  *p.MinimumFileSizeInBytes = 65536


  return p
}




/**
REST response for all response codes in api path /files/v4.0.a2/config/file-server/$actions/unquarantine-infected-files Post operation
*/
type UnQuarantineInfectedFilesApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfUnQuarantineInfectedFilesApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewUnQuarantineInfectedFilesApiResponse() *UnQuarantineInfectedFilesApiResponse {
  p := new(UnQuarantineInfectedFilesApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.UnQuarantineInfectedFilesApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.UnQuarantineInfectedFilesApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *UnQuarantineInfectedFilesApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *UnQuarantineInfectedFilesApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfUnQuarantineInfectedFilesApiResponseData()
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
REST response for all response codes in api path /files/v4.0.a2/config/file-server/admin-users/{adminUserExtId} Put operation
*/
type UpdateAdminUserApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfUpdateAdminUserApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewUpdateAdminUserApiResponse() *UpdateAdminUserApiResponse {
  p := new(UpdateAdminUserApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.UpdateAdminUserApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.UpdateAdminUserApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *UpdateAdminUserApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *UpdateAdminUserApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfUpdateAdminUserApiResponseData()
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
REST response for all response codes in api path /files/v4.0.a2/config/file-server/anti-virus-servers/{antivirusServerExtId} Put operation
*/
type UpdateAntivirusServerApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfUpdateAntivirusServerApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewUpdateAntivirusServerApiResponse() *UpdateAntivirusServerApiResponse {
  p := new(UpdateAntivirusServerApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.UpdateAntivirusServerApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.UpdateAntivirusServerApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *UpdateAntivirusServerApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *UpdateAntivirusServerApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfUpdateAntivirusServerApiResponseData()
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
REST response for all response codes in api path /files/v4.0.a2/config/file-server/data-consolidation-policies/{dataConsolidationPolicyExtId} Put operation
*/
type UpdateDataConsolidationPolicyApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfUpdateDataConsolidationPolicyApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewUpdateDataConsolidationPolicyApiResponse() *UpdateDataConsolidationPolicyApiResponse {
  p := new(UpdateDataConsolidationPolicyApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.UpdateDataConsolidationPolicyApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.UpdateDataConsolidationPolicyApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *UpdateDataConsolidationPolicyApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *UpdateDataConsolidationPolicyApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfUpdateDataConsolidationPolicyApiResponseData()
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
REST response for all response codes in api path /files/v4.0.a2/config/file-server/data-protection-policies/{dataProtectionPolicyExtId} Put operation
*/
type UpdateDataProtectionPolicyApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfUpdateDataProtectionPolicyApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewUpdateDataProtectionPolicyApiResponse() *UpdateDataProtectionPolicyApiResponse {
  p := new(UpdateDataProtectionPolicyApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.UpdateDataProtectionPolicyApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.UpdateDataProtectionPolicyApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *UpdateDataProtectionPolicyApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *UpdateDataProtectionPolicyApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfUpdateDataProtectionPolicyApiResponseData()
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
REST response for all response codes in api path /files/v4.0.a2/config/file-server Put operation
*/
type UpdateFileServerApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfUpdateFileServerApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewUpdateFileServerApiResponse() *UpdateFileServerApiResponse {
  p := new(UpdateFileServerApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.UpdateFileServerApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.UpdateFileServerApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *UpdateFileServerApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *UpdateFileServerApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfUpdateFileServerApiResponseData()
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
REST response for all response codes in api path /files/v4.0.a2/config/file-server/users/{fileServerUserExtId} Put operation
*/
type UpdateFileServerUserApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfUpdateFileServerUserApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewUpdateFileServerUserApiResponse() *UpdateFileServerUserApiResponse {
  p := new(UpdateFileServerUserApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.UpdateFileServerUserApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.UpdateFileServerUserApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *UpdateFileServerUserApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *UpdateFileServerUserApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfUpdateFileServerUserApiResponseData()
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
REST response for all response codes in api path /files/v4.0.a2/config/file-server/blocked-client Put operation
*/
type UpdateGlobalBlockedClientApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfUpdateGlobalBlockedClientApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewUpdateGlobalBlockedClientApiResponse() *UpdateGlobalBlockedClientApiResponse {
  p := new(UpdateGlobalBlockedClientApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.UpdateGlobalBlockedClientApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.UpdateGlobalBlockedClientApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *UpdateGlobalBlockedClientApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *UpdateGlobalBlockedClientApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfUpdateGlobalBlockedClientApiResponseData()
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
REST response for all response codes in api path /files/v4.0.a2/config/file-server/mount-targets/{mountTargetExtId} Put operation
*/
type UpdateMountTargetApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfUpdateMountTargetApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewUpdateMountTargetApiResponse() *UpdateMountTargetApiResponse {
  p := new(UpdateMountTargetApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.UpdateMountTargetApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.UpdateMountTargetApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *UpdateMountTargetApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *UpdateMountTargetApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfUpdateMountTargetApiResponseData()
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
REST response for all response codes in api path /files/v4.0.a2/config/file-server/mount-targets/{mountTargetExtId}/virus-scan-policy/{virusScanPolicyExtId} Put operation
*/
type UpdateMountTargetVirusScanPolicyApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfUpdateMountTargetVirusScanPolicyApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewUpdateMountTargetVirusScanPolicyApiResponse() *UpdateMountTargetVirusScanPolicyApiResponse {
  p := new(UpdateMountTargetVirusScanPolicyApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.UpdateMountTargetVirusScanPolicyApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.UpdateMountTargetVirusScanPolicyApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *UpdateMountTargetVirusScanPolicyApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *UpdateMountTargetVirusScanPolicyApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfUpdateMountTargetVirusScanPolicyApiResponseData()
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
REST response for all response codes in api path /files/v4.0.a2/config/file-server/object-store/profiles/{profileExtId} Put operation
*/
type UpdateObjectStoreApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfUpdateObjectStoreApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewUpdateObjectStoreApiResponse() *UpdateObjectStoreApiResponse {
  p := new(UpdateObjectStoreApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.UpdateObjectStoreApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.UpdateObjectStoreApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *UpdateObjectStoreApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *UpdateObjectStoreApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfUpdateObjectStoreApiResponseData()
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
REST response for all response codes in api path /files/v4.0.a2/config/file-server/quota-email-config Put operation
*/
type UpdateQuotaEmailConfigApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfUpdateQuotaEmailConfigApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewUpdateQuotaEmailConfigApiResponse() *UpdateQuotaEmailConfigApiResponse {
  p := new(UpdateQuotaEmailConfigApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.UpdateQuotaEmailConfigApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.UpdateQuotaEmailConfigApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *UpdateQuotaEmailConfigApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *UpdateQuotaEmailConfigApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfUpdateQuotaEmailConfigApiResponseData()
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
REST response for all response codes in api path /files/v4.0.a2/config/file-server/mount-targets/{mountTargetExtId}/quota-policies/{quotaPolicyExtId} Put operation
*/
type UpdateQuotaPolicyApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfUpdateQuotaPolicyApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewUpdateQuotaPolicyApiResponse() *UpdateQuotaPolicyApiResponse {
  p := new(UpdateQuotaPolicyApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.UpdateQuotaPolicyApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.UpdateQuotaPolicyApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *UpdateQuotaPolicyApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *UpdateQuotaPolicyApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfUpdateQuotaPolicyApiResponseData()
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
REST response for all response codes in api path /files/v4.0.a2/config/file-server/ransomware-config Post operation
*/
type UpdateRansomwareConfigApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfUpdateRansomwareConfigApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewUpdateRansomwareConfigApiResponse() *UpdateRansomwareConfigApiResponse {
  p := new(UpdateRansomwareConfigApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.UpdateRansomwareConfigApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.UpdateRansomwareConfigApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *UpdateRansomwareConfigApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *UpdateRansomwareConfigApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfUpdateRansomwareConfigApiResponseData()
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
REST response for all response codes in api path /files/v4.0.a2/config/file-server/ssr-snapshot-schedules/{ssrSnapshotScheduleExtId} Put operation
*/
type UpdateSsrSnapshotScheduleApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfUpdateSsrSnapshotScheduleApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewUpdateSsrSnapshotScheduleApiResponse() *UpdateSsrSnapshotScheduleApiResponse {
  p := new(UpdateSsrSnapshotScheduleApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.UpdateSsrSnapshotScheduleApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.UpdateSsrSnapshotScheduleApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *UpdateSsrSnapshotScheduleApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *UpdateSsrSnapshotScheduleApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfUpdateSsrSnapshotScheduleApiResponseData()
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
REST response for all response codes in api path /files/v4.0.a2/config/file-server/tier-config/{extId} Put operation
*/
type UpdateTieringConfigApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfUpdateTieringConfigApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewUpdateTieringConfigApiResponse() *UpdateTieringConfigApiResponse {
  p := new(UpdateTieringConfigApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.UpdateTieringConfigApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.UpdateTieringConfigApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *UpdateTieringConfigApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *UpdateTieringConfigApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfUpdateTieringConfigApiResponseData()
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
REST response for all response codes in api path /files/v4.0.a2/config/file-server/tier-policies/{extId} Put operation
*/
type UpdateTieringPolicyApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfUpdateTieringPolicyApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewUpdateTieringPolicyApiResponse() *UpdateTieringPolicyApiResponse {
  p := new(UpdateTieringPolicyApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.UpdateTieringPolicyApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.UpdateTieringPolicyApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *UpdateTieringPolicyApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *UpdateTieringPolicyApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfUpdateTieringPolicyApiResponseData()
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
REST response for all response codes in api path /files/v4.0.a2/config/file-server/user-mapping Put operation
*/
type UpdateUserMappingApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfUpdateUserMappingApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewUpdateUserMappingApiResponse() *UpdateUserMappingApiResponse {
  p := new(UpdateUserMappingApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.UpdateUserMappingApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.UpdateUserMappingApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *UpdateUserMappingApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *UpdateUserMappingApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfUpdateUserMappingApiResponseData()
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
REST response for all response codes in api path /files/v4.0.a2/config/file-server/virus-scan-policy/{virusScanPolicyExtId} Put operation
*/
type UpdateVirusScanPolicyApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfUpdateVirusScanPolicyApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewUpdateVirusScanPolicyApiResponse() *UpdateVirusScanPolicyApiResponse {
  p := new(UpdateVirusScanPolicyApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.UpdateVirusScanPolicyApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.UpdateVirusScanPolicyApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *UpdateVirusScanPolicyApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *UpdateVirusScanPolicyApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfUpdateVirusScanPolicyApiResponseData()
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
User mapping model.
*/
type UserMapping struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  ExplicitMapping *ExplicitIdentityMapping `json:"explicitMapping,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import1.ApiLink `json:"links,omitempty"`
  
  RuleBasedMapping *RuleBasedMapping `json:"ruleBasedMapping,omitempty"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
}


func NewUserMapping() *UserMapping {
  p := new(UserMapping)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.UserMapping"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.UserMapping"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST response for all response codes in api path /files/v4.0.a2/config/file-server/user-mapping Get operation
*/
type UserMappingListApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfUserMappingListApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewUserMappingListApiResponse() *UserMappingListApiResponse {
  p := new(UserMappingListApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.UserMappingListApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.UserMappingListApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *UserMappingListApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *UserMappingListApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfUserMappingListApiResponseData()
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
File server VM object.
*/
type VM struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Acropolis UUID of the file server VM. This is a read-only field.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  List of external networks associated with the file server VM.
  */
  ExternalNetworks []VMNetwork `json:"externalNetworks,omitempty"`
  /**
  FSVM UUID of the file server VM.
  */
  FsvmUuid *string `json:"fsvmUuid,omitempty"`
  /**
  List of internal networks associated with the file server VM.
  */
  InternalNetworks []VMNetwork `json:"internalNetworks,omitempty"`
  /**
  Memory associated with the file server VM.
  */
  MemoryGib *int64 `json:"memoryGib,omitempty"`
  /**
  Name of the file server VM.
  */
  Name *string `json:"name,omitempty"`
  /**
  CPU associated with the file server VM.
  */
  Vcpus *int64 `json:"vcpus,omitempty"`
}


func NewVM() *VM {
  p := new(VM)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.VM"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.VM"}
  p.UnknownFields_ = map[string]interface{}{}

  p.MemoryGib = new(int64)
  *p.MemoryGib = 12
  p.Vcpus = new(int64)
  *p.Vcpus = 4


  return p
}




/**
File server network object.
*/
type VMNetwork struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  IPV4 address range to be used in this network.
  */
  IpAddresses []import3.IPAddress `json:"ipAddresses"`
  /**
  UUID of the network. This is a read-only field.
  */
  VmNetworkExtId *string `json:"vmNetworkExtId,omitempty"`
}

func (p *VMNetwork) MarshalJSON() ([]byte, error) {
  type VMNetworkProxy VMNetwork
  return json.Marshal(struct {
    *VMNetworkProxy
    IpAddresses []import3.IPAddress `json:"ipAddresses,omitempty"`
  }{
    VMNetworkProxy : (*VMNetworkProxy)(p),
    IpAddresses : p.IpAddresses,
  })
}

func NewVMNetwork() *VMNetwork {
  p := new(VMNetwork)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.VMNetwork"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.VMNetwork"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST response for all response codes in api path /files/v4.0.a2/config/file-server/$actions/verify-dns-entries Post operation
*/
type VerifyDnsEntriesApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfVerifyDnsEntriesApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewVerifyDnsEntriesApiResponse() *VerifyDnsEntriesApiResponse {
  p := new(VerifyDnsEntriesApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.VerifyDnsEntriesApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.VerifyDnsEntriesApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *VerifyDnsEntriesApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *VerifyDnsEntriesApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfVerifyDnsEntriesApiResponseData()
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
Antivirus server scan policy information.
*/
type VirusScanPolicy struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Indicates if the file access needs to be blocked for current policy.
  */
  BlockAccessFile *bool `json:"blockAccessFile,omitempty"`
  /**
  Indicates if the antivirus server is enabled for current policy.
  */
  EnableAntiVirus *bool `json:"enableAntiVirus,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  The {extId} of the file server. This is a read-only field.
  */
  FileServerExtId *string `json:"fileServerExtId,omitempty"`
  /**
  File size exclusion. It is the size of a file below which AV scan won't run. Default value is 0 (run on all files).
  */
  FileSizeExclusionInBytes *int64 `json:"fileSizeExclusionInBytes,omitempty"`
  /**
  File type exclusion list.
  */
  FileTypeExclusions []string `json:"fileTypeExclusions,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import1.ApiLink `json:"links,omitempty"`
  /**
  Scan on read policy.
  */
  ScanOnRead *bool `json:"scanOnRead,omitempty"`
  /**
  Scan on write policy.
  */
  ScanOnWrite *bool `json:"scanOnWrite,omitempty"`
  /**
  Scan time interval in seconds.
  */
  ScanTimeoutIntervalInSecs *int `json:"scanTimeoutIntervalInSecs,omitempty"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
}


func NewVirusScanPolicy() *VirusScanPolicy {
  p := new(VirusScanPolicy)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.VirusScanPolicy"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.VirusScanPolicy"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST response for all response codes in api path /files/v4.0.a2/config/file-server/virus-scan-policy/{virusScanPolicyExtId} Get operation
*/
type VirusScanPolicyApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfVirusScanPolicyApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewVirusScanPolicyApiResponse() *VirusScanPolicyApiResponse {
  p := new(VirusScanPolicyApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.VirusScanPolicyApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.VirusScanPolicyApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *VirusScanPolicyApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *VirusScanPolicyApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfVirusScanPolicyApiResponseData()
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
Antivirus server scan policy information.
*/
type VirusScanPolicyMountTarget struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Indicates if the file access needs to be blocked for current policy.
  */
  BlockAccessFile *bool `json:"blockAccessFile,omitempty"`
  /**
  Indicates if the antivirus server is enabled for current policy.
  */
  EnableAntiVirus *bool `json:"enableAntiVirus,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  The {extId} of the file server. This is a read-only field.
  */
  FileServerExtId *string `json:"fileServerExtId,omitempty"`
  /**
  File size exclusion. It is the size of a file below which AV scan won't run. Default value is 0 (run on all files).
  */
  FileSizeExclusionInBytes *int64 `json:"fileSizeExclusionInBytes,omitempty"`
  /**
  File type exclusion list.
  */
  FileTypeExclusions []string `json:"fileTypeExclusions,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import1.ApiLink `json:"links,omitempty"`
  /**
  The {extId} of the mount target under which the virus scan policy is being created. This field will be set to null if it is a global policy.
  */
  MountTargetExtId *string `json:"mountTargetExtId,omitempty"`
  /**
  Scan on read policy.
  */
  ScanOnRead *bool `json:"scanOnRead,omitempty"`
  /**
  Scan on write policy.
  */
  ScanOnWrite *bool `json:"scanOnWrite,omitempty"`
  /**
  Scan time interval in seconds.
  */
  ScanTimeoutIntervalInSecs *int `json:"scanTimeoutIntervalInSecs,omitempty"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
}


func NewVirusScanPolicyMountTarget() *VirusScanPolicyMountTarget {
  p := new(VirusScanPolicyMountTarget)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.VirusScanPolicyMountTarget"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.VirusScanPolicyMountTarget"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Wild card mapping model.
*/
type WildCardMapping struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  IdentityMappingType *IdentityMappingType `json:"identityMappingType,omitempty"`
  
  NfsMapping *NfsIdentitiesMapping `json:"nfsMapping,omitempty"`
  /**
  Identity mapping priority
  */
  Priority *int `json:"priority"`
  
  SmbMapping *SmbIdentitiesMapping `json:"smbMapping,omitempty"`
}

func (p *WildCardMapping) MarshalJSON() ([]byte, error) {
  type WildCardMappingProxy WildCardMapping
  return json.Marshal(struct {
    *WildCardMappingProxy
    Priority *int `json:"priority,omitempty"`
  }{
    WildCardMappingProxy : (*WildCardMappingProxy)(p),
    Priority : p.Priority,
  })
}

func NewWildCardMapping() *WildCardMapping {
  p := new(WildCardMapping)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.WildCardMapping"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.WildCardMapping"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Worm share details.
*/
type WormSpec struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Cooloff interval for worm share.
  */
  CooloffInterval *int64 `json:"cooloffInterval,omitempty"`
  /**
  Retention period for worm share.
  */
  RetentionPeriod *int64 `json:"retentionPeriod,omitempty"`
  
  WormType *WormType `json:"wormType,omitempty"`
}


func NewWormSpec() *WormSpec {
  p := new(WormSpec)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.config.WormSpec"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.config.WormSpec"}
  p.UnknownFields_ = map[string]interface{}{}

  p.CooloffInterval = new(int64)
  *p.CooloffInterval = 600
  p.RetentionPeriod = new(int64)
  *p.RetentionPeriod = 31449600


  return p
}




/**
Worm type for share. The acceptable values are "SHARE_LEVEL" or "DISABLED".
*/
type WormType int

const(
  WORMTYPE_UNKNOWN WormType = 0
  WORMTYPE_REDACTED WormType = 1
  WORMTYPE_SHARE_LEVEL WormType = 2
  WORMTYPE_DISABLED WormType = 3
)

// returns the name of the enum given an ordinal number
func (e *WormType) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "SHARE_LEVEL",
    "DISABLED",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *WormType) index(name string) WormType {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "SHARE_LEVEL",
    "DISABLED",
  }
  for idx := range names {
    if names[idx] == name {
      return WormType(idx)
    }
  }
  return WORMTYPE_UNKNOWN
}

func (e *WormType) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for WormType:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *WormType) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e WormType) Ref() *WormType {
  return &e
}

type OneOfDeleteNetworkApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *import4.TaskReference `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfDeleteNetworkApiResponseData() *OneOfDeleteNetworkApiResponseData {
  p := new(OneOfDeleteNetworkApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfDeleteNetworkApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfDeleteNetworkApiResponseData is nil"))
  }
  switch v.(type) {
    case import4.TaskReference:
      if nil == p.oneOfType0 {p.oneOfType0 = new(import4.TaskReference)}
      *p.oneOfType0 = v.(import4.TaskReference)
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

func (p *OneOfDeleteNetworkApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfDeleteNetworkApiResponseData) UnmarshalJSON(b []byte) error {
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
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteNetworkApiResponseData"))
}

func (p *OneOfDeleteNetworkApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfDeleteNetworkApiResponseData")
}

type OneOfUpdateVirusScanPolicyApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
  oneOfType0 *VirusScanPolicy `json:"-"`
}

func NewOneOfUpdateVirusScanPolicyApiResponseData() *OneOfUpdateVirusScanPolicyApiResponseData {
  p := new(OneOfUpdateVirusScanPolicyApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfUpdateVirusScanPolicyApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfUpdateVirusScanPolicyApiResponseData is nil"))
  }
  switch v.(type) {
    case import2.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = v.(import2.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case VirusScanPolicy:
      if nil == p.oneOfType0 {p.oneOfType0 = new(VirusScanPolicy)}
      *p.oneOfType0 = v.(VirusScanPolicy)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfUpdateVirusScanPolicyApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  return nil
}

func (p *OneOfUpdateVirusScanPolicyApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new(VirusScanPolicy)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "files.v4.config.VirusScanPolicy" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(VirusScanPolicy)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateVirusScanPolicyApiResponseData"))
}

func (p *OneOfUpdateVirusScanPolicyApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfUpdateVirusScanPolicyApiResponseData")
}

type OneOfUpdateDataConsolidationPolicyApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *import4.TaskReference `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfUpdateDataConsolidationPolicyApiResponseData() *OneOfUpdateDataConsolidationPolicyApiResponseData {
  p := new(OneOfUpdateDataConsolidationPolicyApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfUpdateDataConsolidationPolicyApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfUpdateDataConsolidationPolicyApiResponseData is nil"))
  }
  switch v.(type) {
    case import4.TaskReference:
      if nil == p.oneOfType0 {p.oneOfType0 = new(import4.TaskReference)}
      *p.oneOfType0 = v.(import4.TaskReference)
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

func (p *OneOfUpdateDataConsolidationPolicyApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfUpdateDataConsolidationPolicyApiResponseData) UnmarshalJSON(b []byte) error {
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
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateDataConsolidationPolicyApiResponseData"))
}

func (p *OneOfUpdateDataConsolidationPolicyApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfUpdateDataConsolidationPolicyApiResponseData")
}

type OneOfCreateDataConsolidationPolicyApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *import4.TaskReference `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfCreateDataConsolidationPolicyApiResponseData() *OneOfCreateDataConsolidationPolicyApiResponseData {
  p := new(OneOfCreateDataConsolidationPolicyApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfCreateDataConsolidationPolicyApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfCreateDataConsolidationPolicyApiResponseData is nil"))
  }
  switch v.(type) {
    case import4.TaskReference:
      if nil == p.oneOfType0 {p.oneOfType0 = new(import4.TaskReference)}
      *p.oneOfType0 = v.(import4.TaskReference)
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

func (p *OneOfCreateDataConsolidationPolicyApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfCreateDataConsolidationPolicyApiResponseData) UnmarshalJSON(b []byte) error {
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
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateDataConsolidationPolicyApiResponseData"))
}

func (p *OneOfCreateDataConsolidationPolicyApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfCreateDataConsolidationPolicyApiResponseData")
}

type OneOfUpdateSsrSnapshotScheduleApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *SsrSnapshotSchedule `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfUpdateSsrSnapshotScheduleApiResponseData() *OneOfUpdateSsrSnapshotScheduleApiResponseData {
  p := new(OneOfUpdateSsrSnapshotScheduleApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfUpdateSsrSnapshotScheduleApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfUpdateSsrSnapshotScheduleApiResponseData is nil"))
  }
  switch v.(type) {
    case SsrSnapshotSchedule:
      if nil == p.oneOfType0 {p.oneOfType0 = new(SsrSnapshotSchedule)}
      *p.oneOfType0 = v.(SsrSnapshotSchedule)
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

func (p *OneOfUpdateSsrSnapshotScheduleApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfUpdateSsrSnapshotScheduleApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new(SsrSnapshotSchedule)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "files.v4.config.SsrSnapshotSchedule" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(SsrSnapshotSchedule)}
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
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateSsrSnapshotScheduleApiResponseData"))
}

func (p *OneOfUpdateSsrSnapshotScheduleApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfUpdateSsrSnapshotScheduleApiResponseData")
}

type OneOfCreateNetworkApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *import4.TaskReference `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfCreateNetworkApiResponseData() *OneOfCreateNetworkApiResponseData {
  p := new(OneOfCreateNetworkApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfCreateNetworkApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfCreateNetworkApiResponseData is nil"))
  }
  switch v.(type) {
    case import4.TaskReference:
      if nil == p.oneOfType0 {p.oneOfType0 = new(import4.TaskReference)}
      *p.oneOfType0 = v.(import4.TaskReference)
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

func (p *OneOfCreateNetworkApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfCreateNetworkApiResponseData) UnmarshalJSON(b []byte) error {
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
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateNetworkApiResponseData"))
}

func (p *OneOfCreateNetworkApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfCreateNetworkApiResponseData")
}

type OneOfListObjectStoreApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
  oneOfType0 []TieringObjectStoreProfile `json:"-"`
}

func NewOneOfListObjectStoreApiResponseData() *OneOfListObjectStoreApiResponseData {
  p := new(OneOfListObjectStoreApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfListObjectStoreApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfListObjectStoreApiResponseData is nil"))
  }
  switch v.(type) {
    case import2.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = v.(import2.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case []TieringObjectStoreProfile:
      p.oneOfType0 = v.([]TieringObjectStoreProfile)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<files.v4.config.TieringObjectStoreProfile>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<files.v4.config.TieringObjectStoreProfile>"
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfListObjectStoreApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if "List<files.v4.config.TieringObjectStoreProfile>" == *p.Discriminator {
    return p.oneOfType0
  }
  return nil
}

func (p *OneOfListObjectStoreApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new([]TieringObjectStoreProfile)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "files.v4.config.TieringObjectStoreProfile" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<files.v4.config.TieringObjectStoreProfile>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<files.v4.config.TieringObjectStoreProfile>"
      return nil

    }
  }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListObjectStoreApiResponseData"))
}

func (p *OneOfListObjectStoreApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if "List<files.v4.config.TieringObjectStoreProfile>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfListObjectStoreApiResponseData")
}

type OneOfVirusScanPolicyApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
  oneOfType0 *VirusScanPolicy `json:"-"`
}

func NewOneOfVirusScanPolicyApiResponseData() *OneOfVirusScanPolicyApiResponseData {
  p := new(OneOfVirusScanPolicyApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfVirusScanPolicyApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfVirusScanPolicyApiResponseData is nil"))
  }
  switch v.(type) {
    case import2.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = v.(import2.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case VirusScanPolicy:
      if nil == p.oneOfType0 {p.oneOfType0 = new(VirusScanPolicy)}
      *p.oneOfType0 = v.(VirusScanPolicy)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfVirusScanPolicyApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  return nil
}

func (p *OneOfVirusScanPolicyApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new(VirusScanPolicy)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "files.v4.config.VirusScanPolicy" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(VirusScanPolicy)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfVirusScanPolicyApiResponseData"))
}

func (p *OneOfVirusScanPolicyApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfVirusScanPolicyApiResponseData")
}

type OneOfDeleteSsrSnapshotScheduleApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType1 *interface{} `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfDeleteSsrSnapshotScheduleApiResponseData() *OneOfDeleteSsrSnapshotScheduleApiResponseData {
  p := new(OneOfDeleteSsrSnapshotScheduleApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfDeleteSsrSnapshotScheduleApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfDeleteSsrSnapshotScheduleApiResponseData is nil"))
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

func (p *OneOfDeleteSsrSnapshotScheduleApiResponseData) GetValue() interface{} {
  if "EMPTY" == *p.Discriminator {
    return *p.oneOfType1
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfDeleteSsrSnapshotScheduleApiResponseData) UnmarshalJSON(b []byte) error {
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
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteSsrSnapshotScheduleApiResponseData"))
}

func (p *OneOfDeleteSsrSnapshotScheduleApiResponseData) MarshalJSON() ([]byte, error) {
  if "EMPTY" == *p.Discriminator {
    return json.Marshal(p.oneOfType1)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfDeleteSsrSnapshotScheduleApiResponseData")
}

type OneOfQuarantineInfectedFilesApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *import4.TaskReference `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfQuarantineInfectedFilesApiResponseData() *OneOfQuarantineInfectedFilesApiResponseData {
  p := new(OneOfQuarantineInfectedFilesApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfQuarantineInfectedFilesApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfQuarantineInfectedFilesApiResponseData is nil"))
  }
  switch v.(type) {
    case import4.TaskReference:
      if nil == p.oneOfType0 {p.oneOfType0 = new(import4.TaskReference)}
      *p.oneOfType0 = v.(import4.TaskReference)
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

func (p *OneOfQuarantineInfectedFilesApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfQuarantineInfectedFilesApiResponseData) UnmarshalJSON(b []byte) error {
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
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfQuarantineInfectedFilesApiResponseData"))
}

func (p *OneOfQuarantineInfectedFilesApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfQuarantineInfectedFilesApiResponseData")
}

type OneOfCreateSsrSnapshotScheduleApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *SsrSnapshotSchedule `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfCreateSsrSnapshotScheduleApiResponseData() *OneOfCreateSsrSnapshotScheduleApiResponseData {
  p := new(OneOfCreateSsrSnapshotScheduleApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfCreateSsrSnapshotScheduleApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfCreateSsrSnapshotScheduleApiResponseData is nil"))
  }
  switch v.(type) {
    case SsrSnapshotSchedule:
      if nil == p.oneOfType0 {p.oneOfType0 = new(SsrSnapshotSchedule)}
      *p.oneOfType0 = v.(SsrSnapshotSchedule)
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

func (p *OneOfCreateSsrSnapshotScheduleApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfCreateSsrSnapshotScheduleApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new(SsrSnapshotSchedule)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "files.v4.config.SsrSnapshotSchedule" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(SsrSnapshotSchedule)}
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
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateSsrSnapshotScheduleApiResponseData"))
}

func (p *OneOfCreateSsrSnapshotScheduleApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfCreateSsrSnapshotScheduleApiResponseData")
}

type OneOfUpdateRansomwareConfigApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *RansomwareConfig `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfUpdateRansomwareConfigApiResponseData() *OneOfUpdateRansomwareConfigApiResponseData {
  p := new(OneOfUpdateRansomwareConfigApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfUpdateRansomwareConfigApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfUpdateRansomwareConfigApiResponseData is nil"))
  }
  switch v.(type) {
    case RansomwareConfig:
      if nil == p.oneOfType0 {p.oneOfType0 = new(RansomwareConfig)}
      *p.oneOfType0 = v.(RansomwareConfig)
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

func (p *OneOfUpdateRansomwareConfigApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfUpdateRansomwareConfigApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new(RansomwareConfig)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "files.v4.config.RansomwareConfig" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(RansomwareConfig)}
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
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateRansomwareConfigApiResponseData"))
}

func (p *OneOfUpdateRansomwareConfigApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfUpdateRansomwareConfigApiResponseData")
}

type OneOfGetTieringConfigApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *TieringConfig `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfGetTieringConfigApiResponseData() *OneOfGetTieringConfigApiResponseData {
  p := new(OneOfGetTieringConfigApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfGetTieringConfigApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfGetTieringConfigApiResponseData is nil"))
  }
  switch v.(type) {
    case TieringConfig:
      if nil == p.oneOfType0 {p.oneOfType0 = new(TieringConfig)}
      *p.oneOfType0 = v.(TieringConfig)
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

func (p *OneOfGetTieringConfigApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfGetTieringConfigApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new(TieringConfig)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "files.v4.config.TieringConfig" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(TieringConfig)}
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
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetTieringConfigApiResponseData"))
}

func (p *OneOfGetTieringConfigApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfGetTieringConfigApiResponseData")
}

type OneOfClearLoadBalanceFileServerApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType1 *interface{} `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfClearLoadBalanceFileServerApiResponseData() *OneOfClearLoadBalanceFileServerApiResponseData {
  p := new(OneOfClearLoadBalanceFileServerApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfClearLoadBalanceFileServerApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfClearLoadBalanceFileServerApiResponseData is nil"))
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

func (p *OneOfClearLoadBalanceFileServerApiResponseData) GetValue() interface{} {
  if "EMPTY" == *p.Discriminator {
    return *p.oneOfType1
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfClearLoadBalanceFileServerApiResponseData) UnmarshalJSON(b []byte) error {
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
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfClearLoadBalanceFileServerApiResponseData"))
}

func (p *OneOfClearLoadBalanceFileServerApiResponseData) MarshalJSON() ([]byte, error) {
  if "EMPTY" == *p.Discriminator {
    return json.Marshal(p.oneOfType1)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfClearLoadBalanceFileServerApiResponseData")
}

type OneOfAlertApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
  oneOfType0 []Alert `json:"-"`
}

func NewOneOfAlertApiResponseData() *OneOfAlertApiResponseData {
  p := new(OneOfAlertApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfAlertApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfAlertApiResponseData is nil"))
  }
  switch v.(type) {
    case import2.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = v.(import2.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case []Alert:
      p.oneOfType0 = v.([]Alert)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<files.v4.config.Alert>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<files.v4.config.Alert>"
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfAlertApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if "List<files.v4.config.Alert>" == *p.Discriminator {
    return p.oneOfType0
  }
  return nil
}

func (p *OneOfAlertApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new([]Alert)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "files.v4.config.Alert" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<files.v4.config.Alert>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<files.v4.config.Alert>"
      return nil

    }
  }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfAlertApiResponseData"))
}

func (p *OneOfAlertApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if "List<files.v4.config.Alert>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfAlertApiResponseData")
}

type OneOfUpdateDataProtectionPolicyApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *import4.TaskReference `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfUpdateDataProtectionPolicyApiResponseData() *OneOfUpdateDataProtectionPolicyApiResponseData {
  p := new(OneOfUpdateDataProtectionPolicyApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfUpdateDataProtectionPolicyApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfUpdateDataProtectionPolicyApiResponseData is nil"))
  }
  switch v.(type) {
    case import4.TaskReference:
      if nil == p.oneOfType0 {p.oneOfType0 = new(import4.TaskReference)}
      *p.oneOfType0 = v.(import4.TaskReference)
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

func (p *OneOfUpdateDataProtectionPolicyApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfUpdateDataProtectionPolicyApiResponseData) UnmarshalJSON(b []byte) error {
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
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateDataProtectionPolicyApiResponseData"))
}

func (p *OneOfUpdateDataProtectionPolicyApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfUpdateDataProtectionPolicyApiResponseData")
}

type OneOfUpdateFileServerUserApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *FileServerUser `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfUpdateFileServerUserApiResponseData() *OneOfUpdateFileServerUserApiResponseData {
  p := new(OneOfUpdateFileServerUserApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfUpdateFileServerUserApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfUpdateFileServerUserApiResponseData is nil"))
  }
  switch v.(type) {
    case FileServerUser:
      if nil == p.oneOfType0 {p.oneOfType0 = new(FileServerUser)}
      *p.oneOfType0 = v.(FileServerUser)
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

func (p *OneOfUpdateFileServerUserApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfUpdateFileServerUserApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new(FileServerUser)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "files.v4.config.FileServerUser" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(FileServerUser)}
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
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateFileServerUserApiResponseData"))
}

func (p *OneOfUpdateFileServerUserApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfUpdateFileServerUserApiResponseData")
}

type OneOfCreateMountTargetVirusScanPolicyApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *VirusScanPolicyMountTarget `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfCreateMountTargetVirusScanPolicyApiResponseData() *OneOfCreateMountTargetVirusScanPolicyApiResponseData {
  p := new(OneOfCreateMountTargetVirusScanPolicyApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfCreateMountTargetVirusScanPolicyApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfCreateMountTargetVirusScanPolicyApiResponseData is nil"))
  }
  switch v.(type) {
    case VirusScanPolicyMountTarget:
      if nil == p.oneOfType0 {p.oneOfType0 = new(VirusScanPolicyMountTarget)}
      *p.oneOfType0 = v.(VirusScanPolicyMountTarget)
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

func (p *OneOfCreateMountTargetVirusScanPolicyApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfCreateMountTargetVirusScanPolicyApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new(VirusScanPolicyMountTarget)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "files.v4.config.VirusScanPolicyMountTarget" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(VirusScanPolicyMountTarget)}
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
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateMountTargetVirusScanPolicyApiResponseData"))
}

func (p *OneOfCreateMountTargetVirusScanPolicyApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfCreateMountTargetVirusScanPolicyApiResponseData")
}

type OneOfTierDataApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType1 *interface{} `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfTierDataApiResponseData() *OneOfTierDataApiResponseData {
  p := new(OneOfTierDataApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfTierDataApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfTierDataApiResponseData is nil"))
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

func (p *OneOfTierDataApiResponseData) GetValue() interface{} {
  if "EMPTY" == *p.Discriminator {
    return *p.oneOfType1
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfTierDataApiResponseData) UnmarshalJSON(b []byte) error {
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
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfTierDataApiResponseData"))
}

func (p *OneOfTierDataApiResponseData) MarshalJSON() ([]byte, error) {
  if "EMPTY" == *p.Discriminator {
    return json.Marshal(p.oneOfType1)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfTierDataApiResponseData")
}

type OneOfMountTargetApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
  oneOfType0 *MountTarget `json:"-"`
}

func NewOneOfMountTargetApiResponseData() *OneOfMountTargetApiResponseData {
  p := new(OneOfMountTargetApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfMountTargetApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfMountTargetApiResponseData is nil"))
  }
  switch v.(type) {
    case import2.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = v.(import2.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case MountTarget:
      if nil == p.oneOfType0 {p.oneOfType0 = new(MountTarget)}
      *p.oneOfType0 = v.(MountTarget)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfMountTargetApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  return nil
}

func (p *OneOfMountTargetApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new(MountTarget)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "files.v4.config.MountTarget" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(MountTarget)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfMountTargetApiResponseData"))
}

func (p *OneOfMountTargetApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfMountTargetApiResponseData")
}

type OneOfDataProtectionPolicyApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *DataProtectionPolicy `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfDataProtectionPolicyApiResponseData() *OneOfDataProtectionPolicyApiResponseData {
  p := new(OneOfDataProtectionPolicyApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfDataProtectionPolicyApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfDataProtectionPolicyApiResponseData is nil"))
  }
  switch v.(type) {
    case DataProtectionPolicy:
      if nil == p.oneOfType0 {p.oneOfType0 = new(DataProtectionPolicy)}
      *p.oneOfType0 = v.(DataProtectionPolicy)
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

func (p *OneOfDataProtectionPolicyApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfDataProtectionPolicyApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new(DataProtectionPolicy)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "files.v4.config.DataProtectionPolicy" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(DataProtectionPolicy)}
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
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDataProtectionPolicyApiResponseData"))
}

func (p *OneOfDataProtectionPolicyApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfDataProtectionPolicyApiResponseData")
}

type OneOfDataConsolidationPolicyApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
  oneOfType0 *DataConsolidationPolicy `json:"-"`
}

func NewOneOfDataConsolidationPolicyApiResponseData() *OneOfDataConsolidationPolicyApiResponseData {
  p := new(OneOfDataConsolidationPolicyApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfDataConsolidationPolicyApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfDataConsolidationPolicyApiResponseData is nil"))
  }
  switch v.(type) {
    case import2.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = v.(import2.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case DataConsolidationPolicy:
      if nil == p.oneOfType0 {p.oneOfType0 = new(DataConsolidationPolicy)}
      *p.oneOfType0 = v.(DataConsolidationPolicy)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfDataConsolidationPolicyApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  return nil
}

func (p *OneOfDataConsolidationPolicyApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new(DataConsolidationPolicy)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "files.v4.config.DataConsolidationPolicy" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(DataConsolidationPolicy)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDataConsolidationPolicyApiResponseData"))
}

func (p *OneOfDataConsolidationPolicyApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfDataConsolidationPolicyApiResponseData")
}

type OneOfMountTargetListApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 []MountTarget `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfMountTargetListApiResponseData() *OneOfMountTargetListApiResponseData {
  p := new(OneOfMountTargetListApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfMountTargetListApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfMountTargetListApiResponseData is nil"))
  }
  switch v.(type) {
    case []MountTarget:
      p.oneOfType0 = v.([]MountTarget)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<files.v4.config.MountTarget>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<files.v4.config.MountTarget>"
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

func (p *OneOfMountTargetListApiResponseData) GetValue() interface{} {
  if "List<files.v4.config.MountTarget>" == *p.Discriminator {
    return p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfMountTargetListApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new([]MountTarget)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "files.v4.config.MountTarget" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<files.v4.config.MountTarget>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<files.v4.config.MountTarget>"
      return nil

    }
  }
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfMountTargetListApiResponseData"))
}

func (p *OneOfMountTargetListApiResponseData) MarshalJSON() ([]byte, error) {
  if "List<files.v4.config.MountTarget>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfMountTargetListApiResponseData")
}

type OneOfCreateQuotaPolicyApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *import4.TaskReference `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfCreateQuotaPolicyApiResponseData() *OneOfCreateQuotaPolicyApiResponseData {
  p := new(OneOfCreateQuotaPolicyApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfCreateQuotaPolicyApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfCreateQuotaPolicyApiResponseData is nil"))
  }
  switch v.(type) {
    case import4.TaskReference:
      if nil == p.oneOfType0 {p.oneOfType0 = new(import4.TaskReference)}
      *p.oneOfType0 = v.(import4.TaskReference)
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

func (p *OneOfCreateQuotaPolicyApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfCreateQuotaPolicyApiResponseData) UnmarshalJSON(b []byte) error {
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
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateQuotaPolicyApiResponseData"))
}

func (p *OneOfCreateQuotaPolicyApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfCreateQuotaPolicyApiResponseData")
}

type OneOfInfectedFilesApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
  oneOfType0 *InfectedFile `json:"-"`
}

func NewOneOfInfectedFilesApiResponseData() *OneOfInfectedFilesApiResponseData {
  p := new(OneOfInfectedFilesApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfInfectedFilesApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfInfectedFilesApiResponseData is nil"))
  }
  switch v.(type) {
    case import2.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = v.(import2.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case InfectedFile:
      if nil == p.oneOfType0 {p.oneOfType0 = new(InfectedFile)}
      *p.oneOfType0 = v.(InfectedFile)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfInfectedFilesApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  return nil
}

func (p *OneOfInfectedFilesApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new(InfectedFile)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "files.v4.config.InfectedFile" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(InfectedFile)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfInfectedFilesApiResponseData"))
}

func (p *OneOfInfectedFilesApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfInfectedFilesApiResponseData")
}

type OneOfDeleteMountTargetVirusScanPolicyApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType1 *interface{} `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfDeleteMountTargetVirusScanPolicyApiResponseData() *OneOfDeleteMountTargetVirusScanPolicyApiResponseData {
  p := new(OneOfDeleteMountTargetVirusScanPolicyApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfDeleteMountTargetVirusScanPolicyApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfDeleteMountTargetVirusScanPolicyApiResponseData is nil"))
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

func (p *OneOfDeleteMountTargetVirusScanPolicyApiResponseData) GetValue() interface{} {
  if "EMPTY" == *p.Discriminator {
    return *p.oneOfType1
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfDeleteMountTargetVirusScanPolicyApiResponseData) UnmarshalJSON(b []byte) error {
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
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteMountTargetVirusScanPolicyApiResponseData"))
}

func (p *OneOfDeleteMountTargetVirusScanPolicyApiResponseData) MarshalJSON() ([]byte, error) {
  if "EMPTY" == *p.Discriminator {
    return json.Marshal(p.oneOfType1)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfDeleteMountTargetVirusScanPolicyApiResponseData")
}

type OneOfDeleteTieringPolicyApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *import4.TaskReference `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfDeleteTieringPolicyApiResponseData() *OneOfDeleteTieringPolicyApiResponseData {
  p := new(OneOfDeleteTieringPolicyApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfDeleteTieringPolicyApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfDeleteTieringPolicyApiResponseData is nil"))
  }
  switch v.(type) {
    case import4.TaskReference:
      if nil == p.oneOfType0 {p.oneOfType0 = new(import4.TaskReference)}
      *p.oneOfType0 = v.(import4.TaskReference)
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

func (p *OneOfDeleteTieringPolicyApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfDeleteTieringPolicyApiResponseData) UnmarshalJSON(b []byte) error {
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
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteTieringPolicyApiResponseData"))
}

func (p *OneOfDeleteTieringPolicyApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfDeleteTieringPolicyApiResponseData")
}

type OneOfDnsEntriesListApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
  oneOfType0 []DnsDetails `json:"-"`
}

func NewOneOfDnsEntriesListApiResponseData() *OneOfDnsEntriesListApiResponseData {
  p := new(OneOfDnsEntriesListApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfDnsEntriesListApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfDnsEntriesListApiResponseData is nil"))
  }
  switch v.(type) {
    case import2.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = v.(import2.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case []DnsDetails:
      p.oneOfType0 = v.([]DnsDetails)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<files.v4.config.DnsDetails>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<files.v4.config.DnsDetails>"
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfDnsEntriesListApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if "List<files.v4.config.DnsDetails>" == *p.Discriminator {
    return p.oneOfType0
  }
  return nil
}

func (p *OneOfDnsEntriesListApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new([]DnsDetails)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "files.v4.config.DnsDetails" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<files.v4.config.DnsDetails>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<files.v4.config.DnsDetails>"
      return nil

    }
  }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDnsEntriesListApiResponseData"))
}

func (p *OneOfDnsEntriesListApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if "List<files.v4.config.DnsDetails>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfDnsEntriesListApiResponseData")
}

type OneOfDataProtectionPolicyListApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
  oneOfType0 []DataProtectionPolicy `json:"-"`
}

func NewOneOfDataProtectionPolicyListApiResponseData() *OneOfDataProtectionPolicyListApiResponseData {
  p := new(OneOfDataProtectionPolicyListApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfDataProtectionPolicyListApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfDataProtectionPolicyListApiResponseData is nil"))
  }
  switch v.(type) {
    case import2.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = v.(import2.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case []DataProtectionPolicy:
      p.oneOfType0 = v.([]DataProtectionPolicy)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<files.v4.config.DataProtectionPolicy>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<files.v4.config.DataProtectionPolicy>"
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfDataProtectionPolicyListApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if "List<files.v4.config.DataProtectionPolicy>" == *p.Discriminator {
    return p.oneOfType0
  }
  return nil
}

func (p *OneOfDataProtectionPolicyListApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new([]DataProtectionPolicy)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "files.v4.config.DataProtectionPolicy" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<files.v4.config.DataProtectionPolicy>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<files.v4.config.DataProtectionPolicy>"
      return nil

    }
  }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDataProtectionPolicyListApiResponseData"))
}

func (p *OneOfDataProtectionPolicyListApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if "List<files.v4.config.DataProtectionPolicy>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfDataProtectionPolicyListApiResponseData")
}

type OneOfGetTieringConfigListApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 []TieringConfig `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfGetTieringConfigListApiResponseData() *OneOfGetTieringConfigListApiResponseData {
  p := new(OneOfGetTieringConfigListApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfGetTieringConfigListApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfGetTieringConfigListApiResponseData is nil"))
  }
  switch v.(type) {
    case []TieringConfig:
      p.oneOfType0 = v.([]TieringConfig)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<files.v4.config.TieringConfig>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<files.v4.config.TieringConfig>"
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

func (p *OneOfGetTieringConfigListApiResponseData) GetValue() interface{} {
  if "List<files.v4.config.TieringConfig>" == *p.Discriminator {
    return p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfGetTieringConfigListApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new([]TieringConfig)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "files.v4.config.TieringConfig" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<files.v4.config.TieringConfig>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<files.v4.config.TieringConfig>"
      return nil

    }
  }
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetTieringConfigListApiResponseData"))
}

func (p *OneOfGetTieringConfigListApiResponseData) MarshalJSON() ([]byte, error) {
  if "List<files.v4.config.TieringConfig>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfGetTieringConfigListApiResponseData")
}

type OneOfQuotaPolicyApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
  oneOfType0 *QuotaPolicy `json:"-"`
}

func NewOneOfQuotaPolicyApiResponseData() *OneOfQuotaPolicyApiResponseData {
  p := new(OneOfQuotaPolicyApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfQuotaPolicyApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfQuotaPolicyApiResponseData is nil"))
  }
  switch v.(type) {
    case import2.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = v.(import2.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case QuotaPolicy:
      if nil == p.oneOfType0 {p.oneOfType0 = new(QuotaPolicy)}
      *p.oneOfType0 = v.(QuotaPolicy)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfQuotaPolicyApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  return nil
}

func (p *OneOfQuotaPolicyApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new(QuotaPolicy)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "files.v4.config.QuotaPolicy" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(QuotaPolicy)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfQuotaPolicyApiResponseData"))
}

func (p *OneOfQuotaPolicyApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfQuotaPolicyApiResponseData")
}

type OneOfUpdateQuotaPolicyApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *import4.TaskReference `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfUpdateQuotaPolicyApiResponseData() *OneOfUpdateQuotaPolicyApiResponseData {
  p := new(OneOfUpdateQuotaPolicyApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfUpdateQuotaPolicyApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfUpdateQuotaPolicyApiResponseData is nil"))
  }
  switch v.(type) {
    case import4.TaskReference:
      if nil == p.oneOfType0 {p.oneOfType0 = new(import4.TaskReference)}
      *p.oneOfType0 = v.(import4.TaskReference)
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

func (p *OneOfUpdateQuotaPolicyApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfUpdateQuotaPolicyApiResponseData) UnmarshalJSON(b []byte) error {
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
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateQuotaPolicyApiResponseData"))
}

func (p *OneOfUpdateQuotaPolicyApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfUpdateQuotaPolicyApiResponseData")
}

type OneOfSsrSnapshotScheduleApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *SsrSnapshotSchedule `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfSsrSnapshotScheduleApiResponseData() *OneOfSsrSnapshotScheduleApiResponseData {
  p := new(OneOfSsrSnapshotScheduleApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfSsrSnapshotScheduleApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfSsrSnapshotScheduleApiResponseData is nil"))
  }
  switch v.(type) {
    case SsrSnapshotSchedule:
      if nil == p.oneOfType0 {p.oneOfType0 = new(SsrSnapshotSchedule)}
      *p.oneOfType0 = v.(SsrSnapshotSchedule)
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

func (p *OneOfSsrSnapshotScheduleApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfSsrSnapshotScheduleApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new(SsrSnapshotSchedule)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "files.v4.config.SsrSnapshotSchedule" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(SsrSnapshotSchedule)}
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
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfSsrSnapshotScheduleApiResponseData"))
}

func (p *OneOfSsrSnapshotScheduleApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfSsrSnapshotScheduleApiResponseData")
}

type OneOfCreateFileServerUserApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *FileServerUser `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfCreateFileServerUserApiResponseData() *OneOfCreateFileServerUserApiResponseData {
  p := new(OneOfCreateFileServerUserApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfCreateFileServerUserApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfCreateFileServerUserApiResponseData is nil"))
  }
  switch v.(type) {
    case FileServerUser:
      if nil == p.oneOfType0 {p.oneOfType0 = new(FileServerUser)}
      *p.oneOfType0 = v.(FileServerUser)
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

func (p *OneOfCreateFileServerUserApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfCreateFileServerUserApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new(FileServerUser)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "files.v4.config.FileServerUser" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(FileServerUser)}
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
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateFileServerUserApiResponseData"))
}

func (p *OneOfCreateFileServerUserApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfCreateFileServerUserApiResponseData")
}

type OneOfUpdateQuotaEmailConfigApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *QuotaEmailConfig `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfUpdateQuotaEmailConfigApiResponseData() *OneOfUpdateQuotaEmailConfigApiResponseData {
  p := new(OneOfUpdateQuotaEmailConfigApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfUpdateQuotaEmailConfigApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfUpdateQuotaEmailConfigApiResponseData is nil"))
  }
  switch v.(type) {
    case QuotaEmailConfig:
      if nil == p.oneOfType0 {p.oneOfType0 = new(QuotaEmailConfig)}
      *p.oneOfType0 = v.(QuotaEmailConfig)
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

func (p *OneOfUpdateQuotaEmailConfigApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfUpdateQuotaEmailConfigApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new(QuotaEmailConfig)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "files.v4.config.QuotaEmailConfig" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(QuotaEmailConfig)}
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
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateQuotaEmailConfigApiResponseData"))
}

func (p *OneOfUpdateQuotaEmailConfigApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfUpdateQuotaEmailConfigApiResponseData")
}

type OneOfUserMappingListApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
  oneOfType0 []UserMapping `json:"-"`
}

func NewOneOfUserMappingListApiResponseData() *OneOfUserMappingListApiResponseData {
  p := new(OneOfUserMappingListApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfUserMappingListApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfUserMappingListApiResponseData is nil"))
  }
  switch v.(type) {
    case import2.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = v.(import2.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case []UserMapping:
      p.oneOfType0 = v.([]UserMapping)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<files.v4.config.UserMapping>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<files.v4.config.UserMapping>"
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfUserMappingListApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if "List<files.v4.config.UserMapping>" == *p.Discriminator {
    return p.oneOfType0
  }
  return nil
}

func (p *OneOfUserMappingListApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new([]UserMapping)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "files.v4.config.UserMapping" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<files.v4.config.UserMapping>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<files.v4.config.UserMapping>"
      return nil

    }
  }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUserMappingListApiResponseData"))
}

func (p *OneOfUserMappingListApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if "List<files.v4.config.UserMapping>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfUserMappingListApiResponseData")
}

type OneOfGetObjectStoreApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *TieringObjectStoreProfile `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfGetObjectStoreApiResponseData() *OneOfGetObjectStoreApiResponseData {
  p := new(OneOfGetObjectStoreApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfGetObjectStoreApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfGetObjectStoreApiResponseData is nil"))
  }
  switch v.(type) {
    case TieringObjectStoreProfile:
      if nil == p.oneOfType0 {p.oneOfType0 = new(TieringObjectStoreProfile)}
      *p.oneOfType0 = v.(TieringObjectStoreProfile)
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

func (p *OneOfGetObjectStoreApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfGetObjectStoreApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new(TieringObjectStoreProfile)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "files.v4.config.TieringObjectStoreProfile" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(TieringObjectStoreProfile)}
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
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetObjectStoreApiResponseData"))
}

func (p *OneOfGetObjectStoreApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfGetObjectStoreApiResponseData")
}

type OneOfLoadBalanceFileServerApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *import4.TaskReference `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfLoadBalanceFileServerApiResponseData() *OneOfLoadBalanceFileServerApiResponseData {
  p := new(OneOfLoadBalanceFileServerApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfLoadBalanceFileServerApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfLoadBalanceFileServerApiResponseData is nil"))
  }
  switch v.(type) {
    case import4.TaskReference:
      if nil == p.oneOfType0 {p.oneOfType0 = new(import4.TaskReference)}
      *p.oneOfType0 = v.(import4.TaskReference)
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

func (p *OneOfLoadBalanceFileServerApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfLoadBalanceFileServerApiResponseData) UnmarshalJSON(b []byte) error {
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
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfLoadBalanceFileServerApiResponseData"))
}

func (p *OneOfLoadBalanceFileServerApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfLoadBalanceFileServerApiResponseData")
}

type OneOfAcknowledgeAlertsApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
  oneOfType0 []AcknowledgeAlertsResponse `json:"-"`
}

func NewOneOfAcknowledgeAlertsApiResponseData() *OneOfAcknowledgeAlertsApiResponseData {
  p := new(OneOfAcknowledgeAlertsApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfAcknowledgeAlertsApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfAcknowledgeAlertsApiResponseData is nil"))
  }
  switch v.(type) {
    case import2.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = v.(import2.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case []AcknowledgeAlertsResponse:
      p.oneOfType0 = v.([]AcknowledgeAlertsResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<files.v4.config.AcknowledgeAlertsResponse>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<files.v4.config.AcknowledgeAlertsResponse>"
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfAcknowledgeAlertsApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if "List<files.v4.config.AcknowledgeAlertsResponse>" == *p.Discriminator {
    return p.oneOfType0
  }
  return nil
}

func (p *OneOfAcknowledgeAlertsApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new([]AcknowledgeAlertsResponse)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "files.v4.config.AcknowledgeAlertsResponse" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<files.v4.config.AcknowledgeAlertsResponse>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<files.v4.config.AcknowledgeAlertsResponse>"
      return nil

    }
  }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfAcknowledgeAlertsApiResponseData"))
}

func (p *OneOfAcknowledgeAlertsApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if "List<files.v4.config.AcknowledgeAlertsResponse>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfAcknowledgeAlertsApiResponseData")
}

type OneOfAdminUserApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
  oneOfType0 *AdminUser `json:"-"`
}

func NewOneOfAdminUserApiResponseData() *OneOfAdminUserApiResponseData {
  p := new(OneOfAdminUserApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfAdminUserApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfAdminUserApiResponseData is nil"))
  }
  switch v.(type) {
    case import2.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = v.(import2.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case AdminUser:
      if nil == p.oneOfType0 {p.oneOfType0 = new(AdminUser)}
      *p.oneOfType0 = v.(AdminUser)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfAdminUserApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  return nil
}

func (p *OneOfAdminUserApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new(AdminUser)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "files.v4.config.AdminUser" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(AdminUser)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfAdminUserApiResponseData"))
}

func (p *OneOfAdminUserApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfAdminUserApiResponseData")
}

type OneOfConfigureNameServiceApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *import4.TaskReference `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfConfigureNameServiceApiResponseData() *OneOfConfigureNameServiceApiResponseData {
  p := new(OneOfConfigureNameServiceApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfConfigureNameServiceApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfConfigureNameServiceApiResponseData is nil"))
  }
  switch v.(type) {
    case import4.TaskReference:
      if nil == p.oneOfType0 {p.oneOfType0 = new(import4.TaskReference)}
      *p.oneOfType0 = v.(import4.TaskReference)
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

func (p *OneOfConfigureNameServiceApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfConfigureNameServiceApiResponseData) UnmarshalJSON(b []byte) error {
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
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfConfigureNameServiceApiResponseData"))
}

func (p *OneOfConfigureNameServiceApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfConfigureNameServiceApiResponseData")
}

type OneOfListVirusScanPolicyApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
  oneOfType0 []VirusScanPolicy `json:"-"`
}

func NewOneOfListVirusScanPolicyApiResponseData() *OneOfListVirusScanPolicyApiResponseData {
  p := new(OneOfListVirusScanPolicyApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfListVirusScanPolicyApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfListVirusScanPolicyApiResponseData is nil"))
  }
  switch v.(type) {
    case import2.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = v.(import2.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case []VirusScanPolicy:
      p.oneOfType0 = v.([]VirusScanPolicy)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<files.v4.config.VirusScanPolicy>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<files.v4.config.VirusScanPolicy>"
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfListVirusScanPolicyApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if "List<files.v4.config.VirusScanPolicy>" == *p.Discriminator {
    return p.oneOfType0
  }
  return nil
}

func (p *OneOfListVirusScanPolicyApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new([]VirusScanPolicy)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "files.v4.config.VirusScanPolicy" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<files.v4.config.VirusScanPolicy>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<files.v4.config.VirusScanPolicy>"
      return nil

    }
  }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListVirusScanPolicyApiResponseData"))
}

func (p *OneOfListVirusScanPolicyApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if "List<files.v4.config.VirusScanPolicy>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfListVirusScanPolicyApiResponseData")
}

type OneOfEnableDataProtectionPolicyApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *import4.TaskReference `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfEnableDataProtectionPolicyApiResponseData() *OneOfEnableDataProtectionPolicyApiResponseData {
  p := new(OneOfEnableDataProtectionPolicyApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfEnableDataProtectionPolicyApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfEnableDataProtectionPolicyApiResponseData is nil"))
  }
  switch v.(type) {
    case import4.TaskReference:
      if nil == p.oneOfType0 {p.oneOfType0 = new(import4.TaskReference)}
      *p.oneOfType0 = v.(import4.TaskReference)
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

func (p *OneOfEnableDataProtectionPolicyApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfEnableDataProtectionPolicyApiResponseData) UnmarshalJSON(b []byte) error {
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
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfEnableDataProtectionPolicyApiResponseData"))
}

func (p *OneOfEnableDataProtectionPolicyApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfEnableDataProtectionPolicyApiResponseData")
}

type OneOfMountTargetSnapshotApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
  oneOfType0 *Snapshot `json:"-"`
}

func NewOneOfMountTargetSnapshotApiResponseData() *OneOfMountTargetSnapshotApiResponseData {
  p := new(OneOfMountTargetSnapshotApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfMountTargetSnapshotApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfMountTargetSnapshotApiResponseData is nil"))
  }
  switch v.(type) {
    case import2.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = v.(import2.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case Snapshot:
      if nil == p.oneOfType0 {p.oneOfType0 = new(Snapshot)}
      *p.oneOfType0 = v.(Snapshot)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfMountTargetSnapshotApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  return nil
}

func (p *OneOfMountTargetSnapshotApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new(Snapshot)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "files.v4.config.Snapshot" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(Snapshot)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfMountTargetSnapshotApiResponseData"))
}

func (p *OneOfMountTargetSnapshotApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfMountTargetSnapshotApiResponseData")
}

type OneOfMountTargetVirusScanPolicyApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *VirusScanPolicyMountTarget `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfMountTargetVirusScanPolicyApiResponseData() *OneOfMountTargetVirusScanPolicyApiResponseData {
  p := new(OneOfMountTargetVirusScanPolicyApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfMountTargetVirusScanPolicyApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfMountTargetVirusScanPolicyApiResponseData is nil"))
  }
  switch v.(type) {
    case VirusScanPolicyMountTarget:
      if nil == p.oneOfType0 {p.oneOfType0 = new(VirusScanPolicyMountTarget)}
      *p.oneOfType0 = v.(VirusScanPolicyMountTarget)
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

func (p *OneOfMountTargetVirusScanPolicyApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfMountTargetVirusScanPolicyApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new(VirusScanPolicyMountTarget)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "files.v4.config.VirusScanPolicyMountTarget" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(VirusScanPolicyMountTarget)}
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
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfMountTargetVirusScanPolicyApiResponseData"))
}

func (p *OneOfMountTargetVirusScanPolicyApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfMountTargetVirusScanPolicyApiResponseData")
}

type OneOfUpdateGlobalBlockedClientApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *import4.TaskReference `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfUpdateGlobalBlockedClientApiResponseData() *OneOfUpdateGlobalBlockedClientApiResponseData {
  p := new(OneOfUpdateGlobalBlockedClientApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfUpdateGlobalBlockedClientApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfUpdateGlobalBlockedClientApiResponseData is nil"))
  }
  switch v.(type) {
    case import4.TaskReference:
      if nil == p.oneOfType0 {p.oneOfType0 = new(import4.TaskReference)}
      *p.oneOfType0 = v.(import4.TaskReference)
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

func (p *OneOfUpdateGlobalBlockedClientApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfUpdateGlobalBlockedClientApiResponseData) UnmarshalJSON(b []byte) error {
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
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateGlobalBlockedClientApiResponseData"))
}

func (p *OneOfUpdateGlobalBlockedClientApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfUpdateGlobalBlockedClientApiResponseData")
}

type OneOfPrincipalTypeByPrincipalValueApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
  oneOfType0 *PrincipalTypeByPrincipalValue `json:"-"`
}

func NewOneOfPrincipalTypeByPrincipalValueApiResponseData() *OneOfPrincipalTypeByPrincipalValueApiResponseData {
  p := new(OneOfPrincipalTypeByPrincipalValueApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfPrincipalTypeByPrincipalValueApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfPrincipalTypeByPrincipalValueApiResponseData is nil"))
  }
  switch v.(type) {
    case import2.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = v.(import2.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case PrincipalTypeByPrincipalValue:
      if nil == p.oneOfType0 {p.oneOfType0 = new(PrincipalTypeByPrincipalValue)}
      *p.oneOfType0 = v.(PrincipalTypeByPrincipalValue)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfPrincipalTypeByPrincipalValueApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  return nil
}

func (p *OneOfPrincipalTypeByPrincipalValueApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new(PrincipalTypeByPrincipalValue)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "files.v4.config.PrincipalTypeByPrincipalValue" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(PrincipalTypeByPrincipalValue)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfPrincipalTypeByPrincipalValueApiResponseData"))
}

func (p *OneOfPrincipalTypeByPrincipalValueApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfPrincipalTypeByPrincipalValueApiResponseData")
}

type OneOfRemoveInfectedFilesApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *import4.TaskReference `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfRemoveInfectedFilesApiResponseData() *OneOfRemoveInfectedFilesApiResponseData {
  p := new(OneOfRemoveInfectedFilesApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfRemoveInfectedFilesApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfRemoveInfectedFilesApiResponseData is nil"))
  }
  switch v.(type) {
    case import4.TaskReference:
      if nil == p.oneOfType0 {p.oneOfType0 = new(import4.TaskReference)}
      *p.oneOfType0 = v.(import4.TaskReference)
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

func (p *OneOfRemoveInfectedFilesApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfRemoveInfectedFilesApiResponseData) UnmarshalJSON(b []byte) error {
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
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfRemoveInfectedFilesApiResponseData"))
}

func (p *OneOfRemoveInfectedFilesApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfRemoveInfectedFilesApiResponseData")
}

type OneOfCreateDataProtectionPolicyApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *import4.TaskReference `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfCreateDataProtectionPolicyApiResponseData() *OneOfCreateDataProtectionPolicyApiResponseData {
  p := new(OneOfCreateDataProtectionPolicyApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfCreateDataProtectionPolicyApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfCreateDataProtectionPolicyApiResponseData is nil"))
  }
  switch v.(type) {
    case import4.TaskReference:
      if nil == p.oneOfType0 {p.oneOfType0 = new(import4.TaskReference)}
      *p.oneOfType0 = v.(import4.TaskReference)
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

func (p *OneOfCreateDataProtectionPolicyApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfCreateDataProtectionPolicyApiResponseData) UnmarshalJSON(b []byte) error {
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
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateDataProtectionPolicyApiResponseData"))
}

func (p *OneOfCreateDataProtectionPolicyApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfCreateDataProtectionPolicyApiResponseData")
}

type OneOfCreateUserMappingApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *import4.TaskReference `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfCreateUserMappingApiResponseData() *OneOfCreateUserMappingApiResponseData {
  p := new(OneOfCreateUserMappingApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfCreateUserMappingApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfCreateUserMappingApiResponseData is nil"))
  }
  switch v.(type) {
    case import4.TaskReference:
      if nil == p.oneOfType0 {p.oneOfType0 = new(import4.TaskReference)}
      *p.oneOfType0 = v.(import4.TaskReference)
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

func (p *OneOfCreateUserMappingApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfCreateUserMappingApiResponseData) UnmarshalJSON(b []byte) error {
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
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateUserMappingApiResponseData"))
}

func (p *OneOfCreateUserMappingApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfCreateUserMappingApiResponseData")
}

type OneOfUpdateObjectStoreApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *import4.TaskReference `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfUpdateObjectStoreApiResponseData() *OneOfUpdateObjectStoreApiResponseData {
  p := new(OneOfUpdateObjectStoreApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfUpdateObjectStoreApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfUpdateObjectStoreApiResponseData is nil"))
  }
  switch v.(type) {
    case import4.TaskReference:
      if nil == p.oneOfType0 {p.oneOfType0 = new(import4.TaskReference)}
      *p.oneOfType0 = v.(import4.TaskReference)
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

func (p *OneOfUpdateObjectStoreApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfUpdateObjectStoreApiResponseData) UnmarshalJSON(b []byte) error {
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
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateObjectStoreApiResponseData"))
}

func (p *OneOfUpdateObjectStoreApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfUpdateObjectStoreApiResponseData")
}

type OneOfCreateTieringPolicyApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *import4.TaskReference `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfCreateTieringPolicyApiResponseData() *OneOfCreateTieringPolicyApiResponseData {
  p := new(OneOfCreateTieringPolicyApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfCreateTieringPolicyApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfCreateTieringPolicyApiResponseData is nil"))
  }
  switch v.(type) {
    case import4.TaskReference:
      if nil == p.oneOfType0 {p.oneOfType0 = new(import4.TaskReference)}
      *p.oneOfType0 = v.(import4.TaskReference)
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

func (p *OneOfCreateTieringPolicyApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfCreateTieringPolicyApiResponseData) UnmarshalJSON(b []byte) error {
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
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateTieringPolicyApiResponseData"))
}

func (p *OneOfCreateTieringPolicyApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfCreateTieringPolicyApiResponseData")
}

type OneOfProtectedPairsListApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
  oneOfType0 []ProtectedFileServerPair `json:"-"`
}

func NewOneOfProtectedPairsListApiResponseData() *OneOfProtectedPairsListApiResponseData {
  p := new(OneOfProtectedPairsListApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfProtectedPairsListApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfProtectedPairsListApiResponseData is nil"))
  }
  switch v.(type) {
    case import2.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = v.(import2.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case []ProtectedFileServerPair:
      p.oneOfType0 = v.([]ProtectedFileServerPair)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<files.v4.config.ProtectedFileServerPair>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<files.v4.config.ProtectedFileServerPair>"
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfProtectedPairsListApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if "List<files.v4.config.ProtectedFileServerPair>" == *p.Discriminator {
    return p.oneOfType0
  }
  return nil
}

func (p *OneOfProtectedPairsListApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new([]ProtectedFileServerPair)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "files.v4.config.ProtectedFileServerPair" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<files.v4.config.ProtectedFileServerPair>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<files.v4.config.ProtectedFileServerPair>"
      return nil

    }
  }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfProtectedPairsListApiResponseData"))
}

func (p *OneOfProtectedPairsListApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if "List<files.v4.config.ProtectedFileServerPair>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfProtectedPairsListApiResponseData")
}

type OneOfDeleteQuotaPolicyApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *import4.TaskReference `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfDeleteQuotaPolicyApiResponseData() *OneOfDeleteQuotaPolicyApiResponseData {
  p := new(OneOfDeleteQuotaPolicyApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfDeleteQuotaPolicyApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfDeleteQuotaPolicyApiResponseData is nil"))
  }
  switch v.(type) {
    case import4.TaskReference:
      if nil == p.oneOfType0 {p.oneOfType0 = new(import4.TaskReference)}
      *p.oneOfType0 = v.(import4.TaskReference)
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

func (p *OneOfDeleteQuotaPolicyApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfDeleteQuotaPolicyApiResponseData) UnmarshalJSON(b []byte) error {
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
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteQuotaPolicyApiResponseData"))
}

func (p *OneOfDeleteQuotaPolicyApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfDeleteQuotaPolicyApiResponseData")
}

type OneOfReplicationJobApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
  oneOfType0 *ReplicationJob `json:"-"`
}

func NewOneOfReplicationJobApiResponseData() *OneOfReplicationJobApiResponseData {
  p := new(OneOfReplicationJobApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfReplicationJobApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfReplicationJobApiResponseData is nil"))
  }
  switch v.(type) {
    case import2.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = v.(import2.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case ReplicationJob:
      if nil == p.oneOfType0 {p.oneOfType0 = new(ReplicationJob)}
      *p.oneOfType0 = v.(ReplicationJob)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfReplicationJobApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  return nil
}

func (p *OneOfReplicationJobApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new(ReplicationJob)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "files.v4.config.ReplicationJob" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(ReplicationJob)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfReplicationJobApiResponseData"))
}

func (p *OneOfReplicationJobApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfReplicationJobApiResponseData")
}

type OneOfSynchronizationPathsValidationApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
  oneOfType0 *SynchronizationPathsValidationResponse `json:"-"`
}

func NewOneOfSynchronizationPathsValidationApiResponseData() *OneOfSynchronizationPathsValidationApiResponseData {
  p := new(OneOfSynchronizationPathsValidationApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfSynchronizationPathsValidationApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfSynchronizationPathsValidationApiResponseData is nil"))
  }
  switch v.(type) {
    case import2.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = v.(import2.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case SynchronizationPathsValidationResponse:
      if nil == p.oneOfType0 {p.oneOfType0 = new(SynchronizationPathsValidationResponse)}
      *p.oneOfType0 = v.(SynchronizationPathsValidationResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfSynchronizationPathsValidationApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  return nil
}

func (p *OneOfSynchronizationPathsValidationApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new(SynchronizationPathsValidationResponse)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "files.v4.config.SynchronizationPathsValidationResponse" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(SynchronizationPathsValidationResponse)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfSynchronizationPathsValidationApiResponseData"))
}

func (p *OneOfSynchronizationPathsValidationApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfSynchronizationPathsValidationApiResponseData")
}

type OneOfFileServerUserApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *FileServerUser `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfFileServerUserApiResponseData() *OneOfFileServerUserApiResponseData {
  p := new(OneOfFileServerUserApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfFileServerUserApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfFileServerUserApiResponseData is nil"))
  }
  switch v.(type) {
    case FileServerUser:
      if nil == p.oneOfType0 {p.oneOfType0 = new(FileServerUser)}
      *p.oneOfType0 = v.(FileServerUser)
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

func (p *OneOfFileServerUserApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfFileServerUserApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new(FileServerUser)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "files.v4.config.FileServerUser" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(FileServerUser)}
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
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfFileServerUserApiResponseData"))
}

func (p *OneOfFileServerUserApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfFileServerUserApiResponseData")
}

type OneOfResolveAlertsApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 []ResolveAlertsResponse `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfResolveAlertsApiResponseData() *OneOfResolveAlertsApiResponseData {
  p := new(OneOfResolveAlertsApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfResolveAlertsApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfResolveAlertsApiResponseData is nil"))
  }
  switch v.(type) {
    case []ResolveAlertsResponse:
      p.oneOfType0 = v.([]ResolveAlertsResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<files.v4.config.ResolveAlertsResponse>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<files.v4.config.ResolveAlertsResponse>"
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

func (p *OneOfResolveAlertsApiResponseData) GetValue() interface{} {
  if "List<files.v4.config.ResolveAlertsResponse>" == *p.Discriminator {
    return p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfResolveAlertsApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new([]ResolveAlertsResponse)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "files.v4.config.ResolveAlertsResponse" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<files.v4.config.ResolveAlertsResponse>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<files.v4.config.ResolveAlertsResponse>"
      return nil

    }
  }
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfResolveAlertsApiResponseData"))
}

func (p *OneOfResolveAlertsApiResponseData) MarshalJSON() ([]byte, error) {
  if "List<files.v4.config.ResolveAlertsResponse>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfResolveAlertsApiResponseData")
}

type OneOfFileServerUserListApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
  oneOfType0 []FileServerUser `json:"-"`
}

func NewOneOfFileServerUserListApiResponseData() *OneOfFileServerUserListApiResponseData {
  p := new(OneOfFileServerUserListApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfFileServerUserListApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfFileServerUserListApiResponseData is nil"))
  }
  switch v.(type) {
    case import2.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = v.(import2.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case []FileServerUser:
      p.oneOfType0 = v.([]FileServerUser)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<files.v4.config.FileServerUser>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<files.v4.config.FileServerUser>"
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfFileServerUserListApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if "List<files.v4.config.FileServerUser>" == *p.Discriminator {
    return p.oneOfType0
  }
  return nil
}

func (p *OneOfFileServerUserListApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new([]FileServerUser)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "files.v4.config.FileServerUser" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<files.v4.config.FileServerUser>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<files.v4.config.FileServerUser>"
      return nil

    }
  }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfFileServerUserListApiResponseData"))
}

func (p *OneOfFileServerUserListApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if "List<files.v4.config.FileServerUser>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfFileServerUserListApiResponseData")
}

type OneOfInfectedFilesListApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 []InfectedFile `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfInfectedFilesListApiResponseData() *OneOfInfectedFilesListApiResponseData {
  p := new(OneOfInfectedFilesListApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfInfectedFilesListApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfInfectedFilesListApiResponseData is nil"))
  }
  switch v.(type) {
    case []InfectedFile:
      p.oneOfType0 = v.([]InfectedFile)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<files.v4.config.InfectedFile>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<files.v4.config.InfectedFile>"
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

func (p *OneOfInfectedFilesListApiResponseData) GetValue() interface{} {
  if "List<files.v4.config.InfectedFile>" == *p.Discriminator {
    return p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfInfectedFilesListApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new([]InfectedFile)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "files.v4.config.InfectedFile" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<files.v4.config.InfectedFile>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<files.v4.config.InfectedFile>"
      return nil

    }
  }
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfInfectedFilesListApiResponseData"))
}

func (p *OneOfInfectedFilesListApiResponseData) MarshalJSON() ([]byte, error) {
  if "List<files.v4.config.InfectedFile>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfInfectedFilesListApiResponseData")
}

type OneOfActivateFileServerApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *import4.TaskReference `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfActivateFileServerApiResponseData() *OneOfActivateFileServerApiResponseData {
  p := new(OneOfActivateFileServerApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfActivateFileServerApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfActivateFileServerApiResponseData is nil"))
  }
  switch v.(type) {
    case import4.TaskReference:
      if nil == p.oneOfType0 {p.oneOfType0 = new(import4.TaskReference)}
      *p.oneOfType0 = v.(import4.TaskReference)
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

func (p *OneOfActivateFileServerApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfActivateFileServerApiResponseData) UnmarshalJSON(b []byte) error {
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
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfActivateFileServerApiResponseData"))
}

func (p *OneOfActivateFileServerApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfActivateFileServerApiResponseData")
}

type OneOfNetworkGetByExtIdApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *Network `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfNetworkGetByExtIdApiResponseData() *OneOfNetworkGetByExtIdApiResponseData {
  p := new(OneOfNetworkGetByExtIdApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfNetworkGetByExtIdApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfNetworkGetByExtIdApiResponseData is nil"))
  }
  switch v.(type) {
    case Network:
      if nil == p.oneOfType0 {p.oneOfType0 = new(Network)}
      *p.oneOfType0 = v.(Network)
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

func (p *OneOfNetworkGetByExtIdApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfNetworkGetByExtIdApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new(Network)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "files.v4.config.Network" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(Network)}
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
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfNetworkGetByExtIdApiResponseData"))
}

func (p *OneOfNetworkGetByExtIdApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfNetworkGetByExtIdApiResponseData")
}

type OneOfUpdateMountTargetVirusScanPolicyApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *VirusScanPolicyMountTarget `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfUpdateMountTargetVirusScanPolicyApiResponseData() *OneOfUpdateMountTargetVirusScanPolicyApiResponseData {
  p := new(OneOfUpdateMountTargetVirusScanPolicyApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfUpdateMountTargetVirusScanPolicyApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfUpdateMountTargetVirusScanPolicyApiResponseData is nil"))
  }
  switch v.(type) {
    case VirusScanPolicyMountTarget:
      if nil == p.oneOfType0 {p.oneOfType0 = new(VirusScanPolicyMountTarget)}
      *p.oneOfType0 = v.(VirusScanPolicyMountTarget)
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

func (p *OneOfUpdateMountTargetVirusScanPolicyApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfUpdateMountTargetVirusScanPolicyApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new(VirusScanPolicyMountTarget)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "files.v4.config.VirusScanPolicyMountTarget" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(VirusScanPolicyMountTarget)}
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
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateMountTargetVirusScanPolicyApiResponseData"))
}

func (p *OneOfUpdateMountTargetVirusScanPolicyApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfUpdateMountTargetVirusScanPolicyApiResponseData")
}

type OneOfVerifyDnsEntriesApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *import4.TaskReference `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfVerifyDnsEntriesApiResponseData() *OneOfVerifyDnsEntriesApiResponseData {
  p := new(OneOfVerifyDnsEntriesApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfVerifyDnsEntriesApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfVerifyDnsEntriesApiResponseData is nil"))
  }
  switch v.(type) {
    case import4.TaskReference:
      if nil == p.oneOfType0 {p.oneOfType0 = new(import4.TaskReference)}
      *p.oneOfType0 = v.(import4.TaskReference)
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

func (p *OneOfVerifyDnsEntriesApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfVerifyDnsEntriesApiResponseData) UnmarshalJSON(b []byte) error {
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
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfVerifyDnsEntriesApiResponseData"))
}

func (p *OneOfVerifyDnsEntriesApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfVerifyDnsEntriesApiResponseData")
}

type OneOfDisableDataConsolidationPolicyApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *import4.TaskReference `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfDisableDataConsolidationPolicyApiResponseData() *OneOfDisableDataConsolidationPolicyApiResponseData {
  p := new(OneOfDisableDataConsolidationPolicyApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfDisableDataConsolidationPolicyApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfDisableDataConsolidationPolicyApiResponseData is nil"))
  }
  switch v.(type) {
    case import4.TaskReference:
      if nil == p.oneOfType0 {p.oneOfType0 = new(import4.TaskReference)}
      *p.oneOfType0 = v.(import4.TaskReference)
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

func (p *OneOfDisableDataConsolidationPolicyApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfDisableDataConsolidationPolicyApiResponseData) UnmarshalJSON(b []byte) error {
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
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDisableDataConsolidationPolicyApiResponseData"))
}

func (p *OneOfDisableDataConsolidationPolicyApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfDisableDataConsolidationPolicyApiResponseData")
}

type OneOfUpdateFileServerApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *import4.TaskReference `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfUpdateFileServerApiResponseData() *OneOfUpdateFileServerApiResponseData {
  p := new(OneOfUpdateFileServerApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfUpdateFileServerApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfUpdateFileServerApiResponseData is nil"))
  }
  switch v.(type) {
    case import4.TaskReference:
      if nil == p.oneOfType0 {p.oneOfType0 = new(import4.TaskReference)}
      *p.oneOfType0 = v.(import4.TaskReference)
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

func (p *OneOfUpdateFileServerApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfUpdateFileServerApiResponseData) UnmarshalJSON(b []byte) error {
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
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateFileServerApiResponseData"))
}

func (p *OneOfUpdateFileServerApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfUpdateFileServerApiResponseData")
}

type OneOfDeleteDataConsolidationPolicyApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *import4.TaskReference `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfDeleteDataConsolidationPolicyApiResponseData() *OneOfDeleteDataConsolidationPolicyApiResponseData {
  p := new(OneOfDeleteDataConsolidationPolicyApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfDeleteDataConsolidationPolicyApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfDeleteDataConsolidationPolicyApiResponseData is nil"))
  }
  switch v.(type) {
    case import4.TaskReference:
      if nil == p.oneOfType0 {p.oneOfType0 = new(import4.TaskReference)}
      *p.oneOfType0 = v.(import4.TaskReference)
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

func (p *OneOfDeleteDataConsolidationPolicyApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfDeleteDataConsolidationPolicyApiResponseData) UnmarshalJSON(b []byte) error {
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
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteDataConsolidationPolicyApiResponseData"))
}

func (p *OneOfDeleteDataConsolidationPolicyApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfDeleteDataConsolidationPolicyApiResponseData")
}

type OneOfUpdateAntivirusServerApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *AntivirusServer `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfUpdateAntivirusServerApiResponseData() *OneOfUpdateAntivirusServerApiResponseData {
  p := new(OneOfUpdateAntivirusServerApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfUpdateAntivirusServerApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfUpdateAntivirusServerApiResponseData is nil"))
  }
  switch v.(type) {
    case AntivirusServer:
      if nil == p.oneOfType0 {p.oneOfType0 = new(AntivirusServer)}
      *p.oneOfType0 = v.(AntivirusServer)
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

func (p *OneOfUpdateAntivirusServerApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfUpdateAntivirusServerApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new(AntivirusServer)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "files.v4.config.AntivirusServer" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(AntivirusServer)}
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
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateAntivirusServerApiResponseData"))
}

func (p *OneOfUpdateAntivirusServerApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfUpdateAntivirusServerApiResponseData")
}

type OneOfConfigObjectStoreApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *import4.TaskReference `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfConfigObjectStoreApiResponseData() *OneOfConfigObjectStoreApiResponseData {
  p := new(OneOfConfigObjectStoreApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfConfigObjectStoreApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfConfigObjectStoreApiResponseData is nil"))
  }
  switch v.(type) {
    case import4.TaskReference:
      if nil == p.oneOfType0 {p.oneOfType0 = new(import4.TaskReference)}
      *p.oneOfType0 = v.(import4.TaskReference)
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

func (p *OneOfConfigObjectStoreApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfConfigObjectStoreApiResponseData) UnmarshalJSON(b []byte) error {
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
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfConfigObjectStoreApiResponseData"))
}

func (p *OneOfConfigObjectStoreApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfConfigObjectStoreApiResponseData")
}

type OneOfDeleteTieringConfigApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType1 *interface{} `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfDeleteTieringConfigApiResponseData() *OneOfDeleteTieringConfigApiResponseData {
  p := new(OneOfDeleteTieringConfigApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfDeleteTieringConfigApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfDeleteTieringConfigApiResponseData is nil"))
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

func (p *OneOfDeleteTieringConfigApiResponseData) GetValue() interface{} {
  if "EMPTY" == *p.Discriminator {
    return *p.oneOfType1
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfDeleteTieringConfigApiResponseData) UnmarshalJSON(b []byte) error {
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
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteTieringConfigApiResponseData"))
}

func (p *OneOfDeleteTieringConfigApiResponseData) MarshalJSON() ([]byte, error) {
  if "EMPTY" == *p.Discriminator {
    return json.Marshal(p.oneOfType1)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfDeleteTieringConfigApiResponseData")
}

type OneOfDeleteMountTargetApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *import4.TaskReference `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfDeleteMountTargetApiResponseData() *OneOfDeleteMountTargetApiResponseData {
  p := new(OneOfDeleteMountTargetApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfDeleteMountTargetApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfDeleteMountTargetApiResponseData is nil"))
  }
  switch v.(type) {
    case import4.TaskReference:
      if nil == p.oneOfType0 {p.oneOfType0 = new(import4.TaskReference)}
      *p.oneOfType0 = v.(import4.TaskReference)
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

func (p *OneOfDeleteMountTargetApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfDeleteMountTargetApiResponseData) UnmarshalJSON(b []byte) error {
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
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteMountTargetApiResponseData"))
}

func (p *OneOfDeleteMountTargetApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfDeleteMountTargetApiResponseData")
}

type OneOfAntivirusServerListApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
  oneOfType0 []AntivirusServer `json:"-"`
}

func NewOneOfAntivirusServerListApiResponseData() *OneOfAntivirusServerListApiResponseData {
  p := new(OneOfAntivirusServerListApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfAntivirusServerListApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfAntivirusServerListApiResponseData is nil"))
  }
  switch v.(type) {
    case import2.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = v.(import2.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case []AntivirusServer:
      p.oneOfType0 = v.([]AntivirusServer)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<files.v4.config.AntivirusServer>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<files.v4.config.AntivirusServer>"
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfAntivirusServerListApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if "List<files.v4.config.AntivirusServer>" == *p.Discriminator {
    return p.oneOfType0
  }
  return nil
}

func (p *OneOfAntivirusServerListApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new([]AntivirusServer)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "files.v4.config.AntivirusServer" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<files.v4.config.AntivirusServer>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<files.v4.config.AntivirusServer>"
      return nil

    }
  }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfAntivirusServerListApiResponseData"))
}

func (p *OneOfAntivirusServerListApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if "List<files.v4.config.AntivirusServer>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfAntivirusServerListApiResponseData")
}

type OneOfUnQuarantineInfectedFilesApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *import4.TaskReference `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfUnQuarantineInfectedFilesApiResponseData() *OneOfUnQuarantineInfectedFilesApiResponseData {
  p := new(OneOfUnQuarantineInfectedFilesApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfUnQuarantineInfectedFilesApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfUnQuarantineInfectedFilesApiResponseData is nil"))
  }
  switch v.(type) {
    case import4.TaskReference:
      if nil == p.oneOfType0 {p.oneOfType0 = new(import4.TaskReference)}
      *p.oneOfType0 = v.(import4.TaskReference)
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

func (p *OneOfUnQuarantineInfectedFilesApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfUnQuarantineInfectedFilesApiResponseData) UnmarshalJSON(b []byte) error {
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
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUnQuarantineInfectedFilesApiResponseData"))
}

func (p *OneOfUnQuarantineInfectedFilesApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfUnQuarantineInfectedFilesApiResponseData")
}

type OneOfAntivirusServerApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *AntivirusServer `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfAntivirusServerApiResponseData() *OneOfAntivirusServerApiResponseData {
  p := new(OneOfAntivirusServerApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfAntivirusServerApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfAntivirusServerApiResponseData is nil"))
  }
  switch v.(type) {
    case AntivirusServer:
      if nil == p.oneOfType0 {p.oneOfType0 = new(AntivirusServer)}
      *p.oneOfType0 = v.(AntivirusServer)
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

func (p *OneOfAntivirusServerApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfAntivirusServerApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new(AntivirusServer)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "files.v4.config.AntivirusServer" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(AntivirusServer)}
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
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfAntivirusServerApiResponseData"))
}

func (p *OneOfAntivirusServerApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfAntivirusServerApiResponseData")
}

type OneOfRansomwareConfigApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *RansomwareConfig `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfRansomwareConfigApiResponseData() *OneOfRansomwareConfigApiResponseData {
  p := new(OneOfRansomwareConfigApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfRansomwareConfigApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfRansomwareConfigApiResponseData is nil"))
  }
  switch v.(type) {
    case RansomwareConfig:
      if nil == p.oneOfType0 {p.oneOfType0 = new(RansomwareConfig)}
      *p.oneOfType0 = v.(RansomwareConfig)
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

func (p *OneOfRansomwareConfigApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfRansomwareConfigApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new(RansomwareConfig)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "files.v4.config.RansomwareConfig" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(RansomwareConfig)}
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
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfRansomwareConfigApiResponseData"))
}

func (p *OneOfRansomwareConfigApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfRansomwareConfigApiResponseData")
}

type OneOfGlobalBlockedClientApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
  oneOfType0 *GlobalBlockedClient `json:"-"`
}

func NewOneOfGlobalBlockedClientApiResponseData() *OneOfGlobalBlockedClientApiResponseData {
  p := new(OneOfGlobalBlockedClientApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfGlobalBlockedClientApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfGlobalBlockedClientApiResponseData is nil"))
  }
  switch v.(type) {
    case import2.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = v.(import2.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case GlobalBlockedClient:
      if nil == p.oneOfType0 {p.oneOfType0 = new(GlobalBlockedClient)}
      *p.oneOfType0 = v.(GlobalBlockedClient)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfGlobalBlockedClientApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  return nil
}

func (p *OneOfGlobalBlockedClientApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new(GlobalBlockedClient)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "files.v4.config.GlobalBlockedClient" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(GlobalBlockedClient)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGlobalBlockedClientApiResponseData"))
}

func (p *OneOfGlobalBlockedClientApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfGlobalBlockedClientApiResponseData")
}

type OneOfAddDnsEntriesApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *import4.TaskReference `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfAddDnsEntriesApiResponseData() *OneOfAddDnsEntriesApiResponseData {
  p := new(OneOfAddDnsEntriesApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfAddDnsEntriesApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfAddDnsEntriesApiResponseData is nil"))
  }
  switch v.(type) {
    case import4.TaskReference:
      if nil == p.oneOfType0 {p.oneOfType0 = new(import4.TaskReference)}
      *p.oneOfType0 = v.(import4.TaskReference)
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

func (p *OneOfAddDnsEntriesApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfAddDnsEntriesApiResponseData) UnmarshalJSON(b []byte) error {
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
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfAddDnsEntriesApiResponseData"))
}

func (p *OneOfAddDnsEntriesApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfAddDnsEntriesApiResponseData")
}

type OneOfRescanInfectedFilesApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *import4.TaskReference `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfRescanInfectedFilesApiResponseData() *OneOfRescanInfectedFilesApiResponseData {
  p := new(OneOfRescanInfectedFilesApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfRescanInfectedFilesApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfRescanInfectedFilesApiResponseData is nil"))
  }
  switch v.(type) {
    case import4.TaskReference:
      if nil == p.oneOfType0 {p.oneOfType0 = new(import4.TaskReference)}
      *p.oneOfType0 = v.(import4.TaskReference)
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

func (p *OneOfRescanInfectedFilesApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfRescanInfectedFilesApiResponseData) UnmarshalJSON(b []byte) error {
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
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfRescanInfectedFilesApiResponseData"))
}

func (p *OneOfRescanInfectedFilesApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfRescanInfectedFilesApiResponseData")
}

type OneOfTierFilesApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *TierResponseSpec `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfTierFilesApiResponseData() *OneOfTierFilesApiResponseData {
  p := new(OneOfTierFilesApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfTierFilesApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfTierFilesApiResponseData is nil"))
  }
  switch v.(type) {
    case TierResponseSpec:
      if nil == p.oneOfType0 {p.oneOfType0 = new(TierResponseSpec)}
      *p.oneOfType0 = v.(TierResponseSpec)
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

func (p *OneOfTierFilesApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfTierFilesApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new(TierResponseSpec)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "files.v4.config.TierResponseSpec" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(TierResponseSpec)}
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
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfTierFilesApiResponseData"))
}

func (p *OneOfTierFilesApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfTierFilesApiResponseData")
}

type OneOfDeleteAntivirusServerApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType1 *interface{} `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfDeleteAntivirusServerApiResponseData() *OneOfDeleteAntivirusServerApiResponseData {
  p := new(OneOfDeleteAntivirusServerApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfDeleteAntivirusServerApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfDeleteAntivirusServerApiResponseData is nil"))
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

func (p *OneOfDeleteAntivirusServerApiResponseData) GetValue() interface{} {
  if "EMPTY" == *p.Discriminator {
    return *p.oneOfType1
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfDeleteAntivirusServerApiResponseData) UnmarshalJSON(b []byte) error {
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
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteAntivirusServerApiResponseData"))
}

func (p *OneOfDeleteAntivirusServerApiResponseData) MarshalJSON() ([]byte, error) {
  if "EMPTY" == *p.Discriminator {
    return json.Marshal(p.oneOfType1)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfDeleteAntivirusServerApiResponseData")
}

type OneOfSearchUserMappingPOSTResponsesData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
  oneOfType0 *SearchUserMappingResponseSpec `json:"-"`
}

func NewOneOfSearchUserMappingPOSTResponsesData() *OneOfSearchUserMappingPOSTResponsesData {
  p := new(OneOfSearchUserMappingPOSTResponsesData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfSearchUserMappingPOSTResponsesData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfSearchUserMappingPOSTResponsesData is nil"))
  }
  switch v.(type) {
    case import2.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = v.(import2.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case SearchUserMappingResponseSpec:
      if nil == p.oneOfType0 {p.oneOfType0 = new(SearchUserMappingResponseSpec)}
      *p.oneOfType0 = v.(SearchUserMappingResponseSpec)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfSearchUserMappingPOSTResponsesData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  return nil
}

func (p *OneOfSearchUserMappingPOSTResponsesData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new(SearchUserMappingResponseSpec)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "files.v4.config.SearchUserMappingResponseSpec" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(SearchUserMappingResponseSpec)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfSearchUserMappingPOSTResponsesData"))
}

func (p *OneOfSearchUserMappingPOSTResponsesData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfSearchUserMappingPOSTResponsesData")
}

type OneOfUpdateMountTargetApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *import4.TaskReference `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfUpdateMountTargetApiResponseData() *OneOfUpdateMountTargetApiResponseData {
  p := new(OneOfUpdateMountTargetApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfUpdateMountTargetApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfUpdateMountTargetApiResponseData is nil"))
  }
  switch v.(type) {
    case import4.TaskReference:
      if nil == p.oneOfType0 {p.oneOfType0 = new(import4.TaskReference)}
      *p.oneOfType0 = v.(import4.TaskReference)
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

func (p *OneOfUpdateMountTargetApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfUpdateMountTargetApiResponseData) UnmarshalJSON(b []byte) error {
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
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateMountTargetApiResponseData"))
}

func (p *OneOfUpdateMountTargetApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfUpdateMountTargetApiResponseData")
}

type OneOfNameServicesApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *NameServiceSpec `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfNameServicesApiResponseData() *OneOfNameServicesApiResponseData {
  p := new(OneOfNameServicesApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfNameServicesApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfNameServicesApiResponseData is nil"))
  }
  switch v.(type) {
    case NameServiceSpec:
      if nil == p.oneOfType0 {p.oneOfType0 = new(NameServiceSpec)}
      *p.oneOfType0 = v.(NameServiceSpec)
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

func (p *OneOfNameServicesApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfNameServicesApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new(NameServiceSpec)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "files.v4.config.NameServiceSpec" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(NameServiceSpec)}
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
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfNameServicesApiResponseData"))
}

func (p *OneOfNameServicesApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfNameServicesApiResponseData")
}

type OneOfTierRefreshColdDataApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType1 *interface{} `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfTierRefreshColdDataApiResponseData() *OneOfTierRefreshColdDataApiResponseData {
  p := new(OneOfTierRefreshColdDataApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfTierRefreshColdDataApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfTierRefreshColdDataApiResponseData is nil"))
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

func (p *OneOfTierRefreshColdDataApiResponseData) GetValue() interface{} {
  if "EMPTY" == *p.Discriminator {
    return *p.oneOfType1
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfTierRefreshColdDataApiResponseData) UnmarshalJSON(b []byte) error {
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
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfTierRefreshColdDataApiResponseData"))
}

func (p *OneOfTierRefreshColdDataApiResponseData) MarshalJSON() ([]byte, error) {
  if "EMPTY" == *p.Discriminator {
    return json.Marshal(p.oneOfType1)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfTierRefreshColdDataApiResponseData")
}

type OneOfDataConsolidationPolicyListApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 []DataConsolidationPolicy `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfDataConsolidationPolicyListApiResponseData() *OneOfDataConsolidationPolicyListApiResponseData {
  p := new(OneOfDataConsolidationPolicyListApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfDataConsolidationPolicyListApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfDataConsolidationPolicyListApiResponseData is nil"))
  }
  switch v.(type) {
    case []DataConsolidationPolicy:
      p.oneOfType0 = v.([]DataConsolidationPolicy)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<files.v4.config.DataConsolidationPolicy>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<files.v4.config.DataConsolidationPolicy>"
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

func (p *OneOfDataConsolidationPolicyListApiResponseData) GetValue() interface{} {
  if "List<files.v4.config.DataConsolidationPolicy>" == *p.Discriminator {
    return p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfDataConsolidationPolicyListApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new([]DataConsolidationPolicy)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "files.v4.config.DataConsolidationPolicy" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<files.v4.config.DataConsolidationPolicy>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<files.v4.config.DataConsolidationPolicy>"
      return nil

    }
  }
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDataConsolidationPolicyListApiResponseData"))
}

func (p *OneOfDataConsolidationPolicyListApiResponseData) MarshalJSON() ([]byte, error) {
  if "List<files.v4.config.DataConsolidationPolicy>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfDataConsolidationPolicyListApiResponseData")
}

type OneOfNetworkListApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 []Network `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfNetworkListApiResponseData() *OneOfNetworkListApiResponseData {
  p := new(OneOfNetworkListApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfNetworkListApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfNetworkListApiResponseData is nil"))
  }
  switch v.(type) {
    case []Network:
      p.oneOfType0 = v.([]Network)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<files.v4.config.Network>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<files.v4.config.Network>"
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

func (p *OneOfNetworkListApiResponseData) GetValue() interface{} {
  if "List<files.v4.config.Network>" == *p.Discriminator {
    return p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfNetworkListApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new([]Network)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "files.v4.config.Network" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<files.v4.config.Network>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<files.v4.config.Network>"
      return nil

    }
  }
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfNetworkListApiResponseData"))
}

func (p *OneOfNetworkListApiResponseData) MarshalJSON() ([]byte, error) {
  if "List<files.v4.config.Network>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfNetworkListApiResponseData")
}

type OneOfFileServerGetApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *FileServer `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfFileServerGetApiResponseData() *OneOfFileServerGetApiResponseData {
  p := new(OneOfFileServerGetApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfFileServerGetApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfFileServerGetApiResponseData is nil"))
  }
  switch v.(type) {
    case FileServer:
      if nil == p.oneOfType0 {p.oneOfType0 = new(FileServer)}
      *p.oneOfType0 = v.(FileServer)
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

func (p *OneOfFileServerGetApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfFileServerGetApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new(FileServer)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "files.v4.config.FileServer" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(FileServer)}
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
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfFileServerGetApiResponseData"))
}

func (p *OneOfFileServerGetApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfFileServerGetApiResponseData")
}

type OneOfCreateAntivirusServerApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *AntivirusServer `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfCreateAntivirusServerApiResponseData() *OneOfCreateAntivirusServerApiResponseData {
  p := new(OneOfCreateAntivirusServerApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfCreateAntivirusServerApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfCreateAntivirusServerApiResponseData is nil"))
  }
  switch v.(type) {
    case AntivirusServer:
      if nil == p.oneOfType0 {p.oneOfType0 = new(AntivirusServer)}
      *p.oneOfType0 = v.(AntivirusServer)
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

func (p *OneOfCreateAntivirusServerApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfCreateAntivirusServerApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new(AntivirusServer)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "files.v4.config.AntivirusServer" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(AntivirusServer)}
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
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateAntivirusServerApiResponseData"))
}

func (p *OneOfCreateAntivirusServerApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfCreateAntivirusServerApiResponseData")
}

type OneOfDeleteFileServerUserApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType1 *interface{} `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfDeleteFileServerUserApiResponseData() *OneOfDeleteFileServerUserApiResponseData {
  p := new(OneOfDeleteFileServerUserApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfDeleteFileServerUserApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfDeleteFileServerUserApiResponseData is nil"))
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

func (p *OneOfDeleteFileServerUserApiResponseData) GetValue() interface{} {
  if "EMPTY" == *p.Discriminator {
    return *p.oneOfType1
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfDeleteFileServerUserApiResponseData) UnmarshalJSON(b []byte) error {
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
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteFileServerUserApiResponseData"))
}

func (p *OneOfDeleteFileServerUserApiResponseData) MarshalJSON() ([]byte, error) {
  if "EMPTY" == *p.Discriminator {
    return json.Marshal(p.oneOfType1)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfDeleteFileServerUserApiResponseData")
}

type OneOfMountTargetSnapshotListApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
  oneOfType0 []Snapshot `json:"-"`
}

func NewOneOfMountTargetSnapshotListApiResponseData() *OneOfMountTargetSnapshotListApiResponseData {
  p := new(OneOfMountTargetSnapshotListApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfMountTargetSnapshotListApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfMountTargetSnapshotListApiResponseData is nil"))
  }
  switch v.(type) {
    case import2.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = v.(import2.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case []Snapshot:
      p.oneOfType0 = v.([]Snapshot)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<files.v4.config.Snapshot>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<files.v4.config.Snapshot>"
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfMountTargetSnapshotListApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if "List<files.v4.config.Snapshot>" == *p.Discriminator {
    return p.oneOfType0
  }
  return nil
}

func (p *OneOfMountTargetSnapshotListApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new([]Snapshot)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "files.v4.config.Snapshot" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<files.v4.config.Snapshot>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<files.v4.config.Snapshot>"
      return nil

    }
  }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfMountTargetSnapshotListApiResponseData"))
}

func (p *OneOfMountTargetSnapshotListApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if "List<files.v4.config.Snapshot>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfMountTargetSnapshotListApiResponseData")
}

type OneOfTestConnectionAntivirusServerApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
  oneOfType0 *Response `json:"-"`
}

func NewOneOfTestConnectionAntivirusServerApiResponseData() *OneOfTestConnectionAntivirusServerApiResponseData {
  p := new(OneOfTestConnectionAntivirusServerApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfTestConnectionAntivirusServerApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfTestConnectionAntivirusServerApiResponseData is nil"))
  }
  switch v.(type) {
    case import2.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = v.(import2.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case Response:
      if nil == p.oneOfType0 {p.oneOfType0 = new(Response)}
      *p.oneOfType0 = v.(Response)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfTestConnectionAntivirusServerApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  return nil
}

func (p *OneOfTestConnectionAntivirusServerApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new(Response)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "files.v4.config.Response" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(Response)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfTestConnectionAntivirusServerApiResponseData"))
}

func (p *OneOfTestConnectionAntivirusServerApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfTestConnectionAntivirusServerApiResponseData")
}

type OneOfDisableDataProtectionPolicyApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *import4.TaskReference `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfDisableDataProtectionPolicyApiResponseData() *OneOfDisableDataProtectionPolicyApiResponseData {
  p := new(OneOfDisableDataProtectionPolicyApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfDisableDataProtectionPolicyApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfDisableDataProtectionPolicyApiResponseData is nil"))
  }
  switch v.(type) {
    case import4.TaskReference:
      if nil == p.oneOfType0 {p.oneOfType0 = new(import4.TaskReference)}
      *p.oneOfType0 = v.(import4.TaskReference)
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

func (p *OneOfDisableDataProtectionPolicyApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfDisableDataProtectionPolicyApiResponseData) UnmarshalJSON(b []byte) error {
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
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDisableDataProtectionPolicyApiResponseData"))
}

func (p *OneOfDisableDataProtectionPolicyApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfDisableDataProtectionPolicyApiResponseData")
}

type OneOfGetTieringPolicyApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
  oneOfType0 *TieringPolicy `json:"-"`
}

func NewOneOfGetTieringPolicyApiResponseData() *OneOfGetTieringPolicyApiResponseData {
  p := new(OneOfGetTieringPolicyApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfGetTieringPolicyApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfGetTieringPolicyApiResponseData is nil"))
  }
  switch v.(type) {
    case import2.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = v.(import2.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case TieringPolicy:
      if nil == p.oneOfType0 {p.oneOfType0 = new(TieringPolicy)}
      *p.oneOfType0 = v.(TieringPolicy)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfGetTieringPolicyApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  return nil
}

func (p *OneOfGetTieringPolicyApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new(TieringPolicy)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "files.v4.config.TieringPolicy" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(TieringPolicy)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetTieringPolicyApiResponseData"))
}

func (p *OneOfGetTieringPolicyApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfGetTieringPolicyApiResponseData")
}

type OneOfEnableDataConsolidationPolicyApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *import4.TaskReference `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfEnableDataConsolidationPolicyApiResponseData() *OneOfEnableDataConsolidationPolicyApiResponseData {
  p := new(OneOfEnableDataConsolidationPolicyApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfEnableDataConsolidationPolicyApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfEnableDataConsolidationPolicyApiResponseData is nil"))
  }
  switch v.(type) {
    case import4.TaskReference:
      if nil == p.oneOfType0 {p.oneOfType0 = new(import4.TaskReference)}
      *p.oneOfType0 = v.(import4.TaskReference)
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

func (p *OneOfEnableDataConsolidationPolicyApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfEnableDataConsolidationPolicyApiResponseData) UnmarshalJSON(b []byte) error {
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
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfEnableDataConsolidationPolicyApiResponseData"))
}

func (p *OneOfEnableDataConsolidationPolicyApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfEnableDataConsolidationPolicyApiResponseData")
}

type OneOfDeleteAdminUserApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType1 *interface{} `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfDeleteAdminUserApiResponseData() *OneOfDeleteAdminUserApiResponseData {
  p := new(OneOfDeleteAdminUserApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfDeleteAdminUserApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfDeleteAdminUserApiResponseData is nil"))
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

func (p *OneOfDeleteAdminUserApiResponseData) GetValue() interface{} {
  if "EMPTY" == *p.Discriminator {
    return *p.oneOfType1
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfDeleteAdminUserApiResponseData) UnmarshalJSON(b []byte) error {
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
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteAdminUserApiResponseData"))
}

func (p *OneOfDeleteAdminUserApiResponseData) MarshalJSON() ([]byte, error) {
  if "EMPTY" == *p.Discriminator {
    return json.Marshal(p.oneOfType1)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfDeleteAdminUserApiResponseData")
}

type OneOfProtectedPairsApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *ProtectedFileServerPair `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfProtectedPairsApiResponseData() *OneOfProtectedPairsApiResponseData {
  p := new(OneOfProtectedPairsApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfProtectedPairsApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfProtectedPairsApiResponseData is nil"))
  }
  switch v.(type) {
    case ProtectedFileServerPair:
      if nil == p.oneOfType0 {p.oneOfType0 = new(ProtectedFileServerPair)}
      *p.oneOfType0 = v.(ProtectedFileServerPair)
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

func (p *OneOfProtectedPairsApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfProtectedPairsApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new(ProtectedFileServerPair)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "files.v4.config.ProtectedFileServerPair" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(ProtectedFileServerPair)}
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
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfProtectedPairsApiResponseData"))
}

func (p *OneOfProtectedPairsApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfProtectedPairsApiResponseData")
}

type OneOfSsrSnapshotScheduleListApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
  oneOfType0 []SsrSnapshotSchedule `json:"-"`
}

func NewOneOfSsrSnapshotScheduleListApiResponseData() *OneOfSsrSnapshotScheduleListApiResponseData {
  p := new(OneOfSsrSnapshotScheduleListApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfSsrSnapshotScheduleListApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfSsrSnapshotScheduleListApiResponseData is nil"))
  }
  switch v.(type) {
    case import2.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = v.(import2.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case []SsrSnapshotSchedule:
      p.oneOfType0 = v.([]SsrSnapshotSchedule)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<files.v4.config.SsrSnapshotSchedule>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<files.v4.config.SsrSnapshotSchedule>"
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfSsrSnapshotScheduleListApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if "List<files.v4.config.SsrSnapshotSchedule>" == *p.Discriminator {
    return p.oneOfType0
  }
  return nil
}

func (p *OneOfSsrSnapshotScheduleListApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new([]SsrSnapshotSchedule)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "files.v4.config.SsrSnapshotSchedule" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<files.v4.config.SsrSnapshotSchedule>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<files.v4.config.SsrSnapshotSchedule>"
      return nil

    }
  }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfSsrSnapshotScheduleListApiResponseData"))
}

func (p *OneOfSsrSnapshotScheduleListApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if "List<files.v4.config.SsrSnapshotSchedule>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfSsrSnapshotScheduleListApiResponseData")
}

type OneOfQuotaPolicyListApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
  oneOfType0 []QuotaPolicy `json:"-"`
}

func NewOneOfQuotaPolicyListApiResponseData() *OneOfQuotaPolicyListApiResponseData {
  p := new(OneOfQuotaPolicyListApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfQuotaPolicyListApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfQuotaPolicyListApiResponseData is nil"))
  }
  switch v.(type) {
    case import2.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = v.(import2.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case []QuotaPolicy:
      p.oneOfType0 = v.([]QuotaPolicy)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<files.v4.config.QuotaPolicy>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<files.v4.config.QuotaPolicy>"
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfQuotaPolicyListApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if "List<files.v4.config.QuotaPolicy>" == *p.Discriminator {
    return p.oneOfType0
  }
  return nil
}

func (p *OneOfQuotaPolicyListApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new([]QuotaPolicy)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "files.v4.config.QuotaPolicy" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<files.v4.config.QuotaPolicy>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<files.v4.config.QuotaPolicy>"
      return nil

    }
  }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfQuotaPolicyListApiResponseData"))
}

func (p *OneOfQuotaPolicyListApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if "List<files.v4.config.QuotaPolicy>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfQuotaPolicyListApiResponseData")
}

type OneOfResetInfectedFilesApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *import4.TaskReference `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfResetInfectedFilesApiResponseData() *OneOfResetInfectedFilesApiResponseData {
  p := new(OneOfResetInfectedFilesApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfResetInfectedFilesApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfResetInfectedFilesApiResponseData is nil"))
  }
  switch v.(type) {
    case import4.TaskReference:
      if nil == p.oneOfType0 {p.oneOfType0 = new(import4.TaskReference)}
      *p.oneOfType0 = v.(import4.TaskReference)
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

func (p *OneOfResetInfectedFilesApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfResetInfectedFilesApiResponseData) UnmarshalJSON(b []byte) error {
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
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfResetInfectedFilesApiResponseData"))
}

func (p *OneOfResetInfectedFilesApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfResetInfectedFilesApiResponseData")
}

type OneOfDeActivateFileServerApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *import4.TaskReference `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfDeActivateFileServerApiResponseData() *OneOfDeActivateFileServerApiResponseData {
  p := new(OneOfDeActivateFileServerApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfDeActivateFileServerApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfDeActivateFileServerApiResponseData is nil"))
  }
  switch v.(type) {
    case import4.TaskReference:
      if nil == p.oneOfType0 {p.oneOfType0 = new(import4.TaskReference)}
      *p.oneOfType0 = v.(import4.TaskReference)
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

func (p *OneOfDeActivateFileServerApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfDeActivateFileServerApiResponseData) UnmarshalJSON(b []byte) error {
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
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeActivateFileServerApiResponseData"))
}

func (p *OneOfDeActivateFileServerApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfDeActivateFileServerApiResponseData")
}

type OneOfCreateMountTargetApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *import4.TaskReference `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfCreateMountTargetApiResponseData() *OneOfCreateMountTargetApiResponseData {
  p := new(OneOfCreateMountTargetApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfCreateMountTargetApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfCreateMountTargetApiResponseData is nil"))
  }
  switch v.(type) {
    case import4.TaskReference:
      if nil == p.oneOfType0 {p.oneOfType0 = new(import4.TaskReference)}
      *p.oneOfType0 = v.(import4.TaskReference)
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

func (p *OneOfCreateMountTargetApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfCreateMountTargetApiResponseData) UnmarshalJSON(b []byte) error {
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
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateMountTargetApiResponseData"))
}

func (p *OneOfCreateMountTargetApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfCreateMountTargetApiResponseData")
}

type OneOfUpdateTieringConfigApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *TieringConfig `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfUpdateTieringConfigApiResponseData() *OneOfUpdateTieringConfigApiResponseData {
  p := new(OneOfUpdateTieringConfigApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfUpdateTieringConfigApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfUpdateTieringConfigApiResponseData is nil"))
  }
  switch v.(type) {
    case TieringConfig:
      if nil == p.oneOfType0 {p.oneOfType0 = new(TieringConfig)}
      *p.oneOfType0 = v.(TieringConfig)
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

func (p *OneOfUpdateTieringConfigApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfUpdateTieringConfigApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new(TieringConfig)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "files.v4.config.TieringConfig" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(TieringConfig)}
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
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateTieringConfigApiResponseData"))
}

func (p *OneOfUpdateTieringConfigApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfUpdateTieringConfigApiResponseData")
}

type OneOfMimeTypeResponseApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
  oneOfType0 *MimeTypeResponse `json:"-"`
}

func NewOneOfMimeTypeResponseApiResponseData() *OneOfMimeTypeResponseApiResponseData {
  p := new(OneOfMimeTypeResponseApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfMimeTypeResponseApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfMimeTypeResponseApiResponseData is nil"))
  }
  switch v.(type) {
    case import2.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = v.(import2.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case MimeTypeResponse:
      if nil == p.oneOfType0 {p.oneOfType0 = new(MimeTypeResponse)}
      *p.oneOfType0 = v.(MimeTypeResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfMimeTypeResponseApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  return nil
}

func (p *OneOfMimeTypeResponseApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new(MimeTypeResponse)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "files.v4.config.MimeTypeResponse" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(MimeTypeResponse)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfMimeTypeResponseApiResponseData"))
}

func (p *OneOfMimeTypeResponseApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfMimeTypeResponseApiResponseData")
}

type OneOfCreateAdminUserApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
  oneOfType0 *AdminUser `json:"-"`
}

func NewOneOfCreateAdminUserApiResponseData() *OneOfCreateAdminUserApiResponseData {
  p := new(OneOfCreateAdminUserApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfCreateAdminUserApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfCreateAdminUserApiResponseData is nil"))
  }
  switch v.(type) {
    case import2.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = v.(import2.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case AdminUser:
      if nil == p.oneOfType0 {p.oneOfType0 = new(AdminUser)}
      *p.oneOfType0 = v.(AdminUser)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfCreateAdminUserApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  return nil
}

func (p *OneOfCreateAdminUserApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new(AdminUser)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "files.v4.config.AdminUser" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(AdminUser)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateAdminUserApiResponseData"))
}

func (p *OneOfCreateAdminUserApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfCreateAdminUserApiResponseData")
}

type OneOfGetQuotaEmailConfigApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *QuotaEmailConfig `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfGetQuotaEmailConfigApiResponseData() *OneOfGetQuotaEmailConfigApiResponseData {
  p := new(OneOfGetQuotaEmailConfigApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfGetQuotaEmailConfigApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfGetQuotaEmailConfigApiResponseData is nil"))
  }
  switch v.(type) {
    case QuotaEmailConfig:
      if nil == p.oneOfType0 {p.oneOfType0 = new(QuotaEmailConfig)}
      *p.oneOfType0 = v.(QuotaEmailConfig)
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

func (p *OneOfGetQuotaEmailConfigApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfGetQuotaEmailConfigApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new(QuotaEmailConfig)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "files.v4.config.QuotaEmailConfig" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(QuotaEmailConfig)}
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
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetQuotaEmailConfigApiResponseData"))
}

func (p *OneOfGetQuotaEmailConfigApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfGetQuotaEmailConfigApiResponseData")
}

type OneOfUpdateAdminUserApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
  oneOfType0 *AdminUser `json:"-"`
}

func NewOneOfUpdateAdminUserApiResponseData() *OneOfUpdateAdminUserApiResponseData {
  p := new(OneOfUpdateAdminUserApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfUpdateAdminUserApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfUpdateAdminUserApiResponseData is nil"))
  }
  switch v.(type) {
    case import2.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = v.(import2.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case AdminUser:
      if nil == p.oneOfType0 {p.oneOfType0 = new(AdminUser)}
      *p.oneOfType0 = v.(AdminUser)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfUpdateAdminUserApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  return nil
}

func (p *OneOfUpdateAdminUserApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new(AdminUser)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "files.v4.config.AdminUser" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(AdminUser)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateAdminUserApiResponseData"))
}

func (p *OneOfUpdateAdminUserApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfUpdateAdminUserApiResponseData")
}

type OneOfGetTieringPolicyListApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 []TieringPolicy `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfGetTieringPolicyListApiResponseData() *OneOfGetTieringPolicyListApiResponseData {
  p := new(OneOfGetTieringPolicyListApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfGetTieringPolicyListApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfGetTieringPolicyListApiResponseData is nil"))
  }
  switch v.(type) {
    case []TieringPolicy:
      p.oneOfType0 = v.([]TieringPolicy)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<files.v4.config.TieringPolicy>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<files.v4.config.TieringPolicy>"
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

func (p *OneOfGetTieringPolicyListApiResponseData) GetValue() interface{} {
  if "List<files.v4.config.TieringPolicy>" == *p.Discriminator {
    return p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfGetTieringPolicyListApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new([]TieringPolicy)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "files.v4.config.TieringPolicy" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<files.v4.config.TieringPolicy>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<files.v4.config.TieringPolicy>"
      return nil

    }
  }
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetTieringPolicyListApiResponseData"))
}

func (p *OneOfGetTieringPolicyListApiResponseData) MarshalJSON() ([]byte, error) {
  if "List<files.v4.config.TieringPolicy>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfGetTieringPolicyListApiResponseData")
}

type OneOfFailoverApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *import4.TaskReference `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfFailoverApiResponseData() *OneOfFailoverApiResponseData {
  p := new(OneOfFailoverApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfFailoverApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfFailoverApiResponseData is nil"))
  }
  switch v.(type) {
    case import4.TaskReference:
      if nil == p.oneOfType0 {p.oneOfType0 = new(import4.TaskReference)}
      *p.oneOfType0 = v.(import4.TaskReference)
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

func (p *OneOfFailoverApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfFailoverApiResponseData) UnmarshalJSON(b []byte) error {
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
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfFailoverApiResponseData"))
}

func (p *OneOfFailoverApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfFailoverApiResponseData")
}

type OneOfReplicationJobListApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
  oneOfType0 []ReplicationJob `json:"-"`
}

func NewOneOfReplicationJobListApiResponseData() *OneOfReplicationJobListApiResponseData {
  p := new(OneOfReplicationJobListApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfReplicationJobListApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfReplicationJobListApiResponseData is nil"))
  }
  switch v.(type) {
    case import2.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = v.(import2.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case []ReplicationJob:
      p.oneOfType0 = v.([]ReplicationJob)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<files.v4.config.ReplicationJob>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<files.v4.config.ReplicationJob>"
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfReplicationJobListApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if "List<files.v4.config.ReplicationJob>" == *p.Discriminator {
    return p.oneOfType0
  }
  return nil
}

func (p *OneOfReplicationJobListApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new([]ReplicationJob)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "files.v4.config.ReplicationJob" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<files.v4.config.ReplicationJob>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<files.v4.config.ReplicationJob>"
      return nil

    }
  }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfReplicationJobListApiResponseData"))
}

func (p *OneOfReplicationJobListApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if "List<files.v4.config.ReplicationJob>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfReplicationJobListApiResponseData")
}

type OneOfDeleteDataProtectionPolicyApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *import4.TaskReference `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfDeleteDataProtectionPolicyApiResponseData() *OneOfDeleteDataProtectionPolicyApiResponseData {
  p := new(OneOfDeleteDataProtectionPolicyApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfDeleteDataProtectionPolicyApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfDeleteDataProtectionPolicyApiResponseData is nil"))
  }
  switch v.(type) {
    case import4.TaskReference:
      if nil == p.oneOfType0 {p.oneOfType0 = new(import4.TaskReference)}
      *p.oneOfType0 = v.(import4.TaskReference)
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

func (p *OneOfDeleteDataProtectionPolicyApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfDeleteDataProtectionPolicyApiResponseData) UnmarshalJSON(b []byte) error {
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
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteDataProtectionPolicyApiResponseData"))
}

func (p *OneOfDeleteDataProtectionPolicyApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfDeleteDataProtectionPolicyApiResponseData")
}

type OneOfAdminUserListApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 []AdminUser `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfAdminUserListApiResponseData() *OneOfAdminUserListApiResponseData {
  p := new(OneOfAdminUserListApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfAdminUserListApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfAdminUserListApiResponseData is nil"))
  }
  switch v.(type) {
    case []AdminUser:
      p.oneOfType0 = v.([]AdminUser)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<files.v4.config.AdminUser>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<files.v4.config.AdminUser>"
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

func (p *OneOfAdminUserListApiResponseData) GetValue() interface{} {
  if "List<files.v4.config.AdminUser>" == *p.Discriminator {
    return p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfAdminUserListApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new([]AdminUser)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "files.v4.config.AdminUser" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<files.v4.config.AdminUser>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<files.v4.config.AdminUser>"
      return nil

    }
  }
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfAdminUserListApiResponseData"))
}

func (p *OneOfAdminUserListApiResponseData) MarshalJSON() ([]byte, error) {
  if "List<files.v4.config.AdminUser>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfAdminUserListApiResponseData")
}

type OneOfUpdateTieringPolicyApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *import4.TaskReference `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfUpdateTieringPolicyApiResponseData() *OneOfUpdateTieringPolicyApiResponseData {
  p := new(OneOfUpdateTieringPolicyApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfUpdateTieringPolicyApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfUpdateTieringPolicyApiResponseData is nil"))
  }
  switch v.(type) {
    case import4.TaskReference:
      if nil == p.oneOfType0 {p.oneOfType0 = new(import4.TaskReference)}
      *p.oneOfType0 = v.(import4.TaskReference)
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

func (p *OneOfUpdateTieringPolicyApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfUpdateTieringPolicyApiResponseData) UnmarshalJSON(b []byte) error {
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
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateTieringPolicyApiResponseData"))
}

func (p *OneOfUpdateTieringPolicyApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfUpdateTieringPolicyApiResponseData")
}

type OneOfUpdateUserMappingApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *import4.TaskReference `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfUpdateUserMappingApiResponseData() *OneOfUpdateUserMappingApiResponseData {
  p := new(OneOfUpdateUserMappingApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfUpdateUserMappingApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfUpdateUserMappingApiResponseData is nil"))
  }
  switch v.(type) {
    case import4.TaskReference:
      if nil == p.oneOfType0 {p.oneOfType0 = new(import4.TaskReference)}
      *p.oneOfType0 = v.(import4.TaskReference)
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

func (p *OneOfUpdateUserMappingApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfUpdateUserMappingApiResponseData) UnmarshalJSON(b []byte) error {
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
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateUserMappingApiResponseData"))
}

func (p *OneOfUpdateUserMappingApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfUpdateUserMappingApiResponseData")
}

type OneOfNestedMountTargetListApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 []MountTarget `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfNestedMountTargetListApiResponseData() *OneOfNestedMountTargetListApiResponseData {
  p := new(OneOfNestedMountTargetListApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfNestedMountTargetListApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfNestedMountTargetListApiResponseData is nil"))
  }
  switch v.(type) {
    case []MountTarget:
      p.oneOfType0 = v.([]MountTarget)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<files.v4.config.MountTarget>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<files.v4.config.MountTarget>"
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

func (p *OneOfNestedMountTargetListApiResponseData) GetValue() interface{} {
  if "List<files.v4.config.MountTarget>" == *p.Discriminator {
    return p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfNestedMountTargetListApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new([]MountTarget)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "files.v4.config.MountTarget" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<files.v4.config.MountTarget>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<files.v4.config.MountTarget>"
      return nil

    }
  }
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfNestedMountTargetListApiResponseData"))
}

func (p *OneOfNestedMountTargetListApiResponseData) MarshalJSON() ([]byte, error) {
  if "List<files.v4.config.MountTarget>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfNestedMountTargetListApiResponseData")
}

type OneOfRecommendationApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 []Recommendation `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfRecommendationApiResponseData() *OneOfRecommendationApiResponseData {
  p := new(OneOfRecommendationApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfRecommendationApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfRecommendationApiResponseData is nil"))
  }
  switch v.(type) {
    case []Recommendation:
      p.oneOfType0 = v.([]Recommendation)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<files.v4.config.Recommendation>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<files.v4.config.Recommendation>"
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

func (p *OneOfRecommendationApiResponseData) GetValue() interface{} {
  if "List<files.v4.config.Recommendation>" == *p.Discriminator {
    return p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfRecommendationApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new([]Recommendation)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "files.v4.config.Recommendation" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<files.v4.config.Recommendation>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<files.v4.config.Recommendation>"
      return nil

    }
  }
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfRecommendationApiResponseData"))
}

func (p *OneOfRecommendationApiResponseData) MarshalJSON() ([]byte, error) {
  if "List<files.v4.config.Recommendation>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfRecommendationApiResponseData")
}

type OneOfCreateTieringConfigApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *TieringConfig `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfCreateTieringConfigApiResponseData() *OneOfCreateTieringConfigApiResponseData {
  p := new(OneOfCreateTieringConfigApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfCreateTieringConfigApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfCreateTieringConfigApiResponseData is nil"))
  }
  switch v.(type) {
    case TieringConfig:
      if nil == p.oneOfType0 {p.oneOfType0 = new(TieringConfig)}
      *p.oneOfType0 = v.(TieringConfig)
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

func (p *OneOfCreateTieringConfigApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfCreateTieringConfigApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new(TieringConfig)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "files.v4.config.TieringConfig" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(TieringConfig)}
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
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateTieringConfigApiResponseData"))
}

func (p *OneOfCreateTieringConfigApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfCreateTieringConfigApiResponseData")
}

type OneOfDeleteMountTargetSnapshotApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *import4.TaskReference `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfDeleteMountTargetSnapshotApiResponseData() *OneOfDeleteMountTargetSnapshotApiResponseData {
  p := new(OneOfDeleteMountTargetSnapshotApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfDeleteMountTargetSnapshotApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfDeleteMountTargetSnapshotApiResponseData is nil"))
  }
  switch v.(type) {
    case import4.TaskReference:
      if nil == p.oneOfType0 {p.oneOfType0 = new(import4.TaskReference)}
      *p.oneOfType0 = v.(import4.TaskReference)
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

func (p *OneOfDeleteMountTargetSnapshotApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfDeleteMountTargetSnapshotApiResponseData) UnmarshalJSON(b []byte) error {
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
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteMountTargetSnapshotApiResponseData"))
}

func (p *OneOfDeleteMountTargetSnapshotApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfDeleteMountTargetSnapshotApiResponseData")
}

type OneOfRemoveDnsEntriesApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *import4.TaskReference `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfRemoveDnsEntriesApiResponseData() *OneOfRemoveDnsEntriesApiResponseData {
  p := new(OneOfRemoveDnsEntriesApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfRemoveDnsEntriesApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfRemoveDnsEntriesApiResponseData is nil"))
  }
  switch v.(type) {
    case import4.TaskReference:
      if nil == p.oneOfType0 {p.oneOfType0 = new(import4.TaskReference)}
      *p.oneOfType0 = v.(import4.TaskReference)
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

func (p *OneOfRemoveDnsEntriesApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfRemoveDnsEntriesApiResponseData) UnmarshalJSON(b []byte) error {
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
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfRemoveDnsEntriesApiResponseData"))
}

func (p *OneOfRemoveDnsEntriesApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfRemoveDnsEntriesApiResponseData")
}

type OneOfListMountTargetVirusScanPolicyApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
  oneOfType0 []VirusScanPolicyMountTarget `json:"-"`
}

func NewOneOfListMountTargetVirusScanPolicyApiResponseData() *OneOfListMountTargetVirusScanPolicyApiResponseData {
  p := new(OneOfListMountTargetVirusScanPolicyApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfListMountTargetVirusScanPolicyApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfListMountTargetVirusScanPolicyApiResponseData is nil"))
  }
  switch v.(type) {
    case import2.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = v.(import2.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case []VirusScanPolicyMountTarget:
      p.oneOfType0 = v.([]VirusScanPolicyMountTarget)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<files.v4.config.VirusScanPolicyMountTarget>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<files.v4.config.VirusScanPolicyMountTarget>"
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfListMountTargetVirusScanPolicyApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if "List<files.v4.config.VirusScanPolicyMountTarget>" == *p.Discriminator {
    return p.oneOfType0
  }
  return nil
}

func (p *OneOfListMountTargetVirusScanPolicyApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new([]VirusScanPolicyMountTarget)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "files.v4.config.VirusScanPolicyMountTarget" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<files.v4.config.VirusScanPolicyMountTarget>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<files.v4.config.VirusScanPolicyMountTarget>"
      return nil

    }
  }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListMountTargetVirusScanPolicyApiResponseData"))
}

func (p *OneOfListMountTargetVirusScanPolicyApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if "List<files.v4.config.VirusScanPolicyMountTarget>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfListMountTargetVirusScanPolicyApiResponseData")
}

type OneOfAdDnsFailoverApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *import4.TaskReference `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfAdDnsFailoverApiResponseData() *OneOfAdDnsFailoverApiResponseData {
  p := new(OneOfAdDnsFailoverApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfAdDnsFailoverApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfAdDnsFailoverApiResponseData is nil"))
  }
  switch v.(type) {
    case import4.TaskReference:
      if nil == p.oneOfType0 {p.oneOfType0 = new(import4.TaskReference)}
      *p.oneOfType0 = v.(import4.TaskReference)
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

func (p *OneOfAdDnsFailoverApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfAdDnsFailoverApiResponseData) UnmarshalJSON(b []byte) error {
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
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfAdDnsFailoverApiResponseData"))
}

func (p *OneOfAdDnsFailoverApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfAdDnsFailoverApiResponseData")
}

type OneOfCreateVirusScanPolicyApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
  oneOfType0 *VirusScanPolicy `json:"-"`
}

func NewOneOfCreateVirusScanPolicyApiResponseData() *OneOfCreateVirusScanPolicyApiResponseData {
  p := new(OneOfCreateVirusScanPolicyApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfCreateVirusScanPolicyApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfCreateVirusScanPolicyApiResponseData is nil"))
  }
  switch v.(type) {
    case import2.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = v.(import2.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case VirusScanPolicy:
      if nil == p.oneOfType0 {p.oneOfType0 = new(VirusScanPolicy)}
      *p.oneOfType0 = v.(VirusScanPolicy)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfCreateVirusScanPolicyApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  return nil
}

func (p *OneOfCreateVirusScanPolicyApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new(VirusScanPolicy)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "files.v4.config.VirusScanPolicy" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(VirusScanPolicy)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateVirusScanPolicyApiResponseData"))
}

func (p *OneOfCreateVirusScanPolicyApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfCreateVirusScanPolicyApiResponseData")
}

