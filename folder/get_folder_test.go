package folder_test

import (
	"testing"
	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	// "github.com/stretchr/testify/assert"
)

// feel free to change how the unit test is structured
func Test_folder_GetFoldersByOrgID(t *testing.T) {
	t.Parallel()
	tests := [...]struct {
		name    string
		orgID   uuid.UUID
		folders []folder.Folder
		want    []folder.Folder
	}{
		// TODO: your tests here
		{
			"bravo",
			uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
			[]folder.Folder{
				folder.Folder{
					Name: "bravo",
					OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths: "alpha.bravo",
				},
				folder.Folder{
					Name: "charlie",
					OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17v"),
					Paths: "alpha.charlie",
				},
			},
			[]folder.Folder{
				{
					Name: "bravo",
					OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths: "alpha.bravo",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			get := f.GetFoldersByOrgID(tt.orgID)
			t.Logf("\n Folders for orgID: %s", get)
		})
	}
}

func Test_folder_GetAllChildFolders_Basic1(t *testing.T) {
	// TODO: Implement a trivial test with 3
	// OrgIDs present to ensure GetAllChildFolders
}

func Test_folder_GetAllChildFolders_Basic2(t *testing.T) {
	// TODO: Implement a trivial test with 5-6
	// OrgIDs and >8 different folders to ensure GetAllChildFolders
	// works over a basic but longer example (compared to basic1)
}

func Test_folder_GetAllChildFolders_Invalid_OrgID(t *testing.T) {
	// TODO: Test that passing an invalid OrgID results in a 
	// corresponding error being thrown
}

func Test_folder_GetAllChildFolders_Invalid_Path(t *testing.T) {
	// TODO: Test that passing an invalid path results in a 
	// corresponding error being thrown
}

func Test_folder_GetAllChildFolders_SamePath_DiffOrg(t *testing.T) {
	// TODO: Test that passing a path which exists in multiple
	// organisations doesn't result in results being contaminated
	// across invalid orgs.
}

func Test_folder_GetAllChildFolders_QueryEmptyPath(t *testing.T) {
	// TODO: Test that calling GetAllChildFolders with an empty
	// folder name returns a corresponding error
}

func Test_folder_GetAllChildFolders_QueryEmptyOrgID(t *testing.T) {
	// TODO: Test that calling GetAllChildFolders with an empty
	// OrgID returns a corresponding error
}

func Test_folder_GetAllChildFolders_Invalid_ChildPath1(t *testing.T) {
	// TODO: Test that when a child folder incorrectly contains
	// a path pointing to an ancestral child (i.e. One that has already been seen), 
	// the error is handled accordingly. This is to prevent any sort of
	// infinite recursion from happening.
}

func Test_folder_GetAllChildFolders_Invalid_ChildPath2(t *testing.T) {
	// TODO: Test that when a child folder incorrectly contains
	// a path to a non-existent folder, the error is handled
	// accordingly.
}

func Test_folder_GetAllChildFolders_Invalid_ChildPath3(t *testing.T) {
	// TODO: Test that when a child folder contains an invalid path,
	// the error is handled accordingly. An example could be the folder
	// bravo which has the path alpha.charlie (opposed to alpha.bravo)
}

func Test_folder_GetAllChildFolders_Volume1(t *testing.T) {
	// TODO: Test that when a large number of folders are in use,
	// the function still returns the correct results.
}
