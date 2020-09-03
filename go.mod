module github.com/gnharishkumar13/go-grpc--learn

go 1.15

require (
	github.com/golang/protobuf v1.4.1
	google.golang.org/grpc v1.27.0
	google.golang.org/protobuf v1.25.0
)

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.31.1
