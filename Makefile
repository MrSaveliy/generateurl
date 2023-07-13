LOCAL_BIN:=$(CURDIR)/bin

bins:
	GOBIN=$(LOCAL_BIN) go install github.com/bufbuild/buf/cmd/buf@latest && \
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway && \
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 && \
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go && \
	GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc && \
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@latest

generate:
	@PATH=$(LOCAL_BIN):$(PATH) buf generate

db:
	psql postgres -c "drop database if exists generator;"
	createdb generator
	@PATH=$(LOCAL_BIN):$(PATH) goose -allow-missing -dir migrations postgres "dbname=generator sslmode=disable" up