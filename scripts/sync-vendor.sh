#!/usr/bin/env bash

set -e

SCRIPT_NAME="$(basename "$(test -L "$0" && readlink "$0" || echo "$0")")"
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null && pwd)"
ROOT_DIR="$(git rev-parse --show-toplevel)"
TOOLS_DIR=${ROOT_DIR}/tools

go env -w GOPROXY=https://goproxy.io

sync-vendor() {
  go mod tidy -v
  go mod download
  go mod vendor
  go mod verify
}


cd ${ROOT_DIR} || exit 1
echo $(pwd)
sync-vendor

cd ${TOOLS_DIR} || exit 1
echo $(pwd)
sync-vendor
