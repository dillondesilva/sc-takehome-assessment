package folder_test

import (
	"testing"
	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
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

func TestGetAllChildFoldersBasic(t *testing.T) {
	/*
	Trivial test using simple combinations of folders
	*/
	type testCase struct {
		name 	string
		orgID	uuid.UUID
		folders	[]folder.Folder
		want	[]folder.Folder
	}
	assert := assert.New(t)

	testCaseWithThreeFolders := testCase{
		"creative-scalphunter",
		uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
		[]folder.Folder{
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
				OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
				Paths: "stunning-horridus",
			},
		},
		[]folder.Folder{
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
		},
	}

	testCaseWithManyFolders := testCase{
		"helped-blackheart",
		uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385"),
		[]folder.Folder{
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
				OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
				Paths: "stunning-horridus",
			},
			folder.Folder{
				Name: "steady-insect",
				OrgId: uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385"),
				Paths: "steady-insect",
			},
			folder.Folder{
				Name: "helped-blackheart",
				OrgId: uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385"),
				Paths: "steady-insect.helped-blackheart",
			},
			folder.Folder{
				Name: "many-silver-sable",
				OrgId: uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385"),
				Paths: "steady-insect.helped-blackheart.many-silver-sable",
			},
			folder.Folder{
				Name: "happy-emoji",
				OrgId: uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385"),
				Paths: "steady-insect.helped-blackheart.happy-emoji",
			},
			folder.Folder{
				Name: "fried-chicken",
				OrgId: uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385"),
				Paths: "steady-insect.helped-blackheart.fried-chicken",
			},
			folder.Folder{
				Name: "noble-vixen",
				OrgId: uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385"),
				Paths: "steady-insect.helped-blackheart.noble-vixen",
			},
			folder.Folder{
				Name: "creative-scalphunter",
				OrgId: uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385"),
				Paths: "steady-insect.helped-blackheart.noble-vixen.creative-scalphunter",
			},
			folder.Folder{
				Name: "exciting-magma",
				OrgId: uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385"),
				Paths: "steady-insect.helped-blackheart.many-silver-sable.stable-karatecha.exciting-magma",
			},
			folder.Folder{
				Name: "noble-vixen",
				OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
				Paths: "noble-vixen",
			},
			folder.Folder{
				Name: "nearby-secret",
				OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
				Paths: "noble-vixen.nearby-secret",
			},
		},
		[]folder.Folder{
			folder.Folder{
				Name: "helped-blackheart",
				OrgId: uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385"),
				Paths: "steady-insect.helped-blackheart",
			},
			folder.Folder{
				Name: "many-silver-sable",
				OrgId: uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385"),
				Paths: "steady-insect.helped-blackheart.many-silver-sable",
			},
			folder.Folder{
				Name: "happy-emoji",
				OrgId: uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385"),
				Paths: "steady-insect.helped-blackheart.happy-emoji",
			},
			folder.Folder{
				Name: "fried-chicken",
				OrgId: uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385"),
				Paths: "steady-insect.helped-blackheart.fried-chicken",
			},
			folder.Folder{
				Name: "noble-vixen",
				OrgId: uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385"),
				Paths: "steady-insect.helped-blackheart.noble-vixen",
			},
			folder.Folder{
				Name: "creative-scalphunter",
				OrgId: uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385"),
				Paths: "steady-insect.helped-blackheart.noble-vixen.creative-scalphunter",
			},
		},
	}

	tests := [...]testCase{testCaseWithThreeFolders, testCaseWithManyFolders}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			folderDriver := folder.NewDriver(tt.folders)
			implementationResult, error := folderDriver.GetAllChildFolders(tt.orgID, tt.name)

			if assert.Nil(error) {
				assert.Equal(tt.want, implementationResult)
			}
		})
	}
}

