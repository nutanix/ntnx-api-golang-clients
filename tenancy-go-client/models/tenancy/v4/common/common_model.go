/*
 * Generated file models/tenancy/v4/common/common_model.go.
 *
 * Product version: 4.0.1-alpha-1
 *
 * Part of the SP Central Tenant Management
 *
 * (c) 2026 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module tenancy.v4.common of SP Central Tenant Management
*/
package common

import (
	"encoding/json"
)

/*
Base model for action API response.
*/
type ActionBaseResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Success message returned by the action API.
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
	*p.ObjectType_ = "tenancy.v4.common.ActionBaseResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Base state info containing common error details.
*/
type BaseStateInfo struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Error *ErrorResponse `json:"error,omitempty"`
}

func (p *BaseStateInfo) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias BaseStateInfo

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

func (p *BaseStateInfo) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias BaseStateInfo
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewBaseStateInfo()

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

func NewBaseStateInfo() *BaseStateInfo {
	p := new(BaseStateInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "tenancy.v4.common.BaseStateInfo"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Error response containing an error code and message.
*/
type ErrorResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The error code identifying the type of error.
	*/
	Code *string `json:"code"`
	/*
	  The human-readable error message.
	*/
	Message *string `json:"message"`
}

func (p *ErrorResponse) MarshalJSON() ([]byte, error) {
	type ErrorResponseProxy ErrorResponse

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*ErrorResponseProxy
		Code    *string `json:"code,omitempty"`
		Message *string `json:"message,omitempty"`
	}{
		ErrorResponseProxy: (*ErrorResponseProxy)(p),
		Code:               p.Code,
		Message:            p.Message,
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
	*p = *NewErrorResponse()

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
	if known.Message != nil {
		p.Message = known.Message
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "code")
	delete(allFields, "message")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewErrorResponse() *ErrorResponse {
	p := new(ErrorResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "tenancy.v4.common.ErrorResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
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
