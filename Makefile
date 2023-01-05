generate:
	protoc -I=./ --go_out=./ ./proto/*/**.proto

