### Requirements
- Install go1.15.1 darwin/amd64 ( see https://golang.org/doc/install)
- Make sure the file is present under the directory resources/GCP/credentials/google.json
#### How to test
go test ./...

go clean
go build app.go
./app (or "go run app.go")
