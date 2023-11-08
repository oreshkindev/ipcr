fmt:
	go list -f '{{.Dir}}' ./... | grep -v /vendor/ | xargs -L1 gofmt -l

mod:
	go mod tidy
	go mod vendor

build:
	go build -o bin/ipcr main.go

run:
	chmod +x bin/ipcr
	./bin/ipcr