func TestGetAllChildFoldersWithInvalidOrgId(t *testing.T) {
	/*
	Test that passing an invalid OrgID results in an error being thrown
	*/
	type testCase struct {
		name 	string
		orgID	uuid.UUID
		folders	[]folder.Folder
		want	[]folder.Folder
	}
	assert := assert.New(t)

	testCaseInvalidOrgId := testCase{
		"creative-scalphunter",
		uuid.FromStringOrNil("m8b9879b-f73b-4b0e-b9d9-4fc4c23643a1"),
		[]folder.Folder{
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
		},
		[]folder.Folder{},
	}

	t.Run(testCaseInvalidOrgId.name, func(t *testing.T) {
		folderDriver := folder.NewDriver(testCaseInvalidOrgId.folders)
		implementationResult, error := folderDriver.GetAllChildFolders(testCaseInvalidOrgId.orgID, testCaseInvalidOrgId.name)
		assert.Nil(error)
		assert.Equal(testCaseInvalidOrgId.want, implementationResult)
	})
}

func TestGetAllChildFoldersWithInvalidPath(t *testing.T) {
	/*
	Test that passing an invalid path results in an error being thrown
	*/
	type testCase struct {
		name 	string
		orgID	uuid.UUID
		folders	[]folder.Folder
		want	[]folder.Folder
	}
	assert := assert.New(t)

	testCaseInvalidPath := testCase{
		"creative-scalphunter.moonlight",
		uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
		[]folder.Folder{
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
		},
		[]folder.Folder{},
	}

	t.Run(testCaseInvalidPath.name, func(t *testing.T) {
		folderDriver := folder.NewDriver(testCaseInvalidPath.folders)
		implementationResult, error := folderDriver.GetAllChildFolders(testCaseInvalidPath.orgID, testCaseInvalidPath.name)
		assert.NotNil(error)
		assert.Equal(testCaseInvalidPath.want, implementationResult)
	})
}

func TestGetAllChildFoldersWithDuplicatePathsAcrossOrgIds(t *testing.T) {
	/*
	Test that passing a duplicate path across OrgIds results in 
	the correct folders being returned
	*/
	type testCase struct {
		name 	string
		orgID	uuid.UUID
		folders	[]folder.Folder
		want	[]folder.Folder
	}
	assert := assert.New(t)

	testCaseDuplicatePaths := testCase{
		"creative-scalphunter",
		uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
		[]folder.Folder{
			folder.Folder{
				Name: "creative-scalphunter",
				OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
				Paths: "creative-scalphunter",
			},
			folder.Folder{
				Name: "creative-scalphunter",
				OrgId: uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385"),
				Paths: "creative-scalphunter",
			},
			folder.Folder{
				Name: "cool-car",
				OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
				Paths: "creative-scalphunter.cool-car",
			},
			folder.Folder{
				Name: "random-insect",
				OrgId: uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385"),
				Paths: "creative-scalphunter.random-insect",
			},
		},
		[]folder.Folder{
			folder.Folder{
				Name: "creative-scalphunter",
				OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
				Paths: "creative-scalphunter",
			},
			folder.Folder{
				Name: "cool-car",
				OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
				Paths: "creative-scalphunter.cool-car",
			},
		},
	}

	t.Run(testCaseDuplicatePaths.name, func(t *testing.T) {
		folderDriver := folder.NewDriver(testCaseDuplicatePaths.folders)
		implementationResult, error := folderDriver.GetAllChildFolders(testCaseDuplicatePaths.orgID, testCaseDuplicatePaths.name)
		assert.Nil(error)
		assert.Equal(testCaseDuplicatePaths.want, implementationResult)
	})
}

