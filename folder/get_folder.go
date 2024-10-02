package folder

import (
	"github.com/gofrs/uuid"
	"fmt"
	"slices"
	"strings"
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
	// Get subfolders
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

func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) []Folder {
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
	idxOfFolderDataFromParams := slices.IndexFunc(foldersInOrgID, func(folderData Folder) bool {
		if folderData.Name == name {
			return true
		}

		return false
	})
	folderDataFromParams := foldersInOrgID[idxOfFolderDataFromParams]
	fmt.Println(folderDataFromParams)

	fmt.Println("Getting child folders one level down from the root")
	childFoldersOneLevelDown := GetChildFoldersOneLevelDown(foldersInOrgID, folderDataFromParams)
	fmt.Printf("%v\n", childFoldersOneLevelDown)

	// seenPathsDuringTraversal := []Folder{}
	// unseenPathsDuringTraversal := []Folder{ folderDataFromParams }
	// allChildFolders := []Folder{}
	// currentParentPathDuringTraversal := folderDataFromParams.Paths

	// for len(unseenFoldersDuringTraversal) > 0 {
	// 	// TODO: Validate folder path ends in folder name and otherwise throw
	// 	// an error

	// 	// Get neighbouring folders
	// 	// (i.e. folders with same currentParentPathDuringTraversal but are only 
	// 	// one level down)
	// 	// a

	// 	// Filter OrgIds by whatever the currentParentPathDuringTraversal is
	// }
	return []Folder{}
}
