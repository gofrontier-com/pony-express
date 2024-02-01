build:
	go build cmd/pony/pony.go

install:
	go install cmd/pony/pony.go

test:
	go test -v ./...
