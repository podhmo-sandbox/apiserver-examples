#! /usr/bin/make
#
# Makefile for goa cellar example
#
# Targets:
# - clean     delete all generated files
# - generate  (re)generate all goagen-generated files.
# - build     compile executable
#
# Meta targets:
# - all is the default target, it runs all the targets in the order above.
#
DEPEND=	bitbucket.org/pkg/inflect \
	github.com/goadesign/goa \
	github.com/goadesign/goa/goagen \
	github.com/goadesign/goa/logging/logrus \
	github.com/sirupsen/logrus \
	gopkg.in/yaml.v2 \
	golang.org/x/tools/cmd/goimports

CURRENT_DIR := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

all: depend clean generate build

depend:
	@go get $(DEPEND)

clean:
	@rm -rf app
	@rm -rf client
	@rm -rf tool
	@rm -rf public/swagger
	@rm -rf public/schema
	@rm -rf public/js
	@rm -f cellar

generate:
	@goagen app     -d github.com/goadesign/goa-cellar/design
	@goagen client  -d github.com/goadesign/goa-cellar/design
	@goagen schema  -d github.com/goadesign/goa-cellar/design -o public
	@goagen swagger -d github.com/goadesign/goa-cellar/design -o public
	@goagen js      -d github.com/goadesign/goa-cellar/design -o public

build:
	@go build -o cellar
