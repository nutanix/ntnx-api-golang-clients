/*
 * Generated file models/iam/v4/common/common_model.go.
 *
 * Product version: 4.0.1-beta-1
 *
 * Part of the Nutanix Iam Versioned APIs
 *
 * (c) 2023 Nutanix Inc.  All rights reserved
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
)

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
