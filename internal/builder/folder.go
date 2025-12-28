package builder

import (
	"log"
	"os"
)

func (g *Generator) CreateFolder(dirName string) {
	fullPathFolder := os.Getenv("PWD") + "/" + dirName
	log.Println("Creating folder:", fullPathFolder)

	if err := os.Mkdir(fullPathFolder, 0755); err != nil {
		log.Println("Error creating folder:", err)
		log.Fatal("Failed to create folder. Recommended remove .devcontainer folder if exists and try again.")
		os.Exit(1)
	}

	g.Directory = dirName

}

func (g *Generator) CheckFolder() {

	if g.Directory == "" {
		log.Println("No directory specified.")
		panic("Not implemented yet")
	}
}
