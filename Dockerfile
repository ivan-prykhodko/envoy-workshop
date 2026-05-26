ARG GO_VERSION=1.25


#FROM golang:${GO_VERSION}-alpine AS builder
#ENV CGO_ENABLED=0
#ADD . /build
#WORKDIR /build
#RUN cd cmd/httpserver && go build -ldflags "-s -w" -o /build/httpserver
#RUN cd cmd/db && go build -ldflags "-s -w" -o /build/db


#FROM golang:${GO_VERSION}-alpine AS app_prod
#COPY --from=builder /build/httpserver /srv/httpserver
#COPY --from=builder /build/db /srv/db
#WORKDIR /srv
#EXPOSE 8080
#CMD ["/srv/httpserver"]


FROM golang:${GO_VERSION}-alpine AS app_dev
RUN apk add --no-cache protobuf-dev
RUN go install github.com/mitranim/gow@latest
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.11
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.6.1
WORKDIR /srv
EXPOSE 8080
CMD ["gow", "-c", "-v", "run", "cmd/product/httpserver/main.go"]
