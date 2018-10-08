#!/bin/bash

# generate protoc for our serice
echo "Generating grpc-sample.pb.go..."
protoc proto/grpc-sample.proto --go_out=plugins=grpc:.