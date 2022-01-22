#!/bin/bash
cd ../ || exit
go env -w CGO_ENABLED=0
go env -w GOOS=linux
go env -w GOARCH=amd64
go build -ldflags '-w -s' -gcflags '-l' -a -o rsm-linux-amd64
upx -9 rsm-linux-amd64
mv ./rsm-linux-amd64 ./pkg/rsm-linux-amd64
chmod 777 ../pkg/rsm-linux-amd64
go env -w CGO_ENABLED=1
go env -w GOOS=linux
go env -w GOARCH=amd64
cd ./builders/ || exit
echo "rsm-linux-amd64 success"



go env -w CGO_ENABLED=0
go env -w GOOS=windows
go env -w GOARCH=amd64
go build -ldflags '-w -s' -gcflags '-l' -a -o mdnsserver.exe
go env -w CGO_ENABLED=1
go env -w GOOS=linux
go env -w GOARCH=amd64
