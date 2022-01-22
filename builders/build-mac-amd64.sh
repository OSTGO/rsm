#!/bin/bash
cd ../ || exit
go env -w CGO_ENABLED=0
go env -w GOOS=darwin
go env -w GOARCH=amd64
go build -ldflags '-w -s' -gcflags '-l' -a -o rsm-mac-amd64
upx -9 rsm-mac-amd64
mv ./rsm-mac-amd64 ./pkg/rsm-mac-amd64
chmod 777 ../pkg/rsm-mac-amd64
go env -w CGO_ENABLED=1
go env -w GOOS=linux
go env -w GOARCH=amd64
cd ./builders/ || exit
echo "rsm-mac-amd64 success"