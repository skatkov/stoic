#!/bin/sh

# Compile to ./out/stoic
# Takes two positional arguments:
# - The version (e.g.: v1.2)
# - The build hash (7 chars hex)
run_build() {
	go build \
	  -ldflags "-X 'main.BinaryVersion=$1' -X 'main.BinaryBuildHash=$2'" \
	  -o ./out/stoic \
	  stoic.go
}


# Run CLI from sources “on the fly”
# Passes through all input args
run_cli() {
	go run ./stoic.go "$@"
}