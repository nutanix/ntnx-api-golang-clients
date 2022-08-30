/*
 * Generated file models/vmm/v4/config/config_model.go.
 *
 * Product version: 4.0.1-alpha-1
 *
 * Part of the Nutanix Vmm Versioned APIs
 *
 * (c) 2022 Nutanix Inc.  All rights reserved
 *
 */

/*
  Configure Tasks
*/
package config

/**
Object encapsulating Task ID Return Value
*/
type Task struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  The external identifier that can be used to retrieve the task using its URL.
  */
  ExtId *string `json:"extId,omitempty"`
}


func NewTask() *Task {
  p := new(Task)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "vmm.v4.config.Task"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.config.Task"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}



