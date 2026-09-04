# Exact Online Activities CLI

Activities-only Exact Online REST/OData API surface generated from the official Exact Online REST API documentation. Covers the Activities service resources and all documented methods. OAuth bearer token required.

Learn more at [Exact Online Activities](https://start.exactonline.nl/docs/HlpRestAPIResources.aspx?SourceAction=10).

<<<<<<< HEAD
## Install

The recommended path installs both the `exact-online-activities-pp-cli` binary and the `pp-exact-online-activities` agent skill (Claude Code, Codex, Cursor, Gemini CLI, GitHub Copilot, and other agents supported by the upstream [`skills`](https://github.com/vercel-labs/skills) CLI) in one shot:

```bash
npx -y @mvanhorn/printing-press-library install exact-online-activities
=======
Printed by [@Pimmetjeoss](https://github.com/Pimmetjeoss) (Pimmetjeoss).

## Install

The recommended path installs both the `exact-online-activities-pp-cli` binary and the `pp-exact-online-activities` agent skill in one shot:

```bash
npx -y @mvanhorn/printing-press install exact-online-activities
>>>>>>> origin/main
```

For CLI only (no skill):

```bash
<<<<<<< HEAD
npx -y @mvanhorn/printing-press-library install exact-online-activities --cli-only
```

For skill only — installs the skill into the same agents as the default command above, but skips the CLI binary (use this to update or reinstall just the skill):

```bash
npx -y @mvanhorn/printing-press-library install exact-online-activities --skill-only
```

To constrain the skill install to one or more specific agents (repeatable — agent names match the [`skills`](https://github.com/vercel-labs/skills) CLI):

```bash
npx -y @mvanhorn/printing-press-library install exact-online-activities --agent claude-code
npx -y @mvanhorn/printing-press-library install exact-online-activities --agent claude-code --agent codex
```
=======
npx -y @mvanhorn/printing-press install exact-online-activities --cli-only
```

>>>>>>> origin/main

### Without Node

The generated install path is category-agnostic until this CLI is published. If `npx` is not available before publish, install Node or use the category-specific Go fallback from the public-library entry after publish.

### Pre-built binary

Download a pre-built binary for your platform from the [latest release](https://github.com/mvanhorn/printing-press-library/releases/tag/exact-online-activities-current). On macOS, clear the Gatekeeper quarantine: `xattr -d com.apple.quarantine <binary>`. On Unix, mark it executable: `chmod +x <binary>`.

<!-- pp-hermes-install-anchor -->
## Install for Hermes

<<<<<<< HEAD
Install the CLI binary first. The installer writes binaries to a per-user managed bin directory by default: `$HOME/.local/bin` on macOS/Linux and `%LOCALAPPDATA%\Programs\PrintingPress\bin` on Windows.

```bash
npx -y @mvanhorn/printing-press-library install exact-online-activities --cli-only
```

Then install the focused Hermes skill.

=======
>>>>>>> origin/main
From the Hermes CLI:

```bash
hermes skills install mvanhorn/printing-press-library/cli-skills/pp-exact-online-activities --force
```

Inside a Hermes chat session:

```bash
/skills install mvanhorn/printing-press-library/cli-skills/pp-exact-online-activities --force
```

<<<<<<< HEAD
Restart the Hermes session or gateway if the newly installed skill is not visible immediately.

## Install for OpenClaw
Install both the CLI binary and the focused OpenClaw skill. The installer defaults binaries to a per-user bin directory (`$HOME/.local/bin` on macOS/Linux, `%LOCALAPPDATA%\Programs\PrintingPress\bin` on Windows):

```bash
npx -y @mvanhorn/printing-press-library install exact-online-activities --agent openclaw
```

Restart the OpenClaw session or gateway if the newly installed skill is not visible immediately.

## Use with Claude Desktop

This CLI ships an [MCPB](https://github.com/modelcontextprotocol/mcpb) bundle — Claude Desktop's standard format for one-click MCP extension installs (no JSON config required).

To install:

1. Download the `.mcpb` for your platform from the [latest release](https://github.com/mvanhorn/printing-press-library/releases/tag/exact-online-activities-current).
2. Double-click the `.mcpb` file. Claude Desktop opens and walks you through the install.
3. Fill in `EXACT_ONLINE_ACTIVITIES_OAUTH2` when Claude Desktop prompts you.

Requires Claude Desktop 1.0.0 or later. Pre-built bundles ship for macOS Apple Silicon (`darwin-arm64`) and Windows (`amd64`, `arm64`); for other platforms, use the manual config below.

<details>
<summary>Manual JSON config (advanced)</summary>

If you can't use the MCPB bundle (older Claude Desktop, unsupported platform), install the MCP binary and configure it manually.


Install the MCP binary from this CLI's published public-library entry or pre-built release.

Add to your Claude Desktop config (`~/Library/Application Support/Claude/claude_desktop_config.json`):

```json
{
  "mcpServers": {
    "exact-online-activities": {
      "command": "exact-online-activities-pp-mcp",
      "env": {
        "EXACT_ONLINE_ACTIVITIES_OAUTH2": "<your-key>"
      }
    }
  }
}
```

</details>

=======
## Install for OpenClaw

Tell your OpenClaw agent (copy this):

```
Install the pp-exact-online-activities skill from https://github.com/mvanhorn/printing-press-library/tree/main/cli-skills/pp-exact-online-activities. The skill defines how its required CLI can be installed.
```

>>>>>>> origin/main
## Quick Start

### 1. Install

See [Install](#install) above.

### 2. Set Up Credentials

Get your access token from your API provider's developer portal, then store it:

```bash
exact-online-activities-pp-cli auth set-token YOUR_TOKEN_HERE
```

Or set it via environment variable:

```bash
export EXACT_ONLINE_ACTIVITIES_OAUTH2="your-token-here"
```

### 3. Verify Setup

```bash
exact-online-activities-pp-cli doctor
```

This checks your configuration and credentials.

### 4. Try Your First Command

```bash
exact-online-activities-pp-cli activities annual-statements-get mock-value
```

## Usage

Run `exact-online-activities-pp-cli --help` for the full command reference and flag list.

<<<<<<< HEAD
## Paths & environment variables

This CLI separates local files into four path kinds:

| Kind | Contents |
|------|----------|
| `config` | User-editable settings such as `config.toml` and saved profiles |
| `data` | Durable local data: `credentials.toml`, `data.db`, cookies, browser-session proof files, and other auth sidecars |
| `state` | Runtime state such as persisted queries, jobs, and `teach.log` |
| `cache` | Regenerable HTTP/cache files |

Each kind resolves independently. The ladder is:

1. Per-kind env var: `EXACT_ONLINE_ACTIVITIES_CONFIG_DIR`, `EXACT_ONLINE_ACTIVITIES_DATA_DIR`, `EXACT_ONLINE_ACTIVITIES_STATE_DIR`, or `EXACT_ONLINE_ACTIVITIES_CACHE_DIR`
2. `--home <dir>` for this invocation
3. `EXACT_ONLINE_ACTIVITIES_HOME` for a flat relocated root
4. XDG env vars: `XDG_CONFIG_HOME`, `XDG_DATA_HOME`, `XDG_STATE_HOME`, `XDG_CACHE_HOME`
5. Platform defaults matching existing installs

For containers and agent sandboxes, prefer a single relocated root:

```bash
export EXACT_ONLINE_ACTIVITIES_HOME=/srv/exact-online-activities
exact-online-activities-pp-cli doctor
```

Under `EXACT_ONLINE_ACTIVITIES_HOME=/srv/exact-online-activities`, the four dirs resolve to `/srv/exact-online-activities/config`, `/srv/exact-online-activities/data`, `/srv/exact-online-activities/state`, and `/srv/exact-online-activities/cache`.

MCP servers do not receive CLI flags from the host. Put relocation in the host `env` block:

```json
{
  "mcpServers": {
    "exact-online-activities": {
      "command": "exact-online-activities-pp-mcp",
      "env": {
        "EXACT_ONLINE_ACTIVITIES_HOME": "/srv/exact-online-activities"
      }
    }
  }
}
```

Precedence matters in fleets: an ambient per-kind variable such as `EXACT_ONLINE_ACTIVITIES_DATA_DIR` overrides an explicit `--home` for that kind. Use `EXACT_ONLINE_ACTIVITIES_HOME` or the per-kind variables for durable fleet relocation; treat `--home` as the weaker per-invocation lever.

Relocation is one-way. Unsetting `EXACT_ONLINE_ACTIVITIES_HOME` does not move files back to platform defaults, and `doctor` cannot find credentials left under a former root. Move the files manually before unsetting relocation variables.

Existing installs keep working because the platform-default rung matches the legacy layout. On the first auth write, stored secrets leave `config.toml` and are consolidated into `credentials.toml` under the data directory. Run `exact-online-activities-pp-cli doctor --fail-on warn` to check path and credential-location warnings in automation.

=======
>>>>>>> origin/main
## Commands

### activities

Manage activities

- **`exact-online-activities-pp-cli activities annual-statements-get`** - The date indicating by when the action has to be taken; The status of the Annual Statement request.

Official docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ActivitiesAnnualStatements

Scope: Organization workflow
- **`exact-online-activities-pp-cli activities annual-statements-post`** - The date indicating by when the action has to be taken; The status of the Annual Statement request.

Official docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ActivitiesAnnualStatements

Scope: Organization workflow
- **`exact-online-activities-pp-cli activities annual-statements-put`** - The date indicating by when the action has to be taken; The status of the Annual Statement request.

Official docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ActivitiesAnnualStatements

Scope: Organization workflow
- **`exact-online-activities-pp-cli activities communication-notes-get`** - The account that is related to the communication note; The name of the account.

Official docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ActivitiesCommunicationNotes

Scope: Organization workflow
- **`exact-online-activities-pp-cli activities communication-notes-post`** - The account that is related to the communication note; The name of the account.

Official docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ActivitiesCommunicationNotes

Scope: Organization workflow
- **`exact-online-activities-pp-cli activities complaints-get`** - The account that is related to the complaint; The name of the account.

Official docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ActivitiesComplaints

Scope: Organization workflow
- **`exact-online-activities-pp-cli activities complaints-post`** - The account that is related to the complaint; The name of the account.

Official docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ActivitiesComplaints

Scope: Organization workflow
- **`exact-online-activities-pp-cli activities events-get`** - The account that is related to the event; The name of the account.

Official docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ActivitiesEvents

Scope: Organization workflow
- **`exact-online-activities-pp-cli activities events-post`** - The account that is related to the event; The name of the account.

Official docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ActivitiesEvents

Scope: Organization workflow
- **`exact-online-activities-pp-cli activities fiscals-get`** - The date indicating by when the action has to be taken; The user that the request is assigned to.

Official docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ActivitiesFiscals

Scope: Organization workflow
- **`exact-online-activities-pp-cli activities fiscals-post`** - The date indicating by when the action has to be taken; The user that the request is assigned to.

Official docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ActivitiesFiscals

Scope: Organization workflow
- **`exact-online-activities-pp-cli activities fiscals-put`** - The date indicating by when the action has to be taken; The user that the request is assigned to.

Official docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ActivitiesFiscals

Scope: Organization workflow
- **`exact-online-activities-pp-cli activities service-requests-get`** - The account that is related to the service request; The name of the account.

Official docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ActivitiesServiceRequests

Scope: Organization workflow
- **`exact-online-activities-pp-cli activities service-requests-post`** - The account that is related to the service request; The name of the account.

Official docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ActivitiesServiceRequests

Scope: Organization workflow
- **`exact-online-activities-pp-cli activities tasks-get`** - The account that is related to the task; The name of the account.

Official docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ActivitiesTasks

Scope: Organization workflow
- **`exact-online-activities-pp-cli activities tasks-post`** - The account that is related to the task; The name of the account.

Official docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=ActivitiesTasks

Scope: Organization workflow


<<<<<<< HEAD
### Self-learning loop

This CLI caches per-question discovery so repeat queries skip the walk and structurally similar queries get answered via entity substitution. The loop also self-captures: every invocation is journaled locally, and failed-flag corrections plus fresh teaches surface as candidates on the next `recall` for confirm/reject judgment. Agents call `recall` before discovery and fire `teach &` after answering. See the `## Automatic learning` section in `SKILL.md` for the full protocol.

- **`exact-online-activities-pp-cli recall <query>`** - Look up cached resources for a query before running discovery
- **`exact-online-activities-pp-cli teach`** - Record a query -> resource mapping (silent on success, safe to background with `&`)
- **`exact-online-activities-pp-cli learnings list`** - Inspect taught rows
- **`exact-online-activities-pp-cli learnings forget <query>`** - Undo a teach
- **`exact-online-activities-pp-cli learnings candidates`** - List auto-captured candidates awaiting confirm/reject
- **`exact-online-activities-pp-cli learnings stats`** - Local loop metrics: recall hit rate, teach-to-reuse, playbook resolution, candidate counts
- **`exact-online-activities-pp-cli teach-pattern`** - Install a query/resource template up front
- **`exact-online-activities-pp-cli teach-lookup`** - Add an entity mapping (e.g. country code, team alias) for pattern substitution

Pass `--no-learn` or set `EXACT_ONLINE_ACTIVITIES_NO_LEARN=true` to disable the loop for deterministic flows.

The local store's schema version stamp is one-way: once this version of `exact-online-activities-pp-cli` opens the database, older binaries refuse it with a version error — upgrade the binary rather than downgrading.

=======
>>>>>>> origin/main
## Output Formats

```bash
# Human-readable table (default in terminal, JSON when piped)
exact-online-activities-pp-cli activities annual-statements-get mock-value

# JSON for scripting and agents
exact-online-activities-pp-cli activities annual-statements-get mock-value --json
<<<<<<< HEAD
# Filter to specific fields by name
exact-online-activities-pp-cli activities annual-statements-get mock-value --json --select <field>[,<field>...]
=======

# Filter to specific fields
exact-online-activities-pp-cli activities annual-statements-get mock-value --json --select id,name,status
>>>>>>> origin/main

# Dry run — show the request without sending
exact-online-activities-pp-cli activities annual-statements-get mock-value --dry-run

# Agent mode — JSON + compact + no prompts in one flag
exact-online-activities-pp-cli activities annual-statements-get mock-value --agent
```

## Agent Usage

This CLI is designed for AI agent consumption:

- **Non-interactive** - never prompts, every input is a flag
- **Pipeable** - `--json` output to stdout, errors to stderr
<<<<<<< HEAD
- **Filterable** - `--select <field>[,<field>...]` returns only fields you need
- **Previewable** - `--dry-run` shows the request without sending
- **Explicit retries** - add `--idempotent` to create retries when a no-op success is acceptable
- **Explicit confirmation** - `--agent` does not imply `--yes`; pass `--yes` separately only after the target, arguments, and side effects are clear
=======
- **Filterable** - `--select id,name` returns only fields you need
- **Previewable** - `--dry-run` shows the request without sending
- **Explicit retries** - add `--idempotent` to create retries when a no-op success is acceptable
- **Confirmable** - `--yes` for explicit confirmation of destructive actions
>>>>>>> origin/main
- **Piped input** - write commands can accept structured input when their help lists `--stdin`
- **Offline-friendly** - sync/search commands can use the local SQLite store when available
- **Agent-safe by default** - no colors or formatting unless `--human-friendly` is set

Exit codes: `0` success, `2` usage error, `3` not found, `4` auth error, `5` API error, `7` rate limited, `10` config error.

<<<<<<< HEAD
=======
## Use with Claude Code

Install the focused skill — it auto-installs the CLI on first invocation:

```bash
npx skills add mvanhorn/printing-press-library/cli-skills/pp-exact-online-activities -g
```

Then invoke `/pp-exact-online-activities <query>` in Claude Code. The skill is the most efficient path — Claude Code drives the CLI directly without an MCP server in the middle.

<details>
<summary>Use as an MCP server in Claude Code (advanced)</summary>

If you'd rather register this CLI as an MCP server in Claude Code, install the MCP binary first:


Install the MCP binary from this CLI's published public-library entry or pre-built release.

Then register it:

```bash
claude mcp add exact-online-activities exact-online-activities-pp-mcp -e EXACT_ONLINE_ACTIVITIES_OAUTH2=<your-token>
```

</details>

## Use with Claude Desktop

This CLI ships an [MCPB](https://github.com/modelcontextprotocol/mcpb) bundle — Claude Desktop's standard format for one-click MCP extension installs (no JSON config required).

To install:

1. Download the `.mcpb` for your platform from the [latest release](https://github.com/mvanhorn/printing-press-library/releases/tag/exact-online-activities-current).
2. Double-click the `.mcpb` file. Claude Desktop opens and walks you through the install.
3. Fill in `EXACT_ONLINE_ACTIVITIES_OAUTH2` when Claude Desktop prompts you.

Requires Claude Desktop 1.0.0 or later. Pre-built bundles ship for macOS Apple Silicon (`darwin-arm64`) and Windows (`amd64`, `arm64`); for other platforms, use the manual config below.

<details>
<summary>Manual JSON config (advanced)</summary>

If you can't use the MCPB bundle (older Claude Desktop, unsupported platform), install the MCP binary and configure it manually.


Install the MCP binary from this CLI's published public-library entry or pre-built release.

Add to your Claude Desktop config (`~/Library/Application Support/Claude/claude_desktop_config.json`):

```json
{
  "mcpServers": {
    "exact-online-activities": {
      "command": "exact-online-activities-pp-mcp",
      "env": {
        "EXACT_ONLINE_ACTIVITIES_OAUTH2": "<your-key>"
      }
    }
  }
}
```

</details>

>>>>>>> origin/main
## Health Check

```bash
exact-online-activities-pp-cli doctor
```

Verifies configuration, credentials, and connectivity to the API.

## Configuration

<<<<<<< HEAD
Run `exact-online-activities-pp-cli doctor` to see the resolved config, data, state, and cache directories. The platform-default config path is `~/.config/exact-online-activities-pp-cli/config.toml`; `--home`, `EXACT_ONLINE_ACTIVITIES_HOME`, and per-kind env vars can relocate it.
=======
Config file: `~/.config/exact-online-activities-pp-cli/config.toml`
>>>>>>> origin/main

Static request headers can be configured under `headers`; per-command header overrides take precedence.

Environment variables:

| Name | Kind | Required | Description |
| --- | --- | --- | --- |
| `EXACT_ONLINE_ACTIVITIES_OAUTH2` | per_call | Yes | Set to your API credential. |

<<<<<<< HEAD
### agentcookie (optional)

If you use agentcookie to sync secrets across machines, this CLI auto-adopts agentcookie-managed credentials with no extra setup. When the daemon writes to this CLI's config, `exact-online-activities-pp-cli doctor` reports `agentcookie: detected` and `auth-status` labels the source as `agentcookie`. Skip this section if you don't use agentcookie - the CLI works the same as any other.

=======
>>>>>>> origin/main
## Troubleshooting
**Authentication errors (exit code 4)**
- Run `exact-online-activities-pp-cli doctor` to check credentials
- Verify the environment variable is set: `echo $EXACT_ONLINE_ACTIVITIES_OAUTH2`
**Not found errors (exit code 3)**
- Check the resource ID is correct
- Run the `list` command to see available items

---

Generated by [CLI Printing Press](https://github.com/mvanhorn/cli-printing-press)
