all:
	go build -ldflags "-s -w"
install:
	install -m 755 i5 $(DESTDIR)/usr/bin/i5
uninstall:
	rm -rf /usr/bin/i5
test:
	go test ./...
