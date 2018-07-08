GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get
BINARY_NAME=post
OS=linux
ARCH=amd64

.PHONY: all
all: build zip

.PHONY: build
build:
	env GOOS=$(OS) GOARCH=$(ARCH) $(GOBUILD) -o $(BINARY_NAME)

.PHONY: clean
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_NAME).zip
	ls

.PHONY: run
run:
	./$(BINARY_NAME)

.PHONY: zip
zip:
	zip post.zip ./$(BINARY_NAME)
