package builder

import "os"

type Generator struct {
	ProjectName string
	Directory   string
	Files       []File
}

type File struct {
	Name    string
	Content string
}

func NewBuilder() *Generator {

	currentFolder, err := os.Getwd()
	if err != nil {
		panic("Could not get current working directory")
	}

	return &Generator{
		ProjectName: currentFolder,
		Directory:   "",
		Files:       []File{},
	}
}
