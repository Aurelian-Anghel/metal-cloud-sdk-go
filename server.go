package metalcloud

import (
	"fmt"
	"strings"
)

//ServerSearchResult represents a server in a datacenter.
type ServerSearchResult struct {
	ServerID                       int               `json:"server_id,omitempty" yaml:"id,omitempty"`
	ServerUUID                     string            `json:"server_uuid,omitempty" yaml:"uuid,omitempty"`
	ServerStatus                   string            `json:"server_status,omitempty" yaml:"status,omitempty"`
	ServerSerialNumber             string            `json:"server_serial_number,omitempty" yaml:"serial_number,omitempty"`
	ServerVendor                   string            `json:"server_vendor,omitempty" yaml:"vendor,omitempty"`
	ServerNetworkTotalCapacityMbps int               `json:"server_network_total_capacity_mbps,omitempty" yaml:"networkTotalCapacityMbps,omitempty"`
	ServerBootType                 string            `json:"server_boot_type,omitempty" yaml:"bootType,omitempty"`
	ServerPowerStatus              string            `json:"server_power_status,omitempty" yaml:"powerStatus,omitempty"`
	ServerProcessorName            string            `json:"server_processor_name,omitempty" yaml:"processorName,omitempty"`
	ServerProcessorCoreCount       int               `json:"server_processor_core_count,omitempty" yaml:"processorCoreCount,omitempty"`
	ServerProcessorCoreMhz         int               `json:"server_processor_core_mhz,omitempty" yaml:"processorCoreMhz,omitempty"`
	ServerProcessorCount           int               `json:"server_processor_count,omitempty" yaml:"processorCount,omitempty"`
	ServerProcessorThreads         int               `json:"server_processor_threads,omitempty" yaml:"processorThreads,omitempty"`
	ServerProcessorCPUMark         int               `json:"server_processor_cpu_mark" yaml:"processorCPUMark"`
	ServerRAMGbytes                int               `json:"server_ram_gbytes,omitempty" yaml:"ramGbytes,omitempty"`
	ServerDiskCount                int               `json:"server_disk_count,omitempty" yaml:"diskCount,omitempty"`
	ServerDiskSizeMbytes           int               `json:"server_disk_size_mbytes,omitempty" yaml:"diskSizeMbytes,omitempty"`
	ServerDiskType                 string            `json:"server_disk_type,omitempty" yaml:"diskType,omitempty"`
	ServerProductName              string            `json:"server_product_name,omitempty" yaml:"productName,omitempty"`
	ServerClass                    string            `json:"server_class,omitempty" yaml:"class,omitempty"`
	ServerTypeID                   int               `json:"server_type_id,omitempty" yaml:"typeID,omitempty"`
	ServerTypeName                 string            `json:"server_type_name,omitempty" yaml:"type,omitempty"`
	ServerTypeBootType             string            `json:"server_type_boot_type,omitempty" yaml:"serverBootType,omitempty"`
	ServerInterfaces               []ServerInterface `json:"server_interfaces,omitempty" yaml:"interfaces,omitempty"`
	ServerDisks                    []ServerDisk      `json:"server_disks,omitempty" yaml:"disks,omitempty"`
	ServerTags                     []string          `json:"server_tags,omitempty" yaml:"tags,omitempty"`
	ServerIPMIChannel              int               `json:"server_ipmi_channel,omitempty" yaml:"IPMIChannel,omitempty"`
	ServerIPMIHost                 string            `json:"server_ipmi_host,omitempty" yaml:"IPMIHostname,omitempty"`
	ServerIPMInternalUsername      string            `json:"server_ipmi_internal_username,omitempty" yaml:"IPMIUsername,omitempty"`
	ServerIPMInternalPassword      string            `json:"server_ipmi_internal_password,omitempty" yaml:"IPMIPassword,omitempty"`
	ServerIPMCredentialsNeedUpdate bool              `json:"server_ipmi_credentials_need_update,omitempty" yaml:"IPMICredentialsNeedUpdate,omitempty"`
	ServerVendorSKUID              string            `json:"server_vendor_sku_id,omitempty" yaml:"vendorSKUID,omitempty"`
	ServerComments                 string            `json:"server_comments,omitempty" yaml:"comments,omitempty"`
	InstanceLabel                  []string          `json:"instance_label,omitempty" yaml:"instanceLabel,omitempty"`
	InstanceID                     []int             `json:"instance_id,omitempty" yaml:"instanceID,omitempty"`
	InstanceArrayID                []int             `json:"instance_array_id,omitempty" yaml:"instanceArrayID,omitempty"`
	InfrastructureID               []int             `json:"infrastructure_id,omitempty" yaml:"infrastructureID,omitempty"`
	UserEmail                      [][]string        `json:"user_email,omitempty" yaml:"userEmail,omitempty"`
	UserID                         [][]int           `json:"user_id,omitempty" yaml:"users,omitempty"`
	DatacenterName                 string            `json:"datacenter_name,omitempty" yaml:"datacenter,omitempty"`
}

