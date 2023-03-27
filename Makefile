fiber:
	make swagger
	go run cmd/main/main.go
swagger:
	swag fmt
	swag init -g ./internal/swagger/swag.go -o ./docs/
migrate-up:
	migrate -path ./migrations/ -database 'postgres://postgres:admin@localhost:5432/amiddoc_go?sslmode=disable' up
migrate-down:
	migrate -path ./migrations/ -database 'postgres://postgres:admin@localhost:5432/amiddoc_go?sslmode=disable' down
