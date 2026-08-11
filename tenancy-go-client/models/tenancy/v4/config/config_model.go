/*
 * Generated file models/tenancy/v4/config/config_model.go.
 *
 * Product version: 4.0.1-alpha-1
 *
 * Part of the SP Central Tenant Management
 *
 * (c) 2026 Nutanix Inc.  All rights reserved
 *
 */

/*
  Manage tenants
*/
package config

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	import2 "github.com/nutanix/ntnx-api-golang-clients/tenancy-go-client/v4/models/common/v1/response"
	import5 "github.com/nutanix/ntnx-api-golang-clients/tenancy-go-client/v4/models/iam/v4/authz"
	import3 "github.com/nutanix/ntnx-api-golang-clients/tenancy-go-client/v4/models/prism/v4/config"
	import4 "github.com/nutanix/ntnx-api-golang-clients/tenancy-go-client/v4/models/tenancy/v4/common"
	import1 "github.com/nutanix/ntnx-api-golang-clients/tenancy-go-client/v4/models/tenancy/v4/error"
	"time"
)

/*
REST response for all response codes in API path /tenancy/v4.0.a1/config/domains Post operation
*/
type AddDomainApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfAddDomainApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *AddDomainApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias AddDomainApiResponse

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

func (p *AddDomainApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias AddDomainApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewAddDomainApiResponse()

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

func NewAddDomainApiResponse() *AddDomainApiResponse {
	p := new(AddDomainApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "tenancy.v4.config.AddDomainApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *AddDomainApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *AddDomainApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfAddDomainApiResponseData()
	}
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
REST response for all response codes in API path /tenancy/v4.0.a1/config/domain-memberships Post operation
*/
type CreateDomainMembershipApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateDomainMembershipApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *CreateDomainMembershipApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias CreateDomainMembershipApiResponse

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

func (p *CreateDomainMembershipApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias CreateDomainMembershipApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewCreateDomainMembershipApiResponse()

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

func NewCreateDomainMembershipApiResponse() *CreateDomainMembershipApiResponse {
	p := new(CreateDomainMembershipApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "tenancy.v4.config.CreateDomainMembershipApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreateDomainMembershipApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreateDomainMembershipApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreateDomainMembershipApiResponseData()
	}
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
REST response for all response codes in API path /tenancy/v4.0.a1/config/tenants Post operation
*/
type CreateTenantApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateTenantApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *CreateTenantApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias CreateTenantApiResponse

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

func (p *CreateTenantApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias CreateTenantApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewCreateTenantApiResponse()

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

func NewCreateTenantApiResponse() *CreateTenantApiResponse {
	p := new(CreateTenantApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "tenancy.v4.config.CreateTenantApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreateTenantApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreateTenantApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreateTenantApiResponseData()
	}
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
REST response for all response codes in API path /tenancy/v4.0.a1/config/tenants/{tenantExtId}/domain-allocations Post operation
*/
type CreateTenantDomainAllocationApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateTenantDomainAllocationApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *CreateTenantDomainAllocationApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias CreateTenantDomainAllocationApiResponse

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

func (p *CreateTenantDomainAllocationApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias CreateTenantDomainAllocationApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewCreateTenantDomainAllocationApiResponse()

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

func NewCreateTenantDomainAllocationApiResponse() *CreateTenantDomainAllocationApiResponse {
	p := new(CreateTenantDomainAllocationApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "tenancy.v4.config.CreateTenantDomainAllocationApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreateTenantDomainAllocationApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreateTenantDomainAllocationApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreateTenantDomainAllocationApiResponseData()
	}
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
REST response for all response codes in API path /tenancy/v4.0.a1/config/domain-memberships/{extId} Delete operation
*/
type DeleteDomainMembershipApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteDomainMembershipApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *DeleteDomainMembershipApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DeleteDomainMembershipApiResponse

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

func (p *DeleteDomainMembershipApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DeleteDomainMembershipApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewDeleteDomainMembershipApiResponse()

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

func NewDeleteDomainMembershipApiResponse() *DeleteDomainMembershipApiResponse {
	p := new(DeleteDomainMembershipApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "tenancy.v4.config.DeleteDomainMembershipApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteDomainMembershipApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteDomainMembershipApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteDomainMembershipApiResponseData()
	}
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
REST response for all response codes in API path /tenancy/v4.0.a1/config/tenants/{extId} Delete operation
*/
type DeleteTenantApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteTenantApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *DeleteTenantApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DeleteTenantApiResponse

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

func (p *DeleteTenantApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DeleteTenantApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewDeleteTenantApiResponse()

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

func NewDeleteTenantApiResponse() *DeleteTenantApiResponse {
	p := new(DeleteTenantApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "tenancy.v4.config.DeleteTenantApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteTenantApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteTenantApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteTenantApiResponseData()
	}
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
REST response for all response codes in API path /tenancy/v4.0.a1/config/tenants/{tenantExtId}/domain-allocations/{extId} Delete operation
*/
type DeleteTenantDomainAllocationApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteTenantDomainAllocationApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *DeleteTenantDomainAllocationApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DeleteTenantDomainAllocationApiResponse

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

func (p *DeleteTenantDomainAllocationApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DeleteTenantDomainAllocationApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewDeleteTenantDomainAllocationApiResponse()

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

func NewDeleteTenantDomainAllocationApiResponse() *DeleteTenantDomainAllocationApiResponse {
	p := new(DeleteTenantDomainAllocationApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "tenancy.v4.config.DeleteTenantDomainAllocationApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteTenantDomainAllocationApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteTenantDomainAllocationApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteTenantDomainAllocationApiResponseData()
	}
	e := p.Data.SetValue(v)
	if nil == e {
		if nil == p.DataItemDiscriminator_ {
			p.DataItemDiscriminator_ = new(string)
		}
		*p.DataItemDiscriminator_ = *p.Data.Discriminator
	}
	return e
}

type Domain struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  External identifier of the domain.
	*/
	DomainExtId *string `json:"domainExtId"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *Domain) MarshalJSON() ([]byte, error) {
	type DomainProxy Domain

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*DomainProxy
		DomainExtId *string `json:"domainExtId,omitempty"`
	}{
		DomainProxy: (*DomainProxy)(p),
		DomainExtId: p.DomainExtId,
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

func (p *Domain) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Domain
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewDomain()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DomainExtId != nil {
		p.DomainExtId = known.DomainExtId
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "domainExtId")
	delete(allFields, "extId")
	delete(allFields, "links")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewDomain() *Domain {
	p := new(Domain)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "tenancy.v4.config.Domain"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
A domain allocation associates a tenant with a domain and its set of projects.
*/
type DomainAllocation struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  External identifier of the domain.
	*/
	DomainExtId *string `json:"domainExtId,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`

	ExternalNetworkConfig *ExternalNetworkConnection `json:"externalNetworkConfig,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  A project owns the resource allocation and placement settings within a domain allocation.
	*/
	Projects []Project `json:"projects,omitempty"`

	StateInfo *DomainAllocationStateInfo `json:"stateInfo,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *DomainAllocation) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DomainAllocation

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

func (p *DomainAllocation) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DomainAllocation
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewDomainAllocation()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DomainExtId != nil {
		p.DomainExtId = known.DomainExtId
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.ExternalNetworkConfig != nil {
		p.ExternalNetworkConfig = known.ExternalNetworkConfig
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.Projects != nil {
		p.Projects = known.Projects
	}
	if known.StateInfo != nil {
		p.StateInfo = known.StateInfo
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "domainExtId")
	delete(allFields, "extId")
	delete(allFields, "externalNetworkConfig")
	delete(allFields, "links")
	delete(allFields, "projects")
	delete(allFields, "stateInfo")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewDomainAllocation() *DomainAllocation {
	p := new(DomainAllocation)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "tenancy.v4.config.DomainAllocation"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type DomainAllocationStateInfo struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Error *import4.ErrorResponse `json:"error,omitempty"`

	State *DomainEntityState `json:"state,omitempty"`
}

func (p *DomainAllocationStateInfo) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DomainAllocationStateInfo

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

func (p *DomainAllocationStateInfo) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DomainAllocationStateInfo
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewDomainAllocationStateInfo()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Error != nil {
		p.Error = known.Error
	}
	if known.State != nil {
		p.State = known.State
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "error")
	delete(allFields, "state")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewDomainAllocationStateInfo() *DomainAllocationStateInfo {
	p := new(DomainAllocationStateInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "tenancy.v4.config.DomainAllocationStateInfo"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type DomainAllocationStateInfoProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Error *import4.ErrorResponse `json:"error,omitempty"`

	State *DomainEntityState `json:"state,omitempty"`
}

func (p *DomainAllocationStateInfoProjection) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DomainAllocationStateInfoProjection

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

func (p *DomainAllocationStateInfoProjection) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DomainAllocationStateInfoProjection
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewDomainAllocationStateInfoProjection()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Error != nil {
		p.Error = known.Error
	}
	if known.State != nil {
		p.State = known.State
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "error")
	delete(allFields, "state")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewDomainAllocationStateInfoProjection() *DomainAllocationStateInfoProjection {
	p := new(DomainAllocationStateInfoProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "tenancy.v4.config.DomainAllocationStateInfoProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
(Read Only Field) The current state of the domain entity.
*/
type DomainEntityState int

const (
	DOMAINENTITYSTATE_UNKNOWN              DomainEntityState = 0
	DOMAINENTITYSTATE_REDACTED             DomainEntityState = 1
	DOMAINENTITYSTATE_DELETED              DomainEntityState = 2
	DOMAINENTITYSTATE_DELETE_FAILED        DomainEntityState = 3
	DOMAINENTITYSTATE_DELETING_IN_PROGRESS DomainEntityState = 4
	DOMAINENTITYSTATE_PENDING              DomainEntityState = 5
	DOMAINENTITYSTATE_ACTIVE               DomainEntityState = 6
	DOMAINENTITYSTATE_PROVISIONING         DomainEntityState = 7
	DOMAINENTITYSTATE_FAILED               DomainEntityState = 8
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *DomainEntityState) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"DELETED",
		"DELETE_FAILED",
		"DELETING_IN_PROGRESS",
		"PENDING",
		"ACTIVE",
		"PROVISIONING",
		"FAILED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e DomainEntityState) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"DELETED",
		"DELETE_FAILED",
		"DELETING_IN_PROGRESS",
		"PENDING",
		"ACTIVE",
		"PROVISIONING",
		"FAILED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *DomainEntityState) index(name string) DomainEntityState {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"DELETED",
		"DELETE_FAILED",
		"DELETING_IN_PROGRESS",
		"PENDING",
		"ACTIVE",
		"PROVISIONING",
		"FAILED",
	}
	for idx := range names {
		if names[idx] == name {
			return DomainEntityState(idx)
		}
	}
	return DOMAINENTITYSTATE_UNKNOWN
}

func (e *DomainEntityState) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for DomainEntityState:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *DomainEntityState) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e DomainEntityState) Ref() *DomainEntityState {
	return &e
}

/*
(Read Only Field) The provisioning state of the external network configuration.
*/
type ExternalNetworkConfigState int

const (
	EXTERNALNETWORKCONFIGSTATE_UNKNOWN        ExternalNetworkConfigState = 0
	EXTERNALNETWORKCONFIGSTATE_REDACTED       ExternalNetworkConfigState = 1
	EXTERNALNETWORKCONFIGSTATE_NOT_CONFIGURED ExternalNetworkConfigState = 2
	EXTERNALNETWORKCONFIGSTATE_CONFIGURING    ExternalNetworkConfigState = 3
	EXTERNALNETWORKCONFIGSTATE_CONFIGURED     ExternalNetworkConfigState = 4
	EXTERNALNETWORKCONFIGSTATE_FAILED         ExternalNetworkConfigState = 5
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ExternalNetworkConfigState) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NOT_CONFIGURED",
		"CONFIGURING",
		"CONFIGURED",
		"FAILED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ExternalNetworkConfigState) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NOT_CONFIGURED",
		"CONFIGURING",
		"CONFIGURED",
		"FAILED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ExternalNetworkConfigState) index(name string) ExternalNetworkConfigState {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NOT_CONFIGURED",
		"CONFIGURING",
		"CONFIGURED",
		"FAILED",
	}
	for idx := range names {
		if names[idx] == name {
			return ExternalNetworkConfigState(idx)
		}
	}
	return EXTERNALNETWORKCONFIGSTATE_UNKNOWN
}

func (e *ExternalNetworkConfigState) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ExternalNetworkConfigState:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ExternalNetworkConfigState) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ExternalNetworkConfigState) Ref() *ExternalNetworkConfigState {
	return &e
}

type ExternalNetworkConfigStateInfo struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Error *import4.ErrorResponse `json:"error,omitempty"`

	State *ExternalNetworkConfigState `json:"state,omitempty"`
}

func (p *ExternalNetworkConfigStateInfo) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ExternalNetworkConfigStateInfo

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

func (p *ExternalNetworkConfigStateInfo) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ExternalNetworkConfigStateInfo
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewExternalNetworkConfigStateInfo()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Error != nil {
		p.Error = known.Error
	}
	if known.State != nil {
		p.State = known.State
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "error")
	delete(allFields, "state")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewExternalNetworkConfigStateInfo() *ExternalNetworkConfigStateInfo {
	p := new(ExternalNetworkConfigStateInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "tenancy.v4.config.ExternalNetworkConfigStateInfo"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type ExternalNetworkConfigStateInfoProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Error *import4.ErrorResponse `json:"error,omitempty"`

	State *ExternalNetworkConfigState `json:"state,omitempty"`
}

func (p *ExternalNetworkConfigStateInfoProjection) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ExternalNetworkConfigStateInfoProjection

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

func (p *ExternalNetworkConfigStateInfoProjection) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ExternalNetworkConfigStateInfoProjection
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewExternalNetworkConfigStateInfoProjection()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Error != nil {
		p.Error = known.Error
	}
	if known.State != nil {
		p.State = known.State
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "error")
	delete(allFields, "state")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewExternalNetworkConfigStateInfoProjection() *ExternalNetworkConfigStateInfoProjection {
	p := new(ExternalNetworkConfigStateInfoProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "tenancy.v4.config.ExternalNetworkConfigStateInfoProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
External network connection settings for a tenant.
*/
type ExternalNetworkConnConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Nat *NatConfig `json:"nat,omitempty"`

	NoNat *NoNatConfig `json:"noNat,omitempty"`
}

func (p *ExternalNetworkConnConfig) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ExternalNetworkConnConfig

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

func (p *ExternalNetworkConnConfig) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ExternalNetworkConnConfig
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewExternalNetworkConnConfig()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Nat != nil {
		p.Nat = known.Nat
	}
	if known.NoNat != nil {
		p.NoNat = known.NoNat
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "nat")
	delete(allFields, "noNat")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewExternalNetworkConnConfig() *ExternalNetworkConnConfig {
	p := new(ExternalNetworkConnConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "tenancy.v4.config.ExternalNetworkConnConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type ExternalNetworkConnConfigProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Nat *NatConfig `json:"nat,omitempty"`

	NoNat *NoNatConfig `json:"noNat,omitempty"`
}

func (p *ExternalNetworkConnConfigProjection) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ExternalNetworkConnConfigProjection

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

func (p *ExternalNetworkConnConfigProjection) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ExternalNetworkConnConfigProjection
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewExternalNetworkConnConfigProjection()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Nat != nil {
		p.Nat = known.Nat
	}
	if known.NoNat != nil {
		p.NoNat = known.NoNat
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "nat")
	delete(allFields, "noNat")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewExternalNetworkConnConfigProjection() *ExternalNetworkConnConfigProjection {
	p := new(ExternalNetworkConnConfigProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "tenancy.v4.config.ExternalNetworkConnConfigProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Response model for tenant external network configuration, containing NAT/no-NAT settings and provisioning status.
*/
type ExternalNetworkConnection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Config *ExternalNetworkConnConfig `json:"config,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`

	StateInfo *ExternalNetworkConfigStateInfo `json:"stateInfo,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *ExternalNetworkConnection) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ExternalNetworkConnection

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

func (p *ExternalNetworkConnection) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ExternalNetworkConnection
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewExternalNetworkConnection()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Config != nil {
		p.Config = known.Config
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.StateInfo != nil {
		p.StateInfo = known.StateInfo
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "config")
	delete(allFields, "extId")
	delete(allFields, "links")
	delete(allFields, "stateInfo")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewExternalNetworkConnection() *ExternalNetworkConnection {
	p := new(ExternalNetworkConnection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "tenancy.v4.config.ExternalNetworkConnection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type ExternalNetworkConnectionProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Config *ExternalNetworkConnConfig `json:"config,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`

	StateInfo *ExternalNetworkConfigStateInfo `json:"stateInfo,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *ExternalNetworkConnectionProjection) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ExternalNetworkConnectionProjection

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

func (p *ExternalNetworkConnectionProjection) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ExternalNetworkConnectionProjection
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewExternalNetworkConnectionProjection()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Config != nil {
		p.Config = known.Config
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.StateInfo != nil {
		p.StateInfo = known.StateInfo
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "config")
	delete(allFields, "extId")
	delete(allFields, "links")
	delete(allFields, "stateInfo")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewExternalNetworkConnectionProjection() *ExternalNetworkConnectionProjection {
	p := new(ExternalNetworkConnectionProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "tenancy.v4.config.ExternalNetworkConnectionProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Identify the principal via an existing identityExtId or by providing detailed attributes for resolution from an external identity provider.
*/
type ExternalUserDetails struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The external identifier of the identity provider.
	*/
	IdentityProviderExtId *string `json:"identityProviderExtId"`
	/*
	  The unique lookup value for the user, such as a UPN, username, or email address.
	*/
	PrincipalValue *string `json:"principalValue"`

	Type *import5.RmIdentityType `json:"type"`
}

func (p *ExternalUserDetails) MarshalJSON() ([]byte, error) {
	type ExternalUserDetailsProxy ExternalUserDetails

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*ExternalUserDetailsProxy
		IdentityProviderExtId *string                 `json:"identityProviderExtId,omitempty"`
		PrincipalValue        *string                 `json:"principalValue,omitempty"`
		Type                  *import5.RmIdentityType `json:"type,omitempty"`
	}{
		ExternalUserDetailsProxy: (*ExternalUserDetailsProxy)(p),
		IdentityProviderExtId:    p.IdentityProviderExtId,
		PrincipalValue:           p.PrincipalValue,
		Type:                     p.Type,
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

func (p *ExternalUserDetails) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ExternalUserDetails
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewExternalUserDetails()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.IdentityProviderExtId != nil {
		p.IdentityProviderExtId = known.IdentityProviderExtId
	}
	if known.PrincipalValue != nil {
		p.PrincipalValue = known.PrincipalValue
	}
	if known.Type != nil {
		p.Type = known.Type
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "identityProviderExtId")
	delete(allFields, "principalValue")
	delete(allFields, "type")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewExternalUserDetails() *ExternalUserDetails {
	p := new(ExternalUserDetails)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "tenancy.v4.config.ExternalUserDetails"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /tenancy/v4.0.a1/config/domain-memberships/{extId} Get operation
*/
type GetDomainMembershipApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetDomainMembershipApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *GetDomainMembershipApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetDomainMembershipApiResponse

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

func (p *GetDomainMembershipApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetDomainMembershipApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewGetDomainMembershipApiResponse()

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

func NewGetDomainMembershipApiResponse() *GetDomainMembershipApiResponse {
	p := new(GetDomainMembershipApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "tenancy.v4.config.GetDomainMembershipApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetDomainMembershipApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetDomainMembershipApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetDomainMembershipApiResponseData()
	}
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
REST response for all response codes in API path /tenancy/v4.0.a1/config/tenants/{extId} Get operation
*/
type GetTenantApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetTenantApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *GetTenantApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetTenantApiResponse

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

func (p *GetTenantApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetTenantApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewGetTenantApiResponse()

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

func NewGetTenantApiResponse() *GetTenantApiResponse {
	p := new(GetTenantApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "tenancy.v4.config.GetTenantApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetTenantApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetTenantApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetTenantApiResponseData()
	}
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
REST response for all response codes in API path /tenancy/v4.0.a1/config/tenants/{tenantExtId}/domain-allocations/{extId} Get operation
*/
type GetTenantDomainAllocationApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetTenantDomainAllocationApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *GetTenantDomainAllocationApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetTenantDomainAllocationApiResponse

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

func (p *GetTenantDomainAllocationApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetTenantDomainAllocationApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewGetTenantDomainAllocationApiResponse()

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

func NewGetTenantDomainAllocationApiResponse() *GetTenantDomainAllocationApiResponse {
	p := new(GetTenantDomainAllocationApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "tenancy.v4.config.GetTenantDomainAllocationApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetTenantDomainAllocationApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetTenantDomainAllocationApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetTenantDomainAllocationApiResponseData()
	}
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
REST response for all response codes in API path /tenancy/v4.0.a1/config/tenants/{tenantExtId}/domain-allocations/{domainAllocationExtId}/external-network-connection Get operation
*/
type GetTenantExternalNetworkConnectionApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetTenantExternalNetworkConnectionApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *GetTenantExternalNetworkConnectionApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetTenantExternalNetworkConnectionApiResponse

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

func (p *GetTenantExternalNetworkConnectionApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetTenantExternalNetworkConnectionApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewGetTenantExternalNetworkConnectionApiResponse()

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

func NewGetTenantExternalNetworkConnectionApiResponse() *GetTenantExternalNetworkConnectionApiResponse {
	p := new(GetTenantExternalNetworkConnectionApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "tenancy.v4.config.GetTenantExternalNetworkConnectionApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetTenantExternalNetworkConnectionApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetTenantExternalNetworkConnectionApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetTenantExternalNetworkConnectionApiResponseData()
	}
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
REST response for all response codes in API path /tenancy/v4.0.a1/config/domain-memberships Get operation
*/
type ListDomainMembershipsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListDomainMembershipsApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ListDomainMembershipsApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ListDomainMembershipsApiResponse

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

func (p *ListDomainMembershipsApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ListDomainMembershipsApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewListDomainMembershipsApiResponse()

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

func NewListDomainMembershipsApiResponse() *ListDomainMembershipsApiResponse {
	p := new(ListDomainMembershipsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "tenancy.v4.config.ListDomainMembershipsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListDomainMembershipsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListDomainMembershipsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListDomainMembershipsApiResponseData()
	}
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
REST response for all response codes in API path /tenancy/v4.0.a1/config/domains Get operation
*/
type ListDomainsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListDomainsApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ListDomainsApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ListDomainsApiResponse

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

func (p *ListDomainsApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ListDomainsApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewListDomainsApiResponse()

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

func NewListDomainsApiResponse() *ListDomainsApiResponse {
	p := new(ListDomainsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "tenancy.v4.config.ListDomainsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListDomainsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListDomainsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListDomainsApiResponseData()
	}
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
REST response for all response codes in API path /tenancy/v4.0.a1/config/tenants Get operation
*/
type ListTenantsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListTenantsApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ListTenantsApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ListTenantsApiResponse

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

func (p *ListTenantsApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ListTenantsApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewListTenantsApiResponse()

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

func NewListTenantsApiResponse() *ListTenantsApiResponse {
	p := new(ListTenantsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "tenancy.v4.config.ListTenantsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListTenantsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListTenantsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListTenantsApiResponseData()
	}
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
REST response for all response codes in API path /tenancy/v4.0.a1/config/tenants/{tenantExtId}/domain-allocations Get operation
*/
type ListTenantsDomainAllocationsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListTenantsDomainAllocationsApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ListTenantsDomainAllocationsApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ListTenantsDomainAllocationsApiResponse

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

func (p *ListTenantsDomainAllocationsApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ListTenantsDomainAllocationsApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewListTenantsDomainAllocationsApiResponse()

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

func NewListTenantsDomainAllocationsApiResponse() *ListTenantsDomainAllocationsApiResponse {
	p := new(ListTenantsDomainAllocationsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "tenancy.v4.config.ListTenantsDomainAllocationsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListTenantsDomainAllocationsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListTenantsDomainAllocationsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListTenantsDomainAllocationsApiResponseData()
	}
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
Local user information used to provision the initial tenant admin.
*/
type LocalUserInfo struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The email address of the user.
	*/
	Email *string `json:"email,omitempty"`
	/*
	  The first name of the user.
	*/
	FirstName *string `json:"firstName"`
	/*
	  The last name of the user.
	*/
	LastName *string `json:"lastName"`
	/*
	  The password of the user.
	*/
	Password *string `json:"password"`
	/*
	  The username of the user.
	*/
	Username *string `json:"username"`
}

func (p *LocalUserInfo) MarshalJSON() ([]byte, error) {
	type LocalUserInfoProxy LocalUserInfo

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*LocalUserInfoProxy
		FirstName *string `json:"firstName,omitempty"`
		LastName  *string `json:"lastName,omitempty"`
		Password  *string `json:"password,omitempty"`
		Username  *string `json:"username,omitempty"`
	}{
		LocalUserInfoProxy: (*LocalUserInfoProxy)(p),
		FirstName:          p.FirstName,
		LastName:           p.LastName,
		Password:           p.Password,
		Username:           p.Username,
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

func (p *LocalUserInfo) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias LocalUserInfo
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewLocalUserInfo()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Email != nil {
		p.Email = known.Email
	}
	if known.FirstName != nil {
		p.FirstName = known.FirstName
	}
	if known.LastName != nil {
		p.LastName = known.LastName
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
	delete(allFields, "email")
	delete(allFields, "firstName")
	delete(allFields, "lastName")
	delete(allFields, "password")
	delete(allFields, "username")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewLocalUserInfo() *LocalUserInfo {
	p := new(LocalUserInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "tenancy.v4.config.LocalUserInfo"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
A membership assigns a specific role to a principal (user or group) within a particular domain.
*/
type Membership struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The username of the user who created the membership.
	*/
	CreatedBy *string `json:"createdBy,omitempty"`
	/*
	  The date and time when the membership was created.
	*/
	CreatedTime *time.Time `json:"createdTime,omitempty"`
	/*
	  External identifier of the domain.
	*/
	DomainExtId *string `json:"domainExtId"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`

	IdentityDetails *ExternalUserDetails `json:"identityDetails,omitempty"`
	/*
	  The external identifier of an existing principal entity.
	*/
	IdentityExtId *string `json:"identityExtId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  The resolved principal name associated with this membership.
	*/
	Principal *string `json:"principal,omitempty"`
	/*
	  The external identifier of the role assigned to the principal.
	*/
	RoleExtId *string `json:"roleExtId"`

	StateInfo *MembershipStateInfo `json:"stateInfo,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *Membership) MarshalJSON() ([]byte, error) {
	type MembershipProxy Membership

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*MembershipProxy
		DomainExtId *string `json:"domainExtId,omitempty"`
		RoleExtId   *string `json:"roleExtId,omitempty"`
	}{
		MembershipProxy: (*MembershipProxy)(p),
		DomainExtId:     p.DomainExtId,
		RoleExtId:       p.RoleExtId,
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

func (p *Membership) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Membership
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewMembership()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.CreatedBy != nil {
		p.CreatedBy = known.CreatedBy
	}
	if known.CreatedTime != nil {
		p.CreatedTime = known.CreatedTime
	}
	if known.DomainExtId != nil {
		p.DomainExtId = known.DomainExtId
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.IdentityDetails != nil {
		p.IdentityDetails = known.IdentityDetails
	}
	if known.IdentityExtId != nil {
		p.IdentityExtId = known.IdentityExtId
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.Principal != nil {
		p.Principal = known.Principal
	}
	if known.RoleExtId != nil {
		p.RoleExtId = known.RoleExtId
	}
	if known.StateInfo != nil {
		p.StateInfo = known.StateInfo
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "createdBy")
	delete(allFields, "createdTime")
	delete(allFields, "domainExtId")
	delete(allFields, "extId")
	delete(allFields, "identityDetails")
	delete(allFields, "identityExtId")
	delete(allFields, "links")
	delete(allFields, "principal")
	delete(allFields, "roleExtId")
	delete(allFields, "stateInfo")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewMembership() *Membership {
	p := new(Membership)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "tenancy.v4.config.Membership"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (i *Membership) HasIdentityExtId() bool {
	return i.IdentityExtId != nil
}
func (i *Membership) HasIdentityDetails() bool {
	return i.IdentityDetails != nil
}

func (i *Membership) IsValid() bool {
	return i.HasIdentityExtId() || i.HasIdentityDetails()
}

/*
The current lifecycle state of the membership.
*/
type MembershipState int

const (
	MEMBERSHIPSTATE_UNKNOWN  MembershipState = 0
	MEMBERSHIPSTATE_REDACTED MembershipState = 1
	MEMBERSHIPSTATE_PENDING  MembershipState = 2
	MEMBERSHIPSTATE_ACTIVE   MembershipState = 3
	MEMBERSHIPSTATE_ERROR    MembershipState = 4
	MEMBERSHIPSTATE_DELETING MembershipState = 5
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *MembershipState) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PENDING",
		"ACTIVE",
		"ERROR",
		"DELETING",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e MembershipState) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PENDING",
		"ACTIVE",
		"ERROR",
		"DELETING",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *MembershipState) index(name string) MembershipState {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PENDING",
		"ACTIVE",
		"ERROR",
		"DELETING",
	}
	for idx := range names {
		if names[idx] == name {
			return MembershipState(idx)
		}
	}
	return MEMBERSHIPSTATE_UNKNOWN
}

func (e *MembershipState) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for MembershipState:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *MembershipState) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e MembershipState) Ref() *MembershipState {
	return &e
}

type MembershipStateInfo struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Error *import4.ErrorResponse `json:"error,omitempty"`

	State *MembershipState `json:"state,omitempty"`
}

func (p *MembershipStateInfo) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias MembershipStateInfo

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

func (p *MembershipStateInfo) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias MembershipStateInfo
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewMembershipStateInfo()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Error != nil {
		p.Error = known.Error
	}
	if known.State != nil {
		p.State = known.State
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "error")
	delete(allFields, "state")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewMembershipStateInfo() *MembershipStateInfo {
	p := new(MembershipStateInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "tenancy.v4.config.MembershipStateInfo"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type MembershipStateInfoProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Error *import4.ErrorResponse `json:"error,omitempty"`

	State *MembershipState `json:"state,omitempty"`
}

func (p *MembershipStateInfoProjection) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias MembershipStateInfoProjection

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

func (p *MembershipStateInfoProjection) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias MembershipStateInfoProjection
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewMembershipStateInfoProjection()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Error != nil {
		p.Error = known.Error
	}
	if known.State != nil {
		p.State = known.State
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "error")
	delete(allFields, "state")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewMembershipStateInfoProjection() *MembershipStateInfoProjection {
	p := new(MembershipStateInfoProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "tenancy.v4.config.MembershipStateInfoProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
NAT configuration settings.
*/
type NatConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Specifies the total number of floating IP addresses to be allocated in the NAT configuration. If floating IPs are already configured, only the additional required IPs (based on the specified value) will be created. The default value is 0.
	*/
	FloatingIpCount *int `json:"floatingIpCount,omitempty"`
	/*
	  The external identifier of the overlay subnet for NAT configuration.
	*/
	OverlaySubnetExtId *string `json:"overlaySubnetExtId"`
	/*
	  The external identifier of the transit VPC for NAT configuration.
	*/
	TransitVpcExtId *string `json:"transitVpcExtId"`
}

func (p *NatConfig) MarshalJSON() ([]byte, error) {
	type NatConfigProxy NatConfig

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*NatConfigProxy
		OverlaySubnetExtId *string `json:"overlaySubnetExtId,omitempty"`
		TransitVpcExtId    *string `json:"transitVpcExtId,omitempty"`
	}{
		NatConfigProxy:     (*NatConfigProxy)(p),
		OverlaySubnetExtId: p.OverlaySubnetExtId,
		TransitVpcExtId:    p.TransitVpcExtId,
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

func (p *NatConfig) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias NatConfig
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewNatConfig()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.FloatingIpCount != nil {
		p.FloatingIpCount = known.FloatingIpCount
	}
	if known.OverlaySubnetExtId != nil {
		p.OverlaySubnetExtId = known.OverlaySubnetExtId
	}
	if known.TransitVpcExtId != nil {
		p.TransitVpcExtId = known.TransitVpcExtId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "floatingIpCount")
	delete(allFields, "overlaySubnetExtId")
	delete(allFields, "transitVpcExtId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewNatConfig() *NatConfig {
	p := new(NatConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "tenancy.v4.config.NatConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type NatConfigProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Specifies the total number of floating IP addresses to be allocated in the NAT configuration. If floating IPs are already configured, only the additional required IPs (based on the specified value) will be created. The default value is 0.
	*/
	FloatingIpCount *int `json:"floatingIpCount,omitempty"`
	/*
	  The external identifier of the overlay subnet for NAT configuration.
	*/
	OverlaySubnetExtId *string `json:"overlaySubnetExtId"`
	/*
	  The external identifier of the transit VPC for NAT configuration.
	*/
	TransitVpcExtId *string `json:"transitVpcExtId"`
}

func (p *NatConfigProjection) MarshalJSON() ([]byte, error) {
	type NatConfigProjectionProxy NatConfigProjection

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*NatConfigProjectionProxy
		OverlaySubnetExtId *string `json:"overlaySubnetExtId,omitempty"`
		TransitVpcExtId    *string `json:"transitVpcExtId,omitempty"`
	}{
		NatConfigProjectionProxy: (*NatConfigProjectionProxy)(p),
		OverlaySubnetExtId:       p.OverlaySubnetExtId,
		TransitVpcExtId:          p.TransitVpcExtId,
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

func (p *NatConfigProjection) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias NatConfigProjection
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewNatConfigProjection()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.FloatingIpCount != nil {
		p.FloatingIpCount = known.FloatingIpCount
	}
	if known.OverlaySubnetExtId != nil {
		p.OverlaySubnetExtId = known.OverlaySubnetExtId
	}
	if known.TransitVpcExtId != nil {
		p.TransitVpcExtId = known.TransitVpcExtId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "floatingIpCount")
	delete(allFields, "overlaySubnetExtId")
	delete(allFields, "transitVpcExtId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewNatConfigProjection() *NatConfigProjection {
	p := new(NatConfigProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "tenancy.v4.config.NatConfigProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
No-NAT configuration settings.
*/
type NoNatConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Specifies the Externally Routable Prefix (ERP) CIDR IP addresses to be used in a no-NAT configuration. These entries are overwritten with each request. The CIDR addresses are normalized during processing.
	*/
	ExternallyRoutablePrefixes []string `json:"externallyRoutablePrefixes,omitempty"`
	/*
	  The external identifier of the overlay subnet for no-NAT configuration.
	*/
	OverlaySubnetExtId *string `json:"overlaySubnetExtId"`
	/*
	  The external identifier of the transit VPC for no-NAT configuration.
	*/
	TransitVpcExtId *string `json:"transitVpcExtId"`
}

func (p *NoNatConfig) MarshalJSON() ([]byte, error) {
	type NoNatConfigProxy NoNatConfig

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*NoNatConfigProxy
		OverlaySubnetExtId *string `json:"overlaySubnetExtId,omitempty"`
		TransitVpcExtId    *string `json:"transitVpcExtId,omitempty"`
	}{
		NoNatConfigProxy:   (*NoNatConfigProxy)(p),
		OverlaySubnetExtId: p.OverlaySubnetExtId,
		TransitVpcExtId:    p.TransitVpcExtId,
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

func (p *NoNatConfig) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias NoNatConfig
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewNoNatConfig()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ExternallyRoutablePrefixes != nil {
		p.ExternallyRoutablePrefixes = known.ExternallyRoutablePrefixes
	}
	if known.OverlaySubnetExtId != nil {
		p.OverlaySubnetExtId = known.OverlaySubnetExtId
	}
	if known.TransitVpcExtId != nil {
		p.TransitVpcExtId = known.TransitVpcExtId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "externallyRoutablePrefixes")
	delete(allFields, "overlaySubnetExtId")
	delete(allFields, "transitVpcExtId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewNoNatConfig() *NoNatConfig {
	p := new(NoNatConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "tenancy.v4.config.NoNatConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type NoNatConfigProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Specifies the Externally Routable Prefix (ERP) CIDR IP addresses to be used in a no-NAT configuration. These entries are overwritten with each request. The CIDR addresses are normalized during processing.
	*/
	ExternallyRoutablePrefixes []string `json:"externallyRoutablePrefixes,omitempty"`
	/*
	  The external identifier of the overlay subnet for no-NAT configuration.
	*/
	OverlaySubnetExtId *string `json:"overlaySubnetExtId"`
	/*
	  The external identifier of the transit VPC for no-NAT configuration.
	*/
	TransitVpcExtId *string `json:"transitVpcExtId"`
}

func (p *NoNatConfigProjection) MarshalJSON() ([]byte, error) {
	type NoNatConfigProjectionProxy NoNatConfigProjection

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*NoNatConfigProjectionProxy
		OverlaySubnetExtId *string `json:"overlaySubnetExtId,omitempty"`
		TransitVpcExtId    *string `json:"transitVpcExtId,omitempty"`
	}{
		NoNatConfigProjectionProxy: (*NoNatConfigProjectionProxy)(p),
		OverlaySubnetExtId:         p.OverlaySubnetExtId,
		TransitVpcExtId:            p.TransitVpcExtId,
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

func (p *NoNatConfigProjection) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias NoNatConfigProjection
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewNoNatConfigProjection()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ExternallyRoutablePrefixes != nil {
		p.ExternallyRoutablePrefixes = known.ExternallyRoutablePrefixes
	}
	if known.OverlaySubnetExtId != nil {
		p.OverlaySubnetExtId = known.OverlaySubnetExtId
	}
	if known.TransitVpcExtId != nil {
		p.TransitVpcExtId = known.TransitVpcExtId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "externallyRoutablePrefixes")
	delete(allFields, "overlaySubnetExtId")
	delete(allFields, "transitVpcExtId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewNoNatConfigProjection() *NoNatConfigProjection {
	p := new(NoNatConfigProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "tenancy.v4.config.NoNatConfigProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Placement target configuration for cluster and storage resources.
*/
type PlacementTarget struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The UUID of the cluster for resource placement.
	*/
	ClusterUuid *string `json:"clusterUuid,omitempty"`
	/*
	  List of storage container UUIDs for the placement target.
	*/
	StorageContainerUuids []string `json:"storageContainerUuids,omitempty"`
}

func (p *PlacementTarget) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias PlacementTarget

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

func (p *PlacementTarget) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias PlacementTarget
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewPlacementTarget()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ClusterUuid != nil {
		p.ClusterUuid = known.ClusterUuid
	}
	if known.StorageContainerUuids != nil {
		p.StorageContainerUuids = known.StorageContainerUuids
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "clusterUuid")
	delete(allFields, "storageContainerUuids")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewPlacementTarget() *PlacementTarget {
	p := new(PlacementTarget)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "tenancy.v4.config.PlacementTarget"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
A project owns the resource allocation and placement settings within a domain allocation.
*/
type Project struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  External identifier of the project.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  The name of the project.
	*/
	Name *string `json:"name"`

	Quota *Quota `json:"quota,omitempty"`

	ResourceGroup *ResourceGroup `json:"resourceGroup,omitempty"`

	State *DomainEntityState `json:"state,omitempty"`
}

func (p *Project) MarshalJSON() ([]byte, error) {
	type ProjectProxy Project

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*ProjectProxy
		Name *string `json:"name,omitempty"`
	}{
		ProjectProxy: (*ProjectProxy)(p),
		Name:         p.Name,
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

func (p *Project) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Project
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewProject()

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
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.Quota != nil {
		p.Quota = known.Quota
	}
	if known.ResourceGroup != nil {
		p.ResourceGroup = known.ResourceGroup
	}
	if known.State != nil {
		p.State = known.State
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "extId")
	delete(allFields, "name")
	delete(allFields, "quota")
	delete(allFields, "resourceGroup")
	delete(allFields, "state")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewProject() *Project {
	p := new(Project)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "tenancy.v4.config.Project"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Resource quota allocated to a project or aggregated for a tenant.
*/
type Quota struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The external identifier of the quota for the placement target.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  The amount of memory allocated as quota, in bytes.
	*/
	Memory *int64 `json:"memory,omitempty"`
	/*
	  The amount of storage allocated as quota, in bytes.
	*/
	Storage *int64 `json:"storage,omitempty"`
	/*
	  The number of vCPUs allocated as quota.
	*/
	Vcpu *int `json:"vcpu,omitempty"`
}

func (p *Quota) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias Quota

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

func (p *Quota) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Quota
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewQuota()

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
	if known.Memory != nil {
		p.Memory = known.Memory
	}
	if known.Storage != nil {
		p.Storage = known.Storage
	}
	if known.Vcpu != nil {
		p.Vcpu = known.Vcpu
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "extId")
	delete(allFields, "memory")
	delete(allFields, "storage")
	delete(allFields, "vcpu")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewQuota() *Quota {
	p := new(Quota)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "tenancy.v4.config.Quota"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Resource group containing placement targets for resource allocation.
*/
type ResourceGroup struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	PlacementTargets []PlacementTarget `json:"placementTargets,omitempty"`
	/*
	  The UUID of the resource group for the placement target.
	*/
	Uuid *string `json:"uuid,omitempty"`
}

func (p *ResourceGroup) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ResourceGroup

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

func (p *ResourceGroup) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ResourceGroup
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewResourceGroup()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.PlacementTargets != nil {
		p.PlacementTargets = known.PlacementTargets
	}
	if known.Uuid != nil {
		p.Uuid = known.Uuid
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "placementTargets")
	delete(allFields, "uuid")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewResourceGroup() *ResourceGroup {
	p := new(ResourceGroup)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "tenancy.v4.config.ResourceGroup"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
A tenant is an organization or group of users that consumes the underlying Service Provider infrastructure or platform.
*/
type Tenant struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The identifier of the user who created the tenant.
	*/
	AddedBy *string `json:"addedBy,omitempty"`

	AdminInfo *LocalUserInfo `json:"adminInfo,omitempty"`

	AggregatedQuota *Quota `json:"aggregatedQuota,omitempty"`
	/*
	  The timestamp when the tenant was created.
	*/
	CreatedTime *time.Time `json:"createdTime,omitempty"`
	/*
	  A brief description of the tenant.
	*/
	Description *string `json:"description,omitempty"`
	/*
	  List of external identifiers of domains associated with the tenant.
	*/
	Domains []string `json:"domains,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  The name of the tenant. It must start and end with a letter (including Unicode letters with combining marks) or a digit, and can contain hyphens in between.
	*/
	Name *string `json:"name,omitempty"`

	StateInfo *TenantStateInfo `json:"stateInfo,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  The URL of the tenant.
	*/
	Url *string `json:"url,omitempty"`
}

func (p *Tenant) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias Tenant

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

func (p *Tenant) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Tenant
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewTenant()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.AddedBy != nil {
		p.AddedBy = known.AddedBy
	}
	if known.AdminInfo != nil {
		p.AdminInfo = known.AdminInfo
	}
	if known.AggregatedQuota != nil {
		p.AggregatedQuota = known.AggregatedQuota
	}
	if known.CreatedTime != nil {
		p.CreatedTime = known.CreatedTime
	}
	if known.Description != nil {
		p.Description = known.Description
	}
	if known.Domains != nil {
		p.Domains = known.Domains
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
	if known.StateInfo != nil {
		p.StateInfo = known.StateInfo
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}
	if known.Url != nil {
		p.Url = known.Url
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "addedBy")
	delete(allFields, "adminInfo")
	delete(allFields, "aggregatedQuota")
	delete(allFields, "createdTime")
	delete(allFields, "description")
	delete(allFields, "domains")
	delete(allFields, "extId")
	delete(allFields, "links")
	delete(allFields, "name")
	delete(allFields, "stateInfo")
	delete(allFields, "tenantId")
	delete(allFields, "url")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewTenant() *Tenant {
	p := new(Tenant)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "tenancy.v4.config.Tenant"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
(Read Only Field) The current lifecycle state of the tenant.
*/
type TenantState int

const (
	TENANTSTATE_UNKNOWN              TenantState = 0
	TENANTSTATE_REDACTED             TenantState = 1
	TENANTSTATE_ACTIVE               TenantState = 2
	TENANTSTATE_PENDING              TenantState = 3
	TENANTSTATE_INACTIVE             TenantState = 4
	TENANTSTATE_DELETED              TenantState = 5
	TENANTSTATE_FAILED               TenantState = 6
	TENANTSTATE_SETUP_PENDING        TenantState = 7
	TENANTSTATE_DELETING_IN_PROGRESS TenantState = 8
	TENANTSTATE_DELETE_FAILED        TenantState = 9
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *TenantState) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ACTIVE",
		"PENDING",
		"INACTIVE",
		"DELETED",
		"FAILED",
		"SETUP_PENDING",
		"DELETING_IN_PROGRESS",
		"DELETE_FAILED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e TenantState) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ACTIVE",
		"PENDING",
		"INACTIVE",
		"DELETED",
		"FAILED",
		"SETUP_PENDING",
		"DELETING_IN_PROGRESS",
		"DELETE_FAILED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *TenantState) index(name string) TenantState {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ACTIVE",
		"PENDING",
		"INACTIVE",
		"DELETED",
		"FAILED",
		"SETUP_PENDING",
		"DELETING_IN_PROGRESS",
		"DELETE_FAILED",
	}
	for idx := range names {
		if names[idx] == name {
			return TenantState(idx)
		}
	}
	return TENANTSTATE_UNKNOWN
}

func (e *TenantState) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for TenantState:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *TenantState) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e TenantState) Ref() *TenantState {
	return &e
}

type TenantStateInfo struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Error *import4.ErrorResponse `json:"error,omitempty"`

	State *TenantState `json:"state,omitempty"`
}

func (p *TenantStateInfo) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias TenantStateInfo

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

func (p *TenantStateInfo) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias TenantStateInfo
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewTenantStateInfo()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Error != nil {
		p.Error = known.Error
	}
	if known.State != nil {
		p.State = known.State
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "error")
	delete(allFields, "state")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewTenantStateInfo() *TenantStateInfo {
	p := new(TenantStateInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "tenancy.v4.config.TenantStateInfo"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type TenantStateInfoProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Error *import4.ErrorResponse `json:"error,omitempty"`

	State *TenantState `json:"state,omitempty"`
}

func (p *TenantStateInfoProjection) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias TenantStateInfoProjection

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

func (p *TenantStateInfoProjection) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias TenantStateInfoProjection
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewTenantStateInfoProjection()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Error != nil {
		p.Error = known.Error
	}
	if known.State != nil {
		p.State = known.State
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "error")
	delete(allFields, "state")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewTenantStateInfoProjection() *TenantStateInfoProjection {
	p := new(TenantStateInfoProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "tenancy.v4.config.TenantStateInfoProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /tenancy/v4.0.a1/config/tenants/{tenantExtId}/domain-allocations/{extId} Put operation
*/
type UpdateTenantDomainAllocationApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateTenantDomainAllocationApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *UpdateTenantDomainAllocationApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias UpdateTenantDomainAllocationApiResponse

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

func (p *UpdateTenantDomainAllocationApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias UpdateTenantDomainAllocationApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewUpdateTenantDomainAllocationApiResponse()

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

func NewUpdateTenantDomainAllocationApiResponse() *UpdateTenantDomainAllocationApiResponse {
	p := new(UpdateTenantDomainAllocationApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "tenancy.v4.config.UpdateTenantDomainAllocationApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdateTenantDomainAllocationApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdateTenantDomainAllocationApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdateTenantDomainAllocationApiResponseData()
	}
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
REST response for all response codes in API path /tenancy/v4.0.a1/config/tenants/{tenantExtId}/domain-allocations/{domainAllocationExtId}/external-network-connection Put operation
*/
type UpdateTenantExternalNetworkConnectionApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateTenantExternalNetworkConnectionApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *UpdateTenantExternalNetworkConnectionApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias UpdateTenantExternalNetworkConnectionApiResponse

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

func (p *UpdateTenantExternalNetworkConnectionApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias UpdateTenantExternalNetworkConnectionApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewUpdateTenantExternalNetworkConnectionApiResponse()

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

func NewUpdateTenantExternalNetworkConnectionApiResponse() *UpdateTenantExternalNetworkConnectionApiResponse {
	p := new(UpdateTenantExternalNetworkConnectionApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "tenancy.v4.config.UpdateTenantExternalNetworkConnectionApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdateTenantExternalNetworkConnectionApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdateTenantExternalNetworkConnectionApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdateTenantExternalNetworkConnectionApiResponseData()
	}
	e := p.Data.SetValue(v)
	if nil == e {
		if nil == p.DataItemDiscriminator_ {
			p.DataItemDiscriminator_ = new(string)
		}
		*p.DataItemDiscriminator_ = *p.Data.Discriminator
	}
	return e
}

type OneOfUpdateTenantExternalNetworkConnectionApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfUpdateTenantExternalNetworkConnectionApiResponseData() *OneOfUpdateTenantExternalNetworkConnectionApiResponseData {
	p := new(OneOfUpdateTenantExternalNetworkConnectionApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateTenantExternalNetworkConnectionApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateTenantExternalNetworkConnectionApiResponseData is nil"))
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
	case import1.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import1.ErrorResponse)
		}
		*p.oneOfType400 = v.(import1.ErrorResponse)
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

func (p *OneOfUpdateTenantExternalNetworkConnectionApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfUpdateTenantExternalNetworkConnectionApiResponseData) UnmarshalJSON(b []byte) error {
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
					vOneOfType0 := new(import3.TaskReference)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType0)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType0.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
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
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import1.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "tenancy.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
							if nil == p.oneOfType400 {
								p.oneOfType400 = new(import1.ErrorResponse)
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
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if vOneOfType0.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
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
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "tenancy.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import1.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateTenantExternalNetworkConnectionApiResponseData"))
}

func (p *OneOfUpdateTenantExternalNetworkConnectionApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateTenantExternalNetworkConnectionApiResponseData")
}

type OneOfDeleteDomainMembershipApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	oneOfType0    *Membership            `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfDeleteDomainMembershipApiResponseData() *OneOfDeleteDomainMembershipApiResponseData {
	p := new(OneOfDeleteDomainMembershipApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteDomainMembershipApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteDomainMembershipApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import1.ErrorResponse)
		}
		*p.oneOfType400 = v.(import1.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case Membership:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(Membership)
		}
		*p.oneOfType0 = v.(Membership)
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

func (p *OneOfDeleteDomainMembershipApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfDeleteDomainMembershipApiResponseData) UnmarshalJSON(b []byte) error {
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
					vOneOfType400 := new(import1.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "tenancy.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
							if nil == p.oneOfType400 {
								p.oneOfType400 = new(import1.ErrorResponse)
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
					vOneOfType0 := new(Membership)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType0)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType0.ObjectType_ != nil && "tenancy.v4.config.Membership" == *vOneOfType0.ObjectType_ {
							if nil == p.oneOfType0 {
								p.oneOfType0 = new(Membership)
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
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "tenancy.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import1.ErrorResponse)
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
	vOneOfType0 := new(Membership)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if vOneOfType0.ObjectType_ != nil && "tenancy.v4.config.Membership" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(Membership)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteDomainMembershipApiResponseData"))
}

func (p *OneOfDeleteDomainMembershipApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteDomainMembershipApiResponseData")
}

type OneOfListTenantsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    []Tenant               `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfListTenantsApiResponseData() *OneOfListTenantsApiResponseData {
	p := new(OneOfListTenantsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListTenantsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListTenantsApiResponseData is nil"))
	}
	switch v.(type) {
	case []Tenant:
		p.oneOfType0 = v.([]Tenant)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<tenancy.v4.config.Tenant>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<tenancy.v4.config.Tenant>"
	case import1.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import1.ErrorResponse)
		}
		*p.oneOfType400 = v.(import1.ErrorResponse)
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

func (p *OneOfListTenantsApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if "List<tenancy.v4.config.Tenant>" == *p.Discriminator {
		return p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfListTenantsApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["List<tenancy.v4.config.Tenant>"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType0 := new([]Tenant)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType0)
					if unmarshalErr == nil {
						// For arrays, verify the array item ObjectType matches
						if vOneOfType0 == nil || len(*vOneOfType0) == 0 || ((*vOneOfType0)[0].ObjectType_ != nil && "tenancy.v4.config.Tenant" == *((*vOneOfType0)[0].ObjectType_)) {
							p.oneOfType0 = *vOneOfType0
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = "List<tenancy.v4.config.Tenant>"
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = "List<tenancy.v4.config.Tenant>"
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
					vOneOfType400 := new(import1.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "tenancy.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
							if nil == p.oneOfType400 {
								p.oneOfType400 = new(import1.ErrorResponse)
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
	vOneOfType0 := new([]Tenant)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || (vOneOfType0 != nil && (*vOneOfType0)[0].ObjectType_ != nil && "tenancy.v4.config.Tenant" == *((*vOneOfType0)[0].ObjectType_)) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<tenancy.v4.config.Tenant>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<tenancy.v4.config.Tenant>"
			return nil
		}
	}
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "tenancy.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import1.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListTenantsApiResponseData"))
}

func (p *OneOfListTenantsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if "List<tenancy.v4.config.Tenant>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfListTenantsApiResponseData")
}

type OneOfDeleteTenantApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfDeleteTenantApiResponseData() *OneOfDeleteTenantApiResponseData {
	p := new(OneOfDeleteTenantApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteTenantApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteTenantApiResponseData is nil"))
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
	case import1.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import1.ErrorResponse)
		}
		*p.oneOfType400 = v.(import1.ErrorResponse)
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

func (p *OneOfDeleteTenantApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDeleteTenantApiResponseData) UnmarshalJSON(b []byte) error {
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
					vOneOfType0 := new(import3.TaskReference)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType0)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType0.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
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
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import1.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "tenancy.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
							if nil == p.oneOfType400 {
								p.oneOfType400 = new(import1.ErrorResponse)
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
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if vOneOfType0.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
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
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "tenancy.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import1.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteTenantApiResponseData"))
}

func (p *OneOfDeleteTenantApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteTenantApiResponseData")
}

type OneOfGetDomainMembershipApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	oneOfType0    *Membership            `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfGetDomainMembershipApiResponseData() *OneOfGetDomainMembershipApiResponseData {
	p := new(OneOfGetDomainMembershipApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetDomainMembershipApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetDomainMembershipApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import1.ErrorResponse)
		}
		*p.oneOfType400 = v.(import1.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case Membership:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(Membership)
		}
		*p.oneOfType0 = v.(Membership)
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

func (p *OneOfGetDomainMembershipApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfGetDomainMembershipApiResponseData) UnmarshalJSON(b []byte) error {
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
					vOneOfType400 := new(import1.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "tenancy.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
							if nil == p.oneOfType400 {
								p.oneOfType400 = new(import1.ErrorResponse)
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
					vOneOfType0 := new(Membership)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType0)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType0.ObjectType_ != nil && "tenancy.v4.config.Membership" == *vOneOfType0.ObjectType_ {
							if nil == p.oneOfType0 {
								p.oneOfType0 = new(Membership)
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
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "tenancy.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import1.ErrorResponse)
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
	vOneOfType0 := new(Membership)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if vOneOfType0.ObjectType_ != nil && "tenancy.v4.config.Membership" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(Membership)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetDomainMembershipApiResponseData"))
}

func (p *OneOfGetDomainMembershipApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfGetDomainMembershipApiResponseData")
}

type OneOfListDomainMembershipsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    []Membership           `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfListDomainMembershipsApiResponseData() *OneOfListDomainMembershipsApiResponseData {
	p := new(OneOfListDomainMembershipsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListDomainMembershipsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListDomainMembershipsApiResponseData is nil"))
	}
	switch v.(type) {
	case []Membership:
		p.oneOfType0 = v.([]Membership)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<tenancy.v4.config.Membership>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<tenancy.v4.config.Membership>"
	case import1.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import1.ErrorResponse)
		}
		*p.oneOfType400 = v.(import1.ErrorResponse)
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

func (p *OneOfListDomainMembershipsApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if "List<tenancy.v4.config.Membership>" == *p.Discriminator {
		return p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfListDomainMembershipsApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["List<tenancy.v4.config.Membership>"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType0 := new([]Membership)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType0)
					if unmarshalErr == nil {
						// For arrays, verify the array item ObjectType matches
						if vOneOfType0 == nil || len(*vOneOfType0) == 0 || ((*vOneOfType0)[0].ObjectType_ != nil && "tenancy.v4.config.Membership" == *((*vOneOfType0)[0].ObjectType_)) {
							p.oneOfType0 = *vOneOfType0
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = "List<tenancy.v4.config.Membership>"
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = "List<tenancy.v4.config.Membership>"
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
					vOneOfType400 := new(import1.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "tenancy.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
							if nil == p.oneOfType400 {
								p.oneOfType400 = new(import1.ErrorResponse)
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
	vOneOfType0 := new([]Membership)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || (vOneOfType0 != nil && (*vOneOfType0)[0].ObjectType_ != nil && "tenancy.v4.config.Membership" == *((*vOneOfType0)[0].ObjectType_)) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<tenancy.v4.config.Membership>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<tenancy.v4.config.Membership>"
			return nil
		}
	}
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "tenancy.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import1.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListDomainMembershipsApiResponseData"))
}

func (p *OneOfListDomainMembershipsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if "List<tenancy.v4.config.Membership>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfListDomainMembershipsApiResponseData")
}

type OneOfListTenantsDomainAllocationsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    []DomainAllocation     `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfListTenantsDomainAllocationsApiResponseData() *OneOfListTenantsDomainAllocationsApiResponseData {
	p := new(OneOfListTenantsDomainAllocationsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListTenantsDomainAllocationsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListTenantsDomainAllocationsApiResponseData is nil"))
	}
	switch v.(type) {
	case []DomainAllocation:
		p.oneOfType0 = v.([]DomainAllocation)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<tenancy.v4.config.DomainAllocation>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<tenancy.v4.config.DomainAllocation>"
	case import1.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import1.ErrorResponse)
		}
		*p.oneOfType400 = v.(import1.ErrorResponse)
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

func (p *OneOfListTenantsDomainAllocationsApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if "List<tenancy.v4.config.DomainAllocation>" == *p.Discriminator {
		return p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfListTenantsDomainAllocationsApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["List<tenancy.v4.config.DomainAllocation>"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType0 := new([]DomainAllocation)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType0)
					if unmarshalErr == nil {
						// For arrays, verify the array item ObjectType matches
						if vOneOfType0 == nil || len(*vOneOfType0) == 0 || ((*vOneOfType0)[0].ObjectType_ != nil && "tenancy.v4.config.DomainAllocation" == *((*vOneOfType0)[0].ObjectType_)) {
							p.oneOfType0 = *vOneOfType0
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = "List<tenancy.v4.config.DomainAllocation>"
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = "List<tenancy.v4.config.DomainAllocation>"
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
					vOneOfType400 := new(import1.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "tenancy.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
							if nil == p.oneOfType400 {
								p.oneOfType400 = new(import1.ErrorResponse)
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
	vOneOfType0 := new([]DomainAllocation)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || (vOneOfType0 != nil && (*vOneOfType0)[0].ObjectType_ != nil && "tenancy.v4.config.DomainAllocation" == *((*vOneOfType0)[0].ObjectType_)) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<tenancy.v4.config.DomainAllocation>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<tenancy.v4.config.DomainAllocation>"
			return nil
		}
	}
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "tenancy.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import1.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListTenantsDomainAllocationsApiResponseData"))
}

func (p *OneOfListTenantsDomainAllocationsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if "List<tenancy.v4.config.DomainAllocation>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfListTenantsDomainAllocationsApiResponseData")
}

type OneOfDeleteTenantDomainAllocationApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfDeleteTenantDomainAllocationApiResponseData() *OneOfDeleteTenantDomainAllocationApiResponseData {
	p := new(OneOfDeleteTenantDomainAllocationApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteTenantDomainAllocationApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteTenantDomainAllocationApiResponseData is nil"))
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
	case import1.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import1.ErrorResponse)
		}
		*p.oneOfType400 = v.(import1.ErrorResponse)
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

func (p *OneOfDeleteTenantDomainAllocationApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDeleteTenantDomainAllocationApiResponseData) UnmarshalJSON(b []byte) error {
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
					vOneOfType0 := new(import3.TaskReference)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType0)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType0.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
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
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import1.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "tenancy.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
							if nil == p.oneOfType400 {
								p.oneOfType400 = new(import1.ErrorResponse)
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
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if vOneOfType0.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
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
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "tenancy.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import1.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteTenantDomainAllocationApiResponseData"))
}

func (p *OneOfDeleteTenantDomainAllocationApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteTenantDomainAllocationApiResponseData")
}

type OneOfCreateTenantDomainAllocationApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfCreateTenantDomainAllocationApiResponseData() *OneOfCreateTenantDomainAllocationApiResponseData {
	p := new(OneOfCreateTenantDomainAllocationApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateTenantDomainAllocationApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateTenantDomainAllocationApiResponseData is nil"))
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
	case import1.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import1.ErrorResponse)
		}
		*p.oneOfType400 = v.(import1.ErrorResponse)
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

func (p *OneOfCreateTenantDomainAllocationApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfCreateTenantDomainAllocationApiResponseData) UnmarshalJSON(b []byte) error {
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
					vOneOfType0 := new(import3.TaskReference)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType0)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType0.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
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
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import1.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "tenancy.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
							if nil == p.oneOfType400 {
								p.oneOfType400 = new(import1.ErrorResponse)
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
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if vOneOfType0.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
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
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "tenancy.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import1.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateTenantDomainAllocationApiResponseData"))
}

func (p *OneOfCreateTenantDomainAllocationApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfCreateTenantDomainAllocationApiResponseData")
}

type OneOfAddDomainApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	oneOfType0    *Domain                `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfAddDomainApiResponseData() *OneOfAddDomainApiResponseData {
	p := new(OneOfAddDomainApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfAddDomainApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfAddDomainApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import1.ErrorResponse)
		}
		*p.oneOfType400 = v.(import1.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case Domain:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(Domain)
		}
		*p.oneOfType0 = v.(Domain)
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

func (p *OneOfAddDomainApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfAddDomainApiResponseData) UnmarshalJSON(b []byte) error {
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
					vOneOfType400 := new(import1.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "tenancy.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
							if nil == p.oneOfType400 {
								p.oneOfType400 = new(import1.ErrorResponse)
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
					vOneOfType0 := new(Domain)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType0)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType0.ObjectType_ != nil && "tenancy.v4.config.Domain" == *vOneOfType0.ObjectType_ {
							if nil == p.oneOfType0 {
								p.oneOfType0 = new(Domain)
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
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "tenancy.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import1.ErrorResponse)
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
	vOneOfType0 := new(Domain)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if vOneOfType0.ObjectType_ != nil && "tenancy.v4.config.Domain" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(Domain)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfAddDomainApiResponseData"))
}

func (p *OneOfAddDomainApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfAddDomainApiResponseData")
}

type OneOfGetTenantDomainAllocationApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *DomainAllocation      `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfGetTenantDomainAllocationApiResponseData() *OneOfGetTenantDomainAllocationApiResponseData {
	p := new(OneOfGetTenantDomainAllocationApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetTenantDomainAllocationApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetTenantDomainAllocationApiResponseData is nil"))
	}
	switch v.(type) {
	case DomainAllocation:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(DomainAllocation)
		}
		*p.oneOfType0 = v.(DomainAllocation)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import1.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import1.ErrorResponse)
		}
		*p.oneOfType400 = v.(import1.ErrorResponse)
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

func (p *OneOfGetTenantDomainAllocationApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetTenantDomainAllocationApiResponseData) UnmarshalJSON(b []byte) error {
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
					vOneOfType0 := new(DomainAllocation)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType0)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType0.ObjectType_ != nil && "tenancy.v4.config.DomainAllocation" == *vOneOfType0.ObjectType_ {
							if nil == p.oneOfType0 {
								p.oneOfType0 = new(DomainAllocation)
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
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import1.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "tenancy.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
							if nil == p.oneOfType400 {
								p.oneOfType400 = new(import1.ErrorResponse)
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
	vOneOfType0 := new(DomainAllocation)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if vOneOfType0.ObjectType_ != nil && "tenancy.v4.config.DomainAllocation" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(DomainAllocation)
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
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "tenancy.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import1.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetTenantDomainAllocationApiResponseData"))
}

func (p *OneOfGetTenantDomainAllocationApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetTenantDomainAllocationApiResponseData")
}

type OneOfGetTenantExternalNetworkConnectionApiResponseData struct {
	Discriminator *string                              `json:"-"`
	ObjectType_   *string                              `json:"-"`
	oneOfType401  *ExternalNetworkConnectionProjection `json:"-"`
	oneOfType400  *import1.ErrorResponse               `json:"-"`
	oneOfType0    *ExternalNetworkConnection           `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfGetTenantExternalNetworkConnectionApiResponseData() *OneOfGetTenantExternalNetworkConnectionApiResponseData {
	p := new(OneOfGetTenantExternalNetworkConnectionApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetTenantExternalNetworkConnectionApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetTenantExternalNetworkConnectionApiResponseData is nil"))
	}
	switch v.(type) {
	case ExternalNetworkConnectionProjection:
		if nil == p.oneOfType401 {
			p.oneOfType401 = new(ExternalNetworkConnectionProjection)
		}
		*p.oneOfType401 = v.(ExternalNetworkConnectionProjection)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType401.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType401.ObjectType_
	case import1.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import1.ErrorResponse)
		}
		*p.oneOfType400 = v.(import1.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case ExternalNetworkConnection:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(ExternalNetworkConnection)
		}
		*p.oneOfType0 = v.(ExternalNetworkConnection)
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

func (p *OneOfGetTenantExternalNetworkConnectionApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType401 != nil && *p.oneOfType401.ObjectType_ == *p.Discriminator {
		return *p.oneOfType401
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfGetTenantExternalNetworkConnectionApiResponseData) UnmarshalJSON(b []byte) error {
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
					vOneOfType401 := new(ExternalNetworkConnectionProjection)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType401)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType401.ObjectType_ != nil && "tenancy.v4.config.ExternalNetworkConnectionProjection" == *vOneOfType401.ObjectType_ {
							if nil == p.oneOfType401 {
								p.oneOfType401 = new(ExternalNetworkConnectionProjection)
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
					vOneOfType400 := new(import1.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "tenancy.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
							if nil == p.oneOfType400 {
								p.oneOfType400 = new(import1.ErrorResponse)
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
					vOneOfType0 := new(ExternalNetworkConnection)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType0)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType0.ObjectType_ != nil && "tenancy.v4.config.ExternalNetworkConnection" == *vOneOfType0.ObjectType_ {
							if nil == p.oneOfType0 {
								p.oneOfType0 = new(ExternalNetworkConnection)
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
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType401 := new(ExternalNetworkConnectionProjection)
	if err := json.Unmarshal(b, vOneOfType401); err == nil {
		if vOneOfType401.ObjectType_ != nil && "tenancy.v4.config.ExternalNetworkConnectionProjection" == *vOneOfType401.ObjectType_ {
			if nil == p.oneOfType401 {
				p.oneOfType401 = new(ExternalNetworkConnectionProjection)
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
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "tenancy.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import1.ErrorResponse)
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
	vOneOfType0 := new(ExternalNetworkConnection)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if vOneOfType0.ObjectType_ != nil && "tenancy.v4.config.ExternalNetworkConnection" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(ExternalNetworkConnection)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetTenantExternalNetworkConnectionApiResponseData"))
}

func (p *OneOfGetTenantExternalNetworkConnectionApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType401 != nil && *p.oneOfType401.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType401)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfGetTenantExternalNetworkConnectionApiResponseData")
}

type OneOfGetTenantApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *Tenant                `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfGetTenantApiResponseData() *OneOfGetTenantApiResponseData {
	p := new(OneOfGetTenantApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetTenantApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetTenantApiResponseData is nil"))
	}
	switch v.(type) {
	case Tenant:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(Tenant)
		}
		*p.oneOfType0 = v.(Tenant)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import1.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import1.ErrorResponse)
		}
		*p.oneOfType400 = v.(import1.ErrorResponse)
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

func (p *OneOfGetTenantApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetTenantApiResponseData) UnmarshalJSON(b []byte) error {
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
					vOneOfType0 := new(Tenant)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType0)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType0.ObjectType_ != nil && "tenancy.v4.config.Tenant" == *vOneOfType0.ObjectType_ {
							if nil == p.oneOfType0 {
								p.oneOfType0 = new(Tenant)
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
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import1.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "tenancy.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
							if nil == p.oneOfType400 {
								p.oneOfType400 = new(import1.ErrorResponse)
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
	vOneOfType0 := new(Tenant)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if vOneOfType0.ObjectType_ != nil && "tenancy.v4.config.Tenant" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(Tenant)
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
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "tenancy.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import1.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetTenantApiResponseData"))
}

func (p *OneOfGetTenantApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetTenantApiResponseData")
}

type OneOfCreateDomainMembershipApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	oneOfType0    *Membership            `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfCreateDomainMembershipApiResponseData() *OneOfCreateDomainMembershipApiResponseData {
	p := new(OneOfCreateDomainMembershipApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateDomainMembershipApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateDomainMembershipApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import1.ErrorResponse)
		}
		*p.oneOfType400 = v.(import1.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case Membership:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(Membership)
		}
		*p.oneOfType0 = v.(Membership)
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

func (p *OneOfCreateDomainMembershipApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfCreateDomainMembershipApiResponseData) UnmarshalJSON(b []byte) error {
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
					vOneOfType400 := new(import1.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "tenancy.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
							if nil == p.oneOfType400 {
								p.oneOfType400 = new(import1.ErrorResponse)
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
					vOneOfType0 := new(Membership)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType0)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType0.ObjectType_ != nil && "tenancy.v4.config.Membership" == *vOneOfType0.ObjectType_ {
							if nil == p.oneOfType0 {
								p.oneOfType0 = new(Membership)
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
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "tenancy.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import1.ErrorResponse)
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
	vOneOfType0 := new(Membership)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if vOneOfType0.ObjectType_ != nil && "tenancy.v4.config.Membership" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(Membership)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateDomainMembershipApiResponseData"))
}

func (p *OneOfCreateDomainMembershipApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfCreateDomainMembershipApiResponseData")
}

type OneOfCreateTenantApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfCreateTenantApiResponseData() *OneOfCreateTenantApiResponseData {
	p := new(OneOfCreateTenantApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateTenantApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateTenantApiResponseData is nil"))
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
	case import1.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import1.ErrorResponse)
		}
		*p.oneOfType400 = v.(import1.ErrorResponse)
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

func (p *OneOfCreateTenantApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfCreateTenantApiResponseData) UnmarshalJSON(b []byte) error {
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
					vOneOfType0 := new(import3.TaskReference)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType0)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType0.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
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
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import1.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "tenancy.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
							if nil == p.oneOfType400 {
								p.oneOfType400 = new(import1.ErrorResponse)
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
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if vOneOfType0.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
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
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "tenancy.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import1.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateTenantApiResponseData"))
}

func (p *OneOfCreateTenantApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfCreateTenantApiResponseData")
}

type OneOfUpdateTenantDomainAllocationApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfUpdateTenantDomainAllocationApiResponseData() *OneOfUpdateTenantDomainAllocationApiResponseData {
	p := new(OneOfUpdateTenantDomainAllocationApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateTenantDomainAllocationApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateTenantDomainAllocationApiResponseData is nil"))
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
	case import1.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import1.ErrorResponse)
		}
		*p.oneOfType400 = v.(import1.ErrorResponse)
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

func (p *OneOfUpdateTenantDomainAllocationApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfUpdateTenantDomainAllocationApiResponseData) UnmarshalJSON(b []byte) error {
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
					vOneOfType0 := new(import3.TaskReference)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType0)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType0.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
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
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import1.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "tenancy.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
							if nil == p.oneOfType400 {
								p.oneOfType400 = new(import1.ErrorResponse)
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
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if vOneOfType0.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
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
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "tenancy.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import1.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateTenantDomainAllocationApiResponseData"))
}

func (p *OneOfUpdateTenantDomainAllocationApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateTenantDomainAllocationApiResponseData")
}

type OneOfListDomainsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    []Domain               `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfListDomainsApiResponseData() *OneOfListDomainsApiResponseData {
	p := new(OneOfListDomainsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListDomainsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListDomainsApiResponseData is nil"))
	}
	switch v.(type) {
	case []Domain:
		p.oneOfType0 = v.([]Domain)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<tenancy.v4.config.Domain>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<tenancy.v4.config.Domain>"
	case import1.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import1.ErrorResponse)
		}
		*p.oneOfType400 = v.(import1.ErrorResponse)
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

func (p *OneOfListDomainsApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if "List<tenancy.v4.config.Domain>" == *p.Discriminator {
		return p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfListDomainsApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["List<tenancy.v4.config.Domain>"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType0 := new([]Domain)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType0)
					if unmarshalErr == nil {
						// For arrays, verify the array item ObjectType matches
						if vOneOfType0 == nil || len(*vOneOfType0) == 0 || ((*vOneOfType0)[0].ObjectType_ != nil && "tenancy.v4.config.Domain" == *((*vOneOfType0)[0].ObjectType_)) {
							p.oneOfType0 = *vOneOfType0
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = "List<tenancy.v4.config.Domain>"
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = "List<tenancy.v4.config.Domain>"
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
					vOneOfType400 := new(import1.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "tenancy.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
							if nil == p.oneOfType400 {
								p.oneOfType400 = new(import1.ErrorResponse)
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
	vOneOfType0 := new([]Domain)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || (vOneOfType0 != nil && (*vOneOfType0)[0].ObjectType_ != nil && "tenancy.v4.config.Domain" == *((*vOneOfType0)[0].ObjectType_)) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<tenancy.v4.config.Domain>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<tenancy.v4.config.Domain>"
			return nil
		}
	}
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "tenancy.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import1.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListDomainsApiResponseData"))
}

func (p *OneOfListDomainsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if "List<tenancy.v4.config.Domain>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfListDomainsApiResponseData")
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
