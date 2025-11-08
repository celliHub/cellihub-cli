FROM ubuntu:22.04

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
