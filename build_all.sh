#!/usr/bin/bash

CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/pack-windows-amd64.exe .

CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o bin/pack-windows-32.exe .

CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o bin/pack-darwin-amd64 .

CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o bin/pack-darwin-arm64 .

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/pack-linux-amd64 .

CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o bin/pack-linux-arm64 .

CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -o bin/pack-linux-32 .
