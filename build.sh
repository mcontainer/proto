#!/bin/bash

for package in "containers"
do
    protoc -I $package $package/$package.proto --go_out=plugins=grpc:$package
done

