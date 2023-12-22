.PHONY: run
run: frontend
	@swag init --output ./docs/swagger
	@go run main.go serve --verbose

.PHONY: frontend
frontend:
	@pnpm --prefix ./ui run build --emptyOutDir


.PHONY: test
test:
	@go test -v ./...

.PHONY: lint
lint: 
	@golangci-lint run