func TestGetAllChildFoldersWithEmptyPath(t *testing.T) {
	/*
	Test that calling GetAllChildFolders with an empty
	folder name returns a corresponding error
	*/
	type testCase struct {
		name 	string
		orgID	uuid.UUID
		folders	[]folder.Folder
		want	[]folder.Folder
	}
	assert := assert.New(t)

	testCaseDuplicatePaths := testCase{
		"",
		uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
		[]folder.Folder{
			folder.Folder{
				Name: "creative-scalphunter",
				OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
				Paths: "creative-scalphunter",
			},
			folder.Folder{
				Name: "creative-scalphunter",
				OrgId: uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385"),
				Paths: "creative-scalphunter",
			},
			folder.Folder{
				Name: "random-insect",
				OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
				Paths: "creative-scalphunter.random-insect",
			},
			folder.Folder{
				Name: "random-insect",
				OrgId: uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385"),
				Paths: "creative-scalphunter.random-insect",
			},
		},
		[]folder.Folder{},
	}

	t.Run(testCaseDuplicatePaths.name, func(t *testing.T) {
		folderDriver := folder.NewDriver(testCaseDuplicatePaths.folders)
		implementationResult, error := folderDriver.GetAllChildFolders(testCaseDuplicatePaths.orgID, testCaseDuplicatePaths.name)
		assert.NotNil(error)
		assert.Equal(testCaseDuplicatePaths.want, implementationResult)
	})
}

