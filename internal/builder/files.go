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

		var contentFile string
		switch f {
		case "devcontainer.json":
			contentFile = t.Content["devcontainer.json"]
		case "Dockerfile":
			contentFile = t.Content["Dockerfile"]
		case "post-commands.sh":
			contentFile = t.Content["post-commands.sh"]
		default:
			log.Println("File not recognized:", f)
			continue
		}

		file := File{
			Name:    f,
			Content: contentFile,
		}

		filePath, _ := os.Getwd()
		fileFullPath := filePath + string(os.PathSeparator) + ".devcontainer" + string(os.PathSeparator) + file.Name
		log.Printf("[%d]Creating file: %s", i, fileFullPath)

		if _, err := os.Create(fileFullPath); err != nil {
			log.Println("Error creating file:", err)
			panic("Could not create file")
		}

		// Write content to the file
		// %PROJECT_NAME%
		projectFolderName := strings.Split(filePath, string(os.PathSeparator))
		content := strings.ReplaceAll(file.Content, "%PROJECT_NAME%", projectFolderName[len(projectFolderName)-1])
		if err := os.WriteFile(fileFullPath, []byte(content), 0644); err != nil {
			log.Println("Error writing to file:", err)
			panic("Could not write to file")
		}

		g.Files = append(g.Files, file)
	}
}
