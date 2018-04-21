GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get
BINARY_NAME=post
OS=linux
ARCH=amd64

all: build zip
build:
	env GOOS=$(OS) GOARCH=$(ARCH) $(GOBUILD) -o $(BINARY_NAME)
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_NAME).zip
	ls
run:
	./$(BINARY_NAME)
dep:
	dep ensure
zip:
	zip post.zip ./$(BINARY_NAME)
