# Changelog for golden-path-cli

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
