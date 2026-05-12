# Developers Portal — Go SDK

Official Go client for the **Operational API**. Generated from OpenAPI; thin `devportalclient` wrapper in the module root, low-level types in `gen/`.

**Module:** `github.com/Mir-Insight/developers-portal-api-sdk-go`  
**Requires Go:** see [`go.mod`](./go.mod) (currently 1.25.x).

## Install

```bash
go get github.com/Mir-Insight/developers-portal-api-sdk-go@latest
```

## Quick use

```go
import (
  "context"
  "os"

  devportalclient "github.com/Mir-Insight/developers-portal-api-sdk-go"
)

func main() {
  ctx := context.Background()
  c := devportalclient.New(devportalclient.Options{
    Token: os.Getenv("OPERATIONAL_API_TOKEN"),
  })
  me, _, err := c.DefaultAPI().ApiV1MeGet(ctx).Execute()
  _ = me
  _ = err
}
```

| Variable | Purpose |
|----------|---------|
| **`OPERATIONAL_API_TOKEN`** | Convention for apps: Bearer token (`sk_ops_...` or Auth0). Pass into `Options.Token` (not read automatically by the SDK). |
| **`OPERATIONAL_API_BASE_URL`** | Optional API origin when `Options.BaseURL` is empty. |

**Default base URL** when unset: `devportalclient.DefaultPublicAPIBaseURL` in [`defaults_gen.go`](./defaults_gen.go), built from [`defaults.json`](./defaults.json) at generation time.

## Docs

- Human overview (all languages): [developers-portal-api `docs/SDK.md`](https://github.com/Mir-Insight/developers-portal-api/blob/master/docs/SDK.md) (paths may vary by branch).
- **LLM / codegen cheat sheet:** [`LLM_SDK_GUIDE.md`](./LLM_SDK_GUIDE.md).

## Runnable example

See [`example/README.md`](./example/README.md).

## Versioning

This module uses **semantic versioning** (`v0`, `v1`, …). The API server and SDK versions are **not** locked 1:1; see [`CHANGELOG.md`](./CHANGELOG.md) for breaking changes and compatibility notes.

### Migrating from the monorepo path `ops-api/sdks/go`

If you previously used:

```go
import devportalclient "ops-api/sdks/go"
```

with a `replace` in `go.mod`, switch to:

```go
import devportalclient "github.com/Mir-Insight/developers-portal-api-sdk-go"
```

Remove the `replace` directive and run `go get github.com/Mir-Insight/developers-portal-api-sdk-go@<version>`.

## Wire contract

The canonical OpenAPI document lives in [`developers-portal-api`](https://github.com/Mir-Insight/developers-portal-api). This repository **vendors a snapshot** at [`openapi/openapi.yaml`](./openapi/openapi.yaml). CI regenerates `gen/` and fails if the result does not match git (no drift).

Maintainers can refresh code from a local API checkout:

```bash
DEVELOPERS_PORTAL_API_ROOT=/path/to/developers-portal-api bash scripts/sync-from-api-repo.sh
```

## Changelog

See [`CHANGELOG.md`](./CHANGELOG.md).
