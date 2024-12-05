/*
 * Generated file models/files/v4/config/config_model.go.
 *
 * Product version: 4.0.1
 *
 * Part of the Nutanix Files APIs
 *
 * (c) 2024 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module files.v4.config of Nutanix Files APIs
*/
package config

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	import1 "github.com/nutanix/ntnx-api-golang-clients/files-go-client/v4/models/common/v1/config"
	import2 "github.com/nutanix/ntnx-api-golang-clients/files-go-client/v4/models/common/v1/response"
	import4 "github.com/nutanix/ntnx-api-golang-clients/files-go-client/v4/models/files/v4/error"
	import3 "github.com/nutanix/ntnx-api-golang-clients/files-go-client/v4/models/prism/v4/config"
	"time"
)

/*
AWS configuration for tiering profile.
*/
type AWSConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Access key.
	*/
	AccessKey *string `json:"accessKey"`
	/*
	  Bucket location.
	*/
	BucketLocation *string `json:"bucketLocation,omitempty"`
	/*
	  Bucket name.
	*/
	BucketName *string `json:"bucketName,omitempty"`
	/*
	  Secret key.
	*/
	SecretKey *string `json:"secretKey,omitempty"`
}

func (p *AWSConfig) MarshalJSON() ([]byte, error) {
	type AWSConfigProxy AWSConfig
	return json.Marshal(struct {
		*AWSConfigProxy
		AccessKey *string `json:"accessKey,omitempty"`
	}{
		AWSConfigProxy: (*AWSConfigProxy)(p),
		AccessKey:      p.AccessKey,
	})
}

