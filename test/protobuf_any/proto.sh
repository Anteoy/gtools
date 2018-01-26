#!/bin/bash

protoc --go_out=plugins=grpc:. *.proto

ls *.pb.go | xargs -n1 -IX bash -c 'sed s/,omitempty// X > X.tmp && mv X{.tmp,}'