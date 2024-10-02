package folder

import (
	"github.com/gofrs/uuid"
	"slices"
	"strings"
	"errors"
)

func GetAllFolders() []Folder {
	return GetSampleData()
}

func (f *driver) GetFoldersByOrgID(orgID uuid.UUID) []Folder {
	folders := f.folders

	res := []Folder{}
	for _, f := range folders {
		if f.OrgId == orgID {
			res = append(res, f)
		}
	}

	return res
}

func GetChildFoldersOneLevelDown(givenFolders []Folder, rootFolder Folder) []Folder {
	rootPath := rootFolder.Paths
	rootFolderDepth := len(strings.Split(rootPath, "."))
	childFoldersOneLevelDown := []Folder{}

	for _, folder := range givenFolders {
		if strings.HasPrefix(folder.Paths, rootPath) {
			folderDepth := len(strings.Split(folder.Paths, "."))
			if rootFolderDepth + 1 == folderDepth {
				childFoldersOneLevelDown = append(childFoldersOneLevelDown, folder)
			}
		}
	}

	return childFoldersOneLevelDown
}

func IsPathSeen(seenPaths []Folder, newPath string) bool {
	for _, seenPath := range seenPaths {
		if seenPath.Paths == newPath {
			return true
		}
	}

	return false
}

func IsValidFolderName(folderToValidate Folder, parentFolder Folder) bool {
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
	Approach: First get all the folders for the given OrgID by calling
	GetFoldersByOrgID. Then create an empty array for unseen folders,
	seen paths and the child folders. Also create a variable to store
	the traversal's current parent folder which is initially
	set to whatever the given name of a folder's Paths value is.
	Perform a DFS, adding to the child folders array and return
	when there are no more unseen folders.
	*/

	// TODO: Validation of function parameters

	foldersInOrgID := f.GetFoldersByOrgID(orgID)
	if len(foldersInOrgID) == 0 {
		return []Folder{}, errors.New("Invalid or empty OrgID")
	}

	idxOfFolderDataFromParams := slices.IndexFunc(foldersInOrgID, func(folderData Folder) bool {
		if folderData.Name == name {
			return true
		}

		return false
	})
	if idxOfFolderDataFromParams == -1 {
		return []Folder{}, errors.New("Invalid path parameter")
	}
	folderDataFromParams := foldersInOrgID[idxOfFolderDataFromParams]

	seenFoldersInTraversal := []Folder{}
	unseenFoldersInTraversal := []Folder{ folderDataFromParams }
	allChildFolders := []Folder{}

	for len(unseenFoldersInTraversal) > 0 {
		// TODO: Validate folder path ends in folder name and otherwise throw
		// an error

		currParentFolderInTraversal := unseenFoldersInTraversal[0]
		unseenFoldersInTraversal = unseenFoldersInTraversal[1:]

		adjSubfoldersToCurrParent := GetChildFoldersOneLevelDown(foldersInOrgID, currParentFolderInTraversal)

		for _, adjSubfolder := range adjSubfoldersToCurrParent {
			if IsValidFolderName(adjSubfolder, currParentFolderInTraversal) {
				unseenFoldersInTraversal = append(unseenFoldersInTraversal, adjSubfolder)
			} else {
				return []Folder{}, errors.New("Broken folder structure: Invalid folder name for a child during traversal was found.")
			}
		}

		allChildFolders = append(allChildFolders, currParentFolderInTraversal)
		seenFoldersInTraversal = append(seenFoldersInTraversal, currParentFolderInTraversal)

	}

	return allChildFolders, nil
}
