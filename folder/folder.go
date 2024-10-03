package folder

import "github.com/gofrs/uuid"

type IDriver interface {
	// GetFoldersByOrgID returns all folders that belong to a specific orgID.
	GetFoldersByOrgID(orgID uuid.UUID) []Folder

	// GetFolder returns complete folder data under the first match of
	// a given name
	GetFolderByName(folderName string) (Folder, error)

	// IsFolderUniqueToOrgId checks whether a given folder name belongs to
	// only one organisation
	IsFolderUniqueToOrgId(folderToCheck Folder) bool
	
	// CheckFoldersAreUniqueAndSameOrg checks whether two folders belong
	// to the same OrgId, are different and also ensures there are no
	// other OrgIds with the same folder names
	CheckFoldersAreUniqueAndSameOrg(folderA Folder, folderB Folder) bool

	// component 1
	// Implement the following methods:
	// GetAllChildFolders returns all child folders of a specific folder.
	GetAllChildFolders(orgID uuid.UUID, name string) ([]Folder, error)

	// component 2
	// Implement the following methods:
	// MoveFolder moves a folder to a new destination.
	MoveFolder(name string, dst string) ([]Folder, error)

}

type driver struct {
	// define attributes here
	// data structure to store folders
	// or preprocessed data

	// example: feel free to change the data structure, if slice is not what you want
	folders []Folder
}

func NewDriver(folders []Folder) IDriver {
	return &driver{
		// initialize attributes here
		folders: folders,
	}
}
