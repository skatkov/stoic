# Compile to ./out/stoic
# Takes two positional arguments:
# - The version (e.g.: v1.2)
# - The build hash (7 chars hex)
build:
	go build -ldflags "-X 'main.BinaryVersion=$1' -X 'main.BinaryBuildHash=$2'" -o ./out/stoic stoic.go

# Reformat all code
format:
	go fmt ./...

# Run CLI from sources “on the fly”
# Passes through all input args
cli:
	go run stoic.go "$@"

# Execute tests
test:
	go test ./...
