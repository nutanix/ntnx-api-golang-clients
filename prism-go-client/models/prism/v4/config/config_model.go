/*
 * Generated file models/prism/v4/config/config_model.go.
 *
 * Product version: 4.0.1-alpha-1
 *
 * Part of the Nutanix Prism Versioned APIs
 *
 * (c) 2022 Nutanix Inc.  All rights reserved
 *
 */

/*
  Configure Tasks and Monitoring
*/
package config
import (
  "bytes"
  import4 "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/common/v1/config"
  import1 "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/common/v1/response"
  "encoding/json"
  "errors"
  "fmt"
  import3 "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/prism/v4/common"
  import2 "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/prism/v4/error"
  "time"
)

/**
Status of the remote tunnel or service that is running on top of the remote tunnel. For example, pulse, alert email and so on.
*/
type CommunicationStatus struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Last changed timestamp.
  */
  LastChangedTimestamp *int64 `json:"lastChangedTimestamp,omitempty"`
  /**
  Last checked timestamp.
  */
  LastCheckedTimestamp *int64 `json:"lastCheckedTimestamp,omitempty"`
  /**
  Last successful transmission timestamp.
  */
  LastSuccessfulTransmissionTimestamp *int64 `json:"lastSuccessfulTransmissionTimestamp,omitempty"`
  
  Message *ParameterizedMessage `json:"message,omitempty"`
  
  Status *Status `json:"status,omitempty"`
}


func NewCommunicationStatus() *CommunicationStatus {
  p := new(CommunicationStatus)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "prism.v4.config.CommunicationStatus"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a1.config.CommunicationStatus"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}





type EmailConfiguration struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  The default Nutanix email ID to which alert emails would be sent.
  */
  DefaultNutanixEmail *string `json:"defaultNutanixEmail,omitempty"`
  /**
  Rules for email configuration.
  */
  EmailConfigRules []EmailConfigurationRule `json:"emailConfigRules,omitempty"`
  /**
  List of email contacts.
  */
  EmailContactList []string `json:"emailContactList,omitempty"`
  
  EmailTemplate *EmailTemplate `json:"emailTemplate,omitempty"`
  /**
  Indicates whether alert emails should be enabled or not.
  */
  Enable *bool `json:"enable,omitempty"`
  /**
  Indicates whether alert emails should be enabled or not on default Nutanix email ID.
  */
  EnableDefaultNutanixEmail *bool `json:"enableDefaultNutanixEmail,omitempty"`
  /**
  Indicates whether alert email digest should be enabled or not.
  */
  EnableEmailDigest *bool `json:"enableEmailDigest,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import1.ApiLink `json:"links,omitempty"`
  /**
  Send alert email digest only if there are one or more alerts.
  */
  SkipEmptyAlertEmailDigest *bool `json:"skipEmptyAlertEmailDigest,omitempty"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
  
  TunnelDetails *RemoteTunnelDetails `json:"tunnelDetails,omitempty"`
}


func NewEmailConfiguration() *EmailConfiguration {
  p := new(EmailConfiguration)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "prism.v4.config.EmailConfiguration"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a1.config.EmailConfiguration"}
  p.UnknownFields_ = map[string]interface{}{}

  p.Enable = new(bool)
  *p.Enable = false
  p.EnableDefaultNutanixEmail = new(bool)
  *p.EnableDefaultNutanixEmail = false
  p.EnableEmailDigest = new(bool)
  *p.EnableEmailDigest = false
  p.SkipEmptyAlertEmailDigest = new(bool)
  *p.SkipEmptyAlertEmailDigest = false


  return p
}




/**
REST response for all response codes in api path /prism/v4.0.a1/config/email Put operation
*/
type EmailConfigurationApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfEmailConfigurationApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewEmailConfigurationApiResponse() *EmailConfigurationApiResponse {
  p := new(EmailConfigurationApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "prism.v4.config.EmailConfigurationApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a1.config.EmailConfigurationApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *EmailConfigurationApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *EmailConfigurationApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfEmailConfigurationApiResponseData()
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
Status of remote tunnel that is used to send alert emails.
*/
type EmailConfigurationRule struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Cluster UUIDs to which this rule applies to.
  */
  ClusterUuids []string `json:"clusterUuids,omitempty"`
  
  ImpactTypes []import3.ImpactType `json:"impactTypes,omitempty"`
  
  IncludeGlobalEmailContactList *bool `json:"includeGlobalEmailContactList,omitempty"`
  /**
  List of phrases to match the alert with.
  */
  MatchPhrases []string `json:"matchPhrases,omitempty"`
  /**
  List of recipients who will receive emails.
  */
  Recipients []string `json:"recipients,omitempty"`
  
  Severities []import3.Severity `json:"severities,omitempty"`
}


func NewEmailConfigurationRule() *EmailConfigurationRule {
  p := new(EmailConfigurationRule)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "prism.v4.config.EmailConfigurationRule"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a1.config.EmailConfigurationRule"}
  p.UnknownFields_ = map[string]interface{}{}

  p.IncludeGlobalEmailContactList = new(bool)
  *p.IncludeGlobalEmailContactList = false


  return p
}




/**
Email template.
*/
type EmailTemplate struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Suffix for the email body.
  */
  BodySuffix *string `json:"bodySuffix,omitempty"`
  /**
  Prefix for the email subject.
  */
  SubjectPrefix *string `json:"subjectPrefix,omitempty"`
}


