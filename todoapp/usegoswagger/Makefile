test:
	go test -cover ./...

validate:
	swagger validate swagger.yml

generate:
	swagger generate server --exclude-main --model-package=gen/viewmodel --server-package=gen/restapi

setup:
	go get -v github.com/go-swagger/go-swagger/cmd/swagger

init:
	swagger init spec
