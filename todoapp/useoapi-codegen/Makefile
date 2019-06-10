OPTS ?= "--log=DEBUG"
# EXTRA_OPTS := --sort-keys
TOOLSDIR := _tools

## codegen
# https://github.com/deepmap/oapi-codegen
# go get -v github.com/deepmap/oapi-codegen/cmd/oapi-codegen
codegen:
	# types only
	# oapi-codegen -generate types openapi3.yml
	mkdir -p Openapi3
	oapi-codegen -generate server,types,spec openapi3.yml > Openapi3/server.go

run:
	GO111MODULE=on go run cmd/useoapi-server/main.go

## validate
# https://github.com/wework/speccy
node_modules:
	npm install
validate:
	./node_modules/.bin/speccy lint openapi3.yml
.PHONY: validate

## migrate
openapi3.yml: ../usegoswagger/swagger.yml
	python ${TOOLSDIR}/2to3.py ${OPTS} ${EXTRA_OPTS} --src $< --savedir .

