build:
	go build

install:
	go install

clean:
	go clean

check:
	go vet
	golint

defaults:
	build
