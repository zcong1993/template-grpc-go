package main

//go:generate protoc -I pb -I /usr/local/include -I $GOPATH/src -I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:pb pb/pb.proto
//go:generate protoc -I pb -I /usr/local/include -I $GOPATH/src -I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true:pb pb/pb.proto
//go:generate protoc -I pb -I /usr/local/include -I $GOPATH/src -I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --swagger_out=logtostderr=true:pb pb/pb.proto
