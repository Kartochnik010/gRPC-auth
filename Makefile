## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

## run: run the project
.PHONY: run
run:
	@go run ./cmd/sso

.PHONY: gen-grpc
## gen-grpc: generate go grpc code from protofile
gen-grpc:
	@protoc -I protos/proto protos/proto/sso/sso.proto --go_out=./protos/gen/go --go_opt=paths=source_relative  --go-grpc_out=./protos/gen/go --go-grpc_opt=paths=source_relative 

.PHONY: migrate
## migrate: apply migrations stored
migrate:
	@go run ./cmd/migrator --storage-path=./storage/sso.db --migrations-path=./migrations

.PHONY: migrate-test
## migrate-test: apply test migrations
migrate-test:
	@go run ./cmd/migrator --storage-path=./storage/sso.db --migrations-path=./tests/migrations --migrations-table=migrations_test

.PHONY: test
## test: run tests
test: 
	@go test -v ./tests