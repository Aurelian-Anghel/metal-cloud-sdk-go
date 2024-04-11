// Code generated by gen_exports.go DO NOT EDIT

package metalcloud

//InstanceArrayGet 
func (c *Client) InstanceArrayGet(instanceArrayID int) (*InstanceArray, error) {
	return c.instanceArrayGet(instanceArrayID)
}

//InstanceArrayGetByLabel 
func (c *Client) InstanceArrayGetByLabel(instanceArrayLabel string) (*InstanceArray, error) {
	return c.instanceArrayGet(instanceArrayLabel)
}

//InstanceArrays 
func (c *Client) InstanceArrays(infrastructureID int) (*map[string]InstanceArray, error) {
	return c.instanceArrays(infrastructureID)
}

//InstanceArraysByLabel 
func (c *Client) InstanceArraysByLabel(infrastructureLabel string) (*map[string]InstanceArray, error) {
	return c.instanceArrays(infrastructureLabel)
}

//InstanceArrayCreate (colletion of identical instances). Requires Deploy.
func (c *Client) InstanceArrayCreate(infrastructureID int, instanceArray InstanceArray) (*InstanceArray, error) {
	return c.instanceArrayCreate(infrastructureID,instanceArray)
}

//InstanceArrayCreateByLabel (colletion of identical instances). Requires Deploy.
func (c *Client) InstanceArrayCreateByLabel(infrastructureLabel string, instanceArray InstanceArray) (*InstanceArray, error) {
	return c.instanceArrayCreate(infrastructureLabel,instanceArray)
}

//InstanceArrayEdit array. Requires deploy.
func (c *Client) InstanceArrayEdit(instanceArrayID int, instanceArrayOperation InstanceArrayOperation, bSwapExistingInstancesHardware *bool, bKeepDetachingDrives *bool, objServerTypeMatches *ServerTypeMatches, arrInstancesToBeDeleted *[]int) (*InstanceArray, error) {
	return c.instanceArrayEdit(instanceArrayID,instanceArrayOperation,bSwapExistingInstancesHardware,bKeepDetachingDrives,objServerTypeMatches,arrInstancesToBeDeleted)
}

//InstanceArrayEditByLabel array. Requires deploy.
func (c *Client) InstanceArrayEditByLabel(instanceArrayLabel string, instanceArrayOperation InstanceArrayOperation, bSwapExistingInstancesHardware *bool, bKeepDetachingDrives *bool, objServerTypeMatches *ServerTypeMatches, arrInstancesToBeDeleted *[]int) (*InstanceArray, error) {
	return c.instanceArrayEdit(instanceArrayLabel,instanceArrayOperation,bSwapExistingInstancesHardware,bKeepDetachingDrives,objServerTypeMatches,arrInstancesToBeDeleted)
}

//InstanceArrayDelete array. Requires deploy.
func (c *Client) InstanceArrayDelete(instanceArrayID int) error {
	return c.instanceArrayDelete(instanceArrayID)
}

//InstanceArrayDeleteByLabel array. Requires deploy.
func (c *Client) InstanceArrayDeleteByLabel(instanceArrayLabel string) error {
	return c.instanceArrayDelete(instanceArrayLabel)
}

//InstanceArrayStop InstanceArray.
func (c *Client) InstanceArrayStop(instanceArrayID int) (*InstanceArray, error) {
	return c.instanceArrayStop(instanceArrayID)
}

//InstanceArrayStopByLabel InstanceArray.
func (c *Client) InstanceArrayStopByLabel(instanceArrayLabel string) (*InstanceArray, error) {
	return c.instanceArrayStop(instanceArrayLabel)
}

//InstanceArrayStart InstanceArray.
func (c *Client) InstanceArrayStart(instanceArrayID int) (*InstanceArray, error) {
	return c.instanceArrayStart(instanceArrayID)
}

//InstanceArrayStartByLabel InstanceArray.
func (c *Client) InstanceArrayStartByLabel(instanceArrayLabel string) (*InstanceArray, error) {
	return c.instanceArrayStart(instanceArrayLabel)
}
