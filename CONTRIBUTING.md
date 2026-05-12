# Contributing

## Generated code

- **`gen/`** is produced by OpenAPI Generator. Do not hand-edit; change the API spec in [`developers-portal-api`](https://github.com/Mir-Insight/developers-portal-api), then refresh this repo.
- **`defaults_gen.go`** is produced by `go run ./cmd/embed-go-defaults` from **`defaults.json`**.

## Refresh from the API monorepo

With a checkout of `developers-portal-api` next to this repo (or set `DEVELOPERS_PORTAL_API_ROOT`):

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
