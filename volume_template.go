package metalcloud

//go:generate go run helper/gen_exports.go

//OperatingSystem describes an OS
type OperatingSystem struct {
	OperatingSystemType         string `json:"operating_system_type,omitempty"`
	OperatingSystemVersion      string `json:"operating_system_version,omitempty"`
	OperatingSystemArchitecture string `json:"operating_system_architecture,omitempty"`
}

//VolumeTemplate describes an OS template
type VolumeTemplate struct {
	VolumeTemplateID                      int             `json:"volume_template_id,omitempty"`
	VolumeTemplateLabel                   string          `json:"volume_template_label,omitempty"`
	VolumeTemplateSizeMBytes              int             `json:"volume_template_size_mbytes,omitempty"`
	VolumeTemplateDisplayName             string          `json:"volume_template_display_name,omitempty"`
	VolumeTemplateDescription             string          `json:"volume_template_description,omitempty"`
	VolumeTemplateLocalDiskSupported      bool            `json:"volume_template_local_supported,omitempty"`
	VolumeTemplateBootMethodsSupported    string          `json:"volume_template_boot_methods_supported,omitempty"`
	VolumeTemplateBootType                string          `json:"volume_template_boot_type,omitempty"`
	VolumeTemplateDeprecationStatus       string          `json:"volume_template_deprecation_status,omitempty"`
	VolumeTemplateRepoURL                 string          `json:"volume_template_repo_url,omitempty"`
	VolumeTemplateOperatingSystem         OperatingSystem `json:"volume_template_operating_system,omitempty"`
	VolumeTemplateTags                    []string        `json:"volume_template_tags,omitempty"`
	VolumeTemplateOsBootstrapFunctionName string          `json:"volume_template_os_bootstrap_function_name,omitempty"`
}

//VolumeTemplates retrives the list of available templates
func (c *Client) VolumeTemplates() (*map[string]VolumeTemplate, error) {
	res, err := c.rpcClient.Call(
		"volume_templates",
		c.user)

	if err != nil {
		return nil, err
	}

	_, ok := res.Result.([]interface{})
	if ok {
		var m = map[string]VolumeTemplate{}
		return &m, nil
	}

	var createdObject map[string]VolumeTemplate

	err2 := res.GetObject(&createdObject)
	if err2 != nil {
		return nil, err2
	}

	return &createdObject, nil
}

//volumeTemplateGet returns the specified volume template
func (c *Client) volumeTemplateGet(volumeTemplateID id) (*VolumeTemplate, error) {
	var createdObject VolumeTemplate

	if err := checkID(volumeTemplateID); err != nil {
		return nil, err
	}

	err := c.rpcClient.CallFor(
		&createdObject,
		"volume_template_get",
		volumeTemplateID)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

//volumeTemplateCreate creates a private volume template from a drive
func (c *Client) volumeTemplateCreateFromDrive(driveID id, objVolumeTemplate VolumeTemplate) (*VolumeTemplate, error) {
	var createdObject VolumeTemplate

	if err := checkID(driveID); err != nil {
		return nil, err
	}

	err := c.rpcClient.CallFor(
		&createdObject,
		"volume_template_create",
		driveID,
		objVolumeTemplate)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

//VolumeTemplateMakePublic makes a template public
func (c *Client) VolumeTemplateMakePublic(volumeTemplateID int) error {
	_, err := c.rpcClient.Call(
		"volume_template_make_public",
		volumeTemplateID,
	)

	if err != nil {
		return err
	}

	return nil
}

//VolumeTemplateMakePrivate makes a template private
func (c *Client) VolumeTemplateMakePrivate(volumeTemplateID int, userID int) error {
	_, err := c.rpcClient.Call(
		"volume_template_make_private",
		volumeTemplateID,
		userID,
	)

	if err != nil {
		return err
	}

	return nil
}
