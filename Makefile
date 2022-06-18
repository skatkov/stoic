# Compile to ./out/stoic
# Takes two positional arguments:
# - The version (e.g.: v1.2)
# - The build hash (7 chars hex)
build:
	go build -ldflags "-X 'main.BinaryVersion=$1' -X 'main.BinaryBuildHash=$2'" -o ./out/stoic

# Reformat all code
format:
	go fmt ./...

# Run linter
lint:
	golangci-lint run

# Run CLI from sources “on the fly”
# Passes through all input args
cli:
	go run main.go "$@"

# Execute tests
test:
	go test ./...
