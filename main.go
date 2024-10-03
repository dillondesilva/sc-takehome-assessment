package main

import (
	"fmt"
	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
)

func main() {
	// orgID := uuid.FromStringOrNil(folder.DefaultOrgID)

	// res := folder.GetAllFolders()
	testFolders := []folder.Folder{
		folder.Folder{
			Name: "creative-scalphunter",
			OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
			Paths: "creative-scalphunter",
		},
		folder.Folder{
			Name: "clear-arclight",
			OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
			Paths: "creative-scalphunter.clear-arclight",
		},
		folder.Folder{
			Name: "stunning-horridus",
			OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
			Paths: "stunning-horridus",
		},
	}

	// // example usage
	// folderDriver := folder.NewDriver(res)
	folderDriver := folder.NewDriver(testFolders)
	// orgFolder := folderDriver.GetFoldersByOrgID(orgID)
	movedFolders, err := folderDriver.MoveFolder("creative-scalphunter", "stunning-horridus")
	if (err == nil) {
		fmt.Println("MoveFolder returned", movedFolders)
	} else {
		fmt.Println(err)
	}

	// folder.PrettyPrint(res)
	// fmt.Printf("\n Folders for orgID: %s", orgID)
	// folder.PrettyPrint(orgFolder)
}
