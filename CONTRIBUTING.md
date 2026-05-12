# Contributing

This repository is the source of the [`go.sybilion.dev/sybilion`](https://pkg.go.dev/go.sybilion.dev/sybilion) Go module.

## Layout

- Module root — hand-written wrapper exposed as `package sybilion` (`sybilion.New(...)`, `Forecasts().Wait`, ...).
- `api/` — OpenAPI-generated low-level client (`package sybilionapi`). Do not hand-edit.
- `defaults_gen.go` — produced by `go run ./cmd/embed-go-defaults` from `defaults.json`.
- `openapi/openapi.yaml` — vendored public OpenAPI snapshot.
- `vanity/` — static page that resolves `go.sybilion.dev/sybilion` for `go get` (Cloudflare Pages).

## Refresh from the upstream API repository (maintainers)

With a checkout of the upstream API repository next to this repo (set `DEVELOPERS_PORTAL_API_ROOT` to its path):

```bash
bash scripts/sync-from-api-repo.sh
go test ./...
```

Then open a PR with the diff.

## Regenerate locally (Docker)

```bash
make generate   # bash scripts/sdk-generate-go.sh
make verify     # regenerate + fail if git is dirty
```

## CI expectations

`make verify` must pass on the default branch before merging changes that touch `openapi/`, `defaults.json`, or generation-related scripts.

## Release

See [PUBLISHING.md](PUBLISHING.md) for the full publish-to-pkg.go.dev workflow including the vanity-URL deployment. Standard checklist:

1. Update `CHANGELOG.md`.
2. Run `make verify`.
3. Push a semver tag such as `v0.1.0`. The release workflow validates the tree and creates a GitHub release; `pkg.go.dev` discovers the new version once any user runs `go get`.
