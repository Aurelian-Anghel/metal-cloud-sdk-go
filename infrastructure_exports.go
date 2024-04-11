// Code generated by gen_exports.go DO NOT EDIT

package metalcloud

//InfrastructureEdit 
func (c *Client) InfrastructureEdit(infrastructureID int, infrastructureOperation InfrastructureOperation) (*Infrastructure, error) {
	return c.infrastructureEdit(infrastructureID,infrastructureOperation)
}

//InfrastructureEditByLabel 
func (c *Client) InfrastructureEditByLabel(infrastructureLabel string, infrastructureOperation InfrastructureOperation) (*Infrastructure, error) {
	return c.infrastructureEdit(infrastructureLabel,infrastructureOperation)
}

//InfrastructureDelete elements. Requires deploy
func (c *Client) InfrastructureDelete(infrastructureID int) error {
	return c.infrastructureDelete(infrastructureID)
}

//InfrastructureDeleteByLabel elements. Requires deploy
func (c *Client) InfrastructureDeleteByLabel(infrastructureLabel string) error {
	return c.infrastructureDelete(infrastructureLabel)
}

//InfrastructureOperationCancel (undos) alterations done before deploy
func (c *Client) InfrastructureOperationCancel(infrastructureID int) error {
	return c.infrastructureOperationCancel(infrastructureID)
}

//InfrastructureOperationCancelByLabel (undos) alterations done before deploy
func (c *Client) InfrastructureOperationCancelByLabel(infrastructureLabel string) error {
	return c.infrastructureOperationCancel(infrastructureLabel)
}

//InfrastructureDeploy 
func (c *Client) InfrastructureDeploy(infrastructureID int, shutdownOptions ShutdownOptions, allowDataLoss bool, skipAnsible bool) error {
	return c.infrastructureDeploy(infrastructureID,shutdownOptions,allowDataLoss,skipAnsible)
}

//InfrastructureDeployByLabel 
func (c *Client) InfrastructureDeployByLabel(infrastructureLabel string, shutdownOptions ShutdownOptions, allowDataLoss bool, skipAnsible bool) error {
	return c.infrastructureDeploy(infrastructureLabel,shutdownOptions,allowDataLoss,skipAnsible)
}

//InfrastructureDeployWithOptions infrastructure. With options.
func (c *Client) InfrastructureDeployWithOptions(infrastructureID int, shutdownOptions ShutdownOptions, deployOptions *DeployOptions, allowDataLoss bool, skipAnsible bool) error {
	return c.infrastructureDeployWithOptions(infrastructureID,shutdownOptions,deployOptions,allowDataLoss,skipAnsible)
}

//InfrastructureDeployWithOptionsByLabel infrastructure. With options.
func (c *Client) InfrastructureDeployWithOptionsByLabel(infrastructureLabel string, shutdownOptions ShutdownOptions, deployOptions *DeployOptions, allowDataLoss bool, skipAnsible bool) error {
	return c.infrastructureDeployWithOptions(infrastructureLabel,shutdownOptions,deployOptions,allowDataLoss,skipAnsible)
}

//InfrastructureGet 
func (c *Client) InfrastructureGet(infrastructureID int) (*Infrastructure, error) {
	return c.infrastructureGet(infrastructureID)
}

//InfrastructureGetByLabel 
func (c *Client) InfrastructureGetByLabel(infrastructureLabel string) (*Infrastructure, error) {
	return c.infrastructureGet(infrastructureLabel)
}

//InfrastructureUserLimits 
func (c *Client) InfrastructureUserLimits(infrastructureID int) (*map[string]interface{}, error) {
	return c.infrastructureUserLimits(infrastructureID)
}

//InfrastructureUserLimitsByLabel 
func (c *Client) InfrastructureUserLimitsByLabel(infrastructureLabel string) (*map[string]interface{}, error) {
	return c.infrastructureUserLimits(infrastructureLabel)
}
