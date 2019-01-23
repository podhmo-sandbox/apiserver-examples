test:
	go test -cover ./...

# database
docker:
	docker-compose up
psql:
	psql -a --dbname todoapp --host localhost --port 5432 -U docker -W
# iniqsql:
# 	psql -a --dbname db --host localhost --port 5432 -U docker -W -f migrations/00init.sql

# swagger
validate:
	swagger validate swagger.yml
generate:
	swagger generate server --exclude-main --model-package=gen/viewmodel --server-package=gen/restapi

# install
setup:
	go get -v github.com/go-swagger/go-swagger/cmd/swagger

# init:
# 	swagger init spec
