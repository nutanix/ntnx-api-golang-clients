/*
 * Generated file models/iam/v4/common/common_model.go.
 *
 * Product version: 4.0.1
 *
 * Part of the Nutanix IAM Versioned APIs
 *
 * (c) 2024 Nutanix Inc.  All rights reserved
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
	return json.Marshal(struct {
		*ActionBaseResponseProxy
		Message *string `json:"message,omitempty"`
	}{
		ActionBaseResponseProxy: (*ActionBaseResponseProxy)(p),
		Message:                 p.Message,
	})
}

func NewActionBaseResponse() *ActionBaseResponse {
	p := new(ActionBaseResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.common.ActionBaseResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
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

func NewSchemaValidationError() *SchemaValidationError {
	p := new(SchemaValidationError)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.common.SchemaValidationError"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
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

func NewSchemaValidationErrorMessage() *SchemaValidationErrorMessage {
	p := new(SchemaValidationErrorMessage)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.common.SchemaValidationErrorMessage"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
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
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewSchemaValidationErrorResponse() *SchemaValidationErrorResponse {
	p := new(SchemaValidationErrorResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.common.SchemaValidationErrorResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
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

func NewValidationErrorMessage() *ValidationErrorMessage {
	p := new(ValidationErrorMessage)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.common.ValidationErrorMessage"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
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
