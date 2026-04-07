GOOS=linux GOARCH=amd64 go build -o dist/geep-linux
GOOS=darwin GOARCH=arm64 go build -o dist/geep-mac-arm
GOOS=windows GOARCH=amd64 go build -o dist/geep-windows-x64.exe
GOOS=windows GOARCH=386 go build -o dist/geep-windows-x32.exe