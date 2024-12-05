/*
 * Generated file models/prism/v4/management/management_model.go.
 *
 * Product version: 4.0.1
 *
 * Part of the Nutanix Prism APIs
 *
 * (c) 2024 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module prism.v4.management of Nutanix Prism APIs
*/
package management

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	import4 "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/common/v1/config"
	import3 "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/common/v1/response"
	import1 "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/prism/v4/config"
	import2 "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/prism/v4/error"
	"time"
)

/*
Remote cluster specification required for registering an AOS cluster (Prism Element).
*/
type AOSRemoteClusterSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	RemoteCluster *RemoteClusterSpec `json:"remoteCluster"`
}

func (p *AOSRemoteClusterSpec) MarshalJSON() ([]byte, error) {
	type AOSRemoteClusterSpecProxy AOSRemoteClusterSpec
	return json.Marshal(struct {
		*AOSRemoteClusterSpecProxy
		RemoteCluster *RemoteClusterSpec `json:"remoteCluster,omitempty"`
	}{
		AOSRemoteClusterSpecProxy: (*AOSRemoteClusterSpecProxy)(p),
		RemoteCluster:             p.RemoteCluster,
	})
}

func NewAOSRemoteClusterSpec() *AOSRemoteClusterSpec {
	p := new(AOSRemoteClusterSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.management.AOSRemoteClusterSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The base model of S3 object store endpoint where domain manager is backed up.
*/
type AWSS3Config struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The bucket name of the object store endpoint where backup data of domain manager
	is to be stored.
	*/
	BucketName *string `json:"bucketName"`

	Credentials *AccessKeyCredentials `json:"credentials,omitempty"`
	/*
	  The region name of the object store endpoint where backup data of domain manager
	is stored.
	*/
	Region *string `json:"region,omitempty"`
}

func (p *AWSS3Config) MarshalJSON() ([]byte, error) {
	type AWSS3ConfigProxy AWSS3Config
	return json.Marshal(struct {
		*AWSS3ConfigProxy
		BucketName *string `json:"bucketName,omitempty"`
	}{
		AWSS3ConfigProxy: (*AWSS3ConfigProxy)(p),
		BucketName:       p.BucketName,
	})
}

func NewAWSS3Config() *AWSS3Config {
	p := new(AWSS3Config)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.management.AWSS3Config"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	p.Region = new(string)
	*p.Region = "us-east-1"

	return p
}

/*
Secret credentials model for the object store containing access key ID and secret access key.
*/
type AccessKeyCredentials struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Access key Id for the object store provided for backup target.
	*/
	AccessKeyId *string `json:"accessKeyId"`
	/*
	  Secret access key for the object store provided for backup target.
	*/
	SecretAccessKey *string `json:"secretAccessKey"`
}

func (p *AccessKeyCredentials) MarshalJSON() ([]byte, error) {
	type AccessKeyCredentialsProxy AccessKeyCredentials
	return json.Marshal(struct {
		*AccessKeyCredentialsProxy
		AccessKeyId     *string `json:"accessKeyId,omitempty"`
		SecretAccessKey *string `json:"secretAccessKey,omitempty"`
	}{
		AccessKeyCredentialsProxy: (*AccessKeyCredentialsProxy)(p),
		AccessKeyId:               p.AccessKeyId,
		SecretAccessKey:           p.SecretAccessKey,
	})
}

