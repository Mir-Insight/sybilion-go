# Generated Go client (`gen/`)

This tree is produced by **OpenAPI Generator** from the Operational API spec. It is a low-level client (`DefaultApi`, models, `Configuration`).

**Do not follow generic “installation” or `go get github.com/GIT_USER_ID/...` instructions** — those are generator placeholders.

**Do not copy import paths or setup from `docs/*.md` here without checking them** — markdown under `docs/` is regenerated each run; this repo post-processes placeholder GitHub paths to **`ops-api/sdks/go/gen`**, but examples still use the raw generated client (`NewAPIClient`, `DefaultAPI`) instead of the recommended **`devportalclient.New`** wrapper. Prefer:

- **[`../README.md`](../README.md)** — import path `ops-api/sdks/go`, auth, base URL, pagination, forecast polling.
- **[`../LLM_SDK_GUIDE.md`](../LLM_SDK_GUIDE.md)** — copy-paste snippets for LLMs and external apps.

**Wire contract snapshot shipped with this SDK:** [`api/openapi.yaml`](./api/openapi.yaml) (regenerated with `gen/`).

Regenerate from the **monorepo** root (Docker):

```bash
bash scripts/sdk-generate.sh
```
