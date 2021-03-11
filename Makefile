NAME=pixelizer

SHELL := env DOCKER_REPO=$(DOCKER_REPO) $(SHELL)
DOCKER_REPO?=eu.gcr.io/melsoft-infra

SHELL := env VERSION=$(VERSION) $(SHELL)
VERSION ?= $(shell date -u +%Y%m%d.%H.%M.%S)


# COLORS
GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
RESET  := $(shell tput -Txterm sgr0)


TARGET_MAX_CHAR_NUM=20


define colored
	@echo '${GREEN}$1${RESET}'
endef

## Show help
help:
	${call colored, help is running...}
	@echo 'link this Makefile from scripts dir to core root'
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
		helpMessage = match(lastLine, /^## (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 3, RLENGTH); \
			printf "  ${YELLOW}%-$(TARGET_MAX_CHAR_NUM)s${RESET} ${GREEN}%s${RESET}\n", helpCommand, helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

## vet project
vet:
	${call colored, vet is running...}
	./scripts/vet.sh
.PHONY: vet

## Compile executable
compile:
	${call colored, compile is running...}
	./scripts/compile.sh
.PHONY: compile

## Release
release:
	./scripts/release.sh
.PHONY: release

## Release local snapshot
release-local-snapshot:
	${call colored, release is running...}
	./scripts/release-local-snapshot.sh
.PHONY: release-local-snapshot

## Installs tools from vendor.
install-tools: sync-vendor
	./scripts/install-tools.sh
.PHONY: install-tools

## Sync vendor of root project and tools.
sync-vendor:
	./scripts/sync-vendor.sh
.PHONY: sync-vendor
