all:
	go build && echo "run 'sudo make install' to install"
install:
	install -m 555 i5 /usr/bin/i5