func NewAccessKeyCredentials() *AccessKeyCredentials {
	p := new(AccessKeyCredentials)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.management.AccessKeyCredentials"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /prism/v4.0/management/domain-managers/{extId}/$actions/add-root-certificate Post operation
*/
type AddRootCertificateApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfAddRootCertificateApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewAddRootCertificateApiResponse() *AddRootCertificateApiResponse {
	p := new(AddRootCertificateApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.management.AddRootCertificateApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *AddRootCertificateApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *AddRootCertificateApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfAddRootCertificateApiResponseData()
	}
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
Backup policy for the object store provided.
*/
type BackupPolicy struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  RPO interval in minutes at which the backup will be taken
	*/
	RpoInMinutes *int `json:"rpoInMinutes"`
}

func (p *BackupPolicy) MarshalJSON() ([]byte, error) {
	type BackupPolicyProxy BackupPolicy
	return json.Marshal(struct {
		*BackupPolicyProxy
		RpoInMinutes *int `json:"rpoInMinutes,omitempty"`
	}{
		BackupPolicyProxy: (*BackupPolicyProxy)(p),
		RpoInMinutes:      p.RpoInMinutes,
	})
}

func NewBackupPolicy() *BackupPolicy {
	p := new(BackupPolicy)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.management.BackupPolicy"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The backup target for the domain manager, which can be either a cluster or an object store.
*/
type BackupTarget struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Specifies a reason why the backup might have paused. This will be empty if the isBackupPaused field is false.
	*/
	BackupPauseReason *string `json:"backupPauseReason,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Whether the backup is paused on the given cluster or not.
	*/
	IsBackupPaused *bool `json:"isBackupPaused,omitempty"`
	/*
	  Represents the time when the domain manager was last synchronized or copied its
	configuration data to the backup target. This field is updated every 30 minutes.
	*/
	LastSyncTime *time.Time `json:"lastSyncTime,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import3.ApiLink `json:"links,omitempty"`
	/*

	 */
	LocationItemDiscriminator_ *string `json:"$locationItemDiscriminator,omitempty"`
	/*
	  Location of the backup target. For example, a cluster or an object store endpoint, such as AWS s3.
	*/
	Location *OneOfBackupTargetLocation `json:"location,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewBackupTarget() *BackupTarget {
	p := new(BackupTarget)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.management.BackupTarget"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *BackupTarget) GetLocation() interface{} {
	if nil == p.Location {
		return nil
	}
	return p.Location.GetValue()
}

func (p *BackupTarget) SetLocation(v interface{}) error {
	if nil == p.Location {
		p.Location = NewOneOfBackupTargetLocation()
	}
	e := p.Location.SetValue(v)
	if nil == e {
		if nil == p.LocationItemDiscriminator_ {
			p.LocationItemDiscriminator_ = new(string)
		}
		*p.LocationItemDiscriminator_ = *p.Location.Discriminator
	}
	return e
}

/*
Model which contains the information of backup cluster.
*/
type ClusterLocation struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Config *ClusterReference `json:"config"`
}

func (p *ClusterLocation) MarshalJSON() ([]byte, error) {
	type ClusterLocationProxy ClusterLocation
	return json.Marshal(struct {
		*ClusterLocationProxy
		Config *ClusterReference `json:"config,omitempty"`
	}{
		ClusterLocationProxy: (*ClusterLocationProxy)(p),
		Config:               p.Config,
	})
}

func NewClusterLocation() *ClusterLocation {
	p := new(ClusterLocation)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.management.ClusterLocation"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Cluster reference of the remote cluster to be connected.
*/
type ClusterReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Cluster UUID of a remote cluster.
	*/
	ExtId *string `json:"extId"`
	/*
	  Name of the cluster.
	*/
	Name *string `json:"name,omitempty"`
}

func (p *ClusterReference) MarshalJSON() ([]byte, error) {
	type ClusterReferenceProxy ClusterReference
	return json.Marshal(struct {
		*ClusterReferenceProxy
		ExtId *string `json:"extId,omitempty"`
	}{
		ClusterReferenceProxy: (*ClusterReferenceProxy)(p),
		ExtId:                 p.ExtId,
	})
}

func NewClusterReference() *ClusterReference {
	p := new(ClusterReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.management.ClusterReference"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Input specifications to perform the registration with a remote cluster entity.
*/
type ClusterRegistrationSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	RemoteClusterItemDiscriminator_ *string `json:"$remoteClusterItemDiscriminator,omitempty"`
	/*
	  Description of a remote cluster.
	*/
	RemoteCluster *OneOfClusterRegistrationSpecRemoteCluster `json:"remoteCluster"`
}

func (p *ClusterRegistrationSpec) MarshalJSON() ([]byte, error) {
	type ClusterRegistrationSpecProxy ClusterRegistrationSpec
	return json.Marshal(struct {
		*ClusterRegistrationSpecProxy
		RemoteCluster *OneOfClusterRegistrationSpecRemoteCluster `json:"remoteCluster,omitempty"`
	}{
		ClusterRegistrationSpecProxy: (*ClusterRegistrationSpecProxy)(p),
		RemoteCluster:                p.RemoteCluster,
	})
}

func NewClusterRegistrationSpec() *ClusterRegistrationSpec {
	p := new(ClusterRegistrationSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.management.ClusterRegistrationSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ClusterRegistrationSpec) GetRemoteCluster() interface{} {
	if nil == p.RemoteCluster {
		return nil
	}
	return p.RemoteCluster.GetValue()
}

func (p *ClusterRegistrationSpec) SetRemoteCluster(v interface{}) error {
	if nil == p.RemoteCluster {
		p.RemoteCluster = NewOneOfClusterRegistrationSpecRemoteCluster()
	}
	e := p.RemoteCluster.SetValue(v)
	if nil == e {
		if nil == p.RemoteClusterItemDiscriminator_ {
			p.RemoteClusterItemDiscriminator_ = new(string)
		}
		*p.RemoteClusterItemDiscriminator_ = *p.RemoteCluster.Discriminator
	}
	return e
}

/*
Type of cluster to be connected:
- DOMAIN_MANAGER : Domain manager (Prism Central) instance
- AOS : Prism Element cluster instance
*/
type ClusterType int

const (
	CLUSTERTYPE_UNKNOWN        ClusterType = 0
	CLUSTERTYPE_REDACTED       ClusterType = 1
	CLUSTERTYPE_DOMAIN_MANAGER ClusterType = 2
	CLUSTERTYPE_AOS            ClusterType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ClusterType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"DOMAIN_MANAGER",
		"AOS",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ClusterType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"DOMAIN_MANAGER",
		"AOS",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ClusterType) index(name string) ClusterType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"DOMAIN_MANAGER",
		"AOS",
	}
	for idx := range names {
		if names[idx] == name {
			return ClusterType(idx)
		}
	}
	return CLUSTERTYPE_UNKNOWN
}

func (e *ClusterType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ClusterType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ClusterType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ClusterType) Ref() *ClusterType {
	return &e
}

type ClusterUnregistrationSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Cluster UUID of a remote cluster.
	*/
	ExtId *string `json:"extId"`
	/*
	  Name of the cluster.
	*/
	Name *string `json:"name,omitempty"`
}

func NewClusterUnregistrationSpec() *ClusterUnregistrationSpec {
	p := new(ClusterUnregistrationSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.management.ClusterUnregistrationSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /prism/v4.0/management/domain-managers/{extId}/$actions/configure-connection Post operation
*/
type ConfigureConnectionApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfConfigureConnectionApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewConfigureConnectionApiResponse() *ConfigureConnectionApiResponse {
	p := new(ConfigureConnectionApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.management.ConfigureConnectionApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ConfigureConnectionApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ConfigureConnectionApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfConfigureConnectionApiResponseData()
	}
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
Payload to configure connection endpoint. It contains the details of the remote cluster you are connecting to.
*/
type ConnectionConfigurationSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	RemoteCluster *RemoteCluster `json:"remoteCluster"`
}

func (p *ConnectionConfigurationSpec) MarshalJSON() ([]byte, error) {
	type ConnectionConfigurationSpecProxy ConnectionConfigurationSpec
	return json.Marshal(struct {
		*ConnectionConfigurationSpecProxy
		RemoteCluster *RemoteCluster `json:"remoteCluster,omitempty"`
	}{
		ConnectionConfigurationSpecProxy: (*ConnectionConfigurationSpecProxy)(p),
		RemoteCluster:                    p.RemoteCluster,
	})
}

func NewConnectionConfigurationSpec() *ConnectionConfigurationSpec {
	p := new(ConnectionConfigurationSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.management.ConnectionConfigurationSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Payload to unconfigure connection endpoint. It contains the details of the remote cluster to be disconnected.
*/
type ConnectionUnconfigurationSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	RemoteCluster *ClusterReference `json:"remoteCluster"`
}

func (p *ConnectionUnconfigurationSpec) MarshalJSON() ([]byte, error) {
	type ConnectionUnconfigurationSpecProxy ConnectionUnconfigurationSpec
	return json.Marshal(struct {
		*ConnectionUnconfigurationSpecProxy
		RemoteCluster *ClusterReference `json:"remoteCluster,omitempty"`
	}{
		ConnectionUnconfigurationSpecProxy: (*ConnectionUnconfigurationSpecProxy)(p),
		RemoteCluster:                      p.RemoteCluster,
	})
}

func NewConnectionUnconfigurationSpec() *ConnectionUnconfigurationSpec {
	p := new(ConnectionUnconfigurationSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.management.ConnectionUnconfigurationSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /prism/v4.0/management/domain-managers/{domainManagerExtId}/backup-targets Post operation
*/
type CreateBackupTargetApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateBackupTargetApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewCreateBackupTargetApiResponse() *CreateBackupTargetApiResponse {
	p := new(CreateBackupTargetApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.management.CreateBackupTargetApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreateBackupTargetApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreateBackupTargetApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreateBackupTargetApiResponseData()
	}
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
REST response for all response codes in API path /prism/v4.0/management/restore-sources Post operation
*/
type CreateRestoreSourceApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateRestoreSourceApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewCreateRestoreSourceApiResponse() *CreateRestoreSourceApiResponse {
	p := new(CreateRestoreSourceApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.management.CreateRestoreSourceApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreateRestoreSourceApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreateRestoreSourceApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreateRestoreSourceApiResponseData()
	}
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
Credentials to connect to a remote cluster.
*/
type Credentials struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Authentication *import4.BasicAuth `json:"authentication"`
}

func (p *Credentials) MarshalJSON() ([]byte, error) {
	type CredentialsProxy Credentials
	return json.Marshal(struct {
		*CredentialsProxy
		Authentication *import4.BasicAuth `json:"authentication,omitempty"`
	}{
		CredentialsProxy: (*CredentialsProxy)(p),
		Authentication:   p.Authentication,
	})
}

func NewCredentials() *Credentials {
	p := new(Credentials)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.management.Credentials"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /prism/v4.0/management/domain-managers/{domainManagerExtId}/backup-targets/{extId} Delete operation
*/
type DeleteBackupTargetApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteBackupTargetApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDeleteBackupTargetApiResponse() *DeleteBackupTargetApiResponse {
	p := new(DeleteBackupTargetApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.management.DeleteBackupTargetApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteBackupTargetApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteBackupTargetApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteBackupTargetApiResponseData()
	}
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
REST response for all response codes in API path /prism/v4.0/management/restore-sources/{extId} Delete operation
*/
type DeleteRestoreSourceApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteRestoreSourceApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDeleteRestoreSourceApiResponse() *DeleteRestoreSourceApiResponse {
	p := new(DeleteRestoreSourceApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.management.DeleteRestoreSourceApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteRestoreSourceApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteRestoreSourceApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteRestoreSourceApiResponseData()
	}
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
Enum denoting whether the domain manager (Prism Central) instance is reachable with its physical address or reachable through the My Nutanix portal.
Based on the above description, the allowed enum values are:
- ONPREM: Domain manager (Prism Central) reachable on it's physical address.
- NUTANIX_HOSTED_CLOUD: Domain manager (Prism Central) reachable through My Nutanix portal.
*/
type DomainManagerCloudType int

const (
	DOMAINMANAGERCLOUDTYPE_UNKNOWN              DomainManagerCloudType = 0
	DOMAINMANAGERCLOUDTYPE_REDACTED             DomainManagerCloudType = 1
	DOMAINMANAGERCLOUDTYPE_ONPREM_CLOUD         DomainManagerCloudType = 2
	DOMAINMANAGERCLOUDTYPE_NUTANIX_HOSTED_CLOUD DomainManagerCloudType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *DomainManagerCloudType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ONPREM_CLOUD",
		"NUTANIX_HOSTED_CLOUD",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e DomainManagerCloudType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ONPREM_CLOUD",
		"NUTANIX_HOSTED_CLOUD",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *DomainManagerCloudType) index(name string) DomainManagerCloudType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ONPREM_CLOUD",
		"NUTANIX_HOSTED_CLOUD",
	}
	for idx := range names {
		if names[idx] == name {
			return DomainManagerCloudType(idx)
		}
	}
	return DOMAINMANAGERCLOUDTYPE_UNKNOWN
}

func (e *DomainManagerCloudType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for DomainManagerCloudType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *DomainManagerCloudType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e DomainManagerCloudType) Ref() *DomainManagerCloudType {
	return &e
}

/*
Remote cluster specification required for registering a domain manager (Prism Central).
*/
type DomainManagerRemoteClusterSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	CloudType *DomainManagerCloudType `json:"cloudType"`

	RemoteCluster *RemoteClusterSpec `json:"remoteCluster"`
}

func (p *DomainManagerRemoteClusterSpec) MarshalJSON() ([]byte, error) {
	type DomainManagerRemoteClusterSpecProxy DomainManagerRemoteClusterSpec
	return json.Marshal(struct {
		*DomainManagerRemoteClusterSpecProxy
		CloudType     *DomainManagerCloudType `json:"cloudType,omitempty"`
		RemoteCluster *RemoteClusterSpec      `json:"remoteCluster,omitempty"`
	}{
		DomainManagerRemoteClusterSpecProxy: (*DomainManagerRemoteClusterSpecProxy)(p),
		CloudType:                           p.CloudType,
		RemoteCluster:                       p.RemoteCluster,
	})
}

func NewDomainManagerRemoteClusterSpec() *DomainManagerRemoteClusterSpec {
	p := new(DomainManagerRemoteClusterSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.management.DomainManagerRemoteClusterSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /prism/v4.0/management/domain-managers/{domainManagerExtId}/backup-targets/{extId} Get operation
*/
type GetBackupTargetApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetBackupTargetApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetBackupTargetApiResponse() *GetBackupTargetApiResponse {
	p := new(GetBackupTargetApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.management.GetBackupTargetApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetBackupTargetApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetBackupTargetApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetBackupTargetApiResponseData()
	}
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
REST response for all response codes in API path /prism/v4.0/management/restore-sources/{restoreSourceExtId}/restorable-domain-managers/{restorableDomainManagerExtId}/restore-points/{extId} Get operation
*/
type GetRestorePointApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetRestorePointApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetRestorePointApiResponse() *GetRestorePointApiResponse {
	p := new(GetRestorePointApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.management.GetRestorePointApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetRestorePointApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetRestorePointApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetRestorePointApiResponseData()
	}
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
REST response for all response codes in API path /prism/v4.0/management/restore-sources/{extId} Get operation
*/
type GetRestoreSourceApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetRestoreSourceApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetRestoreSourceApiResponse() *GetRestoreSourceApiResponse {
	p := new(GetRestoreSourceApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.management.GetRestoreSourceApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetRestoreSourceApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetRestoreSourceApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetRestoreSourceApiResponseData()
	}
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
REST response for all response codes in API path /prism/v4.0/management/domain-managers/{domainManagerExtId}/backup-targets Get operation
*/
type ListBackupTargetsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListBackupTargetsApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListBackupTargetsApiResponse() *ListBackupTargetsApiResponse {
	p := new(ListBackupTargetsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.management.ListBackupTargetsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListBackupTargetsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListBackupTargetsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListBackupTargetsApiResponseData()
	}
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
REST response for all response codes in API path /prism/v4.0/management/restore-sources/{restoreSourceExtId}/restorable-domain-managers Get operation
*/
type ListRestorableDomainManagersApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListRestorableDomainManagersApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListRestorableDomainManagersApiResponse() *ListRestorableDomainManagersApiResponse {
	p := new(ListRestorableDomainManagersApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.management.ListRestorableDomainManagersApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListRestorableDomainManagersApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListRestorableDomainManagersApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListRestorableDomainManagersApiResponseData()
	}
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
REST response for all response codes in API path /prism/v4.0/management/restore-sources/{restoreSourceExtId}/restorable-domain-managers/{restorableDomainManagerExtId}/restore-points Get operation
*/
type ListRestorePointsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListRestorePointsApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListRestorePointsApiResponse() *ListRestorePointsApiResponse {
	p := new(ListRestorePointsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.management.ListRestorePointsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListRestorePointsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListRestorePointsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListRestorePointsApiResponseData()
	}
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
Model which contains the information of the object store endpoint where
backup is present. It contains information like objectstore
endpoint address, endpoint flavour and last sync timestamp.
*/
type ObjectStoreLocation struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	BackupPolicy *BackupPolicy `json:"backupPolicy,omitempty"`

	ProviderConfig *AWSS3Config `json:"providerConfig"`
}

func (p *ObjectStoreLocation) MarshalJSON() ([]byte, error) {
	type ObjectStoreLocationProxy ObjectStoreLocation
	return json.Marshal(struct {
		*ObjectStoreLocationProxy
		ProviderConfig *AWSS3Config `json:"providerConfig,omitempty"`
	}{
		ObjectStoreLocationProxy: (*ObjectStoreLocationProxy)(p),
		ProviderConfig:           p.ProviderConfig,
	})
}

func NewObjectStoreLocation() *ObjectStoreLocation {
	p := new(ObjectStoreLocation)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.management.ObjectStoreLocation"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /prism/v4.0/management/domain-managers/{extId}/$actions/register Post operation
*/
type RegisterApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfRegisterApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewRegisterApiResponse() *RegisterApiResponse {
	p := new(RegisterApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.management.RegisterApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *RegisterApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *RegisterApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfRegisterApiResponseData()
	}
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
This includes the attributes of a remote cluster, such as the cluster name, cluster type, and address details. The address details comprise the external address (either a virtual IP or FQDN), the port for incoming connections, and the internal addresses (node IP addresses).
*/
type RemoteCluster struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`

	ExternalAddress *import4.IPAddressOrFQDN `json:"externalAddress"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import3.ApiLink `json:"links,omitempty"`
	/*
	  Cluster name of a remote cluster.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  Node IP addresses of a registered cluster.
	*/
	NodeIpAddresses []import4.IPAddressOrFQDN `json:"nodeIpAddresses,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *RemoteCluster) MarshalJSON() ([]byte, error) {
	type RemoteClusterProxy RemoteCluster
	return json.Marshal(struct {
		*RemoteClusterProxy
		ExternalAddress *import4.IPAddressOrFQDN `json:"externalAddress,omitempty"`
	}{
		RemoteClusterProxy: (*RemoteClusterProxy)(p),
		ExternalAddress:    p.ExternalAddress,
	})
}

func NewRemoteCluster() *RemoteCluster {
	p := new(RemoteCluster)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.management.RemoteCluster"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Address configuration of a remote cluster. It requires the address of the remote, that is an IP or domain name along with the basic authentication credentials.
*/
type RemoteClusterSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Address *import4.IPAddressOrFQDN `json:"address"`

	Credentials *Credentials `json:"credentials"`
}

func (p *RemoteClusterSpec) MarshalJSON() ([]byte, error) {
	type RemoteClusterSpecProxy RemoteClusterSpec
	return json.Marshal(struct {
		*RemoteClusterSpecProxy
		Address     *import4.IPAddressOrFQDN `json:"address,omitempty"`
		Credentials *Credentials             `json:"credentials,omitempty"`
	}{
		RemoteClusterSpecProxy: (*RemoteClusterSpecProxy)(p),
		Address:                p.Address,
		Credentials:            p.Credentials,
	})
}

func NewRemoteClusterSpec() *RemoteClusterSpec {
	p := new(RemoteClusterSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.management.RemoteClusterSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /prism/v4.0/management/domain-managers/{extId}/$actions/remove-root-certificate Post operation
*/
type RemoveRootCertificateApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfRemoveRootCertificateApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewRemoveRootCertificateApiResponse() *RemoveRootCertificateApiResponse {
	p := new(RemoveRootCertificateApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.management.RemoveRootCertificateApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *RemoveRootCertificateApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *RemoveRootCertificateApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfRemoveRootCertificateApiResponseData()
	}
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
The backup target for the domain manager, which can be either a cluster or an object store.
*/
type RestorableDomainManager struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Config *import1.DomainManagerClusterConfig `json:"config"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  The external identifier of the cluster hosting the domain manager (Prism Central) instance.
	*/
	HostingClusterExtId *string `json:"hostingClusterExtId,omitempty"`
	/*
	  Boolean value indicating if the domain manager (Prism Central) is registered with the hosting cluster, that is, Prism Element.
	*/
	IsRegisteredWithHostingCluster *bool `json:"isRegisteredWithHostingCluster,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import3.ApiLink `json:"links,omitempty"`

	Network *import1.DomainManagerNetwork `json:"network"`
	/*
	  Domain manager (Prism Central) nodes external identifier.
	*/
	NodeExtIds []string `json:"nodeExtIds,omitempty"`
	/*
	  This configuration enables Prism Central to be deployed in scale-out mode.
	*/
	ShouldEnableHighAvailability *bool `json:"shouldEnableHighAvailability,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *RestorableDomainManager) MarshalJSON() ([]byte, error) {
	type RestorableDomainManagerProxy RestorableDomainManager
	return json.Marshal(struct {
		*RestorableDomainManagerProxy
		Config  *import1.DomainManagerClusterConfig `json:"config,omitempty"`
		Network *import1.DomainManagerNetwork       `json:"network,omitempty"`
	}{
		RestorableDomainManagerProxy: (*RestorableDomainManagerProxy)(p),
		Config:                       p.Config,
		Network:                      p.Network,
	})
}

func NewRestorableDomainManager() *RestorableDomainManager {
	p := new(RestorableDomainManager)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.management.RestorableDomainManager"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	p.ShouldEnableHighAvailability = new(bool)
	*p.ShouldEnableHighAvailability = false

	return p
}

/*
REST response for all response codes in API path /prism/v4.0/management/restore-sources/{restoreSourceExtId}/restorable-domain-managers/{restorableDomainManagerExtId}/restore-points/{extId}/$actions/restore Post operation
*/
type RestoreApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfRestoreApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewRestoreApiResponse() *RestoreApiResponse {
	p := new(RestoreApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.management.RestoreApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *RestoreApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *RestoreApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfRestoreApiResponseData()
	}
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
Restore point represents the state of domain manager at a specific moment. It can be used to restore the domain manager to that state in cases such as disaster. Each restore point captures key details i.e creation time and domain manager config necessary for restoring the system to a prior state.
*/
type RestorePoint struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The UTC date and time in ISO-8601 format when the Restore point was created.
	*/
	CreationTime *time.Time `json:"creationTime,omitempty"`

	DomainManager *import1.DomainManager `json:"domainManager,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import3.ApiLink `json:"links,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewRestorePoint() *RestorePoint {
	p := new(RestorePoint)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.management.RestorePoint"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The backup target for the domain manager, which can be either a cluster or an object store.
*/
type RestoreSource struct {
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
	Links []import3.ApiLink `json:"links,omitempty"`
	/*

	 */
	LocationItemDiscriminator_ *string `json:"$locationItemDiscriminator,omitempty"`
	/*
	  Location of the backup target. For example, a cluster or an object store endpoint, such as AWS s3.
	*/
	Location *OneOfRestoreSourceLocation `json:"location,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewRestoreSource() *RestoreSource {
	p := new(RestoreSource)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.management.RestoreSource"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *RestoreSource) GetLocation() interface{} {
	if nil == p.Location {
		return nil
	}
	return p.Location.GetValue()
}

func (p *RestoreSource) SetLocation(v interface{}) error {
	if nil == p.Location {
		p.Location = NewOneOfRestoreSourceLocation()
	}
	e := p.Location.SetValue(v)
	if nil == e {
		if nil == p.LocationItemDiscriminator_ {
			p.LocationItemDiscriminator_ = new(string)
		}
		*p.LocationItemDiscriminator_ = *p.Location.Discriminator
	}
	return e
}

/*
Restore specs of a domain manager to be deployed on cluster during restore operation.
*/
type RestoreSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	DomainManager *import1.DomainManager `json:"domainManager"`
}

func (p *RestoreSpec) MarshalJSON() ([]byte, error) {
	type RestoreSpecProxy RestoreSpec
	return json.Marshal(struct {
		*RestoreSpecProxy
		DomainManager *import1.DomainManager `json:"domainManager,omitempty"`
	}{
		RestoreSpecProxy: (*RestoreSpecProxy)(p),
		DomainManager:    p.DomainManager,
	})
}

func NewRestoreSpec() *RestoreSpec {
	p := new(RestoreSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.management.RestoreSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Payload to delete root certificate. It contains the external identifier of the remote cluster.
*/
type RootCertRemoveSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The external identifier of the domain manager (Prism Central) resource.
	*/
	ClusterExtId *string `json:"clusterExtId"`
}

func (p *RootCertRemoveSpec) MarshalJSON() ([]byte, error) {
	type RootCertRemoveSpecProxy RootCertRemoveSpec
	return json.Marshal(struct {
		*RootCertRemoveSpecProxy
		ClusterExtId *string `json:"clusterExtId,omitempty"`
	}{
		RootCertRemoveSpecProxy: (*RootCertRemoveSpecProxy)(p),
		ClusterExtId:            p.ClusterExtId,
	})
}

func NewRootCertRemoveSpec() *RootCertRemoveSpec {
	p := new(RootCertRemoveSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.management.RootCertRemoveSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Payload for the add root certificate endpoint. It contains the root certificate of the remote cluster.
*/
type RootCertificateAddSpec struct {
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
	Links []import3.ApiLink `json:"links,omitempty"`
	/*
	  The root certificate of the cluster.
	*/
	RootCertificate *string `json:"rootCertificate"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *RootCertificateAddSpec) MarshalJSON() ([]byte, error) {
	type RootCertificateAddSpecProxy RootCertificateAddSpec
	return json.Marshal(struct {
		*RootCertificateAddSpecProxy
		RootCertificate *string `json:"rootCertificate,omitempty"`
	}{
		RootCertificateAddSpecProxy: (*RootCertificateAddSpecProxy)(p),
		RootCertificate:             p.RootCertificate,
	})
}

func NewRootCertificateAddSpec() *RootCertificateAddSpec {
	p := new(RootCertificateAddSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.management.RootCertificateAddSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The endpoint of the object store s3 where domain manager is backed up.
*/
type S3Config struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The bucket name of the object store endpoint where backup data of domain manager
	is to be stored.
	*/
	BucketName *string `json:"bucketName"`
	/*
	  The region name of the object store endpoint where backup data of domain manager
	is stored.
	*/
	Region *string `json:"region,omitempty"`
}

func (p *S3Config) MarshalJSON() ([]byte, error) {
	type S3ConfigProxy S3Config
	return json.Marshal(struct {
		*S3ConfigProxy
		BucketName *string `json:"bucketName,omitempty"`
	}{
		S3ConfigProxy: (*S3ConfigProxy)(p),
		BucketName:    p.BucketName,
	})
}

func NewS3Config() *S3Config {
	p := new(S3Config)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.management.S3Config"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	p.Region = new(string)
	*p.Region = "us-east-1"

	return p
}

/*
REST response for all response codes in API path /prism/v4.0/management/domain-managers/{extId}/$actions/unconfigure-connection Post operation
*/
type UnconfigureConnectionApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUnconfigureConnectionApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewUnconfigureConnectionApiResponse() *UnconfigureConnectionApiResponse {
	p := new(UnconfigureConnectionApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.management.UnconfigureConnectionApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UnconfigureConnectionApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UnconfigureConnectionApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUnconfigureConnectionApiResponseData()
	}
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
REST response for all response codes in API path /prism/v4.0/management/domain-managers/{extId}/$actions/unregister Post operation
*/
type UnregisterApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUnregisterApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewUnregisterApiResponse() *UnregisterApiResponse {
	p := new(UnregisterApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.management.UnregisterApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UnregisterApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UnregisterApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUnregisterApiResponseData()
	}
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
REST response for all response codes in API path /prism/v4.0/management/domain-managers/{domainManagerExtId}/backup-targets/{extId} Put operation
*/
type UpdateBackupTargetApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateBackupTargetApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewUpdateBackupTargetApiResponse() *UpdateBackupTargetApiResponse {
	p := new(UpdateBackupTargetApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.management.UpdateBackupTargetApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdateBackupTargetApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdateBackupTargetApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdateBackupTargetApiResponseData()
	}
	e := p.Data.SetValue(v)
	if nil == e {
		if nil == p.DataItemDiscriminator_ {
			p.DataItemDiscriminator_ = new(string)
		}
		*p.DataItemDiscriminator_ = *p.Data.Discriminator
	}
	return e
}

type OneOfConfigureConnectionApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfConfigureConnectionApiResponseData() *OneOfConfigureConnectionApiResponseData {
	p := new(OneOfConfigureConnectionApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfConfigureConnectionApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfConfigureConnectionApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import1.TaskReference)
		}
		*p.oneOfType2001 = v.(import1.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
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

func (p *OneOfConfigureConnectionApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfConfigureConnectionApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import1.TaskReference)
			}
			*p.oneOfType2001 = *vOneOfType2001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2001.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2001.ObjectType_
			return nil
		}
	}
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "prism.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfConfigureConnectionApiResponseData"))
}

func (p *OneOfConfigureConnectionApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfConfigureConnectionApiResponseData")
}

type OneOfBackupTargetLocation struct {
	Discriminator *string              `json:"-"`
	ObjectType_   *string              `json:"-"`
	oneOfType0    *ClusterLocation     `json:"-"`
	oneOfType1    *ObjectStoreLocation `json:"-"`
}

func NewOneOfBackupTargetLocation() *OneOfBackupTargetLocation {
	p := new(OneOfBackupTargetLocation)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfBackupTargetLocation) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfBackupTargetLocation is nil"))
	}
	switch v.(type) {
	case ClusterLocation:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(ClusterLocation)
		}
		*p.oneOfType0 = v.(ClusterLocation)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case ObjectStoreLocation:
		if nil == p.oneOfType1 {
			p.oneOfType1 = new(ObjectStoreLocation)
		}
		*p.oneOfType1 = v.(ObjectStoreLocation)
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

func (p *OneOfBackupTargetLocation) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1
	}
	return nil
}

