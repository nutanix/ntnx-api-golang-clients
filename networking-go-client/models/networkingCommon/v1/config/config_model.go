/*
 * Generated file models/networkingCommon/v1/config/config_model.go.
 *
 * Product version: 4.4.1
 *
 * Part of the Nutanix Networking APIs
 *
 * (c) 2026 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module networkingCommon.v1.config of Nutanix Networking APIs
*/
package config

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

/*
Enum model with all possible modes (ETHERNET/INFINIBAND) in which a NIC can operate.
*/
type OperatingMode int

const (
	OPERATINGMODE_UNKNOWN    OperatingMode = 0
	OPERATINGMODE_REDACTED   OperatingMode = 1
	OPERATINGMODE_ETHERNET   OperatingMode = 2
	OPERATINGMODE_INFINIBAND OperatingMode = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *OperatingMode) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ETHERNET",
		"INFINIBAND",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e OperatingMode) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ETHERNET",
		"INFINIBAND",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *OperatingMode) index(name string) OperatingMode {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ETHERNET",
		"INFINIBAND",
	}
	for idx := range names {
		if names[idx] == name {
			return OperatingMode(idx)
		}
	}
	return OPERATINGMODE_UNKNOWN
}

func (e *OperatingMode) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for OperatingMode:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *OperatingMode) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e OperatingMode) Ref() *OperatingMode {
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
