/*
 * Generated file models/networking/v4/error/error_model.go.
 *
 * Product version: 4.1.1
 *
 * Part of the Nutanix Networking APIs
 *
 * (c) 2025 Nutanix Inc.  All rights reserved
 *
 */

/*
  Auto-injected error response object by the dev platform
*/
package error

import (
	"encoding/json"
	"errors"
	"fmt"
	import1 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/common/v1/config"
	"time"
)

/*
Message with associated severity describing status of the current operation.
*/
type AppMessage struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The map of argument name to value.
	*/
	ArgumentsMap map[string]string `json:"argumentsMap,omitempty"`
	/*
	  The code associated with this message. This string is typically prefixed with the namespace to which the endpoint belongs. For example: VMM-40000
	*/
	Code *string `json:"code,omitempty"`
	/*
	  The error group associated with this message of severity ERROR.
	*/
	ErrorGroup *string `json:"errorGroup,omitempty"`
	/*
	  Locale for this message. The default locale would be 'en-US'.
	*/
	Locale *string `json:"locale,omitempty"`
	/*
	  The message string.
	*/
	Message *string `json:"message,omitempty"`

	Severity *import1.MessageSeverity `json:"severity,omitempty"`
}

func (p *AppMessage) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias AppMessage

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

func (p *AppMessage) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias AppMessage
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = AppMessage(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "argumentsMap")
	delete(allFields, "code")
	delete(allFields, "errorGroup")
	delete(allFields, "locale")
	delete(allFields, "message")
	delete(allFields, "severity")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewAppMessage() *AppMessage {
	p := new(AppMessage)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.error.AppMessage"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	p.Locale = new(string)
	*p.Locale = "en_US"

	return p
}

/*
An error response indicates that the operation has failed either due to a client error(4XX) or server error(5XX). Please look at the HTTP status code and namespace specific error code and error message for further details.
*/
type ErrorResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	ErrorItemDiscriminator_ *string `json:"$errorItemDiscriminator,omitempty"`

	Error *OneOfErrorResponseError `json:"error,omitempty"`
}

func (p *ErrorResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ErrorResponse

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

func (p *ErrorResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ErrorResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = ErrorResponse(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$errorItemDiscriminator")
	delete(allFields, "error")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewErrorResponse() *ErrorResponse {
	p := new(ErrorResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.error.ErrorResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ErrorResponse) GetError() interface{} {
	if nil == p.Error {
		return nil
	}
	return p.Error.GetValue()
}

func (p *ErrorResponse) SetError(v interface{}) error {
	if nil == p.Error {
		p.Error = NewOneOfErrorResponseError()
	}
	e := p.Error.SetValue(v)
	if nil == e {
		if nil == p.ErrorItemDiscriminator_ {
			p.ErrorItemDiscriminator_ = new(string)
		}
		*p.ErrorItemDiscriminator_ = *p.Error.Discriminator
	}
	return e
}

/*
This schema is generated from SchemaValidationError.java
*/
type SchemaValidationError struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The generic error message for the response.
	*/
	Error *string `json:"error,omitempty"`
	/*
	  API path on which the request was made.
	*/
	Path *string `json:"path,omitempty"`
	/*
	  The HTTP status code of the response.
	*/
	StatusCode *int `json:"statusCode,omitempty"`
	/*
	  Timestamp of the response.
	*/
	Timestamp *time.Time `json:"timestamp,omitempty"`
	/*
	  List of validation error messages
	*/
	ValidationErrorMessages []SchemaValidationErrorMessage `json:"validationErrorMessages,omitempty"`
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
	*p = SchemaValidationError(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "error")
	delete(allFields, "path")
	delete(allFields, "statusCode")
	delete(allFields, "timestamp")
	delete(allFields, "validationErrorMessages")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewSchemaValidationError() *SchemaValidationError {
	p := new(SchemaValidationError)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.error.SchemaValidationError"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
This schema is generated from SchemaValidationErrorMessage.java
*/
type SchemaValidationErrorMessage struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The path of the attribute that failed validation in the schema.
	*/
	AttributePath *string `json:"attributePath,omitempty"`
	/*
	  The part of the request that failed validation. Validation can fail for path, query parameters, and request body.
	*/
	Location *string `json:"location,omitempty"`
	/*
	  The detailed message for the validation error.
	*/
	Message *string `json:"message,omitempty"`
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
	*p = SchemaValidationErrorMessage(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "attributePath")
	delete(allFields, "location")
	delete(allFields, "message")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewSchemaValidationErrorMessage() *SchemaValidationErrorMessage {
	p := new(SchemaValidationErrorMessage)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.error.SchemaValidationErrorMessage"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type OneOfErrorResponseError struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType202  *SchemaValidationError `json:"-"`
	oneOfType201  []AppMessage           `json:"-"`
}

func NewOneOfErrorResponseError() *OneOfErrorResponseError {
	p := new(OneOfErrorResponseError)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfErrorResponseError) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfErrorResponseError is nil"))
	}
	switch v.(type) {
	case SchemaValidationError:
		if nil == p.oneOfType202 {
			p.oneOfType202 = new(SchemaValidationError)
		}
		*p.oneOfType202 = v.(SchemaValidationError)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType202.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType202.ObjectType_
	case []AppMessage:
		p.oneOfType201 = v.([]AppMessage)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<networking.v4.error.AppMessage>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<networking.v4.error.AppMessage>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfErrorResponseError) GetValue() interface{} {
	if p.oneOfType202 != nil && *p.oneOfType202.ObjectType_ == *p.Discriminator {
		return *p.oneOfType202
	}
	if "List<networking.v4.error.AppMessage>" == *p.Discriminator {
		return p.oneOfType201
	}
	return nil
}

func (p *OneOfErrorResponseError) UnmarshalJSON(b []byte) error {
	vOneOfType202 := new(SchemaValidationError)
	if err := json.Unmarshal(b, vOneOfType202); err == nil {
		if "networking.v4.error.SchemaValidationError" == *vOneOfType202.ObjectType_ {
			if nil == p.oneOfType202 {
				p.oneOfType202 = new(SchemaValidationError)
			}
			*p.oneOfType202 = *vOneOfType202
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType202.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType202.ObjectType_
			return nil
		}
	}
	vOneOfType201 := new([]AppMessage)
	if err := json.Unmarshal(b, vOneOfType201); err == nil {
		if len(*vOneOfType201) == 0 || "networking.v4.error.AppMessage" == *((*vOneOfType201)[0].ObjectType_) {
			p.oneOfType201 = *vOneOfType201
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<networking.v4.error.AppMessage>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<networking.v4.error.AppMessage>"
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfErrorResponseError"))
}

func (p *OneOfErrorResponseError) MarshalJSON() ([]byte, error) {
	if p.oneOfType202 != nil && *p.oneOfType202.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType202)
	}
	if "List<networking.v4.error.AppMessage>" == *p.Discriminator {
		return json.Marshal(p.oneOfType201)
	}
	return nil, errors.New("No value to marshal for OneOfErrorResponseError")
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
