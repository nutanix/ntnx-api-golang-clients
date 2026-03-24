package tasks

// This file holds the request struct for the CancelTask operation.

type CancelTaskRequest struct {
	// (required) A globally unique identifier for a task. It consists of a prefix and a UUID separated by ':'. The 'legacy' prefix can be
	// used with a task UUID provided by previous API families.
	TaskExtId *string
}
