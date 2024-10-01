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
