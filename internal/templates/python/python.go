package python

type Template struct {
	Category string
	Files    []string
	Content  map[string]string
}

var DevcontainerContent = `{
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

var DockerfileContent = `FROM ubuntu:22.04

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
    golang \
    && rm -rf /var/lib/apt/lists/*

# Instala Terraform
RUN curl -fsSL https://releases.hashicorp.com/terraform/1.9.5/terraform_1.9.5_linux_amd64.zip -o terraform.zip && \
    unzip terraform.zip && mv terraform /usr/local/bin/ && rm terraform.zip

# Instala Kubectl
RUN curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl" && \
    install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl && rm kubectl

WORKDIR /workspace
`

var PythonTemplate = Template{
	Category: "python",
	Files:    []string{"devcontainer.json", "Dockerfile"},
	Content: map[string]string{
		"devcontainer.json": DevcontainerContent,
		"Dockerfile":        DockerfileContent,
	},
}
