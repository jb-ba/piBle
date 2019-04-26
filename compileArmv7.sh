#!/usr/bin/env bash
env GOARCH=arm GOARM=7 GOOS=linux go build -o piBle main/main.go
scp piBle piHome: