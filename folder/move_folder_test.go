package folder_test

import (
	"testing"
)

func TestMoveFolderBasic(t *testing.T) {
	/*
	Trivial test with simple folder structures
	*/
	type testCase struct {
		sourceFolder string
		destFolder	 string
		folders		 []folder.Folder
		want		 []folder.Folder
	}
	assert := assert.New(t)

	testCaseWithThreeFolders := testCase{
		"creative-scalphunter",
		"stunning-horridus"
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
				OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
				Paths: "stunning-horridus",
			},
		},
		[]folder.Folder{
			folder.Folder{
				Name: "creative-scalphunter",
				OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
				Paths: "stunning-horridus.creative-scalphunter",
			},
			folder.Folder{
				Name: "clear-arclight",
				OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
				Paths: "stunning-horridus.creative-scalphunter.clear-arclight",
			},
			folder.Folder{
				Name: "stunning-horridus",
				OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
				Paths: "stunning-horridus",
			},
		},
	}

	/* 
	Note: This test case has folder "exciting-magma" but its parent "stable-karatecha"
	is not in structure. This is symbolic of a broken portion of the given folders
	structure which should not be reached by the system. Also tested later.
	*/
	testCaseWithManyFolders := testCase{
		"helped-blackheart",
		"creative-scalphunter",
		[]folder.Folder{
			folder.Folder{
				Name: "creative-scalphunter",
				OrgId: uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385"),
				Paths: "creative-scalphunter",
			},
			folder.Folder{
				Name: "clear-arclight",
				OrgId: uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385"),
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
				Name: "iphone-11",
				OrgId: uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385"),
				Paths: "steady-insect.helped-blackheart.noble-vixen.iphone-11",
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
				Name: "creative-scalphunter",
				OrgId: uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385"),
				Paths: "creative-scalphunter",
			},
			folder.Folder{
				Name: "clear-arclight",
				OrgId: uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385"),
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
				Paths: "creative-scalphunter.steady-insect.helped-blackheart",
			},
			folder.Folder{
				Name: "many-silver-sable",
				OrgId: uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385"),
				Paths: "creative-scalphunter.steady-insect.helped-blackheart.many-silver-sable",
			},
			folder.Folder{
				Name: "happy-emoji",
				OrgId: uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385"),
				Paths: "creative-scalphunter.steady-insect.helped-blackheart.happy-emoji",
			},
			folder.Folder{
				Name: "fried-chicken",
				OrgId: uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385"),
				Paths: "steady-insect.helped-blackheart.fried-chicken",
			},
			folder.Folder{
				Name: "noble-vixen",
				OrgId: uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385"),
				Paths: "creative-scalphunter.steady-insect.helped-blackheart.noble-vixen",
			},
			folder.Folder{
				Name: "iphone-11",
				OrgId: uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385"),
				Paths: "creative-scalphunter.steady-insect.helped-blackheart.noble-vixen.iphone-11",
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
	}

	tests := [...]testCase{testCaseWithThreeFolders, testCaseWithManyFolders}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			folderDriver := folder.NewDriver(tt.folders)
			implementationResult, error := folderDriver.MoveFolder(tt.sourceFolder, tt.destFolder)

			if assert.Nil(error) {
				assert.Equal(tt.want, implementationResult)
			}
		})
	}
}

func TestMoveFolderWithInvalidFolderArguments(t *testing.T) {
	/*
	Test that passing invalid source/dest folders returns an error
	*/
	type testCase struct {
		sourceFolder string
		destFolder	 string
		folders		 []folder.Folder
		want		 []folder.Folder
	}
	assert := assert.New(t)

	nonExistentSourceFolder := testCase{
		"thisIsNonExistent",
		"creative-scalphunter",
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
				OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
				Paths: "stunning-horridus",
			},
		},
		[]folder.Folder{}
	}
	
	nonExistentDestinationFolder := testCase{
		"creative-scalphunter",
		"thisIsNonExistent",
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
				OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
				Paths: "stunning-horridus",
			},
		},
		[]folder.Folder{}
	}

	tests := [...]testCase{nonExistentSourceFolder, nonExistentDestinationFolder}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			folderDriver := folder.NewDriver(tt.folders)
			implementationResult, error := folderDriver.MoveFolder(tt.sourceFolder, tt.destFolder)
			assert.NotNil(error)
			assert.Equal(testCaseSelfMoveFolder.want, implementationResult)
		})
	}
}