func (p *OneOfBackupTargetLocation) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(ClusterLocation)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.management.ClusterLocation" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(ClusterLocation)
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
	vOneOfType1 := new(ObjectStoreLocation)
	if err := json.Unmarshal(b, vOneOfType1); err == nil {
		if "prism.v4.management.ObjectStoreLocation" == *vOneOfType1.ObjectType_ {
			if nil == p.oneOfType1 {
				p.oneOfType1 = new(ObjectStoreLocation)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfBackupTargetLocation"))
}

func (p *OneOfBackupTargetLocation) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	return nil, errors.New("No value to marshal for OneOfBackupTargetLocation")
}

type OneOfListBackupTargetsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
	oneOfType0    []BackupTarget         `json:"-"`
}

func NewOneOfListBackupTargetsApiResponseData() *OneOfListBackupTargetsApiResponseData {
	p := new(OneOfListBackupTargetsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListBackupTargetsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListBackupTargetsApiResponseData is nil"))
	}
	switch v.(type) {
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case []BackupTarget:
		p.oneOfType0 = v.([]BackupTarget)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<prism.v4.management.BackupTarget>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<prism.v4.management.BackupTarget>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListBackupTargetsApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<prism.v4.management.BackupTarget>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfListBackupTargetsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "prism.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	vOneOfType0 := new([]BackupTarget)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "prism.v4.management.BackupTarget" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<prism.v4.management.BackupTarget>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<prism.v4.management.BackupTarget>"
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListBackupTargetsApiResponseData"))
}

func (p *OneOfListBackupTargetsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<prism.v4.management.BackupTarget>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfListBackupTargetsApiResponseData")
}

type OneOfListRestorePointsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
	oneOfType2001 []RestorePoint         `json:"-"`
}

