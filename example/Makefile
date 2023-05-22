server-start:
	make swagger
	go build -ldflags="-s -w" -o build/amiddocs cmd/main/main.go
	MODE=local build/amiddocs
swagger:
	swag fmt
	swag init -g ./internal/swagger/swag.go -o ./docs/ --parseInternal true
migrate-up:
	migrate -path ./migrations/ -database 'postgres://postgres:admin@localhost:54333/amiddoc_go?sslmode=disable' up
migrate-down:
	migrate -path ./migrations/ -database 'postgres://postgres:admin@localhost:54333/amiddoc_go?sslmode=disable' down
postgres-start:
	docker run -d -e POSTGRES_DB=amiddoc_go \
	 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=admin -e POSTGRES_DB=amiddoc_go \
	 -v ./.tmp/db:/var/lib/postgresql/data  \
	 --name amiddoc-localdb -p 54333:5432 postgres:alpine3.17
postgres-finish:
	docker container stop amiddoc-localdb
	docker container rm amiddoc-localdb
docker-stop:
	docker-compose down
docker-start:
	sudo docker-compose up -d --build
test-data:
	go run cmd/testdata/main.go