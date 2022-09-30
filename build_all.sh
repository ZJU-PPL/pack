#!/usr/bin/bash

CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/pack-win64.exe .

CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o bin/pack-win32.exe .

CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o bin/pack-darwin64 .

CGO_ENABLED=0 GOOS=darwin GOARCH=386 go build -o bin/pack-darwin32 .

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/pack-linux64 .

CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -o bin/pack-linux32 .
