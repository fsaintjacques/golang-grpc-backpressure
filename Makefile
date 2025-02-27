.PHONY: proto
proto:
	@ rm -rf proto/gen
	@ mkdir -p proto/gen/echoservice
	cd proto && protoc --go_out=gen/echoservice --go_opt=paths=source_relative --go-grpc_out=gen/echoservice --go-grpc_opt=paths=source_relative echo.proto

.PHONY: server
server:
	@ go run cmd/server/main.go

.PHONY: client
client:
	@ go run cmd/client/main.go