func TestGetAllChildFoldersWithBrokenPathsInChild(t *testing.T) {
	/*
	Test that when a child folder contains an incorrect path, an error is returned
	*/
	t.Parallel()
	type testCase struct {
		name 	string
		orgID	uuid.UUID
		folders	[]folder.Folder
		want	[]folder.Folder
	}
	assert := assert.New(t)
	
	testCaseWithSeenPath := testCase{
		"steady-insect",
		uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
		[]folder.Folder{
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
				OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
				Paths: "stunning-horridus",
			},
			folder.Folder{
				Name: "steady-insect",
				OrgId: uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385"),
				Paths: "steady-insect",
			},
			folder.Folder{
				Name: "helped-blackheart",
				OrgId: uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385"),
				Paths: "steady-insect.helped-blackheart",
			},
			folder.Folder{
				Name: "many-silver-sable",
				OrgId: uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385"),
				Paths: "steady-insect.helped-blackheart.many-silver-sable",
			},
			folder.Folder{
				Name: "formula-one",
				OrgId: uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385"),
				Paths: "steady-insect.helped-blackheart.many-silver-sable",
			},
			folder.Folder{
				Name: "exciting-magma",
				OrgId: uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385"),
				Paths: "steady-insect.helped-blackheart.many-silver-sable.stable-karatecha.exciting-magma",
			},
			folder.Folder{
				Name: "noble-vixen",
				OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
				Paths: "noble-vixen",
			},
			folder.Folder{
				Name: "nearby-secret",
				OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
				Paths: "noble-vixen.nearby-secret",
			},
		},
		[]folder.Folder{},
	}

	testCaseWithFolderNameNotInPath := testCase{
		"steady-insect",
		uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
		[]folder.Folder{
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
				OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
				Paths: "stunning-horridus",
			},
			folder.Folder{
				Name: "steady-insect",
				OrgId: uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385"),
				Paths: "steady-insect",
			},
			folder.Folder{
				Name: "helped-blackheart",
				OrgId: uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385"),
				Paths: "steady-insect.helped-blackheart",
			},
			folder.Folder{
				Name: "many-silver-sable",
				OrgId: uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385"),
				Paths: "steady-insect.helped-blackheart.many-silver-sable",
			},
			folder.Folder{
				Name: "formula-one",
				OrgId: uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385"),
				Paths: "steady-insect.helped-blackheart",
			},
			folder.Folder{
				Name: "exciting-magma",
				OrgId: uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385"),
				Paths: "steady-insect.helped-blackheart.many-silver-sable.stable-karatecha.exciting-magma",
			},
			folder.Folder{
				Name: "noble-vixen",
				OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
				Paths: "noble-vixen",
			},
			folder.Folder{
				Name: "nearby-secret",
				OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
				Paths: "noble-vixen.nearby-secret",
			},
		},
		[]folder.Folder{},
	}

	testCaseWithNonExistentPaths := testCase{
		"steady-insect",
		uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
		[]folder.Folder{
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
				OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
				Paths: "stunning-horridus",
			},
			folder.Folder{
				Name: "steady-insect",
				OrgId: uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385"),
				Paths: "steady-insect",
			},
			folder.Folder{
				Name: "helped-blackheart",
				OrgId: uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385"),
				Paths: "steady-insect.helped-blackheart",
			},
			folder.Folder{
				Name: "many-silver-sable",
				OrgId: uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385"),
				Paths: "steady-insect.helped-blackheart.many-silver-sable",
			},
			folder.Folder{
				Name: "formula-one",
				OrgId: uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385"),
				Paths: "steady-insect.helped-blackheart.non-existent-path",
			},
			folder.Folder{
				Name: "exciting-magma",
				OrgId: uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385"),
				Paths: "steady-insect.helped-blackheart.many-silver-sable.stable-karatecha.exciting-magma",
			},
			folder.Folder{
				Name: "noble-vixen",
				OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
				Paths: "noble-vixen",
			},
			folder.Folder{
				Name: "nearby-secret",
				OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
				Paths: "noble-vixen.nearby-secret",
			},
		},
		[]folder.Folder{},
	}

	testCaseWithPathInOtherOrgId := testCase{
		"steady-insect",
		uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
		[]folder.Folder{
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
				OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
				Paths: "stunning-horridus",
			},
			folder.Folder{
				Name: "steady-insect",
				OrgId: uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385"),
				Paths: "steady-insect",
			},
			folder.Folder{
				Name: "helped-blackheart",
				OrgId: uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385"),
				Paths: "steady-insect.helped-blackheart",
			},
			folder.Folder{
				Name: "many-silver-sable",
				OrgId: uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385"),
				Paths: "steady-insect.helped-blackheart.many-silver-sable",
			},
			folder.Folder{
				Name: "formula-one",
				OrgId: uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385"),
				Paths: "creative-scalphunter.clear-arclight",
			},
			folder.Folder{
				Name: "exciting-magma",
				OrgId: uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385"),
				Paths: "steady-insect.helped-blackheart.many-silver-sable.stable-karatecha.exciting-magma",
			},
			folder.Folder{
				Name: "noble-vixen",
				OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
				Paths: "noble-vixen",
			},
			folder.Folder{
				Name: "nearby-secret",
				OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
				Paths: "noble-vixen.nearby-secret",
			},
		},
		[]folder.Folder{},
	}

	testCaseWithPathJumpingDirectory := testCase{
		"steady-insect",
		uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
		[]folder.Folder{
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
				OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
				Paths: "stunning-horridus",
			},
			folder.Folder{
				Name: "steady-insect",
				OrgId: uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385"),
				Paths: "steady-insect",
			},
			folder.Folder{
				Name: "helped-blackheart",
				OrgId: uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385"),
				Paths: "steady-insect.helped-blackheart",
			},
			folder.Folder{
				Name: "many-silver-sable",
				OrgId: uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385"),
				Paths: "steady-insect.helped-blackheart.many-silver-sable",
			},
			folder.Folder{
				Name: "formula-one",
				OrgId: uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385"),
				Paths: "steady-insect.helped-blackheart.many-silver-sable.stable-karatecha.formula-one",
			},
			folder.Folder{
				Name: "exciting-magma",
				OrgId: uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385"),
				Paths: "steady-insect.helped-blackheart.many-silver-sable.stable-karatecha.exciting-magma",
			},
			folder.Folder{
				Name: "noble-vixen",
				OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
				Paths: "noble-vixen",
			},
			folder.Folder{
				Name: "nearby-secret",
				OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
				Paths: "noble-vixen.nearby-secret",
			},
		},
		[]folder.Folder{},
	}

	tests := [...]testCase{
		testCaseWithSeenPath,
		testCaseWithFolderNameNotInPath,
		testCaseWithNonExistentPaths,
		testCaseWithPathInOtherOrgId,
		testCaseWithPathJumpingDirectory,
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			folderDriver := folder.NewDriver(tt.folders)
			implementationResult, error := folderDriver.GetAllChildFolders(tt.orgID, tt.name)
			assert.NotNil(error)
			assert.Equal(tt.want, implementationResult)
		})
	}
}
