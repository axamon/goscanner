# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
INARY_NAME=mybinary
BINARY_UNIX=$(BINARY_NAME)_unix

all: test build
build: 
	$(GOBUILD) -o $(BINARY_NAME) -v
test: 
	$(GOTEST) -v -cover ./scan
test-file:
	$(GOTEST) -v -cover ./scan -o testscan
profile: test-file
	./testscan --test.v --test.cpuprofile profili/cpu.pprof
	$(GOCMD) tool pprof --pdf eseguibili/goscanner-linux profili/cpu.pprof > profili/cpu.pdf
dockerimage:
	CGO_ENABLED=0 $(GOBUILD) -ldflags="-w -s" -a -installsuffix cgo -o main
	podman build -t goscanner/latest -f Dockerfile
run-docker: dockerimage
	podman run --rm -it --name makefile-goscanner goscanner/latest www.openbsd.org 443

clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)
