set GOPROXY=https://goproxy.io
go mod download
set CGO_ENABLED=0
set GOOS=linux
set GOARCH=amd64
go build -o bin/pic-proxy main.go