func NewEmailTemplate() *EmailTemplate {
  p := new(EmailTemplate)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "prism.v4.config.EmailTemplate"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a1.config.EmailTemplate"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Details of entity.
*/
type EntityReference struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Globally unique identifier of the entity.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  Entity type identified as 'namespace:module[:submodule]:entityType such vmm:ahv:vm'
  */
  Rel *string `json:"rel,omitempty"`
}


func NewEntityReference() *EntityReference {
  p := new(EntityReference)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "prism.v4.config.EntityReference"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a1.config.EntityReference"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Http proxy used to establish remote tunnel.
*/
type HttpProxy struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  AddressValue *import4.IPAddressOrFQDN `json:"addressValue,omitempty"`
  /**
  Proxy name.
  */
  Name *string `json:"name,omitempty"`
  /**
  Password for proxy authentication.
  */
  Password *string `json:"password,omitempty"`
  /**
  Port on which proxy is binding.
  */
  Port *int `json:"port,omitempty"`
  /**
  Proxy types to send applicable traffic.
  */
  ProxyTypes []ProxyType `json:"proxyTypes,omitempty"`
  /**
  Username for proxy authentication.
  */
  Username *string `json:"username,omitempty"`
}


func NewHttpProxy() *HttpProxy {
  p := new(HttpProxy)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "prism.v4.config.HttpProxy"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a1.config.HttpProxy"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Reference to the owner of the task.
*/
type OwnerReference struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Globally unique identifier of the task owner.
  */
  ExtId *string `json:"extId,omitempty"`
}


func NewOwnerReference() *OwnerReference {
  p := new(OwnerReference)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "prism.v4.config.OwnerReference"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a1.config.OwnerReference"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
To store any custom message and attributes.
*/
type ParameterizedMessage struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  Attributes []import4.KVStringPair `json:"attributes,omitempty"`
  
  Message *string `json:"message,omitempty"`
}


func NewParameterizedMessage() *ParameterizedMessage {
  p := new(ParameterizedMessage)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "prism.v4.config.ParameterizedMessage"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a1.config.ParameterizedMessage"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}





type ProxyType int

const(
  PROXYTYPE_UNKNOWN ProxyType = 0
  PROXYTYPE_REDACTED ProxyType = 1
  PROXYTYPE_HTTP ProxyType = 2
  PROXYTYPE_HTTPS ProxyType = 3
  PROXYTYPE_SOCKS ProxyType = 4
)

// returns the name of the enum given an ordinal number
func (e *ProxyType) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "HTTP",
    "HTTPS",
    "SOCKS",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *ProxyType) index(name string) ProxyType {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "HTTP",
    "HTTPS",
    "SOCKS",
  }
  for idx := range names {
    if names[idx] == name {
      return ProxyType(idx)
    }
  }
  return PROXYTYPE_UNKNOWN
}

func (e *ProxyType) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for ProxyType:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *ProxyType) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e ProxyType) Ref() *ProxyType {
  return &e
}


/**
Remote tunnel details associated with the email configuration.
*/
type RemoteTunnelDetails struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  ConnectionStatus *CommunicationStatus `json:"connectionStatus,omitempty"`
  
  HttpProxy *HttpProxy `json:"httpProxy,omitempty"`
  
  ServiceCenter *ServiceCenter `json:"serviceCenter,omitempty"`
  
  TransportStatus *CommunicationStatus `json:"transportStatus,omitempty"`
}


func NewRemoteTunnelDetails() *RemoteTunnelDetails {
  p := new(RemoteTunnelDetails)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "prism.v4.config.RemoteTunnelDetails"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a1.config.RemoteTunnelDetails"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Service center to which remote tunnel is connected at Nutanix's end.
*/
type ServiceCenter struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  IP address of the service center.
  */
  IpAddress *string `json:"ipAddress,omitempty"`
  /**
  Name of service center.
  */
  Name *string `json:"name,omitempty"`
  /**
  Port number.
  */
  Port *int `json:"port,omitempty"`
  /**
  Username.
  */
  Username *string `json:"username,omitempty"`
}


