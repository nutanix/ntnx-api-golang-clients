/*
 * Generated file models/common/v1/response/response_model.go.
 *
 * Product version: 4.1.1
 *
 * Part of the Nutanix Volumes APIs
 *
 * (c) 2025 Nutanix Inc.  All rights reserved
 *
 */

/*
  Nutanix Standard Response Format
*/
package response

import (
	"encoding/json"
	import1 "github.com/nutanix/ntnx-api-golang-clients/volumes-go-client/v4/models/common/v1/config"
)

/*
A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
*/
type ApiLink struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The URL at which the entity described by the link can be accessed.
	*/
	Href *string `json:"href,omitempty"`
	/*
	  A name that identifies the relationship of the link to the object that is returned by the URL.  The unique value of "self" identifies the URL for the object.
	*/
	Rel *string `json:"rel,omitempty"`
}

func (p *ApiLink) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ApiLink

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

func (p *ApiLink) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ApiLink
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = ApiLink(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "href")
	delete(allFields, "rel")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewApiLink() *ApiLink {
	p := new(ApiLink)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "common.v1.response.ApiLink"
	p.Reserved_ = map[string]interface{}{"$fv": "v1.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The metadata associated with an API response. This value is always present and minimally contains the self-link for the API request that produced this response. It also contains pagination data for the paginated requests.
*/
type ApiResponseMetadata struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  An array of entity-specific metadata
	*/
	ExtraInfo []import1.KVPair `json:"extraInfo,omitempty"`
	/*
	  An array of flags that may indicate the status of the response. For example, a flag with the name 'isPaginated' and value 'false', indicates that the response is not paginated.
	*/
	Flags []import1.Flag `json:"flags,omitempty"`
	/*
	  An array of HATEOAS style links for the response that may also include pagination links for list operations.
	*/
	Links []ApiLink `json:"links,omitempty"`
	/*
	  Information, Warning or Error messages that might provide additional contextual information related to the operation.
	*/
	Messages []import1.Message `json:"messages,omitempty"`
	/*
	  The total number of entities that are available on the server for this type.
	*/
	TotalAvailableResults *int `json:"totalAvailableResults,omitempty"`
}

func (p *ApiResponseMetadata) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ApiResponseMetadata

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

func (p *ApiResponseMetadata) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ApiResponseMetadata
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = ApiResponseMetadata(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "extraInfo")
	delete(allFields, "flags")
	delete(allFields, "links")
	delete(allFields, "messages")
	delete(allFields, "totalAvailableResults")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewApiResponseMetadata() *ApiResponseMetadata {
	p := new(ApiResponseMetadata)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "common.v1.response.ApiResponseMetadata"
	p.Reserved_ = map[string]interface{}{"$fv": "v1.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
A model that represents an object instance that is accessible through an API endpoint.  Instances of this type get an extId field that contains the globally unique identifier for that instance.  Externally accessible instances are always tenant aware and, therefore, extend the TenantAwareModel
*/
type ExternalizableAbstractModel struct {
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
	Links []ApiLink `json:"links,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *ExternalizableAbstractModel) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ExternalizableAbstractModel

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

func (p *ExternalizableAbstractModel) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ExternalizableAbstractModel
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = ExternalizableAbstractModel(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "extId")
	delete(allFields, "links")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewExternalizableAbstractModel() *ExternalizableAbstractModel {
	p := new(ExternalizableAbstractModel)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "common.v1.response.ExternalizableAbstractModel"
	p.Reserved_ = map[string]interface{}{"$fv": "v1.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
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