func NewOneOfListRestorePointsApiResponseData() *OneOfListRestorePointsApiResponseData {
	p := new(OneOfListRestorePointsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListRestorePointsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListRestorePointsApiResponseData is nil"))
	}
	switch v.(type) {
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case []RestorePoint:
		p.oneOfType2001 = v.([]RestorePoint)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<prism.v4.management.RestorePoint>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<prism.v4.management.RestorePoint>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListRestorePointsApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<prism.v4.management.RestorePoint>" == *p.Discriminator {
		return p.oneOfType2001
	}
	return nil
}

func (p *OneOfListRestorePointsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "prism.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	vOneOfType2001 := new([]RestorePoint)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if len(*vOneOfType2001) == 0 || "prism.v4.management.RestorePoint" == *((*vOneOfType2001)[0].ObjectType_) {
			p.oneOfType2001 = *vOneOfType2001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<prism.v4.management.RestorePoint>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<prism.v4.management.RestorePoint>"
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListRestorePointsApiResponseData"))
}

func (p *OneOfListRestorePointsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<prism.v4.management.RestorePoint>" == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfListRestorePointsApiResponseData")
}

type OneOfRegisterApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfRegisterApiResponseData() *OneOfRegisterApiResponseData {
	p := new(OneOfRegisterApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfRegisterApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfRegisterApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import1.TaskReference)
		}
		*p.oneOfType2001 = v.(import1.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
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

func (p *OneOfRegisterApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfRegisterApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import1.TaskReference)
			}
			*p.oneOfType2001 = *vOneOfType2001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2001.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2001.ObjectType_
			return nil
		}
	}
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "prism.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfRegisterApiResponseData"))
}