func NewAWSConfig() *AWSConfig {
	p := new(AWSConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.AWSConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Access type of the mount target.
*/
type AccessType int

const (
	ACCESSTYPE_UNKNOWN    AccessType = 0
	ACCESSTYPE_REDACTED   AccessType = 1
	ACCESSTYPE_READ_WRITE AccessType = 2
	ACCESSTYPE_READ_ONLY  AccessType = 3
	ACCESSTYPE_NO_ACCESS  AccessType = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *AccessType) name(index int) string {
	names := [...]string{
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

// Returns the name of the enum
func (e AccessType) GetName() string {
	index := int(e)
	names := [...]string{
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

// Returns the enum type given a string value
func (e *AccessType) index(name string) AccessType {
	names := [...]string{
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

/*
AD domain associated with the file server.
*/
type ActiveDirectoryConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Credential *Credential `json:"credential,omitempty"`
	/*
	  AD domain name associated with the file server. Ex. child4.afs.minerva.com
	*/
	DomainName *string `json:"domainName,omitempty"`
	/*
	  RFC 2307 ENABLED (true, false).
	*/
	IsRfc2307Enabled *bool `json:"isRfc2307Enabled,omitempty"`
	/*
	  AD domain organizational unit associated with the file server. Ex. OU=AFS represents an organizational unit named AFS.
	*/
	OrganizationUnit *string `json:"organizationUnit,omitempty"`
	/*
	  Use a specific domain controller for the join-domain operation in a multi-DC active directory setup. By default, AFS discovers a site local domain controller for join-domain operation. The preferred domain controller cannot be an IP address. It has to be FQDN (example: dc_name.dom.companyname.com)
	*/
	PreferredDomainController *string `json:"preferredDomainController,omitempty"`

	ProtocolType *ProtocolType `json:"protocolType,omitempty"`
	/*
	  AD user or group name as 'name' or 'NETBIOS\name' format.
	*/
	ShouldAddUserAsFsAdmin *bool `json:"shouldAddUserAsFsAdmin,omitempty"`
	/*
	  If the joined active directory is down or not reachable, then set this field to true to leave the domain by bypassing all AD interactions. Ensure to remove the file server machine account manually afterwards.
	*/
	ShouldForceLeaveDomain *bool `json:"shouldForceLeaveDomain,omitempty"`
	/*
	  If a machine account with the same name as file server name is present on the Ad, then overwrite it during join-domain operation.
	*/
	ShouldOverwriteUserAccount *bool `json:"shouldOverwriteUserAccount,omitempty"`
	/*
	  Use the same AD credential for DNS.
	*/
	ShouldUseSameCredentialsForDns *bool `json:"shouldUseSameCredentialsForDns,omitempty"`
	/*
	  Indicates whether the AD credentials will be validated or skipped.
	*/
	ShouldValidateAdCredential *bool `json:"shouldValidateAdCredential,omitempty"`
}

func NewActiveDirectoryConfig() *ActiveDirectoryConfig {
	p := new(ActiveDirectoryConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.ActiveDirectoryConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	p.IsRfc2307Enabled = new(bool)
	*p.IsRfc2307Enabled = false
	p.ShouldAddUserAsFsAdmin = new(bool)
	*p.ShouldAddUserAsFsAdmin = false
	p.ShouldForceLeaveDomain = new(bool)
	*p.ShouldForceLeaveDomain = false
	p.ShouldOverwriteUserAccount = new(bool)
	*p.ShouldOverwriteUserAccount = false
	p.ShouldUseSameCredentialsForDns = new(bool)
	*p.ShouldUseSameCredentialsForDns = false
	p.ShouldValidateAdCredential = new(bool)
	*p.ShouldValidateAdCredential = false

	return p
}

/*
Anonymous identifier of the mount target.
*/
type AnonymousIdentifier struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Group identifier. Default value of the group identifier is -2.
	*/
	Gid *int64 `json:"gid,omitempty"`
	/*
	  User identifier. Default value of the user identifier is -2.
	*/
	Uid *int64 `json:"uid,omitempty"`
}

func NewAnonymousIdentifier() *AnonymousIdentifier {
	p := new(AnonymousIdentifier)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.AnonymousIdentifier"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Antivirus server model.
*/
type AntivirusServer struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Address *import1.IPAddressOrFQDN `json:"address"`

	ConnectionStatus *ConnectionStatus `json:"connectionStatus,omitempty"`
	/*
	  Antivirus server description.
	*/
	Description *string `json:"description,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Internet Content Adaptation Protocol (ICAP) service name.
	*/
	PartnerServiceName *string `json:"partnerServiceName,omitempty"`
	/*
	  Antivirus server port.
	*/
	Port *int `json:"port"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *AntivirusServer) MarshalJSON() ([]byte, error) {
	type AntivirusServerProxy AntivirusServer
	return json.Marshal(struct {
		*AntivirusServerProxy
		Address *import1.IPAddressOrFQDN `json:"address,omitempty"`
		Port    *int                     `json:"port,omitempty"`
	}{
		AntivirusServerProxy: (*AntivirusServerProxy)(p),
		Address:              p.Address,
		Port:                 p.Port,
	})
}

func NewAntivirusServer() *AntivirusServer {
	p := new(AntivirusServer)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.AntivirusServer"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Azure configuration for tiering profile.
*/
type AzureConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Container name for Azure.
	*/
	ContainerName *string `json:"containerName,omitempty"`
	/*
	  Storage account key for Azure.
	*/
	StorageAccountKey *string `json:"storageAccountKey,omitempty"`
	/*
	  Storage account name for Azure.
	*/
	StorageAccountName *string `json:"storageAccountName"`
}

func (p *AzureConfig) MarshalJSON() ([]byte, error) {
	type AzureConfigProxy AzureConfig
	return json.Marshal(struct {
		*AzureConfigProxy
		StorageAccountName *string `json:"storageAccountName,omitempty"`
	}{
		AzureConfigProxy:   (*AzureConfigProxy)(p),
		StorageAccountName: p.StorageAccountName,
	})
}

func NewAzureConfig() *AzureConfig {
	p := new(AzureConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.AzureConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Access type of the mount target.
*/
type BackupAccessType int

const (
	BACKUPACCESSTYPE_UNKNOWN    BackupAccessType = 0
	BACKUPACCESSTYPE_REDACTED   BackupAccessType = 1
	BACKUPACCESSTYPE_READ_WRITE BackupAccessType = 2
	BACKUPACCESSTYPE_READ_ONLY  BackupAccessType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *BackupAccessType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"READ_WRITE",
		"READ_ONLY",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e BackupAccessType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"READ_WRITE",
		"READ_ONLY",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *BackupAccessType) index(name string) BackupAccessType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"READ_WRITE",
		"READ_ONLY",
	}
	for idx := range names {
		if names[idx] == name {
			return BackupAccessType(idx)
		}
	}
	return BACKUPACCESSTYPE_UNKNOWN
}

func (e *BackupAccessType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for BackupAccessType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *BackupAccessType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e BackupAccessType) Ref() *BackupAccessType {
	return &e
}

/*
Backup server details along with access permission type.
*/
type BackupServerAccessControlBlock struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	AccessType *BackupAccessType `json:"accessType"`

	Address *import1.IPAddressOrFQDN `json:"address"`
	/*
	  Mount target external identifier for the partner server.
	*/
	MountTargetExtId *string `json:"mountTargetExtId,omitempty"`

	OperationType *OperationType `json:"operationType"`
}

func (p *BackupServerAccessControlBlock) MarshalJSON() ([]byte, error) {
	type BackupServerAccessControlBlockProxy BackupServerAccessControlBlock
	return json.Marshal(struct {
		*BackupServerAccessControlBlockProxy
		AccessType    *BackupAccessType        `json:"accessType,omitempty"`
		Address       *import1.IPAddressOrFQDN `json:"address,omitempty"`
		OperationType *OperationType           `json:"operationType,omitempty"`
	}{
		BackupServerAccessControlBlockProxy: (*BackupServerAccessControlBlockProxy)(p),
		AccessType:                          p.AccessType,
		Address:                             p.Address,
		OperationType:                       p.OperationType,
	})
}

func NewBackupServerAccessControlBlock() *BackupServerAccessControlBlock {
	p := new(BackupServerAccessControlBlock)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.BackupServerAccessControlBlock"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Configuration for partner server of type BACKUP.
*/
type BackupServerConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of backup client details along with request access permissions.
	*/
	BackupClientAccessConfigs []BackupServerAccessControlBlock `json:"backupClientAccessConfigs,omitempty"`
	/*
	  The expiry time in seconds for the backup client to access the shares. The default value is 24 hours.
	*/
	ExpiryInSecs *int64 `json:"expiryInSecs,omitempty"`
}

func NewBackupServerConfig() *BackupServerConfig {
	p := new(BackupServerConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.BackupServerConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	p.ExpiryInSecs = new(int64)
	*p.ExpiryInSecs = 86400

	return p
}

/*
Blocked clients due to read-only, no access or read-write access.
*/
type BlockedClient struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of clients with no access.
	*/
	NoAccessFilters []ClientBlockingFilter `json:"noAccessFilters,omitempty"`
	/*
	  List of clients with read only access.
	*/
	RoAccessFilters []ClientBlockingFilter `json:"roAccessFilters,omitempty"`
	/*
	  List of clients with read-write access This filter is only applicable for mount target.
	*/
	RwAccessFilters []ClientBlockingFilter `json:"rwAccessFilters,omitempty"`
}

func NewBlockedClient() *BlockedClient {
	p := new(BlockedClient)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.BlockedClient"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Indicates that a client is blocked from notifications. User can specify user and client IP for which notification needs to block.
*/
type BlockedNotificationClient struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	ClientDetails *ClientDetails `json:"clientDetails,omitempty"`

	User *NotificationPolicyUser `json:"user,omitempty"`
}

func NewBlockedNotificationClient() *BlockedNotificationClient {
	p := new(BlockedNotificationClient)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.BlockedNotificationClient"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Operation performed on the object.
*/
type ChangeOperationType int

const (
	CHANGEOPERATIONTYPE_UNKNOWN     ChangeOperationType = 0
	CHANGEOPERATIONTYPE_REDACTED    ChangeOperationType = 1
	CHANGEOPERATIONTYPE_MODIFY      ChangeOperationType = 2
	CHANGEOPERATIONTYPE_META_MODIFY ChangeOperationType = 3
	CHANGEOPERATIONTYPE_ADD         ChangeOperationType = 4
	CHANGEOPERATIONTYPE_DELETE      ChangeOperationType = 5
	CHANGEOPERATIONTYPE_RENAME      ChangeOperationType = 6
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ChangeOperationType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"MODIFY",
		"META_MODIFY",
		"ADD",
		"DELETE",
		"RENAME",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ChangeOperationType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"MODIFY",
		"META_MODIFY",
		"ADD",
		"DELETE",
		"RENAME",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ChangeOperationType) index(name string) ChangeOperationType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"MODIFY",
		"META_MODIFY",
		"ADD",
		"DELETE",
		"RENAME",
	}
	for idx := range names {
		if names[idx] == name {
			return ChangeOperationType(idx)
		}
	}
	return CHANGEOPERATIONTYPE_UNKNOWN
}

func (e *ChangeOperationType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ChangeOperationType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ChangeOperationType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ChangeOperationType) Ref() *ChangeOperationType {
	return &e
}

/*
Snapshot changed content details. This is a read-only field.
*/
type ChangedContent struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Access time of the object in date-time format.
	*/
	AccessTime *time.Time `json:"accessTime,omitempty"`
	/*
	  Change time of the object in date-time format.
	*/
	ChangeTime *time.Time `json:"changeTime,omitempty"`
	/*
	  Create time of the object in date-time format..
	*/
	CreationTime *time.Time `json:"creationTime,omitempty"`
	/*
	  Inode number of the object.
	*/
	InodeNumber *int64 `json:"inodeNumber,omitempty"`

	ObjectType *ObjectType `json:"objectType,omitempty"`
	/*
	  Old path of the object in case of rename.
	*/
	OldPath *string `json:"oldPath,omitempty"`

	OperationType *ChangeOperationType `json:"operationType,omitempty"`
	/*
	  Path of the object.
	*/
	Path *string `json:"path,omitempty"`
	/*
	  Size of the object in bytes.
	*/
	SizeBytes *int64 `json:"sizeBytes,omitempty"`
}

func NewChangedContent() *ChangedContent {
	p := new(ChangedContent)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.ChangedContent"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Blocked clients filter
*/
type ClientBlockingFilter struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of GIDs(Group Identifier).
	*/
	GidList []int64 `json:"gidList,omitempty"`
	/*
	  IP address list.
	*/
	IpList []import1.IPAddress `json:"ipList,omitempty"`
	/*
	  Block all IP addresses.
	*/
	IsAllIpsBlocked *bool `json:"isAllIpsBlocked,omitempty"`
	/*
	  List of SIDs(Security Identifier).
	*/
	SidList []string `json:"sidList,omitempty"`
	/*
	  List of UIDs(User Identifier).
	*/
	UidList []int64 `json:"uidList,omitempty"`
	/*
	  Vendor name of the partner server. For partner type ANTIVIRUS, vendor name denotes the icapServiceName.
	*/
	VendorName *string `json:"vendorName,omitempty"`
}

func NewClientBlockingFilter() *ClientBlockingFilter {
	p := new(ClientBlockingFilter)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.ClientBlockingFilter"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	p.IsAllIpsBlocked = new(bool)
	*p.IsAllIpsBlocked = false

	return p
}

/*
Client details entity.
*/
type ClientDetails struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Address *import1.IPAddressOrFQDN `json:"address,omitempty"`
	/*
	  The flag that defines whether the client IP is a backup server or not.
	*/
	IsBackup *bool `json:"isBackup,omitempty"`
	/*
	  Client port.
	*/
	Port *int `json:"port,omitempty"`
}

func NewClientDetails() *ClientDetails {
	p := new(ClientDetails)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.ClientDetails"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	p.IsBackup = new(bool)
	*p.IsBackup = false

	return p
}

/*
Access exception client of the mount target.
*/
type ClientException struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	AccessType *AccessType `json:"accessType,omitempty"`
	/*
	  Comma-separated client list
	*/
	Clients *string `json:"clients,omitempty"`

	SquashType *SquashType `json:"squashType,omitempty"`
}

func NewClientException() *ClientException {
	p := new(ClientException)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.ClientException"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /files/v4.0/config/file-servers/{fileServerExtId}/mount-targets/{extId}/$actions/clone Post operation
*/
type CloneMountTargetApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCloneMountTargetApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewCloneMountTargetApiResponse() *CloneMountTargetApiResponse {
	p := new(CloneMountTargetApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.CloneMountTargetApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CloneMountTargetApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CloneMountTargetApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCloneMountTargetApiResponseData()
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
REST response for all response codes in API path /files/v4.0/config/file-servers/{fileServerExtId}/mount-targets/{mountTargetExtId}/$actions/compute-snapshot-change Post operation
*/
type ComputeSnapshotChangeApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfComputeSnapshotChangeApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewComputeSnapshotChangeApiResponse() *ComputeSnapshotChangeApiResponse {
	p := new(ComputeSnapshotChangeApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.ComputeSnapshotChangeApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ComputeSnapshotChangeApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ComputeSnapshotChangeApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfComputeSnapshotChangeApiResponseData()
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
Compute snapshot change model details.
*/
type ComputeSnapshotChangeSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Reference snapshot associated with changed files metdata.
	*/
	BaseSnapshotExtId *string `json:"baseSnapshotExtId,omitempty"`
	/*
	  If this flag is set to true, the object relative path is returned in GET API response. Otherwise FQDN path is returned by default.
	*/
	ShouldShowRelativePath *bool `json:"shouldShowRelativePath,omitempty"`
	/*
	  If the smb client used for backup operation does not support DFS referral capability set this flag to false.
	*/
	ShouldUseDFSNamespace *bool `json:"shouldUseDFSNamespace,omitempty"`
	/*
	  Current snapshot associated with changed files metdata.
	*/
	SnapshotExtId *string `json:"snapshotExtId"`
}

func (p *ComputeSnapshotChangeSpec) MarshalJSON() ([]byte, error) {
	type ComputeSnapshotChangeSpecProxy ComputeSnapshotChangeSpec
	return json.Marshal(struct {
		*ComputeSnapshotChangeSpecProxy
		SnapshotExtId *string `json:"snapshotExtId,omitempty"`
	}{
		ComputeSnapshotChangeSpecProxy: (*ComputeSnapshotChangeSpecProxy)(p),
		SnapshotExtId:                  p.SnapshotExtId,
	})
}

func NewComputeSnapshotChangeSpec() *ComputeSnapshotChangeSpec {
	p := new(ComputeSnapshotChangeSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.ComputeSnapshotChangeSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	p.ShouldShowRelativePath = new(bool)
	*p.ShouldShowRelativePath = false
	p.ShouldUseDFSNamespace = new(bool)
	*p.ShouldUseDFSNamespace = true

	return p
}

/*
Connection status values. This is a read-only field.
*/
type ConnectionStatus int

const (
	CONNECTIONSTATUS_UNKNOWN       ConnectionStatus = 0
	CONNECTIONSTATUS_REDACTED      ConnectionStatus = 1
	CONNECTIONSTATUS_NOT_TESTED    ConnectionStatus = 2
	CONNECTIONSTATUS_NOT_REACHABLE ConnectionStatus = 3
	CONNECTIONSTATUS_OK            ConnectionStatus = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ConnectionStatus) name(index int) string {
	names := [...]string{
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

// Returns the name of the enum
func (e ConnectionStatus) GetName() string {
	index := int(e)
	names := [...]string{
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

// Returns the enum type given a string value
func (e *ConnectionStatus) index(name string) ConnectionStatus {
	names := [...]string{
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

/*
REST response for all response codes in API path /files/v4.0/config/file-servers/{fileServerExtId}/anti-virus-servers Post operation
*/
type CreateAntivirusServerApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateAntivirusServerApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewCreateAntivirusServerApiResponse() *CreateAntivirusServerApiResponse {
	p := new(CreateAntivirusServerApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.CreateAntivirusServerApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
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

/*
REST response for all response codes in API path /files/v4.0/config/file-servers/{fileServerExtId}/mount-targets Post operation
*/
type CreateMountTargetApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateMountTargetApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewCreateMountTargetApiResponse() *CreateMountTargetApiResponse {
	p := new(CreateMountTargetApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.CreateMountTargetApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
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

/*
REST response for all response codes in API path /files/v4.0/config/file-servers/{fileServerExtId}/mount-targets/{mountTargetExtId}/snapshots Post operation
*/
type CreateMountTargetSnapshotApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateMountTargetSnapshotApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewCreateMountTargetSnapshotApiResponse() *CreateMountTargetSnapshotApiResponse {
	p := new(CreateMountTargetSnapshotApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.CreateMountTargetSnapshotApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreateMountTargetSnapshotApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreateMountTargetSnapshotApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreateMountTargetSnapshotApiResponseData()
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
REST response for all response codes in API path /files/v4.0/config/file-servers/{fileServerExtId}/notification-policies Post operation
*/
type CreateNotificationPolicyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateNotificationPolicyApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewCreateNotificationPolicyApiResponse() *CreateNotificationPolicyApiResponse {
	p := new(CreateNotificationPolicyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.CreateNotificationPolicyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreateNotificationPolicyApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreateNotificationPolicyApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreateNotificationPolicyApiResponseData()
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
REST response for all response codes in API path /files/v4.0/config/file-servers/{fileServerExtId}/object-store-profiles Post operation
*/
type CreateObjectStoreProfileApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateObjectStoreProfileApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewCreateObjectStoreProfileApiResponse() *CreateObjectStoreProfileApiResponse {
	p := new(CreateObjectStoreProfileApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.CreateObjectStoreProfileApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreateObjectStoreProfileApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreateObjectStoreProfileApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreateObjectStoreProfileApiResponseData()
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
REST response for all response codes in API path /files/v4.0/config/file-servers/{fileServerExtId}/partner-servers Post operation
*/
type CreatePartnerServerApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreatePartnerServerApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewCreatePartnerServerApiResponse() *CreatePartnerServerApiResponse {
	p := new(CreatePartnerServerApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.CreatePartnerServerApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreatePartnerServerApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreatePartnerServerApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreatePartnerServerApiResponseData()
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
REST response for all response codes in API path /files/v4.0/config/file-servers/{fileServerExtId}/mount-targets/{mountTargetExtId}/quota-policies Post operation
*/
type CreateQuotaPolicyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateQuotaPolicyApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewCreateQuotaPolicyApiResponse() *CreateQuotaPolicyApiResponse {
	p := new(CreateQuotaPolicyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.CreateQuotaPolicyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
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

/*
REST response for all response codes in API path /files/v4.0/config/file-servers/{fileServerExtId}/ransomware-configs Post operation
*/
type CreateRansomwareConfigApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateRansomwareConfigApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewCreateRansomwareConfigApiResponse() *CreateRansomwareConfigApiResponse {
	p := new(CreateRansomwareConfigApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.CreateRansomwareConfigApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreateRansomwareConfigApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreateRansomwareConfigApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreateRansomwareConfigApiResponseData()
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
REST response for all response codes in API path /files/v4.0/config/replication-policies Post operation
*/
type CreateReplicationPolicyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateReplicationPolicyApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewCreateReplicationPolicyApiResponse() *CreateReplicationPolicyApiResponse {
	p := new(CreateReplicationPolicyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.CreateReplicationPolicyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreateReplicationPolicyApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreateReplicationPolicyApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreateReplicationPolicyApiResponseData()
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
REST response for all response codes in API path /files/v4.0/config/file-servers/{fileServerExtId}/snapshot-schedules Post operation
*/
type CreateSnapshotScheduleApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateSnapshotScheduleApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewCreateSnapshotScheduleApiResponse() *CreateSnapshotScheduleApiResponse {
	p := new(CreateSnapshotScheduleApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.CreateSnapshotScheduleApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreateSnapshotScheduleApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreateSnapshotScheduleApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreateSnapshotScheduleApiResponseData()
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
REST response for all response codes in API path /files/v4.0/config/file-servers/{fileServerExtId}/tier-configurations Post operation
*/
type CreateTieringConfigurationApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateTieringConfigurationApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewCreateTieringConfigurationApiResponse() *CreateTieringConfigurationApiResponse {
	p := new(CreateTieringConfigurationApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.CreateTieringConfigurationApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreateTieringConfigurationApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreateTieringConfigurationApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreateTieringConfigurationApiResponseData()
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
REST response for all response codes in API path /files/v4.0/config/unified-namespaces Post operation
*/
type CreateUnifiedNamespaceApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateUnifiedNamespaceApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewCreateUnifiedNamespaceApiResponse() *CreateUnifiedNamespaceApiResponse {
	p := new(CreateUnifiedNamespaceApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.CreateUnifiedNamespaceApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreateUnifiedNamespaceApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreateUnifiedNamespaceApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreateUnifiedNamespaceApiResponseData()
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
REST response for all response codes in API path /files/v4.0/config/file-servers/{fileServerExtId}/virus-scan-policies Post operation
*/
type CreateVirusScanPolicyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateVirusScanPolicyApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewCreateVirusScanPolicyApiResponse() *CreateVirusScanPolicyApiResponse {
	p := new(CreateVirusScanPolicyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.CreateVirusScanPolicyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
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

/*
User credential.
*/
type Credential struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Password of the user.
	*/
	Password *string `json:"password,omitempty"`
	/*
	  Name of the user.
	*/
	Username *string `json:"username,omitempty"`
}

func NewCredential() *Credential {
	p := new(Credential)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.Credential"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
To repeat the schedule on a per `Day` basis.
*/
type DailySchedule struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Frequency of snapshots to be taken daily or hourly.
	*/
	Frequency *int `json:"frequency"`
}

func (p *DailySchedule) MarshalJSON() ([]byte, error) {
	type DailyScheduleProxy DailySchedule
	return json.Marshal(struct {
		*DailyScheduleProxy
		Frequency *int `json:"frequency,omitempty"`
	}{
		DailyScheduleProxy: (*DailyScheduleProxy)(p),
		Frequency:          p.Frequency,
	})
}

func NewDailySchedule() *DailySchedule {
	p := new(DailySchedule)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.DailySchedule"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Per day tiering schedule model.
*/
type DaySchedule struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Day of the week for tiering schedule. 1 starting with sunday.
	*/
	DayOfWeek *int64 `json:"dayOfWeek"`

	Schedules []ScheduleDetails `json:"schedules,omitempty"`
}

func (p *DaySchedule) MarshalJSON() ([]byte, error) {
	type DayScheduleProxy DaySchedule
	return json.Marshal(struct {
		*DayScheduleProxy
		DayOfWeek *int64 `json:"dayOfWeek,omitempty"`
	}{
		DayScheduleProxy: (*DayScheduleProxy)(p),
		DayOfWeek:        p.DayOfWeek,
	})
}

func NewDaySchedule() *DaySchedule {
	p := new(DaySchedule)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.DaySchedule"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /files/v4.0/config/file-servers/{fileServerExtId}/anti-virus-servers/{extId} Delete operation
*/
type DeleteAntivirusServerApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteAntivirusServerApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDeleteAntivirusServerApiResponse() *DeleteAntivirusServerApiResponse {
	p := new(DeleteAntivirusServerApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.DeleteAntivirusServerApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
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

/*
REST response for all response codes in API path /files/v4.0/config/file-servers/{fileServerExtId}/infected-files/{extId} Delete operation
*/
type DeleteInfectedFileApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteInfectedFileApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDeleteInfectedFileApiResponse() *DeleteInfectedFileApiResponse {
	p := new(DeleteInfectedFileApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.DeleteInfectedFileApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteInfectedFileApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteInfectedFileApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteInfectedFileApiResponseData()
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
REST response for all response codes in API path /files/v4.0/config/file-servers/{fileServerExtId}/mount-targets/{extId} Delete operation
*/
type DeleteMountTargetApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteMountTargetApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDeleteMountTargetApiResponse() *DeleteMountTargetApiResponse {
	p := new(DeleteMountTargetApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.DeleteMountTargetApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
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

/*
REST response for all response codes in API path /files/v4.0/config/file-servers/{fileServerExtId}/mount-targets/{mountTargetExtId}/snapshots/{extId} Delete operation
*/
type DeleteMountTargetSnapshotApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteMountTargetSnapshotApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDeleteMountTargetSnapshotApiResponse() *DeleteMountTargetSnapshotApiResponse {
	p := new(DeleteMountTargetSnapshotApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.DeleteMountTargetSnapshotApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
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

/*
REST response for all response codes in API path /files/v4.0/config/file-servers/{fileServerExtId}/notification-policies/{extId} Delete operation
*/
type DeleteNotificationPolicyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteNotificationPolicyApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDeleteNotificationPolicyApiResponse() *DeleteNotificationPolicyApiResponse {
	p := new(DeleteNotificationPolicyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.DeleteNotificationPolicyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteNotificationPolicyApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteNotificationPolicyApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteNotificationPolicyApiResponseData()
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
REST response for all response codes in API path /files/v4.0/config/file-servers/{fileServerExtId}/partner-servers/{extId} Delete operation
*/
type DeletePartnerServerApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeletePartnerServerApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDeletePartnerServerApiResponse() *DeletePartnerServerApiResponse {
	p := new(DeletePartnerServerApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.DeletePartnerServerApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeletePartnerServerApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeletePartnerServerApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeletePartnerServerApiResponseData()
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
REST response for all response codes in API path /files/v4.0/config/file-servers/{fileServerExtId}/mount-targets/{mountTargetExtId}/quota-policies/{extId} Delete operation
*/
type DeleteQuotaPolicyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteQuotaPolicyApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDeleteQuotaPolicyApiResponse() *DeleteQuotaPolicyApiResponse {
	p := new(DeleteQuotaPolicyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.DeleteQuotaPolicyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
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

/*
REST response for all response codes in API path /files/v4.0/config/file-servers/{fileServerExtId}/ransomware-configs/{extId} Delete operation
*/
type DeleteRansomwareConfigApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteRansomwareConfigApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDeleteRansomwareConfigApiResponse() *DeleteRansomwareConfigApiResponse {
	p := new(DeleteRansomwareConfigApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.DeleteRansomwareConfigApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteRansomwareConfigApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteRansomwareConfigApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteRansomwareConfigApiResponseData()
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
REST response for all response codes in API path /files/v4.0/config/file-servers/{fileServerExtId}/recommendations/{extId} Delete operation
*/
type DeleteRecommendationApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteRecommendationApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDeleteRecommendationApiResponse() *DeleteRecommendationApiResponse {
	p := new(DeleteRecommendationApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.DeleteRecommendationApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteRecommendationApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteRecommendationApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteRecommendationApiResponseData()
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
REST response for all response codes in API path /files/v4.0/config/replication-policies/{extId} Delete operation
*/
type DeleteReplicationPolicyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteReplicationPolicyApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDeleteReplicationPolicyApiResponse() *DeleteReplicationPolicyApiResponse {
	p := new(DeleteReplicationPolicyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.DeleteReplicationPolicyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteReplicationPolicyApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteReplicationPolicyApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteReplicationPolicyApiResponseData()
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
REST response for all response codes in API path /files/v4.0/config/file-servers/{fileServerExtId}/snapshot-schedules/{extId} Delete operation
*/
type DeleteSnapshotScheduleApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteSnapshotScheduleApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDeleteSnapshotScheduleApiResponse() *DeleteSnapshotScheduleApiResponse {
	p := new(DeleteSnapshotScheduleApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.DeleteSnapshotScheduleApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteSnapshotScheduleApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteSnapshotScheduleApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteSnapshotScheduleApiResponseData()
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
REST response for all response codes in API path /files/v4.0/config/file-servers/{fileServerExtId}/tier-configurations/{extId} Delete operation
*/
type DeleteTieringConfigurationApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteTieringConfigurationApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDeleteTieringConfigurationApiResponse() *DeleteTieringConfigurationApiResponse {
	p := new(DeleteTieringConfigurationApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.DeleteTieringConfigurationApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteTieringConfigurationApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteTieringConfigurationApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteTieringConfigurationApiResponseData()
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
REST response for all response codes in API path /files/v4.0/config/unified-namespaces/{extId} Delete operation
*/
type DeleteUnifiedNamespaceApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteUnifiedNamespaceApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDeleteUnifiedNamespaceApiResponse() *DeleteUnifiedNamespaceApiResponse {
	p := new(DeleteUnifiedNamespaceApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.DeleteUnifiedNamespaceApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteUnifiedNamespaceApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteUnifiedNamespaceApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteUnifiedNamespaceApiResponseData()
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
REST response for all response codes in API path /files/v4.0/config/file-servers/{fileServerExtId}/virus-scan-policies/{extId} Delete operation
*/
type DeleteVirusScanPolicyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteVirusScanPolicyApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDeleteVirusScanPolicyApiResponse() *DeleteVirusScanPolicyApiResponse {
	p := new(DeleteVirusScanPolicyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.DeleteVirusScanPolicyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteVirusScanPolicyApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteVirusScanPolicyApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteVirusScanPolicyApiResponseData()
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
A model for configuring the name and directory services of the file server.
*/
type DirectoryService struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	AdDomain *ActiveDirectoryConfig `json:"adDomain,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`

	LdapDomain *LdapConfig `json:"ldapDomain,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`

	LocalDomain *ProtocolType `json:"localDomain,omitempty"`
	/*
	  NFS V4 domain associated with the file server. Ex. child4.afs.minerva.com
	*/
	NfsV4Domain *string `json:"nfsV4Domain,omitempty"`

	NfsVersion *NfsVersion `json:"nfsVersion,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewDirectoryService() *DirectoryService {
	p := new(DirectoryService)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.DirectoryService"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /files/v4.0/config/file-servers/{fileServerExtId}/mount-targets/{extId}/$actions/disable-worm-legal-hold Post operation
*/
type DisableWormLegalHoldApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDisableWormLegalHoldApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDisableWormLegalHoldApiResponse() *DisableWormLegalHoldApiResponse {
	p := new(DisableWormLegalHoldApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.DisableWormLegalHoldApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DisableWormLegalHoldApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DisableWormLegalHoldApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDisableWormLegalHoldApiResponseData()
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
Displays the DNS action type. This is applicable to revise DNS API only.
*/
type DnsActionType int

const (
	DNSACTIONTYPE_UNKNOWN  DnsActionType = 0
	DNSACTIONTYPE_REDACTED DnsActionType = 1
	DNSACTIONTYPE_ADD      DnsActionType = 2
	DNSACTIONTYPE_REMOVE   DnsActionType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *DnsActionType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ADD",
		"REMOVE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e DnsActionType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ADD",
		"REMOVE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *DnsActionType) index(name string) DnsActionType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ADD",
		"REMOVE",
	}
	for idx := range names {
		if names[idx] == name {
			return DnsActionType(idx)
		}
	}
	return DNSACTIONTYPE_UNKNOWN
}

func (e *DnsActionType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for DnsActionType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *DnsActionType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e DnsActionType) Ref() *DnsActionType {
	return &e
}

/*
DNS detail for the file server.
*/
type DnsRecord struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`

	HostAddress *import1.IPAddressOrFQDN `json:"hostAddress,omitempty"`
	/*
	  This field indicates if the DNS entries are verified on the DNS server or not. This is a read-only field.
	*/
	IsVerified *bool `json:"isVerified,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`

	PtrRecord *import1.IPAddressOrFQDN `json:"ptrRecord,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewDnsRecord() *DnsRecord {
	p := new(DnsRecord)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.DnsRecord"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
DNS detail for the file server.
*/
type DnsRecordSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Action *DnsActionType `json:"action,omitempty"`

	Credential *Credential `json:"credential,omitempty"`
	/*
	  Preferred name server for a file server.
	*/
	PreferredNameServer *string `json:"preferredNameServer,omitempty"`
}

func NewDnsRecordSpec() *DnsRecordSpec {
	p := new(DnsRecordSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.DnsRecordSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /files/v4.0/config/file-servers/{fileServerExtId}/user-mappings Get operation
*/
type DownloadUserMappingsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDownloadUserMappingsApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDownloadUserMappingsApiResponse() *DownloadUserMappingsApiResponse {
	p := new(DownloadUserMappingsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.DownloadUserMappingsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DownloadUserMappingsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DownloadUserMappingsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDownloadUserMappingsApiResponseData()
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
Quota policy email configuration.
*/
type EmailConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Email content.
	*/
	Content *string `json:"content"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Email subject.
	*/
	Subject *string `json:"subject"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *EmailConfig) MarshalJSON() ([]byte, error) {
	type EmailConfigProxy EmailConfig
	return json.Marshal(struct {
		*EmailConfigProxy
		Content *string `json:"content,omitempty"`
		Subject *string `json:"subject,omitempty"`
	}{
		EmailConfigProxy: (*EmailConfigProxy)(p),
		Content:          p.Content,
		Subject:          p.Subject,
	})
}

func NewEmailConfig() *EmailConfig {
	p := new(EmailConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.EmailConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /files/v4.0/config/file-servers/{fileServerExtId}/mount-targets/{extId}/$actions/enable-worm-compliance Post operation
*/
type EnableWormComplianceApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfEnableWormComplianceApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewEnableWormComplianceApiResponse() *EnableWormComplianceApiResponse {
	p := new(EnableWormComplianceApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.EnableWormComplianceApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *EnableWormComplianceApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *EnableWormComplianceApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfEnableWormComplianceApiResponseData()
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
REST response for all response codes in API path /files/v4.0/config/file-servers/{fileServerExtId}/mount-targets/{extId}/$actions/enable-worm-legal-hold Post operation
*/
type EnableWormLegalHoldApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfEnableWormLegalHoldApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewEnableWormLegalHoldApiResponse() *EnableWormLegalHoldApiResponse {
	p := new(EnableWormLegalHoldApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.EnableWormLegalHoldApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *EnableWormLegalHoldApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *EnableWormLegalHoldApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfEnableWormLegalHoldApiResponseData()
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
Enforcement types for the quota policy.
*/
type EnforcementType int

const (
	ENFORCEMENTTYPE_UNKNOWN  EnforcementType = 0
	ENFORCEMENTTYPE_REDACTED EnforcementType = 1
	ENFORCEMENTTYPE_SOFT     EnforcementType = 2
	ENFORCEMENTTYPE_HARD     EnforcementType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *EnforcementType) name(index int) string {
	names := [...]string{
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

// Returns the name of the enum
func (e EnforcementType) GetName() string {
	index := int(e)
	names := [...]string{
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

// Returns the enum type given a string value
func (e *EnforcementType) index(name string) EnforcementType {
	names := [...]string{
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

/*
The mode of the file blocking on the file extension list.
*/
type FileBlockingMode int

const (
	FILEBLOCKINGMODE_UNKNOWN    FileBlockingMode = 0
	FILEBLOCKINGMODE_REDACTED   FileBlockingMode = 1
	FILEBLOCKINGMODE_ALLOW_LIST FileBlockingMode = 2
	FILEBLOCKINGMODE_BLOCK_LIST FileBlockingMode = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *FileBlockingMode) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ALLOW_LIST",
		"BLOCK_LIST",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e FileBlockingMode) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ALLOW_LIST",
		"BLOCK_LIST",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *FileBlockingMode) index(name string) FileBlockingMode {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ALLOW_LIST",
		"BLOCK_LIST",
	}
	for idx := range names {
		if names[idx] == name {
			return FileBlockingMode(idx)
		}
	}
	return FILEBLOCKINGMODE_UNKNOWN
}

func (e *FileBlockingMode) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for FileBlockingMode:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *FileBlockingMode) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e FileBlockingMode) Ref() *FileBlockingMode {
	return &e
}

/*
Indicates the value of the file operations.
*/
type FileOperation int

const (
	FILEOPERATION_UNKNOWN          FileOperation = 0
	FILEOPERATION_REDACTED         FileOperation = 1
	FILEOPERATION_FILE_CREATE      FileOperation = 2
	FILEOPERATION_FILE_DELETE      FileOperation = 3
	FILEOPERATION_FILE_READ        FileOperation = 4
	FILEOPERATION_FILE_WRITE       FileOperation = 5
	FILEOPERATION_FILE_OPEN        FileOperation = 6
	FILEOPERATION_FILE_CLOSE       FileOperation = 7
	FILEOPERATION_DIRECTORY_CREATE FileOperation = 8
	FILEOPERATION_DIRECTORY_DELETE FileOperation = 9
	FILEOPERATION_RENAME           FileOperation = 10
	FILEOPERATION_SETATTR          FileOperation = 11
	FILEOPERATION_SECURITY         FileOperation = 12
	FILEOPERATION_SYMLINK_CREATE   FileOperation = 13
	FILEOPERATION_LINK_CREATE      FileOperation = 14
	FILEOPERATION_TIER             FileOperation = 15
	FILEOPERATION_RECALL           FileOperation = 16
	FILEOPERATION_INLINE_READ      FileOperation = 17
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *FileOperation) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"FILE_CREATE",
		"FILE_DELETE",
		"FILE_READ",
		"FILE_WRITE",
		"FILE_OPEN",
		"FILE_CLOSE",
		"DIRECTORY_CREATE",
		"DIRECTORY_DELETE",
		"RENAME",
		"SETATTR",
		"SECURITY",
		"SYMLINK_CREATE",
		"LINK_CREATE",
		"TIER",
		"RECALL",
		"INLINE_READ",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e FileOperation) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"FILE_CREATE",
		"FILE_DELETE",
		"FILE_READ",
		"FILE_WRITE",
		"FILE_OPEN",
		"FILE_CLOSE",
		"DIRECTORY_CREATE",
		"DIRECTORY_DELETE",
		"RENAME",
		"SETATTR",
		"SECURITY",
		"SYMLINK_CREATE",
		"LINK_CREATE",
		"TIER",
		"RECALL",
		"INLINE_READ",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *FileOperation) index(name string) FileOperation {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"FILE_CREATE",
		"FILE_DELETE",
		"FILE_READ",
		"FILE_WRITE",
		"FILE_OPEN",
		"FILE_CLOSE",
		"DIRECTORY_CREATE",
		"DIRECTORY_DELETE",
		"RENAME",
		"SETATTR",
		"SECURITY",
		"SYMLINK_CREATE",
		"LINK_CREATE",
		"TIER",
		"RECALL",
		"INLINE_READ",
	}
	for idx := range names {
		if names[idx] == name {
			return FileOperation(idx)
		}
	}
	return FILEOPERATION_UNKNOWN
}

func (e *FileOperation) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for FileOperation:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *FileOperation) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e FileOperation) Ref() *FileOperation {
	return &e
}

/*
A model that represents file server resources.
*/
type FileServer struct {
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
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  File server name.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewFileServer() *FileServer {
	p := new(FileServer)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.FileServer"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
File server data synchronization paths
*/
type FileServerSyncPathsSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The external identifier of the file server.
	*/
	FileServerReference *string `json:"fileServerReference,omitempty"`
	/*
	  Data synchronization path entities
	*/
	SyncPaths []string `json:"syncPaths"`
}

func (p *FileServerSyncPathsSpec) MarshalJSON() ([]byte, error) {
	type FileServerSyncPathsSpecProxy FileServerSyncPathsSpec
	return json.Marshal(struct {
		*FileServerSyncPathsSpecProxy
		SyncPaths []string `json:"syncPaths,omitempty"`
	}{
		FileServerSyncPathsSpecProxy: (*FileServerSyncPathsSpecProxy)(p),
		SyncPaths:                    p.SyncPaths,
	})
}

func NewFileServerSyncPathsSpec() *FileServerSyncPathsSpec {
	p := new(FileServerSyncPathsSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.FileServerSyncPathsSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
File server data synchronization paths validation response. Response includes the invalid synchronization paths and the error messages.
*/
type FileServerSyncPathsValidationResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Data synchronization path validation API error message.
	*/
	ErrorMessage *string `json:"errorMessage,omitempty"`
	/*
	  The external identifier of the file server.
	*/
	FileServerExtId *string `json:"fileServerExtId,omitempty"`
	/*
	  Data synchronization path entities validation response.
	*/
	SyncPathEntitiesResponse []SynchronizationPathEntityResponse `json:"syncPathEntitiesResponse,omitempty"`
}

func NewFileServerSyncPathsValidationResponse() *FileServerSyncPathsValidationResponse {
	p := new(FileServerSyncPathsValidationResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.FileServerSyncPathsValidationResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Authentication type of the mount target.
*/
type FilesAuthenticationType int

const (
	FILESAUTHENTICATIONTYPE_UNKNOWN    FilesAuthenticationType = 0
	FILESAUTHENTICATIONTYPE_REDACTED   FilesAuthenticationType = 1
	FILESAUTHENTICATIONTYPE_KERBEROS5  FilesAuthenticationType = 2
	FILESAUTHENTICATIONTYPE_NONE       FilesAuthenticationType = 3
	FILESAUTHENTICATIONTYPE_SYSTEM     FilesAuthenticationType = 4
	FILESAUTHENTICATIONTYPE_KERBEROS5P FilesAuthenticationType = 5
	FILESAUTHENTICATIONTYPE_KERBEROS5I FilesAuthenticationType = 6
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *FilesAuthenticationType) name(index int) string {
	names := [...]string{
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

// Returns the name of the enum
func (e FilesAuthenticationType) GetName() string {
	index := int(e)
	names := [...]string{
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

// Returns the enum type given a string value
func (e *FilesAuthenticationType) index(name string) FilesAuthenticationType {
	names := [...]string{
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

/*
Request body specification for tier APIs.
*/
type FilesSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Path of the files to be tiered/recalled relative to root Mount target.
	*/
	FilePaths []string `json:"filePaths,omitempty"`
	/*
	  List of inode numbers strings. Inode number should be in the format Inode/Inode:GenId.
	*/
	INodeNumbers []string `json:"iNodeNumbers,omitempty"`
	/*
	  Boolean field to indicate if all files belong to the same VG.
	*/
	IsSingleDirTree *bool `json:"isSingleDirTree,omitempty"`
	/*
	  Mount target external identifier for files to tier/recall.
	*/
	MountTargetExtId *string `json:"mountTargetExtId"`
}

func (p *FilesSpec) MarshalJSON() ([]byte, error) {
	type FilesSpecProxy FilesSpec
	return json.Marshal(struct {
		*FilesSpecProxy
		MountTargetExtId *string `json:"mountTargetExtId,omitempty"`
	}{
		FilesSpecProxy:   (*FilesSpecProxy)(p),
		MountTargetExtId: p.MountTargetExtId,
	})
}

func NewFilesSpec() *FilesSpec {
	p := new(FilesSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.FilesSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	p.IsSingleDirTree = new(bool)
	*p.IsSingleDirTree = false

	return p
}

/*
REST response for all response codes in API path /files/v4.0/config/file-servers/{fileServerExtId}/infected-files/{extId}/$actions/fix Post operation
*/
type FixInfectedFileApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfFixInfectedFileApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewFixInfectedFileApiResponse() *FixInfectedFileApiResponse {
	p := new(FixInfectedFileApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.FixInfectedFileApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *FixInfectedFileApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *FixInfectedFileApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfFixInfectedFileApiResponseData()
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
REST response for all response codes in API path /files/v4.0/config/file-servers/{fileServerExtId}/anti-virus-servers/{extId} Get operation
*/
type GetAntivirusServerApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetAntivirusServerApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetAntivirusServerApiResponse() *GetAntivirusServerApiResponse {
	p := new(GetAntivirusServerApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.GetAntivirusServerApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetAntivirusServerApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetAntivirusServerApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetAntivirusServerApiResponseData()
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
REST response for all response codes in API path /files/v4.0/config/file-servers/{fileServerExtId}/email-config Get operation
*/
type GetEmailConfigApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetEmailConfigApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetEmailConfigApiResponse() *GetEmailConfigApiResponse {
	p := new(GetEmailConfigApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.GetEmailConfigApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetEmailConfigApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetEmailConfigApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetEmailConfigApiResponseData()
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
REST response for all response codes in API path /files/v4.0/config/file-servers/{extId} Get operation
*/
type GetFileServerApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetFileServerApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetFileServerApiResponse() *GetFileServerApiResponse {
	p := new(GetFileServerApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.GetFileServerApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetFileServerApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetFileServerApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetFileServerApiResponseData()
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
REST response for all response codes in API path /files/v4.0/config/file-servers/{fileServerExtId}/infected-files/{extId} Get operation
*/
type GetInfectedFileApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetInfectedFileApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetInfectedFileApiResponse() *GetInfectedFileApiResponse {
	p := new(GetInfectedFileApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.GetInfectedFileApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetInfectedFileApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetInfectedFileApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetInfectedFileApiResponseData()
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
REST response for all response codes in API path /files/v4.0/config/file-servers/{fileServerExtId}/mount-targets/{extId} Get operation
*/
type GetMountTargetApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetMountTargetApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetMountTargetApiResponse() *GetMountTargetApiResponse {
	p := new(GetMountTargetApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.GetMountTargetApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetMountTargetApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetMountTargetApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetMountTargetApiResponseData()
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
REST response for all response codes in API path /files/v4.0/config/file-servers/{fileServerExtId}/mount-targets/{mountTargetExtId}/snapshots/{extId} Get operation
*/
type GetMountTargetSnapshotApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetMountTargetSnapshotApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetMountTargetSnapshotApiResponse() *GetMountTargetSnapshotApiResponse {
	p := new(GetMountTargetSnapshotApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.GetMountTargetSnapshotApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetMountTargetSnapshotApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetMountTargetSnapshotApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetMountTargetSnapshotApiResponseData()
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
REST response for all response codes in API path /files/v4.0/config/file-servers/{fileServerExtId}/notification-policies/{extId} Get operation
*/
type GetNotificationPolicyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetNotificationPolicyApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetNotificationPolicyApiResponse() *GetNotificationPolicyApiResponse {
	p := new(GetNotificationPolicyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.GetNotificationPolicyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetNotificationPolicyApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetNotificationPolicyApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetNotificationPolicyApiResponseData()
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
REST response for all response codes in API path /files/v4.0/config/file-servers/{fileServerExtId}/object-store-profiles/{extId} Get operation
*/
type GetObjectStoreProfileApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetObjectStoreProfileApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetObjectStoreProfileApiResponse() *GetObjectStoreProfileApiResponse {
	p := new(GetObjectStoreProfileApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.GetObjectStoreProfileApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetObjectStoreProfileApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetObjectStoreProfileApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetObjectStoreProfileApiResponseData()
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
REST response for all response codes in API path /files/v4.0/config/file-servers/{fileServerExtId}/partner-servers/{extId} Get operation
*/
type GetPartnerServerApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetPartnerServerApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetPartnerServerApiResponse() *GetPartnerServerApiResponse {
	p := new(GetPartnerServerApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.GetPartnerServerApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetPartnerServerApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetPartnerServerApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetPartnerServerApiResponseData()
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
REST response for all response codes in API path /files/v4.0/config/file-servers/{fileServerExtId}/mount-targets/{mountTargetExtId}/quota-policies/{extId} Get operation
*/
type GetQuotaPolicyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetQuotaPolicyApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetQuotaPolicyApiResponse() *GetQuotaPolicyApiResponse {
	p := new(GetQuotaPolicyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.GetQuotaPolicyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetQuotaPolicyApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetQuotaPolicyApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetQuotaPolicyApiResponseData()
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
REST response for all response codes in API path /files/v4.0/config/file-servers/{fileServerExtId}/ransomware-configs/{extId} Get operation
*/
type GetRansomwareConfigApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetRansomwareConfigApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetRansomwareConfigApiResponse() *GetRansomwareConfigApiResponse {
	p := new(GetRansomwareConfigApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.GetRansomwareConfigApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetRansomwareConfigApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetRansomwareConfigApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetRansomwareConfigApiResponseData()
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
REST response for all response codes in API path /files/v4.0/config/file-servers/{fileServerExtId}/recommendations/{extId} Get operation
*/
type GetRecommendationApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetRecommendationApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetRecommendationApiResponse() *GetRecommendationApiResponse {
	p := new(GetRecommendationApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.GetRecommendationApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetRecommendationApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetRecommendationApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetRecommendationApiResponseData()
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
REST response for all response codes in API path /files/v4.0/config/replication-policies/{extId} Get operation
*/
type GetReplicationPolicyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetReplicationPolicyApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetReplicationPolicyApiResponse() *GetReplicationPolicyApiResponse {
	p := new(GetReplicationPolicyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.GetReplicationPolicyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetReplicationPolicyApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetReplicationPolicyApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetReplicationPolicyApiResponseData()
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
REST response for all response codes in API path /files/v4.0/config/file-servers/{fileServerExtId}/mount-targets/{mountTargetExtId}/snapshot-changed-contents/{extId} Get operation
*/
type GetSnapshotChangedContentApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetSnapshotChangedContentApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetSnapshotChangedContentApiResponse() *GetSnapshotChangedContentApiResponse {
	p := new(GetSnapshotChangedContentApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.GetSnapshotChangedContentApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetSnapshotChangedContentApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetSnapshotChangedContentApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetSnapshotChangedContentApiResponseData()
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
REST response for all response codes in API path /files/v4.0/config/file-servers/{fileServerExtId}/snapshot-schedules/{extId} Get operation
*/
type GetSnapshotScheduleApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetSnapshotScheduleApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetSnapshotScheduleApiResponse() *GetSnapshotScheduleApiResponse {
	p := new(GetSnapshotScheduleApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.GetSnapshotScheduleApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetSnapshotScheduleApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetSnapshotScheduleApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetSnapshotScheduleApiResponseData()
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
REST response for all response codes in API path /files/v4.0/config/file-servers/{fileServerExtId}/tier-configurations/{extId} Get operation
*/
type GetTieringConfigurationApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetTieringConfigurationApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetTieringConfigurationApiResponse() *GetTieringConfigurationApiResponse {
	p := new(GetTieringConfigurationApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.GetTieringConfigurationApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetTieringConfigurationApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetTieringConfigurationApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetTieringConfigurationApiResponseData()
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
REST response for all response codes in API path /files/v4.0/config/unified-namespaces/{extId} Get operation
*/
type GetUnifiedNamespaceApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetUnifiedNamespaceApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetUnifiedNamespaceApiResponse() *GetUnifiedNamespaceApiResponse {
	p := new(GetUnifiedNamespaceApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.GetUnifiedNamespaceApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetUnifiedNamespaceApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetUnifiedNamespaceApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetUnifiedNamespaceApiResponseData()
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
REST response for all response codes in API path /files/v4.0/config/file-servers/{fileServerExtId}/replication-policies/{replicationPolicyExtId}/vdi-user-sessions/{extId} Get operation
*/
type GetVdiUserSessionApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetVdiUserSessionApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetVdiUserSessionApiResponse() *GetVdiUserSessionApiResponse {
	p := new(GetVdiUserSessionApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.GetVdiUserSessionApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetVdiUserSessionApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetVdiUserSessionApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetVdiUserSessionApiResponseData()
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
REST response for all response codes in API path /files/v4.0/config/file-servers/{fileServerExtId}/virus-scan-policies/{extId} Get operation
*/
type GetVirusScanPolicyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetVirusScanPolicyApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetVirusScanPolicyApiResponse() *GetVirusScanPolicyApiResponse {
	p := new(GetVirusScanPolicyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.GetVirusScanPolicyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetVirusScanPolicyApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetVirusScanPolicyApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetVirusScanPolicyApiResponseData()
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
Infected files model
*/
type InfectedFile struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Quarantine status of infected files.
	*/
	IsQuarantined *bool `json:"isQuarantined,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  The external identifier of the mount target where infected file belongs to. Example:9c1e537d-6777-4c22-5d41-ddd0c3337aa9.
	*/
	MountTargetExtId *string `json:"mountTargetExtId,omitempty"`
	/*
	  Partner server details.
	*/
	PartnerServer *string `json:"partnerServer,omitempty"`
	/*
	  Path of infected file
	*/
	Path *string `json:"path,omitempty"`
	/*
	  Scan timestamp of infected file. This is a read-only field.
	*/
	ScanTime *time.Time `json:"scanTime,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  Threat details for infected file
	*/
	ThreatDescription *string `json:"threatDescription,omitempty"`
}

func NewInfectedFile() *InfectedFile {
	p := new(InfectedFile)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.InfectedFile"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Security action.
*/
type InfectedFileActionType int

const (
	INFECTEDFILEACTIONTYPE_UNKNOWN      InfectedFileActionType = 0
	INFECTEDFILEACTIONTYPE_REDACTED     InfectedFileActionType = 1
	INFECTEDFILEACTIONTYPE_RESCAN       InfectedFileActionType = 2
	INFECTEDFILEACTIONTYPE_RESET        InfectedFileActionType = 3
	INFECTEDFILEACTIONTYPE_QUARANTINE   InfectedFileActionType = 4
	INFECTEDFILEACTIONTYPE_UNQUARANTINE InfectedFileActionType = 5
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *InfectedFileActionType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"RESCAN",
		"RESET",
		"QUARANTINE",
		"UNQUARANTINE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e InfectedFileActionType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"RESCAN",
		"RESET",
		"QUARANTINE",
		"UNQUARANTINE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *InfectedFileActionType) index(name string) InfectedFileActionType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"RESCAN",
		"RESET",
		"QUARANTINE",
		"UNQUARANTINE",
	}
	for idx := range names {
		if names[idx] == name {
			return InfectedFileActionType(idx)
		}
	}
	return INFECTEDFILEACTIONTYPE_UNKNOWN
}

func (e *InfectedFileActionType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for InfectedFileActionType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *InfectedFileActionType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e InfectedFileActionType) Ref() *InfectedFileActionType {
	return &e
}

/*
Security action to perform on infected files.
*/
type InfectedFileFixSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Action *InfectedFileActionType `json:"action"`
}

func (p *InfectedFileFixSpec) MarshalJSON() ([]byte, error) {
	type InfectedFileFixSpecProxy InfectedFileFixSpec
	return json.Marshal(struct {
		*InfectedFileFixSpecProxy
		Action *InfectedFileActionType `json:"action,omitempty"`
	}{
		InfectedFileFixSpecProxy: (*InfectedFileFixSpecProxy)(p),
		Action:                   p.Action,
	})
}

func NewInfectedFileFixSpec() *InfectedFileFixSpec {
	p := new(InfectedFileFixSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.InfectedFileFixSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
LDAP domain associated with the file server.
*/
type LdapConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  LDAP Base distinguished name (DN) specifies the starting point within an LDAP directory hierarchy for LDAP operations like searches and queries. Ex: DC=ldaps,DC=afs,DC=minerva,DC=com
	*/
	Basedn *string `json:"basedn,omitempty"`
	/*
	  A distinguished name that is used to authenticate and authorize an LDAP client to perform operations on an LDAP directory server. This is an optional field. Ex: CN=example-user,CN=Users,DC=example-domain,DC=com
	*/
	Binddn *string `json:"binddn,omitempty"`
	/*
	  LDAP password
	*/
	Bindpw *string `json:"bindpw,omitempty"`
	/*
	  URI of the LDAP domain. Ex. ldap://10.51.38.36
	*/
	Location *string `json:"location,omitempty"`

	ProtocolType *ProtocolType `json:"protocolType"`
	/*
	  LDAP SSH key
	*/
	TlsCacertContent *string `json:"tlsCacertContent,omitempty"`

	TlsReqCertStatus *TLSReqCertStatus `json:"tlsReqCertStatus,omitempty"`
}

func (p *LdapConfig) MarshalJSON() ([]byte, error) {
	type LdapConfigProxy LdapConfig
	return json.Marshal(struct {
		*LdapConfigProxy
		ProtocolType *ProtocolType `json:"protocolType,omitempty"`
	}{
		LdapConfigProxy: (*LdapConfigProxy)(p),
		ProtocolType:    p.ProtocolType,
	})
}

func NewLdapConfig() *LdapConfig {
	p := new(LdapConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.LdapConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /files/v4.0/config/file-servers/{fileServerExtId}/anti-virus-servers Get operation
*/
type ListAntivirusServersApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListAntivirusServersApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListAntivirusServersApiResponse() *ListAntivirusServersApiResponse {
	p := new(ListAntivirusServersApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.ListAntivirusServersApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListAntivirusServersApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListAntivirusServersApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListAntivirusServersApiResponseData()
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
REST response for all response codes in API path /files/v4.0/config/file-servers/{fileServerExtId}/dns-records Get operation
*/
type ListDnsRecordsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListDnsRecordsApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListDnsRecordsApiResponse() *ListDnsRecordsApiResponse {
	p := new(ListDnsRecordsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.ListDnsRecordsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListDnsRecordsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListDnsRecordsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListDnsRecordsApiResponseData()
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
REST response for all response codes in API path /files/v4.0/config/file-servers Get operation
*/
type ListFileServersApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListFileServersApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListFileServersApiResponse() *ListFileServersApiResponse {
	p := new(ListFileServersApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.ListFileServersApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListFileServersApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListFileServersApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListFileServersApiResponseData()
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
REST response for all response codes in API path /files/v4.0/config/file-servers/{fileServerExtId}/infected-files Get operation
*/
type ListInfectedFilesApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListInfectedFilesApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListInfectedFilesApiResponse() *ListInfectedFilesApiResponse {
	p := new(ListInfectedFilesApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.ListInfectedFilesApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListInfectedFilesApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListInfectedFilesApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListInfectedFilesApiResponseData()
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
REST response for all response codes in API path /files/v4.0/config/file-servers/{fileServerExtId}/mount-targets/{mountTargetExtId}/snapshots Get operation
*/
type ListMountTargetSnapshotsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListMountTargetSnapshotsApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListMountTargetSnapshotsApiResponse() *ListMountTargetSnapshotsApiResponse {
	p := new(ListMountTargetSnapshotsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.ListMountTargetSnapshotsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListMountTargetSnapshotsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListMountTargetSnapshotsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListMountTargetSnapshotsApiResponseData()
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
REST response for all response codes in API path /files/v4.0/config/file-servers/{fileServerExtId}/mount-targets Get operation
*/
type ListMountTargetsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListMountTargetsApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListMountTargetsApiResponse() *ListMountTargetsApiResponse {
	p := new(ListMountTargetsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.ListMountTargetsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListMountTargetsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListMountTargetsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListMountTargetsApiResponseData()
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
REST response for all response codes in API path /files/v4.0/config/file-servers/{fileServerExtId}/notification-policies Get operation
*/
type ListNotificationPoliciesApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListNotificationPoliciesApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListNotificationPoliciesApiResponse() *ListNotificationPoliciesApiResponse {
	p := new(ListNotificationPoliciesApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.ListNotificationPoliciesApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListNotificationPoliciesApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListNotificationPoliciesApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListNotificationPoliciesApiResponseData()
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
REST response for all response codes in API path /files/v4.0/config/file-servers/{fileServerExtId}/object-store-profiles Get operation
*/
type ListObjectStoreProfilesApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListObjectStoreProfilesApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListObjectStoreProfilesApiResponse() *ListObjectStoreProfilesApiResponse {
	p := new(ListObjectStoreProfilesApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.ListObjectStoreProfilesApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListObjectStoreProfilesApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListObjectStoreProfilesApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListObjectStoreProfilesApiResponseData()
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
REST response for all response codes in API path /files/v4.0/config/file-servers/{fileServerExtId}/partner-servers Get operation
*/
type ListPartnerServersApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListPartnerServersApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListPartnerServersApiResponse() *ListPartnerServersApiResponse {
	p := new(ListPartnerServersApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.ListPartnerServersApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListPartnerServersApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListPartnerServersApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListPartnerServersApiResponseData()
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
REST response for all response codes in API path /files/v4.0/config/file-servers/{fileServerExtId}/mount-targets/{mountTargetExtId}/quota-policies Get operation
*/
type ListQuotaPoliciesApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListQuotaPoliciesApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListQuotaPoliciesApiResponse() *ListQuotaPoliciesApiResponse {
	p := new(ListQuotaPoliciesApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.ListQuotaPoliciesApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListQuotaPoliciesApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListQuotaPoliciesApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListQuotaPoliciesApiResponseData()
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
REST response for all response codes in API path /files/v4.0/config/file-servers/{fileServerExtId}/ransomware-configs Get operation
*/
type ListRansomwareConfigsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListRansomwareConfigsApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListRansomwareConfigsApiResponse() *ListRansomwareConfigsApiResponse {
	p := new(ListRansomwareConfigsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.ListRansomwareConfigsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListRansomwareConfigsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListRansomwareConfigsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListRansomwareConfigsApiResponseData()
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
REST response for all response codes in API path /files/v4.0/config/file-servers/{fileServerExtId}/recommendations Get operation
*/
type ListRecommendationsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListRecommendationsApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListRecommendationsApiResponse() *ListRecommendationsApiResponse {
	p := new(ListRecommendationsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.ListRecommendationsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListRecommendationsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListRecommendationsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListRecommendationsApiResponseData()
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
REST response for all response codes in API path /files/v4.0/config/replication-policies Get operation
*/
type ListReplicationPoliciesApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListReplicationPoliciesApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListReplicationPoliciesApiResponse() *ListReplicationPoliciesApiResponse {
	p := new(ListReplicationPoliciesApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.ListReplicationPoliciesApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListReplicationPoliciesApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListReplicationPoliciesApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListReplicationPoliciesApiResponseData()
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
REST response for all response codes in API path /files/v4.0/config/file-servers/{fileServerExtId}/mount-targets/{mountTargetExtId}/snapshot-changed-contents Get operation
*/
type ListSnapshotChangedContentsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListSnapshotChangedContentsApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListSnapshotChangedContentsApiResponse() *ListSnapshotChangedContentsApiResponse {
	p := new(ListSnapshotChangedContentsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.ListSnapshotChangedContentsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListSnapshotChangedContentsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListSnapshotChangedContentsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListSnapshotChangedContentsApiResponseData()
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
REST response for all response codes in API path /files/v4.0/config/file-servers/{fileServerExtId}/snapshot-schedules Get operation
*/
type ListSnapshotSchedulesApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListSnapshotSchedulesApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListSnapshotSchedulesApiResponse() *ListSnapshotSchedulesApiResponse {
	p := new(ListSnapshotSchedulesApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.ListSnapshotSchedulesApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListSnapshotSchedulesApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListSnapshotSchedulesApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListSnapshotSchedulesApiResponseData()
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
REST response for all response codes in API path /files/v4.0/config/file-servers/{fileServerExtId}/tier-configurations Get operation
*/
type ListTieringConfigurationsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListTieringConfigurationsApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListTieringConfigurationsApiResponse() *ListTieringConfigurationsApiResponse {
	p := new(ListTieringConfigurationsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.ListTieringConfigurationsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListTieringConfigurationsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListTieringConfigurationsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListTieringConfigurationsApiResponseData()
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
REST response for all response codes in API path /files/v4.0/config/unified-namespaces Get operation
*/
type ListUnifiedNamespacesApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListUnifiedNamespacesApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListUnifiedNamespacesApiResponse() *ListUnifiedNamespacesApiResponse {
	p := new(ListUnifiedNamespacesApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.ListUnifiedNamespacesApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListUnifiedNamespacesApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListUnifiedNamespacesApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListUnifiedNamespacesApiResponseData()
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
REST response for all response codes in API path /files/v4.0/config/file-servers/{fileServerExtId}/replication-policies/{replicationPolicyExtId}/vdi-user-sessions Get operation
*/
type ListVdiUserSessionsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListVdiUserSessionsApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListVdiUserSessionsApiResponse() *ListVdiUserSessionsApiResponse {
	p := new(ListVdiUserSessionsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.ListVdiUserSessionsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListVdiUserSessionsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListVdiUserSessionsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListVdiUserSessionsApiResponseData()
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
REST response for all response codes in API path /files/v4.0/config/file-servers/{fileServerExtId}/virus-scan-policies Get operation
*/
type ListVirusScanPoliciesApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListVirusScanPoliciesApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListVirusScanPoliciesApiResponse() *ListVirusScanPoliciesApiResponse {
	p := new(ListVirusScanPoliciesApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.ListVirusScanPoliciesApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListVirusScanPoliciesApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListVirusScanPoliciesApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListVirusScanPoliciesApiResponseData()
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
User login status for VDI synchronization.
*/
type LoginStatus int

const (
	LOGINSTATUS_UNKNOWN   LoginStatus = 0
	LOGINSTATUS_REDACTED  LoginStatus = 1
	LOGINSTATUS_FAILED    LoginStatus = 2
	LOGINSTATUS_SUCCEEDED LoginStatus = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *LoginStatus) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"FAILED",
		"SUCCEEDED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e LoginStatus) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"FAILED",
		"SUCCEEDED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *LoginStatus) index(name string) LoginStatus {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"FAILED",
		"SUCCEEDED",
	}
	for idx := range names {
		if names[idx] == name {
			return LoginStatus(idx)
		}
	}
	return LOGINSTATUS_UNKNOWN
}

func (e *LoginStatus) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for LoginStatus:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *LoginStatus) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e LoginStatus) Ref() *LoginStatus {
	return &e
}

/*
Migration plan details
*/
type MigrationPlan struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Description of migration plan
	*/
	Description *string `json:"description,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Migration plan name
	*/
	Name *string `json:"name"`
	/*
	  The extId of migration source to use for migration.
	*/
	SourceExtId *string `json:"sourceExtId"`
	/*
	  List of sub-plan extIds. This is a read-only field.
	*/
	SubPlanExtIds []string `json:"subPlanExtIds,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *MigrationPlan) MarshalJSON() ([]byte, error) {
	type MigrationPlanProxy MigrationPlan
	return json.Marshal(struct {
		*MigrationPlanProxy
		Name        *string `json:"name,omitempty"`
		SourceExtId *string `json:"sourceExtId,omitempty"`
	}{
		MigrationPlanProxy: (*MigrationPlanProxy)(p),
		Name:               p.Name,
		SourceExtId:        p.SourceExtId,
	})
}

func NewMigrationPlan() *MigrationPlan {
	p := new(MigrationPlan)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.MigrationPlan"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Migration source server details to be used by migration plans.
*/
type MigrationSource struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Address *import1.IPAddressOrFQDN `json:"address"`

	Credential *Credential `json:"credential,omitempty"`
	/*
	  Description for migration source
	*/
	Description *string `json:"description,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  List of mount target extIds for migration source server. This is a read-only field.
	*/
	MountTargetExtIds []string `json:"mountTargetExtIds,omitempty"`
	/*
	  Migration source name
	*/
	Name *string `json:"name"`
	/*
	  A flag to indicate if the backend should start mount-target discovery for migration source server. Upon completion of source creation task, list of discovered mount target can be fetched using GET mount-targets API on migration-sources. Note: Credential field is required for SMB mount-targets discovery and `shouldStartDiscoveryDesc` is write only field.
	*/
	ShouldDiscoverMountTargets *bool `json:"shouldDiscoverMountTargets,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *MigrationSource) MarshalJSON() ([]byte, error) {
	type MigrationSourceProxy MigrationSource
	return json.Marshal(struct {
		*MigrationSourceProxy
		Address *import1.IPAddressOrFQDN `json:"address,omitempty"`
		Name    *string                  `json:"name,omitempty"`
	}{
		MigrationSourceProxy: (*MigrationSourceProxy)(p),
		Address:              p.Address,
		Name:                 p.Name,
	})
}

func NewMigrationSource() *MigrationSource {
	p := new(MigrationSource)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.MigrationSource"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	p.ShouldDiscoverMountTargets = new(bool)
	*p.ShouldDiscoverMountTargets = false

	return p
}

/*
Mount target for migration source.
*/
type MigrationSourceMountTarget struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Total bytes present on source path. This is a read-only field.
	*/
	Bytes *int64 `json:"bytes,omitempty"`
	/*
	  Total dirs present on source path. This is a read-only field.
	*/
	Directories *int64 `json:"directories,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Total files present on source path. This is a read-only field.
	*/
	Files *int64 `json:"files,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Name of the source mount target. Only the mount target name is stored in the database after discovery, but to scan you can specify the mount target name with path, in this case name including path is stored in DB.
	*/
	Name *string `json:"name,omitempty"`

	Protocol *MountTargetProtocolType `json:"protocol,omitempty"`

	ScanStatus *ScanStatus `json:"scanStatus,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewMigrationSourceMountTarget() *MigrationSourceMountTarget {
	p := new(MigrationSourceMountTarget)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.MigrationSourceMountTarget"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	p.Bytes = new(int64)
	*p.Bytes = 0
	p.Directories = new(int64)
	*p.Directories = 0
	p.Files = new(int64)
	*p.Files = 0

	return p
}

/*
An enum that represents the migration status
*/
type MigrationStatus int

const (
	MIGRATIONSTATUS_UNKNOWN               MigrationStatus = 0
	MIGRATIONSTATUS_REDACTED              MigrationStatus = 1
	MIGRATIONSTATUS_DEFAULT               MigrationStatus = 2
	MIGRATIONSTATUS_QUEUED                MigrationStatus = 3
	MIGRATIONSTATUS_RUNNING               MigrationStatus = 4
	MIGRATIONSTATUS_SUCCEEDED             MigrationStatus = 5
	MIGRATIONSTATUS_CANCELLED             MigrationStatus = 6
	MIGRATIONSTATUS_PAUSED                MigrationStatus = 7
	MIGRATIONSTATUS_FAILED                MigrationStatus = 8
	MIGRATIONSTATUS_COMPLETED_WITH_ERRORS MigrationStatus = 9
	MIGRATIONSTATUS_COMPLETED             MigrationStatus = 10
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *MigrationStatus) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"DEFAULT",
		"QUEUED",
		"RUNNING",
		"SUCCEEDED",
		"CANCELLED",
		"PAUSED",
		"FAILED",
		"COMPLETED_WITH_ERRORS",
		"COMPLETED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e MigrationStatus) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"DEFAULT",
		"QUEUED",
		"RUNNING",
		"SUCCEEDED",
		"CANCELLED",
		"PAUSED",
		"FAILED",
		"COMPLETED_WITH_ERRORS",
		"COMPLETED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *MigrationStatus) index(name string) MigrationStatus {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"DEFAULT",
		"QUEUED",
		"RUNNING",
		"SUCCEEDED",
		"CANCELLED",
		"PAUSED",
		"FAILED",
		"COMPLETED_WITH_ERRORS",
		"COMPLETED",
	}
	for idx := range names {
		if names[idx] == name {
			return MigrationStatus(idx)
		}
	}
	return MIGRATIONSTATUS_UNKNOWN
}

func (e *MigrationStatus) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for MigrationStatus:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *MigrationStatus) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e MigrationStatus) Ref() *MigrationStatus {
	return &e
}

/*
Migration sub-plan details.
*/
type MigrationSubPlan struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A flag to identify if its final iteration. This flag should be given as true to do the final round of migration. It is recommended to keep the sub-plan source as readonly to migrate all the content from the source to the target. We prioritize such an iteration so that the read-only time for the source can be decreased by utilising migrator resources only for the subplans in final iteration. After the completion of this iteration, if no failed files are found, we need to update the sub-plan status as COMPLETED, and no further action on this sub-plan will be allowed. If there are failed files under max-limit while migration, the subplan status will be changed to succeded and you will have numberOfFilesFailed/numberOfDirectoriesFailed/numberOfStreamsFailed on migration job. These failed files name and error can be retrieved using GET failed-files-summary API, and upon error correction, the user can start another iteration to sync failed files only the flag shouldRetryOnlyFailedFiles flag.
	*/
	IsFinalIteration *bool `json:"isFinalIteration,omitempty"`
	/*
	  Indicates the flag for sync type. If false, then only the source will be synced. The delete and rename are not honored on the target, if anything is deleted or renamed on the source. If true, then the source will be mirrored on the target, that is, every delete and rename will be honored.
	*/
	IsSyncTypeMirror *bool `json:"isSyncTypeMirror,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Migration sub-plan name
	*/
	Name *string `json:"name"`

	Protocol *MountTargetProtocolType `json:"protocol,omitempty"`
	/*
	  A flag to indicate whether unknown SIDs should be allowed during migration instead of failing it.
	*/
	ShouldAllowUnknownSids *bool `json:"shouldAllowUnknownSids,omitempty"`
	/*
	  A flag to indicate whether files at root should be ignored while migration instead of failing it.
	*/
	ShouldIgnoreFilesAtRoot *bool `json:"shouldIgnoreFilesAtRoot,omitempty"`
	/*
	  A flag to indicate whether long TLD name should be ignored while migration instead of failing it.
	*/
	ShouldIgnoreLongTldNames *bool `json:"shouldIgnoreLongTldNames,omitempty"`
	/*
	  A flag to indicate whether the migration iteration is only to migrate failed files/directories from the last iteration. If it is false and status is changed to running from default, then the migration will start for the whole source mount target. If it is true, then only a list of failed files from the last iteration are retried for migration.
	*/
	ShouldRetryOnlyFailedFiles *bool `json:"shouldRetryOnlyFailedFiles,omitempty"`
	/*
	  Source path
	*/
	SourcePath *string `json:"sourcePath"`

	Status *MigrationStatus `json:"status,omitempty"`
	/*
	  Indicates the status message of the migration sub-plan. It updates after each migration step, informing the current migration sub-plan stage. This is a read-only field.
	*/
	StatusMsg *string `json:"statusMsg,omitempty"`
	/*
	  Target mount-target extId.
	*/
	TargetMountTargetExtId *string `json:"targetMountTargetExtId"`
	/*
	  Target path
	*/
	TargetPath *string `json:"targetPath,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *MigrationSubPlan) MarshalJSON() ([]byte, error) {
	type MigrationSubPlanProxy MigrationSubPlan
	return json.Marshal(struct {
		*MigrationSubPlanProxy
		Name                   *string `json:"name,omitempty"`
		SourcePath             *string `json:"sourcePath,omitempty"`
		TargetMountTargetExtId *string `json:"targetMountTargetExtId,omitempty"`
	}{
		MigrationSubPlanProxy:  (*MigrationSubPlanProxy)(p),
		Name:                   p.Name,
		SourcePath:             p.SourcePath,
		TargetMountTargetExtId: p.TargetMountTargetExtId,
	})
}

func NewMigrationSubPlan() *MigrationSubPlan {
	p := new(MigrationSubPlan)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.MigrationSubPlan"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	p.IsFinalIteration = new(bool)
	*p.IsFinalIteration = false
	p.IsSyncTypeMirror = new(bool)
	*p.IsSyncTypeMirror = false
	p.ShouldAllowUnknownSids = new(bool)
	*p.ShouldAllowUnknownSids = true
	p.ShouldIgnoreFilesAtRoot = new(bool)
	*p.ShouldIgnoreFilesAtRoot = true
	p.ShouldIgnoreLongTldNames = new(bool)
	*p.ShouldIgnoreLongTldNames = true
	p.ShouldRetryOnlyFailedFiles = new(bool)
	*p.ShouldRetryOnlyFailedFiles = false

	return p
}

/*
To repeat the schedule on a per `Month` basis.
*/
type MonthlySchedule struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A list of the days of the month on which snapshots should be taken.
	*/
	DaysOfMonth []int `json:"daysOfMonth"`
}

func (p *MonthlySchedule) MarshalJSON() ([]byte, error) {
	type MonthlyScheduleProxy MonthlySchedule
	return json.Marshal(struct {
		*MonthlyScheduleProxy
		DaysOfMonth []int `json:"daysOfMonth,omitempty"`
	}{
		MonthlyScheduleProxy: (*MonthlyScheduleProxy)(p),
		DaysOfMonth:          p.DaysOfMonth,
	})
}

func NewMonthlySchedule() *MonthlySchedule {
	p := new(MonthlySchedule)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.MonthlySchedule"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Mount target object.
*/
type MountTarget struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	BlockedClients *BlockedClient `json:"blockedClients,omitempty"`
	/*
	  Comma-separated list of blocked file extensions.
	*/
	BlockedFileExtensions []string `json:"blockedFileExtensions,omitempty"`
	/*
	  The Connected Shares feature allows you to connect an empty folder, called a junction folder, in one share to another share. This allows you to share content from any part of a share's directory. For example, when creating a connected share, the connectedMountTargetPath is the path of the junction folder.
	*/
	ConnectedMountTargetPath *string `json:"connectedMountTargetPath,omitempty"`
	/*
	  Mount target description.
	*/
	Description *string `json:"description,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Flag to enable compression.
	*/
	IsCompressionEnabled *bool `json:"isCompressionEnabled,omitempty"`
	/*
	  Enable long name support.
	*/
	IsLongNameEnabled *bool `json:"isLongNameEnabled,omitempty"`
	/*
	  Flag to enable windows previous version.
	*/
	IsPreviousVersionEnabled *bool `json:"isPreviousVersionEnabled,omitempty"`
	/*
	  Pause snapshot operations on share.
	*/
	IsSnapshotPaused *bool `json:"isSnapshotPaused,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Maximum size of mount target in GiB.
	*/
	MaxSizeGB *int64 `json:"maxSizeGB,omitempty"`

	MultiProtocolProperties *MultiProtocolProperties `json:"multiProtocolProperties,omitempty"`
	/*
	  Mount target name.
	*/
	Name *string `json:"name"`

	NfsProperties *NfsProtocolProperties `json:"nfsProperties,omitempty"`
	/*
	  external identifier of the parent mount target. This is a read-only field.
	*/
	ParentMountTargetExtId *string `json:"parentMountTargetExtId,omitempty"`
	/*
	  Path of the nested mount target.
	*/
	Path *string `json:"path,omitempty"`

	Protocol *MountTargetProtocolType `json:"protocol,omitempty"`
	/*
	  List of secondary protocol types for the mount target.
	*/
	SecondaryProtocol []MountTargetProtocolType `json:"secondaryProtocol,omitempty"`

	SmbProperties *SmbProtocolProperties `json:"smbProperties,omitempty"`

	State *MountTargetState `json:"state,omitempty"`

	StatusType *StatusType `json:"statusType,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	Type *MountTargetType `json:"type,omitempty"`

	WorkloadType *MountTargetWorkloadType `json:"workloadType,omitempty"`

	WormSpec *WormSpec `json:"wormSpec,omitempty"`
}

func (p *MountTarget) MarshalJSON() ([]byte, error) {
	type MountTargetProxy MountTarget
	return json.Marshal(struct {
		*MountTargetProxy
		Name *string `json:"name,omitempty"`
	}{
		MountTargetProxy: (*MountTargetProxy)(p),
		Name:             p.Name,
	})
}

func NewMountTarget() *MountTarget {
	p := new(MountTarget)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.MountTarget"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	p.IsCompressionEnabled = new(bool)
	*p.IsCompressionEnabled = true
	p.IsLongNameEnabled = new(bool)
	*p.IsLongNameEnabled = false
	p.IsPreviousVersionEnabled = new(bool)
	*p.IsPreviousVersionEnabled = false

	return p
}

/*
Clone mount target model
*/
type MountTargetCloneSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Mount target description.
	*/
	Description *string `json:"description,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Mount target name.
	*/
	Name *string `json:"name"`
	/*
	  The external identifier of the snapshot. Example:48f78959-14a6-4c47-b5db-920460c4b668
	*/
	SnapshotExtId *string `json:"snapshotExtId,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *MountTargetCloneSpec) MarshalJSON() ([]byte, error) {
	type MountTargetCloneSpecProxy MountTargetCloneSpec
	return json.Marshal(struct {
		*MountTargetCloneSpecProxy
		Name *string `json:"name,omitempty"`
	}{
		MountTargetCloneSpecProxy: (*MountTargetCloneSpecProxy)(p),
		Name:                      p.Name,
	})
}

func NewMountTargetCloneSpec() *MountTargetCloneSpec {
	p := new(MountTargetCloneSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.MountTargetCloneSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Protocol type of the mount target.
*/
type MountTargetProtocolType int

const (
	MOUNTTARGETPROTOCOLTYPE_UNKNOWN      MountTargetProtocolType = 0
	MOUNTTARGETPROTOCOLTYPE_REDACTED     MountTargetProtocolType = 1
	MOUNTTARGETPROTOCOLTYPE_SMB          MountTargetProtocolType = 2
	MOUNTTARGETPROTOCOLTYPE_NFS          MountTargetProtocolType = 3
	MOUNTTARGETPROTOCOLTYPE_NONE         MountTargetProtocolType = 4
	MOUNTTARGETPROTOCOLTYPE_INCOMPATIBLE MountTargetProtocolType = 5
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *MountTargetProtocolType) name(index int) string {
	names := [...]string{
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

// Returns the name of the enum
func (e MountTargetProtocolType) GetName() string {
	index := int(e)
	names := [...]string{
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

// Returns the enum type given a string value
func (e *MountTargetProtocolType) index(name string) MountTargetProtocolType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SMB",
		"NFS",
		"NONE",
		"INCOMPATIBLE",
	}
	for idx := range names {
		if names[idx] == name {
			return MountTargetProtocolType(idx)
		}
	}
	return MOUNTTARGETPROTOCOLTYPE_UNKNOWN
}

func (e *MountTargetProtocolType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for MountTargetProtocolType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *MountTargetProtocolType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e MountTargetProtocolType) Ref() *MountTargetProtocolType {
	return &e
}

/*
State of the mount target.
*/
type MountTargetState int

const (
	MOUNTTARGETSTATE_UNKNOWN  MountTargetState = 0
	MOUNTTARGETSTATE_REDACTED MountTargetState = 1
	MOUNTTARGETSTATE_ONLINE   MountTargetState = 2
	MOUNTTARGETSTATE_OFFLINE  MountTargetState = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *MountTargetState) name(index int) string {
	names := [...]string{
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

// Returns the name of the enum
func (e MountTargetState) GetName() string {
	index := int(e)
	names := [...]string{
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

// Returns the enum type given a string value
func (e *MountTargetState) index(name string) MountTargetState {
	names := [...]string{
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

/*
Supported mount target types.
*/
type MountTargetType int

const (
	MOUNTTARGETTYPE_UNKNOWN     MountTargetType = 0
	MOUNTTARGETTYPE_REDACTED    MountTargetType = 1
	MOUNTTARGETTYPE_HOMES       MountTargetType = 2
	MOUNTTARGETTYPE_GENERAL     MountTargetType = 3
	MOUNTTARGETTYPE_DISTRIBUTED MountTargetType = 4
	MOUNTTARGETTYPE_STANDARD    MountTargetType = 5
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *MountTargetType) name(index int) string {
	names := [...]string{
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

// Returns the name of the enum
func (e MountTargetType) GetName() string {
	index := int(e)
	names := [...]string{
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

// Returns the enum type given a string value
func (e *MountTargetType) index(name string) MountTargetType {
	names := [...]string{
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

/*
Supported mount target workload types.
*/
type MountTargetWorkloadType int

const (
	MOUNTTARGETWORKLOADTYPE_UNKNOWN    MountTargetWorkloadType = 0
	MOUNTTARGETWORKLOADTYPE_REDACTED   MountTargetWorkloadType = 1
	MOUNTTARGETWORKLOADTYPE_UNDEFINED  MountTargetWorkloadType = 2
	MOUNTTARGETWORKLOADTYPE_RANDOM     MountTargetWorkloadType = 3
	MOUNTTARGETWORKLOADTYPE_DEFAULT    MountTargetWorkloadType = 4
	MOUNTTARGETWORKLOADTYPE_SEQUENTIAL MountTargetWorkloadType = 5
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *MountTargetWorkloadType) name(index int) string {
	names := [...]string{
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

// Returns the name of the enum
func (e MountTargetWorkloadType) GetName() string {
	index := int(e)
	names := [...]string{
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

// Returns the enum type given a string value
func (e *MountTargetWorkloadType) index(name string) MountTargetWorkloadType {
	names := [...]string{
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

/*
Configuration to enable current and future mount targets.
*/
type MountTargetsEnablementType int

const (
	MOUNTTARGETSENABLEMENTTYPE_UNKNOWN                          MountTargetsEnablementType = 0
	MOUNTTARGETSENABLEMENTTYPE_REDACTED                         MountTargetsEnablementType = 1
	MOUNTTARGETSENABLEMENTTYPE_ALL_CURRENT_MOUNT_TARGETS        MountTargetsEnablementType = 2
	MOUNTTARGETSENABLEMENTTYPE_ALL_CURRENT_FUTURE_MOUNT_TARGETS MountTargetsEnablementType = 3
	MOUNTTARGETSENABLEMENTTYPE_ALL_FUTURE_MOUNT_TARGETS         MountTargetsEnablementType = 4
	MOUNTTARGETSENABLEMENTTYPE_NONE                             MountTargetsEnablementType = 5
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *MountTargetsEnablementType) name(index int) string {
	names := [...]string{
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

// Returns the name of the enum
func (e MountTargetsEnablementType) GetName() string {
	index := int(e)
	names := [...]string{
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

// Returns the enum type given a string value
func (e *MountTargetsEnablementType) index(name string) MountTargetsEnablementType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ALL_CURRENT_MOUNT_TARGETS",
		"ALL_CURRENT_FUTURE_MOUNT_TARGETS",
		"ALL_FUTURE_MOUNT_TARGETS",
		"NONE",
	}
	for idx := range names {
		if names[idx] == name {
			return MountTargetsEnablementType(idx)
		}
	}
	return MOUNTTARGETSENABLEMENTTYPE_UNKNOWN
}

func (e *MountTargetsEnablementType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for MountTargetsEnablementType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *MountTargetsEnablementType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e MountTargetsEnablementType) Ref() *MountTargetsEnablementType {
	return &e
}

/*
Properties for multiprotocol supported mount target.
*/
type MultiProtocolProperties struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Flag to enable case sensitive namespace.
	*/
	IsCaseSensitiveNamespaceEnabled *bool `json:"isCaseSensitiveNamespaceEnabled,omitempty"`
	/*
	  Flag to enable simultaneous access to same files across protocols.
	*/
	IsSimultaneousAccessEnabled *bool `json:"isSimultaneousAccessEnabled,omitempty"`
	/*
	  Flag to enable symlink creation.
	*/
	IsSymlinkCreationEnabled *bool `json:"isSymlinkCreationEnabled,omitempty"`
}

func NewMultiProtocolProperties() *MultiProtocolProperties {
	p := new(MultiProtocolProperties)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.MultiProtocolProperties"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	p.IsCaseSensitiveNamespaceEnabled = new(bool)
	*p.IsCaseSensitiveNamespaceEnabled = false
	p.IsSimultaneousAccessEnabled = new(bool)
	*p.IsSimultaneousAccessEnabled = false
	p.IsSymlinkCreationEnabled = new(bool)
	*p.IsSymlinkCreationEnabled = false

	return p
}

/*
Unified namespace member configuration represents a list of member file servers and a flag to indicate whether it is hosting namespace
*/
type NamespaceMemberConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The external identifier of the member file server which belongs to unified namespace.
	*/
	FileServerExtId *string `json:"fileServerExtId"`
	/*
	  This flag indicates whether the file server is hosting the namespace in unified namespace. Core member's namespace will be used by the clients to access the shares. Core member cannot be changed once the namespace is created.
	*/
	IsCoreMember *bool `json:"isCoreMember,omitempty"`
}

func (p *NamespaceMemberConfig) MarshalJSON() ([]byte, error) {
	type NamespaceMemberConfigProxy NamespaceMemberConfig
	return json.Marshal(struct {
		*NamespaceMemberConfigProxy
		FileServerExtId *string `json:"fileServerExtId,omitempty"`
	}{
		NamespaceMemberConfigProxy: (*NamespaceMemberConfigProxy)(p),
		FileServerExtId:            p.FileServerExtId,
	})
}

func NewNamespaceMemberConfig() *NamespaceMemberConfig {
	p := new(NamespaceMemberConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.NamespaceMemberConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	p.IsCoreMember = new(bool)
	*p.IsCoreMember = false

	return p
}

/*
NFS protocol properties for mount target.
*/
type NfsProtocolProperties struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	AccessType *AccessType `json:"accessType,omitempty"`

	AnonymousIdentifier *AnonymousIdentifier `json:"anonymousIdentifier,omitempty"`

	AuthenticationType *FilesAuthenticationType `json:"authenticationType,omitempty"`
	/*
	  Access exception client list of the mount target.
	*/
	ClientExceptions []ClientException `json:"clientExceptions,omitempty"`

	SquashType *SquashType `json:"squashType,omitempty"`
}

func NewNfsProtocolProperties() *NfsProtocolProperties {
	p := new(NfsProtocolProperties)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.NfsProtocolProperties"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
NFS version associated with the file server.
*/
type NfsVersion int

const (
	NFSVERSION_UNKNOWN  NfsVersion = 0
	NFSVERSION_REDACTED NfsVersion = 1
	NFSVERSION_NFSV4    NfsVersion = 2
	NFSVERSION_NFSV3    NfsVersion = 3
	NFSVERSION_NFSV3V4  NfsVersion = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *NfsVersion) name(index int) string {
	names := [...]string{
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

// Returns the name of the enum
func (e NfsVersion) GetName() string {
	index := int(e)
	names := [...]string{
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

// Returns the enum type given a string value
func (e *NfsVersion) index(name string) NfsVersion {
	names := [...]string{
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

/*
A representation of a notification policy model.
*/
type NotificationPolicy struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A list of users and client IPs, whose notifications needs to be blocked for a partner server.
	*/
	BlockedClients []BlockedNotificationClient `json:"blockedClients,omitempty"`
	/*
	  A description of notification policy entity.
	*/
	Description *string `json:"description,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`

	FileBlockingMode *FileBlockingMode `json:"fileBlockingMode,omitempty"`
	/*
	  List of file blocking extensions. For example: *.mp3.
	*/
	FileExtensions []string `json:"fileExtensions,omitempty"`
	/*
	  Setting this flag ensures that a secure connection is established between AFS and the partner server.
	*/
	HasSecuredConnection *bool `json:"hasSecuredConnection,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  A list of mount target external identifiers.
	*/
	MountTargetExtIds []string `json:"mountTargetExtIds,omitempty"`
	/*
	  Notification policy name.
	*/
	Name *string `json:"name"`
	/*
	  Defines the list of operations on the files.
	*/
	Operations []FileOperation `json:"operations"`
	/*
	  A list of partner server external identifiers.
	*/
	PartnerServerExtIds []string `json:"partnerServerExtIds,omitempty"`
	/*
	  List of mount target protocol type associated to the notification policy.
	*/
	ProtocolTypes []ProtocolType `json:"protocolTypes"`
	/*
	  Setting this flag ensures that a notification policy is applicable to all the mount targets.
	*/
	ShouldIncludeAllMountTargets *bool `json:"shouldIncludeAllMountTargets,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *NotificationPolicy) MarshalJSON() ([]byte, error) {
	type NotificationPolicyProxy NotificationPolicy
	return json.Marshal(struct {
		*NotificationPolicyProxy
		Name          *string         `json:"name,omitempty"`
		Operations    []FileOperation `json:"operations,omitempty"`
		ProtocolTypes []ProtocolType  `json:"protocolTypes,omitempty"`
	}{
		NotificationPolicyProxy: (*NotificationPolicyProxy)(p),
		Name:                    p.Name,
		Operations:              p.Operations,
		ProtocolTypes:           p.ProtocolTypes,
	})
}

func NewNotificationPolicy() *NotificationPolicy {
	p := new(NotificationPolicy)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.NotificationPolicy"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	p.HasSecuredConnection = new(bool)
	*p.HasSecuredConnection = false
	p.ShouldIncludeAllMountTargets = new(bool)
	*p.ShouldIncludeAllMountTargets = false

	return p
}

/*
Indicates the user definition for a notification policy. The user can be either SID, UID or username.
*/
type NotificationPolicyUser struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Notification policy user name.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  Notification policy user security identifier.
	*/
	Sid *string `json:"sid,omitempty"`
	/*
	  User identifier. Default value of the user identifier is -2.
	*/
	Uid *int64 `json:"uid,omitempty"`
}

func NewNotificationPolicyUser() *NotificationPolicyUser {
	p := new(NotificationPolicyUser)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.NotificationPolicyUser"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Object store profile specification.
*/
type ObjectStoreConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Base URL.
	*/
	BaseUrl *string `json:"baseUrl"`
	/*
	  CA certificate.
	*/
	CaCertContent *string `json:"caCertContent,omitempty"`
	/*

	 */
	ConfigurationItemDiscriminator_ *string `json:"$configurationItemDiscriminator,omitempty"`
	/*
	  AWS/Azure configuration for tiering profile.
	User need to provide $objectType : files.v4.config.AWSConfig for AWS/NUTANIX object store type.
	For Azure object store type provide $objectType : files.v4.config.AzureConfig.
	*/
	Configuration *OneOfObjectStoreConfigConfiguration `json:"configuration"`

	ProxyServer *ProxyServer `json:"proxyServer,omitempty"`
}

func (p *ObjectStoreConfig) MarshalJSON() ([]byte, error) {
	type ObjectStoreConfigProxy ObjectStoreConfig
	return json.Marshal(struct {
		*ObjectStoreConfigProxy
		BaseUrl       *string                              `json:"baseUrl,omitempty"`
		Configuration *OneOfObjectStoreConfigConfiguration `json:"configuration,omitempty"`
	}{
		ObjectStoreConfigProxy: (*ObjectStoreConfigProxy)(p),
		BaseUrl:                p.BaseUrl,
		Configuration:          p.Configuration,
	})
}

func NewObjectStoreConfig() *ObjectStoreConfig {
	p := new(ObjectStoreConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.ObjectStoreConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ObjectStoreConfig) GetConfiguration() interface{} {
	if nil == p.Configuration {
		return nil
	}
	return p.Configuration.GetValue()
}

func (p *ObjectStoreConfig) SetConfiguration(v interface{}) error {
	if nil == p.Configuration {
		p.Configuration = NewOneOfObjectStoreConfigConfiguration()
	}
	e := p.Configuration.SetValue(v)
	if nil == e {
		if nil == p.ConfigurationItemDiscriminator_ {
			p.ConfigurationItemDiscriminator_ = new(string)
		}
		*p.ConfigurationItemDiscriminator_ = *p.Configuration.Discriminator
	}
	return e
}

/*
Object store profile specification.
*/
type ObjectStoreProfile struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Enable SSL verify peer certificate.
	*/
	IsSslPeerVerficationEnabled *bool `json:"isSslPeerVerficationEnabled,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Mount target list which should be included in tiering profile.
	*/
	MountTargetExtIds []string `json:"mountTargetExtIds,omitempty"`

	MountTargetsEnablementType *MountTargetsEnablementType `json:"mountTargetsEnablementType,omitempty"`
	/*
	  Object store cloud profile name.
	*/
	Name *string `json:"name"`

	ObjectStoreConfig *ObjectStoreConfig `json:"objectStoreConfig"`

	ObjectStoreType *ObjectStoreType `json:"objectStoreType"`

	RecoveryObjectStoreConfig *ObjectStoreConfig `json:"recoveryObjectStoreConfig,omitempty"`
	/*
	  Indicates time in days for which the data will be maintained in the cloud after it is deleted from the local storage on the file server.
	*/
	RetentionPeriodDays *int64 `json:"retentionPeriodDays,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *ObjectStoreProfile) MarshalJSON() ([]byte, error) {
	type ObjectStoreProfileProxy ObjectStoreProfile
	return json.Marshal(struct {
		*ObjectStoreProfileProxy
		Name              *string            `json:"name,omitempty"`
		ObjectStoreConfig *ObjectStoreConfig `json:"objectStoreConfig,omitempty"`
		ObjectStoreType   *ObjectStoreType   `json:"objectStoreType,omitempty"`
	}{
		ObjectStoreProfileProxy: (*ObjectStoreProfileProxy)(p),
		Name:                    p.Name,
		ObjectStoreConfig:       p.ObjectStoreConfig,
		ObjectStoreType:         p.ObjectStoreType,
	})
}

func NewObjectStoreProfile() *ObjectStoreProfile {
	p := new(ObjectStoreProfile)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.ObjectStoreProfile"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	p.RetentionPeriodDays = new(int64)
	*p.RetentionPeriodDays = 1825

	return p
}

/*
Object store type.
*/
type ObjectStoreType int

const (
	OBJECTSTORETYPE_UNKNOWN         ObjectStoreType = 0
	OBJECTSTORETYPE_REDACTED        ObjectStoreType = 1
	OBJECTSTORETYPE_NUTANIX         ObjectStoreType = 2
	OBJECTSTORETYPE_AWS             ObjectStoreType = 3
	OBJECTSTORETYPE_GCP             ObjectStoreType = 4
	OBJECTSTORETYPE_AZURE           ObjectStoreType = 5
	OBJECTSTORETYPE_AWS_STANDARD_IA ObjectStoreType = 6
	OBJECTSTORETYPE_OTHER           ObjectStoreType = 7
	OBJECTSTORETYPE_AWS_GLACIER_IR  ObjectStoreType = 8
	OBJECTSTORETYPE_AZURE_COOL      ObjectStoreType = 9
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ObjectStoreType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NUTANIX",
		"AWS",
		"GCP",
		"AZURE",
		"AWS_STANDARD_IA",
		"OTHER",
		"AWS_GLACIER_IR",
		"AZURE_COOL",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ObjectStoreType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NUTANIX",
		"AWS",
		"GCP",
		"AZURE",
		"AWS_STANDARD_IA",
		"OTHER",
		"AWS_GLACIER_IR",
		"AZURE_COOL",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ObjectStoreType) index(name string) ObjectStoreType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NUTANIX",
		"AWS",
		"GCP",
		"AZURE",
		"AWS_STANDARD_IA",
		"OTHER",
		"AWS_GLACIER_IR",
		"AZURE_COOL",
	}
	for idx := range names {
		if names[idx] == name {
			return ObjectStoreType(idx)
		}
	}
	return OBJECTSTORETYPE_UNKNOWN
}

func (e *ObjectStoreType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ObjectStoreType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ObjectStoreType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ObjectStoreType) Ref() *ObjectStoreType {
	return &e
}

/*
Object type detail
*/
type ObjectType int

const (
	OBJECTTYPE_UNKNOWN   ObjectType = 0
	OBJECTTYPE_REDACTED  ObjectType = 1
	OBJECTTYPE_FILE      ObjectType = 2
	OBJECTTYPE_DIRECTORY ObjectType = 3
	OBJECTTYPE_FIFO      ObjectType = 4
	OBJECTTYPE_LINK      ObjectType = 5
	OBJECTTYPE_SOCKET    ObjectType = 6
	OBJECTTYPE_PORT      ObjectType = 7
	OBJECTTYPE_DOOR      ObjectType = 8
	OBJECTTYPE_BLOCK     ObjectType = 9
	OBJECTTYPE_CHARACTER ObjectType = 10
	OBJECTTYPE_HARD_LINK ObjectType = 11
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ObjectType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"FILE",
		"DIRECTORY",
		"FIFO",
		"LINK",
		"SOCKET",
		"PORT",
		"DOOR",
		"BLOCK",
		"CHARACTER",
		"HARD_LINK",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ObjectType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"FILE",
		"DIRECTORY",
		"FIFO",
		"LINK",
		"SOCKET",
		"PORT",
		"DOOR",
		"BLOCK",
		"CHARACTER",
		"HARD_LINK",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ObjectType) index(name string) ObjectType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"FILE",
		"DIRECTORY",
		"FIFO",
		"LINK",
		"SOCKET",
		"PORT",
		"DOOR",
		"BLOCK",
		"CHARACTER",
		"HARD_LINK",
	}
	for idx := range names {
		if names[idx] == name {
			return ObjectType(idx)
		}
	}
	return OBJECTTYPE_UNKNOWN
}

func (e *ObjectType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ObjectType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ObjectType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ObjectType) Ref() *ObjectType {
	return &e
}

/*
Backup server (client IP-address/FDQN) management operation.
*/
type OperationType int

const (
	OPERATIONTYPE_UNKNOWN  OperationType = 0
	OPERATIONTYPE_REDACTED OperationType = 1
	OPERATIONTYPE_ADD      OperationType = 2
	OPERATIONTYPE_REMOVE   OperationType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *OperationType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ADD",
		"REMOVE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e OperationType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ADD",
		"REMOVE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *OperationType) index(name string) OperationType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ADD",
		"REMOVE",
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

/*
Configuration to change ownership of the impacted `VDI_SYNC` users.
*/
type OwnershipSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The external identifier of the file server from which ownership of impacted users is being released.
	*/
	CurrentOwnerFileServerExtId *string `json:"currentOwnerFileServerExtId"`
	/*
	  The external identifier of the file server that will take ownership of the impacted users from the other file server.
	*/
	NewOwnerFileServerExtId *string `json:"newOwnerFileServerExtId"`
}

func (p *OwnershipSpec) MarshalJSON() ([]byte, error) {
	type OwnershipSpecProxy OwnershipSpec
	return json.Marshal(struct {
		*OwnershipSpecProxy
		CurrentOwnerFileServerExtId *string `json:"currentOwnerFileServerExtId,omitempty"`
		NewOwnerFileServerExtId     *string `json:"newOwnerFileServerExtId,omitempty"`
	}{
		OwnershipSpecProxy:          (*OwnershipSpecProxy)(p),
		CurrentOwnerFileServerExtId: p.CurrentOwnerFileServerExtId,
		NewOwnerFileServerExtId:     p.NewOwnerFileServerExtId,
	})
}

func NewOwnershipSpec() *OwnershipSpec {
	p := new(OwnershipSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.OwnershipSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Partner server model.
*/
type PartnerServer struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	BackupServerConfig *BackupServerConfig `json:"backupServerConfig,omitempty"`
	/*
	  Partner server details or any third-party vendors to register the server with AFS.
	*/
	Description *string `json:"description,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Partner server name.
	*/
	Name *string `json:"name"`

	PartnerType *PartnerType `json:"partnerType"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  Vendor name of the partner server. For partner type ANTIVIRUS, vendor name denotes the icapServiceName.
	*/
	VendorName *string `json:"vendorName"`

	VendorProperties *VendorProperties `json:"vendorProperties,omitempty"`
}

func (p *PartnerServer) MarshalJSON() ([]byte, error) {
	type PartnerServerProxy PartnerServer
	return json.Marshal(struct {
		*PartnerServerProxy
		Name        *string      `json:"name,omitempty"`
		PartnerType *PartnerType `json:"partnerType,omitempty"`
		VendorName  *string      `json:"vendorName,omitempty"`
	}{
		PartnerServerProxy: (*PartnerServerProxy)(p),
		Name:               p.Name,
		PartnerType:        p.PartnerType,
		VendorName:         p.VendorName,
	})
}

func NewPartnerServer() *PartnerServer {
	p := new(PartnerServer)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.PartnerServer"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Usage type of the partner server.
*/
type PartnerType int

const (
	PARTNERTYPE_UNKNOWN      PartnerType = 0
	PARTNERTYPE_REDACTED     PartnerType = 1
	PARTNERTYPE_NOTIFICATION PartnerType = 2
	PARTNERTYPE_BACKUP       PartnerType = 3
	PARTNERTYPE_MIGRATION    PartnerType = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *PartnerType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NOTIFICATION",
		"BACKUP",
		"MIGRATION",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e PartnerType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NOTIFICATION",
		"BACKUP",
		"MIGRATION",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *PartnerType) index(name string) PartnerType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NOTIFICATION",
		"BACKUP",
		"MIGRATION",
	}
	for idx := range names {
		if names[idx] == name {
			return PartnerType(idx)
		}
	}
	return PARTNERTYPE_UNKNOWN
}

func (e *PartnerType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for PartnerType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *PartnerType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e PartnerType) Ref() *PartnerType {
	return &e
}

/*
Permission type for SMB share access control entry(ACE).
*/
type PermissionType int

const (
	PERMISSIONTYPE_UNKNOWN  PermissionType = 0
	PERMISSIONTYPE_REDACTED PermissionType = 1
	PERMISSIONTYPE_ALLOW    PermissionType = 2
	PERMISSIONTYPE_DENY     PermissionType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *PermissionType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ALLOW",
		"DENY",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e PermissionType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ALLOW",
		"DENY",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *PermissionType) index(name string) PermissionType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ALLOW",
		"DENY",
	}
	for idx := range names {
		if names[idx] == name {
			return PermissionType(idx)
		}
	}
	return PERMISSIONTYPE_UNKNOWN
}

func (e *PermissionType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for PermissionType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *PermissionType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e PermissionType) Ref() *PermissionType {
	return &e
}

/*
Status for the replication policy.
*/
type PolicyStatus int

const (
	POLICYSTATUS_UNKNOWN           PolicyStatus = 0
	POLICYSTATUS_REDACTED          PolicyStatus = 1
	POLICYSTATUS_ENABLED           PolicyStatus = 2
	POLICYSTATUS_DISABLED          PolicyStatus = 3
	POLICYSTATUS_PARTIALLY_ENABLED PolicyStatus = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *PolicyStatus) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ENABLED",
		"DISABLED",
		"PARTIALLY_ENABLED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e PolicyStatus) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ENABLED",
		"DISABLED",
		"PARTIALLY_ENABLED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *PolicyStatus) index(name string) PolicyStatus {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ENABLED",
		"DISABLED",
		"PARTIALLY_ENABLED",
	}
	for idx := range names {
		if names[idx] == name {
			return PolicyStatus(idx)
		}
	}
	return POLICYSTATUS_UNKNOWN
}

func (e *PolicyStatus) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for PolicyStatus:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *PolicyStatus) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e PolicyStatus) Ref() *PolicyStatus {
	return &e
}

/*
Principle type.
*/
type PrincipalType int

const (
	PRINCIPALTYPE_UNKNOWN  PrincipalType = 0
	PRINCIPALTYPE_REDACTED PrincipalType = 1
	PRINCIPALTYPE_USER     PrincipalType = 2
	PRINCIPALTYPE_GROUP    PrincipalType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *PrincipalType) name(index int) string {
	names := [...]string{
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

// Returns the name of the enum
func (e PrincipalType) GetName() string {
	index := int(e)
	names := [...]string{
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

// Returns the enum type given a string value
func (e *PrincipalType) index(name string) PrincipalType {
	names := [...]string{
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

/*
Protocol type for a file server.
*/
type ProtocolType int

const (
	PROTOCOLTYPE_UNKNOWN  ProtocolType = 0
	PROTOCOLTYPE_REDACTED ProtocolType = 1
	PROTOCOLTYPE_SMB      ProtocolType = 2
	PROTOCOLTYPE_NFS      ProtocolType = 3
	PROTOCOLTYPE_SMB_NFS  ProtocolType = 4
	PROTOCOLTYPE_NONE     ProtocolType = 5
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ProtocolType) name(index int) string {
	names := [...]string{
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

// Returns the name of the enum
func (e ProtocolType) GetName() string {
	index := int(e)
	names := [...]string{
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

// Returns the enum type given a string value
func (e *ProtocolType) index(name string) ProtocolType {
	names := [...]string{
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

/*
HTTP proxy server configuration
*/
type ProxyServer struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Credential *Credential `json:"credential,omitempty"`
	/*
	  HTTP proxy server URL.
	*/
	Url *string `json:"url,omitempty"`
}

func NewProxyServer() *ProxyServer {
	p := new(ProxyServer)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.ProxyServer"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Quota policy model.
*/
type QuotaPolicy struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	EnforcementType *EnforcementType `json:"enforcementType"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Enables email notifications for the user or group specified in the principal type.A notification will only be sent if the user or group is close to the quota provided.
	*/
	IsNotificationEnabled *bool `json:"isNotificationEnabled,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  List of recipient's emails.
	*/
	NotificationRecipients []string `json:"notificationRecipients,omitempty"`
	/*
	  Principal name is the name of the user or group, assigned to the principal type.
	*/
	PrincipalName *string `json:"principalName"`

	PrincipalType *PrincipalType `json:"principalType"`
	/*
	  Quota size in bytes.
	*/
	SizeInBytes *int64 `json:"sizeInBytes"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *QuotaPolicy) MarshalJSON() ([]byte, error) {
	type QuotaPolicyProxy QuotaPolicy
	return json.Marshal(struct {
		*QuotaPolicyProxy
		EnforcementType *EnforcementType `json:"enforcementType,omitempty"`
		PrincipalName   *string          `json:"principalName,omitempty"`
		PrincipalType   *PrincipalType   `json:"principalType,omitempty"`
		SizeInBytes     *int64           `json:"sizeInBytes,omitempty"`
	}{
		QuotaPolicyProxy: (*QuotaPolicyProxy)(p),
		EnforcementType:  p.EnforcementType,
		PrincipalName:    p.PrincipalName,
		PrincipalType:    p.PrincipalType,
		SizeInBytes:      p.SizeInBytes,
	})
}

func NewQuotaPolicy() *QuotaPolicy {
	p := new(QuotaPolicy)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.QuotaPolicy"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	p.IsNotificationEnabled = new(bool)
	*p.IsNotificationEnabled = true
	p.SizeInBytes = new(int64)
	*p.SizeInBytes = 1

	return p
}

/*
Ransomware configuration information.
*/
type RansomwareConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of mount targets to be excluded from ransomware configuration.
	*/
	ExcludedMountTargetExtIds []string `json:"excludedMountTargetExtIds,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Ransomware file patterns. Ex - *.txt, ?.db.
	*/
	FileExtensions []string `json:"fileExtensions,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewRansomwareConfig() *RansomwareConfig {
	p := new(RansomwareConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.RansomwareConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /files/v4.0/config/file-servers/{extId}/$actions/recalculate-cold-data Post operation
*/
type RecalculateColdDataApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfRecalculateColdDataApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewRecalculateColdDataApiResponse() *RecalculateColdDataApiResponse {
	p := new(RecalculateColdDataApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.RecalculateColdDataApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *RecalculateColdDataApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *RecalculateColdDataApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfRecalculateColdDataApiResponseData()
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
REST response for all response codes in API path /files/v4.0/config/file-servers/{extId}/$actions/recall-tiered-files Post operation
*/
type RecallTieredFilesApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfRecallTieredFilesApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewRecallTieredFilesApiResponse() *RecallTieredFilesApiResponse {
	p := new(RecallTieredFilesApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.RecallTieredFilesApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *RecallTieredFilesApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *RecallTieredFilesApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfRecallTieredFilesApiResponseData()
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

type RecallTieredFilesSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Path of the files to be tiered/recalled relative to root Mount target.
	*/
	FilePaths []string `json:"filePaths,omitempty"`
	/*
	  List of inode numbers strings. Inode number should be in the format Inode/Inode:GenId.
	*/
	INodeNumbers []string `json:"iNodeNumbers,omitempty"`
	/*
	  Boolean field to indicate if all files belong to the same VG.
	*/
	IsSingleDirTree *bool `json:"isSingleDirTree,omitempty"`
	/*
	  Mount target external identifier for files to tier/recall.
	*/
	MountTargetExtId *string `json:"mountTargetExtId"`
}

func NewRecallTieredFilesSpec() *RecallTieredFilesSpec {
	p := new(RecallTieredFilesSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.RecallTieredFilesSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	p.IsSingleDirTree = new(bool)
	*p.IsSingleDirTree = false

	return p
}

/*
File server recommendation description.
*/
type Recommendation struct {
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
	Links []import2.ApiLink `json:"links,omitempty"`

	RecommendationType *RecommendationType `json:"recommendationType,omitempty"`
	/*
	  The recommended number of file server VMs.
	*/
	RecommendedVmCount *int `json:"recommendedVmCount,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewRecommendation() *Recommendation {
	p := new(Recommendation)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.Recommendation"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
File server recommendation type.
*/
type RecommendationType int

const (
	RECOMMENDATIONTYPE_UNKNOWN   RecommendationType = 0
	RECOMMENDATIONTYPE_REDACTED  RecommendationType = 1
	RECOMMENDATIONTYPE_SCALE_UP  RecommendationType = 2
	RECOMMENDATIONTYPE_SCALE_OUT RecommendationType = 3
	RECOMMENDATIONTYPE_REBALANCE RecommendationType = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *RecommendationType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SCALE_UP",
		"SCALE_OUT",
		"REBALANCE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e RecommendationType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SCALE_UP",
		"SCALE_OUT",
		"REBALANCE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *RecommendationType) index(name string) RecommendationType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SCALE_UP",
		"SCALE_OUT",
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

/*
Replication configuration represents combination of file server entities involved in the replication policy, schedules, policy status and replication summary status. For `SMART_DR` the system supports a single replication configuration object per replication policy. For `DATA_SYNC` or `VDI_SYNC`, replicationConfigurations are allowed between 1 and 10. Note that in `VDI_SYNC`, the primary file server of the first replication entity is the preferred file server for replication.
*/
type ReplicationConfiguration struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Domain manager (ex. Prism central) external identifier where the primary file server of the replication configuration is registered. This is a required field in case of multi Prism central policy. When primary file server is not local to the Prism central.
	*/
	PrimaryDomainManagerExtId *string `json:"primaryDomainManagerExtId,omitempty"`
	/*
	  The external identifier of the primary file server for the replication.
	*/
	PrimaryFileServerExtId *string `json:"primaryFileServerExtId"`
	/*
	  A list of replication entities belonging to the replication configuration.
	*/
	ReplicationEntities []ReplicationEntity `json:"replicationEntities,omitempty"`

	ReplicationSummary *ReplicationSummary `json:"replicationSummary,omitempty"`

	Schedule *ReplicationSchedule `json:"schedule,omitempty"`
	/*
	  Domain manager (ex. Prism central) external identifier where the secondary file server of the replication configuration is registered. This is a required field in case of multi Prism central policy. When secondary file server is not local to the Prism central.
	*/
	SecondaryDomainManagerExtId *string `json:"secondaryDomainManagerExtId,omitempty"`
	/*
	  The external identifier of the secondary file server for the replication.
	*/
	SecondaryFileServerExtId *string `json:"secondaryFileServerExtId"`
	/*
	  This attribute indicates whether all replication jobs associated with the replication configuration need to be aborted during the disable replication configuration operation. This field is ignored if used for any operation other than disabling the policy. It is `false` by default.
	*/
	ShouldCancelOngoingReplicationJobs *bool `json:"shouldCancelOngoingReplicationJobs,omitempty"`

	Status *PolicyStatus `json:"status,omitempty"`
}

func (p *ReplicationConfiguration) MarshalJSON() ([]byte, error) {
	type ReplicationConfigurationProxy ReplicationConfiguration
	return json.Marshal(struct {
		*ReplicationConfigurationProxy
		PrimaryFileServerExtId   *string `json:"primaryFileServerExtId,omitempty"`
		SecondaryFileServerExtId *string `json:"secondaryFileServerExtId,omitempty"`
	}{
		ReplicationConfigurationProxy: (*ReplicationConfigurationProxy)(p),
		PrimaryFileServerExtId:        p.PrimaryFileServerExtId,
		SecondaryFileServerExtId:      p.SecondaryFileServerExtId,
	})
}

func NewReplicationConfiguration() *ReplicationConfiguration {
	p := new(ReplicationConfiguration)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.ReplicationConfiguration"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	p.ShouldCancelOngoingReplicationJobs = new(bool)
	*p.ShouldCancelOngoingReplicationJobs = false

	return p
}

/*
Replication entity belonging to a replication configuration. Here, for `SMART_DR` and `DATA_SYNC`, primary represents source file server or path, and secondary represents target file server or path.
*/
type ReplicationEntity struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Displays a list of wildcard-based patterns to exclude files and directories from replication. Used only for `DATA_SYNC`. The wildcard-based patterns can contain “” which replaces all non-“/“ characters and “*” which replaces any character. Ex. `/share1/subPath/`
	*/
	ExcludeDirPatterns []string `json:"excludeDirPatterns,omitempty"`
	/*
	  A flag indicating whether any of the mount targets of the VDI sync replication entity is unavailable for synchronization. This flag is read-only. This flag will be set to `false` in the following two scenarios: ``` - when the mount target in the replication entity is deleted from the file server - when any file server deletes the sync entity (path) from the policy. ```
	*/
	IsAvailableForReplication *bool `json:"isAvailableForReplication,omitempty"`
	/*
	  The external identifier of the mount target. Example:9c1e537d-6777-4c22-5d41-ddd0c3337aa9.
	*/
	PrimaryFileServerMountTargetExtId *string `json:"primaryFileServerMountTargetExtId,omitempty"`
	/*
	  Path of the nested mount target.
	*/
	PrimaryFileServerMountTargetPath *string `json:"primaryFileServerMountTargetPath,omitempty"`
	/*
	  The external identifier of the mount target. Example:9c1e537d-6777-4c22-5d41-ddd0c3337aa9.
	*/
	SecondaryFileServerMountTargetExtId *string `json:"secondaryFileServerMountTargetExtId,omitempty"`
	/*
	  Path of the nested mount target.
	*/
	SecondaryFileServerMountTargetPath *string `json:"secondaryFileServerMountTargetPath,omitempty"`
}

func NewReplicationEntity() *ReplicationEntity {
	p := new(ReplicationEntity)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.ReplicationEntity"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Replication policy model.
*/
type ReplicationPolicy struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	ChangeUserSessionOwnershipSpec *OwnershipSpec `json:"changeUserSessionOwnershipSpec,omitempty"`
	/*
	  Replication policy description.
	*/
	Description *string `json:"description,omitempty"`
	/*
	  Display a list of file patterns to exclude from the replication. For example, ["*.txt", "*.yaml"].
	This is applicable for `DATA_SYNC`.
	*/
	ExcludeFilePatterns []string `json:"excludeFilePatterns,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A boolean flag to reverse the data replication configuration during a planned failover or to resume replication from the secondary site. Use this configuration only if reverse replication is required for failover. A reverse policy can be created when the file servers are in a failed-over state. No replication entities need to be specified to create a reverse policy. This configuration is applicable only for the `SMART_DR` policy.
	*/
	IsReverse *bool `json:"isReverse,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Name of the replication policy.
	*/
	Name *string `json:"name"`
	/*
	  Replication configuration represents combination of file server entities involved in the replication policy, schedules, policy status and replication summary status. For `SMART_DR` the system supports a single replication configuration object per replication policy. For `DATA_SYNC` or `VDI_SYNC`, replicationConfigurations are allowed between 1 and 10. Note that in `VDI_SYNC`, the primary file server of the first replication entity is the preferred file server for replication.
	*/
	ReplicationConfigurations []ReplicationConfiguration `json:"replicationConfigurations"`
	/*
	  This attribute denotes whether the new mount targets in the file server should be included in the replication policy. By default it is always `true`. If specific mount targets are selected then you should mark this field as `false`. This is applicable for `SMART_DR`.
	*/
	ShouldIncludeNewMountTargets *bool `json:"shouldIncludeNewMountTargets,omitempty"`
	/*
	  Flag to maintain deleted files or folders on the target that are deleted on the source. This is applicable for `DATA_SYNC`.
	*/
	ShouldKeepDeletedFiles *bool `json:"shouldKeepDeletedFiles,omitempty"`

	Status *PolicyStatus `json:"status,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	Type *ReplicationPolicyType `json:"type"`
}

func (p *ReplicationPolicy) MarshalJSON() ([]byte, error) {
	type ReplicationPolicyProxy ReplicationPolicy
	return json.Marshal(struct {
		*ReplicationPolicyProxy
		Name                      *string                    `json:"name,omitempty"`
		ReplicationConfigurations []ReplicationConfiguration `json:"replicationConfigurations,omitempty"`
		Type                      *ReplicationPolicyType     `json:"type,omitempty"`
	}{
		ReplicationPolicyProxy:    (*ReplicationPolicyProxy)(p),
		Name:                      p.Name,
		ReplicationConfigurations: p.ReplicationConfigurations,
		Type:                      p.Type,
	})
}

func NewReplicationPolicy() *ReplicationPolicy {
	p := new(ReplicationPolicy)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.ReplicationPolicy"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	p.IsReverse = new(bool)
	*p.IsReverse = false

	return p
}

/*
Replication policy type value.
*/
type ReplicationPolicyType int

const (
	REPLICATIONPOLICYTYPE_UNKNOWN   ReplicationPolicyType = 0
	REPLICATIONPOLICYTYPE_REDACTED  ReplicationPolicyType = 1
	REPLICATIONPOLICYTYPE_SMART_DR  ReplicationPolicyType = 2
	REPLICATIONPOLICYTYPE_DATA_SYNC ReplicationPolicyType = 3
	REPLICATIONPOLICYTYPE_VDI_SYNC  ReplicationPolicyType = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ReplicationPolicyType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SMART_DR",
		"DATA_SYNC",
		"VDI_SYNC",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ReplicationPolicyType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SMART_DR",
		"DATA_SYNC",
		"VDI_SYNC",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ReplicationPolicyType) index(name string) ReplicationPolicyType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SMART_DR",
		"DATA_SYNC",
		"VDI_SYNC",
	}
	for idx := range names {
		if names[idx] == name {
			return ReplicationPolicyType(idx)
		}
	}
	return REPLICATIONPOLICYTYPE_UNKNOWN
}

func (e *ReplicationPolicyType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ReplicationPolicyType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ReplicationPolicyType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ReplicationPolicyType) Ref() *ReplicationPolicyType {
	return &e
}

/*
This contains the schedule for all the replication and snapshot jobs. This includes the policy’s the start time and replication frequency information.
*/
type ReplicationSchedule struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The frequency at which data should be copied from source to target path. RPO (Integer values are multiplied by schedule type to get the next scheduled replication time).
	*/
	Frequency *int `json:"frequency"`

	ScheduleInterval *ScheduleInterval `json:"scheduleInterval"`
	/*
	  Time of the day, the policy will be started. If the value not set, the policy will be applicable immediately.
	*/
	StartTime *time.Time `json:"startTime,omitempty"`
}

func (p *ReplicationSchedule) MarshalJSON() ([]byte, error) {
	type ReplicationScheduleProxy ReplicationSchedule
	return json.Marshal(struct {
		*ReplicationScheduleProxy
		Frequency        *int              `json:"frequency,omitempty"`
		ScheduleInterval *ScheduleInterval `json:"scheduleInterval,omitempty"`
	}{
		ReplicationScheduleProxy: (*ReplicationScheduleProxy)(p),
		Frequency:                p.Frequency,
		ScheduleInterval:         p.ScheduleInterval,
	})
}

func NewReplicationSchedule() *ReplicationSchedule {
	p := new(ReplicationSchedule)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.ReplicationSchedule"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Replication summary enum values.
*/
type ReplicationSummary int

const (
	REPLICATIONSUMMARY_UNKNOWN           ReplicationSummary = 0
	REPLICATIONSUMMARY_REDACTED          ReplicationSummary = 1
	REPLICATIONSUMMARY_SCHEDULE_MET      ReplicationSummary = 2
	REPLICATIONSUMMARY_SCHEDULE_MISSED   ReplicationSummary = 3
	REPLICATIONSUMMARY_SCHEDULE_PAUSED   ReplicationSummary = 4
	REPLICATIONSUMMARY_REPLICATION_ERROR ReplicationSummary = 5
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ReplicationSummary) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SCHEDULE_MET",
		"SCHEDULE_MISSED",
		"SCHEDULE_PAUSED",
		"REPLICATION_ERROR",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ReplicationSummary) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SCHEDULE_MET",
		"SCHEDULE_MISSED",
		"SCHEDULE_PAUSED",
		"REPLICATION_ERROR",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ReplicationSummary) index(name string) ReplicationSummary {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SCHEDULE_MET",
		"SCHEDULE_MISSED",
		"SCHEDULE_PAUSED",
		"REPLICATION_ERROR",
	}
	for idx := range names {
		if names[idx] == name {
			return ReplicationSummary(idx)
		}
	}
	return REPLICATIONSUMMARY_UNKNOWN
}

func (e *ReplicationSummary) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ReplicationSummary:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ReplicationSummary) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ReplicationSummary) Ref() *ReplicationSummary {
	return &e
}

/*
REST response for all response codes in API path /files/v4.0/config/file-servers/{fileServerExtId}/mount-targets/{mountTargetExtId}/snapshots/{extId}/$actions/restore Post operation
*/
type RestoreMountTargetApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfRestoreMountTargetApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewRestoreMountTargetApiResponse() *RestoreMountTargetApiResponse {
	p := new(RestoreMountTargetApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.RestoreMountTargetApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *RestoreMountTargetApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *RestoreMountTargetApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfRestoreMountTargetApiResponseData()
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
REST response for all response codes in API path /files/v4.0/config/file-servers/{extId}/$actions/revise-dns-records Post operation
*/
type ReviseDnsRecordsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfReviseDnsRecordsApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewReviseDnsRecordsApiResponse() *ReviseDnsRecordsApiResponse {
	p := new(ReviseDnsRecordsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.ReviseDnsRecordsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ReviseDnsRecordsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ReviseDnsRecordsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfReviseDnsRecordsApiResponseData()
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
Access type of the mount target.
*/
type SMBAccessType int

const (
	SMBACCESSTYPE_UNKNOWN      SMBAccessType = 0
	SMBACCESSTYPE_REDACTED     SMBAccessType = 1
	SMBACCESSTYPE_READ         SMBAccessType = 2
	SMBACCESSTYPE_CHANGE       SMBAccessType = 3
	SMBACCESSTYPE_FULL_CONTROL SMBAccessType = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *SMBAccessType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"READ",
		"CHANGE",
		"FULL_CONTROL",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e SMBAccessType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"READ",
		"CHANGE",
		"FULL_CONTROL",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *SMBAccessType) index(name string) SMBAccessType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"READ",
		"CHANGE",
		"FULL_CONTROL",
	}
	for idx := range names {
		if names[idx] == name {
			return SMBAccessType(idx)
		}
	}
	return SMBACCESSTYPE_UNKNOWN
}

func (e *SMBAccessType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for SMBAccessType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *SMBAccessType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e SMBAccessType) Ref() *SMBAccessType {
	return &e
}

/*
Access control entry(ACE) for SMB share.
*/
type SMBShareACE struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	AccessType *SMBAccessType `json:"accessType"`

	PermissionType *PermissionType `json:"permissionType"`
	/*
	  SID of user or group name. This is a read-only field.
	*/
	Sid *string `json:"sid,omitempty"`
	/*
	  User or group name for share permission.
	*/
	UserOrGroupName *string `json:"userOrGroupName"`
}

func (p *SMBShareACE) MarshalJSON() ([]byte, error) {
	type SMBShareACEProxy SMBShareACE
	return json.Marshal(struct {
		*SMBShareACEProxy
		AccessType      *SMBAccessType  `json:"accessType,omitempty"`
		PermissionType  *PermissionType `json:"permissionType,omitempty"`
		UserOrGroupName *string         `json:"userOrGroupName,omitempty"`
	}{
		SMBShareACEProxy: (*SMBShareACEProxy)(p),
		AccessType:       p.AccessType,
		PermissionType:   p.PermissionType,
		UserOrGroupName:  p.UserOrGroupName,
	})
}

func NewSMBShareACE() *SMBShareACE {
	p := new(SMBShareACE)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.SMBShareACE"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Mount target for migration source.
*/
type ScanSourceMountTargetSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Name of the source mount target. Only the mount target name is stored in the database after discovery, but to scan you can specify the mount target name with path, in this case name including path is stored in DB.
	*/
	Path *string `json:"path"`

	Protocol *MountTargetProtocolType `json:"protocol"`
}

func (p *ScanSourceMountTargetSpec) MarshalJSON() ([]byte, error) {
	type ScanSourceMountTargetSpecProxy ScanSourceMountTargetSpec
	return json.Marshal(struct {
		*ScanSourceMountTargetSpecProxy
		Path     *string                  `json:"path,omitempty"`
		Protocol *MountTargetProtocolType `json:"protocol,omitempty"`
	}{
		ScanSourceMountTargetSpecProxy: (*ScanSourceMountTargetSpecProxy)(p),
		Path:                           p.Path,
		Protocol:                       p.Protocol,
	})
}

func NewScanSourceMountTargetSpec() *ScanSourceMountTargetSpec {
	p := new(ScanSourceMountTargetSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.ScanSourceMountTargetSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Status of source mount target path scanning.
*/
type ScanStatus int

const (
	SCANSTATUS_UNKNOWN     ScanStatus = 0
	SCANSTATUS_REDACTED    ScanStatus = 1
	SCANSTATUS_NOT_STARTED ScanStatus = 2
	SCANSTATUS_RUNNING     ScanStatus = 3
	SCANSTATUS_SUCCEEDED   ScanStatus = 4
	SCANSTATUS_FAILED      ScanStatus = 5
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ScanStatus) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NOT_STARTED",
		"RUNNING",
		"SUCCEEDED",
		"FAILED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ScanStatus) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NOT_STARTED",
		"RUNNING",
		"SUCCEEDED",
		"FAILED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ScanStatus) index(name string) ScanStatus {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NOT_STARTED",
		"RUNNING",
		"SUCCEEDED",
		"FAILED",
	}
	for idx := range names {
		if names[idx] == name {
			return ScanStatus(idx)
		}
	}
	return SCANSTATUS_UNKNOWN
}

func (e *ScanStatus) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ScanStatus:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ScanStatus) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ScanStatus) Ref() *ScanStatus {
	return &e
}

/*
Schedule for each day of week.
*/
type ScheduleDetails struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The time difference between the start and the end time of schedule in minutes.
	*/
	DurationMinutes *int64 `json:"durationMinutes"`
	/*
	  Start time hour of the schedule when tiering should start. For Ex. 15 (this means schedule will start at 3PM on that specific day.)
	*/
	StartHours *int64 `json:"startHours"`
	/*
	  Start time minute of the schedule.
	*/
	StartMinutes *int64 `json:"startMinutes"`
}

func (p *ScheduleDetails) MarshalJSON() ([]byte, error) {
	type ScheduleDetailsProxy ScheduleDetails
	return json.Marshal(struct {
		*ScheduleDetailsProxy
		DurationMinutes *int64 `json:"durationMinutes,omitempty"`
		StartHours      *int64 `json:"startHours,omitempty"`
		StartMinutes    *int64 `json:"startMinutes,omitempty"`
	}{
		ScheduleDetailsProxy: (*ScheduleDetailsProxy)(p),
		DurationMinutes:      p.DurationMinutes,
		StartHours:           p.StartHours,
		StartMinutes:         p.StartMinutes,
	})
}

func NewScheduleDetails() *ScheduleDetails {
	p := new(ScheduleDetails)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.ScheduleDetails"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Schedule interval enum values.
*/
type ScheduleInterval int

const (
	SCHEDULEINTERVAL_UNKNOWN  ScheduleInterval = 0
	SCHEDULEINTERVAL_REDACTED ScheduleInterval = 1
	SCHEDULEINTERVAL_YEAR     ScheduleInterval = 2
	SCHEDULEINTERVAL_MONTH    ScheduleInterval = 3
	SCHEDULEINTERVAL_WEEK     ScheduleInterval = 4
	SCHEDULEINTERVAL_DAY      ScheduleInterval = 5
	SCHEDULEINTERVAL_HOUR     ScheduleInterval = 6
	SCHEDULEINTERVAL_MINUTE   ScheduleInterval = 7
	SCHEDULEINTERVAL_SECOND   ScheduleInterval = 8
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ScheduleInterval) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"YEAR",
		"MONTH",
		"WEEK",
		"DAY",
		"HOUR",
		"MINUTE",
		"SECOND",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ScheduleInterval) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"YEAR",
		"MONTH",
		"WEEK",
		"DAY",
		"HOUR",
		"MINUTE",
		"SECOND",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ScheduleInterval) index(name string) ScheduleInterval {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"YEAR",
		"MONTH",
		"WEEK",
		"DAY",
		"HOUR",
		"MINUTE",
		"SECOND",
	}
	for idx := range names {
		if names[idx] == name {
			return ScheduleInterval(idx)
		}
	}
	return SCHEDULEINTERVAL_UNKNOWN
}

func (e *ScheduleInterval) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ScheduleInterval:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ScheduleInterval) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ScheduleInterval) Ref() *ScheduleInterval {
	return &e
}

/*
Type of the server, PRIMARY server type is allowed. This field is only applicable to partner server of type NOTIFICATION.
*/
type ServerType int

const (
	SERVERTYPE_UNKNOWN   ServerType = 0
	SERVERTYPE_REDACTED  ServerType = 1
	SERVERTYPE_PRIMARY   ServerType = 2
	SERVERTYPE_SECONDARY ServerType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ServerType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PRIMARY",
		"SECONDARY",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ServerType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PRIMARY",
		"SECONDARY",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ServerType) index(name string) ServerType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PRIMARY",
		"SECONDARY",
	}
	for idx := range names {
		if names[idx] == name {
			return ServerType(idx)
		}
	}
	return SERVERTYPE_UNKNOWN
}

func (e *ServerType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ServerType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ServerType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ServerType) Ref() *ServerType {
	return &e
}

/*
SMB protocol properties for mount target.
*/
type SmbProtocolProperties struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Flag to enable access based enumeration.
	*/
	IsAccessBasedEnumerationEnabled *bool `json:"isAccessBasedEnumerationEnabled,omitempty"`
	/*
	  Flag to enable continuous availability feature for SMB mount targets.
	*/
	IsCaEnabled *bool `json:"isCaEnabled,omitempty"`
	/*
	  Flag to enable SMB3 encryption.
	*/
	IsSmb3EncryptionEnabled *bool `json:"isSmb3EncryptionEnabled,omitempty"`
	/*
	  Access control list(ACL) for SMB share.
	*/
	ShareACL []SMBShareACE `json:"shareACL,omitempty"`
}

func NewSmbProtocolProperties() *SmbProtocolProperties {
	p := new(SmbProtocolProperties)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.SmbProtocolProperties"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	p.IsAccessBasedEnumerationEnabled = new(bool)
	*p.IsAccessBasedEnumerationEnabled = false
	p.IsCaEnabled = new(bool)
	*p.IsCaEnabled = false
	p.IsSmb3EncryptionEnabled = new(bool)
	*p.IsSmb3EncryptionEnabled = false

	return p
}

/*
Mount target snapshot model.
*/
type Snapshot struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Mount target snapshot created time. This is a read-only field.
	*/
	CreateTime *time.Time `json:"createTime,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Mount target snapshot name.
	*/
	Name *string `json:"name"`
	/*
	  Mount target snapshot reclaimable space in bytes. This is a read-only field.
	*/
	ReclaimableSpaceBytes *int64 `json:"reclaimableSpaceBytes,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  Mount target snapshot total space in bytes. This is a read-only field.
	*/
	TotalSpaceBytes *int64 `json:"totalSpaceBytes,omitempty"`

	Type *SnapshotType `json:"type,omitempty"`
}

func (p *Snapshot) MarshalJSON() ([]byte, error) {
	type SnapshotProxy Snapshot
	return json.Marshal(struct {
		*SnapshotProxy
		Name *string `json:"name,omitempty"`
	}{
		SnapshotProxy: (*SnapshotProxy)(p),
		Name:          p.Name,
	})
}

func NewSnapshot() *Snapshot {
	p := new(Snapshot)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.Snapshot"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	p.ReclaimableSpaceBytes = new(int64)
	*p.ReclaimableSpaceBytes = 0
	p.TotalSpaceBytes = new(int64)
	*p.TotalSpaceBytes = 0

	return p
}

/*
Snapshot change content details.
*/
type SnapshotChangedContent struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Reference snapshot associated with changed files metdata.
	*/
	BaseSnapshotExtId *string `json:"baseSnapshotExtId,omitempty"`
	/*
	  List of changed content. This is a read-only field.
	*/
	ChangedContents []ChangedContent `json:"changedContents,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  If this is set to false, this is the last page of the output. This is a read-only field.
	*/
	HasMoreContent *bool `json:"hasMoreContent,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Current snapshot associated with changed files metdata.
	*/
	SnapshotExtId *string `json:"snapshotExtId"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *SnapshotChangedContent) MarshalJSON() ([]byte, error) {
	type SnapshotChangedContentProxy SnapshotChangedContent
	return json.Marshal(struct {
		*SnapshotChangedContentProxy
		SnapshotExtId *string `json:"snapshotExtId,omitempty"`
	}{
		SnapshotChangedContentProxy: (*SnapshotChangedContentProxy)(p),
		SnapshotExtId:               p.SnapshotExtId,
	})
}

func NewSnapshotChangedContent() *SnapshotChangedContent {
	p := new(SnapshotChangedContent)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.SnapshotChangedContent"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Snapshot schedule model
*/
type SnapshotSchedule struct {
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
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  The maximum number of snapshots to be kept locally.
	*/
	MaxRetentionCount *int `json:"maxRetentionCount"`
	/*

	 */
	ScheduleItemDiscriminator_ *string `json:"$scheduleItemDiscriminator,omitempty"`
	/*
	  Snapshot schedule details for different type of schedules.
	*/
	Schedule *OneOfSnapshotScheduleSchedule `json:"schedule"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	Type *SnapshotScheduleType `json:"type"`
}

func (p *SnapshotSchedule) MarshalJSON() ([]byte, error) {
	type SnapshotScheduleProxy SnapshotSchedule
	return json.Marshal(struct {
		*SnapshotScheduleProxy
		MaxRetentionCount *int                           `json:"maxRetentionCount,omitempty"`
		Schedule          *OneOfSnapshotScheduleSchedule `json:"schedule,omitempty"`
		Type              *SnapshotScheduleType          `json:"type,omitempty"`
	}{
		SnapshotScheduleProxy: (*SnapshotScheduleProxy)(p),
		MaxRetentionCount:     p.MaxRetentionCount,
		Schedule:              p.Schedule,
		Type:                  p.Type,
	})
}

func NewSnapshotSchedule() *SnapshotSchedule {
	p := new(SnapshotSchedule)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.SnapshotSchedule"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *SnapshotSchedule) GetSchedule() interface{} {
	if nil == p.Schedule {
		return nil
	}
	return p.Schedule.GetValue()
}

func (p *SnapshotSchedule) SetSchedule(v interface{}) error {
	if nil == p.Schedule {
		p.Schedule = NewOneOfSnapshotScheduleSchedule()
	}
	e := p.Schedule.SetValue(v)
	if nil == e {
		if nil == p.ScheduleItemDiscriminator_ {
			p.ScheduleItemDiscriminator_ = new(string)
		}
		*p.ScheduleItemDiscriminator_ = *p.Schedule.Discriminator
	}
	return e
}

/*
Snapshot schedule type.
If the user passes this value as WEEKLY, then daysOfWeek becomes a required field.
If the user passes this value as MONTHLY, then daysOfMonth becomes a required field.
*/
type SnapshotScheduleType int

const (
	SNAPSHOTSCHEDULETYPE_UNKNOWN  SnapshotScheduleType = 0
	SNAPSHOTSCHEDULETYPE_REDACTED SnapshotScheduleType = 1
	SNAPSHOTSCHEDULETYPE_HOURLY   SnapshotScheduleType = 2
	SNAPSHOTSCHEDULETYPE_DAILY    SnapshotScheduleType = 3
	SNAPSHOTSCHEDULETYPE_WEEKLY   SnapshotScheduleType = 4
	SNAPSHOTSCHEDULETYPE_MONTHLY  SnapshotScheduleType = 5
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *SnapshotScheduleType) name(index int) string {
	names := [...]string{
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

// Returns the name of the enum
func (e SnapshotScheduleType) GetName() string {
	index := int(e)
	names := [...]string{
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

// Returns the enum type given a string value
func (e *SnapshotScheduleType) index(name string) SnapshotScheduleType {
	names := [...]string{
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

/*
Mount target snapshot type enum values.
*/
type SnapshotType int

const (
	SNAPSHOTTYPE_UNKNOWN              SnapshotType = 0
	SNAPSHOTTYPE_REDACTED             SnapshotType = 1
	SNAPSHOTTYPE_SSR_SNAPSHOT         SnapshotType = 2
	SNAPSHOTTYPE_USER_SNAPSHOT        SnapshotType = 3
	SNAPSHOTTYPE_REPLICATION_SNAPSHOT SnapshotType = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *SnapshotType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SSR_SNAPSHOT",
		"USER_SNAPSHOT",
		"REPLICATION_SNAPSHOT",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e SnapshotType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SSR_SNAPSHOT",
		"USER_SNAPSHOT",
		"REPLICATION_SNAPSHOT",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *SnapshotType) index(name string) SnapshotType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SSR_SNAPSHOT",
		"USER_SNAPSHOT",
		"REPLICATION_SNAPSHOT",
	}
	for idx := range names {
		if names[idx] == name {
			return SnapshotType(idx)
		}
	}
	return SNAPSHOTTYPE_UNKNOWN
}

func (e *SnapshotType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for SnapshotType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *SnapshotType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e SnapshotType) Ref() *SnapshotType {
	return &e
}

/*
Squash type of the mount target.
*/
type SquashType int

const (
	SQUASHTYPE_UNKNOWN     SquashType = 0
	SQUASHTYPE_REDACTED    SquashType = 1
	SQUASHTYPE_ROOT_SQUASH SquashType = 2
	SQUASHTYPE_ALL_SQUASH  SquashType = 3
	SQUASHTYPE_NONE        SquashType = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *SquashType) name(index int) string {
	names := [...]string{
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

// Returns the name of the enum
func (e SquashType) GetName() string {
	index := int(e)
	names := [...]string{
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

// Returns the enum type given a string value
func (e *SquashType) index(name string) SquashType {
	names := [...]string{
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

/*
Status type.
*/
type StatusType int

const (
	STATUSTYPE_UNKNOWN     StatusType = 0
	STATUSTYPE_REDACTED    StatusType = 1
	STATUSTYPE_AVAILABLE   StatusType = 2
	STATUSTYPE_UNAVAILABLE StatusType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *StatusType) name(index int) string {
	names := [...]string{
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

// Returns the name of the enum
func (e StatusType) GetName() string {
	index := int(e)
	names := [...]string{
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

// Returns the enum type given a string value
func (e *StatusType) index(name string) StatusType {
	names := [...]string{
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

/*
Data synchronization path entity validation response.
*/
type SynchronizationPathEntityResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Data synchronization path error message.
	*/
	ErrorMessage *string `json:"errorMessage,omitempty"`
	/*
	  Data synchronization path
	*/
	SyncPath *string `json:"syncPath,omitempty"`
}

func NewSynchronizationPathEntityResponse() *SynchronizationPathEntityResponse {
	p := new(SynchronizationPathEntityResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.SynchronizationPathEntityResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
It denotes how the TLS client should handle the certificate status checks, such as checking for certificate revocation.
*/
type TLSReqCertStatus int

const (
	TLSREQCERTSTATUS_UNKNOWN  TLSReqCertStatus = 0
	TLSREQCERTSTATUS_REDACTED TLSReqCertStatus = 1
	TLSREQCERTSTATUS_ALLOW    TLSReqCertStatus = 2
	TLSREQCERTSTATUS_DEMAND   TLSReqCertStatus = 3
	TLSREQCERTSTATUS_NEVER    TLSReqCertStatus = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *TLSReqCertStatus) name(index int) string {
	names := [...]string{
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

// Returns the name of the enum
func (e TLSReqCertStatus) GetName() string {
	index := int(e)
	names := [...]string{
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

// Returns the enum type given a string value
func (e *TLSReqCertStatus) index(name string) TLSReqCertStatus {
	names := [...]string{
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

/*
REST response for all response codes in API path /files/v4.0/config/file-servers/{fileServerExtId}/anti-virus-servers/{extId}/$actions/test-connection Post operation
*/
type TestConnectionAntivirusServerApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfTestConnectionAntivirusServerApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewTestConnectionAntivirusServerApiResponse() *TestConnectionAntivirusServerApiResponse {
	p := new(TestConnectionAntivirusServerApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.TestConnectionAntivirusServerApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
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

/*
Request body specification for tiering configuration APIs.
*/
type TierConfiguration struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Cool off period in seconds for tiering. Files older than cool off period time will be considered for tiering.
	*/
	CooloffPeriodSeconds *int64 `json:"cooloffPeriodSeconds,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Capacity threshold in percentage for tiering.
	*/
	MemoryThresholdPercent *int64 `json:"memoryThresholdPercent,omitempty"`
	/*
	  Minimum file size for tiering. Files size greater than this will be considered for tiering.
	*/
	MinimumFileSizeBytes *int64 `json:"minimumFileSizeBytes,omitempty"`
	/*
	  Mount target list which should be included in tiering configuration.
	*/
	MountTargetExtIds []string `json:"mountTargetExtIds,omitempty"`

	MountTargetsEnablementType *MountTargetsEnablementType `json:"mountTargetsEnablementType,omitempty"`
	/*
	  Tiering schedule for on-prem tiering. Auto tiering can happen at this specified time.
	*/
	Schedule []DaySchedule `json:"schedule,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewTierConfiguration() *TierConfiguration {
	p := new(TierConfiguration)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.TierConfiguration"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	p.MinimumFileSizeBytes = new(int64)
	*p.MinimumFileSizeBytes = 65536

	return p
}

/*
REST response for all response codes in API path /files/v4.0/config/file-servers/{extId}/$actions/tier-data Post operation
*/
type TierDataApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfTierDataApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewTierDataApiResponse() *TierDataApiResponse {
	p := new(TierDataApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.TierDataApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
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

/*
Model for tier data.
*/
type TierDataSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Duration (in seconds) for which the manual tiering action should run.
	If not provided tiering will run until the memory has reached the threshold capacity value.
	*/
	DurationSeconds *int64 `json:"durationSeconds,omitempty"`

	FilesSpec *FilesSpec `json:"filesSpec,omitempty"`
}

func NewTierDataSpec() *TierDataSpec {
	p := new(TierDataSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.TierDataSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	p.DurationSeconds = new(int64)
	*p.DurationSeconds = 0

	return p
}

/*
Unified namespace is a group of file servers across multiple Nutanix clusters and present unified access to the clients through the namespace of the core member of the group. One of the file servers can be designated as a core member. This is primarily used for expanding capacity across Nutanix Clusters.
*/
type UnifiedNamespace struct {
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
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  This contains a list of unified namespace member configurations. Member configuration represents member file server and a flag to indicate whether it is hosting the namespace.
	*/
	NamespaceMemberConfigs []NamespaceMemberConfig `json:"namespaceMemberConfigs"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *UnifiedNamespace) MarshalJSON() ([]byte, error) {
	type UnifiedNamespaceProxy UnifiedNamespace
	return json.Marshal(struct {
		*UnifiedNamespaceProxy
		NamespaceMemberConfigs []NamespaceMemberConfig `json:"namespaceMemberConfigs,omitempty"`
	}{
		UnifiedNamespaceProxy:  (*UnifiedNamespaceProxy)(p),
		NamespaceMemberConfigs: p.NamespaceMemberConfigs,
	})
}

func NewUnifiedNamespace() *UnifiedNamespace {
	p := new(UnifiedNamespace)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.UnifiedNamespace"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /files/v4.0/config/file-servers/{fileServerExtId}/anti-virus-servers/{extId} Put operation
*/
type UpdateAntivirusServerApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateAntivirusServerApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewUpdateAntivirusServerApiResponse() *UpdateAntivirusServerApiResponse {
	p := new(UpdateAntivirusServerApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.UpdateAntivirusServerApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
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

/*
REST response for all response codes in API path /files/v4.0/config/file-servers/{fileServerExtId}/email-config Put operation
*/
type UpdateEmailConfigApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateEmailConfigApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewUpdateEmailConfigApiResponse() *UpdateEmailConfigApiResponse {
	p := new(UpdateEmailConfigApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.UpdateEmailConfigApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdateEmailConfigApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdateEmailConfigApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdateEmailConfigApiResponseData()
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
REST response for all response codes in API path /files/v4.0/config/file-servers/{fileServerExtId}/mount-targets/{extId} Put operation
*/
type UpdateMountTargetApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateMountTargetApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewUpdateMountTargetApiResponse() *UpdateMountTargetApiResponse {
	p := new(UpdateMountTargetApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.UpdateMountTargetApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
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

/*
REST response for all response codes in API path /files/v4.0/config/file-servers/{fileServerExtId}/notification-policies/{extId} Put operation
*/
type UpdateNotificationPolicyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateNotificationPolicyApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewUpdateNotificationPolicyApiResponse() *UpdateNotificationPolicyApiResponse {
	p := new(UpdateNotificationPolicyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.UpdateNotificationPolicyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdateNotificationPolicyApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdateNotificationPolicyApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdateNotificationPolicyApiResponseData()
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
REST response for all response codes in API path /files/v4.0/config/file-servers/{fileServerExtId}/object-store-profiles/{extId} Put operation
*/
type UpdateObjectStoreProfileApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateObjectStoreProfileApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewUpdateObjectStoreProfileApiResponse() *UpdateObjectStoreProfileApiResponse {
	p := new(UpdateObjectStoreProfileApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.UpdateObjectStoreProfileApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdateObjectStoreProfileApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdateObjectStoreProfileApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdateObjectStoreProfileApiResponseData()
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
REST response for all response codes in API path /files/v4.0/config/file-servers/{fileServerExtId}/partner-servers/{extId} Put operation
*/
type UpdatePartnerServerApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdatePartnerServerApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewUpdatePartnerServerApiResponse() *UpdatePartnerServerApiResponse {
	p := new(UpdatePartnerServerApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.UpdatePartnerServerApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdatePartnerServerApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdatePartnerServerApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdatePartnerServerApiResponseData()
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
REST response for all response codes in API path /files/v4.0/config/file-servers/{fileServerExtId}/mount-targets/{mountTargetExtId}/quota-policies/{extId} Put operation
*/
type UpdateQuotaPolicyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateQuotaPolicyApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewUpdateQuotaPolicyApiResponse() *UpdateQuotaPolicyApiResponse {
	p := new(UpdateQuotaPolicyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.UpdateQuotaPolicyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
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

/*
REST response for all response codes in API path /files/v4.0/config/file-servers/{fileServerExtId}/ransomware-configs/{extId} Put operation
*/
type UpdateRansomwareConfigApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateRansomwareConfigApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewUpdateRansomwareConfigApiResponse() *UpdateRansomwareConfigApiResponse {
	p := new(UpdateRansomwareConfigApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.UpdateRansomwareConfigApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
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

/*
REST response for all response codes in API path /files/v4.0/config/replication-policies/{extId} Put operation
*/
type UpdateReplicationPolicyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateReplicationPolicyApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewUpdateReplicationPolicyApiResponse() *UpdateReplicationPolicyApiResponse {
	p := new(UpdateReplicationPolicyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.UpdateReplicationPolicyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdateReplicationPolicyApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdateReplicationPolicyApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdateReplicationPolicyApiResponseData()
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
REST response for all response codes in API path /files/v4.0/config/file-servers/{fileServerExtId}/snapshot-schedules/{extId} Put operation
*/
type UpdateSnapshotScheduleApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateSnapshotScheduleApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewUpdateSnapshotScheduleApiResponse() *UpdateSnapshotScheduleApiResponse {
	p := new(UpdateSnapshotScheduleApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.UpdateSnapshotScheduleApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdateSnapshotScheduleApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdateSnapshotScheduleApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdateSnapshotScheduleApiResponseData()
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
REST response for all response codes in API path /files/v4.0/config/file-servers/{fileServerExtId}/tier-configurations/{extId} Put operation
*/
type UpdateTieringConfigurationApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateTieringConfigurationApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewUpdateTieringConfigurationApiResponse() *UpdateTieringConfigurationApiResponse {
	p := new(UpdateTieringConfigurationApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.UpdateTieringConfigurationApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdateTieringConfigurationApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdateTieringConfigurationApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdateTieringConfigurationApiResponseData()
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
REST response for all response codes in API path /files/v4.0/config/unified-namespaces/{extId} Put operation
*/
type UpdateUnifiedNamespaceApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateUnifiedNamespaceApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewUpdateUnifiedNamespaceApiResponse() *UpdateUnifiedNamespaceApiResponse {
	p := new(UpdateUnifiedNamespaceApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.UpdateUnifiedNamespaceApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdateUnifiedNamespaceApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdateUnifiedNamespaceApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdateUnifiedNamespaceApiResponseData()
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
REST response for all response codes in API path /files/v4.0/config/file-servers/{fileServerExtId}/replication-policies/{replicationPolicyExtId}/vdi-user-sessions/{extId} Put operation
*/
type UpdateVdiUserSessionApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateVdiUserSessionApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewUpdateVdiUserSessionApiResponse() *UpdateVdiUserSessionApiResponse {
	p := new(UpdateVdiUserSessionApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.UpdateVdiUserSessionApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdateVdiUserSessionApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdateVdiUserSessionApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdateVdiUserSessionApiResponseData()
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
REST response for all response codes in API path /files/v4.0/config/file-servers/{fileServerExtId}/virus-scan-policies/{extId} Put operation
*/
type UpdateVirusScanPolicyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateVirusScanPolicyApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewUpdateVirusScanPolicyApiResponse() *UpdateVirusScanPolicyApiResponse {
	p := new(UpdateVirusScanPolicyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.UpdateVirusScanPolicyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
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

/*
REST response for all response codes in API path /files/v4.0/config/file-servers/{fileServerExtId}/$actions/upload-user-mappings Post operation
*/
type UploadUserMappingsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUploadUserMappingsApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewUploadUserMappingsApiResponse() *UploadUserMappingsApiResponse {
	p := new(UploadUserMappingsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.UploadUserMappingsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UploadUserMappingsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UploadUserMappingsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUploadUserMappingsApiResponseData()
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
VDI synchronization event contains detailed user information. For example - login details like login and logoff time, login status, and more.
*/
type VDIUserLoginEventLog struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The reason why the user cannot login and is considered impacted. This may be due to unsynchronized data, user name conflicts, or if the user is active on other sites.
	*/
	FailureReason *string `json:"failureReason,omitempty"`
	/*
	  External identifier of the file server that user attempted to login.
	*/
	FileServerExtId *string `json:"fileServerExtId,omitempty"`
	/*
	  VDI synchronization user login time in ISO format. This is a read-only field.
	*/
	LoginTime *time.Time `json:"loginTime,omitempty"`
	/*
	  VDI synchronization user logoff time in ISO format. This is a read-only field.
	*/
	LogoutTime *time.Time `json:"logoutTime,omitempty"`

	Status *LoginStatus `json:"status,omitempty"`
}

func NewVDIUserLoginEventLog() *VDIUserLoginEventLog {
	p := new(VDIUserLoginEventLog)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.VDIUserLoginEventLog"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
A VDI session for the user that tried to login.
*/
type VdiUserSession struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	CurrentSession *VDIUserLoginEventLog `json:"currentSession,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  File server external identifier that is or will become the owner file server for the specified users depending on the action performed.
	*/
	OwnerFileServerExtId *string `json:"ownerFileServerExtId"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  Domain and username of the user. Internally it is converted to user SID. This is a read-only field.
	*/
	UserName *string `json:"userName,omitempty"`
}

func (p *VdiUserSession) MarshalJSON() ([]byte, error) {
	type VdiUserSessionProxy VdiUserSession
	return json.Marshal(struct {
		*VdiUserSessionProxy
		OwnerFileServerExtId *string `json:"ownerFileServerExtId,omitempty"`
	}{
		VdiUserSessionProxy:  (*VdiUserSessionProxy)(p),
		OwnerFileServerExtId: p.OwnerFileServerExtId,
	})
}

func NewVdiUserSession() *VdiUserSession {
	p := new(VdiUserSession)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.VdiUserSession"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Partner server configuration for the following types of servers
```
  - BACKUP
  - MIGRATION
  - NOTIFICATION
```
*/
type VendorProperties struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Address *import1.IPAddressOrFQDN `json:"address"`

	ConnectionStatus *ConnectionStatus `json:"connectionStatus,omitempty"`
	/*
	  Indicates the vendor specific custom properties of the partner server. It contains key value pairs separated by pipe(|)operator. For example: kafkatopic=1P1R| brokerlist=10.5.119.71:29092,10.5.119.71:39092,10.5.119.71:49092
	*/
	CustomProperties []import1.KVStringPair `json:"customProperties,omitempty"`
	/*
	  Partner server port.
	*/
	Port *int `json:"port"`

	ServerType *ServerType `json:"serverType,omitempty"`
}

func (p *VendorProperties) MarshalJSON() ([]byte, error) {
	type VendorPropertiesProxy VendorProperties
	return json.Marshal(struct {
		*VendorPropertiesProxy
		Address *import1.IPAddressOrFQDN `json:"address,omitempty"`
		Port    *int                     `json:"port,omitempty"`
	}{
		VendorPropertiesProxy: (*VendorPropertiesProxy)(p),
		Address:               p.Address,
		Port:                  p.Port,
	})
}

func NewVendorProperties() *VendorProperties {
	p := new(VendorProperties)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.VendorProperties"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /files/v4.0/config/file-servers/{extId}/$actions/verify-dns-records Post operation
*/
type VerifyDnsRecordsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfVerifyDnsRecordsApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewVerifyDnsRecordsApiResponse() *VerifyDnsRecordsApiResponse {
	p := new(VerifyDnsRecordsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.VerifyDnsRecordsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *VerifyDnsRecordsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *VerifyDnsRecordsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfVerifyDnsRecordsApiResponseData()
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
Antivirus server scan policy information.
*/
type VirusScanPolicy struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Files with extension in the listed ones will be excluded during antivirus scan.
	*/
	ExcludedFileExtensions []string `json:"excludedFileExtensions,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Indicates if the antivirus server is enabled for the current policy.
	*/
	IsAntiVirusEnabled *bool `json:"isAntiVirusEnabled,omitempty"`
	/*
	  Indicates if the file access needs to be blocked for the current policy.
	*/
	IsFileAccessBlocked *bool `json:"isFileAccessBlocked,omitempty"`
	/*
	  Scan file for presence of virus during read.
	*/
	IsScanOnReadEnabled *bool `json:"isScanOnReadEnabled,omitempty"`
	/*
	  Scan file for presence of virus during write.
	*/
	IsScanOnWriteEnabled *bool `json:"isScanOnWriteEnabled,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Maximum file size in bytes for which file scan is allowed.
	*/
	MaxFileSizeThresholdBytes *int64 `json:"maxFileSizeThresholdBytes,omitempty"`
	/*
	  The external identifier of the mount target under which the virus scan policy is being created. This field will be set to null if it is a global policy.
	*/
	MountTargetReference *string `json:"mountTargetReference,omitempty"`
	/*
	  Time interval to complete the scan. If it is not completed within the specified time, the request is timed out and the default action is taken.
	*/
	ScanTimeoutIntervalSecs *int64 `json:"scanTimeoutIntervalSecs,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewVirusScanPolicy() *VirusScanPolicy {
	p := new(VirusScanPolicy)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.VirusScanPolicy"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	p.IsAntiVirusEnabled = new(bool)
	*p.IsAntiVirusEnabled = true
	p.IsFileAccessBlocked = new(bool)
	*p.IsFileAccessBlocked = false
	p.IsScanOnReadEnabled = new(bool)
	*p.IsScanOnReadEnabled = true
	p.IsScanOnWriteEnabled = new(bool)
	*p.IsScanOnWriteEnabled = true

	return p
}

/*
To repeat the schedule on a per `Week` basis.
*/
type WeeklySchedule struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A list of the days of the week on which snapshots should be taken.
	*/
	DaysOfWeek []int `json:"daysOfWeek"`
}

func (p *WeeklySchedule) MarshalJSON() ([]byte, error) {
	type WeeklyScheduleProxy WeeklySchedule
	return json.Marshal(struct {
		*WeeklyScheduleProxy
		DaysOfWeek []int `json:"daysOfWeek,omitempty"`
	}{
		WeeklyScheduleProxy: (*WeeklyScheduleProxy)(p),
		DaysOfWeek:          p.DaysOfWeek,
	})
}

func NewWeeklySchedule() *WeeklySchedule {
	p := new(WeeklySchedule)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.WeeklySchedule"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
WORM share details.
*/
type WormSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Cooloff interval in seconds for WORM share.
	*/
	CooloffIntervalSeconds *int64 `json:"cooloffIntervalSeconds,omitempty"`
	/*
	  WORM compliance for share. This is a read-only field.
	*/
	IsComplianceEnabled *bool `json:"isComplianceEnabled,omitempty"`
	/*
	  Legal hold for WORM share. This is a read-only field.
	*/
	IsLegalHoldEnabled *bool `json:"isLegalHoldEnabled,omitempty"`
	/*
	  Retention period in seconds for WORM share.
	*/
	RetentionPeriodSeconds *int64 `json:"retentionPeriodSeconds,omitempty"`

	WormType *WormType `json:"wormType,omitempty"`
}

func NewWormSpec() *WormSpec {
	p := new(WormSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.config.WormSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	p.CooloffIntervalSeconds = new(int64)
	*p.CooloffIntervalSeconds = 600
	p.IsComplianceEnabled = new(bool)
	*p.IsComplianceEnabled = false
	p.IsLegalHoldEnabled = new(bool)
	*p.IsLegalHoldEnabled = false
	p.RetentionPeriodSeconds = new(int64)
	*p.RetentionPeriodSeconds = 31449600

	return p
}

/*
WORM type for share.
*/
type WormType int

const (
	WORMTYPE_UNKNOWN     WormType = 0
	WORMTYPE_REDACTED    WormType = 1
	WORMTYPE_SHARE_LEVEL WormType = 2
	WORMTYPE_DISABLED    WormType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *WormType) name(index int) string {
	names := [...]string{
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

// Returns the name of the enum
func (e WormType) GetName() string {
	index := int(e)
	names := [...]string{
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

// Returns the enum type given a string value
func (e *WormType) index(name string) WormType {
	names := [...]string{
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

type OneOfGetReplicationPolicyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *ReplicationPolicy     `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfGetReplicationPolicyApiResponseData() *OneOfGetReplicationPolicyApiResponseData {
	p := new(OneOfGetReplicationPolicyApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetReplicationPolicyApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetReplicationPolicyApiResponseData is nil"))
	}
	switch v.(type) {
	case ReplicationPolicy:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(ReplicationPolicy)
		}
		*p.oneOfType0 = v.(ReplicationPolicy)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfGetReplicationPolicyApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetReplicationPolicyApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(ReplicationPolicy)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "files.v4.config.ReplicationPolicy" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(ReplicationPolicy)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetReplicationPolicyApiResponseData"))
}

func (p *OneOfGetReplicationPolicyApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetReplicationPolicyApiResponseData")
}

type OneOfGetVdiUserSessionApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *VdiUserSession        `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfGetVdiUserSessionApiResponseData() *OneOfGetVdiUserSessionApiResponseData {
	p := new(OneOfGetVdiUserSessionApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetVdiUserSessionApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetVdiUserSessionApiResponseData is nil"))
	}
	switch v.(type) {
	case VdiUserSession:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(VdiUserSession)
		}
		*p.oneOfType0 = v.(VdiUserSession)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfGetVdiUserSessionApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetVdiUserSessionApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(VdiUserSession)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "files.v4.config.VdiUserSession" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(VdiUserSession)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetVdiUserSessionApiResponseData"))
}

func (p *OneOfGetVdiUserSessionApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetVdiUserSessionApiResponseData")
}

type OneOfListMountTargetSnapshotsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    []Snapshot             `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfListMountTargetSnapshotsApiResponseData() *OneOfListMountTargetSnapshotsApiResponseData {
	p := new(OneOfListMountTargetSnapshotsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListMountTargetSnapshotsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListMountTargetSnapshotsApiResponseData is nil"))
	}
	switch v.(type) {
	case []Snapshot:
		p.oneOfType0 = v.([]Snapshot)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<files.v4.config.Snapshot>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<files.v4.config.Snapshot>"
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfListMountTargetSnapshotsApiResponseData) GetValue() interface{} {
	if "List<files.v4.config.Snapshot>" == *p.Discriminator {
		return p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfListMountTargetSnapshotsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new([]Snapshot)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "files.v4.config.Snapshot" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<files.v4.config.Snapshot>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<files.v4.config.Snapshot>"
			return nil
		}
	}
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListMountTargetSnapshotsApiResponseData"))
}

func (p *OneOfListMountTargetSnapshotsApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<files.v4.config.Snapshot>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfListMountTargetSnapshotsApiResponseData")
}

type OneOfUploadUserMappingsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    []import4.AppMessage   `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfUploadUserMappingsApiResponseData() *OneOfUploadUserMappingsApiResponseData {
	p := new(OneOfUploadUserMappingsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUploadUserMappingsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUploadUserMappingsApiResponseData is nil"))
	}
	switch v.(type) {
	case []import4.AppMessage:
		p.oneOfType0 = v.([]import4.AppMessage)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<files.v4.error.AppMessage>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<files.v4.error.AppMessage>"
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfUploadUserMappingsApiResponseData) GetValue() interface{} {
	if "List<files.v4.error.AppMessage>" == *p.Discriminator {
		return p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfUploadUserMappingsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new([]import4.AppMessage)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "files.v4.error.AppMessage" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<files.v4.error.AppMessage>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<files.v4.error.AppMessage>"
			return nil
		}
	}
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUploadUserMappingsApiResponseData"))
}

func (p *OneOfUploadUserMappingsApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<files.v4.error.AppMessage>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfUploadUserMappingsApiResponseData")
}

type OneOfGetFileServerApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *FileServer            `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfGetFileServerApiResponseData() *OneOfGetFileServerApiResponseData {
	p := new(OneOfGetFileServerApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetFileServerApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetFileServerApiResponseData is nil"))
	}
	switch v.(type) {
	case FileServer:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(FileServer)
		}
		*p.oneOfType0 = v.(FileServer)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfGetFileServerApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetFileServerApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(FileServer)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "files.v4.config.FileServer" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(FileServer)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetFileServerApiResponseData"))
}

func (p *OneOfGetFileServerApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetFileServerApiResponseData")
}

type OneOfGetVirusScanPolicyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *VirusScanPolicy       `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfGetVirusScanPolicyApiResponseData() *OneOfGetVirusScanPolicyApiResponseData {
	p := new(OneOfGetVirusScanPolicyApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetVirusScanPolicyApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetVirusScanPolicyApiResponseData is nil"))
	}
	switch v.(type) {
	case VirusScanPolicy:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(VirusScanPolicy)
		}
		*p.oneOfType0 = v.(VirusScanPolicy)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfGetVirusScanPolicyApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetVirusScanPolicyApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(VirusScanPolicy)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "files.v4.config.VirusScanPolicy" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(VirusScanPolicy)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetVirusScanPolicyApiResponseData"))
}

func (p *OneOfGetVirusScanPolicyApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetVirusScanPolicyApiResponseData")
}

type OneOfGetMountTargetApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType0    *MountTarget           `json:"-"`
}

func NewOneOfGetMountTargetApiResponseData() *OneOfGetMountTargetApiResponseData {
	p := new(OneOfGetMountTargetApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetMountTargetApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetMountTargetApiResponseData is nil"))
	}
	switch v.(type) {
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case MountTarget:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(MountTarget)
		}
		*p.oneOfType0 = v.(MountTarget)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfGetMountTargetApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfGetMountTargetApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	vOneOfType0 := new(MountTarget)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "files.v4.config.MountTarget" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(MountTarget)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetMountTargetApiResponseData"))
}

func (p *OneOfGetMountTargetApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfGetMountTargetApiResponseData")
}

type OneOfCreateNotificationPolicyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfCreateNotificationPolicyApiResponseData() *OneOfCreateNotificationPolicyApiResponseData {
	p := new(OneOfCreateNotificationPolicyApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateNotificationPolicyApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateNotificationPolicyApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.TaskReference)
		}
		*p.oneOfType0 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfCreateNotificationPolicyApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfCreateNotificationPolicyApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.TaskReference)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateNotificationPolicyApiResponseData"))
}

func (p *OneOfCreateNotificationPolicyApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfCreateNotificationPolicyApiResponseData")
}

type OneOfListInfectedFilesApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    []InfectedFile         `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfListInfectedFilesApiResponseData() *OneOfListInfectedFilesApiResponseData {
	p := new(OneOfListInfectedFilesApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListInfectedFilesApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListInfectedFilesApiResponseData is nil"))
	}
	switch v.(type) {
	case []InfectedFile:
		p.oneOfType0 = v.([]InfectedFile)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<files.v4.config.InfectedFile>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<files.v4.config.InfectedFile>"
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfListInfectedFilesApiResponseData) GetValue() interface{} {
	if "List<files.v4.config.InfectedFile>" == *p.Discriminator {
		return p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfListInfectedFilesApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new([]InfectedFile)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "files.v4.config.InfectedFile" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<files.v4.config.InfectedFile>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<files.v4.config.InfectedFile>"
			return nil
		}
	}
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListInfectedFilesApiResponseData"))
}

func (p *OneOfListInfectedFilesApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<files.v4.config.InfectedFile>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfListInfectedFilesApiResponseData")
}

type OneOfTierDataApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfTierDataApiResponseData() *OneOfTierDataApiResponseData {
	p := new(OneOfTierDataApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfTierDataApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfTierDataApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.TaskReference)
		}
		*p.oneOfType0 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfTierDataApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfTierDataApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.TaskReference)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfTierDataApiResponseData"))
}

func (p *OneOfTierDataApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfTierDataApiResponseData")
}

type OneOfGetEmailConfigApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *EmailConfig           `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfGetEmailConfigApiResponseData() *OneOfGetEmailConfigApiResponseData {
	p := new(OneOfGetEmailConfigApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetEmailConfigApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetEmailConfigApiResponseData is nil"))
	}
	switch v.(type) {
	case EmailConfig:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(EmailConfig)
		}
		*p.oneOfType0 = v.(EmailConfig)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfGetEmailConfigApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetEmailConfigApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(EmailConfig)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "files.v4.config.EmailConfig" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(EmailConfig)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetEmailConfigApiResponseData"))
}

func (p *OneOfGetEmailConfigApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetEmailConfigApiResponseData")
}

type OneOfListSnapshotSchedulesApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    []SnapshotSchedule     `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfListSnapshotSchedulesApiResponseData() *OneOfListSnapshotSchedulesApiResponseData {
	p := new(OneOfListSnapshotSchedulesApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListSnapshotSchedulesApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListSnapshotSchedulesApiResponseData is nil"))
	}
	switch v.(type) {
	case []SnapshotSchedule:
		p.oneOfType0 = v.([]SnapshotSchedule)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<files.v4.config.SnapshotSchedule>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<files.v4.config.SnapshotSchedule>"
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfListSnapshotSchedulesApiResponseData) GetValue() interface{} {
	if "List<files.v4.config.SnapshotSchedule>" == *p.Discriminator {
		return p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfListSnapshotSchedulesApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new([]SnapshotSchedule)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "files.v4.config.SnapshotSchedule" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<files.v4.config.SnapshotSchedule>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<files.v4.config.SnapshotSchedule>"
			return nil
		}
	}
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListSnapshotSchedulesApiResponseData"))
}

func (p *OneOfListSnapshotSchedulesApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<files.v4.config.SnapshotSchedule>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfListSnapshotSchedulesApiResponseData")
}

type OneOfUpdateNotificationPolicyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfUpdateNotificationPolicyApiResponseData() *OneOfUpdateNotificationPolicyApiResponseData {
	p := new(OneOfUpdateNotificationPolicyApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateNotificationPolicyApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateNotificationPolicyApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.TaskReference)
		}
		*p.oneOfType0 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfUpdateNotificationPolicyApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfUpdateNotificationPolicyApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.TaskReference)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateNotificationPolicyApiResponseData"))
}

func (p *OneOfUpdateNotificationPolicyApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateNotificationPolicyApiResponseData")
}

type OneOfListAntivirusServersApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    []AntivirusServer      `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfListAntivirusServersApiResponseData() *OneOfListAntivirusServersApiResponseData {
	p := new(OneOfListAntivirusServersApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListAntivirusServersApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListAntivirusServersApiResponseData is nil"))
	}
	switch v.(type) {
	case []AntivirusServer:
		p.oneOfType0 = v.([]AntivirusServer)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<files.v4.config.AntivirusServer>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<files.v4.config.AntivirusServer>"
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfListAntivirusServersApiResponseData) GetValue() interface{} {
	if "List<files.v4.config.AntivirusServer>" == *p.Discriminator {
		return p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfListAntivirusServersApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new([]AntivirusServer)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "files.v4.config.AntivirusServer" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<files.v4.config.AntivirusServer>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<files.v4.config.AntivirusServer>"
			return nil
		}
	}
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListAntivirusServersApiResponseData"))
}

func (p *OneOfListAntivirusServersApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<files.v4.config.AntivirusServer>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfListAntivirusServersApiResponseData")
}

type OneOfDeleteAntivirusServerApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfDeleteAntivirusServerApiResponseData() *OneOfDeleteAntivirusServerApiResponseData {
	p := new(OneOfDeleteAntivirusServerApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteAntivirusServerApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteAntivirusServerApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.TaskReference)
		}
		*p.oneOfType0 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfDeleteAntivirusServerApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDeleteAntivirusServerApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.TaskReference)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteAntivirusServerApiResponseData"))
}

func (p *OneOfDeleteAntivirusServerApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteAntivirusServerApiResponseData")
}

type OneOfCreateQuotaPolicyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfCreateQuotaPolicyApiResponseData() *OneOfCreateQuotaPolicyApiResponseData {
	p := new(OneOfCreateQuotaPolicyApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateQuotaPolicyApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateQuotaPolicyApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.TaskReference)
		}
		*p.oneOfType0 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.TaskReference)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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

type OneOfRecalculateColdDataApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    []import4.AppMessage   `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfRecalculateColdDataApiResponseData() *OneOfRecalculateColdDataApiResponseData {
	p := new(OneOfRecalculateColdDataApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfRecalculateColdDataApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfRecalculateColdDataApiResponseData is nil"))
	}
	switch v.(type) {
	case []import4.AppMessage:
		p.oneOfType0 = v.([]import4.AppMessage)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<files.v4.error.AppMessage>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<files.v4.error.AppMessage>"
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfRecalculateColdDataApiResponseData) GetValue() interface{} {
	if "List<files.v4.error.AppMessage>" == *p.Discriminator {
		return p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfRecalculateColdDataApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new([]import4.AppMessage)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "files.v4.error.AppMessage" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<files.v4.error.AppMessage>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<files.v4.error.AppMessage>"
			return nil
		}
	}
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfRecalculateColdDataApiResponseData"))
}

func (p *OneOfRecalculateColdDataApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<files.v4.error.AppMessage>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfRecalculateColdDataApiResponseData")
}

type OneOfRecallTieredFilesApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfRecallTieredFilesApiResponseData() *OneOfRecallTieredFilesApiResponseData {
	p := new(OneOfRecallTieredFilesApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfRecallTieredFilesApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfRecallTieredFilesApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.TaskReference)
		}
		*p.oneOfType0 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfRecallTieredFilesApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfRecallTieredFilesApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.TaskReference)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfRecallTieredFilesApiResponseData"))
}

func (p *OneOfRecallTieredFilesApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfRecallTieredFilesApiResponseData")
}

type OneOfComputeSnapshotChangeApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    []import4.AppMessage   `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfComputeSnapshotChangeApiResponseData() *OneOfComputeSnapshotChangeApiResponseData {
	p := new(OneOfComputeSnapshotChangeApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfComputeSnapshotChangeApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfComputeSnapshotChangeApiResponseData is nil"))
	}
	switch v.(type) {
	case []import4.AppMessage:
		p.oneOfType0 = v.([]import4.AppMessage)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<files.v4.error.AppMessage>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<files.v4.error.AppMessage>"
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfComputeSnapshotChangeApiResponseData) GetValue() interface{} {
	if "List<files.v4.error.AppMessage>" == *p.Discriminator {
		return p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfComputeSnapshotChangeApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new([]import4.AppMessage)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "files.v4.error.AppMessage" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<files.v4.error.AppMessage>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<files.v4.error.AppMessage>"
			return nil
		}
	}
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfComputeSnapshotChangeApiResponseData"))
}

func (p *OneOfComputeSnapshotChangeApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<files.v4.error.AppMessage>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfComputeSnapshotChangeApiResponseData")
}

type OneOfGetNotificationPolicyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *NotificationPolicy    `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfGetNotificationPolicyApiResponseData() *OneOfGetNotificationPolicyApiResponseData {
	p := new(OneOfGetNotificationPolicyApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetNotificationPolicyApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetNotificationPolicyApiResponseData is nil"))
	}
	switch v.(type) {
	case NotificationPolicy:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(NotificationPolicy)
		}
		*p.oneOfType0 = v.(NotificationPolicy)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfGetNotificationPolicyApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetNotificationPolicyApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(NotificationPolicy)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "files.v4.config.NotificationPolicy" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(NotificationPolicy)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetNotificationPolicyApiResponseData"))
}

func (p *OneOfGetNotificationPolicyApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetNotificationPolicyApiResponseData")
}

type OneOfListSnapshotChangedContentsApiResponseData struct {
	Discriminator *string                  `json:"-"`
	ObjectType_   *string                  `json:"-"`
	oneOfType0    []SnapshotChangedContent `json:"-"`
	oneOfType400  *import4.ErrorResponse   `json:"-"`
}

func NewOneOfListSnapshotChangedContentsApiResponseData() *OneOfListSnapshotChangedContentsApiResponseData {
	p := new(OneOfListSnapshotChangedContentsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListSnapshotChangedContentsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListSnapshotChangedContentsApiResponseData is nil"))
	}
	switch v.(type) {
	case []SnapshotChangedContent:
		p.oneOfType0 = v.([]SnapshotChangedContent)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<files.v4.config.SnapshotChangedContent>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<files.v4.config.SnapshotChangedContent>"
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfListSnapshotChangedContentsApiResponseData) GetValue() interface{} {
	if "List<files.v4.config.SnapshotChangedContent>" == *p.Discriminator {
		return p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfListSnapshotChangedContentsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new([]SnapshotChangedContent)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "files.v4.config.SnapshotChangedContent" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<files.v4.config.SnapshotChangedContent>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<files.v4.config.SnapshotChangedContent>"
			return nil
		}
	}
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListSnapshotChangedContentsApiResponseData"))
}

func (p *OneOfListSnapshotChangedContentsApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<files.v4.config.SnapshotChangedContent>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfListSnapshotChangedContentsApiResponseData")
}

type OneOfGetRansomwareConfigApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *RansomwareConfig      `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfGetRansomwareConfigApiResponseData() *OneOfGetRansomwareConfigApiResponseData {
	p := new(OneOfGetRansomwareConfigApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetRansomwareConfigApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetRansomwareConfigApiResponseData is nil"))
	}
	switch v.(type) {
	case RansomwareConfig:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(RansomwareConfig)
		}
		*p.oneOfType0 = v.(RansomwareConfig)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfGetRansomwareConfigApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetRansomwareConfigApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(RansomwareConfig)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "files.v4.config.RansomwareConfig" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(RansomwareConfig)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetRansomwareConfigApiResponseData"))
}

func (p *OneOfGetRansomwareConfigApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetRansomwareConfigApiResponseData")
}

type OneOfObjectStoreConfigConfiguration struct {
	Discriminator *string      `json:"-"`
	ObjectType_   *string      `json:"-"`
	oneOfType0    *AWSConfig   `json:"-"`
	oneOfType1    *AzureConfig `json:"-"`
}

func NewOneOfObjectStoreConfigConfiguration() *OneOfObjectStoreConfigConfiguration {
	p := new(OneOfObjectStoreConfigConfiguration)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfObjectStoreConfigConfiguration) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfObjectStoreConfigConfiguration is nil"))
	}
	switch v.(type) {
	case AWSConfig:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(AWSConfig)
		}
		*p.oneOfType0 = v.(AWSConfig)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case AzureConfig:
		if nil == p.oneOfType1 {
			p.oneOfType1 = new(AzureConfig)
		}
		*p.oneOfType1 = v.(AzureConfig)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType1.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType1.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfObjectStoreConfigConfiguration) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1
	}
	return nil
}

func (p *OneOfObjectStoreConfigConfiguration) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(AWSConfig)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "files.v4.config.AWSConfig" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(AWSConfig)
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
	vOneOfType1 := new(AzureConfig)
	if err := json.Unmarshal(b, vOneOfType1); err == nil {
		if "files.v4.config.AzureConfig" == *vOneOfType1.ObjectType_ {
			if nil == p.oneOfType1 {
				p.oneOfType1 = new(AzureConfig)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfObjectStoreConfigConfiguration"))
}

func (p *OneOfObjectStoreConfigConfiguration) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	return nil, errors.New("No value to marshal for OneOfObjectStoreConfigConfiguration")
}

type OneOfCreateUnifiedNamespaceApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfCreateUnifiedNamespaceApiResponseData() *OneOfCreateUnifiedNamespaceApiResponseData {
	p := new(OneOfCreateUnifiedNamespaceApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateUnifiedNamespaceApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateUnifiedNamespaceApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.TaskReference)
		}
		*p.oneOfType0 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfCreateUnifiedNamespaceApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfCreateUnifiedNamespaceApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.TaskReference)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateUnifiedNamespaceApiResponseData"))
}

func (p *OneOfCreateUnifiedNamespaceApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfCreateUnifiedNamespaceApiResponseData")
}

type OneOfDeleteRansomwareConfigApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfDeleteRansomwareConfigApiResponseData() *OneOfDeleteRansomwareConfigApiResponseData {
	p := new(OneOfDeleteRansomwareConfigApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteRansomwareConfigApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteRansomwareConfigApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.TaskReference)
		}
		*p.oneOfType0 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfDeleteRansomwareConfigApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDeleteRansomwareConfigApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.TaskReference)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteRansomwareConfigApiResponseData"))
}

func (p *OneOfDeleteRansomwareConfigApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteRansomwareConfigApiResponseData")
}

type OneOfGetSnapshotScheduleApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *SnapshotSchedule      `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfGetSnapshotScheduleApiResponseData() *OneOfGetSnapshotScheduleApiResponseData {
	p := new(OneOfGetSnapshotScheduleApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetSnapshotScheduleApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetSnapshotScheduleApiResponseData is nil"))
	}
	switch v.(type) {
	case SnapshotSchedule:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(SnapshotSchedule)
		}
		*p.oneOfType0 = v.(SnapshotSchedule)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfGetSnapshotScheduleApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetSnapshotScheduleApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(SnapshotSchedule)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "files.v4.config.SnapshotSchedule" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(SnapshotSchedule)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetSnapshotScheduleApiResponseData"))
}

func (p *OneOfGetSnapshotScheduleApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetSnapshotScheduleApiResponseData")
}

type OneOfCreateSnapshotScheduleApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *SnapshotSchedule      `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfCreateSnapshotScheduleApiResponseData() *OneOfCreateSnapshotScheduleApiResponseData {
	p := new(OneOfCreateSnapshotScheduleApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateSnapshotScheduleApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateSnapshotScheduleApiResponseData is nil"))
	}
	switch v.(type) {
	case SnapshotSchedule:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(SnapshotSchedule)
		}
		*p.oneOfType0 = v.(SnapshotSchedule)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfCreateSnapshotScheduleApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfCreateSnapshotScheduleApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(SnapshotSchedule)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "files.v4.config.SnapshotSchedule" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(SnapshotSchedule)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateSnapshotScheduleApiResponseData"))
}

func (p *OneOfCreateSnapshotScheduleApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfCreateSnapshotScheduleApiResponseData")
}

type OneOfGetAntivirusServerApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *AntivirusServer       `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfGetAntivirusServerApiResponseData() *OneOfGetAntivirusServerApiResponseData {
	p := new(OneOfGetAntivirusServerApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetAntivirusServerApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetAntivirusServerApiResponseData is nil"))
	}
	switch v.(type) {
	case AntivirusServer:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(AntivirusServer)
		}
		*p.oneOfType0 = v.(AntivirusServer)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfGetAntivirusServerApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetAntivirusServerApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(AntivirusServer)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "files.v4.config.AntivirusServer" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(AntivirusServer)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetAntivirusServerApiResponseData"))
}

func (p *OneOfGetAntivirusServerApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetAntivirusServerApiResponseData")
}

type OneOfDeleteNotificationPolicyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfDeleteNotificationPolicyApiResponseData() *OneOfDeleteNotificationPolicyApiResponseData {
	p := new(OneOfDeleteNotificationPolicyApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteNotificationPolicyApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteNotificationPolicyApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.TaskReference)
		}
		*p.oneOfType0 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfDeleteNotificationPolicyApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDeleteNotificationPolicyApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.TaskReference)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteNotificationPolicyApiResponseData"))
}

func (p *OneOfDeleteNotificationPolicyApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteNotificationPolicyApiResponseData")
}

type OneOfUpdateEmailConfigApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    []import4.AppMessage   `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfUpdateEmailConfigApiResponseData() *OneOfUpdateEmailConfigApiResponseData {
	p := new(OneOfUpdateEmailConfigApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateEmailConfigApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateEmailConfigApiResponseData is nil"))
	}
	switch v.(type) {
	case []import4.AppMessage:
		p.oneOfType0 = v.([]import4.AppMessage)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<files.v4.error.AppMessage>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<files.v4.error.AppMessage>"
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfUpdateEmailConfigApiResponseData) GetValue() interface{} {
	if "List<files.v4.error.AppMessage>" == *p.Discriminator {
		return p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfUpdateEmailConfigApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new([]import4.AppMessage)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "files.v4.error.AppMessage" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<files.v4.error.AppMessage>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<files.v4.error.AppMessage>"
			return nil
		}
	}
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateEmailConfigApiResponseData"))
}

func (p *OneOfUpdateEmailConfigApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<files.v4.error.AppMessage>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateEmailConfigApiResponseData")
}

type OneOfCreateTieringConfigurationApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfCreateTieringConfigurationApiResponseData() *OneOfCreateTieringConfigurationApiResponseData {
	p := new(OneOfCreateTieringConfigurationApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateTieringConfigurationApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateTieringConfigurationApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.TaskReference)
		}
		*p.oneOfType0 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfCreateTieringConfigurationApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfCreateTieringConfigurationApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.TaskReference)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateTieringConfigurationApiResponseData"))
}

func (p *OneOfCreateTieringConfigurationApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfCreateTieringConfigurationApiResponseData")
}

type OneOfUpdateTieringConfigurationApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfUpdateTieringConfigurationApiResponseData() *OneOfUpdateTieringConfigurationApiResponseData {
	p := new(OneOfUpdateTieringConfigurationApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateTieringConfigurationApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateTieringConfigurationApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.TaskReference)
		}
		*p.oneOfType0 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfUpdateTieringConfigurationApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfUpdateTieringConfigurationApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.TaskReference)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateTieringConfigurationApiResponseData"))
}

func (p *OneOfUpdateTieringConfigurationApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateTieringConfigurationApiResponseData")
}

type OneOfListRecommendationsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    []Recommendation       `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfListRecommendationsApiResponseData() *OneOfListRecommendationsApiResponseData {
	p := new(OneOfListRecommendationsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListRecommendationsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListRecommendationsApiResponseData is nil"))
	}
	switch v.(type) {
	case []Recommendation:
		p.oneOfType0 = v.([]Recommendation)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<files.v4.config.Recommendation>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<files.v4.config.Recommendation>"
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfListRecommendationsApiResponseData) GetValue() interface{} {
	if "List<files.v4.config.Recommendation>" == *p.Discriminator {
		return p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfListRecommendationsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new([]Recommendation)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "files.v4.config.Recommendation" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<files.v4.config.Recommendation>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<files.v4.config.Recommendation>"
			return nil
		}
	}
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListRecommendationsApiResponseData"))
}

func (p *OneOfListRecommendationsApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<files.v4.config.Recommendation>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfListRecommendationsApiResponseData")
}

type OneOfReviseDnsRecordsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfReviseDnsRecordsApiResponseData() *OneOfReviseDnsRecordsApiResponseData {
	p := new(OneOfReviseDnsRecordsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfReviseDnsRecordsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfReviseDnsRecordsApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.TaskReference)
		}
		*p.oneOfType0 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfReviseDnsRecordsApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfReviseDnsRecordsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.TaskReference)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfReviseDnsRecordsApiResponseData"))
}

func (p *OneOfReviseDnsRecordsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfReviseDnsRecordsApiResponseData")
}

type OneOfGetPartnerServerApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *PartnerServer         `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfGetPartnerServerApiResponseData() *OneOfGetPartnerServerApiResponseData {
	p := new(OneOfGetPartnerServerApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetPartnerServerApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetPartnerServerApiResponseData is nil"))
	}
	switch v.(type) {
	case PartnerServer:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(PartnerServer)
		}
		*p.oneOfType0 = v.(PartnerServer)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfGetPartnerServerApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetPartnerServerApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(PartnerServer)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "files.v4.config.PartnerServer" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(PartnerServer)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetPartnerServerApiResponseData"))
}

func (p *OneOfGetPartnerServerApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetPartnerServerApiResponseData")
}

type OneOfDeleteReplicationPolicyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfDeleteReplicationPolicyApiResponseData() *OneOfDeleteReplicationPolicyApiResponseData {
	p := new(OneOfDeleteReplicationPolicyApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteReplicationPolicyApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteReplicationPolicyApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.TaskReference)
		}
		*p.oneOfType0 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfDeleteReplicationPolicyApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDeleteReplicationPolicyApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.TaskReference)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteReplicationPolicyApiResponseData"))
}

func (p *OneOfDeleteReplicationPolicyApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteReplicationPolicyApiResponseData")
}

type OneOfUpdatePartnerServerApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfUpdatePartnerServerApiResponseData() *OneOfUpdatePartnerServerApiResponseData {
	p := new(OneOfUpdatePartnerServerApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdatePartnerServerApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdatePartnerServerApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.TaskReference)
		}
		*p.oneOfType0 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfUpdatePartnerServerApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfUpdatePartnerServerApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.TaskReference)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdatePartnerServerApiResponseData"))
}

func (p *OneOfUpdatePartnerServerApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfUpdatePartnerServerApiResponseData")
}

type OneOfListReplicationPoliciesApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    []ReplicationPolicy    `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfListReplicationPoliciesApiResponseData() *OneOfListReplicationPoliciesApiResponseData {
	p := new(OneOfListReplicationPoliciesApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListReplicationPoliciesApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListReplicationPoliciesApiResponseData is nil"))
	}
	switch v.(type) {
	case []ReplicationPolicy:
		p.oneOfType0 = v.([]ReplicationPolicy)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<files.v4.config.ReplicationPolicy>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<files.v4.config.ReplicationPolicy>"
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfListReplicationPoliciesApiResponseData) GetValue() interface{} {
	if "List<files.v4.config.ReplicationPolicy>" == *p.Discriminator {
		return p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfListReplicationPoliciesApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new([]ReplicationPolicy)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "files.v4.config.ReplicationPolicy" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<files.v4.config.ReplicationPolicy>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<files.v4.config.ReplicationPolicy>"
			return nil
		}
	}
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListReplicationPoliciesApiResponseData"))
}

func (p *OneOfListReplicationPoliciesApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<files.v4.config.ReplicationPolicy>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfListReplicationPoliciesApiResponseData")
}

type OneOfGetMountTargetSnapshotApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *Snapshot              `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfGetMountTargetSnapshotApiResponseData() *OneOfGetMountTargetSnapshotApiResponseData {
	p := new(OneOfGetMountTargetSnapshotApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetMountTargetSnapshotApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetMountTargetSnapshotApiResponseData is nil"))
	}
	switch v.(type) {
	case Snapshot:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(Snapshot)
		}
		*p.oneOfType0 = v.(Snapshot)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfGetMountTargetSnapshotApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetMountTargetSnapshotApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(Snapshot)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "files.v4.config.Snapshot" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(Snapshot)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetMountTargetSnapshotApiResponseData"))
}

func (p *OneOfGetMountTargetSnapshotApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetMountTargetSnapshotApiResponseData")
}

type OneOfDeleteQuotaPolicyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfDeleteQuotaPolicyApiResponseData() *OneOfDeleteQuotaPolicyApiResponseData {
	p := new(OneOfDeleteQuotaPolicyApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteQuotaPolicyApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteQuotaPolicyApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.TaskReference)
		}
		*p.oneOfType0 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.TaskReference)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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

type OneOfListObjectStoreProfilesApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    []ObjectStoreProfile   `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfListObjectStoreProfilesApiResponseData() *OneOfListObjectStoreProfilesApiResponseData {
	p := new(OneOfListObjectStoreProfilesApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListObjectStoreProfilesApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListObjectStoreProfilesApiResponseData is nil"))
	}
	switch v.(type) {
	case []ObjectStoreProfile:
		p.oneOfType0 = v.([]ObjectStoreProfile)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<files.v4.config.ObjectStoreProfile>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<files.v4.config.ObjectStoreProfile>"
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfListObjectStoreProfilesApiResponseData) GetValue() interface{} {
	if "List<files.v4.config.ObjectStoreProfile>" == *p.Discriminator {
		return p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfListObjectStoreProfilesApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new([]ObjectStoreProfile)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "files.v4.config.ObjectStoreProfile" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<files.v4.config.ObjectStoreProfile>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<files.v4.config.ObjectStoreProfile>"
			return nil
		}
	}
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListObjectStoreProfilesApiResponseData"))
}

func (p *OneOfListObjectStoreProfilesApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<files.v4.config.ObjectStoreProfile>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfListObjectStoreProfilesApiResponseData")
}

type OneOfListVirusScanPoliciesApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    []VirusScanPolicy      `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfListVirusScanPoliciesApiResponseData() *OneOfListVirusScanPoliciesApiResponseData {
	p := new(OneOfListVirusScanPoliciesApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListVirusScanPoliciesApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListVirusScanPoliciesApiResponseData is nil"))
	}
	switch v.(type) {
	case []VirusScanPolicy:
		p.oneOfType0 = v.([]VirusScanPolicy)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<files.v4.config.VirusScanPolicy>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<files.v4.config.VirusScanPolicy>"
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfListVirusScanPoliciesApiResponseData) GetValue() interface{} {
	if "List<files.v4.config.VirusScanPolicy>" == *p.Discriminator {
		return p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfListVirusScanPoliciesApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new([]VirusScanPolicy)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "files.v4.config.VirusScanPolicy" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<files.v4.config.VirusScanPolicy>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<files.v4.config.VirusScanPolicy>"
			return nil
		}
	}
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListVirusScanPoliciesApiResponseData"))
}

func (p *OneOfListVirusScanPoliciesApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<files.v4.config.VirusScanPolicy>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfListVirusScanPoliciesApiResponseData")
}

type OneOfSnapshotScheduleSchedule struct {
	Discriminator *string          `json:"-"`
	ObjectType_   *string          `json:"-"`
	oneOfType2    *MonthlySchedule `json:"-"`
	oneOfType0    *DailySchedule   `json:"-"`
	oneOfType1    *WeeklySchedule  `json:"-"`
}

func NewOneOfSnapshotScheduleSchedule() *OneOfSnapshotScheduleSchedule {
	p := new(OneOfSnapshotScheduleSchedule)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfSnapshotScheduleSchedule) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfSnapshotScheduleSchedule is nil"))
	}
	switch v.(type) {
	case MonthlySchedule:
		if nil == p.oneOfType2 {
			p.oneOfType2 = new(MonthlySchedule)
		}
		*p.oneOfType2 = v.(MonthlySchedule)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2.ObjectType_
	case DailySchedule:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(DailySchedule)
		}
		*p.oneOfType0 = v.(DailySchedule)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case WeeklySchedule:
		if nil == p.oneOfType1 {
			p.oneOfType1 = new(WeeklySchedule)
		}
		*p.oneOfType1 = v.(WeeklySchedule)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType1.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType1.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfSnapshotScheduleSchedule) GetValue() interface{} {
	if p.oneOfType2 != nil && *p.oneOfType2.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1
	}
	return nil
}

func (p *OneOfSnapshotScheduleSchedule) UnmarshalJSON(b []byte) error {
	vOneOfType2 := new(MonthlySchedule)
	if err := json.Unmarshal(b, vOneOfType2); err == nil {
		if "files.v4.config.MonthlySchedule" == *vOneOfType2.ObjectType_ {
			if nil == p.oneOfType2 {
				p.oneOfType2 = new(MonthlySchedule)
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
	vOneOfType0 := new(DailySchedule)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "files.v4.config.DailySchedule" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(DailySchedule)
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
	vOneOfType1 := new(WeeklySchedule)
	if err := json.Unmarshal(b, vOneOfType1); err == nil {
		if "files.v4.config.WeeklySchedule" == *vOneOfType1.ObjectType_ {
			if nil == p.oneOfType1 {
				p.oneOfType1 = new(WeeklySchedule)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfSnapshotScheduleSchedule"))
}

func (p *OneOfSnapshotScheduleSchedule) MarshalJSON() ([]byte, error) {
	if p.oneOfType2 != nil && *p.oneOfType2.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	return nil, errors.New("No value to marshal for OneOfSnapshotScheduleSchedule")
}

type OneOfDeletePartnerServerApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfDeletePartnerServerApiResponseData() *OneOfDeletePartnerServerApiResponseData {
	p := new(OneOfDeletePartnerServerApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeletePartnerServerApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeletePartnerServerApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.TaskReference)
		}
		*p.oneOfType0 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfDeletePartnerServerApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDeletePartnerServerApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.TaskReference)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeletePartnerServerApiResponseData"))
}

func (p *OneOfDeletePartnerServerApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDeletePartnerServerApiResponseData")
}

type OneOfUpdateUnifiedNamespaceApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfUpdateUnifiedNamespaceApiResponseData() *OneOfUpdateUnifiedNamespaceApiResponseData {
	p := new(OneOfUpdateUnifiedNamespaceApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateUnifiedNamespaceApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateUnifiedNamespaceApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.TaskReference)
		}
		*p.oneOfType0 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfUpdateUnifiedNamespaceApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfUpdateUnifiedNamespaceApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.TaskReference)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateUnifiedNamespaceApiResponseData"))
}

func (p *OneOfUpdateUnifiedNamespaceApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateUnifiedNamespaceApiResponseData")
}

type OneOfDeleteMountTargetApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfDeleteMountTargetApiResponseData() *OneOfDeleteMountTargetApiResponseData {
	p := new(OneOfDeleteMountTargetApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteMountTargetApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteMountTargetApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.TaskReference)
		}
		*p.oneOfType0 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.TaskReference)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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

type OneOfCreateMountTargetSnapshotApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfCreateMountTargetSnapshotApiResponseData() *OneOfCreateMountTargetSnapshotApiResponseData {
	p := new(OneOfCreateMountTargetSnapshotApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateMountTargetSnapshotApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateMountTargetSnapshotApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.TaskReference)
		}
		*p.oneOfType0 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfCreateMountTargetSnapshotApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfCreateMountTargetSnapshotApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.TaskReference)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateMountTargetSnapshotApiResponseData"))
}

func (p *OneOfCreateMountTargetSnapshotApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfCreateMountTargetSnapshotApiResponseData")
}

type OneOfEnableWormLegalHoldApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfEnableWormLegalHoldApiResponseData() *OneOfEnableWormLegalHoldApiResponseData {
	p := new(OneOfEnableWormLegalHoldApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfEnableWormLegalHoldApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfEnableWormLegalHoldApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.TaskReference)
		}
		*p.oneOfType0 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfEnableWormLegalHoldApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfEnableWormLegalHoldApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.TaskReference)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfEnableWormLegalHoldApiResponseData"))
}

func (p *OneOfEnableWormLegalHoldApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfEnableWormLegalHoldApiResponseData")
}

type OneOfTestConnectionAntivirusServerApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfTestConnectionAntivirusServerApiResponseData() *OneOfTestConnectionAntivirusServerApiResponseData {
	p := new(OneOfTestConnectionAntivirusServerApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfTestConnectionAntivirusServerApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfTestConnectionAntivirusServerApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.TaskReference)
		}
		*p.oneOfType0 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfTestConnectionAntivirusServerApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfTestConnectionAntivirusServerApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.TaskReference)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfTestConnectionAntivirusServerApiResponseData"))
}

func (p *OneOfTestConnectionAntivirusServerApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfTestConnectionAntivirusServerApiResponseData")
}

type OneOfUpdateObjectStoreProfileApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfUpdateObjectStoreProfileApiResponseData() *OneOfUpdateObjectStoreProfileApiResponseData {
	p := new(OneOfUpdateObjectStoreProfileApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateObjectStoreProfileApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateObjectStoreProfileApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.TaskReference)
		}
		*p.oneOfType0 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfUpdateObjectStoreProfileApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfUpdateObjectStoreProfileApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.TaskReference)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateObjectStoreProfileApiResponseData"))
}

func (p *OneOfUpdateObjectStoreProfileApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateObjectStoreProfileApiResponseData")
}

type OneOfListNotificationPoliciesApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType0    []NotificationPolicy   `json:"-"`
}

func NewOneOfListNotificationPoliciesApiResponseData() *OneOfListNotificationPoliciesApiResponseData {
	p := new(OneOfListNotificationPoliciesApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListNotificationPoliciesApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListNotificationPoliciesApiResponseData is nil"))
	}
	switch v.(type) {
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case []NotificationPolicy:
		p.oneOfType0 = v.([]NotificationPolicy)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<files.v4.config.NotificationPolicy>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<files.v4.config.NotificationPolicy>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListNotificationPoliciesApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<files.v4.config.NotificationPolicy>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfListNotificationPoliciesApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	vOneOfType0 := new([]NotificationPolicy)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "files.v4.config.NotificationPolicy" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<files.v4.config.NotificationPolicy>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<files.v4.config.NotificationPolicy>"
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListNotificationPoliciesApiResponseData"))
}

func (p *OneOfListNotificationPoliciesApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<files.v4.config.NotificationPolicy>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfListNotificationPoliciesApiResponseData")
}

type OneOfDeleteInfectedFileApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfDeleteInfectedFileApiResponseData() *OneOfDeleteInfectedFileApiResponseData {
	p := new(OneOfDeleteInfectedFileApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteInfectedFileApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteInfectedFileApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.TaskReference)
		}
		*p.oneOfType0 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfDeleteInfectedFileApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDeleteInfectedFileApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.TaskReference)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteInfectedFileApiResponseData"))
}

func (p *OneOfDeleteInfectedFileApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteInfectedFileApiResponseData")
}

type OneOfListMountTargetsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    []MountTarget          `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfListMountTargetsApiResponseData() *OneOfListMountTargetsApiResponseData {
	p := new(OneOfListMountTargetsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListMountTargetsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListMountTargetsApiResponseData is nil"))
	}
	switch v.(type) {
	case []MountTarget:
		p.oneOfType0 = v.([]MountTarget)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<files.v4.config.MountTarget>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<files.v4.config.MountTarget>"
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfListMountTargetsApiResponseData) GetValue() interface{} {
	if "List<files.v4.config.MountTarget>" == *p.Discriminator {
		return p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfListMountTargetsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new([]MountTarget)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "files.v4.config.MountTarget" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<files.v4.config.MountTarget>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<files.v4.config.MountTarget>"
			return nil
		}
	}
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListMountTargetsApiResponseData"))
}

func (p *OneOfListMountTargetsApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<files.v4.config.MountTarget>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfListMountTargetsApiResponseData")
}

type OneOfDeleteTieringConfigurationApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfDeleteTieringConfigurationApiResponseData() *OneOfDeleteTieringConfigurationApiResponseData {
	p := new(OneOfDeleteTieringConfigurationApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteTieringConfigurationApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteTieringConfigurationApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.TaskReference)
		}
		*p.oneOfType0 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfDeleteTieringConfigurationApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDeleteTieringConfigurationApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.TaskReference)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteTieringConfigurationApiResponseData"))
}

func (p *OneOfDeleteTieringConfigurationApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteTieringConfigurationApiResponseData")
}

type OneOfFixInfectedFileApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfFixInfectedFileApiResponseData() *OneOfFixInfectedFileApiResponseData {
	p := new(OneOfFixInfectedFileApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfFixInfectedFileApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfFixInfectedFileApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.TaskReference)
		}
		*p.oneOfType0 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfFixInfectedFileApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfFixInfectedFileApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.TaskReference)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfFixInfectedFileApiResponseData"))
}

func (p *OneOfFixInfectedFileApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfFixInfectedFileApiResponseData")
}

type OneOfRestoreMountTargetApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfRestoreMountTargetApiResponseData() *OneOfRestoreMountTargetApiResponseData {
	p := new(OneOfRestoreMountTargetApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfRestoreMountTargetApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfRestoreMountTargetApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.TaskReference)
		}
		*p.oneOfType0 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfRestoreMountTargetApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfRestoreMountTargetApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.TaskReference)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfRestoreMountTargetApiResponseData"))
}

func (p *OneOfRestoreMountTargetApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfRestoreMountTargetApiResponseData")
}

type OneOfListDnsRecordsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    []DnsRecord            `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfListDnsRecordsApiResponseData() *OneOfListDnsRecordsApiResponseData {
	p := new(OneOfListDnsRecordsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListDnsRecordsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListDnsRecordsApiResponseData is nil"))
	}
	switch v.(type) {
	case []DnsRecord:
		p.oneOfType0 = v.([]DnsRecord)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<files.v4.config.DnsRecord>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<files.v4.config.DnsRecord>"
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfListDnsRecordsApiResponseData) GetValue() interface{} {
	if "List<files.v4.config.DnsRecord>" == *p.Discriminator {
		return p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfListDnsRecordsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new([]DnsRecord)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "files.v4.config.DnsRecord" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<files.v4.config.DnsRecord>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<files.v4.config.DnsRecord>"
			return nil
		}
	}
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListDnsRecordsApiResponseData"))
}

func (p *OneOfListDnsRecordsApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<files.v4.config.DnsRecord>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfListDnsRecordsApiResponseData")
}

type OneOfVerifyDnsRecordsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfVerifyDnsRecordsApiResponseData() *OneOfVerifyDnsRecordsApiResponseData {
	p := new(OneOfVerifyDnsRecordsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfVerifyDnsRecordsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfVerifyDnsRecordsApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.TaskReference)
		}
		*p.oneOfType0 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfVerifyDnsRecordsApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfVerifyDnsRecordsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.TaskReference)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfVerifyDnsRecordsApiResponseData"))
}

func (p *OneOfVerifyDnsRecordsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfVerifyDnsRecordsApiResponseData")
}

type OneOfCreateRansomwareConfigApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfCreateRansomwareConfigApiResponseData() *OneOfCreateRansomwareConfigApiResponseData {
	p := new(OneOfCreateRansomwareConfigApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateRansomwareConfigApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateRansomwareConfigApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.TaskReference)
		}
		*p.oneOfType0 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfCreateRansomwareConfigApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfCreateRansomwareConfigApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.TaskReference)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateRansomwareConfigApiResponseData"))
}

func (p *OneOfCreateRansomwareConfigApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfCreateRansomwareConfigApiResponseData")
}

type OneOfCreateReplicationPolicyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfCreateReplicationPolicyApiResponseData() *OneOfCreateReplicationPolicyApiResponseData {
	p := new(OneOfCreateReplicationPolicyApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateReplicationPolicyApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateReplicationPolicyApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.TaskReference)
		}
		*p.oneOfType0 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfCreateReplicationPolicyApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfCreateReplicationPolicyApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.TaskReference)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateReplicationPolicyApiResponseData"))
}

func (p *OneOfCreateReplicationPolicyApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfCreateReplicationPolicyApiResponseData")
}

type OneOfListVdiUserSessionsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    []VdiUserSession       `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfListVdiUserSessionsApiResponseData() *OneOfListVdiUserSessionsApiResponseData {
	p := new(OneOfListVdiUserSessionsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListVdiUserSessionsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListVdiUserSessionsApiResponseData is nil"))
	}
	switch v.(type) {
	case []VdiUserSession:
		p.oneOfType0 = v.([]VdiUserSession)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<files.v4.config.VdiUserSession>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<files.v4.config.VdiUserSession>"
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfListVdiUserSessionsApiResponseData) GetValue() interface{} {
	if "List<files.v4.config.VdiUserSession>" == *p.Discriminator {
		return p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfListVdiUserSessionsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new([]VdiUserSession)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "files.v4.config.VdiUserSession" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<files.v4.config.VdiUserSession>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<files.v4.config.VdiUserSession>"
			return nil
		}
	}
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListVdiUserSessionsApiResponseData"))
}

func (p *OneOfListVdiUserSessionsApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<files.v4.config.VdiUserSession>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfListVdiUserSessionsApiResponseData")
}

type OneOfCloneMountTargetApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfCloneMountTargetApiResponseData() *OneOfCloneMountTargetApiResponseData {
	p := new(OneOfCloneMountTargetApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCloneMountTargetApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCloneMountTargetApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.TaskReference)
		}
		*p.oneOfType0 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfCloneMountTargetApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfCloneMountTargetApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.TaskReference)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCloneMountTargetApiResponseData"))
}

func (p *OneOfCloneMountTargetApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfCloneMountTargetApiResponseData")
}

type OneOfDownloadUserMappingsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *FileDetail            `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfDownloadUserMappingsApiResponseData() *OneOfDownloadUserMappingsApiResponseData {
	p := new(OneOfDownloadUserMappingsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDownloadUserMappingsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDownloadUserMappingsApiResponseData is nil"))
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
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfDownloadUserMappingsApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && "FileDetail" == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDownloadUserMappingsApiResponseData) UnmarshalJSON(b []byte) error {
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDownloadUserMappingsApiResponseData"))
}

func (p *OneOfDownloadUserMappingsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && "FileDetail" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDownloadUserMappingsApiResponseData")
}

type OneOfUpdateReplicationPolicyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfUpdateReplicationPolicyApiResponseData() *OneOfUpdateReplicationPolicyApiResponseData {
	p := new(OneOfUpdateReplicationPolicyApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateReplicationPolicyApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateReplicationPolicyApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.TaskReference)
		}
		*p.oneOfType0 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfUpdateReplicationPolicyApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfUpdateReplicationPolicyApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.TaskReference)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateReplicationPolicyApiResponseData"))
}

func (p *OneOfUpdateReplicationPolicyApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateReplicationPolicyApiResponseData")
}

type OneOfListRansomwareConfigsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    []RansomwareConfig     `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfListRansomwareConfigsApiResponseData() *OneOfListRansomwareConfigsApiResponseData {
	p := new(OneOfListRansomwareConfigsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListRansomwareConfigsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListRansomwareConfigsApiResponseData is nil"))
	}
	switch v.(type) {
	case []RansomwareConfig:
		p.oneOfType0 = v.([]RansomwareConfig)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<files.v4.config.RansomwareConfig>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<files.v4.config.RansomwareConfig>"
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfListRansomwareConfigsApiResponseData) GetValue() interface{} {
	if "List<files.v4.config.RansomwareConfig>" == *p.Discriminator {
		return p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfListRansomwareConfigsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new([]RansomwareConfig)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "files.v4.config.RansomwareConfig" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<files.v4.config.RansomwareConfig>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<files.v4.config.RansomwareConfig>"
			return nil
		}
	}
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListRansomwareConfigsApiResponseData"))
}

func (p *OneOfListRansomwareConfigsApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<files.v4.config.RansomwareConfig>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfListRansomwareConfigsApiResponseData")
}

type OneOfDeleteUnifiedNamespaceApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfDeleteUnifiedNamespaceApiResponseData() *OneOfDeleteUnifiedNamespaceApiResponseData {
	p := new(OneOfDeleteUnifiedNamespaceApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteUnifiedNamespaceApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteUnifiedNamespaceApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.TaskReference)
		}
		*p.oneOfType0 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfDeleteUnifiedNamespaceApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDeleteUnifiedNamespaceApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.TaskReference)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteUnifiedNamespaceApiResponseData"))
}

func (p *OneOfDeleteUnifiedNamespaceApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteUnifiedNamespaceApiResponseData")
}

type OneOfGetInfectedFileApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *InfectedFile          `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfGetInfectedFileApiResponseData() *OneOfGetInfectedFileApiResponseData {
	p := new(OneOfGetInfectedFileApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetInfectedFileApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetInfectedFileApiResponseData is nil"))
	}
	switch v.(type) {
	case InfectedFile:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(InfectedFile)
		}
		*p.oneOfType0 = v.(InfectedFile)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfGetInfectedFileApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetInfectedFileApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(InfectedFile)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "files.v4.config.InfectedFile" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(InfectedFile)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetInfectedFileApiResponseData"))
}

func (p *OneOfGetInfectedFileApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetInfectedFileApiResponseData")
}

type OneOfUpdateRansomwareConfigApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfUpdateRansomwareConfigApiResponseData() *OneOfUpdateRansomwareConfigApiResponseData {
	p := new(OneOfUpdateRansomwareConfigApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateRansomwareConfigApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateRansomwareConfigApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.TaskReference)
		}
		*p.oneOfType0 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.TaskReference)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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

type OneOfUpdateVirusScanPolicyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfUpdateVirusScanPolicyApiResponseData() *OneOfUpdateVirusScanPolicyApiResponseData {
	p := new(OneOfUpdateVirusScanPolicyApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateVirusScanPolicyApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateVirusScanPolicyApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.TaskReference)
		}
		*p.oneOfType0 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfUpdateVirusScanPolicyApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfUpdateVirusScanPolicyApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.TaskReference)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateVirusScanPolicyApiResponseData"))
}

func (p *OneOfUpdateVirusScanPolicyApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateVirusScanPolicyApiResponseData")
}

type OneOfListQuotaPoliciesApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    []QuotaPolicy          `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfListQuotaPoliciesApiResponseData() *OneOfListQuotaPoliciesApiResponseData {
	p := new(OneOfListQuotaPoliciesApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListQuotaPoliciesApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListQuotaPoliciesApiResponseData is nil"))
	}
	switch v.(type) {
	case []QuotaPolicy:
		p.oneOfType0 = v.([]QuotaPolicy)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<files.v4.config.QuotaPolicy>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<files.v4.config.QuotaPolicy>"
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfListQuotaPoliciesApiResponseData) GetValue() interface{} {
	if "List<files.v4.config.QuotaPolicy>" == *p.Discriminator {
		return p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfListQuotaPoliciesApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new([]QuotaPolicy)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "files.v4.config.QuotaPolicy" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<files.v4.config.QuotaPolicy>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<files.v4.config.QuotaPolicy>"
			return nil
		}
	}
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListQuotaPoliciesApiResponseData"))
}

func (p *OneOfListQuotaPoliciesApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<files.v4.config.QuotaPolicy>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfListQuotaPoliciesApiResponseData")
}

type OneOfDeleteMountTargetSnapshotApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfDeleteMountTargetSnapshotApiResponseData() *OneOfDeleteMountTargetSnapshotApiResponseData {
	p := new(OneOfDeleteMountTargetSnapshotApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteMountTargetSnapshotApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteMountTargetSnapshotApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.TaskReference)
		}
		*p.oneOfType0 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.TaskReference)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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

type OneOfListUnifiedNamespacesApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    []UnifiedNamespace     `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfListUnifiedNamespacesApiResponseData() *OneOfListUnifiedNamespacesApiResponseData {
	p := new(OneOfListUnifiedNamespacesApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListUnifiedNamespacesApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListUnifiedNamespacesApiResponseData is nil"))
	}
	switch v.(type) {
	case []UnifiedNamespace:
		p.oneOfType0 = v.([]UnifiedNamespace)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<files.v4.config.UnifiedNamespace>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<files.v4.config.UnifiedNamespace>"
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfListUnifiedNamespacesApiResponseData) GetValue() interface{} {
	if "List<files.v4.config.UnifiedNamespace>" == *p.Discriminator {
		return p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfListUnifiedNamespacesApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new([]UnifiedNamespace)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "files.v4.config.UnifiedNamespace" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<files.v4.config.UnifiedNamespace>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<files.v4.config.UnifiedNamespace>"
			return nil
		}
	}
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListUnifiedNamespacesApiResponseData"))
}

func (p *OneOfListUnifiedNamespacesApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<files.v4.config.UnifiedNamespace>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfListUnifiedNamespacesApiResponseData")
}

type OneOfCreateAntivirusServerApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfCreateAntivirusServerApiResponseData() *OneOfCreateAntivirusServerApiResponseData {
	p := new(OneOfCreateAntivirusServerApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateAntivirusServerApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateAntivirusServerApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.TaskReference)
		}
		*p.oneOfType0 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.TaskReference)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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

type OneOfDeleteVirusScanPolicyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfDeleteVirusScanPolicyApiResponseData() *OneOfDeleteVirusScanPolicyApiResponseData {
	p := new(OneOfDeleteVirusScanPolicyApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteVirusScanPolicyApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteVirusScanPolicyApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.TaskReference)
		}
		*p.oneOfType0 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfDeleteVirusScanPolicyApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDeleteVirusScanPolicyApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.TaskReference)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteVirusScanPolicyApiResponseData"))
}

func (p *OneOfDeleteVirusScanPolicyApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteVirusScanPolicyApiResponseData")
}

type OneOfGetSnapshotChangedContentApiResponseData struct {
	Discriminator *string                 `json:"-"`
	ObjectType_   *string                 `json:"-"`
	oneOfType0    *SnapshotChangedContent `json:"-"`
	oneOfType400  *import4.ErrorResponse  `json:"-"`
}

func NewOneOfGetSnapshotChangedContentApiResponseData() *OneOfGetSnapshotChangedContentApiResponseData {
	p := new(OneOfGetSnapshotChangedContentApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetSnapshotChangedContentApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetSnapshotChangedContentApiResponseData is nil"))
	}
	switch v.(type) {
	case SnapshotChangedContent:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(SnapshotChangedContent)
		}
		*p.oneOfType0 = v.(SnapshotChangedContent)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfGetSnapshotChangedContentApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetSnapshotChangedContentApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(SnapshotChangedContent)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "files.v4.config.SnapshotChangedContent" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(SnapshotChangedContent)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetSnapshotChangedContentApiResponseData"))
}

func (p *OneOfGetSnapshotChangedContentApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetSnapshotChangedContentApiResponseData")
}

type OneOfGetQuotaPolicyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *QuotaPolicy           `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfGetQuotaPolicyApiResponseData() *OneOfGetQuotaPolicyApiResponseData {
	p := new(OneOfGetQuotaPolicyApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetQuotaPolicyApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetQuotaPolicyApiResponseData is nil"))
	}
	switch v.(type) {
	case QuotaPolicy:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(QuotaPolicy)
		}
		*p.oneOfType0 = v.(QuotaPolicy)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfGetQuotaPolicyApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetQuotaPolicyApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(QuotaPolicy)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "files.v4.config.QuotaPolicy" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(QuotaPolicy)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetQuotaPolicyApiResponseData"))
}

func (p *OneOfGetQuotaPolicyApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetQuotaPolicyApiResponseData")
}

type OneOfDisableWormLegalHoldApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfDisableWormLegalHoldApiResponseData() *OneOfDisableWormLegalHoldApiResponseData {
	p := new(OneOfDisableWormLegalHoldApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDisableWormLegalHoldApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDisableWormLegalHoldApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.TaskReference)
		}
		*p.oneOfType0 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfDisableWormLegalHoldApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDisableWormLegalHoldApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.TaskReference)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDisableWormLegalHoldApiResponseData"))
}

func (p *OneOfDisableWormLegalHoldApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDisableWormLegalHoldApiResponseData")
}

type OneOfGetObjectStoreProfileApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *ObjectStoreProfile    `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfGetObjectStoreProfileApiResponseData() *OneOfGetObjectStoreProfileApiResponseData {
	p := new(OneOfGetObjectStoreProfileApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetObjectStoreProfileApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetObjectStoreProfileApiResponseData is nil"))
	}
	switch v.(type) {
	case ObjectStoreProfile:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(ObjectStoreProfile)
		}
		*p.oneOfType0 = v.(ObjectStoreProfile)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfGetObjectStoreProfileApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetObjectStoreProfileApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(ObjectStoreProfile)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "files.v4.config.ObjectStoreProfile" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(ObjectStoreProfile)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetObjectStoreProfileApiResponseData"))
}

func (p *OneOfGetObjectStoreProfileApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetObjectStoreProfileApiResponseData")
}

type OneOfDeleteRecommendationApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType1    *interface{}           `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfDeleteRecommendationApiResponseData() *OneOfDeleteRecommendationApiResponseData {
	p := new(OneOfDeleteRecommendationApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteRecommendationApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteRecommendationApiResponseData is nil"))
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
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfDeleteRecommendationApiResponseData) GetValue() interface{} {
	if "EMPTY" == *p.Discriminator {
		return *p.oneOfType1
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDeleteRecommendationApiResponseData) UnmarshalJSON(b []byte) error {
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteRecommendationApiResponseData"))
}

func (p *OneOfDeleteRecommendationApiResponseData) MarshalJSON() ([]byte, error) {
	if "EMPTY" == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteRecommendationApiResponseData")
}

type OneOfListPartnerServersApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    []PartnerServer        `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfListPartnerServersApiResponseData() *OneOfListPartnerServersApiResponseData {
	p := new(OneOfListPartnerServersApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListPartnerServersApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListPartnerServersApiResponseData is nil"))
	}
	switch v.(type) {
	case []PartnerServer:
		p.oneOfType0 = v.([]PartnerServer)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<files.v4.config.PartnerServer>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<files.v4.config.PartnerServer>"
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfListPartnerServersApiResponseData) GetValue() interface{} {
	if "List<files.v4.config.PartnerServer>" == *p.Discriminator {
		return p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfListPartnerServersApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new([]PartnerServer)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "files.v4.config.PartnerServer" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<files.v4.config.PartnerServer>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<files.v4.config.PartnerServer>"
			return nil
		}
	}
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListPartnerServersApiResponseData"))
}

func (p *OneOfListPartnerServersApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<files.v4.config.PartnerServer>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfListPartnerServersApiResponseData")
}

type OneOfDeleteSnapshotScheduleApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType1    *interface{}           `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfDeleteSnapshotScheduleApiResponseData() *OneOfDeleteSnapshotScheduleApiResponseData {
	p := new(OneOfDeleteSnapshotScheduleApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteSnapshotScheduleApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteSnapshotScheduleApiResponseData is nil"))
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
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfDeleteSnapshotScheduleApiResponseData) GetValue() interface{} {
	if "EMPTY" == *p.Discriminator {
		return *p.oneOfType1
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDeleteSnapshotScheduleApiResponseData) UnmarshalJSON(b []byte) error {
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteSnapshotScheduleApiResponseData"))
}

func (p *OneOfDeleteSnapshotScheduleApiResponseData) MarshalJSON() ([]byte, error) {
	if "EMPTY" == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteSnapshotScheduleApiResponseData")
}

type OneOfUpdateQuotaPolicyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfUpdateQuotaPolicyApiResponseData() *OneOfUpdateQuotaPolicyApiResponseData {
	p := new(OneOfUpdateQuotaPolicyApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateQuotaPolicyApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateQuotaPolicyApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.TaskReference)
		}
		*p.oneOfType0 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.TaskReference)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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

type OneOfCreatePartnerServerApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfCreatePartnerServerApiResponseData() *OneOfCreatePartnerServerApiResponseData {
	p := new(OneOfCreatePartnerServerApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreatePartnerServerApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreatePartnerServerApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.TaskReference)
		}
		*p.oneOfType0 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfCreatePartnerServerApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfCreatePartnerServerApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.TaskReference)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreatePartnerServerApiResponseData"))
}

func (p *OneOfCreatePartnerServerApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfCreatePartnerServerApiResponseData")
}

type OneOfCreateObjectStoreProfileApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfCreateObjectStoreProfileApiResponseData() *OneOfCreateObjectStoreProfileApiResponseData {
	p := new(OneOfCreateObjectStoreProfileApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateObjectStoreProfileApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateObjectStoreProfileApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.TaskReference)
		}
		*p.oneOfType0 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfCreateObjectStoreProfileApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfCreateObjectStoreProfileApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.TaskReference)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateObjectStoreProfileApiResponseData"))
}

func (p *OneOfCreateObjectStoreProfileApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfCreateObjectStoreProfileApiResponseData")
}

type OneOfGetUnifiedNamespaceApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *UnifiedNamespace      `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfGetUnifiedNamespaceApiResponseData() *OneOfGetUnifiedNamespaceApiResponseData {
	p := new(OneOfGetUnifiedNamespaceApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetUnifiedNamespaceApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetUnifiedNamespaceApiResponseData is nil"))
	}
	switch v.(type) {
	case UnifiedNamespace:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(UnifiedNamespace)
		}
		*p.oneOfType0 = v.(UnifiedNamespace)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfGetUnifiedNamespaceApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetUnifiedNamespaceApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(UnifiedNamespace)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "files.v4.config.UnifiedNamespace" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(UnifiedNamespace)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetUnifiedNamespaceApiResponseData"))
}

func (p *OneOfGetUnifiedNamespaceApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetUnifiedNamespaceApiResponseData")
}

type OneOfUpdateMountTargetApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfUpdateMountTargetApiResponseData() *OneOfUpdateMountTargetApiResponseData {
	p := new(OneOfUpdateMountTargetApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateMountTargetApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateMountTargetApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.TaskReference)
		}
		*p.oneOfType0 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.TaskReference)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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

type OneOfCreateVirusScanPolicyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfCreateVirusScanPolicyApiResponseData() *OneOfCreateVirusScanPolicyApiResponseData {
	p := new(OneOfCreateVirusScanPolicyApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateVirusScanPolicyApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateVirusScanPolicyApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.TaskReference)
		}
		*p.oneOfType0 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfCreateVirusScanPolicyApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfCreateVirusScanPolicyApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.TaskReference)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateVirusScanPolicyApiResponseData"))
}

func (p *OneOfCreateVirusScanPolicyApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfCreateVirusScanPolicyApiResponseData")
}

type OneOfGetTieringConfigurationApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType0    *TierConfiguration     `json:"-"`
}

func NewOneOfGetTieringConfigurationApiResponseData() *OneOfGetTieringConfigurationApiResponseData {
	p := new(OneOfGetTieringConfigurationApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetTieringConfigurationApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetTieringConfigurationApiResponseData is nil"))
	}
	switch v.(type) {
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case TierConfiguration:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(TierConfiguration)
		}
		*p.oneOfType0 = v.(TierConfiguration)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfGetTieringConfigurationApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfGetTieringConfigurationApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	vOneOfType0 := new(TierConfiguration)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "files.v4.config.TierConfiguration" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(TierConfiguration)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetTieringConfigurationApiResponseData"))
}

func (p *OneOfGetTieringConfigurationApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfGetTieringConfigurationApiResponseData")
}

type OneOfCreateMountTargetApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfCreateMountTargetApiResponseData() *OneOfCreateMountTargetApiResponseData {
	p := new(OneOfCreateMountTargetApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateMountTargetApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateMountTargetApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.TaskReference)
		}
		*p.oneOfType0 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.TaskReference)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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

type OneOfListTieringConfigurationsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    []TierConfiguration    `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfListTieringConfigurationsApiResponseData() *OneOfListTieringConfigurationsApiResponseData {
	p := new(OneOfListTieringConfigurationsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListTieringConfigurationsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListTieringConfigurationsApiResponseData is nil"))
	}
	switch v.(type) {
	case []TierConfiguration:
		p.oneOfType0 = v.([]TierConfiguration)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<files.v4.config.TierConfiguration>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<files.v4.config.TierConfiguration>"
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfListTieringConfigurationsApiResponseData) GetValue() interface{} {
	if "List<files.v4.config.TierConfiguration>" == *p.Discriminator {
		return p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfListTieringConfigurationsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new([]TierConfiguration)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "files.v4.config.TierConfiguration" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<files.v4.config.TierConfiguration>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<files.v4.config.TierConfiguration>"
			return nil
		}
	}
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListTieringConfigurationsApiResponseData"))
}

func (p *OneOfListTieringConfigurationsApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<files.v4.config.TierConfiguration>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfListTieringConfigurationsApiResponseData")
}

type OneOfListFileServersApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    []FileServer           `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfListFileServersApiResponseData() *OneOfListFileServersApiResponseData {
	p := new(OneOfListFileServersApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListFileServersApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListFileServersApiResponseData is nil"))
	}
	switch v.(type) {
	case []FileServer:
		p.oneOfType0 = v.([]FileServer)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<files.v4.config.FileServer>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<files.v4.config.FileServer>"
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfListFileServersApiResponseData) GetValue() interface{} {
	if "List<files.v4.config.FileServer>" == *p.Discriminator {
		return p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfListFileServersApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new([]FileServer)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "files.v4.config.FileServer" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<files.v4.config.FileServer>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<files.v4.config.FileServer>"
			return nil
		}
	}
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListFileServersApiResponseData"))
}

func (p *OneOfListFileServersApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<files.v4.config.FileServer>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfListFileServersApiResponseData")
}

type OneOfEnableWormComplianceApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfEnableWormComplianceApiResponseData() *OneOfEnableWormComplianceApiResponseData {
	p := new(OneOfEnableWormComplianceApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfEnableWormComplianceApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfEnableWormComplianceApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.TaskReference)
		}
		*p.oneOfType0 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfEnableWormComplianceApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfEnableWormComplianceApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.TaskReference)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfEnableWormComplianceApiResponseData"))
}

func (p *OneOfEnableWormComplianceApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfEnableWormComplianceApiResponseData")
}

type OneOfUpdateSnapshotScheduleApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    []import4.AppMessage   `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfUpdateSnapshotScheduleApiResponseData() *OneOfUpdateSnapshotScheduleApiResponseData {
	p := new(OneOfUpdateSnapshotScheduleApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateSnapshotScheduleApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateSnapshotScheduleApiResponseData is nil"))
	}
	switch v.(type) {
	case []import4.AppMessage:
		p.oneOfType0 = v.([]import4.AppMessage)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<files.v4.error.AppMessage>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<files.v4.error.AppMessage>"
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfUpdateSnapshotScheduleApiResponseData) GetValue() interface{} {
	if "List<files.v4.error.AppMessage>" == *p.Discriminator {
		return p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfUpdateSnapshotScheduleApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new([]import4.AppMessage)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "files.v4.error.AppMessage" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<files.v4.error.AppMessage>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<files.v4.error.AppMessage>"
			return nil
		}
	}
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateSnapshotScheduleApiResponseData"))
}

func (p *OneOfUpdateSnapshotScheduleApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<files.v4.error.AppMessage>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateSnapshotScheduleApiResponseData")
}

type OneOfUpdateVdiUserSessionApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfUpdateVdiUserSessionApiResponseData() *OneOfUpdateVdiUserSessionApiResponseData {
	p := new(OneOfUpdateVdiUserSessionApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateVdiUserSessionApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateVdiUserSessionApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.TaskReference)
		}
		*p.oneOfType0 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfUpdateVdiUserSessionApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfUpdateVdiUserSessionApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.TaskReference)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateVdiUserSessionApiResponseData"))
}

func (p *OneOfUpdateVdiUserSessionApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateVdiUserSessionApiResponseData")
}

type OneOfGetRecommendationApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *Recommendation        `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfGetRecommendationApiResponseData() *OneOfGetRecommendationApiResponseData {
	p := new(OneOfGetRecommendationApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetRecommendationApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetRecommendationApiResponseData is nil"))
	}
	switch v.(type) {
	case Recommendation:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(Recommendation)
		}
		*p.oneOfType0 = v.(Recommendation)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfGetRecommendationApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetRecommendationApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(Recommendation)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "files.v4.config.Recommendation" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(Recommendation)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetRecommendationApiResponseData"))
}

func (p *OneOfGetRecommendationApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetRecommendationApiResponseData")
}

type OneOfUpdateAntivirusServerApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfUpdateAntivirusServerApiResponseData() *OneOfUpdateAntivirusServerApiResponseData {
	p := new(OneOfUpdateAntivirusServerApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateAntivirusServerApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateAntivirusServerApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.TaskReference)
		}
		*p.oneOfType0 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.TaskReference)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
