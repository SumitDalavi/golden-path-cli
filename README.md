> **NOTE:** This repository is an archival lab or partial prototype. It is not actively maintained and should not be used as a reference for production-grade deployments or performance benchmarks.

# Developer Golden Path CLI

[![CI Pipeline](https://github.com/SumitDalavi/golden-path-cli/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/SumitDalavi/golden-path-cli/actions/workflows/ci.yml)
[![Release](https://img.shields.io/github/v/release/SumitDalavi/golden-path-cli?color=0969da&label=release&logo=github)](https://github.com/SumitDalavi/golden-path-cli/releases)
[![Release Workflow](https://github.com/SumitDalavi/golden-path-cli/actions/workflows/release.yml/badge.svg)](https://github.com/SumitDalavi/golden-path-cli/actions/workflows/release.yml)
[![GHCR Container](https://img.shields.io/badge/ghcr.io-v0.2.0-blue?logo=docker&logoColor=white)](https://github.com/SumitDalavi/golden-path-cli/pkgs/container/golden-path-cli)
[![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go&logoColor=white)](https://golang.org)
[![SBOM](https://img.shields.io/badge/SBOM-SPDX%202.3-blueviolet?logo=spdx)](https://github.com/SumitDalavi/golden-path-cli/releases/tag/v0.2.0)
[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)

[![Go Reference](https://pkg.go.dev/badge/github.com/SumitDalavi/golden-path-cli.svg)](https://pkg.go.dev/github.com/SumitDalavi/golden-path-cli)

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

## 🔒 Release & Security Verification

Every tagged release (
*.*.*) triggers an automated DevSecOps release workflow:
1. **Automated Testing:** Runs go test -v ./... across unit tests.
2. **Binary Packaging:** Compiles standalone CLI binary asset (golden-path-cli-linux-amd64.tar.gz).
3. **Container Image Build:** Builds the CLI container image (ghcr.io/sumitdalavi/golden-path-cli:<tag>).
4. **Vulnerability Scanning:** Scans the CLI container using [Aqua Security Trivy](https://github.com/aquasecurity/trivy-action) for HIGH and CRITICAL CVEs.
5. **SBOM Generation:** Generates a Software Bill of Materials (SBOM) in SPDX JSON format (sbom.spdx.json) using [Anchore Syft](https://github.com/anchore/sbom-action).
6. **GitHub Release Publishing:** Packages release notes and attaches binary tarball and sbom.spdx.json as release assets.

## Author

**Sumit Dalavi — Senior DevSecOps / Platform Engineer**
- [GitHub](https://github.com/SumitDalavi)
- [LinkedIn](https://in.linkedin.com/in/sumit-dalavi-762838129)

## CI & Reliability Updates (August 2026)

- **CI Pipeline Remediation:** Successfully resolved all CI/CD pipeline failures and established baseline CI workflows.
- **Specific Fix:** Added and configured robust GitHub Actions workflows for automated testing, linting, and formatting.
- **Status:** 🟩 Passing

## Release & Security Verification

Every tagged release (`v*.*.*`) triggers an automated DevSecOps release workflow:
1. **Automated Testing:** Runs test suite gates before build.
2. **Container Image Build:** Builds production container image (`ghcr.io/sumitdalavi/golden-path-cli:<tag>`).
3. **Vulnerability Scanning:** Scans the image using [Aqua Security Trivy](https://github.com/aquasecurity/trivy-action) for `HIGH` and `CRITICAL` CVEs.
4. **SBOM Generation:** Generates a Software Bill of Materials (SBOM) in SPDX JSON format (`sbom.spdx.json`) using [Anchore Syft](https://github.com/anchore/sbom-action).
5. **GitHub Release Publishing:** Packages release notes and attaches `sbom.spdx.json` as a verifiable supply-chain asset.

### Verifying Releases via GitHub CLI
```bash
# View release details
gh release view v0.1.0 --repo SumitDalavi/golden-path-cli

# Download and verify SBOM
gh release download v0.1.0 --repo SumitDalavi/golden-path-cli --pattern "sbom.spdx.json"
cat sbom.spdx.json | jq .
```

## 🚀 Progressive Credential Loading & Execution Tiers

This project implements **Progressive Credential Loading (PCL)**, allowing it to run immediately out-of-the-box without external dependencies while seamlessly upgrading as credentials are provided.

### Tier 1: Zero-Config Mock Mode (Default)
Ideal for CI, testing, and rapid local evaluation.
- **Execution:** Run with default configuration (`cp .env.example .env`). Zero API keys, databases, or cloud credentials required.
- **Fallback Behavior:** In-memory caches, deterministic offline heuristics, mock vector embeddings, and local mock storage are automatically used.
- **Health Verification:** Endpoints report `{"mode": "mock"}` for unconfigured dependencies.

### Tier 2: Local Development with Live Services
Connect specific external services while keeping others mocked.
- **Execution:** Provide credentials in `.env` for the services you want active (e.g., `GITHUB_TOKEN, TEMPLATE_CATALOG_URL`).
- **Progressive Upgrade:** The runtime inspects environment variables on boot and dynamically promotes only configured modules to live mode.
- **Safe Fallback:** Any missing credentials automatically remain on mock fallbacks with no crashes or missing service errors.

### Tier 3: Production Hardened
Enterprise deployment with production-grade backends, security scans, and container distributions.
- **OCI Container Image:** Pre-built, multi-arch image on GitHub Container Registry:
  ```bash
  docker pull ghcr.io/sumitdalavi/golden-path-cli:v0.2.0
  ```
- **Production Orchestration:** Mount secrets via Kubernetes Secrets or cloud secret managers with atomic config reload support.
- **Release Verification:** Checksums and SPDX SBOM available on [GitHub Releases](https://github.com/SumitDalavi/golden-path-cli/releases).

## Documentation

- [Architecture](docs/ARCHITECTURE.md)
- [Runbook](docs/runbook.md)
- [Decisions](docs/decisions.md)
- [Changelog](docs/changelog.md)
