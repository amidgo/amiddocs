fiber:
	go run cmd/main/main.go
migrate-up:
	migrate -path ./migrations/ -database 'postgres://postgres:admin@localhost:5432/amiddoc_go?sslmode=disable' up
migrate-down:
	migrate -path ./migrations/ -database 'postgres://postgres:admin@localhost:5432/amiddoc_go?sslmode=disable' down
