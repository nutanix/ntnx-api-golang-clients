/*
 * Generated file models/iam/v4/common/common_model.go.
 *
 * Product version: 4.1.1-beta-2
 *
 * Part of the Nutanix Identity and Access Management APIs
 *
 * (c) 2025 Nutanix Inc.  All rights reserved
 *
 */

/*
  Helper classes
*/
package common

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	import1 "github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/models/common/v1/response"
	"time"
)

/*
Base model for action API response.
*/
type ActionBaseResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Action API successful message.
	*/
	Message *string `json:"message"`
}

func (p *ActionBaseResponse) MarshalJSON() ([]byte, error) {
	type ActionBaseResponseProxy ActionBaseResponse

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*ActionBaseResponseProxy
		Message *string `json:"message,omitempty"`
	}{
		ActionBaseResponseProxy: (*ActionBaseResponseProxy)(p),
		Message:                 p.Message,
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

func (p *ActionBaseResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ActionBaseResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewActionBaseResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Message != nil {
		p.Message = known.Message
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "message")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewActionBaseResponse() *ActionBaseResponse {
	p := new(ActionBaseResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.common.ActionBaseResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Schema Validation Error
*/
type SchemaValidationError struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Error *SchemaValidationErrorMessage `json:"error,omitempty"`
}

func (p *SchemaValidationError) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias SchemaValidationError

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

func (p *SchemaValidationError) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias SchemaValidationError
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewSchemaValidationError()

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

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "error")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewSchemaValidationError() *SchemaValidationError {
	p := new(SchemaValidationError)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.common.SchemaValidationError"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Schema Validation Error Object
*/
type SchemaValidationErrorMessage struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Error code for Schema Validation Error
	*/
	Code *string `json:"code,omitempty"`
	/*
	  The generic error message for the response
	*/
	Error *string `json:"error,omitempty"`
	/*
	  Error group for Schema Validation error code
	*/
	ErrorGroup *string `json:"errorGroup,omitempty"`
	/*
	  API path on which the request was made
	*/
	Path *string `json:"path,omitempty"`
	/*
	  The HTTP status code of the response
	*/
	StatusCode *int `json:"statusCode,omitempty"`
	/*
	  Timestamp of the response
	*/
	Timestamp *time.Time `json:"timestamp,omitempty"`

	ValidationErrorMessages []ValidationErrorMessage `json:"validationErrorMessages,omitempty"`
}

func (p *SchemaValidationErrorMessage) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias SchemaValidationErrorMessage

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

func (p *SchemaValidationErrorMessage) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias SchemaValidationErrorMessage
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewSchemaValidationErrorMessage()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Code != nil {
		p.Code = known.Code
	}
	if known.Error != nil {
		p.Error = known.Error
	}
	if known.ErrorGroup != nil {
		p.ErrorGroup = known.ErrorGroup
	}
	if known.Path != nil {
		p.Path = known.Path
	}
	if known.StatusCode != nil {
		p.StatusCode = known.StatusCode
	}
	if known.Timestamp != nil {
		p.Timestamp = known.Timestamp
	}
	if known.ValidationErrorMessages != nil {
		p.ValidationErrorMessages = known.ValidationErrorMessages
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "code")
	delete(allFields, "error")
	delete(allFields, "errorGroup")
	delete(allFields, "path")
	delete(allFields, "statusCode")
	delete(allFields, "timestamp")
	delete(allFields, "validationErrorMessages")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewSchemaValidationErrorMessage() *SchemaValidationErrorMessage {
	p := new(SchemaValidationErrorMessage)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.common.SchemaValidationErrorMessage"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Api Response to Schema Validation Error
*/
type SchemaValidationErrorResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Data *SchemaValidationError `json:"data,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *SchemaValidationErrorResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias SchemaValidationErrorResponse

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

func (p *SchemaValidationErrorResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias SchemaValidationErrorResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewSchemaValidationErrorResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Data != nil {
		p.Data = known.Data
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
	delete(allFields, "data")
	delete(allFields, "extId")
	delete(allFields, "links")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewSchemaValidationErrorResponse() *SchemaValidationErrorResponse {
	p := new(SchemaValidationErrorResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.common.SchemaValidationErrorResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
DecRef(sortOrderDesc)
*/
type SortOrderType int

const (
	SORTORDERTYPE_UNKNOWN  SortOrderType = 0
	SORTORDERTYPE_REDACTED SortOrderType = 1
	SORTORDERTYPE_ASC      SortOrderType = 2
	SORTORDERTYPE_DESC     SortOrderType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *SortOrderType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ASC",
		"DESC",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e SortOrderType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ASC",
		"DESC",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *SortOrderType) index(name string) SortOrderType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ASC",
		"DESC",
	}
	for idx := range names {
		if names[idx] == name {
			return SortOrderType(idx)
		}
	}
	return SORTORDERTYPE_UNKNOWN
}

func (e *SortOrderType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for SortOrderType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *SortOrderType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e SortOrderType) Ref() *SortOrderType {
	return &e
}

/*
Schema Validation Error Message
*/
type ValidationErrorMessage struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The detailed message for the validation error
	*/
	Message *string `json:"message,omitempty"`
}

func (p *ValidationErrorMessage) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ValidationErrorMessage

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

func (p *ValidationErrorMessage) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ValidationErrorMessage
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewValidationErrorMessage()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Message != nil {
		p.Message = known.Message
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "message")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewValidationErrorMessage() *ValidationErrorMessage {
	p := new(ValidationErrorMessage)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.common.ValidationErrorMessage"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
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
