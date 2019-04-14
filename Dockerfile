# Builder
FROM golang:1.12.4-alpine3.9 as builder

WORKDIR $GOPATH/src/github.com/sunkuet02/go-grpc-rest

COPY . .

RUN apk --update add curl git gcc make unzip
RUN go get -d -v ./...
RUN go install -v ./...

RUN PROTOC_ZIP=protoc-3.7.1-linux-x86_64.zip \
    && curl -OL https://github.com/google/protobuf/releases/download/v3.7.1/$PROTOC_ZIP \
    && unzip -o $PROTOC_ZIP -d /usr/local bin/protoc

RUN make

EXPOSE 9090
