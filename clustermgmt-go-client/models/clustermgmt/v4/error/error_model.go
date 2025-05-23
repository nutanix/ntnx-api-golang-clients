/*
 * Generated file models/clustermgmt/v4/error/error_model.go.
 *
 * Product version: 4.0.2
 *
 * Part of the Nutanix Cluster Management APIs
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
	import1 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/common/v1/config"
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
	  The code associated with this message.This string is typically prefixed by the namespace the endpoint belongs to. For example: VMM-40000
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

func NewAppMessage() *AppMessage {
	p := new(AppMessage)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.error.AppMessage"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
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

func NewErrorResponse() *ErrorResponse {
	p := new(ErrorResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.error.ErrorResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
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

func NewSchemaValidationError() *SchemaValidationError {
	p := new(SchemaValidationError)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.error.SchemaValidationError"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
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

func NewSchemaValidationErrorMessage() *SchemaValidationErrorMessage {
	p := new(SchemaValidationErrorMessage)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.error.SchemaValidationErrorMessage"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
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
		*p.Discriminator = "List<clustermgmt.v4.error.AppMessage>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<clustermgmt.v4.error.AppMessage>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfErrorResponseError) GetValue() interface{} {
	if p.oneOfType202 != nil && *p.oneOfType202.ObjectType_ == *p.Discriminator {
		return *p.oneOfType202
	}
	if "List<clustermgmt.v4.error.AppMessage>" == *p.Discriminator {
		return p.oneOfType201
	}
	return nil
}

func (p *OneOfErrorResponseError) UnmarshalJSON(b []byte) error {
	vOneOfType202 := new(SchemaValidationError)
	if err := json.Unmarshal(b, vOneOfType202); err == nil {
		if "clustermgmt.v4.error.SchemaValidationError" == *vOneOfType202.ObjectType_ {
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
		if len(*vOneOfType201) == 0 || "clustermgmt.v4.error.AppMessage" == *((*vOneOfType201)[0].ObjectType_) {
			p.oneOfType201 = *vOneOfType201
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<clustermgmt.v4.error.AppMessage>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<clustermgmt.v4.error.AppMessage>"
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfErrorResponseError"))
}

func (p *OneOfErrorResponseError) MarshalJSON() ([]byte, error) {
	if p.oneOfType202 != nil && *p.oneOfType202.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType202)
	}
	if "List<clustermgmt.v4.error.AppMessage>" == *p.Discriminator {
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
