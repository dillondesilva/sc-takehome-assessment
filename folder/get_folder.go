package folder

import (
	"github.com/gofrs/uuid"
	"slices"
	"strings"
	"errors"
)

func GetAllFolders() []Folder {
	/*
	Helper function from specification to get data from sample.json 
	*/
	return GetSampleData()
}

func (f *driver) GetFoldersByOrgID(orgID uuid.UUID) []Folder {
	/*
	Helper function from specification to get all folders in an OrgId
	*/
	folders := f.folders

	res := []Folder{}
	for _, f := range folders {
		if f.OrgId == orgID {
			res = append(res, f)
		}
	}

	return res
}

func GetAdjacentSubfolders(givenFolders []Folder, rootFolder Folder) []Folder {
	/*
	Return a slice of Folder objects which are one level down from a given
	root folder (e.g. alpha.bravo is an adjacent subfolder to alpha whereas
	alpha.bravo.charlie is not as it is not adjacent)
	*/
	rootPath := rootFolder.Paths
	rootFolderDepth := len(strings.Split(rootPath, "."))
	adjacentSubfolders := []Folder{}

	for _, folder := range givenFolders {
		if strings.HasPrefix(folder.Paths, rootPath) {
			folderDepth := len(strings.Split(folder.Paths, "."))
			if rootFolderDepth + 1 == folderDepth {
				adjacentSubfolders = append(adjacentSubfolders, folder)
			}
		}
	}

	return adjacentSubfolders
}

func IsPathSeen(seenPaths []Folder, newPath string) bool {
	/*
	Return true if the given path appears as a path in a collection
	of folders that have already been seen. This is used by
	GetAllChildFolders
	*/
	for _, seenPath := range seenPaths {
		if seenPath.Paths == newPath {
			return true
		}
	}

	return false
}

func IsValidFolderName(folderToValidate Folder, parentFolder Folder) bool {
	/*
	Return true if the given folder name has a path that is 
	appropriately constructed with respect to the parent folder. 
	This is used by GetAllChildFolders
	*/
	expectedFolderPath := parentFolder.Paths + "." + folderToValidate.Name
	if expectedFolderPath == folderToValidate.Paths {
		return true
	} else if parentFolder == folderToValidate {
		return true
	}

	return false
}

func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) ([]Folder, error) {
	/*
	Performs a DFS/BFS-style traversal to return all child folders for a 
	given root-folder name. This is to ensure that the edge case 
	of a broken folder structure is handled appropriately 
	(see TestGetAllChildFoldersWithSkippedFoldersInPath)
	*/

	// Filter to only work with folders in the Org
	foldersInOrgID := f.GetFoldersByOrgID(orgID)
	if len(foldersInOrgID) == 0 {
		return []Folder{}, errors.New("Invalid or empty OrgID")
	}

	// Obtaining Folder struct data for the given folder name
	idxOfRootFolderData := slices.IndexFunc(foldersInOrgID, func(folderData Folder) bool {
		if folderData.Name == name {
			return true
		}

		return false
	})
	if idxOfRootFolderData == -1 {
		return []Folder{}, errors.New("Invalid path parameter")
	}
	rootFolderData := foldersInOrgID[idxOfRootFolderData]

	// Setting up parameters for traversal to ensure child folders
	// are stored and any folders seen are maintained (to avoid
	// infinite looping in traversal and handle edge cases) 
	seenFoldersInTraversal := []Folder{}
	unseenFoldersInTraversal := []Folder{ rootFolderData }
	allChildFolders := []Folder{}

	for len(unseenFoldersInTraversal) > 0 {
		currParentFolderInTraversal := unseenFoldersInTraversal[0]
		unseenFoldersInTraversal = unseenFoldersInTraversal[1:]
		adjSubfoldersToCurrParent := GetChildFoldersOneLevelDown(foldersInOrgID, 
			currParentFolderInTraversal)
		
		// Iterate through all adjacent subfolders, pushing them to queue
		// for future processing
		for _, adjSubfolder := range adjSubfoldersToCurrParent {
			if IsValidFolderName(adjSubfolder, currParentFolderInTraversal) {
				unseenFoldersInTraversal = append(unseenFoldersInTraversal, adjSubfolder)
			} else {
				// Child folder is not valid or broken - throw an error
				return []Folder{}, errors.New("Broken folder structure: Invalid folder name for a child during traversal was found.")
			}
		}

		// Append parent folder that was just processed to childFolders
		// unless it is the root folder which we started our traversal
		// from (since we only want child folders)
		if currParentFolderInTraversal != rootFolderData {
			allChildFolders = append(allChildFolders, currParentFolderInTraversal)
		}

		seenFoldersInTraversal = append(seenFoldersInTraversal, currParentFolderInTraversal)
	}

	return allChildFolders, nil
}
