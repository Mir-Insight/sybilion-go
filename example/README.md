# Sybilion Go SDK example (local consumer)

Minimal app that calls **`GET /api/v1/me`** through [`go.sybilion.dev/sybilion`](../) using a `replace` directive (same pattern you can use from another repo until you depend on a published tag).

## Prerequisites

- **Go 1.25+** — same as the parent SDK [`go.mod`](../go.mod).
- A valid **Bearer token**: API key `sk_ops_...` (create one in the Developers Portal) or a dashboard session token.

## Run

From this directory:

```bash
cp .env.example .env
# edit .env — set SYBILION_API_TOKEN (required)

set -a && source .env && set +a && go run .
```

Optional: set **`SYBILION_API_BASE_URL`** in `.env` to hit a non-production host (for example `http://127.0.0.1:8080`). If unset, the SDK reads **`SYBILION_API_BASE_URL`** from the environment, then falls back to the compiled **`DefaultPublicAPIBaseURL`** in the parent module's **`defaults_gen.go`**.

## Use from another module (same machine)

In your app's `go.mod`:

```go
require go.sybilion.dev/sybilion v0.0.0

replace go.sybilion.dev/sybilion => /absolute/path/to/sybilion-go
```

Use a pseudo-version if plain `v0.0.0` is rejected by your toolchain.

## Use from a published tag

```bash
go get go.sybilion.dev/sybilion@v0.1.0
```

```go
import "go.sybilion.dev/sybilion"
```
