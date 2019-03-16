#!/bin/bash -x
set -e

<% if (gateway) { %>MOD_PATH=`go list -f '{{ .Dir }}' -m github.com/grpc-ecosystem/grpc-gateway`<% } %>
OUT_DIR=./pb

rm -rf $OUT_DIR/*.{go,json}

protoc -I pb -I /usr/local/include -I $GOPATH/src<% if (gateway) { %> -I $MOD_PATH/third_party/googleapis<% } %> --go_out=plugins=grpc:$OUT_DIR pb/*.proto<% if (gateway) { %>
protoc -I pb -I /usr/local/include -I $GOPATH/src -I $MOD_PATH/third_party/googleapis --grpc-gateway_out=logtostderr=true:$OUT_DIR pb/*.proto
protoc -I pb -I /usr/local/include -I $GOPATH/src -I $MOD_PATH/third_party/googleapis --swagger_out=logtostderr=true:$OUT_DIR pb/*.proto<% } %>
