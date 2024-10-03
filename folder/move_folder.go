package folder

import (
	"slices"
	"errors"
)

func IsFolderInCollection(givenFolders []Folder, folderToCheck Folder) bool {
	/*
	Check if a folder exists in a given collection of folders
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
	Return the full Folder struct data found in the driver's
	folder collection for a given folder name (if it exists)
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
	/*
	Return true if for a given folder name, it only belongs to one OrgId
	(i.e. if "alpha" is in Org1 and Org2, then this should return false)
	*/
	for _, folder := range f.folders {
		if folder.Name == folderToCheck.Name && folder.OrgId != folderToCheck.OrgId { return false }
	}

	return true
}

func CheckFoldersInSameOrg(folderA Folder, folderB Folder) bool {
	/*
	Return true if given folders are from same OrgId
	*/
	if folderA.OrgId == folderB.OrgId { return true }
	return false
}

func (f* driver) CheckFoldersAreUniqueAndSameOrg(folderA Folder, folderB Folder) bool {
	/*
	Return true if the given folders are from the same organisation and

	*/
	if CheckFoldersInSameOrg(folderA, folderB) && f.IsFolderUniqueToOrgId(folderA) && folderA.Name != folderB.Name { 
		return true 
	}
	return false
}

func (f *driver) MoveFolder(sourceFolder string, destFolder string) ([]Folder, error) {
	/*
	Move contents inside of source folder into the specified destination folder
	*/

	// Creating an empty slice for new folder structure and obtaining struct 
	// data from the source and destination folder names
	updatedFolders := []Folder{}
	sourceFolderData, getSourceErr := f.GetFolderByName(sourceFolder)
	destFolderData, getDestErr := f.GetFolderByName(destFolder)
	if getSourceErr != nil || getDestErr != nil {
		return updatedFolders, errors.New("Unable to get source or dest folder")
	}

	// Validate folders are unique and from same organisation
	if f.CheckFoldersAreUniqueAndSameOrg(sourceFolderData, destFolderData) != true {
		return updatedFolders, errors.New("Source/dest Org do not match or are not unique")
	}

	// Get child folders of source, move them by updating paths and then add to slice
	sourceAndChildFolders, getChildErr := f.GetAllChildFolders(sourceFolderData.OrgId, 
		sourceFolderData.Name)
	if getChildErr != nil {
		return updatedFolders, errors.New("Invalid child folder inside of source detected")
	}

	// Start moving the source folder inside destination folder
	sourceAndChildFolders = append(sourceAndChildFolders, sourceFolderData)
	for _, folderToMove := range sourceAndChildFolders {
		folderToMove.Paths = destFolderData.Paths + "." + folderToMove.Paths
		updatedFolders = append(updatedFolders, folderToMove)
	}

	// Add other folders, ensuring they aren't already in the slice
	// from our earlier operation when moving folders.
	for _, folderToMove := range f.folders {
		if IsFolderInCollection(sourceAndChildFolders, folderToMove) == false {
			updatedFolders = append(updatedFolders, folderToMove)
		}
	}

	return updatedFolders, nil
}
