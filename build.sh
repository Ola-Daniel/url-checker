#!/bin/bash

set -euo pipefail

echo "Building URL Checker..."

GOOS=${GOOS:-linux}
GOARCH=${GOARCH:-amd64}

go build -v -o URL-Checker-$GOOS-$GOARCH