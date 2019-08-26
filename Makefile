all:
	go build -ldflags "-s -w"
install:
	go install
	# install -m 555 i5 /usr/bin/i5
test:
	go test ./...
