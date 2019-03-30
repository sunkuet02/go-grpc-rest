#!/usr/bin/env bash

export SRC_DIR=api/proto/v1
export DST_DIR=pkg/api/v1
export SWGR_DIR=api/swagger/v1

protoc --proto_path=$SRC_DIR --proto_path=third_party --go_out=plugins=grpc:$DST_DIR $SRC_DIR/*.proto
protoc --proto_path=$SRC_DIR --proto_path=third_party --grpc-gateway_out=logtostderr=true:$DST_DIR $SRC_DIR/*.proto
protoc --proto_path=$SRC_DIR --proto_path=third_party --swagger_out=logtostderr=true:$SWGR_DIR $SRC_DIR/*.proto
