# Developer Golden Path CLI

A portfolio project demonstrating Platform Engineering concepts by building a Go-based CLI tool. This CLI scaffolds new microservices ("The Golden Path") with DevSecOps best practices — like distroless containers, non-root security contexts, and CI scanning — baked in by default.

## The Problem
"Shift Left" security often fails when developers are asked to become security experts. If a developer has to manually write a `Dockerfile` and `deployment.yaml` from scratch for every new service, they will inevitably copy-paste insecure defaults (running as root, exposing broad capabilities, skipping CI scanners).

## The Solution
Instead of gating releases with security checks at the end, Platform Engineering provides a "Golden Path": a self-service CLI that generates a fully compliant, secure-by-default boilerplate repository instantly.

```text
+-------------------+       +-----------------------+       +-------------------+
|                   |       |                       |       |   - Distroless    |
|   Developer runs  | ----> |   Golden Path CLI     | ----> |   - Non-root      |
|   `./golden-path` |       |   (Go Templating)     |       |   - SAST/SCA CI   |
|                   |       |                       |       |   - CPU/Mem limits|
+-------------------+       +-----------------------+       +-------------------+
```

## Why This Over the Obvious Alternative
While tools like Backstage provide UI-based scaffolding (via Software Templates), building the underlying generation engine in Go demonstrates a deep understanding of what exactly constitutes a "secure default" and how to dynamically template Kubernetes/Docker configurations.

## Tech Stack
- **Language**: Go 1.21+
- **Templating**: `text/template` & `embed.FS` (Templates baked directly into the binary)
- **Output Target**: Node.js/Express, Docker, Kubernetes, GitHub Actions

## Decision Log

| Component | Decision | Rationale |
| :--- | :--- | :--- |
| **Container Base** | Distroless (`gcr.io/distroless/nodejs20`) | Reduces attack surface by removing package managers and shells, making it significantly harder for an attacker to establish a foothold if the application is compromised. |
| **K8s Security** | `runAsNonRoot: true` | Mandating `fsGroup: 65532` and `drop: ALL` capabilities ensures that container escapes are functionally impossible. |
| **CI/CD** | Semgrep & Trivy | Automatically scaffolds `.github/workflows/ci.yml` so that static analysis and container scanning run on PR #1. |

## Project Structure

```text
golden-path-cli/
├── pkg/scaffold/           # Generator logic for mapping and executing templates
├── templates/              # The "Golden Path" templates (.tmpl) (Embedded in binary)
│   ├── Dockerfile.tmpl
│   ├── k8s-deployment.yaml.tmpl
│   ├── k8s-service.yaml.tmpl
│   ├── ci.yml.tmpl
│   ├── index.js.tmpl
│   └── package.json.tmpl
├── main.go                 # CLI entrypoint parsing flags
├── go.mod                  # Go module definition
└── README.md               # This file
```

## Setup & Usage

### 1. Build the CLI
```bash
go build -o golden-path-cli
```

### 2. Scaffold a New Service
Run the CLI, providing a name and a port:
```bash
./golden-path-cli -name secure-payments-api -port 8080 -out ./my-new-service
```

### 3. Verify the Output
Inspect the generated directory:
```bash
cd my-new-service
ls -la
```
Notice how `k8s/deployment.yaml` and `Dockerfile` have been securely populated with the variables provided.

## Verification

| Check | Expected Result |
| :--- | :--- |
| CLI Execution | Runs without errors and generates all expected files. |
| Security Contexts | `grep runAsNonRoot my-new-service/k8s/deployment.yaml` returns `true`. |
| Port Injection | `grep 8080 my-new-service/index.js` shows the correct templated port. |

## Author

**Sumit Dalavi — Senior DevSecOps / Platform Engineer**
- [GitHub](https://github.com/your-username)
- [LinkedIn](https://linkedin.com/in/your-profile)
