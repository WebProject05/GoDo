# GoDo 🚀

A tiny CLI to manage a local todo list saved in `todos.json`.

## Install

### From source (Go users)

```powershell
# Install the latest release from GitHub
go install github.com/WebProject05/GoDo@latest
```

### From source (build yourself)

```powershell
# Build locally
go build -o godo.exe
.
# or use the provided installer script on Windows
powershell -ExecutionPolicy Bypass -File .\install.ps1
```

### From Releases

Pre-built Windows artifacts are published to GitHub Releases when a tag is pushed (see CI section).

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

- Windows (amd64, arm64)

If you'd like multi-platform builds (macOS / Linux), tell me and I can enable them in the CI config.

## Contributing

PRs are welcome — please include tests for any new behavior.

---

If you want, I can add Homebrew/Scoop/Chocolatey manifests to make the CLI available through package managers. 💡
