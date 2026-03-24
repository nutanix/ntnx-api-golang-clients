package subnets

// This file holds the request struct for the GetSubnetById operation.

type GetSubnetByIdRequest struct {
	// (required) UUID of the subnet.
	ExtId *string
}
