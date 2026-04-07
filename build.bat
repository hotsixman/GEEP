set GOOS=linux
set GOARCH=amd64
go build -o dist/geep-linux

set GOOS=darwin
set GOARCH=arm64
go build -o dist/geep-mac-arm

set GOOS=windows
set GOARCH=amd64
go build -o dist/geep-windows-x64.exe

set GOOS=windows
set GOARCH=386
go build -o dist/geep-windows-x32.exe