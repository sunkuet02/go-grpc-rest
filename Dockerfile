# Builder
FROM golang:1.12.4-alpine3.9 as builder

WORKDIR $GOPATH/src/github.com/sunkuet02/go-grpc-rest

COPY . .

RUN apk update && apk upgrade && \
    apk --update add --no-cache bash curl git gcc make unzip protobuf

RUN go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
    && go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger \
    && go get -u github.com/golang/protobuf/protoc-gen-go

RUN go get -d -v ./...

RUN make

#Deployment
FROM alpine:latest

RUN apk update && apk upgrade && \
    apk --update --no-cache add tzdata && \
    mkdir /deployment

WORKDIR /deployment

EXPOSE 9090

COPY --from=builder /go/src/github.com/sunkuet02/go-grpc-rest /deployment

CMD /deployment/grpc-rest
