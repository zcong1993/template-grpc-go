FROM golang:1.11.1 AS build
WORKDIR /mnt
COPY . .
RUN CGO_ENABLED=0 go build -o ./bin/server main.go

FROM alpine:3.7
WORKDIR /opt
EXPOSE 1234<% if (gateway) { %>
EXPOSE 9009<% } %>
RUN apk add --no-cache ca-certificates
COPY --from=build /mnt/bin/* /usr/bin/
CMD ["server"]
