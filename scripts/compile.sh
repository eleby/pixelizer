#!/usr/bin/env bash
set -e

APP=pixelizer

echo "Building ${APP}..."

SCRIPT_NAME="$(basename "$(test -L "$0" && readlink "$0" || echo "$0")")"
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
ROOT_DIR="$(cd ${SCRIPT_DIR} && git rev-parse --show-toplevel)"
BIN_DIR=${ROOT_DIR}/bin

COMMIT=$(git rev-parse HEAD)
SHORTCOMMIT=$(git rev-parse --short HEAD)
DATE=$(date -u +%Y-%m-%dT%H:%M:%SZ)
VERSION="$(git describe --tags "$(git rev-list --tags --max-count=1)")"

GOVERSION=$(go version | awk '{print $3;}')

if [[ "${VERSION}" = "" ]]; then
  VERSION="v0.0.0"
fi

BIN_OUT=${BIN_DIR}/${APP}

GO_BUILD_PACKAGE="${ROOT_DIR}/cmd/${APP}"
GO_BUILD_LDFLAGS="-s -w -X main.version=${VERSION} -X main.commit=${COMMIT} -X main.date=${DATE} -X main.appname=${APP} -X main.goversion=${GOVERSION}" 

rm -rf ${BIN_OUT}

go build -o ${BIN_OUT} -a -ldflags "${GO_BUILD_LDFLAGS}" "${GO_BUILD_PACKAGE}"

echo "Binary compiled at ${BIN_OUT}"
