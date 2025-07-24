#!/bin/bash
set -e

echo "Building pokedexcli..."
go build -o ./bin/pokedexcli ./cmd/main.go
echo "✅ Done!"
echo
