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
MOD="go.sybilion.dev/sybilion"

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
  --exclude 'PUBLISHING.md' \
  --exclude 'Makefile' \
  --exclude 'go.mod' \
  --exclude 'go.sum' \
  --exclude '.gitignore' \
  --exclude '.github' \
  --exclude 'cmd' \
  --exclude 'scripts' \
  --exclude 'openapi' \
  --exclude 'vanity' \
  "${API}/sdks/go/" "${ROOT}/"
mkdir -p "${ROOT}/openapi"
cp "${API}/internal/httpapi/openapi/openapi.public.yaml" "${ROOT}/openapi/openapi.yaml"
cp "${API}/sdks/defaults.json" "${ROOT}/defaults.json"

# Rewrite any leftover monorepo import paths to the published module path.
find "${ROOT}" -maxdepth 3 -name '*.go' -type f -print0 \
  | xargs -0 perl -pi -e "s|ops-api/sdks/go|${MOD}|g"
perl -pi -e "s|ops-api/sdks/go|${MOD}|g" "${ROOT}/example/main.go"
if [[ -d "${ROOT}/api/docs" ]]; then
  find "${ROOT}/api/docs" -name '*.md' -type f -print0 \
    | xargs -0 perl -pi -e "s|github.com/GIT_USER_ID/GIT_REPO_ID|${MOD}/api|g"
fi

(cd "${ROOT}" && go run ./cmd/embed-go-defaults)
gofmt -w "${ROOT}/defaults_gen.go"

echo "sync-from-api-repo: copied from ${API}; review diff and commit."
