#!/bin/bash

rm -f ./coverage.out ./cover.out
go test -covermode=count -v -coverprofile=coverage.out
go tool cover -func=coverage.out -o coverage-func.txt
go tool cover -html=./coverage.out -o ./coverage.html
go test -test.bench=.* > ./benchmarks.txt
