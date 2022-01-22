#!/bin/bash
cd ../ || exit
go env -w CGO_ENABLED=0
go env -w GOOS=linux
go env -w GOARCH=arm64
go build -ldflags '-w -s' -gcflags '-l' -a -o rsm-linux-arm64
upx -9 rsm-linux-arm64
mv ./rsm-linux-arm64 ./pkg/rsm-linux-arm64
chmod 777 ../pkg/rsm-linux-arm64
go env -w CGO_ENABLED=1
go env -w GOOS=linux
go env -w GOARCH=amd64
cd ./builders/ || exit
echo "rsm-linux-arm64 success"
