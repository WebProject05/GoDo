# GoDo 🚀

A tiny CLI to manage a local todo list saved in `todos.json`.

## Install

### From source (Go users)

```powershell
# Install the latest release from GitHub (replace the path with yours if you fork)
# Example: go install github.com/<your-github-username>/GoDo@latest
# make sure $GOBIN or %USERPROFILE%\go\bin is in your PATH

go install github.com/WebProject05/GoDo@latest
```

### Install globally on Windows (installer script)

Run the installer from the project root in PowerShell:

```powershell
# From PowerShell in the repository root
.\install.ps1

# If your Execution Policy prevents running scripts, use:
powershell -NoProfile -ExecutionPolicy Bypass -File .\install.ps1
```

### Install globally on Unix-like systems (installer script)

```bash
# Build and install to your $GOBIN (defaults to $HOME/go/bin)
./install.sh
# Ensure $GOBIN (or $HOME/go/bin) is in your PATH, for example:
# export PATH="$HOME/go/bin:$PATH"
```

---

## Publishing your repository to GitHub

If you want to push this project to GitHub and make `go install` work from anywhere, follow these steps:

1. If you forked or want to use your own account, update the module path in `go.mod` to match your repo:

```bash
# replace USER and REPO with your GitHub username/repo
go mod edit -module=github.com/USER/REPO
git add go.mod
git commit -m "update module path for publishing"
```

2. Initialize git (if needed) and push to GitHub:

```bash
git init
git add .
git commit -m "initial commit"
# create repo on GitHub (use web UI or gh cli)
# then add remote and push
git remote add origin git@github.com:USER/REPO.git
git branch -M main
git push -u origin main
```

3. Install via `go install` on any machine:

```bash
go install github.com/USER/REPO@latest
```

---

### CI / Releases

A workflow is included (`.github/workflows/build.yml`) that builds the binary for Linux and Windows on pushes to `main` and for releases. On release you can download the generated artifacts from the workflow run or extend the workflow to attach assets to GitHub Releases.


From Command Prompt (CMD):

```cmd
powershell -NoProfile -ExecutionPolicy Bypass -File "%CD%\install.ps1"
```

From PowerShell Core (pwsh):

```bash
pwsh -NoProfile -ExecutionPolicy Bypass -File ./install.ps1
```

If you see an error like "provide path to an existing .ps1", that means PowerShell couldn't find the script at the path you provided. Fixes:
- Make sure you're running the command from the repository root (where `install.ps1` is located) or pass the full absolute path, e.g. `powershell -File C:\full\path\to\GoDo\install.ps1`.
- Confirm the file exists: `Test-Path .\install.ps1` or `dir install.ps1`.
- Use the `-ExecutionPolicy Bypass` form if execution policy blocks script execution.

### Build locally (all platforms)

If you prefer building the binary yourself:

```bash
# clone and build
git clone https://github.com/WebProject05/GoDo.git
cd GoDo
# Windows
go build -o godo.exe
# macOS / Linux
go build -o godo
```

### From Releases

Pre-built artifacts for Windows, Linux, and macOS are published to GitHub Releases when a tag is pushed (see CI section).

## Commands

Below are the available commands and usage examples. All commands operate on the currently selected table (see "Multi-table support" below).

- List todos in the current table:

```powershell
godo -list
```

- Add a todo item (adds to current table):

```powershell
godo -add "Buy milk"
```

- Toggle completion by index:

```powershell
godo -toggle 0
```

- Edit a todo (format: `id:new_title`):

```powershell
godo -edit 0:New title
```

- Delete a todo by index:

```powershell
godo -del 0
```

- Create a new table and switch to it (name is quoted if it contains spaces):

```powershell
godo -newtable "work"
```

- Switch to an existing table:

```powershell
godo -switch "work"
```

- List all available tables (current table is marked with `*`):

```powershell
godo -alltables
```

---

### Multi-table support

GoDo now supports multiple named todo tables (e.g. `default`, `work`, `personal`). Key notes:

- If you have an existing `todos.json` in the old single-list format, the first run of the new binary will automatically migrate those todos into a `default` table.
- Use `-newtable` to create and switch to a new table; use `-switch` to change the active table; use `-alltables` to see all tables and which one is active.

---

> Tip: After installation you can run the binary from anywhere as `godo -list` (see Install section above).

## Development

- Code is written in Go. Dependencies are tracked in `go.mod`.
- Local todo data is stored in `todos.json` in the repo directory.

## CI / Releases

This repository uses GitHub Actions + goreleaser to publish Windows binaries automatically when you push a tag like `v0.1.0`.

Tag a release and push tags:

```bash
git tag v0.1.0
git push origin --tags
```

Goreleaser will build Windows artifacts (zip) and attach them to the GitHub Release automatically.

## Supported platforms

- Windows, Linux, macOS (amd64, arm64)

Multi-platform builds are enabled in CI — releases will include platform-specific archives for Windows (zip) and Unix (tar.gz).

## Contributing

PRs are welcome — please include tests for any new behavior.

---

If you want, I can add Homebrew/Scoop/Chocolatey manifests to make the CLI available through package managers. 💡
