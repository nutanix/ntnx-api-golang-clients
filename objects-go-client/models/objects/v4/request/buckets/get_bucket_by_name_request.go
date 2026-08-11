package buckets

// This file holds the request struct for the GetBucketByName operation.

type GetBucketByNameRequest struct {
	// (required) The UUID of the Object store.
	ObjectStoreExtId *string

	// (required) "The name of the bucket. It must contain only lowercase letters, numbers, hyphens and dots. It must start with a letter
	// or number and end with a letter or number. It must be between 3 and 63 characters long. It must not contain consecutive
	// dots or end with a dot. It must not be an IP address."
	BucketName *string

	// "Name of the federated namespace that this Object store belongs to. It must contain only letters, numbers and hyphens
	// but not consecutive hyphens. It must start with a letter and end with a letter or number. It must be between 1 and 16
	// characters long. If not provided, local namespace is used."
	XNtnxObjectsNamespace *string
}
