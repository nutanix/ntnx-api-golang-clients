package objectstores

// This file holds the request struct for the CreateCertificate operation.

type CreateCertificateRequest struct {
	// (required) The UUID of the Object store.
	ObjectStoreExtId *string

	// (required) A JSON file which contains the public certificates, private key, and CA certificate or chain, along with a list of
	// alternate FQDNs and alternate IPs to create a certificate for the Object store.
	// | | |
	// | ----------- | ----------- |
	// | alternateFqdns | The list of alternate FQDNs for accessing the Object store. The FQDNs must consist of at least 2
	// parts separated by a '.'. Each part can contain upper and lower case letters, digits, hyphens or underscores but must
	// begin and end with a letter. Each part can be up to 63 characters long. For e.g 'objects-0.pc_nutanix.com'. |
	// | alternateIps | A list of the IPs included as Subject Alternative Names (SANs) in the certificate. The IPs must be
	// among the public IPs of the Object store (publicNetworkIps). |
	// | ca | The CA certificate or chain to upload. |
	// | publicCert | The public certificate to upload. |
	// | privateKey | The private key to upload. |
	// | shouldGenerate | If true, a new certificate is generated with the provided alternate FQDNs and IPs. |
	// ```
	// {
	//   "alternateFqdns": [                             // [ 1 .. 999 ] items
	//     {
	// "value": "fqdn1.nutanix.com"                // string
	// ^(?:[a-z0-9][\w\-]*[a-z0-9]*\.)*(?:(?:(?:[a-z0-9][\w\-]*[a-z0-9]*)(?:\.[a-z0-9]+)?))$
	//     },
	//     {
	// "value": "fqdn2.nutanix.com"                // string
	// ^(?:[a-z0-9][\w\-]*[a-z0-9]*\.)*(?:(?:(?:[a-z0-9][\w\-]*[a-z0-9]*)(?:\.[a-z0-9]+)?))$
	//     },
	//     ...
	//   ],
	//   "alternateIps": [                               // [ 1 .. 999 ] items
	//     {
	//       "ipv4": {
	// "value": "92.41.252.152"                  // string
	// ^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$
	//     },
	//     {
	//       "ipv4": {
	// "value": "92.41.252.153"                  // string
	// ^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$
	//     },
	//     ...
	//   ],
	//   "shouldGenerate": true,                         // boolean
	// "ca": "-----BEGIN CERTIFICATE-----\nMIIDzTCCArWgAwIBAgIUI...\n-----END CERTIFICATE-----",                 // string [
	// 1 .. 2000000 ] characters
	// "publicCert": "-----BEGIN CERTIFICATE-----\nMIIDzTCCArWgAwIBAgIUI...\n-----END CERTIFICATE-----",         // string [
	// 1 .. 2000000 ] characters
	// "privateKey": "-----BEGIN RSA PRIVATE KEY-----\nMIIDzTCCArWgAwIBAgIUI...\n-----END RSA PRIVATE KEY-----"  // string [
	// 1 .. 2000000 ] characters
	// }
	// ```
	Path *string
}
