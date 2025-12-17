# GoDo 🚀

A tiny CLI to manage a local todo list saved in `todos.json`.

## Install

### From source (Go users)

```powershell
# Install the latest release from GitHub
go install github.com/WebProject05/GoDo@latest
# make sure $GOBIN or %USERPROFILE%\go\bin is in your PATH
```

### Windows — installer script (recommended for non-Go users)

Run the installer from the project root in PowerShell:

```powershell
# From PowerShell in the repository root
.\install.ps1

# If your Execution Policy prevents running scripts, use:
powershell -NoProfile -ExecutionPolicy Bypass -File .\install.ps1
```

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

## Usage

- List todos:

```powershell
godo -list
```

- Add a todo:

```powershell
godo -add "Buy milk"
```

- Toggle completion:

```powershell
godo -toggle 0
```

- Edit a todo (format: `id:new_title`):

```powershell
godo -edit 0:New title
```

- Delete a todo:

```powershell
godo -del 0
```

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
