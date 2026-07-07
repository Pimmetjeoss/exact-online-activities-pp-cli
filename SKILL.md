---
name: pp-exact-online-activities
description: "Printing Press CLI for Exact Online Activities. Activities-only Exact Online REST/OData API surface generated from the official Exact Online REST API documentation...."
author: "Pimmetjeoss"
license: "Apache-2.0"
argument-hint: "<command> [args] | install cli|mcp"
allowed-tools: "Read Bash"
metadata:
  openclaw:
    requires:
      bins:
        - exact-online-activities-pp-cli
---

# Exact Online Activities — Printing Press CLI

## Prerequisites: Install the CLI

This skill drives the `exact-online-activities-pp-cli` binary. **You must verify the CLI is installed before invoking any command from this skill.** If it is missing, install it first:

1. Install via the Printing Press installer:
   ```bash
   npx -y @mvanhorn/printing-press install exact-online-activities --cli-only
   ```
2. Verify: `exact-online-activities-pp-cli --version`
3. Ensure `$GOPATH/bin` (or `$HOME/go/bin`) is on `$PATH`.

If the `npx` install fails before this CLI has a public-library category, install Node or use the category-specific Go fallback after publish.

If `--version` reports "command not found" after install, the install step did not put the binary on `$PATH`. Do not proceed with skill commands until verification succeeds.

Activities-only Exact Online REST/OData API surface generated from the official Exact Online REST API documentation. Covers the Activities service resources and all documented methods. OAuth bearer token required.

## Command Reference

**activities** — Manage activities

- `exact-online-activities-pp-cli activities annual-statements-get` — The date indicating by when the action has to be taken; The status of the Annual Statement request. Official docs:...
- `exact-online-activities-pp-cli activities annual-statements-post` — The date indicating by when the action has to be taken; The status of the Annual Statement request. Official docs:...
- `exact-online-activities-pp-cli activities annual-statements-put` — The date indicating by when the action has to be taken; The status of the Annual Statement request. Official docs:...
- `exact-online-activities-pp-cli activities communication-notes-get` — The account that is related to the communication note; The name of the account. Official docs:...
- `exact-online-activities-pp-cli activities communication-notes-post` — The account that is related to the communication note; The name of the account. Official docs:...
- `exact-online-activities-pp-cli activities complaints-get` — The account that is related to the complaint; The name of the account. Official docs:...
- `exact-online-activities-pp-cli activities complaints-post` — The account that is related to the complaint; The name of the account. Official docs:...
- `exact-online-activities-pp-cli activities events-get` — The account that is related to the event; The name of the account. Official docs:...
- `exact-online-activities-pp-cli activities events-post` — The account that is related to the event; The name of the account. Official docs:...
- `exact-online-activities-pp-cli activities fiscals-get` — The date indicating by when the action has to be taken; The user that the request is assigned to. Official docs:...
- `exact-online-activities-pp-cli activities fiscals-post` — The date indicating by when the action has to be taken; The user that the request is assigned to. Official docs:...
- `exact-online-activities-pp-cli activities fiscals-put` — The date indicating by when the action has to be taken; The user that the request is assigned to. Official docs:...
- `exact-online-activities-pp-cli activities service-requests-get` — The account that is related to the service request; The name of the account. Official docs:...
- `exact-online-activities-pp-cli activities service-requests-post` — The account that is related to the service request; The name of the account. Official docs:...
- `exact-online-activities-pp-cli activities tasks-get` — The account that is related to the task; The name of the account. Official docs:...
- `exact-online-activities-pp-cli activities tasks-post` — The account that is related to the task; The name of the account. Official docs:...


### Finding the right command

When you know what you want to do but not which command does it, ask the CLI directly:

```bash
exact-online-activities-pp-cli which "<capability in your own words>"
```

`which` resolves a natural-language capability query to the best matching command from this CLI's curated feature index. Exit code `0` means at least one match; exit code `2` means no confident match — fall back to `--help` or use a narrower query.

## Auth Setup

Run `exact-online-activities-pp-cli auth setup` for the URL and steps to obtain a token (add `--launch` to open the URL). Then store it:

```bash
exact-online-activities-pp-cli auth set-token YOUR_TOKEN_HERE
```

Or set `EXACT_ONLINE_ACTIVITIES_OAUTH2` as an environment variable.

Run `exact-online-activities-pp-cli doctor` to verify setup.

## Agent Mode