func (p *OneOfRegisterApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfRegisterApiResponseData")
}

type OneOfRestoreApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfRestoreApiResponseData() *OneOfRestoreApiResponseData {
	p := new(OneOfRestoreApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfRestoreApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfRestoreApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import1.TaskReference)
		}
		*p.oneOfType0 = v.(import1.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
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

func (p *OneOfRestoreApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfRestoreApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import1.TaskReference)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "prism.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfRestoreApiResponseData"))
}

func (p *OneOfRestoreApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfRestoreApiResponseData")
}

type OneOfDeleteBackupTargetApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfDeleteBackupTargetApiResponseData() *OneOfDeleteBackupTargetApiResponseData {
	p := new(OneOfDeleteBackupTargetApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteBackupTargetApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteBackupTargetApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import1.TaskReference)
		}
		*p.oneOfType0 = v.(import1.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
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

func (p *OneOfDeleteBackupTargetApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDeleteBackupTargetApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import1.TaskReference)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "prism.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteBackupTargetApiResponseData"))
}

func (p *OneOfDeleteBackupTargetApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteBackupTargetApiResponseData")
}

type OneOfGetRestorePointApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *RestorePoint          `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfGetRestorePointApiResponseData() *OneOfGetRestorePointApiResponseData {
	p := new(OneOfGetRestorePointApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetRestorePointApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetRestorePointApiResponseData is nil"))
	}
	switch v.(type) {
	case RestorePoint:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(RestorePoint)
		}
		*p.oneOfType2001 = v.(RestorePoint)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
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

func (p *OneOfGetRestorePointApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetRestorePointApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(RestorePoint)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.management.RestorePoint" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(RestorePoint)
			}
			*p.oneOfType2001 = *vOneOfType2001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2001.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2001.ObjectType_
			return nil
		}
	}
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "prism.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetRestorePointApiResponseData"))
}

func (p *OneOfGetRestorePointApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetRestorePointApiResponseData")
}

type OneOfRemoveRootCertificateApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfRemoveRootCertificateApiResponseData() *OneOfRemoveRootCertificateApiResponseData {
	p := new(OneOfRemoveRootCertificateApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfRemoveRootCertificateApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfRemoveRootCertificateApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import1.TaskReference)
		}
		*p.oneOfType2001 = v.(import1.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
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

func (p *OneOfRemoveRootCertificateApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfRemoveRootCertificateApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import1.TaskReference)
			}
			*p.oneOfType2001 = *vOneOfType2001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2001.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2001.ObjectType_
			return nil
		}
	}
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "prism.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfRemoveRootCertificateApiResponseData"))
}

func (p *OneOfRemoveRootCertificateApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfRemoveRootCertificateApiResponseData")
}

type OneOfUnconfigureConnectionApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfUnconfigureConnectionApiResponseData() *OneOfUnconfigureConnectionApiResponseData {
	p := new(OneOfUnconfigureConnectionApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUnconfigureConnectionApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUnconfigureConnectionApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import1.TaskReference)
		}
		*p.oneOfType2001 = v.(import1.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
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

func (p *OneOfUnconfigureConnectionApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfUnconfigureConnectionApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import1.TaskReference)
			}
			*p.oneOfType2001 = *vOneOfType2001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2001.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2001.ObjectType_
			return nil
		}
	}
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "prism.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUnconfigureConnectionApiResponseData"))
}

func (p *OneOfUnconfigureConnectionApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfUnconfigureConnectionApiResponseData")
}

type OneOfAddRootCertificateApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfAddRootCertificateApiResponseData() *OneOfAddRootCertificateApiResponseData {
	p := new(OneOfAddRootCertificateApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfAddRootCertificateApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfAddRootCertificateApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import1.TaskReference)
		}
		*p.oneOfType2001 = v.(import1.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
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

func (p *OneOfAddRootCertificateApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfAddRootCertificateApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import1.TaskReference)
			}
			*p.oneOfType2001 = *vOneOfType2001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2001.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2001.ObjectType_
			return nil
		}
	}
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "prism.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfAddRootCertificateApiResponseData"))
}

func (p *OneOfAddRootCertificateApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfAddRootCertificateApiResponseData")
}

type OneOfDeleteRestoreSourceApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType1    *interface{}           `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfDeleteRestoreSourceApiResponseData() *OneOfDeleteRestoreSourceApiResponseData {
	p := new(OneOfDeleteRestoreSourceApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteRestoreSourceApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteRestoreSourceApiResponseData is nil"))
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
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
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

func (p *OneOfDeleteRestoreSourceApiResponseData) GetValue() interface{} {
	if "EMPTY" == *p.Discriminator {
		return *p.oneOfType1
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDeleteRestoreSourceApiResponseData) UnmarshalJSON(b []byte) error {
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "prism.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteRestoreSourceApiResponseData"))
}

func (p *OneOfDeleteRestoreSourceApiResponseData) MarshalJSON() ([]byte, error) {
	if "EMPTY" == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteRestoreSourceApiResponseData")
}

type OneOfUnregisterApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfUnregisterApiResponseData() *OneOfUnregisterApiResponseData {
	p := new(OneOfUnregisterApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUnregisterApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUnregisterApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import1.TaskReference)
		}
		*p.oneOfType2001 = v.(import1.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
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

func (p *OneOfUnregisterApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfUnregisterApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import1.TaskReference)
			}
			*p.oneOfType2001 = *vOneOfType2001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2001.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2001.ObjectType_
			return nil
		}
	}
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "prism.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUnregisterApiResponseData"))
}

