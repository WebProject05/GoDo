#!/usr/bin/env bash
set -euo pipefail

echo "Building godo..."
go build -o godo

GOBIN=$(go env GOBIN)
if [ -z "$GOBIN" ]; then
  GOBIN="$HOME/go/bin"
fi

mkdir -p "$GOBIN"
cp -f godo "$GOBIN/godo"
rm -f godo

echo "Installed to $GOBIN. Ensure it's in your PATH (e.g. export PATH=\"$GOBIN:$PATH\")."
