#!/usr/bin/env bash


function vet(){
    echo "vet project..."
    declare -a vet_errs=$(go vet $(go list ./...) 2>&1 >/dev/null)
    EXIT_CODE=$?
    if [[ ${EXIT_CODE} -ne 0 ]]; then
        exit 1
    fi
      if [[ ${vet_errs} ]]; then
        echo "fix it:"
        for f in "${vet_errs[@]}"
        do
            echo "$f"

        done
        exit 1

    else
        echo "code is ok"
        echo ${vet_errs}
    fi
    echo ""
    echo ""
}

function fmt(){
    echo "fmt lint..."
    declare -a fmts=$(gofmt -s -l  $(find . -type f -name '*.go' | grep -v 'vendor' |grep -v '.git' ))

    if [[ ${fmts} ]]; then
        echo "fix it:"
        for f in "${fmts[@]}"
        do
            echo "$f"

        done
        exit 1

    else
        echo "code is ok"
        echo ${fmts}
    fi
    echo ""
}

function go-lint(){
    echo "golint..."
    if [[ -f "$(go env GOPATH)/bin/golint" ]] || [[ -f "/usr/local/bin/golint" ]]; then
    declare -a lints=$(golint $(go list ./...))  ## its a hack to not lint generated code
    if [[ ${lints} ]]; then
        echo "fix it:"
        for l in "${lints[@]}"
        do
            echo "$l"

        done
        exit 1

    else
        echo "code is ok"
        echo ${lints}
    fi
    else
     printf "Cannot check golint, please run:
        go get -u -v golang.org/x/lint/golint/... \n"
        exit 1
    fi
    echo ""
}

function go-group(){
echo "gogroup..."
if [[ -f "$(go env GOPATH)/bin/gogroup" ]] || [[ -f "/usr/local/bin/gogroup" ]]; then
    declare -a lints=$(gogroup -order std,other,prefix=melsoft/melmetrics/  $(find . -type f -name "*.go" | grep -v "vendor/" ))

    if [[ ${lints} ]]; then
        echo "fix it:"
        for l in "${lints[@]}"
        do
            echo "$l"

        done
        exit 1

    else
        echo "code is ok"
        echo ${lints}
    fi
    else
    printf "Cannot check gogroup, please run:
        go get -u -v github.com/Bubblyworld/gogroup/... \n"
        exit 1
   fi
    echo ""

}

function golangci(){
    echo "golangci-lint linter running..."
    if [[ -f "$(go env GOPATH)/bin/golangci-lint" ]] || [[ -f "/usr/local/bin/golangci-lint" ]]; then
        golangci-lint run --out-format=colored-line-number ./...
    else
        printf "Cannot check golang-ci, please run:
        curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(go env GOPATH)/bin \n"
        exit 1
    fi
    echo ""
}

function golangci-ci_execute(){
    echo "golangci-lint-ci_execute linter running..."
    if [[ -f "$(go env GOPATH)/bin/golangci-lint" ]] || [[ -f "/usr/local/bin/golangci-lint" ]]; then
        golangci-lint run  ./...  > linters.out
    else
        printf "Cannot check golang-ci, please run:
        curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(go env GOPATH)/bin \n"
        exit 1
    fi
    echo ""
}
