module github.com/firacloudtech/grpc-echo-benchmark

go 1.19

require google.golang.org/protobuf v1.28.1

require (
	buf.build/gen/go/grpc-ecosystem/grpc-gateway/protocolbuffers/go v1.28.1-20221127060915-a1ecdc58eccd.4 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.3.0 // indirect
	golang.org/x/sys v0.3.0 // indirect
	golang.org/x/text v0.5.0 // indirect
)

require (
	buf.build/gen/go/firacloudtech/grpc-echo-benchmark/grpc/go v1.2.0-20230123160103-0a94928fc203.4
	buf.build/gen/go/firacloudtech/grpc-echo-benchmark/protocolbuffers/go v1.28.1-20230123162638-21cf7e84720f.4
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.15.0
	google.golang.org/genproto v0.0.0-20221207170731-23e4bf6bdc37
	google.golang.org/grpc v1.51.0
)
