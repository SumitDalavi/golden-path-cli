# Changelog for golden-path-cli

## [0.2.0] - 2026-10-20
### Added
- **Progressive Credential Loading (PCL):** Dynamic auto-promotion between zero-dependency mock fallbacks and live production backends based on environment credentials.
- **Dual Distribution:** Universal OCI container publishing to GitHub Container Registry (`ghcr.io/sumitdalavi/golden-path-cli`) and Helm OCI registries alongside GitHub Release tarballs.
- **Health Reporting:** Exposed `/health` or `/healthz` endpoint with PCL status detection (`mock` vs `live`).
- **Dynamic Reloading:** Support for atomic ConfigMap / secret reloads without requiring pod restarts.
- **Supply Chain Security:** Trivy vulnerability scan gates and SPDX JSON SBOM (`sbom.spdx.json`) generation for all releases.

## [0.1.0] - 2026-10-18
### Added
- Initial baseline service implementation, unit test suites, and CI workflows.

## [v0.1.0] — 2026-10-18
### Added
- Automated GitHub Actions release pipeline triggered on `v*.*.*` git tags (`.github/workflows/release.yml`).
- CLI standalone binary release asset (`golden-path-cli-linux-amd64.tar.gz`).
- Container image packaging (`ghcr.io/sumitdalavi/golden-path-cli`).
- Container vulnerability scanning via Aqua Security Trivy (`aquasecurity/trivy-action@v0.36.0`).
- SPDX-format Software Bill of Materials (SBOM) generation step (`anchore/sbom-action@v0`).
- Automated GitHub release publishing attaching binary tarball and `sbom.spdx.json` via `softprops/action-gh-release@v2`.

## [Unreleased]
- Added `docs/decisions.md` to document architectural trade-offs.
- Added `docs/runbook.md` with explicit startup and diagnostic steps.
- Identified known limitations in README for transparency.