//Server represents a server in a datacenter.
type Server struct {
	ServerID                       int               `json:"server_id,omitempty" yaml:"id,omitempty"`
	ServerUUID                     string            `json:"server_uuid,omitempty" yaml:"UUID,omitempty"`
	ServerStatus                   string            `json:"server_status,omitempty" yaml:"status,omitempty"`
	ServerSerialNumber             string            `json:"server_serial_number,omitempty" yaml:"serialNumber,omitempty"`
	ServerVendor                   string            `json:"server_vendor,omitempty" yaml:"vendor,omitempty"`
	DatacenterName                 string            `json:"datacenter_name,omitempty" yaml:"datacenter,omitempty"`
	ServerNetworkTotalCapacityMbps int               `json:"server_network_total_capacity_mbps,omitempty" yaml:"networkTotalCapacityMbps,omitempty"`
	ServerBootType                 string            `json:"server_boot_type,omitempty" yaml:"bootType,omitempty"`
	ServerPowerStatus              string            `json:"server_power_status,omitempty" yaml:"powerStatus,omitempty"`
	ServerProcessorName            string            `json:"server_processor_name,omitempty" yaml:"processorName,omitempty"`
	ServerProcessorCoreCount       int               `json:"server_processor_core_count,omitempty" yaml:"processorCoreCount,omitempty"`
	ServerProcessorCoreMhz         int               `json:"server_processor_core_mhz,omitempty" yaml:"processorCoreMhz,omitempty"`
	ServerProcessorCount           int               `json:"server_processor_count,omitempty" yaml:"processorCount,omitempty"`
	ServerProcessorThreads         int               `json:"server_processor_threads,omitempty" yaml:"processorThreads,omitempty"`
	ServerProcessorCPUMark         int               `json:"server_processor_cpu_mark" yaml:"processorCPUMark"`
	ServerRAMGbytes                int               `json:"server_ram_gbytes,omitempty" yaml:"ramGbytes,omitempty"`
	ServerDisks                    []ServerDisk      `json:"server_disks,omitempty" yaml:"disks,omitempty"`
	ServerDiskCount                int               `json:"server_disk_count,omitempty" yaml:"diskCount,omitempty"`
	ServerDiskSizeMbytes           int               `json:"server_disk_size_mbytes,omitempty" yaml:"diskSizeMbytes,omitempty"`
	ServerDiskType                 string            `json:"server_disk_type,omitempty" yaml:"diskType,omitempty"`
	ServerProductName              string            `json:"server_product_name,omitempty" yaml:"productName,omitempty"`
	ServerClass                    string            `json:"server_class,omitempty" yaml:"serverClass,omitempty"`
	ServerTypeID                   int               `json:"server_type_id,omitempty" yaml:"serverTypeID,omitempty"`
	ServerInterfaces               []ServerInterface `json:"server_interfaces,omitempty" yaml:"interfaces,omitempty"`
	ServerTags                     []string          `json:"server_tags,omitempty" yaml:"tags,omitempty"`
	ServerIPMIChannel              int               `json:"server_ipmi_channel" yaml:"IPMIChannel"`
	ServerIPMIHost                 string            `json:"server_ipmi_host,omitempty" yaml:"IPMIHostname,omitempty"`
	ServerIPMInternalUsername      string            `json:"server_ipmi_internal_username,omitempty" yaml:"IPMIUsername,omitempty"`
	ServerIPMInternalPassword      string            `json:"server_ipmi_internal_password,omitempty" yaml:"IPMIPassword,omitempty"`
	ServerIPMCredentialsNeedUpdate bool              `json:"server_ipmi_credentials_need_update,omitempty" yaml:"IPMICredentialsNeedUpdate,omitempty"`
	ServerVendorSKUID              string            `json:"server_vendor_sku_id,omitempty" yaml:"vendorSKU,omitempty"`
	ServerComments                 string            `json:"server_comments,omitempty" yaml:"comments,omitempty"`
	ServerBIOSInfoJSON             string            `json:"server_bios_info_json" yaml:"BIOSInfoJson"`
	ServerCustomJSON               string            `json:"server_custom_json" yaml:"CustomJSON"`
}

