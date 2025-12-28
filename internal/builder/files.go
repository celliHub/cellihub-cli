package builder

import (
	"cellihub-cli/internal/templates/cloud"
	"cellihub-cli/internal/templates/light"
	"log"
	"os"
	"strings"
)

type Template struct {
	Category string
	Files    []string
	Content  map[string]string
}

func (g *Generator) CreateFiles(category string) {

	t := Template{}

	switch strings.ToLower(category) {
	case "cloud":
		t = Template(cloud.CloudTemplate)
	case "light":
		t = Template(light.LightTemplate)
	default:
		log.Fatalf("Unsupported category for files: %s", category)
	}

	for i, f := range t.Files {

		var ContentFile string
		switch f {
		case "devcontainer.json":
			ContentFile = t.Content["devcontainer.json"]
		case "Dockerfile":
			ContentFile = t.Content["Dockerfile"]
		case "post-commands.sh":
			ContentFile = t.Content["post-commands.sh"]
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
