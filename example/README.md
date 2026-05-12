# Go SDK example (local consumer)

Minimal app that calls **`GET /api/v1/me`** through [`github.com/Mir-Insight/developers-portal-api-sdk-go`](../) using a `replace` directive (same pattern you can use from another repo until you depend on a published tag).

## Prerequisites

- **Go 1.25+** — same as the parent SDK [`go.mod`](../go.mod).
- A valid **Bearer token**: API key `sk_ops_...` (create via `POST /api/v1/api-keys` when logged into the developer product, or use the dashboard if your deployment exposes one) or an Auth0 access token for the API audience.

## Run

From this directory:

```bash
cp .env.example .env
# edit .env — set OPERATIONAL_API_TOKEN (required)

set -a && source .env && set +a && go run .
```

Optional: set **`OPERATIONAL_API_BASE_URL`** in `.env` to hit a non-production host (for example `http://127.0.0.1:8080`). If unset, the SDK reads **`OPERATIONAL_API_BASE_URL`** from the environment (if set), then the compiled **`DefaultPublicAPIBaseURL`** in the parent module’s **`defaults_gen.go`**.

## Use from another module (same machine)

In your app’s `go.mod`:

```go
require github.com/Mir-Insight/developers-portal-api-sdk-go v0.0.0

replace github.com/Mir-Insight/developers-portal-api-sdk-go => /absolute/path/to/developers-portal-api-sdk-go
```

Use a pseudo-version if plain `v0.0.0` is rejected by your toolchain.