//ServerDisk describes a disk
type ServerDisk struct {
	ServerDiskID     int    `json:"server_disk_id,omitempty" yaml:"id,omitempty"`
	ServerDiskModel  string `json:"server_disk_model,omitempty" yaml:"model,omitempty"`
	ServerDiskType   string `json:"server_disk_type,omitempty" yaml:"type,omitempty"`
	ServerDiskVendor string `json:"server_disk_vendor,omitempty" yaml:"vendor,omitempty"`
	ServerDiskStatus string `json:"server_disk_status,omitempty" yaml:"status,omitempty"`
	ServerDiskSerial string `json:"server_disk_serial,omitempty" yaml:"serial_number,omitempty"`
	ServerDiskSizeGB int    `json:"server_disk_size_gb,omitempty" yaml:"sizeGB,omitempty"`
}

//ServerInterface contains server connectivity information.
type ServerInterface struct {
	ServerInterfaceMACAddress string `json:"server_interface_mac_address,omitempty" yaml:"macAddress,omitempty"`
}

//SearchResultForServers describes a serach result
type SearchResultForServers struct {
	DurationMilliseconds int                  `json:"duration_millisecnds,omitempty"`
	Rows                 []ServerSearchResult `json:"rows,omitempty"`
	RowsOrder            [][]string           `json:"rows_order,omitempty"`
	RowsTotal            int                  `json:"rows_total,omitempty"`
}

//ServerComponent information about a server's components
type ServerComponent struct {
	ServerComponentID                              int      `json:"server_component_id,omitempty" yaml:"id,omitempty"`
	ServerID                                       int      `json:"server_id,omitempty" yaml:"serverID,omitempty"`
	ServerComponentName                            string   `json:"server_component_name,omitempty" yaml:"componentName,omitempty"`
	ServerComponentFirmwareVersion                 string   `json:"server_component_firmware_version,omitempty" yaml:"firmwareVersion,omitempty"`
	ServerComponentFirmwareUpdateable              bool     `json:"server_component_firmware_updateable,omitempty" yaml:"firmwareUpdateable,omitempty"`
	ServerComponentFirmwareJSON                    string   `json:"server_component_firmware_json,omitempty" yaml:"firmwareJSON,omitempty"`
	ServerComponentFirmwareUpdateAvailableVersions []string `json:"server_component_firmware_update_available_versions,omitempty" yaml:"firmwareUpdateAvailableVersions,omitempty"`
	ServerComponentFirmwareStatus                  string   `json:"server_component_firmware_status,omitempty" yaml:"firmwareStatus,omitempty"`
	ServerComponentType                            string   `json:"server_component_type,omitempty" yaml:"type,omitempty"`
	ServerComponentFirmwareUpdateTimestamp         string   `json:"server_component_firmware_update_timestamp,omitempty" yaml:"firmwareUpdateTimestamp,omitempty"`
	ServerComponentFirmwareTargetVersion           string   `json:"server_component_firmware_target_version,omitempty" yaml:"firmwareTargetVersion,omitempty"`
	ServerComponentFirmwareScheduledTimestamp      string   `json:"server_component_firmware_scheduled_timestamp,omitempty" yaml:"firmwareScheduledTimestamp,omitempty"`
}

//SearchResultForServerComponents describes a search result
type SearchResultForServerComponents struct {
	DurationMilliseconds int               `json:"duration_millisecnds,omitempty"`
	Rows                 []ServerComponent `json:"rows,omitempty"`
	RowsOrder            [][]string        `json:"rows_order,omitempty"`
	RowsTotal            int               `json:"rows_total,omitempty"`
}

