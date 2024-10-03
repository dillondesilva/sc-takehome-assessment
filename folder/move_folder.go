package folder

import (
	// "github.com/gofrs/uuid"
	"slices"
	// "strings"
	"errors"
)

func IsFolderInCollection(givenFolders []Folder, folderToCheck Folder) bool {
	/*
	Check if a folder exists in a given slice of folders
	*/
	for _, folder := range givenFolders {
		if folder.Name == folderToCheck.Name && folder.OrgId == folderToCheck.OrgId {
			return true
		}
	}

	return false
}

func (f* driver) GetFolderByName(folderName string) (Folder, error) {
	/*
	Obtain folder data under a given name
	*/
	
	idxOfFolderData := slices.IndexFunc(f.folders, func(folderData Folder) bool {
		if folderData.Name == folderName {
			return true
		}

		return false
	})
	if idxOfFolderData == -1 {
		return Folder{}, errors.New("Invalid folder name")
	}
	folderData := f.folders[idxOfFolderData]
	return folderData, nil
}

func (f* driver) IsFolderUniqueToOrgId(folderToCheck Folder) bool {
	// Check that for a given folder name, it only belongs to one organisation
	for _, folder := range f.folders {
		if folder.Name == folderToCheck.Name && folder.OrgId != folderToCheck.OrgId { return false }
	}

	return true
}

func CheckFoldersInSameOrg(folderA Folder, folderB Folder) bool {
	if folderA.OrgId == folderB.OrgId { return true }
	return false
}

func (f* driver) CheckFoldersAreUniqueAndSameOrg(folderA Folder, folderB Folder) bool {
	if CheckFoldersInSameOrg(folderA, folderB) && f.IsFolderUniqueToOrgId(folderA) && folderA.Name != folderB.Name { 
		return true 
	}
	return false
}

func (f *driver) MoveFolder(sourceFolder string, destFolder string) ([]Folder, error) {
	/*
	Approach: Create an empty slice to represent new structure. Validate
	both folders exist in the same OrgId (and that folders are unique to OrgId)
	Then get all children of source, copy into the array for the new structure.
	Then copy all the existing code into the  
	*/
	updatedFolders := []Folder{}

	// TODO: Validate same OrgID and folder existence
	sourceFolderData, getSourceErr := f.GetFolderByName(sourceFolder)
	destFolderData, getDestErr := f.GetFolderByName(destFolder)
	if getSourceErr != nil || getDestErr != nil {
		return updatedFolders, errors.New("Unable to get source or dest folder")
	}

	if f.CheckFoldersAreUniqueAndSameOrg(sourceFolderData, destFolderData) != true {
		return updatedFolders, errors.New("Source/dest Org do not match or are not unique")
	}

	// Get child folders of source, move them by updating paths and then add to slice
	sourceAndChildFolders, getChildErr := f.GetAllChildFolders(sourceFolderData.OrgId, sourceFolderData.Name)
	if getChildErr != nil {
		return updatedFolders, errors.New("Invalid child folder inside of source detected")
	}

	sourceAndChildFolders = append(sourceAndChildFolders, sourceFolderData)

	for _, folderToMove := range sourceAndChildFolders {
		// Possible edge case not tested yet: If the dest folder path is invalid
		folderToMove.Paths = destFolderData.Paths + "." + folderToMove.Paths
		updatedFolders = append(updatedFolders, folderToMove)
	}

	// Add all other folders, ensuring they aren't already in the slice
	// from our earlier operation when moving folders.
	for _, folderToMove := range f.folders {
		if IsFolderInCollection(sourceAndChildFolders, folderToMove) == false {
			updatedFolders = append(updatedFolders, folderToMove)
		}
	}

	return updatedFolders, nil
}