func TestMoveFolderWithSelfMove(t *testing.T) {
	/*
	Test that trying to move a folder to itself returns an error
	*/
	type testCase struct {
		sourceFolder string
		destFolder	 string
		folders		 []folder.Folder
		want		 []folder.Folder
	}
	assert := assert.New(t)

	testCaseSelfMoveFolder := testCase{
		"creative-scalphunter",
		"creative-scalphunter",
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
				OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
				Paths: "stunning-horridus",
			},
		},
		[]folder.Folder{}
	}

	t.Run(testCaseSelfMoveFolder.name, func(t *testing.T) {
		folderDriver := folder.NewDriver(testCaseSelfMoveFolder.folders)
		implementationResult, error := folderDriver.MoveFolder(testCaseSelfMoveFolder.sourceFolder, 
			testCaseSelfMoveFolder.destFolder)
		assert.NotNil(error)
		assert.Equal(testCaseSelfMoveFolder.want, implementationResult)
	})
}

func TestMoveFolderOrgDiffersBetweenSourceAndDest(t *testing.T) {
	/*
	Test that trying to move a folder from source to dest
	returns an error if they are in different organisations
	*/
	type testCase struct {
		sourceFolder string
		destFolder	 string
		folders		 []folder.Folder
		want		 []folder.Folder
	}
	assert := assert.New(t)

	testCaseMoveBetweenOrgIDs := testCase{
		"creative-scalphunter",
		"stunning-horridus",
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
				OrgId: uuid.FromStringOrNil("m8b9879b-f73b-4b0e-b9d9-4fc4c23643a2"),
				Paths: "stunning-horridus",
			},
		},
		[]folder.Folder{}
	}

	t.Run(testCaseMoveBetweenOrgIDs.name, func(t *testing.T) {
		folderDriver := folder.NewDriver(testCaseMoveBetweenOrgIDs.folders)
		implementationResult, error := folderDriver.MoveFolder(testCaseMoveBetweenOrgIDs.sourceFolder, 
			testCaseMoveBetweenOrgIDs.destFolder)
		assert.NotNil(error)
		assert.Equal(testCaseMoveBetweenOrgIDs.want, implementationResult)
	})
}

func TestMoveFolderWithSourceAndDestInMultipleOrgIDs() {
	/*
	Test that trying to move a folder which exists in multiple
	organisations returns an error if they are in different organisations
	*/
	type testCase struct {
		sourceFolder string
		destFolder	 string
		folders		 []folder.Folder
		want		 []folder.Folder
	}
	assert := assert.New(t)

	testCaseWithAllFoldersInTwoOrgIDs := testCase{
		"creative-scalphunter",
		"stunning-horridus",
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
				OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
				Paths: "stunning-horridus",
			},
			folder.Folder{
				Name: "creative-scalphunter",
				OrgId: uuid.FromStringOrNil("m8b9879b-f73b-4b0e-b9d9-4fc4c23643a1"),
				Paths: "creative-scalphunter",
			},
			folder.Folder{
				Name: "clear-arclight",
				OrgId: uuid.FromStringOrNil("m8b9879b-f73b-4b0e-b9d9-4fc4c23643a1"),
				Paths: "creative-scalphunter.clear-arclight",
			},
			folder.Folder{
				Name: "stunning-horridus",
				OrgId: uuid.FromStringOrNil("m8b9879b-f73b-4b0e-b9d9-4fc4c23643a1"),
				Paths: "stunning-horridus",
			},
		},
		[]folder.Folder{}
	}

	testCaseWithSourceInTwoOrgIDsAndDestInOne := testCase{
		"creative-scalphunter",
		"stunning-horridus",
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
				OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
				Paths: "stunning-horridus",
			},
			folder.Folder{
				Name: "creative-scalphunter",
				OrgId: uuid.FromStringOrNil("m8b9879b-f73b-4b0e-b9d9-4fc4c23643a1"),
				Paths: "creative-scalphunter",
			},
			folder.Folder{
				Name: "clear-arclight",
				OrgId: uuid.FromStringOrNil("m8b9879b-f73b-4b0e-b9d9-4fc4c23643a1"),
				Paths: "creative-scalphunter.clear-arclight",
			},
		},
		[]folder.Folder{}
	}

	tests := [...]testCase{testCaseWithAllFolderInTwoOrgIDs, testCaseWithSourceInTwoOrgIDsAndDestInOne}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			folderDriver := folder.NewDriver(tt.folders)
			implementationResult, error := folderDriver.MoveFolder(testCaseWithAllFolderInTwoOrgIDs.sourceFolder,
				testCaseWithSourceInTwoOrgIDsAndDestInOne.destFolder)

			assert.NotNil(error)
			assert.Equal(tt.want, implementationResult)
		})
	}

}

