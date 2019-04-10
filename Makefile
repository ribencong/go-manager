SHELL=PATH='$(PATH)' /bin/sh


GOBUILD=CGO_ENABLED=0 go build -ldflags '-w -s'

PLATFORM := $(shell uname -o)

NAME := YPManager.exe
ifeq ($(PLATFORM), Msys)
    INCLUDE := ${shell echo "$(GOPATH)"|sed -e 's/\\/\//g'}
else ifeq ($(PLATFORM), Cygwin)
    INCLUDE := ${shell echo "$(GOPATH)"|sed -e 's/\\/\//g'}
else
	INCLUDE := $(GOPATH)
	NAME=YPManager
endif

# enable second expansion
.SECONDEXPANSION:

.PHONY: all
.PHONY: pbs

BINDIR=$(INCLUDE)/bin


all: pbs  build

pbs:
	cd pbs/ && $(MAKE)

build:
	GOARCH=amd64 $(GOBUILD) -o $(BINDIR)/$(NAME)
clean:
	rm $(BINDIR)/$(NAME)
