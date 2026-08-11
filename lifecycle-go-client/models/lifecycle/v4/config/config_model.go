/*
 * Generated file models/lifecycle/v4/config/config_model.go.
 *
 * Product version: 4.3.1
 *
 * Part of the Nutanix Lifecycle Management APIs
 *
 * (c) 2026 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module lifecycle.v4.config of Nutanix Lifecycle Management APIs
*/
package config

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	import1 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/common/v1/config"
	import2 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/common/v1/response"
	import4 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/lifecycle/v4/error"
	import3 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/prism/v4/config"
	"time"
)

/*
Active-Backup bonding configuration where one NIC is active and others provide failover redundancy
*/
type ActiveBackupBond struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Nics *Nics `json:"nics"`
}

func (p *ActiveBackupBond) MarshalJSON() ([]byte, error) {
	type ActiveBackupBondProxy ActiveBackupBond

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*ActiveBackupBondProxy
		Nics *Nics `json:"nics,omitempty"`
	}{
		ActiveBackupBondProxy: (*ActiveBackupBondProxy)(p),
		Nics:                  p.Nics,
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

func (p *ActiveBackupBond) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ActiveBackupBond
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewActiveBackupBond()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Nics != nil {
		p.Nics = known.Nics
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "nics")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewActiveBackupBond() *ActiveBackupBond {
	p := new(ActiveBackupBond)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.ActiveBackupBond"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
AOS installation details for the node
*/
type AosInstallationDetails struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	AosImageDetailsItemDiscriminator_ *string `json:"$aosImageDetailsItemDiscriminator,omitempty"`
	/*
	  Details of the AOS image
	*/
	AosImageDetails *OneOfAosInstallationDetailsAosImageDetails `json:"aosImageDetails"`
	/*
	  Memory allocation for the CVM in GB. This value will be overridden by FC if it does not fall within the range expected by Foundation Central for the given node
	*/
	CvmMemoryGB *int64 `json:"cvmMemoryGB,omitempty"`
	/*
	  List of DNS servers for the CVM
	*/
	NameServers []import1.IPAddress `json:"nameServers,omitempty"`

	NetworkDetails *AosNetworking `json:"networkDetails"`
	/*
	  List of NTP servers for the CVM
	*/
	NtpServers []import1.IPAddressOrFQDN `json:"ntpServers,omitempty"`
}

func (p *AosInstallationDetails) MarshalJSON() ([]byte, error) {
	type AosInstallationDetailsProxy AosInstallationDetails

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*AosInstallationDetailsProxy
		AosImageDetails *OneOfAosInstallationDetailsAosImageDetails `json:"aosImageDetails,omitempty"`
		NetworkDetails  *AosNetworking                              `json:"networkDetails,omitempty"`
	}{
		AosInstallationDetailsProxy: (*AosInstallationDetailsProxy)(p),
		AosImageDetails:             p.AosImageDetails,
		NetworkDetails:              p.NetworkDetails,
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

func (p *AosInstallationDetails) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias AosInstallationDetails
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewAosInstallationDetails()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.AosImageDetailsItemDiscriminator_ != nil {
		p.AosImageDetailsItemDiscriminator_ = known.AosImageDetailsItemDiscriminator_
	}
	if known.AosImageDetails != nil {
		p.AosImageDetails = known.AosImageDetails
	}
	if known.CvmMemoryGB != nil {
		p.CvmMemoryGB = known.CvmMemoryGB
	}
	if known.NameServers != nil {
		p.NameServers = known.NameServers
	}
	if known.NetworkDetails != nil {
		p.NetworkDetails = known.NetworkDetails
	}
	if known.NtpServers != nil {
		p.NtpServers = known.NtpServers
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$aosImageDetailsItemDiscriminator")
	delete(allFields, "aosImageDetails")
	delete(allFields, "cvmMemoryGB")
	delete(allFields, "nameServers")
	delete(allFields, "networkDetails")
	delete(allFields, "ntpServers")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewAosInstallationDetails() *AosInstallationDetails {
	p := new(AosInstallationDetails)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.AosInstallationDetails"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *AosInstallationDetails) GetAosImageDetails() interface{} {
	if nil == p.AosImageDetails {
		return nil
	}
	return p.AosImageDetails.GetValue()
}

func (p *AosInstallationDetails) SetAosImageDetails(v interface{}) error {
	if nil == p.AosImageDetails {
		p.AosImageDetails = NewOneOfAosInstallationDetailsAosImageDetails()
	}
	e := p.AosImageDetails.SetValue(v)
	if nil == e {
		if nil == p.AosImageDetailsItemDiscriminator_ {
			p.AosImageDetailsItemDiscriminator_ = new(string)
		}
		*p.AosImageDetailsItemDiscriminator_ = *p.AosImageDetails.Discriminator
	}
	return e
}

/*
Networking configuration for AOS including management interface
*/
type AosNetworking struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Management *NetworkDetails `json:"management"`
}

func (p *AosNetworking) MarshalJSON() ([]byte, error) {
	type AosNetworkingProxy AosNetworking

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*AosNetworkingProxy
		Management *NetworkDetails `json:"management,omitempty"`
	}{
		AosNetworkingProxy: (*AosNetworkingProxy)(p),
		Management:         p.Management,
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

func (p *AosNetworking) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias AosNetworking
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewAosNetworking()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Management != nil {
		p.Management = known.Management
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "management")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewAosNetworking() *AosNetworking {
	p := new(AosNetworking)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.AosNetworking"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Authentication details of hardware provider pertaining to API key and API secret
*/
type ApiKeyAuth struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  API key ID for authentication
	*/
	ApiKeyId *string `json:"apiKeyId"`
	/*
	  API key secret for authentication
	*/
	ApiKeySecret *string `json:"apiKeySecret,omitempty"`
}

func (p *ApiKeyAuth) MarshalJSON() ([]byte, error) {
	type ApiKeyAuthProxy ApiKeyAuth

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*ApiKeyAuthProxy
		ApiKeyId *string `json:"apiKeyId,omitempty"`
	}{
		ApiKeyAuthProxy: (*ApiKeyAuthProxy)(p),
		ApiKeyId:        p.ApiKeyId,
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

func (p *ApiKeyAuth) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ApiKeyAuth
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewApiKeyAuth()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ApiKeyId != nil {
		p.ApiKeyId = known.ApiKeyId
	}
	if known.ApiKeySecret != nil {
		p.ApiKeySecret = known.ApiKeySecret
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "apiKeyId")
	delete(allFields, "apiKeySecret")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewApiKeyAuth() *ApiKeyAuth {
	p := new(ApiKeyAuth)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.ApiKeyAuth"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Type of authentication supported by the hardware provider
*/
type AuthType int

const (
	AUTHTYPE_UNKNOWN        AuthType = 0
	AUTHTYPE_REDACTED       AuthType = 1
	AUTHTYPE_BASIC          AuthType = 2
	AUTHTYPE_SIGNED_API_KEY AuthType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *AuthType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"BASIC",
		"SIGNED_API_KEY",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e AuthType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"BASIC",
		"SIGNED_API_KEY",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *AuthType) index(name string) AuthType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"BASIC",
		"SIGNED_API_KEY",
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

/*
Basic authentication details using username and password
*/
type BasicAuth struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Password for authentication
	*/
	Password *string `json:"password,omitempty"`
	/*
	  Username for authentication
	*/
	Username *string `json:"username"`
}

func (p *BasicAuth) MarshalJSON() ([]byte, error) {
	type BasicAuthProxy BasicAuth

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*BasicAuthProxy
		Username *string `json:"username,omitempty"`
	}{
		BasicAuthProxy: (*BasicAuthProxy)(p),
		Username:       p.Username,
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

func (p *BasicAuth) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias BasicAuth
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewBasicAuth()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Password != nil {
		p.Password = known.Password
	}
	if known.Username != nil {
		p.Username = known.Username
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "password")
	delete(allFields, "username")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewBasicAuth() *BasicAuth {
	p := new(BasicAuth)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.BasicAuth"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
BMC IP configuration containing IPv4 and IPv6 addressing options
*/
type BmcIPConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	Ipv4ItemDiscriminator_ *string `json:"$ipv4ItemDiscriminator,omitempty"`
	/*
	  IPv4 configuration for the BMC
	*/
	Ipv4 *OneOfBmcIPConfigIpv4 `json:"ipv4,omitempty"`
	/*

	 */
	Ipv6ItemDiscriminator_ *string `json:"$ipv6ItemDiscriminator,omitempty"`
	/*
	  IPv6 configuration for the BMC
	*/
	Ipv6 *OneOfBmcIPConfigIpv6 `json:"ipv6,omitempty"`
}

func (p *BmcIPConfig) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias BmcIPConfig

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
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

func (p *BmcIPConfig) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias BmcIPConfig
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewBmcIPConfig()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Ipv4ItemDiscriminator_ != nil {
		p.Ipv4ItemDiscriminator_ = known.Ipv4ItemDiscriminator_
	}
	if known.Ipv4 != nil {
		p.Ipv4 = known.Ipv4
	}
	if known.Ipv6ItemDiscriminator_ != nil {
		p.Ipv6ItemDiscriminator_ = known.Ipv6ItemDiscriminator_
	}
	if known.Ipv6 != nil {
		p.Ipv6 = known.Ipv6
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$ipv4ItemDiscriminator")
	delete(allFields, "ipv4")
	delete(allFields, "$ipv6ItemDiscriminator")
	delete(allFields, "ipv6")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewBmcIPConfig() *BmcIPConfig {
	p := new(BmcIPConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.BmcIPConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *BmcIPConfig) GetIpv6() interface{} {
	if nil == p.Ipv6 {
		return nil
	}
	return p.Ipv6.GetValue()
}

func (p *BmcIPConfig) SetIpv6(v interface{}) error {
	if nil == p.Ipv6 {
		p.Ipv6 = NewOneOfBmcIPConfigIpv6()
	}
	e := p.Ipv6.SetValue(v)
	if nil == e {
		if nil == p.Ipv6ItemDiscriminator_ {
			p.Ipv6ItemDiscriminator_ = new(string)
		}
		*p.Ipv6ItemDiscriminator_ = *p.Ipv6.Discriminator
	}
	return e
}

/*
Network bonding configuration. Choose either Active-Backup bonding for simple failover or LACP bonding for dynamic link aggregation
*/
type BondSettings struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	BondConfigItemDiscriminator_ *string `json:"$bondConfigItemDiscriminator,omitempty"`
	/*
	  Bond configuration type
	*/
	BondConfig *OneOfBondSettingsBondConfig `json:"bondConfig"`
}

func (p *BondSettings) MarshalJSON() ([]byte, error) {
	type BondSettingsProxy BondSettings

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*BondSettingsProxy
		BondConfig *OneOfBondSettingsBondConfig `json:"bondConfig,omitempty"`
	}{
		BondSettingsProxy: (*BondSettingsProxy)(p),
		BondConfig:        p.BondConfig,
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

func (p *BondSettings) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias BondSettings
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewBondSettings()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.BondConfigItemDiscriminator_ != nil {
		p.BondConfigItemDiscriminator_ = known.BondConfigItemDiscriminator_
	}
	if known.BondConfig != nil {
		p.BondConfig = known.BondConfig
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$bondConfigItemDiscriminator")
	delete(allFields, "bondConfig")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewBondSettings() *BondSettings {
	p := new(BondSettings)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.BondSettings"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *BondSettings) GetBondConfig() interface{} {
	if nil == p.BondConfig {
		return nil
	}
	return p.BondConfig.GetValue()
}

func (p *BondSettings) SetBondConfig(v interface{}) error {
	if nil == p.BondConfig {
		p.BondConfig = NewOneOfBondSettingsBondConfig()
	}
	e := p.BondConfig.SetValue(v)
	if nil == e {
		if nil == p.BondConfigItemDiscriminator_ {
			p.BondConfigItemDiscriminator_ = new(string)
		}
		*p.BondConfigItemDiscriminator_ = *p.BondConfig.Discriminator
	}
	return e
}

/*
Checksum type
*/
type ChecksumType int

const (
	CHECKSUMTYPE_UNKNOWN  ChecksumType = 0
	CHECKSUMTYPE_REDACTED ChecksumType = 1
	CHECKSUMTYPE_MD5      ChecksumType = 2
	CHECKSUMTYPE_SHA256   ChecksumType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ChecksumType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"MD5",
		"SHA256",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ChecksumType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"MD5",
		"SHA256",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ChecksumType) index(name string) ChecksumType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"MD5",
		"SHA256",
	}
	for idx := range names {
		if names[idx] == name {
			return ChecksumType(idx)
		}
	}
	return CHECKSUMTYPE_UNKNOWN
}

func (e *ChecksumType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ChecksumType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ChecksumType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ChecksumType) Ref() *ChecksumType {
	return &e
}

/*
A secure token used for node authentication and registration with Foundation Central during onboarding
*/
type ClaimToken struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Time when the claim token was created
	*/
	CreatedTime *time.Time `json:"createdTime,omitempty"`
	/*
	  Number of times the claim token has been used for registering nodes
	*/
	CurrentUsageCount *int `json:"currentUsageCount,omitempty"`
	/*
	  Expiry time of the claim token
	*/
	ExpiryTime *time.Time `json:"expiryTime"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Maximum number of times the claim token can be used for registering nodes
	*/
	MaxUsageCount *int `json:"maxUsageCount"`
	/*
	  Name of the claim token
	*/
	Name *string `json:"name"`
	/*
	  External ID of the owner
	*/
	OwnerExtId *string `json:"ownerExtId,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *ClaimToken) MarshalJSON() ([]byte, error) {
	type ClaimTokenProxy ClaimToken

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*ClaimTokenProxy
		ExpiryTime    *time.Time `json:"expiryTime,omitempty"`
		MaxUsageCount *int       `json:"maxUsageCount,omitempty"`
		Name          *string    `json:"name,omitempty"`
	}{
		ClaimTokenProxy: (*ClaimTokenProxy)(p),
		ExpiryTime:      p.ExpiryTime,
		MaxUsageCount:   p.MaxUsageCount,
		Name:            p.Name,
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

func (p *ClaimToken) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ClaimToken
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewClaimToken()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.CreatedTime != nil {
		p.CreatedTime = known.CreatedTime
	}
	if known.CurrentUsageCount != nil {
		p.CurrentUsageCount = known.CurrentUsageCount
	}
	if known.ExpiryTime != nil {
		p.ExpiryTime = known.ExpiryTime
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.MaxUsageCount != nil {
		p.MaxUsageCount = known.MaxUsageCount
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.OwnerExtId != nil {
		p.OwnerExtId = known.OwnerExtId
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "createdTime")
	delete(allFields, "currentUsageCount")
	delete(allFields, "expiryTime")
	delete(allFields, "extId")
	delete(allFields, "links")
	delete(allFields, "maxUsageCount")
	delete(allFields, "name")
	delete(allFields, "ownerExtId")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewClaimToken() *ClaimToken {
	p := new(ClaimToken)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.ClaimToken"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
A node registered with Foundation Central via a claim token, with essential properties captured at registration
*/
type ClaimTokenNode struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	CpuInfo *CpuInfo `json:"cpuInfo,omitempty"`
	/*
	  Time when the node was discovered in Foundation Central
	*/
	CreatedTime *time.Time `json:"createdTime,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`

	HostType *HostType `json:"hostType,omitempty"`
	/*
	  Version of the host installed on the node
	*/
	HostVersion *string `json:"hostVersion,omitempty"`
	/*
	  Unique identifier array for the node (currently limited to 1 item with NODE_SERIAL type)
	*/
	Identifiers []NodeIdentifier `json:"identifiers"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Systems that manage the node
	*/
	ManagedBy []ClaimTokenNodeManager `json:"managedBy,omitempty"`
	/*
	  Manufacturer of the node hardware
	*/
	Manufacturer *string `json:"manufacturer"`
	/*
	  Total memory available in the node (in GB)
	*/
	MemoryGB *int `json:"memoryGB,omitempty"`
	/*
	  Model of the node
	*/
	Model *string `json:"model,omitempty"`

	NetworkDetails *NodeNetworkDetails `json:"networkDetails,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *ClaimTokenNode) MarshalJSON() ([]byte, error) {
	type ClaimTokenNodeProxy ClaimTokenNode

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*ClaimTokenNodeProxy
		Identifiers  []NodeIdentifier `json:"identifiers,omitempty"`
		Manufacturer *string          `json:"manufacturer,omitempty"`
	}{
		ClaimTokenNodeProxy: (*ClaimTokenNodeProxy)(p),
		Identifiers:         p.Identifiers,
		Manufacturer:        p.Manufacturer,
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

func (p *ClaimTokenNode) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ClaimTokenNode
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewClaimTokenNode()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.CpuInfo != nil {
		p.CpuInfo = known.CpuInfo
	}
	if known.CreatedTime != nil {
		p.CreatedTime = known.CreatedTime
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.HostType != nil {
		p.HostType = known.HostType
	}
	if known.HostVersion != nil {
		p.HostVersion = known.HostVersion
	}
	if known.Identifiers != nil {
		p.Identifiers = known.Identifiers
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.ManagedBy != nil {
		p.ManagedBy = known.ManagedBy
	}
	if known.Manufacturer != nil {
		p.Manufacturer = known.Manufacturer
	}
	if known.MemoryGB != nil {
		p.MemoryGB = known.MemoryGB
	}
	if known.Model != nil {
		p.Model = known.Model
	}
	if known.NetworkDetails != nil {
		p.NetworkDetails = known.NetworkDetails
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "cpuInfo")
	delete(allFields, "createdTime")
	delete(allFields, "extId")
	delete(allFields, "hostType")
	delete(allFields, "hostVersion")
	delete(allFields, "identifiers")
	delete(allFields, "links")
	delete(allFields, "managedBy")
	delete(allFields, "manufacturer")
	delete(allFields, "memoryGB")
	delete(allFields, "model")
	delete(allFields, "networkDetails")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewClaimTokenNode() *ClaimTokenNode {
	p := new(ClaimTokenNode)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.ClaimTokenNode"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Systems that manage the node
*/
type ClaimTokenNodeManager int

const (
	CLAIMTOKENNODEMANAGER_UNKNOWN            ClaimTokenNodeManager = 0
	CLAIMTOKENNODEMANAGER_REDACTED           ClaimTokenNodeManager = 1
	CLAIMTOKENNODEMANAGER_FOUNDATION_CENTRAL ClaimTokenNodeManager = 2
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ClaimTokenNodeManager) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"FOUNDATION_CENTRAL",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ClaimTokenNodeManager) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"FOUNDATION_CENTRAL",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ClaimTokenNodeManager) index(name string) ClaimTokenNodeManager {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"FOUNDATION_CENTRAL",
	}
	for idx := range names {
		if names[idx] == name {
			return ClaimTokenNodeManager(idx)
		}
	}
	return CLAIMTOKENNODEMANAGER_UNKNOWN
}

func (e *ClaimTokenNodeManager) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ClaimTokenNodeManager:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ClaimTokenNodeManager) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ClaimTokenNodeManager) Ref() *ClaimTokenNodeManager {
	return &e
}

/*
Claim token secret details
*/
type ClaimTokenSecret struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Claim token secret data
	*/
	Secret *string `json:"secret"`
}

func (p *ClaimTokenSecret) MarshalJSON() ([]byte, error) {
	type ClaimTokenSecretProxy ClaimTokenSecret

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*ClaimTokenSecretProxy
		Secret *string `json:"secret,omitempty"`
	}{
		ClaimTokenSecretProxy: (*ClaimTokenSecretProxy)(p),
		Secret:                p.Secret,
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

func (p *ClaimTokenSecret) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ClaimTokenSecret
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewClaimTokenSecret()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Secret != nil {
		p.Secret = known.Secret
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "secret")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewClaimTokenSecret() *ClaimTokenSecret {
	p := new(ClaimTokenSecret)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.ClaimTokenSecret"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Network details of the component(host or CVM)
*/
type ComponentNetworkDetails struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	ManagementNetwork *NetworkDetails `json:"managementNetwork,omitempty"`
}

func (p *ComponentNetworkDetails) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ComponentNetworkDetails

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
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

func (p *ComponentNetworkDetails) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ComponentNetworkDetails
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewComponentNetworkDetails()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ManagementNetwork != nil {
		p.ManagementNetwork = known.ManagementNetwork
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "managementNetwork")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewComponentNetworkDetails() *ComponentNetworkDetails {
	p := new(ComponentNetworkDetails)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.ComponentNetworkDetails"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /lifecycle/v4.3/config/nodes/$actions/configure Post operation
*/
type ConfigureNodeResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfConfigureNodeResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ConfigureNodeResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ConfigureNodeResponse

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
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

func (p *ConfigureNodeResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ConfigureNodeResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewConfigureNodeResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewConfigureNodeResponse() *ConfigureNodeResponse {
	p := new(ConfigureNodeResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.ConfigureNodeResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ConfigureNodeResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ConfigureNodeResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfConfigureNodeResponseData()
	}
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
Details required to apply node-specific configurations in the hardware provider
*/
type ConfigureNodeSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	ConfigurationItemDiscriminator_ *string `json:"$configurationItemDiscriminator,omitempty"`
	/*
	  Configuration for configuring a node
	*/
	Configuration *OneOfConfigureNodeSpecConfiguration `json:"configuration"`
}

func (p *ConfigureNodeSpec) MarshalJSON() ([]byte, error) {
	type ConfigureNodeSpecProxy ConfigureNodeSpec

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*ConfigureNodeSpecProxy
		Configuration *OneOfConfigureNodeSpecConfiguration `json:"configuration,omitempty"`
	}{
		ConfigureNodeSpecProxy: (*ConfigureNodeSpecProxy)(p),
		Configuration:          p.Configuration,
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

func (p *ConfigureNodeSpec) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ConfigureNodeSpec
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewConfigureNodeSpec()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ConfigurationItemDiscriminator_ != nil {
		p.ConfigurationItemDiscriminator_ = known.ConfigurationItemDiscriminator_
	}
	if known.Configuration != nil {
		p.Configuration = known.Configuration
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$configurationItemDiscriminator")
	delete(allFields, "configuration")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewConfigureNodeSpec() *ConfigureNodeSpec {
	p := new(ConfigureNodeSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.ConfigureNodeSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ConfigureNodeSpec) GetConfiguration() interface{} {
	if nil == p.Configuration {
		return nil
	}
	return p.Configuration.GetValue()
}

func (p *ConfigureNodeSpec) SetConfiguration(v interface{}) error {
	if nil == p.Configuration {
		p.Configuration = NewOneOfConfigureNodeSpecConfiguration()
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
Configuration details for setting up a server node in the hardware provider
*/
type ConfigureServerDetails struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	BmcIPAddressingConfigItemDiscriminator_ *string `json:"$bmcIPAddressingConfigItemDiscriminator,omitempty"`
	/*
	  BMC IP addressing configuration for the node
	*/
	BmcIPAddressingConfig *OneOfConfigureServerDetailsBmcIPAddressingConfig `json:"bmcIPAddressingConfig,omitempty"`
	/*
	  Identifier of the group (e.g., organization or tenant) to which the node belongs. This field is required for Cisco nodes
	*/
	GroupId *string `json:"groupId,omitempty"`
	/*
	  External ID of the node
	*/
	NodeExtId *string `json:"nodeExtId"`

	PhysicalNetworkAdapterConfig *PhysicalNetworkAdapterConfig `json:"physicalNetworkAdapterConfig,omitempty"`

	ServerSettingsConfig *ServerSettingsConfig `json:"serverSettingsConfig,omitempty"`
}

func (p *ConfigureServerDetails) MarshalJSON() ([]byte, error) {
	type ConfigureServerDetailsProxy ConfigureServerDetails

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*ConfigureServerDetailsProxy
		NodeExtId *string `json:"nodeExtId,omitempty"`
	}{
		ConfigureServerDetailsProxy: (*ConfigureServerDetailsProxy)(p),
		NodeExtId:                   p.NodeExtId,
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

func (p *ConfigureServerDetails) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ConfigureServerDetails
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewConfigureServerDetails()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.BmcIPAddressingConfigItemDiscriminator_ != nil {
		p.BmcIPAddressingConfigItemDiscriminator_ = known.BmcIPAddressingConfigItemDiscriminator_
	}
	if known.BmcIPAddressingConfig != nil {
		p.BmcIPAddressingConfig = known.BmcIPAddressingConfig
	}
	if known.GroupId != nil {
		p.GroupId = known.GroupId
	}
	if known.NodeExtId != nil {
		p.NodeExtId = known.NodeExtId
	}
	if known.PhysicalNetworkAdapterConfig != nil {
		p.PhysicalNetworkAdapterConfig = known.PhysicalNetworkAdapterConfig
	}
	if known.ServerSettingsConfig != nil {
		p.ServerSettingsConfig = known.ServerSettingsConfig
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$bmcIPAddressingConfigItemDiscriminator")
	delete(allFields, "bmcIPAddressingConfig")
	delete(allFields, "groupId")
	delete(allFields, "nodeExtId")
	delete(allFields, "physicalNetworkAdapterConfig")
	delete(allFields, "serverSettingsConfig")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewConfigureServerDetails() *ConfigureServerDetails {
	p := new(ConfigureServerDetails)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.ConfigureServerDetails"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ConfigureServerDetails) GetBmcIPAddressingConfig() interface{} {
	if nil == p.BmcIPAddressingConfig {
		return nil
	}
	return p.BmcIPAddressingConfig.GetValue()
}

func (p *ConfigureServerDetails) SetBmcIPAddressingConfig(v interface{}) error {
	if nil == p.BmcIPAddressingConfig {
		p.BmcIPAddressingConfig = NewOneOfConfigureServerDetailsBmcIPAddressingConfig()
	}
	e := p.BmcIPAddressingConfig.SetValue(v)
	if nil == e {
		if nil == p.BmcIPAddressingConfigItemDiscriminator_ {
			p.BmcIPAddressingConfigItemDiscriminator_ = new(string)
		}
		*p.BmcIPAddressingConfigItemDiscriminator_ = *p.BmcIPAddressingConfig.Discriminator
	}
	return e
}

/*
Details of a connection to a hardware provider endpoint, enabling node discovery and management
*/
type Connection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Access details including endpoint and authentication information
	*/
	AccessDetails []ConnectionDetail `json:"accessDetails"`
	/*
	  Creation time of the connection
	*/
	CreatedTime *time.Time `json:"createdTime,omitempty"`

	DeploymentType *ConnectionDeploymentType `json:"deploymentType,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Name of the connection
	*/
	Name *string `json:"name"`
	/*
	  Region for the connection
	*/
	Region *string `json:"region,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *Connection) MarshalJSON() ([]byte, error) {
	type ConnectionProxy Connection

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*ConnectionProxy
		AccessDetails []ConnectionDetail `json:"accessDetails,omitempty"`
		Name          *string            `json:"name,omitempty"`
	}{
		ConnectionProxy: (*ConnectionProxy)(p),
		AccessDetails:   p.AccessDetails,
		Name:            p.Name,
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

func (p *Connection) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Connection
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewConnection()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.AccessDetails != nil {
		p.AccessDetails = known.AccessDetails
	}
	if known.CreatedTime != nil {
		p.CreatedTime = known.CreatedTime
	}
	if known.DeploymentType != nil {
		p.DeploymentType = known.DeploymentType
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.Region != nil {
		p.Region = known.Region
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "accessDetails")
	delete(allFields, "createdTime")
	delete(allFields, "deploymentType")
	delete(allFields, "extId")
	delete(allFields, "links")
	delete(allFields, "name")
	delete(allFields, "region")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewConnection() *Connection {
	p := new(Connection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.Connection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Type of deployment for the hardware provider connection
*/
type ConnectionDeploymentType int

const (
	CONNECTIONDEPLOYMENTTYPE_UNKNOWN    ConnectionDeploymentType = 0
	CONNECTIONDEPLOYMENTTYPE_REDACTED   ConnectionDeploymentType = 1
	CONNECTIONDEPLOYMENTTYPE_SAAS       ConnectionDeploymentType = 2
	CONNECTIONDEPLOYMENTTYPE_ONPREM     ConnectionDeploymentType = 3
	CONNECTIONDEPLOYMENTTYPE_STANDALONE ConnectionDeploymentType = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ConnectionDeploymentType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SAAS",
		"ONPREM",
		"STANDALONE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ConnectionDeploymentType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SAAS",
		"ONPREM",
		"STANDALONE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ConnectionDeploymentType) index(name string) ConnectionDeploymentType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SAAS",
		"ONPREM",
		"STANDALONE",
	}
	for idx := range names {
		if names[idx] == name {
			return ConnectionDeploymentType(idx)
		}
	}
	return CONNECTIONDEPLOYMENTTYPE_UNKNOWN
}

func (e *ConnectionDeploymentType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ConnectionDeploymentType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ConnectionDeploymentType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ConnectionDeploymentType) Ref() *ConnectionDeploymentType {
	return &e
}

/*
Details of a connection including endpoint and authentication
*/
type ConnectionDetail struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	AuthItemDiscriminator_ *string `json:"$authItemDiscriminator,omitempty"`
	/*
	  Authentication configuration for the connection
	*/
	Auth *OneOfConnectionDetailAuth `json:"auth"`
	/*

	 */
	EndpointItemDiscriminator_ *string `json:"$endpointItemDiscriminator,omitempty"`
	/*
	  Endpoint configuration for the connection
	*/
	Endpoint *OneOfConnectionDetailEndpoint `json:"endpoint"`
}

func (p *ConnectionDetail) MarshalJSON() ([]byte, error) {
	type ConnectionDetailProxy ConnectionDetail

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*ConnectionDetailProxy
		Auth     *OneOfConnectionDetailAuth     `json:"auth,omitempty"`
		Endpoint *OneOfConnectionDetailEndpoint `json:"endpoint,omitempty"`
	}{
		ConnectionDetailProxy: (*ConnectionDetailProxy)(p),
		Auth:                  p.Auth,
		Endpoint:              p.Endpoint,
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

func (p *ConnectionDetail) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ConnectionDetail
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewConnectionDetail()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.AuthItemDiscriminator_ != nil {
		p.AuthItemDiscriminator_ = known.AuthItemDiscriminator_
	}
	if known.Auth != nil {
		p.Auth = known.Auth
	}
	if known.EndpointItemDiscriminator_ != nil {
		p.EndpointItemDiscriminator_ = known.EndpointItemDiscriminator_
	}
	if known.Endpoint != nil {
		p.Endpoint = known.Endpoint
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$authItemDiscriminator")
	delete(allFields, "auth")
	delete(allFields, "$endpointItemDiscriminator")
	delete(allFields, "endpoint")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewConnectionDetail() *ConnectionDetail {
	p := new(ConnectionDetail)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.ConnectionDetail"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ConnectionDetail) GetAuth() interface{} {
	if nil == p.Auth {
		return nil
	}
	return p.Auth.GetValue()
}

func (p *ConnectionDetail) SetAuth(v interface{}) error {
	if nil == p.Auth {
		p.Auth = NewOneOfConnectionDetailAuth()
	}
	e := p.Auth.SetValue(v)
	if nil == e {
		if nil == p.AuthItemDiscriminator_ {
			p.AuthItemDiscriminator_ = new(string)
		}
		*p.AuthItemDiscriminator_ = *p.Auth.Discriminator
	}
	return e
}

type ConnectionProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Access details including endpoint and authentication information
	*/
	AccessDetails []ConnectionDetail `json:"accessDetails"`
	/*
	  Creation time of the connection
	*/
	CreatedTime *time.Time `json:"createdTime,omitempty"`

	DeploymentType *ConnectionDeploymentType `json:"deploymentType,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Name of the connection
	*/
	Name *string `json:"name"`
	/*
	  Region for the connection
	*/
	Region *string `json:"region,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *ConnectionProjection) MarshalJSON() ([]byte, error) {
	type ConnectionProjectionProxy ConnectionProjection

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*ConnectionProjectionProxy
		AccessDetails []ConnectionDetail `json:"accessDetails,omitempty"`
		Name          *string            `json:"name,omitempty"`
	}{
		ConnectionProjectionProxy: (*ConnectionProjectionProxy)(p),
		AccessDetails:             p.AccessDetails,
		Name:                      p.Name,
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

func (p *ConnectionProjection) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ConnectionProjection
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewConnectionProjection()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.AccessDetails != nil {
		p.AccessDetails = known.AccessDetails
	}
	if known.CreatedTime != nil {
		p.CreatedTime = known.CreatedTime
	}
	if known.DeploymentType != nil {
		p.DeploymentType = known.DeploymentType
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.Region != nil {
		p.Region = known.Region
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "accessDetails")
	delete(allFields, "createdTime")
	delete(allFields, "deploymentType")
	delete(allFields, "extId")
	delete(allFields, "links")
	delete(allFields, "name")
	delete(allFields, "region")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewConnectionProjection() *ConnectionProjection {
	p := new(ConnectionProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.ConnectionProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
CPU hardware details
*/
type CpuInfo struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  CPU clock speed in GHz
	*/
	CapacityGHz *float32 `json:"capacityGHz,omitempty"`
	/*
	  Number of logical CPU cores
	*/
	LogicalCoreCount *int `json:"logicalCoreCount,omitempty"`
	/*
	  Manufacturer of the CPU (e.g., Intel, AMD)
	*/
	Manufacturer *string `json:"manufacturer,omitempty"`
	/*
	  Model name of the CPU
	*/
	Model *string `json:"model,omitempty"`
	/*
	  Number of CPU sockets in the node
	*/
	SocketCount *int `json:"socketCount,omitempty"`
}

func (p *CpuInfo) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias CpuInfo

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
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

func (p *CpuInfo) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias CpuInfo
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewCpuInfo()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.CapacityGHz != nil {
		p.CapacityGHz = known.CapacityGHz
	}
	if known.LogicalCoreCount != nil {
		p.LogicalCoreCount = known.LogicalCoreCount
	}
	if known.Manufacturer != nil {
		p.Manufacturer = known.Manufacturer
	}
	if known.Model != nil {
		p.Model = known.Model
	}
	if known.SocketCount != nil {
		p.SocketCount = known.SocketCount
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "capacityGHz")
	delete(allFields, "logicalCoreCount")
	delete(allFields, "manufacturer")
	delete(allFields, "model")
	delete(allFields, "socketCount")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewCpuInfo() *CpuInfo {
	p := new(CpuInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.CpuInfo"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /lifecycle/v4.3/config/claim-tokens Post operation
*/
type CreateClaimTokenApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateClaimTokenApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *CreateClaimTokenApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias CreateClaimTokenApiResponse

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
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

func (p *CreateClaimTokenApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias CreateClaimTokenApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewCreateClaimTokenApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewCreateClaimTokenApiResponse() *CreateClaimTokenApiResponse {
	p := new(CreateClaimTokenApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.CreateClaimTokenApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreateClaimTokenApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreateClaimTokenApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreateClaimTokenApiResponseData()
	}
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
REST response for all response codes in API path /lifecycle/v4.3/config/hardware-providers/{hardwareProviderExtId}/connections Post operation
*/
type CreateConnectionApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateConnectionApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *CreateConnectionApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias CreateConnectionApiResponse

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
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

func (p *CreateConnectionApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias CreateConnectionApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewCreateConnectionApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewCreateConnectionApiResponse() *CreateConnectionApiResponse {
	p := new(CreateConnectionApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.CreateConnectionApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreateConnectionApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreateConnectionApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreateConnectionApiResponseData()
	}
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
REST response for all response codes in API path /lifecycle/v4.3/config/images Post operation
*/
type CreateImageApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateImageApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *CreateImageApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias CreateImageApiResponse

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
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

func (p *CreateImageApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias CreateImageApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewCreateImageApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewCreateImageApiResponse() *CreateImageApiResponse {
	p := new(CreateImageApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.CreateImageApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreateImageApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreateImageApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreateImageApiResponseData()
	}
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
REST response for all response codes in API path /lifecycle/v4.3/config/nodes Post operation
*/
type CreateNodeApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateNodeApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *CreateNodeApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias CreateNodeApiResponse

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
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

func (p *CreateNodeApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias CreateNodeApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewCreateNodeApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewCreateNodeApiResponse() *CreateNodeApiResponse {
	p := new(CreateNodeApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.CreateNodeApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreateNodeApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreateNodeApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreateNodeApiResponseData()
	}
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
REST response for all response codes in API path /lifecycle/v4.3/config/patched-images Post operation
*/
type CreatePatchedImageApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreatePatchedImageApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *CreatePatchedImageApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias CreatePatchedImageApiResponse

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
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

func (p *CreatePatchedImageApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias CreatePatchedImageApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewCreatePatchedImageApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewCreatePatchedImageApiResponse() *CreatePatchedImageApiResponse {
	p := new(CreatePatchedImageApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.CreatePatchedImageApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreatePatchedImageApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreatePatchedImageApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreatePatchedImageApiResponseData()
	}
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
Configuration details for the CVM (Controller VM) layer. Includes network settings (IP, gateway, VLAN), DNS servers, and NTP servers
*/
type CvmReconfigurationDetails struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of DNS servers for the CVM
	*/
	NameServers []import1.IPAddress `json:"nameServers,omitempty"`

	NetworkDetails *AosNetworking `json:"networkDetails,omitempty"`
	/*
	  List of NTP servers for the CVM
	*/
	NtpServers []import1.IPAddressOrFQDN `json:"ntpServers,omitempty"`
}

func (p *CvmReconfigurationDetails) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias CvmReconfigurationDetails

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
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

func (p *CvmReconfigurationDetails) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias CvmReconfigurationDetails
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewCvmReconfigurationDetails()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.NameServers != nil {
		p.NameServers = known.NameServers
	}
	if known.NetworkDetails != nil {
		p.NetworkDetails = known.NetworkDetails
	}
	if known.NtpServers != nil {
		p.NtpServers = known.NtpServers
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "nameServers")
	delete(allFields, "networkDetails")
	delete(allFields, "ntpServers")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewCvmReconfigurationDetails() *CvmReconfigurationDetails {
	p := new(CvmReconfigurationDetails)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.CvmReconfigurationDetails"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /lifecycle/v4.3/config/claim-tokens/{extId} Delete operation
*/
type DeleteClaimTokenApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteClaimTokenApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *DeleteClaimTokenApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DeleteClaimTokenApiResponse

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
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

func (p *DeleteClaimTokenApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DeleteClaimTokenApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewDeleteClaimTokenApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewDeleteClaimTokenApiResponse() *DeleteClaimTokenApiResponse {
	p := new(DeleteClaimTokenApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.DeleteClaimTokenApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteClaimTokenApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteClaimTokenApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteClaimTokenApiResponseData()
	}
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
REST response for all response codes in API path /lifecycle/v4.3/config/hardware-providers/{hardwareProviderExtId}/connections/{extId} Delete operation
*/
type DeleteConnectionApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteConnectionApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *DeleteConnectionApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DeleteConnectionApiResponse

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
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

func (p *DeleteConnectionApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DeleteConnectionApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewDeleteConnectionApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewDeleteConnectionApiResponse() *DeleteConnectionApiResponse {
	p := new(DeleteConnectionApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.DeleteConnectionApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteConnectionApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteConnectionApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteConnectionApiResponseData()
	}
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
REST response for all response codes in API path /lifecycle/v4.3/config/images/{extId} Delete operation
*/
type DeleteImageApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteImageApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *DeleteImageApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DeleteImageApiResponse

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
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

func (p *DeleteImageApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DeleteImageApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewDeleteImageApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewDeleteImageApiResponse() *DeleteImageApiResponse {
	p := new(DeleteImageApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.DeleteImageApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteImageApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteImageApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteImageApiResponseData()
	}
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
REST response for all response codes in API path /lifecycle/v4.3/config/nodes/{extId} Delete operation
*/
type DeleteNodeApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteNodeApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *DeleteNodeApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DeleteNodeApiResponse

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
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

func (p *DeleteNodeApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DeleteNodeApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewDeleteNodeApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewDeleteNodeApiResponse() *DeleteNodeApiResponse {
	p := new(DeleteNodeApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.DeleteNodeApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteNodeApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteNodeApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteNodeApiResponseData()
	}
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
REST response for all response codes in API path /lifecycle/v4.3/config/patched-images/{extId} Delete operation
*/
type DeletePatchedImageApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeletePatchedImageApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *DeletePatchedImageApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DeletePatchedImageApiResponse

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
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

func (p *DeletePatchedImageApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DeletePatchedImageApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewDeletePatchedImageApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewDeletePatchedImageApiResponse() *DeletePatchedImageApiResponse {
	p := new(DeletePatchedImageApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.DeletePatchedImageApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeletePatchedImageApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeletePatchedImageApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeletePatchedImageApiResponseData()
	}
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
Specification for discovering nodes from a hardware provider connection
*/
type DiscoverNodesSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of group IDs to refresh resources for. Mandatory param if it's cisco connection
	*/
	GroupIds []string `json:"groupIds,omitempty"`
}

func (p *DiscoverNodesSpec) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DiscoverNodesSpec

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
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

func (p *DiscoverNodesSpec) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DiscoverNodesSpec
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewDiscoverNodesSpec()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.GroupIds != nil {
		p.GroupIds = known.GroupIds
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "groupIds")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewDiscoverNodesSpec() *DiscoverNodesSpec {
	p := new(DiscoverNodesSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.DiscoverNodesSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Information about a discovered node from a hardware provider
*/
type DiscoveredNode struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	BmcIp *import1.IPAddress `json:"bmcIp,omitempty"`

	CpuInfo *CpuInfo `json:"cpuInfo,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Unique identifier array for the discovered node (currently limited to 1 item with NODE_SERIAL type)
	*/
	Identifiers []NodeIdentifier `json:"identifiers,omitempty"`
	/*
	  Time when the discovered node was last refreshed
	*/
	LastRefreshedTime *time.Time `json:"lastRefreshedTime"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Systems that manage this node
	*/
	ManagedBy []ManagedBy `json:"managedBy,omitempty"`
	/*
	  Manufacturer of the discovered node
	*/
	Manufacturer *string `json:"manufacturer,omitempty"`
	/*
	  Memory capacity in GB
	*/
	MemoryGB *int `json:"memoryGB,omitempty"`
	/*
	  Model of the discovered node
	*/
	Model *string `json:"model,omitempty"`

	ProviderData *NodeProviderData `json:"providerData,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *DiscoveredNode) MarshalJSON() ([]byte, error) {
	type DiscoveredNodeProxy DiscoveredNode

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*DiscoveredNodeProxy
		LastRefreshedTime *time.Time `json:"lastRefreshedTime,omitempty"`
	}{
		DiscoveredNodeProxy: (*DiscoveredNodeProxy)(p),
		LastRefreshedTime:   p.LastRefreshedTime,
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

func (p *DiscoveredNode) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DiscoveredNode
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewDiscoveredNode()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.BmcIp != nil {
		p.BmcIp = known.BmcIp
	}
	if known.CpuInfo != nil {
		p.CpuInfo = known.CpuInfo
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.Identifiers != nil {
		p.Identifiers = known.Identifiers
	}
	if known.LastRefreshedTime != nil {
		p.LastRefreshedTime = known.LastRefreshedTime
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.ManagedBy != nil {
		p.ManagedBy = known.ManagedBy
	}
	if known.Manufacturer != nil {
		p.Manufacturer = known.Manufacturer
	}
	if known.MemoryGB != nil {
		p.MemoryGB = known.MemoryGB
	}
	if known.Model != nil {
		p.Model = known.Model
	}
	if known.ProviderData != nil {
		p.ProviderData = known.ProviderData
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "bmcIp")
	delete(allFields, "cpuInfo")
	delete(allFields, "extId")
	delete(allFields, "identifiers")
	delete(allFields, "lastRefreshedTime")
	delete(allFields, "links")
	delete(allFields, "managedBy")
	delete(allFields, "manufacturer")
	delete(allFields, "memoryGB")
	delete(allFields, "model")
	delete(allFields, "providerData")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewDiscoveredNode() *DiscoveredNode {
	p := new(DiscoveredNode)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.DiscoveredNode"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Mapping between unified attribute names and vendor-specific attribute names
*/
type DisplayNameMapping struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Vendor-specific attribute name for domain
	*/
	Domain *string `json:"domain,omitempty"`
	/*
	  Vendor-specific attribute name for groups
	*/
	Groups *string `json:"groups,omitempty"`
	/*
	  Vendor-specific attribute name for mode
	*/
	Mode *string `json:"mode,omitempty"`
	/*
	  Vendor-specific attribute name for tags
	*/
	Tags *string `json:"tags,omitempty"`
}

func (p *DisplayNameMapping) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DisplayNameMapping

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
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

func (p *DisplayNameMapping) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DisplayNameMapping
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewDisplayNameMapping()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Domain != nil {
		p.Domain = known.Domain
	}
	if known.Groups != nil {
		p.Groups = known.Groups
	}
	if known.Mode != nil {
		p.Mode = known.Mode
	}
	if known.Tags != nil {
		p.Tags = known.Tags
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "domain")
	delete(allFields, "groups")
	delete(allFields, "mode")
	delete(allFields, "tags")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewDisplayNameMapping() *DisplayNameMapping {
	p := new(DisplayNameMapping)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.DisplayNameMapping"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Extended provider-specific data for the node including recommended network groups and node position
*/
type ExtendedNodeProviderData struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	DisplayNameMapping *DisplayNameMapping `json:"displayNameMapping,omitempty"`
	/*
	  Domain of the discovered node
	*/
	Domain *string `json:"domain,omitempty"`
	/*
	  Groups associated with the discovered node
	*/
	Groups []Group `json:"groups,omitempty"`
	/*
	  Indicates if the node is available
	*/
	IsAvailable *bool `json:"isAvailable,omitempty"`
	/*
	  Indicates if the node is configured
	*/
	IsConfigured *bool `json:"isConfigured,omitempty"`
	/*
	  Mode of the discovered node
	*/
	Mode *string `json:"mode,omitempty"`
	/*
	  Name of the discovered node
	*/
	Name *string `json:"name,omitempty"`
	/*
	  Physical position of the node in the hardware provider
	*/
	NodePosition *string `json:"nodePosition,omitempty"`

	PhysicalNetworkAdapters *PhysicalNetworkAdapters `json:"physicalNetworkAdapters,omitempty"`
	/*
	  Tags associated with the discovered node
	*/
	Tags []import1.KVStringPair `json:"tags,omitempty"`
}

func (p *ExtendedNodeProviderData) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ExtendedNodeProviderData

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
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

func (p *ExtendedNodeProviderData) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ExtendedNodeProviderData
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewExtendedNodeProviderData()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DisplayNameMapping != nil {
		p.DisplayNameMapping = known.DisplayNameMapping
	}
	if known.Domain != nil {
		p.Domain = known.Domain
	}
	if known.Groups != nil {
		p.Groups = known.Groups
	}
	if known.IsAvailable != nil {
		p.IsAvailable = known.IsAvailable
	}
	if known.IsConfigured != nil {
		p.IsConfigured = known.IsConfigured
	}
	if known.Mode != nil {
		p.Mode = known.Mode
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.NodePosition != nil {
		p.NodePosition = known.NodePosition
	}
	if known.PhysicalNetworkAdapters != nil {
		p.PhysicalNetworkAdapters = known.PhysicalNetworkAdapters
	}
	if known.Tags != nil {
		p.Tags = known.Tags
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "displayNameMapping")
	delete(allFields, "domain")
	delete(allFields, "groups")
	delete(allFields, "isAvailable")
	delete(allFields, "isConfigured")
	delete(allFields, "mode")
	delete(allFields, "name")
	delete(allFields, "nodePosition")
	delete(allFields, "physicalNetworkAdapters")
	delete(allFields, "tags")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewExtendedNodeProviderData() *ExtendedNodeProviderData {
	p := new(ExtendedNodeProviderData)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.ExtendedNodeProviderData"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Status of the image file (for example FILE_AVAILABLE, UPLOADING, URL_AVAILABLE, NOT_AVAILABLE)
*/
type FileStatus int

const (
	FILESTATUS_UNKNOWN        FileStatus = 0
	FILESTATUS_REDACTED       FileStatus = 1
	FILESTATUS_FILE_AVAILABLE FileStatus = 2
	FILESTATUS_UPLOADING      FileStatus = 3
	FILESTATUS_UPLOAD_FAILED  FileStatus = 4
	FILESTATUS_URL_AVAILABLE  FileStatus = 5
	FILESTATUS_NOT_AVAILABLE  FileStatus = 6
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *FileStatus) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"FILE_AVAILABLE",
		"UPLOADING",
		"UPLOAD_FAILED",
		"URL_AVAILABLE",
		"NOT_AVAILABLE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e FileStatus) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"FILE_AVAILABLE",
		"UPLOADING",
		"UPLOAD_FAILED",
		"URL_AVAILABLE",
		"NOT_AVAILABLE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *FileStatus) index(name string) FileStatus {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"FILE_AVAILABLE",
		"UPLOADING",
		"UPLOAD_FAILED",
		"URL_AVAILABLE",
		"NOT_AVAILABLE",
	}
	for idx := range names {
		if names[idx] == name {
			return FileStatus(idx)
		}
	}
	return FILESTATUS_UNKNOWN
}

func (e *FileStatus) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for FileStatus:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *FileStatus) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e FileStatus) Ref() *FileStatus {
	return &e
}

/*
REST response for all response codes in API path /lifecycle/v4.3/config/claim-tokens/{extId} Get operation
*/
type GetClaimTokenApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetClaimTokenApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *GetClaimTokenApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetClaimTokenApiResponse

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
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

func (p *GetClaimTokenApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetClaimTokenApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewGetClaimTokenApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewGetClaimTokenApiResponse() *GetClaimTokenApiResponse {
	p := new(GetClaimTokenApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.GetClaimTokenApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetClaimTokenApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetClaimTokenApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetClaimTokenApiResponseData()
	}
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
REST response for all response codes in API path /lifecycle/v4.3/config/claim-tokens/{extId}/secret Get operation
*/
type GetClaimTokenSecretApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetClaimTokenSecretApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *GetClaimTokenSecretApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetClaimTokenSecretApiResponse

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
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

func (p *GetClaimTokenSecretApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetClaimTokenSecretApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewGetClaimTokenSecretApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewGetClaimTokenSecretApiResponse() *GetClaimTokenSecretApiResponse {
	p := new(GetClaimTokenSecretApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.GetClaimTokenSecretApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetClaimTokenSecretApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetClaimTokenSecretApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetClaimTokenSecretApiResponseData()
	}
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
REST response for all response codes in API path /lifecycle/v4.3/config/hardware-providers/{hardwareProviderExtId}/connections/{extId} Get operation
*/
type GetConnectionApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetConnectionApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *GetConnectionApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetConnectionApiResponse

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
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

func (p *GetConnectionApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetConnectionApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewGetConnectionApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewGetConnectionApiResponse() *GetConnectionApiResponse {
	p := new(GetConnectionApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.GetConnectionApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetConnectionApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetConnectionApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetConnectionApiResponseData()
	}
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
REST response for all response codes in API path /lifecycle/v4.3/config/hardware-providers/{hardwareProviderExtId}/connections/{connectionExtId}/nodes/{extId} Get operation
*/
type GetConnectionNodeApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetConnectionNodeApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *GetConnectionNodeApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetConnectionNodeApiResponse

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
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

func (p *GetConnectionNodeApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetConnectionNodeApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewGetConnectionNodeApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewGetConnectionNodeApiResponse() *GetConnectionNodeApiResponse {
	p := new(GetConnectionNodeApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.GetConnectionNodeApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetConnectionNodeApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetConnectionNodeApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetConnectionNodeApiResponseData()
	}
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
REST response for all response codes in API path /lifecycle/v4.3/config/hardware-providers/{extId} Get operation
*/
type GetHardwareProviderApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetHardwareProviderApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *GetHardwareProviderApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetHardwareProviderApiResponse

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
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

func (p *GetHardwareProviderApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetHardwareProviderApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewGetHardwareProviderApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewGetHardwareProviderApiResponse() *GetHardwareProviderApiResponse {
	p := new(GetHardwareProviderApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.GetHardwareProviderApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetHardwareProviderApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetHardwareProviderApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetHardwareProviderApiResponseData()
	}
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
REST response for all response codes in API path /lifecycle/v4.3/config/images/{extId} Get operation
*/
type GetImageApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetImageApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *GetImageApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetImageApiResponse

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
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

func (p *GetImageApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetImageApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewGetImageApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewGetImageApiResponse() *GetImageApiResponse {
	p := new(GetImageApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.GetImageApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetImageApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetImageApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetImageApiResponseData()
	}
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
REST response for all response codes in API path /lifecycle/v4.3/config/hardware-providers/{hardwareProviderExtId}/connections/{connectionExtId}/ip-pools/{extId} Get operation
*/
type GetIpPoolApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetIpPoolApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *GetIpPoolApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetIpPoolApiResponse

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
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

func (p *GetIpPoolApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetIpPoolApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewGetIpPoolApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewGetIpPoolApiResponse() *GetIpPoolApiResponse {
	p := new(GetIpPoolApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.GetIpPoolApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetIpPoolApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetIpPoolApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetIpPoolApiResponseData()
	}
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
REST response for all response codes in API path /lifecycle/v4.3/config/hardware-providers/{hardwareProviderExtId}/connections/{connectionExtId}/mac-pools/{extId} Get operation
*/
type GetMacPoolApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetMacPoolApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *GetMacPoolApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetMacPoolApiResponse

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
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

func (p *GetMacPoolApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetMacPoolApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewGetMacPoolApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewGetMacPoolApiResponse() *GetMacPoolApiResponse {
	p := new(GetMacPoolApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.GetMacPoolApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetMacPoolApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetMacPoolApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetMacPoolApiResponseData()
	}
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
REST response for all response codes in API path /lifecycle/v4.3/config/nodes/{extId} Get operation
*/
type GetNodeApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetNodeApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *GetNodeApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetNodeApiResponse

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
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

func (p *GetNodeApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetNodeApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewGetNodeApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewGetNodeApiResponse() *GetNodeApiResponse {
	p := new(GetNodeApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.GetNodeApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetNodeApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetNodeApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetNodeApiResponseData()
	}
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
REST response for all response codes in API path /lifecycle/v4.3/config/claim-tokens/{claimTokenExtId}/nodes/{extId} Get operation
*/
type GetNodeByClaimTokenApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetNodeByClaimTokenApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *GetNodeByClaimTokenApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetNodeByClaimTokenApiResponse

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
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

func (p *GetNodeByClaimTokenApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetNodeByClaimTokenApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewGetNodeByClaimTokenApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewGetNodeByClaimTokenApiResponse() *GetNodeByClaimTokenApiResponse {
	p := new(GetNodeByClaimTokenApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.GetNodeByClaimTokenApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetNodeByClaimTokenApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetNodeByClaimTokenApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetNodeByClaimTokenApiResponseData()
	}
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
REST response for all response codes in API path /lifecycle/v4.3/config/patched-images/{extId} Get operation
*/
type GetPatchedImageApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetPatchedImageApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *GetPatchedImageApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetPatchedImageApiResponse

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
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

func (p *GetPatchedImageApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetPatchedImageApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewGetPatchedImageApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewGetPatchedImageApiResponse() *GetPatchedImageApiResponse {
	p := new(GetPatchedImageApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.GetPatchedImageApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetPatchedImageApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetPatchedImageApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetPatchedImageApiResponseData()
	}
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
REST response for all response codes in API path /lifecycle/v4.3/config/hardware-providers/{hardwareProviderExtId}/connections/{connectionExtId}/server-identity-pools/{extId} Get operation
*/
type GetServerIdentityPoolApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetServerIdentityPoolApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *GetServerIdentityPoolApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetServerIdentityPoolApiResponse

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
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

func (p *GetServerIdentityPoolApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetServerIdentityPoolApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewGetServerIdentityPoolApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewGetServerIdentityPoolApiResponse() *GetServerIdentityPoolApiResponse {
	p := new(GetServerIdentityPoolApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.GetServerIdentityPoolApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetServerIdentityPoolApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetServerIdentityPoolApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetServerIdentityPoolApiResponseData()
	}
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
Group information
*/
type Group struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  ID of the group
	*/
	Id *string `json:"id,omitempty"`
	/*
	  Name of the group
	*/
	Name *string `json:"name,omitempty"`
}

func (p *Group) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias Group

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
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

func (p *Group) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Group
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewGroup()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Id != nil {
		p.Id = known.Id
	}
	if known.Name != nil {
		p.Name = known.Name
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "id")
	delete(allFields, "name")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewGroup() *Group {
	p := new(Group)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.Group"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Details of a hardware provider integration, such as a bare-metal or data center management system
*/
type HardwareProvider struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Supported authentication types for the hardware provider
	*/
	AuthTypes []AuthType `json:"authTypes"`
	/*
	  Number of available connections for the hardware provider
	*/
	AvailableConnectionCount *int `json:"availableConnectionCount"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Name of the hardware provider
	*/
	Name *string `json:"name"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	Type *HardwareProviderType `json:"type"`
	/*
	  Vendor of the hardware provider
	*/
	Vendor *string `json:"vendor"`
}

func (p *HardwareProvider) MarshalJSON() ([]byte, error) {
	type HardwareProviderProxy HardwareProvider

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*HardwareProviderProxy
		AuthTypes                []AuthType            `json:"authTypes,omitempty"`
		AvailableConnectionCount *int                  `json:"availableConnectionCount,omitempty"`
		Name                     *string               `json:"name,omitempty"`
		Type                     *HardwareProviderType `json:"type,omitempty"`
		Vendor                   *string               `json:"vendor,omitempty"`
	}{
		HardwareProviderProxy:    (*HardwareProviderProxy)(p),
		AuthTypes:                p.AuthTypes,
		AvailableConnectionCount: p.AvailableConnectionCount,
		Name:                     p.Name,
		Type:                     p.Type,
		Vendor:                   p.Vendor,
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

func (p *HardwareProvider) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias HardwareProvider
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewHardwareProvider()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.AuthTypes != nil {
		p.AuthTypes = known.AuthTypes
	}
	if known.AvailableConnectionCount != nil {
		p.AvailableConnectionCount = known.AvailableConnectionCount
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}
	if known.Type != nil {
		p.Type = known.Type
	}
	if known.Vendor != nil {
		p.Vendor = known.Vendor
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "authTypes")
	delete(allFields, "availableConnectionCount")
	delete(allFields, "extId")
	delete(allFields, "links")
	delete(allFields, "name")
	delete(allFields, "tenantId")
	delete(allFields, "type")
	delete(allFields, "vendor")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewHardwareProvider() *HardwareProvider {
	p := new(HardwareProvider)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.HardwareProvider"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Type of the hardware provider
*/
type HardwareProviderType int

const (
	HARDWAREPROVIDERTYPE_UNKNOWN    HardwareProviderType = 0
	HARDWAREPROVIDERTYPE_REDACTED   HardwareProviderType = 1
	HARDWAREPROVIDERTYPE_MANAGED    HardwareProviderType = 2
	HARDWAREPROVIDERTYPE_STANDALONE HardwareProviderType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *HardwareProviderType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"MANAGED",
		"STANDALONE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e HardwareProviderType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"MANAGED",
		"STANDALONE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *HardwareProviderType) index(name string) HardwareProviderType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"MANAGED",
		"STANDALONE",
	}
	for idx := range names {
		if names[idx] == name {
			return HardwareProviderType(idx)
		}
	}
	return HARDWAREPROVIDERTYPE_UNKNOWN
}

func (e *HardwareProviderType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for HardwareProviderType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *HardwareProviderType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e HardwareProviderType) Ref() *HardwareProviderType {
	return &e
}

/*
Hypervisor installation details for HCI configuration
*/
type HciHypervisorDetails struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Hostname of the hypervisor or host OS
	*/
	Hostname *string `json:"hostname"`
	/*

	 */
	ImageDetailsItemDiscriminator_ *string `json:"$imageDetailsItemDiscriminator,omitempty"`
	/*
	  Details of the hypervisor image to be installed on the node
	*/
	ImageDetails *OneOfHciHypervisorDetailsImageDetails `json:"imageDetails"`
	/*
	  List of nameserver IP addresses to be used by the host
	*/
	NameServers []import1.IPAddress `json:"nameServers,omitempty"`

	NetworkDetails *HciHypervisorNetworking `json:"networkDetails"`
	/*
	  List of NTP server IP addresses to be used by the host
	*/
	NtpServers []import1.IPAddressOrFQDN `json:"ntpServers,omitempty"`
}

func (p *HciHypervisorDetails) MarshalJSON() ([]byte, error) {
	type HciHypervisorDetailsProxy HciHypervisorDetails

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*HciHypervisorDetailsProxy
		Hostname       *string                                `json:"hostname,omitempty"`
		ImageDetails   *OneOfHciHypervisorDetailsImageDetails `json:"imageDetails,omitempty"`
		NetworkDetails *HciHypervisorNetworking               `json:"networkDetails,omitempty"`
	}{
		HciHypervisorDetailsProxy: (*HciHypervisorDetailsProxy)(p),
		Hostname:                  p.Hostname,
		ImageDetails:              p.ImageDetails,
		NetworkDetails:            p.NetworkDetails,
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

func (p *HciHypervisorDetails) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias HciHypervisorDetails
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewHciHypervisorDetails()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Hostname != nil {
		p.Hostname = known.Hostname
	}
	if known.ImageDetailsItemDiscriminator_ != nil {
		p.ImageDetailsItemDiscriminator_ = known.ImageDetailsItemDiscriminator_
	}
	if known.ImageDetails != nil {
		p.ImageDetails = known.ImageDetails
	}
	if known.NameServers != nil {
		p.NameServers = known.NameServers
	}
	if known.NetworkDetails != nil {
		p.NetworkDetails = known.NetworkDetails
	}
	if known.NtpServers != nil {
		p.NtpServers = known.NtpServers
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "hostname")
	delete(allFields, "$imageDetailsItemDiscriminator")
	delete(allFields, "imageDetails")
	delete(allFields, "nameServers")
	delete(allFields, "networkDetails")
	delete(allFields, "ntpServers")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewHciHypervisorDetails() *HciHypervisorDetails {
	p := new(HciHypervisorDetails)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.HciHypervisorDetails"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *HciHypervisorDetails) GetImageDetails() interface{} {
	if nil == p.ImageDetails {
		return nil
	}
	return p.ImageDetails.GetValue()
}

func (p *HciHypervisorDetails) SetImageDetails(v interface{}) error {
	if nil == p.ImageDetails {
		p.ImageDetails = NewOneOfHciHypervisorDetailsImageDetails()
	}
	e := p.ImageDetails.SetValue(v)
	if nil == e {
		if nil == p.ImageDetailsItemDiscriminator_ {
			p.ImageDetailsItemDiscriminator_ = new(string)
		}
		*p.ImageDetailsItemDiscriminator_ = *p.ImageDetails.Discriminator
	}
	return e
}

/*
Networking configuration for HCI hypervisor including management interface
*/
type HciHypervisorNetworking struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Management *ManagementNetwork `json:"management"`
}

func (p *HciHypervisorNetworking) MarshalJSON() ([]byte, error) {
	type HciHypervisorNetworkingProxy HciHypervisorNetworking

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*HciHypervisorNetworkingProxy
		Management *ManagementNetwork `json:"management,omitempty"`
	}{
		HciHypervisorNetworkingProxy: (*HciHypervisorNetworkingProxy)(p),
		Management:                   p.Management,
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

func (p *HciHypervisorNetworking) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias HciHypervisorNetworking
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewHciHypervisorNetworking()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Management != nil {
		p.Management = known.Management
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "management")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewHciHypervisorNetworking() *HciHypervisorNetworking {
	p := new(HciHypervisorNetworking)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.HciHypervisorNetworking"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Hyper Converged Infrastructure imaging configuration including hypervisor and AOS details
*/
type HciImagingDetails struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	AosDetails *AosInstallationDetails `json:"aosDetails"`

	HypervisorDetails *HciHypervisorDetails `json:"hypervisorDetails"`
}

func (p *HciImagingDetails) MarshalJSON() ([]byte, error) {
	type HciImagingDetailsProxy HciImagingDetails

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*HciImagingDetailsProxy
		AosDetails        *AosInstallationDetails `json:"aosDetails,omitempty"`
		HypervisorDetails *HciHypervisorDetails   `json:"hypervisorDetails,omitempty"`
	}{
		HciImagingDetailsProxy: (*HciImagingDetailsProxy)(p),
		AosDetails:             p.AosDetails,
		HypervisorDetails:      p.HypervisorDetails,
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

func (p *HciImagingDetails) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias HciImagingDetails
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewHciImagingDetails()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.AosDetails != nil {
		p.AosDetails = known.AosDetails
	}
	if known.HypervisorDetails != nil {
		p.HypervisorDetails = known.HypervisorDetails
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "aosDetails")
	delete(allFields, "hypervisorDetails")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewHciImagingDetails() *HciImagingDetails {
	p := new(HciImagingDetails)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.HciImagingDetails"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Configuration of the hypervisor or host OS used for patching
*/
type HostConfiguration struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Hostname of the hypervisor or host OS
	*/
	Hostname *string `json:"hostname"`
	/*
	  List of nameserver IP addresses to be used by the host
	*/
	Nameservers []import1.IPAddress `json:"nameservers,omitempty"`

	NetworkDetails *HostNetworkDetails `json:"networkDetails"`
	/*
	  List of NTP server IP addresses to be used by the host
	*/
	Ntpservers []import1.IPAddressOrFQDN `json:"ntpservers,omitempty"`
}

func (p *HostConfiguration) MarshalJSON() ([]byte, error) {
	type HostConfigurationProxy HostConfiguration

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*HostConfigurationProxy
		Hostname       *string             `json:"hostname,omitempty"`
		NetworkDetails *HostNetworkDetails `json:"networkDetails,omitempty"`
	}{
		HostConfigurationProxy: (*HostConfigurationProxy)(p),
		Hostname:               p.Hostname,
		NetworkDetails:         p.NetworkDetails,
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

func (p *HostConfiguration) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias HostConfiguration
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewHostConfiguration()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Hostname != nil {
		p.Hostname = known.Hostname
	}
	if known.Nameservers != nil {
		p.Nameservers = known.Nameservers
	}
	if known.NetworkDetails != nil {
		p.NetworkDetails = known.NetworkDetails
	}
	if known.Ntpservers != nil {
		p.Ntpservers = known.Ntpservers
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "hostname")
	delete(allFields, "nameservers")
	delete(allFields, "networkDetails")
	delete(allFields, "ntpservers")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewHostConfiguration() *HostConfiguration {
	p := new(HostConfiguration)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.HostConfiguration"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Hypervisor or host OS installation details of the node
*/
type HostInstallationDetails struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  External ID of a patched hypervisor or host OS image
	*/
	PatchedImageExtId *string `json:"patchedImageExtId"`
	/*
	  URL of the patched image to be installed on the node. If the URL is provided, nodes download the image from the given URL instead of downloading it from Foundation Central
	*/
	PatchedImageUrl *string `json:"patchedImageUrl,omitempty"`
}

func (p *HostInstallationDetails) MarshalJSON() ([]byte, error) {
	type HostInstallationDetailsProxy HostInstallationDetails

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*HostInstallationDetailsProxy
		PatchedImageExtId *string `json:"patchedImageExtId,omitempty"`
	}{
		HostInstallationDetailsProxy: (*HostInstallationDetailsProxy)(p),
		PatchedImageExtId:            p.PatchedImageExtId,
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

func (p *HostInstallationDetails) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias HostInstallationDetails
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewHostInstallationDetails()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.PatchedImageExtId != nil {
		p.PatchedImageExtId = known.PatchedImageExtId
	}
	if known.PatchedImageUrl != nil {
		p.PatchedImageUrl = known.PatchedImageUrl
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "patchedImageExtId")
	delete(allFields, "patchedImageUrl")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewHostInstallationDetails() *HostInstallationDetails {
	p := new(HostInstallationDetails)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.HostInstallationDetails"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Details of the management network for the workflow
*/
type HostManagementNetwork struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of physical adapters of a server
	*/
	Adapters []PhysicalAdapter `json:"adapters,omitempty"`

	BondSettings *BondSettings `json:"bondSettings,omitempty"`

	Gateway *import1.IPAddress `json:"gateway,omitempty"`

	Ip *import1.IPAddress `json:"ip,omitempty"`
	/*
	  Maximum transmission unit of the management network
	*/
	MtuBytes *int `json:"mtuBytes,omitempty"`
	/*
	  VLAN ID of the network
	*/
	VlanId *int `json:"vlanId,omitempty"`
}

func (p *HostManagementNetwork) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias HostManagementNetwork

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
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

func (p *HostManagementNetwork) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias HostManagementNetwork
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewHostManagementNetwork()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Adapters != nil {
		p.Adapters = known.Adapters
	}
	if known.BondSettings != nil {
		p.BondSettings = known.BondSettings
	}
	if known.Gateway != nil {
		p.Gateway = known.Gateway
	}
	if known.Ip != nil {
		p.Ip = known.Ip
	}
	if known.MtuBytes != nil {
		p.MtuBytes = known.MtuBytes
	}
	if known.VlanId != nil {
		p.VlanId = known.VlanId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "adapters")
	delete(allFields, "bondSettings")
	delete(allFields, "gateway")
	delete(allFields, "ip")
	delete(allFields, "mtuBytes")
	delete(allFields, "vlanId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewHostManagementNetwork() *HostManagementNetwork {
	p := new(HostManagementNetwork)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.HostManagementNetwork"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Details of the host network
*/
type HostNetworkDetails struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Management *ManagementNetwork `json:"management,omitempty"`
}

func (p *HostNetworkDetails) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias HostNetworkDetails

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
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

func (p *HostNetworkDetails) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias HostNetworkDetails
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewHostNetworkDetails()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Management != nil {
		p.Management = known.Management
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "management")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewHostNetworkDetails() *HostNetworkDetails {
	p := new(HostNetworkDetails)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.HostNetworkDetails"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Configuration details for the host layer (hypervisor). Includes network settings (IP, gateway, VLAN, bond), DNS servers, NTP servers, and hostname
*/
type HostReconfigurationDetails struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Hostname of the hypervisor or host OS
	*/
	Hostname *string `json:"hostname,omitempty"`
	/*
	  List of nameserver IP addresses to be used by the host
	*/
	NameServers []import1.IPAddress `json:"nameServers,omitempty"`

	NetworkDetails *HostNetworkDetails `json:"networkDetails,omitempty"`
	/*
	  List of NTP server IP addresses to be used by the host
	*/
	NtpServers []import1.IPAddressOrFQDN `json:"ntpServers,omitempty"`
}

func (p *HostReconfigurationDetails) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias HostReconfigurationDetails

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
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

func (p *HostReconfigurationDetails) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias HostReconfigurationDetails
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewHostReconfigurationDetails()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Hostname != nil {
		p.Hostname = known.Hostname
	}
	if known.NameServers != nil {
		p.NameServers = known.NameServers
	}
	if known.NetworkDetails != nil {
		p.NetworkDetails = known.NetworkDetails
	}
	if known.NtpServers != nil {
		p.NtpServers = known.NtpServers
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "hostname")
	delete(allFields, "nameServers")
	delete(allFields, "networkDetails")
	delete(allFields, "ntpServers")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewHostReconfigurationDetails() *HostReconfigurationDetails {
	p := new(HostReconfigurationDetails)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.HostReconfigurationDetails"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Type of the host installed or to be installed on the node
*/
type HostType int

const (
	HOSTTYPE_UNKNOWN  HostType = 0
	HOSTTYPE_REDACTED HostType = 1
	HOSTTYPE_AHV      HostType = 2
	HOSTTYPE_ESX      HostType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *HostType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"AHV",
		"ESX",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e HostType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"AHV",
		"ESX",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *HostType) index(name string) HostType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"AHV",
		"ESX",
	}
	for idx := range names {
		if names[idx] == name {
			return HostType(idx)
		}
	}
	return HOSTTYPE_UNKNOWN
}

func (e *HostType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for HostType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *HostType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e HostType) Ref() *HostType {
	return &e
}

/*
IP configuration using a pool of IP addresses from the hardware provider
*/
type IPPoolConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Unique identifier (UUID) of the IP pool
	*/
	IpPoolExtId *string `json:"ipPoolExtId,omitempty"`
	/*
	  Configuration for a single VLAN identifier
	*/
	VlanId *int `json:"vlanId,omitempty"`
}

func (p *IPPoolConfig) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias IPPoolConfig

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
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

func (p *IPPoolConfig) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias IPPoolConfig
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewIPPoolConfig()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.IpPoolExtId != nil {
		p.IpPoolExtId = known.IpPoolExtId
	}
	if known.VlanId != nil {
		p.VlanId = known.VlanId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "ipPoolExtId")
	delete(allFields, "vlanId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewIPPoolConfig() *IPPoolConfig {
	p := new(IPPoolConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.IPPoolConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Details of an installer image containing AOS and hypervisor software for node provisioning
*/
type Image struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Certificate chain for the image URL
	*/
	CertificateChain *string `json:"certificateChain,omitempty"`
	/*

	 */
	ChecksumItemDiscriminator_ *string `json:"$checksumItemDiscriminator,omitempty"`
	/*
	  Checksum value of the image. This is required and applicable only for AHV images
	*/
	Checksum *OneOfImageChecksum `json:"checksum,omitempty"`
	/*
	  Time when the image entity was created
	*/
	CreatedTime *time.Time `json:"createdTime,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`

	FileStatus *FileStatus `json:"fileStatus,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  URL from where the image metadata can be downloaded for an AOS image
	*/
	MetadataDownloadUrl *string `json:"metadataDownloadUrl,omitempty"`

	MetadataStatus *MetadataStatus `json:"metadataStatus,omitempty"`
	/*
	  Name of the image. This field is required when creating an image
	*/
	Name *string `json:"name"`

	Source *ImageSource `json:"source"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	Type *ImageType `json:"type"`
	/*
	  URL from where the image can be downloaded. This can be provided only through the update operation and only for images with REMOTE_URL as the source type
	*/
	Url *string `json:"url,omitempty"`
	/*
	  Version of the image
	*/
	Version *string `json:"version,omitempty"`
}

func (p *Image) MarshalJSON() ([]byte, error) {
	type ImageProxy Image

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*ImageProxy
		Name   *string      `json:"name,omitempty"`
		Source *ImageSource `json:"source,omitempty"`
		Type   *ImageType   `json:"type,omitempty"`
	}{
		ImageProxy: (*ImageProxy)(p),
		Name:       p.Name,
		Source:     p.Source,
		Type:       p.Type,
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

func (p *Image) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Image
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewImage()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.CertificateChain != nil {
		p.CertificateChain = known.CertificateChain
	}
	if known.ChecksumItemDiscriminator_ != nil {
		p.ChecksumItemDiscriminator_ = known.ChecksumItemDiscriminator_
	}
	if known.Checksum != nil {
		p.Checksum = known.Checksum
	}
	if known.CreatedTime != nil {
		p.CreatedTime = known.CreatedTime
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.FileStatus != nil {
		p.FileStatus = known.FileStatus
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.MetadataDownloadUrl != nil {
		p.MetadataDownloadUrl = known.MetadataDownloadUrl
	}
	if known.MetadataStatus != nil {
		p.MetadataStatus = known.MetadataStatus
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.Source != nil {
		p.Source = known.Source
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}
	if known.Type != nil {
		p.Type = known.Type
	}
	if known.Url != nil {
		p.Url = known.Url
	}
	if known.Version != nil {
		p.Version = known.Version
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "certificateChain")
	delete(allFields, "$checksumItemDiscriminator")
	delete(allFields, "checksum")
	delete(allFields, "createdTime")
	delete(allFields, "extId")
	delete(allFields, "fileStatus")
	delete(allFields, "links")
	delete(allFields, "metadataDownloadUrl")
	delete(allFields, "metadataStatus")
	delete(allFields, "name")
	delete(allFields, "source")
	delete(allFields, "tenantId")
	delete(allFields, "type")
	delete(allFields, "url")
	delete(allFields, "version")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewImage() *Image {
	p := new(Image)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.Image"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *Image) GetChecksum() interface{} {
	if nil == p.Checksum {
		return nil
	}
	return p.Checksum.GetValue()
}

func (p *Image) SetChecksum(v interface{}) error {
	if nil == p.Checksum {
		p.Checksum = NewOneOfImageChecksum()
	}
	e := p.Checksum.SetValue(v)
	if nil == e {
		if nil == p.ChecksumItemDiscriminator_ {
			p.ChecksumItemDiscriminator_ = new(string)
		}
		*p.ChecksumItemDiscriminator_ = *p.Checksum.Discriminator
	}
	return e
}

/*
MD5 checksum for the image. This is required and applicable only for AHV images
*/
type ImageMd5Checksum struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  MD5 checksum value in hexadecimal format (32 characters)
	*/
	HexDigest *string `json:"hexDigest"`
}

func (p *ImageMd5Checksum) MarshalJSON() ([]byte, error) {
	type ImageMd5ChecksumProxy ImageMd5Checksum

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*ImageMd5ChecksumProxy
		HexDigest *string `json:"hexDigest,omitempty"`
	}{
		ImageMd5ChecksumProxy: (*ImageMd5ChecksumProxy)(p),
		HexDigest:             p.HexDigest,
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

func (p *ImageMd5Checksum) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ImageMd5Checksum
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewImageMd5Checksum()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.HexDigest != nil {
		p.HexDigest = known.HexDigest
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "hexDigest")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewImageMd5Checksum() *ImageMd5Checksum {
	p := new(ImageMd5Checksum)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.ImageMd5Checksum"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /lifecycle/v4.3/config/nodes/{extId}/$actions/image Post operation
*/
type ImageNodeApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfImageNodeApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ImageNodeApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ImageNodeApiResponse

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
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

func (p *ImageNodeApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ImageNodeApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewImageNodeApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewImageNodeApiResponse() *ImageNodeApiResponse {
	p := new(ImageNodeApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.ImageNodeApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ImageNodeApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ImageNodeApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfImageNodeApiResponseData()
	}
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
Request specification to image a node
*/
type ImageNodeSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	ConfigurationItemDiscriminator_ *string `json:"$configurationItemDiscriminator,omitempty"`
	/*
	  Configuration for imaging a node
	*/
	Configuration *OneOfImageNodeSpecConfiguration `json:"configuration"`
}

func (p *ImageNodeSpec) MarshalJSON() ([]byte, error) {
	type ImageNodeSpecProxy ImageNodeSpec

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*ImageNodeSpecProxy
		Configuration *OneOfImageNodeSpecConfiguration `json:"configuration,omitempty"`
	}{
		ImageNodeSpecProxy: (*ImageNodeSpecProxy)(p),
		Configuration:      p.Configuration,
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

func (p *ImageNodeSpec) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ImageNodeSpec
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewImageNodeSpec()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ConfigurationItemDiscriminator_ != nil {
		p.ConfigurationItemDiscriminator_ = known.ConfigurationItemDiscriminator_
	}
	if known.Configuration != nil {
		p.Configuration = known.Configuration
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$configurationItemDiscriminator")
	delete(allFields, "configuration")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewImageNodeSpec() *ImageNodeSpec {
	p := new(ImageNodeSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.ImageNodeSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ImageNodeSpec) GetConfiguration() interface{} {
	if nil == p.Configuration {
		return nil
	}
	return p.Configuration.GetValue()
}

func (p *ImageNodeSpec) SetConfiguration(v interface{}) error {
	if nil == p.Configuration {
		p.Configuration = NewOneOfImageNodeSpecConfiguration()
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
SHA-256 checksum for the image. This is required and applicable only for AHV images
*/
type ImageSha256Checksum struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  SHA-256 checksum value in hexadecimal format (64 characters)
	*/
	HexDigest *string `json:"hexDigest"`
}

func (p *ImageSha256Checksum) MarshalJSON() ([]byte, error) {
	type ImageSha256ChecksumProxy ImageSha256Checksum

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*ImageSha256ChecksumProxy
		HexDigest *string `json:"hexDigest,omitempty"`
	}{
		ImageSha256ChecksumProxy: (*ImageSha256ChecksumProxy)(p),
		HexDigest:                p.HexDigest,
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

func (p *ImageSha256Checksum) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ImageSha256Checksum
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewImageSha256Checksum()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.HexDigest != nil {
		p.HexDigest = known.HexDigest
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "hexDigest")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewImageSha256Checksum() *ImageSha256Checksum {
	p := new(ImageSha256Checksum)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.ImageSha256Checksum"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Source of the image: LOCAL (image managed and hosted by Foundation Central), REMOTE_URL (image hosted outside of Foundation Central, and nodes will download the image from the given URL at the time of imaging)
*/
type ImageSource int

const (
	IMAGESOURCE_UNKNOWN    ImageSource = 0
	IMAGESOURCE_REDACTED   ImageSource = 1
	IMAGESOURCE_LOCAL      ImageSource = 2
	IMAGESOURCE_REMOTE_URL ImageSource = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ImageSource) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"LOCAL",
		"REMOTE_URL",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ImageSource) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"LOCAL",
		"REMOTE_URL",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ImageSource) index(name string) ImageSource {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"LOCAL",
		"REMOTE_URL",
	}
	for idx := range names {
		if names[idx] == name {
			return ImageSource(idx)
		}
	}
	return IMAGESOURCE_UNKNOWN
}

func (e *ImageSource) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ImageSource:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ImageSource) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ImageSource) Ref() *ImageSource {
	return &e
}

/*
Type of the image: AOS (AOS image), AHV (AHV image), or ESX (ESX image). This field is required when creating an image
*/
type ImageType int

const (
	IMAGETYPE_UNKNOWN  ImageType = 0
	IMAGETYPE_REDACTED ImageType = 1
	IMAGETYPE_AOS      ImageType = 2
	IMAGETYPE_AHV      ImageType = 3
	IMAGETYPE_ESX      ImageType = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ImageType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"AOS",
		"AHV",
		"ESX",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ImageType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"AOS",
		"AHV",
		"ESX",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ImageType) index(name string) ImageType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"AOS",
		"AHV",
		"ESX",
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

/*
BMC IP addressing configuration with separate in-band and out-of-band options
*/
type InOutBandBmcIPConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	InBand *BmcIPConfig `json:"inBand,omitempty"`

	OutOfBand *BmcIPConfig `json:"outOfBand,omitempty"`
}

func (p *InOutBandBmcIPConfig) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias InOutBandBmcIPConfig

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
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

func (p *InOutBandBmcIPConfig) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias InOutBandBmcIPConfig
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewInOutBandBmcIPConfig()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.InBand != nil {
		p.InBand = known.InBand
	}
	if known.OutOfBand != nil {
		p.OutOfBand = known.OutOfBand
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "inBand")
	delete(allFields, "outOfBand")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewInOutBandBmcIPConfig() *InOutBandBmcIPConfig {
	p := new(InOutBandBmcIPConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.InOutBandBmcIPConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Endpoint configuration using IP addresses
*/
type IpAddressEndpoint struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of IP addresses for the endpoint
	*/
	IpAddresses []import1.IPAddressOrFQDN `json:"ipAddresses"`
}

func (p *IpAddressEndpoint) MarshalJSON() ([]byte, error) {
	type IpAddressEndpointProxy IpAddressEndpoint

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*IpAddressEndpointProxy
		IpAddresses []import1.IPAddressOrFQDN `json:"ipAddresses,omitempty"`
	}{
		IpAddressEndpointProxy: (*IpAddressEndpointProxy)(p),
		IpAddresses:            p.IpAddresses,
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

func (p *IpAddressEndpoint) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias IpAddressEndpoint
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewIpAddressEndpoint()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.IpAddresses != nil {
		p.IpAddresses = known.IpAddresses
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "ipAddresses")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewIpAddressEndpoint() *IpAddressEndpoint {
	p := new(IpAddressEndpoint)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.IpAddressEndpoint"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Details of a IP address pool available in the hardware provider
*/
type IpPool struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	DisplayNameMapping *IpPoolDisplayNameMapping `json:"displayNameMapping,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  List of groups this IP address pool belongs to
	*/
	Groups []PoolGroup `json:"groups,omitempty"`
	/*
	  Number of available IPv4 addresses in the IP address pool
	*/
	Ipv4AvailableCount *int `json:"ipv4AvailableCount,omitempty"`
	/*
	  Number of available IPv6 addresses in the IP address pool
	*/
	Ipv6AvailableCount *int `json:"ipv6AvailableCount,omitempty"`
	/*
	  Time when the IP address pool was last updated
	*/
	LastUpdatedTime *time.Time `json:"lastUpdatedTime"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Name of the IP address pool
	*/
	Name *string `json:"name,omitempty"`
	/*
	  Reference ID of the IP address pool in the hardware provider
	*/
	ReferenceId *string `json:"referenceId,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *IpPool) MarshalJSON() ([]byte, error) {
	type IpPoolProxy IpPool

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*IpPoolProxy
		LastUpdatedTime *time.Time `json:"lastUpdatedTime,omitempty"`
	}{
		IpPoolProxy:     (*IpPoolProxy)(p),
		LastUpdatedTime: p.LastUpdatedTime,
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

func (p *IpPool) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias IpPool
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewIpPool()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DisplayNameMapping != nil {
		p.DisplayNameMapping = known.DisplayNameMapping
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.Groups != nil {
		p.Groups = known.Groups
	}
	if known.Ipv4AvailableCount != nil {
		p.Ipv4AvailableCount = known.Ipv4AvailableCount
	}
	if known.Ipv6AvailableCount != nil {
		p.Ipv6AvailableCount = known.Ipv6AvailableCount
	}
	if known.LastUpdatedTime != nil {
		p.LastUpdatedTime = known.LastUpdatedTime
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.ReferenceId != nil {
		p.ReferenceId = known.ReferenceId
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "displayNameMapping")
	delete(allFields, "extId")
	delete(allFields, "groups")
	delete(allFields, "ipv4AvailableCount")
	delete(allFields, "ipv6AvailableCount")
	delete(allFields, "lastUpdatedTime")
	delete(allFields, "links")
	delete(allFields, "name")
	delete(allFields, "referenceId")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewIpPool() *IpPool {
	p := new(IpPool)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.IpPool"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Mapping between unified attribute names and vendor-specific attribute names for IP address pools
*/
type IpPoolDisplayNameMapping struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Vendor-specific attribute name for the group of the IP address pool
	*/
	Group *string `json:"group,omitempty"`
}

func (p *IpPoolDisplayNameMapping) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias IpPoolDisplayNameMapping

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
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

func (p *IpPoolDisplayNameMapping) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias IpPoolDisplayNameMapping
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewIpPoolDisplayNameMapping()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Group != nil {
		p.Group = known.Group
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "group")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewIpPoolDisplayNameMapping() *IpPoolDisplayNameMapping {
	p := new(IpPoolDisplayNameMapping)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.IpPoolDisplayNameMapping"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Range of IP addresses
*/
type IpRange struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	EndIp *import1.IPAddress `json:"endIp"`

	StartIp *import1.IPAddress `json:"startIp"`
}

func (p *IpRange) MarshalJSON() ([]byte, error) {
	type IpRangeProxy IpRange

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*IpRangeProxy
		EndIp   *import1.IPAddress `json:"endIp,omitempty"`
		StartIp *import1.IPAddress `json:"startIp,omitempty"`
	}{
		IpRangeProxy: (*IpRangeProxy)(p),
		EndIp:        p.EndIp,
		StartIp:      p.StartIp,
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

func (p *IpRange) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias IpRange
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewIpRange()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.EndIp != nil {
		p.EndIp = known.EndIp
	}
	if known.StartIp != nil {
		p.StartIp = known.StartIp
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "endIp")
	delete(allFields, "startIp")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewIpRange() *IpRange {
	p := new(IpRange)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.IpRange"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Endpoint configuration using IP ranges
*/
type IpRangeEndpoint struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of IP ranges for the endpoint
	*/
	IpRanges []IpRange `json:"ipRanges"`
}

func (p *IpRangeEndpoint) MarshalJSON() ([]byte, error) {
	type IpRangeEndpointProxy IpRangeEndpoint

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*IpRangeEndpointProxy
		IpRanges []IpRange `json:"ipRanges,omitempty"`
	}{
		IpRangeEndpointProxy: (*IpRangeEndpointProxy)(p),
		IpRanges:             p.IpRanges,
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

func (p *IpRangeEndpoint) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias IpRangeEndpoint
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewIpRangeEndpoint()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.IpRanges != nil {
		p.IpRanges = known.IpRanges
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "ipRanges")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewIpRangeEndpoint() *IpRangeEndpoint {
	p := new(IpRangeEndpoint)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.IpRangeEndpoint"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
LACP (Link Aggregation Control Protocol) bonding configuration for dynamic link aggregation
*/
type LacpBond struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	LacpSettings *LacpSettings `json:"lacpSettings"`

	Nics *Nics `json:"nics"`
}

func (p *LacpBond) MarshalJSON() ([]byte, error) {
	type LacpBondProxy LacpBond

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*LacpBondProxy
		LacpSettings *LacpSettings `json:"lacpSettings,omitempty"`
		Nics         *Nics         `json:"nics,omitempty"`
	}{
		LacpBondProxy: (*LacpBondProxy)(p),
		LacpSettings:  p.LacpSettings,
		Nics:          p.Nics,
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

func (p *LacpBond) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias LacpBond
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewLacpBond()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.LacpSettings != nil {
		p.LacpSettings = known.LacpSettings
	}
	if known.Nics != nil {
		p.Nics = known.Nics
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "lacpSettings")
	delete(allFields, "nics")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewLacpBond() *LacpBond {
	p := new(LacpBond)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.LacpBond"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Rate of the LACP
*/
type LacpRate int

const (
	LACPRATE_UNKNOWN  LacpRate = 0
	LACPRATE_REDACTED LacpRate = 1
	LACPRATE_SLOW     LacpRate = 2
	LACPRATE_FAST     LacpRate = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *LacpRate) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SLOW",
		"FAST",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e LacpRate) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SLOW",
		"FAST",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *LacpRate) index(name string) LacpRate {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SLOW",
		"FAST",
	}
	for idx := range names {
		if names[idx] == name {
			return LacpRate(idx)
		}
	}
	return LACPRATE_UNKNOWN
}

func (e *LacpRate) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for LacpRate:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *LacpRate) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e LacpRate) Ref() *LacpRate {
	return &e
}

/*
Details of the LACP settings
*/
type LacpSettings struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Rate *LacpRate `json:"rate,omitempty"`
}

func (p *LacpSettings) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias LacpSettings

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
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

func (p *LacpSettings) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias LacpSettings
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewLacpSettings()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Rate != nil {
		p.Rate = known.Rate
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "rate")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewLacpSettings() *LacpSettings {
	p := new(LacpSettings)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.LacpSettings"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /lifecycle/v4.3/config/claim-tokens Get operation
*/
type ListClaimTokensApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListClaimTokensApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ListClaimTokensApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ListClaimTokensApiResponse

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
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

func (p *ListClaimTokensApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ListClaimTokensApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewListClaimTokensApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewListClaimTokensApiResponse() *ListClaimTokensApiResponse {
	p := new(ListClaimTokensApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.ListClaimTokensApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListClaimTokensApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListClaimTokensApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListClaimTokensApiResponseData()
	}
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
REST response for all response codes in API path /lifecycle/v4.3/config/hardware-providers/{hardwareProviderExtId}/connections/{connectionExtId}/nodes Get operation
*/
type ListConnectionNodesApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListConnectionNodesApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ListConnectionNodesApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ListConnectionNodesApiResponse

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
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

func (p *ListConnectionNodesApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ListConnectionNodesApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewListConnectionNodesApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewListConnectionNodesApiResponse() *ListConnectionNodesApiResponse {
	p := new(ListConnectionNodesApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.ListConnectionNodesApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListConnectionNodesApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListConnectionNodesApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListConnectionNodesApiResponseData()
	}
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
REST response for all response codes in API path /lifecycle/v4.3/config/hardware-providers/{hardwareProviderExtId}/connections Get operation
*/
type ListConnectionsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListConnectionsApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ListConnectionsApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ListConnectionsApiResponse

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
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

func (p *ListConnectionsApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ListConnectionsApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewListConnectionsApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewListConnectionsApiResponse() *ListConnectionsApiResponse {
	p := new(ListConnectionsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.ListConnectionsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListConnectionsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListConnectionsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListConnectionsApiResponseData()
	}
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
REST response for all response codes in API path /lifecycle/v4.3/config/hardware-providers Get operation
*/
type ListHardwareProvidersApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListHardwareProvidersApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ListHardwareProvidersApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ListHardwareProvidersApiResponse

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
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

func (p *ListHardwareProvidersApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ListHardwareProvidersApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewListHardwareProvidersApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewListHardwareProvidersApiResponse() *ListHardwareProvidersApiResponse {
	p := new(ListHardwareProvidersApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.ListHardwareProvidersApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListHardwareProvidersApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListHardwareProvidersApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListHardwareProvidersApiResponseData()
	}
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
REST response for all response codes in API path /lifecycle/v4.3/config/images Get operation
*/
type ListImagesApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListImagesApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ListImagesApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ListImagesApiResponse

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
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

func (p *ListImagesApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ListImagesApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewListImagesApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewListImagesApiResponse() *ListImagesApiResponse {
	p := new(ListImagesApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.ListImagesApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListImagesApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListImagesApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListImagesApiResponseData()
	}
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
REST response for all response codes in API path /lifecycle/v4.3/config/hardware-providers/{hardwareProviderExtId}/connections/{connectionExtId}/ip-pools Get operation
*/
type ListIpPoolsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListIpPoolsApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ListIpPoolsApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ListIpPoolsApiResponse

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
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

func (p *ListIpPoolsApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ListIpPoolsApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewListIpPoolsApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewListIpPoolsApiResponse() *ListIpPoolsApiResponse {
	p := new(ListIpPoolsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.ListIpPoolsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListIpPoolsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListIpPoolsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListIpPoolsApiResponseData()
	}
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
REST response for all response codes in API path /lifecycle/v4.3/config/hardware-providers/{hardwareProviderExtId}/connections/{connectionExtId}/mac-pools Get operation
*/
type ListMacPoolsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListMacPoolsApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ListMacPoolsApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ListMacPoolsApiResponse

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
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

func (p *ListMacPoolsApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ListMacPoolsApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewListMacPoolsApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewListMacPoolsApiResponse() *ListMacPoolsApiResponse {
	p := new(ListMacPoolsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.ListMacPoolsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListMacPoolsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListMacPoolsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListMacPoolsApiResponseData()
	}
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
REST response for all response codes in API path /lifecycle/v4.3/config/nodes Get operation
*/
type ListNodesApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListNodesApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ListNodesApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ListNodesApiResponse

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
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

func (p *ListNodesApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ListNodesApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewListNodesApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewListNodesApiResponse() *ListNodesApiResponse {
	p := new(ListNodesApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.ListNodesApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListNodesApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListNodesApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListNodesApiResponseData()
	}
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
REST response for all response codes in API path /lifecycle/v4.3/config/claim-tokens/{claimTokenExtId}/nodes Get operation
*/
type ListNodesByClaimTokenApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListNodesByClaimTokenApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ListNodesByClaimTokenApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ListNodesByClaimTokenApiResponse

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
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

func (p *ListNodesByClaimTokenApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ListNodesByClaimTokenApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewListNodesByClaimTokenApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewListNodesByClaimTokenApiResponse() *ListNodesByClaimTokenApiResponse {
	p := new(ListNodesByClaimTokenApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.ListNodesByClaimTokenApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListNodesByClaimTokenApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListNodesByClaimTokenApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListNodesByClaimTokenApiResponseData()
	}
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
REST response for all response codes in API path /lifecycle/v4.3/config/patched-images Get operation
*/
type ListPatchedImagesApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListPatchedImagesApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ListPatchedImagesApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ListPatchedImagesApiResponse

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
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

func (p *ListPatchedImagesApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ListPatchedImagesApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewListPatchedImagesApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewListPatchedImagesApiResponse() *ListPatchedImagesApiResponse {
	p := new(ListPatchedImagesApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.ListPatchedImagesApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListPatchedImagesApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListPatchedImagesApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListPatchedImagesApiResponseData()
	}
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
REST response for all response codes in API path /lifecycle/v4.3/config/hardware-providers/{hardwareProviderExtId}/connections/{connectionExtId}/server-identity-pools Get operation
*/
type ListServerIdentityPoolsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListServerIdentityPoolsApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ListServerIdentityPoolsApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ListServerIdentityPoolsApiResponse

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
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

func (p *ListServerIdentityPoolsApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ListServerIdentityPoolsApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewListServerIdentityPoolsApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewListServerIdentityPoolsApiResponse() *ListServerIdentityPoolsApiResponse {
	p := new(ListServerIdentityPoolsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.ListServerIdentityPoolsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListServerIdentityPoolsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListServerIdentityPoolsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListServerIdentityPoolsApiResponseData()
	}
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
Details of the AOS image managed by Foundation Central
*/
type LocalAOSImageDetails struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  External ID of the AOS image managed by Foundation Central
	*/
	ExtId *string `json:"extId"`
}

func (p *LocalAOSImageDetails) MarshalJSON() ([]byte, error) {
	type LocalAOSImageDetailsProxy LocalAOSImageDetails

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*LocalAOSImageDetailsProxy
		ExtId *string `json:"extId,omitempty"`
	}{
		LocalAOSImageDetailsProxy: (*LocalAOSImageDetailsProxy)(p),
		ExtId:                     p.ExtId,
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

func (p *LocalAOSImageDetails) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias LocalAOSImageDetails
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewLocalAOSImageDetails()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "extId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewLocalAOSImageDetails() *LocalAOSImageDetails {
	p := new(LocalAOSImageDetails)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.LocalAOSImageDetails"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Details of the uploaded host image
*/
type LocalHostImageDetails struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  External ID of the uploaded host image
	*/
	ExtId *string `json:"extId"`
}

func (p *LocalHostImageDetails) MarshalJSON() ([]byte, error) {
	type LocalHostImageDetailsProxy LocalHostImageDetails

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*LocalHostImageDetailsProxy
		ExtId *string `json:"extId,omitempty"`
	}{
		LocalHostImageDetailsProxy: (*LocalHostImageDetailsProxy)(p),
		ExtId:                      p.ExtId,
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

func (p *LocalHostImageDetails) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias LocalHostImageDetails
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewLocalHostImageDetails()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "extId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewLocalHostImageDetails() *LocalHostImageDetails {
	p := new(LocalHostImageDetails)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.LocalHostImageDetails"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Details of a MAC address pool available in the hardware provider
*/
type MacPool struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Number of available MAC addresses in the MAC address pool
	*/
	AvailableCount *int `json:"availableCount,omitempty"`

	DisplayNameMapping *MacPoolDisplayNameMapping `json:"displayNameMapping,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  List of groups this MAC address pool belongs to
	*/
	Groups []PoolGroup `json:"groups,omitempty"`
	/*
	  Time when the MAC address pool was last updated
	*/
	LastUpdatedTime *time.Time `json:"lastUpdatedTime"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Name of the MAC address pool
	*/
	Name *string `json:"name,omitempty"`
	/*
	  Reference ID of the MAC address pool in the hardware provider
	*/
	ReferenceId *string `json:"referenceId,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *MacPool) MarshalJSON() ([]byte, error) {
	type MacPoolProxy MacPool

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*MacPoolProxy
		LastUpdatedTime *time.Time `json:"lastUpdatedTime,omitempty"`
	}{
		MacPoolProxy:    (*MacPoolProxy)(p),
		LastUpdatedTime: p.LastUpdatedTime,
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

func (p *MacPool) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias MacPool
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewMacPool()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.AvailableCount != nil {
		p.AvailableCount = known.AvailableCount
	}
	if known.DisplayNameMapping != nil {
		p.DisplayNameMapping = known.DisplayNameMapping
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.Groups != nil {
		p.Groups = known.Groups
	}
	if known.LastUpdatedTime != nil {
		p.LastUpdatedTime = known.LastUpdatedTime
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.ReferenceId != nil {
		p.ReferenceId = known.ReferenceId
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "availableCount")
	delete(allFields, "displayNameMapping")
	delete(allFields, "extId")
	delete(allFields, "groups")
	delete(allFields, "lastUpdatedTime")
	delete(allFields, "links")
	delete(allFields, "name")
	delete(allFields, "referenceId")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewMacPool() *MacPool {
	p := new(MacPool)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.MacPool"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Mapping between unified attribute names and vendor-specific attribute names for MAC address pools
*/
type MacPoolDisplayNameMapping struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Vendor-specific attribute name for the group of the MAC address pool
	*/
	Group *string `json:"group,omitempty"`
}

func (p *MacPoolDisplayNameMapping) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias MacPoolDisplayNameMapping

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
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

func (p *MacPoolDisplayNameMapping) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias MacPoolDisplayNameMapping
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewMacPoolDisplayNameMapping()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Group != nil {
		p.Group = known.Group
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "group")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewMacPoolDisplayNameMapping() *MacPoolDisplayNameMapping {
	p := new(MacPoolDisplayNameMapping)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.MacPoolDisplayNameMapping"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Systems that manage the node
*/
type ManagedBy int

const (
	MANAGEDBY_UNKNOWN            ManagedBy = 0
	MANAGEDBY_REDACTED           ManagedBy = 1
	MANAGEDBY_FOUNDATION_CENTRAL ManagedBy = 2
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ManagedBy) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"FOUNDATION_CENTRAL",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ManagedBy) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"FOUNDATION_CENTRAL",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ManagedBy) index(name string) ManagedBy {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"FOUNDATION_CENTRAL",
	}
	for idx := range names {
		if names[idx] == name {
			return ManagedBy(idx)
		}
	}
	return MANAGEDBY_UNKNOWN
}

func (e *ManagedBy) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ManagedBy:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ManagedBy) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ManagedBy) Ref() *ManagedBy {
	return &e
}

/*
Details of the management network
*/
type ManagementNetwork struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	BondSettings *BondSettings `json:"bondSettings,omitempty"`

	Gateway *import1.IPAddress `json:"gateway,omitempty"`

	Ip *import1.IPAddress `json:"ip,omitempty"`
	/*
	  Maximum transmission unit of the management network
	*/
	MtuBytes *int `json:"mtuBytes,omitempty"`
	/*
	  VLAN ID of the network
	*/
	VlanId *int `json:"vlanId,omitempty"`
}

func (p *ManagementNetwork) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ManagementNetwork

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
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

func (p *ManagementNetwork) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ManagementNetwork
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewManagementNetwork()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.BondSettings != nil {
		p.BondSettings = known.BondSettings
	}
	if known.Gateway != nil {
		p.Gateway = known.Gateway
	}
	if known.Ip != nil {
		p.Ip = known.Ip
	}
	if known.MtuBytes != nil {
		p.MtuBytes = known.MtuBytes
	}
	if known.VlanId != nil {
		p.VlanId = known.VlanId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "bondSettings")
	delete(allFields, "gateway")
	delete(allFields, "ip")
	delete(allFields, "mtuBytes")
	delete(allFields, "vlanId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewManagementNetwork() *ManagementNetwork {
	p := new(ManagementNetwork)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.ManagementNetwork"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Status of the image metadata (for example FILE_AVAILABLE, UPLOADING, NOT_AVAILABLE)
*/
type MetadataStatus int

const (
	METADATASTATUS_UNKNOWN        MetadataStatus = 0
	METADATASTATUS_REDACTED       MetadataStatus = 1
	METADATASTATUS_FILE_AVAILABLE MetadataStatus = 2
	METADATASTATUS_UPLOADING      MetadataStatus = 3
	METADATASTATUS_UPLOAD_FAILED  MetadataStatus = 4
	METADATASTATUS_NOT_AVAILABLE  MetadataStatus = 5
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *MetadataStatus) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"FILE_AVAILABLE",
		"UPLOADING",
		"UPLOAD_FAILED",
		"NOT_AVAILABLE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e MetadataStatus) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"FILE_AVAILABLE",
		"UPLOADING",
		"UPLOAD_FAILED",
		"NOT_AVAILABLE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *MetadataStatus) index(name string) MetadataStatus {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"FILE_AVAILABLE",
		"UPLOADING",
		"UPLOAD_FAILED",
		"NOT_AVAILABLE",
	}
	for idx := range names {
		if names[idx] == name {
			return MetadataStatus(idx)
		}
	}
	return METADATASTATUS_UNKNOWN
}

func (e *MetadataStatus) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for MetadataStatus:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *MetadataStatus) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e MetadataStatus) Ref() *MetadataStatus {
	return &e
}

/*
Forward Error Correction (FEC) mode for the network adaptor
*/
type NetworkAdaptorFecMode int

const (
	NETWORKADAPTORFECMODE_UNKNOWN  NetworkAdaptorFecMode = 0
	NETWORKADAPTORFECMODE_REDACTED NetworkAdaptorFecMode = 1
	NETWORKADAPTORFECMODE_CL91     NetworkAdaptorFecMode = 2
	NETWORKADAPTORFECMODE_CL74     NetworkAdaptorFecMode = 3
	NETWORKADAPTORFECMODE_OFF      NetworkAdaptorFecMode = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *NetworkAdaptorFecMode) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CL91",
		"CL74",
		"OFF",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e NetworkAdaptorFecMode) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CL91",
		"CL74",
		"OFF",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *NetworkAdaptorFecMode) index(name string) NetworkAdaptorFecMode {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CL91",
		"CL74",
		"OFF",
	}
	for idx := range names {
		if names[idx] == name {
			return NetworkAdaptorFecMode(idx)
		}
	}
	return NETWORKADAPTORFECMODE_UNKNOWN
}

func (e *NetworkAdaptorFecMode) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for NetworkAdaptorFecMode:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *NetworkAdaptorFecMode) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e NetworkAdaptorFecMode) Ref() *NetworkAdaptorFecMode {
	return &e
}

/*
Details of the network
*/
type NetworkDetails struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Gateway *import1.IPAddress `json:"gateway,omitempty"`

	Ip *import1.IPAddress `json:"ip,omitempty"`
	/*
	  VLAN ID of the network
	*/
	VlanId *int `json:"vlanId,omitempty"`
}

func (p *NetworkDetails) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias NetworkDetails

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
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

func (p *NetworkDetails) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias NetworkDetails
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewNetworkDetails()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Gateway != nil {
		p.Gateway = known.Gateway
	}
	if known.Ip != nil {
		p.Ip = known.Ip
	}
	if known.VlanId != nil {
		p.VlanId = known.VlanId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "gateway")
	delete(allFields, "ip")
	delete(allFields, "vlanId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewNetworkDetails() *NetworkDetails {
	p := new(NetworkDetails)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.NetworkDetails"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
A network group that defines network segmentation with VLANs and physical port mappings
*/
type NetworkGroup struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Maximum Transmission Unit (MTU) size for the network group
	*/
	MtuBytes *int `json:"mtuBytes,omitempty"`
	/*
	  Name of the network group
	*/
	Name *string `json:"name"`
	/*
	  List of physical port mappings by MAC address
	*/
	PortMappings []PortMapping `json:"portMappings,omitempty"`

	Type *NetworkGroupType `json:"type"`
	/*
	  VLAN IDs or VLAN ranges assigned to the network group (e.g., '100' or '100-200')
	*/
	Vlans []VlanInfo `json:"vlans,omitempty"`
}

func (p *NetworkGroup) MarshalJSON() ([]byte, error) {
	type NetworkGroupProxy NetworkGroup

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*NetworkGroupProxy
		Name *string           `json:"name,omitempty"`
		Type *NetworkGroupType `json:"type,omitempty"`
	}{
		NetworkGroupProxy: (*NetworkGroupProxy)(p),
		Name:              p.Name,
		Type:              p.Type,
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

func (p *NetworkGroup) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias NetworkGroup
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewNetworkGroup()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.MtuBytes != nil {
		p.MtuBytes = known.MtuBytes
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.PortMappings != nil {
		p.PortMappings = known.PortMappings
	}
	if known.Type != nil {
		p.Type = known.Type
	}
	if known.Vlans != nil {
		p.Vlans = known.Vlans
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "mtuBytes")
	delete(allFields, "name")
	delete(allFields, "portMappings")
	delete(allFields, "type")
	delete(allFields, "vlans")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewNetworkGroup() *NetworkGroup {
	p := new(NetworkGroup)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.NetworkGroup"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Type of the network group defining its purpose
*/
type NetworkGroupType int

const (
	NETWORKGROUPTYPE_UNKNOWN  NetworkGroupType = 0
	NETWORKGROUPTYPE_REDACTED NetworkGroupType = 1
	NETWORKGROUPTYPE_INFRA    NetworkGroupType = 2
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *NetworkGroupType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"INFRA",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e NetworkGroupType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"INFRA",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *NetworkGroupType) index(name string) NetworkGroupType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"INFRA",
	}
	for idx := range names {
		if names[idx] == name {
			return NetworkGroupType(idx)
		}
	}
	return NETWORKGROUPTYPE_UNKNOWN
}

func (e *NetworkGroupType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for NetworkGroupType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *NetworkGroupType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e NetworkGroupType) Ref() *NetworkGroupType {
	return &e
}

/*
Details of the NICs
*/
type Nics struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  MAC address of the NIC
	*/
	MacAddress []string `json:"macAddress,omitempty"`
}

func (p *Nics) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias Nics

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
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

func (p *Nics) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Nics
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewNics()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.MacAddress != nil {
		p.MacAddress = known.MacAddress
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "macAddress")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewNics() *Nics {
	p := new(Nics)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.Nics"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Details of a node managed by Foundation Central, including hardware, network, and lifecycle state
*/
type Node struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Version of the AOS installed on the node
	*/
	AosVersion *string `json:"aosVersion,omitempty"`
	/*
	  Block serial number of the node
	*/
	BlockSerialNumber *string `json:"blockSerialNumber,omitempty"`

	CpuInfo *CpuInfo `json:"cpuInfo,omitempty"`
	/*
	  Time when the node was discovered in Foundation Central
	*/
	CreatedTime *time.Time `json:"createdTime,omitempty"`
	/*
	  User tags of the node
	*/
	CustomAttributes []string `json:"customAttributes,omitempty"`

	CvmConnectivityStatus *NodeConnectivityStatus `json:"cvmConnectivityStatus,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`

	HostConnectivityStatus *NodeConnectivityStatus `json:"hostConnectivityStatus,omitempty"`

	HostType *HostType `json:"hostType,omitempty"`
	/*
	  Version of the host installed on the node
	*/
	HostVersion *string `json:"hostVersion,omitempty"`
	/*
	  Hostname of the node
	*/
	Hostname *string `json:"hostname,omitempty"`
	/*
	  Unique identifier array for the node (currently limited to 1 item with NODE_SERIAL type)
	*/
	Identifiers []NodeIdentifier `json:"identifiers"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Manufacturer of the node
	*/
	Manufacturer *string `json:"manufacturer"`
	/*
	  Total memory available in the node (in GB)
	*/
	MemoryGB *int `json:"memoryGB,omitempty"`
	/*
	  Model of the node
	*/
	Model *string `json:"model,omitempty"`

	NetworkDetails *NodeNetworkDetails `json:"networkDetails,omitempty"`
	/*
	  External ID of the owner
	*/
	OwnerExtId *string `json:"ownerExtId,omitempty"`
	/*
	  External ID of the provider connection of the node
	*/
	ProviderConnectionExtId *string `json:"providerConnectionExtId,omitempty"`

	ProviderData *ExtendedNodeProviderData `json:"providerData,omitempty"`
	/*
	  External ID of the provider of the node
	*/
	ProviderExtId *string `json:"providerExtId,omitempty"`

	State *NodeState `json:"state,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *Node) MarshalJSON() ([]byte, error) {
	type NodeProxy Node

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*NodeProxy
		Identifiers  []NodeIdentifier `json:"identifiers,omitempty"`
		Manufacturer *string          `json:"manufacturer,omitempty"`
	}{
		NodeProxy:    (*NodeProxy)(p),
		Identifiers:  p.Identifiers,
		Manufacturer: p.Manufacturer,
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

func (p *Node) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Node
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewNode()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.AosVersion != nil {
		p.AosVersion = known.AosVersion
	}
	if known.BlockSerialNumber != nil {
		p.BlockSerialNumber = known.BlockSerialNumber
	}
	if known.CpuInfo != nil {
		p.CpuInfo = known.CpuInfo
	}
	if known.CreatedTime != nil {
		p.CreatedTime = known.CreatedTime
	}
	if known.CustomAttributes != nil {
		p.CustomAttributes = known.CustomAttributes
	}
	if known.CvmConnectivityStatus != nil {
		p.CvmConnectivityStatus = known.CvmConnectivityStatus
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.HostConnectivityStatus != nil {
		p.HostConnectivityStatus = known.HostConnectivityStatus
	}
	if known.HostType != nil {
		p.HostType = known.HostType
	}
	if known.HostVersion != nil {
		p.HostVersion = known.HostVersion
	}
	if known.Hostname != nil {
		p.Hostname = known.Hostname
	}
	if known.Identifiers != nil {
		p.Identifiers = known.Identifiers
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.Manufacturer != nil {
		p.Manufacturer = known.Manufacturer
	}
	if known.MemoryGB != nil {
		p.MemoryGB = known.MemoryGB
	}
	if known.Model != nil {
		p.Model = known.Model
	}
	if known.NetworkDetails != nil {
		p.NetworkDetails = known.NetworkDetails
	}
	if known.OwnerExtId != nil {
		p.OwnerExtId = known.OwnerExtId
	}
	if known.ProviderConnectionExtId != nil {
		p.ProviderConnectionExtId = known.ProviderConnectionExtId
	}
	if known.ProviderData != nil {
		p.ProviderData = known.ProviderData
	}
	if known.ProviderExtId != nil {
		p.ProviderExtId = known.ProviderExtId
	}
	if known.State != nil {
		p.State = known.State
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "aosVersion")
	delete(allFields, "blockSerialNumber")
	delete(allFields, "cpuInfo")
	delete(allFields, "createdTime")
	delete(allFields, "customAttributes")
	delete(allFields, "cvmConnectivityStatus")
	delete(allFields, "extId")
	delete(allFields, "hostConnectivityStatus")
	delete(allFields, "hostType")
	delete(allFields, "hostVersion")
	delete(allFields, "hostname")
	delete(allFields, "identifiers")
	delete(allFields, "links")
	delete(allFields, "manufacturer")
	delete(allFields, "memoryGB")
	delete(allFields, "model")
	delete(allFields, "networkDetails")
	delete(allFields, "ownerExtId")
	delete(allFields, "providerConnectionExtId")
	delete(allFields, "providerData")
	delete(allFields, "providerExtId")
	delete(allFields, "state")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewNode() *Node {
	p := new(Node)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.Node"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Configuration of the node imaged using the patched image
*/
type NodeConfiguration struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	HostConfiguration *HostConfiguration `json:"hostConfiguration,omitempty"`
	/*
	  External ID of the node
	*/
	NodeExtId *string `json:"nodeExtId"`
}

func (p *NodeConfiguration) MarshalJSON() ([]byte, error) {
	type NodeConfigurationProxy NodeConfiguration

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*NodeConfigurationProxy
		NodeExtId *string `json:"nodeExtId,omitempty"`
	}{
		NodeConfigurationProxy: (*NodeConfigurationProxy)(p),
		NodeExtId:              p.NodeExtId,
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

func (p *NodeConfiguration) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias NodeConfiguration
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewNodeConfiguration()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.HostConfiguration != nil {
		p.HostConfiguration = known.HostConfiguration
	}
	if known.NodeExtId != nil {
		p.NodeExtId = known.NodeExtId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "hostConfiguration")
	delete(allFields, "nodeExtId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewNodeConfiguration() *NodeConfiguration {
	p := new(NodeConfiguration)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.NodeConfiguration"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Connectivity status of the node
*/
type NodeConnectivityStatus int

const (
	NODECONNECTIVITYSTATUS_UNKNOWN             NodeConnectivityStatus = 0
	NODECONNECTIVITYSTATUS_REDACTED            NodeConnectivityStatus = 1
	NODECONNECTIVITYSTATUS_CONNECTED           NodeConnectivityStatus = 2
	NODECONNECTIVITYSTATUS_CONNECTING          NodeConnectivityStatus = 3
	NODECONNECTIVITYSTATUS_DISCONNECTED        NodeConnectivityStatus = 4
	NODECONNECTIVITYSTATUS_CONNECTION_DEGRADED NodeConnectivityStatus = 5
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *NodeConnectivityStatus) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CONNECTED",
		"CONNECTING",
		"DISCONNECTED",
		"CONNECTION_DEGRADED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e NodeConnectivityStatus) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CONNECTED",
		"CONNECTING",
		"DISCONNECTED",
		"CONNECTION_DEGRADED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *NodeConnectivityStatus) index(name string) NodeConnectivityStatus {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CONNECTED",
		"CONNECTING",
		"DISCONNECTED",
		"CONNECTION_DEGRADED",
	}
	for idx := range names {
		if names[idx] == name {
			return NodeConnectivityStatus(idx)
		}
	}
	return NODECONNECTIVITYSTATUS_UNKNOWN
}

func (e *NodeConnectivityStatus) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for NodeConnectivityStatus:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *NodeConnectivityStatus) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e NodeConnectivityStatus) Ref() *NodeConnectivityStatus {
	return &e
}

/*
Node unique identifier object containing type and value
*/
type NodeIdentifier struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Type *NodeIdentifierType `json:"type"`
	/*
	  Unique identifier value of the node
	*/
	Value *string `json:"value"`
}

func (p *NodeIdentifier) MarshalJSON() ([]byte, error) {
	type NodeIdentifierProxy NodeIdentifier

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*NodeIdentifierProxy
		Type  *NodeIdentifierType `json:"type,omitempty"`
		Value *string             `json:"value,omitempty"`
	}{
		NodeIdentifierProxy: (*NodeIdentifierProxy)(p),
		Type:                p.Type,
		Value:               p.Value,
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

func (p *NodeIdentifier) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias NodeIdentifier
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewNodeIdentifier()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Type != nil {
		p.Type = known.Type
	}
	if known.Value != nil {
		p.Value = known.Value
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "type")
	delete(allFields, "value")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewNodeIdentifier() *NodeIdentifier {
	p := new(NodeIdentifier)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.NodeIdentifier"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Type of the unique identifier of the node
*/
type NodeIdentifierType int

const (
	NODEIDENTIFIERTYPE_UNKNOWN       NodeIdentifierType = 0
	NODEIDENTIFIERTYPE_REDACTED      NodeIdentifierType = 1
	NODEIDENTIFIERTYPE_SERIAL_NUMBER NodeIdentifierType = 2
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *NodeIdentifierType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SERIAL_NUMBER",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e NodeIdentifierType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SERIAL_NUMBER",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *NodeIdentifierType) index(name string) NodeIdentifierType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SERIAL_NUMBER",
	}
	for idx := range names {
		if names[idx] == name {
			return NodeIdentifierType(idx)
		}
	}
	return NODEIDENTIFIERTYPE_UNKNOWN
}

func (e *NodeIdentifierType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for NodeIdentifierType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *NodeIdentifierType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e NodeIdentifierType) Ref() *NodeIdentifierType {
	return &e
}

/*
Reference to a node by its external ID
*/
type NodeInfo struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  External ID of the node
	*/
	ExtId *string `json:"extId"`
}

func (p *NodeInfo) MarshalJSON() ([]byte, error) {
	type NodeInfoProxy NodeInfo

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*NodeInfoProxy
		ExtId *string `json:"extId,omitempty"`
	}{
		NodeInfoProxy: (*NodeInfoProxy)(p),
		ExtId:         p.ExtId,
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

func (p *NodeInfo) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias NodeInfo
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewNodeInfo()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "extId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewNodeInfo() *NodeInfo {
	p := new(NodeInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.NodeInfo"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Network details of the host and CVM present on the node, if installed
*/
type NodeNetworkDetails struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Bmc *NetworkDetails `json:"bmc,omitempty"`

	Cvm *ComponentNetworkDetails `json:"cvm,omitempty"`

	Host *ComponentNetworkDetails `json:"host,omitempty"`
}

func (p *NodeNetworkDetails) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias NodeNetworkDetails

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
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

func (p *NodeNetworkDetails) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias NodeNetworkDetails
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewNodeNetworkDetails()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Bmc != nil {
		p.Bmc = known.Bmc
	}
	if known.Cvm != nil {
		p.Cvm = known.Cvm
	}
	if known.Host != nil {
		p.Host = known.Host
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "bmc")
	delete(allFields, "cvm")
	delete(allFields, "host")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewNodeNetworkDetails() *NodeNetworkDetails {
	p := new(NodeNetworkDetails)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.NodeNetworkDetails"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type NodeProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Version of the AOS installed on the node
	*/
	AosVersion *string `json:"aosVersion,omitempty"`
	/*
	  Block serial number of the node
	*/
	BlockSerialNumber *string `json:"blockSerialNumber,omitempty"`

	ConnectionProjection *ConnectionProjection `json:"connectionProjection,omitempty"`

	CpuInfo *CpuInfo `json:"cpuInfo,omitempty"`
	/*
	  Time when the node was discovered in Foundation Central
	*/
	CreatedTime *time.Time `json:"createdTime,omitempty"`
	/*
	  User tags of the node
	*/
	CustomAttributes []string `json:"customAttributes,omitempty"`

	CvmConnectivityStatus *NodeConnectivityStatus `json:"cvmConnectivityStatus,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`

	HostConnectivityStatus *NodeConnectivityStatus `json:"hostConnectivityStatus,omitempty"`

	HostType *HostType `json:"hostType,omitempty"`
	/*
	  Version of the host installed on the node
	*/
	HostVersion *string `json:"hostVersion,omitempty"`
	/*
	  Hostname of the node
	*/
	Hostname *string `json:"hostname,omitempty"`
	/*
	  Unique identifier array for the node (currently limited to 1 item with NODE_SERIAL type)
	*/
	Identifiers []NodeIdentifier `json:"identifiers"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Manufacturer of the node
	*/
	Manufacturer *string `json:"manufacturer"`
	/*
	  Total memory available in the node (in GB)
	*/
	MemoryGB *int `json:"memoryGB,omitempty"`
	/*
	  Model of the node
	*/
	Model *string `json:"model,omitempty"`

	NetworkDetails *NodeNetworkDetails `json:"networkDetails,omitempty"`
	/*
	  External ID of the owner
	*/
	OwnerExtId *string `json:"ownerExtId,omitempty"`
	/*
	  External ID of the provider connection of the node
	*/
	ProviderConnectionExtId *string `json:"providerConnectionExtId,omitempty"`

	ProviderData *ExtendedNodeProviderData `json:"providerData,omitempty"`
	/*
	  External ID of the provider of the node
	*/
	ProviderExtId *string `json:"providerExtId,omitempty"`

	State *NodeState `json:"state,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *NodeProjection) MarshalJSON() ([]byte, error) {
	type NodeProjectionProxy NodeProjection

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*NodeProjectionProxy
		Identifiers  []NodeIdentifier `json:"identifiers,omitempty"`
		Manufacturer *string          `json:"manufacturer,omitempty"`
	}{
		NodeProjectionProxy: (*NodeProjectionProxy)(p),
		Identifiers:         p.Identifiers,
		Manufacturer:        p.Manufacturer,
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

func (p *NodeProjection) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias NodeProjection
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewNodeProjection()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.AosVersion != nil {
		p.AosVersion = known.AosVersion
	}
	if known.BlockSerialNumber != nil {
		p.BlockSerialNumber = known.BlockSerialNumber
	}
	if known.ConnectionProjection != nil {
		p.ConnectionProjection = known.ConnectionProjection
	}
	if known.CpuInfo != nil {
		p.CpuInfo = known.CpuInfo
	}
	if known.CreatedTime != nil {
		p.CreatedTime = known.CreatedTime
	}
	if known.CustomAttributes != nil {
		p.CustomAttributes = known.CustomAttributes
	}
	if known.CvmConnectivityStatus != nil {
		p.CvmConnectivityStatus = known.CvmConnectivityStatus
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.HostConnectivityStatus != nil {
		p.HostConnectivityStatus = known.HostConnectivityStatus
	}
	if known.HostType != nil {
		p.HostType = known.HostType
	}
	if known.HostVersion != nil {
		p.HostVersion = known.HostVersion
	}
	if known.Hostname != nil {
		p.Hostname = known.Hostname
	}
	if known.Identifiers != nil {
		p.Identifiers = known.Identifiers
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.Manufacturer != nil {
		p.Manufacturer = known.Manufacturer
	}
	if known.MemoryGB != nil {
		p.MemoryGB = known.MemoryGB
	}
	if known.Model != nil {
		p.Model = known.Model
	}
	if known.NetworkDetails != nil {
		p.NetworkDetails = known.NetworkDetails
	}
	if known.OwnerExtId != nil {
		p.OwnerExtId = known.OwnerExtId
	}
	if known.ProviderConnectionExtId != nil {
		p.ProviderConnectionExtId = known.ProviderConnectionExtId
	}
	if known.ProviderData != nil {
		p.ProviderData = known.ProviderData
	}
	if known.ProviderExtId != nil {
		p.ProviderExtId = known.ProviderExtId
	}
	if known.State != nil {
		p.State = known.State
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "aosVersion")
	delete(allFields, "blockSerialNumber")
	delete(allFields, "connectionProjection")
	delete(allFields, "cpuInfo")
	delete(allFields, "createdTime")
	delete(allFields, "customAttributes")
	delete(allFields, "cvmConnectivityStatus")
	delete(allFields, "extId")
	delete(allFields, "hostConnectivityStatus")
	delete(allFields, "hostType")
	delete(allFields, "hostVersion")
	delete(allFields, "hostname")
	delete(allFields, "identifiers")
	delete(allFields, "links")
	delete(allFields, "manufacturer")
	delete(allFields, "memoryGB")
	delete(allFields, "model")
	delete(allFields, "networkDetails")
	delete(allFields, "ownerExtId")
	delete(allFields, "providerConnectionExtId")
	delete(allFields, "providerData")
	delete(allFields, "providerExtId")
	delete(allFields, "state")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewNodeProjection() *NodeProjection {
	p := new(NodeProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.NodeProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Provider-specific data for the node
*/
type NodeProviderData struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	DisplayNameMapping *DisplayNameMapping `json:"displayNameMapping,omitempty"`
	/*
	  Domain of the discovered node
	*/
	Domain *string `json:"domain,omitempty"`
	/*
	  Groups associated with the discovered node
	*/
	Groups []Group `json:"groups,omitempty"`
	/*
	  Indicates if the node is available
	*/
	IsAvailable *bool `json:"isAvailable,omitempty"`
	/*
	  Indicates if the node is configured
	*/
	IsConfigured *bool `json:"isConfigured,omitempty"`
	/*
	  Mode of the discovered node
	*/
	Mode *string `json:"mode,omitempty"`
	/*
	  Name of the discovered node
	*/
	Name *string `json:"name,omitempty"`
	/*
	  Tags associated with the discovered node
	*/
	Tags []import1.KVStringPair `json:"tags,omitempty"`
}

func (p *NodeProviderData) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias NodeProviderData

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
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

func (p *NodeProviderData) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias NodeProviderData
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewNodeProviderData()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DisplayNameMapping != nil {
		p.DisplayNameMapping = known.DisplayNameMapping
	}
	if known.Domain != nil {
		p.Domain = known.Domain
	}
	if known.Groups != nil {
		p.Groups = known.Groups
	}
	if known.IsAvailable != nil {
		p.IsAvailable = known.IsAvailable
	}
	if known.IsConfigured != nil {
		p.IsConfigured = known.IsConfigured
	}
	if known.Mode != nil {
		p.Mode = known.Mode
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.Tags != nil {
		p.Tags = known.Tags
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "displayNameMapping")
	delete(allFields, "domain")
	delete(allFields, "groups")
	delete(allFields, "isAvailable")
	delete(allFields, "isConfigured")
	delete(allFields, "mode")
	delete(allFields, "name")
	delete(allFields, "tags")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewNodeProviderData() *NodeProviderData {
	p := new(NodeProviderData)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.NodeProviderData"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
State of the node with respect to its lifecycle in Foundation Central workflows
*/
type NodeState int

const (
	NODESTATE_UNKNOWN                       NodeState = 0
	NODESTATE_REDACTED                      NodeState = 1
	NODESTATE_DISCOVERED                    NodeState = 2
	NODESTATE_DISCOVERY_FAILED              NodeState = 3
	NODESTATE_ONBOARDED                     NodeState = 4
	NODESTATE_HARDWARE_DISCOVERY            NodeState = 5
	NODESTATE_CONFIGURING_HARDWARE_PROVIDER NodeState = 6
	NODESTATE_CONFIGURED_HARDWARE_PROVIDER  NodeState = 7
	NODESTATE_INSTALLING_HOST_OS            NodeState = 8
	NODESTATE_INSTALLED_HOST_OS             NodeState = 9
	NODESTATE_INSTALLING_AOS                NodeState = 10
	NODESTATE_INSTALLED_AOS                 NodeState = 11
	NODESTATE_CLUSTERING                    NodeState = 12
	NODESTATE_CLUSTERED                     NodeState = 13
	NODESTATE_PREPARING                     NodeState = 14
	NODESTATE_PREPARED                      NodeState = 15
	NODESTATE_FAILED                        NodeState = 16
	NODESTATE_IMAGING                       NodeState = 17
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *NodeState) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"DISCOVERED",
		"DISCOVERY_FAILED",
		"ONBOARDED",
		"HARDWARE_DISCOVERY",
		"CONFIGURING_HARDWARE_PROVIDER",
		"CONFIGURED_HARDWARE_PROVIDER",
		"INSTALLING_HOST_OS",
		"INSTALLED_HOST_OS",
		"INSTALLING_AOS",
		"INSTALLED_AOS",
		"CLUSTERING",
		"CLUSTERED",
		"PREPARING",
		"PREPARED",
		"FAILED",
		"IMAGING",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e NodeState) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"DISCOVERED",
		"DISCOVERY_FAILED",
		"ONBOARDED",
		"HARDWARE_DISCOVERY",
		"CONFIGURING_HARDWARE_PROVIDER",
		"CONFIGURED_HARDWARE_PROVIDER",
		"INSTALLING_HOST_OS",
		"INSTALLED_HOST_OS",
		"INSTALLING_AOS",
		"INSTALLED_AOS",
		"CLUSTERING",
		"CLUSTERED",
		"PREPARING",
		"PREPARED",
		"FAILED",
		"IMAGING",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *NodeState) index(name string) NodeState {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"DISCOVERED",
		"DISCOVERY_FAILED",
		"ONBOARDED",
		"HARDWARE_DISCOVERY",
		"CONFIGURING_HARDWARE_PROVIDER",
		"CONFIGURED_HARDWARE_PROVIDER",
		"INSTALLING_HOST_OS",
		"INSTALLED_HOST_OS",
		"INSTALLING_AOS",
		"INSTALLED_AOS",
		"CLUSTERING",
		"CLUSTERED",
		"PREPARING",
		"PREPARED",
		"FAILED",
		"IMAGING",
	}
	for idx := range names {
		if names[idx] == name {
			return NodeState(idx)
		}
	}
	return NODESTATE_UNKNOWN
}

func (e *NodeState) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for NodeState:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *NodeState) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e NodeState) Ref() *NodeState {
	return &e
}

/*
A customized hypervisor image combining a base image with patches for specific node deployments
*/
type PatchedImage struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  ExtId of the claim token used in the patched image
	*/
	ClaimTokenExtId *string `json:"claimTokenExtId"`
	/*
	  Timestamp when the patched image was created
	*/
	CreatedTime *time.Time `json:"createdTime,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`

	HostType *HostType `json:"hostType"`
	/*

	 */
	ImageDetailsItemDiscriminator_ *string `json:"$imageDetailsItemDiscriminator,omitempty"`
	/*
	  Details of the image used for patching
	*/
	ImageDetails *OneOfPatchedImageImageDetails `json:"imageDetails"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Patched image name
	*/
	Name *string `json:"name"`
	/*
	  List of node configurations used for patching the image
	*/
	NodeConfigurations []NodeConfiguration `json:"nodeConfigurations"`
	/*
	  External ID of the owner
	*/
	OwnerExtId *string `json:"ownerExtId,omitempty"`
	/*
	  SHA-256 checksum of the patched image
	*/
	PatchedIsoSha256Checksum *string `json:"patchedIsoSha256Checksum"`
	/*
	  URL for downloading the patched image
	*/
	PatchedIsoUrl *string `json:"patchedIsoUrl"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  Version of the base hypervisor or host OS image used for patching
	*/
	Version *string `json:"version,omitempty"`
}

func (p *PatchedImage) MarshalJSON() ([]byte, error) {
	type PatchedImageProxy PatchedImage

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*PatchedImageProxy
		ClaimTokenExtId          *string                        `json:"claimTokenExtId,omitempty"`
		HostType                 *HostType                      `json:"hostType,omitempty"`
		ImageDetails             *OneOfPatchedImageImageDetails `json:"imageDetails,omitempty"`
		Name                     *string                        `json:"name,omitempty"`
		NodeConfigurations       []NodeConfiguration            `json:"nodeConfigurations,omitempty"`
		PatchedIsoSha256Checksum *string                        `json:"patchedIsoSha256Checksum,omitempty"`
		PatchedIsoUrl            *string                        `json:"patchedIsoUrl,omitempty"`
	}{
		PatchedImageProxy:        (*PatchedImageProxy)(p),
		ClaimTokenExtId:          p.ClaimTokenExtId,
		HostType:                 p.HostType,
		ImageDetails:             p.ImageDetails,
		Name:                     p.Name,
		NodeConfigurations:       p.NodeConfigurations,
		PatchedIsoSha256Checksum: p.PatchedIsoSha256Checksum,
		PatchedIsoUrl:            p.PatchedIsoUrl,
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

func (p *PatchedImage) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias PatchedImage
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewPatchedImage()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ClaimTokenExtId != nil {
		p.ClaimTokenExtId = known.ClaimTokenExtId
	}
	if known.CreatedTime != nil {
		p.CreatedTime = known.CreatedTime
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.HostType != nil {
		p.HostType = known.HostType
	}
	if known.ImageDetailsItemDiscriminator_ != nil {
		p.ImageDetailsItemDiscriminator_ = known.ImageDetailsItemDiscriminator_
	}
	if known.ImageDetails != nil {
		p.ImageDetails = known.ImageDetails
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.NodeConfigurations != nil {
		p.NodeConfigurations = known.NodeConfigurations
	}
	if known.OwnerExtId != nil {
		p.OwnerExtId = known.OwnerExtId
	}
	if known.PatchedIsoSha256Checksum != nil {
		p.PatchedIsoSha256Checksum = known.PatchedIsoSha256Checksum
	}
	if known.PatchedIsoUrl != nil {
		p.PatchedIsoUrl = known.PatchedIsoUrl
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}
	if known.Version != nil {
		p.Version = known.Version
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "claimTokenExtId")
	delete(allFields, "createdTime")
	delete(allFields, "extId")
	delete(allFields, "hostType")
	delete(allFields, "$imageDetailsItemDiscriminator")
	delete(allFields, "imageDetails")
	delete(allFields, "links")
	delete(allFields, "name")
	delete(allFields, "nodeConfigurations")
	delete(allFields, "ownerExtId")
	delete(allFields, "patchedIsoSha256Checksum")
	delete(allFields, "patchedIsoUrl")
	delete(allFields, "tenantId")
	delete(allFields, "version")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewPatchedImage() *PatchedImage {
	p := new(PatchedImage)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.PatchedImage"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *PatchedImage) GetImageDetails() interface{} {
	if nil == p.ImageDetails {
		return nil
	}
	return p.ImageDetails.GetValue()
}

func (p *PatchedImage) SetImageDetails(v interface{}) error {
	if nil == p.ImageDetails {
		p.ImageDetails = NewOneOfPatchedImageImageDetails()
	}
	e := p.ImageDetails.SetValue(v)
	if nil == e {
		if nil == p.ImageDetailsItemDiscriminator_ {
			p.ImageDetailsItemDiscriminator_ = new(string)
		}
		*p.ImageDetailsItemDiscriminator_ = *p.ImageDetails.Discriminator
	}
	return e
}

/*
Details of a physical network adapter of a server
*/
type PhysicalAdapter struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Vendor and model information for the physical network adapter, identifying the exact NIC hardware installed
	*/
	Model *string `json:"model,omitempty"`
	/*
	  Name of the physical network adapter
	*/
	Name *string `json:"name,omitempty"`
	/*
	  Count of ports on the physical network adapter
	*/
	PortCount *int `json:"portCount,omitempty"`
	/*
	  List of physical ports on the network adapter
	*/
	Ports []PhysicalNetworkAdapterPort `json:"ports,omitempty"`
	/*
	  Slot identifier of the physical network adapter. Empty value indicates that the slot could not be determined
	*/
	Slot *string `json:"slot,omitempty"`
}

func (p *PhysicalAdapter) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias PhysicalAdapter

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
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

func (p *PhysicalAdapter) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias PhysicalAdapter
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewPhysicalAdapter()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Model != nil {
		p.Model = known.Model
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.PortCount != nil {
		p.PortCount = known.PortCount
	}
	if known.Ports != nil {
		p.Ports = known.Ports
	}
	if known.Slot != nil {
		p.Slot = known.Slot
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "model")
	delete(allFields, "name")
	delete(allFields, "portCount")
	delete(allFields, "ports")
	delete(allFields, "slot")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewPhysicalAdapter() *PhysicalAdapter {
	p := new(PhysicalAdapter)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.PhysicalAdapter"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Configuration for physical network adapters with network groups and port mappings
*/
type PhysicalNetworkAdapterConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of network groups defining network segments with their associated VLANs and port mappings
	*/
	NetworkGroups []NetworkGroup `json:"networkGroups,omitempty"`
}

func (p *PhysicalNetworkAdapterConfig) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias PhysicalNetworkAdapterConfig

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
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

func (p *PhysicalNetworkAdapterConfig) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias PhysicalNetworkAdapterConfig
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewPhysicalNetworkAdapterConfig()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.NetworkGroups != nil {
		p.NetworkGroups = known.NetworkGroups
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "networkGroups")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewPhysicalNetworkAdapterConfig() *PhysicalNetworkAdapterConfig {
	p := new(PhysicalNetworkAdapterConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.PhysicalNetworkAdapterConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Physical network adapter port of the discovered node
*/
type PhysicalNetworkAdapterPort struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  MAC address associated with this physical port
	*/
	MacAddress *string `json:"macAddress"`
	/*
	  Name of the physical port
	*/
	Name *string `json:"name"`
	/*
	  Position of the port on the physical network adapter
	*/
	Position *int `json:"position"`
	/*
	  Link speed of the physical port. 0 value indicates that the speed could not be determined
	*/
	SpeedGbps *float32 `json:"speedGbps"`

	Status *PhysicalNetworkAdaptorPortStatus `json:"status"`
}

func (p *PhysicalNetworkAdapterPort) MarshalJSON() ([]byte, error) {
	type PhysicalNetworkAdapterPortProxy PhysicalNetworkAdapterPort

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*PhysicalNetworkAdapterPortProxy
		MacAddress *string                           `json:"macAddress,omitempty"`
		Name       *string                           `json:"name,omitempty"`
		Position   *int                              `json:"position,omitempty"`
		SpeedGbps  *float32                          `json:"speedGbps,omitempty"`
		Status     *PhysicalNetworkAdaptorPortStatus `json:"status,omitempty"`
	}{
		PhysicalNetworkAdapterPortProxy: (*PhysicalNetworkAdapterPortProxy)(p),
		MacAddress:                      p.MacAddress,
		Name:                            p.Name,
		Position:                        p.Position,
		SpeedGbps:                       p.SpeedGbps,
		Status:                          p.Status,
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

func (p *PhysicalNetworkAdapterPort) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias PhysicalNetworkAdapterPort
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewPhysicalNetworkAdapterPort()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.MacAddress != nil {
		p.MacAddress = known.MacAddress
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.Position != nil {
		p.Position = known.Position
	}
	if known.SpeedGbps != nil {
		p.SpeedGbps = known.SpeedGbps
	}
	if known.Status != nil {
		p.Status = known.Status
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "macAddress")
	delete(allFields, "name")
	delete(allFields, "position")
	delete(allFields, "speedGbps")
	delete(allFields, "status")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewPhysicalNetworkAdapterPort() *PhysicalNetworkAdapterPort {
	p := new(PhysicalNetworkAdapterPort)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.PhysicalNetworkAdapterPort"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Physical network adapter of the discovered node
*/
type PhysicalNetworkAdapters struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of physical adapters of a server
	*/
	Adapters []PhysicalAdapter `json:"adapters,omitempty"`
	/*
	  List of recommended network groups for the node
	*/
	RecommendedSettings []RecommendedNetworkGroup `json:"recommendedSettings,omitempty"`
}

func (p *PhysicalNetworkAdapters) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias PhysicalNetworkAdapters

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
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

func (p *PhysicalNetworkAdapters) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias PhysicalNetworkAdapters
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewPhysicalNetworkAdapters()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Adapters != nil {
		p.Adapters = known.Adapters
	}
	if known.RecommendedSettings != nil {
		p.RecommendedSettings = known.RecommendedSettings
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "adapters")
	delete(allFields, "recommendedSettings")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewPhysicalNetworkAdapters() *PhysicalNetworkAdapters {
	p := new(PhysicalNetworkAdapters)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.PhysicalNetworkAdapters"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Current operational link state of the physical port
*/
type PhysicalNetworkAdaptorPortStatus int

const (
	PHYSICALNETWORKADAPTORPORTSTATUS_UNKNOWN  PhysicalNetworkAdaptorPortStatus = 0
	PHYSICALNETWORKADAPTORPORTSTATUS_REDACTED PhysicalNetworkAdaptorPortStatus = 1
	PHYSICALNETWORKADAPTORPORTSTATUS_UP       PhysicalNetworkAdaptorPortStatus = 2
	PHYSICALNETWORKADAPTORPORTSTATUS_DOWN     PhysicalNetworkAdaptorPortStatus = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *PhysicalNetworkAdaptorPortStatus) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"UP",
		"DOWN",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e PhysicalNetworkAdaptorPortStatus) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"UP",
		"DOWN",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *PhysicalNetworkAdaptorPortStatus) index(name string) PhysicalNetworkAdaptorPortStatus {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"UP",
		"DOWN",
	}
	for idx := range names {
		if names[idx] == name {
			return PhysicalNetworkAdaptorPortStatus(idx)
		}
	}
	return PHYSICALNETWORKADAPTORPORTSTATUS_UNKNOWN
}

func (e *PhysicalNetworkAdaptorPortStatus) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for PhysicalNetworkAdaptorPortStatus:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *PhysicalNetworkAdaptorPortStatus) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e PhysicalNetworkAdaptorPortStatus) Ref() *PhysicalNetworkAdaptorPortStatus {
	return &e
}

/*
Group reference for resource pools containing only the group ID
*/
type PoolGroup struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  ID of the group
	*/
	Id *string `json:"id,omitempty"`
}

func (p *PoolGroup) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias PoolGroup

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
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

func (p *PoolGroup) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias PoolGroup
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewPoolGroup()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Id != nil {
		p.Id = known.Id
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "id")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewPoolGroup() *PoolGroup {
	p := new(PoolGroup)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.PoolGroup"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Physical port mapping configuration identifying network interfaces by MAC address
*/
type PortMapping struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  MAC address of the physical network interface to be mapped to this network group
	*/
	MacAddress *string `json:"macAddress,omitempty"`
}

func (p *PortMapping) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias PortMapping

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
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

func (p *PortMapping) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias PortMapping
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewPortMapping()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.MacAddress != nil {
		p.MacAddress = known.MacAddress
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "macAddress")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewPortMapping() *PortMapping {
	p := new(PortMapping)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.PortMapping"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Configuration details for post-cluster operations on nodes after cluster creation
*/
type PostClusterConfigDetails struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  External ID of the cluster
	*/
	ClusterExtId *string `json:"clusterExtId"`
	/*
	  List of nodes to perform post-cluster operations on
	*/
	Nodes []NodeInfo `json:"nodes"`
}

func (p *PostClusterConfigDetails) MarshalJSON() ([]byte, error) {
	type PostClusterConfigDetailsProxy PostClusterConfigDetails

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*PostClusterConfigDetailsProxy
		ClusterExtId *string    `json:"clusterExtId,omitempty"`
		Nodes        []NodeInfo `json:"nodes,omitempty"`
	}{
		PostClusterConfigDetailsProxy: (*PostClusterConfigDetailsProxy)(p),
		ClusterExtId:                  p.ClusterExtId,
		Nodes:                         p.Nodes,
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

func (p *PostClusterConfigDetails) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias PostClusterConfigDetails
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewPostClusterConfigDetails()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ClusterExtId != nil {
		p.ClusterExtId = known.ClusterExtId
	}
	if known.Nodes != nil {
		p.Nodes = known.Nodes
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "clusterExtId")
	delete(allFields, "nodes")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewPostClusterConfigDetails() *PostClusterConfigDetails {
	p := new(PostClusterConfigDetails)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.PostClusterConfigDetails"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Configuration details for pre-cluster operations on nodes before cluster creation or expansion
*/
type PreClusterConfigDetails struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of nodes to perform pre-cluster operations on
	*/
	Nodes []NodeInfo `json:"nodes"`
}

func (p *PreClusterConfigDetails) MarshalJSON() ([]byte, error) {
	type PreClusterConfigDetailsProxy PreClusterConfigDetails

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*PreClusterConfigDetailsProxy
		Nodes []NodeInfo `json:"nodes,omitempty"`
	}{
		PreClusterConfigDetailsProxy: (*PreClusterConfigDetailsProxy)(p),
		Nodes:                        p.Nodes,
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

func (p *PreClusterConfigDetails) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias PreClusterConfigDetails
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewPreClusterConfigDetails()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Nodes != nil {
		p.Nodes = known.Nodes
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "nodes")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewPreClusterConfigDetails() *PreClusterConfigDetails {
	p := new(PreClusterConfigDetails)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.PreClusterConfigDetails"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Recommended network group configuration for the node
*/
type RecommendedNetworkGroup struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Name of the network group
	*/
	Name *string `json:"name"`
	/*
	  List of physical port mappings by MAC address
	*/
	PortMappings []PortMapping `json:"portMappings,omitempty"`

	Type *NetworkGroupType `json:"type"`
}

func (p *RecommendedNetworkGroup) MarshalJSON() ([]byte, error) {
	type RecommendedNetworkGroupProxy RecommendedNetworkGroup

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*RecommendedNetworkGroupProxy
		Name *string           `json:"name,omitempty"`
		Type *NetworkGroupType `json:"type,omitempty"`
	}{
		RecommendedNetworkGroupProxy: (*RecommendedNetworkGroupProxy)(p),
		Name:                         p.Name,
		Type:                         p.Type,
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

func (p *RecommendedNetworkGroup) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias RecommendedNetworkGroup
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewRecommendedNetworkGroup()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.PortMappings != nil {
		p.PortMappings = known.PortMappings
	}
	if known.Type != nil {
		p.Type = known.Type
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "name")
	delete(allFields, "portMappings")
	delete(allFields, "type")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewRecommendedNetworkGroup() *RecommendedNetworkGroup {
	p := new(RecommendedNetworkGroup)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.RecommendedNetworkGroup"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Configuration for updating network settings, hostname, DNS, and NTP on nodes with existing installations without reimaging
*/
type ReconfigurationDetails struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	CvmDetails *CvmReconfigurationDetails `json:"cvmDetails,omitempty"`

	HostDetails *HostReconfigurationDetails `json:"hostDetails,omitempty"`
}

func (p *ReconfigurationDetails) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ReconfigurationDetails

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
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

func (p *ReconfigurationDetails) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ReconfigurationDetails
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewReconfigurationDetails()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.CvmDetails != nil {
		p.CvmDetails = known.CvmDetails
	}
	if known.HostDetails != nil {
		p.HostDetails = known.HostDetails
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "cvmDetails")
	delete(allFields, "hostDetails")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewReconfigurationDetails() *ReconfigurationDetails {
	p := new(ReconfigurationDetails)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.ReconfigurationDetails"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /lifecycle/v4.3/config/hardware-providers/{hardwareProviderExtId}/connections/{extId}/$actions/refresh Post operation
*/
type RefreshConnectionApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfRefreshConnectionApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *RefreshConnectionApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias RefreshConnectionApiResponse

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
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

func (p *RefreshConnectionApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias RefreshConnectionApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewRefreshConnectionApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewRefreshConnectionApiResponse() *RefreshConnectionApiResponse {
	p := new(RefreshConnectionApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.RefreshConnectionApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *RefreshConnectionApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *RefreshConnectionApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfRefreshConnectionApiResponseData()
	}
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
Request to refresh nodes or resources from a hardware provider connection
*/
type RefreshConnectionSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	ConfigurationItemDiscriminator_ *string `json:"$configurationItemDiscriminator,omitempty"`
	/*
	  Configuration for refreshing nodes or resources from a hardware provider connection
	*/
	Configuration *OneOfRefreshConnectionSpecConfiguration `json:"configuration"`
}

func (p *RefreshConnectionSpec) MarshalJSON() ([]byte, error) {
	type RefreshConnectionSpecProxy RefreshConnectionSpec

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*RefreshConnectionSpecProxy
		Configuration *OneOfRefreshConnectionSpecConfiguration `json:"configuration,omitempty"`
	}{
		RefreshConnectionSpecProxy: (*RefreshConnectionSpecProxy)(p),
		Configuration:              p.Configuration,
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

func (p *RefreshConnectionSpec) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias RefreshConnectionSpec
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewRefreshConnectionSpec()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ConfigurationItemDiscriminator_ != nil {
		p.ConfigurationItemDiscriminator_ = known.ConfigurationItemDiscriminator_
	}
	if known.Configuration != nil {
		p.Configuration = known.Configuration
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$configurationItemDiscriminator")
	delete(allFields, "configuration")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewRefreshConnectionSpec() *RefreshConnectionSpec {
	p := new(RefreshConnectionSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.RefreshConnectionSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *RefreshConnectionSpec) GetConfiguration() interface{} {
	if nil == p.Configuration {
		return nil
	}
	return p.Configuration.GetValue()
}

func (p *RefreshConnectionSpec) SetConfiguration(v interface{}) error {
	if nil == p.Configuration {
		p.Configuration = NewOneOfRefreshConnectionSpecConfiguration()
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
REST response for all response codes in API path /lifecycle/v4.3/config/nodes/{extId}/$actions/refresh Post operation
*/
type RefreshNodeApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfRefreshNodeApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *RefreshNodeApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias RefreshNodeApiResponse

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
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

func (p *RefreshNodeApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias RefreshNodeApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewRefreshNodeApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewRefreshNodeApiResponse() *RefreshNodeApiResponse {
	p := new(RefreshNodeApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.RefreshNodeApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *RefreshNodeApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *RefreshNodeApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfRefreshNodeApiResponseData()
	}
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
Specification for refreshing resources from a hardware provider connection
*/
type RefreshResourcesSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of group IDs to refresh resources for. Mandatory param if it's cisco connection
	*/
	GroupIds []string `json:"groupIds,omitempty"`
	/*
	  Whether to refresh IP pools from the hardware provider. Default is true
	*/
	ShouldRefreshIpPools *bool `json:"shouldRefreshIpPools,omitempty"`
	/*
	  Whether to refresh MAC pools from the hardware provider. Default is true
	*/
	ShouldRefreshMacPools *bool `json:"shouldRefreshMacPools,omitempty"`
	/*
	  Whether to refresh server identity pools from the hardware provider. Default is true
	*/
	ShouldRefreshServerIdentityPools *bool `json:"shouldRefreshServerIdentityPools,omitempty"`
}

func (p *RefreshResourcesSpec) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias RefreshResourcesSpec

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
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

func (p *RefreshResourcesSpec) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias RefreshResourcesSpec
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewRefreshResourcesSpec()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.GroupIds != nil {
		p.GroupIds = known.GroupIds
	}
	if known.ShouldRefreshIpPools != nil {
		p.ShouldRefreshIpPools = known.ShouldRefreshIpPools
	}
	if known.ShouldRefreshMacPools != nil {
		p.ShouldRefreshMacPools = known.ShouldRefreshMacPools
	}
	if known.ShouldRefreshServerIdentityPools != nil {
		p.ShouldRefreshServerIdentityPools = known.ShouldRefreshServerIdentityPools
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "groupIds")
	delete(allFields, "shouldRefreshIpPools")
	delete(allFields, "shouldRefreshMacPools")
	delete(allFields, "shouldRefreshServerIdentityPools")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewRefreshResourcesSpec() *RefreshResourcesSpec {
	p := new(RefreshResourcesSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.RefreshResourcesSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	p.ShouldRefreshIpPools = new(bool)
	*p.ShouldRefreshIpPools = true
	p.ShouldRefreshMacPools = new(bool)
	*p.ShouldRefreshMacPools = true
	p.ShouldRefreshServerIdentityPools = new(bool)
	*p.ShouldRefreshServerIdentityPools = true

	return p
}

/*
Details of a server identity pool available in the hardware provider
*/
type ServerIdentityPool struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Number of available identities in the server identity pool
	*/
	AvailableCount *int `json:"availableCount,omitempty"`

	DisplayNameMapping *ServerIdentityPoolDisplayNameMapping `json:"displayNameMapping,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  List of groups this server identity pool belongs to
	*/
	Groups []PoolGroup `json:"groups,omitempty"`
	/*
	  Time when the server identity pool was last updated
	*/
	LastUpdatedTime *time.Time `json:"lastUpdatedTime"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Name of the server identity pool
	*/
	Name *string `json:"name,omitempty"`
	/*
	  Reference ID of the server identity pool in the hardware provider
	*/
	ReferenceId *string `json:"referenceId,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *ServerIdentityPool) MarshalJSON() ([]byte, error) {
	type ServerIdentityPoolProxy ServerIdentityPool

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*ServerIdentityPoolProxy
		LastUpdatedTime *time.Time `json:"lastUpdatedTime,omitempty"`
	}{
		ServerIdentityPoolProxy: (*ServerIdentityPoolProxy)(p),
		LastUpdatedTime:         p.LastUpdatedTime,
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

func (p *ServerIdentityPool) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ServerIdentityPool
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewServerIdentityPool()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.AvailableCount != nil {
		p.AvailableCount = known.AvailableCount
	}
	if known.DisplayNameMapping != nil {
		p.DisplayNameMapping = known.DisplayNameMapping
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.Groups != nil {
		p.Groups = known.Groups
	}
	if known.LastUpdatedTime != nil {
		p.LastUpdatedTime = known.LastUpdatedTime
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.ReferenceId != nil {
		p.ReferenceId = known.ReferenceId
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "availableCount")
	delete(allFields, "displayNameMapping")
	delete(allFields, "extId")
	delete(allFields, "groups")
	delete(allFields, "lastUpdatedTime")
	delete(allFields, "links")
	delete(allFields, "name")
	delete(allFields, "referenceId")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewServerIdentityPool() *ServerIdentityPool {
	p := new(ServerIdentityPool)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.ServerIdentityPool"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Mapping between unified attribute names and vendor-specific attribute names for server identity pools
*/
type ServerIdentityPoolDisplayNameMapping struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Vendor-specific attribute name for the group of the server identity pool
	*/
	Group *string `json:"group,omitempty"`
}

func (p *ServerIdentityPoolDisplayNameMapping) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ServerIdentityPoolDisplayNameMapping

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
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

func (p *ServerIdentityPoolDisplayNameMapping) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ServerIdentityPoolDisplayNameMapping
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewServerIdentityPoolDisplayNameMapping()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Group != nil {
		p.Group = known.Group
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "group")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewServerIdentityPoolDisplayNameMapping() *ServerIdentityPoolDisplayNameMapping {
	p := new(ServerIdentityPoolDisplayNameMapping)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.ServerIdentityPoolDisplayNameMapping"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Configuration settings for the server hardware
*/
type ServerSettingsConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	NetworkAdaptorFecMode *NetworkAdaptorFecMode `json:"networkAdaptorFecMode,omitempty"`
	/*
	  UUID of the server identity pool to be used for server configuration
	*/
	ServerIdentityPoolExtId *string `json:"serverIdentityPoolExtId,omitempty"`
}

func (p *ServerSettingsConfig) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ServerSettingsConfig

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
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

func (p *ServerSettingsConfig) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ServerSettingsConfig
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewServerSettingsConfig()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.NetworkAdaptorFecMode != nil {
		p.NetworkAdaptorFecMode = known.NetworkAdaptorFecMode
	}
	if known.ServerIdentityPoolExtId != nil {
		p.ServerIdentityPoolExtId = known.ServerIdentityPoolExtId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "networkAdaptorFecMode")
	delete(allFields, "serverIdentityPoolExtId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewServerSettingsConfig() *ServerSettingsConfig {
	p := new(ServerSettingsConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.ServerSettingsConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
IP configuration using a static IP address
*/
type StaticIPConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Gateway *import1.IPAddress `json:"gateway,omitempty"`

	IpAddress *import1.IPAddress `json:"ipAddress,omitempty"`
	/*
	  Configuration for a single VLAN identifier
	*/
	VlanId *int `json:"vlanId,omitempty"`
}

func (p *StaticIPConfig) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias StaticIPConfig

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
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

func (p *StaticIPConfig) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias StaticIPConfig
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewStaticIPConfig()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Gateway != nil {
		p.Gateway = known.Gateway
	}
	if known.IpAddress != nil {
		p.IpAddress = known.IpAddress
	}
	if known.VlanId != nil {
		p.VlanId = known.VlanId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "gateway")
	delete(allFields, "ipAddress")
	delete(allFields, "vlanId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewStaticIPConfig() *StaticIPConfig {
	p := new(StaticIPConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.StaticIPConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Configuration to unconfigure a node and optionally clean up policies
*/
type UnConfigureServerDetails struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  External ID of the node
	*/
	NodeExtId *string `json:"nodeExtId"`
}

func (p *UnConfigureServerDetails) MarshalJSON() ([]byte, error) {
	type UnConfigureServerDetailsProxy UnConfigureServerDetails

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*UnConfigureServerDetailsProxy
		NodeExtId *string `json:"nodeExtId,omitempty"`
	}{
		UnConfigureServerDetailsProxy: (*UnConfigureServerDetailsProxy)(p),
		NodeExtId:                     p.NodeExtId,
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

func (p *UnConfigureServerDetails) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias UnConfigureServerDetails
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewUnConfigureServerDetails()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.NodeExtId != nil {
		p.NodeExtId = known.NodeExtId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "nodeExtId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewUnConfigureServerDetails() *UnConfigureServerDetails {
	p := new(UnConfigureServerDetails)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.UnConfigureServerDetails"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Unified BMC IP addressing configuration with direct IPv4 and IPv6 settings
*/
type UnifiedBmcIPConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	Ipv4ItemDiscriminator_ *string `json:"$ipv4ItemDiscriminator,omitempty"`
	/*
	  IPv4 configuration for the BMC
	*/
	Ipv4 *OneOfUnifiedBmcIPConfigIpv4 `json:"ipv4,omitempty"`
	/*

	 */
	Ipv6ItemDiscriminator_ *string `json:"$ipv6ItemDiscriminator,omitempty"`
	/*
	  IPv6 configuration for the BMC
	*/
	Ipv6 *OneOfUnifiedBmcIPConfigIpv6 `json:"ipv6,omitempty"`
}

func (p *UnifiedBmcIPConfig) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias UnifiedBmcIPConfig

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
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

func (p *UnifiedBmcIPConfig) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias UnifiedBmcIPConfig
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewUnifiedBmcIPConfig()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Ipv4ItemDiscriminator_ != nil {
		p.Ipv4ItemDiscriminator_ = known.Ipv4ItemDiscriminator_
	}
	if known.Ipv4 != nil {
		p.Ipv4 = known.Ipv4
	}
	if known.Ipv6ItemDiscriminator_ != nil {
		p.Ipv6ItemDiscriminator_ = known.Ipv6ItemDiscriminator_
	}
	if known.Ipv6 != nil {
		p.Ipv6 = known.Ipv6
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$ipv4ItemDiscriminator")
	delete(allFields, "ipv4")
	delete(allFields, "$ipv6ItemDiscriminator")
	delete(allFields, "ipv6")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewUnifiedBmcIPConfig() *UnifiedBmcIPConfig {
	p := new(UnifiedBmcIPConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.UnifiedBmcIPConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /lifecycle/v4.3/config/claim-tokens/{extId} Put operation
*/
type UpdateClaimTokenApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateClaimTokenApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *UpdateClaimTokenApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias UpdateClaimTokenApiResponse

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
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

func (p *UpdateClaimTokenApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias UpdateClaimTokenApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewUpdateClaimTokenApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewUpdateClaimTokenApiResponse() *UpdateClaimTokenApiResponse {
	p := new(UpdateClaimTokenApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.UpdateClaimTokenApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdateClaimTokenApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdateClaimTokenApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdateClaimTokenApiResponseData()
	}
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
REST response for all response codes in API path /lifecycle/v4.3/config/hardware-providers/{hardwareProviderExtId}/connections/{extId} Put operation
*/
type UpdateConnectionApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateConnectionApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *UpdateConnectionApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias UpdateConnectionApiResponse

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
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

func (p *UpdateConnectionApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias UpdateConnectionApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewUpdateConnectionApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewUpdateConnectionApiResponse() *UpdateConnectionApiResponse {
	p := new(UpdateConnectionApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.UpdateConnectionApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdateConnectionApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdateConnectionApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdateConnectionApiResponseData()
	}
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
REST response for all response codes in API path /lifecycle/v4.3/config/images/{extId} Put operation
*/
type UpdateImageApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateImageApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *UpdateImageApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias UpdateImageApiResponse

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
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

func (p *UpdateImageApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias UpdateImageApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewUpdateImageApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewUpdateImageApiResponse() *UpdateImageApiResponse {
	p := new(UpdateImageApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.UpdateImageApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdateImageApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdateImageApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdateImageApiResponseData()
	}
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
REST response for all response codes in API path /lifecycle/v4.3/config/nodes/{extId} Put operation
*/
type UpdateNodeApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateNodeApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *UpdateNodeApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias UpdateNodeApiResponse

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
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

func (p *UpdateNodeApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias UpdateNodeApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewUpdateNodeApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewUpdateNodeApiResponse() *UpdateNodeApiResponse {
	p := new(UpdateNodeApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.UpdateNodeApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdateNodeApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdateNodeApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdateNodeApiResponseData()
	}
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
Endpoint configuration using URL
*/
type UrlEndpoint struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  URL for the endpoint
	*/
	Url *string `json:"url"`
}

func (p *UrlEndpoint) MarshalJSON() ([]byte, error) {
	type UrlEndpointProxy UrlEndpoint

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*UrlEndpointProxy
		Url *string `json:"url,omitempty"`
	}{
		UrlEndpointProxy: (*UrlEndpointProxy)(p),
		Url:              p.Url,
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

func (p *UrlEndpoint) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias UrlEndpoint
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewUrlEndpoint()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Url != nil {
		p.Url = known.Url
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "url")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewUrlEndpoint() *UrlEndpoint {
	p := new(UrlEndpoint)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.UrlEndpoint"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Configuration for a single VLAN identifier
*/
type VlanId struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Configuration for a single VLAN identifier
	*/
	Id *int `json:"id"`
}

func (p *VlanId) MarshalJSON() ([]byte, error) {
	type VlanIdProxy VlanId

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*VlanIdProxy
		Id *int `json:"id,omitempty"`
	}{
		VlanIdProxy: (*VlanIdProxy)(p),
		Id:          p.Id,
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

func (p *VlanId) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias VlanId
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewVlanId()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Id != nil {
		p.Id = known.Id
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "id")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewVlanId() *VlanId {
	p := new(VlanId)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.VlanId"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
VLAN configuration that can be either a single VLAN ID or a VLAN range
*/
type VlanInfo struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	ValueItemDiscriminator_ *string `json:"$valueItemDiscriminator,omitempty"`
	/*
	  Value of the VLAN configuration that can be either a single VLAN ID or a VLAN range
	*/
	Value *OneOfVlanInfoValue `json:"value"`
}

func (p *VlanInfo) MarshalJSON() ([]byte, error) {
	type VlanInfoProxy VlanInfo

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*VlanInfoProxy
		Value *OneOfVlanInfoValue `json:"value,omitempty"`
	}{
		VlanInfoProxy: (*VlanInfoProxy)(p),
		Value:         p.Value,
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

func (p *VlanInfo) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias VlanInfo
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewVlanInfo()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ValueItemDiscriminator_ != nil {
		p.ValueItemDiscriminator_ = known.ValueItemDiscriminator_
	}
	if known.Value != nil {
		p.Value = known.Value
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$valueItemDiscriminator")
	delete(allFields, "value")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewVlanInfo() *VlanInfo {
	p := new(VlanInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.VlanInfo"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *VlanInfo) GetValue() interface{} {
	if nil == p.Value {
		return nil
	}
	return p.Value.GetValue()
}

func (p *VlanInfo) SetValue(v interface{}) error {
	if nil == p.Value {
		p.Value = NewOneOfVlanInfoValue()
	}
	e := p.Value.SetValue(v)
	if nil == e {
		if nil == p.ValueItemDiscriminator_ {
			p.ValueItemDiscriminator_ = new(string)
		}
		*p.ValueItemDiscriminator_ = *p.Value.Discriminator
	}
	return e
}

/*
Configuration for a range of VLAN identifiers
*/
type VlanRange struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Ending VLAN ID for the range (inclusive)
	*/
	EndId *int `json:"endId"`
	/*
	  Starting VLAN ID for the range (inclusive)
	*/
	StartId *int `json:"startId"`
}

func (p *VlanRange) MarshalJSON() ([]byte, error) {
	type VlanRangeProxy VlanRange

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*VlanRangeProxy
		EndId   *int `json:"endId,omitempty"`
		StartId *int `json:"startId,omitempty"`
	}{
		VlanRangeProxy: (*VlanRangeProxy)(p),
		EndId:          p.EndId,
		StartId:        p.StartId,
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

func (p *VlanRange) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias VlanRange
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewVlanRange()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.EndId != nil {
		p.EndId = known.EndId
	}
	if known.StartId != nil {
		p.StartId = known.StartId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "endId")
	delete(allFields, "startId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewVlanRange() *VlanRange {
	p := new(VlanRange)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.VlanRange"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type OneOfGetConnectionNodeApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType1035 *DiscoveredNode        `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfGetConnectionNodeApiResponseData() *OneOfGetConnectionNodeApiResponseData {
	p := new(OneOfGetConnectionNodeApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetConnectionNodeApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetConnectionNodeApiResponseData is nil"))
	}
	switch v.(type) {
	case DiscoveredNode:
		if nil == p.oneOfType1035 {
			p.oneOfType1035 = new(DiscoveredNode)
		}
		*p.oneOfType1035 = v.(DiscoveredNode)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType1035.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType1035.ObjectType_
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

func (p *OneOfGetConnectionNodeApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType1035 != nil && *p.oneOfType1035.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1035
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetConnectionNodeApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType1035 := new(DiscoveredNode)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType1035)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType1035.ObjectType_ != nil && "lifecycle.v4.config.DiscoveredNode" == *vOneOfType1035.ObjectType_ {
							if nil == p.oneOfType1035 {
								p.oneOfType1035 = new(DiscoveredNode)
							}
							*p.oneOfType1035 = *vOneOfType1035
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType1035.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType1035.ObjectType_
							return nil
						}
					}
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import4.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType1035 := new(DiscoveredNode)
	if err := json.Unmarshal(b, vOneOfType1035); err == nil {
		if vOneOfType1035.ObjectType_ != nil && "lifecycle.v4.config.DiscoveredNode" == *vOneOfType1035.ObjectType_ {
			if nil == p.oneOfType1035 {
				p.oneOfType1035 = new(DiscoveredNode)
			}
			*p.oneOfType1035 = *vOneOfType1035
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType1035.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType1035.ObjectType_
			return nil
		}
	}
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetConnectionNodeApiResponseData"))
}

func (p *OneOfGetConnectionNodeApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType1035 != nil && *p.oneOfType1035.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1035)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetConnectionNodeApiResponseData")
}

type OneOfUnifiedBmcIPConfigIpv6 struct {
	Discriminator *string         `json:"-"`
	ObjectType_   *string         `json:"-"`
	oneOfType5137 *IPPoolConfig   `json:"-"`
	oneOfType5138 *StaticIPConfig `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfUnifiedBmcIPConfigIpv6() *OneOfUnifiedBmcIPConfigIpv6 {
	p := new(OneOfUnifiedBmcIPConfigIpv6)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUnifiedBmcIPConfigIpv6) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUnifiedBmcIPConfigIpv6 is nil"))
	}
	switch v.(type) {
	case IPPoolConfig:
		if nil == p.oneOfType5137 {
			p.oneOfType5137 = new(IPPoolConfig)
		}
		*p.oneOfType5137 = v.(IPPoolConfig)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType5137.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType5137.ObjectType_
	case StaticIPConfig:
		if nil == p.oneOfType5138 {
			p.oneOfType5138 = new(StaticIPConfig)
		}
		*p.oneOfType5138 = v.(StaticIPConfig)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType5138.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType5138.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfUnifiedBmcIPConfigIpv6) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType5137 != nil && *p.oneOfType5137.ObjectType_ == *p.Discriminator {
		return *p.oneOfType5137
	}
	if p.oneOfType5138 != nil && *p.oneOfType5138.ObjectType_ == *p.Discriminator {
		return *p.oneOfType5138
	}
	return nil
}

func (p *OneOfUnifiedBmcIPConfigIpv6) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType5137 := new(IPPoolConfig)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType5137)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType5137.ObjectType_ != nil && "lifecycle.v4.config.IPPoolConfig" == *vOneOfType5137.ObjectType_ {
							if nil == p.oneOfType5137 {
								p.oneOfType5137 = new(IPPoolConfig)
							}
							*p.oneOfType5137 = *vOneOfType5137
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType5137.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType5137.ObjectType_
							return nil
						}
					}
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType5138 := new(StaticIPConfig)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType5138)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType5138.ObjectType_ != nil && "lifecycle.v4.config.StaticIPConfig" == *vOneOfType5138.ObjectType_ {
							if nil == p.oneOfType5138 {
								p.oneOfType5138 = new(StaticIPConfig)
							}
							*p.oneOfType5138 = *vOneOfType5138
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType5138.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType5138.ObjectType_
							return nil
						}
					}
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType5137 := new(IPPoolConfig)
	if err := json.Unmarshal(b, vOneOfType5137); err == nil {
		if vOneOfType5137.ObjectType_ != nil && "lifecycle.v4.config.IPPoolConfig" == *vOneOfType5137.ObjectType_ {
			if nil == p.oneOfType5137 {
				p.oneOfType5137 = new(IPPoolConfig)
			}
			*p.oneOfType5137 = *vOneOfType5137
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType5137.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType5137.ObjectType_
			return nil
		}
	}
	vOneOfType5138 := new(StaticIPConfig)
	if err := json.Unmarshal(b, vOneOfType5138); err == nil {
		if vOneOfType5138.ObjectType_ != nil && "lifecycle.v4.config.StaticIPConfig" == *vOneOfType5138.ObjectType_ {
			if nil == p.oneOfType5138 {
				p.oneOfType5138 = new(StaticIPConfig)
			}
			*p.oneOfType5138 = *vOneOfType5138
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType5138.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType5138.ObjectType_
			return nil
		}
	}
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUnifiedBmcIPConfigIpv6"))
}

func (p *OneOfUnifiedBmcIPConfigIpv6) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType5137 != nil && *p.oneOfType5137.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType5137)
	}
	if p.oneOfType5138 != nil && *p.oneOfType5138.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType5138)
	}
	return nil, errors.New("No value to marshal for OneOfUnifiedBmcIPConfigIpv6")
}

type OneOfHciHypervisorDetailsImageDetails struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType3201 *LocalHostImageDetails `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfHciHypervisorDetailsImageDetails() *OneOfHciHypervisorDetailsImageDetails {
	p := new(OneOfHciHypervisorDetailsImageDetails)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfHciHypervisorDetailsImageDetails) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfHciHypervisorDetailsImageDetails is nil"))
	}
	switch v.(type) {
	case LocalHostImageDetails:
		if nil == p.oneOfType3201 {
			p.oneOfType3201 = new(LocalHostImageDetails)
		}
		*p.oneOfType3201 = v.(LocalHostImageDetails)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType3201.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType3201.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfHciHypervisorDetailsImageDetails) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType3201 != nil && *p.oneOfType3201.ObjectType_ == *p.Discriminator {
		return *p.oneOfType3201
	}
	return nil
}

func (p *OneOfHciHypervisorDetailsImageDetails) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType3201 := new(LocalHostImageDetails)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType3201)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType3201.ObjectType_ != nil && "lifecycle.v4.config.LocalHostImageDetails" == *vOneOfType3201.ObjectType_ {
							if nil == p.oneOfType3201 {
								p.oneOfType3201 = new(LocalHostImageDetails)
							}
							*p.oneOfType3201 = *vOneOfType3201
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType3201.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType3201.ObjectType_
							return nil
						}
					}
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType3201 := new(LocalHostImageDetails)
	if err := json.Unmarshal(b, vOneOfType3201); err == nil {
		if vOneOfType3201.ObjectType_ != nil && "lifecycle.v4.config.LocalHostImageDetails" == *vOneOfType3201.ObjectType_ {
			if nil == p.oneOfType3201 {
				p.oneOfType3201 = new(LocalHostImageDetails)
			}
			*p.oneOfType3201 = *vOneOfType3201
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType3201.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType3201.ObjectType_
			return nil
		}
	}
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfHciHypervisorDetailsImageDetails"))
}

func (p *OneOfHciHypervisorDetailsImageDetails) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType3201 != nil && *p.oneOfType3201.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType3201)
	}
	return nil, errors.New("No value to marshal for OneOfHciHypervisorDetailsImageDetails")
}

type OneOfListPatchedImagesApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType1002 []PatchedImage         `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfListPatchedImagesApiResponseData() *OneOfListPatchedImagesApiResponseData {
	p := new(OneOfListPatchedImagesApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListPatchedImagesApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListPatchedImagesApiResponseData is nil"))
	}
	switch v.(type) {
	case []PatchedImage:
		p.oneOfType1002 = v.([]PatchedImage)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<lifecycle.v4.config.PatchedImage>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<lifecycle.v4.config.PatchedImage>"
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

func (p *OneOfListPatchedImagesApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if "List<lifecycle.v4.config.PatchedImage>" == *p.Discriminator {
		return p.oneOfType1002
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfListPatchedImagesApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["List<lifecycle.v4.config.PatchedImage>"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType1002 := new([]PatchedImage)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType1002)
					if unmarshalErr == nil {
						// For arrays, verify the array item ObjectType matches
						if vOneOfType1002 == nil || len(*vOneOfType1002) == 0 || ((*vOneOfType1002)[0].ObjectType_ != nil && "lifecycle.v4.config.PatchedImage" == *((*vOneOfType1002)[0].ObjectType_)) {
							p.oneOfType1002 = *vOneOfType1002
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = "List<lifecycle.v4.config.PatchedImage>"
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = "List<lifecycle.v4.config.PatchedImage>"
							return nil
						}
					}
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import4.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType1002 := new([]PatchedImage)
	if err := json.Unmarshal(b, vOneOfType1002); err == nil {
		if len(*vOneOfType1002) == 0 || (vOneOfType1002 != nil && (*vOneOfType1002)[0].ObjectType_ != nil && "lifecycle.v4.config.PatchedImage" == *((*vOneOfType1002)[0].ObjectType_)) {
			p.oneOfType1002 = *vOneOfType1002
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<lifecycle.v4.config.PatchedImage>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<lifecycle.v4.config.PatchedImage>"
			return nil
		}
	}
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListPatchedImagesApiResponseData"))
}

func (p *OneOfListPatchedImagesApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if "List<lifecycle.v4.config.PatchedImage>" == *p.Discriminator {
		return json.Marshal(p.oneOfType1002)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfListPatchedImagesApiResponseData")
}

type OneOfListServerIdentityPoolsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType1143 []ServerIdentityPool   `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfListServerIdentityPoolsApiResponseData() *OneOfListServerIdentityPoolsApiResponseData {
	p := new(OneOfListServerIdentityPoolsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListServerIdentityPoolsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListServerIdentityPoolsApiResponseData is nil"))
	}
	switch v.(type) {
	case []ServerIdentityPool:
		p.oneOfType1143 = v.([]ServerIdentityPool)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<lifecycle.v4.config.ServerIdentityPool>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<lifecycle.v4.config.ServerIdentityPool>"
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

func (p *OneOfListServerIdentityPoolsApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if "List<lifecycle.v4.config.ServerIdentityPool>" == *p.Discriminator {
		return p.oneOfType1143
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfListServerIdentityPoolsApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["List<lifecycle.v4.config.ServerIdentityPool>"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType1143 := new([]ServerIdentityPool)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType1143)
					if unmarshalErr == nil {
						// For arrays, verify the array item ObjectType matches
						if vOneOfType1143 == nil || len(*vOneOfType1143) == 0 || ((*vOneOfType1143)[0].ObjectType_ != nil && "lifecycle.v4.config.ServerIdentityPool" == *((*vOneOfType1143)[0].ObjectType_)) {
							p.oneOfType1143 = *vOneOfType1143
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = "List<lifecycle.v4.config.ServerIdentityPool>"
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = "List<lifecycle.v4.config.ServerIdentityPool>"
							return nil
						}
					}
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import4.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType1143 := new([]ServerIdentityPool)
	if err := json.Unmarshal(b, vOneOfType1143); err == nil {
		if len(*vOneOfType1143) == 0 || (vOneOfType1143 != nil && (*vOneOfType1143)[0].ObjectType_ != nil && "lifecycle.v4.config.ServerIdentityPool" == *((*vOneOfType1143)[0].ObjectType_)) {
			p.oneOfType1143 = *vOneOfType1143
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<lifecycle.v4.config.ServerIdentityPool>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<lifecycle.v4.config.ServerIdentityPool>"
			return nil
		}
	}
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListServerIdentityPoolsApiResponseData"))
}

func (p *OneOfListServerIdentityPoolsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if "List<lifecycle.v4.config.ServerIdentityPool>" == *p.Discriminator {
		return json.Marshal(p.oneOfType1143)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfListServerIdentityPoolsApiResponseData")
}

type OneOfConfigureNodeResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType1013 *import3.TaskReference `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfConfigureNodeResponseData() *OneOfConfigureNodeResponseData {
	p := new(OneOfConfigureNodeResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfConfigureNodeResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfConfigureNodeResponseData is nil"))
	}
	switch v.(type) {
	case import3.TaskReference:
		if nil == p.oneOfType1013 {
			p.oneOfType1013 = new(import3.TaskReference)
		}
		*p.oneOfType1013 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType1013.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType1013.ObjectType_
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

func (p *OneOfConfigureNodeResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType1013 != nil && *p.oneOfType1013.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1013
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfConfigureNodeResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType1013 := new(import3.TaskReference)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType1013)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType1013.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType1013.ObjectType_ {
							if nil == p.oneOfType1013 {
								p.oneOfType1013 = new(import3.TaskReference)
							}
							*p.oneOfType1013 = *vOneOfType1013
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType1013.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType1013.ObjectType_
							return nil
						}
					}
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import4.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType1013 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType1013); err == nil {
		if vOneOfType1013.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType1013.ObjectType_ {
			if nil == p.oneOfType1013 {
				p.oneOfType1013 = new(import3.TaskReference)
			}
			*p.oneOfType1013 = *vOneOfType1013
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType1013.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType1013.ObjectType_
			return nil
		}
	}
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfConfigureNodeResponseData"))
}

func (p *OneOfConfigureNodeResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType1013 != nil && *p.oneOfType1013.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1013)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfConfigureNodeResponseData")
}

type OneOfGetNodeApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType401  *NodeProjection        `json:"-"`
	oneOfType1003 *Node                  `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfGetNodeApiResponseData() *OneOfGetNodeApiResponseData {
	p := new(OneOfGetNodeApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetNodeApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetNodeApiResponseData is nil"))
	}
	switch v.(type) {
	case NodeProjection:
		if nil == p.oneOfType401 {
			p.oneOfType401 = new(NodeProjection)
		}
		*p.oneOfType401 = v.(NodeProjection)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType401.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType401.ObjectType_
	case Node:
		if nil == p.oneOfType1003 {
			p.oneOfType1003 = new(Node)
		}
		*p.oneOfType1003 = v.(Node)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType1003.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType1003.ObjectType_
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

func (p *OneOfGetNodeApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType401 != nil && *p.oneOfType401.ObjectType_ == *p.Discriminator {
		return *p.oneOfType401
	}
	if p.oneOfType1003 != nil && *p.oneOfType1003.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1003
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetNodeApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType401 := new(NodeProjection)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType401)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType401.ObjectType_ != nil && "lifecycle.v4.config.NodeProjection" == *vOneOfType401.ObjectType_ {
							if nil == p.oneOfType401 {
								p.oneOfType401 = new(NodeProjection)
							}
							*p.oneOfType401 = *vOneOfType401
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType401.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType401.ObjectType_
							return nil
						}
					}
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType1003 := new(Node)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType1003)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType1003.ObjectType_ != nil && "lifecycle.v4.config.Node" == *vOneOfType1003.ObjectType_ {
							if nil == p.oneOfType1003 {
								p.oneOfType1003 = new(Node)
							}
							*p.oneOfType1003 = *vOneOfType1003
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType1003.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType1003.ObjectType_
							return nil
						}
					}
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import4.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType401 := new(NodeProjection)
	if err := json.Unmarshal(b, vOneOfType401); err == nil {
		if vOneOfType401.ObjectType_ != nil && "lifecycle.v4.config.NodeProjection" == *vOneOfType401.ObjectType_ {
			if nil == p.oneOfType401 {
				p.oneOfType401 = new(NodeProjection)
			}
			*p.oneOfType401 = *vOneOfType401
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType401.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType401.ObjectType_
			return nil
		}
	}
	vOneOfType1003 := new(Node)
	if err := json.Unmarshal(b, vOneOfType1003); err == nil {
		if vOneOfType1003.ObjectType_ != nil && "lifecycle.v4.config.Node" == *vOneOfType1003.ObjectType_ {
			if nil == p.oneOfType1003 {
				p.oneOfType1003 = new(Node)
			}
			*p.oneOfType1003 = *vOneOfType1003
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType1003.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType1003.ObjectType_
			return nil
		}
	}
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetNodeApiResponseData"))
}

func (p *OneOfGetNodeApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType401 != nil && *p.oneOfType401.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType401)
	}
	if p.oneOfType1003 != nil && *p.oneOfType1003.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1003)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetNodeApiResponseData")
}

type OneOfVlanInfoValue struct {
	Discriminator *string    `json:"-"`
	ObjectType_   *string    `json:"-"`
	oneOfType5061 *VlanId    `json:"-"`
	oneOfType5062 *VlanRange `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfVlanInfoValue() *OneOfVlanInfoValue {
	p := new(OneOfVlanInfoValue)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfVlanInfoValue) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfVlanInfoValue is nil"))
	}
	switch v.(type) {
	case VlanId:
		if nil == p.oneOfType5061 {
			p.oneOfType5061 = new(VlanId)
		}
		*p.oneOfType5061 = v.(VlanId)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType5061.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType5061.ObjectType_
	case VlanRange:
		if nil == p.oneOfType5062 {
			p.oneOfType5062 = new(VlanRange)
		}
		*p.oneOfType5062 = v.(VlanRange)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType5062.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType5062.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfVlanInfoValue) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType5061 != nil && *p.oneOfType5061.ObjectType_ == *p.Discriminator {
		return *p.oneOfType5061
	}
	if p.oneOfType5062 != nil && *p.oneOfType5062.ObjectType_ == *p.Discriminator {
		return *p.oneOfType5062
	}
	return nil
}

func (p *OneOfVlanInfoValue) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType5061 := new(VlanId)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType5061)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType5061.ObjectType_ != nil && "lifecycle.v4.config.VlanId" == *vOneOfType5061.ObjectType_ {
							if nil == p.oneOfType5061 {
								p.oneOfType5061 = new(VlanId)
							}
							*p.oneOfType5061 = *vOneOfType5061
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType5061.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType5061.ObjectType_
							return nil
						}
					}
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType5062 := new(VlanRange)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType5062)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType5062.ObjectType_ != nil && "lifecycle.v4.config.VlanRange" == *vOneOfType5062.ObjectType_ {
							if nil == p.oneOfType5062 {
								p.oneOfType5062 = new(VlanRange)
							}
							*p.oneOfType5062 = *vOneOfType5062
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType5062.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType5062.ObjectType_
							return nil
						}
					}
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType5061 := new(VlanId)
	if err := json.Unmarshal(b, vOneOfType5061); err == nil {
		if vOneOfType5061.ObjectType_ != nil && "lifecycle.v4.config.VlanId" == *vOneOfType5061.ObjectType_ {
			if nil == p.oneOfType5061 {
				p.oneOfType5061 = new(VlanId)
			}
			*p.oneOfType5061 = *vOneOfType5061
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType5061.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType5061.ObjectType_
			return nil
		}
	}
	vOneOfType5062 := new(VlanRange)
	if err := json.Unmarshal(b, vOneOfType5062); err == nil {
		if vOneOfType5062.ObjectType_ != nil && "lifecycle.v4.config.VlanRange" == *vOneOfType5062.ObjectType_ {
			if nil == p.oneOfType5062 {
				p.oneOfType5062 = new(VlanRange)
			}
			*p.oneOfType5062 = *vOneOfType5062
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType5062.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType5062.ObjectType_
			return nil
		}
	}
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfVlanInfoValue"))
}

func (p *OneOfVlanInfoValue) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType5061 != nil && *p.oneOfType5061.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType5061)
	}
	if p.oneOfType5062 != nil && *p.oneOfType5062.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType5062)
	}
	return nil, errors.New("No value to marshal for OneOfVlanInfoValue")
}

type OneOfConnectionDetailEndpoint struct {
	Discriminator *string            `json:"-"`
	ObjectType_   *string            `json:"-"`
	oneOfType2026 *UrlEndpoint       `json:"-"`
	oneOfType2027 *IpRangeEndpoint   `json:"-"`
	oneOfType2028 *IpAddressEndpoint `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfConnectionDetailEndpoint() *OneOfConnectionDetailEndpoint {
	p := new(OneOfConnectionDetailEndpoint)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfConnectionDetailEndpoint) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfConnectionDetailEndpoint is nil"))
	}
	switch v.(type) {
	case UrlEndpoint:
		if nil == p.oneOfType2026 {
			p.oneOfType2026 = new(UrlEndpoint)
		}
		*p.oneOfType2026 = v.(UrlEndpoint)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2026.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2026.ObjectType_
	case IpRangeEndpoint:
		if nil == p.oneOfType2027 {
			p.oneOfType2027 = new(IpRangeEndpoint)
		}
		*p.oneOfType2027 = v.(IpRangeEndpoint)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2027.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2027.ObjectType_
	case IpAddressEndpoint:
		if nil == p.oneOfType2028 {
			p.oneOfType2028 = new(IpAddressEndpoint)
		}
		*p.oneOfType2028 = v.(IpAddressEndpoint)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2028.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2028.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfConnectionDetailEndpoint) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType2026 != nil && *p.oneOfType2026.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2026
	}
	if p.oneOfType2027 != nil && *p.oneOfType2027.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2027
	}
	if p.oneOfType2028 != nil && *p.oneOfType2028.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2028
	}
	return nil
}

func (p *OneOfConnectionDetailEndpoint) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType2026 := new(UrlEndpoint)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType2026)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType2026.ObjectType_ != nil && "lifecycle.v4.config.UrlEndpoint" == *vOneOfType2026.ObjectType_ {
							if nil == p.oneOfType2026 {
								p.oneOfType2026 = new(UrlEndpoint)
							}
							*p.oneOfType2026 = *vOneOfType2026
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType2026.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType2026.ObjectType_
							return nil
						}
					}
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType2027 := new(IpRangeEndpoint)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType2027)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType2027.ObjectType_ != nil && "lifecycle.v4.config.IpRangeEndpoint" == *vOneOfType2027.ObjectType_ {
							if nil == p.oneOfType2027 {
								p.oneOfType2027 = new(IpRangeEndpoint)
							}
							*p.oneOfType2027 = *vOneOfType2027
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType2027.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType2027.ObjectType_
							return nil
						}
					}
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType2028 := new(IpAddressEndpoint)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType2028)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType2028.ObjectType_ != nil && "lifecycle.v4.config.IpAddressEndpoint" == *vOneOfType2028.ObjectType_ {
							if nil == p.oneOfType2028 {
								p.oneOfType2028 = new(IpAddressEndpoint)
							}
							*p.oneOfType2028 = *vOneOfType2028
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType2028.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType2028.ObjectType_
							return nil
						}
					}
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType2026 := new(UrlEndpoint)
	if err := json.Unmarshal(b, vOneOfType2026); err == nil {
		if vOneOfType2026.ObjectType_ != nil && "lifecycle.v4.config.UrlEndpoint" == *vOneOfType2026.ObjectType_ {
			if nil == p.oneOfType2026 {
				p.oneOfType2026 = new(UrlEndpoint)
			}
			*p.oneOfType2026 = *vOneOfType2026
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2026.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2026.ObjectType_
			return nil
		}
	}
	vOneOfType2027 := new(IpRangeEndpoint)
	if err := json.Unmarshal(b, vOneOfType2027); err == nil {
		if vOneOfType2027.ObjectType_ != nil && "lifecycle.v4.config.IpRangeEndpoint" == *vOneOfType2027.ObjectType_ {
			if nil == p.oneOfType2027 {
				p.oneOfType2027 = new(IpRangeEndpoint)
			}
			*p.oneOfType2027 = *vOneOfType2027
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2027.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2027.ObjectType_
			return nil
		}
	}
	vOneOfType2028 := new(IpAddressEndpoint)
	if err := json.Unmarshal(b, vOneOfType2028); err == nil {
		if vOneOfType2028.ObjectType_ != nil && "lifecycle.v4.config.IpAddressEndpoint" == *vOneOfType2028.ObjectType_ {
			if nil == p.oneOfType2028 {
				p.oneOfType2028 = new(IpAddressEndpoint)
			}
			*p.oneOfType2028 = *vOneOfType2028
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2028.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2028.ObjectType_
			return nil
		}
	}
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfConnectionDetailEndpoint"))
}

func (p *OneOfConnectionDetailEndpoint) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType2026 != nil && *p.oneOfType2026.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2026)
	}
	if p.oneOfType2027 != nil && *p.oneOfType2027.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2027)
	}
	if p.oneOfType2028 != nil && *p.oneOfType2028.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2028)
	}
	return nil, errors.New("No value to marshal for OneOfConnectionDetailEndpoint")
}

type OneOfDeleteClaimTokenApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType1007 *import3.TaskReference `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfDeleteClaimTokenApiResponseData() *OneOfDeleteClaimTokenApiResponseData {
	p := new(OneOfDeleteClaimTokenApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteClaimTokenApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteClaimTokenApiResponseData is nil"))
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
	case import3.TaskReference:
		if nil == p.oneOfType1007 {
			p.oneOfType1007 = new(import3.TaskReference)
		}
		*p.oneOfType1007 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType1007.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType1007.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfDeleteClaimTokenApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType1007 != nil && *p.oneOfType1007.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1007
	}
	return nil
}

func (p *OneOfDeleteClaimTokenApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import4.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType1007 := new(import3.TaskReference)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType1007)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType1007.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType1007.ObjectType_ {
							if nil == p.oneOfType1007 {
								p.oneOfType1007 = new(import3.TaskReference)
							}
							*p.oneOfType1007 = *vOneOfType1007
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType1007.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType1007.ObjectType_
							return nil
						}
					}
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType1007 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType1007); err == nil {
		if vOneOfType1007.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType1007.ObjectType_ {
			if nil == p.oneOfType1007 {
				p.oneOfType1007 = new(import3.TaskReference)
			}
			*p.oneOfType1007 = *vOneOfType1007
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType1007.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType1007.ObjectType_
			return nil
		}
	}
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteClaimTokenApiResponseData"))
}

func (p *OneOfDeleteClaimTokenApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType1007 != nil && *p.oneOfType1007.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1007)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteClaimTokenApiResponseData")
}

type OneOfListImagesApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType4003 []Image                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfListImagesApiResponseData() *OneOfListImagesApiResponseData {
	p := new(OneOfListImagesApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListImagesApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListImagesApiResponseData is nil"))
	}
	switch v.(type) {
	case []Image:
		p.oneOfType4003 = v.([]Image)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<lifecycle.v4.config.Image>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<lifecycle.v4.config.Image>"
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

func (p *OneOfListImagesApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if "List<lifecycle.v4.config.Image>" == *p.Discriminator {
		return p.oneOfType4003
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfListImagesApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["List<lifecycle.v4.config.Image>"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType4003 := new([]Image)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType4003)
					if unmarshalErr == nil {
						// For arrays, verify the array item ObjectType matches
						if vOneOfType4003 == nil || len(*vOneOfType4003) == 0 || ((*vOneOfType4003)[0].ObjectType_ != nil && "lifecycle.v4.config.Image" == *((*vOneOfType4003)[0].ObjectType_)) {
							p.oneOfType4003 = *vOneOfType4003
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = "List<lifecycle.v4.config.Image>"
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = "List<lifecycle.v4.config.Image>"
							return nil
						}
					}
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import4.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType4003 := new([]Image)
	if err := json.Unmarshal(b, vOneOfType4003); err == nil {
		if len(*vOneOfType4003) == 0 || (vOneOfType4003 != nil && (*vOneOfType4003)[0].ObjectType_ != nil && "lifecycle.v4.config.Image" == *((*vOneOfType4003)[0].ObjectType_)) {
			p.oneOfType4003 = *vOneOfType4003
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<lifecycle.v4.config.Image>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<lifecycle.v4.config.Image>"
			return nil
		}
	}
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListImagesApiResponseData"))
}

func (p *OneOfListImagesApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if "List<lifecycle.v4.config.Image>" == *p.Discriminator {
		return json.Marshal(p.oneOfType4003)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfListImagesApiResponseData")
}

type OneOfConfigureNodeSpecConfiguration struct {
	Discriminator *string                   `json:"-"`
	ObjectType_   *string                   `json:"-"`
	oneOfType2805 *UnConfigureServerDetails `json:"-"`
	oneOfType2802 *ConfigureServerDetails   `json:"-"`
	oneOfType2803 *PreClusterConfigDetails  `json:"-"`
	oneOfType2804 *PostClusterConfigDetails `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfConfigureNodeSpecConfiguration() *OneOfConfigureNodeSpecConfiguration {
	p := new(OneOfConfigureNodeSpecConfiguration)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfConfigureNodeSpecConfiguration) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfConfigureNodeSpecConfiguration is nil"))
	}
	switch v.(type) {
	case UnConfigureServerDetails:
		if nil == p.oneOfType2805 {
			p.oneOfType2805 = new(UnConfigureServerDetails)
		}
		*p.oneOfType2805 = v.(UnConfigureServerDetails)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2805.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2805.ObjectType_
	case ConfigureServerDetails:
		if nil == p.oneOfType2802 {
			p.oneOfType2802 = new(ConfigureServerDetails)
		}
		*p.oneOfType2802 = v.(ConfigureServerDetails)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2802.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2802.ObjectType_
	case PreClusterConfigDetails:
		if nil == p.oneOfType2803 {
			p.oneOfType2803 = new(PreClusterConfigDetails)
		}
		*p.oneOfType2803 = v.(PreClusterConfigDetails)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2803.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2803.ObjectType_
	case PostClusterConfigDetails:
		if nil == p.oneOfType2804 {
			p.oneOfType2804 = new(PostClusterConfigDetails)
		}
		*p.oneOfType2804 = v.(PostClusterConfigDetails)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2804.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2804.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfConfigureNodeSpecConfiguration) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType2805 != nil && *p.oneOfType2805.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2805
	}
	if p.oneOfType2802 != nil && *p.oneOfType2802.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2802
	}
	if p.oneOfType2803 != nil && *p.oneOfType2803.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2803
	}
	if p.oneOfType2804 != nil && *p.oneOfType2804.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2804
	}
	return nil
}

func (p *OneOfConfigureNodeSpecConfiguration) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType2805 := new(UnConfigureServerDetails)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType2805)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType2805.ObjectType_ != nil && "lifecycle.v4.config.UnConfigureServerDetails" == *vOneOfType2805.ObjectType_ {
							if nil == p.oneOfType2805 {
								p.oneOfType2805 = new(UnConfigureServerDetails)
							}
							*p.oneOfType2805 = *vOneOfType2805
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType2805.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType2805.ObjectType_
							return nil
						}
					}
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType2802 := new(ConfigureServerDetails)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType2802)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType2802.ObjectType_ != nil && "lifecycle.v4.config.ConfigureServerDetails" == *vOneOfType2802.ObjectType_ {
							if nil == p.oneOfType2802 {
								p.oneOfType2802 = new(ConfigureServerDetails)
							}
							*p.oneOfType2802 = *vOneOfType2802
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType2802.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType2802.ObjectType_
							return nil
						}
					}
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType2803 := new(PreClusterConfigDetails)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType2803)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType2803.ObjectType_ != nil && "lifecycle.v4.config.PreClusterConfigDetails" == *vOneOfType2803.ObjectType_ {
							if nil == p.oneOfType2803 {
								p.oneOfType2803 = new(PreClusterConfigDetails)
							}
							*p.oneOfType2803 = *vOneOfType2803
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType2803.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType2803.ObjectType_
							return nil
						}
					}
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType2804 := new(PostClusterConfigDetails)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType2804)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType2804.ObjectType_ != nil && "lifecycle.v4.config.PostClusterConfigDetails" == *vOneOfType2804.ObjectType_ {
							if nil == p.oneOfType2804 {
								p.oneOfType2804 = new(PostClusterConfigDetails)
							}
							*p.oneOfType2804 = *vOneOfType2804
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType2804.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType2804.ObjectType_
							return nil
						}
					}
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType2805 := new(UnConfigureServerDetails)
	if err := json.Unmarshal(b, vOneOfType2805); err == nil {
		if vOneOfType2805.ObjectType_ != nil && "lifecycle.v4.config.UnConfigureServerDetails" == *vOneOfType2805.ObjectType_ {
			if nil == p.oneOfType2805 {
				p.oneOfType2805 = new(UnConfigureServerDetails)
			}
			*p.oneOfType2805 = *vOneOfType2805
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2805.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2805.ObjectType_
			return nil
		}
	}
	vOneOfType2802 := new(ConfigureServerDetails)
	if err := json.Unmarshal(b, vOneOfType2802); err == nil {
		if vOneOfType2802.ObjectType_ != nil && "lifecycle.v4.config.ConfigureServerDetails" == *vOneOfType2802.ObjectType_ {
			if nil == p.oneOfType2802 {
				p.oneOfType2802 = new(ConfigureServerDetails)
			}
			*p.oneOfType2802 = *vOneOfType2802
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2802.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2802.ObjectType_
			return nil
		}
	}
	vOneOfType2803 := new(PreClusterConfigDetails)
	if err := json.Unmarshal(b, vOneOfType2803); err == nil {
		if vOneOfType2803.ObjectType_ != nil && "lifecycle.v4.config.PreClusterConfigDetails" == *vOneOfType2803.ObjectType_ {
			if nil == p.oneOfType2803 {
				p.oneOfType2803 = new(PreClusterConfigDetails)
			}
			*p.oneOfType2803 = *vOneOfType2803
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2803.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2803.ObjectType_
			return nil
		}
	}
	vOneOfType2804 := new(PostClusterConfigDetails)
	if err := json.Unmarshal(b, vOneOfType2804); err == nil {
		if vOneOfType2804.ObjectType_ != nil && "lifecycle.v4.config.PostClusterConfigDetails" == *vOneOfType2804.ObjectType_ {
			if nil == p.oneOfType2804 {
				p.oneOfType2804 = new(PostClusterConfigDetails)
			}
			*p.oneOfType2804 = *vOneOfType2804
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2804.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2804.ObjectType_
			return nil
		}
	}
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfConfigureNodeSpecConfiguration"))
}

func (p *OneOfConfigureNodeSpecConfiguration) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType2805 != nil && *p.oneOfType2805.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2805)
	}
	if p.oneOfType2802 != nil && *p.oneOfType2802.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2802)
	}
	if p.oneOfType2803 != nil && *p.oneOfType2803.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2803)
	}
	if p.oneOfType2804 != nil && *p.oneOfType2804.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2804)
	}
	return nil, errors.New("No value to marshal for OneOfConfigureNodeSpecConfiguration")
}

type OneOfCreateClaimTokenApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType1000 *import3.TaskReference `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfCreateClaimTokenApiResponseData() *OneOfCreateClaimTokenApiResponseData {
	p := new(OneOfCreateClaimTokenApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateClaimTokenApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateClaimTokenApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.TaskReference:
		if nil == p.oneOfType1000 {
			p.oneOfType1000 = new(import3.TaskReference)
		}
		*p.oneOfType1000 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType1000.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType1000.ObjectType_
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

func (p *OneOfCreateClaimTokenApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType1000 != nil && *p.oneOfType1000.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1000
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfCreateClaimTokenApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType1000 := new(import3.TaskReference)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType1000)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType1000.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType1000.ObjectType_ {
							if nil == p.oneOfType1000 {
								p.oneOfType1000 = new(import3.TaskReference)
							}
							*p.oneOfType1000 = *vOneOfType1000
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType1000.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType1000.ObjectType_
							return nil
						}
					}
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import4.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType1000 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType1000); err == nil {
		if vOneOfType1000.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType1000.ObjectType_ {
			if nil == p.oneOfType1000 {
				p.oneOfType1000 = new(import3.TaskReference)
			}
			*p.oneOfType1000 = *vOneOfType1000
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType1000.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType1000.ObjectType_
			return nil
		}
	}
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateClaimTokenApiResponseData"))
}

func (p *OneOfCreateClaimTokenApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType1000 != nil && *p.oneOfType1000.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1000)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfCreateClaimTokenApiResponseData")
}

type OneOfRefreshNodeApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType1018 *import3.TaskReference `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfRefreshNodeApiResponseData() *OneOfRefreshNodeApiResponseData {
	p := new(OneOfRefreshNodeApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfRefreshNodeApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfRefreshNodeApiResponseData is nil"))
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
	case import3.TaskReference:
		if nil == p.oneOfType1018 {
			p.oneOfType1018 = new(import3.TaskReference)
		}
		*p.oneOfType1018 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType1018.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType1018.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfRefreshNodeApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType1018 != nil && *p.oneOfType1018.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1018
	}
	return nil
}

func (p *OneOfRefreshNodeApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import4.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType1018 := new(import3.TaskReference)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType1018)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType1018.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType1018.ObjectType_ {
							if nil == p.oneOfType1018 {
								p.oneOfType1018 = new(import3.TaskReference)
							}
							*p.oneOfType1018 = *vOneOfType1018
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType1018.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType1018.ObjectType_
							return nil
						}
					}
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType1018 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType1018); err == nil {
		if vOneOfType1018.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType1018.ObjectType_ {
			if nil == p.oneOfType1018 {
				p.oneOfType1018 = new(import3.TaskReference)
			}
			*p.oneOfType1018 = *vOneOfType1018
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType1018.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType1018.ObjectType_
			return nil
		}
	}
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfRefreshNodeApiResponseData"))
}

