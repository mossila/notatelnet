#!/bin/sh
env GOOS=windows GOARCH=amd64 go build -ldflags "-w" -o notatelnet.exe