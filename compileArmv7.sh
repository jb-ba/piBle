#!/usr/bin/env bash
#env CGO_ENABLED=0 GOARCH=arm GOARM=7 GOOS=linux go build -ldflags "-s -w" -a -installsuffix cgo -o piBle main/main.go
env CGO_ENABLED=0 GOARCH=arm GOARM=7 GOOS=linux go build -o piBle main/main.go
scp piBle deactiveBluetooth.sh piHome:
