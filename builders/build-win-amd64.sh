#!/bin/bash
cd ../ || exit
go env -w CGO_ENABLED=0
go env -w GOOS=windows
go env -w GOARCH=amd64
go build -ldflags '-w -s' -gcflags '-l' -a -o pkg/rsm-win-amd64.exe
chmod 777 pkg/rsm-win-amd64.exe
go env -w CGO_ENABLED=1
go env -w GOOS=linux
go env -w GOARCH=amd64
cd ./builders/ || exit
echo "rsm-win-amd64 success"
