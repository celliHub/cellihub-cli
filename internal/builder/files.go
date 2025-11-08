package builder

import (
	"cellihub-cli/internal/templates/cloud"
	"log"
	"os"
	"strings"
)

func (g *Generator) CreateFiles(files []string, category string) {

	for i, f := range files {

		var ContentFile string
		switch f {
		case "devcontainer.json":
			ContentFile = cloud.DevcontainerContent
		case "Dockerfile":
			ContentFile = cloud.DockerfileContent
		default:
			log.Println("File not recognized:", f)
			continue
		}

		file := File{
			Name:    f,
			Content: ContentFile,
		}

		FilePath, _ := os.Getwd()
		FileFullPath := FilePath + string(os.PathSeparator) + ".devcontainer" + string(os.PathSeparator) + file.Name
		log.Printf("[%d]Creating file: %s", i, FileFullPath)

		if _, err := os.Create(FileFullPath); err != nil {
			log.Println("Error creating file:", err)
			panic("Could not create file")
		}

		// Write content to the file
		// %PROJECT_NAME%
		ProjectFolderName := strings.Split(FilePath, string(os.PathSeparator))
		Content := strings.ReplaceAll(file.Content, "%PROJECT_NAME%", ProjectFolderName[len(ProjectFolderName)-1])
		if err := os.WriteFile(FileFullPath, []byte(Content), 0644); err != nil {
			log.Println("Error writing to file:", err)
			panic("Could not write to file")
		}

		g.Files = append(g.Files, file)
	}
}