func (p *OneOfRefreshNodeApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType1018 != nil && *p.oneOfType1018.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1018)
	}
	return nil, errors.New("No value to marshal for OneOfRefreshNodeApiResponseData")
}

type OneOfGetServerIdentityPoolApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType1148 *ServerIdentityPool    `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfGetServerIdentityPoolApiResponseData() *OneOfGetServerIdentityPoolApiResponseData {
	p := new(OneOfGetServerIdentityPoolApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetServerIdentityPoolApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetServerIdentityPoolApiResponseData is nil"))
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
	case ServerIdentityPool:
		if nil == p.oneOfType1148 {
			p.oneOfType1148 = new(ServerIdentityPool)
		}
		*p.oneOfType1148 = v.(ServerIdentityPool)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType1148.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType1148.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfGetServerIdentityPoolApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType1148 != nil && *p.oneOfType1148.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1148
	}
	return nil
}

func (p *OneOfGetServerIdentityPoolApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import4.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType1148 := new(ServerIdentityPool)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType1148)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType1148.ObjectType_ != nil && "lifecycle.v4.config.ServerIdentityPool" == *vOneOfType1148.ObjectType_ {
							if nil == p.oneOfType1148 {
								p.oneOfType1148 = new(ServerIdentityPool)
							}
							*p.oneOfType1148 = *vOneOfType1148
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType1148.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType1148.ObjectType_
							return nil
						}
					}
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType1148 := new(ServerIdentityPool)
	if err := json.Unmarshal(b, vOneOfType1148); err == nil {
		if vOneOfType1148.ObjectType_ != nil && "lifecycle.v4.config.ServerIdentityPool" == *vOneOfType1148.ObjectType_ {
			if nil == p.oneOfType1148 {
				p.oneOfType1148 = new(ServerIdentityPool)
			}
			*p.oneOfType1148 = *vOneOfType1148
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType1148.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType1148.ObjectType_
			return nil
		}
	}
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetServerIdentityPoolApiResponseData"))
}

func (p *OneOfGetServerIdentityPoolApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType1148 != nil && *p.oneOfType1148.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1148)
	}
	return nil, errors.New("No value to marshal for OneOfGetServerIdentityPoolApiResponseData")
}

type OneOfAosInstallationDetailsAosImageDetails struct {
	Discriminator *string               `json:"-"`
	ObjectType_   *string               `json:"-"`
	oneOfType2605 *LocalAOSImageDetails `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfAosInstallationDetailsAosImageDetails() *OneOfAosInstallationDetailsAosImageDetails {
	p := new(OneOfAosInstallationDetailsAosImageDetails)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfAosInstallationDetailsAosImageDetails) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfAosInstallationDetailsAosImageDetails is nil"))
	}
	switch v.(type) {
	case LocalAOSImageDetails:
		if nil == p.oneOfType2605 {
			p.oneOfType2605 = new(LocalAOSImageDetails)
		}
		*p.oneOfType2605 = v.(LocalAOSImageDetails)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2605.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2605.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfAosInstallationDetailsAosImageDetails) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType2605 != nil && *p.oneOfType2605.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2605
	}
	return nil
}

