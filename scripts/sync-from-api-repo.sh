#!/usr/bin/env bash
# Copy the Go SDK slice, OpenAPI spec, and defaults from a developers-portal-api checkout,
# then rewrite imports to this module. Intended for maintainers (run from this repo root).
#
# Usage:
#   DEVELOPERS_PORTAL_API_ROOT=/path/to/developers-portal-api bash scripts/sync-from-api-repo.sh
#
# Default API root: sibling directory ../developers-portal-api
set -euo pipefail
ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
API="${DEVELOPERS_PORTAL_API_ROOT:-$(cd "${ROOT}/.." && pwd)/developers-portal-api}"
MOD="github.com/Mir-Insight/developers-portal-api-sdk-go"

if [[ ! -d "${API}/sdks/go" ]]; then
  echo "sync-from-api-repo: missing ${API}/sdks/go (set DEVELOPERS_PORTAL_API_ROOT)" >&2
  exit 1
fi

rsync -a --delete \
  --exclude '.git' \
  --exclude 'README.md' \
  --exclude 'CHANGELOG.md' \
  --exclude 'CONTRIBUTING.md' \
  --exclude 'LICENSE' \
  --exclude 'SECURITY.md' \
  --exclude 'Makefile' \
  --exclude '.gitignore' \
  --exclude '.github' \
  --exclude 'cmd' \
  --exclude 'scripts' \
  --exclude 'openapi' \
  "${API}/sdks/go/" "${ROOT}/"
mkdir -p "${ROOT}/openapi"
cp "${API}/internal/httpapi/openapi/openapi.yaml" "${ROOT}/openapi/openapi.yaml"
cp "${API}/sdks/defaults.json" "${ROOT}/defaults.json"

find "${ROOT}" -maxdepth 3 -name '*.go' -type f -exec sed -i "s|ops-api/sdks/go|${MOD}|g" {} +
sed -i "s|^module ops-api/sdks/go\$|module ${MOD}|" "${ROOT}/go.mod"
sed -i "s|^module ops-api/sdks/go/example\$|module ${MOD}/example|" "${ROOT}/example/go.mod"
sed -i "s|require ops-api/sdks/go v0.0.0|require ${MOD} v0.0.0|" "${ROOT}/example/go.mod"
sed -i "s|replace ops-api/sdks/go => ..|replace ${MOD} => ..|" "${ROOT}/example/go.mod"
sed -i "s|ops-api/sdks/go|${MOD}|g" "${ROOT}/example/main.go"
if [[ -d "${ROOT}/gen/docs" ]]; then
  find "${ROOT}/gen/docs" -name '*.md' -type f -exec sed -i "s|github.com/GIT_USER_ID/GIT_REPO_ID|${MOD}/gen|g" {} +
fi

(cd "${ROOT}" && go run ./cmd/embed-go-defaults)
gofmt -w "${ROOT}/defaults_gen.go"

echo "sync-from-api-repo: copied from ${API}; review diff and commit."
