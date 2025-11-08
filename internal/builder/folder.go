package builder

import (
	"log"
	"os"
)

func (g *Generator) CreateFolder(dirName string) {
	FullPathFolder := os.Getenv("PWD") + "/" + dirName
	log.Println("Creating folder:", FullPathFolder)

	if err := os.Mkdir(FullPathFolder, 0755); err != nil {
		log.Println("Error creating folder:", err)
		panic("Could not create folder")
	}

	g.Directory = dirName

}

func (g *Generator) CheckFolder() {

	if g.Directory == "" {
		log.Println("No directory specified.")
		panic("Not implemented yet")
	}
}
