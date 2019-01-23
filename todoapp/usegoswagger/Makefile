validate:
	swagger validate swagger.yml

generate:
	rm -rf gen
	swagger generate server --model-package=gen/models --server-package=gen/restapi

setup:
	go get -v github.com/go-swagger/go-swagger/cmd/swagger

init:
	swagger init spec