func NewServiceCenter() *ServiceCenter {
  p := new(ServiceCenter)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "prism.v4.config.ServiceCenter"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a1.config.ServiceCenter"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}





type Status int

const(
  STATUS_UNKNOWN Status = 0
  STATUS_REDACTED Status = 1
  STATUS_SUCCESS Status = 2
  STATUS_FAILED Status = 3
  STATUS_RETRYING Status = 4
  STATUS_DISABLED Status = 5
)

// returns the name of the enum given an ordinal number
func (e *Status) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "SUCCESS",
    "FAILED",
    "RETRYING",
    "DISABLED",
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
    "SUCCESS",
    "FAILED",
    "RETRYING",
    "DISABLED",
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
The task object tracking an async operation.
*/
type Task struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  UTC date and time in RFC-3339 format when task was completed.
  */
  CompletedTime *time.Time `json:"completedTime,omitempty"`
  /**
  Additional details post operation completion to guide further actions.
  */
  CompletionDetails []import4.KVPair `json:"completionDetails,omitempty"`
  /**
  UTC date and time in RFC-3339 format when task was created.
  */
  CreatedTime *time.Time `json:"createdTime,omitempty"`
  /**
  Reference to entities associated with the task.
  */
  EntitiesAffected []EntityReference `json:"entitiesAffected,omitempty"`
  /**
  Error details in case of task failure.
  */
  ErrorMessages []import2.AppMessage `json:"errorMessages,omitempty"`
  /**
  Globally unique identifier of a task.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  Signifies if task can be cancelled.
  */
  IsCancelable *bool `json:"isCancelable,omitempty"`
  /**
  Provides the error in the absence of well defined error message for tasks created via legacy APIs.
  */
  LegacyErrorMessage *string `json:"legacyErrorMessage,omitempty"`
  /**
  Operation being tracked by the task.
  */
  Operation *string `json:"operation,omitempty"`
  /**
  Description of the operation being tracked by the task.
  */
  OperationDescription *string `json:"operationDescription,omitempty"`
  
  OwnedBy *OwnerReference `json:"ownedBy,omitempty"`
  
  ParentTask *TaskReferenceInternal `json:"parentTask,omitempty"`
  /**
  Task progress expressed as a percentage.
  */
  ProgressPercentage *int `json:"progressPercentage,omitempty"`
  /**
  UTC date and time in RFC-3339 format when task was started.
  */
  StartedTime *time.Time `json:"startedTime,omitempty"`
  
  Status *TaskStatus `json:"status,omitempty"`
  /**
  List of steps completed as part of task.
  */
  SubSteps []TaskStep `json:"subSteps,omitempty"`
  /**
  Reference to tasks spawned as children of current task.
  */
  SubTasks []TaskReferenceInternal `json:"subTasks,omitempty"`
}


func NewTask() *Task {
  p := new(Task)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "prism.v4.config.Task"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a1.config.Task"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST response for all response codes in api path /prism/v4.0.a1/config/tasks/{taskExtId}/$actions/cancel Post operation
*/
type TaskCancelResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfTaskCancelResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewTaskCancelResponse() *TaskCancelResponse {
  p := new(TaskCancelResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "prism.v4.config.TaskCancelResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a1.config.TaskCancelResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *TaskCancelResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *TaskCancelResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfTaskCancelResponseData()
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
REST response for all response codes in api path /prism/v4.0.a1/config/tasks/{taskExtId} Get operation
*/
type TaskGetResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfTaskGetResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewTaskGetResponse() *TaskGetResponse {
  p := new(TaskGetResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "prism.v4.config.TaskGetResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a1.config.TaskGetResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *TaskGetResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *TaskGetResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfTaskGetResponseData()
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
Reference to a task tracking the async operation.
*/
type TaskReference struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Globally unique identifier of a task.
  */
  ExtId *string `json:"extId,omitempty"`
}


func NewTaskReference() *TaskReference {
  p := new(TaskReference)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "prism.v4.config.TaskReference"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a1.config.TaskReference"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Reference to task associated with current task.
*/
type TaskReferenceInternal struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Glocally unique identifier of the task.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  The URL at which the entity described by this link can be accessed.
  */
  Href *string `json:"href,omitempty"`
  /**
  A name that identifies the relationship of this link to the object that is returned by the URL.  The special value of "self" identifies the URL for the object.
  */
  Rel *string `json:"rel,omitempty"`
}


func NewTaskReferenceInternal() *TaskReferenceInternal {
  p := new(TaskReferenceInternal)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "prism.v4.config.TaskReferenceInternal"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a1.config.TaskReferenceInternal"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Status of the task.
*/
type TaskStatus int

const(
  TASKSTATUS_UNKNOWN TaskStatus = 0
  TASKSTATUS_REDACTED TaskStatus = 1
  TASKSTATUS_QUEUED TaskStatus = 2
  TASKSTATUS_RUNNING TaskStatus = 3
  TASKSTATUS_CANCELING TaskStatus = 4
  TASKSTATUS_SUCCEEDED TaskStatus = 5
  TASKSTATUS_FAILED TaskStatus = 6
  TASKSTATUS_CANCELED TaskStatus = 7
)

// returns the name of the enum given an ordinal number
func (e *TaskStatus) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "QUEUED",
    "RUNNING",
    "CANCELING",
    "SUCCEEDED",
    "FAILED",
    "CANCELED",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *TaskStatus) index(name string) TaskStatus {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "QUEUED",
    "RUNNING",
    "CANCELING",
    "SUCCEEDED",
    "FAILED",
    "CANCELED",
  }
  for idx := range names {
    if names[idx] == name {
      return TaskStatus(idx)
    }
  }
  return TASKSTATUS_UNKNOWN
}

func (e *TaskStatus) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for TaskStatus:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *TaskStatus) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e TaskStatus) Ref() *TaskStatus {
  return &e
}


/**
A single step in the task.
*/
type TaskStep struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Message describing the step completed of the task.
  */
  Name *string `json:"name,omitempty"`
}


