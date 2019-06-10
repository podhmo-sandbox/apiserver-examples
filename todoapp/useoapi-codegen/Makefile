OPTS ?= "--log=DEBUG"
# EXTRA_OPTS := --sort-keys
TOOLSDIR := _tools

swagger.yml: ../usegoswagger/swagger.yml
	python ${TOOLSDIR}/2to3.py ${OPTS} ${EXTRA_OPTS} --src $< --savedir .