func TestMoveFolderWithInvalidPathInChildFolder(t *testing.T) {
	/*
	Test that moving folders with invalid paths in child returns error
	*/
	type testCase struct {
		sourceFolder string
		destFolder	 string
		folders		 []folder.Folder
		want		 []folder.Folder
	}

	testCaseWithSeenPath := testCase{
		"steady-insect",
		"creative-scalphunter",
		[]folder.Folder{
			folder.Folder{
				Name: "creative-scalphunter",
				OrgId: uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385"),
				Paths: "creative-scalphunter",
			},
			folder.Folder{
				Name: "clear-arclight",
				OrgId: uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385"),
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
		"creative-scalphunter",
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
				OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
				Paths: "stunning-horridus",
			},
			folder.Folder{
				Name: "steady-insect",
				OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
				Paths: "steady-insect",
			},
			folder.Folder{
				Name: "helped-blackheart",
				OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
				Paths: "steady-insect.helped-blackheart",
			},
			folder.Folder{
				Name: "many-silver-sable",
				OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
				Paths: "steady-insect.helped-blackheart.many-silver-sable",
			},
			folder.Folder{
				Name: "formula-one",
				OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
				Paths: "steady-insect.helped-blackheart",
			},
			folder.Folder{
				Name: "exciting-magma",
				OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
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
		"creative-scalphunter",
		[]folder.Folder{
			folder.Folder{
				Name: "creative-scalphunter",
				OrgId: uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385"),
				Paths: "creative-scalphunter",
			},
			folder.Folder{
				Name: "clear-arclight",
				OrgId: uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385"),
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

	tests := [...]testCase{
		testCaseWithSeenPath,
		testCaseWithFolderNameNotInPath,
		testCaseWithNonExistentPaths,
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			folderDriver := folder.NewDriver(tt.folders)
			implementationResult, error := folderDriver.GetAllChildFolders(tt.sourceFolder, tt.destFolder)
			assert.NotNil(error)
			assert.Equal(tt.want, implementationResult)
		})
	}
}

func TestMoveFolderWithSkippedFoldersInPath(t *testing.T) {
	/*
	Test that folders are moved without including any
	that skip over non-existent folders (i.e. a folder is missing in
	between, could be due to corrupt files in a practical scenario).
	*/
	type testCase struct {
		name 	string
		orgID	uuid.UUID
		folders	[]folder.Folder
		want	[]folder.Folder
	}
	assert := assert.New(t)

	testCaseWithSkippedFolders := testCase{
		"steady-insect",
		"creative-scalphunter",
		[]folder.Folder{
			folder.Folder{
				Name: "creative-scalphunter",
				OrgId: uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385"),
				Paths: "creative-scalphunter",
			},
			folder.Folder{
				Name: "clear-arclight",
				OrgId: uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385"),
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
		[]folder.Folder{
			folder.Folder{
				Name: "creative-scalphunter",
				OrgId: uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385"),
				Paths: "creative-scalphunter",
			},
			folder.Folder{
				Name: "clear-arclight",
				OrgId: uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385"),
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
				Paths: "creative-scalphunter.steady-insect",
			},
			folder.Folder{
				Name: "helped-blackheart",
				OrgId: uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385"),
				Paths: "creative-scalphunter.steady-insect.helped-blackheart",
			},
			folder.Folder{
				Name: "many-silver-sable",
				OrgId: uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385"),
				Paths: "creative-scalphunter.steady-insect.helped-blackheart.many-silver-sable",
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
	}

	t.Run(testCaseWithSkippedFolders.name, func(t *testing.T) {
		folderDriver := folder.NewDriver(testCaseWithSkippedFolders.folders)
		implementationResult, error := folderDriver.MoveFolder("steady-insect", "creative-scalphunter")
		assert.NotNil(error)
		assert.Equal(testCaseWithSkippedFolders.want, implementationResult)
	})
}
