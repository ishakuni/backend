#!/bin/bash -eu

PATH=$PATH:$GOPATH/bin
protodir=../../proto

protoc --go_out=plugins=grpc:genproto -I $protodir $protodir/ishakuni.proto
