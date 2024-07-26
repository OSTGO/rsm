#!/bin/bash
cd ../ || exit
go env -w CGO_ENABLED=0
go env -w GOOS=linux
go env -w GOARCH=amd64
go build -ldflags '-w -s' -gcflags '-l' -a -o pkg/rsm-linux-amd64
upx -9 pkg/rsm-linux-amd64
chmod 777 pkg/rsm-linux-amd64
go env -w CGO_ENABLED=1
go env -w GOOS=linux
go env -w GOARCH=amd64
cd ./builders/ || exit
echo "rsm-linux-amd64 success"

