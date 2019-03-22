// Package moodel contains models used along the excersice
package model

import (
	"errors"
	"fmt"
	"os"

	"github.com/lfernandezlo/hackrn/constant"
)

// Directory struct represents terminal abstraction
type Directory struct {
	Folders          []*Folder
	CurrentDirectory string
	CurrentFolder    *Folder
}

// Touch create file inside current folder files
func (d *Directory) Touch(f string) (*string, error) {
	// Validate file length
	var file *string

	if len(f) > 100 {
		return file, errors.New(constant.ErrorInvalidFileOrFolderName)
	}

	// Should have current folder (root or created ones) this only happens when trying to touch without a root folder instanciated
	if d.CurrentFolder == nil {
		return file, errors.New(constant.ErrorNoRootFolder)
	}

	files := &d.CurrentFolder.Files

	// Validate existance
	var exists bool

	if len(*files) >= 1 {
		for _, file := range *files {
			if file == f {
				exists = true
				break
			}
		}
	}

	if exists {
		return file, errors.New(constant.ErrorFileAlreadyExists)
	}

	file = &f

	*files = append(*files, *file)

	return file, nil
}

// Mkdir creates new dir in working directory
func (d *Directory) Mkdir(dir string) (*Folder, error) {
	// Validate dir length
	if len(dir) > 100 {
		return &Folder{}, errors.New(constant.ErrorInvalidFileOrFolderName)
	}

	// Check destination
	des := &d.Folders

	// If root folder has not been instanciated, create it
	if len(*des) == 0 {
		rootFolder := &Folder{Name: "root", Path: "/root"}
		d.CurrentDirectory = rootFolder.Path
		d.CurrentFolder = rootFolder

		*des = append(*des, rootFolder)
	}

	des = &d.CurrentFolder.Folders

	// Validate existance
	var exists bool

	for _, f := range *des {
		if f.Name == dir {
			exists = true
			break
		}
	}

	if exists {
		return &Folder{}, errors.New(constant.ErrorDirectoryAlreadyExists)
	}

	// Create folder using working directory
	workingDirectory := d.Pwd()

	path := fmt.Sprintf("%v/%v", workingDirectory, dir)
	newFolder := &Folder{Name: dir, Path: path, ParentFolder: d.CurrentFolder}

	*des = append(*des, newFolder)

	return newFolder, nil
}

// Ls list the contents (directories and files) from current folder
func (d *Directory) Ls() string {
	var content string

	currentFolder := d.CurrentFolder

	if currentFolder != nil {
		for i, f := range d.CurrentFolder.Folders {
			newLine := "\n"

			if i == len(d.CurrentFolder.Folders)-1 && len(currentFolder.Files) == 0 {
				newLine = ""
			}

			if f == nil {
				continue
			}

			content += f.Name + newLine
		}
	}

	if len(currentFolder.Files) > 0 {
		for i, f := range currentFolder.Files {
			newLine := "\n"

			if i == len(currentFolder.Files)-1 {
				newLine = ""
			}

			content += f + newLine
		}
	}

	return content
}

// Pwd prints full path of current directory
func (d *Directory) Pwd() string {
	return d.CurrentDirectory
}

// Cd Changes the current path to a sub directory by name
func (d *Directory) Cd(dir string) (string, error) {
	// Folders to check
	folders := d.Folders

	if d.CurrentFolder != nil {
		folders = d.CurrentFolder.Folders
	}

	// If dir its .. and there is a parent folder, move to it
	if dir == ".." && d.CurrentFolder.ParentFolder != nil {
		parent := d.CurrentFolder.ParentFolder
		d.CurrentDirectory = parent.Path
		d.CurrentFolder = parent
		return d.Pwd(), nil
	}

	// If dir its .. and there is not parent folder, it means that terminal is already on root
	if dir == ".." && d.CurrentFolder.ParentFolder == nil {
		return "", errors.New(constant.ErrorAlreadyOnRoot)
	}

	// Check existance
	var exists bool

	for _, f := range folders {
		if f.Name == dir {
			d.CurrentDirectory = f.Path
			d.CurrentFolder = f
			exists = true
		}
	}

	if !exists {
		return d.Pwd(), errors.New(constant.ErrorDirectoryNotFound)
	}

	return d.Pwd(), nil
}

// Quit exits from terminal
func (d *Directory) Quit(status int) {
	os.Exit(status)
}