func (p *OneOfAosInstallationDetailsAosImageDetails) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType2605 := new(LocalAOSImageDetails)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType2605)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType2605.ObjectType_ != nil && "lifecycle.v4.config.LocalAOSImageDetails" == *vOneOfType2605.ObjectType_ {
							if nil == p.oneOfType2605 {
								p.oneOfType2605 = new(LocalAOSImageDetails)
							}
							*p.oneOfType2605 = *vOneOfType2605
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType2605.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType2605.ObjectType_
							return nil
						}
					}
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType2605 := new(LocalAOSImageDetails)
	if err := json.Unmarshal(b, vOneOfType2605); err == nil {
		if vOneOfType2605.ObjectType_ != nil && "lifecycle.v4.config.LocalAOSImageDetails" == *vOneOfType2605.ObjectType_ {
			if nil == p.oneOfType2605 {
				p.oneOfType2605 = new(LocalAOSImageDetails)
			}
			*p.oneOfType2605 = *vOneOfType2605
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2605.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2605.ObjectType_
			return nil
		}
	}
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfAosInstallationDetailsAosImageDetails"))
}

func (p *OneOfAosInstallationDetailsAosImageDetails) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType2605 != nil && *p.oneOfType2605.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2605)
	}
	return nil, errors.New("No value to marshal for OneOfAosInstallationDetailsAosImageDetails")
}

type OneOfCreatePatchedImageApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType1000 *import3.TaskReference `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfCreatePatchedImageApiResponseData() *OneOfCreatePatchedImageApiResponseData {
	p := new(OneOfCreatePatchedImageApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreatePatchedImageApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreatePatchedImageApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.TaskReference:
		if nil == p.oneOfType1000 {
			p.oneOfType1000 = new(import3.TaskReference)
		}
		*p.oneOfType1000 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType1000.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType1000.ObjectType_
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

func (p *OneOfCreatePatchedImageApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType1000 != nil && *p.oneOfType1000.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1000
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfCreatePatchedImageApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType1000 := new(import3.TaskReference)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType1000)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType1000.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType1000.ObjectType_ {
							if nil == p.oneOfType1000 {
								p.oneOfType1000 = new(import3.TaskReference)
							}
							*p.oneOfType1000 = *vOneOfType1000
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType1000.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType1000.ObjectType_
							return nil
						}
					}
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import4.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType1000 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType1000); err == nil {
		if vOneOfType1000.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType1000.ObjectType_ {
			if nil == p.oneOfType1000 {
				p.oneOfType1000 = new(import3.TaskReference)
			}
			*p.oneOfType1000 = *vOneOfType1000
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType1000.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType1000.ObjectType_
			return nil
		}
	}
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreatePatchedImageApiResponseData"))
}

func (p *OneOfCreatePatchedImageApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType1000 != nil && *p.oneOfType1000.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1000)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfCreatePatchedImageApiResponseData")
}

type OneOfConnectionDetailAuth struct {
	Discriminator *string     `json:"-"`
	ObjectType_   *string     `json:"-"`
	oneOfType2029 *ApiKeyAuth `json:"-"`
	oneOfType2030 *BasicAuth  `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfConnectionDetailAuth() *OneOfConnectionDetailAuth {
	p := new(OneOfConnectionDetailAuth)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfConnectionDetailAuth) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfConnectionDetailAuth is nil"))
	}
	switch v.(type) {
	case ApiKeyAuth:
		if nil == p.oneOfType2029 {
			p.oneOfType2029 = new(ApiKeyAuth)
		}
		*p.oneOfType2029 = v.(ApiKeyAuth)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2029.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2029.ObjectType_
	case BasicAuth:
		if nil == p.oneOfType2030 {
			p.oneOfType2030 = new(BasicAuth)
		}
		*p.oneOfType2030 = v.(BasicAuth)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2030.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2030.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfConnectionDetailAuth) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType2029 != nil && *p.oneOfType2029.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2029
	}
	if p.oneOfType2030 != nil && *p.oneOfType2030.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2030
	}
	return nil
}

