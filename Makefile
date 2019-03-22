build:
	@ vgo build

gox:
	@ gox -output "dist/sum_{{.OS}}_{{.Arch}}"

test:
	@ richgo test -v -cover ./...

coverage:
	@ richgo test -v -coverprofile=/tmp/profile -covermode=atomic ./...
	@ go tool cover -html=/tmp/profile

validate-ci-config:
	@ circleci config validate -c .circleci/config.yml

