BIN=sum
JOB=build

build:
	@ go build

dist:
	@ gox -output "dist/$(BIN)_{{.OS}}_{{.Arch}}" --osarch "linux/amd64 darwin/amd64 windows/amd64"

test:
	@ go vet ./...
	@ richgo test -v -cover ./...

coverage:
	@ richgo test -v -coverprofile=/tmp/profile -covermode=atomic ./...
	@ go tool cover -html=/tmp/profile

validate-ci-config:
	@ circleci config validate -c .circleci/config.yml

local-ci:
	@ circleci local execute --job $(JOB)
