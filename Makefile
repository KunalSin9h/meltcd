run:
	@swag init --output ./docs/swagger
	@go run main.go serve --verbose
