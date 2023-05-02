server-start:
	make swagger
	go build -ldflags="-s -w" -o build/fiber cmd/main/main.go
	build/fiber
swagger:
	swag fmt
	swag init -g ./internal/swagger/swag.go -o ./docs/ --parseInternal true
migrate-up:
	migrate -path ./migrations/ -database 'postgres://postgres:admin@localhost:54333/amiddoc_go?sslmode=disable' up
migrate-down:
	migrate -path ./migrations/ -database 'postgres://postgres:admin@localhost:54333/amiddoc_go?sslmode=disable' down
database-start:
	docker run -d -e POSTGRES_DB=amiddoc_go \
	 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=admin -v ./.tmp/db:/var/lib/postgresql/data  \
	 --name amiddoc-localdb -p 54333:5432 postgres:alpine3.17
database-stop:
	docker container stop amiddoc-localdb
	docker container rm amiddoc-localdb
docker-stop:
	docker compose down
	docker image rm amiddocs-server
	cp config/local/.env config/
docker-start:
	cp config/docker/.env config/
	sudo docker compose up -d
docker-rebuild:
	make docker-stop
	make docker-start