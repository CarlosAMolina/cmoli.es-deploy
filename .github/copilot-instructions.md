# Copilot Instructions for cmoli.es-deploy

## Project Overview

**cmoli.es-deploy** is a Go CLI tool that automates the website deployment pipeline for [cmoli.es](https://cmoli.es). It orchestrates multiple repositories and conversion processes to create and publish website content to a VPS.

### Core Workflow
1. **Pull Git repositories** - Fetches latest from multiple source repos (cmoli.es, checkIframe, wiki, tools)
2. **Prepare content** - Consolidates markdown/assets from various repos into a staging directory (`cfg.WebContentPath`, default `/tmp/www`)
3. **Convert markdown to HTML** - Delegates to the external Go project `md-to-html-go` by passing `cfg.WebContentPath` as the target directory
4. **Post-process HTML** - Applies specific modifications (e.g., CSS classes) to generated HTML
5. **Create media symlinks** - Links media files from `cfg.MediaContentPath` to match markdown structure
6. **Send to VPS** - Syncs content to VPS using rsync

### Build & Run

**Commands (via makefile):**
- `make run` - Run the CLI interactively with menu options
- `make build` - Cross-compile to Linux (CGO disabled to avoid glibc version conflicts on older VPS systems)
- `make format` - Run `go fmt`
- `make dependencies` - Run `go mod tidy`

**No tests exist** - Project is straightforward CLI orchestration without test files.

## Architecture & Key Files

### Module Structure (Single Package)
All code is in the `main` package. Each responsibility is split across files:

| File | Purpose |
|------|---------|
| `main.go` | CLI entry point, menu system, orchestration of deployment steps |
| `config.go` | Loads `config-deploy.json` into global `cfg` variable |
| `git.go` | Pulls from multiple repos (cmoli.es, checkIframe, wiki, tools, md-to-html-go) |
| `md_to_html.go` | Copies content from source repos to staging dir, invokes `md-to-html-go` CLI |
| `modify_html.go` | Post-processes HTML (currently only modifies Rust comparison table styling) |
| `vps.go` | Rsync content to VPS |
| `path_utils.go` | Path helpers (`getPathSoftware()`, `getCurrentPath()`) and media symlink creation |
| `command.go` | Bash command execution wrapper |

### Configuration
- **`config-deploy.json`** - Deployment paths (staging, VPS destination, media content location)

Note: `config-md-to-html.json` is no longer used. The md-to-html converter is now invoked with the `cfg.WebContentPath` directory as argument (from `config.json`).

## Key Conventions

### Global Configuration
The `cfg` variable (of type `deployConfig`) is loaded at package init and used throughout. Always reference `cfg.WebContentPath`, `cfg.VpsAlias`, etc. rather than hardcoding paths.

### Repository Dependencies
The tool manages multiple repositories in `~/Software/`:
- **cmoli.es** - Main website content (src → copied to staging)
- **checkIframe** - Project docs (docs → staging/projects/check-iframe)
- **wiki** - Wiki content (src → staging/wiki)
- **md-to-html-go** - External converter invoked via `go run`
- **Tools** (open-urls, job-check-lambda-name, job-modify-issue-name) - Copied to staging/tools

These are pulled via `git pull` (if exist) or cloned fresh (if not). Updates to repo URLs must be made in `git.go`.

### Media Symlinks
The `setMedia()` function creates symlinks recursively from `cfg.MediaContentPath` to mirror the structure in `cfg.WebContentPath`. Source and destination paths must match for symlinks to resolve correctly. See README for setup requirements.

### HTML Post-Processing
Currently only modifies one specific file: `projects/rust-vs-other-languages/02-results-summary.html` (adds CSS class to `<table>`). Use `modify_html.go` for additional post-conversion tweaks; avoid editing the HTML post-generation in other workflows.

### Error Handling
The tool uses `panic()` for config loading errors and `exitIfError()` for runtime operations. Exit code 1 on error, 0 on success.

### Command Execution
All shell commands go through the `run()` function in `command.go`, which:
- Prints the command before execution
- Panics if execution fails
- Prints non-empty output (trimming trailing newlines)

## Requirements

- **Go 1.25.5+** (as specified in go.mod)
- **Git** - For pulling repositories
- **rsync** - For VPS sync
- **Media directory** - `$HOME/Software/cmoli-media-content` must exist with matching paths to markdown structure
- **VPS SSH access** - Configured via SSH config alias (set in config-deploy.json)

## Versioning & Release

Use semantic versioning. Update `CHANGELOG.md` before tagging:
```bash
git tag -a X.Y.Z -m "Release description"
```
