Run Go Lang `go run file.go`

Initialize Go mod `go mod init`

Build Go to exe
- `go build`
- `go build -o app_name.exe`

Compiling to different OS
Windows: `GOOS=windows GOARCH=amd64 go build -o winapp.exe`
Linux: `GOOS=linux GOARCH=amd64 go build -o linuxapp`
Mac: `GOOS=darwin GOARCH=amd64 go build -o macapp`

Formatting Go source code(gofmt)
- `gofmt -w file.go`
- `go fmt`