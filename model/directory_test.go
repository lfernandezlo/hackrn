// package model_test provides tests for models methods
package model_test

import (
	"strings"
	"testing"

	"github.com/lfernandezlo/hackrn/constant"
	"github.com/lfernandezlo/hackrn/model"
	"github.com/stretchr/testify/assert"
)

// Mkdir tests
func TestDirectory_Mkdir(t *testing.T) {
	type args struct {
		dir string
	}

	tests := []struct {
		name               string
		directory          model.Directory
		args               args
		wantErr            bool
		errorDetail        string
		expectedFolderName string
		expectedFolderPath string
	}{
		{
			name: "TestMkdir",
			args: args{
				dir: "test",
			},
			expectedFolderName: "test",
			expectedFolderPath: "/root/test",
		},
		{
			name: "TestShouldFailOnInvalidLen",
			args: args{
				dir: strings.Repeat("x", 101),
			},
			wantErr:     true,
			errorDetail: constant.ErrorInvalidFileOrFolderName,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &model.Directory{
				Folders:          tt.directory.Folders,
				CurrentDirectory: tt.directory.CurrentDirectory,
				CurrentFolder:    tt.directory.CurrentFolder,
			}
			assert := assert.New(t)

			f, err := d.Mkdir(tt.args.dir)

			if tt.wantErr {
				assert.NotNil(err)
				assert.Equal(tt.errorDetail, err.Error())
				t.Skip()
			}

			assert.Nil(err)
			assert.Equal(tt.expectedFolderPath, f.Path)
			assert.Equal(tt.expectedFolderName, f.Name)
		})
	}
}

func TestDirectory_Mkdir_ShouldFailOnExistentDirectory(t *testing.T) {
	d := model.Directory{}
	_, err := d.Mkdir("test")

	assert := assert.New(t)

	assert.Nil(err)

	_, err = d.Mkdir("test")

	assert.NotNil(err)
	assert.EqualError(err, constant.ErrorDirectoryAlreadyExists)
}

// Cd tests

func TestCd(t *testing.T) {
	d := model.Directory{}
	_, err := d.Mkdir("cdtest")

	assert := assert.New(t)

	assert.Nil(err)

	path, err := d.Cd("cdtest")

	assert.Nil(err)

	assert.Equal("/root/cdtest", path)
}

func TestCdShouldFailOnUnfoundDir(t *testing.T) {
	d := model.Directory{}

	assert := assert.New(t)

	_, err := d.Cd("cdtest")

	assert.NotNil(err)

	assert.EqualError(err, constant.ErrorDirectoryNotFound)
}

// Touch tests

func TestTouch(t *testing.T) {
	type args struct {
		f string
	}

	rootFolder := model.Folder{Name: "root", Path: "/root"}

	tests := []struct {
		name         string
		directory    model.Directory
		args         args
		wantErr      bool
		errorDetail  string
		expectedFile string
	}{
		{
			name: "TestShouldFailOnInvalidLen",
			args: args{
				f: strings.Repeat("x", 101),
			},
			directory: model.Directory{
				Folders:          []*model.Folder{&rootFolder},
				CurrentFolder:    &rootFolder,
				CurrentDirectory: "/root",
			},
			wantErr:     true,
			errorDetail: constant.ErrorInvalidFileOrFolderName,
		},
		{
			name: "TestShouldFailOnUnfoundRootFolder",
			args: args{
				f: "x",
			},
			directory: model.Directory{
				Folders: []*model.Folder{},
			},
			wantErr:     true,
			errorDetail: constant.ErrorNoRootFolder,
		},
		{
			name: "TestShouldCreateFileX",
			args: args{
				f: "x",
			},
			directory: model.Directory{
				Folders:          []*model.Folder{&rootFolder},
				CurrentFolder:    &rootFolder,
				CurrentDirectory: "/root",
			},
			expectedFile: "x",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &model.Directory{
				Folders:          tt.directory.Folders,
				CurrentDirectory: tt.directory.CurrentDirectory,
				CurrentFolder:    tt.directory.CurrentFolder,
			}

			assert := assert.New(t)

			f, err := d.Touch(tt.args.f)

			if tt.wantErr {
				assert.NotNil(err)
				assert.Equal(tt.errorDetail, err.Error())
				t.Skip()
			}

			assert.Nil(err)
			assert.Equal(tt.expectedFile, *f)
		})
	}
}

func TestCdShouldReturnErrorOnDuplicatedFile(t *testing.T) {
	rootFolder := model.Folder{Name: "root", Path: "/root"}

	d := model.Directory{
		Folders:          []*model.Folder{&rootFolder},
		CurrentFolder:    &rootFolder,
		CurrentDirectory: "/root",
	}

	f, err := d.Touch("test")

	assert := assert.New(t)

	assert.Nil(err)
	assert.Equal("test", *f)

	f, err = d.Touch("test")

	assert.NotNil(err)

	assert.EqualError(err, constant.ErrorFileAlreadyExists)
}

// Ls tests

func TestLs(t *testing.T) {
	rootFolder := model.Folder{
		Name: "root",
		Path: "/root",
		Folders: []*model.Folder{
			&model.Folder{Name: "test", Path: "/root/test"},
			&model.Folder{Name: "example", Path: "/root/example"},
		},
		Files: []string{"testfile", "testfile2"},
	}

	d := model.Directory{
		Folders:          []*model.Folder{&rootFolder},
		CurrentFolder:    &rootFolder,
		CurrentDirectory: "/root",
	}

	output := d.Ls()
	expectedOutput := "test\nexample\ntestfile\ntestfile2"
	assert := assert.New(t)

	assert.Equal(expectedOutput, output)
}
