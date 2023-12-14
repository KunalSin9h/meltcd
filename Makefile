run:
	@swag init --output ./docs/swagger
	@go run main.go serve --verbose


test:
	@go test -v ./...