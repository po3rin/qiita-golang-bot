GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get
BINARY_NAME=qiitter

all: build run
build:
	$(GOBUILD) -o $(BINARY_NAME)
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
run:
	./$(BINARY_NAME)
dep:
	dep ensure
