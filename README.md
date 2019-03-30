###Go-gRPC-REST
This is a very basic project containing only few mock methods to 
test gRPC and REST with golang.

Mainly [go-grpc-http-rest-microservice-tutorial](https://github.com/amsokol/go-grpc-http-rest-microservice-tutorial/tree/master)
is followed to see basics and project structure.

***
###Compile and Build
You can generate the protocol buffers and gateway go files by running the
below command in the root directory of the project.
```bash
./third_party/protoc-gen.sh 
``` 

After generating above files you can build the server codes by running the below command
in the /cmd/server
```bash
go build .
```

Then you can run the server file by executing 
```bash
./server --grpc-port=7171 --http-port=7172 --log-time-format=2006-01-02T15:04:05
```



###Used Libraries
- [golang/protobuf](https://github.com/golang/protobuf) for Google's protocol buffers
- [grpc-ecosystem/grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway) for gRPC to JSON proxy generator
- [google/uuid](https://github.com/google/uuid) for uuid genenaration

