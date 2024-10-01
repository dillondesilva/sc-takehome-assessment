package folder_test

import (
	"testing"
)

func Test_folder_MoveFolder_Basic1(t *testing.T) {
	// TODO: Test that with a small set of folders,
	// they are correctly moved
}

func Test_folder_MoveFolder_Basic2(t *testing.T) {
	// TODO: Test that with about 5-6 OrgIds and,
	// > 8 different folders they are correctly moved
}

func Test_folder_MoveFolder_Invalid_SourceFolder(t *testing.T) {
	// TODO: Test that passing an invalid source folder results in a 
	// corresponding error being thrown
}

func Test_folder_MoveFolder_Invalid_DestinationFolder(t *testing.T) {
	// TODO: Test that passing an invalid destination folder results
	// in a corresponding error being thrown
}

func Test_folder_MoveFolder_QueryEmptyPath(t *testing.T) {
	// TODO: Test that calling MoveFolder with an empty
	// folder name returns a corresponding error
}

func Test_folder_MoveFolder_QueryEmptyOrgID(t *testing.T) {
	// TODO: Test that calling MoveFolder with an empty
	// OrgID returns a corresponding error
}

func Test_folder_MoveFolder_SelfMove(t *testing.T) {
	// TODO: Test that trying to move a folder to itself results
	// in an error
}

func Test_folder_MoveFolder_OrgDiffersBetweenSourceAndDest(t *testing.T) {
	// TODO: Test that error is thrown if the source folder and destination folder
	// belong to different organisations
}

func Test_folder_MoveFolder_OrgDiffersInTransit(t *testing.T) {
	// TODO: Test that while trying to move the children of the source
	// folder, if any folder doesn't match the OrgID of the
	// destination folder, an error is thrown
}

func Test_folder_MoveFolder_PathDiffersInTransit1(t *testing.T) {
	// TODO: Test that while trying to move the children of the source
	// folder, if any folder has a non-existent path, then an error
	// is returned
}

func Test_folder_MoveFolder_PathDiffersInTransit2(t *testing.T) {
	// TODO: Test that while trying to move the children of the source
	// folder, if any folder has an incorrect path (e.g. bravo is alpha.charlie), 
	// then an error is returned
}

func Test_folder_MoveFolder_DuplicateFolderInTransit(t *testing.T) {
	// TODO: Test that while trying to move the folder in transit,
	// if it already appears in the destination folder, an error is thrown
}

func Test_folder_MoveFolder_VolumeTest(t *testing.T) {
	// TODO: Test that when a large number of folders are in use,
	// the function still returns the correct results.
}
