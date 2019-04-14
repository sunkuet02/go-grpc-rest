BINARY_NAME=grpc-rest

.PHONY:
	proto build

all: clean build

proto:
	./third_party/protoc-gen.sh

proto-clean:
	rm pkg/api/v1/*

clean:
	if [ -f ${BINARY_NAME} ] ; then rm ${BINARY_NAME} ; fi

build: proto
	go build -o ${BINARY_NAME} ./cmd/server

run: build
	./${BINARY_NAME}

docker-build:
	docker build -t go-grpc-rest .

docker-run: docker-build
	docker run -it -p 9090:9090 -d --name=grpc-rest --network=host go-grpc-rest