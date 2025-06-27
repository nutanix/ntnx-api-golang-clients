/*
 * Generated file models/networking/v4/aws/config/config_model.go.
 *
 * Product version: 4.1.1
 *
 * Part of the Nutanix Cluster Management APIs
 *
 * (c) 2025 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module networking.v4.aws.config of Nutanix Cluster Management APIs
*/
package config

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	import1 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/common/v1/response"
)

/*
NC2A Subnet in the given VPC.
*/
type AwsSubnet struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Annotation string for cloud resources.
	*/
	Annotation *string `json:"annotation,omitempty"`
	/*
	  Availability Zone in which resource is situated.
	*/
	AvailabilityZone *string `json:"availabilityZone"`
	/*
	  Cloud subnet mask.
	*/
	Cidr *string `json:"cidr"`

	CloudType *CloudType `json:"cloudType"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Cloud subnet gateway IP.
	*/
	GatewayIp *string `json:"gatewayIp"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  Cloud subnet Id.
	*/
	SubnetId *string `json:"subnetId"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  AWS VPC ID where cluster is deployed.
	*/
	VpcId *string `json:"vpcId"`
}

func (p *AwsSubnet) MarshalJSON() ([]byte, error) {
	type AwsSubnetProxy AwsSubnet

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*AwsSubnetProxy
		AvailabilityZone *string    `json:"availabilityZone,omitempty"`
		Cidr             *string    `json:"cidr,omitempty"`
		CloudType        *CloudType `json:"cloudType,omitempty"`
		GatewayIp        *string    `json:"gatewayIp,omitempty"`
		SubnetId         *string    `json:"subnetId,omitempty"`
		VpcId            *string    `json:"vpcId,omitempty"`
	}{
		AwsSubnetProxy:   (*AwsSubnetProxy)(p),
		AvailabilityZone: p.AvailabilityZone,
		Cidr:             p.Cidr,
		CloudType:        p.CloudType,
		GatewayIp:        p.GatewayIp,
		SubnetId:         p.SubnetId,
		VpcId:            p.VpcId,
	}

	known, err := json.Marshal(baseStruct)
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *AwsSubnet) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias AwsSubnet
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = AwsSubnet(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "annotation")
	delete(allFields, "availabilityZone")
	delete(allFields, "cidr")
	delete(allFields, "cloudType")
	delete(allFields, "extId")
	delete(allFields, "gatewayIp")
	delete(allFields, "links")
	delete(allFields, "subnetId")
	delete(allFields, "tenantId")
	delete(allFields, "vpcId")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewAwsSubnet() *AwsSubnet {
	p := new(AwsSubnet)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.aws.config.AwsSubnet"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Enum defining different cloud platforms.
*/
type CloudType int

const (
	CLOUDTYPE_UNKNOWN  CloudType = 0
	CLOUDTYPE_REDACTED CloudType = 1
	CLOUDTYPE_AWS      CloudType = 2
	CLOUDTYPE_AZURE    CloudType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *CloudType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"AWS",
		"AZURE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e CloudType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"AWS",
		"AZURE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *CloudType) index(name string) CloudType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"AWS",
		"AZURE",
	}
	for idx := range names {
		if names[idx] == name {
			return CloudType(idx)
		}
	}
	return CLOUDTYPE_UNKNOWN
}

func (e *CloudType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for CloudType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *CloudType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e CloudType) Ref() *CloudType {
	return &e
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
