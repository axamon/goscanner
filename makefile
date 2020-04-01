APPNAME=goscanner
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
	$(GOTEST) -v -cover ./...
test-file:
	$(GOTEST) -v -cover -o test
profile: test-file
	./testscan --test.v --test.cpuprofile profili/cpu.pprof
	$(GOCMD) tool pprof --pdf eseguibili/goscanner-linux profili/cpu.pprof > profili/cpu.pdf
dockerimage: test
	CGO_ENABLED=0 $(GOBUILD) -ldflags="-w -s" -a -installsuffix cgo -o main
	podman build -t $(APPNAME):latest -f Dockerfile
dockerrun:
	podman run --rm -it --name $(APPNAME) $(APPNAME):latest www.openbsd.org 443

clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)
