#!/bin/bash
set -x
mkdir ./bin
echo "-----------------------------"
FILE=cmd/auth/main.go
APP="auth"
echo $APP
GOOS=windows GOARCH=386 go build -o ./bin/${APP}_win32.exe $FILE
GOOS=windows GOARCH=amd64 go build -o ./bin/${APP}_win64.exe $FILE
go build -o ./bin/${APP}_elf_64 $FILE
echo "-----------------------------"
FILE=cmd/negotiations/main.go
APP="negotiations"
echo $APP
GOOS=windows GOARCH=386 go build -o ./bin/${APP}_win32.exe  $FILE
GOOS=windows GOARCH=amd64 go build -o ./bin/${APP}_win64.exe  $FILE
go build -o ./bin/${APP}_elf64 $FILE