# Instagram Username Monitor

A Go username monitor that checks configured Instagram usernames through GraphQL and reports availability signals.

## How It Works

- Loads target usernames, active sessions, and proxies from `./data`.
- Builds GraphQL username-check requests for every configured target.
- Sends checks through the proxy client with the selected number of goroutines.
- Keeps TLS connections open for the monitoring workers.
- Detects availability signals from GraphQL responses.

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
