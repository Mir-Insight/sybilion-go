# Changelog

All notable changes to this Go SDK module are documented here.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

**Versioning:** SDK releases use SemVer (`v0`, `v1`, …) independently of the `developers-portal-api` server image or chart version. Breaking Go API changes bump the major version (or minor while `v0`).

## [Unreleased]

### Added

- Standalone repository `github.com/Mir-Insight/developers-portal-api-sdk-go` with vendored `openapi/openapi.yaml`, `defaults.json`, `scripts/sdk-generate-go.sh`, `scripts/sdk-verify-go.sh`, and `scripts/sync-from-api-repo.sh`.
- `cmd/embed-go-defaults` for `defaults_gen.go`.
- CI (test, vet, OpenAPI drift check) and release workflow on semver tags.

### Changed

- **Module path** from monorepo `ops-api/sdks/go` to `github.com/Mir-Insight/developers-portal-api-sdk-go` (see README migration).

## [0.1.0] - 2026-05-11

### Added

- Initial generated client (`gen/`) from OpenAPI 0.1.0.
- `devportalclient` wrapper with Bearer auth, `WaitForecast`, and pagination helpers.
