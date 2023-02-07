/*
 * Generated file models/storage/v4/common/common_model.go.
 *
 * Product version: 4.0.2-alpha-3
 *
 * Part of the Nutanix Storage Versioned APIs
 *
 * (c) 2023 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module storage.v4.common of Nutanix Storage Versioned APIs
*/
package common

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type OperationType int

const (
	OPERATIONTYPE_UNKNOWN    OperationType = 0
	OPERATIONTYPE_REDACTED   OperationType = 1
	OPERATIONTYPE_ADD        OperationType = 2
	OPERATIONTYPE_REMOVE     OperationType = 3
	OPERATIONTYPE_UPGRADE    OperationType = 4
	OPERATIONTYPE_REGISTER   OperationType = 5
	OPERATIONTYPE_START      OperationType = 6
	OPERATIONTYPE_STOP       OperationType = 7
	OPERATIONTYPE_DESTROY    OperationType = 8
	OPERATIONTYPE_DOWNLOAD   OperationType = 9
	OPERATIONTYPE_CONFIGURE  OperationType = 10
	OPERATIONTYPE_MIGRATE    OperationType = 11
	OPERATIONTYPE_ACTIVATE   OperationType = 12
	OPERATIONTYPE_DEACTIVATE OperationType = 13
	OPERATIONTYPE_RESTORE    OperationType = 14
	OPERATIONTYPE_REPLICATE  OperationType = 15
)

// returns the name of the enum given an ordinal number
func (e *OperationType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ADD",
		"REMOVE",
		"UPGRADE",
		"REGISTER",
		"START",
		"STOP",
		"DESTROY",
		"DOWNLOAD",
		"CONFIGURE",
		"MIGRATE",
		"ACTIVATE",
		"DEACTIVATE",
		"RESTORE",
		"REPLICATE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// returns the enum type given a string value
func (e *OperationType) index(name string) OperationType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ADD",
		"REMOVE",
		"UPGRADE",
		"REGISTER",
		"START",
		"STOP",
		"DESTROY",
		"DOWNLOAD",
		"CONFIGURE",
		"MIGRATE",
		"ACTIVATE",
		"DEACTIVATE",
		"RESTORE",
		"REPLICATE",
	}
	for idx := range names {
		if names[idx] == name {
			return OperationType(idx)
		}
	}
	return OPERATIONTYPE_UNKNOWN
}

func (e *OperationType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for OperationType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *OperationType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e OperationType) Ref() *OperationType {
	return &e
}

type ProgressStatus int

const (
	PROGRESSSTATUS_UNKNOWN   ProgressStatus = 0
	PROGRESSSTATUS_REDACTED  ProgressStatus = 1
	PROGRESSSTATUS_QUEUED    ProgressStatus = 2
	PROGRESSSTATUS_RUNNING   ProgressStatus = 3
	PROGRESSSTATUS_SUCCEEDED ProgressStatus = 4
	PROGRESSSTATUS_ABORTED   ProgressStatus = 5
	PROGRESSSTATUS_SUSPENDED ProgressStatus = 6
	PROGRESSSTATUS_FAILED    ProgressStatus = 7
)

// returns the name of the enum given an ordinal number
func (e *ProgressStatus) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"QUEUED",
		"RUNNING",
		"SUCCEEDED",
		"ABORTED",
		"SUSPENDED",
		"FAILED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// returns the enum type given a string value
func (e *ProgressStatus) index(name string) ProgressStatus {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"QUEUED",
		"RUNNING",
		"SUCCEEDED",
		"ABORTED",
		"SUSPENDED",
		"FAILED",
	}
	for idx := range names {
		if names[idx] == name {
			return ProgressStatus(idx)
		}
	}
	return PROGRESSSTATUS_UNKNOWN
}

func (e *ProgressStatus) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ProgressStatus:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ProgressStatus) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ProgressStatus) Ref() *ProgressStatus {
	return &e
}

type Task struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  The date-time string (RFC 3339) when the task was completed.
	*/
	CompletedTime *time.Time `json:"completedTime,omitempty"`
	/**
	  The date-time string (RFC 3339) when the task was created.
	*/
	CreatedTime *time.Time `json:"createdTime,omitempty"`
	/**
	  A string consisting of the description of the category as defined by the user.

	The server does not validate this value nor does it enforce the uniqueness or any other constraints.<br>
	It is the responsibility of the user to ensure that any semantic or syntactic constraints are retained when mutating
	this field.
	*/
	Description *string `json:"description,omitempty"`
	/**
	  The short name of this task. It may not be unique for each task.<br>
	This field is immutable.
	*/
	DisplayName *string `json:"displayName,omitempty"`
	/**
	  The extId referencing this task
	This field is immutable.
	*/
	ExtId *string `json:"extId,omitempty"`
	/**
	  The date-time string (RFC 3339) when the task was last updated.
	*/
	LastUpdatedTime *time.Time `json:"lastUpdatedTime,omitempty"`

	OperationType *OperationType `json:"operationType,omitempty"`
	/**
	  The extId referencing the parent task of this task (may be null if this task is a root task).<br>
	Each task can have at most one parent.<br>
	This field is immutable.
	*/
	ParentTaskExtId *string `json:"parentTaskExtId,omitempty"`
	/**
	  This will be a float value indicating the percentage completion of the task
	*/
	PercentageComplete *float32 `json:"percentageComplete,omitempty"`

	ProgressStatus *ProgressStatus `json:"progressStatus,omitempty"`
	/**
	  The date-time string (RFC 3339) when the task was started.
	*/
	StartedTime *time.Time `json:"startedTime,omitempty"`
	/**
	  The list of extIds referencing the sub-tasks of this task (may be empty if there are no sub-tasks).<br>
	This field is immutable.
	*/
	SubtaskExtIds []string `json:"subtaskExtIds,omitempty"`
}

func NewTask() *Task {
	p := new(Task)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "storage.v4.common.Task"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "storage.v4.r0.a3.common.Task"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}