Add `--agent` to any command. Expands to: `--json --compact --no-input --no-color --yes`.

- **Pipeable** — JSON on stdout, errors on stderr
- **Filterable** — `--select` keeps a subset of fields. Dotted paths descend into nested structures; arrays traverse element-wise. Critical for keeping context small on verbose APIs:

  ```bash
  exact-online-activities-pp-cli activities annual-statements-get mock-value --agent --select id,name,status
  ```
- **Previewable** — `--dry-run` shows the request without sending
- **Offline-friendly** — sync/search commands can use the local SQLite store when available
- **Non-interactive** — never prompts, every input is a flag
- **Explicit retries** — use `--idempotent` only when an already-existing create should count as success

### Response envelope

Commands that read from the local store or the API wrap output in a provenance envelope:

```json
{
  "meta": {"source": "live" | "local", "synced_at": "...", "reason": "..."},
  "results": <data>
}
```

Parse `.results` for data and `.meta.source` to know whether it's live or local. A human-readable `N results (live)` summary is printed to stderr only when stdout is a terminal — piped/agent consumers get pure JSON on stdout.

## Agent Feedback

When you (or the agent) notice something off about this CLI, record it:

```
exact-online-activities-pp-cli feedback "the --since flag is inclusive but docs say exclusive"
exact-online-activities-pp-cli feedback --stdin < notes.txt
exact-online-activities-pp-cli feedback list --json --limit 10
```

Entries are stored locally at `~/.exact-online-activities-pp-cli/feedback.jsonl`. They are never POSTed unless `EXACT_ONLINE_ACTIVITIES_FEEDBACK_ENDPOINT` is set AND either `--send` is passed or `EXACT_ONLINE_ACTIVITIES_FEEDBACK_AUTO_SEND=true`. Default behavior is local-only.

Write what *surprised* you, not a bug report. Short, specific, one line: that is the part that compounds.

## Output Delivery

Every command accepts `--deliver <sink>`. The output goes to the named sink in addition to (or instead of) stdout, so agents can route command results without hand-piping. Three sinks are supported:

| Sink | Effect |
|------|--------|
| `stdout` | Default; write to stdout only |
| `file:<path>` | Atomically write output to `<path>` (tmp + rename) |
| `webhook:<url>` | POST the output body to the URL (`application/json` or `application/x-ndjson` when `--compact`) |

Unknown schemes are refused with a structured error naming the supported set. Webhook failures return non-zero and log the URL + HTTP status on stderr.

## Named Profiles

A profile is a saved set of flag values, reused across invocations. Use it when a scheduled agent calls the same command every run with the same configuration - HeyGen's "Beacon" pattern.

```
exact-online-activities-pp-cli profile save briefing --json
exact-online-activities-pp-cli --profile briefing activities annual-statements-get mock-value
exact-online-activities-pp-cli profile list --json
exact-online-activities-pp-cli profile show briefing
exact-online-activities-pp-cli profile delete briefing --yes
```

Explicit flags always win over profile values; profile values win over defaults. `agent-context` lists all available profiles under `available_profiles` so introspecting agents discover them at runtime.

## Exit Codes

| Code | Meaning |
|------|---------|
| 0 | Success |
| 2 | Usage error (wrong arguments) |
| 3 | Resource not found |
| 4 | Authentication required |
| 5 | API error (upstream issue) |
| 7 | Rate limited (wait and retry) |
| 10 | Config error |

## Argument Parsing

Parse `$ARGUMENTS`:

1. **Empty, `help`, or `--help`** → show `exact-online-activities-pp-cli --help` output
2. **Starts with `install`** → ends with `mcp` → MCP installation; otherwise → see Prerequisites above
3. **Anything else** → Direct Use (execute as CLI command with `--agent`)

## MCP Server Installation

Install the MCP binary from this CLI's published public-library entry or pre-built release, then register it:

```bash
claude mcp add exact-online-activities-pp-mcp -- exact-online-activities-pp-mcp
```

Verify: `claude mcp list`

## Direct Use

1. Check if installed: `which exact-online-activities-pp-cli`
   If not found, offer to install (see Prerequisites at the top of this skill).
2. Match the user query to the best command from the Unique Capabilities and Command Reference above.
3. Execute with the `--agent` flag:
   ```bash
   exact-online-activities-pp-cli <command> [subcommand] [args] --agent
   ```
4. If ambiguous, drill into subcommand help: `exact-online-activities-pp-cli <command> --help`.
