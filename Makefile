build:
	go mod vendor
	go build -o build/formatlog -mod=vendor pkg/formatlog.go

run:
	cat assets/example_log | ./build/formatlog

install:
	sudo cp build/formatlog /usr/local/bin/formatlog

.PHONY: build
