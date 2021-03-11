#!/usr/bin/env bash

set -e

SCRIPT_NAME="$(basename "$(test -L "$0" && readlink "$0" || echo "$0")")"
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd -P)"
ROOT_DIR="$(cd ${SCRIPT_DIR} && git rev-parse --show-toplevel)"
BIN_DIR=${ROOT_DIR}/bin
cd ${ROOT_DIR}/tools || exit 1

export GO111MODULE=on

go generate -tags=tools ./...

cd - || exit 1
