// DO NOT EDIT: this file is automatically generated by docgen
package metalcloud

import (
	"github.com/projectdiscovery/yamldoc-go/encoder"
)

var (
	NetworkProfileDoc           encoder.Doc
	NetworkProfileVLANDoc       encoder.Doc
	NetworkProfileSubnetPoolDoc encoder.Doc
)

func init() {
	NetworkProfileDoc.Type = "NetworkProfile"
	NetworkProfileDoc.Comments[encoder.LineComment] = "NetworkProfile  A network profile modifies the default network configuration of an instance array when attached to a specific Network."
	NetworkProfileDoc.Description = "NetworkProfile  A network profile modifies the default network configuration of an instance array when attached to a specific Network."
	NetworkProfileDoc.Fields = make([]encoder.Doc, 8)
	NetworkProfileDoc.Fields[0].Name = "id"
	NetworkProfileDoc.Fields[0].Type = "int"
	NetworkProfileDoc.Fields[0].Note = ""
	NetworkProfileDoc.Fields[0].Description = "The id of the network profile"
	NetworkProfileDoc.Fields[0].Comments[encoder.LineComment] = "The id of the network profile"
	NetworkProfileDoc.Fields[1].Name = "label"
	NetworkProfileDoc.Fields[1].Type = "string"
	NetworkProfileDoc.Fields[1].Note = ""
	NetworkProfileDoc.Fields[1].Description = "The label of the network profile"
	NetworkProfileDoc.Fields[1].Comments[encoder.LineComment] = "The label of the network profile"
	NetworkProfileDoc.Fields[2].Name = "dc"
	NetworkProfileDoc.Fields[2].Type = "string"
	NetworkProfileDoc.Fields[2].Note = ""
	NetworkProfileDoc.Fields[2].Description = "The label of the datacenter on which this network profile applies"
	NetworkProfileDoc.Fields[2].Comments[encoder.LineComment] = "The label of the datacenter on which this network profile applies"
	NetworkProfileDoc.Fields[3].Name = "networkProfileIsPublic"
	NetworkProfileDoc.Fields[3].Type = "bool"
	NetworkProfileDoc.Fields[3].Note = ""
	NetworkProfileDoc.Fields[3].Description = "description: If set to true any of the users can use this network profile. If set to false the network profile needs to be\nexplicitly allowed in the user limits for a user to be able to use it\n"
	NetworkProfileDoc.Fields[3].Comments[encoder.LineComment] = "description: If set to true any of the users can use this network profile. If set to false the network profile needs to be"
	NetworkProfileDoc.Fields[4].Name = "networkType"
	NetworkProfileDoc.Fields[4].Type = "string"
	NetworkProfileDoc.Fields[4].Note = ""
	NetworkProfileDoc.Fields[4].Description = "Type of network to which this network profile can be applied"
	NetworkProfileDoc.Fields[4].Comments[encoder.LineComment] = "Type of network to which this network profile can be applied"
	NetworkProfileDoc.Fields[4].Values = []string{
		"wan",
		"lan",
		"san",
	}
	NetworkProfileDoc.Fields[5].Name = "vlans"
	NetworkProfileDoc.Fields[5].Type = "[]NetworkProfileVLAN"
	NetworkProfileDoc.Fields[5].Note = ""
	NetworkProfileDoc.Fields[5].Description = "VLAN (L2 network) entries"
	NetworkProfileDoc.Fields[5].Comments[encoder.LineComment] = "VLAN (L2 network) entries"
	NetworkProfileDoc.Fields[6].Name = "createdTimestamp"
	NetworkProfileDoc.Fields[6].Type = "string"
	NetworkProfileDoc.Fields[6].Note = ""
	NetworkProfileDoc.Fields[6].Description = "ISO 8601 timestamp which holds the date and time when the network profile was created."
	NetworkProfileDoc.Fields[6].Comments[encoder.LineComment] = "ISO 8601 timestamp which holds the date and time when the network profile was created."
	NetworkProfileDoc.Fields[7].Name = "updatedTimestamp"
	NetworkProfileDoc.Fields[7].Type = "string"
	NetworkProfileDoc.Fields[7].Note = ""
	NetworkProfileDoc.Fields[7].Description = "ISO 8601 timestamp which holds the date and time when the network profile was updated."
	NetworkProfileDoc.Fields[7].Comments[encoder.LineComment] = "ISO 8601 timestamp which holds the date and time when the network profile was updated."

	NetworkProfileVLANDoc.Type = "NetworkProfileVLAN"
	NetworkProfileVLANDoc.Comments[encoder.LineComment] = "NetworkProfileVLAN object describes an overlay L2 network which usually translates into a VLAN and"
	NetworkProfileVLANDoc.Description = "NetworkProfileVLAN object describes an overlay L2 network which usually translates into a VLAN and\nassociated VXLAN VNIs or EPGs depending on the vendor\n"
	NetworkProfileVLANDoc.AppearsIn = []encoder.Appearance{
		{
			TypeName:  "NetworkProfile",
			FieldName: "vlans",
		},
	}
	NetworkProfileVLANDoc.Fields = make([]encoder.Doc, 9)
	NetworkProfileVLANDoc.Fields[0].Name = "vlanID"
	NetworkProfileVLANDoc.Fields[0].Type = "int"
	NetworkProfileVLANDoc.Fields[0].Note = ""
	NetworkProfileVLANDoc.Fields[0].Description = "description: the VLAN ID to use (on EVPNVXLAN provisioner this translates to a VLAN ID to be used as VTEP for VXLAN).\nOn VLAN provisioner the VLAN will be used end-to-end. On SDN provisioner with CISCO ACI it will be translated into a VLAN attached to an EPG.\nIf set to null it will be automatically allocated based on the settings in the Datacenter config object.\n"
	NetworkProfileVLANDoc.Fields[0].Comments[encoder.LineComment] = "description: the VLAN ID to use (on EVPNVXLAN provisioner this translates to a VLAN ID to be used as VTEP for VXLAN)."
	NetworkProfileVLANDoc.Fields[1].Name = "label"
	NetworkProfileVLANDoc.Fields[1].Type = "int"
	NetworkProfileVLANDoc.Fields[1].Note = ""
	NetworkProfileVLANDoc.Fields[1].Description = "the label of this VLAN entry"
	NetworkProfileVLANDoc.Fields[1].Comments[encoder.LineComment] = "the label of this VLAN entry"
	NetworkProfileVLANDoc.Fields[2].Name = "portMode"
	NetworkProfileVLANDoc.Fields[2].Type = "string"
	NetworkProfileVLANDoc.Fields[2].Note = ""
	NetworkProfileVLANDoc.Fields[2].Description = "Type of VLAN"
	NetworkProfileVLANDoc.Fields[2].Comments[encoder.LineComment] = "Type of VLAN"
	NetworkProfileVLANDoc.Fields[2].Values = []string{
		"native",
		"trunk",
		"access",
	}
	NetworkProfileVLANDoc.Fields[3].Name = "provisionSubnetGateways"
	NetworkProfileVLANDoc.Fields[3].Type = "bool"
	NetworkProfileVLANDoc.Fields[3].Note = ""
	NetworkProfileVLANDoc.Fields[3].Description = "description: if true will provision subnet gateways on the switches. This depends on the provisioner and vendor: Symmetric IRB, VRRP-based redundant gateway configuration etc."
	NetworkProfileVLANDoc.Fields[3].Comments[encoder.LineComment] = "description: if true will provision subnet gateways on the switches. This depends on the provisioner and vendor: Symmetric IRB, VRRP-based redundant gateway configuration etc."
	NetworkProfileVLANDoc.Fields[4].Name = "provisionVXLAN"
	NetworkProfileVLANDoc.Fields[4].Type = "bool"
	NetworkProfileVLANDoc.Fields[4].Note = ""
	NetworkProfileVLANDoc.Fields[4].Description = "if false it will not configure the vxlan tunnel. Equivalent to the VLAN provisioner."
	NetworkProfileVLANDoc.Fields[4].Comments[encoder.LineComment] = "if false it will not configure the vxlan tunnel. Equivalent to the VLAN provisioner."
	NetworkProfileVLANDoc.Fields[5].Name = "extConnectionIDs"
	NetworkProfileVLANDoc.Fields[5].Type = "[]int"
	NetworkProfileVLANDoc.Fields[5].Note = ""
	NetworkProfileVLANDoc.Fields[5].Description = "If any external connection id is configured in the array the L2 network will be extended to that external connection point"
	NetworkProfileVLANDoc.Fields[5].Comments[encoder.LineComment] = "If any external connection id is configured in the array the L2 network will be extended to that external connection point"
	NetworkProfileVLANDoc.Fields[6].Name = "subnetPools"
	NetworkProfileVLANDoc.Fields[6].Type = "[]NetworkProfileSubnetPool"
	NetworkProfileVLANDoc.Fields[6].Note = ""
	NetworkProfileVLANDoc.Fields[6].Description = "Subnet pools to allocate IPs from"
	NetworkProfileVLANDoc.Fields[6].Comments[encoder.LineComment] = "Subnet pools to allocate IPs from"
	NetworkProfileVLANDoc.Fields[7].Name = "vlanAutoAllocationIndex"
	NetworkProfileVLANDoc.Fields[7].Type = "int"
	NetworkProfileVLANDoc.Fields[7].Note = ""
	NetworkProfileVLANDoc.Fields[7].Description = "If set to a non-null value this will help ensure that two different network profiles can get the same auto-allocated VLAN ID"
	NetworkProfileVLANDoc.Fields[7].Comments[encoder.LineComment] = "If set to a non-null value this will help ensure that two different network profiles can get the same auto-allocated VLAN ID"
	NetworkProfileVLANDoc.Fields[8].Name = "vrfID"
	NetworkProfileVLANDoc.Fields[8].Type = "int"
	NetworkProfileVLANDoc.Fields[8].Note = ""
	NetworkProfileVLANDoc.Fields[8].Description = "The VRF ID to use. If set to null it will fall back to automatically allocated VRFs"
	NetworkProfileVLANDoc.Fields[8].Comments[encoder.LineComment] = "The VRF ID to use. If set to null it will fall back to automatically allocated VRFs"

	NetworkProfileSubnetPoolDoc.Type = "NetworkProfileSubnetPool"
	NetworkProfileSubnetPoolDoc.Comments[encoder.LineComment] = ""
	NetworkProfileSubnetPoolDoc.Description = ""
	NetworkProfileSubnetPoolDoc.AppearsIn = []encoder.Appearance{
		{
			TypeName:  "NetworkProfileVLAN",
			FieldName: "subnetPools",
		},
	}
	NetworkProfileSubnetPoolDoc.Fields = make([]encoder.Doc, 3)
	NetworkProfileSubnetPoolDoc.Fields[0].Name = "subnetPoolID"
	NetworkProfileSubnetPoolDoc.Fields[0].Type = "int"
	NetworkProfileSubnetPoolDoc.Fields[0].Note = ""
	NetworkProfileSubnetPoolDoc.Fields[0].Description = "The ID of the subnet pool to use. Set to null for auto in which case the type of subnet pool will be used to pick a subnet pool"
	NetworkProfileSubnetPoolDoc.Fields[0].Comments[encoder.LineComment] = "The ID of the subnet pool to use. Set to null for auto in which case the type of subnet pool will be used to pick a subnet pool"
	NetworkProfileSubnetPoolDoc.Fields[1].Name = "subnetPoolType"
	NetworkProfileSubnetPoolDoc.Fields[1].Type = "string"
	NetworkProfileSubnetPoolDoc.Fields[1].Note = ""
	NetworkProfileSubnetPoolDoc.Fields[1].Description = "description: The type of subnet pool to use\n- ipv4\n- ipv6\n"
	NetworkProfileSubnetPoolDoc.Fields[1].Comments[encoder.LineComment] = "description: The type of subnet pool to use"
	NetworkProfileSubnetPoolDoc.Fields[2].Name = "subnetPoolProvidesDefaultRoute"
	NetworkProfileSubnetPoolDoc.Fields[2].Type = "bool"
	NetworkProfileSubnetPoolDoc.Fields[2].Note = ""
	NetworkProfileSubnetPoolDoc.Fields[2].Description = "If set to true this subnet pool will be used to set the default route"
	NetworkProfileSubnetPoolDoc.Fields[2].Comments[encoder.LineComment] = "If set to true this subnet pool will be used to set the default route"
}

func (NetworkProfile) Doc() *encoder.Doc {
	return &NetworkProfileDoc
}

func (NetworkProfileVLAN) Doc() *encoder.Doc {
	return &NetworkProfileVLANDoc
}

func (NetworkProfileSubnetPool) Doc() *encoder.Doc {
	return &NetworkProfileSubnetPoolDoc
}

// GetNetworkProfileDoc returns documentation for the file ./.
func GetNetworkProfileDoc() *encoder.FileDoc {
	return &encoder.FileDoc{
		Name:        "NetworkProfile",
		Description: "",
		Structs: []*encoder.Doc{
			&NetworkProfileDoc,
			&NetworkProfileVLANDoc,
			&NetworkProfileSubnetPoolDoc,
		},
	}
}
