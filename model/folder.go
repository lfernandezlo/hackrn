package model

// Folder struct represents folder information
type Folder struct {
	Name         string
	Path         string
	Folders      []*Folder
	Files        []string
	ParentFolder *Folder
}
