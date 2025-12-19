/*
 * Generated file models/dataprotection/v4/common/common_model.go.
 *
 * Product version: 4.2.1
 *
 * Part of the Nutanix Data Policies APIs
 *
 * (c) 2025 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module dataprotection.v4.common of Nutanix Data Policies APIs
*/
package common

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

/*
Type of the Recovery point.
*/
type RecoveryPointType int

const (
	RECOVERYPOINTTYPE_UNKNOWN                RecoveryPointType = 0
	RECOVERYPOINTTYPE_REDACTED               RecoveryPointType = 1
	RECOVERYPOINTTYPE_CRASH_CONSISTENT       RecoveryPointType = 2
	RECOVERYPOINTTYPE_APPLICATION_CONSISTENT RecoveryPointType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *RecoveryPointType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CRASH_CONSISTENT",
		"APPLICATION_CONSISTENT",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e RecoveryPointType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CRASH_CONSISTENT",
		"APPLICATION_CONSISTENT",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *RecoveryPointType) index(name string) RecoveryPointType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CRASH_CONSISTENT",
		"APPLICATION_CONSISTENT",
	}
	for idx := range names {
		if names[idx] == name {
			return RecoveryPointType(idx)
		}
	}
	return RECOVERYPOINTTYPE_UNKNOWN
}

func (e *RecoveryPointType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for RecoveryPointType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *RecoveryPointType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e RecoveryPointType) Ref() *RecoveryPointType {
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