func (p *OneOfConnectionDetailAuth) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType2029 := new(ApiKeyAuth)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType2029)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType2029.ObjectType_ != nil && "lifecycle.v4.config.ApiKeyAuth" == *vOneOfType2029.ObjectType_ {
							if nil == p.oneOfType2029 {
								p.oneOfType2029 = new(ApiKeyAuth)
							}
							*p.oneOfType2029 = *vOneOfType2029
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType2029.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType2029.ObjectType_
							return nil
						}
					}
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType2030 := new(BasicAuth)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType2030)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType2030.ObjectType_ != nil && "lifecycle.v4.config.BasicAuth" == *vOneOfType2030.ObjectType_ {
							if nil == p.oneOfType2030 {
								p.oneOfType2030 = new(BasicAuth)
							}
							*p.oneOfType2030 = *vOneOfType2030
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType2030.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType2030.ObjectType_
							return nil
						}
					}
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType2029 := new(ApiKeyAuth)
	if err := json.Unmarshal(b, vOneOfType2029); err == nil {
		if vOneOfType2029.ObjectType_ != nil && "lifecycle.v4.config.ApiKeyAuth" == *vOneOfType2029.ObjectType_ {
			if nil == p.oneOfType2029 {
				p.oneOfType2029 = new(ApiKeyAuth)
			}
			*p.oneOfType2029 = *vOneOfType2029
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2029.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2029.ObjectType_
			return nil
		}
	}
	vOneOfType2030 := new(BasicAuth)
	if err := json.Unmarshal(b, vOneOfType2030); err == nil {
		if vOneOfType2030.ObjectType_ != nil && "lifecycle.v4.config.BasicAuth" == *vOneOfType2030.ObjectType_ {
			if nil == p.oneOfType2030 {
				p.oneOfType2030 = new(BasicAuth)
			}
			*p.oneOfType2030 = *vOneOfType2030
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2030.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2030.ObjectType_
			return nil
		}
	}
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfConnectionDetailAuth"))
}

