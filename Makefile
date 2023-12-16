.PHONY: run
run:
	@swag init --output ./docs/swagger
	@go run main.go serve --verbose


.PHONY: test
test:
	@go test -v ./...

.PHONY: lint
lint: 
	@golangci-lint run