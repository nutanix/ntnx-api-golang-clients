package vpnconnections

// This file holds the request struct for the GetVpnApplianceForVpnConnectionById operation.

type GetVpnApplianceForVpnConnectionByIdRequest struct {
	// (required) VPN connection UUID.
	VpnConnectionExtId *string

	// (required) Third-party VPN appliance UUID
	ExtId *string
}
