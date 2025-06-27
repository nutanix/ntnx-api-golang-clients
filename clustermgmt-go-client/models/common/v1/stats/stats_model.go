/*
 * Generated file models/common/v1/stats/stats_model.go.
 *
 * Product version: 4.1.1
 *
 * Part of the Nutanix Cluster Management APIs
 *
 * (c) 2025 Nutanix Inc.  All rights reserved
 *
 */

/*
  Nutanix Stats Configuration
*/
package stats

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

/*
The operator to use while performing down-sampling on stats data. Allowed values are SUM, MIN, MAX, AVG, COUNT and LAST.
*/
type DownSamplingOperator int

const (
	DOWNSAMPLINGOPERATOR_UNKNOWN  DownSamplingOperator = 0
	DOWNSAMPLINGOPERATOR_REDACTED DownSamplingOperator = 1
	DOWNSAMPLINGOPERATOR_SUM      DownSamplingOperator = 2
	DOWNSAMPLINGOPERATOR_MIN      DownSamplingOperator = 3
	DOWNSAMPLINGOPERATOR_MAX      DownSamplingOperator = 4
	DOWNSAMPLINGOPERATOR_AVG      DownSamplingOperator = 5
	DOWNSAMPLINGOPERATOR_COUNT    DownSamplingOperator = 6
	DOWNSAMPLINGOPERATOR_LAST     DownSamplingOperator = 7
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *DownSamplingOperator) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SUM",
		"MIN",
		"MAX",
		"AVG",
		"COUNT",
		"LAST",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e DownSamplingOperator) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SUM",
		"MIN",
		"MAX",
		"AVG",
		"COUNT",
		"LAST",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *DownSamplingOperator) index(name string) DownSamplingOperator {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SUM",
		"MIN",
		"MAX",
		"AVG",
		"COUNT",
		"LAST",
	}
	for idx := range names {
		if names[idx] == name {
			return DownSamplingOperator(idx)
		}
	}
	return DOWNSAMPLINGOPERATOR_UNKNOWN
}

func (e *DownSamplingOperator) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for DownSamplingOperator:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *DownSamplingOperator) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e DownSamplingOperator) Ref() *DownSamplingOperator {
	return &e
}

/*
A time value pair representing a stat associated with a given entity at a given point of date and time represented in extended ISO-8601 format."
*/
type TimeIntValuePair struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The date and time at which the stat was recorded.The value should be in extended ISO-8601 format. For example, start time of 2022-04-23T01:23:45.678+09:00 would consider all stats starting at 1:23:45.678 on the 23rd of April 2022. Details around ISO-8601 format can be found at https://www.iso.org/standard/70907.html
	*/
	Timestamp *time.Time `json:"timestamp,omitempty"`
	/*
	  Value of the stat at the recorded date and time in extended ISO-8601 format."
	*/
	Value *int64 `json:"value,omitempty"`
}

func (p *TimeIntValuePair) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias TimeIntValuePair

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

func (p *TimeIntValuePair) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias TimeIntValuePair
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = TimeIntValuePair(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "timestamp")
	delete(allFields, "value")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewTimeIntValuePair() *TimeIntValuePair {
	p := new(TimeIntValuePair)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "common.v1.stats.TimeIntValuePair"
	p.Reserved_ = map[string]interface{}{"$fv": "v1.r0"}
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
