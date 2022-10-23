#!/bin/bash

rm cover.html cover.out
go test -cover ./... -coverprofile=cover.out
go tool cover -html=cover.out -o cover.html