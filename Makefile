buf-generate:
	buf generate

run-server:
	cd ./grpc-server && go run main.go

run-client:
	cd ./grpc-client && go run main.go

buf-push:
	buf build && cd ./proto && buf push