//ServersSearch searches for servers matching certain filter
func (c *Client) ServersSearch(filter string) (*[]ServerSearchResult, error) {

	tables := []string{"_servers_instances"}
	columns := map[string][]string{
		"_servers_instances": {
			"server_id",
			"server_type_name",
			"server_type_boot_type",
			"server_product_name",
			"datacenter_name",
			"server_status",
			"server_class",
			"server_created_timestamp",
			"server_vendor",
			"server_serial_number",
			"server_uuid",
			"server_bios_version",
			"server_vendor_sku_id",
			"server_boot_type",
			"server_allocation_timestamp",
			"instance_label",
			"instance_id",
			"instance_array_id",
			"infrastructure_id",
			"server_ipmi_host",
			"server_custom_json",
			"server_ipmi_internal_username",
			"server_ipmi_internal_password",
			"server_processor_name",
			"server_processor_count",
			"server_processor_core_count",
			"server_processor_core_mhz",
			"server_processor_threads",
			"server_processor_cpu_mark",
			"server_disk_type",
			"server_disk_count",
			"server_disk_size_mbytes",
			"server_ram_gbytes",
			"server_network_total_capacity_mbps",
			"server_dhcp_status",
			"server_dhcp_packet_sniffing_is_enabled",
			"server_dhcp_relay_security_is_enabled",
			"server_disk_wipe",
			"server_power_status",
			"server_power_status_last_update_timestamp",
			"user_id",
			"user_id_owner",
			"user_email",
		},
	}

	userID := c.GetUserID()

	collapseType := "array_row_span"

	res, err := c.rpcClient.Call(
		"search",
		userID,
		filter,
		tables,
		columns,
		collapseType)

	if err != nil {
		return nil, err
	}

	var createdObject map[string]SearchResultForServers

	err2 := res.GetObject(&createdObject)
	if err2 != nil {
		return nil, err2
	}

	servers := []ServerSearchResult{}
	for _, s := range createdObject[tables[0]].Rows {
		servers = append(servers, s)
	}

	return &servers, nil
}

//ServerGetByUUID retrieves information about a specified Server by using the server's UUID
func (c *Client) ServerGetByUUID(serverUUID string, decryptPasswd bool) (*Server, error) {

	var createdObject Server

	err := c.rpcClient.CallFor(
		&createdObject,
		"server_with_uuid_get",
		serverUUID)

	if err != nil {
		return nil, err
	}

	if decryptPasswd {
		passwdComponents := strings.Split(createdObject.ServerIPMInternalPassword, ":")
		if len(passwdComponents) != 2 {
			return nil, fmt.Errorf("Password not returned with proper components")
		}
		var passwd string
		err = c.rpcClient.CallFor(
			&passwd,
			"password_decrypt",
			passwdComponents[1],
		)
		if err != nil {
			return nil, err
		}
		createdObject.ServerIPMInternalPassword = passwd
	} else {
		createdObject.ServerIPMInternalPassword = ""
	}

	return &createdObject, nil

}

//ServerGet returns a server's details
func (c *Client) ServerGet(serverID int, decryptPasswd bool) (*Server, error) {

	var createdObject Server

	err := c.rpcClient.CallFor(
		&createdObject,
		"server_get",
		serverID)

	if err != nil {
		return nil, err
	}

	if decryptPasswd {
		var internalSrvObject Server
		err := c.rpcClient.CallFor(
			&internalSrvObject,
			"server_get_internal",
			serverID)

		if err != nil {
			return nil, err
		}

		passwdComponents := strings.Split(internalSrvObject.ServerIPMInternalPassword, ":")
		if len(passwdComponents) != 2 {
			return nil, fmt.Errorf("Password not returned with proper components")
		}
		var passwd string
		err = c.rpcClient.CallFor(
			&passwd,
			"password_decrypt",
			passwdComponents[1],
		)
		if err != nil {
			return nil, err
		}
		createdObject.ServerIPMInternalPassword = passwd
	} else {
		createdObject.ServerIPMInternalPassword = ""
	}

	return &createdObject, nil
}

//ServerCreate manually creates a server record
func (c *Client) ServerCreate(server Server, autoGenerate bool) (int, error) {

	var createdObject int

	err := c.rpcClient.CallFor(
		&createdObject,
		"server_create",
		server,
		autoGenerate,
	)

	if err != nil {
		return 0, err
	}

	return createdObject, nil
}

//ServerEditComplete - perform a complete edit
func (c *Client) ServerEditComplete(serverID int, server Server) (*Server, error) {
	return c.ServerEdit(serverID, "complete", server)
}

//ServerEditIPMI - edit only IPMI settings
func (c *Client) ServerEditIPMI(serverID int, server Server) (*Server, error) {
	return c.ServerEdit(serverID, "ipmi", server)
}

//ServerEditAvailability - edit only server availability settings
func (c *Client) ServerEditAvailability(serverID int, server Server) (*Server, error) {
	return c.ServerEdit(serverID, "availability", server)
}

