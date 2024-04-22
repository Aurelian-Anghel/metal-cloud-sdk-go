// DO NOT EDIT: this file is automatically generated by docgen
package metalcloud

import (
	"github.com/projectdiscovery/yamldoc-go/encoder"
)

var (
	SharedDriveDoc            encoder.Doc
	SharedDriveCredentialsDoc encoder.Doc
	SharedDriveOperationDoc   encoder.Doc
)

func init() {
	SharedDriveDoc.Type = "SharedDrive"
	SharedDriveDoc.Comments[encoder.LineComment] = "SharedDrive represents a drive that can be shared between instances"
	SharedDriveDoc.Description = "SharedDrive represents a drive that can be shared between instances"
	SharedDriveDoc.Fields = make([]encoder.Doc, 18)
	SharedDriveDoc.Fields[0].Name = "label"
	SharedDriveDoc.Fields[0].Type = "string"
	SharedDriveDoc.Fields[0].Note = ""
	SharedDriveDoc.Fields[0].Description = "Label of the shared drive"
	SharedDriveDoc.Fields[0].Comments[encoder.LineComment] = "Label of the shared drive"
	SharedDriveDoc.Fields[1].Name = "subdomain"
	SharedDriveDoc.Fields[1].Type = "string"
	SharedDriveDoc.Fields[1].Note = ""
	SharedDriveDoc.Fields[1].Description = "Unique string representing the shared drive"
	SharedDriveDoc.Fields[1].Comments[encoder.LineComment] = "Unique string representing the shared drive"
	SharedDriveDoc.Fields[2].Name = "id"
	SharedDriveDoc.Fields[2].Type = "int"
	SharedDriveDoc.Fields[2].Note = ""
	SharedDriveDoc.Fields[2].Description = "Id of the shared drive"
	SharedDriveDoc.Fields[2].Comments[encoder.LineComment] = "Id of the shared drive"
	SharedDriveDoc.Fields[3].Name = "storageType"
	SharedDriveDoc.Fields[3].Type = "string"
	SharedDriveDoc.Fields[3].Note = ""
	SharedDriveDoc.Fields[3].Description = "Type of the shared drive"
	SharedDriveDoc.Fields[3].Comments[encoder.LineComment] = "Type of the shared drive"
	SharedDriveDoc.Fields[3].Values = []string{
		"iscsi_ssd",
		"iscsi_hdd",
	}
	SharedDriveDoc.Fields[4].Name = "infrastructureID"
	SharedDriveDoc.Fields[4].Type = "int"
	SharedDriveDoc.Fields[4].Note = ""
	SharedDriveDoc.Fields[4].Description = "ID of the infrastructure"
	SharedDriveDoc.Fields[4].Comments[encoder.LineComment] = "ID of the infrastructure"
	SharedDriveDoc.Fields[5].Name = "serviceStatus"
	SharedDriveDoc.Fields[5].Type = "string"
	SharedDriveDoc.Fields[5].Note = ""
	SharedDriveDoc.Fields[5].Description = "Service status of the shared drive"
	SharedDriveDoc.Fields[5].Comments[encoder.LineComment] = "Service status of the shared drive"
	SharedDriveDoc.Fields[5].Values = []string{
		"ordered",
		"active",
		"stopped",
		"deleted",
	}
	SharedDriveDoc.Fields[6].Name = "createdTimestamp"
	SharedDriveDoc.Fields[6].Type = "string"
	SharedDriveDoc.Fields[6].Note = ""
	SharedDriveDoc.Fields[6].Description = "ISO 8601 timestamp which holds the date and time when the SharedDrive was created."
	SharedDriveDoc.Fields[6].Comments[encoder.LineComment] = "ISO 8601 timestamp which holds the date and time when the SharedDrive was created."
	SharedDriveDoc.Fields[7].Name = "updatedTimestamp"
	SharedDriveDoc.Fields[7].Type = "string"
	SharedDriveDoc.Fields[7].Note = ""
	SharedDriveDoc.Fields[7].Description = "ISO 8601 timestamp which holds the date and time when the SharedDrive was updated."
	SharedDriveDoc.Fields[7].Comments[encoder.LineComment] = "ISO 8601 timestamp which holds the date and time when the SharedDrive was updated."
	SharedDriveDoc.Fields[8].Name = "sizeMBytes"
	SharedDriveDoc.Fields[8].Type = "int"
	SharedDriveDoc.Fields[8].Note = ""
	SharedDriveDoc.Fields[8].Description = "Size of the drive"
	SharedDriveDoc.Fields[8].Comments[encoder.LineComment] = "Size of the drive"
	SharedDriveDoc.Fields[9].Name = "attachedInstaceArrays"
	SharedDriveDoc.Fields[9].Type = "[]int"
	SharedDriveDoc.Fields[9].Note = ""
	SharedDriveDoc.Fields[9].Description = "An array of the instance array ids attached to this drive"
	SharedDriveDoc.Fields[9].Comments[encoder.LineComment] = "An array of the instance array ids attached to this drive"
	SharedDriveDoc.Fields[10].Name = "operation"
	SharedDriveDoc.Fields[10].Type = "SharedDriveOperation"
	SharedDriveDoc.Fields[10].Note = ""
	SharedDriveDoc.Fields[10].Description = "The operation object"
	SharedDriveDoc.Fields[10].Comments[encoder.LineComment] = "The operation object"
	SharedDriveDoc.Fields[11].Name = "credentials"
	SharedDriveDoc.Fields[11].Type = "SharedDriveCredentials"
	SharedDriveDoc.Fields[11].Note = ""
	SharedDriveDoc.Fields[11].Description = "Credentials of the shared drive"
	SharedDriveDoc.Fields[11].Comments[encoder.LineComment] = "Credentials of the shared drive"
	SharedDriveDoc.Fields[12].Name = "changeID"
	SharedDriveDoc.Fields[12].Type = "int"
	SharedDriveDoc.Fields[12].Note = ""
	SharedDriveDoc.Fields[12].Description = "The operation id"
	SharedDriveDoc.Fields[12].Comments[encoder.LineComment] = "The operation id"
	SharedDriveDoc.Fields[13].Name = "targetsJSON"
	SharedDriveDoc.Fields[13].Type = "string"
	SharedDriveDoc.Fields[13].Note = ""
	SharedDriveDoc.Fields[13].Description = "Details on the ISCSI or FC targets"
	SharedDriveDoc.Fields[13].Comments[encoder.LineComment] = "Details on the ISCSI or FC targets"
	SharedDriveDoc.Fields[14].Name = "ioLimit"
	SharedDriveDoc.Fields[14].Type = "string"
	SharedDriveDoc.Fields[14].Note = ""
	SharedDriveDoc.Fields[14].Description = "Used by certain storage types"
	SharedDriveDoc.Fields[14].Comments[encoder.LineComment] = "Used by certain storage types"
	SharedDriveDoc.Fields[15].Name = "wwn"
	SharedDriveDoc.Fields[15].Type = "string"
	SharedDriveDoc.Fields[15].Note = ""
	SharedDriveDoc.Fields[15].Description = "WWN of the drive as reported by the storage appliance"
	SharedDriveDoc.Fields[15].Comments[encoder.LineComment] = "WWN of the drive as reported by the storage appliance"
	SharedDriveDoc.Fields[16].Name = "storagePoolID"
	SharedDriveDoc.Fields[16].Type = "int"
	SharedDriveDoc.Fields[16].Note = ""
	SharedDriveDoc.Fields[16].Description = "The storage pool id to use if set."
	SharedDriveDoc.Fields[16].Comments[encoder.LineComment] = "The storage pool id to use if set."
	SharedDriveDoc.Fields[17].Name = "affinity"
	SharedDriveDoc.Fields[17].Type = "string"
	SharedDriveDoc.Fields[17].Note = ""
	SharedDriveDoc.Fields[17].Description = "Used to control if drives in this drive array should be allocated on the same storage pool or different storage pools"
	SharedDriveDoc.Fields[17].Comments[encoder.LineComment] = "Used to control if drives in this drive array should be allocated on the same storage pool or different storage pools"
	SharedDriveDoc.Fields[17].Values = []string{
		"same_storage",
		"different_storage",
	}

	SharedDriveCredentialsDoc.Type = "SharedDriveCredentials"
	SharedDriveCredentialsDoc.Comments[encoder.LineComment] = "SharedDriveCredentials iscsi or other forms of connection details"
	SharedDriveCredentialsDoc.Description = "SharedDriveCredentials iscsi or other forms of connection details"
	SharedDriveCredentialsDoc.AppearsIn = []encoder.Appearance{
		{
			TypeName:  "SharedDrive",
			FieldName: "credentials",
		},
	}
	SharedDriveCredentialsDoc.Fields = make([]encoder.Doc, 1)
	SharedDriveCredentialsDoc.Fields[0].Name = "iscsi"
	SharedDriveCredentialsDoc.Fields[0].Type = "ISCSI"
	SharedDriveCredentialsDoc.Fields[0].Note = ""
	SharedDriveCredentialsDoc.Fields[0].Description = "details"
	SharedDriveCredentialsDoc.Fields[0].Comments[encoder.LineComment] = "details"

	SharedDriveOperationDoc.Type = "SharedDriveOperation"
	SharedDriveOperationDoc.Comments[encoder.LineComment] = "SharedDriveOperation represents an ongoing or new operation on a shared drive"
	SharedDriveOperationDoc.Description = "SharedDriveOperation represents an ongoing or new operation on a shared drive"
	SharedDriveOperationDoc.AppearsIn = []encoder.Appearance{
		{
			TypeName:  "SharedDrive",
			FieldName: "operation",
		},
	}
	SharedDriveOperationDoc.Fields = make([]encoder.Doc, 12)
	SharedDriveOperationDoc.Fields[0].Name = "deployStatus"
	SharedDriveOperationDoc.Fields[0].Type = "string"
	SharedDriveOperationDoc.Fields[0].Note = ""
	SharedDriveOperationDoc.Fields[0].Description = "Deploy status"
	SharedDriveOperationDoc.Fields[0].Comments[encoder.LineComment] = "Deploy status"
	SharedDriveOperationDoc.Fields[0].Values = []string{
		"not_started",
		"ongoing",
		"finished",
	}
	SharedDriveOperationDoc.Fields[1].Name = "deployType"
	SharedDriveOperationDoc.Fields[1].Type = "string"
	SharedDriveOperationDoc.Fields[1].Note = ""
	SharedDriveOperationDoc.Fields[1].Description = "Type of operation"
	SharedDriveOperationDoc.Fields[1].Comments[encoder.LineComment] = "Type of operation"
	SharedDriveOperationDoc.Fields[1].Values = []string{
		"create",
		"edit",
		"delete",
		"start",
		"stop",
	}
	SharedDriveOperationDoc.Fields[2].Name = "label"
	SharedDriveOperationDoc.Fields[2].Type = "string"
	SharedDriveOperationDoc.Fields[2].Note = ""
	SharedDriveOperationDoc.Fields[2].Description = "Label"
	SharedDriveOperationDoc.Fields[2].Comments[encoder.LineComment] = "Label"
	SharedDriveOperationDoc.Fields[3].Name = "subdomain"
	SharedDriveOperationDoc.Fields[3].Type = "string"
	SharedDriveOperationDoc.Fields[3].Note = ""
	SharedDriveOperationDoc.Fields[3].Description = "Unique string describing this shared drive"
	SharedDriveOperationDoc.Fields[3].Comments[encoder.LineComment] = "Unique string describing this shared drive"
	SharedDriveOperationDoc.Fields[4].Name = "id"
	SharedDriveOperationDoc.Fields[4].Type = "int"
	SharedDriveOperationDoc.Fields[4].Note = ""
	SharedDriveOperationDoc.Fields[4].Description = "ID of the shared drive"
	SharedDriveOperationDoc.Fields[4].Comments[encoder.LineComment] = "ID of the shared drive"
	SharedDriveOperationDoc.Fields[5].Name = "sizeMBytes"
	SharedDriveOperationDoc.Fields[5].Type = "int"
	SharedDriveOperationDoc.Fields[5].Note = ""
	SharedDriveOperationDoc.Fields[5].Description = "Size of the drive"
	SharedDriveOperationDoc.Fields[5].Comments[encoder.LineComment] = "Size of the drive"
	SharedDriveOperationDoc.Fields[6].Name = "storageType"
	SharedDriveOperationDoc.Fields[6].Type = "string"
	SharedDriveOperationDoc.Fields[6].Note = ""
	SharedDriveOperationDoc.Fields[6].Description = "Type of the shared drive. Readonly."
	SharedDriveOperationDoc.Fields[6].Comments[encoder.LineComment] = "Type of the shared drive. Readonly."
	SharedDriveOperationDoc.Fields[6].Values = []string{
		"iscsi_ssd",
		"iscsi_hdd",
	}
	SharedDriveOperationDoc.Fields[7].Name = "infrastructureID"
	SharedDriveOperationDoc.Fields[7].Type = "int"
	SharedDriveOperationDoc.Fields[7].Note = ""
	SharedDriveOperationDoc.Fields[7].Description = "ID of the infrastructure"
	SharedDriveOperationDoc.Fields[7].Comments[encoder.LineComment] = "ID of the infrastructure"
	SharedDriveOperationDoc.Fields[8].Name = "serviceStatus"
	SharedDriveOperationDoc.Fields[8].Type = "string"
	SharedDriveOperationDoc.Fields[8].Note = ""
	SharedDriveOperationDoc.Fields[8].Description = "Status of the service"
	SharedDriveOperationDoc.Fields[8].Comments[encoder.LineComment] = "Status of the service"
	SharedDriveOperationDoc.Fields[8].Values = []string{
		"active",
		"ordered",
		"deleted",
		"suspended",
		"stopped",
	}
	SharedDriveOperationDoc.Fields[9].Name = "attachedInstanceArrays"
	SharedDriveOperationDoc.Fields[9].Type = "[]int"
	SharedDriveOperationDoc.Fields[9].Note = ""
	SharedDriveOperationDoc.Fields[9].Description = "List of instance arrays to which this shared drive is attached"
	SharedDriveOperationDoc.Fields[9].Comments[encoder.LineComment] = "List of instance arrays to which this shared drive is attached"
	SharedDriveOperationDoc.Fields[10].Name = "changeID"
	SharedDriveOperationDoc.Fields[10].Type = "int"
	SharedDriveOperationDoc.Fields[10].Note = ""
	SharedDriveOperationDoc.Fields[10].Description = "ID of the operation"
	SharedDriveOperationDoc.Fields[10].Comments[encoder.LineComment] = "ID of the operation"
	SharedDriveOperationDoc.Fields[11].Name = "ioLimit"
	SharedDriveOperationDoc.Fields[11].Type = "string"
	SharedDriveOperationDoc.Fields[11].Note = ""
	SharedDriveOperationDoc.Fields[11].Description = "Used for certain SAN appliances"
	SharedDriveOperationDoc.Fields[11].Comments[encoder.LineComment] = "Used for certain SAN appliances"
}

func (SharedDrive) Doc() *encoder.Doc {
	return &SharedDriveDoc
}

func (SharedDriveCredentials) Doc() *encoder.Doc {
	return &SharedDriveCredentialsDoc
}

func (SharedDriveOperation) Doc() *encoder.Doc {
	return &SharedDriveOperationDoc
}

// GetSharedDriveDoc returns documentation for the file ./.
func GetSharedDriveDoc() *encoder.FileDoc {
	return &encoder.FileDoc{
		Name:        "SharedDrive",
		Description: "",
		Structs: []*encoder.Doc{
			&SharedDriveDoc,
			&SharedDriveCredentialsDoc,
			&SharedDriveOperationDoc,
		},
	}
}
