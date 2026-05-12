#!/usr/bin/env bash
# Regenerate Go client + defaults, then fail if the tree differs from git (prevents drift vs OpenAPI).
set -euo pipefail
ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "${ROOT}"

bash scripts/sdk-generate-go.sh

if ! git rev-parse --verify HEAD >/dev/null 2>&1; then
  echo "sdk-verify-go: OK (no commits yet — skipping drift check; commit tracked files, then re-run)"
  exit 0
fi

# Ignore untracked files so optional local files do not fail verify; tracked drift still fails.
if [[ -n "$(git status --porcelain --untracked-files=no)" ]]; then
  echo "sdk-verify-go: tracked files differ after regeneration. Run 'bash scripts/sdk-generate-go.sh', commit," >&2
  echo "or revert edits under gen/ and defaults_gen.go." >&2
  git status --porcelain --untracked-files=no
  git diff --stat
  exit 1
fi

echo "sdk-verify-go: OK (matches openapi/openapi.yaml + defaults.json)"
