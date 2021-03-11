// +build tools

package tools

//go:generate go clean
//go:generate go install -v -mod=vendor github.com/golangci/golangci-lint/cmd/golangci-lint
//go:generate go install -v -mod=vendor github.com/vasi-stripe/gogroup/cmd/gogroup
//go:generate go install -v -mod=vendor github.com/goreleaser/goreleaser

import (
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/goreleaser/goreleaser"
	_ "github.com/vasi-stripe/gogroup/cmd/gogroup"
)
