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
			Name: "alpha",
			Paths: "alpha",
			OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
		},
		folder.Folder{
			Name: "bravo",
			Paths: "alpha.bravo",
			OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
		},
		folder.Folder{
			Name: "charlie",
			Paths: "alpha.bravo.charlie",
			OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
		},
		folder.Folder{
			Name: "delta",
			Paths: "alpha.delta",
			OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
		},
		folder.Folder{
			Name: "echo",
			Paths: "alpha.delta.echo",
			OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
		},
		folder.Folder{
			Name: "foxtrot",
			Paths: "foxtrot",
			OrgId: uuid.FromStringOrNil("q1556e17-b7c0-45a3-a6ae-9546248fb17d"),
		},
		folder.Folder{
			Name: "golf",
			Paths: "golf",
			OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
		},
	}

	// // example usage
	// folderDriver := folder.NewDriver(res)
	folderDriver := folder.NewDriver(testFolders)
	// orgFolder := folderDriver.GetFoldersByOrgID(orgID)

	// ADJUST THIS TO EXPERIMENT
	childFolders, errChild := folderDriver.GetAllChildFolders(uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), "invalid_folder")
	if (errChild == nil) {
		fmt.Println("POST GET ALL CHILD FOLDERS")
		for _, folder := range childFolders {
			fmt.Println(folder)
		}
	} else {
		fmt.Println(errChild)
	}

	// ADJUST THIS TO EXPERIMENT
	movedFolders, err := folderDriver.MoveFolder("bravo", "golf")
	if (err == nil) {
		fmt.Println("POST MOVE FOLDER STRUCTURE")
		for _, folder := range movedFolders {
			fmt.Println(folder)
		}
	} else {
		fmt.Println(err)
	}

	// folder.PrettyPrint(res)
	// fmt.Printf("\n Folders for orgID: %s", orgID)
	// folder.PrettyPrint(orgFolder)
}