func NewTaskStep() *TaskStep {
  p := new(TaskStep)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "prism.v4.config.TaskStep"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a1.config.TaskStep"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}



type OneOfEmailConfigurationApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
  oneOfType0 *EmailConfiguration `json:"-"`
}

func NewOneOfEmailConfigurationApiResponseData() *OneOfEmailConfigurationApiResponseData {
  p := new(OneOfEmailConfigurationApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfEmailConfigurationApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfEmailConfigurationApiResponseData is nil"))
  }
  switch v.(type) {
    case import2.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = v.(import2.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case EmailConfiguration:
      if nil == p.oneOfType0 {p.oneOfType0 = new(EmailConfiguration)}
      *p.oneOfType0 = v.(EmailConfiguration)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfEmailConfigurationApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  return nil
}

func (p *OneOfEmailConfigurationApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "prism.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new(EmailConfiguration)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "prism.v4.config.EmailConfiguration" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(EmailConfiguration)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfEmailConfigurationApiResponseData"))
}

func (p *OneOfEmailConfigurationApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfEmailConfigurationApiResponseData")
}

type OneOfTaskCancelResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
  oneOfType2001 *interface{} `json:"-"`
}

func NewOneOfTaskCancelResponseData() *OneOfTaskCancelResponseData {
  p := new(OneOfTaskCancelResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfTaskCancelResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfTaskCancelResponseData is nil"))
  }
  if nil == v {
    if nil == p.oneOfType2001 {p.oneOfType2001 = new(interface {})}
    *p.oneOfType2001 = nil
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

func (p *OneOfTaskCancelResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if "EMPTY" == *p.Discriminator {
    return *p.oneOfType2001
  }
  return nil
}

func (p *OneOfTaskCancelResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType2001 := new(interface {})
  if err := json.Unmarshal(b, vOneOfType2001); err == nil {
    if nil == *vOneOfType2001 {
      if nil == p.oneOfType2001 {p.oneOfType2001 = new(interface {})}
      *p.oneOfType2001 = nil
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "EMPTY"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "EMPTY"
      return nil
    }
  }
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "prism.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfTaskCancelResponseData"))
}

func (p *OneOfTaskCancelResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if "EMPTY" == *p.Discriminator {
    return json.Marshal(p.oneOfType2001)
  }
  return nil, errors.New("No value to marshal for OneOfTaskCancelResponseData")
}

type OneOfTaskGetResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType2001 *Task `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfTaskGetResponseData() *OneOfTaskGetResponseData {
  p := new(OneOfTaskGetResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfTaskGetResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfTaskGetResponseData is nil"))
  }
  switch v.(type) {
    case Task:
      if nil == p.oneOfType2001 {p.oneOfType2001 = new(Task)}
      *p.oneOfType2001 = v.(Task)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType2001.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType2001.ObjectType_
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

func (p *OneOfTaskGetResponseData) GetValue() interface{} {
  if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
    return *p.oneOfType2001
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfTaskGetResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType2001 := new(Task)
  if err := json.Unmarshal(b, vOneOfType2001); err == nil {
    if "prism.v4.config.Task" == *vOneOfType2001.ObjectType_ {
          if nil == p.oneOfType2001 {p.oneOfType2001 = new(Task)}
      *p.oneOfType2001 = *vOneOfType2001
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType2001.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType2001.ObjectType_
      return nil
    }
    }
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "prism.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfTaskGetResponseData"))
}

func (p *OneOfTaskGetResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType2001)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfTaskGetResponseData")
}

