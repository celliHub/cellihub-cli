package aws

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GenerateDevContainer() {
	dir := ".devcontainer"

	// Get current project name from directory
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}
	projectName := filepath.Base(currentDir)

	// Create the directory
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}

	// File paths
	devcontainerPath := filepath.Join(dir, "devcontainer.json")
	dockerfilePath := filepath.Join(dir, "Dockerfile")

	// Template with placeholder
	devcontainerTemplate := `{
  "name": "%PROJECT_NAME% Dev Container",
  "build": {
    "dockerfile": "Dockerfile"
  },
  "customizations": {
    "vscode": {
      "settings": {
        "terminal.integrated.shell.linux": "/bin/bash"
      },
      "extensions": [
        "golang.go",
        "ms-python.python",
        "ms-azuretools.vscode-docker",
        "dbaeumer.vscode-eslint",
        "ms-vscode.makefile-tools",
        "hashicorp.terraform"
      ]
    }
  },
  "postCreateCommand": "go version && node -v && npm -v && python3 --version && terraform --version && aws --version && kubectl version --client"
}`

	// Replace the placeholder
	devcontainerContent := strings.ReplaceAll(devcontainerTemplate, "%PROJECT_NAME%", projectName)

	// Dockerfile
	dockerfileContent := `FROM golang:1.23-bullseye

# Instala dependências
RUN apt-get update && apt-get install -y \
    unzip \
    curl \
    git \
    python3 \
    python3-pip \
    nodejs \
    npm \
    awscli \
    && rm -rf /var/lib/apt/lists/*

# Instala Terraform
RUN curl -fsSL https://releases.hashicorp.com/terraform/1.9.5/terraform_1.9.5_linux_amd64.zip -o terraform.zip && \
    unzip terraform.zip && mv terraform /usr/local/bin/ && rm terraform.zip

# Instala Kubectl
RUN curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl" && \
    install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl && rm kubectl

WORKDIR /workspace
`

	// Write the files
	writeFile(devcontainerPath, devcontainerContent)
	writeFile(dockerfilePath, dockerfileContent)

	fmt.Printf("✅ DevContainer created successfully at '%s' (project: %s)\n", dir, projectName)
}

func writeFile(path, content string) {
	err := os.WriteFile(path, []byte(content), 0644)
	if err != nil {
		fmt.Println("Error creating file", path, ":", err)
	}
}
