/*
 * Generated file models/storage/v4/config/config_model.go.
 *
 * Product version: 4.0.1-alpha-4
 *
 * Part of the Nutanix Dataprotection Versioned APIs
 *
 * (c) 2023 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module storage.v4.config of Nutanix Dataprotection Versioned APIs
*/
package config

/*
Object encapsulating Task ID return value.
*/
type Task struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The external identifier of the task.
	*/
	ExtId *string `json:"extId,omitempty"`
}

func NewTask() *Task {
	p := new(Task)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "storage.v4.config.Task"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "storage.v4.r0.a1.config.Task"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}