func (p *OneOfUnregisterApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfUnregisterApiResponseData")
}

type OneOfClusterRegistrationSpecRemoteCluster struct {
	Discriminator  *string                         `json:"-"`
	ObjectType_    *string                         `json:"-"`
	oneOfType12002 *AOSRemoteClusterSpec           `json:"-"`
	oneOfType12003 *ClusterReference               `json:"-"`
	oneOfType12001 *DomainManagerRemoteClusterSpec `json:"-"`
}

func NewOneOfClusterRegistrationSpecRemoteCluster() *OneOfClusterRegistrationSpecRemoteCluster {
	p := new(OneOfClusterRegistrationSpecRemoteCluster)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfClusterRegistrationSpecRemoteCluster) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfClusterRegistrationSpecRemoteCluster is nil"))
	}
	switch v.(type) {
	case AOSRemoteClusterSpec:
		if nil == p.oneOfType12002 {
			p.oneOfType12002 = new(AOSRemoteClusterSpec)
		}
		*p.oneOfType12002 = v.(AOSRemoteClusterSpec)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType12002.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType12002.ObjectType_
	case ClusterReference:
		if nil == p.oneOfType12003 {
			p.oneOfType12003 = new(ClusterReference)
		}
		*p.oneOfType12003 = v.(ClusterReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType12003.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType12003.ObjectType_
	case DomainManagerRemoteClusterSpec:
		if nil == p.oneOfType12001 {
			p.oneOfType12001 = new(DomainManagerRemoteClusterSpec)
		}
		*p.oneOfType12001 = v.(DomainManagerRemoteClusterSpec)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType12001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType12001.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfClusterRegistrationSpecRemoteCluster) GetValue() interface{} {
	if p.oneOfType12002 != nil && *p.oneOfType12002.ObjectType_ == *p.Discriminator {
		return *p.oneOfType12002
	}
	if p.oneOfType12003 != nil && *p.oneOfType12003.ObjectType_ == *p.Discriminator {
		return *p.oneOfType12003
	}
	if p.oneOfType12001 != nil && *p.oneOfType12001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType12001
	}
	return nil
}

