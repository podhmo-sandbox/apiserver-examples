OPTS ?= "--log=DEBUG"
# EXTRA_OPTS := --sort-keys
TOOLSDIR := _tools

openapi3.yml: ../usegoswagger/swagger.yml
	python ${TOOLSDIR}/2to3.py ${OPTS} ${EXTRA_OPTS} --src $< --savedir .

node_modules:
	npm install
validate:
	./node_modules/.bin/speccy lint openapi3.yml
.PHONY: validate