func (p *OneOfConnectionDetailAuth) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType2029 != nil && *p.oneOfType2029.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2029)
	}
	if p.oneOfType2030 != nil && *p.oneOfType2030.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2030)
	}
	return nil, errors.New("No value to marshal for OneOfConnectionDetailAuth")
}

type OneOfCreateImageApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType4001 *import3.TaskReference `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfCreateImageApiResponseData() *OneOfCreateImageApiResponseData {
	p := new(OneOfCreateImageApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateImageApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateImageApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.TaskReference:
		if nil == p.oneOfType4001 {
			p.oneOfType4001 = new(import3.TaskReference)
		}
		*p.oneOfType4001 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType4001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType4001.ObjectType_
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

func (p *OneOfCreateImageApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType4001 != nil && *p.oneOfType4001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType4001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfCreateImageApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType4001 := new(import3.TaskReference)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType4001)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType4001.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType4001.ObjectType_ {
							if nil == p.oneOfType4001 {
								p.oneOfType4001 = new(import3.TaskReference)
							}
							*p.oneOfType4001 = *vOneOfType4001
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType4001.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType4001.ObjectType_
							return nil
						}
					}
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import4.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType4001 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType4001); err == nil {
		if vOneOfType4001.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType4001.ObjectType_ {
			if nil == p.oneOfType4001 {
				p.oneOfType4001 = new(import3.TaskReference)
			}
			*p.oneOfType4001 = *vOneOfType4001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType4001.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType4001.ObjectType_
			return nil
		}
	}
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateImageApiResponseData"))
}

func (p *OneOfCreateImageApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType4001 != nil && *p.oneOfType4001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType4001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfCreateImageApiResponseData")
}

type OneOfDeleteConnectionApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType1024 *import3.TaskReference `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfDeleteConnectionApiResponseData() *OneOfDeleteConnectionApiResponseData {
	p := new(OneOfDeleteConnectionApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteConnectionApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteConnectionApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.TaskReference:
		if nil == p.oneOfType1024 {
			p.oneOfType1024 = new(import3.TaskReference)
		}
		*p.oneOfType1024 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType1024.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType1024.ObjectType_
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

func (p *OneOfDeleteConnectionApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType1024 != nil && *p.oneOfType1024.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1024
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDeleteConnectionApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType1024 := new(import3.TaskReference)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType1024)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType1024.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType1024.ObjectType_ {
							if nil == p.oneOfType1024 {
								p.oneOfType1024 = new(import3.TaskReference)
							}
							*p.oneOfType1024 = *vOneOfType1024
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType1024.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType1024.ObjectType_
							return nil
						}
					}
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import4.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType1024 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType1024); err == nil {
		if vOneOfType1024.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType1024.ObjectType_ {
			if nil == p.oneOfType1024 {
				p.oneOfType1024 = new(import3.TaskReference)
			}
			*p.oneOfType1024 = *vOneOfType1024
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType1024.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType1024.ObjectType_
			return nil
		}
	}
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteConnectionApiResponseData"))
}

func (p *OneOfDeleteConnectionApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType1024 != nil && *p.oneOfType1024.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1024)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteConnectionApiResponseData")
}

type OneOfConfigureServerDetailsBmcIPAddressingConfig struct {
	Discriminator *string               `json:"-"`
	ObjectType_   *string               `json:"-"`
	oneOfType3022 *UnifiedBmcIPConfig   `json:"-"`
	oneOfType3021 *InOutBandBmcIPConfig `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfConfigureServerDetailsBmcIPAddressingConfig() *OneOfConfigureServerDetailsBmcIPAddressingConfig {
	p := new(OneOfConfigureServerDetailsBmcIPAddressingConfig)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfConfigureServerDetailsBmcIPAddressingConfig) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfConfigureServerDetailsBmcIPAddressingConfig is nil"))
	}
	switch v.(type) {
	case UnifiedBmcIPConfig:
		if nil == p.oneOfType3022 {
			p.oneOfType3022 = new(UnifiedBmcIPConfig)
		}
		*p.oneOfType3022 = v.(UnifiedBmcIPConfig)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType3022.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType3022.ObjectType_
	case InOutBandBmcIPConfig:
		if nil == p.oneOfType3021 {
			p.oneOfType3021 = new(InOutBandBmcIPConfig)
		}
		*p.oneOfType3021 = v.(InOutBandBmcIPConfig)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType3021.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType3021.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfConfigureServerDetailsBmcIPAddressingConfig) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType3022 != nil && *p.oneOfType3022.ObjectType_ == *p.Discriminator {
		return *p.oneOfType3022
	}
	if p.oneOfType3021 != nil && *p.oneOfType3021.ObjectType_ == *p.Discriminator {
		return *p.oneOfType3021
	}
	return nil
}

func (p *OneOfConfigureServerDetailsBmcIPAddressingConfig) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType3022 := new(UnifiedBmcIPConfig)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType3022)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType3022.ObjectType_ != nil && "lifecycle.v4.config.UnifiedBmcIPConfig" == *vOneOfType3022.ObjectType_ {
							if nil == p.oneOfType3022 {
								p.oneOfType3022 = new(UnifiedBmcIPConfig)
							}
							*p.oneOfType3022 = *vOneOfType3022
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType3022.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType3022.ObjectType_
							return nil
						}
					}
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType3021 := new(InOutBandBmcIPConfig)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType3021)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType3021.ObjectType_ != nil && "lifecycle.v4.config.InOutBandBmcIPConfig" == *vOneOfType3021.ObjectType_ {
							if nil == p.oneOfType3021 {
								p.oneOfType3021 = new(InOutBandBmcIPConfig)
							}
							*p.oneOfType3021 = *vOneOfType3021
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType3021.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType3021.ObjectType_
							return nil
						}
					}
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType3022 := new(UnifiedBmcIPConfig)
	if err := json.Unmarshal(b, vOneOfType3022); err == nil {
		if vOneOfType3022.ObjectType_ != nil && "lifecycle.v4.config.UnifiedBmcIPConfig" == *vOneOfType3022.ObjectType_ {
			if nil == p.oneOfType3022 {
				p.oneOfType3022 = new(UnifiedBmcIPConfig)
			}
			*p.oneOfType3022 = *vOneOfType3022
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType3022.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType3022.ObjectType_
			return nil
		}
	}
	vOneOfType3021 := new(InOutBandBmcIPConfig)
	if err := json.Unmarshal(b, vOneOfType3021); err == nil {
		if vOneOfType3021.ObjectType_ != nil && "lifecycle.v4.config.InOutBandBmcIPConfig" == *vOneOfType3021.ObjectType_ {
			if nil == p.oneOfType3021 {
				p.oneOfType3021 = new(InOutBandBmcIPConfig)
			}
			*p.oneOfType3021 = *vOneOfType3021
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType3021.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType3021.ObjectType_
			return nil
		}
	}
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfConfigureServerDetailsBmcIPAddressingConfig"))
}

func (p *OneOfConfigureServerDetailsBmcIPAddressingConfig) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType3022 != nil && *p.oneOfType3022.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType3022)
	}
	if p.oneOfType3021 != nil && *p.oneOfType3021.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType3021)
	}
	return nil, errors.New("No value to marshal for OneOfConfigureServerDetailsBmcIPAddressingConfig")
}

type OneOfListNodesByClaimTokenApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType1014 []ClaimTokenNode       `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfListNodesByClaimTokenApiResponseData() *OneOfListNodesByClaimTokenApiResponseData {
	p := new(OneOfListNodesByClaimTokenApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListNodesByClaimTokenApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListNodesByClaimTokenApiResponseData is nil"))
	}
	switch v.(type) {
	case []ClaimTokenNode:
		p.oneOfType1014 = v.([]ClaimTokenNode)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<lifecycle.v4.config.ClaimTokenNode>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<lifecycle.v4.config.ClaimTokenNode>"
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

func (p *OneOfListNodesByClaimTokenApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if "List<lifecycle.v4.config.ClaimTokenNode>" == *p.Discriminator {
		return p.oneOfType1014
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfListNodesByClaimTokenApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["List<lifecycle.v4.config.ClaimTokenNode>"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType1014 := new([]ClaimTokenNode)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType1014)
					if unmarshalErr == nil {
						// For arrays, verify the array item ObjectType matches
						if vOneOfType1014 == nil || len(*vOneOfType1014) == 0 || ((*vOneOfType1014)[0].ObjectType_ != nil && "lifecycle.v4.config.ClaimTokenNode" == *((*vOneOfType1014)[0].ObjectType_)) {
							p.oneOfType1014 = *vOneOfType1014
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = "List<lifecycle.v4.config.ClaimTokenNode>"
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = "List<lifecycle.v4.config.ClaimTokenNode>"
							return nil
						}
					}
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import4.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType1014 := new([]ClaimTokenNode)
	if err := json.Unmarshal(b, vOneOfType1014); err == nil {
		if len(*vOneOfType1014) == 0 || (vOneOfType1014 != nil && (*vOneOfType1014)[0].ObjectType_ != nil && "lifecycle.v4.config.ClaimTokenNode" == *((*vOneOfType1014)[0].ObjectType_)) {
			p.oneOfType1014 = *vOneOfType1014
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<lifecycle.v4.config.ClaimTokenNode>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<lifecycle.v4.config.ClaimTokenNode>"
			return nil
		}
	}
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListNodesByClaimTokenApiResponseData"))
}

func (p *OneOfListNodesByClaimTokenApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if "List<lifecycle.v4.config.ClaimTokenNode>" == *p.Discriminator {
		return json.Marshal(p.oneOfType1014)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfListNodesByClaimTokenApiResponseData")
}

type OneOfDeletePatchedImageApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType1007 *import3.TaskReference `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfDeletePatchedImageApiResponseData() *OneOfDeletePatchedImageApiResponseData {
	p := new(OneOfDeletePatchedImageApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeletePatchedImageApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeletePatchedImageApiResponseData is nil"))
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
	case import3.TaskReference:
		if nil == p.oneOfType1007 {
			p.oneOfType1007 = new(import3.TaskReference)
		}
		*p.oneOfType1007 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType1007.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType1007.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfDeletePatchedImageApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType1007 != nil && *p.oneOfType1007.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1007
	}
	return nil
}

func (p *OneOfDeletePatchedImageApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import4.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType1007 := new(import3.TaskReference)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType1007)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType1007.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType1007.ObjectType_ {
							if nil == p.oneOfType1007 {
								p.oneOfType1007 = new(import3.TaskReference)
							}
							*p.oneOfType1007 = *vOneOfType1007
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType1007.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType1007.ObjectType_
							return nil
						}
					}
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType1007 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType1007); err == nil {
		if vOneOfType1007.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType1007.ObjectType_ {
			if nil == p.oneOfType1007 {
				p.oneOfType1007 = new(import3.TaskReference)
			}
			*p.oneOfType1007 = *vOneOfType1007
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType1007.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType1007.ObjectType_
			return nil
		}
	}
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeletePatchedImageApiResponseData"))
}

func (p *OneOfDeletePatchedImageApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType1007 != nil && *p.oneOfType1007.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1007)
	}
	return nil, errors.New("No value to marshal for OneOfDeletePatchedImageApiResponseData")
}

type OneOfListHardwareProvidersApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType1002 []HardwareProvider     `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfListHardwareProvidersApiResponseData() *OneOfListHardwareProvidersApiResponseData {
	p := new(OneOfListHardwareProvidersApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListHardwareProvidersApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListHardwareProvidersApiResponseData is nil"))
	}
	switch v.(type) {
	case []HardwareProvider:
		p.oneOfType1002 = v.([]HardwareProvider)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<lifecycle.v4.config.HardwareProvider>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<lifecycle.v4.config.HardwareProvider>"
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

func (p *OneOfListHardwareProvidersApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if "List<lifecycle.v4.config.HardwareProvider>" == *p.Discriminator {
		return p.oneOfType1002
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfListHardwareProvidersApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["List<lifecycle.v4.config.HardwareProvider>"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType1002 := new([]HardwareProvider)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType1002)
					if unmarshalErr == nil {
						// For arrays, verify the array item ObjectType matches
						if vOneOfType1002 == nil || len(*vOneOfType1002) == 0 || ((*vOneOfType1002)[0].ObjectType_ != nil && "lifecycle.v4.config.HardwareProvider" == *((*vOneOfType1002)[0].ObjectType_)) {
							p.oneOfType1002 = *vOneOfType1002
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = "List<lifecycle.v4.config.HardwareProvider>"
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = "List<lifecycle.v4.config.HardwareProvider>"
							return nil
						}
					}
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import4.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType1002 := new([]HardwareProvider)
	if err := json.Unmarshal(b, vOneOfType1002); err == nil {
		if len(*vOneOfType1002) == 0 || (vOneOfType1002 != nil && (*vOneOfType1002)[0].ObjectType_ != nil && "lifecycle.v4.config.HardwareProvider" == *((*vOneOfType1002)[0].ObjectType_)) {
			p.oneOfType1002 = *vOneOfType1002
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<lifecycle.v4.config.HardwareProvider>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<lifecycle.v4.config.HardwareProvider>"
			return nil
		}
	}
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListHardwareProvidersApiResponseData"))
}

func (p *OneOfListHardwareProvidersApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if "List<lifecycle.v4.config.HardwareProvider>" == *p.Discriminator {
		return json.Marshal(p.oneOfType1002)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfListHardwareProvidersApiResponseData")
}

type OneOfUpdateClaimTokenApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType1009 *import3.TaskReference `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfUpdateClaimTokenApiResponseData() *OneOfUpdateClaimTokenApiResponseData {
	p := new(OneOfUpdateClaimTokenApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateClaimTokenApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateClaimTokenApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.TaskReference:
		if nil == p.oneOfType1009 {
			p.oneOfType1009 = new(import3.TaskReference)
		}
		*p.oneOfType1009 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType1009.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType1009.ObjectType_
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

func (p *OneOfUpdateClaimTokenApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType1009 != nil && *p.oneOfType1009.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1009
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfUpdateClaimTokenApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType1009 := new(import3.TaskReference)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType1009)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType1009.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType1009.ObjectType_ {
							if nil == p.oneOfType1009 {
								p.oneOfType1009 = new(import3.TaskReference)
							}
							*p.oneOfType1009 = *vOneOfType1009
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType1009.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType1009.ObjectType_
							return nil
						}
					}
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import4.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType1009 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType1009); err == nil {
		if vOneOfType1009.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType1009.ObjectType_ {
			if nil == p.oneOfType1009 {
				p.oneOfType1009 = new(import3.TaskReference)
			}
			*p.oneOfType1009 = *vOneOfType1009
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType1009.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType1009.ObjectType_
			return nil
		}
	}
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateClaimTokenApiResponseData"))
}

func (p *OneOfUpdateClaimTokenApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType1009 != nil && *p.oneOfType1009.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1009)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateClaimTokenApiResponseData")
}

type OneOfRefreshConnectionApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType1052 *import3.TaskReference `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfRefreshConnectionApiResponseData() *OneOfRefreshConnectionApiResponseData {
	p := new(OneOfRefreshConnectionApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfRefreshConnectionApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfRefreshConnectionApiResponseData is nil"))
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
	case import3.TaskReference:
		if nil == p.oneOfType1052 {
			p.oneOfType1052 = new(import3.TaskReference)
		}
		*p.oneOfType1052 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType1052.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType1052.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfRefreshConnectionApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType1052 != nil && *p.oneOfType1052.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1052
	}
	return nil
}

func (p *OneOfRefreshConnectionApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import4.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType1052 := new(import3.TaskReference)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType1052)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType1052.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType1052.ObjectType_ {
							if nil == p.oneOfType1052 {
								p.oneOfType1052 = new(import3.TaskReference)
							}
							*p.oneOfType1052 = *vOneOfType1052
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType1052.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType1052.ObjectType_
							return nil
						}
					}
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType1052 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType1052); err == nil {
		if vOneOfType1052.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType1052.ObjectType_ {
			if nil == p.oneOfType1052 {
				p.oneOfType1052 = new(import3.TaskReference)
			}
			*p.oneOfType1052 = *vOneOfType1052
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType1052.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType1052.ObjectType_
			return nil
		}
	}
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfRefreshConnectionApiResponseData"))
}

func (p *OneOfRefreshConnectionApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType1052 != nil && *p.oneOfType1052.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1052)
	}
	return nil, errors.New("No value to marshal for OneOfRefreshConnectionApiResponseData")
}

type OneOfGetConnectionApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType1017 *Connection            `json:"-"`
	oneOfType401  *ConnectionProjection  `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfGetConnectionApiResponseData() *OneOfGetConnectionApiResponseData {
	p := new(OneOfGetConnectionApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetConnectionApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetConnectionApiResponseData is nil"))
	}
	switch v.(type) {
	case Connection:
		if nil == p.oneOfType1017 {
			p.oneOfType1017 = new(Connection)
		}
		*p.oneOfType1017 = v.(Connection)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType1017.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType1017.ObjectType_
	case ConnectionProjection:
		if nil == p.oneOfType401 {
			p.oneOfType401 = new(ConnectionProjection)
		}
		*p.oneOfType401 = v.(ConnectionProjection)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType401.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType401.ObjectType_
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

func (p *OneOfGetConnectionApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType1017 != nil && *p.oneOfType1017.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1017
	}
	if p.oneOfType401 != nil && *p.oneOfType401.ObjectType_ == *p.Discriminator {
		return *p.oneOfType401
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetConnectionApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType1017 := new(Connection)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType1017)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType1017.ObjectType_ != nil && "lifecycle.v4.config.Connection" == *vOneOfType1017.ObjectType_ {
							if nil == p.oneOfType1017 {
								p.oneOfType1017 = new(Connection)
							}
							*p.oneOfType1017 = *vOneOfType1017
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType1017.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType1017.ObjectType_
							return nil
						}
					}
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType401 := new(ConnectionProjection)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType401)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType401.ObjectType_ != nil && "lifecycle.v4.config.ConnectionProjection" == *vOneOfType401.ObjectType_ {
							if nil == p.oneOfType401 {
								p.oneOfType401 = new(ConnectionProjection)
							}
							*p.oneOfType401 = *vOneOfType401
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType401.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType401.ObjectType_
							return nil
						}
					}
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import4.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType1017 := new(Connection)
	if err := json.Unmarshal(b, vOneOfType1017); err == nil {
		if vOneOfType1017.ObjectType_ != nil && "lifecycle.v4.config.Connection" == *vOneOfType1017.ObjectType_ {
			if nil == p.oneOfType1017 {
				p.oneOfType1017 = new(Connection)
			}
			*p.oneOfType1017 = *vOneOfType1017
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType1017.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType1017.ObjectType_
			return nil
		}
	}
	vOneOfType401 := new(ConnectionProjection)
	if err := json.Unmarshal(b, vOneOfType401); err == nil {
		if vOneOfType401.ObjectType_ != nil && "lifecycle.v4.config.ConnectionProjection" == *vOneOfType401.ObjectType_ {
			if nil == p.oneOfType401 {
				p.oneOfType401 = new(ConnectionProjection)
			}
			*p.oneOfType401 = *vOneOfType401
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType401.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType401.ObjectType_
			return nil
		}
	}
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetConnectionApiResponseData"))
}

func (p *OneOfGetConnectionApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType1017 != nil && *p.oneOfType1017.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1017)
	}
	if p.oneOfType401 != nil && *p.oneOfType401.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType401)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetConnectionApiResponseData")
}

type OneOfCreateNodeApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType1016 *import3.TaskReference `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfCreateNodeApiResponseData() *OneOfCreateNodeApiResponseData {
	p := new(OneOfCreateNodeApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateNodeApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateNodeApiResponseData is nil"))
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
	case import3.TaskReference:
		if nil == p.oneOfType1016 {
			p.oneOfType1016 = new(import3.TaskReference)
		}
		*p.oneOfType1016 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType1016.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType1016.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfCreateNodeApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType1016 != nil && *p.oneOfType1016.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1016
	}
	return nil
}

func (p *OneOfCreateNodeApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import4.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType1016 := new(import3.TaskReference)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType1016)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType1016.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType1016.ObjectType_ {
							if nil == p.oneOfType1016 {
								p.oneOfType1016 = new(import3.TaskReference)
							}
							*p.oneOfType1016 = *vOneOfType1016
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType1016.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType1016.ObjectType_
							return nil
						}
					}
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType1016 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType1016); err == nil {
		if vOneOfType1016.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType1016.ObjectType_ {
			if nil == p.oneOfType1016 {
				p.oneOfType1016 = new(import3.TaskReference)
			}
			*p.oneOfType1016 = *vOneOfType1016
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType1016.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType1016.ObjectType_
			return nil
		}
	}
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateNodeApiResponseData"))
}

func (p *OneOfCreateNodeApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType1016 != nil && *p.oneOfType1016.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1016)
	}
	return nil, errors.New("No value to marshal for OneOfCreateNodeApiResponseData")
}

type OneOfCreateConnectionApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType1010 *import3.TaskReference `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfCreateConnectionApiResponseData() *OneOfCreateConnectionApiResponseData {
	p := new(OneOfCreateConnectionApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateConnectionApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateConnectionApiResponseData is nil"))
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
	case import3.TaskReference:
		if nil == p.oneOfType1010 {
			p.oneOfType1010 = new(import3.TaskReference)
		}
		*p.oneOfType1010 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType1010.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType1010.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfCreateConnectionApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType1010 != nil && *p.oneOfType1010.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1010
	}
	return nil
}

func (p *OneOfCreateConnectionApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import4.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType1010 := new(import3.TaskReference)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType1010)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType1010.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType1010.ObjectType_ {
							if nil == p.oneOfType1010 {
								p.oneOfType1010 = new(import3.TaskReference)
							}
							*p.oneOfType1010 = *vOneOfType1010
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType1010.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType1010.ObjectType_
							return nil
						}
					}
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType1010 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType1010); err == nil {
		if vOneOfType1010.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType1010.ObjectType_ {
			if nil == p.oneOfType1010 {
				p.oneOfType1010 = new(import3.TaskReference)
			}
			*p.oneOfType1010 = *vOneOfType1010
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType1010.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType1010.ObjectType_
			return nil
		}
	}
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateConnectionApiResponseData"))
}

func (p *OneOfCreateConnectionApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType1010 != nil && *p.oneOfType1010.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1010)
	}
	return nil, errors.New("No value to marshal for OneOfCreateConnectionApiResponseData")
}

type OneOfGetMacPoolApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType1048 *MacPool               `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfGetMacPoolApiResponseData() *OneOfGetMacPoolApiResponseData {
	p := new(OneOfGetMacPoolApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetMacPoolApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetMacPoolApiResponseData is nil"))
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
	case MacPool:
		if nil == p.oneOfType1048 {
			p.oneOfType1048 = new(MacPool)
		}
		*p.oneOfType1048 = v.(MacPool)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType1048.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType1048.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfGetMacPoolApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType1048 != nil && *p.oneOfType1048.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1048
	}
	return nil
}

func (p *OneOfGetMacPoolApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import4.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType1048 := new(MacPool)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType1048)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType1048.ObjectType_ != nil && "lifecycle.v4.config.MacPool" == *vOneOfType1048.ObjectType_ {
							if nil == p.oneOfType1048 {
								p.oneOfType1048 = new(MacPool)
							}
							*p.oneOfType1048 = *vOneOfType1048
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType1048.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType1048.ObjectType_
							return nil
						}
					}
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType1048 := new(MacPool)
	if err := json.Unmarshal(b, vOneOfType1048); err == nil {
		if vOneOfType1048.ObjectType_ != nil && "lifecycle.v4.config.MacPool" == *vOneOfType1048.ObjectType_ {
			if nil == p.oneOfType1048 {
				p.oneOfType1048 = new(MacPool)
			}
			*p.oneOfType1048 = *vOneOfType1048
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType1048.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType1048.ObjectType_
			return nil
		}
	}
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetMacPoolApiResponseData"))
}

func (p *OneOfGetMacPoolApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType1048 != nil && *p.oneOfType1048.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1048)
	}
	return nil, errors.New("No value to marshal for OneOfGetMacPoolApiResponseData")
}

type OneOfUpdateNodeApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType1007 *import3.TaskReference `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfUpdateNodeApiResponseData() *OneOfUpdateNodeApiResponseData {
	p := new(OneOfUpdateNodeApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateNodeApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateNodeApiResponseData is nil"))
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
	case import3.TaskReference:
		if nil == p.oneOfType1007 {
			p.oneOfType1007 = new(import3.TaskReference)
		}
		*p.oneOfType1007 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType1007.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType1007.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfUpdateNodeApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType1007 != nil && *p.oneOfType1007.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1007
	}
	return nil
}

func (p *OneOfUpdateNodeApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import4.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType1007 := new(import3.TaskReference)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType1007)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType1007.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType1007.ObjectType_ {
							if nil == p.oneOfType1007 {
								p.oneOfType1007 = new(import3.TaskReference)
							}
							*p.oneOfType1007 = *vOneOfType1007
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType1007.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType1007.ObjectType_
							return nil
						}
					}
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType1007 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType1007); err == nil {
		if vOneOfType1007.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType1007.ObjectType_ {
			if nil == p.oneOfType1007 {
				p.oneOfType1007 = new(import3.TaskReference)
			}
			*p.oneOfType1007 = *vOneOfType1007
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType1007.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType1007.ObjectType_
			return nil
		}
	}
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateNodeApiResponseData"))
}

func (p *OneOfUpdateNodeApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType1007 != nil && *p.oneOfType1007.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1007)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateNodeApiResponseData")
}

type OneOfBmcIPConfigIpv4 struct {
	Discriminator *string         `json:"-"`
	ObjectType_   *string         `json:"-"`
	oneOfType5135 *IPPoolConfig   `json:"-"`
	oneOfType5136 *StaticIPConfig `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfBmcIPConfigIpv4() *OneOfBmcIPConfigIpv4 {
	p := new(OneOfBmcIPConfigIpv4)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfBmcIPConfigIpv4) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfBmcIPConfigIpv4 is nil"))
	}
	switch v.(type) {
	case IPPoolConfig:
		if nil == p.oneOfType5135 {
			p.oneOfType5135 = new(IPPoolConfig)
		}
		*p.oneOfType5135 = v.(IPPoolConfig)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType5135.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType5135.ObjectType_
	case StaticIPConfig:
		if nil == p.oneOfType5136 {
			p.oneOfType5136 = new(StaticIPConfig)
		}
		*p.oneOfType5136 = v.(StaticIPConfig)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType5136.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType5136.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfBmcIPConfigIpv4) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType5135 != nil && *p.oneOfType5135.ObjectType_ == *p.Discriminator {
		return *p.oneOfType5135
	}
	if p.oneOfType5136 != nil && *p.oneOfType5136.ObjectType_ == *p.Discriminator {
		return *p.oneOfType5136
	}
	return nil
}

func (p *OneOfBmcIPConfigIpv4) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType5135 := new(IPPoolConfig)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType5135)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType5135.ObjectType_ != nil && "lifecycle.v4.config.IPPoolConfig" == *vOneOfType5135.ObjectType_ {
							if nil == p.oneOfType5135 {
								p.oneOfType5135 = new(IPPoolConfig)
							}
							*p.oneOfType5135 = *vOneOfType5135
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType5135.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType5135.ObjectType_
							return nil
						}
					}
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType5136 := new(StaticIPConfig)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType5136)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType5136.ObjectType_ != nil && "lifecycle.v4.config.StaticIPConfig" == *vOneOfType5136.ObjectType_ {
							if nil == p.oneOfType5136 {
								p.oneOfType5136 = new(StaticIPConfig)
							}
							*p.oneOfType5136 = *vOneOfType5136
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType5136.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType5136.ObjectType_
							return nil
						}
					}
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType5135 := new(IPPoolConfig)
	if err := json.Unmarshal(b, vOneOfType5135); err == nil {
		if vOneOfType5135.ObjectType_ != nil && "lifecycle.v4.config.IPPoolConfig" == *vOneOfType5135.ObjectType_ {
			if nil == p.oneOfType5135 {
				p.oneOfType5135 = new(IPPoolConfig)
			}
			*p.oneOfType5135 = *vOneOfType5135
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType5135.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType5135.ObjectType_
			return nil
		}
	}
	vOneOfType5136 := new(StaticIPConfig)
	if err := json.Unmarshal(b, vOneOfType5136); err == nil {
		if vOneOfType5136.ObjectType_ != nil && "lifecycle.v4.config.StaticIPConfig" == *vOneOfType5136.ObjectType_ {
			if nil == p.oneOfType5136 {
				p.oneOfType5136 = new(StaticIPConfig)
			}
			*p.oneOfType5136 = *vOneOfType5136
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType5136.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType5136.ObjectType_
			return nil
		}
	}
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfBmcIPConfigIpv4"))
}

func (p *OneOfBmcIPConfigIpv4) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType5135 != nil && *p.oneOfType5135.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType5135)
	}
	if p.oneOfType5136 != nil && *p.oneOfType5136.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType5136)
	}
	return nil, errors.New("No value to marshal for OneOfBmcIPConfigIpv4")
}

type OneOfDeleteNodeApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType1017 *import3.TaskReference `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfDeleteNodeApiResponseData() *OneOfDeleteNodeApiResponseData {
	p := new(OneOfDeleteNodeApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteNodeApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteNodeApiResponseData is nil"))
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
	case import3.TaskReference:
		if nil == p.oneOfType1017 {
			p.oneOfType1017 = new(import3.TaskReference)
		}
		*p.oneOfType1017 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType1017.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType1017.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfDeleteNodeApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType1017 != nil && *p.oneOfType1017.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1017
	}
	return nil
}

func (p *OneOfDeleteNodeApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import4.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType1017 := new(import3.TaskReference)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType1017)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType1017.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType1017.ObjectType_ {
							if nil == p.oneOfType1017 {
								p.oneOfType1017 = new(import3.TaskReference)
							}
							*p.oneOfType1017 = *vOneOfType1017
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType1017.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType1017.ObjectType_
							return nil
						}
					}
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType1017 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType1017); err == nil {
		if vOneOfType1017.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType1017.ObjectType_ {
			if nil == p.oneOfType1017 {
				p.oneOfType1017 = new(import3.TaskReference)
			}
			*p.oneOfType1017 = *vOneOfType1017
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType1017.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType1017.ObjectType_
			return nil
		}
	}
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteNodeApiResponseData"))
}

func (p *OneOfDeleteNodeApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType1017 != nil && *p.oneOfType1017.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1017)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteNodeApiResponseData")
}

type OneOfRefreshConnectionSpecConfiguration struct {
	Discriminator *string               `json:"-"`
	ObjectType_   *string               `json:"-"`
	oneOfType2104 *RefreshResourcesSpec `json:"-"`
	oneOfType2103 *DiscoverNodesSpec    `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfRefreshConnectionSpecConfiguration() *OneOfRefreshConnectionSpecConfiguration {
	p := new(OneOfRefreshConnectionSpecConfiguration)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfRefreshConnectionSpecConfiguration) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfRefreshConnectionSpecConfiguration is nil"))
	}
	switch v.(type) {
	case RefreshResourcesSpec:
		if nil == p.oneOfType2104 {
			p.oneOfType2104 = new(RefreshResourcesSpec)
		}
		*p.oneOfType2104 = v.(RefreshResourcesSpec)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2104.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2104.ObjectType_
	case DiscoverNodesSpec:
		if nil == p.oneOfType2103 {
			p.oneOfType2103 = new(DiscoverNodesSpec)
		}
		*p.oneOfType2103 = v.(DiscoverNodesSpec)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2103.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2103.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfRefreshConnectionSpecConfiguration) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType2104 != nil && *p.oneOfType2104.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2104
	}
	if p.oneOfType2103 != nil && *p.oneOfType2103.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2103
	}
	return nil
}

func (p *OneOfRefreshConnectionSpecConfiguration) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType2104 := new(RefreshResourcesSpec)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType2104)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType2104.ObjectType_ != nil && "lifecycle.v4.config.RefreshResourcesSpec" == *vOneOfType2104.ObjectType_ {
							if nil == p.oneOfType2104 {
								p.oneOfType2104 = new(RefreshResourcesSpec)
							}
							*p.oneOfType2104 = *vOneOfType2104
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType2104.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType2104.ObjectType_
							return nil
						}
					}
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType2103 := new(DiscoverNodesSpec)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType2103)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType2103.ObjectType_ != nil && "lifecycle.v4.config.DiscoverNodesSpec" == *vOneOfType2103.ObjectType_ {
							if nil == p.oneOfType2103 {
								p.oneOfType2103 = new(DiscoverNodesSpec)
							}
							*p.oneOfType2103 = *vOneOfType2103
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType2103.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType2103.ObjectType_
							return nil
						}
					}
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType2104 := new(RefreshResourcesSpec)
	if err := json.Unmarshal(b, vOneOfType2104); err == nil {
		if vOneOfType2104.ObjectType_ != nil && "lifecycle.v4.config.RefreshResourcesSpec" == *vOneOfType2104.ObjectType_ {
			if nil == p.oneOfType2104 {
				p.oneOfType2104 = new(RefreshResourcesSpec)
			}
			*p.oneOfType2104 = *vOneOfType2104
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2104.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2104.ObjectType_
			return nil
		}
	}
	vOneOfType2103 := new(DiscoverNodesSpec)
	if err := json.Unmarshal(b, vOneOfType2103); err == nil {
		if vOneOfType2103.ObjectType_ != nil && "lifecycle.v4.config.DiscoverNodesSpec" == *vOneOfType2103.ObjectType_ {
			if nil == p.oneOfType2103 {
				p.oneOfType2103 = new(DiscoverNodesSpec)
			}
			*p.oneOfType2103 = *vOneOfType2103
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2103.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2103.ObjectType_
			return nil
		}
	}
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfRefreshConnectionSpecConfiguration"))
}

func (p *OneOfRefreshConnectionSpecConfiguration) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType2104 != nil && *p.oneOfType2104.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2104)
	}
	if p.oneOfType2103 != nil && *p.oneOfType2103.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2103)
	}
	return nil, errors.New("No value to marshal for OneOfRefreshConnectionSpecConfiguration")
}

type OneOfGetIpPoolApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType1048 *IpPool                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfGetIpPoolApiResponseData() *OneOfGetIpPoolApiResponseData {
	p := new(OneOfGetIpPoolApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetIpPoolApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetIpPoolApiResponseData is nil"))
	}
	switch v.(type) {
	case IpPool:
		if nil == p.oneOfType1048 {
			p.oneOfType1048 = new(IpPool)
		}
		*p.oneOfType1048 = v.(IpPool)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType1048.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType1048.ObjectType_
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

func (p *OneOfGetIpPoolApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType1048 != nil && *p.oneOfType1048.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1048
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetIpPoolApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType1048 := new(IpPool)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType1048)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType1048.ObjectType_ != nil && "lifecycle.v4.config.IpPool" == *vOneOfType1048.ObjectType_ {
							if nil == p.oneOfType1048 {
								p.oneOfType1048 = new(IpPool)
							}
							*p.oneOfType1048 = *vOneOfType1048
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType1048.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType1048.ObjectType_
							return nil
						}
					}
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import4.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType1048 := new(IpPool)
	if err := json.Unmarshal(b, vOneOfType1048); err == nil {
		if vOneOfType1048.ObjectType_ != nil && "lifecycle.v4.config.IpPool" == *vOneOfType1048.ObjectType_ {
			if nil == p.oneOfType1048 {
				p.oneOfType1048 = new(IpPool)
			}
			*p.oneOfType1048 = *vOneOfType1048
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType1048.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType1048.ObjectType_
			return nil
		}
	}
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetIpPoolApiResponseData"))
}

func (p *OneOfGetIpPoolApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType1048 != nil && *p.oneOfType1048.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1048)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetIpPoolApiResponseData")
}

type OneOfImageNodeSpecConfiguration struct {
	Discriminator *string                  `json:"-"`
	ObjectType_   *string                  `json:"-"`
	oneOfType2453 *HciImagingDetails       `json:"-"`
	oneOfType2454 *ReconfigurationDetails  `json:"-"`
	oneOfType2451 *HostInstallationDetails `json:"-"`
	oneOfType2452 *AosInstallationDetails  `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfImageNodeSpecConfiguration() *OneOfImageNodeSpecConfiguration {
	p := new(OneOfImageNodeSpecConfiguration)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfImageNodeSpecConfiguration) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfImageNodeSpecConfiguration is nil"))
	}
	switch v.(type) {
	case HciImagingDetails:
		if nil == p.oneOfType2453 {
			p.oneOfType2453 = new(HciImagingDetails)
		}
		*p.oneOfType2453 = v.(HciImagingDetails)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2453.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2453.ObjectType_
	case ReconfigurationDetails:
		if nil == p.oneOfType2454 {
			p.oneOfType2454 = new(ReconfigurationDetails)
		}
		*p.oneOfType2454 = v.(ReconfigurationDetails)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2454.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2454.ObjectType_
	case HostInstallationDetails:
		if nil == p.oneOfType2451 {
			p.oneOfType2451 = new(HostInstallationDetails)
		}
		*p.oneOfType2451 = v.(HostInstallationDetails)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2451.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2451.ObjectType_
	case AosInstallationDetails:
		if nil == p.oneOfType2452 {
			p.oneOfType2452 = new(AosInstallationDetails)
		}
		*p.oneOfType2452 = v.(AosInstallationDetails)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2452.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2452.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfImageNodeSpecConfiguration) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType2453 != nil && *p.oneOfType2453.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2453
	}
	if p.oneOfType2454 != nil && *p.oneOfType2454.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2454
	}
	if p.oneOfType2451 != nil && *p.oneOfType2451.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2451
	}
	if p.oneOfType2452 != nil && *p.oneOfType2452.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2452
	}
	return nil
}

func (p *OneOfImageNodeSpecConfiguration) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType2453 := new(HciImagingDetails)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType2453)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType2453.ObjectType_ != nil && "lifecycle.v4.config.HciImagingDetails" == *vOneOfType2453.ObjectType_ {
							if nil == p.oneOfType2453 {
								p.oneOfType2453 = new(HciImagingDetails)
							}
							*p.oneOfType2453 = *vOneOfType2453
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType2453.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType2453.ObjectType_
							return nil
						}
					}
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType2454 := new(ReconfigurationDetails)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType2454)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType2454.ObjectType_ != nil && "lifecycle.v4.config.ReconfigurationDetails" == *vOneOfType2454.ObjectType_ {
							if nil == p.oneOfType2454 {
								p.oneOfType2454 = new(ReconfigurationDetails)
							}
							*p.oneOfType2454 = *vOneOfType2454
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType2454.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType2454.ObjectType_
							return nil
						}
					}
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType2451 := new(HostInstallationDetails)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType2451)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType2451.ObjectType_ != nil && "lifecycle.v4.config.HostInstallationDetails" == *vOneOfType2451.ObjectType_ {
							if nil == p.oneOfType2451 {
								p.oneOfType2451 = new(HostInstallationDetails)
							}
							*p.oneOfType2451 = *vOneOfType2451
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType2451.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType2451.ObjectType_
							return nil
						}
					}
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType2452 := new(AosInstallationDetails)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType2452)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType2452.ObjectType_ != nil && "lifecycle.v4.config.AosInstallationDetails" == *vOneOfType2452.ObjectType_ {
							if nil == p.oneOfType2452 {
								p.oneOfType2452 = new(AosInstallationDetails)
							}
							*p.oneOfType2452 = *vOneOfType2452
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType2452.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType2452.ObjectType_
							return nil
						}
					}
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType2453 := new(HciImagingDetails)
	if err := json.Unmarshal(b, vOneOfType2453); err == nil {
		if vOneOfType2453.ObjectType_ != nil && "lifecycle.v4.config.HciImagingDetails" == *vOneOfType2453.ObjectType_ {
			if nil == p.oneOfType2453 {
				p.oneOfType2453 = new(HciImagingDetails)
			}
			*p.oneOfType2453 = *vOneOfType2453
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2453.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2453.ObjectType_
			return nil
		}
	}
	vOneOfType2454 := new(ReconfigurationDetails)
	if err := json.Unmarshal(b, vOneOfType2454); err == nil {
		if vOneOfType2454.ObjectType_ != nil && "lifecycle.v4.config.ReconfigurationDetails" == *vOneOfType2454.ObjectType_ {
			if nil == p.oneOfType2454 {
				p.oneOfType2454 = new(ReconfigurationDetails)
			}
			*p.oneOfType2454 = *vOneOfType2454
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2454.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2454.ObjectType_
			return nil
		}
	}
	vOneOfType2451 := new(HostInstallationDetails)
	if err := json.Unmarshal(b, vOneOfType2451); err == nil {
		if vOneOfType2451.ObjectType_ != nil && "lifecycle.v4.config.HostInstallationDetails" == *vOneOfType2451.ObjectType_ {
			if nil == p.oneOfType2451 {
				p.oneOfType2451 = new(HostInstallationDetails)
			}
			*p.oneOfType2451 = *vOneOfType2451
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2451.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2451.ObjectType_
			return nil
		}
	}
	vOneOfType2452 := new(AosInstallationDetails)
	if err := json.Unmarshal(b, vOneOfType2452); err == nil {
		if vOneOfType2452.ObjectType_ != nil && "lifecycle.v4.config.AosInstallationDetails" == *vOneOfType2452.ObjectType_ {
			if nil == p.oneOfType2452 {
				p.oneOfType2452 = new(AosInstallationDetails)
			}
			*p.oneOfType2452 = *vOneOfType2452
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2452.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2452.ObjectType_
			return nil
		}
	}
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfImageNodeSpecConfiguration"))
}

func (p *OneOfImageNodeSpecConfiguration) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType2453 != nil && *p.oneOfType2453.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2453)
	}
	if p.oneOfType2454 != nil && *p.oneOfType2454.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2454)
	}
	if p.oneOfType2451 != nil && *p.oneOfType2451.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2451)
	}
	if p.oneOfType2452 != nil && *p.oneOfType2452.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2452)
	}
	return nil, errors.New("No value to marshal for OneOfImageNodeSpecConfiguration")
}

type OneOfListIpPoolsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType1043 []IpPool               `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfListIpPoolsApiResponseData() *OneOfListIpPoolsApiResponseData {
	p := new(OneOfListIpPoolsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListIpPoolsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListIpPoolsApiResponseData is nil"))
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
	case []IpPool:
		p.oneOfType1043 = v.([]IpPool)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<lifecycle.v4.config.IpPool>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<lifecycle.v4.config.IpPool>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListIpPoolsApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<lifecycle.v4.config.IpPool>" == *p.Discriminator {
		return p.oneOfType1043
	}
	return nil
}

func (p *OneOfListIpPoolsApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import4.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["List<lifecycle.v4.config.IpPool>"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType1043 := new([]IpPool)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType1043)
					if unmarshalErr == nil {
						// For arrays, verify the array item ObjectType matches
						if vOneOfType1043 == nil || len(*vOneOfType1043) == 0 || ((*vOneOfType1043)[0].ObjectType_ != nil && "lifecycle.v4.config.IpPool" == *((*vOneOfType1043)[0].ObjectType_)) {
							p.oneOfType1043 = *vOneOfType1043
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = "List<lifecycle.v4.config.IpPool>"
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = "List<lifecycle.v4.config.IpPool>"
							return nil
						}
					}
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType1043 := new([]IpPool)
	if err := json.Unmarshal(b, vOneOfType1043); err == nil {
		if len(*vOneOfType1043) == 0 || (vOneOfType1043 != nil && (*vOneOfType1043)[0].ObjectType_ != nil && "lifecycle.v4.config.IpPool" == *((*vOneOfType1043)[0].ObjectType_)) {
			p.oneOfType1043 = *vOneOfType1043
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<lifecycle.v4.config.IpPool>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<lifecycle.v4.config.IpPool>"
			return nil
		}
	}
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListIpPoolsApiResponseData"))
}

func (p *OneOfListIpPoolsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<lifecycle.v4.config.IpPool>" == *p.Discriminator {
		return json.Marshal(p.oneOfType1043)
	}
	return nil, errors.New("No value to marshal for OneOfListIpPoolsApiResponseData")
}

type OneOfGetImageApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType4006 *Image                 `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfGetImageApiResponseData() *OneOfGetImageApiResponseData {
	p := new(OneOfGetImageApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetImageApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetImageApiResponseData is nil"))
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
	case Image:
		if nil == p.oneOfType4006 {
			p.oneOfType4006 = new(Image)
		}
		*p.oneOfType4006 = v.(Image)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType4006.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType4006.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfGetImageApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType4006 != nil && *p.oneOfType4006.ObjectType_ == *p.Discriminator {
		return *p.oneOfType4006
	}
	return nil
}

func (p *OneOfGetImageApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import4.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType4006 := new(Image)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType4006)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType4006.ObjectType_ != nil && "lifecycle.v4.config.Image" == *vOneOfType4006.ObjectType_ {
							if nil == p.oneOfType4006 {
								p.oneOfType4006 = new(Image)
							}
							*p.oneOfType4006 = *vOneOfType4006
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType4006.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType4006.ObjectType_
							return nil
						}
					}
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType4006 := new(Image)
	if err := json.Unmarshal(b, vOneOfType4006); err == nil {
		if vOneOfType4006.ObjectType_ != nil && "lifecycle.v4.config.Image" == *vOneOfType4006.ObjectType_ {
			if nil == p.oneOfType4006 {
				p.oneOfType4006 = new(Image)
			}
			*p.oneOfType4006 = *vOneOfType4006
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType4006.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType4006.ObjectType_
			return nil
		}
	}
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetImageApiResponseData"))
}

func (p *OneOfGetImageApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType4006 != nil && *p.oneOfType4006.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType4006)
	}
	return nil, errors.New("No value to marshal for OneOfGetImageApiResponseData")
}

type OneOfGetPatchedImageApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType1005 *PatchedImage          `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfGetPatchedImageApiResponseData() *OneOfGetPatchedImageApiResponseData {
	p := new(OneOfGetPatchedImageApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetPatchedImageApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetPatchedImageApiResponseData is nil"))
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
	case PatchedImage:
		if nil == p.oneOfType1005 {
			p.oneOfType1005 = new(PatchedImage)
		}
		*p.oneOfType1005 = v.(PatchedImage)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType1005.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType1005.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfGetPatchedImageApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType1005 != nil && *p.oneOfType1005.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1005
	}
	return nil
}

func (p *OneOfGetPatchedImageApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import4.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType1005 := new(PatchedImage)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType1005)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType1005.ObjectType_ != nil && "lifecycle.v4.config.PatchedImage" == *vOneOfType1005.ObjectType_ {
							if nil == p.oneOfType1005 {
								p.oneOfType1005 = new(PatchedImage)
							}
							*p.oneOfType1005 = *vOneOfType1005
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType1005.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType1005.ObjectType_
							return nil
						}
					}
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType1005 := new(PatchedImage)
	if err := json.Unmarshal(b, vOneOfType1005); err == nil {
		if vOneOfType1005.ObjectType_ != nil && "lifecycle.v4.config.PatchedImage" == *vOneOfType1005.ObjectType_ {
			if nil == p.oneOfType1005 {
				p.oneOfType1005 = new(PatchedImage)
			}
			*p.oneOfType1005 = *vOneOfType1005
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType1005.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType1005.ObjectType_
			return nil
		}
	}
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetPatchedImageApiResponseData"))
}

func (p *OneOfGetPatchedImageApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType1005 != nil && *p.oneOfType1005.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1005)
	}
	return nil, errors.New("No value to marshal for OneOfGetPatchedImageApiResponseData")
}

type OneOfListNodesApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType1000 []Node                 `json:"-"`
	oneOfType401  []NodeProjection       `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfListNodesApiResponseData() *OneOfListNodesApiResponseData {
	p := new(OneOfListNodesApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListNodesApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListNodesApiResponseData is nil"))
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
	case []Node:
		p.oneOfType1000 = v.([]Node)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<lifecycle.v4.config.Node>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<lifecycle.v4.config.Node>"
	case []NodeProjection:
		p.oneOfType401 = v.([]NodeProjection)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<lifecycle.v4.config.NodeProjection>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<lifecycle.v4.config.NodeProjection>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListNodesApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<lifecycle.v4.config.Node>" == *p.Discriminator {
		return p.oneOfType1000
	}
	if "List<lifecycle.v4.config.NodeProjection>" == *p.Discriminator {
		return p.oneOfType401
	}
	return nil
}

func (p *OneOfListNodesApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import4.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["List<lifecycle.v4.config.Node>"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType1000 := new([]Node)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType1000)
					if unmarshalErr == nil {
						// For arrays, verify the array item ObjectType matches
						if vOneOfType1000 == nil || len(*vOneOfType1000) == 0 || ((*vOneOfType1000)[0].ObjectType_ != nil && "lifecycle.v4.config.Node" == *((*vOneOfType1000)[0].ObjectType_)) {
							p.oneOfType1000 = *vOneOfType1000
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = "List<lifecycle.v4.config.Node>"
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = "List<lifecycle.v4.config.Node>"
							return nil
						}
					}
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["List<lifecycle.v4.config.NodeProjection>"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType401 := new([]NodeProjection)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType401)
					if unmarshalErr == nil {
						// For arrays, verify the array item ObjectType matches
						if vOneOfType401 == nil || len(*vOneOfType401) == 0 || ((*vOneOfType401)[0].ObjectType_ != nil && "lifecycle.v4.config.NodeProjection" == *((*vOneOfType401)[0].ObjectType_)) {
							p.oneOfType401 = *vOneOfType401
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = "List<lifecycle.v4.config.NodeProjection>"
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = "List<lifecycle.v4.config.NodeProjection>"
							return nil
						}
					}
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType1000 := new([]Node)
	if err := json.Unmarshal(b, vOneOfType1000); err == nil {
		if len(*vOneOfType1000) == 0 || (vOneOfType1000 != nil && (*vOneOfType1000)[0].ObjectType_ != nil && "lifecycle.v4.config.Node" == *((*vOneOfType1000)[0].ObjectType_)) {
			p.oneOfType1000 = *vOneOfType1000
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<lifecycle.v4.config.Node>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<lifecycle.v4.config.Node>"
			return nil
		}
	}
	vOneOfType401 := new([]NodeProjection)
	if err := json.Unmarshal(b, vOneOfType401); err == nil {
		if len(*vOneOfType401) == 0 || (vOneOfType401 != nil && (*vOneOfType401)[0].ObjectType_ != nil && "lifecycle.v4.config.NodeProjection" == *((*vOneOfType401)[0].ObjectType_)) {
			p.oneOfType401 = *vOneOfType401
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<lifecycle.v4.config.NodeProjection>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<lifecycle.v4.config.NodeProjection>"
			return nil
		}
	}
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListNodesApiResponseData"))
}

func (p *OneOfListNodesApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<lifecycle.v4.config.Node>" == *p.Discriminator {
		return json.Marshal(p.oneOfType1000)
	}
	if "List<lifecycle.v4.config.NodeProjection>" == *p.Discriminator {
		return json.Marshal(p.oneOfType401)
	}
	return nil, errors.New("No value to marshal for OneOfListNodesApiResponseData")
}

type OneOfListClaimTokensApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType1002 []ClaimToken           `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfListClaimTokensApiResponseData() *OneOfListClaimTokensApiResponseData {
	p := new(OneOfListClaimTokensApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListClaimTokensApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListClaimTokensApiResponseData is nil"))
	}
	switch v.(type) {
	case []ClaimToken:
		p.oneOfType1002 = v.([]ClaimToken)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<lifecycle.v4.config.ClaimToken>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<lifecycle.v4.config.ClaimToken>"
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

func (p *OneOfListClaimTokensApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if "List<lifecycle.v4.config.ClaimToken>" == *p.Discriminator {
		return p.oneOfType1002
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfListClaimTokensApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["List<lifecycle.v4.config.ClaimToken>"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType1002 := new([]ClaimToken)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType1002)
					if unmarshalErr == nil {
						// For arrays, verify the array item ObjectType matches
						if vOneOfType1002 == nil || len(*vOneOfType1002) == 0 || ((*vOneOfType1002)[0].ObjectType_ != nil && "lifecycle.v4.config.ClaimToken" == *((*vOneOfType1002)[0].ObjectType_)) {
							p.oneOfType1002 = *vOneOfType1002
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = "List<lifecycle.v4.config.ClaimToken>"
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = "List<lifecycle.v4.config.ClaimToken>"
							return nil
						}
					}
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import4.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType1002 := new([]ClaimToken)
	if err := json.Unmarshal(b, vOneOfType1002); err == nil {
		if len(*vOneOfType1002) == 0 || (vOneOfType1002 != nil && (*vOneOfType1002)[0].ObjectType_ != nil && "lifecycle.v4.config.ClaimToken" == *((*vOneOfType1002)[0].ObjectType_)) {
			p.oneOfType1002 = *vOneOfType1002
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<lifecycle.v4.config.ClaimToken>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<lifecycle.v4.config.ClaimToken>"
			return nil
		}
	}
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListClaimTokensApiResponseData"))
}

func (p *OneOfListClaimTokensApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if "List<lifecycle.v4.config.ClaimToken>" == *p.Discriminator {
		return json.Marshal(p.oneOfType1002)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfListClaimTokensApiResponseData")
}

type OneOfListConnectionNodesApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType1030 []DiscoveredNode       `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfListConnectionNodesApiResponseData() *OneOfListConnectionNodesApiResponseData {
	p := new(OneOfListConnectionNodesApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListConnectionNodesApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListConnectionNodesApiResponseData is nil"))
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
	case []DiscoveredNode:
		p.oneOfType1030 = v.([]DiscoveredNode)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<lifecycle.v4.config.DiscoveredNode>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<lifecycle.v4.config.DiscoveredNode>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListConnectionNodesApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<lifecycle.v4.config.DiscoveredNode>" == *p.Discriminator {
		return p.oneOfType1030
	}
	return nil
}

func (p *OneOfListConnectionNodesApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import4.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["List<lifecycle.v4.config.DiscoveredNode>"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType1030 := new([]DiscoveredNode)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType1030)
					if unmarshalErr == nil {
						// For arrays, verify the array item ObjectType matches
						if vOneOfType1030 == nil || len(*vOneOfType1030) == 0 || ((*vOneOfType1030)[0].ObjectType_ != nil && "lifecycle.v4.config.DiscoveredNode" == *((*vOneOfType1030)[0].ObjectType_)) {
							p.oneOfType1030 = *vOneOfType1030
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = "List<lifecycle.v4.config.DiscoveredNode>"
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = "List<lifecycle.v4.config.DiscoveredNode>"
							return nil
						}
					}
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType1030 := new([]DiscoveredNode)
	if err := json.Unmarshal(b, vOneOfType1030); err == nil {
		if len(*vOneOfType1030) == 0 || (vOneOfType1030 != nil && (*vOneOfType1030)[0].ObjectType_ != nil && "lifecycle.v4.config.DiscoveredNode" == *((*vOneOfType1030)[0].ObjectType_)) {
			p.oneOfType1030 = *vOneOfType1030
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<lifecycle.v4.config.DiscoveredNode>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<lifecycle.v4.config.DiscoveredNode>"
			return nil
		}
	}
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListConnectionNodesApiResponseData"))
}

func (p *OneOfListConnectionNodesApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<lifecycle.v4.config.DiscoveredNode>" == *p.Discriminator {
		return json.Marshal(p.oneOfType1030)
	}
	return nil, errors.New("No value to marshal for OneOfListConnectionNodesApiResponseData")
}

type OneOfListMacPoolsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType1043 []MacPool              `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfListMacPoolsApiResponseData() *OneOfListMacPoolsApiResponseData {
	p := new(OneOfListMacPoolsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListMacPoolsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListMacPoolsApiResponseData is nil"))
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
	case []MacPool:
		p.oneOfType1043 = v.([]MacPool)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<lifecycle.v4.config.MacPool>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<lifecycle.v4.config.MacPool>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListMacPoolsApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<lifecycle.v4.config.MacPool>" == *p.Discriminator {
		return p.oneOfType1043
	}
	return nil
}

func (p *OneOfListMacPoolsApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import4.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["List<lifecycle.v4.config.MacPool>"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType1043 := new([]MacPool)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType1043)
					if unmarshalErr == nil {
						// For arrays, verify the array item ObjectType matches
						if vOneOfType1043 == nil || len(*vOneOfType1043) == 0 || ((*vOneOfType1043)[0].ObjectType_ != nil && "lifecycle.v4.config.MacPool" == *((*vOneOfType1043)[0].ObjectType_)) {
							p.oneOfType1043 = *vOneOfType1043
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = "List<lifecycle.v4.config.MacPool>"
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = "List<lifecycle.v4.config.MacPool>"
							return nil
						}
					}
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType1043 := new([]MacPool)
	if err := json.Unmarshal(b, vOneOfType1043); err == nil {
		if len(*vOneOfType1043) == 0 || (vOneOfType1043 != nil && (*vOneOfType1043)[0].ObjectType_ != nil && "lifecycle.v4.config.MacPool" == *((*vOneOfType1043)[0].ObjectType_)) {
			p.oneOfType1043 = *vOneOfType1043
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<lifecycle.v4.config.MacPool>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<lifecycle.v4.config.MacPool>"
			return nil
		}
	}
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListMacPoolsApiResponseData"))
}

func (p *OneOfListMacPoolsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<lifecycle.v4.config.MacPool>" == *p.Discriminator {
		return json.Marshal(p.oneOfType1043)
	}
	return nil, errors.New("No value to marshal for OneOfListMacPoolsApiResponseData")
}

type OneOfGetClaimTokenSecretApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType1012 *ClaimTokenSecret      `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfGetClaimTokenSecretApiResponseData() *OneOfGetClaimTokenSecretApiResponseData {
	p := new(OneOfGetClaimTokenSecretApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetClaimTokenSecretApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetClaimTokenSecretApiResponseData is nil"))
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
	case ClaimTokenSecret:
		if nil == p.oneOfType1012 {
			p.oneOfType1012 = new(ClaimTokenSecret)
		}
		*p.oneOfType1012 = v.(ClaimTokenSecret)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType1012.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType1012.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfGetClaimTokenSecretApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType1012 != nil && *p.oneOfType1012.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1012
	}
	return nil
}

func (p *OneOfGetClaimTokenSecretApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import4.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType1012 := new(ClaimTokenSecret)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType1012)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType1012.ObjectType_ != nil && "lifecycle.v4.config.ClaimTokenSecret" == *vOneOfType1012.ObjectType_ {
							if nil == p.oneOfType1012 {
								p.oneOfType1012 = new(ClaimTokenSecret)
							}
							*p.oneOfType1012 = *vOneOfType1012
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType1012.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType1012.ObjectType_
							return nil
						}
					}
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType1012 := new(ClaimTokenSecret)
	if err := json.Unmarshal(b, vOneOfType1012); err == nil {
		if vOneOfType1012.ObjectType_ != nil && "lifecycle.v4.config.ClaimTokenSecret" == *vOneOfType1012.ObjectType_ {
			if nil == p.oneOfType1012 {
				p.oneOfType1012 = new(ClaimTokenSecret)
			}
			*p.oneOfType1012 = *vOneOfType1012
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType1012.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType1012.ObjectType_
			return nil
		}
	}
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetClaimTokenSecretApiResponseData"))
}

func (p *OneOfGetClaimTokenSecretApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType1012 != nil && *p.oneOfType1012.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1012)
	}
	return nil, errors.New("No value to marshal for OneOfGetClaimTokenSecretApiResponseData")
}

type OneOfBondSettingsBondConfig struct {
	Discriminator *string           `json:"-"`
	ObjectType_   *string           `json:"-"`
	oneOfType4002 *ActiveBackupBond `json:"-"`
	oneOfType4003 *LacpBond         `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfBondSettingsBondConfig() *OneOfBondSettingsBondConfig {
	p := new(OneOfBondSettingsBondConfig)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfBondSettingsBondConfig) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfBondSettingsBondConfig is nil"))
	}
	switch v.(type) {
	case ActiveBackupBond:
		if nil == p.oneOfType4002 {
			p.oneOfType4002 = new(ActiveBackupBond)
		}
		*p.oneOfType4002 = v.(ActiveBackupBond)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType4002.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType4002.ObjectType_
	case LacpBond:
		if nil == p.oneOfType4003 {
			p.oneOfType4003 = new(LacpBond)
		}
		*p.oneOfType4003 = v.(LacpBond)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType4003.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType4003.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfBondSettingsBondConfig) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType4002 != nil && *p.oneOfType4002.ObjectType_ == *p.Discriminator {
		return *p.oneOfType4002
	}
	if p.oneOfType4003 != nil && *p.oneOfType4003.ObjectType_ == *p.Discriminator {
		return *p.oneOfType4003
	}
	return nil
}

func (p *OneOfBondSettingsBondConfig) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType4002 := new(ActiveBackupBond)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType4002)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType4002.ObjectType_ != nil && "lifecycle.v4.config.ActiveBackupBond" == *vOneOfType4002.ObjectType_ {
							if nil == p.oneOfType4002 {
								p.oneOfType4002 = new(ActiveBackupBond)
							}
							*p.oneOfType4002 = *vOneOfType4002
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType4002.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType4002.ObjectType_
							return nil
						}
					}
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType4003 := new(LacpBond)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType4003)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType4003.ObjectType_ != nil && "lifecycle.v4.config.LacpBond" == *vOneOfType4003.ObjectType_ {
							if nil == p.oneOfType4003 {
								p.oneOfType4003 = new(LacpBond)
							}
							*p.oneOfType4003 = *vOneOfType4003
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType4003.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType4003.ObjectType_
							return nil
						}
					}
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType4002 := new(ActiveBackupBond)
	if err := json.Unmarshal(b, vOneOfType4002); err == nil {
		if vOneOfType4002.ObjectType_ != nil && "lifecycle.v4.config.ActiveBackupBond" == *vOneOfType4002.ObjectType_ {
			if nil == p.oneOfType4002 {
				p.oneOfType4002 = new(ActiveBackupBond)
			}
			*p.oneOfType4002 = *vOneOfType4002
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType4002.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType4002.ObjectType_
			return nil
		}
	}
	vOneOfType4003 := new(LacpBond)
	if err := json.Unmarshal(b, vOneOfType4003); err == nil {
		if vOneOfType4003.ObjectType_ != nil && "lifecycle.v4.config.LacpBond" == *vOneOfType4003.ObjectType_ {
			if nil == p.oneOfType4003 {
				p.oneOfType4003 = new(LacpBond)
			}
			*p.oneOfType4003 = *vOneOfType4003
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType4003.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType4003.ObjectType_
			return nil
		}
	}
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfBondSettingsBondConfig"))
}

func (p *OneOfBondSettingsBondConfig) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType4002 != nil && *p.oneOfType4002.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType4002)
	}
	if p.oneOfType4003 != nil && *p.oneOfType4003.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType4003)
	}
	return nil, errors.New("No value to marshal for OneOfBondSettingsBondConfig")
}

type OneOfPatchedImageImageDetails struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType3401 *LocalHostImageDetails `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfPatchedImageImageDetails() *OneOfPatchedImageImageDetails {
	p := new(OneOfPatchedImageImageDetails)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfPatchedImageImageDetails) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfPatchedImageImageDetails is nil"))
	}
	switch v.(type) {
	case LocalHostImageDetails:
		if nil == p.oneOfType3401 {
			p.oneOfType3401 = new(LocalHostImageDetails)
		}
		*p.oneOfType3401 = v.(LocalHostImageDetails)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType3401.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType3401.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfPatchedImageImageDetails) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType3401 != nil && *p.oneOfType3401.ObjectType_ == *p.Discriminator {
		return *p.oneOfType3401
	}
	return nil
}

func (p *OneOfPatchedImageImageDetails) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType3401 := new(LocalHostImageDetails)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType3401)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType3401.ObjectType_ != nil && "lifecycle.v4.config.LocalHostImageDetails" == *vOneOfType3401.ObjectType_ {
							if nil == p.oneOfType3401 {
								p.oneOfType3401 = new(LocalHostImageDetails)
							}
							*p.oneOfType3401 = *vOneOfType3401
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType3401.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType3401.ObjectType_
							return nil
						}
					}
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType3401 := new(LocalHostImageDetails)
	if err := json.Unmarshal(b, vOneOfType3401); err == nil {
		if vOneOfType3401.ObjectType_ != nil && "lifecycle.v4.config.LocalHostImageDetails" == *vOneOfType3401.ObjectType_ {
			if nil == p.oneOfType3401 {
				p.oneOfType3401 = new(LocalHostImageDetails)
			}
			*p.oneOfType3401 = *vOneOfType3401
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType3401.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType3401.ObjectType_
			return nil
		}
	}
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfPatchedImageImageDetails"))
}

func (p *OneOfPatchedImageImageDetails) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType3401 != nil && *p.oneOfType3401.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType3401)
	}
	return nil, errors.New("No value to marshal for OneOfPatchedImageImageDetails")
}

type OneOfBmcIPConfigIpv6 struct {
	Discriminator *string         `json:"-"`
	ObjectType_   *string         `json:"-"`
	oneOfType5137 *IPPoolConfig   `json:"-"`
	oneOfType5138 *StaticIPConfig `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfBmcIPConfigIpv6() *OneOfBmcIPConfigIpv6 {
	p := new(OneOfBmcIPConfigIpv6)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfBmcIPConfigIpv6) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfBmcIPConfigIpv6 is nil"))
	}
	switch v.(type) {
	case IPPoolConfig:
		if nil == p.oneOfType5137 {
			p.oneOfType5137 = new(IPPoolConfig)
		}
		*p.oneOfType5137 = v.(IPPoolConfig)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType5137.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType5137.ObjectType_
	case StaticIPConfig:
		if nil == p.oneOfType5138 {
			p.oneOfType5138 = new(StaticIPConfig)
		}
		*p.oneOfType5138 = v.(StaticIPConfig)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType5138.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType5138.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfBmcIPConfigIpv6) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType5137 != nil && *p.oneOfType5137.ObjectType_ == *p.Discriminator {
		return *p.oneOfType5137
	}
	if p.oneOfType5138 != nil && *p.oneOfType5138.ObjectType_ == *p.Discriminator {
		return *p.oneOfType5138
	}
	return nil
}

func (p *OneOfBmcIPConfigIpv6) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType5137 := new(IPPoolConfig)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType5137)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType5137.ObjectType_ != nil && "lifecycle.v4.config.IPPoolConfig" == *vOneOfType5137.ObjectType_ {
							if nil == p.oneOfType5137 {
								p.oneOfType5137 = new(IPPoolConfig)
							}
							*p.oneOfType5137 = *vOneOfType5137
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType5137.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType5137.ObjectType_
							return nil
						}
					}
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType5138 := new(StaticIPConfig)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType5138)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType5138.ObjectType_ != nil && "lifecycle.v4.config.StaticIPConfig" == *vOneOfType5138.ObjectType_ {
							if nil == p.oneOfType5138 {
								p.oneOfType5138 = new(StaticIPConfig)
							}
							*p.oneOfType5138 = *vOneOfType5138
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType5138.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType5138.ObjectType_
							return nil
						}
					}
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType5137 := new(IPPoolConfig)
	if err := json.Unmarshal(b, vOneOfType5137); err == nil {
		if vOneOfType5137.ObjectType_ != nil && "lifecycle.v4.config.IPPoolConfig" == *vOneOfType5137.ObjectType_ {
			if nil == p.oneOfType5137 {
				p.oneOfType5137 = new(IPPoolConfig)
			}
			*p.oneOfType5137 = *vOneOfType5137
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType5137.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType5137.ObjectType_
			return nil
		}
	}
	vOneOfType5138 := new(StaticIPConfig)
	if err := json.Unmarshal(b, vOneOfType5138); err == nil {
		if vOneOfType5138.ObjectType_ != nil && "lifecycle.v4.config.StaticIPConfig" == *vOneOfType5138.ObjectType_ {
			if nil == p.oneOfType5138 {
				p.oneOfType5138 = new(StaticIPConfig)
			}
			*p.oneOfType5138 = *vOneOfType5138
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType5138.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType5138.ObjectType_
			return nil
		}
	}
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfBmcIPConfigIpv6"))
}

func (p *OneOfBmcIPConfigIpv6) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType5137 != nil && *p.oneOfType5137.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType5137)
	}
	if p.oneOfType5138 != nil && *p.oneOfType5138.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType5138)
	}
	return nil, errors.New("No value to marshal for OneOfBmcIPConfigIpv6")
}

type OneOfGetHardwareProviderApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType1005 *HardwareProvider      `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfGetHardwareProviderApiResponseData() *OneOfGetHardwareProviderApiResponseData {
	p := new(OneOfGetHardwareProviderApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetHardwareProviderApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetHardwareProviderApiResponseData is nil"))
	}
	switch v.(type) {
	case HardwareProvider:
		if nil == p.oneOfType1005 {
			p.oneOfType1005 = new(HardwareProvider)
		}
		*p.oneOfType1005 = v.(HardwareProvider)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType1005.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType1005.ObjectType_
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

func (p *OneOfGetHardwareProviderApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType1005 != nil && *p.oneOfType1005.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1005
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetHardwareProviderApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType1005 := new(HardwareProvider)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType1005)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType1005.ObjectType_ != nil && "lifecycle.v4.config.HardwareProvider" == *vOneOfType1005.ObjectType_ {
							if nil == p.oneOfType1005 {
								p.oneOfType1005 = new(HardwareProvider)
							}
							*p.oneOfType1005 = *vOneOfType1005
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType1005.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType1005.ObjectType_
							return nil
						}
					}
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import4.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType1005 := new(HardwareProvider)
	if err := json.Unmarshal(b, vOneOfType1005); err == nil {
		if vOneOfType1005.ObjectType_ != nil && "lifecycle.v4.config.HardwareProvider" == *vOneOfType1005.ObjectType_ {
			if nil == p.oneOfType1005 {
				p.oneOfType1005 = new(HardwareProvider)
			}
			*p.oneOfType1005 = *vOneOfType1005
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType1005.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType1005.ObjectType_
			return nil
		}
	}
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetHardwareProviderApiResponseData"))
}

func (p *OneOfGetHardwareProviderApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType1005 != nil && *p.oneOfType1005.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1005)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetHardwareProviderApiResponseData")
}

type OneOfImageNodeApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType1010 *import3.TaskReference `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfImageNodeApiResponseData() *OneOfImageNodeApiResponseData {
	p := new(OneOfImageNodeApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfImageNodeApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfImageNodeApiResponseData is nil"))
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
	case import3.TaskReference:
		if nil == p.oneOfType1010 {
			p.oneOfType1010 = new(import3.TaskReference)
		}
		*p.oneOfType1010 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType1010.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType1010.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfImageNodeApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType1010 != nil && *p.oneOfType1010.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1010
	}
	return nil
}

func (p *OneOfImageNodeApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import4.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType1010 := new(import3.TaskReference)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType1010)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType1010.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType1010.ObjectType_ {
							if nil == p.oneOfType1010 {
								p.oneOfType1010 = new(import3.TaskReference)
							}
							*p.oneOfType1010 = *vOneOfType1010
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType1010.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType1010.ObjectType_
							return nil
						}
					}
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType1010 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType1010); err == nil {
		if vOneOfType1010.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType1010.ObjectType_ {
			if nil == p.oneOfType1010 {
				p.oneOfType1010 = new(import3.TaskReference)
			}
			*p.oneOfType1010 = *vOneOfType1010
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType1010.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType1010.ObjectType_
			return nil
		}
	}
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfImageNodeApiResponseData"))
}

func (p *OneOfImageNodeApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType1010 != nil && *p.oneOfType1010.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1010)
	}
	return nil, errors.New("No value to marshal for OneOfImageNodeApiResponseData")
}

type OneOfGetNodeByClaimTokenApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType1018 *ClaimTokenNode        `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfGetNodeByClaimTokenApiResponseData() *OneOfGetNodeByClaimTokenApiResponseData {
	p := new(OneOfGetNodeByClaimTokenApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetNodeByClaimTokenApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetNodeByClaimTokenApiResponseData is nil"))
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
	case ClaimTokenNode:
		if nil == p.oneOfType1018 {
			p.oneOfType1018 = new(ClaimTokenNode)
		}
		*p.oneOfType1018 = v.(ClaimTokenNode)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType1018.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType1018.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfGetNodeByClaimTokenApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType1018 != nil && *p.oneOfType1018.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1018
	}
	return nil
}

func (p *OneOfGetNodeByClaimTokenApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import4.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType1018 := new(ClaimTokenNode)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType1018)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType1018.ObjectType_ != nil && "lifecycle.v4.config.ClaimTokenNode" == *vOneOfType1018.ObjectType_ {
							if nil == p.oneOfType1018 {
								p.oneOfType1018 = new(ClaimTokenNode)
							}
							*p.oneOfType1018 = *vOneOfType1018
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType1018.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType1018.ObjectType_
							return nil
						}
					}
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType1018 := new(ClaimTokenNode)
	if err := json.Unmarshal(b, vOneOfType1018); err == nil {
		if vOneOfType1018.ObjectType_ != nil && "lifecycle.v4.config.ClaimTokenNode" == *vOneOfType1018.ObjectType_ {
			if nil == p.oneOfType1018 {
				p.oneOfType1018 = new(ClaimTokenNode)
			}
			*p.oneOfType1018 = *vOneOfType1018
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType1018.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType1018.ObjectType_
			return nil
		}
	}
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetNodeByClaimTokenApiResponseData"))
}

func (p *OneOfGetNodeByClaimTokenApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType1018 != nil && *p.oneOfType1018.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1018)
	}
	return nil, errors.New("No value to marshal for OneOfGetNodeByClaimTokenApiResponseData")
}

type OneOfUpdateConnectionApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType1020 *import3.TaskReference `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfUpdateConnectionApiResponseData() *OneOfUpdateConnectionApiResponseData {
	p := new(OneOfUpdateConnectionApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateConnectionApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateConnectionApiResponseData is nil"))
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
	case import3.TaskReference:
		if nil == p.oneOfType1020 {
			p.oneOfType1020 = new(import3.TaskReference)
		}
		*p.oneOfType1020 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType1020.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType1020.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfUpdateConnectionApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType1020 != nil && *p.oneOfType1020.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1020
	}
	return nil
}

func (p *OneOfUpdateConnectionApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import4.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType1020 := new(import3.TaskReference)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType1020)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType1020.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType1020.ObjectType_ {
							if nil == p.oneOfType1020 {
								p.oneOfType1020 = new(import3.TaskReference)
							}
							*p.oneOfType1020 = *vOneOfType1020
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType1020.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType1020.ObjectType_
							return nil
						}
					}
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType1020 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType1020); err == nil {
		if vOneOfType1020.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType1020.ObjectType_ {
			if nil == p.oneOfType1020 {
				p.oneOfType1020 = new(import3.TaskReference)
			}
			*p.oneOfType1020 = *vOneOfType1020
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType1020.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType1020.ObjectType_
			return nil
		}
	}
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateConnectionApiResponseData"))
}

func (p *OneOfUpdateConnectionApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType1020 != nil && *p.oneOfType1020.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1020)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateConnectionApiResponseData")
}

type OneOfUpdateImageApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType4016 *import3.TaskReference `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfUpdateImageApiResponseData() *OneOfUpdateImageApiResponseData {
	p := new(OneOfUpdateImageApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateImageApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateImageApiResponseData is nil"))
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
	case import3.TaskReference:
		if nil == p.oneOfType4016 {
			p.oneOfType4016 = new(import3.TaskReference)
		}
		*p.oneOfType4016 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType4016.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType4016.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfUpdateImageApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType4016 != nil && *p.oneOfType4016.ObjectType_ == *p.Discriminator {
		return *p.oneOfType4016
	}
	return nil
}

func (p *OneOfUpdateImageApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import4.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType4016 := new(import3.TaskReference)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType4016)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType4016.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType4016.ObjectType_ {
							if nil == p.oneOfType4016 {
								p.oneOfType4016 = new(import3.TaskReference)
							}
							*p.oneOfType4016 = *vOneOfType4016
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType4016.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType4016.ObjectType_
							return nil
						}
					}
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType4016 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType4016); err == nil {
		if vOneOfType4016.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType4016.ObjectType_ {
			if nil == p.oneOfType4016 {
				p.oneOfType4016 = new(import3.TaskReference)
			}
			*p.oneOfType4016 = *vOneOfType4016
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType4016.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType4016.ObjectType_
			return nil
		}
	}
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateImageApiResponseData"))
}

func (p *OneOfUpdateImageApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType4016 != nil && *p.oneOfType4016.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType4016)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateImageApiResponseData")
}

type OneOfGetClaimTokenApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType1005 *ClaimToken            `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfGetClaimTokenApiResponseData() *OneOfGetClaimTokenApiResponseData {
	p := new(OneOfGetClaimTokenApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetClaimTokenApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetClaimTokenApiResponseData is nil"))
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
	case ClaimToken:
		if nil == p.oneOfType1005 {
			p.oneOfType1005 = new(ClaimToken)
		}
		*p.oneOfType1005 = v.(ClaimToken)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType1005.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType1005.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfGetClaimTokenApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType1005 != nil && *p.oneOfType1005.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1005
	}
	return nil
}

func (p *OneOfGetClaimTokenApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import4.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType1005 := new(ClaimToken)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType1005)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType1005.ObjectType_ != nil && "lifecycle.v4.config.ClaimToken" == *vOneOfType1005.ObjectType_ {
							if nil == p.oneOfType1005 {
								p.oneOfType1005 = new(ClaimToken)
							}
							*p.oneOfType1005 = *vOneOfType1005
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType1005.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType1005.ObjectType_
							return nil
						}
					}
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType1005 := new(ClaimToken)
	if err := json.Unmarshal(b, vOneOfType1005); err == nil {
		if vOneOfType1005.ObjectType_ != nil && "lifecycle.v4.config.ClaimToken" == *vOneOfType1005.ObjectType_ {
			if nil == p.oneOfType1005 {
				p.oneOfType1005 = new(ClaimToken)
			}
			*p.oneOfType1005 = *vOneOfType1005
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType1005.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType1005.ObjectType_
			return nil
		}
	}
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetClaimTokenApiResponseData"))
}

func (p *OneOfGetClaimTokenApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType1005 != nil && *p.oneOfType1005.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1005)
	}
	return nil, errors.New("No value to marshal for OneOfGetClaimTokenApiResponseData")
}

type OneOfImageChecksum struct {
	Discriminator *string              `json:"-"`
	ObjectType_   *string              `json:"-"`
	oneOfType1020 *ImageSha256Checksum `json:"-"`
	oneOfType1021 *ImageMd5Checksum    `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfImageChecksum() *OneOfImageChecksum {
	p := new(OneOfImageChecksum)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfImageChecksum) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfImageChecksum is nil"))
	}
	switch v.(type) {
	case ImageSha256Checksum:
		if nil == p.oneOfType1020 {
			p.oneOfType1020 = new(ImageSha256Checksum)
		}
		*p.oneOfType1020 = v.(ImageSha256Checksum)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType1020.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType1020.ObjectType_
	case ImageMd5Checksum:
		if nil == p.oneOfType1021 {
			p.oneOfType1021 = new(ImageMd5Checksum)
		}
		*p.oneOfType1021 = v.(ImageMd5Checksum)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType1021.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType1021.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfImageChecksum) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType1020 != nil && *p.oneOfType1020.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1020
	}
	if p.oneOfType1021 != nil && *p.oneOfType1021.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1021
	}
	return nil
}

func (p *OneOfImageChecksum) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType1020 := new(ImageSha256Checksum)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType1020)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType1020.ObjectType_ != nil && "lifecycle.v4.config.ImageSha256Checksum" == *vOneOfType1020.ObjectType_ {
							if nil == p.oneOfType1020 {
								p.oneOfType1020 = new(ImageSha256Checksum)
							}
							*p.oneOfType1020 = *vOneOfType1020
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType1020.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType1020.ObjectType_
							return nil
						}
					}
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType1021 := new(ImageMd5Checksum)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType1021)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType1021.ObjectType_ != nil && "lifecycle.v4.config.ImageMd5Checksum" == *vOneOfType1021.ObjectType_ {
							if nil == p.oneOfType1021 {
								p.oneOfType1021 = new(ImageMd5Checksum)
							}
							*p.oneOfType1021 = *vOneOfType1021
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType1021.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType1021.ObjectType_
							return nil
						}
					}
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType1020 := new(ImageSha256Checksum)
	if err := json.Unmarshal(b, vOneOfType1020); err == nil {
		if vOneOfType1020.ObjectType_ != nil && "lifecycle.v4.config.ImageSha256Checksum" == *vOneOfType1020.ObjectType_ {
			if nil == p.oneOfType1020 {
				p.oneOfType1020 = new(ImageSha256Checksum)
			}
			*p.oneOfType1020 = *vOneOfType1020
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType1020.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType1020.ObjectType_
			return nil
		}
	}
	vOneOfType1021 := new(ImageMd5Checksum)
	if err := json.Unmarshal(b, vOneOfType1021); err == nil {
		if vOneOfType1021.ObjectType_ != nil && "lifecycle.v4.config.ImageMd5Checksum" == *vOneOfType1021.ObjectType_ {
			if nil == p.oneOfType1021 {
				p.oneOfType1021 = new(ImageMd5Checksum)
			}
			*p.oneOfType1021 = *vOneOfType1021
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType1021.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType1021.ObjectType_
			return nil
		}
	}
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfImageChecksum"))
}

func (p *OneOfImageChecksum) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType1020 != nil && *p.oneOfType1020.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1020)
	}
	if p.oneOfType1021 != nil && *p.oneOfType1021.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1021)
	}
	return nil, errors.New("No value to marshal for OneOfImageChecksum")
}

type OneOfUnifiedBmcIPConfigIpv4 struct {
	Discriminator *string         `json:"-"`
	ObjectType_   *string         `json:"-"`
	oneOfType5135 *IPPoolConfig   `json:"-"`
	oneOfType5136 *StaticIPConfig `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfUnifiedBmcIPConfigIpv4() *OneOfUnifiedBmcIPConfigIpv4 {
	p := new(OneOfUnifiedBmcIPConfigIpv4)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUnifiedBmcIPConfigIpv4) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUnifiedBmcIPConfigIpv4 is nil"))
	}
	switch v.(type) {
	case IPPoolConfig:
		if nil == p.oneOfType5135 {
			p.oneOfType5135 = new(IPPoolConfig)
		}
		*p.oneOfType5135 = v.(IPPoolConfig)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType5135.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType5135.ObjectType_
	case StaticIPConfig:
		if nil == p.oneOfType5136 {
			p.oneOfType5136 = new(StaticIPConfig)
		}
		*p.oneOfType5136 = v.(StaticIPConfig)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType5136.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType5136.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfUnifiedBmcIPConfigIpv4) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType5135 != nil && *p.oneOfType5135.ObjectType_ == *p.Discriminator {
		return *p.oneOfType5135
	}
	if p.oneOfType5136 != nil && *p.oneOfType5136.ObjectType_ == *p.Discriminator {
		return *p.oneOfType5136
	}
	return nil
}

func (p *OneOfUnifiedBmcIPConfigIpv4) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType5135 := new(IPPoolConfig)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType5135)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType5135.ObjectType_ != nil && "lifecycle.v4.config.IPPoolConfig" == *vOneOfType5135.ObjectType_ {
							if nil == p.oneOfType5135 {
								p.oneOfType5135 = new(IPPoolConfig)
							}
							*p.oneOfType5135 = *vOneOfType5135
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType5135.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType5135.ObjectType_
							return nil
						}
					}
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType5136 := new(StaticIPConfig)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType5136)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType5136.ObjectType_ != nil && "lifecycle.v4.config.StaticIPConfig" == *vOneOfType5136.ObjectType_ {
							if nil == p.oneOfType5136 {
								p.oneOfType5136 = new(StaticIPConfig)
							}
							*p.oneOfType5136 = *vOneOfType5136
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType5136.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType5136.ObjectType_
							return nil
						}
					}
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType5135 := new(IPPoolConfig)
	if err := json.Unmarshal(b, vOneOfType5135); err == nil {
		if vOneOfType5135.ObjectType_ != nil && "lifecycle.v4.config.IPPoolConfig" == *vOneOfType5135.ObjectType_ {
			if nil == p.oneOfType5135 {
				p.oneOfType5135 = new(IPPoolConfig)
			}
			*p.oneOfType5135 = *vOneOfType5135
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType5135.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType5135.ObjectType_
			return nil
		}
	}
	vOneOfType5136 := new(StaticIPConfig)
	if err := json.Unmarshal(b, vOneOfType5136); err == nil {
		if vOneOfType5136.ObjectType_ != nil && "lifecycle.v4.config.StaticIPConfig" == *vOneOfType5136.ObjectType_ {
			if nil == p.oneOfType5136 {
				p.oneOfType5136 = new(StaticIPConfig)
			}
			*p.oneOfType5136 = *vOneOfType5136
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType5136.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType5136.ObjectType_
			return nil
		}
	}
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUnifiedBmcIPConfigIpv4"))
}

func (p *OneOfUnifiedBmcIPConfigIpv4) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType5135 != nil && *p.oneOfType5135.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType5135)
	}
	if p.oneOfType5136 != nil && *p.oneOfType5136.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType5136)
	}
	return nil, errors.New("No value to marshal for OneOfUnifiedBmcIPConfigIpv4")
}

type OneOfDeleteImageApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType4008 *import3.TaskReference `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfDeleteImageApiResponseData() *OneOfDeleteImageApiResponseData {
	p := new(OneOfDeleteImageApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteImageApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteImageApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.TaskReference:
		if nil == p.oneOfType4008 {
			p.oneOfType4008 = new(import3.TaskReference)
		}
		*p.oneOfType4008 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType4008.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType4008.ObjectType_
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

func (p *OneOfDeleteImageApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType4008 != nil && *p.oneOfType4008.ObjectType_ == *p.Discriminator {
		return *p.oneOfType4008
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDeleteImageApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType4008 := new(import3.TaskReference)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType4008)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType4008.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType4008.ObjectType_ {
							if nil == p.oneOfType4008 {
								p.oneOfType4008 = new(import3.TaskReference)
							}
							*p.oneOfType4008 = *vOneOfType4008
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType4008.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType4008.ObjectType_
							return nil
						}
					}
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import4.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType4008 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType4008); err == nil {
		if vOneOfType4008.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType4008.ObjectType_ {
			if nil == p.oneOfType4008 {
				p.oneOfType4008 = new(import3.TaskReference)
			}
			*p.oneOfType4008 = *vOneOfType4008
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType4008.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType4008.ObjectType_
			return nil
		}
	}
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteImageApiResponseData"))
}

func (p *OneOfDeleteImageApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType4008 != nil && *p.oneOfType4008.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType4008)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteImageApiResponseData")
}

type OneOfListConnectionsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType401  []ConnectionProjection `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType1013 []Connection           `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfListConnectionsApiResponseData() *OneOfListConnectionsApiResponseData {
	p := new(OneOfListConnectionsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListConnectionsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListConnectionsApiResponseData is nil"))
	}
	switch v.(type) {
	case []ConnectionProjection:
		p.oneOfType401 = v.([]ConnectionProjection)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<lifecycle.v4.config.ConnectionProjection>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<lifecycle.v4.config.ConnectionProjection>"
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
	case []Connection:
		p.oneOfType1013 = v.([]Connection)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<lifecycle.v4.config.Connection>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<lifecycle.v4.config.Connection>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListConnectionsApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if "List<lifecycle.v4.config.ConnectionProjection>" == *p.Discriminator {
		return p.oneOfType401
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<lifecycle.v4.config.Connection>" == *p.Discriminator {
		return p.oneOfType1013
	}
	return nil
}

func (p *OneOfListConnectionsApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["List<lifecycle.v4.config.ConnectionProjection>"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType401 := new([]ConnectionProjection)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType401)
					if unmarshalErr == nil {
						// For arrays, verify the array item ObjectType matches
						if vOneOfType401 == nil || len(*vOneOfType401) == 0 || ((*vOneOfType401)[0].ObjectType_ != nil && "lifecycle.v4.config.ConnectionProjection" == *((*vOneOfType401)[0].ObjectType_)) {
							p.oneOfType401 = *vOneOfType401
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = "List<lifecycle.v4.config.ConnectionProjection>"
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = "List<lifecycle.v4.config.ConnectionProjection>"
							return nil
						}
					}
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import4.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["List<lifecycle.v4.config.Connection>"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType1013 := new([]Connection)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType1013)
					if unmarshalErr == nil {
						// For arrays, verify the array item ObjectType matches
						if vOneOfType1013 == nil || len(*vOneOfType1013) == 0 || ((*vOneOfType1013)[0].ObjectType_ != nil && "lifecycle.v4.config.Connection" == *((*vOneOfType1013)[0].ObjectType_)) {
							p.oneOfType1013 = *vOneOfType1013
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = "List<lifecycle.v4.config.Connection>"
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = "List<lifecycle.v4.config.Connection>"
							return nil
						}
					}
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType401 := new([]ConnectionProjection)
	if err := json.Unmarshal(b, vOneOfType401); err == nil {
		if len(*vOneOfType401) == 0 || (vOneOfType401 != nil && (*vOneOfType401)[0].ObjectType_ != nil && "lifecycle.v4.config.ConnectionProjection" == *((*vOneOfType401)[0].ObjectType_)) {
			p.oneOfType401 = *vOneOfType401
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<lifecycle.v4.config.ConnectionProjection>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<lifecycle.v4.config.ConnectionProjection>"
			return nil
		}
	}
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType1013 := new([]Connection)
	if err := json.Unmarshal(b, vOneOfType1013); err == nil {
		if len(*vOneOfType1013) == 0 || (vOneOfType1013 != nil && (*vOneOfType1013)[0].ObjectType_ != nil && "lifecycle.v4.config.Connection" == *((*vOneOfType1013)[0].ObjectType_)) {
			p.oneOfType1013 = *vOneOfType1013
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<lifecycle.v4.config.Connection>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<lifecycle.v4.config.Connection>"
			return nil
		}
	}
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListConnectionsApiResponseData"))
}

func (p *OneOfListConnectionsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if "List<lifecycle.v4.config.ConnectionProjection>" == *p.Discriminator {
		return json.Marshal(p.oneOfType401)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<lifecycle.v4.config.Connection>" == *p.Discriminator {
		return json.Marshal(p.oneOfType1013)
	}
	return nil, errors.New("No value to marshal for OneOfListConnectionsApiResponseData")
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
