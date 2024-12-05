/*
 * Generated file models/security/v4/common/common_model.go.
 *
 * Product version: 4.0.1-beta-1
 *
 * Part of the Nutanix Security APIs
 *
 * (c) 2024 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module security.v4.common of Nutanix Security APIs
*/
package common

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

type Cluster struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  UUID for the cluster.
	*/
	ClusterExtId *string `json:"clusterExtId,omitempty"`
	/*
	  Human readable name for the cluster.
	*/
	ClusterName *string `json:"clusterName,omitempty"`
}

func NewCluster() *Cluster {
	p := new(Cluster)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "security.v4.common.Cluster"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type ClusterProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  UUID for the cluster.
	*/
	ClusterExtId *string `json:"clusterExtId,omitempty"`
	/*
	  Human readable name for the cluster.
	*/
	ClusterName *string `json:"clusterName,omitempty"`
}

func NewClusterProjection() *ClusterProjection {
	p := new(ClusterProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "security.v4.common.ClusterProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Contains potential different duration of time.
*/
type Timescale int

const (
	TIMESCALE_UNKNOWN  Timescale = 0
	TIMESCALE_REDACTED Timescale = 1
	TIMESCALE_DAY      Timescale = 2
	TIMESCALE_WEEK     Timescale = 3
	TIMESCALE_MONTH    Timescale = 4
	TIMESCALE_QUARTER  Timescale = 5
	TIMESCALE_YEAR     Timescale = 6
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *Timescale) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"DAY",
		"WEEK",
		"MONTH",
		"QUARTER",
		"YEAR",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e Timescale) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"DAY",
		"WEEK",
		"MONTH",
		"QUARTER",
		"YEAR",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *Timescale) index(name string) Timescale {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"DAY",
		"WEEK",
		"MONTH",
		"QUARTER",
		"YEAR",
	}
	for idx := range names {
		if names[idx] == name {
			return Timescale(idx)
		}
	}
	return TIMESCALE_UNKNOWN
}

func (e *Timescale) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for Timescale:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *Timescale) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e Timescale) Ref() *Timescale {
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
