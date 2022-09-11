fmt:
	go fmt main.go

vet:
	go vet main.go

build:
    go build -o bin/main main.go

run:
    go run main.go

compile:
    GOOS=windows GOARCH=amd64 go build -o bin/main-windows-amd64 main.go

godoc:
    godoc -http=:6060 using GOPATH mode
