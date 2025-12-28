# CelliHub CLI

**CelliHub CLI** is a lightweight command-line interface to help developers quickly scaffold and configure local development environments using **Dev Containers**.

It provides ready-to-use templates for cloud environments (AWS, GCP, and Azure) and aims to seamlessly integrate with the upcoming **CelliHub SaaS platform**, allowing users to synchronize infrastructure and environment configurations directly from the CLI.

---

## Overview

The **CelliHub CLI** automates the creation of configuration files and folders that define a development container environment

With a single command, it can generate a `.devcontainer/` folder containing:
- `devcontainer.json` — fully preconfigured for your cloud provider and language.
- `Dockerfile` — ready to build your local container image with preinstalled tools (Terraform, AWS CLI, Go, Node.js, Python, etc.).

---

## 🧱 Example Usage

To scaffold a DevContainer for an AWS-based project:

```
cellihub-cli cloud aws devcontainer
```