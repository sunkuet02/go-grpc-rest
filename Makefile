.PHONY:
	clean proto run

all: clean proto run

proto:
	./third_party/protoc-gen.sh

proto-clean:
	rm pkg/api/v1/*

clean:
	rm server

run:
	go build ./cmd/server
	./server