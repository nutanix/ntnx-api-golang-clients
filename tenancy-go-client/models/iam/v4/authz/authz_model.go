/*
 * Generated file models/iam/v4/authz/authz_model.go.
 *
 * Product version: 4.0.1-alpha-1
 *
 * Part of the SP Central Tenant Management
 *
 * (c) 2026 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module iam.v4.authz of SP Central Tenant Management
*/
package authz

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

/*
Type of identity associated with the role membership.
*/
type RmIdentityType int

const (
	RMIDENTITYTYPE_UNKNOWN  RmIdentityType = 0
	RMIDENTITYTYPE_REDACTED RmIdentityType = 1
	RMIDENTITYTYPE_USER     RmIdentityType = 2
	RMIDENTITYTYPE_GROUP    RmIdentityType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *RmIdentityType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"USER",
		"GROUP",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e RmIdentityType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"USER",
		"GROUP",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *RmIdentityType) index(name string) RmIdentityType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"USER",
		"GROUP",
	}
	for idx := range names {
		if names[idx] == name {
			return RmIdentityType(idx)
		}
	}
	return RMIDENTITYTYPE_UNKNOWN
}

func (e *RmIdentityType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for RmIdentityType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *RmIdentityType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e RmIdentityType) Ref() *RmIdentityType {
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
