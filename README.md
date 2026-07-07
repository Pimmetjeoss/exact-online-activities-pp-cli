# Exact Online Activities CLI

Activities-only Exact Online REST/OData API surface generated from the official Exact Online REST API documentation. Covers the Activities service resources and all documented methods. OAuth bearer token required.

Learn more at [Exact Online Activities](https://start.exactonline.nl/docs/HlpRestAPIResources.aspx?SourceAction=10).

Printed by [@Pimmetjeoss](https://github.com/Pimmetjeoss) (Pimmetjeoss).

## Install

The recommended path installs both the `exact-online-activities-pp-cli` binary and the `pp-exact-online-activities` agent skill in one shot:

```bash
npx -y @mvanhorn/printing-press install exact-online-activities
```

For CLI only (no skill):

```bash
npx -y @mvanhorn/printing-press install exact-online-activities --cli-only
```


### Without Node

The generated install path is category-agnostic until this CLI is published. If `npx` is not available before publish, install Node or use the category-specific Go fallback from the public-library entry after publish.

### Pre-built binary

Download a pre-built binary for your platform from the [latest release](https://github.com/mvanhorn/printing-press-library/releases/tag/exact-online-activities-current). On macOS, clear the Gatekeeper quarantine: `xattr -d com.apple.quarantine <binary>`. On Unix, mark it executable: `chmod +x <binary>`.

<!-- pp-hermes-install-anchor -->
## Install for Hermes

From the Hermes CLI:

```bash
hermes skills install mvanhorn/printing-press-library/cli-skills/pp-exact-online-activities --force
```

Inside a Hermes chat session:

```bash
/skills install mvanhorn/printing-press-library/cli-skills/pp-exact-online-activities --force
```

## Install for OpenClaw

Tell your OpenClaw agent (copy this):

```
Install the pp-exact-online-activities skill from https://github.com/mvanhorn/printing-press-library/tree/main/cli-skills/pp-exact-online-activities. The skill defines how its required CLI can be installed.
```

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


## Output Formats

```bash
# Human-readable table (default in terminal, JSON when piped)
exact-online-activities-pp-cli activities annual-statements-get mock-value

# JSON for scripting and agents
exact-online-activities-pp-cli activities annual-statements-get mock-value --json

# Filter to specific fields
exact-online-activities-pp-cli activities annual-statements-get mock-value --json --select id,name,status

# Dry run — show the request without sending
exact-online-activities-pp-cli activities annual-statements-get mock-value --dry-run

# Agent mode — JSON + compact + no prompts in one flag
exact-online-activities-pp-cli activities annual-statements-get mock-value --agent
```

## Agent Usage

This CLI is designed for AI agent consumption:

- **Non-interactive** - never prompts, every input is a flag
- **Pipeable** - `--json` output to stdout, errors to stderr
- **Filterable** - `--select id,name` returns only fields you need
- **Previewable** - `--dry-run` shows the request without sending
- **Explicit retries** - add `--idempotent` to create retries when a no-op success is acceptable
- **Confirmable** - `--yes` for explicit confirmation of destructive actions
- **Piped input** - write commands can accept structured input when their help lists `--stdin`
- **Offline-friendly** - sync/search commands can use the local SQLite store when available
- **Agent-safe by default** - no colors or formatting unless `--human-friendly` is set

Exit codes: `0` success, `2` usage error, `3` not found, `4` auth error, `5` API error, `7` rate limited, `10` config error.

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

## Health Check

```bash
exact-online-activities-pp-cli doctor
```

Verifies configuration, credentials, and connectivity to the API.

## Configuration

Config file: `~/.config/exact-online-activities-pp-cli/config.toml`

Static request headers can be configured under `headers`; per-command header overrides take precedence.

Environment variables:

| Name | Kind | Required | Description |
| --- | --- | --- | --- |
| `EXACT_ONLINE_ACTIVITIES_OAUTH2` | per_call | Yes | Set to your API credential. |

## Troubleshooting
**Authentication errors (exit code 4)**
- Run `exact-online-activities-pp-cli doctor` to check credentials
- Verify the environment variable is set: `echo $EXACT_ONLINE_ACTIVITIES_OAUTH2`
**Not found errors (exit code 3)**
- Check the resource ID is correct
- Run the `list` command to see available items

---

Generated by [CLI Printing Press](https://github.com/mvanhorn/cli-printing-press)
