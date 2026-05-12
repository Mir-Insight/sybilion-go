# Vanity URL — `go.sybilion.dev`

This directory is a tiny static site whose only job is to make `go get go.sybilion.dev/sybilion` work. The Go toolchain fetches `https://go.sybilion.dev/sybilion?go-get=1` and reads the `<meta name="go-import">` tag in [`index.html`](./index.html) to learn where to clone from. Browsers visiting the page get a `<meta http-equiv="refresh">` that bounces them to the docs.

## What this contains

- `index.html` — the vanity page with `go-import` and `go-source` meta tags.
- `_redirects` — every path serves `index.html` so the toolchain finds the tags regardless of subpath.

That's it. No build step.

## Deploy with Cloudflare Pages (recommended)

Cloudflare Pages is the lightest path: free tier, fast cold start, supports `_redirects`, integrates with Cloudflare DNS.

1. Install Wrangler if you don't have it: `npm i -g wrangler` (or use `npx`).
2. From the repo root, deploy:

   ```bash
   npx wrangler pages deploy vanity --project-name=sybilion-vanity
   ```

   On the first run Wrangler will ask you to log in and create the project. Subsequent deploys are one command.

3. Point the custom domain `go.sybilion.dev` at the Pages project:
   - In the Cloudflare dashboard go to **Pages -> sybilion-vanity -> Custom domains -> Set up a custom domain**.
   - Enter `go.sybilion.dev`. If `sybilion.dev` is already on Cloudflare DNS, the CNAME is added automatically.
   - If `sybilion.dev` lives at another registrar, add a `CNAME` record `go -> sybilion-vanity.pages.dev` there.

4. Verify:

   ```bash
   curl 'https://go.sybilion.dev/sybilion?go-get=1' | grep go-import
   ```

   You should see the `<meta name="go-import" ...>` line.

5. Try a real install from any clean directory:

   ```bash
   GOPROXY=direct go get go.sybilion.dev/sybilion@latest
   ```

   `GOPROXY=direct` bypasses caches so you see real-time resolution.

## Alternative hosts

Because the page is one static HTML file, anywhere that serves static files works:

- **Netlify** — `netlify deploy --dir=vanity --prod`. Honors `_redirects`.
- **Vercel** — `vercel deploy vanity --prod`. Use `vercel.json` if you need redirects.
- **GitHub Pages** — drop `index.html` in a `gh-pages` branch of any repo, point a CNAME at `go.sybilion.dev`. Note: GitHub Pages doesn't honor `_redirects`; rely on the path always serving `index.html` (404s do).

## When to redeploy

Only when:

- The vanity meta tag's target repo URL changes (e.g. you rename the GitHub repo).
- The body copy / docs link changes.

Tagging a new SDK release (`v0.2.0`, etc.) does **not** require redeploying the vanity page — `go get` only reads it once per major version to learn where the source repo is.
