/*
 * Generated file models/volumes/v4/config/config_model.go.
 *
 * Product version: 4.2.1
 *
 * Part of the Nutanix Data Policies APIs
 *
 * (c) 2025 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module volumes.v4.config of Nutanix Data Policies APIs
*/
package config

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

/*
The field indicates whether a VG has a VM or an external attachment associated with it.
*/
type AttachmentType int

const (
	ATTACHMENTTYPE_UNKNOWN  AttachmentType = 0
	ATTACHMENTTYPE_REDACTED AttachmentType = 1
	ATTACHMENTTYPE_NONE     AttachmentType = 2
	ATTACHMENTTYPE_DIRECT   AttachmentType = 3
	ATTACHMENTTYPE_EXTERNAL AttachmentType = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *AttachmentType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NONE",
		"DIRECT",
		"EXTERNAL",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e AttachmentType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NONE",
		"DIRECT",
		"EXTERNAL",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *AttachmentType) index(name string) AttachmentType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NONE",
		"DIRECT",
		"EXTERNAL",
	}
	for idx := range names {
		if names[idx] == name {
			return AttachmentType(idx)
		}
	}
	return ATTACHMENTTYPE_UNKNOWN
}

func (e *AttachmentType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for AttachmentType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *AttachmentType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e AttachmentType) Ref() *AttachmentType {
	return &e
}

/*
The authentication type enabled for the Volume Group. This is an optional field. If omitted, the authentication is not configured for the Volume Group. If this is set to CHAP, the target/client secret must be provided.
*/
type AuthenticationType int

const (
	AUTHENTICATIONTYPE_UNKNOWN  AuthenticationType = 0
	AUTHENTICATIONTYPE_REDACTED AuthenticationType = 1
	AUTHENTICATIONTYPE_CHAP     AuthenticationType = 2
	AUTHENTICATIONTYPE_NONE     AuthenticationType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *AuthenticationType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CHAP",
		"NONE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e AuthenticationType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CHAP",
		"NONE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *AuthenticationType) index(name string) AuthenticationType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CHAP",
		"NONE",
	}
	for idx := range names {
		if names[idx] == name {
			return AuthenticationType(idx)
		}
	}
	return AUTHENTICATIONTYPE_UNKNOWN
}

func (e *AuthenticationType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for AuthenticationType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *AuthenticationType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e AuthenticationType) Ref() *AuthenticationType {
	return &e
}

/*
Type of protocol to be used for Volume Group.
*/
type Protocol int

const (
	PROTOCOL_UNKNOWN      Protocol = 0
	PROTOCOL_REDACTED     Protocol = 1
	PROTOCOL_NOT_ASSIGNED Protocol = 2
	PROTOCOL_ISCSI        Protocol = 3
	PROTOCOL_NVMF         Protocol = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *Protocol) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NOT_ASSIGNED",
		"ISCSI",
		"NVMF",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e Protocol) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NOT_ASSIGNED",
		"ISCSI",
		"NVMF",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *Protocol) index(name string) Protocol {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NOT_ASSIGNED",
		"ISCSI",
		"NVMF",
	}
	for idx := range names {
		if names[idx] == name {
			return Protocol(idx)
		}
	}
	return PROTOCOL_UNKNOWN
}

func (e *Protocol) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for Protocol:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *Protocol) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e Protocol) Ref() *Protocol {
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
