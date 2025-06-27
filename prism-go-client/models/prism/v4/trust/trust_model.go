/*
 * Generated file models/prism/v4/trust/trust_model.go.
 *
 * Product version: 4.1.1
 *
 * Part of the Nutanix Prism APIs
 *
 * (c) 2025 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module prism.v4.trust of Nutanix Prism APIs
*/
package trust

import (
	"encoding/json"
)

/*
HATEOAS links for the request. For paginated requests includes
prev/next/first and last links
*/
type ApiLink struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The URL that points to the relationship
	*/
	Href *string `json:"href,omitempty"`
	/*
	  The name of the relationship
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
	*p.ObjectType_ = "prism.v4.trust.ApiLink"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Metadata associated with API responses
*/
type ApiResponseMetadata struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Links []ApiLink `json:"links,omitempty"`
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
	delete(allFields, "links")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewApiResponseMetadata() *ApiResponseMetadata {
	p := new(ApiResponseMetadata)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.trust.ApiResponseMetadata"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
An Object capturing metadata related to authentication and authorization.
*/
type AuthMetadata struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  An Auth token generated to be compatible with Component Registry.
	In request it will be auth token of client and likewise in reponse it will
	be of PC component.
	*/
	AuthToken *string `json:"authToken,omitempty"`
}

func (p *AuthMetadata) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias AuthMetadata

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

func (p *AuthMetadata) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias AuthMetadata
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = AuthMetadata(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "authToken")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewAuthMetadata() *AuthMetadata {
	p := new(AuthMetadata)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.trust.AuthMetadata"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type SignedCertDetails struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	AuthMetadata *AuthMetadata `json:"authMetadata,omitempty"`
	/*
	  PC Component's cert chain in PEM format.
	*/
	CaCertChain *string `json:"caCertChain,omitempty"`
	/*
	  Public key of the client in Certificate Signing Request format encoded in
	PEM format.
	*/
	CertificateSigningRequest *string `json:"certificateSigningRequest,omitempty"`
	/*
	  An error string capturing any errors faced during trust setup,
	will be empty if operation is successful.
	*/
	Error *string `json:"error,omitempty"`
	/*
	  Clients Public key signed by PC's intermedicate Certificate in PEM format.
	*/
	SignedCertificate *string `json:"signedCertificate,omitempty"`
}

func (p *SignedCertDetails) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias SignedCertDetails

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

func (p *SignedCertDetails) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias SignedCertDetails
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = SignedCertDetails(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "authMetadata")
	delete(allFields, "caCertChain")
	delete(allFields, "certificateSigningRequest")
	delete(allFields, "error")
	delete(allFields, "signedCertificate")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewSignedCertDetails() *SignedCertDetails {
	p := new(SignedCertDetails)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.trust.SignedCertDetails"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
An Object representing a Trust Setup request.
*/
type Trust struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	AuthMetadata *AuthMetadata `json:"authMetadata,omitempty"`
	/*
	  Public key of the client in Certificate Signing Request format encoded in
	PEM format.
	*/
	CertificateSigningRequest *string `json:"certificateSigningRequest,omitempty"`
}

func (p *Trust) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias Trust

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

func (p *Trust) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Trust
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = Trust(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "authMetadata")
	delete(allFields, "certificateSigningRequest")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewTrust() *Trust {
	p := new(Trust)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.trust.Trust"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
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
