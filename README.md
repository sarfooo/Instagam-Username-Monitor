# GraphQL Monitor Example

A trimmed, public-facing Go example that demonstrates the GraphQL monitoring request builders from the original project. It intentionally excludes bulk-route monitoring, credentials, webhooks, runtime data, and account-changing logic.

Create `data/usernames.txt` with `username:id` lines and `data/proxies.txt` with one HTTP proxy per line, then run `go run .`.
