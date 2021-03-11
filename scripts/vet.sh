#!/usr/bin/env bash

set -e

go vet $(go list ./...)

echo "Done."
