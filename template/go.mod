module <%= projectPath %>

require (
	github.com/golang/protobuf v1.2.0<% if (gateway) { %>
	github.com/grpc-ecosystem/grpc-gateway v1.8.5<% } %>
	golang.org/x/net v0.0.0-20181220203305-927f97764cc3
	google.golang.org/genproto v0.0.0-20180817151627-c66870c02cf8
	google.golang.org/grpc v1.19.0
)
