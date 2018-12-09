#!/bin/bash

# Ensure you have a Docker Hub account and belongs to hwsc organization
docker build -t dev-hwsc-grpc-sample-svc .
docker tag dev-hwsc-grpc-sample-svc hwsc/dev-hwsc-grpc-sample-svc
docker push hwsc/dev-hwsc-grpc-sample-svc