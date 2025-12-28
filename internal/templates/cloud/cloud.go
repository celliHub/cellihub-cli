package cloud

type Template struct {
	Category string
	Files    []string
	Content  map[string]string
}

var DevcontainerContent = `{
  "name": "%PROJECT_NAME% Dev Container",
  "build": {
    "dockerfile": "Dockerfile",
    "args": {
      "DOCKER_GID": "80"
    },
    "options": ["--platform=linux/arm64"]
  },
  "features": {
    "ghcr.io/devcontainers/features/common-utils:2": {},
    "ghcr.io/devcontainers/features/terraform:1": {
      "version": "latest"
    }
  },
  "customizations": {
    "vscode": {
      "settings": {
        "terminal.integrated.shell.linux": "/bin/bash"
      },
      "extensions": [
        "ms-azuretools.vscode-docker",
        "hashicorp.terraform",
        "amazonwebservices.aws-toolkit-vscode",
        "redhat.vscode-extension-bpmn-editor",
        "openai.chatgpt",
        "ms-azuretools.vscode-containers",
        "golang.go",
        "ms-vscode.vscode-typescript-next",
        "esbenp.prettier-vscode",
        "ms-vscode.remote-server",
        "rangav.vscode-thunder-client"
      ]
    }
  },
  "mounts": [
    "source=${localWorkspaceFolder},target=/workspace,type=bind",
    // Se desejar acesso ao Docker do host, descomente a linha abaixo
    // "source=/var/run/docker.sock,target=/var/run/docker.sock,type=bind"

    // Volume isolado opcional
    // "source=clienteA-data,target=/workspace,type=volume"
  ],
  "runArgs": [
    "--name=dev-${localWorkspaceFolderBasename}"
  ],
  "remoteUser": "vscode",
  "postStartCommand": "./.devcontainer/post-commands.sh"
}`

var DockerfileContent = `FROM mcr.microsoft.com/devcontainers/base:ubuntu-22.04

# -----------------------------
# Basic dependencies
# -----------------------------
RUN apt-get update && apt-get install -y \
    unzip \
    curl \
    less \
    groff \
    tar \
    ca-certificates \
    gnupg \
    && rm -rf /var/lib/apt/lists/*

ARG DOCKER_GID=80
RUN groupadd -for -g ${DOCKER_GID} docker && usermod -aG docker vscode

# -----------------------------
# Install Golang
# -----------------------------
ARG GO_VERSION=1.25.5
RUN wget https://go.dev/dl/go${GO_VERSION}.linux-arm64.tar.gz \
    && tar -C /usr/local -xzf go${GO_VERSION}.linux-arm64.tar.gz \
    && rm go${GO_VERSION}.linux-arm64.tar.gz && \
    echo 'export PATH=$PATH:/usr/local/go/bin:$HOME/go/bin' >> /etc/profile.d/go.sh
ENV PATH="/usr/local/go/bin:${PATH}"
ENV GOPATH="/home/vscode/go"

# -----------------------------
# Docker CLI
# -----------------------------
# RUN apt-get update && apt-get install -y docker-cli && rm -rf /var/lib/apt/lists/*

# -----------------------------
# AWS CLI v2 (ARM64)
# -----------------------------
RUN curl "https://awscli.amazonaws.com/awscli-exe-linux-aarch64.zip" -o "awscliv2.zip" \
    && unzip awscliv2.zip \
    && ./aws/install \
    && rm -rf aws awscliv2.zip

# -----------------------------
# AWS SSM Session Manager Plugin
# -----------------------------
RUN curl -s "https://s3.amazonaws.com/session-manager-downloads/plugin/latest/ubuntu_arm64/session-manager-plugin.deb" \
      -o "session-manager-plugin.deb" \
    && apt-get install -y ./session-manager-plugin.deb \
    && rm -f session-manager-plugin.deb

# -----------------------------
# Install kubectl (official repo)
# -----------------------------
RUN mkdir -p /etc/apt/keyrings \
    && curl -fsSL https://pkgs.k8s.io/core:/stable:/v1.30/deb/Release.key \
        | gpg --dearmor -o /etc/apt/keyrings/kubernetes-apt-keyring.gpg \
    && echo "deb [signed-by=/etc/apt/keyrings/kubernetes-apt-keyring.gpg] \
        https://pkgs.k8s.io/core:/stable:/v1.30/deb/ /" \
        | tee /etc/apt/sources.list.d/kubernetes.list \
    && apt-get update \
    && apt-get install -y kubectl \
    && rm -rf /var/lib/apt/lists/*

# -----------------------------
# Install NVM + Latest Node + Latest npm
# -----------------------------
USER vscode
ENV NVM_DIR=/home/vscode/.nvm
ENV NODE_VERSION=stable

RUN curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.6/install.sh | bash \
    && . "$NVM_DIR/nvm.sh" \
    && nvm install $NODE_VERSION \
    && nvm use $NODE_VERSION \
    && nvm alias default $NODE_VERSION \
    && npm install -g firebase-tools \
    && npm install -g npm

RUN echo 'export NVM_DIR="$HOME/.nvm"' >> /home/vscode/.bashrc \
    && echo '[ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh"' >> /home/vscode/.bashrc

USER root

# -----------------------------
# Install k9s (latest ARM64)
# -----------------------------
RUN wget https://github.com/derailed/k9s/releases/latest/download/k9s_linux_arm64.deb \
    && apt install ./k9s_linux_arm64.deb && rm k9s_linux_arm64.deb

# Ensure vscode user is in docker group
RUN groupadd -f docker && usermod -aG docker vscode`

var PostCommandsContent = `#!/bin/bash
echo "Post-commands script."

aws --version
terraform --version
session-manager-plugin --version
    
echo "Post-commands script completed."`

var CloudTemplate = Template{
	Category: "cloud",
	Files:    []string{"devcontainer.json", "Dockerfile", "post-commands.sh"},
	Content: map[string]string{
		"devcontainer.json": DevcontainerContent,
		"Dockerfile":        DockerfileContent,
		"post-commands.sh":  PostCommandsContent,
	},
}