//ServerEdit edits a server record
func (c *Client) ServerEdit(serverID int, serverEditType string, server Server) (*Server, error) {

	var createdObject Server

	err := c.rpcClient.CallFor(
		&createdObject,
		"server_edit",
		serverID,
		serverEditType,
		server,
	)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

//ServerDelete deletes all the information about a specified Server.
func (c *Client) ServerDelete(serverID int, skipIPMI bool) error {

	resp, err := c.rpcClient.Call("server_delete", serverID, skipIPMI)

	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	return nil
}

//ServerDecomission decomissions the server row and deletes all child rows.
func (c *Client) ServerDecomission(serverID int, skipIPMI bool) error {

	resp, err := c.rpcClient.Call("server_decomission", serverID, skipIPMI)

	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	return nil
}

//ServerFirmwareComponentUpgrade Creates a firmware upgrading session for the specified component.
//If no strServerComponentFirmwareNewVersion or strFirmwareBinaryURL are provided the system will use the values from the database which should have been previously added
func (c *Client) ServerFirmwareComponentUpgrade(serverID int, serverComponentID int, serverComponentFirmwareNewVersion string, firmwareBinaryURL string) error {

	_, err := c.rpcClient.Call(
		"server_firmware_component_upgrade",
		serverID,
		serverComponentID,
		serverComponentFirmwareNewVersion,
		firmwareBinaryURL,
	)
	return err
}

//ServerFirmwareUpgrade creates a firmware upgrading session that affects all components from the specified server that have a target version set and are updatable.
func (c *Client) ServerFirmwareUpgrade(serverID int) error {

	_, err := c.rpcClient.Call(
		"server_firmware_upgrade",
		serverID,
	)
	return err
}

//ServerFirmwareComponentTargetVersionSet Sets a firmware target version for the upgrading process. The system will apply the upgrade at the next upgrading session.
func (c *Client) ServerFirmwareComponentTargetVersionSet(serverComponentID int, serverComponentFirmwareNewVersion string) error {

	_, err := c.rpcClient.Call(
		"server_firmware_component_target_version_set",
		serverComponentID,
		serverComponentFirmwareNewVersion,
	)
	return err
}

//ServerFirmwareComponentTargetVersionUpdate Updates for every component of the specified server the available firmware versions that can be used as target by the firmware upgrading process. The available versions are extracted from a vendor specific catalog.
func (c *Client) ServerFirmwareComponentTargetVersionUpdate(serverComponentID int) error {

	_, err := c.rpcClient.Call(
		"server_firmware_component_available_versions_update",
		serverComponentID,
	)
	return err
}

//ServerFirmwareComponentTargetVersionAdd Adds a new available firmware version for a server component along with the url of the binary. If the version already exists the old url will be overwritten.
func (c *Client) ServerFirmwareComponentTargetVersionAdd(serverComponentID int, version string, firmareBinaryURL string) error {

	_, err := c.rpcClient.Call(
		"server_firmware_component_available_versions_add",
		serverComponentID,
		version,
		firmareBinaryURL,
	)
	return err
}

//ServerComponentGet returns a server's component's details
func (c *Client) ServerComponentGet(serverComponentID int) (*ServerComponent, error) {

	var createdObject ServerComponent

	err := c.rpcClient.CallFor(
		&createdObject,
		"server_get_internal",
		serverComponentID)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

//ServerComponents searches for servers matching certain filter
func (c *Client) ServerComponents(serverID int, filter string) (*[]ServerComponent, error) {

	tables := []string{"_server_components"}
	columns := map[string][]string{
		"_server_components": {
			"server_component_name",
			"server_component_id",
			"server_component_firmware_json",
			"server_component_type",
			"server_component_firmware_update_timestamp",
			"server_component_firmware_status",
			"server_component_firmware_update_available_versions",
			"server_component_firmware_updateable",
			"server_component_firmware_version",
			"server_component_firmware_status",
		},
	}

	userID := c.GetUserID()

	collapseType := "none"
	filterWithServerID := fmt.Sprintf("+server_id:%d %s", serverID, filter)
	res, err := c.rpcClient.Call(
		"search",
		userID,
		filterWithServerID,
		tables,
		columns,
		collapseType)

	if err != nil {
		return nil, err
	}

	var createdObject map[string]SearchResultForServerComponents

	err2 := res.GetObject(&createdObject)
	if err2 != nil {
		return nil, err2
	}

	list := []ServerComponent{}
	for _, s := range createdObject[tables[0]].Rows {
		list = append(list, s)
	}

	return &list, nil
}

//ServerPowerSet reboots or powers on a server
func (c *Client) ServerPowerSet(serverID int, operation string) error {
	if err := checkID(serverID); err != nil {
		return err
	}

	resp, err := c.rpcClient.Call(
		"server_power_set",
		serverID,
		operation)

	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	return nil
}