func (p *OneOfClusterRegistrationSpecRemoteCluster) UnmarshalJSON(b []byte) error {
	vOneOfType12002 := new(AOSRemoteClusterSpec)
	if err := json.Unmarshal(b, vOneOfType12002); err == nil {
		if "prism.v4.management.AOSRemoteClusterSpec" == *vOneOfType12002.ObjectType_ {
			if nil == p.oneOfType12002 {
				p.oneOfType12002 = new(AOSRemoteClusterSpec)
			}
			*p.oneOfType12002 = *vOneOfType12002
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType12002.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType12002.ObjectType_
			return nil
		}
	}
	vOneOfType12003 := new(ClusterReference)
	if err := json.Unmarshal(b, vOneOfType12003); err == nil {
		if "prism.v4.management.ClusterReference" == *vOneOfType12003.ObjectType_ {
			if nil == p.oneOfType12003 {
				p.oneOfType12003 = new(ClusterReference)
			}
			*p.oneOfType12003 = *vOneOfType12003
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType12003.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType12003.ObjectType_
			return nil
		}
	}
	vOneOfType12001 := new(DomainManagerRemoteClusterSpec)
	if err := json.Unmarshal(b, vOneOfType12001); err == nil {
		if "prism.v4.management.DomainManagerRemoteClusterSpec" == *vOneOfType12001.ObjectType_ {
			if nil == p.oneOfType12001 {
				p.oneOfType12001 = new(DomainManagerRemoteClusterSpec)
			}
			*p.oneOfType12001 = *vOneOfType12001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType12001.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType12001.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfClusterRegistrationSpecRemoteCluster"))
}

func (p *OneOfClusterRegistrationSpecRemoteCluster) MarshalJSON() ([]byte, error) {
	if p.oneOfType12002 != nil && *p.oneOfType12002.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType12002)
	}
	if p.oneOfType12003 != nil && *p.oneOfType12003.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType12003)
	}
	if p.oneOfType12001 != nil && *p.oneOfType12001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType12001)
	}
	return nil, errors.New("No value to marshal for OneOfClusterRegistrationSpecRemoteCluster")
}

type OneOfCreateRestoreSourceApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
	oneOfType0    *RestoreSource         `json:"-"`
}

func NewOneOfCreateRestoreSourceApiResponseData() *OneOfCreateRestoreSourceApiResponseData {
	p := new(OneOfCreateRestoreSourceApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateRestoreSourceApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateRestoreSourceApiResponseData is nil"))
	}
	switch v.(type) {
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case RestoreSource:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(RestoreSource)
		}
		*p.oneOfType0 = v.(RestoreSource)
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

func (p *OneOfCreateRestoreSourceApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfCreateRestoreSourceApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "prism.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	vOneOfType0 := new(RestoreSource)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.management.RestoreSource" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(RestoreSource)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateRestoreSourceApiResponseData"))
}

func (p *OneOfCreateRestoreSourceApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfCreateRestoreSourceApiResponseData")
}

type OneOfUpdateBackupTargetApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfUpdateBackupTargetApiResponseData() *OneOfUpdateBackupTargetApiResponseData {
	p := new(OneOfUpdateBackupTargetApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateBackupTargetApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateBackupTargetApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import1.TaskReference)
		}
		*p.oneOfType0 = v.(import1.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
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

func (p *OneOfUpdateBackupTargetApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfUpdateBackupTargetApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import1.TaskReference)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "prism.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateBackupTargetApiResponseData"))
}

func (p *OneOfUpdateBackupTargetApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateBackupTargetApiResponseData")
}

type OneOfGetBackupTargetApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
	oneOfType0    *BackupTarget          `json:"-"`
}

func NewOneOfGetBackupTargetApiResponseData() *OneOfGetBackupTargetApiResponseData {
	p := new(OneOfGetBackupTargetApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetBackupTargetApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetBackupTargetApiResponseData is nil"))
	}
	switch v.(type) {
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case BackupTarget:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(BackupTarget)
		}
		*p.oneOfType0 = v.(BackupTarget)
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

func (p *OneOfGetBackupTargetApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfGetBackupTargetApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "prism.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	vOneOfType0 := new(BackupTarget)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.management.BackupTarget" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(BackupTarget)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetBackupTargetApiResponseData"))
}

func (p *OneOfGetBackupTargetApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfGetBackupTargetApiResponseData")
}

type OneOfGetRestoreSourceApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
	oneOfType0    *RestoreSource         `json:"-"`
}

func NewOneOfGetRestoreSourceApiResponseData() *OneOfGetRestoreSourceApiResponseData {
	p := new(OneOfGetRestoreSourceApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetRestoreSourceApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetRestoreSourceApiResponseData is nil"))
	}
	switch v.(type) {
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case RestoreSource:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(RestoreSource)
		}
		*p.oneOfType0 = v.(RestoreSource)
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

func (p *OneOfGetRestoreSourceApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfGetRestoreSourceApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "prism.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	vOneOfType0 := new(RestoreSource)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.management.RestoreSource" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(RestoreSource)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetRestoreSourceApiResponseData"))
}

func (p *OneOfGetRestoreSourceApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfGetRestoreSourceApiResponseData")
}

type OneOfCreateBackupTargetApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfCreateBackupTargetApiResponseData() *OneOfCreateBackupTargetApiResponseData {
	p := new(OneOfCreateBackupTargetApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateBackupTargetApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateBackupTargetApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import1.TaskReference)
		}
		*p.oneOfType0 = v.(import1.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
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

func (p *OneOfCreateBackupTargetApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfCreateBackupTargetApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import1.TaskReference)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "prism.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateBackupTargetApiResponseData"))
}

func (p *OneOfCreateBackupTargetApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfCreateBackupTargetApiResponseData")
}

type OneOfListRestorableDomainManagersApiResponseData struct {
	Discriminator *string                   `json:"-"`
	ObjectType_   *string                   `json:"-"`
	oneOfType2001 []RestorableDomainManager `json:"-"`
	oneOfType400  *import2.ErrorResponse    `json:"-"`
}

func NewOneOfListRestorableDomainManagersApiResponseData() *OneOfListRestorableDomainManagersApiResponseData {
	p := new(OneOfListRestorableDomainManagersApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListRestorableDomainManagersApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListRestorableDomainManagersApiResponseData is nil"))
	}
	switch v.(type) {
	case []RestorableDomainManager:
		p.oneOfType2001 = v.([]RestorableDomainManager)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<prism.v4.management.RestorableDomainManager>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<prism.v4.management.RestorableDomainManager>"
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
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

func (p *OneOfListRestorableDomainManagersApiResponseData) GetValue() interface{} {
	if "List<prism.v4.management.RestorableDomainManager>" == *p.Discriminator {
		return p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfListRestorableDomainManagersApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new([]RestorableDomainManager)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if len(*vOneOfType2001) == 0 || "prism.v4.management.RestorableDomainManager" == *((*vOneOfType2001)[0].ObjectType_) {
			p.oneOfType2001 = *vOneOfType2001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<prism.v4.management.RestorableDomainManager>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<prism.v4.management.RestorableDomainManager>"
			return nil
		}
	}
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "prism.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListRestorableDomainManagersApiResponseData"))
}

func (p *OneOfListRestorableDomainManagersApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<prism.v4.management.RestorableDomainManager>" == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfListRestorableDomainManagersApiResponseData")
}

type OneOfRestoreSourceLocation struct {
	Discriminator *string              `json:"-"`
	ObjectType_   *string              `json:"-"`
	oneOfType0    *ClusterLocation     `json:"-"`
	oneOfType1    *ObjectStoreLocation `json:"-"`
}

func NewOneOfRestoreSourceLocation() *OneOfRestoreSourceLocation {
	p := new(OneOfRestoreSourceLocation)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfRestoreSourceLocation) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfRestoreSourceLocation is nil"))
	}
	switch v.(type) {
	case ClusterLocation:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(ClusterLocation)
		}
		*p.oneOfType0 = v.(ClusterLocation)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case ObjectStoreLocation:
		if nil == p.oneOfType1 {
			p.oneOfType1 = new(ObjectStoreLocation)
		}
		*p.oneOfType1 = v.(ObjectStoreLocation)
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

func (p *OneOfRestoreSourceLocation) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1
	}
	return nil
}

func (p *OneOfRestoreSourceLocation) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(ClusterLocation)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.management.ClusterLocation" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(ClusterLocation)
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
	vOneOfType1 := new(ObjectStoreLocation)
	if err := json.Unmarshal(b, vOneOfType1); err == nil {
		if "prism.v4.management.ObjectStoreLocation" == *vOneOfType1.ObjectType_ {
			if nil == p.oneOfType1 {
				p.oneOfType1 = new(ObjectStoreLocation)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfRestoreSourceLocation"))
}

func (p *OneOfRestoreSourceLocation) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	return nil, errors.New("No value to marshal for OneOfRestoreSourceLocation")
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
