#!/bin/bash

GOPATH=`go env GOPATH`
PATH=$PATH:$GOPATH/bin
git clone --branch review https://github.com/cernbox/cs3apis build && cd build && make && cd ..
cat prototool_gen_go.yaml >> build/prototool.yaml
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
go get -u github.com/golang/protobuf/protoc-gen-go
cd build && prototool generate && cd ..
rm -rf build

# change import paths
egrep -lR 'github.com/cernbox/cs3/' ./cs3 | xargs sed -i 's|github.com/cernbox/cs3/|github.com/cernbox/go-cs3apis/cs3/|g'
