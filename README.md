# Instagram Autoclaimer

A Go username autoclaimer that monitors configured Instagram usernames through GraphQL and sends username-change requests when a target becomes available.

## How It Works

- Loads target usernames, active sessions, and proxies from `./data`.
- Builds GraphQL username-check requests for every configured target.
- Sends checks through the proxy client with the selected number of goroutines.
- Keeps TLS connections open for the claim workers.
- Switches the claim workers from the dummy request to the detected username's prebuilt request when the GraphQL response signals availability.

## Files

Create the following files locally. The `data` directory is ignored by Git and must not be committed.

- `data/usernames.txt` — one `username:id` record per line.
- `data/active_sessions.txt` — active-session records using the format expected by the application.
- `data/proxies.txt` — one proxy per line.

## Run

```bash
go run .
```

Enter the number of monitor goroutines when prompted.

## Public Source

This repository contains the GraphQL monitoring and claim flow. The bulk-route monitor, bulk endpoints, banned-session handling, runtime data, and credentials are intentionally not included.
