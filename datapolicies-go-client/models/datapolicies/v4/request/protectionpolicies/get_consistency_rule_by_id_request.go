package protectionpolicies

// This file holds the request struct for the GetConsistencyRuleById operation.

type GetConsistencyRuleByIdRequest struct {
	// (required) The external identifier of the protection policy.
	ProtectionPolicyExtId *string

	// (required) The external identifier of the consistency rule.
	ExtId *string
}
