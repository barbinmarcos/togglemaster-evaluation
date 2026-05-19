# ToggleMaster Evaluation Service

Microsserviço responsável pela avaliação das regras de Feature Flags.

## Responsabilidades

- Avaliação de regras
- Processamento de decisões
- Retorno de flags ativas
- Engine de decisão

## Stack

- Golang
- Docker
- Kubernetes
- GitHub Actions
- Amazon ECR
- Amazon EKS

## Execução local

```bash
go run cmd/main.go
Endpoint de Health Check
GET /health
Docker

Build:

docker build -t togglemaster-evaluation .

Run:

docker run -p 8083:8083 togglemaster-evaluation
CI/CD

Pipeline DevSecOps com:

Build
Unit Tests
GolangCI-Lint
Gosec
Trivy
Docker Build
Push para Amazon ECR
Deploy

Deploy automatizado via GitOps utilizando ArgoCD.


