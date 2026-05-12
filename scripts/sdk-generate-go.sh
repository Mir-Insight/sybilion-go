#!/usr/bin/env bash
# Regenerate the Go client in gen/ from vendored openapi/openapi.yaml and refresh defaults_gen.go.
# Requires Docker.
set -euo pipefail
ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
SPEC="/local/openapi/openapi.yaml"
IMG="${OPENAPI_GENERATOR_IMAGE:-openapitools/openapi-generator-cli:v7.11.0}"
MOD="github.com/Mir-Insight/developers-portal-api-sdk-go"
UID_LOCAL="$(id -u)"
GID_LOCAL="$(id -g)"
GEN_USER="${GEN_USER:-${UID_LOCAL}:${GID_LOCAL}}"

run_gen() {
  local generator="$1" out_dir="$2"
  shift 2
  docker run --rm --user "${GEN_USER}" -v "${ROOT}:/local" "${IMG}" generate \
    -i "${SPEC}" -g "${generator}" -o "/local/${out_dir}" "$@"
}

mkdir -p "${ROOT}/gen"

run_gen go gen \
  --additional-properties=packageName=devportal,withGoMod=false,generateInterfaces=false,enumClassPrefix=true

if [[ -d "${ROOT}/gen/docs" ]]; then
  find "${ROOT}/gen/docs" -name '*.md' -type f -exec sed -i "s|github.com/GIT_USER_ID/GIT_REPO_ID|${MOD}/gen|g" {} +
fi

rm -rf "${ROOT}/gen/test" "${ROOT}/gen/.travis.yml" "${ROOT}/gen/git_push.sh" 2>/dev/null || true

if [[ ! -w "${ROOT}/gen/client.go" ]]; then
  docker run --rm -v "${ROOT}:/local" alpine:3.19 sh -c "chown -R ${UID_LOCAL}:${GID_LOCAL} /local/gen" || true
fi

(cd "${ROOT}" && go run ./cmd/embed-go-defaults)
gofmt -w "${ROOT}/defaults_gen.go"

echo "sdk-generate-go: complete."
