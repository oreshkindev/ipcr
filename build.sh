#!/bin/sh
GOOS=linux GOARCH=amd64 go build -o bin/ipcr ./main.go
